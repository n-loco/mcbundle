import { kill, pid, stdin, stdout } from "node:process";
import readline, { emitKeypressEvents, Key } from "node:readline";
import { TextNode, TextSpan } from "./text.js";
import { ArrowSequence, BEL, BOLD, colorIndex, CR, DIM, EOT, EREASE_EOE, EREASE_EOL, ITALLIC, LF, RESET } from "./ansi.js";
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

    export function lineFeed() {
        stdout.write(LF);
    }

    export function ereaseScreen() {
        stdout.write(EREASE_EOE);
    }

    export function ereaseLine() {
        stdout.write(EREASE_EOL);
    }

    export function moveCursor(X: number, y: number) {
        stdout.write(ArrowSequence.seq(X, y));
    }

    export function carriageReturn() {
        stdout.write(CR);
    }

    export function verticalReturn(lines: number) {
        stdout.write(CR + ArrowSequence.seq(0, lines));
    }

    let stdinLock = false;
    export async function keypress(): Promise<string> {
        if (stdinLock) {
            throw new Error("stdin locked");
        }

        stdinLock = true;
        emitKeypressEvents(stdin);
        stdin.setRawMode(true);
        stdin.resume();

        const keyPromise = new Promise<string>((resolve) => {
            stdin.once("keypress", (_: string, key: Key) => {
                stdin.pause();
                stdin.setRawMode(false);
                stdinLock = false;

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
