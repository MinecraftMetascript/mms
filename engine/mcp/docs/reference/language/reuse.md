# Constants, templates and loops

Four tools cut repetition: `const` names a value, `Template` names a value with parameters, comprehensions build lists, and a top-level `for` repeats declarations. They all run at compile time and write nothing of their own: only the declarations that use them reach the pack. `private` keeps a name to the file it is written in.

## Constants

`const name = value` names a value. Any value can be a constant: a number, a list, a loot pool, a predicate, a selector, a text. A constant never becomes a file. It is checked against its type, written after the name or taken from the first place it is used, and copied in wherever it is named, so shorthand and completion work inside it as they do in place.

```mms
/** Chunks between dungeon starts. */
const dungeon_spacing: int = 32
const woods = [oak, spruce, birch, jungle, acacia, dark_oak, mangrove, cherry]

ItemTag gems = [diamond, emerald, amethyst_shard]

const gem_pool: LootPool = Pool(1..3, #gems)
const stormy: Predicate = WeatherCheck(thundering = true) && TimeCheck(clock = overworld, value = 13000..)

LootTable chests/gems { pools = [gem_pool] }
```

Constants work at file level, in templates, in Blueprints and in function bodies. In a function, a constant is folded into each command that uses it, so `const boss = @e[tag=siege.boss, limit=1]` writes the selector itself. Hover on a constant's name shows its value.

## Templates

A `Template` is a helper you write. A call looks like a built-in helper's, and completion, signature help and hover treat it the same way. The return type is the type its body is checked against, so shorthand applies inside it:

```mms
/** A pool of one item, rolled a number of times. */
Template single(item: Id<item>, rolls: int = 1): LootPool = Pool(rolls, item)

LootTable chests/apples {
  pools = [single(apple, 3), single(golden_apple)]
}
```

Parameters have a type, and an optional default. Arguments are positional or named, as with helpers: `single(rolls = 2, item = bread)`.

A Template that returns `Commands` writes function code where it is called, and one that returns `Text` builds text. See [Functions](functions.md#command-templates) and [Text](text.md#named-text). The prelude's own templates, such as `Piece` for jigsaw pools, are listed in the [Helpers](https://minecraftmetascript.github.io/mms/docs/reference/helpers) reference.

## Comprehensions and spread

`[for x in xs: f(x)]` builds a list from another one, and `[for x in xs if cond: f(x)]` keeps only some items. `...list` inside a list puts that list's items in its place:

```mms
const woods = [oak, spruce, birch, jungle, acacia, dark_oak, mangrove, cherry]

ItemTag planks_and_logs = [
  ...[for w in woods: "${w}_planks"],
  ...[for w in woods if w != mangrove: "${w}_log"],
]
```

A string with `${...}` in it, such as `"${w}_planks"`, makes an id where an id is expected.

## Repeating declarations

A `for` at file level repeats the declarations in its body. Each name is a template over the loop variable, so every copy has an ordinary id of its own:

```mms
const woods = [oak, spruce, birch, jungle, acacia, dark_oak, mangrove, cherry]

for w in woods {
  /** A crate made of ${w}. */
  Recipe ${w}_crate = Shaped(["PPP", "P P", "PPP"], { P = "${w}_planks" }, barrel)

  LootTable blocks/${w}_crate {
    pools = [Pool(1, barrel)]
  }
}
```

That declares `oak_crate`, `spruce_crate` and so on. Go to definition on `oak_crate` lands on the loop, with `w = oak` shown inline, and find references and rename work as on any other name. Nothing is addressed by index, as `crate["oak"]` would be.

The list must be known at compile time, so the editor always knows every id the loop makes. `for k, v in map` walks a map's names and values. Loops can be nested.

## Private names

Every file in a project sees every other file's constants, templates, Types and Blueprints. Write `private` before one to keep it to its own file:

```mms
private const gap: int = 4
private Template spaced(n: int): int = n * gap

LootTable chests/spaced { pools = [Pool(spaced(2), stone)] }
```

Another file can't name `gap` or `spaced`, and it may keep a private `gap` of its own: the two never clash. A template or Blueprint still sees its own file's private names when another file uses it, so a public `Template` can be built from private helpers.

A private name that hides a public one declared in another file is fine. The file means its own, and a hint (`mms/shadowed-private`) points it out. The same name declared both public and private in one file is a duplicate.

Only names the compiler expands away can be private: `const`, `Template`, `Type` and `Blueprint`. A declaration that writes a file into the pack, such as a `LootTable` or a `Function`, lands in the namespace every file shares, so `private` on one is an error.
