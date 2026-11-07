Test second_wave_spawns_four {
  setup {
    game.wave = 1
    /summon marker 0 64 0 {Tags:["siege.spawn"]}
  }
  run next_wave()
  expect game.wave == 2
  expect count(@e[tag=siege.enemy]) == 4
}
