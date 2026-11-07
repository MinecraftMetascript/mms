package sema

import (
	"sort"
	"strings"

	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// Conversion is one `as T = expr` of a Blueprint instance: what the
// instance's name means where a T is expected.
type Conversion struct {
	Type *spec.Type
	// X is the converted value; it is checked in place of the name.
	X syntax.Expr
	// Rule and Hint, when set, report a hint at the name.
	Rule, Hint string
	// Sig fingerprints the conversion for incremental checking.
	Sig string
	// Kind is the instance's Blueprint and Decl its declaration, for the
	// editor.
	Kind string
	Decl *syntax.Decl
	Path string
}

// instanceCandidates lists the Blueprint instances whose names convert,
// without a hint, to what v expects.
func (c *Checker) instanceCandidates(v tview) []Candidate {
	if c.Syms == nil || v.T == nil {
		return nil
	}
	var out []Candidate
	for id, cs := range c.Syms.Conversions {
		for _, cv := range cs {
			// A conversion with a hint, such as a CustomItem standing for
			// its base item, is still offered: the hint explains it if
			// the place takes it differently than meant.
			if cv.Type == nil || !c.convFits(cv.Type, v, 0) {
				continue
			}
			label := id
			if ns, path, ok := strings.Cut(id, ":"); ok && ns == c.NS {
				label = path
			}
			cand := Candidate{Label: label, Kind: CandID, Detail: cv.Kind, Rank: 0}
			if cv.Decl != nil {
				cand.Doc = DocMarkdown(cv.Decl.Doc)
				cand.Deprecated = Deprecated(cv.Decl)
			}
			out = append(out, cand)
			break
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Label < out[j].Label })
	return out
}

// ConversionsSig fingerprints the conversions of an id.
func ConversionsSig(cs []*Conversion) string {
	parts := make([]string, len(cs))
	for i, cv := range cs {
		parts[i] = cv.Sig
	}
	return strings.Join(parts, ";")
}

// conversions returns the conversions of the instance id names,
// recording the lookup.
func (c *Checker) conversions(id string) []*Conversion {
	if c.Syms == nil {
		return nil
	}
	cs := c.Syms.Conversions[id]
	if c.Deps != nil {
		c.Deps["c|"+id+"|"+ConversionsSig(cs)] = true
	}
	return cs
}

// convert checks a Blueprint instance's name where one of its `as`
// types is expected, as the converted value. ok is false when e is not
// such a name or no conversion fits v.
func (c *Checker) convert(e syntax.Expr, v tview) (jsonv.Value, bool) {
	x := c.converted(e, func(t *spec.Type) bool { return c.convFits(t, v, 0) })
	if x == nil {
		return nil, false
	}
	return c.checkConverted(c.instanceID(e), e.Span(), x, v), true
}

// mergeConvert checks an object or helper call at v one of whose fields
// names a Blueprint instance that converts to v: `{ id = ruby, count = 4 }`
// or `Item(name = ruby, modifier = [SetCount(2)])`. The converted value
// comes first and what is written goes on top of it: objects merge,
// lists append, and anything else written wins.
func (c *Checker) mergeConvert(e syntax.Expr, v tview) (jsonv.Value, bool) {
	if c.Syms == nil || c.merging[e] {
		return nil, false
	}
	var vals []syntax.Expr
	switch x := e.(type) {
	case *syntax.Object:
		for _, it := range x.Body.Items {
			if f, ok := it.(*syntax.FieldStmt); ok {
				vals = append(vals, f.Value)
			}
		}
	case *syntax.Call:
		for _, a := range x.Args {
			vals = append(vals, a.Value)
		}
	default:
		return nil, false
	}
	for _, fv := range vals {
		id := c.instanceID(fv)
		if id == "" {
			continue
		}
		for _, cv := range c.conversions(id) {
			if cv.Type == nil || !c.convFits(cv.Type, v, 0) {
				continue
			}
			base := c.checkConverted(id, fv.Span(), syntax.CloneExprAt(cv.X, fv.Span()), v)
			if c.merging == nil {
				c.merging = map[syntax.Node]bool{}
			}
			c.merging[e] = true
			top := c.check(e, v)
			delete(c.merging, e)
			// The name brings its whole value here, so the hint that it
			// stands for less does not apply.
			c.info(fv).Converted = cv
			return mergeJSON(base, top), true
		}
	}
	return nil, false
}

func mergeJSON(base, top jsonv.Value) jsonv.Value {
	switch t := top.(type) {
	case *jsonv.Object:
		b, ok := base.(*jsonv.Object)
		if !ok {
			return top
		}
		// What is written keeps its order, and the rest of the
		// converted value follows.
		out := &jsonv.Object{}
		for _, m := range t.Members {
			out.Set(m.Key, mergeJSON(b.Get(m.Key), m.Value))
		}
		for _, m := range b.Members {
			if out.Get(m.Key) == nil {
				out.Set(m.Key, m.Value)
			}
		}
		return out
	case jsonv.Array:
		if b, ok := base.(jsonv.Array); ok {
			return append(append(jsonv.Array{}, b...), t...)
		}
	}
	return top
}

// ConvertedStack returns the item stack a Blueprint instance's name
// stands for in `${...}` in a command, from its conversion to an item
// stack, or nil.
func (c *Checker) ConvertedStack(e syntax.Expr) *jsonv.Object {
	for _, path := range []string{tItemStack, "::java::world::item::ItemStack"} {
		t := &spec.Type{Kind: "reference", Path: path}
		x := c.converted(e, func(ct *spec.Type) bool { return ct.Kind == "reference" && ct.Path == path })
		if x == nil {
			continue
		}
		st, _ := c.checkConverted(c.instanceID(e), e.Span(), x, tview{T: t}).(*jsonv.Object)
		return st
	}
	return nil
}

// checkConverted checks the converted value x of instance id, which
// must not contain itself: a CustomItem in its own bundle.
func (c *Checker) checkConverted(id string, at syntax.Span, x syntax.Expr, v tview) jsonv.Value {
	if c.converting[id] {
		c.errorf("mms/conversion-cycle", at, "'%s' contains itself", id)
		return nil
	}
	if c.converting == nil {
		c.converting = map[string]bool{}
	}
	c.converting[id] = true
	defer delete(c.converting, id)
	return c.check(x, v)
}

// instanceID is the id a name expression gives, or "".
func (c *Checker) instanceID(e syntax.Expr) string {
	switch x := e.(type) {
	case *syntax.Ident:
		return c.NS + ":" + x.Name
	case *syntax.ResLoc:
		if strings.Contains(x.Text, ":") {
			return x.Text
		}
		return c.NS + ":" + x.Text
	}
	return ""
}

// convertedID is the converted value of a Blueprint instance's name
// where an id of registry reg is expected, or nil.
func (c *Checker) convertedID(e syntax.Expr, reg string) syntax.Expr {
	var isID func(t *spec.Type) bool
	isID = func(t *spec.Type) bool {
		if t.Kind == "union" {
			for _, m := range t.Members {
				if isID(m) {
					return true
				}
			}
		}
		return t.Kind == "string" && idOf(t).Registry == reg
	}
	return c.converted(e, isID)
}

// converted returns the value of the first conversion of the instance e
// names whose type fits, built at e and with its hint reported, or nil.
func (c *Checker) converted(e syntax.Expr, fits func(*spec.Type) bool) syntax.Expr {
	// No early return when the project has no conversions: the lookup
	// must be recorded, so that declaring an instance later checks e
	// again.
	if c.Syms == nil {
		return nil
	}
	id := c.instanceID(e)
	if id == "" {
		return nil
	}
	for _, cv := range c.conversions(id) {
		if cv.Type == nil || !fits(cv.Type) {
			continue
		}
		// The hint, if any, is a lint rule's finding: see Info.Converted.
		c.info(e).Converted = cv
		return syntax.CloneExprAt(cv.X, e.Span())
	}
	return nil
}

// convFits reports whether a value of conversion type t is what v expects:
// the same schema or project type, an id of the same registry, or a
// member of a union v allows.
func (c *Checker) convFits(t *spec.Type, v tview, depth int) bool {
	if t == nil || v.T == nil || depth > 8 {
		return false
	}
	switch t.Kind {
	case "union":
		for _, m := range t.Members {
			if c.convFits(m, v, depth+1) {
				return true
			}
		}
		return false
	case "reference":
		if v.Name == t.Path {
			return true
		}
	case "string":
		if v.T.Kind == "string" {
			want, got := idOf(t), idOf(v.T)
			if want.Present || got.Present {
				return want.Present && got.Present && want.Registry == got.Registry
			}
			return true
		}
	case "boolean", "byte", "short", "int", "long", "float", "double":
		if v.T.Kind == t.Kind {
			return true
		}
	}
	if v.T.Kind == "union" {
		for _, m := range v.T.Members {
			if c.convFits(t, c.norm(tview{T: m, Env: v.Env}), depth+1) {
				return true
			}
		}
	}
	return false
}

// StackCandidates lists the Blueprint instances whose names stand for an
// item stack, for `${...}` in a command.
func (c *Checker) StackCandidates() []Candidate {
	var out []Candidate
	for _, path := range []string{tItemStack, "::java::world::item::ItemStack"} {
		out = append(out, c.instanceCandidates(tview{T: &spec.Type{Kind: "reference", Path: path}, Name: path})...)
	}
	return dedupe(out)
}
