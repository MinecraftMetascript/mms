// Noticing that a player chopped a log with an axe.
//
// No advancement trigger fires when a player breaks a block, but the
// game counts every block a player mines in a statistic, one per block,
// and a scoreboard objective can watch a statistic. So there is one
// objective per kind of log that grows on a tree. The drop tells us
// where the log was: a broken block's item appears inside the block's
// space, and the tick function runs before items move for the first time.

/** A wood that grows on a tree. */
Type Wood {
  /** Names its objective: chop_<name>. */
  name: string
  /** The log block, whose mined statistic the objective watches. */
  log: Id<block>
  /** The block tag of every log of this wood, which the fill follows. */
  tag: string
  /** Whether it is the stem of a huge fungus, capped with wart blocks. */
  nether: bool = false
}

/** Every wood that grows on a tree. */
const woods: [Wood] = [
  { name = "oak", log = oak_log, tag = "minecraft:oak_logs" },
  { name = "spruce", log = spruce_log, tag = "minecraft:spruce_logs" },
  { name = "birch", log = birch_log, tag = "minecraft:birch_logs" },
  { name = "jungle", log = jungle_log, tag = "minecraft:jungle_logs" },
  { name = "acacia", log = acacia_log, tag = "minecraft:acacia_logs" },
  { name = "dark_oak", log = dark_oak_log, tag = "minecraft:dark_oak_logs" },
  { name = "mangrove", log = mangrove_log, tag = "minecraft:mangrove_logs" },
  { name = "cherry", log = cherry_log, tag = "minecraft:cherry_logs" },
  { name = "pale_oak", log = pale_oak_log, tag = "minecraft:pale_oak_logs" },
  { name = "crimson", log = crimson_stem, tag = "minecraft:crimson_stems", nether = true },
  { name = "warped", log = warped_stem, tag = "minecraft:warped_stems", nether = true },
]

for w in woods {
  /** How many ${w.log} blocks the player has mined since the last tick. */
  Objective chop_${w.name}: "minecraft.mined:minecraft.${w.log}"
}

/** Whether the player is sneaking. Sneaking chops one log, as in vanilla. */
Predicate sneaking = EntityProperties(entity = this, predicate = { flags = { is_sneaking = true } })

OnTick {
  // Each kind of log passes on its block tag, so that only logs of the
  // same wood fall: an oak growing into a birch leaves the birch alone.
  for w in woods {
    As(@a[scores={ "chop_${w.name}" = 1.. }]).At(@s) { chopped(w.tag, w.nether) }
  }
}

/**
 * Runs as and at a player who mined a log since the last tick. wood is the
 * block tag of that kind of log, such as minecraft:oak_logs, and nether
 * is true for the stems of huge fungi.
 */
Function chopped(wood: string, nether_tree: bool) {
  // Reset every count, so a player who mined two kinds of log in one
  // tick is handled once.
  for w in woods {
    /scoreboard players reset @s chop_${w.name}
  }
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
