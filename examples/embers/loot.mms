/**
 * A cache of shards: give it to yourself with
 * /loot give @s loot embers:ember_cache
 */
LootTable ember_cache {
  pools = [
    Pool(
      1,
      [
        Item(
          name = amethyst_shard,
          modifier = [
            SetCount(2..5),
            SetComponents(
              components = {
                item_model = ember_shard
                item_name = { translate = "item.embers.ember_shard" }
              },
            ),
          ],
        ),
      ],
    ),
  ]
}
