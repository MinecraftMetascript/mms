# Making your own kinds

`CustomItem`, `Book` and `JigsawStructure` are not built into the compiler. Each is a `Blueprint`: a kind of declaration written in MMS, which turns a few fields into the declarations the game needs. You can write your own, and yours get the same completion, checking and hover as the built-in ones.

This guide builds one step by step. [Blueprints](../reference/language/blueprints.md) in the language reference lists every part.

## When a Blueprint helps

A [`Template`](repetition.md#name-a-shape) stands for one value. A Blueprint is for when a thing in your pack is **several declarations that belong together**. A wand, for example, is a custom item, a cooldown score, a line in the tick function and a recipe. Write the second wand by copying the first, and the two will drift apart.

With a Blueprint, a wand is this:

```
Wand fire_wand {
  name = "Fire Wand"
  cast = shoot_fireball
  cooldown = 60
  reagent = blaze_powder
}
```

## 1. Fields and emit

A Blueprint starts with its fields, and an `emit` block that says what an instance writes:

```mms
/** A wand: a stick that casts a spell on right click. */
Blueprint Wand {
  /** The wand's name. */
  name: string
  /** The function that casts the spell. */
  cast: Function

  emit {
    primary CustomItem ${self.path} {
      base = stick
      name = name
      on_right_click = cast
    }
  }
}

Function shoot_fireball {
  /summon fireball ^ ^1.5 ^1
}

Wand fire_wand {
  name = "Fire Wand"
  cast = shoot_fireball
}

Function give {
  /give @s ${fire_wand}
}
```

- **Fields** are `name: Type`. `cast: Function` takes the name of a function, and the editor completes the project's functions there. The doc comments show on hover and in completion, in every file that declares a `Wand`.
- **`emit`** holds ordinary declarations. Field names stand for the instance's values, and `${self.path}` is the instance's own name, so `Wand fire_wand` writes `CustomItem fire_wand`.
- **`primary`** marks the declaration the instance stands for. That is why `${fire_wand}` works in the `give` command: the wand's name means what the custom item's name means.

## 2. Defaults and optional fields

A field with `= value` has a default, and one with `?` after its type can be left out. Anything else is required, and leaving it out is an error with a fix that adds it.

```
  /** Ticks between casts. */
  cooldown: int = 20
  /** What the wand is crafted from, with a stick. Left out, it has no recipe. */
  reagent: Id<item>?
```

`emit` can use `if`, `for` and `let`, so an optional field can decide whether a declaration is written at all:

```
    if (reagent != none) {
      Recipe ${self.path} = Shapeless([stick, reagent], ${self.path})
    }
```

## 3. Code in emit

An emitted function is written like any other, with the fields in reach. `run { }` holds commands where a function's name is expected, which keeps the cooldown logic inside the Blueprint:

```
    /** Ticks until a player can use ${self.path} again. */
    Objective ${self.path}_cooldown

    primary CustomItem ${self.path} {
      base = stick
      name = name
      on_right_click = run {
        if (${self.path}_cooldown[@s] > 0) {
          actionbar(@s, "{red}${name} is cooling down")
          return
        }
        ${self.path}_cooldown[@s] = cooldown
        cast()
      }
    }

    OnTick {
      /scoreboard players remove @a[scores={${self.path}_cooldown=1..}] ${self.path}_cooldown 1
    }
```

Each wand gets its own objective, `fire_wand_cooldown`, so two wands cool down separately. `cooldown` is written into the function as a number, and `cast()` calls whichever function the instance names.

`OnTick` is a kind that merges: every instance's block is added to the pack's one tick function, along with any `OnTick` the project writes itself. Tags, `Lang` and `OnLoad` merge in the same way.

## 4. Rules of your own

`check` runs on every instance at compile time, and reports problems on the field that caused them:

```
  check {
    if (cooldown < 1) {
      error(cooldown, "a cooldown is at least one tick")
    }
    if (cooldown > 200) {
      warn(cooldown, "slow-wand", "a cooldown of over ten seconds feels broken")
    }
  }
```

A `Wand` with `cooldown = 0` now gets an error on that line, in the editor and in `mms check`, with your message. `warn` and `hint` report softer problems. The name between the two, `"slow-wand"`, makes the rule one that `// mms-ignore slow-wand` silences and the `lint` settings in `mms.json` can turn off or promote, like the built-in [lint rules](lint.md).

Mistakes inside the instance's own values are caught without any `check`: a `cast` that names no function, a `cooldown` given as text, or an unknown `reagent` are errors from the field's type.

## The whole Blueprint

```mms
/** A wand: a stick that casts a spell, with a cooldown. */
Blueprint Wand {
  /** The wand's name. */
  name: string
  /** The function that casts the spell. */
  cast: Function
  /** Ticks between casts. */
  cooldown: int = 20
  /** What the wand is crafted from, with a stick. Left out, it has no recipe. */
  reagent: Id<item>?

  check {
    if (cooldown < 1) {
      error(cooldown, "a cooldown is at least one tick")
    }
    if (cooldown > 200) {
      warn(cooldown, "slow-wand", "a cooldown of over ten seconds feels broken")
    }
  }

  emit {
    /** Ticks until a player can use ${self.path} again. */
    Objective ${self.path}_cooldown

    primary CustomItem ${self.path} {
      base = stick
      name = name
      on_right_click = run {
        if (${self.path}_cooldown[@s] > 0) {
          actionbar(@s, "{red}${name} is cooling down")
          return
        }
        ${self.path}_cooldown[@s] = cooldown
        cast()
      }
    }

    OnTick {
      /scoreboard players remove @a[scores={${self.path}_cooldown=1..}] ${self.path}_cooldown 1
    }

    if (reagent != none) {
      Recipe ${self.path} = Shapeless([stick, reagent], ${self.path})
    }
  }
}

Function shoot_fireball {
  /summon fireball ^ ^1.5 ^1
}

Function blink {
  /tp @s ^ ^ ^6
}

Wand fire_wand {
  name = "Fire Wand"
  cast = shoot_fireball
  cooldown = 60
  reagent = blaze_powder
}

Wand blink_wand {
  name = "Blink Wand"
  cast = blink
}

LootTable chests/wizard {
  pools = [Pool(1, fire_wand, blink_wand)]
}
```

Two wands are now eleven lines. Between them they write two custom items with their right-click hooks, two objectives, two lines of the tick function, a recipe and the translations.

## Seeing what an instance writes

- **Expanded.** With the cursor in `Wand fire_wand`, the web editor's Expanded tab shows the declarations it emits, with the field values filled in. In VS Code, click the **Expand** lens above it.
- **Errors point both ways.** A problem in emitted code is reported on each instance that runs into it, and the editor also marks the line of the Blueprint that wrote it.
- **Go to definition** on `Wand` opens the Blueprint, and on `fire_wand` the instance.

## What a name means elsewhere

`primary` forwards the instance's name to one emitted declaration. For finer control, `as Type = value` says what the name means where that type is expected. The prelude's `CustomItem` uses it to be an item stack in a recipe result, a loot entry in a pool and its base item in a tag. See [as](../reference/language/blueprints.md#as).

## Sharing and hiding

A Blueprint can be used from any file of the project. Put `private` before it to keep it to its own file. The prelude's Blueprints are a good place to read real ones: each page under [Declaration kinds](https://minecraftmetascript.github.io/mms/docs/reference/kinds) lists the fields, and the [`shop`](../../examples/shop/) example has a Blueprint that emits a function, with tests.
