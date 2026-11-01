package ide

import (
	"strings"
	"testing"
)

func TestCustomItemEditor(t *testing.T) {
	const item = "CustomItem ruby = { base = emerald, name = \"Ruby\" }\n"
	cases := []struct {
		name, src string
		want      []string
	}{
		{"kind", "Custom|", []string{"CustomItem"}},
		{"fields", "CustomItem ruby = {\n  |\n}", []string{"base", "name", "texture", "model", "item_model", "components"}},
		{"base", "CustomItem ruby = { base = emer| }", []string{"emerald"}},
		{"components", "CustomItem ruby = { base = emerald, components = { |} }", []string{"rarity", "max_stack_size"}},
		{"result", item + "Recipe r = Shapeless([emerald], |)", []string{"ruby"}},
	}
	for _, tc := range cases {
		s, off := session(t, tc.src)
		items, _ := s.Completion("main.mms", off)
		if got := labels(items); !has(got, tc.want...) {
			t.Errorf("%s: completions %v lack %v", tc.name, got, tc.want)
		}
	}

	for _, tc := range []struct{ src, want string }{
		{"CustomItem ruby = { ba|se = emerald }", "vanilla item"},
		{"Custom|Item ruby = { base = emerald }", "CustomItem"},
	} {
		s, off := session(t, tc.src)
		if h, _ := s.Hover("main.mms", off); !strings.Contains(h, tc.want) {
			t.Errorf("%q: hover = %q, want %q", tc.src, h, tc.want)
		}
	}

	for _, src := range []string{
		item + "Recipe r = Shapeless([emerald], ru|by * 2)",
		item + "LootTable l { pools = [Pool(1, [Item(name = ru|by)])] }",
		item + "Function f {\n  /give @s ${ru|by}\n}",
	} {
		s, off := session(t, src)
		defs := s.Definition("main.mms", off)
		if len(defs) != 1 || defs[0].Span.Start != strings.Index(src, "ruby") {
			t.Errorf("%q: definition = %+v", src, defs)
		}
	}
}
