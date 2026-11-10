<p align="center">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="branding/logo-dark.svg">
    <img src="branding/logo-light.svg" alt="MMS Metascript" height="80">
  </picture>
</p>

# Minecraft Metascript (MMS)

MMS is a language that compiles to Minecraft Java Edition data packs (and the resource pack they need). One `.mms` file stands in for the many JSON and `.mcfunction` files a feature takes: worldgen, loot tables, recipes, advancements, tags, custom items, books, jigsaw structures and functions. Every value is checked against the game's own schemas and command tree as you type, and the editors preview what you are building.

Target: Minecraft Java Edition 26.3 (data pack format 121).

```mms
/** Gems that drop from the gem chest. */
ItemTag gems = [diamond, emerald, amethyst_shard]

LootTable chests/gems {
  pools = [Pool(1..3, #gems)]
}

Recipe gem_block = Shaped(["GG", "GG"], { G = #gems }, emerald_block)

Objective found

Function find {
  found[@s] += 1
  /loot give @s loot gems:chests/gems
  tell(@s, "You have found ${found[@s]} gem chests")
}

/** Counts a find when right-clicked. */
CustomItem gem_finder {
  base = compass
  name = "Gem Finder"
  on_right_click = find
}
```

`mms build` turns that into a loot table, a recipe, an item tag, the function, the advancement and hook functions behind the right click, and the load and tick functions that set up the scoreboard.

Beyond declarations, the language has `const` values and `for` loops that run at build time, Templates and Blueprints for reuse (`CustomItem`, `Book` and `JigsawStructure` are Blueprints written in MMS itself), typed `Storage` records and `let` variables in functions, named selectors, chat and book text with inline markup (`.mmstext` and `.mmsbook` files), and `Test` declarations that run functions without starting the game.

## Get it

- **Web editor:** [minecraftmetascript.github.io/mms](https://minecraftmetascript.github.io/mms/) runs in the browser with nothing to install, and downloads the pack as a zip or builds it into a world folder.
- **VS Code:** install **Minecraft Metascript** (`minecraftmetascript.mms`) from the VS Code Marketplace, or from Open VSX for Cursor, VSCodium and other editors. It includes `mms` for its own use, so the editor needs nothing else; to run the commands below in a terminal, install the command line too. See [its README](editors/vscode/README.md) for what it does.
- **Command line:** download `mms` for Linux, macOS or Windows from the [releases page](https://github.com/MinecraftMetascript/mms/releases), or run `go install github.com/minecraftmetascript/mms/engine/cmd/mms@latest`.

```sh
mms init -namespace gems gems   # mms.json and a starter file
mms build gems                  # check, then write the data pack to gems/build
mms build -world ~/.minecraft/saves/MyWorld gems   # or install it into a world
mms check gems                  # report problems without writing anything
mms test gems                   # run the project's Test declarations
mms fix gems                    # apply every safe lint fix
mms fmt gems                    # format every .mms file in gems
mms import pack.zip proj        # turn an existing data pack (folder or zip) into MMS
mms lsp                         # language server over stdio, for other editors
mms mcp                         # MCP server over stdio, for AI agents that write MMS (-check tests it)
mms reference                   # every kind, helper and prelude declaration, as JSON
```

## Documentation

Everything is on the [docs site](https://minecraftmetascript.github.io/mms/docs):

- [Getting started](https://minecraftmetascript.github.io/mms/docs/start/getting-started), then a tutorial: [a magic wand](https://minecraftmetascript.github.io/mms/docs/start/tutorial-custom-item), [a wave arena](https://minecraftmetascript.github.io/mms/docs/start/tutorial-minigame), [a guide book](https://minecraftmetascript.github.io/mms/docs/start/tutorial-book), [a custom dimension](https://minecraftmetascript.github.io/mms/docs/start/tutorial-dimension) or [a jigsaw dungeon](https://minecraftmetascript.github.io/mms/docs/start/tutorial-structure).
- Guides to [resources](https://minecraftmetascript.github.io/mms/docs/guides/resources), [functions](https://minecraftmetascript.github.io/mms/docs/guides/functions), [state](https://minecraftmetascript.github.io/mms/docs/guides/state), [constants, templates and loops](https://minecraftmetascript.github.io/mms/docs/guides/repetition), [Blueprints](https://minecraftmetascript.github.io/mms/docs/guides/blueprints), [testing](https://minecraftmetascript.github.io/mms/docs/guides/testing), [terrain](https://minecraftmetascript.github.io/mms/docs/guides/terrain), [building and playing](https://minecraftmetascript.github.io/mms/docs/guides/building), [the editors](https://minecraftmetascript.github.io/mms/docs/guides/editors) and [lint rules](https://minecraftmetascript.github.io/mms/docs/guides/lint).
- The [reference](https://minecraftmetascript.github.io/mms/docs/reference): the language itself, then every kind, helper, Blueprint and lint rule.
- For AI agents: the [`mms mcp` server](https://minecraftmetascript.github.io/mms/docs/guides/agents), and [llms.txt](https://minecraftmetascript.github.io/mms/llms.txt) and [llms-full.txt](https://minecraftmetascript.github.io/mms/llms-full.txt).

The [examples](examples/) are whole projects (worldgen, custom items, books, jigsaw structures, minigames). The web editor lists them, and the tests build every one.

## Repository

| Folder | Contents |
| --- | --- |
| `engine/` | Go module: spec pack, parser, evaluator, type checker, compiler, prelude, test interpreter, language server, CLI and WASM build |
| `web/` | Web editor and docs site (SvelteKit, CodeMirror 6), deployed to GitHub Pages |
| `editors/vscode/` | VS Code extension, a client for `mms lsp` |
| `packages/preview/` | Svelte previews shared by the web editor and the VS Code extension |
| `site/` | The docs site's pages, as Markdown that mirrors the `/docs` URLs |
| `examples/` | Example projects |
| `docs/` | The plan and task list behind the language |
| `branding/` | Logo, mark and favicon |

To work on it, run `go test ./...` in `engine/`, and `npm ci` in the repository root for the npm workspace (`packages/preview`, `web` and `editors/vscode`). The [Taskfile](Taskfile.yml) has shortcuts for the builds CI runs.

## Credits

MMS stands on other people's work:

- [Deepslate](https://github.com/misode/deepslate) by [Misode](https://github.com/misode) (MIT) renders every worldgen preview in the web editor and the VS Code panel.
- [mcmeta](https://github.com/misode/mcmeta) by Misode provides the vanilla data, assets, registries and command tree that previews and checks read.
- [vanilla-mcdoc](https://github.com/SpyglassMC/vanilla-mcdoc) by [SpyglassMC](https://github.com/SpyglassMC) (MIT) is the source of the data pack schemas the engine checks against.

License notices for what each build bundles: `engine/THIRD_PARTY_NOTICES.txt` (copied into every CLI release), `third-party-licenses.txt` on the web editor (generated by `npm run build`), and `out/third-party-licenses.txt` in the VS Code extension.
