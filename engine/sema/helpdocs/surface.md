# Surface rules: material conditions and material rules

Each `## heading` names a helper by dispatcher and type; see helpdoc.go.

## minecraft:material_condition/stone_depth
True for blocks close to a surface: just under the ground's top (`Floor`) or just above a cave roof's underside (`Ceiling`). Vanilla puts grass on the first block and dirt a few blocks deeper with it.

- surface_type: Which surface depth is counted from: `floor` counts down from the top of the ground, `ceiling` counts up from the underside of an overhang or cave roof.
- offset: Extra depth that still passes. With `0` only the surface block itself passes; with `2` the top three blocks do.
- add_surface_depth: Adds the surface depth to `offset`: a noise of a few blocks that varies across the world, so the layer has an uneven bottom like vanilla dirt.
- secondary_depth_range: Adds a second noise, from 0 up to this many blocks, to the depth. Vanilla uses it for patchy layers such as badlands terracotta.

```mms
SurfaceRule grass = If(StoneDepth(Floor)) { Block(grass_block) }
SurfaceRule dirt = If(StoneDepth(Floor).Add()) { Block(dirt) }
```

## minecraft:material_condition/water
True for blocks at or above the water surface over them, shifted by `offset`, and also where there is no water above at all. Vanilla uses it so grass doesn't grow under lakes and oceans.

- offset: Blocks added to the water height before comparing. `-1` also passes one block below the water.
- surface_depth_multiplier: Multiplies the surface depth noise and adds it to the water height, for an uneven shoreline. From -20 to 20.
- add_stone_depth: Adds the block's depth below the surface to its Y before comparing.

```mms
SurfaceRule dry_grass = If(AboveWater()) { Block(grass_block) }
```

## minecraft:material_condition/above_preliminary_surface
True above the preliminary surface: the rough terrain height worked out from the noise before caves are carved. Vanilla wraps its whole surface layer in it so cave floors deep underground don't get grass.

```mms
SurfaceRule top = If(AboveSurface() && StoneDepth(Floor)) { Block(grass_block) }
```

## minecraft:material_condition/biome
True when the block's biome is one of these. Takes any number of biomes.

- biome_is: The biomes, by id.

```mms
SurfaceRule desert_sand = If(Biome(desert, beach)) { Block(sand) }
```

## minecraft:material_condition/hole
True where the surface depth noise is zero or less: spots where the surface layer is missing. Vanilla badlands leave these as plain terracotta.

```mms
SurfaceRule bands = If(!Hole()) { Block(orange_terracotta) }
```

## minecraft:material_condition/noise_threshold
True where a noise, sampled at the block's x and z, lies between `min_threshold` and `max_threshold`. Use it for patches such as coarse dirt in savannas.

- noise: The noise to sample, by id (a `Noise` declaration or a vanilla one such as `surface`).
- min_threshold: Lowest noise value that passes. Written as `.Min(x)`.
- max_threshold: Highest noise value that passes. Written as `.Max(x)`.

```mms
SurfaceRule patches = If(NoiseThreshold(surface).Min(0.2).Max(1)) { Block(coarse_dirt) }
```

## minecraft:material_condition/not
True when `invert` is false. Usually written with `!`, as in `!Hole()`.

- invert: The condition to invert.

## minecraft:material_condition/steep
True on steep slopes, where the terrain height jumps by four or more blocks between neighbouring columns. Vanilla puts stone on steep mountainsides with it.

```mms
SurfaceRule cliffs = If(Steep()) { Block(stone) }
```

## minecraft:material_condition/temperature
True where the biome is cold enough at this height for snow to fall.

```mms
SurfaceRule snowy = If(Temperature()) { Block(snow_block) }
```

## minecraft:material_condition/vertical_gradient
True at and below `true_at_and_below`, false at and above `false_at_and_above`, and true at random in between, less often the higher the block. Vanilla makes the bedrock floor and the deepslate transition with it.

- random_name: Seeds the randomness. Give each gradient its own name, such as `"bedrock_floor"`, so they don't line up.
- true_at_and_below: The height where the condition is always true, and below.
- false_at_and_above: The height where the condition is always false, and above.

```mms
SurfaceRule bedrock = If(VerticalGradient("bedrock_floor", AboveBottom(0), AboveBottom(5))) { Block(bedrock) }
```

## minecraft:material_condition/y_above
True for blocks at or above a height. The height can shift with the surface depth noise for an uneven edge.

- anchor: The height, as a number (absolute Y) or an anchor such as `AboveBottom(5)`.
- surface_depth_multiplier: Multiplies the surface depth noise and adds it to the height. From -20 to 20. Written as `.Mul(x)`.
- add_stone_depth: Adds the block's depth below the surface to its Y before comparing. Written as `.Add()`.

```mms
SurfaceRule peaks = If(YAbove(200)) { Block(snow_block) }
```

## minecraft:material_rule/block
Places one block. Writing just a block id where a rule is expected does the same thing.

- result_state: The block, with optional state such as `oak_log[axis = y]`.

```mms
SurfaceRule floor = Block(stone)
```

## minecraft:material_rule/condition
Runs `then_run` only where `if_true` passes. Usually written `If(condition) { rules }`.

- if_true: The condition to test.
- then_run: The rule to run when it passes.

```mms
SurfaceRule grass = If(StoneDepth(Floor)) { Block(grass_block) }
```

## minecraft:material_rule/sequence
Tries each rule in order and stops at the first one that places a block. A list `[a, b]`, or several rules in an `If { }` body, makes the same thing.

- sequence: The rules, in order.

```mms
SurfaceRule top = [
  If(StoneDepth(Floor)) { Block(grass_block) },
  Block(dirt),
]
```

## minecraft:material_rule/bandlands
Places the striped terracotta of badlands, with the colour picked by height.

```mms
SurfaceRule stripes = If(YAbove(63)) { Bandlands() }
```
