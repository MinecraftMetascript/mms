Namespace mms_demo

SurfaceCondition in_friendly_biome = Biome(forest, plains, beach)
SurfaceCondition in_unfriendly_biome = Biome(desert, badlands, deep_ocean)

SurfaceRule honey_surface = Block(honey_block)
SurfaceRule slime_surface = Block(slime_block)

SurfaceRule my_strange_surface = [
  If(in_friendly_biome) { honey_surface }
  If(in_unfriendly_biome) { slime_surface }
  Block(magma_block)
]

SurfaceRule terracotta_bands = If(StoneDepth(Floor)) {
  If(
    NoiseThreshold(minecraft:surface).Min(-0.909).Max(-0.5454)
      || NoiseThreshold(minecraft:surface, -0.1818, 0.1818)
  ) { Block(terracotta) }
  Bandlands()
}

SurfaceRule white_terracotta = If(AboveWater().Offset(-6).Mul(-1).Add()) { Block(white_terracotta) }

SurfaceRule edge = If(YAbove(63) && !YAbove(74).Mul(1).Add()) { Block(orange_terracotta) }

SurfaceRule demorgan = If(!(Hole() || Steep())) { Block(oak_log[axis=y]) }

SurfaceRule bad_state = Block(oak_log[axis=q, colour=red])
SurfaceRule bad_cond = If(Biome(nowhere)) { Block(stone) }
