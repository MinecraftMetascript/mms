package ide

import (
	"sort"
	"strings"

	"github.com/minecraftmetascript/mms/engine/eval"
	"github.com/minecraftmetascript/mms/engine/format"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// Expansion is what a Template call or a Blueprint instance writes, as
// MMS source.
type Expansion struct {
	Title string      // "Template particle_ring" or "CustomItem ruby"
	Text  string      // the expanded source
	Sp    syntax.Span // the call or declaration
}

// Expand returns the expansion of the innermost Template call at off, or
// else of the Blueprint instance declared there.
func (s *Session) Expand(path string, off int) (Expansion, bool) {
	if s.R.Eval == nil {
		return Expansion{}, false
	}
	var best *eval.Call
	for _, c := range s.R.Eval.Calls {
		if c.Path != path || c.X == nil || !c.X.Sp.Contains(off) {
			continue
		}
		if best == nil || c.X.Sp.Len() < best.X.Sp.Len() {
			best = c
		}
	}
	if best != nil {
		x := Expansion{Title: "Template " + best.Template.D.NameText(), Sp: best.X.Sp}
		switch {
		case best.Expanded != nil:
			x.Text = format.Stmts(best.Expanded)
		case best.Result != nil:
			x.Text = format.Expr(best.Result)
		}
		return x, true
	}
	f := s.R.Eval.Files[path]
	if f == nil {
		return Expansion{}, false
	}
	var from *syntax.Decl
	var items []syntax.Stmt
	for _, d := range f.Decls {
		if d.Instance == nil || d.From == nil || d.D == d.From || !d.From.Sp.Contains(off) {
			continue
		}
		if from != nil && d.From != from {
			continue
		}
		from = d.From
		items = append(items, d.D)
	}
	if from == nil {
		return Expansion{}, false
	}
	return Expansion{
		Title: from.Kind.Name + " " + from.NameText(),
		Text:  strings.TrimRight(format.File(&syntax.File{Items: items}), "\n"),
		Sp:    from.Sp,
	}, true
}

// Expansions lists what the Expand lens offers in a file: each Blueprint
// instance and each Commands Template call, in order.
func (s *Session) Expansions(path string) []Expansion {
	var out []Expansion
	for _, c := range s.commandCalls(path) {
		if x, ok := s.Expand(path, c.X.Sp.Start); ok {
			out = append(out, x)
		}
	}
	if s.R.Eval != nil && s.R.Eval.Files[path] != nil {
		seen := map[*syntax.Decl]bool{}
		for _, d := range s.R.Eval.Files[path].Decls {
			if d.Instance == nil || d.From == nil || d.D == d.From || seen[d.From] {
				continue
			}
			seen[d.From] = true
			if x, ok := s.Expand(path, d.From.Kind.Sp.Start); ok {
				out = append(out, x)
			}
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Sp.Start < out[j].Sp.Start })
	return out
}
