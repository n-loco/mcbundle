export const EOT = "\x03";
export const BEL = "\x07";
export const CR = "\r";
export const LF = "\n";
export const ESC = "\x1b";
export const SPACE = " ";
export const BP = "\x08";
export const DEL = "\x7f";

export const RESET = `${ESC}[0m`;
export const RESET_BD = `${ESC}[22m`;

export const BOLD = `${ESC}[1m`;
export const DIM = `${ESC}[2m`;
export const ITALLIC = `${ESC}[3m`;

export const EREASE_EOE = `${ESC}[0J`;
export const EREASE_EOL = `${ESC}[0K`;

export namespace ArrowSequence {
    const UP = `${ESC}[A`;
    const DOWN = `${ESC}[B`;
    const RIGHT = `${ESC}[C`;
    const LEFT = `${ESC}[D`;

    export function vec(seq?: string): [number, number] | null {
        switch (seq) {
            case UP:
                return [0, 1];
            case DOWN:
                return [0, -1];
            case RIGHT:
                return [1, 0];
            case LEFT:
                return [-1, 0];
        }

        return null;
    }

    export function seq(x: number, y: number): string {
        const ySeq = y !== 0
            ? `${ESC}[${Math.abs(y)}${y > 0 ? "A" : "B"}`
            : "";

        const xSeq = x !== 0
            ? `${ESC}[${Math.abs(x)}${x > 0 ? "C" : "D"}`
            : "";

        return `${xSeq}${ySeq}`;
    }
}

function commaRGB(color: number, fakeDim = false): string {
    const red = (color & 0xff0000) >> 16;
    const green = (color & 0x00ff00) >> 8;
    const blue = color & 0x0000ff;

    if (fakeDim) {
        return `2;${Math.floor(red / 2)};${Math.floor(green / 2)};${Math.floor(blue / 2)}`;
    }

    return `2;${red};${green};${blue}`;
}

export function colorSequence(color: number, fakeDim = false) {
    return `${ESC}[38;${commaRGB(color, fakeDim)}m`;
}

export function bgColorSequence(color: number) {
    return `${ESC}[48;${commaRGB(color)}m`;
}

export function isControl(char: string) {
    const fCodePoint = char.codePointAt(0) || 0;
    if (0 <= fCodePoint && fCodePoint <= 0x1f) return true
    if (fCodePoint === 0x7f) return true;
    return false;
}
