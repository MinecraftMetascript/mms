package project

import (
	"context"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/spec"
)

// privateProject analyzes files in namespace demo.
func privateProject(t *testing.T, files map[string]string) *Result {
	t.Helper()
	p := New(spec.MustDefault(), Config{Namespace: "demo"})
	for path, src := range files {
		p.SetFile(path, src)
	}
	r, err := p.Analyze(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	return r
}

func codes(r *Result, path string) []string {
	var out []string
	for _, d := range r.Diagnostics(path) {
		out = append(out, d.Code+" "+d.Message)
	}
	for _, d := range r.Lints()[path] {
		out = append(out, d.Code+" "+d.Message)
	}
	return out
}

func hasCode(ds []string, code string) bool {
	for _, d := range ds {
		if strings.HasPrefix(d, code+" ") {
			return true
		}
	}
	return false
}

// Two files each keep a private const, Template, Type and Blueprint of the
// same names; each file's own one is the one it means, with no clash.
func TestPrivateNamesPerFile(t *testing.T) {
	file := func(n string) string {
		return "private const rolls = " + n + "\n" +
			"private Template twice(x: int): int = x * 2\n" +
			"private Type Pick { n: int = " + n + " }\n" +
			"private Blueprint Crate { emit { LootTable ${self.path} { pools = [Pool(twice(rolls), stone)] } } }\n" +
			"Crate crate_" + n + " { }\n" +
			"Storage picks_" + n + ": Pick\n"
	}
	r := privateProject(t, map[string]string{"a.mms": file("1"), "b.mms": file("3")})
	for _, p := range []string{"a.mms", "b.mms"} {
		if ds := codes(r, p); len(ds) > 0 {
			t.Errorf("%s: %q", p, ds)
		}
	}
	out := map[string]string{}
	for _, f := range r.Emit() {
		out[f.Path] = string(f.Data)
	}
	for n, rolls := range map[string]string{"1": "2", "3": "6"} {
		got := out["data/demo/loot_table/crate_"+n+".json"]
		if !strings.Contains(got, `"rolls": `+rolls) {
			t.Errorf("crate_%s:\n%s", n, got)
		}
	}
	load := ""
	for p, d := range out {
		if strings.Contains(p, "function/") {
			load += d
		}
	}
	for _, n := range []string{"1", "3"} {
		if !strings.Contains(load, "picks_"+n+".n set value "+n) {
			t.Errorf("picks_%s does not default n to its own file's Pick:\n%s", n, load)
		}
	}
}

// A private name is not visible from another file, and a Template keeps
// seeing its own file's private consts when it is called from elsewhere.
func TestPrivateNotVisibleElsewhere(t *testing.T) {
	r := privateProject(t, map[string]string{
		"lib.mms": "private const base = 4\nTemplate scaled(x: int): int = x * base\n",
		"use.mms": "LootTable a { pools = [Pool(scaled(2), stone)] }\nLootTable b { pools = [Pool(base, stone)] }\n",
	})
	ds := codes(r, "use.mms")
	if len(ds) != 1 || !strings.Contains(ds[0], "base") {
		t.Errorf("use.mms: %q", ds)
	}
	for _, f := range r.Emit() {
		if strings.HasSuffix(f.Path, "loot_table/a.json") && !strings.Contains(string(f.Data), `"rolls": 8`) {
			t.Errorf("a.json:\n%s", f.Data)
		}
	}
	if ds := codes(r, "lib.mms"); len(ds) > 0 {
		t.Errorf("lib.mms: %q", ds)
	}
}

// A private name that hides a public one in another file is a hint, not an
// error; the same name public and private in one file is a duplicate; and
// private on a declaration that writes a data pack file is an error.
func TestPrivateClashesAndMisuse(t *testing.T) {
	r := privateProject(t, map[string]string{
		"a.mms": "const n = 1\nLootTable a { pools = [Pool(n, stone)] }\n",
		"b.mms": "private const n = 2\nLootTable b { pools = [Pool(n, stone)] }\n",
		"c.mms": "const m = 1\nprivate const m = 2\nLootTable c { pools = [Pool(m, stone)] }\n",
		"d.mms": "private LootTable d { pools = [] }\n",
	})
	if ds := codes(r, "a.mms"); len(ds) > 0 {
		t.Errorf("a.mms: %q", ds)
	}
	if ds := codes(r, "b.mms"); len(ds) != 1 || !hasCode(ds, "mms/shadowed-private") {
		t.Errorf("b.mms: %q", ds)
	}
	for _, d := range r.Diagnostics("b.mms") {
		if d.Severity != 4 && d.Code == "mms/shadowed-private" {
			t.Errorf("shadowed-private severity %v, want a hint", d.Severity)
		}
	}
	if ds := codes(r, "c.mms"); !hasCode(ds, "mms/duplicate") {
		t.Errorf("c.mms: %q", ds)
	}
	if ds := codes(r, "d.mms"); !hasCode(ds, "mms/private") {
		t.Errorf("d.mms: %q", ds)
	}
	for _, f := range r.Emit() {
		if strings.HasSuffix(f.Path, "loot_table/b.json") && !strings.Contains(string(f.Data), `"rolls": 2`) {
			t.Errorf("b.json uses the public n:\n%s", f.Data)
		}
		if strings.HasSuffix(f.Path, "loot_table/a.json") && !strings.Contains(string(f.Data), `"rolls": 1`) {
			t.Errorf("a.json uses b's private n:\n%s", f.Data)
		}
	}
}

// An unused private const is reported even when another file uses a
// public one of the same name.
func TestPrivateUnused(t *testing.T) {
	r := privateProject(t, map[string]string{
		"a.mms": "const n = 1\n",
		"b.mms": "private const n = 2\n",
		"c.mms": "LootTable c { pools = [Pool(n, stone)] }\n",
	})
	if ds := codes(r, "b.mms"); !hasCode(ds, "lint/unused") {
		t.Errorf("b.mms: %q", ds)
	}
	if ds := codes(r, "a.mms"); hasCode(ds, "lint/unused") {
		t.Errorf("a.mms: %q", ds)
	}
}
