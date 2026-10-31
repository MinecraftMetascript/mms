SurfaceCondition in_friendly_biome = Biome(forest, plains, beach)
SurfaceCondition in_unfriendly_biome = Biome(desert, badlands, deep_ocean)

SurfaceRule honey_surface = Block(honey_block)
SurfaceRule slime_surface = Block(slime_block)

SurfaceRule my_strange_surface = [
  If(in_friendly_biome) { honey_surface },
  If(in_unfriendly_biome) { slime_surface },
  Block(magma_block),
]
