# Project Coding Rules (Non-Obvious Only)

- Use build tags `//go:build !js && !wasm` for native code, `//go:build js && wasm` for WASM code
- Project diagnostics accessed via `project.Diagnostics()` method, not field access
- File parsing requires both AST nodes and file tree structure generation for export
- Custom `FileTreeLike` structure used for JSON export to Minecraft format
- Namespace declarations must use specific `namespace:name` syntax for cross-references
- Keep all imports at the top of files; never insert imports mid-file
- After dependency changes: `go mod tidy && go mod vendor` (required)