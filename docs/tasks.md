# Tasks

The overhaul is split into eight workstreams, each a series of pull requests into the `overhaul` branch. The [implementation plan](https://claude.ai/artifact/WbZaZqkoNX1WxctKpipZ4t) has the detail behind each one, including the editor DX parity table that gates the final merge.

W1 to W4 run in order. W5 to W8 can run side by side after W4, and W8 starts as soon as W2 lands so the editors track the engine.

| # | Workstream | Covers | Done when |
| --- | --- | --- | --- |
| W1 | Grammar and parser | New syntax in the ANTLR grammar, AST, error recovery and chunks; `engine/format` prints every new form. Delete the Zed extension and `zed.yml`. | Every snippet in the design doc parses and round-trips through `mms fmt` unchanged; fuzz and recovery tests pass. |
| W2 | Evaluator and provenance | New `eval` package: `const`, `for`, Template calls, Blueprint `emit` with `primary`, a step budget, and provenance from every value to the user span that caused it. | Golden tests show expanded output and provenance for each design doc example. |
| W3 | Type checker | New `types` package: vanilla-mcdoc types, `Type` declarations, first-match unions, blueprint fields and `check`, `as T = expr hint(...)`; diagnostics mapped through provenance. | The `sema` golden tests pass against `types`, and `sema/check.go` and its helpers are deleted. |
| W4 | Prelude core | Embedded prelude loaded like a library, `extern` intrinsics, Namespace, Function, Objective, Global, On* hooks, CustomItem, all former `sugar.go` rules, renamed lint ids. | The hardcoded kind map and `sugar.go` are gone; the custom item examples are rewritten and build cleanly. |
| W5 | Text | New `text` package, type-directed markup, `.mmsbook` and `.mmstext`, `tell`/`title`/`actionbar`, native score and NBT components, the Book blueprint, accidental-markup and page-break-in-text lints. | Book examples are rewritten and build cleanly; the book renderer and a new chat preview share one component model. |
| W6 | Jigsaw | JigsawStructure blueprint and StructureFile, jigsaw intrinsics, jigsaw lints and fixes, jigsaw graph and give command. | `examples/jigsaw` and cave-siege are rewritten with JigsawStructure; every jigsaw lint has a fixture and its fix applies. |
| W7 | Function side | `Storage` records, selector values, Commands templates with the `inlined` inlay, `run { }`, `Test` declarations, `mms test`, simulator support. | Function examples are rewritten with `Test` declarations, and `mms test` gives the same results in Go and in the web simulator. |
| W8 | Editors and docs | Every row of the parity table, `mms/expand`, VS Code (lenses, views, `mmstext` language, TextMate grammars, a new extension host test), web editor (tokenizer, previews, reference), importer, `docs/*.md`, refreshed examples. | Every parity table row is checked off by the LSP, Playwright and VS Code tests. |

The earlier rewrite milestones (M0 to M6) are tracked under issue [#22](https://github.com/MinecraftMetascript/mms/issues/22).
