# Project Architecture Rules (Non-Obvious Only)

- Single ANTLR4 grammar file generates all language constructs (unusual for complex DSLs)
- WASM build embeds LSP server via custom stream implementation (not standard stdio)
- Examples in `examples/test_files/` serve as integration tests, not separate test files
- Project generates both AST nodes and file tree structures for Minecraft JSON export
- LSP protocol runs over custom WASM stream bridge requiring global JS functions
- New constructs register via `traversal.ConstructRegistry.Register(...)`
- Generated grammar files in `lang/grammar/` must never be edited manually