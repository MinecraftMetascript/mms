// The custom-dimension tutorial (docs/tutorial-dimension.md) builds this
// project. The dimension, its type and its biome are here; the terrain is
// in terrain.mms and the way in is in visit.mms. Names resolve across
// files, so `settings = islands` below finds the NoiseSettings there.

// 1. How the world behaves: height, light, sky.
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

// 2. The one biome.
Biome meadow = {
  temperature = 0.7
  downfall = 0.6
  has_precipitation = true
  effects = { water_color = "#3f76e4" }
  carvers = []
  features = []
}

// 5. The dimension: ties the type, the biomes and the terrain together.
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
