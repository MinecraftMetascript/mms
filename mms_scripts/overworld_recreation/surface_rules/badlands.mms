namespace MySpace;


surface {
    // --- BADLANDS SURFACE RULES ---
    // Main entry: Bandlands
    // This file defines the surface rules for the Badlands biome, including terracotta bands, sand, and stone layers.

    // --- BIOME CONDITION ---
    condition InBadlands
        biome [ minecraft:badlands minecraft:eroded_badlands minecraft:wooded_badlands ]

    // --- MAIN RULE ENTRY ---
    rule Bandlands
        if (InBadlands)
            sequence [
                if (y_above absolute 63 0 sub)
                    sequence [
                        SkyTerracotta
                        TerracottaBands
                        SurfaceSands
                        NonHoleOrangeTerracotta
                        WhiteTerracotta
                        StoneAndGravel
                    ]
                if (y_above absolute 63 -1 add)
                    sequence [
                        OrangeTerracottaEdge
                        bandlands
                    ]
                if (stone_depth floor 0 add 0)
                    WhiteTerracotta
            ]

    // --- TERRACOTTA LAYERS ---
    rule SkyTerracotta
        if (y_above absolute 256 0 sub)
            block minecraft:orange_terracotta

    rule TerracottaBands
        if (stone_depth floor 0 sub 0)
            sequence [
                if (y_above absolute 74 1 add)
                    sequence [
                        if (noise minecraft:surface [-0.909, -0.5454]) block minecraft:terracotta
                        if (noise minecraft:surface [-0.1818, 0.1818]) block minecraft:terracotta
                        if (noise minecraft:surface [0.5454, 0.909]) block minecraft:terracotta
                        bandlands
                    ]
            ]

    rule NonHoleOrangeTerracotta
        if !(hole)
            block minecraft:orange_terracotta

    // --- SANDY LAYERS ---
    rule SurfaceSands
        if (above_water -1 0 sub)
            sequence [
                if (stone_depth ceiling 0 sub 0) block minecraft:red_sandstone
                block minecraft:red_sand
            ]

    rule WhiteTerracotta
        if (above_water -6 -1 add)
            block minecraft:white_terracotta

    // --- ORANGE TERRACOTTA EDGE ---
    rule OrangeTerracottaEdge
        if (and (
            y_above absolute 63 0 sub
            !y_above absolute 74 1 add
        ))
            block minecraft:orange_terracotta

    // --- STONE LAYERS ---
    rule StoneAndGravel
        sequence [
            if (stone_depth ceiling 0 sub 0) block minecraft:stone
            block minecraft:gravel
        ]


}
