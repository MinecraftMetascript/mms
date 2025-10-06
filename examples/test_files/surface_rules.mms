Namespace test_files {
/*
	Noise {
		Xyz = Noise(-1).Amplitudes(1,2,3)


		Zyx = Noise(-2).Amplitudes(1,2,3)
	}
*/
	Surface {
		Cond = If(!Frozen() || Hole()) Bandlands()
		/*Seq = [
			Block(minecraft:air)
			Block(minecraft:stone)

		]*/
	}
}
