#!/usr/bin/env node

import fs from "node:fs";

const { atime, mtime } = fs.statSync(".og.package.json");
const ogPackageJSON = fs.readFileSync(".og.package.json", "utf-8");

fs.rmSync(".og.package.json");

fs.writeFileSync("package.json", new TextEncoder().encode(ogPackageJSON));
fs.utimesSync("package.json", atime, mtime);
