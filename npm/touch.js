import fs from "node:fs";

const { atime } = fs.statSync("pnpm-lock.yaml");
fs.utimesSync("pnpm-lock.yaml", atime, new Date());
