# Jigsaw dungeon

An underground dungeon of stone-brick rooms and halls, grown by the game's jigsaw placement from the pieces in `structures/`.

| Pool | Pieces | Doorways (jigsaw blocks) |
| --- | --- | --- |
| `start` | `start` | four `room_exit`, drawing from `halls` and targeting `hall_entry` |
| `halls` | `halls/hall` (weight 3), `halls/short_hall` (1) | `hall_entry` to join on, and a `hall_exit` drawing from `rooms` and targeting `room_entry` |
| `rooms` | `rooms/room` (2), `rooms/vault` (1) | `room_entry` to join on; a room has three more `room_exit` doors, a vault none |
| `caps` | `caps/hall_cap`, `caps/room_cap` | one jigsaw named `hall_entry` or `room_entry`, which turns into stone bricks |

`caps` is the fallback of `halls` and `rooms`: when no hall or room fits at a doorway, or the dungeon is `size` (6) pieces deep, the game puts a cap there instead and the doorway is walled up.

Put the cursor in a `TemplatePool` or in `Structure dungeon` and open the Preview tab (in VS Code, click Preview above it). The preview draws each pool as a column by depth, with a line from each piece to the pieces that can join its jigsaw blocks. Try changing a `target` in a piece, or renaming a pool, to see dead ends turn red.

Try it in a world:

```
mms build --world ~/.minecraft/saves/MyWorld
```

Then make new chunks and run `/locate structure dungeon:dungeon`.

The structure files are plain hollow boxes written by a script, not saved from the game, so they have not been tried in a world yet. To make your own pieces, save them with structure blocks and import them from the VS Code World view.
