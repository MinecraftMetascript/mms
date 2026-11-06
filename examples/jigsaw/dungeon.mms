// An underground dungeon grown from jigsaw pieces. Each piece is a
// structure file under structures/, saved with a structure block, and
// each doorway in it is a jigsaw block. A jigsaw block names the pool
// the next piece comes from and the target: the name of a jigsaw block
// in that piece to join onto.
//
//   start    four room_exit doors   -> halls, joining at hall_entry
//   halls    hall_exit at the far end -> rooms, joining at room_entry
//   rooms    room_exit doors        -> halls again
//   caps     walls that close a doorway when nothing else fits
//
// JigsawStructure writes the structure set, the structure and a template
// pool for each entry of pools. Piece("halls/hall", 3) is the structure
// file structures/halls/hall.nbt with weight 3. Put the cursor in the
// dungeon below and open the Preview tab to see which pieces join which.
//
// Try it: make a new world with the pack, then /locate structure dungeon:dungeon.

/** About one dungeon per 32 by 32 chunks, six pieces deep, 20 blocks above the bottom of the world. */
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
    /** A hall leads on to a room; the short one is rarer. */
    halls = {
      fallback = caps
      elements = [
        /** @label Long hall */
        Piece("halls/hall", 3),
        /** @label Short hall */
        Piece("halls/short_hall"),
      ]
    }
    /** A room has three more doors; a vault is a dead end on purpose. */
    rooms = { fallback = caps, elements = [Piece("rooms/room", 2), Piece("rooms/vault")] }
    /**
     * Walls that close a doorway, tried when no hall or room fits there or
     * the dungeon is six pieces deep. One of each, so both kinds of doorway
     * can be closed.
     */
    caps = [Piece("caps/hall_cap"), Piece("caps/room_cap")]
  }
}

// The names the pieces' jigsaw blocks join by. Declaring them lists them
// in the editor's Jigsaw view, and the jigsaw-hole lint reports a jigsaw
// block whose name or target is not one of them, such as a typo made in
// the game's jigsaw screen.

/** The doorways of the start room and of rooms, leading on to a hall. */
JigsawHole room_exit
/** The near end of a hall, where it joins a room_exit. */
JigsawHole hall_entry
/** The far end of a hall, leading on to a room. */
JigsawHole hall_exit
/** The doorway a room joins a hall_exit by. */
JigsawHole room_entry
