// The Storage Barrel item. Data packs cannot add blocks, so the item is
// an item frame that looks like a barrel (item_model) and places an
// invisible frame tagged barrels.place (entity_data). The tick function
// in place.mms swaps that frame for a real barrel block and the entities
// that make up its front.

/** A barrel that holds up to 32 stacks of one kind of item. */
CustomItem storage_barrel = {
  base = item_frame
  name = "Storage Barrel"
  components = {
    item_model = "minecraft:barrel"
    entity_data = {
      id = "minecraft:item_frame"
      Tags = ["barrels.place"]
      Invisible = true
    }
  }
}

/** A barrel ringed by four chests makes a Storage Barrel. */
Recipe storage_barrel = Shaped(
  [" C ", "CBC", " C "],
  {
    C = chest
    B = barrel
  },
  storage_barrel,
)

/** What a broken Storage Barrel drops, and /loot give @s loot barrels:storage_barrel. */
LootTable storage_barrel {
  pools = [Pool(1, [Item(name = storage_barrel)])]
}

/** Data packs cannot add items to the creative menu: /function barrels:give */
Function give {
  /give @s ${storage_barrel} 4
}
