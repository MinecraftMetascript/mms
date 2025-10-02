Namespace test_files {
	Noise {
		Xyz = Noise(-1).Amplitudes(1,2,3)
	}
	Surface {
		Thresh = NoiseThreshold(test_files:Xyz)
		X = NoiseThreshold(test_files:Xyz).Max()
	}
}
