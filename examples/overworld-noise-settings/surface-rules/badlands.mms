namespace minecraft {
  Surface {
    InBadlands = Biome(
      minecraft:badlands,
      minecraft:eroded_badlands,
      minecraft:wooded_badlands
    )

    SkyTerracotta = If (AboveSurface()) Block(minecraft:stone)

    NonHoleOrangeTerracotta = If (!Hole()) Block(minecraft:orange_terracotta)

    TerracottaBands =
        If (StoneDepth(Floor)) [
          If (
            Or (
              NoiseThreshold(minecraft:surface).Min(-0.909).Max(-0.5454)
              NoiseThreshold(minecraft:surface).Min(-0.1818).Max(0.1818)
              NoiseThreshold(minecraft:surface).Min(0.5454).Max(0.909)
            )
          ) Block(minecraft:terracotta)
          Bandlands()
        ]

    WhiteTerracotta =
      If (AboveWater().Offset(-6).Mul(-1).Add())
        Block(minecraft:white_terracotta)

    OrangeTerracotta =
      If (YAbove(63))
        Block(minecraft:orange_terracotta)

    StoneAndGravel = [
        If (StoneDepth(Ceiling))
          Block(minecraft:stone)
        Block(minecraft:gravel)
      ]

    SurfaceSands =
      If (AboveWater().Offset(-1)) [
        If (StoneDepth(Ceiling))
          Block(minecraft:red_sandstone)
        Block(minecraft:red_sand)
      ]

    OrangeTerracottaEdge = If (
      And (
        YAbove(63)
        !YAbove(74).Mul(1).Add()
      )
    ) Block(minecraft:orange_terracotta)

    Badlands = If (InBadlands) [
      If (YAbove(63)) [
        SkyTerracotta
        TerracottaBands
        SurfaceSands
        NonHoleOrangeTerracotta
        WhiteTerracotta
        StoneAndGravel
      ]

      If (YAbove(63).Mul(-1).Add()) [
        OrangeTerracottaEdge
        Bandlands()
      ]

      If (StoneDepth(Floor)) [
        WhiteTerracotta
      ]
    ]
  }
}
