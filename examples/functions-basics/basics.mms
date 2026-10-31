// A small minigame: players earn points for kills, and every few kills
// they get a reward. Run /function basics:status to see the totals.

/** Total kills across all players since the pack was installed. */
Global total_kills: int = 0

/** Kills needed for each reward. */
Global reward_every: int = 5

Objective kills = playerKillCount
Objective points = dummy
Objective online = dummy

OnLoad {
  /tellraw @a "Points game loaded"
}

OnTick {
  /execute as @a[scores={kills=1..}] run function basics:on_kill
}

/** Runs as a player who has just killed another player. */
Function on_kill {
  total_kills += kills[@s]
  points[@s] += kills[@s] * 10
  kills[@s] = 0

  if (points[@s] % (reward_every * 10) == 0) {
    reward()
  }
}

Function reward {
  /give @s diamond
  /tellraw @s "Reward! You have ${points[@s]} points"
}

/** Counts down from 3, then reports the totals. */
Function status {
  for i in 1..3 {
    /say ${4 - i}...
  }
  // A command can write a score that the code then reads.
  /execute store result score #players online if entity @a
  let players = online["#players"]
  let average = 0
  if (players > 0) {
    average = total_kills / players
  }
  /tellraw @s "${total_kills} kills in total, ${average} per player"
}
