// A shop where players spend coins. `Offer` is a Blueprint: a kind of
// declaration this project adds, with fields the editor completes and
// checks, rules of its own in `check`, and the function it writes in
// `emit`.

/** Coins a player has to spend. */
Objective coins

/**
 * Something for sale. Buying it runs `reward` as the buyer, once they
 * have paid. Its name is the function that buys it:
 * `/function shop:<name>`.
 */
Blueprint Offer {
  /** What it costs, in coins. */
  price: int
  /** What the buyer gets: a function run as them. */
  reward: Function
  /** How many each player can buy; 0 for no limit. */
  limit: int = 0
  /** Shown when they buy it. */
  thanks: string = "Thanks!"

  check {
    if (price < 1) {
      error(price, "an offer costs at least one coin")
    }
    if (limit < 0) {
      error(limit, "a limit is 0, for none, or more")
    }
    if (price > 1000) {
      warn(price, "pricey", "few players save over a thousand coins")
    }
  }

  emit {
    /** How many ${self.path} each player has bought. */
    Objective bought_${self.path}

    primary Function ${self.path} {
      if (limit > 0 && bought_${self.path}[@s] >= limit) {
        tell(@s, "{red}You can't buy any more of these")
        return
      }
      if (coins[@s] < price) {
        tell(@s, "{red}That costs ${price} coins")
        return
      }
      coins[@s] -= price
      bought_${self.path}[@s] += 1
      reward()
      tell(@s, "{green}${thanks}")
    }
  }
}
