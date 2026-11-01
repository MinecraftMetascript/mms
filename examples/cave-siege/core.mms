// Cave Siege, phase 1: the siege core prototype.
//
// Try it in a flat creative world:
//   /function siege:arena_maze      build a test maze where you stand
//   /function siege:give_banner     get the Siege Banner, then place it on the gold block
//   /function siege:spawn_zombies   stand on the emerald block and spawn 20 zombies
//   /function siege:seal_banner     wall the banner in, to test the breakers
//   /function siege:status          banner health and enemy count
//   /function siege:reset           remove the banner, the core and every enemy

/** Whether a Siege Banner is planted. */
Global banner_placed: bool = false

/** The banner's health. It falls when this reaches 0. */
Global banner_hp: int = 100

/** The banner's health when planted. */
Global banner_max: int = 100

/** Counts ticks while the banner is planted, to spread work over ticks. */
Global ticks: int = 0

// Objective names are global to the world, so they carry a cs_ prefix.

/** Scratch values: #cx #cy #cz (core position), #ex #ey #ez, #lost, #hp, #tgt, #need, #near. */
Objective cs_tmp = dummy
/** Which of the 4 update groups an enemy is in, 0 to 3. */
Objective cs_group = dummy
/** Updates an enemy has had, to run some checks less often. */
Objective cs_age = dummy
/** An enemy's distance to the banner at the last stuck check (blocks, x + y + z). */
Objective cs_prev = dummy
/** Updates left in breaker mode. Sappers start with a very large value. */
Objective cs_stuck = dummy
/** Progress on the block a breaker is digging, in updates. */
Objective cs_break = dummy

OnLoad {
  /team add cs_defenders "Defenders"
  /team modify cs_defenders friendlyFire false
  /bossbar add siege:banner "Siege Banner"
  /bossbar set siege:banner color red
  /bossbar set siege:banner style notched_10
}

OnTick {
  /team join cs_defenders @a[team=]
  if (banner_placed) {
    siege_tick()
  }
}

/** Runs every tick while the banner is planted. */
Function siege_tick {
  ticks += 1
  // Each enemy updates every 4 ticks; a quarter of them run each tick.
  let g = ticks % 4
  if (g == 0) {
    /execute as @e[tag=cs_enemy,scores={cs_group=0}] at @s run function siege:enemy_update
  } else if (g == 1) {
    /execute as @e[tag=cs_enemy,scores={cs_group=1}] at @s run function siege:enemy_update
  } else if (g == 2) {
    /execute as @e[tag=cs_enemy,scores={cs_group=2}] at @s run function siege:enemy_update
  } else {
    /execute as @e[tag=cs_enemy,scores={cs_group=3}] at @s run function siege:enemy_update
  }
  if (path_phase > 0) {
    path_tick()
  }
  if (ticks % 20 == 0) {
    banner_second()
  }
}

/** Prints the banner's health and how many enemies are alive. */
Function status {
  /execute store result score #enemies cs_tmp if entity @e[tag=cs_enemy]
  let enemies = cs_tmp["#enemies"]
  if (banner_placed) {
    /tellraw @s "Banner: ${banner_hp} / ${banner_max} health, ${enemies} enemies alive"
  } else {
    /tellraw @s "No banner is planted, ${enemies} enemies alive"
  }
}

/** Removes the banner, the core, the waypoints and every enemy. */
Function reset {
  /kill @e[tag=cs_enemy]
  /execute at @e[tag=cs_core] run setblock ~ ~ ~ air
  /execute as @e[tag=cs_carrot] run function siege:remove_carrot
  // Forget the saved banner and unload the scratch area.
  /execute at @e[tag=cs_core,limit=1] in siege:scratch run setblock ~ 319 ~ minecraft:air
  /execute in siege:scratch run forceload remove all
  path_phase = 0
  /kill @e[tag=cs_core]
  /bossbar set siege:banner players
  banner_placed = false
  banner_hp = banner_max
  /tellraw @a "Cave Siege reset"
}
