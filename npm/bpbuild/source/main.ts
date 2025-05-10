#!/usr/bin/env node

import url from "node:url";
import childProcess from "node:child_process";

function findExecutable(): string {
    let suffix = process.platform == "win32" ? ".exe" : "";
    let exeName = (Debug && `debug/bpbuild${suffix}`) || `bpbuild${suffix}`;
    let packageName = `@bpbuild/${process.platform}-${process.arch}`;
    let fileURL = import.meta.resolve(`${packageName}/${exeName}`);
    return url.fileURLToPath(fileURL);
}

let exePath = findExecutable();

let bpbuildProcess = childProcess.spawn(exePath, process.argv.slice(2), { stdio: "inherit", detached: true });

bpbuildProcess.addListener("exit", exitCode => {
    process.exitCode = exitCode || 0;
});
