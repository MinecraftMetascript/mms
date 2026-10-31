// Enemies: spawning, steering them to the lure, and noticing when they are stuck.

/** A damage type that only exists to point mobs at the lure. It does no knockback. */
DamageType lure_ping = {
  message_id = "generic"
  exhaustion = 0
  scaling = "never"
}

/** Without this, the 0.01 damage of a lure ping knocks enemies away from the banner. */
DamageTypeTag minecraft:no_knockback = [siege:lure_ping]

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
  retarget()
}

/** Runs as and at each enemy, every 4 ticks. */
Function enemy_update {
  cs_age[@s] += 1
  // Every 2 seconds: point it at the lure again.
  if (cs_age[@s] % 10 == 1) {
    retarget()
  }
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
 * Makes the enemy target the lure, by having the lure "hurt" it: most hostile
 * mobs turn on whatever hurt them. Enemies close to a player are left alone so
 * they can fight.
 */
Function retarget {
  /execute unless entity @a[distance=..8,gamemode=!spectator,gamemode=!creative] run damage @s 0.01 siege:lure_ping by @e[tag=cs_lure,limit=1]
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
