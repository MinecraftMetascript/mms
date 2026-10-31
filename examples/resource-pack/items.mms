/** The item model a stack picks with the item_model component. */
ItemModel lantern_charm = { model = Model(model = block/glowstone_warm) }

/** Names for the item, in the resource pack's English translations. */
Lang en_us = { "item.glow.lantern_charm" = "Lantern Charm" }

/** A torch that looks like warm glowstone and carries the charm's name. */
Recipe lantern_charm = Shapeless(
  [torch, glowstone_dust],
  {
    id = torch
    components = {
      item_model = lantern_charm
      item_name = { translate = "item.glow.lantern_charm" }
    }
  },
)
