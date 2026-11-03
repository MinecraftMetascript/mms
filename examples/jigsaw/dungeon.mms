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
// Put the cursor in a TemplatePool or the Structure below and open the
// Preview tab to see which pieces join which.
//
// Try it: make a new world with the pack, then /locate structure dungeon:dungeon.

/** About one dungeon per 32 by 32 chunks. */
StructureSet dungeon = {
  structures = [
    {
      structure = dungeon
      weight = 1
    },
  ]
  placement = RandomSpread(spacing = 32, separation = 12, salt = 20261102)
}

/** Grows up to six pieces deep from the start room, 20 blocks above the bottom of the world. */
Structure dungeon = Jigsaw(
  biomes = #minecraft:is_overworld,
  step = underground_structures,
  spawn_overrides = {},
  start_pool = start,
  size = 6,
  start_height = Constant({ above_bottom = 20 }),
  max_distance_from_center = 80,
  use_expansion_hack = false,
)

TemplatePool start = {
  fallback = minecraft:empty
  elements = [
    {
      weight = 1
      element = SinglePoolElement(
        projection = rigid,
        location = Structure("structures/start.nbt", "start"),
        processors = minecraft:empty,
      )
    },
  ]
}

/** A hall leads on to a room; the short one is rarer. */
TemplatePool halls = {
  fallback = caps
  elements = [
    /** @label Long hall */
    {
      weight = 3
      element = SinglePoolElement(
        projection = rigid,
        location = Structure("structures/halls/hall.nbt", "halls/hall"),
        processors = minecraft:empty,
      )
    },
    /** @label Short hall */
    {
      weight = 1
      element = SinglePoolElement(
        projection = rigid,
        location = Structure("structures/halls/short_hall.nbt", "halls/short_hall"),
        processors = minecraft:empty,
      )
    },
  ]
}

/** A room has three more doors; a vault is a dead end on purpose. */
TemplatePool rooms = {
  fallback = caps
  elements = [
    {
      weight = 2
      element = SinglePoolElement(
        projection = rigid,
        location = Structure("structures/rooms/room.nbt", "rooms/room"),
        processors = minecraft:empty,
      )
    },
    {
      weight = 1
      element = SinglePoolElement(
        projection = rigid,
        location = Structure("structures/rooms/vault.nbt", "rooms/vault"),
        processors = minecraft:empty,
      )
    },
  ]
}

/**
 * Walls that close a doorway, tried when no hall or room fits there or
 * the dungeon is six pieces deep. One of each, so both kinds of doorway
 * can be closed.
 */
TemplatePool caps = {
  fallback = minecraft:empty
  elements = [
    {
      weight = 1
      element = SinglePoolElement(
        projection = rigid,
        location = Structure("structures/caps/hall_cap.nbt", "caps/hall_cap"),
        processors = minecraft:empty,
      )
    },
    {
      weight = 1
      element = SinglePoolElement(
        projection = rigid,
        location = Structure("structures/caps/room_cap.nbt", "caps/room_cap"),
        processors = minecraft:empty,
      )
    },
  ]
}
