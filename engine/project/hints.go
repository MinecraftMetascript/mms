package project

import (
	"sort"

	"github.com/minecraftmetascript/mms/engine/eval"
	"github.com/minecraftmetascript/mms/engine/format"
	"github.com/minecraftmetascript/mms/engine/prelude"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// A Blueprint's `as T = expr hint("rule", "message")` reports where an
// instance's name is checked as that conversion. Each hint in the
// prelude is a lint rule of its own, documented by the `as` doc comment;
// a project's own hints are reported under their rule name as hints.

// hintRules are the rules made from the prelude's hints.
var hintRules = map[string]bool{}

func isHintRule(r *Rule) bool { return hintRules[r.ID] }

// emittedDecl reports whether a Blueprint instance emitted d.
func (r *Result) emittedDecl(d *syntax.Decl) bool {
	if r.Eval == nil {
		return false
	}
	if r.emitted == nil {
		r.emitted = map[*syntax.Decl]bool{}
		for _, f := range r.Eval.Files {
			for _, od := range f.Decls {
				if od.Instance != nil {
					r.emitted[od.D] = true
				}
			}
		}
	}
	return r.emitted[d]
}

// preludeEmitted are the declarations the prelude's Blueprints emitted.
func preludeEmitted(ev *eval.Result) map[*syntax.Decl]bool {
	out := map[*syntax.Decl]bool{}
	written := map[*syntax.Decl]map[string]bool{}
	for _, f := range ev.Files {
		for _, od := range f.Decls {
			if od.Instance == nil || !prelude.Is(od.Instance.Blueprint.Path) {
				continue
			}
			// A name the user wrote in the instance, such as a JigsawStructure's
			// pool `halls`, is one they refer to, so it stays offered.
			names, ok := written[od.From]
			if !ok {
				names = writtenNames(od.From)
				written[od.From] = names
			}
			if od.D.Name == nil || !names[od.D.Name.Text] {
				out[od.D] = true
			}
		}
	}
	return out
}

// writtenNames lists the names in a Blueprint instance as written: its own
// name and the keys of its body.
func writtenNames(d *syntax.Decl) map[string]bool {
	out := map[string]bool{}
	if d == nil {
		return out
	}
	for _, n := range d.Names() {
		out[n.Text] = true
	}
	if d.Body != nil {
		syntax.Walk(d.Body, func(n syntax.Node) bool {
			if f, ok := n.(*syntax.FieldStmt); ok && f.Key != nil {
				out[f.Key.Text] = true
			}
			return true
		})
	}
	return out
}

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
				hintRules[id] = true
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
			sev := syntax.SeverityHint
			if s, ok := l.r.cfg.Lint[cv.Rule]; ok {
				sev, _ = ParseSeverity(s)
			}
			if sev != 0 {
				*l.cur = append(*l.cur, syntax.Diagnostic{Code: "lint/" + cv.Rule, Severity: sev,
					Message: cv.Hint, Sp: n.Span(), Fixes: []syntax.Fix{fix}})
			}
			return true
		}
		l.report(d, n.Span(), cv.Hint, fix)
		return true
	})
}

func sortRules() {
	sort.Slice(Rules, func(i, j int) bool { return Rules[i].ID < Rules[j].ID })
}
