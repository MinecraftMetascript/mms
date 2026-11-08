package eval

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/format"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

var update = flag.Bool("update", false, "rewrite the golden files")

// TestGolden expands each case in testdata and compares the output
// declarations, their provenance, the diagnostics and the provenance of
// copied nodes with the case's golden file. A case is a .txt file of
// `-- name.mms --` sections; a section can also name a file under
// ../syntax/testdata/design with `-- name.mms -- design`.
func TestGolden(t *testing.T) {
	cases, _ := filepath.Glob("testdata/*.txt")
	if len(cases) == 0 {
		t.Fatal("no cases")
	}
	for _, path := range cases {
		name := strings.TrimSuffix(filepath.Base(path), ".txt")
		t.Run(name, func(t *testing.T) {
			units, srcs := readCase(t, path)
			r := Run(units, Options{Namespace: "demo"})
			got := render(units, srcs, r)
			golden := strings.TrimSuffix(path, ".txt") + ".golden"
			if *update {
				if err := os.WriteFile(golden, []byte(got), 0o644); err != nil {
					t.Fatal(err)
				}
				return
			}
			want, err := os.ReadFile(golden)
			if err != nil {
				t.Fatalf("%v (run with -update)", err)
			}
			if got != string(want) {
				t.Errorf("output differs from %s; rerun with -update to accept:\n%s", golden, got)
			}
		})
	}
}

func readCase(t *testing.T, path string) ([]Unit, map[string][]rune) {
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	var units []Unit
	srcs := map[string][]rune{}
	var name string
	var body strings.Builder
	flush := func() {
		if name == "" {
			return
		}
		f := syntax.Parse(body.String())
		for _, d := range f.Diags {
			t.Fatalf("%s: %s: parse error at %d: %s", path, name, d.Sp.Start, d.Message)
		}
		units = append(units, Unit{Path: name, File: f})
		srcs[name] = f.Src
	}
	for _, line := range strings.SplitAfter(string(data), "\n") {
		if strings.HasPrefix(line, "-- ") {
			flush()
			fields := strings.Fields(strings.TrimSpace(line))
			name = fields[1]
			body.Reset()
			if len(fields) > 3 && fields[3] == "design" {
				b, err := os.ReadFile(filepath.Join("../syntax/testdata/design", name))
				if err != nil {
					t.Fatal(err)
				}
				body.Write(b)
			}
			continue
		}
		body.WriteString(line)
	}
	flush()
	return units, srcs
}

func lineCol(src []rune, off int) string {
	line, col := 1, 1
	for i := 0; i < off && i < len(src); i++ {
		if src[i] == '\n' {
			line, col = line+1, 1
		} else {
			col++
		}
	}
	return fmt.Sprintf("%d:%d", line, col)
}

func render(units []Unit, srcs map[string][]rune, r *Result) string {
	var b strings.Builder
	at := func(o Origin) string { return o.Path + ":" + lineCol(srcs[o.Path], o.Sp.Start) }
	frames := func(fr []Frame) string {
		var s strings.Builder
		for i := 0; i < len(fr); {
			f := fr[i]
			j := i + 1
			for j < len(fr) && fr[j] == f {
				j++
			}
			fmt.Fprintf(&s, " < %s %s (at %s, def %s)", f.Kind, f.Name, at(f.Call), at(f.Def))
			if j-i > 1 {
				fmt.Fprintf(&s, " x%d", j-i)
			}
			i = j
		}
		return s.String()
	}
	for _, u := range units {
		out := r.Files[u.Path]
		fmt.Fprintf(&b, "=== %s (changed: %v)\n", u.Path, out.Changed)
		// Layout tokens only: the output has no comments of its own.
		var toks []syntax.Token
		for _, t := range u.File.Tokens {
			if !t.IsComment() {
				toks = append(toks, t)
			}
		}
		for _, d := range out.Decls {
			line := "# at " + at(d.Prov.At)
			if d.Prov.Def != nil {
				line += ", def " + at(*d.Prov.Def)
			}
			line += frames(d.Prov.Frames)
			if d.Primary {
				line += " [primary]"
			}
			b.WriteString(line + "\n")
			b.WriteString(format.File(&syntax.File{Items: []syntax.Stmt{d.D}, Src: srcs[u.Path], Tokens: toks}))
		}
		if len(out.Diags) > 0 {
			b.WriteString("--- diagnostics\n")
		}
		for _, d := range out.Diags {
			fmt.Fprintf(&b, "%s %s %s: %s\n", lineCol(srcs[u.Path], d.Sp.Start), sevName(d.Severity), d.Code, d.Message)
		}
	}
	// Copied nodes, one line per place, source line and frames: how
	// many output nodes there are, and the widest of them.
	type group struct {
		n     int
		label string
	}
	groups := map[string]*group{}
	for n, p := range r.Nodes {
		def := p.Def.Path + ":" + strings.SplitN(lineCol(srcs[p.Def.Path], p.Def.Sp.Start), ":", 2)[0]
		key := fmt.Sprintf("%s from %s%s", at(p.At), def, frames(p.Frames))
		g := groups[key]
		if g == nil {
			g = &group{}
			groups[key] = g
		}
		g.n++
		if l := nodeLabel(n); len(l) > len(g.label) || len(l) == len(g.label) && l < g.label {
			g.label = l
		}
	}
	var lines []string
	for k, g := range groups {
		lines = append(lines, fmt.Sprintf("%s: %d node%s, %s", k, g.n, map[bool]string{true: "", false: "s"}[g.n == 1], g.label))
	}
	if len(lines) > 0 {
		sort.Strings(lines)
		b.WriteString("=== provenance of copied nodes\n")
		for _, l := range lines {
			b.WriteString(l + "\n")
		}
	}
	return b.String()
}

func nodeLabel(n syntax.Node) string {
	x, ok := n.(syntax.Expr)
	if !ok {
		return fmt.Sprintf("%T", n)
	}
	s := format.Expr(x)
	if strings.Contains(s, "\n") || len(s) > 48 {
		s = strings.SplitN(s, "\n", 2)[0]
		if len(s) > 45 {
			s = s[:45]
		}
		s += "..."
	}
	return s
}

func sevName(s syntax.Severity) string {
	switch s {
	case syntax.SeverityError:
		return "error"
	case syntax.SeverityWarning:
		return "warning"
	case syntax.SeverityHint:
		return "hint"
	}
	return "info"
}

// TestSpreadReducedOnce checks that a Template call spread into an object
// is evaluated and recorded once, however deep the spreads nest.
func TestSpreadReducedOnce(t *testing.T) {
	src := "Template base(): any = { a = 1 }\n" +
		"const one = { ...base(), b = 2 }\n" +
		"const two = { ...{ ...{ ...base() } }, c = 3 }\n" +
		"Function f {\n  let x = one.a + two.a\n}\n"
	f := syntax.Parse(src)
	if len(f.Diags) > 0 {
		t.Fatal(f.Diags)
	}
	r := Run([]Unit{{Path: "main.mms", File: f}}, Options{Namespace: "demo"})
	if out := r.Files["main.mms"]; out != nil && len(out.Diags) > 0 {
		t.Fatal(out.Diags)
	}
	if len(r.Calls) != 2 {
		t.Errorf("recorded %d calls of base, want 2 (one per use)", len(r.Calls))
	}
}
