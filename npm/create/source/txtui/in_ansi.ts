import { kill, pid, platform as os, stdin } from "node:process";
import fs from "node:fs";

export const enum CharCodes {
    EndOfText = 0x03,
    CarriageReturn = 0x0d,
    Escape = 0x1b,
    A = 0x41,
    B = 0x42,
    C = 0x43,
    D = 0x44,
    OpenSquareBracket = 0x5b,
}

export function handleEnter(data: Uint8Array): boolean {
    return data.length === 1 && data[0] === CharCodes.CarriageReturn;
}

export enum Arrow {
    Up,
    Down,
    Left,
    Right,
}

export namespace Arrow {
    export function x(arrow: Arrow): number {
        switch (arrow) {
            case Arrow.Right:
                return 1;
            case Arrow.Left:
                return -1;
        }
        return 0;
    }

    export function y(arrow: Arrow): number {
        switch (arrow) {
            case Arrow.Up:
                return -1;
            case Arrow.Down:
                return 1;
        }
        return 0;
    }
}

export function handleArrow(data: Uint8Array): Arrow | null {
    if (data.length !== 3) return null;

    const isControl =
        data[0] === CharCodes.Escape &&
        data[1] === CharCodes.OpenSquareBracket;

    const isArrow =
        isControl &&
        (CharCodes.A <= data[2] || data[2] <= CharCodes.D);

    if (isArrow) {
        switch (data[2]) {
            case CharCodes.A:
                return Arrow.Up;
            case CharCodes.B:
                return Arrow.Down;
            case CharCodes.C:
                return Arrow.Right;
            case CharCodes.D:
                return Arrow.Left;
        }
    }

    return null;
}

const fd = os == "win32" ? stdin.fd : fs.openSync("/dev/tty", "rs");

export function interact(react: (data: Uint8Array, done: () => void) => void, autoHandleEnter = true) {
    stdin.resume();
    const ogMode = stdin.isRaw;
    stdin.setRawMode(true);

    const buffer = Buffer.alloc(3);

    let keepInteracting = true;
    const done = () => {
        keepInteracting = false;
    }

    while (keepInteracting) {
        const bytesRead = fs.readSync(fd, buffer, 0, 3, null);
        const data = new Uint8Array(buffer.subarray(0, bytesRead));

        if (data.length === 1) {
            if (data[0] === CharCodes.EndOfText) {
                kill(pid, "SIGINT");
            }
            if (autoHandleEnter && handleEnter(data)) {
                break;
            }
        }

        react(data, done);
    }

    stdin.setRawMode(ogMode);
    stdin.pause();
}
