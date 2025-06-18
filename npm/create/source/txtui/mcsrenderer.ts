import { isEmpty } from "../utils.js";
import { BOLD, colorIndex, DIM, ITALLIC, RESET_BD, RESET } from "./ansi.js";

const ssCodeRegExp = /§./g;

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

function ssCodeToAnsi(wasBold: VRef<boolean>, ssCode: string, preview: boolean): string {
    const boldSuffix = wasBold.value ? BOLD : "";

    switch (ssCode) {
        case "§r":
            wasBold.value = false;
            return preview ? `${RESET}${DIM}§r${RESET}` : `${RESET}`;
        case "§l":
            wasBold.value = true;
            return preview ? `${DIM}${BOLD}§l${RESET_BD}${BOLD}` : `${BOLD}`;
        case "§o":
            return preview ? `${DIM}${ITALLIC}§o${RESET_BD}${boldSuffix}` : `${ITALLIC}`;
        // case "§k": TODO!
    }

    const colorI = ssColorIndexMap.get(ssCode);
    if (!isEmpty(colorI)) {
        const ansiColor = colorIndex(colorI);
        return preview ? `${DIM}${ansiColor}${ssCode}${RESET_BD}${boldSuffix}` : ansiColor;
    }

    return preview ? `${DIM}${ssCode}${RESET_BD}${boldSuffix}` : "";
}

export function renderSSCodes(ssStr: string, preview: boolean): string {
    const wasBold: VRef<boolean> = {
        value: false,
    }

    return ssStr.replaceAll(ssCodeRegExp, c => {
        return ssCodeToAnsi(wasBold, c, preview);
    });
}

interface VRef<T> {
    value: T;
}
