#!/usr/bin/env node

import fs from "node:fs";
import { ProgramVersion } from "../lib/build/index.js";

let packageJSONData = fs.readFileSync("package.json", "utf-8");
let packageJSON = JSON.parse(packageJSONData);

packageJSON.version = ProgramVersion;
packageJSONData = JSON.stringify(packageJSON, null, "  ");

fs.writeFileSync("package.json", new TextEncoder().encode(packageJSONData))
