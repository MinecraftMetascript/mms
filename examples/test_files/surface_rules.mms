Namespace test_files {
	Noise {
		Xyz = Noise(-1).Amplitudes(1)
	}
	Surface {
		Thresh = NoiseThreshold(test_files:Xyz).Max()
		X = NoiseThreshold(test_files:Xyz).Min(5)
	}
}
