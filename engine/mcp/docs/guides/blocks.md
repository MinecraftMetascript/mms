# Retexturing and remodeling blocks

A pack can change how any vanilla block looks: its textures, and the shape of its model. It cannot change how the block behaves. Collision, the outline you aim at, light, sounds and what drops come from the game's code for each block, and neither a data pack nor a resource pack can change them. This page covers what a pack can change, what it can't, and the usual ways around the limits.

## New textures, same shape

[`Retexture`](resources.md#retexturing-a-vanilla-block) is the short way. Name the block, then give its textures by slot:

```mms
// Oak logs that look like birch.
Retexture minecraft:oak_log {
  textures = { side = block/birch_log, end = block/birch_log_top }
}
```

A texture is a vanilla texture id, as here, or a file of your own, such as `Texture("art/log_side.png")`. It can also be [animated](resources.md#animated-textures). `Retexture` covers every model the block shows: a log's upright and sideways models, each shape of a stair, each half of a door. When the block's item shows the block, it covers the item too. In `textures`, completion lists the block's slots. A slot the block doesn't have is an error that lists the slots it does have.

## New shapes

To change a block's shape, write the `Model` with its own `elements`, then a `BlockState` that shows it. A model's elements are boxes measured in sixteenths of a block, so `[0, 0, 0]` to `[16, 16, 16]` is the whole block:

```mms
/** A glass block with a thin oak post through the middle. */
Model block/glass_post {
  parent = block/glass
  elements = [
    { from = [0, 0, 0], to = [16, 16, 16], faces = {
      north = { texture = "#all", cullface = north }, south = { texture = "#all", cullface = south },
      east = { texture = "#all", cullface = east }, west = { texture = "#all", cullface = west },
      up = { texture = "#all", cullface = up }, down = { texture = "#all", cullface = down },
    } },
    { from = [6, 0, 6], to = [10, 16, 10], faces = {
      north = { texture = "#post" }, south = { texture = "#post" },
      east = { texture = "#post" }, west = { texture = "#post" },
    } },
  ]
  textures = { post = block/oak_planks }
}

BlockState minecraft:glass {
  variants = { "" = { model = block/glass_post } }
}
```

The model changes only the picture. The block keeps everything else its type has. Glass is still glass, solid stone is still solid stone, and a slab is still a slab. Two consequences decide which blocks you can reshape safely.

**Collision and the outline don't follow the model.** You walk into a stone block reshaped as a fence post as if it were a full cube, and you aim at the full cube when you break it. You can't remove a block's collision or give it a new one. A shape works best on a vanilla block with about the same collision: a slab for a low shape, a fence or wall for a post, a carpet or pressure plate for something flat, and a flower for something you walk through.

**Solid blocks hide their neighbours' faces.** A full, opaque block such as stone tells the game that the faces touching it can't be seen, so the game skips drawing them. Reshape stone as something smaller, and those faces stay hidden. You see through the gap into caves, or into the void. Reshape blocks the game already draws around, such as glass, leaves, slabs, stairs, fences and plants. A solid block can take a new shape only when that shape still covers each whole side of the cube.

Light follows the block too: stone blocks light whatever its model looks like, and glass lets it through.

## More than one look for a block

A pack has one block state file per block, so a block takes one `Retexture`. A second `Retexture` of the same block, or a `Retexture` together with your own `BlockState` for it, is a duplicate. There are three ways to get more looks out of one block.

**Random looks by position.** A block state can give a weighted list of models, and the game picks one for each block from its position. A wall of stone then shows a mix of looks, as vanilla does by rotating grass, dirt and sand. You can't choose which block gets which look.

```mms
Model block/mossy_patch { parent = block/cube_all, textures = { all = block/mossy_cobblestone } }

// Three in four stone blocks keep vanilla's look; the rest are mossy.
BlockState minecraft:stone {
  variants = {
    "" = [
      { model = block/stone, weight = 3 },
      { model = block/mossy_patch },
    ]
  }
}
```

**Looks by state.** A block state gives a model to each combination of the block's properties, such as `facing`, `axis` or `powered`, so each state can show a different model. This only gives you separate looks when nothing in the game changes those properties under you. Most properties do change, from placement, neighbours or redstone, which is why packs that do this pick blocks with many states that players rarely see, and set the state they want with `/setblock`.

**Display entities.** For as many looks as you want, at any size, leave the blocks alone and show the look with an `item_display` entity. Its item can carry an `item_model` component that names an `ItemModel` of your own. Put a block under it for collision: a barrier for a full cube, or a vanilla block of the right shape. A [custom item](resources.md#custom-items)'s `on_use_on_block` can place both:

```mms
Model block/crate {
  parent = block/cube_column
  textures = { end = block/barrel_top, side = block/barrel_side }
}
ItemModel crate { model = { type = minecraft:model, model = block/crate } }

/** Puts a crate on top of the block the player clicked. */
Function place_crate {
  /setblock ~ ~1 ~ barrier
  /summon item_display ~ ~1.5 ~ {Tags:["crate"],item:{id:"minecraft:stick",count:1,components:{"minecraft:item_model":"${crate}"}}}
}
```

Entities cost more than blocks, so keep this for things a world holds tens or hundreds of, not thousands. The block under the display is still that block. A barrier can't be broken in survival, so packs that use one add an `interaction` entity to detect the hit, and remove the barrier and the display in a function.

## Summary

| You want | Use |
| --- | --- |
| New textures, same shape | `Retexture` |
| A new shape where the vanilla collision fits | your own `Model` and `BlockState` |
| Random variety for one block | a weighted list in its `BlockState` |
| Several looks the player chooses | display entities, or states nothing changes |
| A different hitbox or behaviour | not possible from a pack; pick a vanilla block that already has it |
