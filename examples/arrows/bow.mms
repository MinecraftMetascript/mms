// The Thunder Bow's enchantment.
//
// A bow is the weapon an arrow was "fired from", and the game keeps it
// on the arrow. When the arrow hits, the game runs the bow's enchantment
// effects at the exact spot: hit_block when it sticks in a block, and
// post_attack when it hurts a mob or player. So a custom bow needs no
// tick function at all, and it works with any arrow, even plain ones.
//
// Nothing else hands out this enchantment: it is in none of the vanilla
// enchantment tags, so enchanting tables, villagers and loot never pick
// it. A grindstone takes it off, like any enchantment that isn't a curse.

/** Calls lightning wherever an arrow from the bow hits. */
Enchantment thunder = {
  description = {
    translate = "enchantment.arrows.thunder"
    fallback = "Thunder"
  }
  supported_items = bow
  weight = 1
  max_level = 1
  min_cost = {
    base = 30
    per_level_above_first = 0
  }
  max_cost = {
    base = 50
    per_level_above_first = 0
  }
  anvil_cost = 8
  slots = [mainhand, offhand]
  effects = {
    hit_block = [
      {
        effect = {
          type = summon_entity
          entity = lightning_bolt
        }
      },
    ]
    post_attack = [
      {
        // Only arrows: hitting a mob with the bow itself does nothing.
        requirements = DamageSourceProperties(predicate = { is_direct = false })
        enchanted = attacker
        affected = victim
        effect = {
          type = summon_entity
          entity = lightning_bolt
        }
      },
    ]
  }
}
