# Treecapitator

Chop a tree's log with an axe and every log of the same wood joined to it, at its height or above, falls too, along with the tree's leaves. The drops gather where the chopped log was, and the axe wears as if each log had been chopped by hand.

## Try it

```
mms build --world ~/.minecraft/saves/MyWorld
```

Then chop a tree with any axe in survival.

- **Sneak** while chopping to take one log, as in vanilla.
- **Chop part way up** a trunk and the stump below stays.
- **Only trees fall.** Natural leaves must touch the logs on at least 4 of their 5 open sides (east, west, south, north, top), counted over the whole tree. A huge fungus needs one wart block or shroomlight beside its stem. Leaves a player places never count, so a log cabin stays standing.
- **Only one wood.** An oak growing into a birch leaves the birch standing.
- **Big trees** fall too: up to 256 logs, enough for a 2x2 jungle or spruce tree. A bigger group of logs is left alone, with a message on the action bar.
- **The axe** loses a point of durability per extra log. Unbreaking skips points with its vanilla odds, and an axe that runs out breaks.
- **Leaves** of the tree break with it and drop saplings, sticks and apples as usual, up to 600 of them; any past that decay on their own. Set `break_leaves` to false in `fell.mms` to leave them all to decay. Fungus caps stay.

## How it works

| File | What it does |
| --- | --- |
| `detect.mms` | Notices a player who mined a log with an axe, and finds where the log was. |
| `fell.mms` | The flood fill, the felling and the axe's wear. |

**Noticing the chop.** No advancement trigger fires when a block is broken, but the game keeps a "mined" statistic for every block, and an objective can watch one (`Objective chop_oak: minecraft.mined:minecraft.oak_log`). There is one per log that grows on trees. The tick function looks for players with a count and calls `chopped` with that wood's block tag (`"minecraft:oak_logs"`), which resets the counts and checks that they hold something in `#minecraft:axes` and are not sneaking.

**Finding the log.** A broken block's drop appears inside the block's space, and the tick function runs before new items move. So the nearest new log item (`Age:0s`) near the player, aligned to its block, is where the log was.

**The flood fill.** `spread` looks at the 8 blocks around a position and the 9 above it, never below, for blocks in the wood's tag (`#${kind}`, which MMS turns into a macro line). Each log there without a marker gets one and is visited: `visit` notes which of its sides touch natural leaves (`#minecraft:leaves[persistent=false]`) and spreads from it. The recursion goes through raw `/function` commands, so it keeps no stack; the counts are Globals every visit shares.

**The leaves.** A natural leaf's `distance` state is how many steps it is from the nearest log. A second fill starts at the distance 1 leaves beside the felled logs and only steps to leaves exactly one further away, with a marker per leaf that keeps its distance in a score. That keeps to this tree's leaves: a neighbouring tree's leaves get closer to its own logs again, so the fill stops there.

**Felling.** If the logs look like a tree, every marked log and leaf is broken with `setblock ~ ~ ~ air destroy`, which drops it and plays the breaking effects. The new items next to each marker are moved to the chopped log's place, and the markers are removed.

**Wearing the axe.** The axe's damage is read from a copy on a scratch item display. The new damage is written back with an inline `set_components` item modifier, a macro line that MMS writes from `${damage}`. If the `damage` predicate then finds no durability left, the axe is replaced with air and the break sound plays.

## Limits

- **Creative mode** does nothing: the game only counts mined blocks in survival and adventure.
- **Branches that dip** below the log they grow from are left behind, since the fill never goes down.
- **Trees of the same wood** are felled together when their logs touch.
- **A log structure** of the tree's wood, joined to its logs, falls with it.
- **Hunger** is not charged for the extra logs, and leaves cost the axe nothing.
- **Command limit.** At the limits, 256 logs and 600 leaves, a chop runs about 35,000 commands in one tick, under the default `maxCommandChainLength` of 65,536. Raising `max_logs` or `max_leaves` (`/scoreboard players set $max_leaves mms.treecap 1200`) may need that gamerule raised too.
