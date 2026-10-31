// A multi-noise source without a preset lists its biomes inline (the
// %none case of minecraft:multi_noise_biome_source).
Dimension inline = {
  type = overworld
  generator = {
    type = noise
    settings = overworld
    biome_source = {
      type = multi_noise
      biomes = [
        { biome = plains, parameters = { temperature = [-1, 0], humidity = 0, continentalness = 0, erosion = 0, weirdness = 0, depth = 0, offset = 0 } },
      ]
    }
  }
}

Dimension preset = {
  type = overworld
  generator = {
    type = noise
    settings = overworld
    biome_source = { type = multi_noise, preset = overworld }
  }
}

// biomes needs no preset, so leaving both out is an error.
Dimension neither = {
  type = overworld
  generator = { type = noise, settings = overworld, biome_source = { type = multi_noise } }
}
