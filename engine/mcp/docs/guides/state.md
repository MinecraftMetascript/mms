# Keeping state

A data pack has a few places to remember things between commands: scoreboards, command storage and tags on entities. MMS gives each a name you can read and write like a variable. This guide is about choosing between them.

| You need | Use | Lives in |
| --- | --- | --- |
| a value while one function runs | `let` | a scratch score |
| a number or flag for the whole pack | `Global` | a score |
| a number for each player or entity | `Objective` | a scoreboard objective |
| several values that belong together, or text | `Storage` | command storage |
| a yes or no for each entity | a tag | the entity |

## One value for the whole pack

A `Global` is a single value that every function sees:

```mms
/** The round being played. */
Global round: int = 1
/** Whether new players may join the round. */
Global open: bool = true

Function next_round {
  round += 1
  open = round < 10
  tell(@a, "Round ${round} begins")
}
```

The value after `=` is where it starts. It is set once, when the pack first loads, and a `/reload` keeps whatever it has become. An `int` or `bool` Global is a score, the cheapest thing to read and change, so it is the right home for a counter or a timer that runs every tick. A `string` Global is kept in storage.

To look at an `int` or `bool` Global in the game, run `/scoreboard players get $round mms.<namespace>`. A `string` one is in `/data get storage <namespace>:mms g`.

## A number for each player

An `Objective` is a scoreboard objective. `objective[holder]` is one holder's score:

```mms
/** Mana each player has, up to 20. */
Objective mana
/** The game counts these by itself. */
Objective kills: playerKillCount

OnFirstJoin {
  mana[@s] = 20
}

Function cast {
  if (mana[@s] >= 5) {
    mana[@s] -= 5
    /effect give @s speed 5 1
  } else {
    actionbar(@s, "{red}Not enough mana")
  }
}

OnTick {
  /scoreboard players add @a[scores={mana=..19}] mana 1
}

/** Adds a player's kills to the server's total. */
Function bank_kills {
  kills["#total"] += kills[@s]
  kills[@s] = 0
}
```

- **The holder** is a selector, such as `@s`, or a name in quotes. A name that starts with `#`, such as `"#total"`, is a made-up holder that never shows on the sidebar, which makes it a handy pack-wide slot on an objective.
- **A criterion** after the colon makes the game keep the score: `deathCount`, `playerKillCount`, `minecraft.used:minecraft.carrot_on_a_stick` and the rest. Without one the objective is `dummy`, changed only by your code.
- **Mobs have scores too.** `hp[@e[tag=boss, limit=1]]` works like a player's, and the score goes away with the entity.
- **A score nobody has set is not 0.** It counts as 0 in arithmetic, but `<`, `<=`, `>`, `>=` and `==` are all false for it and no selector matches it: `mana[@s] < 5` is false for a player whose mana was never set, and `@a[scores={mana=..19}]` skips them. Give each player a starting value in `OnFirstJoin`, as above, and prefer asking whether a player has enough (`>=`) over whether they have too little.

Raw commands and score variables mix freely. One selector that changes many players at once, as in the `OnTick` above, is cheaper than a loop.

## Values that belong together

A `Storage` keeps a record in command storage. The record's shape is a `Type`, so fields are completed, checked and documented:

```mms
/** The settings an admin can change. */
Type Settings {
  difficulty: "easy" | "hard" = "easy"
  max_players: int = 8
  /** Where players start. */
  spawn: { x: int, y: int = 64, z: int }
  /** Shown to players who join. */
  motd: string = "Welcome"
}

Storage settings: Settings

Function harder {
  settings.difficulty = "hard"
  settings.max_players = 4
  settings.spawn.y += 10
  tell(@a, "${settings.motd}. The game is now ${settings.difficulty}.")
}
```

Reach for Storage when:

- the values are one thing, such as a game's phase, wave and timer, and reading `game.phase` beats three Globals with a shared prefix
- a value is text, or one of a few words such as `"easy" | "hard"`, where a wrong word should be an error
- other commands need the data: the record is at `<namespace>:mms <name>`, so `/data get storage demo:mms settings` shows it, and macros and `data modify` can use it

Each read or write of a field is a `data` command, which costs a little more than a score. In code that runs for every entity on every tick, copy the field into a `let` first, or keep that value in a `Global`.

A field added later starts at its default the next time the pack loads, and existing values are kept, so a record can grow without resetting a world.

## Locals

`let` names a value inside one function. Its type comes from the value:

```mms
Objective mana

Function refill {
  let missing = 20 - mana[@s]
  let full: bool = missing <= 0
  if (!full) {
    mana[@s] += missing
  }
}
```

A local is not per entity. Inside `As(@a) { ... }` there is still one `missing`, which each player's turn overwrites. Anything that must be remembered per entity goes in an objective.

## Flags on entities

For a yes or no per entity, the game's own tags are enough, and selectors read them directly:

```mms
const frozen = @a[tag=frozen]

Function freeze {
  /tag @s add frozen
}

OnTick {
  /effect give ${frozen} slowness 1 255 true
}
```

A [named selector](functions.md#selectors) keeps the tag's name in one place. Prefix tags with your namespace, such as `siege.boss`, when other packs may be loaded.

## Values the game already has

Not everything needs storing. A command after `=` reads a value straight from the game:

```mms
Function report {
  let time = /time query gametime
  let online = /execute if entity @a
  let holding_sword: bool = /execute if items entity @s weapon.mainhand #swords
  tell(@s, "Tick ${time}, with ${online} players online")
}
```

`exists(selector)` and `count(selector)` cover the common questions about entities, and a [`Predicate`](functions.md#predicates) is a condition by name, such as `if (sneaking)`.

## Starting values and resets

- `OnLoad { ... }` runs each time the pack loads, including every `/reload`, after Globals and Storage fields have their values. Defaults are given only once, so use `OnLoad` for anything that must be put back on every load.
- `OnFirstJoin { ... }` runs once per player, which is the place to give a starting score.
- A reset is a function that assigns the starting values again:

```mms
Global round: int = 1
Objective mana

Function reset {
  round = 1
  mana[@a] = 20
}
```

## Checking state without the game

A [`Test`](testing.md) can set state up, run a function and check what changed, with Globals, scores and Storage fields all readable in `expect` lines. State is where most pack bugs live, so it is the part most worth testing.
