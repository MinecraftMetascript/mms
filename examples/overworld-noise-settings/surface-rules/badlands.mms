namespace MySpace;

InBadlands := SurfaceCondition {
    Biome [
        minecraft:badlands
        minecraft:eroded_badlands
        minecraft:wooded_badlands
    ]
}

SkyTerracotta := SurfaceRule {
    If (AboveSurface) Block stone
}

NonHoleOrangeTerracotta := SurfaceRule { If (Not(Hole)) Block orange_terracotta }

TerracottaBands := SurfaceRule {
    If (StoneDepth Floor 0 Sub 0)
        Sequence [
            If (YAbove 74 1 Add)
                Sequence [
                    If (Or (
                        Noise minecraft:surface [-0.909 -0.5454]
                        Noise minecraft:surface [-0.1818 0.1818]
                        Noise minecraft:surface [0.5454 0.909]
                    ))
                        Block minecraft:terracotta

                    Bandlands
                ]
        ]
}

WhiteTerracotta := SurfaceRule { If (AboveWater -6 -1 Add) Block white_terracotta }
OrangeTerracotta := SurfaceRule { If (YAbove 63 0 Sub ) Block orange_terracotta }
StoneAndGravel := SurfaceRule {
    Sequence [
        If (StoneDepth Ceiling 0 Sub 0) Block stone
        Block Gravel
    ]
}

SurfaceSands := SurfaceRule {
    If (AboveWater -1 0 Sub) Sequence [
        If (StoneDepth Ceiling 0 Sub 0) Block minecraft:red_sandstone
        Block minecraft:red_sand
    ]
}

OrangeTerracottaEdge := SurfaceRule {
    If ( And (
        YAbove 63 0 Sub
        Not(YAbove 74 1 Add)
    ) ) Block minecraft:orange_terracotta
}

Badlands := SurfaceRule {
    If (InBadlands) Sequence [
        If (YAbove 63 0 Sub) Sequence [
            SkyTerracotta
            TerracottaBands
            SurfaceSands
            NonHoleOrangeTerracotta
            WhiteTerracotta
            StoneAndGravel
        ]
        If (YAbove 63 -1 Add) Sequence [
            OrangeTerracottaEdge
            Bandlands
        ]
        If (StoneDepth Floor 0 Add 0) WhiteTerracotta
    ]
}
