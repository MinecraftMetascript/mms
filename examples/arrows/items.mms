// The custom arrows are tipped arrows. on_projectile_hit_block and
// on_projectile_hit_entity run a function where the arrow hits (see
// effects.mms). A tipped arrow as the base also:
//
// - takes its colour from potion_contents, on the item and in flight,
//   where it trails particles of that colour (custom_color is
//   RED << 16 | GREEN << 8 | BLUE, so #E0441B is 14697499);
// - is never free: Infinity only spares plain arrows;
// - takes its name from potion_contents.custom_name, as
//   item.minecraft.tipped_arrow.effect.<custom_name> (see lang.mms).

/** Explodes where it lands, whether it hits a block or a mob. */
CustomItem explosive_arrow = {
  base = tipped_arrow
  name = "Explosive Arrow"
  components = {
    potion_contents = {
      custom_color = 14697499
      custom_name = "arrows.explosive"
    }
  }
  on_projectile_hit_block = explode_here
  on_projectile_hit_entity = explode
}

/** Teleports whoever shot it to where it lands, like an ender pearl. */
CustomItem ender_arrow = {
  base = tipped_arrow
  name = "Ender Arrow"
  components = {
    potion_contents = {
      custom_color = 2062958
      custom_name = "arrows.ender"
    }
  }
  on_projectile_hit_block = ender_block
  on_projectile_hit_entity = ender_entity
}

/**
 * Lifts what it hits into the air for three seconds. This one needs no
 * commands at all: it is a tipped arrow with an effect of its own.
 * potion_duration_scale is 1 so that the effect lasts as long as written
 * (a tipped arrow normally gives an eighth of its potion's time).
 */
CustomItem levitation_arrow = {
  base = tipped_arrow
  name = "Arrow of Levitation"
  components = {
    potion_contents = {
      custom_color = 13565951
      custom_name = "arrows.levitation"
      custom_effects = [
        {
          id = levitation
          duration = 60
          amplifier = 1
        },
      ]
    }
    potion_duration_scale = 1
  }
}

/**
 * Every arrow it shoots calls lightning where it hits, whatever the
 * arrow is. On a bow or crossbow, the projectile hooks follow the arrows
 * it shoots.
 */
CustomItem thunder_bow = {
  base = bow
  name = "Thunder Bow"
  on_projectile_hit_block = strike
  on_projectile_hit_entity = strike
}

/**
 * A splash potion that explodes where it lands. Thrown potions keep no
 * bow, so on_projectile_land follows the potion itself instead.
 */
CustomItem blast_potion = {
  base = splash_potion
  name = "Blast Potion"
  components = {
    potion_contents = {
      custom_color = 14697499
      custom_name = "arrows.blast"
    }
  }
  on_projectile_land = explode
}

/** A stick that fires cows. See launch.mms. */
CustomItem cow_wand = {
  base = stick
  name = "Cow Wand"
  on_right_click = shoot_cow
}

Recipe explosive_arrow = Shapeless([arrow, gunpowder], explosive_arrow)

Recipe ender_arrow = Shapeless([arrow, arrow, ender_pearl], ender_arrow * 2)

Recipe levitation_arrow = Shaped(
  ["AAA", "ASA", "AAA"],
  {
    A = arrow
    S = shulker_shell
  },
  levitation_arrow * 8,
)

Recipe thunder_bow = Shapeless([bow, lightning_rod], thunder_bow)

Recipe blast_potion = Shapeless([glass_bottle, tnt], blast_potion)

Recipe cow_wand = Shapeless([stick, stick], cow_wand)
