// An arena of waves. Mark spawn points, start the game, and each wave
// summons its mobs at every spawn point after a rest. When the last enemy
// of a wave dies the next one starts, and after the last wave the players
// win.
//
//   /function waves:mark_spawn   a spawn point where you stand
//   /function waves:start        the first wave, after its rest

/** One wave of the game. */
Type Wave {
  /** The mob it summons. */
  mob: Id<entity_type>
  /** How many at each spawn point. */
  count: int = 1
  /** Ticks of rest before it starts: 100 is five seconds. */
  rest: int = 100
}

/** Where the game is. */
Type Arena {
  phase: "idle" | "resting" | "fighting" | "won" = "idle"
  /** The wave being fought or rested for, from 1. */
  wave: int = 0
  /** Ticks of rest left. */
  timer: int = 0
}

Storage arena: Arena

/** The waves, in order. Fields left out take Wave's defaults. */
const waves: [Wave] = [
  { mob = zombie, count = 2 },
  { mob = skeleton, count = 2 },
  { mob = zombie, count = 4, rest = 60 },
  { mob = witch, rest = 200 },
]

/** The wave each enemy came from. */
Objective wave_of

/** Says something to everyone, in chat and as a title. */
Template announce(text: Text): Commands {
  tell(@a, text)
  title(@a, text)
}

/** A spawn point where you stand. */
Function mark_spawn {
  /summon marker ~ ~ ~ {Tags:["waves.spawn"]}
  tell(@s, "{gray}Spawn point marked")
}

/** Clears the arena and rests for the first wave. */
Function start {
  /kill @e[tag=waves.enemy]
  arena.wave = 0
  next_wave()
}

/** Rests for the next wave, or ends the game after the last one. */
Function next_wave {
  arena.wave += 1
  if (arena.wave > len(waves)) {
    arena.phase = "won"
    announce("{gold}The arena is yours!")
    return
  }
  arena.phase = "resting"
  rest_for_wave()
  announce("{yellow}Wave ${arena.wave} of ${len(waves)} is coming")
}

/**
 * Sets the rest timer from the current wave. The loop unrolls at compile
 * time, into one line per wave.
 */
Template rest_for_wave(): Commands {
  for i in 1..len(waves) {
    if (arena.wave == i) {
      arena.timer = waves[i - 1].rest
    }
  }
}

/** Summons the current wave at every spawn point. */
Function spawn {
  summon_current()
  wave_of[@e[tag=waves.new]] = arena.wave
  /tag @e[tag=waves.new] remove waves.new
  arena.phase = "fighting"
}

/** The current wave's mobs, one line per wave once it unrolls. */
Template summon_current(): Commands {
  for i in 1..len(waves) {
    if (arena.wave == i) {
      summon_wave(waves[i - 1])
    }
  }
}

/** One wave's mobs at each spawn point, tagged waves.new. */
Template summon_wave(w: Wave): Commands {
  As(@e[tag=waves.spawn]).At(@s) {
    for n in 1..w.count {
      /summon ${w.mob} ~ ~ ~ {Tags:["waves.enemy","waves.new"]}
    }
  }
}

/** Counts down the rest, and moves on once a wave is beaten. */
Function tick {
  if (arena.phase == "resting") {
    arena.timer -= 1
    if (arena.timer <= 0) {
      spawn()
    }
  } else if (arena.phase == "fighting") {
    let alive = /execute if entity @e[tag=waves.enemy]
    if (alive == 0) {
      next_wave()
    }
  }
}

OnTick {
  tick()
}
