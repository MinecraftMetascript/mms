# Density functions

A density function gives a number for every block position. In noise
settings, blocks where `final_density` is above 0 are solid. MMS writes
`a + b`, `a - b`, `a * b` and `a / b` for the arithmetic helpers, and a
plain number for a constant.

## minecraft:density_function/noise
Samples a noise at the block's position, scaled per axis. The result is roughly between -1 and 1. Most terrain starts here.

- noise: The noise, by id: a `Noise` declaration or a vanilla one such as `minecraft:ridge`.
- xz_scale: Multiplies x and z before sampling. Below 1 stretches the noise out horizontally; above 1 squeezes it.
- y_scale: Multiplies y before sampling. `0` makes the noise the same at every height (2D).
- shift_x: Added to x before sampling, so the noise doesn't line up with others. Vanilla uses `ShiftA`.
- shift_y: Added to y before sampling.
- shift_z: Added to z before sampling. Vanilla uses `ShiftB`.

```mms
Noise hills = { base_octave = -7, amplitude_modifiers = [1, 1] }
DensityFunction land = Noise(hills, 1, 0) * 0.5
```

## minecraft:density_function/constant
The same number everywhere. A plain number does the same thing.

- value: The number.

```mms
DensityFunction half = Constant(0.5)
```

## minecraft:density_function/add
`left + right`. Usually written with `+`.

```mms
DensityFunction raised = Noise(minecraft:ridge, 1, 0) + 0.2
```

## minecraft:density_function/sub
`left - right`. Usually written with `-`.

## minecraft:density_function/mul
`left * right`. Usually written with `*`. Vanilla stops computing `right` where `left` is 0.

## minecraft:density_function/div
`left / right`. Usually written with `/`.

## minecraft:density_function/min
The smaller of the two inputs. `Min(a, b)` keeps the lower of two terrains.

## minecraft:density_function/max
The larger of the two inputs. `Max(a, b)` joins two terrains.

```mms
DensityFunction both = Max(Noise(minecraft:ridge, 1, 0), Noise(minecraft:jagged, 1, 0))
```

## minecraft:density_function/abs
The absolute value of `input`. Turns a noise into sharp ridges along its zero line.

## minecraft:density_function/negate
`-input`. Usually written with a leading `-`.

## minecraft:density_function/square
`input * input`. Always 0 or more, and flattens values near 0.

## minecraft:density_function/cube
`input * input * input`. Keeps the sign and flattens values near 0.

## minecraft:density_function/sqrt
The square root of `input`.

## minecraft:density_function/log
The natural logarithm of `input`.

## minecraft:density_function/pow
`base` raised to the power `exponent`.

## minecraft:density_function/reciprocal
`1 / input`.

## minecraft:density_function/sign
`-1`, `0` or `1`, following the sign of `input`.

## minecraft:density_function/half_negative
`input`, with negative values halved. Vanilla uses it to soften cave walls.

## minecraft:density_function/quarter_negative
`input`, with negative values divided by four.

## minecraft:density_function/squeeze
Clamps `input` to -1..1, then bends it to `x / 2 - x³ / 24`. Vanilla applies it near the end of `final_density` to smooth sharp values.

## minecraft:density_function/clamp
Keeps `input` between `min` and `max`.

- min: The lowest value let through.
- max: The highest value let through.

```mms
DensityFunction capped = Clamp(Noise(minecraft:ridge, 1, 1), -0.5, 0.5)
```

## minecraft:density_function/ceil
Rounds `input` up to a whole multiple of `multiple`.

## minecraft:density_function/floor
Rounds `input` down to a whole multiple of `multiple`. Rounding a gradient makes terraces.

## minecraft:density_function/round
Rounds `input` to the nearest whole multiple of `multiple`.

## minecraft:density_function/truncate
Rounds `input` toward zero, to a whole multiple of `multiple`.

## minecraft:density_function/gradient
Changes linearly along one axis: `from_value` at `from_coordinate`, `to_value` at `to_coordinate`. Beyond the two coordinates it holds the end values, unless `tiling` says otherwise. Use it to make terrain solid at the bottom and air at the top.

- axis: The axis the gradient runs along: `x`, `y` or `z`.
- from_coordinate: Where the gradient starts, in blocks.
- to_coordinate: Where the gradient ends, in blocks.
- from_value: The value at `from_coordinate`.
- to_value: The value at `to_coordinate`.
- tiling: What happens beyond the two coordinates: `clamp_to_edge` holds the end values, `repeat` starts over, `mirrored_repeat` runs back and forth.

```mms
DensityFunction ground = Gradient(
  axis = y,
  from_coordinate = 0,
  to_coordinate = 128,
  from_value = 1,
  to_value = -1,
)
```

## minecraft:density_function/interpolated
Computes `input` only at the corners of cells and blends between them. Smoother and much faster than sampling every block; vanilla wraps its terrain in it.

- input: The function to interpolate.
- cell_size_xz: Cell width, in blocks.
- cell_size_y: Cell height, in blocks.

## minecraft:density_function/cache
Computes `input` once per position, for a function used in several places. The value is unchanged.

## minecraft:density_function/range_choice
Picks between two functions by the value of `input`: `when_in_range` where `min_inclusive <= input < max_exclusive`, `when_out_of_range` elsewhere. Only the chosen one is computed.

- input: The value to test.
- min_inclusive: The lowest value that counts as in range.
- max_exclusive: The value above the range; it counts as out of range.
- when_in_range: Used where `input` is in range.
- when_out_of_range: Used everywhere else.

## minecraft:density_function/interval_select
Picks one of `functions` by which interval of `thresholds` the value of `input` falls in: below the first threshold picks the first function, and so on.

## minecraft:density_function/lerp
Blends from `first` to `second` by `alpha`: `first` where `alpha` is 0, `second` where it is 1.

- alpha: How far to blend, usually from 0 to 1.
- first: The value where `alpha` is 0.
- second: The value where `alpha` is 1.

## minecraft:density_function/spline
Maps a value through a cubic spline: a curve through points, each with a `location`, a `value` and a `derivative`. Vanilla shapes continents, erosion and peaks into terrain height with splines.

## minecraft:density_function/shift
Samples `noise` at a quarter of the block's x, y and z and multiplies it by 4. Used as a shift for another noise.

## minecraft:density_function/shift_a
Samples `noise` at a quarter of the block's x and z (y is 0) and multiplies it by 4. Vanilla uses it as `shift_x` of its climate noises.

## minecraft:density_function/shift_b
Like `ShiftA`, with x and z swapped. Vanilla uses it as `shift_z` of its climate noises.

## minecraft:density_function/old_blended_noise
The 3D terrain noise from before 1.18, which vanilla still uses as the base of overworld and nether terrain.

## minecraft:density_function/end_outer_islands
The End's island terrain: the main island in the middle, and scattered islands beyond 1024 blocks.

## minecraft:density_function/beardifier
Adds density around structures that adapt the terrain, such as villages and ancient cities. Vanilla adds it to `final_density`.

## minecraft:density_function/blend_alpha
Blending with chunks from before 1.18: 1 in new chunks, lower near old ones.

## minecraft:density_function/blend_offset
Blending with chunks from before 1.18: 0 in new chunks.

## minecraft:density_function/blend_density
Blends `input` with the terrain of chunks from before 1.18. In new chunks it is `input` unchanged.
