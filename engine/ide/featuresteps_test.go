package ide

import (
	"strings"
	"testing"
)

// Each list in a biome's features is labelled with its generation step,
// in a Biome and in a CustomBiome's biome alike. A list a const made is
// not labelled, nor anything after a spread.
func TestFeatureStepHints(t *testing.T) {
	src := `Biome b {
  has_precipitation = false
  temperature = 2
  downfall = 0
  effects = { water_color = "#3F76E4" }
  carvers = []
  features = [[], [], [minecraft:ore_coal_upper]]
}
CustomBiome c {
  climate = { temperature = 0, humidity = 0, continentalness = 0, erosion = 0, weirdness = 0, depth = 0, offset = 0 }
  biome = { has_precipitation = false, temperature = 2, downfall = 0, effects = { water_color = "#3F76E4" }, carvers = [], features = [[]] }
}
const steps = [[], []]
Biome d {
  has_precipitation = false
  temperature = 2
  downfall = 0
  effects = { water_color = "#3F76E4" }
  carvers = []
  features = steps
}
Biome e {
  has_precipitation = false
  temperature = 2
  downfall = 0
  effects = { water_color = "#3F76E4" }
  carvers = []
  features = [...steps, []]
}
`
	s, _ := session(t, src)
	var got []string
	for _, h := range s.InlayHints("main.mms") {
		if h.Before {
			got = append(got, h.Label+"@"+src[h.Off:h.Off+2])
		}
	}
	want := "raw_generation:@[] lakes:@[] local_modifications:@[m raw_generation:@[]"
	if strings.Join(got, " ") != want {
		t.Errorf("hints:\n got %s\nwant %s", strings.Join(got, " "), want)
	}
}
