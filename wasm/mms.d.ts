declare global {
    function updateFile(filename: string, content: string, callback: ProjectUpdateHook): void

    function getFileDiag(filename: string, callback: (serial: string) => void)
}

export type ProjectUpdateHook =
/**
 * @param {string} root JSON Serialized `FileTreeLike`
 * @param {string} symbols JSON Serialized `Record<string, MmsSymbol>`
 */
    (root: string, symbols: string) => unknown

export type FileTreeLike = {
    name: string,
    data?: unknown
} & ({ isDir: true, children?: Record<string, FileTreeLike> } | { isDir: false, content?: string })

export type MmsSymbol = {
    nameLocation: unknown,
    contentLocation: unknown,
    value: unkonwn,
    ref: unknown
}

export class Go {
    argv: string[]
    env: Record<string, string>
    exit: (code: number) => void
    readonly importObject: WebAssembly.Imports

    constructor()

    run(mod: WebAssembly.Instance): Promise<void>
}