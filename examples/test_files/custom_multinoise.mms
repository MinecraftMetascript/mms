Namespace test_dimensions {

    DimensionType {
        CustomType = DimensionType(
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
        // Custom multi-noise biome configuration
        CustomBiomes = test_dimensions:CustomType.Generator("noise").Settings(
            NoiseSettings(
                SeaLevel(63),
                DefaultBlock("minecraft:stone"),
                DefaultFluid("minecraft:water"),
                MinY(-64),
                Height(320),
                Size(2, 1),
                BiomeSource("multi_noise")
                    .Biomes([
                        MultiNoiseBiome("minecraft:plains").Parameters(
                            MultiNoiseParameters()
                                .Temperature(0.8)
                                .Humidity(0.4)
                                .Continentalness(0.3)
                                .Erosion(0.5)
                                .Weirdness(0.2)
                                .Depth(0.1)
                        ),
                        MultiNoiseBiome("minecraft:desert").Parameters(
                            MultiNoiseParameters()
                                .Temperature(2.0)
                                .Humidity(0.0)
                                .Continentalness(0.5)
                                .Erosion(0.8)
                                .Weirdness(0.1)
                                .Depth(0.1)
                        ),
                        MultiNoiseBiome("minecraft:savanna").Parameters(
                            MultiNoiseParameters()
                                .Temperature(1.2)
                                .Humidity(0.0)
                                .Continentalness(0.4)
                                .Erosion(0.3)
                                .Weirdness(0.3)
                                .Depth(0.1)
                        ),
                        MultiNoiseBiome("minecraft:forest").Parameters(
                            MultiNoiseParameters()
                                .Temperature(0.7)
                                .Humidity(0.8)
                                .Continentalness(0.2)
                                .Erosion(0.4)
                                .Weirdness(0.1)
                                .Depth(0.1)
                        ),
                        MultiNoiseBiome("minecraft:taiga").Parameters(
                            MultiNoiseParameters()
                                .Temperature(0.25)
                                .Humidity(0.8)
                                .Continentalness(0.3)
                                .Erosion(0.6)
                                .Weirdness(0.2)
                                .Depth(0.1)
                        ),
                        MultiNoiseBiome("minecraft:snowy_tundra").Parameters(
                            MultiNoiseParameters()
                                .Temperature(0.0)
                                .Humidity(0.5)
                                .Continentalness(0.4)
                                .Erosion(0.7)
                                .Weirdness(0.1)
                                .Depth(0.1)
                        ),
                        MultiNoiseBiome("minecraft:jungle").Parameters(
                            MultiNoiseParameters()
                                .Temperature(0.95)
                                .Humidity(0.9)
                                .Continentalness(0.1)
                                .Erosion(0.2)
                                .Weirdness(0.4)
                                .Depth(0.1)
                        ),
                        MultiNoiseBiome("minecraft:ocean").Parameters(
                            MultiNoiseParameters()
                                .Temperature(0.5)
                                .Humidity(0.5)
                                .Continentalness(-0.5)
                                .Erosion(0.3)
                                .Weirdness(0.1)
                                .Depth(-0.8)
                        )
                    ])
            )
        )

        // Checkerboard with biome tags
        BiomeTagCheckerboard = test_dimensions:CustomType.Generator("noise").Settings(
            NoiseSettings(
                SeaLevel(63),
                DefaultBlock("minecraft:stone"),
                DefaultFluid("minecraft:water"),
                MinY(-64),
                Height(320),
                Size(2, 1),
                BiomeSource("checkerboard")
                    .Biomes("#minecraft:is_forest", "#minecraft:is_mountain", "#minecraft:is_ocean")
                    .Scale(5)
            )
        )

        // Mixed biome sources using tags
        MixedBiomes = test_dimensions:CustomType.Generator("noise").Settings(
            NoiseSettings(
                SeaLevel(63),
                DefaultBlock("minecraft:stone"),
                DefaultFluid("minecraft:water"),
                MinY(-64),
                Height(320),
                Size(2, 1),
                BiomeSource("checkerboard")
                    .Biomes("minecraft:plains", "#minecraft:is_mountain", "minecraft:ocean")
                    .Scale(2)
            )
        )
    }
}