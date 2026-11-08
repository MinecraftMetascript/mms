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

## How the overhaul landed

All eight workstreams are merged into `main`. Some done conditions above describe a design that changed on the way, and the code is the reference:

- **`types` resolves, `sema` checks.** `engine/types` turns MMS types into mcdoc types, matches unions and checks Blueprint fields, Template arguments and typed consts at the user's spans, but it hands the checking of values against vanilla-mcdoc to `sema`. `sema/check.go` stays as that checker; it was not deleted (W3).
- **Kinds come from the spec and the prelude.** `sema/kinds.go` builds the kind table from vanilla-mcdoc folders and the prelude's `extern Blueprint`s, with a short alias table, instead of the old hardcoded map. `sugar.go` became `sema/shorthand.go` (W4).
- **Markup lives in `sema`.** There is no `text` package: markup is parsed and checked in `sema/bookmarkup.go`, `book.go`, `text.go` and `tell.go` (W5).
- **One simulator.** `mms test`, the LSP and both editors run the Go interpreter in `engine/internal/mcfn`. The web editor's function-run preview, which used a second simulator, was removed (W7).

The double type check of typed consts and instances (one in `types` for diagnostics, another in `project` for hover info) should become one: `types.Env.Check` returning its checker so `project` reuses it.

The earlier rewrite milestones (M0 to M6) are tracked under issue [#22](https://github.com/MinecraftMetascript/mms/issues/22).
