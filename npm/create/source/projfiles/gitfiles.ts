import { PackageManager } from "./package.js";

export const gitattributes = `* text=auto eol=crlf`;

export function gitignoreLines(pm: PackageManager): string[] {
    const output = [
        "# Output",
        "dist/",
    ];

    const depDirs = pm === PackageManager.YARN
        ? [
            "# Dependency directories",
            ".yarn/*",
            "!.yarn/patches",
            "!.yarn/plugins",
            "!.yarn/releases",
            "!.yarn/sdks",
            "!.yarn/versions",
            ".pnp.*",
        ]
        : [
            "# Dependency directories",
            "node_modules/",
        ];

    const localRepo = [
        "# Local repo",
        "/*[Uu][Ss][Ee][Rr]*",
    ];

    const result = new Array<string>();

    result.push(...output, "");
    result.push(...depDirs, "");
    result.push(...localRepo);

    return result;
}
