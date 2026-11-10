# Shaping terrain

A world's shape comes from one number per block, the **density**. Where it is above 0 the generator places stone, and where it is below 0 it leaves air, or water below sea level. A `DensityFunction` is the formula for that number, and shaping terrain means writing formulas.

In JSON a formula is a deep tree of `add` and `mul` objects. In MMS it is arithmetic, and the editor draws the result as you type. This guide is a set of recipes to combine. For the resources around the terrain (dimension, biome, surface), follow the [custom dimension tutorial](../start/tutorial-dimension.md) first.

## See what you are doing

Put the cursor on a `Noise`, a `DensityFunction` or a `NoiseSettings` and open the **Preview** tab of the web editor (in VS Code, click **Preview** above it). It draws the function as a slice of the world: the raw density, the blocks it makes, or a 3D view. The picture redraws on every edit, so changing a number and watching is the normal way to work.

Every sample on this page is a file of its own. Paste one into a new project to see it.

## Ground with a height

Almost every terrain starts from a gradient: solid at the bottom, air at the top.

```mms
/** 1 at the bottom of the world, 0 at y 64, -1 at y 192. */
DensityFunction slope = Gradient(
  axis = y,
  from_coordinate = -64,
  to_coordinate = 192,
  from_value = 1,
  to_value = -1,
)
```

On its own this is a flat world with its surface at y = 64, where the value crosses 0. Everything else is added to it.

## Hills

Noise is a smooth random value between about -1 and 1. Adding it to the slope moves the surface up where the noise is high and down where it is low:

```mms
/** Rolling hills, a few hundred blocks across. */
Noise rolling {
  base_octave = -7
  amplitude_modifiers = [1, 0.5, 0.25]
}

DensityFunction slope = Gradient(
  axis = y,
  from_coordinate = -64,
  to_coordinate = 192,
  from_value = 1,
  to_value = -1,
)

/** The surface rises and falls about 32 blocks around y 64. */
DensityFunction hills = slope + Noise(rolling, 1, 0) * 0.25
```

Three numbers do the work:

- **`base_octave`** sets the size of the shapes. Each step down doubles it: `-7` gives features about 128 blocks across, `-8` about 256.
- **`amplitude_modifiers`** adds detail. Each entry is another layer of noise at half the size, with that strength. `[1]` is smooth, and `[1, 0.5, 0.25]` has bumps on its bumps.
- **The multiplier** sets the height. The slope changes by 1 over 128 blocks, so `* 0.25` moves the surface up to 32 blocks either way.

The arguments after the noise are its scale across (`xz_scale`) and up (`y_scale`). A `y_scale` of `0` makes the noise the same at every height, so the ground has no overhangs. With `1`, the noise changes with height too, and hills grow overhangs and floating bits.

## Caves

Caves are a second formula that is below 0 where there should be air. `Min` takes the lower of two values, so a block is solid only when both say so:

```mms
Noise cave_noise {
  base_octave = -5
  amplitude_modifiers = [1, 0.5]
}

DensityFunction slope = Gradient(
  axis = y,
  from_coordinate = -64,
  to_coordinate = 192,
  from_value = 1,
  to_value = -1,
)

/** Below 0 inside a cave. The y scale of 2 makes caves wider than tall. */
DensityFunction caves = 0.45 - Noise(cave_noise, 1, 2)

DensityFunction carved = Min(slope, caves)
```

Lower `0.45` for more and bigger caves. `Max` does the opposite and joins two terrains, such as islands floating above a ground.

A few more helpers cover most shapes:

| Helper | Use it to |
| --- | --- |
| `Abs(x)` | turn a noise into sharp ridges, or into thin tunnels along the lines where it is 0 |
| `Square(x)` | flatten values near 0 |
| `Clamp(x, min, max)` | keep a value in a range, so that it cannot overpower the rest |
| `Lerp(alpha, a, b)` | blend from one terrain to another, by a third function |
| `IntervalSelect(x, [...], [...])` | switch between terrains by a value, such as one terrain per region |

Hover a helper for what it does. The [`cave-siege`](../../examples/cave-siege/) example builds a whole cave world from these, with comments on each step.

## Plains and cliffs with a spline

Adding noise gives hills that are equally bumpy everywhere. Real landscapes have flat places and steep ones. A `Spline` maps a value through a curve you draw, so the same noise can mean "flat" over most of its range and "cliff" in a narrow band:

```mms
Noise rolling {
  base_octave = -7
  amplitude_modifiers = [1, 0.5, 0.25]
}

DensityFunction slope = Gradient(
  axis = y,
  from_coordinate = -64,
  to_coordinate = 192,
  from_value = 1,
  to_value = -1,
)

/** How far the ground is lifted: plains, a cliff, then a plateau. */
DensityFunction lift = Spline({
  coordinate = Noise(rolling, 1, 0)
  points = [
    SplinePoint(-1, 0, -0.1),
    SplinePoint(0.1, 0, 0),
    SplinePoint(0.3, 0, 0.35),
    SplinePoint(1, 0, 0.4),
  ]
})

DensityFunction plateaus = slope + lift
```

Each `SplinePoint` is a location on the coordinate, the slope of the curve there, and the value. Between noise values of 0.1 and 0.3 the lift jumps from 0 to 0.35, which is a cliff about 45 blocks high. Everywhere else it barely changes.

Typing the numbers is the hard way. With the cursor in the spline, the Preview tab shows the curve: drag a point to move it, drag its handle to change the slope, double-click to add a point and press Delete to remove one. Each drag rewrites the numbers in the source. In VS Code, the **Edit spline** lens above `Spline(` opens the same editor. [Splines](resources.md#splines) has the details, including splines inside splines, which is how vanilla mixes continents, erosion and ridges.

## Keep the world closed

A world with caves needs a floor, or players fall out of it. Blend towards solid near the bottom:

```mms
DensityFunction terrain = Gradient(axis = y, from_coordinate = -64, to_coordinate = 192, from_value = 1, to_value = -1)

/** 0 at the bottom of the world, 1 from y -48 up. */
DensityFunction floor_fade = Gradient(
  axis = y,
  from_coordinate = -64,
  to_coordinate = -48,
  from_value = 0,
  to_value = 1,
)

/** Solid at the bottom, the terrain from y -48 up. */
DensityFunction closed = Lerp(floor_fade, 1, terrain)
```

The fade starts at the bottom of the world, which is -64 here. Use your own dimension's `min_y`: in a world that starts at y = 0, this fade would be 1 everywhere and do nothing. The same trick at the top, blending towards `-1`, keeps terrain from being cut off flat at the height limit.

## Use it in a world

The terrain is the `final_density` of a `NoiseSettings`, which a `Dimension` names as its generator with `Noise(settings = hills, ...)`. The sample below uses the flat slope, to stay short. Put the name of your own formula in its place, and keep `noise` in step with the dimension type's `min_y` and `height`:

```mms
DensityFunction terrain = Gradient(axis = y, from_coordinate = -64, to_coordinate = 192, from_value = 1, to_value = -1)

NoiseSettings hills {
  sea_level = 48
  default_block = stone
  default_fluid = water
  disable_mob_generation = false
  legacy_random_source = false
  noise = { min_y = -64, height = 384 }
  material_rule = Block(stone)
  spawn_target = []
  noise_router = {
    final_density = Interpolated(terrain, 4, 8)
    temperature = 0
    vegetation = 0
    continents = 0
    erosion = 0
    depth = 0
    ridges = 0
    chunk_surface_level = 0
  }
}
```

`Interpolated(terrain, 4, 8)` works the formula out once per 4×8×4 cell and blends between, as vanilla does. It is much faster and smooths out single-block noise. The preview of the `NoiseSettings` shows the finished blocks, with water and the surface rule applied.

Give each step a name, as the samples here do. A named `DensityFunction` can be previewed on its own, which is how you find the step that went wrong, and its doc comment says what its numbers mean.

## Try it in the game

```sh
mms build --world ~/.minecraft/saves/MyWorld --reset-dimensions
```

The game never generates a chunk twice, so a change to the terrain only shows in new chunks. `--reset-dimensions` deletes the saved chunks of the project's dimensions whose worldgen changed, so the same place generates again with the new shape. Close the world first. [Seeing worldgen changes](../start/tutorial-dimension.md#seeing-worldgen-changes) explains what is reset and how to get the old chunks back.

## Borrow from vanilla

Vanilla's own density functions and noises can be used by id, with no declaration: `overworld/continents`, `overworld/erosion`, or the noise `ridge`. Go to definition on one opens the game's JSON for it. To study a whole vanilla dimension as MMS, run `mms import` on the vanilla data pack.

The [`sky-islands`](../../examples/sky-islands/), [`badlands`](../../examples/badlands/) and [`cave-siege`](../../examples/cave-siege/) examples are complete terrains to take apart.
