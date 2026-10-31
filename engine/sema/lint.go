package sema

import (
	"regexp"
	"strings"

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

// Simplify returns sugar that means the same as an object literal at its
// position (see sugar.go), written from the source text src: an item entry
// as its bare id, a uniform provider or bounds as a range, a stack as
// `id * count`, a criterion as a trigger call and all_of, any_of and
// inverted conditions as &&, || and !. ok is false when there is none.
func Simplify(in *Info, o *syntax.Object, src []rune) (text string, op string, ok bool) {
	text, op = simplify(in, o, src)
	return text, op, text != ""
}

// simplify returns the sugar and its top-level operator ("" for none).
func simplify(in *Info, o *syntax.Object, src []rune) (string, string) {
	if in == nil {
		return "", ""
	}
	fields := map[string]syntax.Expr{}
	var order []string
	for _, it := range o.Body.Items {
		f, ok := it.(*syntax.FieldStmt)
		if !ok || f.KeyQuoted {
			return "", ""
		}
		fields[f.Key.Text] = f.Value
		order = append(order, f.Key.Text)
	}
	text := func(e syntax.Expr) string { return string(src[e.Span().Start:e.Span().End]) }
	is := func(e syntax.Expr, name string) bool {
		switch x := e.(type) {
		case *syntax.Ident:
			return x.Name == name
		case *syntax.ResLoc:
			return x.Text == "minecraft:"+name
		case *syntax.String:
			return x.Value == name || x.Value == "minecraft:"+name
		}
		return false
	}
	only := func(keys ...string) bool {
		if len(keys) != len(fields) {
			return false
		}
		for _, k := range keys {
			if fields[k] == nil {
				return false
			}
		}
		return true
	}
	bound := func(e syntax.Expr) bool { return e == nil || isNumber(e) }
	for _, name := range typeNames(in.Type) {
		if text, op := simplifyAs(name, fields, only, is, bound, text); text != "" {
			return text, op
		}
	}
	return "", ""
}

// typeNames lists the named types a position accepts: its own name, or
// the names of an anonymous union's members.
func typeNames(v tview) []string {
	if v.Name != "" {
		return []string{v.Name}
	}
	var out []string
	if v.T != nil && v.T.Kind == "union" {
		for _, m := range v.T.Members {
			if m.Kind == "reference" {
				out = append(out, m.Path)
			}
		}
	}
	return out
}

func simplifyAs(name string, fields map[string]syntax.Expr, only func(...string) bool, is func(syntax.Expr, string) bool,
	bound func(syntax.Expr) bool, text func(syntax.Expr) string) (string, string) {
	switch name {
	case tLootEntry:
		if only("type", "name") && is(fields["type"], "item") && isID(fields["name"]) {
			return text(fields["name"]), ""
		}
	case tIntProvider, tFloatProvider:
		if only("type", "min", "max") && is(fields["type"], "uniform") && isNumber(fields["min"]) && isNumber(fields["max"]) {
			return text(fields["min"]) + ".." + text(fields["max"]), ".."
		}
	case tIntRange, tFloatRange, tMinMax:
		lo, hi := fields["min"], fields["max"]
		if len(fields) > 0 && len(fields) == btoi(lo != nil)+btoi(hi != nil) && bound(lo) && bound(hi) {
			s := ".."
			if lo != nil {
				s = text(lo) + s
			}
			if hi != nil {
				s += text(hi)
			}
			return s, ".."
		}
	case tItemStack:
		if only("id", "count") && isID(fields["id"]) && isNumber(fields["count"]) {
			return text(fields["id"]) + " * " + text(fields["count"]), "*"
		}
	case tCriterion:
		trig, ok := fields["trigger"].(*syntax.Ident)
		if !ok || (len(fields) != 1 && !only("trigger", "conditions")) {
			return "", ""
		}
		var args []string
		if cond := fields["conditions"]; cond != nil {
			co, ok := cond.(*syntax.Object)
			if !ok {
				return "", ""
			}
			for _, it := range co.Body.Items {
				f, ok := it.(*syntax.FieldStmt)
				if !ok || f.KeyQuoted {
					return "", ""
				}
				args = append(args, f.Key.Text+" = "+text(f.Value))
			}
		}
		return Pascal(trig.Name) + "(" + strings.Join(args, ", ") + ")", ""
	case tDensity, tDensityRef:
		if !only("type", "left", "right") {
			return "", ""
		}
		l, r := fields["left"], fields["right"]
		if isNumber(l) && isNumber(r) {
			return "", "" // `2 * 3` would fold to a constant
		}
		for op, name := range densityOps {
			if is(fields["type"], strings.ToLower(name)) {
				return operand(op, false, l, text(l)) + " " + op + " " + operand(op, true, r, text(r)), op
			}
		}
	case tLootCondition, tLootConditionRef, tPredicateRef, tPredicate:
		switch {
		case only("type", "terms") && (is(fields["type"], "all_of") || is(fields["type"], "any_of")):
			l, ok := fields["terms"].(*syntax.List)
			if !ok || len(l.Elems) < 2 {
				return "", ""
			}
			op := "&&"
			if is(fields["type"], "any_of") {
				op = "||"
			}
			parts := make([]string, len(l.Elems))
			for i, t := range l.Elems {
				parts[i] = operand(op, i > 0, t, text(t))
			}
			return strings.Join(parts, " "+op+" "), op
		case only("type", "term") && is(fields["type"], "inverted"):
			t := fields["term"]
			if _, ok := t.(*syntax.Binary); ok {
				return "!(" + text(t) + ")", "!"
			}
			return "!" + text(t), "!"
		}
	}
	return "", ""
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
