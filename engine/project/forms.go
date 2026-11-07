package project

import (
	"strconv"
	"strings"

	"github.com/minecraftmetascript/mms/engine/prelude"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// valueKinds are the declarations whose `= value` is a value, not the
// resource's body: a Template's result, a Global's or Storage's start
// value and the like. Every other kind writes an object body as
// `Kind name { ... }`.
var valueKinds = map[string]bool{
	"Template": true, "Type": true, "Blueprint": true,
	"Global": true, "Storage": true, "Test": true,
	// A Function's `= { }` has its own error, which asks for run { }.
	"Function": true,
}

// checkForms reports the declaration forms the language has dropped, in a
// file as written: `Kind name = { ... }`, which is `Kind name { ... }`,
// and `Objective name = criterion`, which is `Objective name: criterion`.
// Each comes with a safe fix, so `mms fix` rewrites old code.
func checkForms(path string, f *syntax.File) []syntax.Diagnostic {
	if f == nil || prelude.Is(path) {
		return nil
	}
	var out []syntax.Diagnostic
	syntax.Walk(f, func(n syntax.Node) bool {
		d, ok := n.(*syntax.Decl)
		if !ok || d.Kind == nil || d.Value == nil || d.Body != nil {
			return true
		}
		switch {
		case d.Kind.Name == "Objective":
			out = append(out, objectiveForm(f.Src, d))
		case !valueKinds[d.Kind.Name]:
			if _, ok := d.Value.(*syntax.Object); ok {
				if dg, ok := bodyForm(f.Src, d); ok {
					out = append(out, dg)
				}
			}
		}
		return true
	})
	return out
}

// head returns where the part of d before its `=` ends.
func head(d *syntax.Decl) int {
	end := d.Kind.Sp.End
	for _, n := range d.Names() {
		end = max(end, n.Sp.End)
	}
	if d.Type != nil {
		end = max(end, d.Type.Span().End)
	}
	if d.Params != nil {
		end = max(end, d.Params.Sp.End)
	}
	return end
}

// equals returns the span of d's `=`, or false when something other
// than space sits between the head and the value.
func equals(src []rune, d *syntax.Decl) (syntax.Span, bool) {
	from, to := head(d), d.Value.Span().Start
	if from > to || to > len(src) {
		return syntax.Span{}, false
	}
	between := string(src[from:to])
	if strings.TrimSpace(between) != "=" {
		return syntax.Span{}, false
	}
	at := from + len([]rune(between[:strings.Index(between, "=")]))
	return syntax.Span{Start: at, End: at + 1}, true
}

func bodyForm(src []rune, d *syntax.Decl) (syntax.Diagnostic, bool) {
	eq, ok := equals(src, d)
	if !ok {
		return syntax.Diagnostic{}, false
	}
	name := d.Kind.Name + " " + string(src[d.Name.Sp.Start:d.Name.Sp.End])
	return syntax.Diagnostic{
		Code: "mms/syntax", Severity: syntax.SeverityError, Sp: eq,
		Message: "a body is written without '=': " + name + " { ... }",
		Fixes: []syntax.Fix{{Title: "Remove '='", Safe: true, Edits: []syntax.TextEdit{
			{Sp: syntax.Span{Start: head(d), End: d.Value.Span().Start}, Text: " "},
		}}},
	}, true
}

func objectiveForm(src []rune, d *syntax.Decl) syntax.Diagnostic {
	dg := syntax.Diagnostic{
		Code: "mms/syntax", Severity: syntax.SeverityError, Sp: d.Value.Span(),
		Message: "an objective's criterion is written after ':', as in Objective kills: deathCount",
	}
	if eq, ok := equals(src, d); ok {
		dg.Sp = syntax.Span{Start: eq.Start, End: d.Value.Span().End}
	}
	crit := ""
	switch v := d.Value.(type) {
	case *syntax.Ident:
		crit = v.Name
	case *syntax.ResLoc:
		crit = v.Text
	case *syntax.String:
		if criterionRe(v.Value) {
			crit = v.Value
		}
	}
	if crit != "" && d.Type == nil {
		if _, ok := equals(src, d); ok {
			dg.Fixes = []syntax.Fix{{Title: "Write Objective " + d.NameText() + ": " + crit, Safe: true, Edits: []syntax.TextEdit{
				{Sp: syntax.Span{Start: head(d), End: d.Value.Span().End}, Text: ": " + crit},
			}}}
		}
	}
	return dg
}

// criterionRe reports whether s can be written as a bare criterion:
// words joined by '.' and ':', as in minecraft.killed:minecraft.zombie.
func criterionRe(s string) bool {
	if s == "" {
		return false
	}
	for _, part := range strings.FieldsFunc(s, func(r rune) bool { return r == '.' || r == ':' }) {
		if part == "" {
			return false
		}
		for _, r := range part {
			if !(r == '_' || r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9') {
				return false
			}
		}
		if _, err := strconv.Atoi(part[:1]); err == nil {
			return false
		}
	}
	return !strings.HasPrefix(s, ".") && !strings.HasSuffix(s, ".") && !strings.Contains(s, "..") && !strings.HasPrefix(s, ":") && !strings.HasSuffix(s, ":")
}
