// Use the sidebar to select one of the noise symbols to open the preview panel
// In the future, there will be a button next to the variable name in the editor
// to open the preview more quickly
Namespace mms_demo {
  Density {
    X = OldBlendedNoise().XzFactor(7)

    Y = Noise(mms_demo:BigNoise).XzScale(5).YScale(5)
  }
  Noise {
    BigNoise = Noise(-5).Amplitudes(5)
    SmallNoise = Noise(-3).Amplitudes(1, 5, 10)
  }

  Surface {
    MyRule = Frozen()

    MySequence = If( AboveSurface() && mms_demo:MyRule ) Block(stone)

    ConditionalNoise = [
    If(Biome(forest)) Block(sand)
    ]

    ConditionalBlock = Block(s)
    X = NoiseThreshold(Noise(-10).Amplitudes(-1,2,5)).Min(2).Max(5)
    Woah = Frozen()
    Lol = StoneDepth(floor).AddSurfaceDepth().Offset(5).SecondaryDepthRange(2)
  }
}
