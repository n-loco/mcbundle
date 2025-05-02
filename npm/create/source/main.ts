#!/usr/bin/env node

import { stdout } from "node:process";
import { getPackageManager, PackageManagerName } from "./package_manager.js";

const packageManager = getPackageManager();

if (packageManager.name == PackageManagerName.Unknown) {
    stdout.write(`No package manager was detected\n`);
} else {
    stdout.write(`Using ${packageManager.name} - v${packageManager.version}\n`);
}

const packageJson = {
    packageManager: (() => {
        if (packageManager.name === PackageManagerName.Unknown
            || packageManager.name === PackageManagerName.Bun) {
            return undefined;
        }
        return `${packageManager.name}@${packageManager.version}`;
    })(),
    devDependencies: {
        "bpbuild": BPBuildSpecifier,
    }
};

stdout.write(`${JSON.stringify(packageJson, null, "  ")}\n`);
