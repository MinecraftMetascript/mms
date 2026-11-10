# Tutorial: a wave arena

This tutorial builds a small game: mark some spawn points, start the arena, and waves of mobs come at you, each after a short rest. Beat the last wave and you win. Along the way it covers typed data, state that survives a `/reload`, function code, compile-time loops and tests that run without the game.

The [`waves`](../../examples/waves/) example is a slightly bigger version of the same game. In the web editor, open it from **Examples**.

## 1. Set up the project

```sh
mkdir arena && cd arena
mms init -namespace arena
```

Everything below goes in `main.mms`.

## 2. Describe the waves

Start with the data. A `Type` describes a record, and a `const` holds a value the compiler knows:

```mms
/** One wave of the game. */
Type Wave {
  /** The mob it summons. */
  mob: Id<entity_type>
  /** How many at each spawn point. */
  count: int = 1
}

/** The waves, in order. */
const waves: [Wave] = [
  { mob = zombie, count = 2 },
  { mob = skeleton, count = 2 },
  { mob = witch },
]

/** Ticks of rest before each wave: five seconds. */
const rest = 100
```

`mob` is an `Id<entity_type>`, so the editor completes mob names and `zombi` is an error. `count` has a default, so the witch comes alone. Neither the type nor the constants write a file: they exist for the code below to use.

## 3. Keep the game's state

The game needs to remember which wave it is on and whether players are resting or fighting. A `Storage` keeps a record in command storage, where functions read and write its fields by name:

```mms
/** Where the game is. */
Type Game {
  phase: "idle" | "resting" | "fighting" | "won" = "idle"
  /** The wave being rested for or fought, from 1. */
  wave: int = 0
  /** Ticks of rest left. */
  timer: int = 0
}

Storage game: Game
```

`phase` is one of four words. Comparing it with a word that is not on the list, such as `game.phase == "restng"`, is an error. In the game, `/data get storage arena:mms game` shows the record.

[Keeping state](../guides/state.md) compares Storage with scores and Globals.

## 4. Start the game

Three functions get the game going. They use the `waves`, `rest` and `game` declared above:

```
/** Marks a spawn point where you stand. */
Function mark_spawn {
  /summon marker ~ ~ ~ {Tags:["arena.spawn"]}
  tell(@s, "{gray}Spawn point marked")
}

/** Clears the arena and rests for the first wave. */
Function start {
  /kill @e[tag=arena.enemy]
  game.wave = 0
  next_wave()
}

/** Rests for the next wave, or ends the game after the last one. */
Function next_wave {
  game.wave += 1
  if (game.wave > len(waves)) {
    game.phase = "won"
    title(@a, "{gold}The arena is yours!")
    return
  }
  game.phase = "resting"
  game.timer = rest
  title(@a, "{yellow}Wave ${game.wave} of ${len(waves)}", subtitle = "{gray}Get ready")
}
```

- A marker is an entity that does nothing, which makes it a good bookmark for a place.
- `len(waves)` is worked out when compiling, so adding a wave to the list is all it takes to make the game longer.
- In the title, `${game.wave}` is read from storage when the title shows, and `${len(waves)}` is written in as `3`.

## 5. Summon a wave

Summoning is where the wave list meets the game. The mob is part of a command, and the game has no way to pick a command by number. So the pack needs one `summon` line per wave, and MMS can write them.

A `Template` that returns `Commands` is code that is written out where it is called. Its `for` loops run at compile time:

```
/** Summons the mobs of one wave at every spawn point. */
Template summon_wave(w: Wave): Commands {
  As(@e[tag=arena.spawn]).At(@s) {
    for n in 1..w.count {
      /summon ${w.mob} ~ ~ ~ {Tags:["arena.enemy"]}
    }
  }
}

/** Summons the current wave: one `if` per wave once the loop unrolls. */
Template summon_current(): Commands {
  for i in 1..len(waves) {
    if (game.wave == i) {
      summon_wave(waves[i - 1])
    }
  }
}

/** Starts the fight. */
Function spawn {
  summon_current()
  game.phase = "fighting"
}
```

`As(...).At(@s) { }` runs its body as each marker, at that marker. The built `spawn.mcfunction` tests the wave three times, and each branch holds that wave's own `summon` lines:

```
execute store result score #spawn.-t0 mms.arena run data get storage arena:mms game.wave
execute if score #spawn.-t0 mms.arena matches 1 run function arena:spawn/_if0
...
```

The same `for` written inside a `Function` would be a loop the game runs, and a `Wave` cannot be looked up as the game runs. The editor says so if you try. Hover a template call to see the commands it writes.

## 6. Run it every tick

```
OnTick {
  if (game.phase == "resting") {
    game.timer -= 1
    if (game.timer <= 0) {
      spawn()
    }
  } else if (game.phase == "fighting" && !exists(@e[tag=arena.enemy])) {
    next_wave()
  }
}
```

`OnTick` runs twenty times a second. While resting it counts the timer down. While fighting it waits for the last enemy to die: `exists(selector)` is true when the selector finds anything.

## 7. Test it

A `Test` runs functions in an interpreter built into `mms`, so there is no need to start the game to find out whether the logic holds:

```
Test start_rests_before_the_first_wave {
  run start()
  expect game.phase == "resting"
  expect game.wave == 1
  expect game.timer == 100
}

Test a_wave_spawns_at_every_spawn_point {
  setup {
    /summon marker 0 64 0 {Tags:["arena.spawn"]}
    /summon marker 8 64 0 {Tags:["arena.spawn"]}
    game.wave = 2
  }
  run spawn()
  expect game.phase == "fighting"
  expect count(@e[type=skeleton, tag=arena.enemy]) == 4
}

Test beating_the_last_wave_wins {
  setup {
    game.wave = 3
  }
  run next_wave()
  expect game.phase == "won"
}
```

```sh
mms test -v
```

```
ok    start_rests_before_the_first_wave
    said: Get ready
    said: Wave 1 of 3
ok    a_wave_spawns_at_every_spawn_point
ok    beating_the_last_wave_wins
    said: The arena is yours!
ok: 3 tests passed
```

Each test starts in a fresh world. `setup` puts the world in the state the test needs, `run` calls a function and `expect` checks the result. The web editor's Tests tab runs them after every edit. [Testing functions](../guides/testing.md) has more.

## 8. Play

```sh
mms build --world ~/.minecraft/saves/MyWorld
```

In the world, run `/reload`, stand where mobs should appear and run `/function arena:mark_spawn` a few times, then `/function arena:start`.

## The whole file

```mms
/** One wave of the game. */
Type Wave {
  /** The mob it summons. */
  mob: Id<entity_type>
  /** How many at each spawn point. */
  count: int = 1
}

/** The waves, in order. */
const waves: [Wave] = [
  { mob = zombie, count = 2 },
  { mob = skeleton, count = 2 },
  { mob = witch },
]

/** Ticks of rest before each wave: five seconds. */
const rest = 100

/** Where the game is. */
Type Game {
  phase: "idle" | "resting" | "fighting" | "won" = "idle"
  /** The wave being rested for or fought, from 1. */
  wave: int = 0
  /** Ticks of rest left. */
  timer: int = 0
}

Storage game: Game

/** Marks a spawn point where you stand. */
Function mark_spawn {
  /summon marker ~ ~ ~ {Tags:["arena.spawn"]}
  tell(@s, "{gray}Spawn point marked")
}

/** Clears the arena and rests for the first wave. */
Function start {
  /kill @e[tag=arena.enemy]
  game.wave = 0
  next_wave()
}

/** Rests for the next wave, or ends the game after the last one. */
Function next_wave {
  game.wave += 1
  if (game.wave > len(waves)) {
    game.phase = "won"
    title(@a, "{gold}The arena is yours!")
    return
  }
  game.phase = "resting"
  game.timer = rest
  title(@a, "{yellow}Wave ${game.wave} of ${len(waves)}", subtitle = "{gray}Get ready")
}

/** Summons the mobs of one wave at every spawn point. */
Template summon_wave(w: Wave): Commands {
  As(@e[tag=arena.spawn]).At(@s) {
    for n in 1..w.count {
      /summon ${w.mob} ~ ~ ~ {Tags:["arena.enemy"]}
    }
  }
}

/** Summons the current wave: one `if` per wave once the loop unrolls. */
Template summon_current(): Commands {
  for i in 1..len(waves) {
    if (game.wave == i) {
      summon_wave(waves[i - 1])
    }
  }
}

/** Starts the fight. */
Function spawn {
  summon_current()
  game.phase = "fighting"
}

OnTick {
  if (game.phase == "resting") {
    game.timer -= 1
    if (game.timer <= 0) {
      spawn()
    }
  } else if (game.phase == "fighting" && !exists(@e[tag=arena.enemy])) {
    next_wave()
  }
}

Test start_rests_before_the_first_wave {
  run start()
  expect game.phase == "resting"
  expect game.wave == 1
  expect game.timer == 100
}

Test a_wave_spawns_at_every_spawn_point {
  setup {
    /summon marker 0 64 0 {Tags:["arena.spawn"]}
    /summon marker 8 64 0 {Tags:["arena.spawn"]}
    game.wave = 2
  }
  run spawn()
  expect game.phase == "fighting"
  expect count(@e[type=skeleton, tag=arena.enemy]) == 4
}

Test beating_the_last_wave_wins {
  setup {
    game.wave = 3
  }
  run next_wave()
  expect game.phase == "won"
}
```

## Next steps

- Give each wave its own rest, as the [`waves`](../../examples/waves/) example does, or a reward when it is beaten.
- Show the wave on a boss bar or the action bar. [Text commands](../guides/functions.md#text-commands) covers `tell`, `title` and `actionbar`.
- Read [How functions compile](../concepts/how-functions-compile.md) to see what the pack looks like inside.
