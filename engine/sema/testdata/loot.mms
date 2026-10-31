Namespace example

LootTable blocks/ruby_ore {
  type = block
  pools = [
    {
      rolls = 1
      entries = [{ type = item, name = diamond }]
    },
    Pool(2, entries = [Item(emerald)]),
  ]
}

LootTable broken {
  pols = []
  pools = [{ rolls = "x", entries = [] }]
}
