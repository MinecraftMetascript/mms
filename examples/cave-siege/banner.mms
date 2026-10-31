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

/**
 * Summons the core at the banner: a marker that holds its position, and the
 * lure, a hidden baby villager that enemies are told to attack. Damage the
 * lure takes is the banner's damage.
 */
Function create_core {
  /summon minecraft:marker ~ ~ ~ {Tags:["cs_core"]}
  /summon minecraft:villager ~ ~ ~ {Tags:["cs_lure"],NoAI:1b,Silent:1b,PersistenceRequired:1b,Age:-2147483648,CanPickUpLoot:0b,attributes:[{id:"minecraft:max_health",base:1024d},{id:"minecraft:knockback_resistance",base:1d}],Health:1024f,active_effects:[{id:"minecraft:invisibility",duration:-1,show_particles:0b}]}
  // On the players' team, so players cannot hurt it; enemies are on no team.
  /team join cs_defenders @e[tag=cs_lure]
  /execute store result score #cx cs_tmp run data get entity @e[tag=cs_core,limit=1] Pos[0]
  /execute store result score #cy cs_tmp run data get entity @e[tag=cs_core,limit=1] Pos[1]
  /execute store result score #cz cs_tmp run data get entity @e[tag=cs_core,limit=1] Pos[2]
  /spawnpoint @a[team=cs_defenders] ~ ~ ~
  banner_placed = true
  banner_hp = banner_max
  ticks = 0
  /bossbar set siege:banner max ${banner_max}
  /bossbar set siege:banner value ${banner_hp}
  /bossbar set siege:banner players @a
  /tellraw @a "The Siege Banner is planted. Defend it!"
}

/** Runs once a second while the banner is planted. */
Function banner_second {
  // A missing banner block (broken, burned, pushed) is put back.
  /execute at @e[tag=cs_core,limit=1] unless block ~ ~ ~ #minecraft:banners run setblock ~ ~ ~ minecraft:white_banner
  // Damage the lure took this second comes off the banner; then the lure heals.
  /execute store result score #lure_hp cs_tmp run data get entity @e[tag=cs_lure,limit=1] Health
  let lost = 1024 - cs_tmp["#lure_hp"]
  /data modify entity @e[tag=cs_lure,limit=1] Health set value 1024f
  /execute as @e[tag=cs_lure] at @e[tag=cs_core,limit=1] run tp @s ~ ~ ~
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

/** The banner reached 0 health: the siege is lost. */
Function banner_falls {
  /title @a title "The banner has fallen"
  /execute at @e[tag=cs_core,limit=1] run playsound minecraft:entity.wither.death master @a ~ ~ ~ 1 0.6
  reset()
}
