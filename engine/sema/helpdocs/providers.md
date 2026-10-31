# Int, float and height providers

A provider is a number picked at random each time it is used. MMS writes
a range such as `2..5` for `Uniform`, and a plain number for `Constant`.

## minecraft:int_provider/constant
Always the same whole number. A plain number does the same thing.

## minecraft:int_provider/uniform
A random whole number from `min_inclusive` to `max_inclusive`, each equally likely. Usually written as a range, `2..5`.

- min_inclusive: The smallest value.
- max_inclusive: The largest value.

```mms
PlacedFeature some_iron = { feature = minecraft:ore_iron, placement = [Count(Uniform(2, 5)), InSquare()] }
```

## minecraft:int_provider/biased_to_bottom
A random whole number from `min_inclusive` to `max_inclusive`, with low values more likely.

## minecraft:int_provider/very_biased_to_bottom
Like `BiasedToBottom`, with low values even more likely.

## minecraft:int_provider/clamped
Another provider's value, kept from `min_inclusive` to `max_inclusive`.

- source: The provider to clamp.

## minecraft:int_provider/clamped_normal
A whole number from a bell curve around `mean`, kept from `min_inclusive` to `max_inclusive`.

- mean: The most likely value.
- deviation: How widely values spread around `mean`.

## minecraft:int_provider/trapezoid
A whole number from `min` to `max`, most likely in a middle band `plateau` wide and less likely toward the ends.

- plateau: Width of the most likely band in the middle. `0` gives a triangle.

## minecraft:int_provider/weighted_list
Picks one of several providers by weight.

## minecraft:float_provider/constant
Always the same number. A plain number does the same thing.

## minecraft:float_provider/uniform
A random number from `min_inclusive` up to (not including) `max_exclusive`. Usually written as a range, `0.5..1.5`.

## minecraft:float_provider/clamped_normal
A number from a bell curve around `mean`, kept from `min` to `max`.

- mean: The most likely value.
- deviation: How widely values spread around `mean`.

## minecraft:float_provider/trapezoid
A number from `min` to `max`, most likely in a middle band `plateau` wide.

## minecraft:height_provider/constant
Always the same height.

## minecraft:height_provider/uniform
A random height from `min_inclusive` to `max_inclusive`, each equally likely.

```mms
PlacedFeature low_iron = {
  feature = minecraft:ore_iron
  placement = [Count(10), InSquare(), HeightRange(Uniform(AboveBottom(0), Absolute(64)))]
}
```

## minecraft:height_provider/trapezoid
A random height from `min_inclusive` to `max_inclusive`, most likely in the middle. Vanilla ores that peak at one depth use it.

- plateau: Width of the most likely band in the middle. Defaults to 0, a triangle.

## minecraft:height_provider/biased_to_bottom
A random height from `min_inclusive` to `max_inclusive`, with low heights more likely.


## minecraft:height_provider/very_biased_to_bottom
Like `BiasedToBottom`, with low heights even more likely.


## minecraft:height_provider/weighted_list
Picks one of several height providers by weight.
