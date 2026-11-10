# Cutting repetition

Data packs repeat themselves. Sixteen colours need sixteen recipes, every mob's loot table has the same Looting rule, and a number that matters is typed into ten files. MMS has four tools for this, and they all run when the pack is compiled: the game only ever sees plain files.

| When you have | Use |
| --- | --- |
| the same value in several places | `const` |
| the same shape with different values | `Template` |
| the same declaration for each item of a list | a top-level `for` |
| a list that follows from another list | a comprehension |

This guide works through each with a small example. [Constants, templates and loops](../reference/language/reuse.md) is the reference.

## Name a value

A `const` gives a value a name. Any value will do: a number, a list, a selector, a loot pool, a piece of text.

```mms
/** Ticks a round lasts: three minutes. */
const round_length = 3 * 60 * 20
/** Everyone taking part. */
const players = @a[tag=playing]

Global timer: int = 0

Function start_round {
  timer = round_length
  tell(players, "The round lasts ${round_length / 20} seconds")
}
```

The value is written into each place that names it, so a `const` costs nothing in the game. Hover the name to see the value, and rename it everywhere with one rename.

A constant can have a type, as in `const gem_pool: LootPool = Pool(1..3, #gems)`. Without one, it takes the type of the first place it is used, and the shorthand of that place works inside it.

## Name a shape

When the values differ but the shape is the same, write a `Template`: a helper of your own, with parameters. Here every mob drop has the same Looting rule:

```mms
/** A drop of up to `max` of one item, with more for Looting. */
Template drop(item: Id<item>, max: int = 2): LootPool = Pool(1, [
  Item(name = item, modifier = [
    SetCount(0..max),
    EnchantedCountIncrease(enchantment = looting, count = 0..1),
  ]),
])

LootTable mobs/guard {
  pools = [drop(iron_nugget), drop(bread, 1)]
}

LootTable mobs/captain {
  pools = [drop(iron_ingot), drop(emerald, max = 1)]
}
```

- **Parameters are typed.** `item: Id<item>` completes item names at each call and rejects `drop(zombie)`.
- **The return type** is what the body is checked against. `LootPool` here means the body is a loot pool, so `Pool(...)` and `0..max` work as they do in a loot table.
- **Calls look like built-in helpers.** Arguments go by position or by name, and a parameter with a default can be left out.

A template is expanded where it is called, and the JSON tab of the web editor shows the finished loot table. Hover a call to see its signature and doc comment.

## Repeat a declaration

A `for` at the top of a file repeats the declarations inside it, once for each item of a list. `${...}` in a name gives each copy its own id:

```mms
const colors = [white, orange, magenta, light_blue, yellow, lime, pink, gray]

for c in colors {
  /** Unravels ${c} wool into string. */
  Recipe ${c}_wool_to_string = Shapeless([${c}_wool], string * 4)
}
```

That is eight recipes, `white_wool_to_string` to `gray_wool_to_string`. Each is an ordinary declaration: completion offers it, go to definition lands on the loop with `c = white` shown beside it, and rename works.

The list has to be known at compile time, which is what lets the editor know every id the loop makes. Loops nest, and `for k, v in map` walks the names and values of an object.

## Build a list from a list

A comprehension makes a list out of another one, with an optional `if`. `...` spreads a list into the one around it:

```mms
const colors = [white, orange, magenta, light_blue, yellow, lime, pink, gray]

ItemTag bright_wool = [for c in colors if c != gray: ${c}_wool]

ItemTag soft_things = [
  ...[for c in colors: ${c}_wool],
  ...[for c in colors: ${c}_carpet],
  hay_block,
]
```

## Drive it from data

These tools are strongest together. Describe the things once, as a list of records, and let loops write everything that follows from them. Here each kit becomes a function:

```mms
/** A starting kit. */
Type Kit {
  name: string
  /** What the player is given. */
  items: [Id<item>]
  /** Coins it costs. */
  price: int = 0
}

const kits: [Kit] = [
  { name = "miner", items = [iron_pickaxe, torch, bread] },
  { name = "archer", items = [bow, arrow], price = 10 },
]

Objective coins

OnFirstJoin {
  coins[@s] = 0
}

for k in kits {
  /** Gives the ${k.name} kit, for ${k.price} coins. */
  Function kit/${k.name} {
    if (coins[@s] >= k.price) {
      coins[@s] -= k.price
      for item in k.items {
        /give @s ${item}
      }
    } else {
      tell(@s, "{red}The ${k.name} kit costs ${k.price} coins")
    }
  }
}
```

A third kit is one more line in `kits`. The `Type` means that line is completed and checked: a missing `name` or a misspelled item is an error on the line itself. For a player with ten coins, `/function demo:kit/archer` takes them and runs two `give` commands. Nothing of the list or the loops is left in the pack.

## Repeat commands

Inside a `Function`, a `for` over a list the compiler knows is written out, as `for item in k.items` was above. A `for` over a range, such as `for i in 1..10`, is a loop the game runs. To write a range out at compile time too, or to reuse a run of commands, put them in a `Template` that returns `Commands`:

```mms
/** A circle of particles around whoever runs it. */
Template ring(p: Id<particle>, radius: int, points: int = 12): Commands {
  for i in 0..points - 1 {
    /particle ${p} ^${radius * cos(i * 360 / points)} ^1 ^${radius * sin(i * 360 / points)}
  }
}

Function celebrate {
  As(@a).At(@s) {
    ring(flame, 2)
    ring(end_rod, 3, points = 16)
  }
}
```

`celebrate` becomes 28 `particle` commands with the positions worked out, and no loop runs in the game. The editor marks each call with an `inlined` hint, and hovering it shows the commands. See [Command templates](functions.md#command-templates).

## Keep helpers to their file

Every file sees every other file's constants and templates. A helper that only makes sense where it is written can be `private`:

```mms
private const gap = 4
private Template spaced(n: int): int = n * gap

LootTable chests/spaced {
  pools = [Pool(spaced(2), stone)]
}
```

Another file cannot name `gap`, and can have a private `gap` of its own. Only things the compiler expands away can be private: a `LootTable` or a `Function` is a file in the pack, which every file shares.

## When to stop

A loop over two items is longer than the two declarations, and harder to read. Reach for these tools when a third copy turns up, or when one number has to agree in several places.

When a group of declarations keeps appearing together, such as an item, its recipe and its cooldown, the next step is a [Blueprint](blueprints.md): a kind of declaration of your own, with fields the editor checks.
