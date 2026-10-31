// Breakers dig through whatever stands between them and the banner.

/** Blocks a mob walks through: nothing to break. */
BlockTag walkable = [
  #minecraft:air,
  #minecraft:replaceable,
  #minecraft:wall_post_override,
  #minecraft:climbable,
  #minecraft:banners,
  minecraft:light,
]

/** Never broken. */
BlockTag unbreakable = [
  minecraft:bedrock,
  minecraft:reinforced_deepslate,
  minecraft:end_portal_frame,
  minecraft:barrier,
  minecraft:command_block,
  minecraft:structure_block,
]

/** 1 second: soft blocks. */
BlockTag break_fast = [
  #minecraft:mineable/axe,
  #minecraft:mineable/shovel,
  #minecraft:mineable/hoe,
  #minecraft:wool,
  #minecraft:impermeable,
]

/** 6 seconds: iron. */
BlockTag break_metal = [
  minecraft:iron_block,
  minecraft:iron_door,
  minecraft:iron_bars,
  minecraft:iron_trapdoor,
  #minecraft:anvil,
]

/** 15 seconds: obsidian and similar. */
BlockTag break_obsidian = [
  minecraft:obsidian,
  minecraft:crying_obsidian,
  minecraft:respawn_anchor,
  minecraft:netherite_block,
  minecraft:ancient_debris,
]

/** Runs as and at a breaker: looks at the block in front of it, level with the banner's direction. */
Function break_step {
  /execute facing entity @e[tag=cs_lure,limit=1] feet rotated ~ 0 positioned ^ ^ ^0.8 align xyz run function siege:break_at
}

/** Runs at the block in front of a breaker's feet: digs the feet block first, then the head block. */
Function break_at {
  /execute unless block ~ ~ ~ #siege:walkable run return run function siege:break_block
  /execute unless block ~ ~1 ~ #siege:walkable positioned ~ ~1 ~ run return run function siege:break_block
  cs_break[@s] = 0
}

/** Runs at the block being dug. Stone and everything else not listed takes 3 seconds. */
Function break_block {
  /execute if block ~ ~ ~ #siege:unbreakable run return 0
  cs_tmp["#need"] = 15
  /execute if block ~ ~ ~ #siege:break_fast run scoreboard players set #need cs_tmp 5
  /execute if block ~ ~ ~ #siege:break_metal run scoreboard players set #need cs_tmp 30
  /execute if block ~ ~ ~ #siege:break_obsidian run scoreboard players set #need cs_tmp 75
  cs_break[@s] += 1
  /particle minecraft:crit ~0.5 ~0.5 ~0.5 0.3 0.3 0.3 0.1 4
  /playsound minecraft:block.stone.hit hostile @a ~ ~ ~ 0.6 0.8
  if (cs_break[@s] >= cs_tmp["#need"]) {
    /setblock ~ ~ ~ air destroy
    cs_break[@s] = 0
  }
}
