// The Siege Banner: placing it, the core entities, and its health.

/** Gives you the Siege Banner. */
Function give_banner {
  /give @s minecraft:white_banner[custom_name="Siege Banner",custom_data={siege_banner:1b},enchantment_glint_override=true,banner_patterns=[{pattern:"minecraft:flower",color:"red"},{pattern:"minecraft:border",color:"black"}]]
}

/** Granted when a player places the Siege Banner; it runs on_place and is revoked straight away. */
Advancement placed_banner = {
  criteria = {
    placed = PlacedBlock(
      location = MatchTool(
        predicate = {
          items = [white_banner]
          predicates = { "minecraft:custom_data" = "{siege_banner:1b}" }
        },
      ),
    )
  }
  rewards = { function = siege:on_place }
}

/** Runs as the player who placed the banner. Advancements do not say where the block went, so a ray from the eyes finds it. */
Function on_place {
  /advancement revoke @s only siege:placed_banner
  cs_tmp["#ray"] = 0
  /execute at @s anchored eyes positioned ^ ^ ^ run function siege:find_banner
}

/** One step of the ray: stops at the first banner block, up to 6 blocks out. */
Function find_banner {
  /execute if block ~ ~ ~ #minecraft:banners run return run function siege:found_banner
  cs_tmp["#ray"] += 1
  if (cs_tmp["#ray"] < 60) {
    /execute positioned ^ ^ ^0.1 run function siege:find_banner
  }
}

/** Runs at the placed banner block. */
Function found_banner {
  if (banner_placed) {
    /setblock ~ ~ ~ air
    give_banner()
    /tellraw @s "There is already a Siege Banner planted"
    return
  }
  /execute align xyz positioned ~0.5 ~ ~0.5 run function siege:create_core
}

/** Runs at the centre of the placed banner's cell: the core marker holds the banner's position. */
Function create_core {
  /summon minecraft:marker ~ ~ ~ {Tags:["cs_core"]}
  /execute store result score #cx cs_tmp run data get entity @e[tag=cs_core,limit=1] Pos[0]
  /execute store result score #cy cs_tmp run data get entity @e[tag=cs_core,limit=1] Pos[1]
  /execute store result score #cz cs_tmp run data get entity @e[tag=cs_core,limit=1] Pos[2]
  /spawnpoint @a[team=cs_defenders] ~ ~ ~
  banner_placed = true
  banner_hp = banner_max
  ticks = 0
  repath()
  /bossbar set siege:banner max ${banner_max}
  /bossbar set siege:banner value ${banner_hp}
  /bossbar set siege:banner players @a
  /tellraw @a "The Siege Banner is planted. Defend it!"
}

/** Runs once a second while the banner is planted. */
Function banner_second {
  // A missing banner block (broken, burned, pushed) is put back as it was, and the banner it dropped is removed.
  /execute at @e[tag=cs_core,limit=1] unless block ~ ~ ~ #minecraft:banners run function siege:restore_banner
  /execute at @e[tag=cs_core,limit=1] run kill @e[type=minecraft:item,distance=..3,nbt={Item:{id:"minecraft:white_banner"}}]
  // Rebuild the field after a breaker opened a wall, and every 10 seconds for blocks players placed or broke.
  if (path_phase == 0 && (path_dirty || ticks % 200 == 0)) {
    repath()
  }
  carrot_second()
  // Damage taken by waypoints parked at the banner this second comes off the banner.
  cs_tmp["#lost"] = 0
  /execute as @e[tag=cs_parked] run function siege:parked_damage
  let lost = cs_tmp["#lost"]
  if (lost > 0) {
    banner_hp -= lost
    /execute at @e[tag=cs_core,limit=1] run particle minecraft:damage_indicator ~ ~1 ~ 0.3 0.5 0.3 0 5
  }
  if (banner_hp <= 0) {
    banner_falls()
    return
  }
  /bossbar set siege:banner value ${banner_hp}
  /bossbar set siege:banner players @a
}

/** Runs as each parked waypoint: adds the health it lost to #lost. */
Function parked_damage {
  /execute store result score #hp cs_tmp run data get entity @s Health
  cs_tmp["#lost"] += 1024 - cs_tmp["#hp"]
  /data modify entity @s Health set value 1024f
  /execute at @e[tag=cs_core,limit=1] run tp @s ~ ~ ~
}

/** Waypoint upkeep, once a second: heal them all, and remove ones no enemy has chased for 5 seconds. */
Function carrot_second {
  /execute as @e[tag=cs_carrot,tag=!cs_parked] run data modify entity @s Health set value 1024f
  /scoreboard players add @e[tag=cs_carrot] cs_age 1
  /execute as @e[tag=cs_carrot,scores={cs_age=6..}] run function siege:remove_carrot
}

/** Removes a waypoint without a death animation or drops. */
Function remove_carrot {
  /execute at @s run tp @s ~ -1000 ~
  /kill @s
}

/** The banner reached 0 health: the siege is lost. */
Function banner_falls {
  /title @a title "The banner has fallen"
  /execute at @e[tag=cs_core,limit=1] run playsound minecraft:entity.wither.death master @a ~ ~ ~ 1 0.6
  reset()
}
