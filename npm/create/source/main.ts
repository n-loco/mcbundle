#!/usr/bin/env node

import { bell, boolDialog, checkMenu, println, selectMenu, stringInput } from "./txtui/index.js";

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

console.debug(await selectMenu({
    label: [{ content: "Test9", bold: true }],
    options: [
        {
            normal: "Op 1",
            highlighted: [
                { content: "Op 1", bold: true }, " ",
                { content: "useful hint 1", dim: true },
            ],
            choosen: [{ content: "Op 1", color: 0x2bef76 }]
        },
        {
            normal: "Op 2",
            highlighted: [
                { content: "Op 2", bold: true }, " ",
                { content: "useful hint 2", dim: true },
            ],
            choosen: [{ content: "Op 2", color: 0x2bef76 }]
        },
        {
            normal: "Op 3",
            highlighted: [
                { content: "Op 3", bold: true }, " ",
                { content: "useful hint 3", dim: true },
            ],
            choosen: [{ content: "Op 3", color: 0x2bef76 }]
        }
    ],
}));

console.debug(await checkMenu({
    label: [{ content: "Test10", bold: true }],
    options: [
        {
            normal: "Data",
            highlighted: [{ content: "Data", bold: true }],
            choosen: [{ content: "Data", color: 0x2bef76 }]
        },
        {
            normal: "Server",
            highlighted: [{ content: "Server", bold: true }],
            choosen: [{ content: "Server", color: 0x2bef76 }]
        }
    ]
}));

console.debug(await boolDialog({
    label: [{ content: "Test11", bold: true }]
}));

console.debug(await stringInput({
    label: [{ content: "Test12", bold: true }],
    defaultValue: "Minceraft 2",
}));
