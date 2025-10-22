# MMS Language Specification

This document describes the language specification for Minecraft Metascript (MMS). It provides a comprehensive reference for all block types, functions, and their usage patterns.

## Table of Contents

- [Language Basics](#language-basics)
- [Surface Rules](#surface-rules)
- [Density Functions](#density-functions)
- [Dimension Configuration](#dimension-configuration)
- [Noise Functions](#noise-functions)

## Language Basics

This section covers the fundamental syntax and structure of MMS (Minecraft Metascript) files.

### File Structure

MMS files are organized into namespaces that contain block definitions, variables, and functions:

```mms
Namespace namespace_name {
    // Block definitions go here
    BlockType {
        variable_name = value
        function_name = Function(args)
    }
}
```

### Basic Syntax Elements

#### Namespaces
Namespaces organize related definitions and allow cross-referencing between files:

```mms
Namespace test_files {
    Noise {
        MyNoise = Noise(-1).Amplitudes(1, 2, 0.5, -1)
    }
}
```

#### Variable Assignment
Variables are defined using the `=` operator:

```mms
variable_name = value
```

#### References
Reference other namespaces using the `namespace:name` syntax:

```mms
mms_demo:InFriendlyBiome  // References InFriendlyBiome from mms_demo namespace
```

#### Function Calls
Functions use parentheses for arguments with method chaining support:

```mms
FunctionName(arg1, arg2)
  .Method1(value)
  .Method2(value)
```

#### Lists
Lists are defined using square brackets:

```mms
[ item1, item2, item3 ]
```

#### Conditionals
Conditional expressions use the `If` keyword:

```mms
If (condition) value_if_true
```

#### Comments
- Single line comments: `// comment text`
- Multi-line comments: `/* comment text */`
- Documentation comments: `/** documentation */`
  - Documentation comments appear when use "hover" functionality with a supported editor

### Grammar Overview

The MMS language follows this basic grammar structure:

```
file: (namedBlock)*
namedBlock: Identifier Identifier '{' (varDecl | block)* '}'
block: Identifier '{' (varDecl)* '}'
varDecl: Identifier '=' value
value: number | String | function | reference | conditional | list | resourceTag
```

### Common Patterns

#### Biome-based Surface Rules
```mms
Namespace mms_demo {
  Surface {
    InFriendlyBiome = Biome(forest, plains, beach)
    InUnfriendlyBiome = Biome(desert, badlands, deep_ocean)

    HoneySurface = Block(honey)
    SlimeSurface = Block(slime)

    MyStrangeSurface = [
      If (mms_demo:InFriendlyBiome) HoneySurface
      If (mms_demo:InUnfriendlyBiome) SlimeSurface
      Block(magma_block)
    ]
  }
}
```

#### Noise Function Definition
```mms
Namespace test_files {
    Noise {
        MyNoise = Noise(-1).Amplitudes(1, 2, 0.5, -1)
    }
}
```

### Best Practices

1. **Use descriptive names** for namespaces, variables, and functions
2. **Organize related definitions** within the same namespace
3. **Use references** to avoid duplicating complex definitions
4. **Leverage conditionals** for dynamic behavior based on biome, height, or other factors
5. **Document complex logic** with comments for maintainability

## Surface Rules

Surface rules define how blocks are placed on the terrain surface based on various conditions.

### Surface

#### Block Selection Functions

**Bandlands()**
- **Description**: Used in badlands to place terracotta
- **Usage**:
  ```mms
  Bandlands()
  ```

**Block(ref)**
- **Description**: Selects a block to place
- **Parameters**:
  - `ref`: Block reference
- **Usage**:
  ```mms
  Block(ref)
  ```

#### Surface Rule Lists

**SurfaceRule List**
- **Description**: A list of surface rules that can be combined
- **Usage**:
  ```mms
  [{ ref(SurfaceRule) } | { SurfaceRule }]
  ```

#### Conditional Functions

**AboveSurface()**
- **Description**: Checks if the current position is above the preliminary surface level, ignoring noise caves
- **Usage**:
  ```mms
  AboveSurface()
  ```

**Biome(ref, ref...)**
- **Description**: Passes for columns where the biome matches the specified biomes
- **Parameters**:
  - `ref`: Biome reference(s)
- **Usage**:
  ```mms
  Biome(ref, ref...)
  ```

**Hole()**
- **Description**: Passes for columns where the surface depth is 0
- **Usage**:
  ```mms
  Hole()
  ```

**NoiseThreshold(ref(Noise) | Noise())**
- **Description**: Passes for columns where the input noise is between the minimum and maximum values
- **Parameters**:
  - `ref(Noise) | Noise()`: Noise function reference or inline definition
- **Configuration**:
  - `.Min(float)`: Minimum threshold value
  - `.Max(float)`: Maximum threshold value
- **Usage**:
  ```mms
  NoiseThreshold(ref(Noise) | Noise())
    .Min(float)
    .Max(float)
  ```

**Steep()**
- **Description**: Checks if the current position is a steep face on the north or east sides of a mountain
- **Usage**:
  ```mms
  Steep()
  ```

#### Depth and Height Functions

**StoneDepth(floor | ceiling)**
- **Description**: Checks if the current position is within a specified distance from the surface
- **Parameters**:
  - `floor | ceiling`: Depth calculation mode
- **Configuration**:
  - `.Offset(float)`: Offset from surface
  - `.AddSurfaceDepth()`: Include surface depth in calculation
  - `.SecondaryDepthRange(float)`: Secondary depth range
- **Usage**:
  ```mms
  StoneDepth(floor | ceiling)
    .Offset(float)
    .AddSurfaceDepth()
    .SecondaryDepthRange(float)
  ```

**Water()**
- **Description**: Checks if the current position is above water, based on terrain depth
- **Configuration**:
  - `.Offset(float)`: Height offset for water check
  - `.DepthMultiplier(float)`: Depth calculation multiplier
  - `.AddStoneDepth()`: Include stone depth in calculation
- **Usage**:
  ```mms
  Water()
    .Offset(float)
    .DepthMultiplier(float)
    .AddStoneDepth()
  ```

**VerticalGradient(string)**
- **Description**: Compares the current Y position with a transition similar to deepslate and bedrock
- **Parameters**:
  - `string`: Gradient type identifier
- **Configuration**:
  - `.Lower(int | Abs())`: Lower bound (absolute or relative)
  - `.Upper(int | Abs())`: Upper bound (absolute or relative)
- **Usage**:
  ```mms
  VerticalGradient(string)
    .Lower(int | Abs())
    .Upper(int | Abs())
  ```

**YAbove(int | Abs())**
- **Description**: Checks if the current position is above a specified height (exclusive)
- **Parameters**:
  - `int | Abs()`: Height value or absolute reference
- **Configuration**:
  - `.DepthMultiplier(float)`: Depth calculation multiplier
  - `.AddStoneDepth()`: Include stone depth in calculation
- **Usage**:
  ```mms
  YAbove(int | Abs())
    .DepthMultiplier(float)
    .AddStoneDepth()
  ```

## Density Functions

Density functions control how terrain density is calculated across different dimensions.

### Density

#### Basic Operations

**Number**
- **Description**: Numeric literal value
- **Usage**: Direct numeric values

**Reference: DensityFunction**
- **Description**: Reference to a predefined density function
- **Usage**:
  ```mms
  namespace:name
  ```

#### Transformation Functions

**Interpolated(DensityFunction)**
- **Description**: Interpolates density values
- **Usage**:
  ```mms
  Interpolated(DensityFunction)
  ```

**Cache(flat | _2d | once | cell, DensityFunction)**
- **Description**: Caches density function results for performance
- **Parameters**:
  - `flat | _2d | once | cell`: Cache strategy
  - `DensityFunction`: Function to cache
- **Usage**:
  ```mms
  Cache(flat | _2d | once | cell, DensityFunction)
  ```

**Shift(a | b, ref(Noise) | Noise())**
- **Description**: Shifts density values using noise
- **Usage**:
  ```mms
  Shift(a | b, ref(Noise) | Noise())
  Shift(ref(Noise) | Noise())
  ```

#### Mathematical Operations

**Abs(DensityFunction)**
- **Description**: Absolute value of density function
- **Usage**:
  ```mms
  Abs(DensityFunction)
  ```

**Cube(DensityFunction)**
- **Description**: Cubes the density function values
- **Usage**:
  ```mms
  Cube(DensityFunction)
  ```

**Square(DensityFunction)**
- **Description**: Squares the density function values
- **Usage**:
  ```mms
  Square(DensityFunction)
  ```

**HalfNegative(DensityFunction)**
- **Description**: Applies half-negative transformation
- **Usage**:
  ```mms
  HalfNegative(DensityFunction)
  ```

**QuarterNegative(DensityFunction)**
- **Description**: Applies quarter-negative transformation
- **Usage**:
  ```mms
  QuarterNegative(DensityFunction)
  ```

**Squeeze(DensityFunction)**
- **Description**: Applies squeeze transformation
- **Usage**:
  ```mms
  Squeeze(DensityFunction)
  ```

**Invert(DensityFunction)**
- **Description**: Inverts density function values
- **Usage**:
  ```mms
  Invert(DensityFunction)
  ```

#### Arithmetic Operations

**Add(DensityFunction, DensityFunction)**
- **Description**: Adds two density functions
- **Usage**:
  ```mms
  Add(DensityFunction, DensityFunction)
  ```

**Mul(DensityFunction, DensityFunction)**
- **Description**: Multiplies two density functions
- **Usage**:
  ```mms
  Mul(DensityFunction, DensityFunction)
  ```

**Min(DensityFunction, DensityFunction)**
- **Description**: Minimum value of two density functions
- **Usage**:
  ```mms
  Min(DensityFunction, DensityFunction)
  ```

**Max(DensityFunction, DensityFunction)**
- **Description**: Maximum value of two density functions
- **Usage**:
  ```mms
  Max(DensityFunction, DensityFunction)
  ```

#### Noise-Based Functions

**OldBlendedNoise()**
- **Description**: Legacy blended noise implementation
- **Configuration**:
  - `.XzScale(float)`: X/Z axis scale factor
  - `.YScale(float)`: Y axis scale factor
  - `.XzFactor(float)`: X/Z multiplication factor
  - `.YFactor(float)`: Y multiplication factor
  - `.SmearScaleMul(float)`: Smearing scale multiplier
- **Usage**:
  ```mms
  OldBlendedNoise()
    .XzScale(float)
    .YScale(float)
    .XzFactor(float)
    .YFactor(float)
    .SmearScaleMul(float)
  ```

**Noise(ref(Noise) | Noise())**
- **Description**: Standard noise function
- **Configuration**:
  - `.XzScale(float)`: X/Z axis scale factor
  - `.YScale(float)`: Y axis scale factor
- **Usage**:
  ```mms
  Noise(ref(Noise) | Noise())
    .XzScale(float)
    .YScale(float)
  ```

**EndIslands()**
- **Description**: Specialized noise for End dimension islands
- **Usage**:
  ```mms
  EndIslands()
  ```

**WeirdScaledSampler(type_1 | type_2, ref(Noise) | Noise(), DensityFunction)**
- **Description**: Complex scaled sampling with multiple strategies
- **Usage**:
  ```mms
  WeirdScaledSampler(type_1 | type_2, ref(Noise) | Noise(), DensityFunction)
  ```

**ShiftedNoise(ref(Noise) | Noise())**
- **Description**: Noise function with configurable shifting
- **Configuration**:
  - `.XzScale(float)`: X/Z axis scale factor
  - `.YScale(float)`: Y axis scale factor
  - `.ShiftX(DensityFunction)`: X-axis shift function
  - `.ShiftY(DensityFunction)`: Y-axis shift function
  - `.ShiftZ(DensityFunction)`: Z-axis shift function
- **Usage**:
  ```mms
  ShiftedNoise(ref(Noise) | Noise())
    .XzScale(float)
    .YScale(float)
    .ShiftX(DensityFunction)
    .ShiftY(DensityFunction)
    .ShiftZ(DensityFunction)
  ```

#### Range and Selection Functions

**RangeChoice(DensityFunction)**
- **Description**: Chooses values based on input ranges
- **Configuration**:
  - `.Min(DensityFunction, DensityFunction)`: Minimum range mapping
  - `.Max(DensityFunction, DensityFunction)`: Maximum range mapping
  - `.InRange(DensityFunction)`: Value when input is in range
  - `.OutRange(DensityFunction)`: Value when input is out of range
- **Usage**:
  ```mms
  RangeChoice(DensityFunction)
    .Min(DensityFunction, DensityFunction)
    .Max(DensityFunction, DensityFunction)
    .InRange(DensityFunction)
    .OutRange(DensityFunction)
  ```

**Clamp(DensityFunction)**
- **Description**: Clamps density values to specified range
- **Configuration**:
  - `.Min(float)`: Minimum allowed value
  - `.Max(float)`: Maximum allowed value
- **Usage**:
  ```mms
  Clamp(DensityFunction)
    .Min(float)
    .Max(float)
  ```

#### Interpolation Functions

**Spline(DensityFunction)**
- **Description**: Creates smooth interpolation using spline curves
- **Configuration**:
  - `.Point(float, float, DensityFunction | Spline())`: Spline control points
- **Usage**:
  ```mms
  Spline(DensityFunction)
    .Point(float, float, DensityFunction | Spline())
  ```

**YClampedGradient()**
- **Description**: Creates gradient clamped to Y-axis bounds
- **Configuration**:
  - `.From(int, float)`: Starting point (Y, value)
  - `.To(int, float)`: Ending point (Y, value)
- **Usage**:
  ```mms
  YClampedGradient()
    .From(int, float)
    .To(int, float)
  ```

## Dimension Configuration

Dimension configuration defines how Minecraft dimensions are structured and generated.

### Dimension

**Dimension(ref(DimensionType) | overworld | the_nether | the_end | overworld_caves)**
- **Description**: Defines a Minecraft dimension with a type and generator settings
- **Parameters**:
  - `ref(DimensionType) | overworld | the_nether | the_end | overworld_caves`: Dimension type reference or preset
- **Configuration**:
  - `.Generator(debug)`: Generator type (currently only debug supported)
- **Usage**:
  ```mms
  Dimension(ref(DimensionType) | overworld | the_nether | the_end | overworld_caves)
    .Generator(debug)
  ```

**NoiseSettings()**
- **Description**: Configures noise generation parameters for a dimension
- **Configuration**:
  - `.SeaLevel(int)`: Sea level height
  - `.DisableMobGen()`: Disable mob generation
  - `.EnableOreVeins()`: Enable ore vein generation
  - `.Aquifers()`: Enable aquifer generation
  - `.DefaultBlock(ref)`: Default block type
  - `.DefaultFluid(ref)`: Default fluid type
  - `.SpawnTarget()`: Configure spawn targeting
  - `.MinY(float)`: Minimum Y coordinate
  - `.Height(float)`: Dimension height
  - `.Size(int, int)`: Horizontal and vertical size
  - `.NoiseRouter(ref(NoiseRouter))`: Noise router configuration
  - `.SurfaceRule(SurfaceRule | ref(SurfaceRule))`: Surface generation rules
- **Usage**:
  ```mms
  NoiseSettings()
    .SeaLevel(int)
    .DisableMobGen()
    .EnableOreVeins()
    .Aquifers()
    .DefaultBlock(ref)
    .DefaultFluid(ref)
    .SpawnTarget()
    .MinY(float)
    .Height(float)
    .Size(int, int)
    .NoiseRouter(ref(NoiseRouter))
    .SurfaceRule(SurfaceRule | ref(SurfaceRule))
  ```

**DimensionType()**
- **Description**: Defines properties of a dimension type
- **Configuration**:
  - `.Ultrawarm()`: Enable ultra-warm temperature effects
  - `.Natural()`: Enable natural dimension properties
  - `.Skylight()`: Enable sky light calculations
  - `.Ceiling()`: Enable ceiling height limits
  - `.PiglinSafe()`: Safe for piglin mob spawning
  - `.BedsWork()`: Allow beds to be used
  - `.AnchorsWork()`: Allow respawn anchors to work
  - `.HasRaids()`: Enable raid events
  - `.CoordinateScale(int)`: Coordinate scaling factor
  - `.AmbientLight(float)`: Base ambient light level
  - `.FixedTime(int)`: Fixed time of day (if applicable)
  - `.MonsterLightLevel(int)`: Light level for monster spawning
  - `.MonsterLightLimit(int)`: Maximum light for monster spawning
  - `.LogicalHeight(int)`: Logical height of dimension
  - `.CloudHeight(int)`: Cloud rendering height
  - `.MinY(int)`: Minimum Y coordinate
  - `.Height(int)`: Total dimension height
  - `.Infiniburn(tag(Block))`: Blocks that burn infinitely
  - `.Effects(minecraft:overworld | minecraft:the_nether | minecraft:the_end)`: Dimension effects
- **Usage**:
  ```mms
  DimensionType()
    .Ultrawarm()
    .Natural()
    .Skylight()
    .Ceiling()
    .PiglinSafe()
    .BedsWork()
    .AnchorsWork()
    .HasRaids()
    .CoordinateScale(int)
    .AmbientLight(float)
    .FixedTime(int)
    .MonsterLightLevel(int)
    .MonsterLightLimit(int)
    .LogicalHeight(int)
    .CloudHeight(int)
    .MinY(int)
    .Height(int)
    .Infiniburn(tag(Block))
    .Effects(minecraft:overworld | minecraft:the_nether | minecraft:the_end)
  ```

**NoiseRouter()**
- **Description**: Routes different noise functions for terrain generation
- **Configuration**:
  - `.PreliminarySurfaceLevel(DensityFunction)`: Initial surface level calculation
  - `.FinalDensity(DensityFunction)`: Final density values
  - `.Barrier(DensityFunction)`: Barrier placement density
  - `.FluidLevelFloodedness(DensityFunction)`: Fluid flooding calculations
  - `.FluidLevelSpread(DensityFunction)`: Fluid spread patterns
  - `.Lava(DensityFunction)`: Lava placement density
  - `.VeinToggle(DensityFunction)`: Ore vein generation toggle
  - `.VeinRidged(DensityFunction)`: Ore vein ridge patterns
  - `.VeinGap(DensityFunction)`: Ore vein gap calculations
  - `.Temperature(DensityFunction)`: Temperature variations
  - `.Vegetation(DensityFunction)`: Vegetation placement
  - `.Continents(DensityFunction)`: Continental landmass shapes
  - `.Erosion(DensityFunction)`: Terrain erosion patterns
  - `.Depth(DensityFunction)`: Depth-based calculations
  - `.Ridges(DensityFunction)`: Mountain ridge formations
- **Usage**:
  ```mms
  NoiseRouter()
    .PreliminarySurfaceLevel(DensityFunction)
    .FinalDensity(DensityFunction)
    .Barrier(DensityFunction)
    .FluidLevelFloodedness(DensityFunction)
    .FluidLevelSpread(DensityFunction)
    .Lava(DensityFunction)
    .VeinToggle(DensityFunction)
    .VeinRidged(DensityFunction)
    .VeinGap(DensityFunction)
    .Temperature(DensityFunction)
    .Vegetation(DensityFunction)
    .Continents(DensityFunction)
    .Erosion(DensityFunction)
    .Depth(DensityFunction)
    .Ridges(DensityFunction)
  ```

## Noise Functions

Noise functions generate the fundamental patterns used in terrain generation.

### Noise

**Noise(int)**
- **Description**: Defines a noise function with specified parameters
- **Parameters**:
  - `int`: Noise function identifier
- **Configuration**:
  - `.Amplitudes(float, float...)`: Amplitude values for different octaves
- **Usage**:
  ```mms
  Noise(int)
    .Amplitudes(float, float...)