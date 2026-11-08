package project

import (
	"sort"
	"strings"

	"github.com/minecraftmetascript/mms/engine/internal/near"
	"github.com/minecraftmetascript/mms/engine/prelude"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// A Blueprint names its own lint rules: warn(field, "rule", message) in
// its check, and `as T = expr hint("rule", message)`. They are reported
// as lint/<rule>, and mms.json's "lint" sets their severity like a
// built-in rule's.

// blueprintRules lists the rule names the Blueprints in files use.
func blueprintRules(files []*syntax.File, into map[string]bool) {
	for _, f := range files {
		for _, d := range f.Decls {
			if d.Kind == nil || d.Kind.Name != "Blueprint" || d.Body == nil {
				continue
			}
			syntax.Walk(d.Body, func(n syntax.Node) bool {
				var c *syntax.Call
				switch n := n.(type) {
				case *syntax.AsStmt:
					c = n.Hint
				case *syntax.Call:
					if id, ok := n.Fn.(*syntax.Ident); ok && id.Name == "warn" && len(n.Args) == 3 {
						n := *n
						n.Args = n.Args[1:]
						c = &n
					}
				}
				if c != nil && len(c.Args) > 0 {
					if s, ok := c.Args[0].Value.(*syntax.String); ok {
						into[s.Value] = true
					}
				}
				return true
			})
		}
	}
}

// lintRules are every rule name mms.json can configure for this result.
func (r *Result) lintRules() map[string]bool {
	out := map[string]bool{}
	for _, rule := range Rules {
		out[rule.ID] = true
	}
	files := make([]*syntax.File, 0, len(r.Source))
	for _, p := range prelude.Parse() {
		files = append(files, p.File)
	}
	for _, f := range r.Source {
		files = append(files, f)
	}
	blueprintRules(files, out)
	return out
}

// configure applies mms.json's severities and `// mms-ignore` comments
// to the Blueprint check findings in path's diagnostics, and offers the
// ignore comment as a fix. Lint rules, hints included, get this as they
// run.
func (r *Result) configure(path string, ds []syntax.Diagnostic) []syntax.Diagnostic {
	u := r.Unit(path)
	if u == nil {
		return ds
	}
	var lineOf func(int) int
	var ig ignoreSet
	out := ds[:0]
	for _, d := range ds {
		id, ok := strings.CutPrefix(d.Code, "lint/")
		if !ok || RuleByID(id) != nil {
			out = append(out, d)
			continue
		}
		if s, ok := r.cfg.Lint[id]; ok {
			sev, _ := ParseSeverity(s)
			if sev == 0 {
				continue
			}
			d.Severity = sev
		}
		if lineOf == nil {
			lineOf = lineIndex(u.File.Src)
			ig = ignores(u.File, lineOf)
		}
		line := lineOf(d.Sp.Start)
		if ig.covers(d, line) {
			continue
		}
		d.Fixes = append(d.Fixes[:len(d.Fixes):len(d.Fixes)], ignoreFix(u.File.Src, lineOf, d, line))
		out = append(out, d)
	}
	return out
}

// ConfigWarnings are problems with mms.json that only analysis can see:
// lint settings for rules that neither MMS nor a Blueprint has.
func (r *Result) ConfigWarnings() []string {
	known := r.lintRules()
	var out []string
	for id := range r.cfg.Lint {
		if known[id] {
			continue
		}
		msg := "lint: unknown rule \"" + id + "\""
		if s := near.ClosestKey(id, known); s != "" {
			msg += "; did you mean \"" + s + "\"?"
		}
		out = append(out, msg)
	}
	sort.Strings(out)
	return out
}
