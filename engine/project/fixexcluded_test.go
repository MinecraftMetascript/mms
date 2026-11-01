package project

import (
	"context"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/spec"
)

// A biome attribute that only a dimension type can set is an error, with a
// fix that removes the line.
func TestNonPositionalAttributeFix(t *testing.T) {
	src := "Biome b = {\n  temperature = 1\n  downfall = 0\n  has_precipitation = false\n  effects = { water_color = 1 }\n  carvers = []\n  features = []\n  attributes = {\n    \"minecraft:gameplay/fast_lava\" = true\n    \"minecraft:visual/fog_color\" = \"#330808\"\n  }\n}\n"
	p := New(spec.MustDefault(), Config{Namespace: "demo"})
	p.SetFile("a.mms", src)
	r, _ := p.Analyze(context.Background())
	var found bool
	for _, d := range r.Diagnostics("a.mms") {
		if d.Code != "mms/not-allowed-here" {
			continue
		}
		found = true
		if !strings.Contains(d.Message, "dimension type") {
			t.Errorf("message = %q", d.Message)
		}
		if len(d.Fixes) != 1 {
			t.Fatalf("fixes = %v", d.Fixes)
		}
		e := d.Fixes[0].Edits[0]
		rs := []rune(src)
		got := string(rs[:e.Sp.Start]) + e.Text + string(rs[e.Sp.End:])
		want := strings.Replace(src, "    \"minecraft:gameplay/fast_lava\" = true\n", "", 1)
		if got != want {
			t.Errorf("fix:\n got %q\nwant %q", got, want)
		}
	}
	if !found {
		t.Fatalf("no not-allowed-here diagnostic in %v", r.Diagnostics("a.mms"))
	}
}
