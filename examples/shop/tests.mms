// Tests of the shop. Run them with `mms test`, in the web editor's Tests
// tab, or in VS Code's Testing view. They are never written to the pack.
//
// There are no players in a test, so an armor stand tagged shopper does
// the shopping: it has coins like a player, and the offers run as it.

/** The shopper's coins. */
const wallet = @e[tag=shopper, limit=1]

Test buying_takes_the_price {
  setup {
    /summon armor_stand 0 64 0 {Tags:["shopper"]}
    coins[wallet] = 50
    As(wallet) { bread() }
  }
  expect coins[wallet] == 45
  expect bought_bread[wallet] == 1
}

Test buying_without_enough_coins_takes_nothing {
  setup {
    /summon armor_stand 0 64 0 {Tags:["shopper"]}
    coins[wallet] = 10
    As(wallet) { pickaxe() }
  }
  expect coins[wallet] == 10
  // Never bought: the score is not even set.
  expect !(bought_pickaxe[wallet] >= 1)
}

Test a_limit_stops_a_second_purchase {
  setup {
    /summon armor_stand 0 64 0 {Tags:["shopper"]}
    coins[wallet] = 100
    As(wallet) {
      pickaxe()
      pickaxe()
    }
  }
  expect coins[wallet] == 60
  expect bought_pickaxe[wallet] == 1
}

Test offers_without_a_limit_sell_again {
  setup {
    /summon armor_stand 0 64 0 {Tags:["shopper"]}
    coins[wallet] = 50
    As(wallet) {
      potion()
      potion()
      potion()
    }
  }
  expect coins[wallet] == 5
  expect bought_potion[wallet] == 3
}

Test each_shopper_pays_for_themselves {
  setup {
    /summon armor_stand 0 64 0 {Tags:["shopper", "rich"]}
    /summon armor_stand 0 64 0 {Tags:["shopper", "poor"]}
    coins[@e[tag=rich]] = 100
    coins[@e[tag=poor]] = 3
    As(@e[tag=shopper]) { bread() }
  }
  expect coins[@e[tag=rich, limit=1]] == 95
  expect coins[@e[tag=poor, limit=1]] == 3
}
