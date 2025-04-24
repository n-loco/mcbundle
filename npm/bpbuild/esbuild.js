import esbuild from "esbuild";

esbuild.buildSync({
    entryPoints: ["./source/main.ts"],
    outfile: "./dist/bpbuild.mjs",
    minifyWhitespace: true,
    target: "node22"
});
