// Nested helper calls step down one level per line once their parent breaks.
DensityFunction compound = RangeChoice(y, 128, 256, l1, RangeChoice(y, 0, 128, l1, RangeChoice(y, -128, 0, l1, RangeChoice(y, -256, -128, l1, 0))))

// The same chain stays on one line when it fits.
DensityFunction short = RangeChoice(y, 0, 1, a, RangeChoice(y, 1, 2, b, c))

// Short nested calls stay on one line inside a broken call.
DensityFunction mixed = Add(Mul(base_noise_value, 2), Clamp(Noise(another_noise, 1, 1), -1, 1), Abs(offset_value))

// A long nested argument breaks even though it would fit on its own line.
DensityFunction named = Interpolated(argument = Add(Noise(first_noise, 1, 1), Mul(second_value, 2)), xz = 4, y = 8)

// One-argument wrappers keep the plain fits-or-breaks rule.
DensityFunction wrapped = Spline(Not(MatchingBlockTag(tag = #minecraft:features_cannot_replace)), extra_value)
