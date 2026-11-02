# Embers

A small pack that brings its own art. Everything under `art/` and `sounds/` was made for this example.

| File | Used by | Becomes |
| --- | --- | --- |
| `art/ember_shard.png`, `art/ember_disc.png` | the custom items in `items.mms` | `textures/item/`, with their models |
| `art/hearth.png` (32x16) | the `hearth` painting in `painting.mms` | `textures/painting/hearth.png` |
| `sounds/embers.ogg` | the `embers` jukebox song in `music.mms` | `sounds/embers.ogg` and `sounds.json` |
| `art/icon.png` | `icon` in `mms.json` | `pack.png` in both packs |

Try it in a world:

```
mms build --world ~/.minecraft/saves/MyWorld
```

That installs the data pack and writes the resource pack to the world's `resources.zip`. Data packs cannot add items to the creative menu, so run `/function embers:give` for one of everything (commands need cheats on, or operator permissions on a server). In survival, craft blaze powder with an amethyst shard for two Ember Shards, or ring a creeper-drop music disc with four blaze powder for the Embers disc. `/loot give @s loot embers:ember_cache` gives a few shards. Hand-placed paintings can now pick the Hearth.

The shards are amethyst shards with an `item_model` and a name: each `CustomItem` in `items.mms` writes the model, the item model and the translation, and its name stands for the item in recipes, loot and `/give`. Recipe ingredients match by item id only, so a recipe cannot ask for an Ember Shard instead of any amethyst shard.

Right-click with an Ember Shard to puff out flames and get a few seconds of fire resistance. Each use takes one shard. Right-click a block with the Ember Wand to light a fire on top of it. Each fire costs the wand one point of its 32 durability, and it breaks when that runs out. In `items.mms`, the shard sets `on_right_click = warm` and `consume_on_use = 1`. The wand sets `on_use_on_block = kindle`, `damage_on_use = 1` and `max_damage = 32`. Both get a `consumable` component that never finishes, so right click counts as a use, and a `using_item` advancement runs the functions once per press.
