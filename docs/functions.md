# Functions

`Function`, `OnLoad` and `OnTick` bodies compile to `.mcfunction` files.

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

Parameters are `int` or `bool`.

## How it compiles

Values live on one scoreboard objective per namespace, `mms.<namespace>`:

| Holder | Value |
|---|---|
| `$<name>` | a Global |
| `#<function>.<name>` | a parameter or local |
| `#<function>._t<n>`, `#<function>._f<n>` | temporaries and condition flags |
| `#c<n>` | the constant `n`, for multiplication, division and comparison |
| `#cf`, `#rv` | control flow, see below |

**Branches and loops.** `if` with a one-command body becomes `execute if ... run <command>`. Longer bodies become sub-functions next to the function, such as `demo:check/_if0`.

A loop is a function that checks its condition, runs its body and calls itself. `break` returns from it, and `continue` calls it again.

**Leaving early.** When `break`, `continue` or `return` happens inside a sub-function, it sets `#cf`, and the calling function passes it on. The return value travels in `#rv`.

**Results.** Functions return values with the `return` command, and callers read them with `execute store result score`.

**Load.** The generated load function creates the objective, sets the constants, and gives each global its first value.

## Generated functions

`OnLoad { ... }` and `OnTick { ... }` bodies, and `Objective` declarations, are collected into `<namespace>:mms_gen/load` and `<namespace>:mms_gen/tick`. These are registered in `#minecraft:load` and `#minecraft:tick`.

An `OnLoad` or `OnTick` body that holds only commands is copied into these functions as it is. A body with variables or control flow compiles to its own function, `<namespace>:mms_gen/load/<hash>`, which the generated function calls.
