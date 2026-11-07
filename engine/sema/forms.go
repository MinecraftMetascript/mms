package sema

import (
	"strconv"
	"strings"
	"sync"

	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/prelude"
	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// The built-in forms of expression, such as `1..3`, `!a` and `a + b`,
// mean what the prelude's `extern Type` declarations for them say at each
// type (prelude/forms.mms). formOf names an expression's form and gives
// its parts; formConvert checks the form's conversion for the expected
// type in its place.

// formOf returns the name of e's form and its parts, by the field names
// the prelude gives them, or "".
func formOf(e syntax.Expr) (string, map[string]syntax.Expr) {
	switch x := e.(type) {
	case *syntax.Range:
		return "Range", map[string]syntax.Expr{"lo": x.Lo, "hi": x.Hi}
	case *syntax.Unary:
		switch {
		case x.Op == "!":
			return "Not", map[string]syntax.Expr{"operand": x.X}
		case x.Op == "-" && !isNumber(x):
			return "Negation", map[string]syntax.Expr{"operand": x.X}
		case isNumber(x):
			// `-128` is a number, such as a height below zero.
			return "NumberLiteral", map[string]syntax.Expr{"value": e}
		}
	case *syntax.Binary:
		switch x.Op {
		case "&&", "||":
			form := "And"
			if x.Op == "||" {
				form = "Or"
			}
			return form, map[string]syntax.Expr{"terms": &syntax.List{Elems: chain(x), Sp: x.Sp}}
		case "+", "-", "*", "/":
			form := map[string]string{"+": "Sum", "-": "Difference", "*": "Product", "/": "Quotient"}[x.Op]
			return form, map[string]syntax.Expr{"left": x.L, "right": x.R}
		}
	case *syntax.Ident, *syntax.ResLoc:
		return "BareId", map[string]syntax.Expr{"id": e}
	case *syntax.Tag:
		return "TagRef", map[string]syntax.Expr{"tag": e}
	case *syntax.Number:
		return "NumberLiteral", map[string]syntax.Expr{"value": e}
	case *syntax.List:
		return "ListLiteral", map[string]syntax.Expr{"items": flatList(x)}
	}
	return "", nil
}

// chain lists the terms of a chain of one operator: `a && b && c` has
// three. Parentheses keep a group: `(a && b) && c` has two.
func chain(b *syntax.Binary) []syntax.Expr {
	var terms []syntax.Expr
	var add func(x syntax.Expr)
	add = func(x syntax.Expr) {
		if inner, ok := x.(*syntax.Binary); ok && inner.Op == b.Op {
			add(inner.L)
			add(inner.R)
			return
		}
		terms = append(terms, x)
	}
	add(b)
	return terms
}

// flatList flattens lists inside l into it.
func flatList(l *syntax.List) *syntax.List {
	flat := &syntax.List{Sp: l.Sp}
	var add func(l *syntax.List)
	add = func(l *syntax.List) {
		for _, x := range l.Elems {
			if inner, ok := x.(*syntax.List); ok {
				add(inner)
				continue
			}
			flat.Elems = append(flat.Elems, x)
		}
	}
	add(l)
	return flat
}

// formConversion returns the conversion of e's form for v, and the
// converted value built at e, or nil.
func (c *Checker) formConversion(e syntax.Expr, v tview) syntax.Expr {
	if v.Name == "" || c.Pack == nil {
		return nil
	}
	form, parts := formOf(e)
	if form == "" {
		return nil
	}
	for _, cv := range Forms(c.Pack)[form] {
		if !c.formFits(cv, v) {
			continue
		}
		return substitute(cv.X, parts, e.Span())
	}
	return nil
}

// formFits reports whether a form conversion is for the type v names:
// the same type, or one that is an alias of the other.
func (c *Checker) formFits(cv *Conversion, v tview) bool {
	return cv.Type != nil && cv.Type.Kind == "reference" && sameType(c.Pack, cv.Type.Path, v.Name)
}

// sameType reports whether two type paths name the same type, following
// aliases (references) from either side.
func sameType(p *spec.Pack, a, b string) bool {
	if a == "" || b == "" {
		return false
	}
	seen := map[string]bool{}
	for path, i := a, 0; path != "" && i < 8; i++ {
		seen[path] = true
		path = aliasOf(p, path)
	}
	for path, i := b, 0; path != "" && i < 8; i++ {
		if seen[path] {
			return true
		}
		path = aliasOf(p, path)
	}
	return false
}

func aliasOf(p *spec.Pack, path string) string {
	if p == nil {
		return ""
	}
	if t := p.Types[path]; t != nil && t.Kind == "reference" {
		return t.Path
	}
	return ""
}

// formConvert checks e as its form's conversion for v, if it has one.
func (c *Checker) formConvert(e syntax.Expr, v tview) (jsonv.Value, bool) {
	x := c.formConversion(e, v)
	if x == nil {
		return nil, false
	}
	// Arithmetic on numbers alone is the number it comes to.
	if b, ok := e.(*syntax.Binary); ok {
		if f, ok := fold(b); ok {
			return c.check(&syntax.Number{Text: strconv.FormatFloat(f, 'g', -1, 64), Sp: b.Sp}, v), true
		}
	}
	return c.check(x, v), true
}

// substitute builds a form conversion's value at sp, with each
// `self.part` replaced by the part as written. A field whose part was
// left out is dropped.
func substitute(x syntax.Expr, parts map[string]syntax.Expr, sp syntax.Span) syntax.Expr {
	part := func(e syntax.Expr) (syntax.Expr, bool, bool) {
		m, ok := e.(*syntax.Member)
		if !ok || m.Name == nil {
			return nil, false, false
		}
		if id, ok := m.X.(*syntax.Ident); !ok || id.Name != "self" {
			return nil, false, false
		}
		p, ok := parts[m.Name.Text]
		return p, ok && p != nil, true
	}
	if o, ok := x.(*syntax.Object); ok && o.Body != nil {
		b := &syntax.Block{Closed: true, Sp: sp}
		for _, it := range o.Body.Items {
			if f, ok := it.(*syntax.FieldStmt); ok {
				if _, present, isPart := part(f.Value); isPart && !present {
					continue
				}
			}
			b.Items = append(b.Items, syntax.CloneStmtWith(it, sp, func(e syntax.Expr) (syntax.Expr, bool) {
				p, present, _ := part(e)
				return p, present
			}))
		}
		return &syntax.Object{Body: b, Sp: sp}
	}
	return syntax.CloneExprWith(x, sp, func(e syntax.Expr) (syntax.Expr, bool) {
		p, present, _ := part(e)
		return p, present
	})
}

// fold evaluates an arithmetic expression made only of numbers.
func fold(e syntax.Expr) (float64, bool) {
	switch x := e.(type) {
	case *syntax.Paren:
		return fold(x.X)
	case *syntax.Number, *syntax.Unary:
		if t, ok := numberText(e); ok {
			f, err := strconv.ParseFloat(t, 64)
			return f, err == nil
		}
	case *syntax.Binary:
		l, ok1 := fold(x.L)
		r, ok2 := fold(x.R)
		if !ok1 || !ok2 {
			return 0, false
		}
		switch x.Op {
		case "+":
			return l + r, true
		case "-":
			return l - r, true
		case "*":
			return l * r, true
		case "/":
			if r != 0 {
				return l / r, true
			}
		}
	}
	return 0, false
}

// shorthand is the reverse of formConvert, for the use-shorthand lint: the
// form an object written at a type could be written as, and the form's
// top-level operator ("" for none). The object must hold exactly the
// fields of a form conversion for its type, with the same words where the
// conversion writes one and a fitting part where it reads `self`.
func shorthand(p *spec.Pack, names []string, o *syntax.Object, text func(syntax.Expr) string) (string, string) {
	if p == nil {
		return "", ""
	}
	fields := map[string]syntax.Expr{}
	for _, it := range o.Body.Items {
		f, ok := it.(*syntax.FieldStmt)
		if !ok || f.KeyQuoted {
			return "", ""
		}
		fields[f.Key.Text] = f.Value
	}
	names = unionNames(p, names)
	for _, form := range formOrder {
		for _, cv := range Forms(p)[form] {
			if cv.Type == nil || cv.Type.Kind != "reference" {
				continue
			}
			fits := false
			for _, n := range names {
				fits = fits || sameType(p, cv.Type.Path, n)
			}
			if !fits {
				continue
			}
			parts, ok := matchForm(cv.X, fields)
			if !ok {
				continue
			}
			if s, op := printForm(form, parts, text); s != "" {
				return s, op
			}
		}
	}
	return "", ""
}

// formOrder is the order shorthands are tried in.
var formOrder = []string{"Range", "Not", "And", "Or", "Sum", "Difference", "Product", "Quotient", "Negation", "BareId", "TagRef", "NumberLiteral", "ListLiteral"}

// matchForm matches the written fields against a conversion's object,
// returning the parts it reads.
func matchForm(x syntax.Expr, fields map[string]syntax.Expr) (map[string]syntax.Expr, bool) {
	o, ok := x.(*syntax.Object)
	if !ok || o.Body == nil {
		return nil, false
	}
	parts := map[string]syntax.Expr{}
	used := 0
	for _, it := range o.Body.Items {
		f, ok := it.(*syntax.FieldStmt)
		if !ok || f.Key == nil {
			return nil, false
		}
		got := fields[f.Key.Text]
		if m, ok := f.Value.(*syntax.Member); ok && m.Name != nil {
			if id, ok := m.X.(*syntax.Ident); ok && id.Name == "self" {
				if got != nil {
					parts[m.Name.Text] = got
					used++
				}
				continue
			}
		}
		if got == nil || !sameWord(f.Value, got) {
			return nil, false
		}
		used++
	}
	return parts, used == len(fields) && used > 0
}

// sameWord reports whether a written value is the word or literal a
// conversion writes: `item`, "item" or "minecraft:item" for `item`.
func sameWord(want, got syntax.Expr) bool {
	w, ok := literalText(want)
	if !ok {
		return false
	}
	g, ok := literalText(got)
	return ok && (g == w || g == "minecraft:"+w)
}

func literalText(e syntax.Expr) (string, bool) {
	switch x := e.(type) {
	case *syntax.Ident:
		return x.Name, true
	case *syntax.ResLoc:
		return x.Text, true
	case *syntax.String:
		return x.Value, true
	case *syntax.Bool:
		return strconv.FormatBool(x.Value), true
	case *syntax.Number, *syntax.Unary:
		return numberText(e)
	}
	return "", false
}

// printForm writes a form from its parts, or "" when the parts would not
// read back as that form.
func printForm(form string, parts map[string]syntax.Expr, text func(syntax.Expr) string) (string, string) {
	bin := map[string]string{"Sum": "+", "Difference": "-", "Product": "*", "Quotient": "/"}
	switch form {
	case "Range":
		lo, hi := parts["lo"], parts["hi"]
		if (lo == nil && hi == nil) || (lo != nil && !isNumber(lo)) || (hi != nil && !isNumber(hi)) {
			return "", ""
		}
		s := ".."
		if lo != nil {
			s = text(lo) + s
		}
		if hi != nil {
			s += text(hi)
		}
		return s, ".."
	case "Not":
		t := parts["operand"]
		if t == nil {
			return "", ""
		}
		switch t.(type) {
		case *syntax.Binary, *syntax.Range:
			return "!(" + text(t) + ")", "!"
		}
		return "!" + text(t), "!"
	case "And", "Or":
		l, ok := parts["terms"].(*syntax.List)
		if !ok || len(l.Elems) < 2 {
			return "", ""
		}
		op := "&&"
		if form == "Or" {
			op = "||"
		}
		out := make([]string, len(l.Elems))
		for i, t := range l.Elems {
			out[i] = operand(op, i > 0, t, text(t))
		}
		return strings.Join(out, " "+op+" "), op
	case "Sum", "Difference", "Product", "Quotient":
		l, r := parts["left"], parts["right"]
		if l == nil || r == nil || isNumber(l) && isNumber(r) {
			return "", "" // `2 * 3` would fold to a constant
		}
		op := bin[form]
		return operand(op, false, l, text(l)) + " " + op + " " + operand(op, true, r, text(r)), op
	case "Negation":
		t := parts["operand"]
		if t == nil || isNumber(t) {
			return "", ""
		}
		if _, ok := t.(*syntax.Binary); ok {
			return "-(" + text(t) + ")", "unary"
		}
		return "-" + text(t), "unary"
	case "BareId":
		if id := parts["id"]; id != nil && isID(id) {
			return text(id), ""
		}
	}
	return "", ""
}

var (
	formsMu sync.Mutex
	forms   = map[*spec.Pack]map[string][]*Conversion{}
)

// Forms returns the conversions of the built-in forms, by form name, from
// the prelude's `extern Type` declarations, with their types resolved in
// pack p. The prelude names only vanilla types there, by their short
// names.
func Forms(p *spec.Pack) map[string][]*Conversion {
	formsMu.Lock()
	defer formsMu.Unlock()
	if m, ok := forms[p]; ok {
		return m
	}
	short := map[string]string{}
	for path := range p.Types {
		i := strings.LastIndex(path, "::")
		name := path[i+2:]
		if old, ok := short[name]; !ok || len(path) < len(old) || len(path) == len(old) && path < old {
			short[name] = path
		}
	}
	m := map[string][]*Conversion{}
	for _, f := range prelude.Parse() {
		for _, d := range f.File.Decls {
			if !d.Extern || d.Kind == nil || d.Kind.Name != "Type" || d.Body == nil {
				continue
			}
			for _, it := range d.Body.Items {
				as, ok := it.(*syntax.AsStmt)
				if !ok || as.Value == nil {
					continue
				}
				n, ok := as.Type.(*syntax.TypeName)
				if !ok || short[n.Name] == "" {
					continue
				}
				m[d.NameText()] = append(m[d.NameText()], &Conversion{
					Type: &spec.Type{Kind: "reference", Path: short[n.Name]}, X: as.Value,
					Kind: d.NameText(), Decl: d, Path: f.Path,
				})
			}
		}
	}
	forms[p] = m
	return m
}

// unionNames adds to names the named members of the unions they name,
// such as the predicate an id-or-predicate position also takes.
func unionNames(p *spec.Pack, names []string) []string {
	seen := map[string]bool{}
	var out []string
	var add func(name string, depth int)
	add = func(name string, depth int) {
		if seen[name] || depth > 4 {
			return
		}
		seen[name] = true
		out = append(out, name)
		if t := p.Types[name]; t != nil && t.Kind == "union" {
			for _, m := range t.Members {
				if m.Kind == "reference" {
					add(m.Path, depth+1)
				}
			}
		}
	}
	for _, n := range names {
		add(n, 0)
	}
	return out
}
