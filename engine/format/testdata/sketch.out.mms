Namespace example

/** Ridges used by the custom continent shape. */
Noise ridges = Noise(-7).Amplitudes(1, 1, 0.5)

DensityFunction terrain = Clamp(Abs(ridges) * 0.5 + YGradient(-64 -> 1, 320 -> -1), -1, 1)

SurfaceRule badlands_top = If(Biome(badlands, eroded_badlands) && !YAbove(74)) {
  Block(orange_terracotta[axis=y])
}

LootTable blocks/ruby_ore {
  pools = [Pool(1).Entry(Item(ruby).When(!SilkTouch()).Count(1..3))]
}

Recipe ruby_block = Shaped("RRR", "RRR", "RRR", R = ruby) -> ruby_block

Objective kills = deathCount

Global deaths_total: int = 0

Function on_kill() {
  let streak = Score(@s, kills)
  deaths_total += 1
  if (streak >= 5) {
    /say ${@s} is on a streak
  } else if (streak == 0) {
    return
  } else {
    As(@a).At(@s) {
      /particle flame ~ ~1 ~
    }
  }
}

OnLoad {
  /say loaded
}
