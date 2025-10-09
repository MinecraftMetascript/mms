// Use the sidebar to select one of the noise symbols to open the preview panel
// In the future, there will be a button next to the variable name in the editor
// to open the preview more quickly
Namespace mms_demo {
	Dimension {
		TheZone = Dimension(overworld).Generator(debug)
		T = NoiseRouter().Barrier(EndIslands())
		
		S = NoiseSettings()
			.NoiseRouter().Barrier().FinalDensity()
			.SurfaceRule([])
	}

  Density {
	/**
	 * This is something that will need to be addressed.
	 */
	ZZZ = Noise(BigNoise).XzScale(1).YScale(2)
	Woah = Clamp(Noise(Noise(5).Amplitudes(1,2,3)).XzScale(1)).Max(5).Min(-5)
  /**
   * This is a docstring
   *       `Multiple lines are in play!`
   */
    SomethingWonderful = 5
  }
  Noise {
	/** Test Doc */
    BigNoise = Noise(-5).Amplitudes(5)
    // SmallNoise = Noise(-3).Amplitudes(1, 5, 10)
  }
}
