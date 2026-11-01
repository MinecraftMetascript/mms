# Embers

A small pack that brings its own art. Everything under `art/` and `sounds/` was made for this example.

| File | Used by | Becomes |
| --- | --- | --- |
| `art/ember_shard.png`, `art/ember_disc.png` | item models in `items.mms` | `textures/item/` |
| `art/hearth.png` (32x16) | the `hearth` painting in `painting.mms` | `textures/painting/hearth.png` |
| `sounds/embers.ogg` | the `embers` jukebox song in `music.mms` | `sounds/embers.ogg` and `sounds.json` |
| `art/icon.png` | `icon` in `mms.json` | `pack.png` in both packs |

Try it in a world:

```
mms build --world ~/.minecraft/saves/MyWorld
```

That installs the data pack and writes the resource pack to the world's `resources.zip`. Data packs cannot add items to the creative menu, so run `/function embers:give` for one of everything. In survival, craft blaze powder with an amethyst shard for two Ember Shards, or ring any music disc with four blaze powder for the Embers disc. `/loot give @s loot embers:ember_cache` gives a few shards. Hand-placed paintings can now pick the Hearth.

The shards are amethyst shards with an `item_model` and a name. Recipe ingredients match by item id only, so a recipe cannot ask for an Ember Shard instead of any amethyst shard.
