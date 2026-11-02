# Arrows

Custom projectiles: three new arrows, a bow that calls lightning, and a wand that fires cows.

## Try it

```
mms build --world ~/.minecraft/saves/MyWorld
```

Then run `/function arrows:give` for one of everything, or craft them:

| Item | Recipe | What it does |
| --- | --- | --- |
| Explosive Arrow | arrow + gunpowder | Explodes where it lands, at half the strength of TNT. |
| Ender Arrow (2) | 2 arrows + ender pearl | Teleports whoever shot it to where it lands. |
| Arrow of Levitation (8) | 8 arrows around a shulker shell | Lifts what it hits for three seconds. |
| Thunder Bow | bow + lightning rod | Calls lightning wherever its arrows hit, block or mob. Any arrow works. |
| Cow Wand | 2 sticks | Right-click to fire a cow the way you are looking. |

The arrows work from bows, crossbows and dispensers.

## How it works

| File | What it does |
| --- | --- |
| `items.mms` | The custom items and their recipes. |
| `flight.mms` | Finds each custom arrow when it is shot and notices where it lands. |
| `effects.mms` | What the Explosive and Ender arrows do where they land. |
| `bow.mms` | The Thunder Bow's enchantment. |
| `launch.mms` | The Cow Wand, and a `launch` function that throws an entity the way it faces. |
| `lang.mms` | Names for the arrows and the enchantment. |

There are three ways to give a projectile an effect, and this example uses each one.

**A tipped arrow's own effects** (Arrow of Levitation). A tipped arrow applies its `potion_contents` to what it hits, and `custom_effects` can be any effects at all. No commands run. `potion_duration_scale = 1` makes the effect last as long as written; tipped arrows normally give an eighth.

**A bow's enchantment** (Thunder Bow). An arrow remembers the bow it was shot from, and when it hits, the game runs that bow's enchantment effects at the exact spot: `hit_block` for a block, `post_attack` for a mob. `bow.mms` declares an enchantment that summons lightning in both, and the bow has it in its `enchantments`. Enchantment effects can also explode, apply effects, play sounds, or run a function. This is exact and needs no tick function, but it belongs to the bow, not the arrow.

**Watching the arrow** (Explosive and Ender arrows). The custom arrows are tipped arrows marked with `custom_data`, which the arrow entity keeps in its `item`. Each tick, `flight.mms`:

1. looks once at each new arrow. A custom one gets a marker riding on it as a stand-in, and notes who shot it (`execute on origin`) as a number in the `shooter_id` objective.
2. finds custom arrows stuck in a block (`inGround`), and runs the arrow's effect there.
3. finds stand-ins that are no longer riding anything. The game removes an arrow that hits a mob and drops its passengers where they are, so the stand-in marks where the arrow was.

Tipped arrows were picked as the base on purpose. Their colour shows on the item and as a particle trail in flight, a bow with Infinity does not make them free (it only spares plain `arrow` items), and they take their name from `potion_contents.custom_name` (see `lang.mms`).

**Launching an entity** (Cow Wand). An entity flies when its `Motion` is set. The usual trick for a vector that points where the player looks is to summon a marker at `^ ^ ^speed` from the world origin and read its `Pos`. `launch.mms` does that with three fixes that the trick often misses:

- The origin chunk is kept loaded (`forceload add 0 0` in the overworld on load), and the marker is always summoned in the overworld. Spawn chunks are no longer kept loaded, so this can't be assumed.
- The marker is killed once read. Without that every shot leaves one behind.
- `launch` reads `^ ^ ^` with `anchored feet`. Under `anchored eyes`, local coordinates start at the eyes of whoever runs the command, which tilts the vector up.

`launch` runs as the entity being thrown (`execute summon cow run ...`), so it never has to find it again by tag. It also turns the entity to face the way it flies.

## Limits

- **Not tested in game yet.**
- **Mob hits land short.** The stand-in sits where the arrow was at the end of the tick before, so on a mob hit the explosion or teleport happens up to one tick of flight before the mob: about three blocks at full draw. Block hits are exact. The Thunder Bow's way is exact for both.
- **Point-blank shots** that hit a mob in the tick they are fired are never seen, so they do nothing special.
- **Arrows that leave the world** (into the void, or despawned) set off their effect where they vanish.
- **The Thunder Bow** loses its enchantment in a grindstone, like any enchantment that isn't a curse.
