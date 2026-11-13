package ide

import (
	"sort"

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
	// The name of a record's only field may be a map's key type being
	// typed: `{st|: string}` is on its way to `{string: string}`.
	if _, ok := inner.(*syntax.Name); ok && len(path) >= 3 {
		if r, ok := path[len(path)-3].(*syntax.TypeRecord); ok && len(r.Fields) == 1 && inTypes(path) {
			return s.typeNameCandidates(a.u.Path), true
		}
	}
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
			if before := prevToken(f, prev.Sp.Start); before != nil && (before.Text == "Id" || before.Text == "Ids" || before.Text == "Tag") {
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
		if outer, ok := path[len(path)-2].(*syntax.TypeName); ok && (outer.Name == "Id" || outer.Name == "Ids" || outer.Name == "Tag") {
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
	return s.typeNameCandidates(a.u.Path), true
}

// typeNameCandidates lists the names a type expression in file path can
// use.
func (s *Session) typeNameCandidates(path string) []sema.Candidate {
	var out []sema.Candidate
	briefs := sema.TypeBriefs(s.P.Pack, s.P.Kinds)
	for i, t := range s.R.TypeEnv.Names(path) {
		c := sema.Candidate{Label: t.Name, Kind: sema.CandKind, Detail: t.Detail, Rank: i}
		if b, ok := briefs[t.Name]; ok && t.Detail == "schema type" {
			// Its shape, and the types pack files use before those only
			// the game's own data does.
			c.Detail, c.Doc = b.Summary, b.Doc
			if !b.Used {
				c.Rank += 100000
			}
		}
		if t.Args {
			c.Insert = t.Name + "<$1>"
		}
		out = append(out, c)
	}
	sort.SliceStable(out, func(i, j int) bool { return out[i].Rank < out[j].Rank })
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
// function code: in a Template's or Type's head, a Blueprint's fields,
// parameters, result and `as` types, or a file-level const's annotation,
// but not in a value or body.
func inTypes(path []syntax.Node) bool {
	if len(path) < 2 {
		return false
	}
	for _, n := range path[2:] {
		if _, ok := n.(*syntax.Decl); ok {
			// A declaration in an emit block, such as `Objective o: |`,
			// completes like one at the top level.
			return false
		}
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
			// Its fields, `as` types, parameters (`Blueprint b(label: |)`)
			// and result type.
			// A `name: type` inside a section, such as a half-typed
			// `Objective o: |` in emit, is not a field.
			for _, n := range path {
				switch n.(type) {
				case *syntax.SectionStmt:
					return false
				case *syntax.FieldDecl, *syntax.AsStmt, *syntax.ParamList, *syntax.Param:
					ok = true
				}
			}
			if len(path) > 2 && d.Type != nil && path[2] == syntax.Node(d.Type) {
				ok = true
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
		case *syntax.TypeName, *syntax.TypeList, *syntax.TypeMap, *syntax.TypeRecord, *syntax.TypeUnion, *syntax.TypeOptional, *syntax.TypeParen:
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
	for _, t := range s.R.TypeEnv.Names(path) {
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
			md = "```mms\n" + head + "\n```\n" + s.schemaTypeDoc(t.Name)
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

// schemaTypeDoc describes the schema type MMS code names name: its
// shape, examples, values or fields, and a link to the reference.
func (s *Session) schemaTypeDoc(name string) string {
	c := sema.NewChecker(s.P.Pack, s.P.Kinds, s.R.Syms, s.P.Config.Namespace)
	path := c.SchemaTypePath(name)
	e := c.DescribeType(path)
	if e == nil {
		return "A type from the game's data schema."
	}
	if !c.HasTypeEntry(path) {
		path = ""
	}
	return sema.TypeMarkdown(e, path)
}
