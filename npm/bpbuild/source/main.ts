#!/usr/bin/env node

import process, { argv, platform as os, arch as cpu } from "node:process";
import childProcess from "node:child_process";
import url from "node:url";

function findExecutable(): string {
    let suffix = os == "win32" ? ".exe" : "";
    let exeName = `bpbuild${suffix}`;
    let packageName = `@bpbuild/${os}-${cpu}`;
    let fileURL = import.meta.resolve(`${packageName}/${exeName}`);
    return url.fileURLToPath(fileURL);
}

let exePath = findExecutable();

let bpbuildProcess = childProcess.spawn(exePath, argv.slice(2), { stdio: "inherit", detached: true });

bpbuildProcess.addListener("exit", exitCode => {
    process.exitCode = exitCode || 0;
});
