// Glowstone, retextured as a block of flickering embers. The image is a
// strip of four 16x16 frames (16x64); Texture's animation argument writes
// textures/block/ember_glow.png.mcmeta beside it, so the game plays them
// in turn, three ticks each, blending between frames.

/** A full cube showing the ember strip on every face. */
Model block/ember_glow {
  parent = block/cube_all
  textures = { all = Texture("art/ember_glow.png", animation = { frametime = 3, interpolate = true }) }
}

/** Every glowstone block shows the ember model. */
BlockState minecraft:glowstone {
  variants = { "" = { model = block/ember_glow } }
}
