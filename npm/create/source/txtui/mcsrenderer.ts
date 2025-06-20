import { isEmpty } from "../utils.js";
import { BOLD, DIM, ITALLIC, RESET_BD, RESET, colorSequence } from "./ansi.js";

const ssColorIndexMap = new Map<string, number>([
    ["§0", 0x000000],
    ["§1", 0x0000aa],
    ["§2", 0x00aa00],
    ["§3", 0x00aaaa],
    ["§4", 0xaa0000],
    ["§5", 0xaa00aa],
    ["§6", 0xffaa00],
    ["§7", 0xc5c5c5],
    ["§8", 0x545454],
    ["§9", 0x5454ff],
    ["§a", 0x54ff54],
    ["§b", 0x54ffff],
    ["§c", 0xff5454],
    ["§d", 0xff54ff],
    ["§e", 0xffff54],
    ["§f", 0xffffff],
    ["§g", 0xefce16],
    ["§h", 0xe2d3d1],
    ["§i", 0xcec9c9],
    ["§j", 0x44393a],
    ["§m", 0x961506],
    ["§n", 0xb4684d],
    ["§p", 0xdeb02c],
    ["§q", 0x119f36],
    ["§s", 0x2cb9a8],
    ["§t", 0x20487a],
    ["§u", 0x9a5cc5],
    ["§v", 0xea7113],
]);

function ssCodeToAnsi(renderRecord: RenderRecord, ssCode: string, preview: boolean): string {
    const restoreSuffix = renderRecord.prevColor > -1
        ? colorSequence(renderRecord.prevColor, false)
        : renderRecord.wasBold
            ? `${RESET_BD}${BOLD}`
            : RESET_BD;

    const dimPrefix = renderRecord.prevColor > -1
        ? colorSequence(renderRecord.prevColor, true)
        : DIM;

    const o = (s: string): string => {
        if (renderRecord.obfuscated) {
            return obfuscateChar(s[0], renderRecord.obfuscationNoise) +
                obfuscateChar(s[1], renderRecord.obfuscationNoise + 1);
        } else {
            return s;
        }
    }

    switch (ssCode) {
        case "§r":
            renderRecord.wasBold = false;
            renderRecord.obfuscated = false;
            renderRecord.prevColor = -1;
            return preview ? `${RESET}${DIM}§r${RESET}` : RESET;
        case "§l":
            {
                let ssprev = `${dimPrefix}${BOLD}${o("§l")}`;
                if (renderRecord.prevColor > -1) {
                    ssprev += colorSequence(renderRecord.prevColor);
                } else {
                    ssprev += `${RESET_BD}${BOLD}`;
                }
                renderRecord.wasBold = true;
                return preview
                    ? ssprev
                    : BOLD;
            }
        case "§o":
            return preview ? `${dimPrefix}${ITALLIC}${o("§o")}${restoreSuffix}` : ITALLIC;
        case "§k":
            renderRecord.obfuscated = true;
    }

    const colorI = ssColorIndexMap.get(ssCode);
    if (!isEmpty(colorI)) {
        renderRecord.prevColor = colorI;
        const ansiColor = colorSequence(colorI);
        const dimAnsiColor = colorSequence(colorI, true);
        return preview ? `${dimAnsiColor}${o(ssCode)}${ansiColor}` : ansiColor;
    }

    return preview ? `${dimPrefix}${o(ssCode)}${restoreSuffix}` : "";
}

const obfuscatorSeed = Math.round(Math.random() * 100);

function obfuscateChar(char: string, noise: number): string {
    const particle = ((char.codePointAt(0) || 0) * 2) + noise + obfuscatorSeed;
    const transformed = (particle % (0x02b0 - 0x00a1)) + 0x00a1;
    return String.fromCodePoint(transformed);
}

function obfuscator(str: string): string {
    let obfuscated = "";
    let obfuscate = false;

    for (let i = 0; i < str.length; i++) {
        const char = str[i];
        const char2 = str[i + 1];

        if (char === "§" && char2 != undefined) {
            i++;
            if (char2 === "k") obfuscate = true;
            if (char2 === "r") obfuscate = false;
            obfuscated += char + char2;
        } else {
            obfuscated += obfuscate ? obfuscateChar(char, i) : char;
        }
    }

    return obfuscated;
}

export function renderSSCodes(ssStr: string, preview: boolean): string {
    const record: RenderRecord = {
        wasBold: false,
        obfuscated: false,
        obfuscationNoise: 0,
        prevColor: -1,
    };

    return obfuscator(ssStr).replaceAll(/§./g, (c, i) => {
        record.obfuscationNoise = i;
        return ssCodeToAnsi(record, c, preview);
    });
}

interface RenderRecord {
    wasBold: boolean,
    obfuscated: boolean,
    obfuscationNoise: number,
    prevColor: number
}
