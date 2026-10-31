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

## Setup

The extension runs the `mms` language server. Download `mms` for your platform from the [releases page](https://github.com/MinecraftMetascript/mms/releases), and put it on your `PATH`. Or set `mms.path` to its location.

Open a folder that contains an `mms.json` and the extension starts. Use **MMS: Restart Language Server** after updating `mms`.

## License

Same as the MMS repository.
