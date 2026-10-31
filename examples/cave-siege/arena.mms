// Test arenas for the prototype. Run these in a flat creative world.

/**
 * Builds a stone maze 17 by 25 blocks from where you stand, towards +x and +z.
 * The path from the emerald block (start) to the gold block (banner) is about 66 blocks.
 */
Function arena_maze {
  /fill ~ ~-1 ~ ~16 ~-1 ~24 minecraft:stone
  /fill ~ ~ ~ ~16 ~3 ~24 minecraft:air
  /fill ~ ~ ~ ~16 ~2 ~ minecraft:stone
  /fill ~ ~ ~24 ~16 ~2 ~24 minecraft:stone
  /fill ~ ~ ~ ~ ~2 ~24 minecraft:stone
  /fill ~16 ~ ~ ~16 ~2 ~24 minecraft:stone
  /fill ~ ~ ~6 ~12 ~2 ~6 minecraft:stone
  /fill ~4 ~ ~12 ~16 ~2 ~12 minecraft:stone
  /fill ~ ~ ~18 ~12 ~2 ~18 minecraft:stone
  /setblock ~2 ~-1 ~3 minecraft:emerald_block
  /setblock ~2 ~-1 ~21 minecraft:gold_block
  /time set midnight
  /tellraw @s "Maze built. Place the banner on the gold block, then spawn enemies on the emerald block."
}

/** Walls the planted banner in with a 5 by 5 cobblestone box, to test the breakers. */
Function seal_banner {
  /execute at @e[tag=cs_core,limit=1] run fill ~-2 ~-1 ~-2 ~2 ~2 ~2 minecraft:cobblestone outline
  /tellraw @s "Banner sealed in cobblestone. Enemies should get stuck, then dig in."
}

/** Walls the planted banner in with obsidian, the slowest block breakers can dig. */
Function seal_banner_obsidian {
  /execute at @e[tag=cs_core,limit=1] run fill ~-2 ~-1 ~-2 ~2 ~2 ~2 minecraft:obsidian outline
}
