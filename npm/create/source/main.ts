#!/usr/bin/env node

import { createDirNode, createFileNode, FSNode, writeFS } from "./mktree.js";
import {
    createPackageJSON,
    Language,
    PackageJSON,
    PackageManager,
    serverMainScript,
    availableModules,
    createRecipe,
    createRecipeConfig,
    createRecipeHeader,
    createRecipeModule,
    displayStrRecipeMod,
    displayStrRecipeType,
    Recipe,
    RecipeConfig,
    RecipeModuleType,
    RecipeType,
    gitattributes,
    gitignoreLines,
    createTSConfig,
} from "./projfiles/index.js";
import { Prompt } from "./txtui/index.js";
import { MetaData } from "./utils.js";
import childProcess from "node:child_process";

function askProjectType(): RecipeType {
    const options = [
        RecipeType.BEHAVIOR_PACK,
        RecipeType.RESOURCE_PACK,
        RecipeType.ADDON,
    ];

    const menuDisplayOptions = options.map(displayStrRecipeType);
    const choice = Prompt.selectionMenu("● Project type", menuDisplayOptions);

    return options[choice];
}

function askProjectName(): string {
    const recipeTypeStr = displayStrRecipeType(config.type);
    const projName = Prompt.stringInput(`● ${recipeTypeStr} name`);
    return projName;
}

function askInitialModules(config: RecipeConfig): RecipeModuleType[] {
    const options = availableModules(config.type);

    if (options.length === 1) {
        return [options[0]];
    }

    const menuDisplayOptions = options.map(displayStrRecipeMod);
    const choices = Prompt.checkMenu("● Initial modules", menuDisplayOptions);

    return options.filter((_, i) => choices[i]);
}

function askLanguage(recipe: Recipe): Language | null {
    if (!recipe[MetaData].scripting) return null

    const options = [Language.JAVASCRIPT, Language.TYPESCRIPT];
    const choice = Prompt.selectionMenu("● Language", options);

    return options[choice];
}

function askInsallPkgs(packageJSON: PackageJSON): boolean {
    const pmName = packageJSON[MetaData].packageManaer;

    if (pmName === PackageManager.UNKNOWN) {
        return false;
    }

    return Prompt.confirmDialog(`● Install packages (${pmName})`);
}

const config = createRecipeConfig(askProjectType());
const header = createRecipeHeader(config, askProjectName());
const modules = askInitialModules(config).map(createRecipeModule);

const recipe = createRecipe(config, header, modules);
const packageJSON = createPackageJSON(recipe);

const language = askLanguage(recipe);
const installPkgs = askInsallPkgs(packageJSON);

const sourceDirNodes: FSNode[] = [];


for (const mod of modules) {
    if (mod[MetaData] === "scripting") {
        const ext = (() => {
            switch (language) {
                case Language.JAVASCRIPT:
                    return ".js";
                case Language.TYPESCRIPT:
                    return ".ts";
            }
        })();

        sourceDirNodes.push(createDirNode(mod.type, [
            createFileNode(`main${ext}`, serverMainScript),
        ]));
    } else {
        sourceDirNodes.push(createDirNode(mod.type, []));
    }
}

const projFS = createDirNode(".", [
    createDirNode("source", sourceDirNodes),
    createFileNode("package.json", JSON.stringify(packageJSON, null, "  ")),
    createFileNode("recipe.json", JSON.stringify(recipe, null, "  ")),
    createFileNode(".gitattributes", gitattributes),
    createFileNode(".gitignore", gitignoreLines(packageJSON[MetaData].packageManaer).join("\n")),
]);

if (language === Language.TYPESCRIPT) {
    projFS.nodes.push(createFileNode("tsconfig.json", JSON.stringify(createTSConfig(), null, "  ")));
}

writeFS(projFS);

if (installPkgs) {
    childProcess.spawnSync(packageJSON[MetaData].packageManaer, ["install"], { stdio: "inherit" });
}
