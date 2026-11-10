# Testing functions

Finding out whether a function works usually means building the pack, opening a world, running `/reload`, setting up some scores by hand and watching chat. A `Test` does that in a fraction of a second, without the game: `mms` has an interpreter for the commands functions compile to, and a test runs your functions in it and checks what they did.

## A first test

```mms
Function double(n: int): int {
  return n * 2
}

Test doubling {
  expect double(21) == 42
  expect double(-1) == -2
}
```

```sh
mms test
```

```
ok: 1 test passed
```

A test of a function that returns a value needs nothing but `expect` lines. Tests are never written into the pack, so they can sit next to the code they test or in a file of their own, such as `tests.mms`.

## Where tests run

- **`mms test`** runs every test and exits with 1 when one fails. `-run <regexp>` keeps the tests whose names match, and `-v` lists passing tests too, with everything the functions said in chat, titles and the action bar.
- **The web editor** shows a Tests tab once a project has a test, and runs them all again after every edit.
- **VS Code** lists them in its Testing view, and marks a failing `expect` at its line.

All three use the same interpreter, so they agree.

## Setup, run, expect

Most functions change state instead of returning something. A test of one has three parts:

```mms
Type SiegeState {
  wave: int = 0
  boss_alive: bool
}

Storage game: SiegeState

Function next_wave {
  game.wave += 1
  game.boss_alive = game.wave % 5 == 0
}

Test every_fifth_wave_has_a_boss {
  setup {
    game.wave = 4
  }
  run next_wave()
  expect game.wave == 5
  expect game.boss_alive
}
```

| Part | Does |
| --- | --- |
| `setup { }` | puts the world in the state the test needs. It is function code: assignments, commands, calls, `let`. |
| `run f()` | calls a function |
| `expect cond` | checks a condition, and stops the test when it is false |

Every test starts in a fresh world, right after the pack's load function, so Globals and Storage fields hold their defaults and no test sees what another one did. `run` and `expect` lines can alternate, which suits anything that happens over time:

```mms
Global timer: int = 2
Global rang: bool = false

Function tick {
  timer -= 1
  if (timer == 0) {
    rang = true
  }
}

Test rings_on_the_second_tick {
  run tick()
  expect !rang
  run tick()
  expect rang
}
```

## Standing in for a player

There are no players in a test, so `@a` and `@p` find no one. Most function code only needs *something* to run as, and a summoned entity will do. Give it a tag, name the selector, and run the function as it:

```mms
/** Mana each player has. */
Objective mana
/** Spells cast since the pack loaded. */
Global casts: int = 0

/** Takes the cost from whoever runs it. False when they can't pay. */
Function pay(cost: int): bool {
  if (mana[@s] >= cost) {
    mana[@s] -= cost
    casts += 1
    return true
  }
  return false
}

/** Stands in for a player. */
const caster = @e[tag=caster, limit=1]

Test paying_takes_mana {
  setup {
    /summon armor_stand 0 64 0 {Tags:["caster"]}
    mana[caster] = 12
    let paid = false
    As(caster) {
      paid = pay(5)
    }
  }
  expect paid
  expect mana[caster] == 7
  expect casts == 1
}

Test the_poor_pay_nothing {
  setup {
    /summon armor_stand 0 64 0 {Tags:["caster"]}
    mana[caster] = 3
    let paid = true
    As(caster) {
      paid = pay(5)
    }
  }
  expect !paid
  expect mana[caster] == 3
}
```

A `let` in `setup` stays in scope for the `expect` lines, which is how `paid` gets out of the `As` block. Entities keep their type, tags, position and scores, so `count(@e[type=zombie, tag=wave])` and `exists(...)` can check what a function summoned or killed.

## When a test fails

```
FAIL  paying_takes_mana
    main.mms:39:3: expected mana[caster] == 8, but mana[caster] is 7
1 of 3 tests failed
```

A failed comparison shows both sides with their values. The line is the `expect`, and the editors mark it in place.

## What the interpreter runs

The interpreter knows the commands that function code compiles to, and enough of the game to check logic:

| Works | |
| --- | --- |
| scores, Globals, Storage, strings | `scoreboard`, `data` on storage |
| control flow and calls | `execute if`/`unless`/`as`/`at`/`positioned`/`store`, `function`, `return` |
| entities | `summon`, `kill`, `tag`, and `@e` and `@s` with `type`, `tag`, `scores`, `limit` |
| messages | `say`, `tellraw`, `title`, which `-v` prints |
| `give` | listed by `-v` |

There are no blocks, no players and no ticks. A command the interpreter does not know, such as `particle`, `effect` or `setblock`, fails the test and names the command. That is on purpose: a test never passes by skipping something.

## Writing code that can be tested

The limit suggests a habit that is good anyway: keep **deciding** apart from **showing**. The function that works out what happens changes scores and storage, and returns the answer. The function the game calls asks it, and then does the particles, sounds and effects.

```mms
Objective mana

/** Deciding: can be tested. */
Function pay(cost: int): bool {
  if (mana[@s] >= cost) {
    mana[@s] -= cost
    return true
  }
  return false
}

/** Showing: a wand's right-click hook. */
Function cast {
  if (!pay(5)) {
    actionbar(@s, "{red}Not enough mana")
    return
  }
  /particle flame ~ ~1 ~ 0.2 0.2 0.2 0.05 20
  /effect give @s speed 5 1
}
```

`pay` holds everything that can go wrong, and the tests above cover it. `cast` is short enough to check by eye, and a test can still run its refusing branch, since that stops before the first particle.

`pay` asks whether the player *has* enough, not whether they have too little, and that matters. A score nobody has set is not a number yet, so every comparison with it is false: `mana[@s] < cost` would let a brand new player cast for free. A test with a caster whose mana is never set would catch it.

## What to test

- **Rules with edges.** Costs, limits, cooldowns, the last wave, an empty list: anywhere `<` could have been `<=`.
- **State machines.** Set a phase, run the tick function, expect the next phase. The [wave arena tutorial](../start/tutorial-minigame.md) tests its game this way.
- **Blueprints.** A test can call the functions an instance emits, which checks the Blueprint itself. The [`shop`](../../examples/shop/) example does.
- **Bugs you have fixed.** A test named after the bug keeps it fixed.

`mms test` needs only the project folder, so it fits a CI step next to `mms check`. An AI agent can run the same tests through [`mms mcp`](agents.md).

The [`functions-advanced`](../../examples/functions-advanced/), [`waves`](../../examples/waves/), [`shop`](../../examples/shop/) and [`treecapitator`](../../examples/treecapitator/) examples all have a `tests.mms` to read.
