import { env } from "node:process";
import { isEmpty, MetaData } from "../utils.js";
import { Recipe } from "./recipe.js";

export interface PackageJSON {
    type: "module" | "commonjs",
    private: boolean,
    packageManager: string | undefined,
    devDependencies: { [key: string]: string },
    dependencies: { [key: string]: string } | undefined,
    [MetaData]: {
        packageManaer: PackageManager,
    },
}

export const enum PackageManager {
    NPM = "npm",
    YARN = "yarn",
    PNPM = "pnpm",
    BUN = "bun",

    UNKNOWN = "unknown",
}

export const enum Language {
    JAVASCRIPT = "JavaScript",
    TYPESCRIPT = "TypeScript",
}

export function createPackageJSON(recipe: Recipe): PackageJSON {
    const [pmName, pmVersion] = getPackageManager();

    const isNodePM = pmName !== PackageManager.UNKNOWN && pmVersion !== PackageManager.BUN;

    return {
        type: "module",
        private: true,
        packageManager: isNodePM ? `${pmName}@${pmVersion}` : undefined,
        dependencies: recipe[MetaData].scripting ? { 
            "@minecraft/server": "^1.9.0",
            "@minecraft/vanilla-data": "^1.21.0",
         } : undefined,
        devDependencies: {
            "bpbuild": BPBuildSpecifier,
        },
        [MetaData]: {
            packageManaer: pmName,
        },
    }
}

function isPMKnown(name: string): name is PackageManager {
    switch (name) {
        case PackageManager.NPM:
        case PackageManager.YARN:
        case PackageManager.PNPM:
        case PackageManager.BUN:
            return true;
    }
    return false;
}

const unknownPackageManager: [name: PackageManager, version: string] = [PackageManager.UNKNOWN, "?"]

function getPackageManager(): [name: PackageManager, version: string] {
    const npmConfigUserAgent = env.npm_config_user_agent;

    if (isEmpty(npmConfigUserAgent)) {
        return unknownPackageManager;
    }

    const firstSlashIndex = npmConfigUserAgent.indexOf("/");

    const packageManagerName = npmConfigUserAgent.slice(0, firstSlashIndex);

    if (!isPMKnown(packageManagerName)) {
        return unknownPackageManager;
    }

    const packageManagerVersion = npmConfigUserAgent.slice(firstSlashIndex + 1, npmConfigUserAgent.indexOf(" "));

    return [packageManagerName, packageManagerVersion];
}