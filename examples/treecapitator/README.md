# Treecapitator

Chop a tree's log with an axe and every log joined to it, at its height or above, falls too. The drops gather where the chopped log was, and the axe wears as if each log had been chopped by hand.

## Try it

```
mms build --world ~/.minecraft/saves/MyWorld
```

Then chop a tree with any axe in survival.

- **Sneak** while chopping to take one log, as in vanilla.
- **Chop part way up** a trunk and the stump below stays.
- **Only trees fall.** The logs must touch at least 3 natural leaves (or, for huge fungi, wart blocks or shroomlights). Leaves a player places never count, so a log cabin stays standing.
- **Big trees** fall too: up to 256 logs, enough for a 2x2 jungle or spruce tree. A bigger group of logs is left alone, with a message on the action bar.
- **The axe** loses a point of durability per extra log. Unbreaking skips points with its vanilla odds, and an axe that runs out breaks.
- **Leaves** are left in place and decay on their own, dropping saplings and apples as usual.

## How it works

| File | What it does |
| --- | --- |
| `detect.mms` | Notices a player who mined a log with an axe, and finds where the log was. |
| `fell.mms` | The flood fill, the felling and the axe's wear. |

**Noticing the chop.** No advancement trigger fires when a block is broken, but the game keeps a "mined" statistic for every block, and an objective can watch one (`Objective chop_oak = minecraft.mined:minecraft.oak_log`). There is one per log that grows on trees. The tick function looks for players with a count, resets it, and checks that they hold something in `#minecraft:axes` and are not sneaking.

**Finding the log.** A broken block's drop appears inside the block's space, and the tick function runs before new items move. So the nearest new log item (`Age:0s`) near the player, aligned to its block, is where the log was.

**The flood fill.** `spread` looks at the 8 blocks around a position and the 9 above it, never below. Each log there without a marker gets one and is visited: `visit` counts the natural leaves (`#minecraft:leaves[persistent=false]`) on its six faces and spreads from it. The recursion goes through raw `/function` commands, so it keeps no stack; the counts are Globals every visit shares.

**Felling.** If the logs look like a tree, every marked log is broken with `setblock ~ ~ ~ air destroy`, which drops it and plays the breaking effects. The new items next to each marker are moved to the chopped log's place, and the markers are removed.

**Wearing the axe.** The axe's damage is read from a copy on a scratch item display. The new damage is written back with an inline `set_components` item modifier, a macro line that MMS writes from `${damage}`. If the `damage` predicate then finds no durability left, the axe is replaced with air and the break sound plays.

## Limits

- **Creative mode** does nothing: the game only counts mined blocks in survival and adventure.
- **Branches that dip** below the log they grow from are left behind, since the fill never goes down.
- **Trees that touch** are felled together when their logs touch.
- **A log structure** joined to a real tree's logs, and touching its leaves, falls with it.
- **Hunger** is not charged for the extra logs.
- **Command limit.** A 256-log tree runs about 9,000 commands in one tick, well under the default `maxCommandChainLength` of 65,536. Raising `max_logs` (`/scoreboard players set $max_logs mms.treecap 512`) may need that gamerule raised too.
