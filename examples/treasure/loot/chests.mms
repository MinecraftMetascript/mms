// Loot tables. Names can have folders: this is treasure:chests/buried.
//
// A loot table does nothing by itself. Something has to roll it:
// - chests/buried fills the chests in the wrecks from worldgen.mms, and the
//   chests that hide_chest places (see chests.mms). A chest rolls it the
//   first time a player opens it.
// - chests/storm_bonus is rolled by found_gem in progress.mms, which gives
//   its items straight to the player with /loot.

/** Rolled when a player first opens a wreck chest or a chest from hide_chest. */
LootTable chests/buried {
  pools = [
    // One to three entries from the gem tag, each equally likely.
    Pool(1..3, #gems),
    Pool(
      1,
      [
        Item(name = gold_ingot).Weight(5),
        Item(name = golden_apple).Weight(1),
        Item(name = enchanted_book).Weight(2),
      ],
    ),
  ]
}

/** Rolled by found_gem. Gives a trident, but only on a stormy night. */
LootTable chests/storm_bonus {
  pools = [
    {
      rolls = 1
      condition = stormy_night
      entries = [Item(name = trident)]
    },
  ]
}
