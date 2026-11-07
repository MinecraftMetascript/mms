// Surface rules built from named conditions and rules: the first rule
// whose condition holds places its block.

SurfaceCondition in_friendly_biome = Biome(forest, plains, beach)
SurfaceCondition in_unfriendly_biome = Biome(desert, badlands, deep_ocean)

SurfaceRule honey_surface = Block(honey_block)
SurfaceRule slime_surface = Block(slime_block)

/** Honey in friendly biomes, slime in unfriendly ones, magma elsewhere. */
// mms-ignore unused
SurfaceRule my_strange_surface = [
  If(in_friendly_biome) { honey_surface },
  If(in_unfriendly_biome) { slime_surface },
  Block(magma_block),
]
