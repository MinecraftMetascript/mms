# Minecraft Metascript for Zed

Language support for [Minecraft Metascript](https://github.com/MinecraftMetascript/mms) (`.mms`) files in [Zed](https://zed.dev).

- **Highlighting, bracket matching, indentation, outline and text objects** come from the tree-sitter grammar in [`editors/tree-sitter-mms`](../tree-sitter-mms).
- **Everything else** comes from the MMS language server (`mms lsp`), the same one the VS Code extension and the web editor use:
  - diagnostics for errors, lint findings and raw commands
  - completion, hover, signature help
  - go to definition, find references, rename
  - quick fixes, and "Fix all" (`source.fixAll.mms`)
  - formatting

## Install

Until the extension is in Zed's registry, install it from a checkout:

1. Install Rust with [rustup](https://rustup.rs) (Zed needs it to build dev extensions).
2. In Zed, run **zed: install dev extension** from the command palette and pick the `editors/zed` folder.

Open a folder with an `mms.json` and any `.mms` file to start the server.

## Finding `mms`

The extension uses the first of these that exists:

1. `lsp.mms.binary.path` in your Zed settings.
2. `mms` on your `PATH`.
3. The latest [GitHub release](https://github.com/MinecraftMetascript/mms/releases) for your platform, downloaded into the extension's folder. It checks for a newer release each time the server starts, and keeps the last download for when you are offline.

To pin a build, or pass extra flags:

```jsonc
// settings.json
{
  "lsp": {
    "mms": {
      "binary": {
        "path": "/path/to/mms",
        "arguments": ["lsp"]
      }
    }
  }
}
```

Use **editor: restart language server** after updating `mms`.

Format on save follows Zed's normal `format_on_save` setting. The formatter has no options, so it is safe to turn on for the language:

```jsonc
{
  "languages": {
    "Minecraft Metascript": { "format_on_save": "on" }
  }
}
```

## Developing

```sh
cd editors/zed
rustup target add wasm32-wasip2
cargo build --release --target wasm32-wasip2   # what Zed runs when it builds the extension
```

After installing it as a dev extension, the **Rebuild** button on its entry in the Extensions page picks up changes. `zed --foreground` prints the extension's log, including server start-up errors.

### Changing the grammar

Zed does not read the grammar from your checkout. It clones the repository at the `rev` in `extension.toml` and builds `editors/tree-sitter-mms` from there. So, after changing the grammar:

1. Regenerate and test it (`cd editors/tree-sitter-mms && npx tree-sitter generate && npm test`).
2. Commit it.
3. Set `rev` in `extension.toml` to that commit, and commit again.

CI fails while `rev` points at a commit whose grammar differs from the current one. To try a grammar change locally before committing, point the grammar at your checkout for a moment: `repository = "file:///path/to/mms"` (it still needs a commit for `rev`).

The queries live in `languages/mms/`. `highlights.scm` is a copy of `editors/tree-sitter-mms/queries/highlights.scm`; edit both (CI checks that they match). As in Zed's own languages, a later pattern in `highlights.scm` wins over an earlier one.

## Publishing to the Zed registry

This needs a maintainer, and the repository needs a license first: Zed's registry only accepts extensions with an accepted open source license (MIT, Apache 2.0, BSD, GPL and similar) in the extension folder or the repository root. The repository has none yet.

1. Add a `LICENSE` file to the repository root (or `editors/zed/`), and set `license` in `editors/tree-sitter-mms/package.json` and `tree-sitter.json` to match.
2. Make sure `rev` in `extension.toml` is a commit on `main`, and `version` is the version to publish.
3. Fork [zed-industries/extensions](https://github.com/zed-industries/extensions) and add this repository as a submodule:

   ```sh
   git submodule add https://github.com/MinecraftMetascript/mms.git extensions/minecraft-metascript
   ```

4. Add an entry to its `extensions.toml`, pointing at the extension's folder inside the submodule:

   ```toml
   [minecraft-metascript]
   submodule = "extensions/minecraft-metascript"
   path = "editors/zed"
   version = "1.0.0"
   ```

5. Run `pnpm sort-extensions` in the fork, then open a pull request against zed-industries/extensions.

For each later release: bump `version` in `extension.toml` (and `rev` if the grammar changed), merge to `main`, then open a PR on zed-industries/extensions that moves the submodule to the new commit and updates `version` in `extensions.toml`.
