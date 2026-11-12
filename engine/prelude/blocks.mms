// Retexture: new textures for a vanilla block, written in MMS on top of
// the models, block state and item model it emits.

/**
 * New textures for a vanilla block, by texture slot:
 * `Retexture minecraft:oak_log { textures = { side = Texture("art/log_side.png"), end = Texture("art/log_end.png") } }`.
 * It writes a model of the project's for each vanilla model the block
 * shows, such as a log's upright and sideways ones or each shape of a
 * stair, with that model as its parent and these textures. The block's
 * block state, and its item's model when that shows the block, then
 * show those models instead. Other blocks keep their look, even those
 * that share the vanilla models, such as waxed copper.
 */
Blueprint Retexture {
  /**
   * The new textures, by the slot of the block's models they fill:
   * `all` for a block with one texture, such as stone; `side` and `end`
   * for a log; `top`, `bottom` and `side` for a stair. Completion lists
   * the block's slots.
   */
  textures: {string: Texture}

  let vanilla = vanilla_block(self.id)
  let slot_list = join(vanilla.slots, ", ")
  let folder = "block/${self.path}/"

  check {
    if (len(vanilla.models) == 0) {
      error(self, "Retexture names a vanilla block, such as minecraft:${self.path}, and '${self.id}' is not one")
    } else {
      for slot, texture in textures {
        if (!contains(vanilla.slots, slot) && !contains(vanilla.other_slots, slot)) {
          error(textures, "${self.id} has no texture slot '${slot}'; its faces show ${slot_list}")
        }
      }
    }
  }

  emit {
    if (len(vanilla.models) > 0) {
      for path in vanilla.models {
        Model ${replace(path, "block/", folder)} {
          parent = "minecraft:${path}"
          textures = textures
        }
      }
      primary BlockState ${self.id} = rebase(vanilla.state, "minecraft:block/", folder)
      if (vanilla.item != none) {
        ItemModel ${self.id} = rebase(vanilla.item, "minecraft:block/", folder)
      }
    }
  }
}
