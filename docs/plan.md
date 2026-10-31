# Plan

The full architecture and implementation plan is the doc [MMS rewrite: compiler, LSP, and web editor plan](https://claude.ai/artifact/9kdQU1hDLuseV4E8ydVjM8). This file keeps the decisions the code depends on, so the repo stands on its own.

## Decided

- Go plus ANTLR 4. One engine of memoised queries serves three hosts: the CLI, the LSP over stdio, and the same LSP compiled to WASM inside a Web Worker.
- Target the latest stable Minecraft release, 26.3 (pack format 121). Version selection comes later.
- Scope: all data pack content (worldgen, loot tables, predicates, item modifiers, recipes, advancements, tags, functions).
- Output schemas come from vanilla-mcdoc `symbols.json` filtered to 26.3; registries, block states and the command tree come from mcmeta `26.3-summary`.
- Sugar is a hand-written lowering layer above the generated schema, chosen by expected type.
- Functions get a real backend: typed IR, slot allocation on scoreboards and storage, control flow lowered to sub-functions, a generated load function.
- Casing (Brian, 2026-10-31): declaration kinds and helpers are PascalCase (`Namespace`, `DensityFunction`, `Noise(...)`). Resource names stay lowercase snake_case. Statement keywords in function bodies (`let`, `if`, `else`, `while`, `for`, `return`) are lowercase.
- Repo layout (Brian): Go module in `engine/`, editor app in `web/`.
- Infix math is in (Brian), guarded by a precedence lint and formatter parentheses.
- Formatter: two-space indent, 100 columns, no other options.
- Hosting (Brian): the web editor is deployed to GitHub Pages, with no telemetry.
- Web stack (Brian): SvelteKit 3 (currently 3.1.x) with Tailwind CSS (currently 4.x), built with the static adapter.

## Milestones

M0 spec pack, M1 parser and formatter, M2 worldgen parity and core LSP, M3 web editor and previews, M4 all JSON resources and lint, M5 function backend, M6 1.0 release. Details and status are in [tasks.md](tasks.md).
