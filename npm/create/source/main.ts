#!/usr/bin/env node

import { createPackageJSON, Language, PackageJSON, PackageManager } from "./projfiles/package.js";
import {
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
} from "./projfiles/recipe.js";
import { print, Prompt } from "./txtui/index.js";
import { MetaData } from "./utils.js";

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

askLanguage(recipe);
askInsallPkgs(packageJSON);

print("\n" + JSON.stringify(recipe, null, "  "));
print("\n" + JSON.stringify(packageJSON, null, "  "));
