const gem_pool: LootPool = Pool(1..3, #gems)

const stormy: Predicate = WeatherCheck(thundering = true)
  && TimeCheck(clock = overworld, value = 13000..)

const grass_top: SurfaceRule = If(StoneDepth(Floor)) { Block(grass_block) }

LootTable chests/gems { pools = [gem_pool] }
LootTable chests/storm_gems { pools = [{ ...gem_pool, conditions = [stormy] }] }
