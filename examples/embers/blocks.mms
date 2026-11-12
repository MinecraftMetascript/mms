// Glowstone, retextured as a block of flickering embers. The image is a
// strip of four 16x16 frames (16x64); Texture's animation argument writes
// textures/block/ember_glow.png.mcmeta beside it, so the game plays them
// in turn, three ticks each, blending between frames.

/**
 * Glowstone shows the ember strip on every face. Retexture writes a model
 * with vanilla's glowstone model as its parent, and glowstone's block
 * state and item model, which show it.
 */
Retexture minecraft:glowstone {
  textures = { all = Texture("art/ember_glow.png", animation = { frametime = 3, interpolate = true }) }
}
