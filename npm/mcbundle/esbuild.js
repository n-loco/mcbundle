import { IsDebug, IsRelease, DistDirName } from "root/build";
import esbuild from "esbuild";
import path from "node:path";

esbuild.buildSync({
    bundle: true,
    packages: "external",
    platform: "node",
    format: "esm",
    entryPoints: ["./source/main.ts"],
    outfile: `./${DistDirName}/mcbundle.mjs`,
    target: "node22",
    minifyWhitespace: IsRelease,
    sourcemap: IsDebug && "inline",
    sourceRoot: path.join(import.meta.dirname, "debug"),
    sourcesContent: false,
    define: { // defines.d.ts
        "DistDir": `"${DistDirName}"`,
    },
});
