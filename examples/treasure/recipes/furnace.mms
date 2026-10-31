// Cooking recipes, written as helpers named after the recipe type.

Recipe smelt_raw_gold = Smelting(
  ingredient = raw_gold_block,
  result = { id = gold_block },
  experience = 2,
  cookingtime = 400,
)

Recipe blast_raw_gold = Blasting(
  ingredient = raw_gold_block,
  result = { id = gold_block },
  experience = 2,
  cookingtime = 200,
)

Recipe cut_bricks = Stonecutting(ingredient = stone, result = stone_bricks * 2)
