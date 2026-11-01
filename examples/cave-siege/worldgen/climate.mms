// Where each biome goes. The router's climate functions are flat noises
// (they don't change with y), except depth, which falls from 1.5 at the
// bedrock floor to -1 at the roof. So biomes form wide regions that change
// with height too: Deep Dark and Darkroot Grove sit low, Glowmeadow and
// Rootwood Hollow sit in the middle, and Sunken Lake and Ash Vents can
// show up at any height.
//
// The noises are scaled by 1.6 so that the -0.25 and 0.25 splits used in
// dimension.mms cut each climate into roughly equal thirds.

/**
 * 0 at x 0, z 0, rising to 1 about 200 blocks out. The climate is faded
 * towards 0 near the origin, which is Glowmeadow's climate, so the
 * starter cavern is always in a lit farm biome.
 */
DensityFunction spawn_fade = Clamp(
  Square(
    Gradient(axis = x, from_coordinate = -200, to_coordinate = 200, from_value = -1, to_value = 1),
  ) +
    Square(
      Gradient(
        axis = z,
        from_coordinate = -200,
        to_coordinate = 200,
        from_value = -1,
        to_value = 1,
      ),
    ),
  0,
  1,
)

Noise climate_temperature = {
  base_octave = -10
  amplitude_modifiers = [1, 1, 0.5]
}

Noise climate_humidity = {
  base_octave = -9
  amplitude_modifiers = [1, 1, 0.5]
}

/** Picks the rare regions: Mycelium Deeps where it is low. */
Noise climate_continents = {
  base_octave = -9
  amplitude_modifiers = [1, 0.5]
}

Noise climate_erosion = {
  base_octave = -8
  amplitude_modifiers = [1, 0.5]
}

DensityFunction climate_temperature = Noise(climate_temperature, 1, 0) * 1.6 * spawn_fade

DensityFunction climate_humidity = Noise(climate_humidity, 1, 0) * 1.6 * spawn_fade

DensityFunction climate_continents = Noise(climate_continents, 1, 0) * 1.6 * spawn_fade

DensityFunction climate_erosion = Noise(climate_erosion, 1, 0) * 1.6

/** 1.5 at y -64, 0 at y 90, -1 at y 192. */
DensityFunction climate_depth = Gradient(
  axis = y,
  from_coordinate = -64,
  to_coordinate = 192,
  from_value = 1.5,
  to_value = -1,
)
