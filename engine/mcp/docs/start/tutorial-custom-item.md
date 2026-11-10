# Tutorial: a magic wand

This tutorial makes a Blink Wand: a stick that throws its holder six blocks forward on right click, with a cooldown, durability, a recipe and a place in a loot table. It takes about fifteen minutes, and each step builds, so you can try the wand in a world as you go.

Data packs cannot add items to the game. What they can do is give a vanilla item a name, a look and components of its own, and watch for players using it. By hand that takes an item model, a translation, an advancement per event and the same list of components at every `/give`, recipe and loot table. A `CustomItem` declaration writes all of it.

## 1. Set up the project

```sh
mkdir wands && cd wands
mms init -namespace wands
```

Everything below goes in the starter `main.mms`. Each step shows the declarations it changes: where a step shows `CustomItem blink_wand` or `Function blink` again, change the one you have instead of adding a second, and keep the fields from earlier steps. [The whole file](#the-whole-file) at the end is where the steps arrive. In the web editor, start a new project instead and set its namespace to `wands`.

## 2. Declare the item

```mms
/** A stick that throws its holder six blocks forward. */
CustomItem blink_wand {
  base = stick
  name = "Blink Wand"
  components = {
    rarity = epic
    lore = ["{gray}Right-click to blink forward."]
  }
}

/** /function wands:give hands one over. */
Function give {
  /give @s ${blink_wand}
}
```

- `base` is the vanilla item the wand is made from. It decides how the item behaves, so pick one that does nothing on right click, such as `stick`, `paper` or `amethyst_shard`.
- `name` is the item's name. It goes into the lang file, so it is plain text: the `rarity` component colours it.
- `components` holds any other [item components](https://minecraft.wiki/w/Data_component_format). Completion lists them, and `lore` lines are [text markup](../guides/resources.md#text-markup).
- `${blink_wand}` in a command stands for the whole item: `stick[item_name=..., rarity="epic", lore=[...]]`. Hover it to see what it writes.

Build into a world and try it:

```sh
mms build --world ~/.minecraft/saves/MyWorld
```

Open the world (or run `/reload` if it is open), then run `/function wands:give`.

## 3. Make it do something

An `on_*` field names a function to run when something happens with the item. `on_right_click` runs once per press, as the player and at the player:

```mms
CustomItem blink_wand {
  base = stick
  name = "Blink Wand"
  on_right_click = blink
}

/** Throws the player six blocks the way they are looking. */
Function blink {
  /particle portal ~ ~1 ~ 0.3 0.5 0.3 0.1 40
  /tp @s ^ ^ ^6
  /playsound entity.enderman.teleport player @a ~ ~ ~
}
```

A line that starts with `/` is a command, checked against the game's command tree as you type. `^ ^ ^6` is six blocks along the player's line of sight.

There are hooks for hitting, killing, eating, picking up, breaking and shooting too. [Item hooks](../guides/functions.md#item-hooks) lists them all and explains how right click is detected.

## 4. Add a cooldown

The wand is too strong when it can be spammed. A cooldown needs a number per player, which is what a scoreboard objective is:

```mms
/** Ticks until a player can blink again. */
Objective blink_cooldown

Function blink {
  if (blink_cooldown[@s] > 0) {
    actionbar(@s, "{red}The wand is cooling down")
    return
  }
  blink_cooldown[@s] = 40
  /particle portal ~ ~1 ~ 0.3 0.5 0.3 0.1 40
  /tp @s ^ ^ ^6
  /playsound entity.enderman.teleport player @a ~ ~ ~
}

OnTick {
  /scoreboard players remove @a[scores={blink_cooldown=1..}] blink_cooldown 1
}
```

- `blink_cooldown[@s]` is the score of whoever runs the function. It reads and writes like a variable.
- `OnTick` runs its body twenty times a second. Here it counts every running cooldown down by one, so 40 is two seconds.
- `actionbar` shows a line above the hotbar. It takes the same markup as `lore`.

## 5. Make it wear out

Two fields make each use cost something. `damage_on_use` takes durability, and the item breaks when it runs out. `consume_on_use` takes items from the stack instead.

```mms
CustomItem blink_wand {
  base = stick
  name = "Blink Wand"
  on_right_click = blink
  damage_on_use = 1
  components = {
    max_damage = 64
    rarity = epic
    lore = ["{gray}Right-click to blink forward."]
  }
}

Function blink {
  /tp @s ^ ^ ^6
}
```

A stick has no durability of its own, so the wand sets `max_damage`. Leave it out and MMS says so, on the `damage_on_use` line.

## 6. Hand it out

The wand's name works wherever an item stack goes, so a recipe result and a loot entry are both just `blink_wand`:

```mms
CustomItem blink_wand {
  base = stick
  name = "Blink Wand"
}

/** An ender pearl on a stick. */
Recipe blink_wand = Shaped([" E", "S "], { E = ender_pearl, S = stick }, blink_wand)

/** One chest in five holds a wand. */
LootTable chests/wizard {
  pools = [
    Pool(1, [Item(name = blink_wand, weight = 1), Item(name = ender_pearl, weight = 4)]),
  ]
}

/** Places a wizard's chest where you stand. */
Function chest {
  /setblock ~ ~ ~ chest{LootTable:"wands:chests/wizard"}
}
```

A loot table does nothing until something uses it. `/function wands:chest` places a chest that rolls this one when it is opened. A `Recipe` and a `CustomItem` can share a name, because different kinds never clash.

One thing to know: the game matches recipe ingredients and item tags by item id only. A custom item written there means its base item, so a recipe that takes `blink_wand` as an ingredient accepts any stick. The editor marks each such use with a hint.

## 7. Give it a look

Without a texture the wand looks like a stick. To draw your own, save a 16×16 `.png` in the project, here as `art/blink_wand.png`, and name it by its path:

```
CustomItem blink_wand {
  base = stick
  name = "Blink Wand"
  texture = Texture("art/blink_wand.png")
  parent = item/handheld
}
```

`parent = item/handheld` holds the item like a tool. MMS writes the model, the item model definition and the texture into a resource pack next to the data pack, and `mms build --world` installs both. In the web editor, **+ Add texture or sound** adds the file under its own name, so the path there is `Texture("blink_wand.png")`. See [Custom items](../guides/resources.md#custom-items) for the other ways to pick a model.

## The whole file

```mms
/** Ticks until a player can blink again. */
Objective blink_cooldown

/** A stick that throws its holder six blocks forward. */
CustomItem blink_wand {
  base = stick
  name = "Blink Wand"
  on_right_click = blink
  damage_on_use = 1
  components = {
    max_damage = 64
    rarity = epic
    lore = ["{gray}Right-click to blink forward."]
  }
}

/** Throws the player forward, unless the wand is still cooling down. */
Function blink {
  if (blink_cooldown[@s] > 0) {
    actionbar(@s, "{red}The wand is cooling down")
    return
  }
  blink_cooldown[@s] = 40
  /particle portal ~ ~1 ~ 0.3 0.5 0.3 0.1 40
  /tp @s ^ ^ ^6
  /playsound entity.enderman.teleport player @a ~ ~ ~
}

OnTick {
  /scoreboard players remove @a[scores={blink_cooldown=1..}] blink_cooldown 1
}

/** An ender pearl on a stick. */
Recipe blink_wand = Shaped([" E", "S "], { E = ender_pearl, S = stick }, blink_wand)

/** One chest in five holds a wand. */
LootTable chests/wizard {
  pools = [
    Pool(1, [Item(name = blink_wand, weight = 1), Item(name = ender_pearl, weight = 4)]),
  ]
}

/** Places a wizard's chest where you stand. */
Function chest {
  /setblock ~ ~ ~ chest{LootTable:"wands:chests/wizard"}
}

/** /function wands:give hands one over. */
Function give {
  /give @s ${blink_wand}
}
```

## Next steps

- Make a second wand. Once two items share their cooldown code, a [Blueprint](../guides/blueprints.md) turns "a wand" into a kind of declaration of its own.
- [Keeping state](../guides/state.md) covers scores, Globals and Storage, for when a cooldown is not enough.
- The [`embers`](../../examples/embers/) example has items with their own textures, a use cost and a block hook, and [`arrows`](../../examples/arrows/) has arrows that do things where they land.
