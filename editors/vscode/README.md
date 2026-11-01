# Minecraft Metascript for VS Code

Language support for [Minecraft Metascript](https://github.com/MinecraftMetascript/mms) (`.mms`) files.

- **Diagnostics** for errors, lint findings and every raw command, checked against the Minecraft command tree.
- **Editing help:**
  - completion, hover and signature help
  - go to definition, find references and rename
- **Fixes:**
  - quick fixes
  - "Fix all" (`source.fixAll.mms`)
- **Formatting** with the opinionated formatter (format on save is on by default for `.mms`).
- **Semantic highlighting.**
- **Previews:** click **Preview** above a noise, density function, noise settings, surface rule or dimension to see it rendered beside the editor (density, blocks, 3D, biome map), or above a loot table, recipe or biome to inspect it. The panel stays on that declaration while you edit elsewhere and redraws as you type; click another **Preview** to switch. **MMS: Open Preview to the Side** (the button in the editor title, or `Ctrl+K V`) opens it on the declaration at the cursor. Vanilla worldgen the pack refers to is fetched from GitHub (mcmeta).
- **`mms.json` support:** completion, hover docs and validation for the project file, including every lint rule id and severity. This works without the language server.

## Setup

The extension runs the `mms` language server. Download `mms` for your platform from the [releases page](https://github.com/MinecraftMetascript/mms/releases) and put it on your `PATH`, or set `mms.path` to its location.

To use a local build, run `go install ./cmd/mms` in `engine/` (the extension also looks in `~/go/bin`, which VS Code launched from the macOS Dock may not have on its `PATH`), or run `go build ./cmd/mms` and set `mms.path` to the full path of the binary, such as `/Users/you/mms/engine/mms`. `~` and `${workspaceFolder}` work in `mms.path`.

The extension starts when a folder contains an `mms.json` or when you open a `.mms` file. If the server cannot start, a notification says why; fix it and run **MMS: Restart Language Server** (also use it after updating `mms`). **MMS: Show Language Server Output** shows the server log.

## Building

The preview panel uses the same Svelte components as the web editor, from `packages/preview`. `npm run build` bundles the extension (esbuild) and the panel (Vite, into `out/webview`).

## License

Same as the MMS repository.
