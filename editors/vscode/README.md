# Minecraft Metascript for VS Code

Language support for [Minecraft Metascript](https://github.com/MinecraftMetascript/mms) (`.mms`) files.

- **Diagnostics** for errors, lint findings and every raw command, checked against the Minecraft command tree.
- **Editing help:**
  - completion, hover and signature help
  - go to definition, find references and rename
- **Fixes:**
  - quick fixes
  - "Fix all" (`source.fixAll.mms`)
  - "Extract to a declaration" for inline values (`refactor.extract`), which then opens the rename box on the new name
- **Formatting** with the opinionated formatter (format on save is on by default for `.mms`).
- **Semantic highlighting.**
- **Previews:** click **Preview** above a noise, density function, noise settings, surface rule or dimension to see it rendered beside the editor (density, blocks, 3D, biome map), or above a loot table, recipe or biome to inspect it. The panel stays on that declaration while you edit elsewhere and redraws as you type; click another **Preview** to switch. **MMS: Open Preview to the Side** (the button in the editor title, or `Ctrl+K V`) opens it on the declaration at the cursor. Vanilla worldgen the pack refers to is fetched from GitHub (mcmeta).
- **Build into a world:** the MMS icon in the activity bar opens a side bar. In **World**, choose a Minecraft world (worlds from the Minecraft launcher's, Prism Launcher's and MultiMC's saves folders are listed, most recently played first, or browse to any world folder) and click **Build into world**. It saves the project's files and runs `mms build --world`: the data pack goes to `<world>/datapacks/<namespace>` and the resource pack, if there is one, is merged into `<world>/resourcepacks/resources.zip`, keeping the files other projects built into the same world. The result stays in the view under the button, with what to do in the game next, worked out from what changed: run `/reload` when only functions, loot tables, recipes, advancements and similar changed (with a **Copy /reload** button), or exit and rejoin the world when worldgen, dimensions, other registry entries or the resource pack changed. There are no pop-up notifications while the view is open; the status bar says it briefly, and a failed build only pops up a notification when the side bar is closed. Each project remembers its world. Tick **Build on save** (the `mms.buildOnSave` setting) to build on every save, and `mms.savesFolders` to list worlds from other launchers.
- **Structures from the world:** under the world, **Structures** lists what structure blocks saved in it (`<world>/generated/<namespace>/structure`), each marked new, imported or changed in world, with buttons to import it (or **Import N** for all the new and changed ones), copy its `Structure(...)` call, such as `Structure("structures/tower/base.nbt", "tower/base")`, and show the project copy. Imports go to the project's `structures` folder (`mms.structuresFolder`). A structure that is imported but that no `Structure(...)` call uses (the `unused-asset` lint) is marked **Not used** in yellow, with the call to add to a template pool; unused `.nbt` files in the project that this world did not save are listed too. After you save a piece again in the game it shows as changed; import it again to update it. This makes building jigsaw pieces in game and wiring them up in MMS a loop.
- **Pack explorer:** **Pack** in the same side bar lists the files the project generates, as the web editor's Pack tab does: the data pack, the resource pack and the custom items. Click a file to read it (it follows your edits); the buttons beside a file open its preview or go to the declaration that wrote it.
- **`mms.json` support:** completion, hover docs and validation for the project file, including every lint rule id and severity. This works without the language server.

## Setup

The extension runs the `mms` language server. Download `mms` for your platform from the [releases page](https://github.com/MinecraftMetascript/mms/releases) and put it on your `PATH`, or set `mms.path` to its location.

To use a local build, run `go install ./cmd/mms` in `engine/` (the extension also looks in `~/go/bin`, which VS Code launched from the macOS Dock may not have on its `PATH`), or run `go build ./cmd/mms` and set `mms.path` to the full path of the binary, such as `/Users/you/mms/engine/mms`. `~` and `${workspaceFolder}` work in `mms.path`.

The extension starts when a folder contains an `mms.json` or when you open a `.mms` file. If the server cannot start, a notification says why; fix it and run **MMS: Restart Language Server** (also use it after updating `mms`). **MMS: Show Language Server Output** shows the server log.

## Building

The preview panel uses the same Svelte components as the web editor, from `packages/preview`. `npm run build` bundles the extension (esbuild), the panel (Vite, into `out/webview`) and the World view's page (`webview/World.svelte`, into `out/world`).

## Credits

Previews are rendered with [Deepslate](https://github.com/misode/deepslate) by [Misode](https://github.com/misode), and read vanilla worldgen from Misode's [mcmeta](https://github.com/misode/mcmeta). The license notices of everything the extension bundles are in `out/third-party-licenses.txt`, written by `npm run build`.

## License

Same as the MMS repository.
