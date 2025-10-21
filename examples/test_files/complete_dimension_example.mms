Namespace test_dimensions {

    // Complete example showing all dimension features

    Dimension {
        // Advanced custom dimension type
        AdvancedType = DimensionType().
            // Ultrawarm(false).
            Natural().
            Skylight().
            // Ceiling(false).
            // PiglinSafe(false).
            BedsWork().
            AnchorsWork().
            HasRaids().
            CoordinateScale(1).
            AmbientLight(0).
            FixedTime(6000).  // Always daytime
            MonsterLightLevel(0).
            MonsterLightLimit(7).
            LogicalHeight(256).
            CloudHeight(128).
            MinY(-32).
            Height(224).
            Infiniburn("#minecraft:infiniburn_overworld").
            Effects(minecraft:overworld)
        
            // Debug dimension for testing
            DebugDimension = AdvancedType.Generator("debug")

            // Complex flat dimension with multiple layers
            ComplexFlat = AdvancedType.Generator("flat").Settings(
                FlatSettings(
                    Layers([
                        FlatLayer("minecraft:bedrock", 1),
                        FlatLayer("minecraft:deepslate", 5),
                        FlatLayer("minecraft:tuff", 3),
                        FlatLayer("minecraft:stone", 15),
                        FlatLayer("minecraft:calcite", 2),
                        FlatLayer("minecraft:dirt", 4),
                        FlatLayer("minecraft:grass_block", 1)
                    ]),
                    Structures("#minecraft:village_plains", "#minecraft:pillager_outpost", "#minecraft:ruined_portal"),
                    Lakes("#minecraft:lakes"),
                    Features("#minecraft:trees_plains", "#minecraft:flowers_plains", "#minecraft:ore_coal_upper"),
                    Biome("minecraft:plains")
                )
            )

            // Advanced noise dimension with custom settings
            AdvancedNoise = test_dimensions:AdvancedType.Generator("noise").Settings(
                NoiseSettings(
                    SeaLevel(80),  // Higher sea level for more islands
                    DisableMobGen(),  // Peaceful for building
                    EnableOreVeins(),  // More interesting ore generation
                    Aquifers(),  // Enable aquifer generation
                    DefaultBlock("minecraft:stone"),
                    DefaultFluid("minecraft:water"),
                    MinY(-32),
                    Height(224),
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
                                    .Offset(0.0)
                            ),
                            MultiNoiseBiome("minecraft:forest").Parameters(
                                MultiNoiseParameters()
                                    .Temperature(0.7)
                                    .Humidity(0.8)
                                    .Continentalness(0.2)
                                    .Erosion(0.4)
                                    .Weirdness(0.1)
                                    .Depth(0.1)
                                    .Offset(0.1)
                            ),
                            MultiNoiseBiome("minecraft:mountains").Parameters(
                                MultiNoiseParameters()
                                    .Temperature(0.2)
                                    .Humidity(0.3)
                                    .Continentalness(0.8)
                                    .Erosion(0.9)
                                    .Weirdness(0.3)
                                    .Depth(0.4)
                                    .Offset(0.2)
                            ),
                            MultiNoiseBiome("minecraft:ocean").Parameters(
                                MultiNoiseParameters()
                                    .Temperature(0.5)
                                    .Humidity(0.5)
                                    .Continentalness(-0.8)
                                    .Erosion(0.3)
                                    .Weirdness(0.1)
                                    .Depth(-0.9)
                                    .Offset(-0.1)
                            )
                        ])
                )
            )

            // Sky islands dimension using checkerboard
            SkyIslands = AdvancedType.Generator("noise").Settings(
                NoiseSettings(
                    SeaLevel(100),  // High sea level creates floating islands
                    DefaultBlock("minecraft:stone"),
                    DefaultFluid("minecraft:air"),  // No water
                    MinY(64),
                    Height(128),
                    Size(2, 1),
                    BiomeSource("checkerboard")
                        .Biomes("minecraft:plains", "minecraft:forest")
                        .Scale(8)  // Large islands
                )
            )
    }
}