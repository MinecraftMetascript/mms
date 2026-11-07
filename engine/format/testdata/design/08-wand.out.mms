/** A wand: a stick with a texture, a spell and a cooldown. */
Blueprint Wand {
  name: string
  texture: Texture
  cast: Function
  /** Ticks between casts. */
  cooldown: int = 20

  check {
    if (cooldown < 1) error(cooldown, "A cooldown is at least one tick")
  }

  emit {
    primary CustomItem ${self.path} {
      base = stick
      name = name
      texture = texture
      parent = item/handheld
      on_right_click = run {
        if (wand_cooldown[@s] > 0) {
          return
        }
        wand_cooldown[@s] = cooldown
        cast()
      }
    }
  }
}

Objective wand_cooldown

OnTick {
  /scoreboard players remove @a[scores={wand_cooldown=1..}] wand_cooldown 1
}

Wand fire_wand {
  name = "Fire Wand"
  texture = Texture("art/fire_wand.png")
  cast = shoot_fireball
}

Recipe fire_wand = Shapeless([stick, blaze_powder], fire_wand) // the primary CustomItem's stack
