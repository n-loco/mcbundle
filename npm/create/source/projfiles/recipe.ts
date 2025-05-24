import { MetaData } from "../utils.js";

export const enum RecipeType {
    BEHAVIOR_PACK = "behavior_pack",
    RESOURCE_PACK = "resource_pack",
    ADDON = "addon",
}

export const enum RecipeModuleType {
    RESOURCES = "resources",
    DATA = "data",
    SERVER = "server",
}

export interface RecipeModule {
    type: RecipeModuleType,
    version: string,
    uuid: string,
    [MetaData]: undefined | "scripting",
}

export interface RecipeConfig {
    type: RecipeType,
    artifact: string,
}

export interface RecipeHeader {
    name: string,
    version: string,
    uuid: string | undefined,
    uuids: [string, string] | undefined,
}

export interface Recipe {
    config: RecipeConfig,
    header: RecipeHeader,
    modules: RecipeModule[],
    [MetaData]: {
        scripting: boolean,
    },
}

export function availableModules(recipeType: RecipeType): RecipeModuleType[] {
    switch (recipeType) {
        case RecipeType.BEHAVIOR_PACK:
            return [
                RecipeModuleType.DATA,
                RecipeModuleType.SERVER,
            ];
        case RecipeType.RESOURCE_PACK:
            return [
                RecipeModuleType.RESOURCES,
            ];
        case RecipeType.ADDON:
            return [
                RecipeModuleType.RESOURCES,
                RecipeModuleType.DATA,
                RecipeModuleType.SERVER,
            ];
    }
}

export function displayStrRecipeType(recipeType: RecipeType): string {
    switch (recipeType) {
        case RecipeType.BEHAVIOR_PACK:
            return "Behavior Pack";
        case RecipeType.RESOURCE_PACK:
            return "Resource Pack";
        case RecipeType.ADDON:
            return "Add-On";
    }
}

export function displayStrRecipeMod(moduleType: RecipeModuleType): string {
    switch (moduleType) {
        case RecipeModuleType.RESOURCES:
            return "Resources";
        case RecipeModuleType.DATA:
            return "Data";
        case RecipeModuleType.SERVER:
            return "Server";
    }
}

export function createRecipeConfig(recipeType: RecipeType): RecipeConfig {
    return {
        type: recipeType,
        artifact: "{{short_name}}_{{version}}",
    }
}

export function createRecipeHeader(config: RecipeConfig, name: string): RecipeHeader {
    const isAddon = config.type === RecipeType.ADDON;
    return {
        name,
        version: "0.1.0",
        uuid: isAddon ? undefined : crypto.randomUUID(),
        uuids: isAddon ? [crypto.randomUUID(), crypto.randomUUID()] : undefined,
    }
}

export function createRecipeModule(modType: RecipeModuleType): RecipeModule {
    return {
        type: modType,
        version: "0.1.0",
        uuid: crypto.randomUUID(),
        [MetaData]: modType === RecipeModuleType.SERVER ? "scripting": undefined,
    }
}

export function createRecipe(config: RecipeConfig, header: RecipeHeader, modules: RecipeModule[]): Recipe {
    const scripting = (() => {
        for (const mod of modules) {
            if (mod[MetaData] === "scripting") return true;
        }
        return false;
    })();

    return {
        config,
        header,
        modules,
        [MetaData]: {
            scripting,
        }
    }
}
