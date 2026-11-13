// The biomes. Each custom biome supplies something a surface playthrough
// gets from the sky; Lush Caves, Dripstone Caves and Deep Dark are vanilla
// and kept as they are. Each CustomBiome says where it goes (`climate`,
// against the noises in climate.mms), what its ground is made of
// (`surface_rule`, from surface.mms) and which vanilla tags it joins, so
// that strongholds, mineshafts and trial chambers still generate in it.
// dimension.mms puts their noise entries in the biome source.
//
// Features run in 11 steps, one list per step; the editor names each
// step beside its list.
//
// Hostile mobs spawn in small groups (1 or 2) so waves stay the main threat.

/**
 * Oak, birch and spruce woods. Glowstone in the roof and glow berries
 * light the floor, so saplings regrow.
 */
CustomBiome rootwood_hollow {
  climate = [
    {
      temperature = [0.25, 1]
      humidity = [-0.25, 1]
      continentalness = [-0.75, 1]
      erosion = [-1, 1]
      weirdness = 0
      depth = [-1, 0.6]
      offset = 0
    },
    {
      temperature = [-1, -0.25]
      humidity = [-0.25, 0.25]
      continentalness = [-0.75, 1]
      erosion = [-1, 1]
      weirdness = 0
      depth = [-1, 1.1]
      offset = 0
    },
  ]
  surface_rule = meadow_ground
  tags = [
    #is_overworld,
    #has_structure/stronghold,
    #has_structure/mineshaft,
    #has_structure/trial_chambers,
  ]
  biome = {
    temperature = 0.6
    downfall = 0.7
    has_precipitation = false
    effects = {
      water_color = "#3f76e4"
      grass_color = "#6a9c3a"
      foliage_color = "#5e9a2c"
    }
    attributes = {
      "minecraft:visual/fog_color" = "#24301c"
      "minecraft:visual/ambient_particles" = {
        modifier = append
        argument = [
          {
            particle = SporeBlossomAir()
            probability = 0.004
          },
        ]
      }
      "minecraft:audio/background_music" = {
        default = {
          sound = minecraft:music.overworld.forest
          min_delay = 12000
          max_delay = 24000
        }
      }
      "minecraft:gameplay/natural_mob_spawns" = {
        modifier = overlay
        argument = {
          spawn_costs = {}
          spawns_by_category = {
            creature = [
              {
                type = rabbit
                count = 2
                weight = 4
              },
              {
                type = fox
                count = 2
                weight = 4
              },
            ]
            ambient = [
              {
                type = bat
                count = 4
                weight = 10
              },
            ]
            monster = [
              {
                type = zombie
                count = 1
                weight = 100
              },
              {
                type = skeleton
                count = 1
                weight = 100
              },
              {
                type = spider
                count = 1
                weight = 100
              },
              {
                type = creeper
                count = 1
                weight = 80
              },
            ]
          }
        }
      }
    }
    carvers = []
    features = [
      [],
      [],
      #underground_extras,
      #dungeons,
      [],
      [],
      #ores,
      [roof_glowstone, glow_lichen],
      #water_springs,
      [rootwood_trees, floor_grass, floor_flowers, cave_vines, glow_caps],
      [],
    ]
  }
}

/**
 * The farm biome: grass floors lit to 10 or more by glowstone and ochre
 * froglights above and glow caps below, with cows, sheep, pigs and
 * chickens, wheat seeds from the grass, flowers and pumpkins.
 */
CustomBiome glowmeadow {
  climate = {
    temperature = [-0.25, 0.25]
    humidity = [-1, 0.25]
    continentalness = [-0.75, 1]
    erosion = [-1, 1]
    weirdness = 0
    depth = [-1, 1.1]
    offset = 0
  }
  surface_rule = meadow_ground
  tags = [
    #is_overworld,
    #has_structure/stronghold,
    #has_structure/mineshaft,
    #has_structure/trial_chambers,
  ]
  biome = {
    temperature = 0.8
    downfall = 0.4
    has_precipitation = false
    effects = {
      water_color = "#3f76e4"
      grass_color = "#7cbd4a"
    }
    attributes = {
      "minecraft:visual/fog_color" = "#3a3320"
      "minecraft:visual/ambient_particles" = {
        modifier = append
        argument = [
          {
            particle = WhiteAsh()
            probability = 0.002
          },
        ]
      }
      "minecraft:audio/background_music" = {
        default = {
          sound = minecraft:music.overworld.meadow
          min_delay = 12000
          max_delay = 24000
        }
      }
      // Animals at chunk generation, more often than the overworld's 0.1.
      "minecraft:gameplay/creature_world_gen_spawn_probability" = 0.4
      "minecraft:gameplay/natural_mob_spawns" = {
        modifier = overlay
        argument = {
          spawn_costs = {}
          spawns_by_category = {
            creature = [
              {
                type = cow
                count = 4
                weight = 8
              },
              {
                type = sheep
                count = 4
                weight = 12
              },
              {
                type = pig
                count = 4
                weight = 10
              },
              {
                type = chicken
                count = 4
                weight = 10
              },
            ]
            ambient = [
              {
                type = bat
                count = 4
                weight = 10
              },
            ]
            monster = [
              {
                type = zombie
                count = 1
                weight = 100
              },
              {
                type = skeleton
                count = 1
                weight = 100
              },
              {
                type = creeper
                count = 1
                weight = 60
              },
            ]
          }
        }
      }
    }
    carvers = []
    features = [
      [],
      [],
      #underground_extras,
      #dungeons,
      [],
      [],
      #ores,
      [roof_glowstone, ochre_crystals, starter_lights],
      #water_springs,
      [floor_grass, floor_tall_grass, floor_flowers, floor_pumpkins, glow_caps],
      [],
    ]
  }
}

/** Mushrooms and mooshrooms: food without light. Hostile mobs don't spawn here, as in vanilla. */
CustomBiome mycelium_deeps {
  climate = {
    temperature = [-1, 1]
    humidity = [-1, 1]
    continentalness = [-1, -0.75]
    erosion = [-1, 1]
    weirdness = 0
    depth = [0, 1.1]
    offset = 0
  }
  surface_rule = mycelium_ground
  tags = [#minecraft:is_overworld, #minecraft:has_structure/stronghold]
  biome = {
    temperature = 0.9
    downfall = 1
    has_precipitation = false
    effects = { water_color = "#3f76e4" }
    attributes = {
      "minecraft:visual/fog_color" = "#2a2030"
      "minecraft:visual/ambient_particles" = {
        modifier = append
        argument = [
          {
            particle = Mycelium()
            probability = 0.02
          },
        ]
      }
      "minecraft:gameplay/natural_mob_spawns" = {
        modifier = overlay
        argument = {
          spawn_costs = {}
          spawns_by_category = {
            creature = [
              {
                type = mooshroom
                count = 4
                weight = 8
              },
            ]
            ambient = [
              {
                type = bat
                count = 4
                weight = 10
              },
            ]
          }
        }
      }
    }
    carvers = []
    features = [
      [],
      [],
      #underground_extras,
      [],
      [],
      [],
      #ores,
      [roof_glowstone_sparse],
      #water_springs,
      [huge_mushrooms, floor_mushrooms, floor_red_mushrooms],
      [],
    ]
  }
}

/**
 * Deep and rare: dark oak, jungle and cherry trees, cocoa on the jungle
 * logs, bamboo and melons, lit by verdant froglight crystals.
 */
CustomBiome darkroot_grove {
  climate = {
    temperature = [0.25, 1]
    humidity = [0.25, 1]
    continentalness = [-0.75, 1]
    erosion = [-1, 1]
    weirdness = 0
    depth = [0.75, 1.5]
    offset = 0
  }
  surface_rule = darkroot_ground
  tags = [
    #is_overworld,
    #has_structure/stronghold,
    #has_structure/mineshaft,
    #has_structure/trial_chambers,
  ]
  biome = {
    temperature = 0.95
    downfall = 0.9
    has_precipitation = false
    effects = {
      water_color = "#3f76e4"
      grass_color = "#4f7a2a"
      foliage_color = "#3f7a1c"
    }
    attributes = {
      "minecraft:visual/fog_color" = "#15241a"
      "minecraft:audio/background_music" = {
        default = {
          sound = minecraft:music.overworld.jungle
          min_delay = 12000
          max_delay = 24000
        }
      }
      "minecraft:gameplay/natural_mob_spawns" = {
        modifier = overlay
        argument = {
          spawn_costs = {}
          spawns_by_category = {
            creature = [
              {
                type = parrot
                count = 1
                weight = 4
              },
              {
                type = ocelot
                count = 1
                weight = 2
              },
              {
                type = chicken
                count = 2
                weight = 6
              },
            ]
            ambient = [
              {
                type = bat
                count = 4
                weight = 10
              },
            ]
            monster = [
              {
                type = zombie
                count = 1
                weight = 100
              },
              {
                type = skeleton
                count = 1
                weight = 100
              },
              {
                type = spider
                count = 2
                weight = 100
              },
              {
                type = creeper
                count = 1
                weight = 80
              },
              {
                type = witch
                count = 1
                weight = 5
              },
            ]
          }
        }
      }
    }
    carvers = []
    features = [
      [],
      [],
      #underground_extras,
      #dungeons,
      [],
      [],
      #ores,
      [verdant_crystals, roof_glowstone_sparse, ore_emerald],
      #water_springs,
      [darkroot_trees, floor_bamboo, floor_melons, floor_grass, glow_caps],
      [],
    ]
  }
}

/**
 * Lakes with sandy shores: sand and sandstone floors, sugar cane, kelp,
 * fish and squid. Clay comes from the shared ore list.
 */
CustomBiome sunken_lake {
  climate = {
    temperature = [-1, -0.25]
    humidity = [0.25, 1]
    continentalness = [-0.75, 1]
    erosion = [-1, 1]
    weirdness = 0
    depth = [-1, 1.1]
    offset = 0
  }
  surface_rule = lake_ground
  tags = [
    #is_overworld,
    #has_structure/stronghold,
    #has_structure/mineshaft,
    #has_structure/trial_chambers,
  ]
  biome = {
    temperature = 0.7
    downfall = 0.8
    has_precipitation = false
    effects = {
      water_color = "#2f6fbf"
      grass_color = "#6a9c3a"
    }
    attributes = {
      "minecraft:visual/fog_color" = "#16222e"
      "minecraft:visual/water_fog_color" = "#0e2a40"
      "minecraft:gameplay/natural_mob_spawns" = {
        modifier = overlay
        argument = {
          spawn_costs = {}
          spawns_by_category = {
            water_creature = [
              {
                type = squid
                count = 4
                weight = 4
              },
            ]
            water_ambient = [
              {
                type = cod
                count = 4
                weight = 10
              },
              {
                type = salmon
                count = 4
                weight = 5
              },
            ]
            underground_water_creature = [
              {
                type = glow_squid
                count = 4
                weight = 10
              },
            ]
            ambient = [
              {
                type = bat
                count = 4
                weight = 10
              },
            ]
            monster = [
              {
                type = drowned
                count = 1
                weight = 100
              },
              {
                type = zombie
                count = 1
                weight = 50
              },
              {
                type = skeleton
                count = 1
                weight = 50
              },
            ]
          }
        }
      }
    }
    carvers = []
    features = [
      [],
      [water_lakes],
      #underground_extras,
      [],
      [],
      [],
      #ores,
      [roof_glowstone_sparse],
      #water_springs,
      [lake_kelp, shore_sugar_cane, floor_grass],
      [],
    ]
  }
}

/**
 * Lava pools on basalt and blackstone floors, for early obsidian. Hot and
 * dangerous, so Rifts should avoid it.
 */
CustomBiome ash_vents {
  climate = {
    temperature = [0.5, 1]
    humidity = [-1, -0.35]
    continentalness = [-0.75, 1]
    erosion = [-1, 1]
    weirdness = 0
    depth = [-1, 1.5]
    offset = 0
  }
  surface_rule = ash_ground
  tags = [
    #is_overworld,
    #has_structure/stronghold,
    #has_structure/mineshaft,
  ]
  biome = {
    temperature = 2
    downfall = 0
    has_precipitation = false
    effects = { water_color = "#3f76e4" }
    attributes = {
      "minecraft:visual/fog_color" = "#3a1a10"
      "minecraft:visual/ambient_particles" = {
        modifier = append
        argument = [
          {
            particle = Ash()
            probability = 0.03
          },
        ]
      }
      "minecraft:visual/default_dripstone_particle" = DrippingDripstoneLava()
      "minecraft:gameplay/natural_mob_spawns" = {
        modifier = overlay
        argument = {
          spawn_costs = {}
          spawns_by_category = {
            monster = [
              {
                type = magma_cube
                count = 1
                weight = 50
              },
              {
                type = skeleton
                count = 1
                weight = 100
              },
              {
                type = zombie
                count = 1
                weight = 50
              },
            ]
          }
        }
      }
    }
    carvers = []
    features = [
      [],
      [lava_lakes],
      #underground_extras,
      #dungeons,
      [],
      [],
      #ores,
      [ore_magma, ore_emerald],
      #water_springs,
      [],
      [],
    ]
  }
}
