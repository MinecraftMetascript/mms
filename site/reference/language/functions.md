# Functions

A `Function` body is code that compiles to a `.mcfunction` file. It mixes raw commands, which start with `/`, with statements that compile to commands. The [Functions guide](../../guides/functions.md) explains each part at length, and [How functions compile](../../concepts/how-functions-compile.md) shows what it becomes.

```mms
/** The state of a siege, kept in storage. */
Type SiegeState {
  wave: int = 0
  boss_alive: bool = false
}

Storage game: SiegeState

const max_wave = 10
const boss = @e[tag=siege.boss, limit=1]
const wave_mobs = [zombie, skeleton, spider, creeper]

Function next_wave {
  if (game.wave >= max_wave) {
    tell(@a, "{gold}**Victory!**{/} The banner holds.")
    return
  }
  game.wave += 1
  for mob in wave_mobs {
    /execute at @e[tag=siege.spawn] run summon ${mob} ~ ~ ~ {Tags:["siege.enemy"]}
  }
  if (game.wave % 5 == 0 && !exists(boss)) {
    /summon wither_skeleton ~ ~ ~ {Tags:["siege.boss"]}
    game.boss_alive = true
  }
  title(@a, "{gold}Wave ${game.wave}", subtitle = "{gray}*hold the line*")
}
```

## Statements

| Write | Does |
| --- | --- |
| `/say hi` | a command, checked against the game's command tree |
| `let n = 3`, `let ok: bool = n > 2` | a local variable |
| `let n = /data get block ~ ~ ~ Items`, `let ok: bool = /execute if block ~ ~ ~ chest` | a command's result, or for a bool whether it succeeded |
| `const n = 3` | a value folded into every use |
| `n = 1`, `n += 1`, `-=`, `*=`, `/=`, `%=` | assignment |
| `if (c) { } else if (c) { } else { }` | a branch |
| `while (c) { }` | a loop |
| `for i in 1..10 { }` | a loop over a range, both ends included |
| `for x in [a, b, c] { }` | the body once per item, unrolled at compile time |
| `break`, `continue`, `return`, `return value` | leave a loop or the function |
| `f()`, `f(1, 2)`, `other:f()` | a call |
| `As(@a).At(@s) { }` | the block under `execute` |
| `run { }` | an inline function, as a value, such as a custom item's `on_right_click` |

## Values

Function code computes with `int` (32-bit, wrapping like a scoreboard), `bool` and `string`.

| Write | Is |
| --- | --- |
| `Global deaths: int = 0` | one value for the pack, set when it first loads |
| `Objective kills: deathCount`, `kills[@s]` | a score per holder: a selector, or a name such as `kills["#total"]` |
| `Storage game: SiegeState`, `game.wave` | a record's field in command storage |
| `let n = 0` | a local of the function |
| `@e[tag=x]`, `exists(sel)`, `count(sel)` | a selector, whether it finds any entity, and how many |

`Objective name` with no criterion is a `dummy` score, and a comma list declares several: `Objective wins, losses`.

## Parameters and results

```mms
Function classify(x: int): int {
  if (x > 100) {
    return 2
  } else if (x > 0) {
    return 1
  }
  return 0
}

Function report {
  let c = classify(150)
  /say class ${c}
}
```

Parameters are `int`, `bool` or `string`, and a function returns `int` or `bool`. A `Type` that is one of these, such as `Type Count = int` or a union of strings like `Type Mode = "easy" | "hard"`, works in their place, and in `let` and `Global` too. A function with a return type must return on every path.

## Interpolation

`${expr}` puts a value into a command. A value known at compile time is written in. A run-time value, such as a score or a parameter, turns the line into a macro, and the editor marks it with a `macro` inlay. In text, `${...}` never needs a macro: see [Text](text.md#values-in-text).

## Command templates

A `Template` that returns `Commands` writes its body where it is called. Its loops unroll and its `${...}` pieces are worked out at compile time, so it costs no function call:

```mms
/** Draws a ring of particles around the executor. */
Template particle_ring(p: Id<particle>, radius: float, points: int = 16): Commands {
  for i in 0..points - 1 {
    let a = i * 360 / points
    /particle ${p} ^${radius * cos(a)} ^1 ^${radius * sin(a)}
  }
}

Function celebrate {
  As(@a).At(@s) {
    particle_ring(flame, 2)
  }
}
```

The editor marks each call with an `inlined` inlay, and hover shows the commands it writes. `cos` and `sin` take degrees.

## Execute blocks

`As`, `At`, `Positioned`, `Rotated`, `Facing`, `In`, `Anchored`, `Align`, `On`, `Summon` and `Execute("...")` chain into an `execute` command whose `run` is the block. The [guide](../../guides/functions.md#execute-blocks) lists each one.

## Events

| Declaration | Runs |
| --- | --- |
| `OnLoad { }` | when the pack loads, and on `/reload` |
| `OnTick { }` | every tick |
| `OnFirstJoin { }` | once per player, the first time they join |
| `OnRespawn { }` | when a player respawns |
| `OnAdvancement name = Trigger(...) { }` | each time a player meets the trigger, such as `PlacedBlock(...)` |

The body runs as the player, at their position, except `OnLoad` and `OnTick`, which run as the server. The unnamed ones merge: several `OnTick` blocks in a project become one tick function. A [custom item](../../guides/resources.md#custom-items) has its own hooks, such as `on_right_click` and `on_hit`.

```mms
Objective banners

/** Counts every banner a player places. */
OnAdvancement placed_banner = PlacedBlock(location = LocationCheck(predicate = { block = { blocks = #banners } })) {
  banners[@s] += 1
}

OnRespawn {
  /effect give @s minecraft:regeneration 5 1
}
```
