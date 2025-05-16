#!/usr/bin/env node

import { exit } from "node:process";
import { getPackageManager, PackageManagerName } from "./package_manager.js";
import { print, Prompt } from "./txtui/index.js";

const projTypes = ["behavior_pack", "resource_pack", "addon"];

const projTypeIndex = Prompt.selectionMenu("● Project type", [
    "Behavior Pack",
    "Resource Pack",
    "Add-On",
]);

const projName = Prompt.stringInput(`● ${projTypes[projTypeIndex] === "addon" ? "Add-On" : "Pack"} name`);

print("\nrecipe.json:\n" + JSON.stringify({
    config: {
        type: projTypes[projTypeIndex],
    },
    header: {
        name: projName,
        uuid: projTypes[projTypeIndex] !== "addon" ? crypto.randomUUID() : undefined,
        uuids: projTypes[projTypeIndex] === "addon" ? [crypto.randomUUID(), crypto.randomUUID()] : undefined,
        version: "0.1.0",
    }
}, null, "  "));

print("\npackage.json:\n" + JSON.stringify({
    type: "module",
    private: true,
    packageManager: (() => {
        const pkgManager = getPackageManager();
        if (pkgManager.name !== PackageManagerName.Unknown && pkgManager.name !== PackageManagerName.Bun) {
            return `${pkgManager.name}@${pkgManager.version}`;
        }
        return undefined;
    })(),
    devDependencies: {
        "bpbuild": BPBuildSpecifier,
    },
}, null, "  "));
