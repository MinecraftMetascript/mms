declare global {
    function updateFile(filename: string, content: string, callback: ProjectUpdateHook): void

    function getFileDiag(filename: string, callback: (serial: string) => void)
}

export type ProjectUpdateHook =
    (serial: string) => unknown

export type FileTreeLike = {
    name: string,
    data?: unknown
} & ({ isDir: true, children?: Record<string, FileTreeLike> } | { isDir: false, content?: string })


export type MmsTextLocation = {
    Start: {
        Line: number,
        Column: number
    },
    StartIdx: number,
    Stop: {
        Line: number,
        Column: number
    },
    StopIdx: number,
    Text: string,
    Filename: string
}

export type MmsReference = `${string}:${string}`

export type MmsSymbol = {
    nameLocation: MmsTextLocation,
    contentLocation: MmsTextLocation,
    value: unknown,
    ref: MmsReference
    type: string
}


export class Go {
    argv: string[]
    env: Record<string, string>
    exit: (code: number) => void
    readonly importObject: WebAssembly.Imports

    constructor()

    run(mod: WebAssembly.Instance): Promise<void>
}