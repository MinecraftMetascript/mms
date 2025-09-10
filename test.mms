Namespace xyz {
    NoiseSettings {
        someDimension = NoiseSettings()
            .SeaLevel(5)
            .DefaultBlock(stone)
            .DefaultFluid(Block(water))
            .NoiseSize(3,3)
    }

    Noise {
        n = Noise(-10).Amplitudes(5,6,7)
    }
    DensityFn {
        r = Noise(myNoise)
    }
    NoiseRouter {
        o = Router()
            .FinalDensity(
                YClampedGradient()
                    .Min(0)
                    .Max(100)
                    .Bottom(-10)
                    .Top(10)
            )
            .Temperature(
                Noise(-5).Amplitudes(1,2,3)
            )
            .Erosion(n)
    }
}

