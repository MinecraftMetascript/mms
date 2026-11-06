package sema

import (
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
	if c.Syms == nil || c.Syms.Conversions == nil {
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
	return c.check(x, v), true
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
	if c.Syms == nil || len(c.Syms.Conversions) == 0 {
		return nil
	}
	var id string
	switch x := e.(type) {
	case *syntax.Ident:
		id = c.NS + ":" + x.Name
	case *syntax.ResLoc:
		id = x.Text
		if !strings.Contains(id, ":") {
			id = c.NS + ":" + id
		}
	default:
		return nil
	}
	for _, cv := range c.conversions(id) {
		if cv.Type == nil || !fits(cv.Type) {
			continue
		}
		if cv.Hint != "" && !c.trial {
			c.Diags = append(c.Diags, syntax.Diagnostic{Code: cv.Rule, Severity: syntax.SeverityHint, Message: cv.Hint, Sp: e.Span()})
		}
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
