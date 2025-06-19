export function isEmpty(value: unknown): value is (null | undefined) {
    return value === null || value === undefined;
}

export type MetaData = typeof MetaData;
export const MetaData = Symbol(".internal.metadata");

export type Prettify<T> = {
    [k in keyof T]: T[k];
} & {};
