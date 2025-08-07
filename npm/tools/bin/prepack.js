#!/usr/bin/env node

import fs from "node:fs";
import path from "node:path";

const licenseData = fs.readFileSync(path.join(import.meta.dirname, "..", "..", "..", "LICENSE"));

const { atime, mtime } = fs.statSync("package.json");
let packageJSONData = fs.readFileSync("package.json", "utf-8");
let packagePublishJSONData = fs.readFileSync("package.publish.json", "utf-8");

fs.writeFileSync(".og.package.json", new TextEncoder().encode(packageJSONData));
fs.utimesSync(".og.package.json", atime, mtime);

let packageJSON = JSON.parse(packageJSONData);
let packagePublishJSON = JSON.parse(packagePublishJSONData)

Object.assign(packageJSON, packagePublishJSON);

packageJSONData = JSON.stringify(packageJSON, null, "  ");

fs.writeFileSync("package.json", new TextEncoder().encode(packageJSONData))

fs.writeFileSync("LICENSE", licenseData);
