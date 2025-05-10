import esbuild from "esbuild";
import path from "node:path";
import fs from "node:fs";
import url from "node:url";

const BuildModeValues = new Set(["debug", "release"]);

const BuildMode = process.env.BUILD_MODE;

if (!BuildModeValues.has(BuildMode)) {
    throw new Error(`Invalid build mode: ${BuildMode}`);
}

const IsDebug = BuildMode == "debug";
const IsRelease = BuildMode == "release";

const DbgPath = IsDebug ? "debug/" : "";

esbuild.buildSync({
    bundle: true,
    packages: "external",
    platform: "node",
    format: "esm",
    entryPoints: ["./source/main.ts"],
    outfile: `./${DbgPath}dist/bpbuild.mjs`,
    target: "node22",
    minifyWhitespace: IsRelease,
    sourcemap: IsDebug && "linked",
    sourcesContent: false,
    sourceRoot: import.meta.dirname,
    define: { // defines.d.ts
        "Debug": `${IsDebug}`,
    },
});

if (IsDebug) {
    const packageJsonContent = fs.readFileSync("package.json", "utf-8");
    const packageJsonObj = JSON.parse(packageJsonContent);

    for (const dep of Object.getOwnPropertyNames(packageJsonObj.optionalDependencies)) {
        const depPath = path.dirname((url.fileURLToPath(import.meta.resolve(`${dep}/package.json`))));
        packageJsonObj.optionalDependencies[dep] = `file:${depPath.replaceAll("\\", "\\\\")}`;
    }

    fs.writeFileSync("./debug/package.json", JSON.stringify(packageJsonObj, null, "  "), "utf-8");
}
