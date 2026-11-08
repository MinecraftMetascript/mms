// Each kind and helper has one spelling. The names an alias replaces are
// errors that say which spelling to use.
MaterialRule a = Block(stone)
MaterialCondition b = Biome(plains)
SurfaceCondition c = Water(0, 0)
SurfaceCondition d = AboveWater(0, 0)
Recipe e = CraftingShapeless(ingredients = [stone], result = { id = "stone" })
Recipe f = Shapeless(ingredients = [stone], result = { id = "stone" })
LootTable g = { pools = [LootPool(1, [])] }
LootTable h = { pools = [Pool(1, [])] }
