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

## Inline values where an id is needed

Some fields need an id, and the game will not load an inline value there. Examples are a surface rule's `NoiseThreshold` noise, a biome source's biome, and a template pool's fallback. A density function's noise is in the same group. The schema allows it inline, but the game seeds each noise from its id, so an inline noise fails when the world loads.

You can still write the value inline in those fields. MMS checks it as its own declaration, writes it under `mms_gen/`, and puts its id in the field:

```mms
DensityFunction hills = Noise({ base_octave = -6, amplitude_modifiers = [1, 1] }, 1, 1)
```

This writes `worldgen/noise/mms_gen/hills/noise.json` and gives `hills` the noise `<ns>:mms_gen/hills/noise`. A second inline value of the same kind in the same declaration is named `noise_2`, and so on. Each inline noise gets its own id, so two identical inline noises still produce different noise. To share one noise, declare it with a name. Where the game does take inline values, such as density functions, features and surface rules, they stay inline.

## Asset files

Textures, sounds and structure templates are files, not declarations. Keep them anywhere in the project and name them where they are used:

```
PaintingVariant sunset = { width = 2, height = 1, asset_id = Texture("art/sunset.png") }
JukeboxSong tune = { sound_event = Sound("music/tune.ogg"), description = "Tune", length_in_seconds = 30, comparator_output = 3 }
TemplatePool houses = {
  fallback = minecraft:empty
  elements = [{ weight = 1, element = { element_type = "minecraft:single_pool_element", location = Structure("builds/house.nbt"), projection = "rigid", processors = "minecraft:empty" } }]
}
```

Each call stands for the id the file gets in your namespace: its base name, lowercased, or the second argument, as in `Texture("art/Sunset Final.png", "sunset")`. Only files that code uses are copied, to the path the game expects:

| Helper | File | Copied to |
| --- | --- | --- |
| `Texture` | `.png` | `assets/<ns>/textures/<prefix><name>.png`, where the field sets the prefix, such as `painting/` |
| `Sound` | `.ogg` | `assets/<ns>/sounds/<name>.ogg`, plus an entry in a generated `sounds.json` |
| `Structure` | `.nbt` | `data/<ns>/structure/<name>.nbt` |

Inside the quotes, completion lists the project's files of the helper's type, in VS Code, Zed and the web editor. A path to a file that is not there is an error.

Textures, sounds and [resource pack declarations](#resource-packs) go into a resource pack built next to the data pack: `build-resources/` (set `resourcePack.output` in `mms.json` to change it), `<name>-resources.zip` with `mms build --zip`, and `<world>/resources.zip` with `mms build --world`, which Minecraft loads as that world's resource pack. `"icon": "art/icon.png"` in `mms.json` becomes `pack.png` in both packs.

Where the game takes a sound event either by id or inline, as a jukebox song's `sound_event` does, `Sound(...)` writes it inline, `{ "sound_id": "my_pack:tune" }`: the game's sound event registry is fixed, so a pack's own sound cannot be named there by id.

Once a file is included, its id can also be written plainly elsewhere, with its namespace, such as `asset_id = my_pack:sunset`. An id in your namespace that nothing provides is reported by the [`external-asset`](lint.md) lint rule.

The web editor holds asset files too: **+ Add texture or sound** under the sources adds `.png`, `.ogg` and `.nbt` files, which keep their file name as their path. **Download resources** appears next to **Download pack** once there is a resource pack. Project zips carry the asset files both ways.

The [`embers`](../examples/embers/) example ships its own art: [custom items](#custom-items) with their own textures, a painting, a music disc with its own song, translations and a pack icon.

## Resource packs

Block states, models and the other JSON files of a resource pack are declarations too. They are written to `assets/<ns>/...` in the resource pack, and the schema checks them like any other resource.

```
BlockState minecraft:glowstone = {
  variants = {
    "" = [
      { model = block/glowstone_1, weight = 3 },
      { model = block/glowstone_2, y = 90 },
    ]
  }
}
Model block/glowstone_1 = { parent = block/cube_all, textures = { all = Texture("art/glow1.png", "block/glowstone_1") } }
Model block/glowstone_2 = { parent = block/cube_all, textures = { all = Texture("art/glow2.png", "block/glowstone_2") } }
```

| Kind | Writes | For |
| --- | --- | --- |
| `BlockState` | `blockstates/` | which model a block shows in each state: `variants`, weighted lists, `multipart` |
| `Model` | `models/` | block and item models, with `parent`, `textures` and `elements` |
| `ItemModel` | `items/` | the item model definition that the `item_model` component names |
| `Lang` | `lang/` | translations, such as `Lang en_us = { "item.my_pack.charm" = "Charm" }` |
| `Equipment` | `equipment/` | worn armour layers, named by the `equippable` component's `asset_id` |
| `Font`, `Atlas`, `Particle`, `PostEffect`, `WaypointStyle` | `font/`, `atlases/`, `particles/`, `post_effect/`, `waypoint_style/` | as in the game |

A block state is named after the block it replaces, such as `minecraft:glowstone`: the game only reads block states of its own blocks. Variant keys (`"facing=north,lit=true"`, or `""` for every state) and multipart `when` conditions are checked against the block's properties, and the editor completes property names in `when`.

In a model's `textures`, a value is a texture id, `"#name"` for another slot, or `Texture("file.png", "block/name")`, which copies the file to `textures/block/name.png`. A texture id can name a vanilla texture, such as `block/stone`.

`sounds.json` is still generated from `Sound(...)`, and textures cannot have `.png.mcmeta` animation yet. The [`resource-pack`](../examples/resource-pack/) example shows each kind with vanilla textures; [`embers`](../examples/embers/) uses files of its own.

## Custom items

Data packs cannot add items, but they can give a vanilla item a new look and name with its `item_model` and `item_name` components. That takes a model, an item model, a translation and the same components at every use. `CustomItem` writes all of it from one declaration:

```
CustomItem ember_shard = {
  base = amethyst_shard
  name = "Ember Shard"
  texture = Texture("art/ember_shard.png")
  components = { rarity = rare }
}
```

This writes `models/item/ember_shard.json` (a flat `item/generated` model of the texture), `items/ember_shard.json`, and `"item.embers.ember_shard": "Ember Shard"` in `lang/en_us.json`, merged with any `Lang en_us` declaration in the namespace. `Texture("...")` copies the file to `textures/item/ember_shard.png`.

| Field | |
| --- | --- |
| `base` | the vanilla item it is made from, which decides how it behaves |
| `name` | its name in English, or one per language: `{ en_us = "Ember Shard", de_de = "Glutsplitter" }`. The `item_name` component also carries the English text as a fallback, so the name shows without the resource pack |
| `texture` | the picture for a flat item model |
| `parent` | the generated model's parent: `item/generated` by default, `item/handheld` for tools |
| `model` | an existing model to show instead of one made from `texture` |
| `item_model` | an `ItemModel` id, or an item model body such as `{ model = Select(...) }`, instead of the generated one |
| `components` | more item components, given to every use |

Without `texture`, `model` or `item_model`, the item keeps its base item's look and only gets its name and components.

The item's name stands for the base item with its components:

```
Recipe ember_shard = Shapeless([blaze_powder, amethyst_shard], ember_shard * 2)
LootTable cache { pools = [Pool(1, [Item(name = ember_shard, modifier = [SetCount(2..5)])])] }
Function give {
  /give @s ${ember_shard} 16
}
```

- **Item stacks**, such as recipe results: `ember_shard`, `ember_shard * 2`, or `{ id = ember_shard, count = 2, components = { ... } }`, whose components are added to the item's own.
- **Loot entries**: `ember_shard` or `Item(name = ember_shard, ...)` drop the base item with a `set_components` modifier ahead of the others.
- **Commands**: `${ember_shard}` writes `amethyst_shard[item_model="embers:ember_shard",item_name={...}]`.

Anywhere else that takes an item id, such as a recipe ingredient or an item tag, a custom item is an error: the game matches those by item id only, so they need the base item. A `Model`, `ItemModel` or `Lang` entry that the custom item also writes is reported as a duplicate.
