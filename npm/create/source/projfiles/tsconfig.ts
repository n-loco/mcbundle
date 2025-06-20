export interface CompilerOptions {
    esModuleInterop: boolean,
    module: string,
    moduleResolution: string,
    strict: boolean,
    allowJs: boolean,
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
            allowJs: true,
        },
    }
}
