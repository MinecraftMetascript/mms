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