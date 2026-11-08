// Tests of the arena. Run them with `mms test`, in the web editor's Tests
// tab, or in VS Code's Testing view. They are never written to the pack.

Test start_rests_before_the_first_wave {
  run start()
  expect arena.phase == "resting"
  expect arena.wave == 1
  expect arena.timer == 100
}

Test a_wave_spawns_at_every_spawn_point {
  setup {
    /summon marker 0 64 0 {Tags:["waves.spawn"]}
    /summon marker 8 64 0 {Tags:["waves.spawn"]}
    arena.wave = 3
  }
  run spawn()
  expect arena.phase == "fighting"
  expect count(@e[type=zombie, tag=waves.enemy]) == 8
  // Each enemy remembers its wave.
  expect count(@e[scores={ wave_of = 3 }]) == 8
}

Test the_rest_counts_down_to_the_wave {
  setup {
    /summon marker 0 64 0 {Tags:["waves.spawn"]}
    arena.phase = "resting"
    arena.wave = 2
    arena.timer = 2
  }
  run tick()
  expect arena.timer == 1
  expect count(@e[tag=waves.enemy]) == 0
  run tick()
  expect arena.phase == "fighting"
  expect count(@e[type=skeleton]) == 2
}

Test a_wave_lasts_while_an_enemy_lives {
  setup {
    /summon zombie 0 64 0 {Tags:["waves.enemy"]}
    arena.phase = "fighting"
    arena.wave = 1
  }
  run tick()
  expect arena.phase == "fighting"
  expect arena.wave == 1
}

Test beating_a_wave_rests_for_the_next {
  setup {
    arena.phase = "fighting"
    arena.wave = 2
  }
  run tick()
  expect arena.phase == "resting"
  expect arena.wave == 3
  // Wave 3 sets its own, shorter rest.
  expect arena.timer == 60
}

Test beating_the_last_wave_wins {
  setup {
    arena.phase = "fighting"
    arena.wave = 4
  }
  run tick()
  expect arena.phase == "won"
}
