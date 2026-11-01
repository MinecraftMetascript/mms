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
 * Ember shards around any music disc make the Embers disc, which plays the
 * song in music.mms. The ingredient is the plain amethyst shard: recipes
 * cannot ask for components, so the shard recipe is the only way in.
 */
Recipe ember_disc = Shaped(
  [" A ", "ADA", " A "],
  {
    A = amethyst_shard
    D = #minecraft:creeper_drop_music_discs
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
