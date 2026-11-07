/** Gems that drop from the gem chest. */
ItemTag gems = [diamond, emerald, amethyst_shard]

LootTable chests/gems {
  pools = [Pool(1..3, #gems)]
}

Recipe gem_block = Shaped(["GG", "GG"], { G = #gems }, emerald_block)

Advancement found {
  criteria = { got = InventoryChanged(items = [{ items = #gems }]) }
  rewards = { function = run {
    /say found it
  } }
}
