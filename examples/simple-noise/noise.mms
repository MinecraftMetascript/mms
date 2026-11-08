// A noise and the terrain density it shapes. Put the cursor on either to
// preview it. 26.3 noises use base_octave and amplitude_modifiers, not
// firstOctave and amplitudes.

/** Rolling noise: one octave above the base, with a strong second one. */
Noise rolling {
  base_octave = -1
  amplitude_modifiers = [1, 2, 0.5, 1]
}

/** Gentle hills: the noise, halved and lifted so most of it is solid. */
// mms-ignore unused
DensityFunction hills = Noise(rolling, 1, 1) * 0.5 + 0.25
