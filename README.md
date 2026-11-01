# Minecraft Metascript (MMS)

MMS is a small language that compiles to Minecraft data packs. One MMS file can stand in for the many JSON files a feature needs (noise, density functions, surface rules, loot tables, recipes, functions and more). It also gives functions helpers for globals, variables and scoreboards.

This branch is a full rewrite. The plan is in [docs/plan.md](docs/plan.md), progress is in [docs/tasks.md](docs/tasks.md). Judgement calls are explained in pull request descriptions and commit messages. The older [decision log](docs/decisions.md) (D1 to D94) is frozen and kept for reference.

| Folder | Contents |
| --- | --- |
| `engine/` | Go module: spec pack, parser, analysis engine, compiler, language server, CLI and WASM entry |
| `web/` | Browser editor (SvelteKit, CodeMirror 6, Deepslate previews), deployed to GitHub Pages |
| `docs/` | Guides, plan, task list, frozen decision log |
| `editors/vscode/` | VS Code extension (a client for `mms lsp`) |
| `examples/` | Example projects, built by the tests and listed in the web editor |

Target: Minecraft Java Edition 26.3 (data pack format 121).

## Command line

```sh
cd engine && go install ./cmd/mms
mms init my_pack          # mms.json and a starter file
mms build my_pack         # check and write the data pack to my_pack/build
mms fmt my_pack           # format every .mms file
mms import pack.zip proj  # turn an existing data pack (folder or zip) into MMS
mms lsp                   # language server over stdio, for editors
mms reference             # every kind and helper, as JSON
mms version               # the version and the Minecraft version it targets
```

Release builds of `mms` for Linux, macOS and Windows, and the VS Code extension (`editors/vscode`), are attached to each [GitHub release](https://github.com/MinecraftMetascript/mms/releases).

## Documentation

The guides are in `docs/` and on the [docs site](https://minecraftmetascript.github.io/mms/docs), which also has the generated reference:

- [Getting started](docs/getting-started.md)
- [Writing resources](docs/resources.md)
- [Functions](docs/functions.md)
- [Lint rules](docs/lint.md)
- [Tutorial: a custom dimension](docs/tutorial-dimension.md)
