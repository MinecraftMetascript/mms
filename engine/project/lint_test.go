package project

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/spec"
)

func lintCodes(t *testing.T, cfg Config, src string) []string {
	t.Helper()
	p := New(spec.MustDefault(), cfg)
	p.SetFile("a.mms", src)
	r, _ := p.Analyze(context.Background())
	var out []string
	for _, d := range r.Diagnostics("a.mms") {
		line, _ := r.Unit("a.mms").File.LineCol(d.Sp.Start)
		out = append(out, d.Code+"@"+string(rune('0'+line)))
	}
	return out
}

func TestLintRules(t *testing.T) {
	cfg := Config{Namespace: "demo"}
	src := `Noise used = { base_octave = -1, amplitude_modifiers = [1] }
Noise spare = { base_octave = -1, amplitude_modifiers = [1] }
DensityFunction d = Noise(used, 1, 1) * (2 * 0.5)
LootTable t = { pools = [{ rolls = 1, entries = [{ type = "minecraft:item", name = "minecraft:diamond" }] }] }
Recipe minecraft:stick = { type = crafting_shapeless, ingredients = [oak_planks], result = { id = stick } }
PlacedFeature by_command = { feature = minecraft:oak, placement = [] }
Function f {
  /place feature demo:by_command
}
`
	got := strings.Join(lintCodes(t, cfg, src), " ")
	want := "lint/unused@2 lint/unused@3 lint/constant-expression@3 lint/quoted-name@4 lint/quoted-name@4 lint/overrides-vanilla@5 lint/use-sugar@5 lint/redundant-namespace@6"
	if got != want {
		t.Errorf("lint = %s\nwant %s", got, want)
	}

	// Severities come from mms.json, and comments silence rules.
	cfg.Lint = map[string]string{"quoted-name": "off", "unused": "warning"}
	src2 := "// mms-ignore constant-expression\nDensityFunction d = 1 + 1\nNoise spare = { base_octave = -1, amplitude_modifiers = [1] } // mms-ignore\nLootTable t = { type = \"minecraft:block\" }\n"
	if got := strings.Join(lintCodes(t, cfg, src2), " "); got != "lint/unused@2" {
		t.Errorf("configured lint = %s", got)
	}
	if err := (Config{Namespace: "x", Lint: map[string]string{"nope": "off"}}).Validate(); err == nil {
		t.Error("unknown rule accepted")
	}
}

func TestFixAll(t *testing.T) {
	p := New(spec.MustDefault(), Config{Namespace: "demo"})
	src := `LootTable t = { pools = [{ rolls = 1, entries = [{ type = "minecraft:item", name = "minecraft:diamond" }] }] }
DensityFunction d = (1 + 2) * 0.5 + Noise(minecraft:ridge, 1, 1)
`
	p.SetFile("a.mms", src)
	r0, _ := p.Analyze(context.Background())
	before := dumpFiles(r0.Emit())
	files, n, err := p.FixAll(context.Background(), "a.mms")
	text := files["a.mms"]
	if err != nil || n != 5 {
		t.Fatalf("FixAll = %d fixes, %v", n, err)
	}
	want := `LootTable t = { pools = [{ rolls = 1, entries = [diamond] }] }
DensityFunction d = 1.5 + Noise(ridge, 1, 1)
`
	if text != want {
		t.Errorf("fixed:\n%s\nwant\n%s", text, want)
	}
	r, _ := p.Analyze(context.Background())
	if after := dumpFiles(r.Emit()); after != before {
		t.Errorf("fixes changed the output:\n%s\nwas\n%s", after, before)
	}
}

func TestUseSugar(t *testing.T) {
	p := New(spec.MustDefault(), Config{Namespace: "demo"})
	src := `LootTable t = { pools = [{ rolls = { type = uniform, min = 1, max = 3 }, entries = [{ type = item, name = diamond }], condition = { type = inverted, term = { type = any_of, terms = [{ type = survives_explosion }, { type = random_chance, chance = 0.5 }] } } }] }
DensityFunction d = { type = mul, left = { type = sub, left = { type = noise, noise = ridge, xz_scale = 1, y_scale = 1 }, right = 8.0 }, right = 0.5 }
Recipe r = { type = crafting_shapeless, ingredients = [emerald_block], result = { id = emerald, count = 9 } }
`
	p.SetFile("a.mms", src)
	r0, _ := p.Analyze(context.Background())
	before := dumpFiles(r0.Emit())
	files, _, err := p.FixAll(context.Background(), "a.mms")
	if err != nil {
		t.Fatal(err)
	}
	want := `LootTable t = { pools = [{ rolls = 1..3, entries = [diamond], condition = !(SurvivesExplosion() || RandomChance(chance = 0.5)) }] }
DensityFunction d = (Noise(noise = ridge, xz_scale = 1, y_scale = 1) - 8.0) * 0.5
Recipe r = Shapeless(ingredients = [emerald_block], result = emerald * 9)
`
	if files["a.mms"] != want {
		t.Errorf("fixed:\n%s\nwant\n%s", files["a.mms"], want)
	}
	r, _ := p.Analyze(context.Background())
	if after := dumpFiles(r.Emit()); after != before {
		t.Errorf("sugar changed the output:\n%s\nwas\n%s", after, before)
	}
}

func TestLintDeprecated(t *testing.T) {
	src := `/**
 * Old way to reward.
 * @deprecated use reward
 */
Function old_reward {
  /give @s diamond
}
Function reward {
  old_reward()
  /function demo:old_reward
}
`
	p := New(spec.MustDefault(), Config{Namespace: "demo"})
	p.SetFile("a.mms", src)
	r, _ := p.Analyze(context.Background())
	var got []string
	for _, d := range r.Diagnostics("a.mms") {
		if d.Code == "lint/deprecated" {
			if !d.Deprecated {
				t.Error("finding is not tagged deprecated")
			}
			line, _ := r.Unit("a.mms").File.LineCol(d.Sp.Start)
			got = append(got, fmt.Sprint(line)+" "+d.Message)
		}
	}
	want := "9 Function 'demo:old_reward' is deprecated: use reward|10 Function 'demo:old_reward' is deprecated: use reward"
	if g := strings.Join(got, "|"); g != want {
		t.Errorf("got %s\nwant %s", g, want)
	}
}
