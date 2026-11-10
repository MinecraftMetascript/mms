package project

import (
	"strings"
	"testing"
)

// TestInlinePlacedFeatureInBiome checks a Biome whose features list holds
// an inline PlacedFeature with an inline tree feature. The tree's
// `oak_log[axis=y]` is written as { id, properties }, the block state
// form 26.x uses; the old { Name, Properties } failed the schema and
// surfaced as an internal error on the Biome.
func TestInlinePlacedFeatureInBiome(t *testing.T) {
	out := emitText(t, `
Biome cave_meadow {
  temperature = 0
  downfall = 0
  has_precipitation = false
  effects = { water_color = "#4553cb" }
  carvers = []
  features = [
    [],
    [cave_vines],
    [
      PlacedFeature(
        {
          type = tree
          minimum_size = TwoLayersFeatureSize(6)
          below_trunk_provider = { type = rule_based, rules = [] }
          trunk_provider = Simple(oak_log[axis=y])
          foliage_provider = Simple(oak_leaves)
          trunk_placer = StraightTrunkPlacer(2, 1, 2)
          foliage_placer = PineFoliagePlacer(
            { type = uniform, min_inclusive = 1, max_inclusive = 3 },
            0,
            { type = biased_to_bottom, max_inclusive = 5, min_inclusive = 1 },
          )
          decorators = []
        },
        [
          Count(2),
          InSquare(),
          EnvironmentScan(direction_of_search = down, 8, MatchingBlocks([grass_block, dirt])),
        ],
      ),
    ],
  ]
}
`)
	biome := compact(out["data/demo/worldgen/biome/cave_meadow.json"])
	for _, want := range []string{
		`"trunk_provider":{"type":"minecraft:simple","state":{"id":"minecraft:oak_log","properties":{"axis":"y"}}}`,
		`"foliage_provider":{"type":"minecraft:simple","state":"minecraft:oak_leaves"}`,
	} {
		if !strings.Contains(biome, want) {
			t.Errorf("missing %s in\n%s", want, biome)
		}
	}
}
