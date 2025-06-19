import { Prettify } from "../utils.js";

export type TextStyle = {
    useMCSSCodes?: false;
    color?: number;
    backgroundColor?: number;
    bold?: boolean;
    dim?: boolean;
    itallic?: boolean;
} | {
    useMCSSCodes: true;
    ssCodesPreview?: boolean;
};

export type TextNode = Prettify<{ content: string } & TextStyle>;

export type TextSpan = (TextNode | string)[]
