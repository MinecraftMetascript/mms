# Declarations

A file is a list of declarations. Each one starts with its kind, a PascalCase word, then its name:

```mms
ItemTag gems = [diamond, emerald, amethyst_shard]

LootTable chests/gems {
  pools = [Pool(1..3, #gems)]
}
```

Write `Kind name { fields }` when the value is an object, and `Kind name = value` for anything else: a list, a helper call, a number. There is one spelling for each.

The [Declaration kinds](https://minecraftmetascript.github.io/mms/docs/reference/kinds) page lists every kind. Lowercase words such as `const`, `let` and `for` are not kinds: they declare values the compiler works with, and never write a file. See [Constants, templates and loops](reuse.md).

## Names and ids

A declaration's name is its id without the namespace. Names can have folders: `chests/gems` in the namespace `gems` is `gems:chests/gems`, and a `LootTable` of that name is written to `data/gems/loot_table/chests/gems.json`.

Where a field takes an id, write the name bare. A bare name means the project's own declaration when there is one, and the vanilla one otherwise, so `diamond` is `minecraft:diamond` and `#gems` is your `ItemTag gems`. Write `minecraft:` only to reach vanilla when the project has a declaration of the same name.

Different kinds can share a name, such as a `Recipe` and a `LootTable` both called `gem_block`. Two declarations of one kind with the same id are an error, except tags, which merge.

## Namespaces

A file starts in the namespace from `mms.json`. A `Namespace` line changes it for the rest of that file:

```mms
Namespace minecraft

/** Adds to vanilla's swords tag. */
ItemTag swords = [diamond_sword]
```

## Several at once

Kinds whose declarations carry no value take a comma list, which declares each name:

```mms
JigsawHole room_exit, hall_entry, hall_exit
Objective wins, losses
```

## Doc comments

A `/** ... */` comment right above a declaration documents it. The text is Markdown, and hover shows it on the declaration and on every use, including ids inside commands. JSDoc tags such as `@param`, `@returns` and `@deprecated` are understood:

```mms
Objective found

/**
 * Gives a player their reward and counts it.
 *
 * @param count how many diamonds to give
 * @deprecated use `reward_all`
 */
Function reward(count: int) {
  /give @s diamond ${count}
  found[@s] += 1
}
```

A `@deprecated` tag makes every use a `deprecated` lint warning. [Getting started](../../start/getting-started.md#doc-comments) lists every tag.

## Names made from values

Inside a top-level `for` or a Blueprint's `emit`, a name can be a template, so one declaration writes many real ids:

```mms
const colors = [red, green, blue]

for c in colors {
  /** A ${c} banner from wool. */
  Recipe ${c}_banner_from_wool = Shaped(["WWW", "WWW", " S "], { W = "${c}_wool", S = stick }, "${c}_banner")
}
```

That declares `red_banner_from_wool`, `green_banner_from_wool` and `blue_banner_from_wool`. Each is an ordinary id: go to definition, find references and rename all work on it. See [Constants, templates and loops](reuse.md#repeating-declarations).
