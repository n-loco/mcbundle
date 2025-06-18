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
