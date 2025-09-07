namespace xyz {
    Noise {
        NoiseRef = Noise(-6)
    }

  DensityFn {
        old = WeirdScaledSampler(5).Type1().Noise(NoiseRef)
        old2 = WeirdScaledSampler(5).Type1().Noise(-6).Amplitudes(2)
  }
}