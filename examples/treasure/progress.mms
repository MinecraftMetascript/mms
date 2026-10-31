// Predicates, advancements and a function that ties them together.

/** Thundering, and night (or a coin flip). */
Predicate stormy_night = WeatherCheck(thundering = true)
  && (TimeCheck(clock = overworld, value = 13000..23000) || RandomChance(0.5))

/** Granted the first time a player picks up a gem. */
Advancement first_gem = {
  criteria = { has_gem = InventoryChanged(items = [{ items = #gems }]) }
  rewards = { function = treasure:found_gem }
}

/** Runs as the player who earned first_gem, and rolls chests/storm_bonus for them. */
Function found_gem {
  /tellraw @s "You found your first gem. Four gems craft into a gem block!"
  /execute if predicate treasure:stormy_night run loot give @s loot treasure:chests/storm_bonus
}
