# Functions

`Function` bodies and event hooks (`OnLoad`, `OnTick`, `OnFirstJoin`, `OnRespawn`, `OnAdvancement`) compile to `.mcfunction` files.

## Raw commands

A line that starts with `/` is a command. It is written to the output as it is, without the slash:

```
Function greet {
  /say hello
  /execute as @a at @s run particle minecraft:heart ~ ~2 ~
}
```

Every command is checked against the 26.3 command tree, the same one the game uses. The checker reports:

- unknown commands and options, with the alternatives
- missing and extra arguments
- argument syntax: numbers and their bounds, selectors and their options, coordinates, ids, ranges, and balanced NBT and text
- unknown vanilla ids, such as `minecraft:diamnd`
- ids in your own namespaces that nothing declares, such as `function demo:missing`

Completion inside a command offers the next literal and values for the argument being typed. These include selectors, game modes, registry ids, and your project's functions, tags and objectives.

The checker checks syntax, not meaning. Whether an entity exists, or whether an NBT path fits the entity, is left to the game.

## Values and variables

Function code computes with whole numbers (`int`, 32-bit, wrapping like the game's scoreboards) and `bool`s.

```
Global deaths: int = 0          // one value for the whole pack
Objective kills = deathCount    // a score per player

Function on_death {
  deaths += 1
  kills[@s] = 0                 // a score: objective[holder]
  let streak = kills[@s] * 2    // a local; its type is inferred
  let lucky: bool = streak > 10
}
```

- **Globals** are set once, when the pack first loads. A `/reload` keeps their values.
- **Locals** (`let`) belong to their function. A name in an inner block can shadow an outer one.
- **Objective scores** are written `objective[holder]`. The holder is a selector, or a fake player name in quotes, such as `kills["#total"]`.

### Operators

| Operators | Meaning |
|---|---|
| `+ - * / %` | Arithmetic. Division rounds down and `%` takes the divisor's sign, as the scoreboard does. |
| `< <= > >= == !=` | Comparisons. |
| `&& \|\| !` | Logic. The right side of `&&` and `\|\|` only runs when it is needed. |

Constant parts are folded at compile time.

### Strings

`string` values live in command storage (`<namespace>:mms`). They can be joined with `+`, appended with `+=`, and compared with `==` and `!=`. Ints and bools join as their text.

```
Global motd: string = "Welcome"

Function welcome(name: string) {
  let line = motd + ", " + name + "!"
  if (name == "Notch") {
    line += " Nice to meet you."
  }
  /tellraw @a "${line}"
}
```

## Interpolation

`${expr}` puts a value into a command:

```
Function report(points: int) {
  /say ${points} points, ${points * 2} doubled
  /scoreboard players set @s kills ${points}
}
```

- **Constant values** (literals, constant arithmetic) are written into the command at compile time.
- **Run-time values** turn the line into a [macro](https://minecraft.wiki/w/Function_(Java_Edition)#Macros): the values are copied into storage, and a one-line macro function runs the command.
- **Quotes:** a string with a `"` in it will break a command that wraps it in quotes, just as it would in a hand-written macro.

## Control flow

```
Function classify(x: int): int {
  if (x > 100) {
    return 2
  } else if (x > 0) {
    return 1
  }
  return 0
}

Function count {
  for i in 1..10 {          // both ends included, like 1..10 in a selector
    if (i % 2 == 0) {
      continue
    }
    /say odd
  }
  let n = 0
  while (n < 5) {
    n += 1
    if (n == 3) {
      break
    }
  }
}
```

A function with a return type must end with `return` on every path.

## Calls

A call looks like `add(1, 2)`, or `other_ns:add(1, 2)` for a function in another namespace. Arguments are passed by position. A call can stand on its own, or be used as a value when the function has a return type.

Parameters are `int`, `bool` or `string`. Functions return `int` or `bool`. To hand a string back, assign it to a Global.

## Execute blocks

A chain of execute helpers followed by a block runs the block under `execute`:

```
Objective hits

Function zap {
  As(@e[tag=tower]).At(@s) {
    /particle minecraft:flame ~ ~1 ~
    hits[@s] += 1
  }
  Positioned(~ ~1 ~) {
    /say above
  }
  Execute("if entity @e[tag=boss]").In(the_nether) {
    /say the boss is here
  }
}
```

| Helper | Subcommand |
|---|---|
| `As(targets)` | `as` |
| `At(targets)` | `at` |
| `Positioned(pos)`, `Positioned(targets)` | `positioned`, `positioned as` |
| `Rotated(rot)`, `Rotated(targets)` | `rotated`, `rotated as` |
| `Facing(pos)`, `Facing(targets, anchor)` | `facing`, `facing entity` (the anchor defaults to `eyes`) |
| `In(dimension)` | `in` |
| `Anchored(eyes)` | `anchored` |
| `Align(xyz)` | `align` |
| `On(relation)` | `on` |
| `Summon(entity)` | `summon` |
| `Execute("...")` | any subcommands, written as text |

Arguments are written into the command as they appear: selectors as `@e[tag=x]`, coordinates as `~ ~1 ~` or `0, 64, 0`, ids as `the_nether`. The chain is checked against the command tree like a raw command, and `Execute("...")` covers conditions such as `if block` and anything without a helper.

A body with one command is written inline, as `execute as @a at @s run particle ...`. A longer body becomes a sub-function, such as `demo:zap/_exec0`.

The body can use the function's variables. They are not per entity: a body that runs for five entities updates one `n` five times. Keep per-entity values in an objective, such as `hits[@s]`.

`break`, `continue` and `return` cannot leave an execute block, because the block may run once per entity. Set a variable inside it and act on it afterwards.

## Conditions on the world

The tests of `execute if` are bool values, written as helpers with the arguments in the order the subcommand takes them:

```
Function check {
  if (Block(~ ~-1 ~, stone) && !Entity(@a[distance=..8])) {
    /say nobody is near this stone
  }
  let armed: bool = Items(@s, weapon.mainhand, bow) || Data(@s::Inventory[0])
  if (At(@p).Entity(@e[type=zombie, distance=..4])) {
    /say a zombie is near the closest player
  }
}
```

| Helper | Subcommand |
|---|---|
| `Block(pos, block)` | `if block` |
| `Blocks(start, end, destination, all \| masked)` | `if blocks` |
| `Entity(targets)` | `if entity` |
| `Biome(pos, biome)` | `if biome` |
| `Dimension(dimension)` | `if dimension` |
| `Loaded(pos)` | `if loaded` |
| `Data(target::path)`, `Data(target, path)` | `if data entity`, `storage` or `block`, by how the target is written |
| `Items(target, slot, item)` | `if items entity` or `block` |
| `Predicate(predicate)` | `if predicate` (a bare name is one of this project's predicates) |
| `Function(function)` | `if function` |
| `Execute("if ...")` | any test, written as text |

A condition may start with context helpers, such as `At(@p).Entity(...)` or `Positioned(~ ~-1 ~).Block(~ ~ ~, air)`. The test then runs in that context, and the statement it guards does not: `if (At(@p).Block(~ ~-1 ~, lava)) { /say danger }` says it where the function runs. A chain of tests alone, such as `Block(...) && Entity(...)`, becomes one `execute if ... if ...`; a chain with context becomes a flag that is set under `execute at @p ...` and tested after.

The same helpers work as execute blocks: `Entity(@e[tag=boss]) { ... }` runs the block when the boss exists.

`Count(targets)` is the number of entities a selector matches, as an int.

## NBT

`target::path` reads and writes NBT. The target is an entity (a selector), a storage (`namespace:name`, or a bare name for a storage in this namespace) or a block (coordinates). Reads are ints: `data get` rounds decimals down, so multiply first to keep them.

```
Function nbt {
  let y = @s::Pos[1]                       // data get entity @s Pos[1]
  let cents = @s::Pos[1] * 100             // data get entity @s Pos[1] 100
  let next = siege:path::queue[0].x        // data get storage siege:path queue[0].x
  let items = ~ ~-1 ~::Items               // the number of entries in a list

  path::count = next + 1                   // storage demo:path, written as an int
  path::name = "bob"                       // a string
  path::done = y > 60                      // a bool, written as a byte
  @s::Health = Float(20)                   // a typed write
  @s::Motion[1] = Double(cents)            // from a variable, through execute store
  @s::Pos = @e[tag=m, limit=1]::Pos        // a copy keeps the source's type
  path::count += 1                         // read, add, write back as an int
}
```

`Byte(x)`, `Short(x)`, `Int(x)`, `Long(x)`, `Float(x)` and `Double(x)` set the type of a written number; without one, numbers are ints and bools are bytes. A compound operator such as `+=` writes an int.

A list index is a whole number, counting from the end when negative: `Inventory[-1]`. For any other path form, write the path in quotes: `@s::"Inventory[{Slot:0b}].id"`. Lists and compounds are not written from code yet; use `/data modify`.

The checker checks the selector and the path's syntax. Whether the path fits the entity is left to the game.

## Commands as values

A command on the right of `=` or after `return` runs and gives its result, or its success when the value is a bool:

```
Function count(): int {
  let n = /execute if entity @e[tag=enemy]       // store result
  let has_path: bool = /data get storage demo:path queue  // store success
  return /scoreboard players get #best demo_score
}
```

The command runs to the end of the line, so it cannot sit inside a longer expression; store it in a variable first. `${...}` is not supported inside one.

## Built-in functions

| Function | Value |
|---|---|
| `Count(targets)` | the number of entities a selector matches |
| `Abs(x)` | the absolute value |
| `Min(a, b, ...)`, `Max(a, b, ...)` | the smallest or largest of the values |

A declared `Function` with one of these names takes precedence.

## Unset scores

A score on an `Objective` that was never set reads as 0: a comparison such as `kills[@s] == 0` or `kills[@s] < 5` holds for a player with no score, and `let n = kills[@s]` gives 0. The scoreboard itself fails such a test, so the generated command is the complement (`execute unless score @s kills matches 1..`), and a copy from a score sets 0 first. Comparing two objective scores with each other still needs both to be set, as `execute if score a < b` does.

## Event hooks

These declarations run their body when something happens, as the player it happens to, at their position:

```
Objective banners

/** Counts every banner a player places. */
OnAdvancement placed_banner = PlacedBlock(location = LocationCheck(predicate = { block = { blocks = #banners } })) {
  banners[@s] += 1
}

OnFirstJoin {
  /give @s minecraft:bread 4
}

OnRespawn {
  /effect give @s minecraft:regeneration 5 1
}
```

- **`OnAdvancement name = Trigger(...) { ... }`** runs each time a player meets a trigger. The trigger is one of the criterion helpers, such as `PlacedBlock`, `PlayerKilledEntity`, `InventoryChanged` or `Tick`, and its named arguments are the conditions. It writes the advancement `<namespace>:<name>` and a reward function, `<namespace>:mms_gen/advancement/<name>`, which revokes the advancement and then runs the body, so it fires again next time. Commands can refer to it, as in `/advancement grant @s only demo:placed_banner`.
- **`OnFirstJoin { ... }`** runs once per player, the first time they are in the world with the pack. It uses a tick advancement, `<namespace>:mms_gen/first_join`, that is never revoked; `/advancement revoke @s only <namespace>:mms_gen/first_join` makes it run again.
- **`OnRespawn { ... }`** runs right after a player respawns. A `deathCount` objective, `mms_respawn.<namespace>`, notes the death, and the tick function runs the body once the player is alive again.

### Item hooks

A [custom item](resources.md)'s `on_*` fields each name a function to run when something happens with the item. The functions take no parameters and can be in any file. A field can also hold the commands themselves in a `{ ... }` block, written like a function body (each command on its own line). That block is compiled as a function of its own, under `mms_gen/<item>/`. They run as the player, at the player unless the table says otherwise:

```
CustomItem ember_wand = {
  base = stick
  name = "Ember Wand"
  on_right_click = cast_ember
  on_use_on_block = light_block
}

Function cast_ember {
  /particle flame ^ ^1.5 ^1 0.2 0.2 0.2 0.05 20
}

Function light_block {
  /setblock ~ ~1 ~ fire keep
}
```

| Field | Runs when the player... |
| --- | --- |
| `on_right_click` | right-clicks with the item, once per press |
| `on_hold_right_click` | holds right click with the item: every tick, from the first |
| `on_use_on_block` | right-clicks a block with the item, once per press. It runs positioned at the centre of the block |
| `on_right_click_entity` | right-clicks an entity with the item, once per press. It runs at the entity, which has the tag `mms_target` while the function runs, so `@e[tag=mms_target,limit=1]` is the entity |
| `on_hit` | hurts an entity while holding the item in their main hand |
| `on_kill` | kills an entity while holding the item in their main hand |
| `on_consume` | finishes eating or drinking the item. It needs a `consumable` component, or a base item that is food or a potion |
| `on_pickup` | picks the item up from the ground |
| `on_break` | uses up the item's durability, so it breaks. It needs a base with durability, or a `max_damage` component |
| `on_projectile_hit_block` | shoots an arrow that hits a block (see [Projectile hooks](#projectile-hooks)). It runs as the arrow, at the point it hit |
| `on_projectile_hit_entity` | shoots an arrow that hurts an entity. It runs as and at the entity hit |

Every hook marks the item with `mms_item: "<namespace>:<name>"` in its `custom_data`, wherever it is used: in `/give`, recipes and loot. A `custom_data` of your own keeps its keys and gains the marker. Each hook is an advancement, `<namespace>:mms_gen/item_hook/<name>/<hook>`, whose trigger matches the base item with that marker, and a reward function of the same name that revokes it and calls your function.

**Right click.** The four right-click hooks use the consumable method from the [datapack.wiki guide](https://datapack.wiki/guide/right-click/consumable_component):

- The item gets a `consumable` component that takes about eleven days to eat, with no animation, sound or particles, so right click "uses" it without using it up. It also gets `use_effects`, so the player is not slowed down while holding the button. Set `consumable` or `use_effects` in the item's `components` to use your own instead.
- A `using_item` advancement, `.../press`, fires on every tick the button is held. `on_hold_right_click` runs on each of those ticks. For the others, a score, `mms_click.<namespace>`, notes that the player is still holding the button. The tick function counts it down, and they run only once it has reached 0.
- `on_use_on_block` and `on_right_click_entity` step a ray from the player's eyes, 0.1 blocks at a time, out to their `block_interaction_range` and `entity_interaction_range`. It stops at the first block or entity it meets. The ray counts a whole block as solid, so it can stop short of a slab or a flower's actual shape, and it skips items, arrows, markers and display entities.
- `on_right_click` runs on every press, whether or not the player is looking at a block or entity.

**Use cost.** Two more fields make each right-click use cost something:

```
CustomItem ember_wand = {
  base = blaze_rod
  on_use_on_block = kindle
  damage_on_use = 1
  components = { max_damage = 32 }
}
```

- `consume_on_use = n` takes `n` of the item from the stack.
- `damage_on_use = n` costs `n` points of durability. When the durability runs out, the item breaks with the break sound, and `on_break` runs. The item needs a `max_damage` component. MMS doesn't know vanilla items' durability, so a base item's own durability is not enough. A damageable item cannot stack, so the item gets `max_stack_size = 1` unless it sets one, and any other stack size is an error.

The cost is paid once per press, from the hand that holds the item. With `on_right_click` it is paid on every press. Otherwise it is paid only when `on_use_on_block` or `on_right_click_entity` runs, so a click at the air costs nothing. `on_hold_right_click` alone pays no cost.

A base item that already does something on right click keeps doing it, and then none of the right-click hooks runs. A block is still placed and a bow still draws, and a chest still opens when clicked. Pick a base such as `stick`, `paper` or `amethyst_shard` that does nothing on its own. A `use_cooldown` component has no effect, because the item is never finished being used. Keep a cooldown in a score instead. `on_consume` cannot be combined with the right-click hooks, because they stop the item from ever being finished.

### Projectile hooks

`on_projectile_hit_block` and `on_projectile_hit_entity` follow arrows. What they follow depends on the item's base:

- **A bow or crossbow**: every arrow shot from the item, whatever the arrow is.
- **An `arrow`, `tipped_arrow` or `spectral_arrow`**: the item itself, once it is shot from any bow, crossbow or dispenser. Prefer `tipped_arrow`: Infinity makes plain `arrow` items free, and a tipped arrow can be coloured with `potion_contents.custom_color`.

```
CustomItem explosive_arrow = {
  base = tipped_arrow
  name = "Explosive Arrow"
  on_projectile_hit_block = boom
  on_projectile_hit_entity = boom
}

Function boom {
  /summon tnt ~ ~ ~ {fuse:0s,explosion_power:2f}
}
```

On a block hit the function runs as the arrow, at the point it hit, so `execute on origin` is whoever shot it, and `kill @s` keeps the arrow from being picked up. On an entity hit it runs as and at the entity hurt, and `execute on attacker` is the shooter.

**How it works.** An arrow keeps the bow or crossbow it was shot from as its `weapon`, and when it hits, the game runs the weapon's enchantment effects at the exact spot: `hit_block` as the arrow, `post_attack` as the entity it hurt. Each item with projectile hooks gets an enchantment, `<namespace>:mms_gen/item_hook/<name>/projectile`, whose effects run the functions. The tick function looks once at each new arrow (it tags them `mms_shot.<namespace>`), and adds the enchantment to the weapon of an arrow that is the item, or that was shot from it. An arrow from a dispenser has no weapon, so it gets a plain bow. The enchantment is never on the item, so it shows in no tooltip and gives no glint. Several items' hooks combine: an Explosive Arrow shot from a bow with hooks runs both.

## How it compiles

Values live on one scoreboard objective per namespace, `mms.<namespace>`:

| Holder | Value |
|---|---|
| `$<name>` | a Global |
| `#<function>.<name>` | a parameter or local |
| `#<function>._t<n>`, `#<function>._f<n>` | temporaries and condition flags |
| `#c<n>` | the constant `n`, for multiplication, division and comparison |
| `#cf`, `#rv` | control flow, see below |

String values and macro arguments live in storage `<namespace>:mms`: `g.<name>` for globals, `l."<function>".<name>` for locals, and `m` for the arguments of the macro being called.

**Branches and loops.** `if` with a one-command body becomes `execute if ... run <command>`. Longer bodies become sub-functions next to the function, such as `demo:check/_if0`.

A loop is a function that checks its condition, runs its body and calls itself. `break` returns from it, and `continue` calls it again.

**Leaving early.** When `break`, `continue` or `return` happens inside a sub-function, it sets `#cf`, and the calling function passes it on. The return value travels in `#rv`.

**Results.** Functions return values with the `return` command, and callers read them with `execute store result score`. NBT reads, `Count` and commands used as values also come through `execute store result score`; a command that gives a bool uses `store success`.

**Recursion.** A call that can lead back to the calling function, directly or through other functions, pushes the caller's scores onto a stack in storage, `<namespace>:mms stack`, and pops them afterwards. Only the scores read after the call are saved. Every other call costs nothing extra.

**Optimisation.** Temporaries that only carry a value from one place to the next are removed, so `x = x + 1` becomes one `scoreboard players add`.

**Load.** The generated load function creates the objective, sets the constants, and gives each global its first value.

## Generated functions

`OnLoad { ... }` and `OnTick { ... }` bodies, and `Objective` declarations, are collected into `<namespace>:mms_gen/load` and `<namespace>:mms_gen/tick`. These are registered in `#minecraft:load` and `#minecraft:tick`. `OnFirstJoin` and `OnRespawn` bodies are collected the same way, into `<namespace>:mms_gen/first_join` and `<namespace>:mms_gen/respawn`.

A hook body that holds only commands is copied into these functions as it is. A body with variables or control flow compiles to its own function, such as `<namespace>:mms_gen/load/<hash>`, which the generated function calls.
