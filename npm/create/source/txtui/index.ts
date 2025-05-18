import { stdout } from "node:process";
import { Arrow, handleArrow, handleEnter, interact } from "./in_ansi.js";
import { bold, colorIndex, ereaseFromCursorToEOE, ereaseFromCursorToEOL, italic, moveCursorH, moveCursorV, reset } from "./out_ansi.js";
import { isEmpty } from "../utils.js";

export function print(str: string, newLine = true) {
    stdout.write(str + (newLine ? "\n" : ""));
}

export namespace Prompt {
    export function stringInput(title: string): string {
        print(bold + title + reset + ": ", false);

        let userStr = "";
        const decoder = new TextDecoder("utf-8");

        const render = (preview = true) => {
            stdout.write("\r" + moveCursorH(title.length + 2));
            stdout.write(ereaseFromCursorToEOL);
            stdout.write(userStr.replaceAll(/§./g, ssCode => ssCodeToAnsi(ssCode, preview)) + reset);
        }

        interact((data) => {
            if (handleArrow(data) != null) {
                return;
            }

            if (data.length === 1 && data[0] == 0x7f) {
                userStr = userStr.slice(0, userStr.length - 1);
            } else {
                userStr += decoder.decode(data);
            }

            render();
        });

        render(false);
        stdout.write("\n");

        return userStr;
    }

    export function selectionMenu(title: string, options: string[]): number {
        print(bold + title + reset);

        let pointer = 0;

        const render = (rerendering = false) => {
            if (rerendering) {
                stdout.write(moveCursorV(-options.length) + "\r");
                stdout.write(ereaseFromCursorToEOE);
            }
            options.forEach((option, i) => {
                const prefix = i === pointer ? "    ▶ " + bold : "      ";
                stdout.write(prefix + option + reset + "\n");
            });
        }

        render();

        interact((data) => {
            const arrow = handleArrow(data);

            if (arrow != null) {
                pointer = (options.length + pointer + Arrow.y(arrow)) % options.length;
            }

            render(true);
        });

        stdout.write(moveCursorV(-options.length - 1) + "\r");
        stdout.write(moveCursorH(title.length) + ": ");
        stdout.write(ereaseFromCursorToEOE);
        stdout.write(colorIndex(0x6bceff) + options[pointer] + reset + "\n");

        return pointer;
    }

    export function checkMenu(title: string, options: string[]): boolean[] {
        print(bold + title + reset);

        let pointer = 0;
        let checkedOpts = new Array<boolean>(options.length).fill(false);

        const render = (rerendering = false) => {
            if (rerendering) {
                stdout.write(moveCursorV(-options.length) + "\r");
                stdout.write(ereaseFromCursorToEOE);
            }
            options.forEach((option, i) => {
                const checkSStr = checkedOpts[i] ? "◼" : "◻"
                const ptr = i === pointer ? "▶" : " ";
                const prefix = `  ${ptr} ${checkSStr} ${i === pointer ? bold : ""}`;
                stdout.write(prefix + option + reset + "\n");
            });
        }

        render();

        interact((data) => {
            const arrow = handleArrow(data);

            if (arrow != null) {
                pointer = (options.length + pointer + Arrow.y(arrow)) % options.length;
            }

            if (data.length === 1 && data[0] === 0x20) {
                checkedOpts[pointer] = !checkedOpts[pointer];
            }

            render(true);
        });

        stdout.write(moveCursorV(-options.length - 1) + "\r");
        stdout.write(moveCursorH(title.length) + ": ");
        stdout.write(ereaseFromCursorToEOE);
        stdout.write(colorIndex(0x6bceff) + options.filter((_, i) => checkedOpts[i]).join(", ") + reset + "\n");

        return checkedOpts;
    }

    export function confirmDialog(title: string): boolean {
        print(bold + title + reset + ": ", false);

        let choice = true;

        const yesDisplay = `${colorIndex(0x38e578)}Yes${reset}`;
        const noDisplay = `${colorIndex(0x535de9)}No ${reset}`;

        const render = (result = false) => {
            stdout.write("\r" + moveCursorH(title.length + 2));

            if (result) {
                stdout.write(ereaseFromCursorToEOL);
                const finalChoiceStr = choice ? "Yes" : "No";
                stdout.write(`${colorIndex(0x6bceff)}${finalChoiceStr}${reset}`)
                return;
            }

            const choiceDisplay = choice ? yesDisplay : noDisplay;
            stdout.write(`${bold}< ${choiceDisplay}${bold} >${reset}`);
        }

        render();

        interact(data => {
            const arrow = handleArrow(data);

            if (arrow != null) {
                choice = Boolean((2 + Number(choice) + Arrow.x(arrow)) % 2);
            }

            render();
        });

        render(true);
        stdout.write("\n");

        return choice;
    }
}

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

function ssCodeToAnsi(ssCode: string, preview: boolean): string {
    switch (ssCode) {
        case "§r":
            return preview ? `${reset}§r` : `${reset}`;
        case "§l":
            return preview ? `${bold}§l` : `${bold}`;
        case "§o":
            return preview ? `${italic}§o` : `${italic}`;
        // case "§k": TODO!
    }

    const colorI = ssColorIndexMap.get(ssCode);
    if (!isEmpty(colorI)) {
        const ansiColor = colorIndex(colorI);
        return preview ? `${ansiColor}${ssCode}` : ansiColor;
    }

    return preview ? ssCode : "";
}
