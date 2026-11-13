# Blueprints

A `Blueprint` declares a new kind of declaration. Its fields are checked and completed like the game's own, its `check` block reports problems, and its `emit` block writes the declarations it stands for. `CustomItem`, `Book`, `JigsawStructure` and `CustomBiome` are Blueprints written this way, in the prelude: their pages under [Declaration kinds](https://minecraftmetascript.github.io/mms/docs/reference/kinds) list their fields.

```mms
Objective wand_cooldown

/** A wand: a stick that casts a spell, with a cooldown. */
Blueprint Wand {
  /** The wand's name. */
  name: string
  /** The function that casts the spell. */
  cast: Function
  /** Ticks between casts. */
  cooldown: int = 20

  check {
    if (cooldown < 1) {
      error(cooldown, "a cooldown is at least one tick")
    }
    if (cooldown > 200) {
      warn(cooldown, "slow-wand", "a cooldown of over ten seconds feels broken")
    }
  }

  emit {
    primary CustomItem ${self.path} {
      base = stick
      name = name
      on_right_click = run {
        if (wand_cooldown[@s] > 0) {
          return
        }
        wand_cooldown[@s] = cooldown
        cast()
      }
    }
  }
}

OnTick {
  /scoreboard players remove @a[scores={wand_cooldown=1..}] wand_cooldown 1
}

Function shoot_fireball {
  /summon fireball ^ ^1.5 ^1
}

Wand fire_wand {
  name = "Fire Wand"
  cast = shoot_fireball
}

Recipe fire_wand = Shapeless([stick, blaze_powder], fire_wand)
```

A Blueprint has four parts, each optional.

## Fields

A field is `name: Type`, with `= default` or a `?` on the type when it can be left out. A field without either is required, and leaving it out is an error with a fix that adds it. Doc comments on fields show in completion and hover. See [Values and types](values.md#types) for the types.

`let name = value` between the fields names a value worked out from them, for `check`, `as` and `emit` to share.

## check

`check { }` runs at compile time on each declaration of the kind. `error(field, "message")`, `warn(...)` and `hint(...)` report a problem on that field's value. A rule name between the two, as in `warn(cooldown, "slow-wand", "...")`, lets `// mms-ignore slow-wand` and the `lint` settings in `mms.json` turn it off or change its level. A rule without a name is reported under the kind's name, such as `mms/wand`.

## as

`as Type = value` says what the declaration's name means where a `Type` is expected. This is how a custom item's name becomes an item stack in a recipe result and a loot entry, and its base item in a tag:

```mms
/** A trophy: a gold ingot with a name. */
Blueprint Trophy {
  /** What it was won for. */
  title: string

  let trophy_components = { item_name = title, rarity = epic }

  /** In a recipe result or a give command: the trophy itself. */
  as ItemStack = { id = gold_ingot, components = trophy_components }
  as ItemStackTemplate = { id = gold_ingot, components = trophy_components }
  as LootPoolEntry = Item(name = gold_ingot, modifier = [{ type = set_components, components = trophy_components }])
  /** Where only an item id fits, it is any gold ingot. */
  as Id<item> = gold_ingot hint("trophy-as-item", "'${self.path}' stands for gold_ingot here")
}

Trophy dragon_slayer {
  title = "Dragon Slayer"
}

LootTable trophies {
  pools = [Pool(1, dragon_slayer)]
}
```

The first `as` that fits the expected type is used. `hint(...)` after the value reports a lint hint wherever that `as` is used, for a meaning a reader might not expect.

## emit

`emit { }` holds the declarations the kind writes. It can use `if`, `for` and `let`, and the fields by name. `self.path` is the declaration's name and `self.id` its id with the namespace, and names can be templates, as in a [top-level for](reuse.md#repeating-declarations). Emitting into a kind that merges, such as a tag, `Lang`, `OnLoad` or `OnTick`, merges with the project's own.

One emitted declaration can be marked `primary`. It must have the declaration's own id, and the name then forwards to it: wherever none of the Blueprint's `as` fits, `fire_wand` means what the custom item `fire_wand` means. That is why `fire_wand` above works as a recipe result. A Blueprint has at most one primary.

### Consts of an instance

A `const` at the top of an `emit` block belongs to each instance, and other code reads it as `name.const`. Use one for a value that has no file of its own. The prelude's `CustomBiome` has one, the biome's entries for a multi-noise biome source:

```
Blueprint CustomBiome {
  climate: ClimateParameters | [ClimateParameters]
  biome: Biome
  surface_rule: SurfaceRule?
  tags: [Tag<worldgen/biome>] = []

  let climates = if (kind_of(climate) == "list") climate else [climate]

  emit {
    primary Biome ${self.path} = biome

    const noise_entries: [BiomeNoiseEntry] = [
      for c in climates:
        { biome = self.id, parameters = c }
    ]

    for tag in tags {
      BiomeTag ${tag} = [self.id]
    }

    if (surface_rule != none) {
      SurfaceRule ${self.path} = If(Biome(self.id)) { surface_rule }
    }
  }
}
```

Each `CustomBiome` then has its own `noise_entries`:

```mms
CustomBiome ash_fields {
  climate = {
    temperature = [0.5, 1]
    humidity = [-1, 0]
    continentalness = 0
    erosion = 0
    weirdness = 0
    depth = 0
    offset = 0
  }
  biome = {
    has_precipitation = false
    temperature = 2
    downfall = 0
    effects = { water_color = "#3F76E4" }
    carvers = []
    features = []
  }
}

CustomBiome frost_fields {
  climate = {
    temperature = [-1, -0.5]
    humidity = 0
    continentalness = 0
    erosion = 0
    weirdness = 0
    depth = 0
    offset = 0
  }
  biome = {
    has_precipitation = true
    temperature = -0.5
    downfall = 0.5
    effects = { water_color = "#3D57D6" }
    carvers = []
    features = []
  }
}

const fields_source: BiomeSource = MultiNoise(
  biomes = [...ash_fields.noise_entries, ...frost_fields.noise_entries]
)
```

`ash_fields.noise_entries` are the entries that `CustomBiome ash_fields` worked out. They can be read anywhere a const can, in any file and before the instance is written, and in a command as `${ash_fields.noise_entries}`. Only the consts at the top of an `emit` block are read this way. A `let`, a field, or a const inside an `if` or a `for` stays inside the Blueprint. An instance that reads its own consts while it is being declared is an error.
