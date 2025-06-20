#!/usr/bin/env node

import { cwd } from "node:process";
import childProcess from "node:child_process";
import { availableModules, createPackageJSON, createRecipe, createRecipeConfig, createRecipeHeader, createRecipeModule, createTSConfig, displayStrRecipeMod, displayStrRecipeType, gitattributes, gitignoreLines, Language, PackageManager, RecipeModule, RecipeModuleType, RecipeType, serverMainScript } from "./projfiles/index.js";
import { boolDialog, BoolDialogOptions, checkMenu, CheckMenuOptions, Colors, Selectable, selectMenu, SelectMenuOptions, stringInput, StringInputOptions, TextNode } from "./txtui/index.js";
import path from "node:path";
import { MetaData } from "./utils.js";
import { createDirNode, createFileNode, DirNode, writeFS } from "./mktree.js";

const recipeType = await (async () => {
    const options = [RecipeType.RESOURCE_PACK, RecipeType.BEHAVIOR_PACK, RecipeType.ADD_ON]
    const displayOptions = options.map<Selectable>(o => {
        const displayStr = displayStrRecipeType(o);
        return {
            normal: displayStr,
            highlighted: [{ content: displayStr, bold: true }],
            choosen: [{ content: displayStr, color: Colors.TYPE_VALUE }]
        };
    });
    const menuOptions: SelectMenuOptions = {
        label: [{ content: "Project Type", bold: true }],
        options: displayOptions,
    };

    const choice = await selectMenu(menuOptions);

    return options[choice];
})();

const projName = await (async () => {
    const defaultName = path.basename(cwd()).split(/(?:[ _-]|(?<=[a-z])(?=[A-Z]))+/g)
        .map(s => s.slice(0, 1).toUpperCase() + s.slice(1).toLowerCase())
        .join(" ");
    
    const inputOptions: StringInputOptions = {
        label: [{ content: "Project Name", bold: true }],
        defaultValue: defaultName
    };

    return await stringInput(inputOptions);
})()

const initialMods = await (async () => {
    const options = availableModules(recipeType);
    const displayOptions = options.map<Selectable>(m => {
        const displayStr = displayStrRecipeMod(m);
        const hNode: TextNode = { content: displayStr, bold: true };
        return {
            normal: displayStr,
            highlighted: m === RecipeModuleType.SERVER
                ? [hNode, { content: " - the \"script\" module in manifest.json", dim: true }]
                : [hNode],
            choosen: [{ content: displayStr, color: Colors.TYPE_VALUE }],
        };
    });

    const menuOptions: CheckMenuOptions = {
        label: [{ content: "Initial Modules", bold: true }],
        options: displayOptions,
        defaultValues: (() => {
            const arr: boolean[] = [];
            arr.length = options.length;
            arr.fill(true);
            return arr;
        })(),
    };

    const choices = await checkMenu(menuOptions);

    return options.filter((_, i) => choices[i]);
})()

const recipeConfig = createRecipeConfig(recipeType);
const recipeHeader = createRecipeHeader(recipeConfig, projName);
const recipeModules: RecipeModule[] = [];
for (const rt of initialMods) {
    recipeModules.push(createRecipeModule(rt));
}

const recipe = createRecipe(recipeConfig, recipeHeader, recipeModules);
const hasScripting = recipe[MetaData].scripting;

const language = await (async () => {
    if (!hasScripting) return null

    const options = [Language.TYPESCRIPT, Language.JAVASCRIPT];
    const displayOptions: Selectable[] = [
        {
            normal: "TypeScript",
            highlighted: [{ content: "TypeScript", bold: true }],
            choosen: [{ content: "TypeScript", color: Colors.LANGUAGE }],
        },
        {
            normal: "JavaScript",
            highlighted: [{ content: "JavaScript", bold: true }],
            choosen: [{ content: "JavaScript", color: Colors.LANGUAGE }],
        }
    ];

    const menuOptions: SelectMenuOptions = {
        label: [{ content: "Language", bold: true }],
        options: displayOptions,
    };

    const choice = await selectMenu(menuOptions);

    return options[choice];
})();

const packageJSON = createPackageJSON(recipe);

const pkgMngr = packageJSON[MetaData].packageManaer;

const projTree: DirNode = createDirNode(".", []);

projTree.nodes.push(
    createFileNode("recipe.json", JSON.stringify(recipe, null, "  ")),
    createFileNode("package.json", JSON.stringify(packageJSON, null, "  ")),
    createFileNode(".gitattributes", gitattributes),
    createFileNode(".gitignore", gitignoreLines(pkgMngr).join("\n")),
);

const sourceDir = createDirNode("source", []);

for (const mod of initialMods) {
    sourceDir.nodes.push(
        createDirNode(mod, mod === RecipeModuleType.SERVER
            ? [createFileNode(language === Language.TYPESCRIPT ? "main.ts" : "main.js", serverMainScript)]
            : []
        ),
    );
}

projTree.nodes.push(sourceDir);

if (hasScripting) {
    projTree.nodes.push(
        createFileNode("tsconfig.json", JSON.stringify(createTSConfig(), null, "  ")),
    );
}

writeFS(projTree);

if (pkgMngr !== PackageManager.UNKNOWN) {
    const installPkgs = await (async () => {
        const dialogOptions: BoolDialogOptions = {
            label: [{ content: "Install Packages", bold: true }],
            truthy: {
                normal: ["Yes ", { content: `(${pkgMngr})`, dim: true }],
                highlighted: [
                    { content: "Yes ", bold: true },
                    { content: `(${pkgMngr})`, dim: true },
                ],
                choosen: [{ content: "Yes", color: Colors.BOOL_VALUE }],
            },
            defaultValue: true,
        };

        return await boolDialog(dialogOptions);
    })();

    if (installPkgs) {
        childProcess.spawnSync(pkgMngr, ["install"], { stdio: "inherit" });
    }
}
