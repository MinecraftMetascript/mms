// Placing treasure chests. This is where chests/buried gets used.
//
// Try it in game: stand on some planks and run /function treasure:hide_chest,
// then open the chest. Its contents are rolled from chests/buried right then.

/**
 * Places a chest at your feet that is filled from treasure:chests/buried,
 * if you are standing on one of the chest_floors blocks.
 */
Function hide_chest {
  /execute if block ~ ~-1 ~ #treasure:chest_floors run setblock ~ ~ ~ chest{LootTable:"treasure:chests/buried"}
  /execute unless block ~ ~-1 ~ #treasure:chest_floors run tellraw @s "Stand on logs, planks or stone bricks to hide a chest"
}

/** Rolls chests/buried straight into your inventory, to preview the loot. */
Function preview_loot {
  /loot give @s loot treasure:chests/buried
}
