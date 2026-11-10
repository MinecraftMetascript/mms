# Building jigsaw pieces

A jigsaw structure is only as good as the agreement between its pieces. The game joins two pieces when their jigsaw blocks face each other and their names match, and checks nothing else: a doorway one block too high, a floor at the wrong level or a wall of the wrong stone still joins, and looks wrong. Write the agreement down once, as a short spec, and anyone can build pieces to it: you, other people, or several AI agents at once, each building a few pieces in parallel. Pieces built to the same spec fit together the first time.

This page is a template for that spec. Copy it, fill it in, and hand the same copy to everyone who builds pieces. The [`jigsaw`](../../examples/jigsaw/) example is a small dungeon built this way.

## Why names make variety

Each jigsaw block has a **name**, the name it is joined by, and a **target** and **pool**: the next piece comes from the pool, and joins by its jigsaw block named the target. So any piece with a `hall_entry` block can join any `room_exit` that draws from the `halls` pool. Nothing else needs to know about it.

That is what makes variety cheap. Keep the set of names small and fixed, and a new piece is one more file in a pool: the more pieces share the same names, the more layouts the game can grow. Decide the names before anyone builds, and declare each one as a [`JigsawHole`](resources.md#jigsaw-holes), so that a typo in the game's jigsaw screen is a `jigsaw-unknown-name` warning instead of a structure that quietly stops growing:

```mms
/** The doorways of a room, leading on to a hall. */
JigsawHole room_exit
/** The near end of a hall, where it joins a room_exit. */
JigsawHole hall_entry
```

## The spec

Fill in each section. Everything in angle brackets is yours to choose; the rest is a starting point that works.

```
# <Structure name>: piece spec

Namespace: <ns>            Structure: <ns>:<structure>
Pieces are built in a flat creative world with the data pack loaded.

## Size
- The grid is <7> blocks. Every piece's footprint is a whole number of
  cells: 7x7, 7x14, 14x14. Height is <6> blocks per storey.
- The piece's floor is its bottom layer (y = 0 in the structure block).
  Ground level outside is <y = 0, the floor is flush with it>.
- No piece is bigger than 48x48x48, the most a structure block saves.

## Doorways and jigsaw blocks
- A doorway is <3> wide and <4> high, centred on a cell's edge, with its
  sill on the floor.
- The jigsaw block sits in the outermost layer of the piece, in the
  middle of the doorway's bottom row, facing out of the piece.
- Its final state is <minecraft:air> for an open doorway; a cap's is
  <minecraft:stone_bricks>, so a closed doorway becomes wall.
- Names (each one is a JigsawHole in the project):

  | Name        | On                   | Pool  | Target      |
  | ----------- | -------------------- | ----- | ----------- |
  | room_exit   | rooms, each doorway  | halls | hall_entry  |
  | hall_entry  | halls, near end      | (none, it is joined onto) |
  | hall_exit   | halls, far end       | rooms | room_entry  |
  | room_entry  | rooms, one doorway   | (none) |            |

- A piece that is joined onto has exactly one block with its entry name.
  Every other doorway is an exit.

## Pools and pieces
  | Pool  | Fallback | Pieces (file, weight)                  | Builder |
  | ----- | -------- | -------------------------------------- | ------- |
  | start | none     | start (1)                              | <A>     |
  | halls | caps     | halls/hall (3), halls/short_hall (1)   | <B>     |
  | rooms | caps     | rooms/room (2), rooms/vault (1)        | <C>     |
  | caps  | none     | caps/hall_cap, caps/room_cap           | <A>     |
- Every pool that a jigsaw draws from has a cap in its fallback, so a
  doorway with nothing to join is walled up rather than left open.

## Materials
- Walls: <stone_bricks>, with <mossy_stone_bricks, cracked_stone_bricks>
  for at most a quarter of them.
- Floor: <polished_andesite>. Ceiling: <stone_brick_slab, top half>.
- Trim and doorframes: <chiseled_stone_bricks>.
- Never: <glass, wood, anything that burns>.
- Fill the parts of the piece that are not the building with structure
  void, so the piece does not carve out the terrain around it.

## Lighting
- Hostile mobs spawn where block light is 0. Every floor block a mob
  could stand on is within <7> blocks of a light source.
- Light sources: <lanterns hung from the ceiling, soul lanterns in
  vaults>. No torches on doorframes, where the next piece's wall goes.

## Containers
- Chests and barrels carry a loot table: <ns>:chests/<pool>. Each table
  named here is declared in the project before the piece is saved.

## Saving
- Save each piece with a structure block named <ns>:<file>, such as
  <ns>:halls/hall. Include entities: <off>.
- The game writes it to <world>/generated/<ns>/structure/<file>.nbt.
  Import it into the project as structures/<file>.nbt.

## Done means
- mms check passes with no warnings: jigsaw-unknown-name,
  jigsaw-unknown-pool, jigsaw-unreachable-target, jigsaw-facing and
  jigsaw-fit all clean.
- The piece's file is in its pool, so unused-asset does not report it.
- The structure's Simulation preview grows through the piece, with every
  jigsaw block green (joined) or grey (an end on purpose).
```

## Notes on each section

**Size.** A grid is what makes pieces interchangeable: if every doorway sits in the middle of a cell's edge, any two pieces meet doorway to doorway, and walls line up. Pick the cell size from the narrowest piece you want, such as a corridor. Pieces taller than one storey, such as a tower, are still a whole number of storeys.

**Jigsaw blocks.** The game places the next piece so that its jigsaw block is the block in front of this one, so keep jigsaw blocks in the outer layer facing out, and the two pieces meet without overlapping. A jigsaw block inside its own piece is an *internal* one, and only pieces that fit inside the parent are placed there; `jigsaw-fit` reports one whose pool has pieces too big. The game turns pieces only about the y axis, so a jigsaw facing up joins only one facing down (`jigsaw-facing`). For those, the joint is `rollable` (the next piece may turn) or `aligned` (it keeps the turn of this one).

**Pools.** Writing down who builds which piece lets several builders work at once without two of them making the same file. The pools and weights become the `pools` of the [`JigsawStructure`](https://minecraftmetascript.github.io/mms/docs/reference/kinds/JigsawStructure), with a `Piece` for each file. A `@label` doc comment on a `Piece` names it in the editor's preview.

**Materials and lighting.** These are the parts the game does not check, so they are the parts pieces disagree on when the spec leaves them out. Name blocks by their ids, so there is no question which stone is meant.

**Containers.** A loot table id in a chest is not checked when the piece is saved, and a missing table gives an empty chest. Declaring the tables first, in the spec and in the project, keeps builders from inventing names. A `Rule` processor with `AppendLoot` is another way, which puts the loot table in the pool instead of in the piece: see [Jigsaw holes](resources.md#jigsaw-holes).

**Saving.** The structure block's name is the file's path, so naming it `<ns>:halls/hall` puts it where `Piece("halls/hall")` reads it. The **World** tab in the web editor and the World view in VS Code list the structures saved in a world, import them into `structures/`, and mark the ones no pool uses yet. See [Asset files](resources.md#asset-files).

**Done means.** The jigsaw lints read the pieces' `.nbt` files, so a builder can check their own pieces before anyone puts the structure together. An AI agent can run the same checks through [`mms mcp`](agents.md).
