package project

import (
	"context"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// TestRelated checks that a problem with text a Blueprint or const wrote
// stays where the user's code put it in, and points at the line that
// wrote it.
func TestRelated(t *testing.T) {
	for _, c := range []struct{ src, at, want, msg string }{
		{
			src: "Blueprint Thing {\n  n: int\n  emit {\n    Noise ${self.path} { base_octave = n, amplitude_modifiers = \"oops\" }\n  }\n}\nThing a { n = 1 }\n",
			at:  "a", want: `"oops"`, msg: "written in Blueprint Thing",
		},
		{
			src: "const bad = { base_octave = 1, amplitude_modifiers = \"oops\" }\nNoise a = bad\n",
			at:  "bad\n", want: "const bad = {", msg: "written in const bad",
		},
	} {
		p := New(spec.MustDefault(), Config{Namespace: "demo"})
		p.SetFile("a.mms", c.src)
		r, _ := p.Analyze(context.Background())
		var ds, echoes []syntax.Diagnostic
		for _, d := range r.Diagnostics("a.mms") {
			if d.Echo {
				echoes = append(echoes, d)
			} else {
				ds = append(ds, d)
			}
		}
		if len(ds) != 1 || len(ds[0].Related) != 1 || len(echoes) != 1 {
			t.Errorf("%q: diagnostics = %+v, echoes = %+v", c.src, ds, echoes)
			continue
		}
		src := []rune(c.src)
		d, rel := ds[0], ds[0].Related[0]
		// The problem is shown again where it was written, pointing back
		// at the use.
		if e := echoes[0]; e.Sp != rel.Sp || e.Message != d.Message || len(e.Related) != 1 || e.Related[0].Sp != d.Sp || len(e.Fixes) > 0 {
			t.Errorf("%q: echo = %+v", c.src, e)
		}
		if got := string(src[d.Sp.Start:]); !strings.HasPrefix(got, c.at) {
			t.Errorf("reported at %q, want %q", got, c.at)
		}
		if got := string(src[rel.Sp.Start:rel.Sp.End]); !strings.HasPrefix(got, c.want) || rel.Path != "a.mms" || rel.Message != c.msg {
			t.Errorf("related = %q in %s (%s), want %q (%s)", got, rel.Path, rel.Message, c.want, c.msg)
		}
	}
}

// TestRelatedLine checks that a problem in a Function a Blueprint emits
// points at the line of the body that wrote it, not the Blueprint, and
// is shown on that line too.
func TestRelatedLine(t *testing.T) {
	src := `Blueprint Throttle {
  label: string
  rate: int = 20
  emit {
    Objective cd: dummy
    Function ${self.path} {
      if (cd[@s] > 0) {
        cd[@s] -= 1
      }
      cd = rate
      ${label}_cd[@s] = ${rat}
    }
  }
}
Throttle wand { label = "w" }
`
	p := New(spec.MustDefault(), Config{Namespace: "demo"})
	p.SetFile("a.mms", src)
	r, _ := p.Analyze(context.Background())
	lines := strings.Split(src, "\n")
	f := r.Unit("a.mms").File
	type got struct {
		line      string
		echo      bool
		code, rel string
	}
	var all []got
	for _, d := range r.Diagnostics("a.mms") {
		l, _ := f.LineCol(d.Sp.Start)
		g := got{line: strings.TrimSpace(lines[l-1]), echo: d.Echo, code: d.Code}
		if len(d.Related) == 1 {
			rl, _ := f.LineCol(d.Related[0].Sp.Start)
			g.rel = strings.TrimSpace(lines[rl-1])
		}
		all = append(all, g)
	}
	want := []got{
		{line: "cd = rate", echo: true, code: "mms/type-mismatch", rel: `Throttle wand { label = "w" }`},
		{line: "${label}_cd[@s] = ${rat}", echo: true, code: "mms/name-piece", rel: `Throttle wand { label = "w" }`},
		{line: `Throttle wand { label = "w" }`, code: "mms/type-mismatch", rel: "cd = rate"},
		{line: `Throttle wand { label = "w" }`, code: "mms/name-piece", rel: "${label}_cd[@s] = ${rat}"},
		// w_cd is not declared: the templated name was read.
		{line: "${label}_cd[@s] = ${rat}", echo: true, code: "mms/unknown-name", rel: `Throttle wand { label = "w" }`},
		{line: `Throttle wand { label = "w" }`, code: "mms/unknown-name", rel: "${label}_cd[@s] = ${rat}"},
	}
	for _, w := range want {
		found := false
		for _, g := range all {
			found = found || g == w
		}
		if !found {
			t.Errorf("missing %+v in %+v", w, all)
		}
	}
	if len(all) != len(want) {
		t.Errorf("diagnostics = %+v", all)
	}
	for _, d := range r.Diagnostics("a.mms") {
		if d.Code == "mms/type-mismatch" && !strings.Contains(d.Message, "set a score with cd[@s] = ...") {
			t.Errorf("message = %q", d.Message)
		}
		if d.Code == "mms/name-piece" && !strings.Contains(d.Message, "nothing at compile time is named 'rat'") {
			t.Errorf("message = %q", d.Message)
		}
	}
}

// TestRelatedLineAfterEdit checks that the line is still found when the
// analysis reuses the checked declarations of an earlier one.
func TestRelatedLineAfterEdit(t *testing.T) {
	bp := "Blueprint B {\n  emit {\n    Objective cd: dummy\n    Function ${self.path} {\n      cd = 1\n    }\n  }\n}\n"
	p := New(spec.MustDefault(), Config{Namespace: "demo"})
	p.SetFile("b.mms", bp)
	p.SetFile("a.mms", "B one { }\n")
	p.Analyze(context.Background())
	p.SetFile("a.mms", "B one { }\nFunction other {\n}\n")
	r, _ := p.Analyze(context.Background())
	var echoes int
	for _, d := range r.Diagnostics("b.mms") {
		if d.Echo {
			echoes++
			if l, _ := r.Unit("b.mms").File.LineCol(d.Sp.Start); l != 5 {
				t.Errorf("echo on line %d, want 5", l)
			}
		}
	}
	if echoes != 1 {
		t.Errorf("echoes = %d: %+v", echoes, r.Diagnostics("b.mms"))
	}
}

// TestUnusedInstanceNotEchoed checks that an unused Blueprint instance
// is reported at the instance, by the kind the user wrote, and not
// echoed over the whole Blueprint, whose body no single line of wrote it.
func TestUnusedInstanceNotEchoed(t *testing.T) {
	src := `Tier upper {
  tier = 1
}
Blueprint Tier {
  tier: int
  emit {
    primary DensityFunction ${self.path} = Constant(tier)
  }
}
`
	p := New(spec.MustDefault(), Config{Namespace: "demo"})
	p.SetFile("a.mms", src)
	r, _ := p.Analyze(context.Background())
	var unused []syntax.Diagnostic
	for _, d := range r.Diagnostics("a.mms") {
		if d.Code == "lint/unused" {
			unused = append(unused, d)
		}
	}
	if len(unused) != 1 {
		t.Fatalf("unused = %+v", unused)
	}
	d := unused[0]
	if d.Echo || string([]rune(src)[d.Sp.Start:d.Sp.End]) != "upper" || d.Message != "Tier 'demo:upper' is never used" {
		t.Errorf("unused = %+v", d)
	}
}
