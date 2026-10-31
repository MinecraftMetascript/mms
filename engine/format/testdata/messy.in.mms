Namespace   example
// Objectives used by the kill tracker
Objective kills=deathCount
Objective deaths =   deathCount


/** Very long density function that must wrap. */
DensityFunction long_one=Add(Mul(Noise(minecraft:continentalness).XzScale(0.25).YScale(0.0), 0.75), Clamp(Abs(ridges), -1, 1), Constant(4)) // trailing note
LootTable blocks/ruby_ore{pools=[{rolls=1,entries=[{type="minecraft:item",name='minecraft:ruby'}]},{rolls=2}]
  // functions apply to every pool
  functions = [ ]
  "minecraft:weird key" = 3
}
SurfaceRule s = If(Biome(badlands, eroded_badlands, wooded_badlands) && !YAbove(74) && StoneDepth(0, false, 0, floor) || Water(-1, 0)) { Block(red_sand) }
Function   f(a : int,b) {
    let x=1+2*3
    x+=   ( 4 - 1 ) * 2


    /say   hello    "two  spaces"   {"a":  1}
    while(x>0){x-=1}
    for i in 0..10 {
  /say ${i}
}
}
/* final
   comment */
