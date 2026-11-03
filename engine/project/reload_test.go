package project

import (
	"reflect"
	"testing"
)

func TestPlanApply(t *testing.T) {
	old := map[string][]byte{
		"pack.mcmeta":                         []byte("{}"),
		"data/a/function/load.mcfunction":     []byte("say hi"),
		"data/a/tags/block/ores.json":         []byte("{}"),
		"data/a/worldgen/biome/plains.json":   []byte("{}"),
		"data/a/tags/worldgen/biome/hot.json": []byte("{}"),
	}
	same := []OutFile{
		{Path: "pack.mcmeta", Data: []byte("{}")},
		{Path: "data/a/function/load.mcfunction", Data: []byte("say hi")},
		{Path: "data/a/tags/block/ores.json", Data: []byte("{}")},
		{Path: "data/a/worldgen/biome/plains.json", Data: []byte("{}")},
		{Path: "data/a/tags/worldgen/biome/hot.json", Data: []byte("{}")},
	}
	for _, tc := range []struct {
		name  string
		files []OutFile
		res   bool
		want  ApplyPlan
	}{
		{"same", same, false, ApplyPlan{Need: ApplyNone, Reload: []string{}, Rejoin: []string{}}},
		{"function edited", with(same, 1, "say bye"), false, ApplyPlan{Need: ApplyReload, Reload: []string{"function"}, Rejoin: []string{}}},
		{"block tag edited", with(same, 2, "[]"), false, ApplyPlan{Need: ApplyReload, Reload: []string{"tags/block"}, Rejoin: []string{}}},
		{"biome tag edited", with(same, 4, "[]"), false, ApplyPlan{Need: ApplyRejoin, Reload: []string{}, Rejoin: []string{"tags/worldgen/biome"}}},
		{"biome removed", append(append([]OutFile{}, same[:3]...), same[4]), false, ApplyPlan{Need: ApplyRejoin, Reload: []string{}, Rejoin: []string{"worldgen/biome"}}},
		{"recipe added, resources changed", append(append([]OutFile{}, same...), OutFile{Path: "data/a/recipe/x.json"}), true, ApplyPlan{Need: ApplyRejoin, Reload: []string{"recipe"}, Rejoin: []string{ResourcePackKind}}},
		{"enchantment added", append(append([]OutFile{}, same...), OutFile{Path: "data/a/enchantment/x.json"}), false, ApplyPlan{Need: ApplyRejoin, Reload: []string{}, Rejoin: []string{"enchantment"}}},
	} {
		if got := PlanApply(old, tc.files, tc.res); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("%s: got %+v, want %+v", tc.name, got, tc.want)
		}
	}
}

func with(files []OutFile, i int, data string) []OutFile {
	out := append([]OutFile{}, files...)
	out[i].Data = []byte(data)
	return out
}
