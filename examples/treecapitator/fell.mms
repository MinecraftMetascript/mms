// Felling the tree. A flood fill from the chopped log finds every log of
// the same wood joined to it at its height or above, marking each with a
// marker entity so it is visited once, and notes which sides of the logs
// touch natural leaves. Only when the logs look like a tree does anything
// break: then a second fill finds the tree's own leaves, and logs and
// leaves break together.

/** The most logs one chop can fell. A bigger group of logs is left alone. */
Global max_logs: int = 256
/** Whether the tree's leaves break with it. Otherwise they decay on their own. */
Global break_leaves: bool = true
/** The most leaves one chop breaks. Leaves past this decay on their own. */
Global max_leaves: int = 600

/** The block tag of the wood being felled, such as minecraft:oak_logs. */
Global kind: string = ""
/** Whether the tree is a huge fungus, whose cap is wart blocks instead of leaves. */
Global nether: bool = false

// Working values for the tree being felled. Raw commands below write some
// of them as the scores $<name> on mms.treecap.
Global logs: int = 0
Global leaf_count: int = 0
/** Set to 1 when a log has natural leaves to its east, west, south, north or top. */
Global east: int = 0
Global west: int = 0
Global south: int = 0
Global north: int = 0
Global top: int = 0
Global unbreaking: int = 0
Global roll: int = 0
Global damage: int = 0

/** The distance value of the leaf block under each leaf marker, 1 to 7. */
Objective leaf_dist

/**
 * Leaves a player has placed are persistent and never decay, so only
 * leaves that grew with a tree count. Huge fungi in the nether have caps
 * of wart blocks instead.
 */
BlockTag caps = [nether_wart_block, warped_wart_block, shroomlight]

/** The six blocks that share a face with this one. */
const faces = ["~1 ~ ~", "~-1 ~ ~", "~ ~1 ~", "~ ~-1 ~", "~ ~ ~1", "~ ~ ~-1"]

/** The 8 blocks around this one, at its height. */
const around = [
  "~1 ~ ~",
  "~-1 ~ ~",
  "~ ~ ~1",
  "~ ~ ~-1",
  "~1 ~ ~1",
  "~1 ~ ~-1",
  "~-1 ~ ~1",
  "~-1 ~ ~-1",
]

/** The 8 blocks around this one and the 9 above it: never below. */
const upward = [...around, "~ ~1 ~", ...[for at in around: replace(at, " ~ ", " ~1 ")]]

/**
 * Runs as the player who chopped a log, positioned at the centre of the
 * block where the log was.
 */
Function fell {
  logs = 0
  leaf_count = 0
  east = 0
  west = 0
  south = 0
  north = 0
  top = 0
  /summon marker ~ ~ ~ {Tags:["treecap.node","treecap.origin"]}
  spread()
  if (logs > max_logs) {
    actionbar(@s, "{gray}That is too big to fell at once")
  } else if (logs > 0 && is_tree()) {
    if (break_leaves && !nether) {
      /execute as @e[type=marker,tag=treecap.node] at @s run function treecap:leaves_by_log
    }
    /execute as @e[type=marker,tag=treecap.node,tag=!treecap.origin] at @s run setblock ~ ~ ~ air destroy
    /execute as @e[type=marker,tag=treecap.leaf] at @s run setblock ~ ~ ~ air destroy
    // Gather the drops where the chopped log was, so they do not land
    // all over the canopy.
    /execute as @e[type=marker,tag=treecap.node,tag=!treecap.origin] at @s as @e[type=item,distance=..1,nbt={Age:0s}] run tp @s @e[type=marker,tag=treecap.origin,limit=1]
    /execute as @e[type=marker,tag=treecap.leaf] at @s as @e[type=item,distance=..1,nbt={Age:0s}] run tp @s @e[type=marker,tag=treecap.origin,limit=1]
    wear_axe()
  }
  /kill @e[type=marker,tag=treecap.node]
  /kill @e[type=marker,tag=treecap.leaf]
}

/**
 * Whether the logs look like a tree: natural leaves on at least 4 of
 * their 5 open sides (east, west, south, north and top, counted over all
 * the logs), or for a huge fungus, any cap block beside or above them.
 */
Function is_tree(): bool {
  let sides = east + west + south + north + top
  if (nether) {
    return sides >= 1
  }
  return sides >= 4
}

/**
 * Runs at the centre of a log, or of the chopped log's space. Visits the
 * logs of the same wood beside it and above it: the 8 around it and the
 * 9 above, but never below, so the stump stays when a trunk is chopped
 * part way up.
 */
Function spread {
  for at in upward {
    /execute positioned ${at} if block ~ ~ ~ #${kind} unless entity @e[type=marker,tag=treecap.node,distance=..0.5] run function treecap:visit
  }
}

/**
 * Runs at the centre of a log that has not been visited yet: marks it,
 * notes which of its sides touch leaves and carries on from it. The calls
 * go through raw /function commands, so the flood fill keeps no stack:
 * the counts are Globals that every visit shares.
 */
Function visit {
  /summon marker ~ ~ ~ {Tags:["treecap.node"]}
  logs += 1
  if (logs > max_logs) {
    return
  }
  if (nether) {
    // One cap block is enough for a fungus, so the checks at the
    // stem's corners, where wart blocks also grow, just set top.
    /execute if block ~1 ~ ~ #treecap:caps run scoreboard players set $east mms.treecap 1
    /execute if block ~-1 ~ ~ #treecap:caps run scoreboard players set $west mms.treecap 1
    /execute if block ~ ~ ~1 #treecap:caps run scoreboard players set $south mms.treecap 1
    /execute if block ~ ~ ~-1 #treecap:caps run scoreboard players set $north mms.treecap 1
    /execute if block ~ ~1 ~ #treecap:caps run scoreboard players set $top mms.treecap 1
    /execute if block ~1 ~ ~1 #treecap:caps run scoreboard players set $top mms.treecap 1
    /execute if block ~1 ~ ~-1 #treecap:caps run scoreboard players set $top mms.treecap 1
    /execute if block ~-1 ~ ~1 #treecap:caps run scoreboard players set $top mms.treecap 1
    /execute if block ~-1 ~ ~-1 #treecap:caps run scoreboard players set $top mms.treecap 1
  } else {
    /execute if block ~1 ~ ~ #minecraft:leaves[persistent=false] run scoreboard players set $east mms.treecap 1
    /execute if block ~-1 ~ ~ #minecraft:leaves[persistent=false] run scoreboard players set $west mms.treecap 1
    /execute if block ~ ~ ~1 #minecraft:leaves[persistent=false] run scoreboard players set $south mms.treecap 1
    /execute if block ~ ~ ~-1 #minecraft:leaves[persistent=false] run scoreboard players set $north mms.treecap 1
    /execute if block ~ ~1 ~ #minecraft:leaves[persistent=false] run scoreboard players set $top mms.treecap 1
  }
  /function treecap:spread
}

// The tree's own leaves. A natural leaf block's distance is how many
// steps it is from the nearest log, up to 7. Starting from leaves at
// distance 1 next to the felled logs and only stepping to leaves exactly
// one further away keeps to this tree's leaves, and stops at a
// neighbouring tree's, whose distances fall again toward its own logs.

/** Runs as and at a log's marker: starts a leaf fill from each of its six faces. */
Function leaves_by_log {
  for at in faces {
    /execute positioned ${at} if block ~ ~ ~ #minecraft:leaves[persistent=false,distance=1] unless entity @e[type=marker,tag=treecap.leaf,distance=..0.5] summon marker run function treecap:leaf
  }
}

/**
 * Runs as a new marker at the centre of a leaf block of this tree: marks
 * it, notes its distance and carries on to the leaves one step further.
 */
Function leaf {
  /tag @s add treecap.leaf
  leaf_count += 1
  if (leaf_count > max_leaves) {
    return
  }
  for d in [1, 2, 3, 4, 5, 6] {
    /execute if block ~ ~ ~ #minecraft:leaves[distance=${d}] run scoreboard players set @s leaf_dist ${d}
  }
  /execute if block ~ ~ ~ #minecraft:leaves[distance=7] run return 0
  for at in faces {
    /execute positioned ${at} if block ~ ~ ~ #minecraft:leaves[persistent=false] unless entity @e[type=marker,tag=treecap.leaf,distance=..0.5] if function treecap:farther summon marker run function treecap:leaf
  }
}

/**
 * Runs as a leaf marker, positioned at a neighbouring leaf block: whether
 * that block is exactly one step further from the logs.
 */
Function farther {
  for d in [1, 2, 3, 4, 5, 6] {
    /execute if score @s leaf_dist matches ${d} if block ~ ~ ~ #minecraft:leaves[distance=${d + 1}] run return 1
  }
  /return fail
}

/**
 * Takes a point of durability from the held axe for each extra log, as if
 * each had been chopped by hand. Leaves cost nothing. Unbreaking skips
 * each point with the same odds it would in vanilla: 1 in (level + 1)
 * points is taken.
 */
Function wear_axe {
  /execute if items entity @s weapon.mainhand *[minecraft:unbreakable] run return 0
  unbreaking = 0
  for level in [1, 2, 3] {
    /execute if items entity @s weapon.mainhand *[minecraft:enchantments~[{enchantments:"minecraft:unbreaking",levels:${level}}]] run scoreboard players set $unbreaking mms.treecap ${level}
  }
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
