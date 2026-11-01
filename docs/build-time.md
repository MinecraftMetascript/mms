# Build-time loops and values

Some data packs repeat one shape many times: a tag per biome, a function per shrine, a dimension per planet. MMS can write that once. A build-time loop copies its body once per value while the pack is built, so the output is ordinary JSON and `.mcfunction` files, with nothing left to run in the game.

## Constants

`Const` names a value that is known while the pack is built:

```
Const biomes = [meadow, forest, cherry_grove]
Const title = "Shrine"
Const planets = [
  { name = "ash", kind = "desert", sea = 40 },
  { name = "rime", kind = "ice", sea = 63 },
]
```

A constant is a number, a string, a bool, an id, a tag, a list, or a record (`{ key = value }`). Constants can use other constants, in any order and across files, as long as no constant depends on itself (`mms/const-cycle`). A `Const` belongs at the top of a file, not inside a loop or a function (`mms/misplaced`).

A constant can stand anywhere a value goes: a resource field, a tag's list, a command, or function code.

```
BiomeTag shrine_biomes = biomes

Function hello {
  /say ${title}
}
```

Hovering a constant, or any use of one, shows its value.

## Top-level `for`

A `for` at the top of a file repeats the declarations inside it once per element:

```
for b in biomes {
  BiomeTag "has_shrine/${b}" = [b]

  Function "shrine/enter_${b}" {
    /title @s title "${b}"
  }
}
```

This writes three tags and three functions. The name of a declaration inside a loop is a string, and `${...}` in it is filled in for each copy. Two copies with the same name are an error (`mms/duplicate`) that says which values made them.

`for i, x in xs` also binds `i`, the position from 0. A loop can go over a list, a range such as `1..8` (both ends included), or a record's keys with `Keys(r)`. Loops can nest.

### `if` and `let`

`if` keeps or drops declarations by a build-time condition, and `let` names a value for the rest of its block:

```
for i, p in planets {
  let label = "${p.name} (${i})"
  if (p.kind == "ice") {
    Function "cold/${p.name}" {
      /say ${label} sea ${p.sea + 1}
    }
  } else {
    Function "warm/${p.name}" {
      /say ${label}
    }
  }
}
```

### Assertions

`Assert` stops the build with its message when a condition is false:

```
Assert "every biome is listed once" = Unique(biomes)
Assert "at most 16 planets" = Len(planets) <= 16
```

### Helpers

| Helper | Gives |
| --- | --- |
| `Len(x)` | the length of a list, record or string |
| `Keys(r)` | a record's keys, as a list of strings |
| `Unique(xs)` | `true` when no element appears twice |
| `Concat(xs, ys, ...)` | the lists joined in order |
| `Rand("key")` | a seeded generator; see below |

Lists index from 0 (`biomes[0]`) and records use `.` (`p.name`). A misspelled field gets a suggestion.

## The same `for` in functions

There is one `for`. Inside a function, what it does depends on what it loops over:

- A loop over a list or a constant is **unrolled**: the body is written once per element, and the loop variable is a build-time value, so it can go in `${...}` in a command.
- A loop over a range such as `1..n` stays a **scoreboard loop** that runs in the game, as described in [Functions](functions.md#control-flow).

```
Const effects = [speed, haste, regeneration]

Function buff_all {
  for e in effects {
    /effect give @a minecraft:${e} 30 1
  }
  for i in 1..10 {   // runs in the game
    /say ${i}
  }
}
```

`break` and `continue` work in both. Code that needs a value only the game knows, such as `effects[n]` where `n` is a score, is an error (`mms/not-constant`), because a build-time list does not exist in the game.

## Seeded randomness

`Rand("key")` makes a generator. Its draws are the same in every build, on every machine, and depend on the namespace, so two packs that share code still differ. There is no seed setting: change the key to get other values.

```
for i in 1..4 {
  let r = Rand("planet/${i}")
  Function "planet/${i}/info" {
    /say sea level ${r.Int("sea", 40, 80)}
    /say ${r.Pick("sky", [red, green, blue])} sky
  }
}
```

| Method | Gives |
| --- | --- |
| `Int(lo, hi)` | a whole number from `lo` to `hi`, both included |
| `Float(lo, hi)` | a number from `lo` up to, not including, `hi` |
| `Chance(p)` | `true` with probability `p` |
| `Pick(xs)` | one element of a list |
| `Weighted(pairs)` | one element by weight: `[[x, 3], [y, 1]]` |
| `Sample(xs, n)` | `n` distinct elements of a list |
| `Shuffle(xs)` | the list in a random order |
| `Child("label")` | a new generator for a part, keyed under this one |

Every method can take a label first, as in `r.Int("sea", 40, 80)`. A labeled draw depends only on its key and label, so adding, removing or moving other draws never changes it. Without a label, draws are numbered in the order they appear in each declaration, which is shorter to write but shifts when draws are added above. Use labels for anything you want to keep stable.

## Limits

To keep a typo from producing a million files, the build warns when it goes past these and fails at twice them:

| Limit | Warns above |
| --- | --- |
| iterations of one top-level loop | 4,096 |
| declarations made by loops | 10,000 |
| commands from one unrolled function loop | 256 |

## Editor support

The editor checks each copy and reports a problem once on the template. When only some copies have it, the message says which, as in "(when b = forest, and 2 more)". Hovering a loop variable or a `let` shows its values, and outlines list the declarations a loop makes.
