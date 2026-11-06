# Lint rules

The engine lints every declaration that checks without errors. Findings appear in the editor, in `mms check` (hints only with `-hints`), and as LSP diagnostics with quick fixes. `mms fix` and the editor's "Fix all" (Ctrl/⌘+Alt+F) apply every *safe* fix. In the web editor, Ctrl/⌘+. lists the quick fixes at the cursor, and Format is Shift+Alt+F (⌘+Shift+I on macOS) or Ctrl/⌘+S. A safe fix never changes the generated pack.

| Rule | Default | Fix | What it reports |
| --- | --- | --- | --- |
| `unused` | hint | none | A density function, noise, feature, carver, processor list, surface rule or condition, predicate, item modifier, parameter list or template pool that nothing references, by id or in a command. A template pool also counts as used when a jigsaw block in one of the project's structure files draws from it. The game never loads these on its own. Declarations in `minecraft` are skipped, because they override vanilla files. The editor fades the name. |
| `quoted-name` | hint | safe | An id, dispatch key or enum name written in quotes: `"minecraft:stone"` becomes `stone`. |
| `redundant-namespace` | hint | safe | A `minecraft:` prefix that can be dropped. It is kept when the project declares its own resource with the short name. |
| `use-shorthand` | hint | safe | An object literal with a shorter form: a dispatched object as its helper call (`{ type = noise, noise = ridge }` becomes `Noise(noise = ridge)`), and the shorthand forms in [resources.md](resources.md): bare item entries, `1..3`, `diamond * 4`, trigger calls, `&&`, `||`, `!` and density arithmetic. Parentheses are added where precedence needs them. |
| `custom-item-as-base` | hint | quick fix | A [custom item](resources.md#custom-items) in a field that takes a plain item id, such as a recipe ingredient or an item tag. It is written as its base item, so the field also matches the base item without the custom components. The quick fix writes the base item instead. |
| `custom-item-base` | hint | quick fix | The same for a Book, which stands for `written_book` there. |
| `accidental-markup` | hint | safe | A style mark in text outside a book that is never closed, such as the `*` in `"5 * 3"`. It is written as the character itself; the fix puts a backslash before it to say so. |
| `constant-expression` | hint | safe | Arithmetic on number literals, such as `2 * 0.25`, which is replaced by its value. |
| `deprecated` | warning | none | A use of a declaration whose [doc comment](getting-started.md#doc-comments) has a `@deprecated` tag, anywhere in the project, commands included. The message carries the tag's text, and the editor strikes the use out. |
| `external-asset` | warning | none | A texture, sound or structure id in a namespace the project declares that no `Texture`, `Sound` or `Structure` call in the project provides. The pack still loads, so this is a warning: the file has to come from another pack. See [Asset files](resources.md#asset-files). |
| `unused-asset` | info | none | A `.png`, `.ogg` or `.nbt` file in the project that no `Texture`, `Sound` or `Structure` call and no `icon` uses, so the packs leave it out. It has no source line, so it shows on the file itself: in VS Code's Problems panel, in `mms check` output, and as an "unused" badge in the web editor's asset list. Nothing is reported while the project has errors. |
| `jigsaw-hole` | warning | quick fix | A jigsaw block in a structure file whose name or target is in a namespace that declares [`JigsawHole`](resources.md#jigsaw-holes)s but is not one of them, most likely a typo made in the game's jigsaw screen. The message names the closest declared hole; the quick fix declares the name. Namespaces without any `JigsawHole` are not checked. |
| `book-overflow` | warning | none | A [Book](resources.md#books) page with more text than the book screen shows, 14 lines of 114 pixels. The game cuts the rest off without a warning. Widths come from the vanilla font, so text outside ASCII is a close guess. In a `.mmsbook` file the warning is in the file. |
| `overrides-vanilla` | info | none | A declaration that replaces a vanilla resource, such as `LootTable minecraft:blocks/stone`. A tag only counts when it sets `replace`, since otherwise it adds to vanilla's entries. It is skipped when the project namespace is `minecraft`. |

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

```mms
DensityFunction d = Clamp(1)            // missing fields 'min', 'max'
DensityFunction d = Clamp(1, min = 0, max = 0)
```

In an object that spans lines, each field goes on its own line before the closing brace. A placeholder is `0` (or the nearest number the field's range allows), `false`, the first enum value, `[]`, `""` or `{}`. For a value that can take several forms, the simplest form is used, so a loot pool's `rolls` becomes `0`. An inserted `{}` reports its own missing fields, with its own fix. Placeholders are guesses, so these fixes are not safe and "Fix all" leaves them alone.
