Namespace test_dimensions

    DimensionType {
        NormalType = DimensionType(
            Ultrawarm(false),
            Natural(true),
            Skylight(true),
            Ceiling(false),
            PiglinSafe(false),
            BedsWork(true),
            AnchorsWork(true),
            HasRaids(true),
            CoordinateScale(1),
            AmbientLight(0),
            MonsterLightLevel(0),
            MonsterLightLimit(7),
            LogicalHeight(256),
            CloudHeight(192),
            MinY(-64),
            Height(320),
            Infiniburn("#minecraft:infiniburn_overworld"),
            Effects(minecraft:overworld)
        )
    }

    Dimension {
        // Simple fixed biome dimension
        PlainsOnly = NormalType.Generator("noise").Settings(
            NoiseSettings(
                SeaLevel(63)
                DisableMobGen()
                DefaultBlock("minecraft:stone")
                DefaultFluid("minecraft:water")
                MinY(-64)
                Height(320)
                Size(2, 1)
                BiomeSource("fixed")
                    Biome("minecraft:plains")
            )
        )

        // Checkerboard biome dimension
        CheckerboardBiomes = NormalType.Generator("noise").Settings(
            NoiseSettings(
                SeaLevel(63)
                DefaultBlock("minecraft:stone")
                DefaultFluid("minecraft:water")
                MinY(-64)
                Height(320)
                Size(2, 1)
                BiomeSource("checkerboard")
                    Biomes("minecraft:plains", "minecraft:forest", "minecraft:desert")
                    Scale(3)
            )
        )

        // Multi-noise with overworld preset
        OverworldPreset = NormalType.Generator("noise").Settings(
            NoiseSettings(
                SeaLevel(63)
                DefaultBlock("minecraft:stone")
                DefaultFluid("minecraft:water")
                MinY(-64)
                Height(320)
                Size(2, 1)
                BiomeSource("multi_noise")
                    Preset("overworld")
            )
        )

        // Multi-noise with nether preset
        NetherPreset = NormalType.Generator("noise").Settings(
            NoiseSettings(
                SeaLevel(32)
                DisableMobGen()
                DefaultBlock("minecraft:netherrack")
                DefaultFluid("minecraft:lava")
                MinY(0)
                Height(128)
                Size(2, 1)
                BiomeSource("multi_noise")
                    Preset("nether")
            )
        )

        // The End dimension
        TheEnd = NormalType.Generator("noise").Settings(
            NoiseSettings(
                SeaLevel(0)
                DisableMobGen()
                DefaultBlock("minecraft:end_stone")
                DefaultFluid("minecraft:air")
                MinY(0)
                Height(256)
                Size(2, 1)
                BiomeSource("the_end")
            )
        )

        // Custom noise settings with specific noise router
        CustomNoise = NormalType.Generator("noise").Settings(
            NoiseSettings(
                SeaLevel(63)
                DisableMobGen()
                DefaultBlock("minecraft:stone")
                DefaultFluid("minecraft:water")
                MinY(-64)
                Height(320)
                Size(2, 1)
                NoiseRouter(test_files:MyNoise)
                BiomeSource("fixed")
                    Biome("minecraft:plains")
            )
        )
    }