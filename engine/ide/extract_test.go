package ide

import (
	"context"
	"sort"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/project"
	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// applyAll applies edits to src, given in rune offsets of src.
func applyAll(src string, edits []Edit) string {
	r := []rune(src)
	sort.Slice(edits, func(i, j int) bool { return edits[i].Span.Start > edits[j].Span.Start })
	for _, e := range edits {
		r = append(r[:e.Span.Start], append([]rune(e.Text), r[e.Span.End:]...)...)
	}
	return string(r)
}

// extracted runs the extraction titled with prefix at the cursor and
// returns the new source, failing when the result does not check.
func extracted(t *testing.T, src, prefix string) (string, Extraction) {
	t.Helper()
	s, off := session(t, src)
	plain := strings.Replace(src, "|", "", 1)
	var titles []string
	for _, x := range s.Extractions("main.mms", off, off) {
		titles = append(titles, x.Title)
		if !strings.HasPrefix(x.Title, prefix) {
			continue
		}
		out := applyAll(plain, x.Edits)
		if got := string([]rune(out)[x.NameAt : x.NameAt+len([]rune(x.Name))]); got != x.Name {
			t.Errorf("NameAt points at %q, want %q", got, x.Name)
		}
		p := project.New(spec.MustDefault(), project.Config{Namespace: "test"})
		p.SetFile("main.mms", out)
		r, err := p.Analyze(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		for _, d := range r.Diagnostics("main.mms") {
			if d.Severity == syntax.SeverityError {
				t.Errorf("after extracting: %s %s\n%s", d.Code, d.Message, out)
			}
		}
		return out, x
	}
	t.Fatalf("no extraction %q; got %q", prefix, titles)
	return "", Extraction{}
}

func TestExtractHoistedNoise(t *testing.T) {
	out, x := extracted(t, `/** Hills. */
DensityFunction hills = Noise({
  base_octave = -6
  amplitude_modifiers = [1]
|}, 1, 1)
`, "Extract { ... } to a Noise")
	if !strings.HasPrefix(out, "Noise hills_noise = {\n  base_octave = -6\n  amplitude_modifiers = [1]\n}\n\n/** Hills. */\nDensityFunction hills = Noise(hills_noise, 1, 1)") {
		t.Errorf("got\n%s", out)
	}
	if x.Name != "hills_noise" {
		t.Errorf("name = %q", x.Name)
	}
}

// TestExtractNested checks a value nested in a body: its lines move left
// to the top level, it is named after its field, and a name in use gets
// a number.
func TestExtractNested(t *testing.T) {
	out, _ := extracted(t, `Noise custom_final_density = { base_octave = -4, amplitude_modifiers = [1] }

NoiseSettings custom {
  default_block = stone
  default_fluid = water
  disable_mob_generation = false
  material_rule = [Block(stone)]
  sea_level = 0
  noise = { min_y = -64, height = 384 }
  legacy_random_source = false
  noise_router = {
    temperature = 0
    vegetation = 0
    continents = 0
    erosion = 0
    depth = 0
    ridges = 0
    chunk_surface_level = 0
    final_density = Max(
      Gradient(axis = y, from_coordinate = 0, to_coordinate = 256, from_value = 0, to_value = 1),
      |0.5,
    )
  }
  aquifers = { barrier = 0, fluid_level_floodedness = 0, fluid_level_spread = 0, lava = 0, exclusion = 0, surface_level = 0 }
  spawn_target = []
}
`, "Extract Max(...) to a DensityFunction")
	if !strings.Contains(out, "\nDensityFunction custom_final_density_2 = Max(\n  Gradient(axis = y, from_coordinate = 0, to_coordinate = 256, from_value = 0, to_value = 1),\n  0.5,\n)\n\nNoiseSettings custom {") {
		t.Errorf("got\n%s", out)
	}
	if !strings.Contains(out, "final_density = custom_final_density_2\n") {
		t.Errorf("value not replaced:\n%s", out)
	}
}

func TestExtractInlineFunction(t *testing.T) {
	out, _ := extracted(t, `Advancement found {
  criteria = { tick = { trigger = tick } }
  rewards = { function = run {
    /say |found it
  } }
}
`, "Extract run { /say found it } to a Function declaration")
	if !strings.HasPrefix(out, "Function found_function {\n  /say found it\n}\n\nAdvancement found {") {
		t.Errorf("got\n%s", out)
	}
}

// TestExtractNone checks that ids and plain numbers are not offered.
func TestExtractNone(t *testing.T) {
	for _, src := range []string{
		"Noise n = { base_octave = -4, amplitude_modifiers = [1] }\nDensityFunction d = Noise(|n, 1, 1)\n",
		"DensityFunction d = Add(|1, 2)\n",
	} {
		s, off := session(t, src)
		for _, x := range s.Extractions("main.mms", off, off) {
			t.Errorf("%q: unexpected %q", src, x.Title)
		}
	}
}
