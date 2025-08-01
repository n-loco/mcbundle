import { IsDebug, IsRelease, ProgramVersion, DistDirName } from "root/build";
import esbuild from "esbuild";
import path from "node:path";

const McbundlePath = path.join(path.resolve(import.meta.dirname), "..", "mcbundle");

const mcbundleSpecVal = IsDebug
    ? `file:${McbundlePath.replaceAll("\\", "\\\\")}`
    : `^${ProgramVersion}`;

esbuild.buildSync({
    bundle: true,
    packages: "external",
    platform: "node",
    format: "esm",
    entryPoints: ["./source/main.ts"],
    outfile: `./${DistDirName}/create.mjs`,
    target: "node22",
    minify: IsRelease,
    sourcemap: IsDebug && "inline",
    sourcesContent: false,
    sourceRoot: path.join(import.meta.dirname, "debug"),
    define: { // ./defines.d.ts
        "mcbundleSpecifier": `"${mcbundleSpecVal}"`,
    },
});
