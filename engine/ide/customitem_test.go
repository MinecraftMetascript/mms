package ide

import (
	"strings"
	"testing"
)

func TestCustomItemEditor(t *testing.T) {
	const item = "CustomItem ruby { base = emerald, name = \"Ruby\" }\n"
	cases := []struct {
		name, src string
		want      []string
	}{
		{"kind", "Custom|", []string{"CustomItem"}},
		{"fields", "CustomItem ruby {\n  |\n}", []string{"base", "name", "texture", "model", "item_model", "components"}},
		{"base", "CustomItem ruby { base = emer| }", []string{"emerald"}},
		{"components", "CustomItem ruby { base = emerald, components = { |} }", []string{"rarity", "max_stack_size"}},
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
		{"CustomItem ruby { ba|se = emerald }", "vanilla item"},
		{"Custom|Item ruby { base = emerald }", "Blueprint CustomItem"},
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

// A custom item used as a plain item id shows its base item after it.
func TestCustomItemInlayHint(t *testing.T) {
	s, _ := session(t, "CustomItem ruby { base = emerald }\nRecipe r = Shapeless([ruby, stick], ruby * 2)\n|")
	hints := s.InlayHints("main.mms")
	if len(hints) != 1 || hints[0].Label != "→ emerald" {
		t.Fatalf("hints = %+v", hints)
	}
	src := []rune(s.R.Unit("main.mms").File.Src)
	if got := string(src[hints[0].Off-4 : hints[0].Off]); got != "ruby" {
		t.Errorf("hint after %q", got)
	}
}

func TestColorHover(t *testing.T) {
	for _, tc := range []struct{ src, want string }{
		{`Book b { title = "T", author = "A", pages = [{ text = "x", shadow_co|lor = "#FF0000" }] }`, `"#AARRGGBB"`},
		{`Biome b { effects = { water_co|lor = "#3F76E4" } }`, `"#RRGGBB"`},
	} {
		s, off := session(t, tc.src)
		if h, _ := s.Hover("main.mms", off); !strings.Contains(h, tc.want) || !strings.Contains(h, "from 0 to 1") {
			t.Errorf("%q: hover = %q, want %q", tc.src, h, tc.want)
		}
	}
}

// What the prelude's Blueprints write for an instance is not offered by
// name, and neither are its @internal Blueprints.
func TestPreludeOutputHidden(t *testing.T) {
	s, off := session(t, "CustomItem w {\n  base = stick\n  on_right_click = run {\n    /say hi\n  }\n}\nAdvancement a = { rewards = { function = |} }")
	items, _ := s.Completion("main.mms", off)
	for _, l := range labels(items) {
		if strings.Contains(l, "mms_gen") {
			t.Errorf("offered %q", l)
		}
	}
	s, off = session(t, "Item|")
	if got := labels(func() []Item { it, _ := s.Completion("main.mms", off); return it }()); has(got, "ItemHook") || !has(got, "ItemModel") {
		t.Errorf("kinds = %v", got)
	}
}
