import esbuild from "esbuild";

esbuild.buildSync({
    bundle: true,
    packages: "external",
    platform: "node",
    format: "esm",
    entryPoints: ["./source/main.ts"],
    outfile: "./dist/bpbuild.mjs",
    target: "node22",
    minifyWhitespace: true,
});
