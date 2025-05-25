export interface CompilerOptions {
    esModuleInterop: boolean,
    module: string,
    moduleResolution: string,
    strict: boolean,
}

export interface TSConfig {
    compilerOptions: CompilerOptions,
}

export function createTSConfig(): TSConfig {
    return {
        compilerOptions: {
            esModuleInterop: true,
            module: "NodeNext",
            moduleResolution: "nodenext",
            strict: true,
        },
    }
}
