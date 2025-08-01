import { SupportedPlatforms } from "root/build";
import fs from "node:fs";

const nativePackages = SupportedPlatforms.map(t => `@mcbundle/${t}`);
const nativePackageSpecifiers = nativePackages.map(p => `workspace:../${p}`);

/** @type {any} */
const optionalDependencies = {};

for (let i = 0; i < SupportedPlatforms.length; i++) {
    optionalDependencies[nativePackages[i]] = nativePackageSpecifiers[i];
}

let packageJSONData = fs.readFileSync("package.json", "utf-8");
let packageJSON = JSON.parse(packageJSONData);

packageJSON.optionalDependencies = optionalDependencies;

packageJSONData = JSON.stringify(packageJSON, null, "  ");

fs.writeFileSync("package.json", new TextEncoder().encode(packageJSONData));
