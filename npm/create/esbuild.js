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

const DebugMcbundlePath = path.join(path.resolve(import.meta.dirname), "..", "mcbundle", "debug");
const mcbundleVersion = JSON.parse(fs.readFileSync("package.json", "utf-8")).version;

const mcbundleSpecVal = BuildMode == "debug"
    ? `file:${DebugMcbundlePath.replaceAll("\\", "\\\\")}`
    : `^${mcbundleVersion}`;

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
    sourceRoot: path.join(import.meta.dirname, "debug", "dist"),
    define: { // ./defines.d.ts
        "mcbundleSpecifier": `"${mcbundleSpecVal}"`,
    },
});

if (IsDebug) {
    fs.copyFileSync("package.json", "./debug/package.json");
}

