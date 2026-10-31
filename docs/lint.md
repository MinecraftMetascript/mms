# Lint rules

The engine lints every declaration that checks without errors. Findings appear in the editor, in `mms check` (hints only with `-hints`), and as LSP diagnostics with quick fixes. `mms fix` and the editor's "Fix all" (Ctrl/⌘+Alt+F) apply every *safe* fix. A safe fix never changes the generated pack.

| Rule | Default | Fix | What it reports |
| --- | --- | --- | --- |
| `unused` | hint | none | A density function, noise, feature, carver, processor list, surface rule or condition, predicate, item modifier or parameter list that nothing references, by id or in a command. The game never loads these on its own. Declarations in `minecraft` are skipped, because they override vanilla files. The editor fades the name. |
| `quoted-name` | hint | safe | An id, dispatch key or enum name written in quotes: `"minecraft:stone"` becomes `stone`. |
| `redundant-namespace` | hint | safe | A `minecraft:` prefix that can be dropped. It is kept when the project declares its own resource with the short name. |
| `use-sugar` | hint | safe | An object literal with a shorter form: a dispatched object as its helper call (`{ type = noise, noise = ridge }` becomes `Noise(noise = ridge)`), and the sugar in [resources.md](resources.md): bare item entries, `1..3`, `diamond * 4`, trigger calls, `&&`, `||`, `!` and density arithmetic. Parentheses are added where precedence needs them. |
| `constant-expression` | hint | safe | Arithmetic on number literals, such as `2 * 0.25`, which is replaced by its value. |
| `external-asset` | warning | none | A texture, sound or structure id in a namespace the project declares that no `Texture`, `Sound` or `Structure` call in the project provides. The pack still loads, so this is a warning: the file has to come from another pack. See [Asset files](resources.md#asset-files). |
| `overrides-vanilla` | info | none | A declaration that replaces a vanilla resource, such as `LootTable minecraft:blocks/stone`. It is skipped when the project namespace is `minecraft`. |

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
