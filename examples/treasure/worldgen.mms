// World generation: wrecks that spawn in new chunks with chests/buried
// inside. No .nbt file is needed, since the pool reuses a vanilla template
// and a processor swaps in this pack's loot table.
//
// JigsawStructure writes the three resources a structure needs: the
// structure set (how often and where wrecks may start), the structure
// (which biomes, and at what height) and the template pool (which
// template to place, and the processors to run). The ProcessorList holds
// rules that edit blocks as the template is placed.
//
// Try it: make a new world with the pack, then /locate structure treasure:wreck.

/** Biomes that wrecks spawn in. */
BiomeTag wreck_biomes = [plains, sunflower_plains, savanna, desert]

/** Half a ship, sunk two blocks into the ground, about one per 24 by 24 chunks and never closer than 8 chunks apart. */
JigsawStructure wreck {
  biomes = #wreck_biomes
  step = surface_structures
  spacing = 24
  separation = 8
  height = Constant(Absolute(-2))
  heightmap = WORLD_SURFACE_WG
  depth = 1
  start = wreck
  pools = {
    wreck = [
      WeightedElement(
        1,
        SinglePoolElement(rigid, minecraft:shipwreck/rightsideup_backhalf, wreck_loot),
      ),
    ]
  }
}

/**
 * Fills the template's two chests from chests/buried, and rots some planks.
 * A rule replaces the whole block state, so each rule matches one chest's
 * state and writes the same state back, to keep the way the chest faces.
 * Properties left out, such as waterlogged, take their default values.
 */
ProcessorList wreck_loot {
  processors = [
    Rule(
      [
        {
          input_predicate = BlockstateMatch({
            id = chest
            properties = { facing = "west" }
          })
          location_predicate = AlwaysTrue()
          output_state = {
            id = chest
            properties = { facing = "west" }
          }
          block_entity_modifier = AppendLoot(chests/buried)
        },
        {
          input_predicate = BlockstateMatch({
            id = chest
            properties = { facing = "north" }
          })
          location_predicate = AlwaysTrue()
          output_state = {
            id = chest
            properties = { facing = "north" }
          }
          block_entity_modifier = AppendLoot(chests/buried)
        },
      ],
    ),
    BlockRot(integrity = 0.85, rottable_blocks = #planks),
  ]
}
