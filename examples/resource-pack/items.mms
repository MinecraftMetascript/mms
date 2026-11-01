/**
 * A torch that looks like warm glowstone and carries the charm's name.
 * The CustomItem writes the item model definition that points at the
 * block model, and the "Lantern Charm" name in the English lang file.
 */
CustomItem lantern_charm = {
  base = torch
  name = "Lantern Charm"
  model = block/glowstone_warm
}

Recipe lantern_charm = Shapeless([torch, glowstone_dust], lantern_charm)
