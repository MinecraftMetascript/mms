# Dimension Test Files

This directory contains comprehensive test files demonstrating all aspects of the Minecraft Metascript dimension system.

## Files Overview

### 1. `dimension_types.mms`
**Purpose**: Demonstrates different dimension type configurations

**Examples**:
- **OverworldType**: Standard overworld-like dimension with normal lighting and physics
- **NetherType**: Nether-like dimension with ultrawarm properties and different coordinate scaling
- **EndType**: End-like dimension with no natural lighting and void-like properties
- **SkyDimension**: Custom dimension with fixed daytime and elevated minimum Y

**Key Features Shown**:
- All basic dimension type properties (ultrawarm, natural, skylight, etc.)
- Coordinate scaling and ambient lighting
- Monster spawn limits and logical height
- Infiniburn and effects configuration

### 2. `flat_dimensions.mms`
**Purpose**: Shows various flat world configurations

**Examples**:
- **ClassicFlat**: Traditional superflat with bedrock, stone, dirt, and grass
- **DesertFlat**: Desert-themed flat world with sandstone and sand
- **OceanFlat**: Water world with deep water layer
- **SnowFlat**: Snowy flat world for winter themes
- **LayeredFlat**: Complex multi-layer geological flat world

**Key Features Shown**:
- Flat layer definitions with blocks and heights
- Structure placement using tags
- Lake and feature configuration
- Biome assignment for flat worlds

### 3. `noise_dimensions.mms`
**Purpose**: Demonstrates noise-based dimensions with different biome sources

**Examples**:
- **PlainsOnly**: Single biome dimension using fixed biome source
- **CheckerboardBiomes**: Multiple biomes in checkerboard pattern
- **OverworldPreset**: Standard overworld biome distribution
- **NetherPreset**: Nether biome distribution
- **TheEnd**: End dimension with void biome source
- **CustomNoise**: Using custom noise settings

**Key Features Shown**:
- All four biome source types (fixed, checkerboard, multi_noise, the_end)
- Preset biome configurations (overworld, nether)
- Custom noise settings integration
- Biome source parameter tuning

### 4. `custom_multinoise.mms`
**Purpose**: Advanced multi-noise biome configuration

**Examples**:
- **CustomBiomes**: Full custom biome placement with specific noise parameters
- **BiomeTagCheckerboard**: Using biome tags instead of individual biome IDs
- **MixedBiomes**: Mixing individual biomes and biome tags

**Key Features Shown**:
- Custom noise parameter definitions for biomes
- Temperature, humidity, continentalness, erosion, weirdness, depth, and offset parameters
- Biome tag usage for dynamic biome selection
- Complex biome placement strategies

### 5. `complete_dimension_example.mms`
**Purpose**: Comprehensive examples showing advanced configurations

**Examples**:
- **DebugDimension**: Simple debug world for testing
- **ComplexFlat**: Multi-layer flat world with rich structure generation
- **AdvancedNoise**: Complex noise world with ore veins and aquifers
- **SkyIslands**: Floating island world using high sea level

**Key Features Shown**:
- Advanced noise settings (ore veins, aquifers, mob generation control)
- Complex flat layer geology
- High sea level for floating island generation
- Integration of multiple systems

## Usage

Each file can be used independently or in combination:

```mms
// Use individual dimension types
UseNamespace "test_dimensions"

// Reference specific dimensions
Dimension myDimension = test_dimensions:CustomBiomes
```

## Testing Checklist

When testing dimension generation, verify:

- [ ] Dimension types export correct JSON properties
- [ ] Flat generators create expected layer structures
- [ ] Noise generators use correct biome sources
- [ ] Multi-noise biomes respect parameter boundaries
- [ ] Checkerboard patterns generate correctly
- [ ] Custom noise settings integrate properly
- [ ] File exports to correct worldgen directories

## Generated Files

These test files will generate JSON files in:
- `worldgen/dimension_type/` - Custom dimension types
- `worldgen/flat/` - Flat world settings
- `worldgen/biome_source/` - Biome source configurations
- `worldgen/dimension/` - Complete dimension definitions

## Best Practices Demonstrated

- Consistent namespace usage
- Proper dimension type configuration
- Layered approach to complexity (simple to advanced)
- Documentation of parameters and their effects
- Reusable dimension types across multiple dimensions