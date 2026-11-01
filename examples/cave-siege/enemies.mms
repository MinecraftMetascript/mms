// Enemies: spawning, steering them along the path field, and noticing when they are stuck.
//
// Each enemy chases a waypoint: an invisible, tiny, silent iron golem that walks the path
// field a few blocks ahead of it. Zombies, skeletons and spiders attack iron golems on
// sight, so they pick up a waypoint by themselves and never take damage for it. Creepers
// only attack players, so a creeper is pointed at its first waypoint with one silent ping.
// Waypoints that reach the banner park there: damage they take is the banner's damage.

/** A damage type that only exists to point creepers at a waypoint. It does no knockback. */
DamageType creeper_ping = {
  message_id = "generic"
  exhaustion = 0
  scaling = "never"
}

/** Without this, the 0.01 damage of a ping knocks a creeper away from its waypoint. */
DamageTypeTag minecraft:no_knockback = [siege:creeper_ping]

/** Spawns 20 zombies where you stand. */
Function spawn_zombies {
  spawn("minecraft:zombie", 20)
}

/** Spawns 5 each of zombies, skeletons, spiders and creepers where you stand. */
Function spawn_mixed {
  spawn("minecraft:zombie", 5)
  spawn("minecraft:skeleton", 5)
  spawn("minecraft:spider", 5)
  spawn("minecraft:creeper", 5)
}

/** Spawns 3 sappers: zombies with pickaxes that always dig towards the banner. */
Function spawn_sappers {
  for i in 1..3 {
    /summon minecraft:zombie ~ ~ ~ {Tags:["cs_enemy","cs_new","cs_sapper"],PersistenceRequired:1b,CustomName:"Sapper",equipment:{mainhand:{id:"minecraft:iron_pickaxe",count:1}}}
  }
  /execute as @e[tag=cs_new] run function siege:enemy_init
}

/** Spawns count enemies of an entity type where the function runs. */
Function spawn(kind: string, count: int) {
  for i in 1..count {
    /summon ${kind} ~ ~ ~ {Tags:["cs_enemy","cs_new"],PersistenceRequired:1b}
  }
  /execute as @e[tag=cs_new] run function siege:enemy_init
}

/** Runs once as each new enemy. */
Function enemy_init {
  /tag @s remove cs_new
  /attribute @s minecraft:follow_range base set 64
  /execute store result score @s cs_group run random value 0..3
  cs_age[@s] = 0
  cs_break[@s] = 0
  cs_prev[@s] = 1000000
  cs_stuck[@s] = 0
  /execute if entity @s[tag=cs_sapper] run scoreboard players set @s cs_stuck 1000000
}

/** Runs as and at each enemy, every 4 ticks. */
Function enemy_update {
  cs_age[@s] += 1
  steer()
  // Every 5 seconds: has it got closer?
  if (cs_age[@s] % 25 == 0) {
    stuck_check()
  }
  if (cs_stuck[@s] > 0) {
    cs_stuck[@s] -= 1
    break_step()
  }
}

/**
 * Keeps the enemy chasing a waypoint. An enemy that is after a player is left alone, so it
 * fights. One with no target gets a new waypoint, unless there is one near it to pick up.
 */
Function steer {
  // 0: no target, 1: a player or anything else, 2: a waypoint, 3: a waypoint parked at the banner.
  cs_tmp["#tgt"] = 0
  /execute on target run scoreboard players set #tgt cs_tmp 1
  /execute on target if entity @s[tag=cs_carrot] run scoreboard players set #tgt cs_tmp 2
  /execute on target if entity @s[tag=cs_parked] run scoreboard players set #tgt cs_tmp 3
  let tgt = cs_tmp["#tgt"]
  if (tgt == 0) {
    if (cs_age[@s] % 5 == 1) {
      /execute unless entity @e[tag=cs_carrot,distance=..5] run function siege:give_carrot
    }
  }
  if (tgt >= 2) {
    // The waypoint is in use; unused ones are removed after 5 seconds.
    /execute on target run scoreboard players set @s cs_age 0
    // Creepers start to swell 3 blocks from their target, so theirs stays further ahead.
    /execute if entity @s[type=minecraft:creeper] on target if entity @s[distance=..5] at @s run function siege:carrot_far
    /execute unless entity @s[type=minecraft:creeper] on target if entity @s[distance=..3] at @s run function siege:carrot_near
  }
  bow_check(tgt == 3)
}

/** Skeletons with bows stop to shoot their waypoint, so they only get the bow at the banner. */
Function bow_check(parked: bool) {
  /execute unless entity @s[type=#minecraft:skeletons] run return 0
  if (parked) {
    /execute if entity @s[tag=cs_bowless] run item replace entity @s weapon.mainhand with minecraft:bow
    /tag @s remove cs_bowless
  } else {
    /execute if entity @s[tag=!cs_bowless] if items entity @s weapon.mainhand minecraft:bow run function siege:take_bow
  }
}

/** Runs as a skeleton away from the banner. */
Function take_bow {
  /item replace entity @s weapon.mainhand with minecraft:air
  /tag @s add cs_bowless
}

/** Runs as and at an enemy with no target: summons a waypoint in its cell and walks it a few cells ahead. */
Function give_carrot {
  /execute align xyz positioned ~0.5 ~ ~0.5 run summon minecraft:iron_golem ~ ~ ~ {Tags:["cs_carrot","cs_new_carrot"],NoAI:1b,NoGravity:1b,Silent:1b,PersistenceRequired:1b,PlayerCreated:1b,DeathLootTable:"minecraft:empty",attributes:[{id:"minecraft:max_health",base:1024d},{id:"minecraft:scale",base:0.25d},{id:"minecraft:knockback_resistance",base:1d}],Health:1024f}
  /effect give @e[tag=cs_new_carrot] minecraft:invisibility infinite 0 true
  // On the players' team, so players cannot hurt it; enemies are on no team.
  /team join cs_defenders @e[tag=cs_new_carrot]
  /scoreboard players set @e[tag=cs_new_carrot] cs_age 0
  /execute as @e[tag=cs_new_carrot] at @s run function siege:carrot_near
  // A creeper does not attack golems: one silent ping points it at the waypoint.
  /execute if entity @s[type=minecraft:creeper] run function siege:ping_creeper
  /tag @e[tag=cs_new_carrot] remove cs_new_carrot
}

/** Runs as a creeper. */
Function ping_creeper {
  /data merge entity @s {Silent:1b}
  /damage @s 0.01 siege:creeper_ping by @e[tag=cs_new_carrot,limit=1]
  /data merge entity @s {Silent:0b}
}

/** Runs as and at a waypoint: 3 cells ahead, within the 4 blocks at which mobs notice an invisible golem. */
Function carrot_near {
  for i in 1..3 {
    /execute at @s run function siege:carrot_step
  }
}

/** Runs as and at a creeper's waypoint: 7 cells ahead. */
Function carrot_far {
  for i in 1..7 {
    /execute at @s run function siege:carrot_step
  }
}

/**
 * An enemy that is not 2 blocks closer than 5 seconds ago, with no player near, becomes a breaker
 * for 10 seconds. That includes one right next to the banner behind a 1-block wall; one already at
 * the banner only "digs" walkable blocks, which does nothing.
 */
Function stuck_check {
  /execute store result score #ex cs_tmp run data get entity @s Pos[0]
  /execute store result score #ey cs_tmp run data get entity @s Pos[1]
  /execute store result score #ez cs_tmp run data get entity @s Pos[2]
  let dx = cs_tmp["#ex"] - cs_tmp["#cx"]
  let dy = cs_tmp["#ey"] - cs_tmp["#cy"]
  let dz = cs_tmp["#ez"] - cs_tmp["#cz"]
  if (dx < 0) {
    dx = -dx
  }
  if (dy < 0) {
    dy = -dy
  }
  if (dz < 0) {
    dz = -dz
  }
  let d = dx + dy + dz
  /execute store result score #near cs_tmp if entity @a[distance=..8,gamemode=!spectator,gamemode=!creative]
  if (d > cs_prev[@s] - 2 && cs_stuck[@s] <= 0 && cs_tmp["#near"] == 0) {
    cs_stuck[@s] = 50
  }
  cs_prev[@s] = d
}
