# Tutorial: a jigsaw dungeon

This tutorial grows a dungeon out of rooms and halls, the way the game grows villages and bastions. You build the pieces in the game, MMS wires them together, and the editor shows the layouts the game will make before you open a world.

The finished project is the [`jigsaw`](../../examples/jigsaw/) example. Its pieces are plain stone boxes, so you can also open it from **Examples** in the web editor and follow along without building anything.

## How jigsaw structures work

A jigsaw structure starts from one piece. Every piece is a saved build (a `.nbt` structure file), and its doorways are **jigsaw blocks**. A jigsaw block says three things:

- its **name**, which other blocks join it by
- a **pool**, the list of pieces the next piece is picked from
- a **target**, the name of the jigsaw block in that next piece to join onto

The game places the start piece, then for each jigsaw block picks a piece from its pool and lines up the two blocks, and so on until the pieces run out, overlap, or the structure is deep enough.

## 1. Plan the names

Decide the names before building anything, because they are typed into every piece. This dungeon has rooms and halls that alternate:

| Name | On | Draws from pool | Joins onto |
| --- | --- | --- | --- |
| `room_exit` | each doorway of a room | `halls` | `hall_entry` |
| `hall_entry` | the near end of a hall | | |
| `hall_exit` | the far end of a hall | `rooms` | `room_entry` |
| `room_entry` | one doorway of a room | | |

Set up the project and declare the names:

```sh
mkdir dungeon && cd dungeon
mms init -namespace dungeon
```

```mms
/** The doorways of the start room and of rooms, leading on to a hall. */
JigsawHole room_exit
/** The near end of a hall, where it joins a room_exit. */
JigsawHole hall_entry
/** The far end of a hall, leading on to a room. */
JigsawHole hall_exit
/** The doorway a room joins a hall_exit by. */
JigsawHole room_entry
```

A `JigsawHole` writes no file. It is there so that the editor can offer the names, and so that a typo made in the game's jigsaw screen is a warning in the editor instead of a dungeon that stops growing.

## 2. Build the pieces

Build in a flat creative world. Choose it in the editor first: the **World** tab in the web editor (Chrome and Edge), or **Choose World** in the MMS side bar in VS Code. Build the project into it once, so the game knows the `dungeon` namespace.

For each piece:

1. **Build it.** Keep to a grid, so that any two pieces meet doorway to doorway. The example uses rooms of 7×7 and halls 5 wide. [Building jigsaw pieces](../guides/jigsaw-pieces.md) is a checklist of the things worth agreeing on first: sizes, doorways, materials and lighting.
2. **Place the jigsaw blocks.** Each goes in the outer layer of the piece, in a doorway, facing out. The editor's **Jigsaw** tab (the **Jigsaw** view in VS Code) sets one up: pick the pool, the target and the name from the lists, and it copies a `/give` command for a jigsaw block that already has them. Paste it in the game and place the block.
3. **Save it.** Run `/give @s structure_block`, place it at a corner of the piece, switch it to **Save** mode, set the size to cover the piece and name it after its file: `dungeon:halls/hall`. Press **SAVE**.

The pieces of this dungeon:

| File | Jigsaw blocks |
| --- | --- |
| `start` | four `room_exit` doorways |
| `halls/hall`, `halls/short_hall` | a `hall_entry` at one end and a `hall_exit` at the other |
| `rooms/room` | a `room_entry`, and three `room_exit` doorways |
| `rooms/vault` | a `room_entry` only: a dead end on purpose |
| `caps/hall_cap`, `caps/room_cap` | one `hall_entry` or `room_entry`, turning into stone bricks |

The caps are small plugs of wall. They close a doorway when no hall or room fits there.

## 3. Bring them into the project

The game saves each piece under the world's `generated` folder. The **Structures** list, in the World tab or the VS Code side bar, shows what the world has saved and marks each one new, imported or changed. **Import** copies it to `structures/` in the project, so `dungeon:halls/hall` becomes `structures/halls/hall.nbt`.

After changing a piece in the game, save it again with its structure block and import it again.

## 4. Wire it up

One declaration turns the files into a structure:

```mms in=examples/jigsaw
JigsawStructure dungeon {
  biomes = #minecraft:is_overworld
  step = underground_structures
  spacing = 32
  separation = 12
  height = AboveBottom(20)
  depth = 6
  start = start

  pools = {
    start = [Piece("start")]
    halls = { fallback = caps, elements = [Piece("halls/hall", 3), Piece("halls/short_hall")] }
    rooms = { fallback = caps, elements = [Piece("rooms/room", 2), Piece("rooms/vault")] }
    caps = [Piece("caps/hall_cap"), Piece("caps/room_cap")]
  }
}

JigsawHole room_exit, hall_entry, hall_exit, room_entry
```

- **Where it generates.** `biomes` is the biomes it may start in, `step` the stage of world generation it belongs to, and `height` where the start piece goes. `spacing` and `separation` are in chunks: one try per 32×32 chunks, at least 12 apart.
- **How big it gets.** `depth = 6` stops the dungeon six pieces from the start.
- **The pools.** Each entry of `pools` is a pool that jigsaw blocks name. `Piece("halls/hall", 3)` is the file `structures/halls/hall.nbt` with weight 3, so long halls are three times as common as short ones.
- **Fallbacks.** When nothing in `halls` fits, or the dungeon is at its depth, the game tries the pool's `fallback`. Here that is `caps`, which walls the doorway up.

By hand this is a structure set, a structure and four template pools, each in its own file and each naming the next by id.

## 5. Look before you load

Put the cursor in the `JigsawStructure` and open the **Preview** tab (in VS Code, click **Preview** above it).

- **Connections** draws the pools and which pieces can join which. A jigsaw block with nothing to join shows red.
- **Simulation** grows one layout the way the game does, in 3D. Each jigsaw block is green (joined), red (found nothing) or grey (an end on purpose). **Re-roll** grows another, and **Slice** cuts the view open.

`mms check` reads the piece files too. The `jigsaw-` lint rules report a block whose name or pool is unknown, a target no piece in the pool has, and pieces that cannot fit. Fixing those here is much faster than hunting for a dungeon in a world and finding half of it missing.

The simulation leaves the terrain out, and a layout seed is not a world seed. It shows what kind of dungeon the pieces make, not the dungeon at one place in one world.

## 6. Find one in the game

```sh
mms build --world ~/.minecraft/saves/MyWorld
```

Structures are read when a world opens, so leave the world and open it again. Then:

- `/place structure dungeon:dungeon` builds one where you stand, which is the quickest way to see it.
- `/locate structure dungeon:dungeon` finds one that generated. Only chunks made after the pack was added have them, so travel somewhere new.

## Next steps

- Add a room. A new piece is one more file and one more `Piece(...)` in a pool, and every layout can use it.
- Fill the chests: a chest saved in a piece can name a loot table, and the `structure-unknown-loot-table` lint checks that the table exists.
- [Structures](../guides/resources.md#structures) in the resources guide covers processors, `spawn_overrides`, pool aliases and the other fields.
- The [`treasure`](../../examples/treasure/) example reuses a vanilla template with its own loot.
