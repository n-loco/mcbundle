import { ArrowSequence, BP, CR, SPACE } from "./ansi.js";
import Figures from "./figures.js";
import StdIO from "./stdio.js";
import { TextSpan } from "./text.js";

export * from "./text.js";

export function print(str: string | TextSpan) {
    StdIO.write(str);
}

export function println(str: string | TextSpan) {
    let rStr = str;

    if (typeof rStr === "string") {
        rStr += "\n";
    } else {
        rStr.push("\n");
    }

    StdIO.write(rStr);
}

export function bell() {
    StdIO.bell();
}

export interface Selectable {
    normal: string | TextSpan;
    highlighted: string | TextSpan;
    choosen?: string | TextSpan;
}

export interface SelectMenuOptions {
    label: string | TextSpan;
    options: Selectable[];
}

export async function selectMenu(opts: SelectMenuOptions): Promise<number> {
    let choice = 0;

    renderLabel({
        label: opts.label,
        done: false,
        lf: true,
        hint: `${Figures.ARROW_UP} ${Figures.ARROW_DOWN} = move; ENTER = submit`,
    });

    const renderOptions = () => {
        StdIO.ereaseScreen();
        for (let i = 0; i < opts.options.length; i++) {
            const option = opts.options[i];
            const prefix = i === choice ? `   ${Figures.POINTER} ` : "     ";

            StdIO.write(prefix);
            StdIO.write(i === choice ? option.highlighted : option.normal);
            StdIO.lineFeed();
        }
    };

    renderOptions();

    while (true) {
        const key = await StdIO.keypress();
        if (key === CR) break;

        const move = ArrowSequence.vec(key);
        if (move === null || move[1] === 0) {
            StdIO.bell();
            continue;
        }

        let [_, y] = move;
        choice = (choice - y + opts.options.length) % opts.options.length;
        StdIO.verticalReturn(opts.options.length);
        renderOptions();
    }

    StdIO.verticalReturn(opts.options.length + 1);
    
    renderLabel({
        label: opts.label,
        done: true,
        lf: false,
    });

    StdIO.write(opts.options[choice].choosen || opts.options[choice].normal);
    StdIO.lineFeed();

    return choice;
}

export interface CheckMenuOptions {
    label: string | TextSpan;
    options: Selectable[];
    defaultValues?: boolean[];
}

export async function checkMenu(opts: CheckMenuOptions): Promise<boolean[]> {
    let pointer = 0;
    const choices = (() => {
        let chs = opts.defaultValues || [];
        const ogLen = chs.length;
        chs.length = opts.options.length;
        return chs.fill(false, ogLen, opts.options.length);
    })();

    renderLabel({
        label: opts.label,
        done: false,
        lf: true,
        hint: `${Figures.ARROW_UP} ${Figures.ARROW_DOWN} = move; SPACE = select/unselect; ENTER = submit`,
    });

    const renderOptions = () => {
        StdIO.ereaseScreen();
        for (let i = 0; i < opts.options.length; i++) {
            const option = opts.options[i];
            const prefix = i === pointer ? `   ${Figures.POINTER} ` : "     ";
            const checkBox = choices[i]
                ? `${Figures.CHECKED_BOX} `
                : `${Figures.UNCHECKED_BOX} `;

            StdIO.write(prefix + checkBox);
            StdIO.write(i === pointer ? option.highlighted : option.normal);
            StdIO.lineFeed();
        }
    };

    renderOptions();

    while (true) {
        const key = await StdIO.keypress();
        if (key === CR) break;

        if (key === SPACE) {
            choices[pointer] = !choices[pointer];
            StdIO.verticalReturn(opts.options.length);
            renderOptions()
            continue;
        }

        const move = ArrowSequence.vec(key);
        if (move === null || move[1] === 0) {
            StdIO.bell();
            continue;
        }

        let [_, y] = move;
        pointer = (pointer - y + opts.options.length) % opts.options.length;
        StdIO.verticalReturn(opts.options.length);
        renderOptions();
    }

    StdIO.verticalReturn(opts.options.length + 1);
    
    renderLabel({
        label: opts.label,
        done: true,
        lf: false,
    });

    const choosenOptions = opts.options.filter((_, i) => choices[i]);

    if (choosenOptions.length !== 0) {
        for (let i = 0; i < choosenOptions.length; i++) {
            const option = choosenOptions[i];
            StdIO.write(option.choosen || option.normal);
            if (i < choosenOptions.length - 1) {
                StdIO.write([{ content: ", ", dim: true }]);
            }
        }
    } else {
        StdIO.write([{ content: "none", itallic: true, dim: true }]);
    }

    StdIO.lineFeed();

    return choices;
}

export interface BoolDialogOptions {
    label: string | TextSpan,
    truthy?: Selectable,
    falsy?: Selectable,
    defaultValue?: boolean
}

export async function boolDialog(opts: BoolDialogOptions): Promise<boolean> {
    let choice = opts.defaultValue || false;

    const truthy = opts.truthy || {
        normal: "Yes",
        highlighted: [{ content: "Yes", bold: true }],
        choosen: [{ content: "Yes", color: 0x29a0f5 }]
    };

    const falsy = opts.falsy || {
        normal: "No",
        highlighted: [{ content: "No", bold: true }],
        choosen: [{ content: "No", color: 0x29a0f5 }]
    };

    renderLabel({
        label: opts.label,
        done: false,
        lf: true,
        hint: `${Figures.ARROW_LEFT} ${Figures.ARROW_RIGHT} = move; ENTER = submit`,
    });

    const renderOptions = () => {
        StdIO.carriageReturn();
        StdIO.ereaseLine()
        StdIO.write([
            " ",
            choice ? { content: Figures.SMALL_PTR_L, dim: true } : " ",
            " "
        ]);
        StdIO.write(choice ? truthy.highlighted : truthy.normal)
        StdIO.write([
            " ",
            choice ? { content: Figures.SMALL_PTR_R, dim: true } : " ",
            " "
        ]);
        StdIO.write([{ content: Figures.SEP, dim: true }]);
        StdIO.write([
            " ",
            !choice ? { content: Figures.SMALL_PTR_L, dim: true } : " ",
            " "
        ]);
        StdIO.write(!choice ? falsy.highlighted : falsy.normal)
        StdIO.write([
            " ",
            !choice ? { content: Figures.SMALL_PTR_R, dim: true } : " ",
            " "
        ]);
    }

    renderOptions();

    while (true) {
        const key = await StdIO.keypress();
        if (key === CR) break;

        const move = ArrowSequence.vec(key);
        if (move === null || move[0] === 0) {
            StdIO.bell();
            continue;
        }

        let [x, _] = move;
        choice = Boolean((Number(choice) + x + 2) % 2);

        renderOptions();
    }

    StdIO.verticalReturn(1);
    
    renderLabel({
        label: opts.label,
        done: true,
        lf: false,
    });

    StdIO.write(choice
        ? truthy.choosen || truthy.normal
        : falsy.choosen || falsy.normal
    )
    StdIO.lineFeed();

    return choice;
}

export interface StringInputOptions {
    label: string | TextSpan;
    defaultValue?: string;
}

export async function stringInput(opts: StringInputOptions): Promise<string> {
    let strInput = "";

    renderLabel({
        label: opts.label,
        done: false,
        lf: true,
        hint: `${Figures.ARROW_LEFT} ${Figures.ARROW_RIGHT} = move; ENTER = submit`,
    });

    const renderStr = () => {
        StdIO.carriageReturn();
        StdIO.ereaseLine();
        StdIO.write([" ", { content: Figures.SMALL_POINTER, dim: true }, " "]);
        if (opts.defaultValue && strInput.length === 0) {
            StdIO.write([{ content: opts.defaultValue, dim: true }]);
            StdIO.moveCursor(-opts.defaultValue.length, 0);
        } else if (strInput.length > 0) {
            StdIO.write([{ content: strInput, useMCSSCodes: true, ssCodesPreview: true }]);
        }
    }

    renderStr();

    while (true) {
        const key = await StdIO.keypress();
        if (key === CR) break;

        const move = ArrowSequence.vec(key);
        if (move !== null && move[0] === 0) {
            StdIO.bell();
            continue;
        } else if (move !== null) {
            continue;
        }

        if (key === BP) {
            if (strInput.length > 0) {
                strInput = strInput.slice(0, strInput.length - 1);
            } else {
                StdIO.bell();
            }
        } else {
            strInput += key;
        }

        renderStr();
    }

    if (opts.defaultValue && strInput.length === 0) {
        strInput = opts.defaultValue;
    }

    StdIO.verticalReturn(1);
    renderLabel({
        label: opts.label,
        done: true,
        lf: false,
    })
    StdIO.write([{ content: strInput, useMCSSCodes: true }]);
    StdIO.lineFeed();

    return strInput;
}

interface RenderLabelOptions {
    label: string | TextSpan;
    done: boolean;
    lf: boolean;
    hint?: string;
}

function renderLabel(opts: RenderLabelOptions) {
    StdIO.ereaseScreen();
    if (opts.done) {
        StdIO.write([
            " ",
            { content: Figures.CHECK, color: 0x2bef76, bold: true },
            " ",
        ]);
    } else {
        StdIO.write([
            " ",
            { content: Figures.QUESTION, color: 0xffef0d, bold: true },
            " ",
        ]);
    }

    StdIO.write(opts.label);

    if (!opts.done && opts.hint) {
        StdIO.write(["   ", { content: opts.hint, dim: true }]);
    }

    if (opts.done) {
        StdIO.write(": ");
    }

    if (opts.lf) {
        StdIO.lineFeed();
    }
}
