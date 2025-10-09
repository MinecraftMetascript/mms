declare global {
    function updateFile(filename: string, content: string, dst: (a: Uint8Array) => void): void

    function getFileDiag(filename: string, callback: (serial: string) => void)

    /**
     * Exported by mms.wasm
     * @param input
     */
    function mmsLspWrite(input: string): void

    /**
     * Must be defined for mms.wasm to load properly.
     * @param output
     */
    function mmsLspRead(output: string): void
}

export type ProjectUpdateHook =
    (serial: string) => unknown

export type FileTreeLike = {
    name: string,
    data?: unknown
} & ({ isDir: true, children?: Record<string, FileTreeLike> } | { isDir: false, content?: string })


export type MmsSourceLocation = {
    start: {
        line: number,
        column: number,
        index: number
    },
    stop: {
        line: number,
        column: number,
        index: number
    },
    file: string
}

export type MmsReference = `${string}:${string}`

export type MmsSymbol = {
    nameLocation: MmsSourceLocation,
    location: MmsSourceLocation,
    value: object,
    ref: MmsReference
    kind: string
}


export class Go {
    argv: string[]
    env: Record<string, string>
    exit: (code: number) => void
    readonly importObject: WebAssembly.Imports

    constructor()

    run(mod: WebAssembly.Instance): Promise<void>
}