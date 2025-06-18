import { kill, pid, stdin, stdout } from "node:process";
import readline, { emitKeypressEvents, Key } from "node:readline";
import { TextNode, TextSpan } from "./text.js";
import { BEL, BOLD, colorIndex, DIM, EOT, ITALLIC, RESET } from "./ansi.js";
import { isEmpty } from "../utils.js";
import { renderSSCodes } from "./mcsrenderer.js";

namespace StdIO {
    export function bell() {
        stdout.write(BEL);
    }

    export function write(s: string | TextSpan) {
        if (typeof s === "string") {
            stdout.write(s);
        } else {
            let wstr = "";
            for (const node of s) {
                if (typeof node === "string") {
                    wstr += node;
                } else {
                    wstr += renderTextNode(node);
                }
            }
            stdout.write(wstr);
        }
    }

    export async function keypress(): Promise<string> {
        emitKeypressEvents(stdin);
        stdin.setRawMode(true);
        stdin.resume();

        const keyPromise = new Promise<string>((resolve) => {
            stdin.once("keypress", (_: string, key: Key) => {
                stdin.pause();
                stdin.setRawMode(false);

                if (key.sequence === EOT) {
                    kill(pid, "SIGINT");
                }

                resolve(key.sequence || "");
            });
        });

        return keyPromise;
    }
}

export default StdIO;

function renderTextNode(node: TextNode): string {
    let text = node.content;

    if (node.useMCSSCodes) {
        return renderSSCodes(text, node.ssCodesPreview || false) + RESET;
    }

    let hasStyle = false;

    if (!isEmpty(node.color)) {
        text = colorIndex(node.color) + text;
        hasStyle = true;
    }

    if (!isEmpty(node.backgroundColor)) {
        text = colorIndex(node.backgroundColor, true) + text;
        hasStyle = true;
    }

    if (node.bold) {
        text = BOLD + text;
        hasStyle = true;
    }

    if (node.dim) {
        text = DIM + text;
        hasStyle = true;
    }

    if (node.itallic) {
        text = ITALLIC + text;
        hasStyle = true;
    }

    if (hasStyle) {
        return text + RESET;
    } else {
        return text;
    }
}
