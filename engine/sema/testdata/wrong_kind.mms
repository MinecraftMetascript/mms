// A bare name that matches a local declaration of another kind says so,
// instead of reporting an unknown minecraft:<name>.
Noise ridges = { base_octave = -7, octave_count = 2 }

DensityFunction terrain = Abs(ridges)

DensityFunction sampled = Abs(Noise(ridges, 1, 1))

// A name with no local declaration still falls back to minecraft.
DensityFunction missing = Abs(nowhere)
