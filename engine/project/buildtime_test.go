package project

import (
	"context"
	"sort"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// build analyses src as namespace demo and returns the emitted files by
// path and the diagnostics.
func build(t *testing.T, src string) (map[string]string, []syntax.Diagnostic) {
	t.Helper()
	p := New(spec.MustDefault(), Config{Namespace: "demo", Description: "test"})
	p.SetFile("main.mms", src)
	r, err := p.Analyze(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	files := map[string]string{}
	for _, f := range r.Emit() {
		files[f.Path] = string(f.Data)
	}
	return files, r.Diagnostics("main.mms")
}

func buildOK(t *testing.T, src string) map[string]string {
	t.Helper()
	files, diags := build(t, src)
	for _, d := range diags {
		if d.Severity == syntax.SeverityError {
			t.Fatalf("%s at %d: %s", d.Code, d.Sp.Start, d.Message)
		}
	}
	return files
}

func paths(files map[string]string, prefix string) []string {
	var out []string
	for p := range files {
		if strings.HasPrefix(p, prefix) {
			out = append(out, strings.TrimPrefix(p, prefix))
		}
	}
	sort.Strings(out)
	return out
}

func TestTopLevelLoop(t *testing.T) {
	files := buildOK(t, `
Const biomes = ["meadow", "forest", "nether"]

for b in biomes {
  Function "shrine/enter_${b}" {
    /say Welcome to ${b}
  }
}
`)
	got := paths(files, "data/demo/function/shrine/")
	want := []string{"enter_forest.mcfunction", "enter_meadow.mcfunction", "enter_nether.mcfunction"}
	if strings.Join(got, " ") != strings.Join(want, " ") {
		t.Fatalf("functions: %v", got)
	}
	if f := files["data/demo/function/shrine/enter_nether.mcfunction"]; !strings.Contains(f, "say Welcome to nether") {
		t.Fatalf("enter_nether: %q", f)
	}
}

func TestTopLevelIfLetAndRecords(t *testing.T) {
	files := buildOK(t, `
Const planets = [
  { name = "ash", kind = "desert", sea = 40 },
  { name = "rime", kind = "ice", sea = 63 },
]

for i, p in planets {
  let label = "${p.name} (${i})"
  if (p.kind == "ice") {
    Function "cold/${p.name}" {
      /say ${label} sea ${p.sea + 1}
    }
  } else {
    Function "warm/${p.name}" {
      /say ${label}
    }
  }
}
`)
	if got := files["data/demo/function/cold/rime.mcfunction"]; !strings.Contains(got, "say rime (1) sea 64") {
		t.Fatalf("cold/rime: %q", got)
	}
	if got := files["data/demo/function/warm/ash.mcfunction"]; !strings.Contains(got, "say ash (0)") {
		t.Fatalf("warm/ash: %q", got)
	}
	if _, ok := files["data/demo/function/cold/ash.mcfunction"]; ok {
		t.Fatal("the if kept a branch it should have dropped")
	}
}

func TestConstInResources(t *testing.T) {
	files := buildOK(t, `
Const shrine_biomes = [meadow, forest]
Const title = "Shrine"

BiomeTag shrines = shrine_biomes

for b in shrine_biomes {
  BiomeTag "has/${b}" = [b, desert]
}
`)
	if got := files["data/demo/tags/worldgen/biome/shrines.json"]; !strings.Contains(got, `"minecraft:meadow"`) || !strings.Contains(got, `"minecraft:forest"`) {
		t.Fatalf("shrines: %s", got)
	}
	if got := files["data/demo/tags/worldgen/biome/has/forest.json"]; !strings.Contains(got, `"minecraft:forest"`) || !strings.Contains(got, `"minecraft:desert"`) {
		t.Fatalf("has/forest: %s", got)
	}
}

func TestUnrolledLoopInFunction(t *testing.T) {
	vm := world(t, `
Const arenas = ["red", "blue", "green"]
Global total: int = 0
Global seen: int = 0

Function count {
  for a in arenas {
    total += 1
  }
}

Function first_two {
  for i, a in arenas {
    if (i == 2) {
      break
    }
    seen += 10
  }
}

Function skip_blue {
  for a in arenas {
    if (a == "blue") {
      continue
    }
    seen += 1
  }
}
`)
	for _, f := range []string{"demo:count", "demo:first_two", "demo:skip_blue"} {
		if _, _, err := vm.Run(f); err != nil {
			t.Fatal(err)
		}
	}
	if got := score(t, vm, "$total", "mms.demo"); got != 3 {
		t.Fatalf("total = %d, want 3", got)
	}
	if got := score(t, vm, "$seen", "mms.demo"); got != 22 {
		t.Fatalf("seen = %d, want 22", got)
	}
}

func TestUnrolledCommands(t *testing.T) {
	files := buildOK(t, `
Const biomes = ["meadow", "forest", "nether"]

Function reset {
  for b in biomes {
    /scoreboard players set ${b} shrines 0
  }
}
`)
	got := files["data/demo/function/reset.mcfunction"]
	for _, b := range []string{"meadow", "forest", "nether"} {
		if !strings.Contains(got, "scoreboard players set "+b+" shrines 0") {
			t.Fatalf("reset: %q", got)
		}
	}
	if strings.Contains(got, "$") {
		t.Fatalf("build-time values should not use macros: %q", got)
	}
}

func TestRandIsStable(t *testing.T) {
	src := `
Const kinds = ["desert", "ice", "jungle", "ocean"]

for i in 0..5 {
  let r = Rand("planet/${i}")
  Function "planet_${i}" {
    /say ${r.Pick(kinds)} ${r.Int("sea", 40, 80)} ${r.Int(1, 6)}
  }
}
`
	a := buildOK(t, src)
	b := buildOK(t, src)
	for p, f := range a {
		if b[p] != f {
			t.Fatalf("%s differs between builds", p)
		}
	}
	// A new unlabeled draw before a labeled one leaves the labeled one alone.
	c := buildOK(t, strings.Replace(src, "${r.Pick(kinds)} ", "${r.Pick(kinds)} ${r.Int(0, 9)} ", 1))
	sea := func(files map[string]string) string {
		return strings.Fields(files["data/demo/function/planet_3.mcfunction"])[2]
	}
	seaC := strings.Fields(c["data/demo/function/planet_3.mcfunction"])[3]
	if sea(a) != seaC {
		t.Fatalf("labeled draw moved: %s vs %s", sea(a), seaC)
	}
}

// TestRandPinned pins the generator's output: projects depend on it, so
// a change here is a breaking change.
func TestRandPinned(t *testing.T) {
	files := buildOK(t, `
Function roll {
  /say ${Rand("a").Int("x", 1, 1000)} ${Rand("a").Int(1, 1000)} ${Rand("b").Pick("p", [1, 2, 3, 4, 5, 6, 7, 8])}
}
`)
	got := strings.TrimSpace(files["data/demo/function/roll.mcfunction"])
	if got != "say 434 987 5" {
		t.Fatalf("roll = %q", got)
	}
}

func TestBuildTimeErrors(t *testing.T) {
	cases := []struct{ name, src, code, msg string }{
		{"not a list", "for b in nowhere {\n  Function \"f_${b}\" {\n    /say hi\n  }\n}\n", "mms/not-constant", "'nowhere' is not a Const"},
		{"duplicate", "for i in [1, 2] {\n  Function same {\n    /say hi\n  }\n}\n", "mms/duplicate", "made twice"},
		{"assert", "Const xs = [1, 1]\nAssert \"xs are unique\" = Unique(xs)\n", "mms/assert", "xs are unique"},
		{"unknown field", "Const ps = [{ a = 1 }]\nfor p in ps {\n  Function \"f_${p.b}\" {\n    /say hi\n  }\n}\n", "mms/unknown-field", "no field 'b'"},
		{"cycle", "Const a = b\nConst b = a\n", "mms/const-cycle", "depends on itself"},
		{"runtime index", "Const biomes = [\"a\", \"b\"]\nGlobal n: int = 0\nFunction f {\n  /say ${biomes[n]}\n}\n", "mms/not-constant", ""},
		{"const in loop", "for i in [1] {\n  Const x = 1\n}\n", "mms/misplaced", "top level"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, diags := build(t, tc.src)
			for _, d := range diags {
				if d.Code == tc.code && strings.Contains(d.Message, tc.msg) {
					return
				}
			}
			t.Fatalf("want %s %q, got %v", tc.code, tc.msg, diags)
		})
	}
}

func TestLoopDiagnosticsCollapse(t *testing.T) {
	_, diags := build(t, `
for b in ["a", "b", "c"] {
  Function "f_${b}" {
    /say ${missing}
  }
}
`)
	n := 0
	for _, d := range diags {
		if d.Code == "mms/unknown-name" {
			n++
		}
	}
	if n != 1 {
		t.Fatalf("want one unknown-name error, got %d: %v", n, diags)
	}
}

func TestBuildTimeValuesInResourceFields(t *testing.T) {
	files := buildOK(t, `
Const powers = [{ name = "zap", power = 4 }, { name = "jolt", power = 9 }]

for p in powers {
  DamageType "${p.name}" = {
    message_id = "mms.${p.name}"
    exhaustion = p.power * 0.5
    scaling = never
  }
}
`)
	got := files["data/demo/damage_type/jolt.json"]
	if !strings.Contains(got, `"message_id": "mms.jolt"`) || !strings.Contains(got, `"exhaustion": 4.5`) {
		t.Fatalf("jolt: %s", got)
	}
}

func TestPartialIterationDiagnostic(t *testing.T) {
	_, diags := build(t, `
for n in [1, 3000000000] {
  Function "set_${n}" {
    /scoreboard players set x y ${n + 0}
  }
}
`)
	for _, d := range diags {
		if strings.Contains(d.Message, "(when n = 3000000000)") {
			return
		}
	}
	t.Fatalf("want an error that names the iteration, got %v", diags)
}

func TestConstEditRechecks(t *testing.T) {
	p := New(spec.MustDefault(), Config{Namespace: "demo", Description: "test"})
	run := func(src string) map[string]string {
		p.SetFile("main.mms", src)
		r, err := p.Analyze(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		files := map[string]string{}
		for _, f := range r.Emit() {
			files[f.Path] = string(f.Data)
		}
		return files
	}
	body := "\n\nFunction greet {\n  /say ${word}\n}\n\nfor b in biomes {\n  Function \"f_${b}\" {\n    /say ${word} ${b}\n  }\n}\n"
	run(`Const word = "hi"` + "\nConst biomes = [\"a\"]" + body)
	files := run(`Const word = "bye"` + "\nConst biomes = [\"a\", \"b\"]" + body)
	if got := files["data/demo/function/greet.mcfunction"]; !strings.Contains(got, "say bye") {
		t.Fatalf("greet after the edit: %q", got)
	}
	if got := files["data/demo/function/f_a.mcfunction"]; !strings.Contains(got, "say bye a") {
		t.Fatalf("f_a after the edit: %q", got)
	}
	if _, ok := files["data/demo/function/f_b.mcfunction"]; !ok {
		t.Fatal("f_b missing after the edit")
	}
}
