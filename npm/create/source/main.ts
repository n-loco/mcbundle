#!/usr/bin/env node

import { getPackageManager, PackageManagerName } from "./package_manager.js";
import { print, Prompt } from "./txtui/index.js";

const pkgManager = getPackageManager();
const isPkgManager = pkgManager.name !== PackageManagerName.Unknown && pkgManager.name !== PackageManagerName.Bun;

const projTypes = ["behavior_pack", "resource_pack", "addon"];

const projTypeIndex = Prompt.selectionMenu("● Project type", [
    "Behavior Pack",
    "Resource Pack",
    "Add-On",
]);

const sProjT = projTypes[projTypeIndex];
const isAddon = sProjT === "addon";

const projName = Prompt.stringInput(`● ${isAddon ? "Add-On" : "Pack"} name`);

const modules = (() => {
    if (isAddon) {
        const addonModules = [
            {
                type: "resources",
                version: "0.1.0",
                uuid: crypto.randomUUID(),
            },
            {
                type: "data",
                version: "0.1.0",
                uuid: crypto.randomUUID(),
            },
            {
                type: "server",
                version: "0.1.0",
                uuid: crypto.randomUUID(),
            },
        ];
        const opts = Prompt.checkMenu("● Start with modules", [
            "Resources",
            "Data",
            "Server",
        ]);

        return addonModules.filter((_, i) => opts[i]);
    }

    if (sProjT === "behavior_pack") {
        const bpModules = [
            {
                type: "data",
                version: "0.1.0",
                uuid: crypto.randomUUID(),
            },
            {
                type: "server",
                version: "0.1.0",
                uuid: crypto.randomUUID(),
            },
        ];

        const opts = Prompt.checkMenu("● Start with modules", [
            "Data",
            "Server",
        ]);

        return bpModules.filter((_, i) => opts[i]);
    }

    if (sProjT === "resource_pack") {
        return [{
            type: "resources",
            version: "0.1.0",
            uuid: crypto.randomUUID(),
        }];
    }
})();

const installPkgs = (() => {
    if (pkgManager.name === PackageManagerName.Unknown) return false;
    return Prompt.confirmDialog(`● Install packages (${pkgManager.name})`);
})();

if (installPkgs) {
    print(`\ninstalling with ${pkgManager.name}...\n`);
}

print("\nrecipe.json:\n" + JSON.stringify({
    config: {
        type: projTypes[projTypeIndex],
    },
    header: {
        name: projName,
        uuid: !isAddon ? crypto.randomUUID() : undefined,
        uuids: isAddon ? [crypto.randomUUID(), crypto.randomUUID()] : undefined,
        version: "0.1.0",
    },
    modules,
}, null, "  "));

print("\npackage.json:\n" + JSON.stringify({
    type: "module",
    private: true,
    packageManager: (() => {
        if (isPkgManager) {
            return `${pkgManager.name}@${pkgManager.version}`;
        }
        return undefined;
    })(),
    devDependencies: {
        "bpbuild": BPBuildSpecifier,
    },
}, null, "  "));
