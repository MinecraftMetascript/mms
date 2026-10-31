// The custom-dimension tutorial (docs/tutorial-dimension.md) builds this
// project step by step.

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

// 3. The terrain: islands where 3D noise is high, fading out above and
// below y = 128.
Noise island_noise = {
  base_octave = -6
  amplitude_modifiers = [1, 1, 0.5]
}

// 0 at y = 128, rising to 1 at y = 32 and y = 224.
DensityFunction height = Gradient(
  axis = y,
  from_coordinate = 32,
  to_coordinate = 224,
  from_value = -1,
  to_value = 1,
)

DensityFunction islands = Noise(island_noise, 1, 2) - height * height - 0.3

// 4. What the surface is made of.
SurfaceRule surface = [
  If(StoneDepth(Floor)) { Block(grass_block) },
  If(StoneDepth(Floor).Offset(3)) { Block(dirt) },
]

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

// 6. A way in: /function sky:visit
Function visit {
  /execute in sky:islands run tp @s 0 160 0
  /tellraw @s "Welcome to the sky islands"
}
