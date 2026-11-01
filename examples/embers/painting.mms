// A painting is data (its size and name) plus a texture (its picture).
// Texture(...) in asset_id copies the file to textures/painting/hearth.png:
// the painting_variant field sets the painting/ prefix.

/** A 2x1 hearth fire, drawn at 16 pixels per block (32x16). */
PaintingVariant hearth = {
  width = 2
  height = 1
  asset_id = Texture("art/hearth.png")
  title = { translate = "painting.embers.hearth.title" }
  author = { translate = "painting.embers.hearth.author" }
}

/** Paintings placed by hand pick from this tag, so hearth can show up on a wall. */
PaintingVariantTag minecraft:placeable = [hearth]
