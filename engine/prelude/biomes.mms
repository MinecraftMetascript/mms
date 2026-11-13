// CustomBiome: a biome with where it goes and what its ground is made of,
// written in MMS on top of the Biome and SurfaceRule it emits.

/**
 * A biome with its place in a multi-noise biome source and, optionally,
 * its own surface:
 * `CustomBiome ash_fields { climate = { temperature = [0.5, 1], humidity = [-1, 0], continentalness = 0, erosion = 0, weirdness = 0, depth = 0, offset = 0 }, biome = { ... }, surface_rule = Block(basalt) }`.
 * Its name stands for the biome. `ash_fields.noise_entries` are its
 * entries for a biome source, as in
 * `MultiNoise(biomes = [...ash_fields.noise_entries, ...frost_fields.noise_entries])`,
 * and with a `surface_rule` the surface rule `ash_fields` is that rule
 * kept to the biome, for the noise settings' `material_rule`.
 */
Blueprint CustomBiome {
  /**
   * Where the biome goes: the climate it wins at in a multi-noise biome
   * source, or a list of them for a biome found in more than one place.
   * Each parameter is a number or a `[min, max]` range.
   */
  climate: ClimateParameters | [ClimateParameters]
  /** The biome itself: its weather, colors, mobs, carvers and features. */
  biome: Biome
  /**
   * The blocks on its ground, such as `If(StoneDepth(Floor)) { Block(grass_block) }`.
   * It writes the surface rule of the biome's name, which applies this
   * rule only in the biome. Put that name in the noise settings'
   * `material_rule` with the rules of the other biomes.
   */
  surface_rule: SurfaceRule?

  let climates = if (kind_of(climate) == "list") climate else [climate]

  emit {
    primary Biome ${self.path} = biome

    /** The biome's entries for a multi-noise biome source's `biomes`, one for each climate. */
    const noise_entries: [BiomeNoiseEntry] = [
      for c in climates:
        { biome = self.id, parameters = c }
    ]

    if (surface_rule != none) {
      SurfaceRule ${self.path} = If(Biome(self.id)) { surface_rule }
    }
  }
}
