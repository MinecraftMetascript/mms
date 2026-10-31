/** The three badlands biomes. */
SurfaceCondition in_badlands = Biome(badlands, eroded_badlands, wooded_badlands)

SurfaceRule sky_terracotta = If(AboveSurface()) { Block(stone) }

SurfaceRule non_hole_orange_terracotta = If(!Hole()) { Block(orange_terracotta) }

SurfaceRule terracotta_bands = If(StoneDepth(Floor)) {
  If(
    NoiseThreshold(surface).Min(-0.909).Max(-0.5454)
      || NoiseThreshold(surface).Min(-0.1818).Max(0.1818)
      || NoiseThreshold(surface).Min(0.5454).Max(0.909),
  ) { Block(terracotta) }
  Bandlands()
}

SurfaceRule white_terracotta = If(AboveWater().Offset(-6).Mul(-1).Add()) { Block(white_terracotta) }

SurfaceRule orange_terracotta = If(YAbove(63)) { Block(orange_terracotta) }

SurfaceRule stone_and_gravel = [
  If(StoneDepth(Ceiling)) { Block(stone) },
  Block(gravel),
]

SurfaceRule surface_sands = If(AboveWater().Offset(-1)) {
  If(StoneDepth(Ceiling)) { Block(red_sandstone) }
  Block(red_sand)
}

SurfaceRule orange_terracotta_edge = If(YAbove(63) && !YAbove(74).Mul(1).Add()) {
  Block(orange_terracotta)
}

SurfaceRule badlands = If(in_badlands) {
  If(YAbove(63)) {
    sky_terracotta
    terracotta_bands
    surface_sands
    non_hole_orange_terracotta
    white_terracotta
    stone_and_gravel
  }
  If(YAbove(63).Mul(-1).Add()) {
    orange_terracotta_edge
    Bandlands()
  }
  If(StoneDepth(Floor)) { white_terracotta }
}
