import path from "node:path";
import fs from "node:fs";

/**
 * @typedef {"release"|"debug"} BuildType
 */

const __buildMode = process.env.BUILD_MODE || "debug";

if (__buildMode != "release" && __buildMode != "debug") {
    throw new Error("invalid build mode");
}

/**
 * This constant is defined in the top-level `Makefile`.
 * But it can be defined manually by setting the `BUILD_MODE` environment variable
 * @type {BuildType}
 */
export const BuildMode = __buildMode;

export const IsDebug = BuildMode == "debug";
export const IsRelease = BuildMode == "release";

/** 
 * Standardized distribution directory name
 * @type {string}
 */
export const DistDirName = IsDebug ? "debug" : "dist";


const __programVersionPath = path.join(import.meta.dirname, "..", "..", "..", "..", "assets", "program_version.txt");

/**
 * This constant is defined in the `assets/program_version.txt` file
 */
export const ProgramVersion = fs.readFileSync(__programVersionPath, "utf-8");

/**
 * This constant is defined in the top-level `platforms.txt` file.
 * But it can be overriden by defining the `SUPPORTED_PLATFORMS` environment variable
 * 
 * All the items in the array matches the pattern `${process.platform}-${process.arch}`
 */
export const SupportedPlatforms = (() => {
    const __supportedPlatforms = process.env.SUPPORTED_PLATFORMS || "";

    if (__supportedPlatforms.length > 0) {
        return __supportedPlatforms.split(" ");
    }

    const rawFilePath = path.join(import.meta.dirname, "..", "..", "..", "..", "platforms.txt");
    const rawData = fs.readFileSync(rawFilePath, "utf-8");

    return rawData.split(/\s+/gm).sort();
})();
