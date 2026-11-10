// JigsawStructure: a structure grown from jigsaw pieces, written in MMS on
// top of the StructureSet, Structure and TemplatePool resources it emits.

/**
 * One structure file as a weighted pool entry: `Piece("halls/hall", 3)`
 * reads structures/halls/hall.nbt and gives it weight 3. The file becomes
 * the structure `<namespace>:halls/hall` in the pack. `processors` is a
 * processor list's id, or the processors written inline:
 * `Piece("halls/hall", processors = [BlockRot(integrity = 0.8)])`.
 */
Template Piece(
  path: string,
  weight: int = 1,
  processors: Id<worldgen/processor_list> | ProcessorList = minecraft:empty,
  projection: Projection = rigid,
): WeightedElement =
  WeightedElement(weight, SinglePoolElement(projection, StructureFile("structures/${path}.nbt", path), processors))

/**
 * A structure grown from jigsaw pieces, with how often it appears:
 * `JigsawStructure dungeon { biomes = #minecraft:is_overworld, step = underground_structures, spacing = 32, separation = 12, height = AboveBottom(20), start = start, pools = { start = [Piece("start")] } }`.
 * It writes the structure set, the structure and one template pool per
 * entry of `pools`. Its name stands for the structure, as in
 * `/locate structure ${dungeon}`.
 */
Blueprint JigsawStructure {
  /** Where it can start: a biome, a tag such as `#minecraft:is_overworld`, or a list of biomes. */
  biomes: Ids<worldgen/biome>
  /** The generation step it is placed in, such as `surface_structures` or `underground_structures`. */
  step: DecorationStep
  /** The size in chunks of the grid cells it is spread over; about one structure per cell. */
  spacing: int
  /** The fewest chunks between two of them; less than spacing. */
  separation: int
  /** Seeds where in each cell it goes. Made from the structure's id unless set, so two structures never share one by accident. */
  salt: int?
  /** The height of the start piece, such as `AboveBottom(20)` or `Constant(64)`. */
  height: HeightProvider
  /** The start piece's y is taken from the ground's surface instead, with `height` added to it, as villages do. */
  heightmap: Heightmap?
  /** How many pieces deep it grows from the start piece, from 1 to 20. */
  depth: int = 7
  /** How far in blocks a piece can be from the start piece, from 1 to 128. */
  max_distance: int = 80
  /** How the terrain fits around it: `beard_thin` and `beard_box` raise the ground under it, `bury` sinks it, `encapsulate` fills around it. */
  terrain_adaptation: TerrainAdaptation?
  /** The pool the start piece comes from, usually one of `pools`. */
  start: Id<worldgen/template_pool>
  /**
   * The template pools, by name: a list of pieces, such as `start = [Piece("start")]`,
   * or a whole pool, such as `halls = { fallback = caps, elements = [Piece("halls/hall")] }`.
   * Each one is a TemplatePool of that name, so jigsaw blocks draw from
   * `<namespace>:<name>`.
   */
  pools: {string: [WeightedElement] | TemplatePool}
  /** The jigsaw block in the start piece that sits at the structure's center, by its name, such as `minecraft:city_anchor`. The start piece's own center is used unless set. */
  start_jigsaw_name: string?
  /**
   * The mobs that spawn in it instead of the biome's, by mob category:
   * `monster = { bounding_box = piece, spawns = [] }` stops hostile mobs
   * spawning inside its pieces. `full` covers its whole bounding box.
   */
  spawn_overrides: {MobCategory: SpawnOverride} = {}
  /** Whether pieces placed in water are waterlogged: `apply_waterlogging` (the game's default) or `ignore_waterlogging`. */
  liquid_settings: LiquidSettings?
  /**
   * Stand-in pool names that jigsaw blocks can draw from, each bound to a
   * real pool when the structure starts, as trial chambers do:
   * `[{ type = random, alias = my:spawner, targets = [{ data = zombies, weight = 1 }] }]`.
   */
  pool_aliases: [PoolAlias]?
  /** How many blocks from the bottom and top of the world its pieces keep: a number for both, or `{ bottom = 10, top = 0 }`. */
  dimension_padding: (int | DimensionPaddingConfig)?

  check {
    if (separation >= spacing) {
      error(separation, "separation must be less than spacing")
    }
    if (depth < 1 || depth > 20) {
      error(depth, "depth is from 1 to 20")
    }
    if (max_distance < 1 || max_distance > 128) {
      error(max_distance, "max_distance is from 1 to 128")
    }
  }

  emit {
    StructureSet ${self.path} {
      structures = [{ structure = self.id, weight = 1 }]
      placement = RandomSpread(spacing = spacing, separation = separation, salt = salt ?? hash(self.id))
    }
    primary Structure ${self.path} {
      type = jigsaw
      biomes = biomes
      step = step
      terrain_adaptation = terrain_adaptation
      spawn_overrides = spawn_overrides
      start_pool = start
      start_jigsaw_name = start_jigsaw_name
      size = depth
      start_height = height
      project_start_to_heightmap = heightmap
      max_distance_from_center = max_distance
      use_expansion_hack = false
      pool_aliases = pool_aliases
      dimension_padding = dimension_padding
      liquid_settings = liquid_settings
    }
    for name, pool in pools {
      TemplatePool ${name} = if (kind_of(pool) == "list") { fallback = minecraft:empty, elements = pool } else pool
    }
  }
}
