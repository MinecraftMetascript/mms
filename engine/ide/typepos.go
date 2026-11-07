package ide

import (
	"strings"

	"github.com/minecraftmetascript/mms/engine/sema"
	"github.com/minecraftmetascript/mms/engine/syntax"
	"github.com/minecraftmetascript/mms/engine/syntax/parser"
	"github.com/minecraftmetascript/mms/engine/types"
)

// typeCompletion completes a type expression outside function code: a
// Template's parameter and return types, a Blueprint's field types, a
// Type declaration and a typed const. Inside Id<...> or Ids<...> it
// lists registries. ok is false when off is not in a type.
func (s *Session) typeCompletion(a *at, off int) ([]sema.Candidate, bool) {
	if s.R.TypeEnv == nil {
		return nil, false
	}
	f := a.u.File
	start := off
	if sp, ok := wordAt(f, off); ok {
		start = sp.Start
	}
	prev := prevToken(f, start)
	path := a.path
	inner := path[len(path)-1]
	if _, ok := inner.(*syntax.TypeName); !ok {
		if prev == nil {
			return nil, false
		}
		switch prev.Kind {
		case parser.MMSLexerCOLON, parser.MMSLexerLT, parser.MMSLexerPIPE, parser.MMSLexerLBRACK:
		case parser.MMSLexerASSIGN:
			// `Type T = |`
			if d := typeDecl(path); d == nil || d.Kind == nil || d.Kind.Name != "Type" {
				return nil, false
			}
		default:
			return nil, false
		}
		// `): |` after a Template's parameters is past the declaration the
		// parser kept; the token before the colon still belongs to it.
		path = syntax.PathAt(f, prev.Sp.Start)
		if prev.Kind == parser.MMSLexerLT {
			if before := prevToken(f, prev.Sp.Start); before != nil && (before.Text == "Id" || before.Text == "Ids") {
				if inTypes(path) {
					return s.registryCandidates(), true
				}
			}
			return nil, false
		}
		if prev.Kind != parser.MMSLexerCOLON && prev.Kind != parser.MMSLexerASSIGN {
			if _, ok := path[len(path)-1].(*syntax.TypeName); !ok && !inTypeExpr(path) {
				return nil, false
			}
		}
	} else if len(path) >= 2 {
		if outer, ok := path[len(path)-2].(*syntax.TypeName); ok && (outer.Name == "Id" || outer.Name == "Ids") {
			for _, arg := range outer.Args {
				if arg == inner {
					return s.registryCandidates(), true
				}
			}
		}
	}
	if !inTypes(path) {
		return nil, false
	}
	return s.typeNameCandidates(), true
}

// typeNameCandidates lists the names a type expression can use.
func (s *Session) typeNameCandidates() []sema.Candidate {
	var out []sema.Candidate
	for i, t := range s.R.TypeEnv.Names() {
		c := sema.Candidate{Label: t.Name, Kind: sema.CandKind, Detail: t.Detail, Rank: i}
		if t.Args {
			c.Insert = t.Name + "<$1>"
		}
		out = append(out, c)
	}
	return out
}

func (s *Session) registryCandidates() []sema.Candidate {
	var out []sema.Candidate
	for i, r := range s.R.TypeEnv.Registries() {
		out = append(out, sema.Candidate{Label: r, Kind: sema.CandEnum, Detail: "registry", Rank: i})
	}
	return out
}

// typeDecl is the declaration path is in, or nil.
func typeDecl(path []syntax.Node) *syntax.Decl {
	if len(path) > 1 {
		d, _ := path[1].(*syntax.Decl)
		return d
	}
	return nil
}

// inTypes reports whether path is where types are written outside
// function code: in a Template's or Type's head, a Blueprint's fields and
// `as` types,
// or a file-level const's annotation, but not in a value or body.
func inTypes(path []syntax.Node) bool {
	if len(path) < 2 {
		return false
	}
	ok := false
	if _, isConst := path[1].(*syntax.ConstStmt); isConst {
		ok = true
	} else if d := typeDecl(path); d != nil && d.Kind != nil {
		switch d.Kind.Name {
		case "Type":
			return true
		case "Template":
			ok = true
		case "Blueprint":
			for _, n := range path {
				switch n.(type) {
				case *syntax.FieldDecl, *syntax.AsStmt:
					ok = true
				}
			}
		}
	}
	if !ok {
		return false
	}
	// Values hold no types: a default, a const's value, a Template's
	// result or a Commands Template's code.
	var values []syntax.Node
	for _, n := range path {
		switch n := n.(type) {
		case *syntax.Decl:
			if n.Value != nil {
				values = append(values, n.Value)
			}
			if n.Body != nil && n.Kind != nil && n.Kind.Name == "Template" {
				values = append(values, n.Body)
			}
		case *syntax.ConstStmt:
			values = append(values, n.Value)
		case *syntax.FieldDecl:
			values = append(values, n.Default)
		case *syntax.Param:
			values = append(values, n.Default)
		case *syntax.AsStmt:
			values = append(values, n.Value)
			if n.Hint != nil {
				values = append(values, n.Hint)
			}
		}
	}
	for _, n := range path {
		for _, v := range values {
			if v != nil && n == v {
				return false
			}
		}
	}
	return true
}

// inTypeExpr reports whether path ends inside a type expression.
func inTypeExpr(path []syntax.Node) bool {
	for _, n := range path {
		switch n.(type) {
		case *syntax.TypeName, *syntax.TypeList, *syntax.TypeMap, *syntax.TypeUnion, *syntax.TypeOptional, *syntax.TypeParen:
			return true
		}
	}
	return false
}

// typeHover describes a built-in, resource or schema type named in a
// type position. The project's own Types and Blueprints have evalHover.
func (s *Session) typeHover(path string, off int) (string, syntax.Span, bool) {
	u := s.R.Unit(path)
	if u == nil || u.File == nil || s.R.TypeEnv == nil {
		return "", syntax.Span{}, false
	}
	p := syntax.PathAt(u.File, off)
	tn, ok := p[len(p)-1].(*syntax.TypeName)
	if !ok || !within(typeNameSpan(tn), off) || !inTypes(p) {
		return "", syntax.Span{}, false
	}
	for _, t := range s.R.TypeEnv.Names() {
		if t.Name != tn.Name {
			continue
		}
		head := t.Name
		if t.Args {
			head += "<registry>"
		}
		md := "```mms\n" + head + "\n```\n" + typeDetailDoc(t)
		if doc := sema.KindDoc(t.Name); doc != "" {
			md += "\n\n" + doc
		}
		if t.Detail == "schema type" {
			md += schemaFields(s.R.TypeEnv, path, tn)
		}
		return md, typeNameSpan(tn), true
	}
	return "", syntax.Span{}, false
}

// typeDetailDoc says in a sentence what kind of type t is.
func typeDetailDoc(t types.TypeName) string {
	switch t.Detail {
	case "value type":
		return "A value type."
	case "language type":
		return "A type of the language's own."
	case "id in a registry":
		return "An id in the registry or resource folder it names, such as `Id<item>`."
	case "ids or a tag in a registry":
		return "Ids or a `#tag` in the registry or resource folder it names, such as `Ids<block>`."
	case "resource: an id or an inline value":
		return "A resource: the id of one, or one written inline."
	case "schema type":
		return "A type from the game's data schema."
	}
	return t.Detail
}

// schemaFields lists a schema struct's fixed fields, for its hover.
func schemaFields(e *types.Env, path string, tn *syntax.TypeName) string {
	rt := e.Lookup(path, tn)
	for i := 0; rt != nil && rt.Kind == "reference" && i < 8; i++ {
		rt = e.Pack.Types[rt.Path]
	}
	if rt == nil || rt.Kind != "struct" {
		return ""
	}
	var sb strings.Builder
	n := 0
	for _, f := range rt.Fields {
		if f.Kind != "pair" || f.Key == "" {
			continue
		}
		if n == 20 {
			sb.WriteString("- …\n")
			break
		}
		if n == 0 {
			sb.WriteString("\n\n**Fields**\n\n")
		}
		n++
		opt := ""
		if f.Optional {
			opt = " *(optional)*"
		}
		sb.WriteString("- `" + f.Key + "`" + opt + "\n")
	}
	return strings.TrimRight(sb.String(), "\n")
}
