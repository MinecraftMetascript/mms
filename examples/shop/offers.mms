// What the shop sells.

Function give_bread {
  /give @s bread 8
}

Function give_pick {
  /give @s iron_pickaxe
}

Function give_potion {
  /give @s potion[potion_contents={potion:"minecraft:healing"}]
}

/** A stack of bread for the road. */
Offer bread {
  price = 5
  reward = give_bread
}

/** One iron pickaxe per player, at a discount. */
Offer pickaxe {
  price = 40
  reward = give_pick
  limit = 1
  thanks = "Mind the lava"
}

/** A potion of healing. */
Offer potion {
  price = 15
  reward = give_potion
}

/** Every player starts with some coins. */
OnFirstJoin {
  coins[@s] = 50
}
