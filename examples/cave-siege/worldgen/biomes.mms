// The biomes. Each custom biome supplies something a surface playthrough
// gets from the sky; Lush Caves, Dripstone Caves and Deep Dark are vanilla
// and kept as they are. Where they go is in climate.mms and the biome
// source in dimension.mms.
//
// Features run in 11 steps; each biome lists one list per step:
//  1 raw_generation  2 lakes  3 local_modifications  4 underground_structures
//  5 surface_structures  6 strongholds  7 underground_ores
//  8 underground_decoration  9 fluid_springs  10 vegetal_decoration
//  11 top_layer_modification
//
// Hostile mobs spawn in small groups (1 or 2) so waves stay the main threat.

/**
 * Oak, birch and spruce woods. Glowstone in the roof and glow berries
 * light the floor, so saplings regrow.
 */
Biome rootwood_hollow = {
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
          particle = { type = spore_blossom_air }
          probability = 0.004
        },
      ]
    }
    "minecraft:audio/background_music" = {
      default = {
        sound = "minecraft:music.overworld.forest"
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
    [roof_glowstone, minecraft:glow_lichen],
    #water_springs,
    [rootwood_trees, floor_grass, floor_flowers, minecraft:cave_vines, glow_caps],
    [],
  ]
}

/**
 * The farm biome: grass floors lit to 10 or more by glowstone and ochre
 * froglights above and glow caps below, with cows, sheep, pigs and
 * chickens, wheat seeds from the grass, flowers and pumpkins.
 */
Biome glowmeadow = {
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
          particle = { type = white_ash }
          probability = 0.002
        },
      ]
    }
    "minecraft:audio/background_music" = {
      default = {
        sound = "minecraft:music.overworld.meadow"
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

/** Mushrooms and mooshrooms: food without light. Hostile mobs don't spawn here, as in vanilla. */
Biome mycelium_deeps = {
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
          particle = { type = mycelium }
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

/**
 * Deep and rare: dark oak, jungle and cherry trees, cocoa on the jungle
 * logs, bamboo and melons, lit by verdant froglight crystals.
 */
Biome darkroot_grove = {
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
        sound = "minecraft:music.overworld.jungle"
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
    [verdant_crystals, roof_glowstone_sparse, minecraft:ore_emerald],
    #water_springs,
    [darkroot_trees, floor_bamboo, floor_melons, floor_grass, glow_caps],
    [],
  ]
}

/**
 * Lakes with sandy shores: sand and sandstone floors, sugar cane, kelp,
 * fish and squid. Clay comes from the shared ore list.
 */
Biome sunken_lake = {
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

/**
 * Lava pools on basalt and blackstone floors, for early obsidian. Hot and
 * dangerous, so Rifts should avoid it.
 */
Biome ash_vents = {
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
          particle = { type = ash }
          probability = 0.03
        },
      ]
    }
    "minecraft:visual/default_dripstone_particle" = { type = dripping_dripstone_lava }
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
    [minecraft:ore_magma, minecraft:ore_emerald],
    #water_springs,
    [],
    [],
  ]
}

// Vanilla structures that should still generate in our biomes. These add
// to vanilla's tags (no `replace`), so they don't override them.

/** Strongholds, so the End is reachable. */
BiomeTag minecraft:has_structure/stronghold = [
  rootwood_hollow,
  glowmeadow,
  mycelium_deeps,
  darkroot_grove,
  sunken_lake,
  ash_vents,
]

BiomeTag minecraft:has_structure/mineshaft = [
  rootwood_hollow,
  glowmeadow,
  darkroot_grove,
  sunken_lake,
  ash_vents,
]

BiomeTag minecraft:has_structure/trial_chambers = [
  rootwood_hollow,
  glowmeadow,
  darkroot_grove,
  sunken_lake,
]

/** Counts as the overworld for vanilla checks such as goat horns and wandering traders. */
BiomeTag minecraft:is_overworld = [
  rootwood_hollow,
  glowmeadow,
  mycelium_deeps,
  darkroot_grove,
  sunken_lake,
  ash_vents,
]
