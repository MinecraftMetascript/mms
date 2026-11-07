/** The kinds of ore this pack adds, each with its own drop. */
Type Ore {
  name: string
  drops: Id<item>
  most: int = 3
}

const ores: [Ore] = [
{name="ruby",drops=emerald},  { name = "sapphire", drops = lapis_lazuli, most = 5 },
]

/** Ridges used by the custom continent shape. */
DensityFunction ridges = Noise(ore_veininess, 2, 1)

DensityFunction terrain = Clamp(Abs(ridges)*0.5+Gradient(axis = y, from_coordinate = -64, to_coordinate = 320, from_value = 1, to_value = -1),-1,1)

for o in ores {
  LootTable blocks/${o.name}_ore {
    pools = [Pool(1, [Item(name = o.drops)])]
  }
}

Template gem_block(gem: Id<item>): Recipe = Shaped(["GG", "GG"], { G = gem }, emerald_block)

Recipe ruby_block = gem_block(emerald)

Objective kills: "playerKillCount"

Global   deaths_total :int=0

Function on_kill {
let streak=kills[@s]
  deaths_total += 1
  if (streak >= 5) {
    tell(@a, "{gold}A player is on a streak")
  }
  else if(streak==0){
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
