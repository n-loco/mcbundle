import esbuild from "esbuild";
import path from "node:path";
import fs from "node:fs";

const BuildModeValues = new Set(["debug", "release"]);

const BuildMode = process.env.BUILD_MODE;

if (!BuildModeValues.has(BuildMode)) {
    throw new Error(`Invalid build mode: ${BuildMode}`);
}

const IsDebug = BuildMode == "debug";
const IsRelease = BuildMode == "release";

const DbgPath = IsDebug ? "debug/" : "";

const DebugBpBuildPath = path.join(path.resolve(import.meta.dirname), "..", "bpbuild", "debug");
const BpBuildVersion = JSON.parse(fs.readFileSync("package.json", "utf-8")).version;

const BpBuildSpecifier = BuildMode == "debug"
    ? `file:${DebugBpBuildPath.replaceAll("\\", "\\\\")}`
    : `^${BpBuildVersion}`;

esbuild.buildSync({
    bundle: true,
    packages: "external",
    platform: "node",
    format: "esm",
    entryPoints: ["./source/main.ts"],
    outfile: `./${DbgPath}dist/create.mjs`,
    target: "node22",
    minify: IsRelease,
    sourcemap: IsDebug && "inline",
    sourcesContent: false,
    define: { // ./defines.d.ts
        "BPBuildSpecifier": `"${BpBuildSpecifier}"`,
    },
});

if (IsDebug) {
    fs.copyFileSync("package.json", "./debug/package.json");
}

