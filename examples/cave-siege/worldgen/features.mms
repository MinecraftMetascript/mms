// Features: light, plants, trees, lakes and the shared ore list.
//
// Vanilla's surface features find the ground with a heightmap, which in a
// world with a roof is the top of the roof. So every floor feature here
// uses CountOnEveryLayer instead, which finds each floor in the column,
// and every ceiling feature scans up from a random height to the roof.
//
// Light comes from blocks the player can see and mine, never from hidden
// light blocks: glowstone in the roof, shroomlight "glow caps" on the floor
// and froglight crystals. Saplings and animals need light 9 or more.

// ---- Shared placement pieces -----------------------------------------------

/** Ores, clay and gravel, as in the vanilla overworld. Every custom biome uses this. */
PlacedFeatureTag ores = [
  minecraft:ore_dirt,
  minecraft:ore_gravel,
  minecraft:ore_granite_upper,
  minecraft:ore_granite_lower,
  minecraft:ore_diorite_upper,
  minecraft:ore_diorite_lower,
  minecraft:ore_andesite_upper,
  minecraft:ore_andesite_lower,
  minecraft:ore_tuff,
  minecraft:ore_coal_upper,
  minecraft:ore_coal_lower,
  minecraft:ore_iron_upper,
  minecraft:ore_iron_middle,
  minecraft:ore_iron_small,
  minecraft:ore_gold,
  minecraft:ore_gold_lower,
  minecraft:ore_redstone,
  minecraft:ore_redstone_lower,
  minecraft:ore_diamond,
  minecraft:ore_diamond_medium,
  minecraft:ore_diamond_large,
  minecraft:ore_diamond_buried,
  minecraft:ore_lapis,
  minecraft:ore_lapis_buried,
  minecraft:ore_copper,
  minecraft:ore_clay,
]

/** Geodes and dungeons. */
PlacedFeatureTag underground_extras = [minecraft:amethyst_geode]

PlacedFeatureTag dungeons = [minecraft:monster_room, minecraft:monster_room_deep]

/** Water springs only: lava springs would pour lava across the tunnels waves walk. */
PlacedFeatureTag water_springs = [minecraft:spring_water]

// ---- Light -----------------------------------------------------------------

/** A hanging cluster of glowstone, about 7 blocks across. */
Feature glowstone_cluster = RandomNeighborSpread(
  block = { id = glowstone },
  accepted_neighbors = [glowstone, stone, deepslate, tuff, granite, diorite, andesite, dirt],
  can_replace = MatchingBlockTag(tag = #minecraft:air),
  attempts = 400,
  xz_offset = Trapezoid(min = -4, max = 4, plateau = 0),
  y_offset = Uniform(min_inclusive = -6, max_inclusive = 0),
)

/** A short mushroom stem with a shroomlight on top. */
Feature glow_cap = BlockColumn(
  direction = up,
  allowed_placement = MatchingBlockTag(tag = #minecraft:air),
  prioritize_tip = true,
  layers = [
    {
      height = Uniform(min_inclusive = 1, max_inclusive = 3)
      provider = { id = mushroom_stem }
    },
    {
      height = 1
      provider = { id = shroomlight }
    },
  ],
)

Feature ochre_crystals = RandomNeighborSpread(
  block = { id = ochre_froglight },
  accepted_neighbors = [ochre_froglight, stone, deepslate, tuff],
  can_replace = MatchingBlockTag(tag = #minecraft:air),
  attempts = 60,
  xz_offset = Trapezoid(min = -2, max = 2, plateau = 0),
  y_offset = Uniform(min_inclusive = -3, max_inclusive = 0),
)

Feature verdant_crystals = RandomNeighborSpread(
  block = { id = verdant_froglight },
  accepted_neighbors = [verdant_froglight, stone, deepslate, tuff],
  can_replace = MatchingBlockTag(tag = #minecraft:air),
  attempts = 60,
  xz_offset = Trapezoid(min = -2, max = 2, plateau = 0),
  y_offset = Uniform(min_inclusive = -3, max_inclusive = 0),
)

/**
 * Glowstone under the roof of a chamber: picks heights between the floor
 * and the roof, then scans up to the ceiling above.
 */
PlacedFeature roof_glowstone = {
  feature = glowstone_cluster
  placement = [
    Count(Uniform(4, 8)),
    InSquare(),
    HeightRange(Uniform(AboveBottom(10), BelowTop(24))),
    EnvironmentScan(
      direction_of_search = up,
      max_steps = 32,
      target_condition = Solid(),
      allowed_search_condition = MatchingBlockTag(tag = #minecraft:air),
    ),
    Offset(0, -1, 0),
    Biome(),
  ]
}

/** Fewer clusters, for biomes that only need some light. */
PlacedFeature roof_glowstone_sparse = {
  feature = glowstone_cluster
  placement = [
    Count(Uniform(1, 3)),
    InSquare(),
    HeightRange(Uniform(AboveBottom(10), BelowTop(24))),
    EnvironmentScan(
      direction_of_search = up,
      max_steps = 32,
      target_condition = Solid(),
      allowed_search_condition = MatchingBlockTag(tag = #minecraft:air),
    ),
    Offset(0, -1, 0),
    Biome(),
  ]
}

/** Glow caps on every floor, a few per chunk. */
PlacedFeature glow_caps = {
  feature = glow_cap
  placement = [
    CountOnEveryLayer(2),
    BlockPredicateFilter(
      MatchingBlocks(
        offset = [0, -1, 0],
        blocks = [grass_block, moss_block, podzol, mycelium, dirt, stone, deepslate],
      ),
    ),
    Biome(),
  ]
}

PlacedFeature ochre_crystals = {
  feature = ochre_crystals
  placement = [
    RarityFilter(2),
    InSquare(),
    HeightRange(Uniform(AboveBottom(10), BelowTop(24))),
    EnvironmentScan(
      direction_of_search = up,
      max_steps = 32,
      target_condition = Solid(),
      allowed_search_condition = MatchingBlockTag(tag = #minecraft:air),
    ),
    Offset(0, -1, 0),
    Biome(),
  ]
}

PlacedFeature verdant_crystals = {
  feature = verdant_crystals
  placement = [
    Count(2),
    InSquare(),
    HeightRange(Uniform(AboveBottom(10), BelowTop(24))),
    EnvironmentScan(
      direction_of_search = up,
      max_steps = 32,
      target_condition = Solid(),
      allowed_search_condition = MatchingBlockTag(tag = #minecraft:air),
    ),
    Offset(0, -1, 0),
    Biome(),
  ]
}

/**
 * Lights the starter cavern at x 0, z 0 (see terrain.mms): five glowstone
 * clusters under its roof. FixedPlacement only places in the chunk being
 * decorated, so every biome can list this and it runs once.
 */
PlacedFeature starter_lights = {
  feature = glowstone_cluster
  placement = [
    FixedPlacement([[0, 40, 0], [18, 40, 8], [-16, 40, -10], [10, 40, -18], [-12, 40, 16]]),
    EnvironmentScan(
      direction_of_search = up,
      max_steps = 20,
      target_condition = Solid(),
      allowed_search_condition = MatchingBlockTag(tag = #minecraft:air),
    ),
    Offset(0, -1, 0),
  ]
}

// ---- Plants ----------------------------------------------------------------

/** Grass on every lit floor. */
PlacedFeature floor_grass = {
  feature = minecraft:grass
  placement = [
    CountOnEveryLayer(12),
    BlockPredicateFilter(
      AllOf([MatchingBlockTag(tag = #minecraft:air), WouldSurvive(state = { id = short_grass })]),
    ),
    Biome(),
  ]
}

PlacedFeature floor_tall_grass = {
  feature = minecraft:tall_grass
  placement = [
    CountOnEveryLayer(2),
    BlockPredicateFilter(
      AllOf([MatchingBlockTag(tag = #minecraft:air), WouldSurvive(state = { id = tall_grass })]),
    ),
    Biome(),
  ]
}

PlacedFeature floor_flowers = {
  feature = minecraft:flower_default
  placement = [
    CountOnEveryLayer(2),
    BlockPredicateFilter(
      AllOf([MatchingBlockTag(tag = #minecraft:air), WouldSurvive(state = { id = poppy })]),
    ),
    Biome(),
  ]
}

PlacedFeature floor_pumpkins = {
  feature = minecraft:pumpkin
  placement = [
    RarityFilter(4),
    CountOnEveryLayer(1),
    BlockPredicateFilter(
      AllOf(
        [
          MatchingBlockTag(tag = #minecraft:air),
          MatchingBlocks(offset = [0, -1, 0], blocks = grass_block),
        ],
      ),
    ),
    Biome(),
  ]
}

PlacedFeature floor_melons = {
  feature = minecraft:melon
  placement = [
    RarityFilter(3),
    CountOnEveryLayer(1),
    BlockPredicateFilter(
      AllOf(
        [
          MatchingBlockTag(tag = #minecraft:air),
          MatchingBlocks(offset = [0, -1, 0], blocks = [grass_block, podzol]),
        ],
      ),
    ),
    Biome(),
  ]
}

/** Sugar cane on lake shores. WouldSurvive checks for water next to the block below. */
PlacedFeature shore_sugar_cane = {
  feature = minecraft:sugar_cane
  placement = [
    CountOnEveryLayer(8),
    BlockPredicateFilter(
      AllOf([MatchingBlockTag(tag = #minecraft:air), WouldSurvive(state = { id = sugar_cane })]),
    ),
    Biome(),
  ]
}

PlacedFeature lake_kelp = {
  feature = minecraft:kelp
  placement = [
    CountOnEveryLayer(4),
    BlockPredicateFilter(
      AllOf([MatchingBlocks(blocks = water), MatchingBlocks(offset = [0, 1, 0], blocks = water)]),
    ),
    Biome(),
  ]
}

PlacedFeature floor_mushrooms = {
  feature = minecraft:brown_mushroom
  placement = [
    CountOnEveryLayer(2),
    BlockPredicateFilter(
      AllOf(
        [MatchingBlockTag(tag = #minecraft:air), WouldSurvive(state = { id = brown_mushroom })],
      ),
    ),
    Biome(),
  ]
}

PlacedFeature floor_red_mushrooms = {
  feature = minecraft:red_mushroom
  placement = [
    CountOnEveryLayer(1),
    BlockPredicateFilter(
      AllOf([MatchingBlockTag(tag = #minecraft:air), WouldSurvive(state = { id = red_mushroom })]),
    ),
    Biome(),
  ]
}

PlacedFeature huge_mushrooms = {
  feature = minecraft:mushroom_island_vegetation
  placement = [
    CountOnEveryLayer(1),
    BlockPredicateFilter(MatchingBlocks(offset = [0, -1, 0], blocks = mycelium)),
    Biome(),
  ]
}

PlacedFeature floor_bamboo = {
  feature = minecraft:bamboo_no_podzol
  placement = [
    RarityFilter(2),
    CountOnEveryLayer(3),
    BlockPredicateFilter(
      AllOf([MatchingBlockTag(tag = #minecraft:air), WouldSurvive(state = { id = bamboo })]),
    ),
    Biome(),
  ]
}

// ---- Trees -----------------------------------------------------------------

/** Oak, birch and spruce, as one pick per spot. */
Feature rootwood_trees = RandomSelector(
  features = [
    {
      chance = 0.3
      feature = {
        feature = minecraft:birch
        placement = []
      }
    },
    {
      chance = 0.25
      feature = {
        feature = minecraft:spruce
        placement = []
      }
    },
    {
      chance = 0.15
      feature = {
        feature = minecraft:fancy_oak
        placement = []
      }
    },
  ],
  default = {
    feature = minecraft:oak
    placement = []
  },
)

/** Dark oak, jungle (with cocoa) and cherry. */
Feature darkroot_trees = RandomSelector(
  features = [
    {
      chance = 0.4
      feature = {
        feature = minecraft:jungle_tree
        placement = []
      }
    },
    {
      chance = 0.2
      feature = {
        feature = minecraft:cherry
        placement = []
      }
    },
  ],
  default = {
    feature = minecraft:dark_oak
    placement = []
  },
)

PlacedFeature rootwood_trees = {
  feature = rootwood_trees
  placement = [
    CountOnEveryLayer(3),
    BlockPredicateFilter(WouldSurvive(state = { id = oak_sapling })),
    Biome(),
  ]
}

PlacedFeature darkroot_trees = {
  feature = darkroot_trees
  placement = [
    CountOnEveryLayer(4),
    BlockPredicateFilter(WouldSurvive(state = { id = oak_sapling })),
    Biome(),
  ]
}

// ---- Lakes -----------------------------------------------------------------

Feature water_lake = Lake(
  fluid = { id = water },
  barrier = { id = sandstone },
  can_place_feature = True(),
  can_replace_with_air_or_fluid = Not(MatchingBlockTag(tag = #minecraft:features_cannot_replace)),
  can_replace_with_barrier = Not(MatchingBlockTag(tag = #minecraft:lava_pool_stone_cannot_replace)),
)

/** Finds a floor below a random height, then digs a lake into it. */
PlacedFeature water_lakes = {
  feature = water_lake
  placement = [
    Count(3),
    InSquare(),
    HeightRange(Uniform(AboveBottom(12), BelowTop(24))),
    EnvironmentScan(
      direction_of_search = down,
      max_steps = 32,
      target_condition = Solid(),
      allowed_search_condition = MatchingBlockTag(tag = #minecraft:air),
    ),
    Biome(),
  ]
}

/** Shallow lava pools, for early obsidian. */
PlacedFeature lava_lakes = {
  feature = minecraft:lake_lava
  placement = [
    Count(2),
    InSquare(),
    HeightRange(Uniform(AboveBottom(12), BelowTop(24))),
    EnvironmentScan(
      direction_of_search = down,
      max_steps = 32,
      target_condition = Solid(),
      allowed_search_condition = MatchingBlockTag(tag = #minecraft:air),
    ),
    Biome(),
  ]
}
