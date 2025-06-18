#!/usr/bin/env node

import { bell, println } from "./txtui/index.js";

println("Test1");

println([
    {
        content: "Test2",
        color: 0xbbddff,
        itallic: true,
        dim: true,
    },
    " with some nice plain text",
]);
println([
    {
        content: "Test3",
        backgroundColor: 0x00cc00,
        bold: true,
    }
]);

bell(); // "Test4"

println([
    {
        content: "§b§oTest5§r",
        useMCSSCodes: true,
    }
]);

println([
    {
        content: "§b§oTest6§r",
        useMCSSCodes: true,
        ssCodesPreview: true,
    }
]);

println([
    {
        content: "§l§e§oTest7§r pickaxe",
        useMCSSCodes: true,
        ssCodesPreview: true,
    }
]);

println([
    {
        content: "§l§c§oTest8§r §kcreeper§r",
        useMCSSCodes: true,
        ssCodesPreview: true,
    }
]);
