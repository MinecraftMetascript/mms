Namespace test_dimensions

    DimensionType {
        FlatType = DimensionType(
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
            MinY(0),
            Height(256),
            Infiniburn("#minecraft:infiniburn_overworld"),
            Effects(minecraft:overworld)
        )
    }

    Dimension {
        // Classic superflat world
        ClassicFlat = FlatType.Generator("flat").Settings(
            FlatSettings(
                Layers([
                    FlatLayer("minecraft:bedrock", 1)
                    FlatLayer("minecraft:stone", 3)
                    FlatLayer("minecraft:dirt", 1)
                    FlatLayer("minecraft:grass_block", 1)
                ])
                Structures("#minecraft:village_plains", "#minecraft:pillager_outpost")
                Lakes("#minecraft:lakes")
                Features("#minecraft:trees_plains", "#minecraft:flowers_plains")
                Biome("minecraft:plains")
            )
        )

        // Desert superflat
        DesertFlat = FlatType.Generator("flat").Settings(
            FlatSettings(
                Layers([
                    FlatLayer("minecraft:bedrock", 1)
                    FlatLayer("minecraft:stone", 3)
                    FlatLayer("minecraft:sandstone", 2)
                    FlatLayer("minecraft:sand", 1)
                ])
                Structures("#minecraft:village_desert")
                Biome("minecraft:desert")
            )
        )

        // Ocean superflat (for water worlds)
        OceanFlat = FlatType.Generator("flat").Settings(
            FlatSettings(
                Layers([
                    FlatLayer("minecraft:bedrock", 1)
                    FlatLayer("minecraft:stone", 3)
                    FlatLayer("minecraft:dirt", 1)
                    FlatLayer("minecraft:water", 50)
                ])
                Biome("minecraft:ocean")
            )
        )

        // Snow superflat
        SnowFlat = FlatType.Generator("flat").Settings(
            FlatSettings(
                Layers([
                    FlatLayer("minecraft:bedrock", 1)
                    FlatLayer("minecraft:stone", 3)
                    FlatLayer("minecraft:snow_block", 1)
                ])
                Structures("#minecraft:village_snowy")
                Biome("minecraft:snowy_tundra")
            )
        )

        // Complex layered world
        LayeredFlat = FlatType.Generator("flat").Settings(
            FlatSettings(
                Layers([
                    FlatLayer("minecraft:bedrock", 1)
                    FlatLayer("minecraft:deepslate", 10)
                    FlatLayer("minecraft:stone", 20)
                    FlatLayer("minecraft:dirt", 5)
                    FlatLayer("minecraft:grass_block", 1)
                ])
                Structures("#minecraft:village_plains", "#minecraft:ruined_portal")
                Lakes("#minecraft:lakes")
                Features("#minecraft:trees_plains", "#minecraft:ore_dirt")
                Biome("minecraft:plains")
            )
        )
    }