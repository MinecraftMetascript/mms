// Two new-looking items, built on vanilla ones. The item_model component
// points a stack at our item model; item_name gives it our translated name.
//
// Texture("art/ember_shard.png", "item/ember_shard") copies the file to
// assets/embers/textures/item/ember_shard.png in the resource pack. Only
// files that code names are copied, so art/ can hold drafts too.

/** A flat item model drawn from one texture, like most vanilla items. */
Model item/ember_shard = {
  parent = item/generated
  textures = { layer0 = Texture("art/ember_shard.png", "item/ember_shard") }
}

Model item/ember_disc = {
  parent = item/generated
  textures = { layer0 = Texture("art/ember_disc.png", "item/ember_disc") }
}

/** The item models the item_model component names. */
ItemModel ember_shard = { model = Model(model = item/ember_shard) }

ItemModel ember_disc = { model = Model(model = item/ember_disc) }

/** Blaze powder and an amethyst shard make two ember shards. */
Recipe ember_shard = Shapeless(
  [blaze_powder, amethyst_shard],
  {
    id = amethyst_shard
    count = 2
    components = {
      item_model = ember_shard
      item_name = { translate = "item.embers.ember_shard" }
    }
  },
)

/**
 * Blaze powder around a music disc that creepers drop (#creeper_drop_music_discs,
 * most of the vanilla discs) makes the Embers disc, which plays the
 * song in music.mms. Recipe ingredients match by item id only, so a recipe
 * cannot ask for an Ember Shard: any amethyst shard would do. This one uses
 * plain vanilla items instead.
 */
Recipe ember_disc = Shaped(
  [" B ", "BDB", " B "],
  {
    B = blaze_powder
    D = #creeper_drop_music_discs
  },
  {
    id = music_disc_13
    components = {
      item_model = ember_disc
      item_name = { translate = "item.embers.ember_disc" }
      jukebox_playable = embers
    }
  },
)
