Namespace xyz {
    Noise {
        x = Noise(4)
    }

  /* DensityFn {
    TestNoise = Noise(-5).Amplitudes(5)
    TestFn = YClampedGradient().Min(0).Max(1).Bottom(1).Top(4)
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
  } */
}