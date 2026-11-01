// Ids the registry has but the game rejects at that spot
// (#[id(exclude=[...])] in vanilla-mcdoc).

Biome hell = {
  temperature = 2
  downfall = 0
  has_precipitation = false
  effects = { water_color = "#905957" }
  attributes = {
    "minecraft:visual/fog_color" = "#330808"
    "minecraft:gameplay/fast_lava" = true
    "gameplay/sky_light_level" = 4
  }
  carvers = []
  features = []
}

DimensionType deep = {
  height = 256
  logical_height = 256
  min_y = -64
  coordinate_scale = 1
  ambient_light = 0
  has_skylight = false
  has_ceiling = true
  has_ender_dragon_fight = false
  has_fixed_time = false
  infiniburn = "#minecraft:infiniburn_overworld"
  monster_spawn_block_light_limit = 0
  monster_spawn_light_level = 0
  attributes = {
    "minecraft:gameplay/fast_lava" = true
  }
}

Recipe nothing = Shapeless([stick], air)

LootTable empty = {
  pools = [Pool(1, [Item(air)])]
}

// Block state variants may not use the builtin models.
BlockState minecraft:stone = {
  variants = { "" = { model = "builtin/entity" } }
}

// Heights come in 16-block sections (#[divisible_by=16]).
DimensionType short = {
  height = 100
  logical_height = 100
  min_y = -8
  coordinate_scale = 1
  ambient_light = 0
  has_skylight = false
  has_ceiling = true
  has_ender_dragon_fight = false
  has_fixed_time = false
  infiniburn = "#minecraft:infiniburn_overworld"
  monster_spawn_block_light_limit = 0
  monster_spawn_light_level = 0
}
