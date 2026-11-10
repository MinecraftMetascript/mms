# Lint rules

The engine lints every declaration that checks without errors. Findings appear in the editor, in `mms check` (hints only with `-hints`), and as LSP diagnostics with quick fixes. `mms fix` and the editor's "Fix all" (Ctrl/⌘+Alt+F) apply every *safe* fix. In the web editor, Ctrl/⌘+. lists the quick fixes at the cursor, and Format is Shift+Alt+F (⌘+Shift+I on macOS) or Ctrl/⌘+S. A safe fix never changes the generated pack.

Every rule, with its default severity and what it reports, is listed in the [lint rule reference](https://minecraftmetascript.github.io/mms/docs/reference/lint), which is generated from the engine. The ones you will meet most:

- `unused` (hint): a declaration nothing references, by id or in a command. The editor fades the name. Declarations in `minecraft` are skipped, because they override vanilla files.
- `quoted-name`, `redundant-namespace`, `use-shorthand` and `constant-expression` (hint, safe fixes): a shorter way to write the same thing, such as `stone` for `"minecraft:stone"` or `Noise(noise = ridge)` for `{ type = noise, noise = ridge }`.
- `deprecated` (warning): a use of a declaration whose [doc comment](../start/getting-started.md#doc-comments) has a `@deprecated` tag.
- `external-asset` (warning) and `unused-asset` (info): an asset id that no `Texture`, `Sound` or `StructureFile` call provides, and an asset file nothing uses. See [Asset files](resources.md#asset-files).
- `texture-outside-atlas` (warning): a texture a model names that neither the blocks atlas (`textures/block/`) nor the items atlas (`textures/item/`) stitches, so the game shows the missing texture. See [Resource packs](resources.md#resource-packs).
- `book-page-overflow` (warning): a [Book](resources.md#books) page with more text than the book screen shows.
- The `jigsaw-` rules (warning or info): jigsaw blocks in structure files that can never join, name an unknown pool or hole, or place pieces that stick out or overlap.
- `structure-unknown-loot-table` (warning): a chest, barrel, chest minecart or other container in a structure file whose `LootTable` tag names a table that is not vanilla and that no `LootTable` declares, so the game leaves it empty. The quick fix declares the table.

Blueprints add rules of their own: the prelude's `custom-item-as-base` hint on a [custom item](resources.md#custom-items) written where a plain item id is expected, and any rule a project's Blueprint names in `warn(field, "rule", "message")` or `hint("rule", "message")`. They are configured and ignored like the built-in rules. A setting for a rule that neither MMS nor a Blueprint has is a warning on `mms.json`, with the closest rule name.

## Configuration

Set severities in `mms.json`. Each value is `error`, `warning`, `info`, `hint` or `off`:

```json
{
  "namespace": "my_pack",
  "lint": { "unused": "warning", "quoted-name": "off" }
}
```

## Suppression

- `// mms-ignore <rule>[, <rule>]` silences findings on its own line and on the next line.
- `// mms-ignore` with no rule silences every rule.
- `// mms-ignore-file <rule>` applies to the whole file.

Every lint finding offers a quick fix that inserts the comment for you.

## Quick fixes for errors

Some checker errors have a quick fix too. A "missing field" error on an object or a helper call offers **Add missing fields**, which writes every required field with a placeholder value:

```mms error
DensityFunction d = Clamp(1)            // missing fields 'min', 'max'
DensityFunction d = Clamp(1, min = 0, max = 0)
```

In an object that spans lines, each field goes on its own line before the closing brace. A placeholder is `0` (or the nearest number the field's range allows), `false`, the first enum value, `[]`, `""` or `{}`. For a value that can take several forms, the simplest form is used, so a loot pool's `rolls` becomes `0`. An inserted `{}` reports its own missing fields, with its own fix. Placeholders are guesses, so these fixes are not safe and "Fix all" leaves them alone.
