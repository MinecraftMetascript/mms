# AGENTS.md

This file provides guidance to agents when working with code in this repository.

## Additional AI Rules

- Project-specific AI rules available in `.windsurfrules` file
- Includes build preferences, testing guidance, and coding style requirements

## Build System & Commands

- Run `nix build` for native binary (auto-generates ANTLR4 parser), not `go build`
- WASM build requires `GOOS=js GOARCH=wasm` and vendored dependencies
- Manual ANTLR4 regeneration: `antlr4 -Dlanguage=Go grammar/MinecraftMetascript.g4 -o lang -package grammar`
- LSP server runs embedded in WASM build via custom stream implementation

## Language & Project Structure

- Single ANTLR4 grammar file generates all language constructs (unusual for complex DSLs)
- Examples in `examples/test_files/` serve as both documentation and integration tests
- Generated Minecraft configs use specific directory structure: `worldgen/dimension_type/`, `worldgen/biome_source/`, etc.
- Namespace references use `namespace:name` syntax across files

## WASM Integration

- WASM build exports `updateFile()`, `getFileDiag()`, `getMmsSpec()` functions to JavaScript
- LSP protocol runs over custom WASM stream bridge (not standard stdio)
- Global `mmsLspRead` and `mmsLspWrite` functions must be available in JS environment

## Code Patterns

- Build tags: `//go:build !js && !wasm` for native, `//go:build js && wasm` for WASM
- Project diagnostics accessed via `project.Diagnostics()` method, not field
- File parsing returns both AST nodes and generates file tree structure for export
- Custom `FileTreeLike` structure used for JSON export to Minecraft format
- Keep all imports at the top of files; never insert imports mid-file

## Testing & Development

- No formal test files (uses examples as integration tests)
- LSP server requires specific trigger characters: `.`, `(`, `)`, `:`, `=`, ` `, `,`
- Project rebuilds symbol table on every file change via `AddFile()` method
- Debug builds show serialized project symbols in logs (not just diagnostics)
- Testing pattern: Create Project, add File, call Parse(), assert on Diagnostics()
- After dependency changes: `go mod tidy && go mod vendor` (required)