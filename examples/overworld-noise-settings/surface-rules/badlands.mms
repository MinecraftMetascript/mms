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

NonHoleOrangeTerracotta := SurfaceRule { If (!Hole) Block orange_terracotta }

TerracottaBands := SurfaceRule {
    If (StoneDepth Floor 0 sub 0)
        Sequence [
            If (YAbove Absolute(74) 1 add)
                Sequence [
                    If (or (
                        Noise minecraft:surface [-0.909, -0.5454]
                        Noise minecraft:surface [-0.1818, 0.1818]
                        Noise minecraft:surface [0.5454, 0.909]
                    ))
                        Block minecraft:terracotta

                    Bandlands
                ]
        ]
}

WhiteTerracotta := SurfaceRule { If (AboveWater -6 -1 add) Block white_terracotta }
OrangeTerracotta := SurfaceRule { If (YAbove Absolute(63) 0 sub ) Block orange_terracotta }
StoneAndGravel := SurfaceRule {
    Sequence [
        If (StoneDepth Ceiling 0 sub 0) Block stone
        Block Gravel
    ]
}

SurfaceSands := SurfaceRule {
    If (AboveWater -1 0 sub) Sequence [
        If (StoneDepth Ceiling 0 sub 0) Block minecraft:red_sandstone
        Block minecraft:red_sand
    ]
}

OrangeTerracottaEdge := SurfaceRule {
    If ( and (
        YAbove Absolute(63) 0 sub
        !YAbove Absolute(74) 1 add
    ) ) Block minecraft:orange_terracotta
}

Badlands := SurfaceRule {
    If (InBadlands) Sequence [
        If (YAbove Absolute(63) 0 sub) Sequence [
            SkyTerracotta
            TerracottaBands
            SurfaceSands
            NonHoleOrangeTerracotta
            WhiteTerracotta
            StoneAndGravel
        ]
        If (YAbove Absolute(63) -1 add) Sequence [
            OrangeTerracottaEdge
            Bandlands
        ]
        If (StoneDepth Floor 0 add 0) WhiteTerracotta
    ]
}
