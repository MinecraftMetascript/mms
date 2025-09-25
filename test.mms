Namespace numberTwo {
  Noise {
    blah = Noise(10)
  }
}

Namespace xyz {
    Noise {
        x = Noise(4).Amplitudes(5)

        another_noise = Noise(5)
    }

  DensityFn {
    Xyz = Noise(xyz:another_noise).XZScale(123).YScale(123)
    TestNoise = Noise(-5).Amplitudes(5)
    TestFn = YClampedGradient().Min(0).Max(1).Bottom(1).Top(4)
  }
/*
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
