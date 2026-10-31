// Crafting recipes. Each name becomes treasure:<name> in data/treasure/recipe/.

/** Four of any gem make a gem block. `#gems` is the tag in tags.mms. */
Recipe gem_block = Shaped(["GG", "GG"], { G = #gems }, emerald_block)

/** Break a gem block back down. */
Recipe gems_from_block = Shapeless([emerald_block], emerald * 4)

/** A compass that points to treasure. */
Recipe treasure_compass = Shaped(
  [" G ", "GCG", " G "],
  {
    G = gold_ingot
    C = compass
  },
  recovery_compass,
)
