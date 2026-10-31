// Vanilla 26.3 worldgen files rewritten in MMS. TestVanillaParity builds
// this file and compares each output with the vanilla file of the same id.
Namespace minecraft

Noise ridge = {
  base_octave = -7
  base_amplitude = 0.9147152149950137
  amplitude_modifiers = [1, 2, 1, 0, 0, 0]
  octave_count = 6
}

DensityFunction overworld/ridges = Cache(Noise(ridge, 0.25, 0).ShiftX(shift_x).ShiftZ(shift_z))

DensityFunction overworld/ridges_folded = (Abs(Abs(overworld/ridges) + -0.6666667) + -0.33333334) *
  -3

DensityFunction overworld/ore_vein/gap = -0.3 - Noise(ore_gap, 1, 1)

DensityFunction overworld/ore_vein/toggle = Cache(
  Interpolated(RangeChoice(y, -64, 57, Noise(ore_veininess, 1.5, 1.5), 0))
    .CellSizeXz(4)
    .CellSizeY(8),
)

DensityFunction end/base_3d_noise = OldBlendedNoise(0.25, 0.25, 80, 160, 4)
