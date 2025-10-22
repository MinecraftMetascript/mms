# Project Debug Rules (Non-Obvious Only)

- LSP server requires specific trigger characters: `.`, `(`, `)`, `:`, `=`, ` `, `,`
- Debug builds show serialized project symbols in logs (not just diagnostics)
- Project rebuilds symbol table on every file change via `AddFile()` method
- WASM build logs prefixed with `[MMS:WASM]:` and uses custom logging setup
- Native build uses standard log package, WASM build uses custom LSP logging
- Testing pattern: Create Project, add File, call Parse(), assert on Diagnostics()