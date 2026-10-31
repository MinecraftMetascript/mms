# Tasks

The rewrite broken into milestones and tasks. This file stands in for GitHub issues until issue creation is allowed for the integration (see D2 in [decisions.md](decisions.md)). `[x]` is done, `[~]` is in progress.

## M0 · Spec pack spike

- [~] Repo skeleton: `engine/` Go module, `web/` app folder, docs, CI workflow
- [ ] `specgen`: fetch vanilla-mcdoc `symbols.json` and mcmeta `26.3-summary` at pinned commits
- [ ] `specgen`: filter version gates to 26.3, keep data-pack kinds, write `engine/spec/data/26.3.json.gz`
- [ ] `spec`: loader and query API (types by path, dispatchers, resource kinds, registries, block states, commands)
- [ ] Coverage test: all density function types present; old hand-written constructs present
- [ ] Size and load-time report; WASM hello-world that loads the pack in a worker

## M1 · Parser and formatter

- [ ] Lexer and parser grammars (ANTLR 4, Go target), generated code checked in, CI regen check
- [ ] CST with trivia (hidden channel), stable node spans
- [ ] Custom error strategy, `Missing` and `Error` nodes, recovery anchors
- [ ] Recovery corpus of half-typed files with golden CST and diagnostics
- [ ] Doc IR printer and formatter rules; idempotence and token-equality tests
- [ ] `mms fmt` and `mms fmt --check`
- [ ] Fuzz targets for parser and formatter

## M2 · Worldgen parity and core LSP

- [ ] Query engine: inputs, memoised queries, fingerprints, cancellation, panic isolation
- [ ] Declarations, namespaces, name resolution rule
- [ ] Expected-type checker over the spec pack, with union narrowing
- [ ] Lowering registry and worldgen rules (conditions, density math, anchors, gradients, sequences, inlining)
- [ ] Inline definition lifting with stable names
- [ ] Emit: pack layout, `pack.mcmeta`, deterministic JSON, source maps, schema re-validation
- [ ] `mms.json` project file; CLI `build`, `check`, `init`
- [ ] LSP: sync, diagnostics, completion, hover, signature help, definition, references, rename, formatting, semantic tokens, symbols
- [ ] Old examples build to valid 26.3 JSON; native latency benchmarks

## M3 · Web editor and previews

- [ ] WASM entry with worker transport; supervisor and restart
- [ ] SvelteKit app, CodeMirror 6 with LSP client, multi-file projects in IndexedDB
- [ ] Sidebar trees: Sources, Files, Symbols; JSON preview with source links
- [ ] Deepslate preview worker: noise, density function, noise settings, with pan, zoom, hover, lock marker
- [ ] Zip export and import; example gallery
- [ ] GitHub Pages deploy workflow; Playwright smoke tests

## M4 · All JSON resources and lint

- [ ] Generic syntax and checking for every data-pack resource kind
- [ ] Sugar for loot tables, predicates, item modifiers, recipes, advancements, tags
- [ ] Lint framework and initial rules, suppressions, fix-all
- [ ] `mms import` and Paste JSON as MMS
- [ ] Vanilla 26.3 round-trip test
- [ ] Biome, loot and recipe previews

## M5 · Function backend

- [ ] Slice 1: raw commands checked against the command tree; `OnLoad`, `OnTick`
- [ ] Slice 2: `Global`, `Objective`, int expressions, `if`
- [ ] Slice 3: loops, calls, returns
- [ ] Slice 4: storage values, macros, interpolation
- [ ] Slice 5: recursion, optimiser passes
- [ ] IR interpreter tests; in-server differential tests

## M6 · 1.0 release

- [ ] VS Code extension
- [ ] Native binaries for Linux, macOS, Windows
- [ ] Docs site generated from the spec pack and overlays
- [ ] Custom-dimension tutorial built end to end
