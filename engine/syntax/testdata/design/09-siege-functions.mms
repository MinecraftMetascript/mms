/** The state of a siege, kept in storage siege:mms game. */
Type SiegeState {
  wave: int = 0
  /** Mobs still to spawn this wave. */
  queue: [Id<entity_type>] = []
  boss_alive: bool = false
}

Storage game: SiegeState

const max_wave = 10
const boss = @e[tag=siege.boss, limit=1]
const wave_mobs = [zombie, skeleton, spider, creeper]

/** Draws a ring of particles around the executor, at compile time. */
Template particle_ring(p: Id<particle>, radius: float, points: int = 16): Commands {
  for i in 0..points - 1 {
    let a = i * 360 / points
    /particle ${p} ^${radius * cos(a)} ^1 ^${radius * sin(a)}
  }
}

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
  As(@a).At(@s) {
    particle_ring(flame, 2)
  }
  tell(@a, "Wave {red}${game.wave}{/} of ${max_wave} begins")
}
