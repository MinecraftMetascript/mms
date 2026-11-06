const welcome: Text = "Welcome to **the Library**, ${@s}! Press {key:inventory} to open your bag."

/** Shown when a wave starts. */
Template wave_start(wave: int): Text =
  "{red}Wave ${wave}{/} begins. [Rules](run:/function siege:rules \"Show the rules\")"

const rules: Text = TextFile("text/rules.mmstext")

OnFirstJoin {
  tell(@s, welcome)
  title(@s, "{gold}The Library", subtitle = "{gray}*mind the shelves*")
}

Function start_wave {
  tell(@a, wave_start(game.wave))
}

CustomItem ember_shard {
  base = amethyst_shard
  name = "{#ff8800}Ember Shard"
  components = { lore = ["{gray}*Still warm.*"] }
}
