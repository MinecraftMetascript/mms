// A small minigame: players earn points for kills, and every few kills
// they get a reward. Run /function basics:status to see the totals.

/** Total kills across all players since the pack was installed. */
Global total_kills: int = 0

/** Kills needed for each reward. */
Global reward_every: int = 5

// Objective names are global to the world, so they carry the pack's name.
Objective basics_kills = playerKillCount
Objective basics_points = dummy
Objective basics_online = dummy

OnLoad {
  /tellraw @a "Points game loaded"
}

OnTick {
  /execute as @a[scores={basics_kills=1..}] run function basics:on_kill
}

/** Runs as a player who has just killed another player. */
Function on_kill {
  let before = basics_points[@s]
  total_kills += basics_kills[@s]
  basics_points[@s] += basics_kills[@s] * 10
  basics_kills[@s] = 0

  // Several kills can land in one tick, so give a reward for every
  // threshold passed, not only when the score is an exact multiple.
  let step = reward_every * 10
  let rewards = basics_points[@s] / step - before / step
  while (rewards > 0) {
    reward()
    rewards -= 1
  }
}

Function reward {
  /give @s diamond
  /tellraw @s "Reward! You have ${basics_points[@s]} points"
}

/** Counts down from 3, then reports the totals. */
Function status {
  for i in 1..3 {
    /say ${4 - i}...
  }
  // A command can write a score that the code then reads.
  /execute store result score #players basics_online if entity @a
  let players = basics_online["#players"]
  let average = 0
  if (players > 0) {
    average = total_kills / players
  }
  /tellraw @s "${total_kills} kills in total, ${average} per player"
}
