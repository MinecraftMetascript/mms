// Noticing that a player chopped a log with an axe.
//
// No advancement trigger fires when a player breaks a block, but the
// game counts every block a player mines in a statistic, one per block,
// and a scoreboard objective can watch a statistic. So there is one
// objective per kind of log that grows on a tree. The drop tells us
// where the log was: a broken block's item appears inside the block's
// space, and the tick function runs before items move for the first time.

Objective chop_oak: minecraft.mined:minecraft.oak_log
Objective chop_spruce: minecraft.mined:minecraft.spruce_log
Objective chop_birch: minecraft.mined:minecraft.birch_log
Objective chop_jungle: minecraft.mined:minecraft.jungle_log
Objective chop_acacia: minecraft.mined:minecraft.acacia_log
Objective chop_dark_oak: minecraft.mined:minecraft.dark_oak_log
Objective chop_mangrove: minecraft.mined:minecraft.mangrove_log
Objective chop_cherry: minecraft.mined:minecraft.cherry_log
Objective chop_pale_oak: minecraft.mined:minecraft.pale_oak_log
Objective chop_crimson: minecraft.mined:minecraft.crimson_stem
Objective chop_warped: minecraft.mined:minecraft.warped_stem

/** Whether the player is sneaking. Sneaking chops one log, as in vanilla. */
Predicate sneaking = EntityProperties(entity = this, predicate = { flags = { is_sneaking = true } })

OnTick {
  // Each kind of log passes on its block tag, so that only logs of the
  // same wood fall: an oak growing into a birch leaves the birch alone.
  As(@a[scores={ chop_oak = 1.. }]).At(@s) { chopped("minecraft:oak_logs", false) }
  As(@a[scores={ chop_spruce = 1.. }]).At(@s) { chopped("minecraft:spruce_logs", false) }
  As(@a[scores={ chop_birch = 1.. }]).At(@s) { chopped("minecraft:birch_logs", false) }
  As(@a[scores={ chop_jungle = 1.. }]).At(@s) { chopped("minecraft:jungle_logs", false) }
  As(@a[scores={ chop_acacia = 1.. }]).At(@s) { chopped("minecraft:acacia_logs", false) }
  As(@a[scores={ chop_dark_oak = 1.. }]).At(@s) { chopped("minecraft:dark_oak_logs", false) }
  As(@a[scores={ chop_mangrove = 1.. }]).At(@s) { chopped("minecraft:mangrove_logs", false) }
  As(@a[scores={ chop_cherry = 1.. }]).At(@s) { chopped("minecraft:cherry_logs", false) }
  As(@a[scores={ chop_pale_oak = 1.. }]).At(@s) { chopped("minecraft:pale_oak_logs", false) }
  As(@a[scores={ chop_crimson = 1.. }]).At(@s) { chopped("minecraft:crimson_stems", true) }
  As(@a[scores={ chop_warped = 1.. }]).At(@s) { chopped("minecraft:warped_stems", true) }
}

/**
 * Runs as and at a player who mined a log since the last tick. wood is the
 * block tag of that kind of log, such as minecraft:oak_logs, and nether
 * is true for the stems of huge fungi.
 */
Function chopped(wood: string, nether_tree: bool) {
  // Reset every count, so a player who mined two kinds of log in one
  // tick is handled once.
  /scoreboard players reset @s chop_oak
  /scoreboard players reset @s chop_spruce
  /scoreboard players reset @s chop_birch
  /scoreboard players reset @s chop_jungle
  /scoreboard players reset @s chop_acacia
  /scoreboard players reset @s chop_dark_oak
  /scoreboard players reset @s chop_mangrove
  /scoreboard players reset @s chop_cherry
  /scoreboard players reset @s chop_pale_oak
  /scoreboard players reset @s chop_crimson
  /scoreboard players reset @s chop_warped
  /execute if predicate treecap:sneaking run return 0
  // The axe is still in hand, unless that log used up its last point of
  // durability.
  /execute unless items entity @s weapon.mainhand #minecraft:axes run return 0
  // The log's drop is a brand new item (Age 0) near the player.
  kind = wood
  nether = nether_tree
  /execute as @e[type=item,distance=..8,nbt={Age:0s}] if items entity @s contents #${kind} run tag @s add treecap.drop
  At(@e[type=item, tag=treecap.drop, limit=1, sort=nearest])
    .Align(xyz)
    .Positioned(~0.5 ~0.5 ~0.5) { fell() }
  /tag @e[type=item,tag=treecap.drop] remove treecap.drop
}
