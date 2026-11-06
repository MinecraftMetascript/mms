package sema

import (
	"regexp"
	"strings"

	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// This file is the checker's API for lint rules.

var (
	bareIdent  = regexp.MustCompile(`^[a-z_][a-z0-9_]*$`)
	barePath   = regexp.MustCompile(`^[a-z_][a-z0-9_]*(/[a-z0-9_./-]+)+$`)
	bareResLoc = regexp.MustCompile(`^[a-z0-9_.-]+:[a-z0-9_./-]+$`)
	reserved   = map[string]bool{"let": true, "if": true, "else": true, "while": true, "for": true, "in": true,
		"return": true, "break": true, "continue": true, "true": true, "false": true}
)

// BareString returns how a string literal could be written without
// quotes and mean the same, using what the checker recorded about its
// position: an id (`minecraft:stone` becomes `stone` when no declaration
// in ns would capture the short name), a dispatch key or a string enum
// name. ok is false when the quotes are needed.
func BareString(in *Info, syms *Table, ns string, s *syntax.String) (string, bool) {
	if in == nil || s.Value == "" {
		return "", false
	}
	switch in.Bare {
	case BareKey, BareEnum:
		v := strings.TrimPrefix(s.Value, "minecraft:")
		return v, bareIdent.MatchString(v) && !reserved[v]
	case BareID:
	default:
		return "", false
	}
	text := s.Value
	tag := strings.HasPrefix(text, "#")
	if tag {
		text = text[1:]
	}
	// A quoted id resolves exactly like the bare one, so only the
	// spelling matters; the minecraft: prefix is dropped when the short
	// name cannot mean a declaration instead.
	if path, ok := strings.CutPrefix(text, "minecraft:"); ok && in.Folder != "" && syms.Find(in.Folder, ns+":"+path) == nil && lexesBare(path) {
		text = path
	}
	if !lexesBare(text) {
		return "", false
	}
	if tag {
		return "#" + text, true
	}
	return text, true
}

// ShortName returns name (an id, dispatch key or enum name written bare,
// without '#') with its minecraft: prefix dropped, when the short form
// means the same at this position.
func ShortName(in *Info, syms *Table, ns, name string) (string, bool) {
	path, ok := strings.CutPrefix(name, "minecraft:")
	if in == nil || !ok || !lexesBare(path) {
		return "", false
	}
	switch in.Bare {
	case BareKey, BareEnum:
		return path, bareIdent.MatchString(path) && !reserved[path]
	case BareID:
		return path, in.Folder != "" && syms.Find(in.Folder, ns+":"+path) == nil
	}
	return "", false
}

func lexesBare(s string) bool {
	return (bareIdent.MatchString(s) && !reserved[s]) || barePath.MatchString(s) || bareResLoc.MatchString(s)
}

// Simplify returns the shorthand that means the same as an object
// literal at its position, written from the source text src: the form of
// expression whose conversion in the prelude (prelude/forms.mms) builds
// that object, such as `1..3` for a uniform provider or `a && b` for
// all_of, or a criterion as a trigger call. ok is false when there is
// none.
func Simplify(p *spec.Pack, in *Info, o *syntax.Object, src []rune) (text string, op string, ok bool) {
	if in == nil {
		return "", "", false
	}
	names := typeNames(in.Type)
	if s, op := shorthand(p, names, o, txt(src)); s != "" {
		return s, op, true
	}
	for _, n := range names {
		if n == tCriterion {
			if s := criterionShorthand(o, txt(src)); s != "" {
				return s, "", true
			}
		}
	}
	return "", "", false
}

func txt(src []rune) func(syntax.Expr) string {
	return func(e syntax.Expr) string { return string(src[e.Span().Start:e.Span().End]) }
}

// typeNames lists the named types a position accepts: its own name, or
// the names of an anonymous union's members.
func typeNames(v tview) []string {
	var out []string
	if v.Name != "" {
		out = append(out, v.Name)
	}
	if v.T != nil && v.T.Kind == "union" {
		for _, m := range v.T.Members {
			if m.Kind == "reference" {
				out = append(out, m.Path)
			}
		}
	}
	return out
}

// criterionShorthand writes `{ trigger = t, conditions = {...} }` as the
// trigger's call, `T(field = value, ...)`.
func criterionShorthand(o *syntax.Object, text func(syntax.Expr) string) string {
	fields := map[string]syntax.Expr{}
	for _, it := range o.Body.Items {
		f, ok := it.(*syntax.FieldStmt)
		if !ok || f.KeyQuoted {
			return ""
		}
		fields[f.Key.Text] = f.Value
	}
	trig, ok := fields["trigger"].(*syntax.Ident)
	if !ok || len(fields) > 2 || (len(fields) == 2 && fields["conditions"] == nil) {
		return ""
	}
	var args []string
	if cond := fields["conditions"]; cond != nil {
		co, ok := cond.(*syntax.Object)
		if !ok {
			return ""
		}
		for _, it := range co.Body.Items {
			f, ok := it.(*syntax.FieldStmt)
			if !ok || f.KeyQuoted {
				return ""
			}
			args = append(args, f.Key.Text+" = "+text(f.Value))
		}
	}
	return Pascal(trig.Name) + "(" + strings.Join(args, ", ") + ")"
}

// operand parenthesises an operator expression used as an operand of op.
func operand(op string, right bool, e syntax.Expr, text string) string {
	switch x := e.(type) {
	case *syntax.Binary:
		if NeedsParens(op, right, x.Op) {
			return "(" + text + ")"
		}
	case *syntax.Range:
		return "(" + text + ")"
	}
	return text
}

// NeedsParens reports whether an expression whose top-level operator is
// child needs parentheses as the left or right operand of parent. Binary
// operators are written as themselves, unary ones as "unary" (as a
// parent) or "!" (as a child): they bind tighter than any binary
// operator. A range ("..") always needs them.
func NeedsParens(parent string, right bool, child string) bool {
	switch child {
	case "", "!":
		return false
	case "..":
		return true
	}
	if parent == "unary" {
		return true
	}
	pc, pp := precedence[child], precedence[parent]
	if pc != pp {
		return pc < pp
	}
	// Same precedence: arithmetic is left-associative, so only the right
	// side needs them (a + (b + c) lowers to a different tree). A chain of
	// && or || flattens into one all_of or any_of, so a nested one keeps
	// its parentheses on either side.
	if parent == "&&" || parent == "||" {
		return true
	}
	return right
}

var precedence = map[string]int{"||": 1, "&&": 2, "+": 3, "-": 3, "*": 4, "/": 4}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

// AsCall returns a dispatched object literal written as its helper call,
// `{ type = noise, noise = ridge }` as `Noise(noise = ridge)`, when the call
// checks to the same value. in is the object's Info and body its body's.
func AsCall(in, body *Info, o *syntax.Object, src []rune) (string, bool) {
	if in == nil || body == nil || body.Disp == nil || (body.Disp.NoCase && len(o.Body.Items) > 1) {
		return "", false
	}
	for _, name := range typeNames(in.Type) {
		if name == tCriterion {
			return "", false // calls there are trigger sugar (criterionCall)
		}
	}
	d := body.Disp
	var key string
	var args []string
	for _, it := range o.Body.Items {
		f, ok := it.(*syntax.FieldStmt)
		if !ok || f.KeyQuoted || !bareIdent.MatchString(f.Key.Text) || reserved[f.Key.Text] {
			return "", false
		}
		if f.Key.Text == d.KeyField {
			switch x := f.Value.(type) {
			case *syntax.Ident:
				key = x.Name
			case *syntax.ResLoc:
				key = strings.TrimPrefix(x.Text, "minecraft:")
			}
			continue
		}
		args = append(args, f.Key.Text+" = "+string(src[f.Value.Span().Start:f.Value.Span().End]))
	}
	if key == "" || !bareIdent.MatchString(key) {
		return "", false
	}
	// Overlay defaults and methods change what a call means; leave those
	// objects alone. Positional order alone does not matter here.
	if ov, ok := overlays[d.Dispatcher+"/"+key]; ok && (len(ov.defaults) > 0 || len(ov.methods) > 0) {
		return "", false
	}
	name := Pascal(key)
	for alias, k := range callAliases[d.Dispatcher] {
		if k == key {
			name = alias
		}
	}
	if a, ok := callAliases[d.Dispatcher][name]; (ok && a != key) || (!ok && Snake(name) != key) {
		return "", false
	}
	return name + "(" + strings.Join(args, ", ") + ")", true
}
