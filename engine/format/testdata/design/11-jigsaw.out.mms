JigsawHole room_exit, hall_entry, hall_exit, room_entry

/** About one dungeon per 32 by 32 chunks, six pieces deep. */
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
    halls = Pool(
      fallback = caps,
      /** @label Long hall */
      Piece("halls/hall", 3),
      /** @label Short hall */
      Piece("halls/short_hall"),
    )
    rooms = Pool(fallback = caps, Piece("rooms/room", 2), Piece("rooms/vault"))
    caps = [Piece("caps/hall_cap"), Piece("caps/room_cap")]
  }
}
