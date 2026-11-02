# Barrels

A storage barrel in the style of the drawer and barrel mods: it holds a lot of one kind of item, and its front shows the item and how many there are.

Data packs cannot add blocks, so a Storage Barrel is a vanilla barrel block with three entities on its front face:

| Entity | Tag | Job |
| --- | --- | --- |
| interaction | `barrels.hit` | Catches right clicks and punches on the front. Keeps the count and stack size as scores. |
| item display | `barrels.face` | Shows the stored item, like an item frame would. It is the one place the item itself is kept. |
| text display | `barrels.label` | Shows the count under the item: grey "Empty", white while there is room, gold when full. |

A shared `barrel_id` score ties the three together.

## Try it

```
mms build --world ~/.minecraft/saves/MyWorld
```

Then `/function barrels:give` for four Storage Barrels, or craft one from a barrel with a chest on each side.

- **Place** it like an item frame. On a wall it faces away from the wall; on a floor or ceiling it faces you.
- **Right-click** the front with an item to put the whole stack in. An empty barrel takes the kind of the first item that goes in. Items must match exactly (id and components), so a named sword and a plain one are different items.
- **Right-click with an empty hand** to put in every matching item from your inventory.
- **Punch** the front to take out a stack; **sneak and punch** to take out one. They drop at your feet and are picked up straight away.
- **Break** the barrel by mining its sides or edges. Everything inside spills out, and in survival it drops a Storage Barrel.

A barrel holds 32 stacks: 2048 of an item that stacks to 64, or 512 of one that stacks to 16. Items that stack to 1, or to some other size, do not fit.

## How it works

| File | What it does |
| --- | --- |
| `item.mms` | The Storage Barrel item (a `CustomItem`), its recipe and its loot table. |
| `place.mms` | Turns a placed Storage Barrel into a barrel block and its front, and notices when the block is gone. |
| `use.mms` | Putting items in and taking them out. |
| `state.mms` | The scores every barrel keeps, and the Globals for the one being used. |

**Placing.** The item is an item frame that looks like a barrel (`item_model`) and places an invisible frame tagged `barrels.place` (`entity_data`). Item frames are a common trick for this, because the frame lands in the block in front of the clicked face and records which way it faces. The tick function swaps the frame for a barrel block, facing out, and summons the three entities.

**Clicking.** Right-clicking an interaction entity fires the `player_interacted_with_entity` trigger, and punching it fires `player_hurt_entity`. Each `OnAdvancement` finds which barrel the player clicked with `execute on target` or `on attacker`, the relations an interaction entity keeps for its last click.

**Comparing items.** The player's stack is copied to a scratch item display and set to a count of 1. Then `data modify ... set from` copies the stored item over it, and `execute store success` reports whether that changed anything: no change means the same item, components included.

**Counting.** `SetCount` item modifiers take their count from a score (`Score(Fixed("$amount"), "mms.barrels")`), which is how a stack loses exactly what fit, and how a stack of the right size comes out.

**Breaking.** Each tick checks that every barrel still has its block. When one is gone, the contents drop as stacks and the plain barrel item the block dropped is swapped for a Storage Barrel.

The barrel block is locked with a key no item matches, so its own 27 slots stay shut: right-clicking its sides says "Barrel is locked".

## Limits

- **Hoppers** do not feed or empty a Storage Barrel. A hopper can still push items into the locked barrel block's own slots, and they drop when the barrel breaks.
- **Stack sizes** other than 16 and 64 are refused, because an item predicate can test for a stack size but a command cannot read it.
- **The front** only works from the front: the clickable box covers the middle of that face.
