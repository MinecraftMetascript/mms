package project

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/spec"
)

var update = os.Getenv("MMS_UPDATE") != ""

func loadDir(t *testing.T, dir string) *Project {
	t.Helper()
	cfg, err := LoadConfig(dir)
	if err != nil {
		t.Fatal(err)
	}
	p := New(spec.MustDefault(), cfg)
	paths, err := cfg.SourceFiles(dir)
	if err != nil {
		t.Fatal(err)
	}
	for _, rel := range paths {
		b, _ := os.ReadFile(filepath.Join(dir, rel))
		p.SetFile(rel, string(b))
	}
	return p
}

func dumpFiles(files []OutFile) string {
	var sb strings.Builder
	for _, f := range files {
		fmt.Fprintf(&sb, "=== %s\n%s", f.Path, f.Data)
	}
	return sb.String()
}

// TestExamples builds every project under examples/ and compares the pack
// with testdata/examples/<name>.golden.
func TestExamples(t *testing.T) {
	dirs, _ := filepath.Glob("../../examples/*")
	if len(dirs) == 0 {
		t.Fatal("no examples found")
	}
	for _, dir := range dirs {
		name := filepath.Base(dir)
		t.Run(name, func(t *testing.T) {
			p := loadDir(t, dir)
			r, err := p.Analyze(context.Background())
			if err != nil {
				t.Fatal(err)
			}
			for _, u := range r.Units {
				for _, d := range r.Diagnostics(u.Path) {
					t.Errorf("%s: %s %s", u.Path, d.Code, d.Message)
				}
			}
			got := dumpFiles(r.Emit())
			golden := filepath.Join("testdata", "examples", name+".golden")
			if update {
				os.MkdirAll(filepath.Dir(golden), 0o755)
				os.WriteFile(golden, []byte(got), 0o644)
				return
			}
			want, _ := os.ReadFile(golden)
			if got != string(want) {
				t.Errorf("output differs from %s; rerun with MMS_UPDATE=1 to accept:\n%s", golden, got)
			}
		})
	}
}

func TestCachingAndDeterminism(t *testing.T) {
	p := New(spec.MustDefault(), DefaultConfig())
	p.SetFile("a.mms", "Namespace demo\nDensityFunction a = Noise(continentalness, 1, 0) * 2\n")
	p.SetFile("b.mms", "Namespace demo\nDensityFunction b = a + 1\n")
	r1, _ := p.Analyze(context.Background())
	out1 := r1.Emit()

	// Editing b's text re-checks b only.
	p.SetFile("b.mms", "Namespace demo\nDensityFunction b = a + 2\n")
	r2, _ := p.Analyze(context.Background())
	if r1.Order[0] != r2.Order[0] {
		t.Error("a was re-checked although only b changed")
	}
	if r1.Order[1] == r2.Order[1] {
		t.Error("b was not re-checked after its text changed")
	}
	// Adding a name a does not look up keeps a's result.
	p.SetFile("c.mms", "Namespace demo\nDensityFunction c = 1\n")
	r3, _ := p.Analyze(context.Background())
	if r2.Order[0] != r3.Order[0] {
		t.Error("a was re-checked although no name it uses changed")
	}
	// Declaring demo:continentalness changes what a's noise resolves to.
	p.SetFile("c.mms", "Namespace demo\nNoise continentalness = { base_octave = -1, amplitude_modifiers = [1] }\n")
	r3b, _ := p.Analyze(context.Background())
	if r3.Order[0] == r3b.Order[0] {
		t.Error("a kept its result although a name it uses was declared")
	}
	if !strings.Contains(string(jsonv.Compact(r3b.Order[0].Value)), "demo:continentalness") {
		t.Errorf("a = %s", jsonv.Compact(r3b.Order[0].Value))
	}

	p.SetFile("b.mms", "Namespace demo\nDensityFunction b = a + 1\n")
	p.RemoveFile("c.mms")
	r4, _ := p.Analyze(context.Background())
	if !bytes.Equal([]byte(dumpFiles(out1)), []byte(dumpFiles(r4.Emit()))) {
		t.Error("rebuilding the same sources gave different output")
	}
}

func TestCancel(t *testing.T) {
	p := New(spec.MustDefault(), DefaultConfig())
	p.SetFile("a.mms", "DensityFunction a = 1\n")
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := p.Analyze(ctx); err == nil {
		t.Error("Analyze ignored a cancelled context")
	}
}

func TestDiagnosticsAndTags(t *testing.T) {
	p := New(spec.MustDefault(), DefaultConfig())
	p.SetFile("a.mms", `Namespace demo
BlockTag ores = [diamond_ore, #minecraft:coal_ores]
BlockTag ores = [emerald_ore, diamond_ore]
FunctionTag minecraft:tick = [demo:tick]
Function tick {
  /say hi
}
Objective kills = deathCount
Noise Bad_Name = { base_octave = 0 }
DensityFunction dup = 1
DensityFunction dup = 2
`)
	r, _ := p.Analyze(context.Background())
	var codes []string
	for _, d := range r.Diagnostics("a.mms") {
		codes = append(codes, d.Code)
	}
	if got := strings.Join(codes, " "); got != "mms/invalid-name mms/duplicate" {
		t.Errorf("diagnostics = %s", got)
	}
	got := dumpFiles(r.Emit())
	for _, want := range []string{
		"=== data/demo/tags/block/ores.json\n{\n  \"values\": [\n    \"minecraft:diamond_ore\",\n    \"#minecraft:coal_ores\",\n    \"minecraft:emerald_ore\"\n  ]\n}\n",
		"=== data/minecraft/tags/function/tick.json\n{\n  \"values\": [\n    \"demo:tick\"\n  ]\n}\n",
		"=== data/demo/function/mms_gen/load.mcfunction\nscoreboard objectives add kills deathCount\n",
		"=== data/demo/function/tick.mcfunction\nsay hi\n",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("missing %q in output:\n%s", want, got)
		}
	}
}

// TestIncrementalMatchesFresh applies a series of edits and checks that
// the incrementally analysed project reports exactly what a fresh one does.
func TestIncrementalMatchesFresh(t *testing.T) {
	base, err := os.ReadFile("../sema/testdata/density.mms")
	if err != nil {
		t.Fatal(err)
	}
	src := string(base)
	edits := []func(string) string{
		func(s string) string { return "// top\n" + s },
		func(s string) string { return strings.Replace(s, "Clmp", "Clamp", 1) },
		func(s string) string { return s + "\nNoise extra = { base_octave = -2, amplitude_modifiers = [1] }\n" },
		func(s string) string { return strings.Replace(s, "// top\n", "", 1) },
		func(s string) string { return strings.Replace(s, "Namespace example", "Namespace other", 1) },
		func(s string) string {
			return "Noise continentalness = { base_octave = -1, amplitude_modifiers = [1] }\n" + s
		},
	}
	p := New(spec.MustDefault(), DefaultConfig())
	p.SetFile("main.mms", src)
	p.Analyze(context.Background())
	describe := func(r *Result) string {
		var sb strings.Builder
		for _, d := range r.Diagnostics("main.mms") {
			fmt.Fprintf(&sb, "%v %s %s\n", d.Sp, d.Code, d.Message)
		}
		sb.WriteString(dumpFiles(r.Emit()))
		for _, d := range r.Order {
			for n := range d.Info {
				if !d.D.Sp.Contains(n.Span().Start) {
					fmt.Fprintf(&sb, "stale info node %T at %v outside %v\n", n, n.Span(), d.D.Sp)
				}
			}
		}
		return sb.String()
	}
	for i, edit := range edits {
		src = edit(src)
		p.SetFile("main.mms", src)
		got, _ := p.Analyze(context.Background())
		fresh := New(spec.MustDefault(), DefaultConfig())
		fresh.SetFile("main.mms", src)
		want, _ := fresh.Analyze(context.Background())
		if g, w := describe(got), describe(want); g != w {
			t.Errorf("edit %d: incremental result differs\n--- incremental\n%s\n--- fresh\n%s", i, g, w)
		}
	}
}
