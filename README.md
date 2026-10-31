# Minecraft Metascript (MMS)

MMS is a small language that compiles to Minecraft data packs. One MMS file can stand in for the many JSON files a feature needs (noise, density functions, surface rules, loot tables, recipes, functions and more). It also gives functions helpers for globals, variables and scoreboards.

This branch is a full rewrite. The plan is in [docs/plan.md](docs/plan.md), progress is in [docs/tasks.md](docs/tasks.md), and every judgement call is logged in [docs/decisions.md](docs/decisions.md).

| Folder | Contents |
| --- | --- |
| `engine/` | Go module: spec pack, parser, analysis engine, compiler, language server, CLI and WASM entry |
| `web/` | Browser editor (SvelteKit, CodeMirror 6, Deepslate previews), deployed to GitHub Pages |
| `docs/` | Plan, task list, decision log, language reference |

Target: Minecraft Java Edition 26.3 (data pack format 121).
