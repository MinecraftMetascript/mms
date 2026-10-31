// World generation: wrecks that spawn in new chunks with chests/buried
// inside. No .nbt file is needed, since the pool reuses a vanilla template
// and a processor swaps in this pack's loot table.
//
// Four resources work together:
// - StructureSet: how often and where in the world wrecks may start.
// - Structure: what kind of structure it is, which biomes, and at what height.
// - TemplatePool: which template (.nbt) to place, and the processors to run.
// - ProcessorList: rules that edit blocks as the template is placed.
//
// Try it: make a new world with the pack, then /locate structure treasure:wreck.

/** Biomes that wrecks spawn in. */
BiomeTag wreck_biomes = [plains, sunflower_plains, savanna, desert]

/** About one wreck per 24 by 24 chunks, never closer than 8 chunks apart. */
StructureSet wreck = {
  structures = [
    {
      structure = wreck
      weight = 1
    },
  ]
  placement = RandomSpread(spacing = 24, separation = 8, salt = 20261031)
}

/** Half a ship, sunk two blocks into the ground. */
Structure wreck = Jigsaw(
  biomes = #wreck_biomes,
  step = surface_structures,
  spawn_overrides = {},
  start_pool = wreck,
  size = 1,
  start_height = Constant({ absolute = -2 }),
  project_start_to_heightmap = WORLD_SURFACE_WG,
  max_distance_from_center = 80,
  use_expansion_hack = false,
)

TemplatePool wreck = {
  fallback = minecraft:empty
  elements = [
    {
      weight = 1
      element = SinglePoolElement(
        projection = rigid,
        location = minecraft:shipwreck/rightsideup_backhalf,
        processors = wreck_loot,
      )
    },
  ]
}

/**
 * Fills the template's two chests from chests/buried, and rots some planks.
 * A rule replaces the whole block state, so each rule matches one chest's
 * state and writes the same state back, to keep the way the chest faces.
 * Properties left out, such as waterlogged, take their default values.
 */
ProcessorList wreck_loot = {
  processors = [
    Rule(
      [
        {
          input_predicate = BlockstateMatch(
            {
              id = chest
              properties = { facing = "west" }
            },
          )
          location_predicate = AlwaysTrue()
          output_state = {
            id = chest
            properties = { facing = "west" }
          }
          block_entity_modifier = AppendLoot(chests/buried)
        },
        {
          input_predicate = BlockstateMatch(
            {
              id = chest
              properties = { facing = "north" }
            },
          )
          location_predicate = AlwaysTrue()
          output_state = {
            id = chest
            properties = { facing = "north" }
          }
          block_entity_modifier = AppendLoot(chests/buried)
        },
      ],
    ),
    BlockRot(integrity = 0.85, rottable_blocks = #minecraft:planks),
  ]
}
