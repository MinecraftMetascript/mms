# Biomes and features

A biome is a `Biome` declaration: its weather, its colours, its mobs and the features that decorate it. A dimension with more than one biome also needs a biome source that says where each one goes. This guide covers those pieces, in the form the game's 26.3 data packs use. For a whole dimension around them, follow the [custom dimension tutorial](../start/tutorial-dimension.md), and for the ground itself, [shaping terrain](terrain.md).

## A biome

```mms
Biome meadow {
  temperature = 0.7
  downfall = 0.6
  has_precipitation = true
  effects = {
    water_color = "#3f76e4"
    grass_color = "#79c05a"
  }
  attributes = {
    "minecraft:visual/sky_color" = "#78a7ff"
    "minecraft:visual/fog_color" = "#c0d8ff"
  }
  carvers = []
  features = []
}
```

| Field | What it holds |
| --- | --- |
| `temperature`, `downfall`, `has_precipitation` | the weather: snow or rain, and how grass and leaves are tinted when `effects` gives no colour |
| `effects` | colours of water, grass and leaves |
| `attributes` | environment attributes: sky, fog, music, ambient particles, mob spawns and more |
| `carvers` | the caves and canyons cut into the ground, such as `[cave, canyon]` |
| `features` | what decorates it: ores, trees, flowers, springs, one list per step |

Hover a field in the editor to read what it does, and completion lists the fields that are left.

## Sky, fog and other colours

On 26.3 a biome's colours are split between two fields.

- **`effects`** holds only the colours that tint blocks: `water_color` (required), `grass_color`, `foliage_color`, `dry_foliage_color` and `grass_color_modifier`.
- **`attributes`** holds everything about the environment, the same environment attributes a `DimensionType` has. Sky and fog are `"minecraft:visual/sky_color"` and `"minecraft:visual/fog_color"`, and underwater fog is `"minecraft:visual/water_fog_color"`.

`sky_color` or `fog_color` inside `effects` is an error, since that is where older versions kept them. An attribute is a string key, so it is quoted, and a misspelt one is reported with the nearest name. Completion inside `attributes` lists them all.

A biome's attribute replaces the dimension's for players inside the biome. Some attributes take a modifier that combines with the dimension's value instead, written as `{ modifier = ..., argument = ... }`, as the mob spawns below do.

## Mobs

There is no `spawners` or `spawn_costs` field on 26.3. Natural mob spawns are the attribute `"minecraft:gameplay/natural_mob_spawns"`:

```mms
Biome meadow {
  temperature = 0.7
  downfall = 0.6
  has_precipitation = true
  effects = { water_color = "#3f76e4" }
  attributes = {
    "minecraft:gameplay/natural_mob_spawns" = {
      modifier = overlay
      argument = {
        spawn_costs = {}
        spawns_by_category = {
          creature = [
            { type = sheep, count = 4, weight = 12 },
            { type = horse, count = Uniform(2, 6), weight = 5 },
          ]
          monster = [
            { type = zombie, count = 4, weight = 100 },
            { type = skeleton, count = 4, weight = 100 },
          ]
          ambient = [{ type = bat, count = 8, weight = 10 }]
        }
      }
    }
  }
  carvers = []
  features = []
}
```

- **`spawns_by_category`** has one list for each mob category: `monster`, `creature`, `ambient`, `axolotls`, `underground_water_creature`, `water_creature`, `water_ambient` and `misc`. A category left out spawns nothing.
- Each entry is a mob `type`, a `weight` against the others in its list, and a group `count`: a number, or a range such as `Uniform(2, 6)`.
- **`spawn_costs`** limits crowding, as in soul sand valleys: `spawn_costs = { ghast = { energy_budget = 0.7, charge = 0.15 } }`. Leave it `{}` for none.
- **`modifier`** says how the spawns combine with the dimension's. Vanilla biomes use `overlay`, and `override` is the other choice.

## Features

A feature is something placed into the world after the terrain: a tree, an ore vein, a patch of flowers. It takes two declarations:

- a **`Feature`** says what to place: a tree of oak logs and leaves, a vein of 17 coal ore.
- a **`PlacedFeature`** says where and how often: 30 tries per chunk, anywhere from y 136 up, only in the biomes that list it.

### Feature

A `Feature` is written either with a helper, after `=`:

```mms
/** An oak tree, as vanilla's `oak`. */
Feature small_oak = Tree(
  minimum_size = TwoLayersFeatureSize(),
  below_trunk_provider = soil_beneath_tree,
  trunk_provider = { id = oak_log },
  foliage_provider = { id = oak_leaves },
  trunk_placer = StraightTrunkPlacer(4, 2, 0),
  foliage_placer = BlobFoliagePlacer(radius = 2, offset = 0, height = 3),
  decorators = [],
  ignore_vines = true,
)
```

or as a block, with its `type` and the same fields:

```mms
/** The same tree, as a block. */
Feature small_oak {
  type = tree
  minimum_size = TwoLayersFeatureSize()
  below_trunk_provider = soil_beneath_tree
  trunk_provider = { id = oak_log }
  foliage_provider = { id = oak_leaves }
  trunk_placer = StraightTrunkPlacer(4, 2, 0)
  foliage_placer = BlobFoliagePlacer(radius = 2, offset = 0, height = 3)
  decorators = []
  ignore_vines = true
}
```

Both write the same `worldgen/feature/small_oak.json`. The helper is the shorter form, and its name is the feature's type: `Tree`, `Ore`, `RandomPatch`, `SimpleBlock`, `Lake` and so on. The [helpers reference](https://minecraftmetascript.github.io/mms/docs/reference/helpers) lists them all. `Feature x = { type = tree, ... }` is an error: a body is written without `=`.

Vanilla's features can be used by id with no declaration, such as `oak`, `fancy_oak` or `ore_coal`. Go to definition on one opens the game's JSON.

### PlacedFeature

```mms
/** An oak in about every other chunk, on dry ground. */
PlacedFeature meadow_oaks {
  feature = oak
  placement = [
    RarityFilter(2),
    InSquare(),
    SurfaceWaterDepthFilter(0),
    Heightmap(OCEAN_FLOOR),
    BlockPredicateFilter(WouldSurvive(state = { id = oak_sapling })),
    Biome(),
  ]
}
```

`feature` names a `Feature`, or holds one inline. Each placement step takes the positions from the one before: `Count` or `RarityFilter` sets how many tries, `InSquare` spreads them across the chunk, `Heightmap` or `HeightRange` picks the height, and the filters drop the ones that don't fit. End with `Biome()`, which keeps the feature to the biomes that list it.

### The biome's features list

A biome's `features` is a list of lists, one for each generation step, in this order:

| # | Step | # | Step |
| --- | --- | --- | --- |
| 1 | `raw_generation` | 7 | `underground_ores` |
| 2 | `lakes` | 8 | `underground_decoration` |
| 3 | `local_modifications` | 9 | `fluid_springs` |
| 4 | `underground_structures` | 10 | `vegetal_decoration` |
| 5 | `surface_structures` | 11 | `top_layer_modification` |
| 6 | `strongholds` | | |

The steps have no names in the file, only their place in the list, so steps with nothing in them are `[]`. Trailing steps can be left out, so a biome with ores and trees needs the first ten:

```mms
PlacedFeature meadow_oaks {
  feature = oak
  placement = [RarityFilter(2), InSquare(), Heightmap(OCEAN_FLOOR), Biome()]
}

Biome meadow {
  temperature = 0.7
  downfall = 0.6
  has_precipitation = true
  effects = { water_color = "#3f76e4" }
  carvers = [cave, canyon]
  features = [
    [],
    [],
    [],
    [],
    [],
    [],
    [ore_coal_upper, ore_coal_lower, ore_iron_middle],
    [],
    [spring_water],
    [meadow_oaks, flower_plains, patch_grass_plain],
  ]
}
```

Each entry is a placed feature: one of yours, a vanilla one by id such as `ore_coal_upper`, or one written inline as `{ feature = oak, placement = [...] }`. A step's list can also be a `PlacedFeatureTag` in place of the list, such as `#ores`, which is how several biomes share one list. A feature that several biomes list has to keep the same step and the same order against the others in every one of them: otherwise the game finds a feature order cycle and the world fails to load.

## Where biomes go

A dimension with one biome uses `Fixed(biome = meadow)` as its biome source. For more, a `MultiNoise` biome source lists each biome with the climate it belongs in. At each spot, the game samples the climate from the noise settings' router and picks the biome whose climate is nearest:

```mms
const meadow_and_desert: BiomeSource = MultiNoise(
  biomes = [
    {
      biome = plains
      parameters = {
        temperature = [-0.45, 0.2]
        humidity = [-1, 1]
        continentalness = [-0.11, 1]
        erosion = [-1, 1]
        weirdness = [-1, 1]
        depth = 0
        offset = 0
      }
    },
    {
      biome = desert
      parameters = ClimateParameters(
        temperature = [0.55, 1],
        humidity = [-1, 1],
        continentalness = [-0.11, 1],
        erosion = [-1, 1],
        weirdness = [-1, 1],
        depth = 0,
        offset = 0,
      )
    },
  ]
)
```

- Each entry is `{ biome = ..., parameters = ... }`. `parameters` is a `ClimateParameters`, written as a record or with the `ClimateParameters(...)` helper. The two are the same.
- All seven fields are required. Each of the first six is a number, for one point, or a `[min, max]` range, both between -2 and 2. The biome is nearest anywhere inside its range.
- Each one is matched against a function of the noise settings' `noise_router`: `temperature`, `vegetation` for `humidity`, `continents` for `continentalness`, `erosion`, `ridges` for `weirdness`, and `depth`. In vanilla's router only `depth` changes with height: it is 0 at the surface and grows underground, which is how cave biomes sit below the others.
- `offset` is a plain number from 0 to 1, added to the biome's distance everywhere, so a biome with a higher offset wins less often.
- `MultiNoise(preset = overworld)` uses vanilla's own table instead of `biomes`.

The biome source goes in the dimension's generator, as in `Noise(settings = ..., biome_source = MultiNoise(...))`. A `BiomeSource` is not a declaration of its own, since it has no file: name one with `const`, as here, or write it inline. The dimension's **Biomes** preview shows where each biome wins, and [why a biome doesn't appear](editors.md#why-a-biome-doesnt-appear) helps when one never does.
