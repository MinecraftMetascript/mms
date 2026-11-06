# Plan

MMS is being overhauled into a more cohesive language. Two docs hold the full picture:

- [MMS language design proposal](https://claude.ai/artifact/GBf7rorbHR1BKW4wjqUczj): the syntax and semantics, with examples.
- [MMS overhaul implementation plan](https://claude.ai/artifact/WbZaZqkoNX1WxctKpipZ4t): the architecture, the editor DX parity table, workstreams, testing and risks.

This file keeps the decisions the code depends on, so the repo stands on its own.

## Foundations (unchanged by the overhaul)

- Go plus ANTLR 4. One engine of memoised queries serves three hosts: the CLI, the LSP over stdio, and the same LSP compiled to WASM inside a Web Worker.
- Target the latest stable Minecraft release, 26.3 (pack format 121). Version selection comes later.
- Scope: all data pack content (worldgen, loot tables, predicates, item modifiers, recipes, advancements, tags, functions).
- Output schemas come from vanilla-mcdoc `symbols.json` filtered to 26.3; registries, block states and the command tree come from mcmeta `26.3-summary`.
- Functions get a real backend: typed IR, slot allocation on scoreboards and storage, control flow lowered to sub-functions, a generated load function.
- Casing (Brian, 2026-10-31): declaration kinds and helpers are PascalCase (`Namespace`, `DensityFunction`, `Noise(...)`). Resource names stay lowercase snake_case. Keywords (`const`, `let`, `if`, `else`, `while`, `for`, `return`) are lowercase.
- Repo layout (Brian): Go module in `engine/`, editor app in `web/`.
- Infix math is in (Brian), guarded by a precedence lint and formatter parentheses.
- Formatter: two-space indent, 100 columns, no other options.
- Hosting (Brian): the web editor is deployed to GitHub Pages, with no telemetry.
- Web stack (Brian): SvelteKit 3 (currently 3.1.x) with Tailwind CSS (currently 4.x), built with the static adapter.

## Overhaul decisions (Brian, 2026-11-06)

Process:

- One big rewrite. There are no users yet, so there is no phased rollout and no migration path.
- No deprecations: when syntax or a name is replaced, the old form is removed outright, with no aliases.
- Work lands on a long-lived `overhaul` branch through ordinary pull requests. `main` gets only urgent fixes until `overhaul` merges, which happens once every row of the editor DX parity table works in VS Code and the web editor.
- Editor DX parity with today (lints, quick fixes, completion, hover, rename, previews and the rest) is a merge requirement.
- Zed support is dropped: the Zed extension, its tree-sitter grammar and `zed.yml` are deleted.
- Lint rule ids are renamed where they name a removed concept or use jargon (for example `jigsaw-hole` becomes `jigsaw-unknown-name` and `use-sugar` becomes `use-shorthand`). Old ids are unknown-rule errors.
- Built output does not have to match today's byte for byte. The examples are refreshed as part of the overhaul, and pull requests that change one show its built-pack diff for review.

Language:

- `const` (lowercase, like `let`) names any value, not only values that become their own JSON file: pools, predicates, pool elements, text, records.
- `Template` for reusable values, `Type` for types, and `Blueprint` (with `fields`, `check`, `as T = expr hint(...)` and `emit`) for user-defined declaration kinds. Union types match the first arm in order; ambiguity is not an error.
- A blueprint forwards to another kind with an explicit `emit { primary Kind ${self.path} { ... } }`, at most one per blueprint, with the same id. `emit` only needs to cover CustomItem, Book and JigsawStructure for now.
- Top-level `for` with templated names, list comprehensions and spread. No HCL-style `for_each` and no module system.
- The prelude, written in MMS and embedded in the engine, defines the built-in kinds (Namespace, Function, Objective, Global, the On* hooks, CustomItem, Book, JigsawStructure, StructureFile) and every former `sugar.go` rule. What MMS cannot express stays in Go as `extern` intrinsics.
- Every string in a text-typed position is markup, chosen by type; `\{` escapes. There is no `Text` declaration kind: named text is `const x: Text = ...` or `Template f(...): Text`.
- Two text file kinds: `.mmsbook` keeps `---` page breaks and the page-overflow lint; `.mmstext` has neither, and `---` there is an error.
- Inline function bodies always use `run { }`. Commands templates use plain call spelling, marked in the editor by an `inlined` inlay and hover. Template bodies always unroll; range loops in functions stay run-time.
- Deferred: libraries ([#189](https://github.com/MinecraftMetascript/mms/issues/189); the prelude loads the way a library will, and the `libraries` key in `mms.json` is reserved) and translations ([#188](https://github.com/MinecraftMetascript/mms/issues/188)).

## Workstreams

W1 grammar, W2 evaluator and provenance, W3 type checker, W4 prelude core, W5 text, W6 jigsaw, W7 function side, W8 editors and docs. Details and done conditions are in [tasks.md](tasks.md).
