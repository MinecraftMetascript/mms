Namespace example

Noise ridges_noise = { base_octave = -7, octave_count = 2 }

DensityFunction continents = Cache(Noise(continentalness, 0.25, 0))

DensityFunction shaped = Clamp(continents * 2 + 0.5 - Abs(ridges_noise_df), -1, 1)

DensityFunction ridges_noise_df = Noise(ridges_noise).XzScale(1).YScale(0)

DensityFunction constant = 1 + 2 * 3

DensityFunction typo = Clmp(continents, -1, 1)
DensityFunction missing = Noise(minecraft:nope, 1, 1)
DensityFunction neg = -continents
