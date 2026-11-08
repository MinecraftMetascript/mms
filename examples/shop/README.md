# Shop

A shop where players spend coins. Each thing for sale is an `Offer`, a kind of declaration this project adds with a `Blueprint`, so writing a new one is four lines and the editor checks every field.

## Try it

```
mms build --world ~/.minecraft/saves/MyWorld
```

Every player starts with 50 coins. Then, in the world:

- `/function shop:bread` buys a stack of bread for 5 coins.
- `/function shop:pickaxe` buys an iron pickaxe for 40, once per player.
- `/function shop:potion` buys a potion of healing for 15.
- `/scoreboard players add @s coins 100` gives you more to spend.

Run the tests with `mms test`, or open the example in the web editor and look at the Tests tab.

## How it works

| File | What it does |
| --- | --- |
| `offer.mms` | The `Offer` Blueprint: its fields, its checks, and the function it writes. |
| `offers.mms` | What the shop sells, and the coins players start with. |
| `tests.mms` | Tests of buying, limits and running out of coins, run without the game. |

**A kind of declaration.** `Blueprint Offer` declares the fields an offer has. `price` and `reward` are required, and leaving one out is an error with a fix that adds it. `limit` and `thanks` have defaults. Completion and hover on `Offer bread { ... }` show the fields and their doc comments, as for the game's own kinds.

**Its own rules.** The `check` block runs on every offer as you type: a price below one coin is an error on the price, and one over a thousand is a warning named `pricey`, which `// mms-ignore pricey` or the `lint` settings in `mms.json` can turn off.

**What it writes.** The `emit` block writes an objective that counts each player's purchases of that offer, and the function that buys it. The function is `primary`, so the offer's name means that function: `/function shop:bread` runs it, and so does `bread()` in another function. Hover on an offer's name, or open the Expanded tab in the web editor, to see what it wrote.

## Tests

There are no players in a test, so an armor stand tagged `shopper` does the shopping: it has a coins score like a player, and the offers run as it with `As(wallet) { ... }`. Each entity keeps its own scores, so the last test checks that two shoppers pay for themselves.
