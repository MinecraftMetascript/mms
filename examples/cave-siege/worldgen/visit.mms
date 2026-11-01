// Ways into the cave world.

/** Takes you to siege:caves, the cave world as its own dimension, in the starter cavern. */
Function visit_caves {
  /execute in siege:caves run spreadplayers 0 0 0 8 under 40 false @s
  /tellraw @s "Welcome to the caves. The starter cavern is at x 0, z 0."
}

/**
 * In a world made with the Cave Siege world type, puts you on the floor of
 * the starter cavern and makes it the world spawn. Spawn search in a world
 * with a roof can pick a poor spot, so the game calls this on first join.
 */
Function go_to_starter_cavern {
  /execute in minecraft:overworld run spreadplayers 0 0 0 8 under 40 false @s
  /execute at @s run setworldspawn ~ ~ ~
}
