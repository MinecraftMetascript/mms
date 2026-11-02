// The custom arrows are tipped arrows. A tipped arrow keeps its item on
// the arrow entity once it is shot, components and all, so the tick
// function in flight.mms can tell which kind of arrow is flying by its
// custom_data. A tipped arrow also:
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
    custom_data = { arrow = "explosive" }
  }
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
    custom_data = { arrow = "ender" }
  }
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
 * Every arrow it shoots calls lightning where it hits. The thunder
 * enchantment in bow.mms does the work, so it fires on the exact block
 * or mob the arrow hits, whatever the arrow is.
 */
CustomItem thunder_bow = {
  base = bow
  name = "Thunder Bow"
  components = { enchantments = { "arrows:thunder" = 1 } }
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

Recipe cow_wand = Shapeless([stick, stick], cow_wand)
