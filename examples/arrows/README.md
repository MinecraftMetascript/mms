# Arrows

Custom projectiles: three new arrows, a bow that calls lightning, an exploding splash potion, and a wand that fires cows.

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
| Blast Potion | glass bottle + TNT | A splash potion that explodes where it lands. |
| Cow Wand | 2 sticks | Right-click to fire a cow the way you are looking. |

The arrows work from bows, crossbows and dispensers.

## How it works

| File | What it does |
| --- | --- |
| `items.mms` | The custom items and their recipes. |
| `effects.mms` | What the Explosive and Ender arrows and the Thunder Bow do where an arrow hits. |
| `launch.mms` | The Cow Wand, and a `launch` function that summons an entity already flying the way the player looks. |
| `lang.mms` | Names for the arrows. |

**Effects in the arrow** (Arrow of Levitation). A tipped arrow applies its `potion_contents` to what it hits, and `custom_effects` can be any effects at all. No commands run. `potion_duration_scale = 1` makes the effect last as long as written; tipped arrows normally give an eighth.

**Projectile hooks** (Explosive and Ender arrows, Thunder Bow). `on_projectile_hit_block` and `on_projectile_hit_entity` name a function to run where an arrow hits. On an arrow base they follow that arrow; on a bow or crossbow base they follow every arrow it shoots. A block hit runs as the arrow at the point it hit (`execute on origin` is the shooter); an entity hit runs as and at the entity (`execute on attacker` is the shooter).

Under the hood, an arrow keeps the bow it was shot from as its `weapon`, and the game runs that weapon's enchantment effects at the exact spot it hits. MMS writes an enchantment per item that runs the hook functions, and the tick function adds it to the weapon of each new arrow that is the item or came from it. The hooks combine: an Explosive Arrow from the Thunder Bow explodes and calls lightning. See [Projectile hooks](../../docs/functions.md#projectile-hooks).

**Thrown potions** (Blast Potion). A thrown potion keeps no bow, so no enchantment runs when it lands. `on_projectile_land` puts a marker on each one instead, riding it; the game drops a rider where it removes the potion, and the hook runs at that spot. The spot can be up to one tick of flight short, about half a block for a thrown potion, and the hook can't tell what the potion hit or who threw it.

The custom arrows are tipped arrows on purpose. Their colour shows on the item and as a particle trail in flight, a bow with Infinity does not make them free (it only spares plain `arrow` items), and they take their name from `potion_contents.custom_name` (see `lang.mms`).

**Launching an entity** (Cow Wand). An entity flies when its `Motion` is set. The usual trick for a vector that points where the player looks is to summon a marker at `^ ^ ^speed` from the world origin and read its `Pos`. `launch.mms` does that with three fixes that the trick often misses:

- The origin chunk is kept loaded (`forceload add 0 0` in the overworld on load), and the marker is always summoned in the overworld. Spawn chunks are no longer kept loaded, so this can't be assumed.
- The marker is killed once read. Without that every shot leaves one behind.
- `launch` reads `^ ^ ^` with `anchored feet`. Under `anchored eyes`, local coordinates start at the eyes of whoever runs the command, which tilts the vector up.

`launch` then summons the entity with that vector as its `Motion`, through a macro, and faced the way the player looks. Setting `Motion` with `/data` after summoning also works on the server, but players were told about the entity when it appeared standing still, so it seems to hang for a moment before it moves.

## Limits

- **Not tested in game yet.**
- **Dispensers.** An arrow from a dispenser is followed, but the entity-hit hook may not run for it, since the game may want a shooter to run a weapon's attack effects. Untested.
- **The Thunder Bow** works only on arrows. A trident or a crossbow's firework is not followed.
