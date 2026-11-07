# Values and types

Values look like the JSON they become, with less punctuation. Every value is checked against the type its place expects, and that type also decides what shorthand means there.

## Literals

| Write | Means |
| --- | --- |
| `12`, `-3`, `0.5` | numbers |
| `true`, `false` | booleans |
| `"text"` | a string. In a text position it is [markup](text.md). |
| `diamond`, `minecraft:stone`, `chests/gems` | an id, or an enum word such as `rigid` |
| `#gems`, `#minecraft:logs` | a tag |
| `1..3`, `5..`, `..10` | a range: a number provider, or bounds |
| `[a, b, c]` | a list. Commas between items are optional. |
| `{ key = value }` | an object. Fields are separated by newlines or commas. |
| `none` | no value: a field set to `none` is left out |

```mms
LootTable chests/gems {
  pools = [
    {
      rolls = 1..3
      entries = [Item(diamond), Item(emerald)]
    },
  ]
}
```

## Helpers

A helper builds one typed object. `Pool(...)` is a loot pool, and `Noise(...)` is `{ type = noise, ... }`. Arguments are positional in the order the reference lists, or named:

```mms
LootTable chests/gems {
  pools = [Pool(rolls = 1..3, entries = [diamond, emerald])]
}
```

Every type the game picks by a `type` field has a helper named after it. The [Helpers](https://minecraftmetascript.github.io/mms/docs/reference/helpers) reference lists them by what they make. Your own [Templates](reuse.md#templates) are called the same way.

## Operators

| Operators | On |
| --- | --- |
| `+`, `-`, `*`, `/`, `%` | numbers. On density functions, `+`, `-`, `*` and `/` build `add`, `sub`, `mul` and `div`. |
| `==`, `!=`, `<`, `<=`, `>`, `>=` | numbers and other values |
| `&&`, `\|\|`, `!` | booleans, and loot and advancement predicates |
| `a ?? b` | `a`, or `b` when `a` is `none` |
| `if (c) a else b` | `a` when `c` holds, else `b` |
| `xs[i]`, `x.field` | an item of a list, a field of an object or record |

Arithmetic on numbers known at compile time is folded, so `Constant(16 * 4)` writes `64`.

```mms
const raining = WeatherCheck(raining = true)

Predicate dry_night = !raining && TimeCheck(clock = overworld, value = 13000..23000)

DensityFunction hills = Noise(minecraft:ridge, 1, 1) * 0.5 + 0.25
```

## Shorthand

The expected type decides what a short form means, so the same spelling works wherever it makes sense:

- A bare id in a loot pool's entries is an item entry, and `#tag` is every item of the tag.
- `diamond * 4` is a stack of four where an item stack goes.
- `1..3` is a uniform number provider, or bounds in a predicate.
- A trigger helper such as `InventoryChanged(...)` is a whole advancement criterion.

The `use-shorthand` lint hint points at an object that has a shorter form, with a fix. [Writing resources](../../guides/resources.md) shows each form.

## Types

Types come from three places: the game's own schema (`ItemStack`, `LootPool`, `HeightProvider`), the language, and `Type` declarations. The language's types are:

- Numbers and words: `int`, `float`, `double`, `long`, `short`, `byte`, `bool`, `string`, and `any` for a value of any type.
- `Text`: a text component, written as [text markup](text.md).
- `BookText`: a book's text, a list of markup lines or `TextFile("books/guide.mmsbook")`, where `---` starts a page.
- `Commands`: the statements a Template writes where it is called.
- `Function`: a function, by its name.
- `Selector`: a target selector such as `@e[tag=boss]`.
- `Texture`, `Sound` and `StructureFile`: an asset file in the project, read by the helper of the same name.

| Write | Means |
| --- | --- |
| `T?` | optional: the value can be `none` |
| `[T]` | a list |
| `{string: T}` | a map from names to values |
| `A \| B` | either. The first one that matches is used. |
| `Id<item>` | an id in a registry, vanilla or the project's |
| `Ids<worldgen/biome>` | an id, a tag, or a list of ids |
| `"rigid" \| "terrain_matching"` | one of these words |
| `LootTable`, any kind's name | an id of one, or a value of one written in place |

A `Type` declaration names a record or an alias:

```mms
/** One wave of a siege. */
Type Wave {
  mob: Id<entity_type>
  /** How many to spawn. */
  count: int = 1
}

/** Text in English, or one text per language. */
Type Label = string | {string: string}

const waves: [Wave] = [{ mob = zombie, count = 3 }, { mob = skeleton }]
```

A field can hold a record written in place, with its fields between braces as in a `Type` body: `spawn: { x: int, z: int = 0 }`. A lone field named `string`, as in `{string: string}`, is a map instead.

A record's fields get completion, hover and the missing-field fix, like the game's own types. Records are used by [Templates](reuse.md#templates), [Blueprints](blueprints.md) and [Storage](functions.md#values).
