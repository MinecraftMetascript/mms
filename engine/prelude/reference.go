package prelude

import (
	"strings"

	"github.com/minecraftmetascript/mms/engine/format"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// RefDecl is one Template, Blueprint or Type of the prelude, for the
// language reference.
type RefDecl struct {
	// Kind is Template, Blueprint or Type.
	Kind string `json:"kind"`
	Name string `json:"name"`
	// Signature is how it is written: `tell(target: Selector, message: Text): Commands`
	// for a Template, `LangText = string | {string: string}` for a Type alias.
	Signature string `json:"signature"`
	Doc       string `json:"doc,omitempty"`
	// Fields are a Blueprint's or a Type's fields, or a Template's
	// parameters.
	Fields []RefField `json:"fields,omitempty"`
	// As is the type a Blueprint instance's name stands for, as in
	// `as Id<item>`.
	As string `json:"as,omitempty"`
	// Uses are every `as` of a Blueprint or Type, in order: what its
	// name, or for a built-in form the expression, means where each type
	// is expected.
	Uses []RefAs `json:"uses,omitempty"`
	// Builtin is set for an `extern` declaration, whose meaning the
	// compiler supplies.
	Builtin bool `json:"builtin,omitempty"`
}

// RefAs is one `as T = value`: the type, and its doc comment.
type RefAs struct {
	Type string `json:"type"`
	Doc  string `json:"doc,omitempty"`
}

// RefField is a field or parameter.
type RefField struct {
	Name     string `json:"name"`
	Type     string `json:"type,omitempty"`
	Default  string `json:"default,omitempty"`
	Optional bool   `json:"optional,omitempty"`
	Doc      string `json:"doc,omitempty"`
}

// Reference lists the prelude's Templates, Blueprints and Types, in file
// order. Blueprints the compiler builds itself (Function, Objective and
// the like) are listed as kinds elsewhere and left out here, and so are
// declarations marked @internal.
func Reference() []RefDecl {
	var out []RefDecl
	for _, f := range Parse() {
		for _, it := range f.File.Items {
			d, ok := it.(*syntax.Decl)
			if !ok || d.Kind == nil || d.Name == nil {
				continue
			}
			doc := syntax.ParseDoc(d.Doc)
			if internal(doc) {
				continue
			}
			r := RefDecl{Kind: d.Kind.Name, Name: d.Name.Text, Doc: doc.Text, Builtin: d.Extern}
			switch d.Kind.Name {
			case "Template":
				var params []string
				if d.Params != nil {
					for _, p := range d.Params.Params {
						if p.Name == nil {
							continue
						}
						fd := RefField{Name: p.Name.Text, Doc: doc.Param(p.Name.Text)}
						text := p.Name.Text
						if p.Type != nil {
							fd.Type = syntax.TypeText(p.Type)
							fd.Optional = strings.HasSuffix(fd.Type, "?")
							text += ": " + fd.Type
						}
						if p.Default != nil {
							fd.Default = format.Expr(p.Default)
							fd.Optional = true
							text += " = " + fd.Default
						}
						params = append(params, text)
						r.Fields = append(r.Fields, fd)
					}
				}
				r.Signature = d.Name.Text + "(" + strings.Join(params, ", ") + ")"
				if d.Type != nil {
					r.Signature += ": " + syntax.TypeText(d.Type)
				}
			case "Blueprint":
				if d.Extern {
					continue // a kind the compiler builds
				}
				r.Signature = d.Name.Text
				r.Fields, r.Uses = fields(d.Body)
				if len(r.Uses) > 0 {
					r.As = r.Uses[0].Type
				}
			case "Type":
				r.Signature = d.Name.Text
				if d.TypeValue != nil {
					r.Signature += " = " + syntax.TypeText(d.TypeValue)
				}
				r.Fields, r.Uses = fields(d.Body)
			default:
				continue
			}
			out = append(out, r)
		}
	}
	return out
}

// fields lists the fields of a Blueprint or Type body, and its `as` types.
func fields(b *syntax.Block) ([]RefField, []RefAs) {
	if b == nil {
		return nil, nil
	}
	var out []RefField
	var as []RefAs
	for _, it := range b.Items {
		switch it := it.(type) {
		case *syntax.FieldDecl:
			if it.Name == nil {
				continue
			}
			f := RefField{Name: it.Name.Text, Doc: syntax.ParseDoc(it.Doc).Text}
			if it.Type != nil {
				f.Type = syntax.TypeText(it.Type)
				f.Optional = strings.HasSuffix(f.Type, "?")
			}
			if it.Default != nil {
				f.Default = format.Expr(it.Default)
				f.Optional = true
			}
			out = append(out, f)
		case *syntax.AsStmt:
			if it.Type != nil {
				as = append(as, RefAs{Type: syntax.TypeText(it.Type), Doc: syntax.ParseDoc(it.Doc).Text})
			}
		}
	}
	return out, as
}

// internal reports whether a doc comment has the @internal tag: a
// declaration other prelude code uses, not one for projects.
func internal(d syntax.Doc) bool {
	for _, t := range d.Other {
		if t.Name == "internal" {
			return true
		}
	}
	return false
}
