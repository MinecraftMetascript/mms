# Functions

`Function` bodies and event hooks (`OnLoad`, `OnTick`, `OnFirstJoin`, `OnRespawn`, `OnAdvancement`) compile to `.mcfunction` files.

## Raw commands

A line that starts with `/` is a command. It is written to the output as it is, without the slash:

```
Function greet {
  /say hello
  /execute as @a at @s run particle minecraft:heart ~ ~2 ~
}
```

Every command is checked against the 26.3 command tree, the same one the game uses. The checker reports:

- unknown commands and options, with the alternatives
- missing and extra arguments
- argument syntax: numbers and their bounds, selectors and their options, coordinates, ids, ranges, and balanced NBT and text
- unknown vanilla ids, such as `minecraft:diamnd`
- ids in your own namespaces that nothing declares, such as `function demo:missing`

Completion inside a command offers the next literal and values for the argument being typed. These include selectors, game modes, registry ids, and your project's functions, tags and objectives.

The checker checks syntax, not meaning. Whether an entity exists, or whether an NBT path fits the entity, is left to the game.

## Values and variables

Function code computes with whole numbers (`int`, 32-bit, wrapping like the game's scoreboards) and `bool`s.

```
Global deaths: int = 0          // one value for the whole pack
Objective kills = deathCount    // a score per player

Function on_death {
  deaths += 1
  kills[@s] = 0                 // a score: objective[holder]
  let streak = kills[@s] * 2    // a local; its type is inferred
  let lucky: bool = streak > 10
}
```

- **Globals** are set once, when the pack first loads. A `/reload` keeps their values.
- **Locals** (`let`) belong to their function. A name in an inner block can shadow an outer one.
- **Objective scores** are written `objective[holder]`. The holder is a selector, or a fake player name in quotes, such as `kills["#total"]`.

### Operators

| Operators | Meaning |
|---|---|
| `+ - * / %` | Arithmetic. Division rounds down and `%` takes the divisor's sign, as the scoreboard does. |
| `< <= > >= == !=` | Comparisons. |
| `&& \|\| !` | Logic. The right side of `&&` and `\|\|` only runs when it is needed. |

Constant parts are folded at compile time.

### Strings

`string` values live in command storage (`<namespace>:mms`). They can be joined with `+`, appended with `+=`, and compared with `==` and `!=`. Ints and bools join as their text.

```
Global motd: string = "Welcome"

Function welcome(name: string) {
  let line = motd + ", " + name + "!"
  if (name == "Notch") {
    line += " Nice to meet you."
  }
  /tellraw @a "${line}"
}
```

## Interpolation

`${expr}` puts a value into a command:

```
Function report(points: int) {
  /say ${points} points, ${points * 2} doubled
  /scoreboard players set @s kills ${points}
}
```

- **Constant values** (literals, constant arithmetic) are written into the command at compile time.
- **Run-time values** turn the line into a [macro](https://minecraft.wiki/w/Function_(Java_Edition)#Macros): the values are copied into storage, and a one-line macro function runs the command.
- **Quotes:** a string with a `"` in it will break a command that wraps it in quotes, just as it would in a hand-written macro.

## Control flow

```
Function classify(x: int): int {
  if (x > 100) {
    return 2
  } else if (x > 0) {
    return 1
  }
  return 0
}

Function count {
  for i in 1..10 {          // both ends included, like 1..10 in a selector
    if (i % 2 == 0) {
      continue
    }
    /say odd
  }
  let n = 0
  while (n < 5) {
    n += 1
    if (n == 3) {
      break
    }
  }
}
```

A function with a return type must end with `return` on every path.

## Calls

A call looks like `add(1, 2)`, or `other_ns:add(1, 2)` for a function in another namespace. Arguments are passed by position. A call can stand on its own, or be used as a value when the function has a return type.

Parameters are `int`, `bool` or `string`. Functions return `int` or `bool`. To hand a string back, assign it to a Global.

## Execute blocks

A chain of execute helpers followed by a block runs the block under `execute`:

```
Objective hits

Function zap {
  As(@e[tag=tower]).At(@s) {
    /particle minecraft:flame ~ ~1 ~
    hits[@s] += 1
  }
  Positioned(~ ~1 ~) {
    /say above
  }
  Execute("if entity @e[tag=boss]").In(the_nether) {
    /say the boss is here
  }
}
```

| Helper | Subcommand |
|---|---|
| `As(targets)` | `as` |
| `At(targets)` | `at` |
| `Positioned(pos)`, `Positioned(targets)` | `positioned`, `positioned as` |
| `Rotated(rot)`, `Rotated(targets)` | `rotated`, `rotated as` |
| `Facing(pos)`, `Facing(targets, anchor)` | `facing`, `facing entity` (the anchor defaults to `eyes`) |
| `In(dimension)` | `in` |
| `Anchored(eyes)` | `anchored` |
| `Align(xyz)` | `align` |
| `On(relation)` | `on` |
| `Summon(entity)` | `summon` |
| `Execute("...")` | any subcommands, written as text |

Arguments are written into the command as they appear: selectors as `@e[tag=x]`, coordinates as `~ ~1 ~` or `0, 64, 0`, ids as `the_nether`. The chain is checked against the command tree like a raw command, and `Execute("...")` covers conditions such as `if block` and anything without a helper.

A body with one command is written inline, as `execute as @a at @s run particle ...`. A longer body becomes a sub-function, such as `demo:zap/_exec0`.

The body can use the function's variables. They are not per entity: a body that runs for five entities updates one `n` five times. Keep per-entity values in an objective, such as `hits[@s]`.

`break`, `continue` and `return` cannot leave an execute block, because the block may run once per entity. Set a variable inside it and act on it afterwards.

## Event hooks

These declarations run their body when something happens, as the player it happens to, at their position:

```
Objective banners

/** Counts every banner a player places. */
OnAdvancement placed_banner = PlacedBlock(location = LocationCheck(predicate = { block = { blocks = #banners } })) {
  banners[@s] += 1
}

OnFirstJoin {
  /give @s minecraft:bread 4
}

OnRespawn {
  /effect give @s minecraft:regeneration 5 1
}
```

- **`OnAdvancement name = Trigger(...) { ... }`** runs each time a player meets a trigger. The trigger is one of the criterion helpers, such as `PlacedBlock`, `PlayerKilledEntity`, `InventoryChanged` or `Tick`, and its named arguments are the conditions. It writes the advancement `<namespace>:<name>` and a reward function, `<namespace>:mms_gen/advancement/<name>`, which revokes the advancement and then runs the body, so it fires again next time. Commands can refer to it, as in `/advancement grant @s only demo:placed_banner`.
- **`OnFirstJoin { ... }`** runs once per player, the first time they are in the world with the pack. It uses a tick advancement, `<namespace>:mms_gen/first_join`, that is never revoked; `/advancement revoke @s only <namespace>:mms_gen/first_join` makes it run again.
- **`OnRespawn { ... }`** runs right after a player respawns. A `deathCount` objective, `mms_respawn.<namespace>`, notes the death, and the tick function runs the body once the player is alive again.

### Right-click

A [custom item](resources.md)'s `on_right_click` names a function to run when a player right-clicks with the item, as and at that player:

```
CustomItem ember_wand = {
  base = stick
  name = "Ember Wand"
  on_right_click = cast_ember
}

Function cast_ember {
  /particle flame ^ ^1.5 ^1 0.2 0.2 0.2 0.05 20
  /effect give @s fire_resistance 5
}
```

The function takes no parameters and can be in any file. It runs once per press: holding the button down does not repeat it. This uses the consumable method from the [datapack.wiki guide](https://datapack.wiki/guide/right-click/consumable_component):

- The item gets three more components wherever it is used, in `/give`, recipes and loot: a `consumable` that takes about eleven days to eat, with no animation, sound or particles, so right click "uses" it without using it up; `use_effects`, so the player is not slowed down while holding the button; and `mms_item: "<namespace>:<name>"` in its `custom_data`. Set `consumable` or `use_effects` in the item's `components` to use your own instead. A `custom_data` of your own keeps its keys and gains the marker.
- A `using_item` advancement, `<namespace>:mms_gen/right_click/<name>`, matches the base item with that marker. Its reward function, of the same name, revokes it and calls your function.
- `using_item` fires on every tick the button is held, so a score, `mms_click.<namespace>`, notes that the player is still holding it. The tick function counts it down, and your function runs only once it has reached 0.

A base item that already does something on right click keeps doing it: a block is still placed and a bow still draws. Pick a base such as `stick`, `paper` or `amethyst_shard` that does nothing on its own. A `use_cooldown` component has no effect, because the item is never finished being used; keep a cooldown in a score instead.

## How it compiles

Values live on one scoreboard objective per namespace, `mms.<namespace>`:

| Holder | Value |
|---|---|
| `$<name>` | a Global |
| `#<function>.<name>` | a parameter or local |
| `#<function>._t<n>`, `#<function>._f<n>` | temporaries and condition flags |
| `#c<n>` | the constant `n`, for multiplication, division and comparison |
| `#cf`, `#rv` | control flow, see below |

String values and macro arguments live in storage `<namespace>:mms`: `g.<name>` for globals, `l."<function>".<name>` for locals, and `m` for the arguments of the macro being called.

**Branches and loops.** `if` with a one-command body becomes `execute if ... run <command>`. Longer bodies become sub-functions next to the function, such as `demo:check/_if0`.

A loop is a function that checks its condition, runs its body and calls itself. `break` returns from it, and `continue` calls it again.

**Leaving early.** When `break`, `continue` or `return` happens inside a sub-function, it sets `#cf`, and the calling function passes it on. The return value travels in `#rv`.

**Results.** Functions return values with the `return` command, and callers read them with `execute store result score`.

**Recursion.** A call that can lead back to the calling function, directly or through other functions, pushes the caller's scores onto a stack in storage, `<namespace>:mms stack`, and pops them afterwards. Only the scores read after the call are saved. Every other call costs nothing extra.

**Optimisation.** Temporaries that only carry a value from one place to the next are removed, so `x = x + 1` becomes one `scoreboard players add`.

**Load.** The generated load function creates the objective, sets the constants, and gives each global its first value.

## Generated functions

`OnLoad { ... }` and `OnTick { ... }` bodies, and `Objective` declarations, are collected into `<namespace>:mms_gen/load` and `<namespace>:mms_gen/tick`. These are registered in `#minecraft:load` and `#minecraft:tick`. `OnFirstJoin` and `OnRespawn` bodies are collected the same way, into `<namespace>:mms_gen/first_join` and `<namespace>:mms_gen/respawn`.

A hook body that holds only commands is copied into these functions as it is. A body with variables or control flow compiles to its own function, such as `<namespace>:mms_gen/load/<hash>`, which the generated function calls.
