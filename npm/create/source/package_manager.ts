import { env } from "node:process";
import { isEmpty } from "./utils.js";

export const enum PackageManagerName {
    Npm = "npm",
    Yarn = "yarn",
    Pnpm = "pnpm",
    Bun = "bun",

    Unknown = "unknown"
}

function isNameKnown(name: string): name is PackageManagerName {
    switch (name) {
        case PackageManagerName.Npm:
        case PackageManagerName.Yarn:
        case PackageManagerName.Pnpm:
        case PackageManagerName.Bun:
            return true;
    }
    return false;
}

export interface PackageManager {
    readonly name: PackageManagerName,
    readonly version: string,
}

const unknownPackageManager: PackageManager = {
    name: PackageManagerName.Unknown,
    version: "",
}

export function getPackageManager(): PackageManager {
    const npmConfigUserAgent = env.npm_config_user_agent;

    if (isEmpty(npmConfigUserAgent)) {
        return unknownPackageManager;
    }

    const firstSlashIndex = npmConfigUserAgent.indexOf("/");

    const packageManagerName = npmConfigUserAgent.slice(0, firstSlashIndex);

    if (!isNameKnown(packageManagerName)) {
        return unknownPackageManager;
    }

    const packageManagerVersion = npmConfigUserAgent.slice(firstSlashIndex + 1, npmConfigUserAgent.indexOf(" "));

    return {
        name: packageManagerName,
        version: packageManagerVersion,
    };
}
