# Placed features, block predicates, block state providers and rule tests

A placed feature's `placement` is a list of modifiers that run in order.
Each one turns every position it gets into zero, one or more positions,
starting from the corner of the chunk.

## minecraft:placement_modifier/count
Repeats each position `count` times. Follow it with `InSquare()` to spread the copies over the chunk.

- count: How many: a number, or a range such as `2..5`.

```mms
PlacedFeature trees = { feature = minecraft:oak, placement = [Count(3), InSquare(), Heightmap(WORLD_SURFACE_WG), Biome()] }
```

## minecraft:placement_modifier/in_square
Moves each position to a random x and z within its chunk.

## minecraft:placement_modifier/heightmap
Moves each position up or down onto the surface of a heightmap, such as `WORLD_SURFACE_WG` for the top block or `OCEAN_FLOOR_WG` for the sea floor.

- heightmap: Which surface to use.

## minecraft:placement_modifier/height_range
Sets each position's Y to a height picked by a height provider.

- height: The height, such as `Uniform(AboveBottom(0), Absolute(64))`.

## minecraft:placement_modifier/biome
Drops positions whose biome doesn't list this feature. Vanilla ends most placements with it.

## minecraft:placement_modifier/rarity_filter
Keeps each position with a chance of 1 in `chance`.

- chance: Keep one position in this many, on average.

```mms
PlacedFeature rare_rocks = { feature = minecraft:forest_rock, placement = [RarityFilter(16), InSquare(), Heightmap(MOTION_BLOCKING), Biome()] }
```

## minecraft:placement_modifier/random_chance
Keeps each position with probability `chance`, from 0 to 1.

## minecraft:placement_modifier/offset
Moves each position by the given number of blocks.

## minecraft:placement_modifier/block_predicate_filter
Keeps positions where a block predicate passes, such as `Solid()` or `MatchingBlocks(air)`.

## minecraft:placement_modifier/count_on_every_layer
Places `count` positions on each layer of the terrain, from the top down, so features appear under overhangs too. Vanilla nether forests use it.

## minecraft:placement_modifier/environment_scan
Moves each position up or down, at most `max_steps` blocks, until `target_condition` passes. Drops it if nothing is found.

- direction_of_search: `up` or `down`.
- max_steps: How far to search, in blocks.
- target_condition: The block predicate to search for.
- allowed_search_condition: The scan stops if a block on the way fails this.

## minecraft:placement_modifier/noise_based_count
Repeats each position a number of times that follows a noise, so features cluster in some areas. The count is `ceil((noise + noise_offset) * noise_to_count_ratio)`.

- noise_to_count_ratio: Multiplies the noise to get the count.
- noise_factor: Divides x and z before sampling; larger values give larger clusters.

## minecraft:placement_modifier/noise_threshold_count
Repeats each position `below_noise` times where a noise is below `noise_level`, and `above_noise` times where it is above.

## minecraft:placement_modifier/surface_relative_threshold_filter
Keeps positions whose Y, measured from a heightmap's surface, is from `min_inclusive` to `max_inclusive`.

## minecraft:placement_modifier/surface_water_depth_filter
Keeps positions where the water above the surface is at most `max_water_depth` blocks deep.

## minecraft:placement_modifier/fixed_placement
Replaces each position with a fixed list of positions.

## minecraft:block_predicate/matching_blocks
Passes when the block is one of these.

- blocks: The blocks that pass.

## minecraft:block_predicate/matching_block_tag
Passes when the block is in a block tag.

## minecraft:block_predicate/matching_fluids
Passes when the fluid at the block is one of these.

## minecraft:block_predicate/matching_biomes
Passes when the block's biome is one of these.

## minecraft:block_predicate/solid
Passes when the block is solid.

## minecraft:block_predicate/replaceable
Passes when the block can be replaced, as air, water and grass can.

## minecraft:block_predicate/would_survive
Passes when a block would stay put here, as a sapling needs dirt below.

- state: The block to test.

## minecraft:block_predicate/has_sturdy_face
Passes when the block has a full face on one side.

- direction: The side to check.

## minecraft:block_predicate/inside_world_bounds
Passes when the block is inside the world's height limits.

## minecraft:block_predicate/unobstructed
Passes when no entity is in the way of the block.

## minecraft:block_predicate/all_of
Passes when every predicate passes.

## minecraft:block_predicate/any_of
Passes when any predicate passes.

## minecraft:block_predicate/not
Passes when `predicate` fails.

## minecraft:block_predicate/true
Always passes.

## minecraft:block_state_provider/simple
Always the same block.

- state: The block, with optional state such as `oak_log[axis = y]`.

## minecraft:block_state_provider/weighted
Picks a block at random, by weight.

## minecraft:block_state_provider/random_block
Picks one of these blocks at random, each equally likely.

## minecraft:block_state_provider/rotated
A block with an `axis` property, turned to match the direction it is placed in, as logs are.

## minecraft:block_state_provider/noise
Picks from `states` by a noise, so the same block appears in patches.

- seed: Seeds the noise.
- scale: Scales x and z before sampling; smaller values give larger patches.
- states: The blocks to pick from.

## minecraft:block_state_provider/noise_threshold
Uses `default_state` where a noise is below `threshold`, and one of `low_states` or `high_states` above it. Vanilla flower forests use it.

## minecraft:block_state_provider/dual_noise
Like `Noise`, with a second, slower noise that picks how many of `states` can appear in each area.

## minecraft:block_state_provider/randomized_int
Another provider's block, with an integer property set at random, such as a flower's `flower_amount`.

- property: The property to set.
- values: The values to pick from.
- source: The provider that picks the block.

## minecraft:block_state_provider/rule_based
Uses the first rule whose block predicate passes at the position, or `fallback` when none does.

## minecraft:rule_test/always_true
Always passes.

## minecraft:rule_test/block_match
Passes when the block is this block.

## minecraft:rule_test/blockstate_match
Passes when the block is exactly this block state.

## minecraft:rule_test/tag_match
Passes when the block is in this block tag. Vanilla ores test `#stone_ore_replaceables` this way.

## minecraft:rule_test/random_block_match
Passes when the block is this block, with probability `probability`.

## minecraft:rule_test/random_blockstate_match
Passes when the block is exactly this block state, with probability `probability`.

## minecraft:biome_source/fixed
The same biome everywhere.

- biome: The biome.

## minecraft:biome_source/checkerboard
Biomes in a checkerboard of square tiles, repeating the list in order.

- scale: Tile size: each tile is 2^(scale + 2) blocks across. From 0 to 62; defaults to 2.
- biomes: The biomes, in order.

## minecraft:biome_source/multi_noise
Picks biomes by climate noise, as the overworld and nether do. `preset` uses vanilla's tables; otherwise list biomes with their climate parameters.

## minecraft:biome_source/the_end
The End's biomes: the central island, highlands, midlands, small islands and barrens.

## minecraft:chunk_generator/noise
Generates terrain from noise settings and density functions, and biomes from a biome source. Vanilla dimensions use it.

- settings: The noise settings: a `NoiseSettings` declaration or a vanilla one such as `minecraft:overworld`.
- biome_source: Where biomes come from.

## minecraft:chunk_generator/flat
A superflat world: layers of blocks and one biome.

## minecraft:chunk_generator/debug
The debug world: every block state in a grid.
