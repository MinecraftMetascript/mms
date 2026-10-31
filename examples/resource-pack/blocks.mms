/**
 * Glowstone picks one of three looks per block. Each variant names a
 * model; weight makes the plain one twice as likely, and y turns it.
 */
BlockState minecraft:glowstone = {
  variants = {
    "" = [
      {
        model = block/glowstone_plain
        weight = 2
      },
      {
        model = block/glowstone_plain
        y = 90
      },
      { model = block/glowstone_warm },
    ]
  }
}

// A cube with one texture on every face. A file of your own would be
// Texture("art/glow.png", "block/glow") instead of a vanilla texture id.
Model block/glowstone_plain = {
  parent = block/cube_all
  textures = { all = block/glowstone }
}

Model block/glowstone_warm = {
  parent = block/cube_all
  textures = { all = block/shroomlight }
}

/** Redstone lamps keep their lit and unlit looks, chosen by the lit property. */
BlockState minecraft:redstone_lamp = {
  variants = {
    "lit=false" = { model = block/redstone_lamp }
    "lit=true" = { model = block/glowstone_warm }
  }
}

/** Multipart: the post always shows, and each side shows when it connects. */
BlockState minecraft:oak_fence = {
  multipart = [
    { apply = { model = block/oak_fence_post } },
    {
      when = { north = "true" }
      apply = {
        model = block/oak_fence_side
        uvlock = true
      }
    },
    {
      when = { east = "true" }
      apply = {
        model = block/oak_fence_side
        y = 90
        uvlock = true
      }
    },
    {
      when = { south = "true" }
      apply = {
        model = block/oak_fence_side
        y = 180
        uvlock = true
      }
    },
    {
      when = { west = "true" }
      apply = {
        model = block/oak_fence_side
        y = 270
        uvlock = true
      }
    },
  ]
}
