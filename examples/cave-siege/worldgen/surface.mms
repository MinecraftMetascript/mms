// What the caves are made of. The rules are tried in order, and the first
// one that matches places its block; anything left over stays stone.
//
// StoneDepth(Floor) is the top block of a floor, StoneDepth(Ceiling) the
// bottom block of a roof. `.Offset(n)` reaches n blocks further in.

SurfaceCondition grassy = Biome(glowmeadow, rootwood_hollow)

/** A floor block, or one of the 3 blocks under it. */
SurfaceCondition topsoil = StoneDepth(Floor).Offset(3)

/** Bedrock at the bottom: always y -64, thinning out by y -60. */
SurfaceRule bedrock_floor = If(
  VerticalGradient("siege:bedrock_floor", AboveBottom(0), AboveBottom(5)),
) { Block(bedrock) }

/** Bedrock at the top, so the roof can't be mined through to the build limit. */
SurfaceRule bedrock_roof = If(!VerticalGradient("siege:bedrock_roof", BelowTop(5), BelowTop(0))) {
  Block(bedrock)
}

/** Grass on top and dirt below, in the farm and wood biomes. */
SurfaceRule meadow_floor = If(grassy) {
  If(StoneDepth(Floor)) { Block(grass_block) }
  If(topsoil) { Block(dirt) }
}

/** Podzol and coarse dirt patches among the grass, then rooted dirt. */
SurfaceRule darkroot_floor = If(Biome(darkroot_grove)) {
  If(StoneDepth(Floor)) {
    If(NoiseThreshold(minecraft:surface).Min(0.2).Max(1)) { Block(podzol) }
    If(NoiseThreshold(minecraft:surface).Min(-1).Max(-0.6)) { Block(coarse_dirt) }
    Block(grass_block)
  }
  If(topsoil) { Block(dirt) }
  If(StoneDepth(Ceiling)) { Block(rooted_dirt) }
}

SurfaceRule mycelium_floor = If(Biome(mycelium_deeps)) {
  If(StoneDepth(Floor)) { Block(mycelium) }
  If(topsoil) { Block(dirt) }
}

/** Sand shores 4 deep on sandstone. Roofs are sandstone so they don't fall. */
SurfaceRule lake_floor = If(Biome(sunken_lake)) {
  If(StoneDepth(Floor).Offset(3)) { Block(sand) }
  If(StoneDepth(Floor).Offset(6)) { Block(sandstone) }
  If(StoneDepth(Ceiling)) { Block(sandstone) }
}

/** Basalt and blackstone floors with magma seams; blackstone roofs. */
SurfaceRule ash_floor = If(Biome(ash_vents)) {
  If(StoneDepth(Floor)) {
    If(NoiseThreshold(minecraft:surface).Min(0.4).Max(1)) { Block(magma_block) }
    If(NoiseThreshold(minecraft:surface).Min(-0.2).Max(0.4)) { Block(basalt) }
    Block(blackstone)
  }
  If(StoneDepth(Floor).Offset(2)) { Block(basalt) }
  If(StoneDepth(Ceiling).Offset(1)) { Block(blackstone) }
}

/** Deepslate below y 0, blending over y 0 to 8, as in the overworld. */
SurfaceRule deepslate = If(VerticalGradient("siege:deepslate", Absolute(0), Absolute(8))) {
  Block(deepslate)
}

SurfaceRule surface = [
  bedrock_floor,
  bedrock_roof,
  meadow_floor,
  darkroot_floor,
  mycelium_floor,
  lake_floor,
  ash_floor,
  deepslate,
]
