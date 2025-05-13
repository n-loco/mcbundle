const ESC = "\x1b";

export const reset = `${ESC}[0m`;

export const bold = `${ESC}[1m`;
export const italic = `${ESC}[3m`;
export const underline = `${ESC}[4m`;
export const strikethrough = `${ESC}[9m`;

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

export function moveCursorV(lines: number): string {
    const absLines = Math.round(lines < 0 ? -lines : lines);
    const arrow = lines < 0 ? "A" : "B";
    return `${ESC}[${absLines}${arrow}`;
}

export function moveCursorH(chars: number): string {
    const absLines = Math.round(chars < 0 ? -chars : chars);
    const arrow = chars < 0 ? "D" : "C";
    return `${ESC}[${absLines}${arrow}`;
}

export const ereaseFromCursorToEOE = `${ESC}[0J`;
