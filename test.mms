Namespace xyz {
  DensityFn {
    TestNoise = Noise(-5).Amplitudes(5)
  }

  NoiseSettings {
    MyNoise = NoiseSettings()
      .NoiseRouter(
        Router()
          .FinalDensity(1)
      )
      .SurfaceRule(Block(stone))
      .DefaultBlock(stone)
      .DefaultFluid(water)
  }
}