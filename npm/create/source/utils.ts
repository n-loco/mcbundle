export function isEmpty(value: unknown): value is (null | undefined) {
    return value === null || value === undefined;
}

export type MetaData = typeof MetaData;
export const MetaData = Symbol(".internal.metadata");
