export const EOT = "\x03";
export const BEL = "\x07";
export const CR = "\r";
export const LF = "\n";
export const ESC = "\x1b";
export const SPACE = " ";
export const BP = "\x08";

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

export function colorRGB(r: number, g: number, b: number, bg = false): string {
    const red = `${Math.round(r)}`;
    const green = `${Math.round(g)}`;
    const blue = `${Math.round(b)}`;

    const space = bg ? "48" : "38";

    return `${ESC}[${space};2;${red};${green};${blue}m`
}

export function colorIndex(color: number, bg = false): string {
    color = Math.round(color);

    const red = (color & 0xff0000) >> 16;
    const green = (color & 0x00ff00) >> 8;
    const blue = color & 0x0000ff;

    return colorRGB(red, green, blue, bg);
}
