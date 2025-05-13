import { stdout } from "node:process";
import { Arrow, handleArrow, handleEnter, interact } from "./in_ansi.js";
import { bold, colorIndex, ereaseFromCursorToEOE, moveCursorH, moveCursorV, reset } from "./out_ansi.js";

export function print(str: string, newLine = true) {
    stdout.write(str + (newLine ? "\n" : ""));
}

export namespace prompt {
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

        interact((data, done) => {
            if (handleEnter(data)) {
                stdout.write(moveCursorV(-options.length - 1) + "\r");
                stdout.write(moveCursorH(title.length) + ": ");
                stdout.write(ereaseFromCursorToEOE);
                stdout.write(colorIndex(0x6bceff) + options[pointer] + reset + "\n");
                done();
                return;
            }

            const arrow = handleArrow(data);

            if (arrow != null) {
                pointer = (options.length + pointer + Arrow.y(arrow)) % options.length;
            }

            render(true);
        }, false);

        return pointer;
    }
}
