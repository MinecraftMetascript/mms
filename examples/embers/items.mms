// Two new-looking items, built on vanilla ones. Each CustomItem writes a
// flat item model drawn from its texture, the item model definition that
// the item_model component names, and its name in the lang file. Its
// name then stands for the base item with those components, wherever an
// item stack goes: recipe results here, loot entries in loot.mms and
// /give in give.mms.
//
// Texture("art/ember_shard.png") copies the file to
// assets/embers/textures/item/ember_shard.png in the resource pack. Only
// files that code names are copied, so art/ can hold drafts too.

/** An amethyst shard that glows like a coal. */
CustomItem ember_shard = {
  base = amethyst_shard
  name = "Ember Shard"
  texture = Texture("art/ember_shard.png")
}

/** A music disc that plays the song in music.mms. */
CustomItem ember_disc = {
  base = music_disc_13
  name = "Music Disc"
  texture = Texture("art/ember_disc.png")
  components = { jukebox_playable = embers }
}

/** Blaze powder and an amethyst shard make two ember shards. */
Recipe ember_shard = Shapeless([blaze_powder, amethyst_shard], ember_shard * 2)

/**
 * Blaze powder around a music disc that creepers drop (#creeper_drop_music_discs,
 * most of the vanilla discs) makes the Embers disc. Recipe ingredients
 * match by item id only, so a recipe cannot ask for an Ember Shard: any
 * amethyst shard would do. This one uses plain vanilla items instead.
 */
Recipe ember_disc = Shaped(
  [" B ", "BDB", " B "],
  {
    B = blaze_powder
    D = #creeper_drop_music_discs
  },
  ember_disc,
)

/**
 * Right-clicking with an Ember Shard puffs out flames and keeps the
 * player safe from fire for a few seconds. It runs once per press, as
 * and at the player; holding the button down does not repeat it.
 */
OnRightClick ember_shard {
  /particle flame ^ ^1.5 ^1 0.2 0.2 0.2 0.05 20
  /playsound item.firecharge.use player @a
  /effect give @s fire_resistance 5
}
