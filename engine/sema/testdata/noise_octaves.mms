// octave_count is filled in from amplitude_modifiers when left out.
Noise inferred = { base_octave = -7, amplitude_modifiers = [1, 0.6, 0.3, 0.15] }
Noise matching = { base_octave = -7, octave_count = 2, amplitude_modifiers = [1, 1] }
Noise empty = { base_octave = -7, amplitude_modifiers = [] }
Noise plain = { base_octave = -7 }
Noise wrong = { base_octave = -7, octave_count = 1, amplitude_modifiers = [1, 1] }
