# Waves

An arena of waves. Mark spawn points, start the game, and each wave summons its mobs at every spawn point after a short rest. When the last enemy of a wave dies, the next wave comes; after the last one, the arena is yours.

## Try it

```
mms build --world ~/.minecraft/saves/MyWorld
```

Then, in the world:

- `/function waves:mark_spawn` marks a spawn point where you stand. Mark a few.
- `/function waves:start` starts the first wave.

Run the tests with `mms test`, or open the example in the web editor and look at the Tests tab.

## How it works

| File | What it does |
| --- | --- |
| `arena.mms` | The waves, the game's state, and the functions that run it. |
| `tests.mms` | Tests of starting, spawning, resting and winning, run without the game. |

**Typed data.** `Type Wave` describes one wave: its mob, how many come from each spawn point, and the rest before it. `count` and `rest` have defaults, so `{ mob = witch, rest = 200 }` comes one at a time. `const waves: [Wave]` lists them, and the editor completes and checks every field.

**State in Storage.** `Storage arena: Arena` keeps the phase, the wave and the rest timer in command storage. `phase` is one of four words, so `arena.phase == "resting"` is checked against them, and a typo is an error.

**Loops that unroll.** A `Template` that returns `Commands` writes its body where it is called. Its `for i in 1..len(waves)` loop runs at compile time, so `summon_current()` becomes one `if` per wave, each with that wave's mob and count written in, and `rest_for_wave()` one line per wave.

**Per-entity scores.** Every enemy keeps the wave it came from in `wave_of[@e[tag=waves.new]]`. The tests check it with `count(@e[scores={ wave_of = 3 }])`.

**Command values.** `let alive = /execute if entity @e[tag=waves.enemy]` keeps the number of enemies left, so the tick function knows when a wave is beaten.

**Text with values.** `"Wave ${arena.wave} of ${len(waves)}"` shows the wave from storage when the message is sent, and the number of waves, worked out at compile time.

## Tests

Each test starts from a fresh world after the pack's load function, so `arena` holds its defaults. The interpreter has no players, so `@a` selects nobody and the messages only show in `mms test -v`. Markers stand in for spawn points, and the tests count what the functions summon.
