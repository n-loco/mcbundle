export interface TextNode {
    content: string;
    useMCSSCodes?: boolean;
    ssCodesPreview?: boolean;
    color?: number | null;
    backgroundColor?: number | null;
    bold?: boolean;
    dim?: boolean;
    itallic?: boolean;
}

export type TextSpan = Array<TextNode | string>
