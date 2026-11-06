package project

import (
	"sort"

	"github.com/minecraftmetascript/mms/engine/format"
	"github.com/minecraftmetascript/mms/engine/prelude"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// A Blueprint's `as T = expr hint("rule", "message")` reports where an
// instance's name is checked as that conversion. Each hint in the
// prelude is a lint rule of its own, documented by the `as` doc comment;
// a project's own hints are reported under their rule name as hints.

func init() {
	for _, f := range prelude.Parse() {
		for _, d := range f.File.Decls {
			if d.Kind == nil || d.Kind.Name != "Blueprint" || d.Body == nil {
				continue
			}
			for _, it := range d.Body.Items {
				as, ok := it.(*syntax.AsStmt)
				if !ok || as.Hint == nil || len(as.Hint.Args) == 0 {
					continue
				}
				s, ok := as.Hint.Args[0].Value.(*syntax.String)
				if !ok || RuleByID(s.Value) != nil {
					continue
				}
				id := s.Value
				Rules = append(Rules, &Rule{ID: id, Severity: syntax.SeverityHint, Doc: as.Doc,
					run: func(l *lintRun, d *Decl) { lintHints(l, d, id) }})
			}
		}
	}
	sortRules()
}

// lintHints reports the uses of one hint rule, or with rule "" those of
// rules no lint rule knows: a project's own blueprints.
func lintHints(l *lintRun, d *Decl, rule string) {
	syntax.Walk(d.D, func(n syntax.Node) bool {
		in := d.Info[n]
		if in == nil || in.Converted == nil || in.Converted.Hint == "" {
			return true
		}
		cv := in.Converted
		if rule == "" && RuleByID(cv.Rule) != nil || rule != "" && cv.Rule != rule {
			return true
		}
		to := format.Expr(cv.X)
		fix := syntax.Fix{Title: "Replace with " + to, Edits: []syntax.TextEdit{{Sp: n.Span(), Text: to}}}
		if rule == "" {
			*l.cur = append(*l.cur, syntax.Diagnostic{Code: "lint/" + cv.Rule, Severity: syntax.SeverityHint,
				Message: cv.Hint, Sp: n.Span(), Fixes: []syntax.Fix{fix}})
			return true
		}
		l.report(d, n.Span(), cv.Hint, fix)
		return true
	})
}

func sortRules() {
	sort.Slice(Rules, func(i, j int) bool { return Rules[i].ID < Rules[j].ID })
}
