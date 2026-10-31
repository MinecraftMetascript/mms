// 26.3 noises use base_octave and amplitude_modifiers, not firstOctave and
// amplitudes (see docs/decisions.md, D23).
Noise my_noise = {
  base_octave = -1
  amplitude_modifiers = [1, 2, 0.5, 1]
}
