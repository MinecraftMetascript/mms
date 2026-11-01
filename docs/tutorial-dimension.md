# Tutorial: a custom dimension

This tutorial builds a dimension of floating islands, step by step. The finished project is [`examples/sky-islands`](../examples/sky-islands/), split into `dimension.mms`, `terrain.mms` and `visit.mms`. In the web editor, open it from **Examples**. The tests build it on every change, so it stays in step with the language.

A dimension is made of several resources that refer to each other by id:

| Resource | Kind | Says |
|---|---|---|
| `dimension_type` | `DimensionType` | how the world behaves: height, light, sky |
| `worldgen/biome` | `Biome` | climate, colours, mobs and features |
| `worldgen/noise` and `worldgen/density_function` | `Noise`, `DensityFunction` | the terrain's shape |
| `worldgen/material_rule` | `SurfaceRule` | which blocks cover the terrain |
| `worldgen/noise_settings` | `NoiseSettings` | the generator: terrain, surface, sea level |
| `dimension` | `Dimension` | the dimension itself, tying the rest together |

In MMS they can share one file or be split across several, since every file in the project sees every declaration. A name such as `islands` refers to the declaration of that kind in your namespace, and every reference is checked. Completion inside each `{ }` lists the fields the game accepts.

## 1. Set up the project

```sh
mkdir sky-islands && cd sky-islands
mms init -namespace sky
```

This writes `mms.json` and a starter `main.mms`. You can put everything below in that file, or split it up the way the example does: the dimension type, biome and dimension in `dimension.mms`, steps 4 to 6 in `terrain.mms`, and the function in `visit.mms`. The order of files and declarations does not matter.

## 2. The dimension type

```
DimensionType islands = {
  height = 256
  logical_height = 256
  min_y = 0
  coordinate_scale = 1
  ambient_light = 0.1
  has_skylight = true
  has_ceiling = false
  has_ender_dragon_fight = false
  infiniburn = "#minecraft:infiniburn_overworld"
  default_clock = overworld
  monster_spawn_block_light_limit = 0
  monster_spawn_light_level = 0
  attributes = {
    "minecraft:visual/sky_color" = "#9ad0ff"
    "minecraft:visual/fog_color" = "#d8ecff"
    "minecraft:visual/cloud_height" = 40
  }
}
```

The world runs from y = 0 to y = 255. Low clouds at y = 40 drift below the islands. `attributes` holds environment attributes, the 26.3 home of sky, fog and music settings. Hover a field to read what it does.

## 3. A biome

```
Biome meadow = {
  temperature = 0.7
  downfall = 0.6
  has_precipitation = true
  effects = { water_color = "#3f76e4" }
  carvers = []
  features = []
}
```

With no carvers and no features, the islands are bare stone, grass and dirt. To add trees or ores later, list placed features in `features`, one list per generation step.

## 4. The terrain

Terrain comes from a density function. Where it is above 0 the generator places `default_block`, and where it is below 0 it leaves air (or `default_fluid` below the sea level).

```
Noise island_noise = {
  base_octave = -6
  amplitude_modifiers = [1, 1, 0.5]
}
```

Each entry in `amplitude_modifiers` is one octave. Minecraft also needs `octave_count`, which must match the length of the list. MMS fills it in for you when you leave it out, and reports an error if you set it to a different number.

```
// 0 at y = 128, rising to 1 at y = 32 and y = 224.
DensityFunction height = Gradient(
  axis = y,
  from_coordinate = 32,
  to_coordinate = 224,
  from_value = -1,
  to_value = 1,
)

DensityFunction islands = Noise(island_noise, 1, 2) - height * height - 0.3
```

- `island_noise` is smooth 3D noise. A lower `base_octave` gives bigger shapes.
- `height` is a gradient from -1 to 1 over the height of the world. Squared, it is 0 in the middle and 1 at the ends.
- `islands` subtracts that from the noise. Solid blocks appear only where the noise is high and the height is near y = 128. Raising `0.3` makes the islands smaller and further apart.

The infix arithmetic compiles to `sub` and `mul` density functions. See [Writing resources](resources.md#density-functions).

## 5. The surface

```
SurfaceRule surface = [
  If(StoneDepth(Floor)) { Block(grass_block) },
  If(StoneDepth(Floor).Offset(3)) { Block(dirt) },
]
```

The rules are tried in order. The top block of each floor becomes grass, and the three blocks below it become dirt. Everything else stays stone.

## 6. The generator settings

```
NoiseSettings islands = {
  sea_level = 0
  default_block = stone
  default_fluid = air
  disable_mob_generation = false
  legacy_random_source = false
  noise = {
    min_y = 0
    height = 256
  }
  material_rule = surface
  spawn_target = []
  noise_router = {
    final_density = islands
    temperature = 0
    vegetation = 0
    continents = 0
    erosion = 0
    depth = 0
    ridges = 0
    chunk_surface_level = 0
  }
}
```

`noise` must match the dimension type's `min_y` and `height`. The router's other entries drive biome placement and aquifers. With a single biome they can all be `0`.

## 7. The dimension

```
Dimension islands = {
  type = islands
  generator = {
    type = noise
    settings = islands
    biome_source = {
      type = fixed
      biome = meadow
    }
  }
}
```

The names on the right are resolved by the field they are in: `type` names a `DimensionType`, `settings` names a `NoiseSettings`, and `biome` names a `Biome`. They all become `sky:islands` and `sky:meadow` in the JSON. If you mistype one, the editor reports it at once.

## 8. A way in

```
Function visit {
  /execute in sky:islands run tp @s 0 160 0
  /tellraw @s "Welcome to the sky islands"
}
```

The command checker knows your project's dimensions, so `in sky:islandz` would be an error.

## 9. Build and play

```sh
mms build --world ~/.minecraft/saves/MyWorld
```

Custom dimensions load when a world starts, so restart the world, not just `/reload`, then run `/function sky:visit`.

## Next steps

- Use a `multi_noise` biome source and the router's `temperature` and `vegetation` entries for several biomes.
- Add features to the biome's `features` lists.
- Turn on aquifers by adding an `aquifers` block to the settings.
- Import a vanilla dimension with `mms import` to see a complete one written in MMS.
