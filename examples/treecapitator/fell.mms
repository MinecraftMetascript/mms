// Felling the tree. A flood fill from the chopped log finds every log
// joined to it at its height or above, marking each with a marker entity
// so it is visited once, and counts the natural leaves touching them.
// Only when the logs look like a tree does anything break.

/** The most logs one chop can fell. A bigger group of logs is left alone. */
Global max_logs: int = 256
/** How many natural leaves (or nether wart blocks) must touch the logs for them to count as a tree. */
Global min_leaves: int = 3

// Working values for the tree being felled. Raw commands below write some
// of them as the scores $<name> on mms.treecap.
Global logs: int = 0
Global leaves: int = 0
Global unbreaking: int = 0
Global roll: int = 0
Global damage: int = 0

/**
 * Leaves a player has placed are persistent and never decay, so only
 * leaves that grew with a tree count. Huge fungi in the nether have caps
 * of wart blocks instead.
 */
BlockTag caps = [nether_wart_block, warped_wart_block, shroomlight]

/**
 * Runs as the player who chopped a log, positioned at the centre of the
 * block where the log was.
 */
Function fell {
  logs = 0
  leaves = 0
  /summon marker ~ ~ ~ {Tags:["treecap.node","treecap.origin"]}
  spread()
  if (logs > max_logs) {
    /title @s actionbar {text:"That is too big to fell at once",color:"gray"}
  } else if (logs > 0 && leaves >= min_leaves) {
    /execute as @e[type=marker,tag=treecap.node,tag=!treecap.origin] at @s run setblock ~ ~ ~ air destroy
    // Gather the drops where the chopped log was, so they do not land
    // all over the canopy.
    /execute as @e[type=marker,tag=treecap.node,tag=!treecap.origin] at @s as @e[type=item,distance=..1,nbt={Age:0s}] run tp @s @e[type=marker,tag=treecap.origin,limit=1]
    wear_axe()
  }
  /kill @e[type=marker,tag=treecap.node]
}

/**
 * Runs at the centre of a log, or of the chopped log's space. Visits the
 * logs beside it and above it: the 8 around it and the 9 above, but never
 * below, so the stump stays when a trunk is chopped part way up.
 */
Function spread {
  /execute positioned ~1 ~ ~ if block ~ ~ ~ #minecraft:logs unless entity @e[type=marker,tag=treecap.node,distance=..0.5] run function treecap:visit
  /execute positioned ~-1 ~ ~ if block ~ ~ ~ #minecraft:logs unless entity @e[type=marker,tag=treecap.node,distance=..0.5] run function treecap:visit
  /execute positioned ~ ~ ~1 if block ~ ~ ~ #minecraft:logs unless entity @e[type=marker,tag=treecap.node,distance=..0.5] run function treecap:visit
  /execute positioned ~ ~ ~-1 if block ~ ~ ~ #minecraft:logs unless entity @e[type=marker,tag=treecap.node,distance=..0.5] run function treecap:visit
  /execute positioned ~1 ~ ~1 if block ~ ~ ~ #minecraft:logs unless entity @e[type=marker,tag=treecap.node,distance=..0.5] run function treecap:visit
  /execute positioned ~1 ~ ~-1 if block ~ ~ ~ #minecraft:logs unless entity @e[type=marker,tag=treecap.node,distance=..0.5] run function treecap:visit
  /execute positioned ~-1 ~ ~1 if block ~ ~ ~ #minecraft:logs unless entity @e[type=marker,tag=treecap.node,distance=..0.5] run function treecap:visit
  /execute positioned ~-1 ~ ~-1 if block ~ ~ ~ #minecraft:logs unless entity @e[type=marker,tag=treecap.node,distance=..0.5] run function treecap:visit
  /execute positioned ~ ~1 ~ if block ~ ~ ~ #minecraft:logs unless entity @e[type=marker,tag=treecap.node,distance=..0.5] run function treecap:visit
  /execute positioned ~1 ~1 ~ if block ~ ~ ~ #minecraft:logs unless entity @e[type=marker,tag=treecap.node,distance=..0.5] run function treecap:visit
  /execute positioned ~-1 ~1 ~ if block ~ ~ ~ #minecraft:logs unless entity @e[type=marker,tag=treecap.node,distance=..0.5] run function treecap:visit
  /execute positioned ~ ~1 ~1 if block ~ ~ ~ #minecraft:logs unless entity @e[type=marker,tag=treecap.node,distance=..0.5] run function treecap:visit
  /execute positioned ~ ~1 ~-1 if block ~ ~ ~ #minecraft:logs unless entity @e[type=marker,tag=treecap.node,distance=..0.5] run function treecap:visit
  /execute positioned ~1 ~1 ~1 if block ~ ~ ~ #minecraft:logs unless entity @e[type=marker,tag=treecap.node,distance=..0.5] run function treecap:visit
  /execute positioned ~1 ~1 ~-1 if block ~ ~ ~ #minecraft:logs unless entity @e[type=marker,tag=treecap.node,distance=..0.5] run function treecap:visit
  /execute positioned ~-1 ~1 ~1 if block ~ ~ ~ #minecraft:logs unless entity @e[type=marker,tag=treecap.node,distance=..0.5] run function treecap:visit
  /execute positioned ~-1 ~1 ~-1 if block ~ ~ ~ #minecraft:logs unless entity @e[type=marker,tag=treecap.node,distance=..0.5] run function treecap:visit
}

/**
 * Runs at the centre of a log that has not been visited yet: marks it,
 * counts the leaves on its six faces and carries on from it. The calls go
 * through raw /function commands, so the flood fill keeps no stack: the
 * counts are Globals that every visit shares.
 */
Function visit {
  /summon marker ~ ~ ~ {Tags:["treecap.node"]}
  logs += 1
  if (logs > max_logs) {
    return
  }
  /execute if block ~1 ~ ~ #minecraft:leaves[persistent=false] run scoreboard players add $leaves mms.treecap 1
  /execute if block ~-1 ~ ~ #minecraft:leaves[persistent=false] run scoreboard players add $leaves mms.treecap 1
  /execute if block ~ ~1 ~ #minecraft:leaves[persistent=false] run scoreboard players add $leaves mms.treecap 1
  /execute if block ~ ~-1 ~ #minecraft:leaves[persistent=false] run scoreboard players add $leaves mms.treecap 1
  /execute if block ~ ~ ~1 #minecraft:leaves[persistent=false] run scoreboard players add $leaves mms.treecap 1
  /execute if block ~ ~ ~-1 #minecraft:leaves[persistent=false] run scoreboard players add $leaves mms.treecap 1
  /execute if block ~1 ~ ~ #treecap:caps run scoreboard players add $leaves mms.treecap 1
  /execute if block ~-1 ~ ~ #treecap:caps run scoreboard players add $leaves mms.treecap 1
  /execute if block ~ ~1 ~ #treecap:caps run scoreboard players add $leaves mms.treecap 1
  /execute if block ~ ~-1 ~ #treecap:caps run scoreboard players add $leaves mms.treecap 1
  /execute if block ~ ~ ~1 #treecap:caps run scoreboard players add $leaves mms.treecap 1
  /execute if block ~ ~ ~-1 #treecap:caps run scoreboard players add $leaves mms.treecap 1
  /function treecap:spread
}

/**
 * Takes a point of durability from the held axe for each extra log, as if
 * each had been chopped by hand. Unbreaking skips each point with the
 * same odds it would in vanilla: 1 in (level + 1) points is taken.
 */
Function wear_axe {
  /execute if items entity @s weapon.mainhand *[minecraft:unbreakable] run return 0
  unbreaking = 0
  /execute if items entity @s weapon.mainhand *[minecraft:enchantments~[{enchantments:"minecraft:unbreaking",levels:1}]] run scoreboard players set $unbreaking mms.treecap 1
  /execute if items entity @s weapon.mainhand *[minecraft:enchantments~[{enchantments:"minecraft:unbreaking",levels:2}]] run scoreboard players set $unbreaking mms.treecap 2
  /execute if items entity @s weapon.mainhand *[minecraft:enchantments~[{enchantments:"minecraft:unbreaking",levels:3}]] run scoreboard players set $unbreaking mms.treecap 3
  // Whole shares of the points, then one roll for what is left over.
  let cost = logs / (unbreaking + 1)
  let rest = logs % (unbreaking + 1)
  if (rest > 0) {
    /execute store result score $roll mms.treecap run random value 0..${unbreaking}
    if (roll < rest) {
      cost += 1
    }
  }
  if (cost == 0) {
    return
  }
  // Read the axe's damage from a copy on a scratch item display. A
  // missing damage component reads as 0, the same as an unused axe.
  /summon item_display ~ ~ ~ {Tags:["treecap.tmp"]}
  /item replace entity @e[type=item_display,tag=treecap.tmp,limit=1] contents from entity @s weapon.mainhand
  /execute store result score $damage mms.treecap run data get entity @e[type=item_display,tag=treecap.tmp,limit=1] item.components."minecraft:damage"
  /kill @e[type=item_display,tag=treecap.tmp]
  damage += cost
  /item modify entity @s weapon.mainhand {function:"minecraft:set_components",components:{"minecraft:damage":${damage}}}
  /execute if items entity @s weapon.mainhand *[minecraft:damage~{durability:..0}] run function treecap:break_axe
}

/** Runs as a player whose axe has no durability left. */
Function break_axe {
  /item replace entity @s weapon.mainhand with air
  /playsound minecraft:entity.item.break player @a
}
