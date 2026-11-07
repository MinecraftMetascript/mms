# Tests

A `Test` runs functions in an interpreter, without starting the game, and checks what they did:

```mms
Type SiegeState { wave: int = 0 }
Storage game: SiegeState

Function next_wave {
  game.wave += 1
  for i in 1..game.wave * 2 {
    As(@e[tag=siege.spawn]).At(@s) {
      /summon zombie ~ ~ ~ {Tags:["siege.enemy"]}
    }
  }
}

Function fib(n: int): int {
  if (n < 2) {
    return n
  }
  return fib(n - 1) + fib(n - 2)
}

Test second_wave_spawns_four {
  setup {
    game.wave = 1
    /summon marker 0 64 0 {Tags:["siege.spawn"]}
  }
  run next_wave()
  expect game.wave == 2
  expect count(@e[tag=siege.enemy]) == 4
}

Test fibonacci {
  expect fib(10) == 55
}
```

| Part | Does |
| --- | --- |
| `setup { }` | function code and commands that run first. Its `let` variables are in scope for the lines after it. |
| `run f()` | calls a function |
| `expect cond` | checks a condition. The test stops at the first one that fails, and a failed comparison shows both sides. |

Each test starts from a fresh world after the pack's load functions. `run` and `expect` lines can come in any order after `setup`.

`mms test` runs every test, the web editor's Tests tab runs them after every edit, and VS Code lists them in its Testing view. The [Functions guide](../../guides/functions.md#tests) lists what the interpreter can run.
