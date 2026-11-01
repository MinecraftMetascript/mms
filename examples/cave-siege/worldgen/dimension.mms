// The cave world: a dimension type, the generator settings, the dimension,
// and a world preset that puts it in place of the overworld.
//
// Two ways to play it:
// - New world: pick "Cave Siege" as the world type (More World Options >
//   World Type). The overworld itself is caves, so Nether portals, beds,
//   strongholds and the End all keep working.
// - Existing world: /function siege:visit_caves takes you to siege:caves,
//   the same world as a separate dimension.

/**
 * No sky and a solid roof. Caves are dark but not black, and monsters only
 * spawn in full darkness, as in the vanilla overworld. The clock and
 * timelines are the overworld's, so a day still lasts 20 minutes: that is
 * the siege clock.
 */
DimensionType caves = {
  height = 256
  logical_height = 256
  min_y = -64
  coordinate_scale = 1
  ambient_light = 0.04
  has_skylight = false
  has_ceiling = true
  has_ender_dragon_fight = false
  has_fixed_time = false
  skybox = none
  infiniburn = "#minecraft:infiniburn_overworld"
  default_clock = overworld
  timelines = "#minecraft:in_overworld"
  monster_spawn_block_light_limit = 0
  monster_spawn_light_level = 0
  attributes = {
    "minecraft:visual/fog_color" = "#1b1a22"
    "minecraft:visual/fog_start_distance" = 24
    "minecraft:visual/fog_end_distance" = 120
    "minecraft:visual/ambient_light_color" = "#1a1714"
    "minecraft:visual/sky_light_factor" = 0
    "minecraft:audio/ambient_sounds" = {
      mood = {
        sound = "minecraft:ambient.cave"
        tick_delay = 6000
        block_search_extent = 8
        offset = 2
      }
    }
    "minecraft:audio/background_music" = {
      default = {
        sound = "minecraft:music.game"
        min_delay = 12000
        max_delay = 24000
      }
    }
    // Beds set your spawn, but there is no night to sleep through: the day
    // is the siege clock, so sleeping would skip the build phase.
    "minecraft:gameplay/bed_rule" = {
      can_set_spawn = always
      can_sleep = never
      error_message = {
        translate = "siege.bed.no_sleep"
        fallback = "There is no night underground. The siege clock keeps running."
      }
    }
    "minecraft:gameplay/nether_portal_spawns_piglin" = true
    "minecraft:gameplay/respawn_anchor_works" = false
  }
}

/** The generator: our terrain, our surface rules, no oceans and no aquifers. */
NoiseSettings caves = {
  sea_level = -64
  default_block = stone
  default_fluid = water
  disable_mob_generation = false
  legacy_random_source = false
  noise = {
    min_y = -64
    height = 256
  }
  material_rule = surface
  // Players arriving in a new world look for a spawn in Glowmeadow climate.
  spawn_target = [
    {
      "siege:climate_temperature" = [-0.2, 0.2]
      "siege:climate_humidity" = [-0.2, 0.2]
      "siege:climate_continents" = [-0.75, 1]
    },
  ]
  noise_router = {
    final_density = caves
    temperature = climate_temperature
    vegetation = climate_humidity
    continents = climate_continents
    erosion = climate_erosion
    depth = climate_depth
    ridges = 0
    chunk_surface_level = 0
  }
}

/** The cave world as its own dimension, for existing worlds. */
Dimension caves = {
  type = caves
  generator = {
    type = noise
    settings = caves
    biome_source = {
      type = multi_noise
      biomes = [
        // Rare regions where continents is low
        {
          biome = mycelium_deeps
          parameters = {
            temperature = [-1, 1]
            humidity = [-1, 1]
            continentalness = [-1, -0.75]
            erosion = [-1, 1]
            weirdness = 0
            depth = [0, 1.1]
            offset = 0
          }
        },
        // The deep layer, below y -35
        {
          biome = minecraft:deep_dark
          parameters = {
            temperature = [-1, 0.25]
            humidity = [-1, 1]
            continentalness = [-0.75, 1]
            erosion = [-1, 1]
            weirdness = 0
            depth = [1.1, 1.5]
            offset = 0
          }
        },
        {
          biome = ash_vents
          parameters = {
            temperature = [0.5, 1]
            humidity = [-1, -0.35]
            continentalness = [-0.75, 1]
            erosion = [-1, 1]
            weirdness = 0
            depth = [-1, 1.5]
            offset = 0
          }
        },
        // Below y 10
        {
          biome = darkroot_grove
          parameters = {
            temperature = [0.25, 1]
            humidity = [0.25, 1]
            continentalness = [-0.75, 1]
            erosion = [-1, 1]
            weirdness = 0
            depth = [0.75, 1.5]
            offset = 0
          }
        },
        {
          biome = rootwood_hollow
          parameters = {
            temperature = [0.25, 1]
            humidity = [-0.25, 1]
            continentalness = [-0.75, 1]
            erosion = [-1, 1]
            weirdness = 0
            depth = [-1, 0.6]
            offset = 0
          }
        },
        {
          biome = rootwood_hollow
          parameters = {
            temperature = [-1, -0.25]
            humidity = [-0.25, 0.25]
            continentalness = [-0.75, 1]
            erosion = [-1, 1]
            weirdness = 0
            depth = [-1, 1.1]
            offset = 0
          }
        },
        {
          biome = minecraft:dripstone_caves
          parameters = {
            temperature = [-1, -0.25]
            humidity = [-1, -0.25]
            continentalness = [-0.75, 1]
            erosion = [-1, 1]
            weirdness = 0
            depth = [-1, 1.1]
            offset = 0
          }
        },
        {
          biome = sunken_lake
          parameters = {
            temperature = [-1, -0.25]
            humidity = [0.25, 1]
            continentalness = [-0.75, 1]
            erosion = [-1, 1]
            weirdness = 0
            depth = [-1, 1.1]
            offset = 0
          }
        },
        {
          biome = glowmeadow
          parameters = {
            temperature = [-0.25, 0.25]
            humidity = [-1, 0.25]
            continentalness = [-0.75, 1]
            erosion = [-1, 1]
            weirdness = 0
            depth = [-1, 1.1]
            offset = 0
          }
        },
        {
          biome = minecraft:lush_caves
          parameters = {
            temperature = [-0.25, 0.25]
            humidity = [0.25, 1]
            continentalness = [-0.75, 1]
            erosion = [-1, 1]
            weirdness = 0
            depth = [-1, 1.1]
            offset = 0
          }
        },
      ]
    }
  }
}

/**
 * "Cave Siege" on the world type list. The overworld becomes caves; the
 * Nether and the End are vanilla.
 */
WorldPreset cave_siege = {
  dimensions = {
    "minecraft:overworld" = {
      type = caves
      generator = {
        type = noise
        settings = caves
        biome_source = {
          type = multi_noise
          biomes = [
            // Rare regions where continents is low
            {
              biome = mycelium_deeps
              parameters = {
                temperature = [-1, 1]
                humidity = [-1, 1]
                continentalness = [-1, -0.75]
                erosion = [-1, 1]
                weirdness = 0
                depth = [0, 1.1]
                offset = 0
              }
            },
            // The deep layer, below y -35
            {
              biome = minecraft:deep_dark
              parameters = {
                temperature = [-1, 0.25]
                humidity = [-1, 1]
                continentalness = [-0.75, 1]
                erosion = [-1, 1]
                weirdness = 0
                depth = [1.1, 1.5]
                offset = 0
              }
            },
            {
              biome = ash_vents
              parameters = {
                temperature = [0.5, 1]
                humidity = [-1, -0.35]
                continentalness = [-0.75, 1]
                erosion = [-1, 1]
                weirdness = 0
                depth = [-1, 1.5]
                offset = 0
              }
            },
            // Below y 10
            {
              biome = darkroot_grove
              parameters = {
                temperature = [0.25, 1]
                humidity = [0.25, 1]
                continentalness = [-0.75, 1]
                erosion = [-1, 1]
                weirdness = 0
                depth = [0.75, 1.5]
                offset = 0
              }
            },
            {
              biome = rootwood_hollow
              parameters = {
                temperature = [0.25, 1]
                humidity = [-0.25, 1]
                continentalness = [-0.75, 1]
                erosion = [-1, 1]
                weirdness = 0
                depth = [-1, 0.6]
                offset = 0
              }
            },
            {
              biome = rootwood_hollow
              parameters = {
                temperature = [-1, -0.25]
                humidity = [-0.25, 0.25]
                continentalness = [-0.75, 1]
                erosion = [-1, 1]
                weirdness = 0
                depth = [-1, 1.1]
                offset = 0
              }
            },
            {
              biome = minecraft:dripstone_caves
              parameters = {
                temperature = [-1, -0.25]
                humidity = [-1, -0.25]
                continentalness = [-0.75, 1]
                erosion = [-1, 1]
                weirdness = 0
                depth = [-1, 1.1]
                offset = 0
              }
            },
            {
              biome = sunken_lake
              parameters = {
                temperature = [-1, -0.25]
                humidity = [0.25, 1]
                continentalness = [-0.75, 1]
                erosion = [-1, 1]
                weirdness = 0
                depth = [-1, 1.1]
                offset = 0
              }
            },
            {
              biome = glowmeadow
              parameters = {
                temperature = [-0.25, 0.25]
                humidity = [-1, 0.25]
                continentalness = [-0.75, 1]
                erosion = [-1, 1]
                weirdness = 0
                depth = [-1, 1.1]
                offset = 0
              }
            },
            {
              biome = minecraft:lush_caves
              parameters = {
                temperature = [-0.25, 0.25]
                humidity = [0.25, 1]
                continentalness = [-0.75, 1]
                erosion = [-1, 1]
                weirdness = 0
                depth = [-1, 1.1]
                offset = 0
              }
            },
          ]
        }
      }
    }
    "minecraft:the_nether" = {
      type = minecraft:the_nether
      generator = {
        type = noise
        settings = minecraft:nether
        biome_source = {
          type = multi_noise
          preset = minecraft:nether
        }
      }
    }
    "minecraft:the_end" = {
      type = minecraft:the_end
      generator = {
        type = noise
        settings = minecraft:end
        biome_source = { type = the_end }
      }
    }
  }
}

/** Lists Cave Siege on the world type button. */
WorldPresetTag minecraft:normal = [cave_siege]
