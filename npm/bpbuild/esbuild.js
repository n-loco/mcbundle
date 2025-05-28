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
    sourcemap: IsDebug && "inline",
    sourcesContent: false,
});

if (IsDebug) {
    const packageJsonContent = fs.readFileSync("package.json", "utf-8");
    const packageJsonObj = JSON.parse(packageJsonContent);

    const bpBuildNodeModules = path.join("debug", "node_modules", "@bpbuild");

    fs.mkdirSync(bpBuildNodeModules, { recursive: true });

    for (const dep of Object.getOwnPropertyNames(packageJsonObj.optionalDependencies)) {
        const taget_double = dep.split("/")[1];
        const depPath = path.join(url.fileURLToPath(import.meta.resolve(`${dep}/package.json`)), "..", "debug");
        const nodeModPath = path.join(bpBuildNodeModules, taget_double);

        packageJsonObj.optionalDependencies[dep] = `file:${depPath.replaceAll("\\", "\\\\")}`;
        fs.rmSync(nodeModPath, { force: true });
        fs.symlinkSync(depPath, nodeModPath);
    }

    fs.writeFileSync(path.join("debug", "package.json"), JSON.stringify(packageJsonObj, null, "  "), "utf-8");
}
