# Writing resources

Every JSON resource kind can be written as an object literal that mirrors the JSON. Two kinds of shorthand sit on top of that.

**Helpers from the schema.** Every dispatched type has a helper named after its case:

- `Noise(noise = ridge, xz_scale = 1, y_scale = 1)` is `{ type = noise, ... }`.
- `SurvivesExplosion()` is `{ type = survives_explosion }`.
- Positional arguments fill fields in schema order. A list field takes the remaining arguments.

**Sugar.** Shorthand for common shapes, listed below. Each one lowers to exactly the JSON shown, and the schema still checks every field.

The `use-sugar` lint rule suggests the short forms, and imported packs use them automatically.

## Loot tables

```
LootTable chests/gems {
  pools = [
    Pool(1..3, diamond, emerald, #minecraft:logs),
    Pool(1, [Item(name = gold_ingot).Weight(3)]),
  ]
}
```

| Write | Means |
| --- | --- |
| `Pool(rolls, entries...)` | `{ rolls, entries = [...] }` |
| `diamond` in `entries` | `{ type = item, name = diamond }` |
| `#minecraft:logs` in `entries` | `{ type = tag, items = "#minecraft:logs", expand = true }`. The tag is expanded, so each item is equally likely. |
| `1..3` as a number provider (`rolls`, `count`, ...) | `{ type = uniform, min = 1, max = 3 }` |

## Conditions and predicates

```
Predicate stormy_night = WeatherCheck(thundering = true) && !(TimeCheck(clock = overworld, value = 0..12000) || RandomChance(0.5))
```

| Write | Means |
| --- | --- |
| `a && b && c` | `{ type = all_of, terms = [a, b, c] }` |
| `a \|\| b` | `{ type = any_of, terms = [a, b] }` |
| `!a` | `{ type = inverted, term = a }` |

A chain of the same operator becomes one list of terms. Parentheses keep a group nested: `(a && b) && c` is `all_of(all_of(a, b), c)`. The same operators work wherever a loot condition is expected, such as an entry's or pool's `condition`.

## Bounds

| Write | Means |
| --- | --- |
| `10..` | `{ min = 10 }` |
| `..2` | `{ max = 2 }` |
| `1..5` | `{ min = 1, max = 5 }` |

These work for int and float ranges and for min/max bounds (scores, counts, distances).

## Recipes and item stacks

```
Recipe ruby_block = Shaped(["RRR", "RRR", "RRR"], { R = emerald }, emerald_block)
Recipe rubies = Shapeless([emerald_block], emerald * 9)
```

| Write | Means |
| --- | --- |
| `Shaped(pattern, key, result)` | `{ type = crafting_shaped, pattern, key, result }` |
| `Shapeless(ingredients, result)` | `{ type = crafting_shapeless, ingredients, result }` |
| `emerald * 9` as an item stack | `{ id = emerald, count = 9 }` |

## Advancements

```
Advancement first_gem = {
  criteria = { has_gem = InventoryChanged(items = [{ items = [diamond, emerald] }]) }
}
```

A trigger helper's named arguments are the criterion's conditions. `InventoryChanged(items = ...)` is `{ trigger = inventory_changed, conditions = { items = ... } }`.

## Tags

```
ItemTag gems = [diamond, #minecraft:logs, Optional(emerald)]
```

| Write | Means |
| --- | --- |
| `#minecraft:logs` | `"#minecraft:logs"`: another tag's entries |
| `Optional(emerald)` | `{ id = emerald, required = false }`: the tag still loads if the entry is missing |

## Structures

A structure that generates in new chunks takes four resources, each naming the next:

| Kind | Says |
| --- | --- |
| `StructureSet` | how far apart structures start, with `RandomSpread(spacing, separation, salt)` |
| `Structure` | the biomes, the height, and for `Jigsaw(...)` the pool to start from |
| `TemplatePool` | the templates (`.nbt` files) to place, with their processors |
| `ProcessorList` | rules that change blocks as a template is placed |

A template's chests can get a loot table from a `Rule` processor with `block_entity_modifier = AppendLoot(chests/buried)`. That works on vanilla templates too, so a structure can reuse one without its own `.nbt` file. The [`treasure`](../examples/treasure/) example does this in `worldgen.mms`.

## Density functions

`+`, `-`, `*` and `/` build `add`, `sub`, `mul` and `div`. Arithmetic on constants is folded. A unary `-x` is `x * -1`.
