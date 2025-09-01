declare global {

}

export class Go {
    argv: string[]
    env: Record<string,string>
    exit: (code: number) => void

    constructor()
    run: Promise<void>
}