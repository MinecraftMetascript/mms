Namespace test_dimensions {

    // Basic dimension types for different environments

    Dimension {
        // Overworld-like dimension
        OverworldType = DimensionType(
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

        // Nether-like dimension
        NetherType = DimensionType(
            Ultrawarm(true),
            Natural(false),
            Skylight(false),
            Ceiling(true),
            PiglinSafe(true),
            BedsWork(false),
            AnchorsWork(false),
            HasRaids(false),
            CoordinateScale(8),
            AmbientLight(0),
            MonsterLightLevel(7),
            MonsterLightLimit(15),
            LogicalHeight(128),
            CloudHeight(128),
            MinY(0),
            Height(128),
            Infiniburn("#minecraft:infiniburn_nether"),
            Effects(minecraft:the_nether)
        )

        // End-like dimension
        EndType = DimensionType(
            Ultrawarm(false),
            Natural(false),
            Skylight(false),
            Ceiling(false),
            PiglinSafe(true),
            BedsWork(false),
            AnchorsWork(false),
            HasRaids(false),
            CoordinateScale(1),
            AmbientLight(0),
            MonsterLightLevel(0),
            MonsterLightLimit(7),
            LogicalHeight(256),
            CloudHeight(256),
            MinY(0),
            Height(256),
            Infiniburn("#minecraft:infiniburn_end"),
            Effects(minecraft:the_end)
        )

        // Custom sky dimension
        SkyDimension = DimensionType(
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
            FixedTime(6000),  // Always daytime
            MonsterLightLevel(0),
            MonsterLightLimit(7),
            LogicalHeight(256),
            CloudHeight(128),
            MinY(64),
            Height(192),
            Infiniburn("#minecraft:infiniburn_overworld"),
            Effects(minecraft:overworld)
        )
    }
    }