import esbuild from "esbuild";
import path from "node:path";

const bpbuildPath = path.join(path.resolve(import.meta.dirname), "..", "bpbuild").replaceAll("\\", "\\\\");
const bpbuildSpecifier = `file:${bpbuildPath}`;

esbuild.buildSync({
    bundle: true,
    packages: "external",
    platform: "node",
    format: "esm",
    entryPoints: ["./source/main.ts"],
    outfile: "./dist/create.mjs",
    target: "node22",
    minify: true,
    define: { // ./defines.d.ts
        "BPBuildSpecifier": `"${bpbuildSpecifier}"`,
    },
});
