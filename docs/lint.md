# Lint rules

The engine lints every declaration that checks without errors. Findings appear in the editor, in `mms check` (hints only with `-hints`), and as LSP diagnostics with quick fixes. `mms fix` and the editor's "Fix all" (Ctrl/⌘+Alt+F) apply every *safe* fix. A safe fix never changes the generated pack.

| Rule | Default | Fix | What it reports |
| --- | --- | --- | --- |
| `unused` | hint | none | A density function, noise, feature, carver, processor list, surface rule or condition, predicate, item modifier or parameter list that nothing references, by id or in a command. The game never loads these on its own. Declarations in `minecraft` are skipped, because they override vanilla files. The editor fades the name. |
| `quoted-name` | hint | safe | An id, dispatch key or enum name written in quotes: `"minecraft:stone"` becomes `stone`. |
| `redundant-namespace` | hint | safe | A `minecraft:` prefix that can be dropped. It is kept when the project declares its own resource with the short name. |
| `use-sugar` | hint | safe | An object literal with a shorter form: a dispatched object as its helper call (`{ type = noise, noise = ridge }` becomes `Noise(noise = ridge)`), and the sugar in [resources.md](resources.md): bare item entries, `1..3`, `diamond * 4`, trigger calls, `&&`, `||`, `!` and density arithmetic. Parentheses are added where precedence needs them. |
| `constant-expression` | hint | safe | Arithmetic on number literals, such as `2 * 0.25`, which is replaced by its value. |
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
