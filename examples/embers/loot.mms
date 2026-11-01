/**
 * A cache of shards: give it to yourself with
 * /loot give @s loot embers:ember_cache
 */
LootTable ember_cache {
  pools = [Pool(1, [Item(name = ember_shard, modifier = [SetCount(2..5)])])]
}
