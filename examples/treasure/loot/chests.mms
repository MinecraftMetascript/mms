// Loot tables. Names can have folders: this is treasure:chests/buried.

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

/** Extra loot only on a stormy night. */
LootTable chests/storm_bonus {
  pools = [
    {
      rolls = 1
      condition = stormy_night
      entries = [Item(name = trident)]
    },
  ]
}
