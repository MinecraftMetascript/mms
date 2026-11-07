package ide

import (
	"github.com/minecraftmetascript/mms/engine/spec"
	"regexp"
	"sort"
	"strings"

	"github.com/minecraftmetascript/mms/engine/eval"
	"github.com/minecraftmetascript/mms/engine/format"
	"github.com/minecraftmetascript/mms/engine/prelude"
	"github.com/minecraftmetascript/mms/engine/sema"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// The names the evaluator binds (consts, Templates, Blueprints and Types)
// are not resources: they have no symbol in the checker, and expansion
// replaces their uses. Navigation and rename for them work on the source
// as written.

// nameKind says what an evaluator name is.
type nameKind int

const (
	nameConst nameKind = iota + 1
	nameTemplate
	nameBlueprint
	nameType
)

// evalName is a const, Template, Blueprint or Type.
type evalName struct {
	Kind nameKind
	Name string
	Path string      // the file that declares it
	Def  syntax.Span // its name there
}

// evalNameAt finds the evaluator name at off, at its declaration or at a
// use, or reports false.
func (s *Session) evalNameAt(path string, off int) (evalName, bool) {
	src := s.R.Source[path]
	if src == nil || s.R.Eval == nil {
		return evalName{}, false
	}
	var found evalName
	ok := false
	for _, it := range src.Items {
		if !it.Span().Contains(off) && it.Span().End != off {
			continue
		}
		bound := boundNames(it)
		syntax.Walk(it, func(n syntax.Node) bool {
			if !n.Span().Contains(off) && n.Span().End != off {
				return true
			}
			switch n := n.(type) {
			case *syntax.Decl:
				if n.Name != nil && within(n.Name.Sp, off) {
					if k := declNameKind(n); k != 0 {
						found, ok = s.lookupName(k, n.Name.Text)
					}
				}
			case *syntax.ConstStmt:
				if n.Name != nil && within(n.Name.Sp, off) {
					found, ok = s.lookupName(nameConst, n.Name.Text)
				}
			case *syntax.Ident:
				if within(n.Sp, off) {
					found, ok = s.identName(n.Name, bound)
				}
			case *syntax.TypeName:
				if within(typeNameSpan(n), off) {
					found, ok = s.typeRef(n.Name)
				}
			case *syntax.String, *syntax.CommandStmt:
				if name, ok2 := pieceAt(src.Src, n.Span(), off); ok2 {
					found, ok = s.identName(name, bound)
				}
			}
			return true
		})
	}
	return found, ok
}

func within(sp syntax.Span, off int) bool { return sp.Contains(off) || sp.End == off }

// declNameKind says what a declaration's name declares, if it is one of
// the evaluator's.
func declNameKind(d *syntax.Decl) nameKind {
	if d.Kind == nil {
		return 0
	}
	switch d.Kind.Name {
	case "Template":
		return nameTemplate
	case "Blueprint":
		return nameBlueprint
	case "Type":
		return nameType
	}
	return 0
}

// lookupName finds a declared evaluator name.
func (s *Session) lookupName(k nameKind, name string) (evalName, bool) {
	e := s.R.Eval
	var path string
	var sp syntax.Span
	switch k {
	case nameConst:
		c := e.Consts[name]
		if c == nil || c.S.Name == nil {
			return evalName{}, false
		}
		path, sp = c.Path, c.S.Name.Sp
	case nameTemplate:
		t := e.Templates[name]
		if t == nil || t.D.Name == nil {
			return evalName{}, false
		}
		path, sp = t.Path, t.D.Name.Sp
	case nameBlueprint:
		b := e.Blueprints[name]
		if b == nil || b.D.Name == nil {
			return evalName{}, false
		}
		path, sp = b.Path, b.D.Name.Sp
	case nameType:
		t := e.Types[name]
		if t == nil || t.D.Name == nil {
			return evalName{}, false
		}
		path, sp = t.Path, t.D.Name.Sp
	}
	return evalName{Kind: k, Name: name, Path: path, Def: sp}, true
}

// identName resolves a name used as a value: a const or a Template, or a
// Blueprint used as a kind, unless the item binds it itself.
func (s *Session) identName(name string, bound map[string]bool) (evalName, bool) {
	if bound[name] {
		return evalName{}, false
	}
	for _, k := range []nameKind{nameConst, nameTemplate, nameBlueprint} {
		if n, ok := s.lookupName(k, name); ok {
			return n, true
		}
	}
	return evalName{}, false
}

// typeRef resolves a type name: a Type or a Blueprint.
func (s *Session) typeRef(name string) (evalName, bool) {
	if n, ok := s.lookupName(nameType, name); ok {
		return n, true
	}
	return s.lookupName(nameBlueprint, name)
}

// typeNameSpan is the span of a type's name, without its arguments.
func typeNameSpan(n *syntax.TypeName) syntax.Span {
	end := min(n.Sp.Start+len([]rune(n.Name)), n.Sp.End)
	return syntax.Span{Start: n.Sp.Start, End: end}
}

// boundNames are the names an item binds itself: parameters, fields, loop
// and let variables. A use of one of them is not the const or Template of
// that name.
func boundNames(it syntax.Node) map[string]bool {
	out := map[string]bool{}
	add := func(ns ...*syntax.Name) {
		for _, n := range ns {
			if n != nil {
				out[n.Text] = true
			}
		}
	}
	syntax.Walk(it, func(n syntax.Node) bool {
		switch n := n.(type) {
		case *syntax.Param:
			add(n.Name)
		case *syntax.FieldDecl:
			add(n.Name)
		case *syntax.ForStmt:
			add(n.Var, n.Var2)
		case *syntax.Comprehension:
			add(n.Var, n.Var2)
		case *syntax.LetStmt:
			add(n.Name)
		}
		return true
	})
	return out
}

// pieceRe finds the `${...}` pieces of a string or command.
var pieceRe = regexp.MustCompile(`\$\{[^}]*\}`)

// wordRe finds a name inside a piece, and not after a '.', where it would
// be a field.
var wordRe = regexp.MustCompile(`[A-Za-z_][A-Za-z0-9_]*`)

// pieces returns the spans of the names inside `${...}` in the text at sp,
// in runes.
func pieces(src []rune, sp syntax.Span) []syntax.Span {
	if sp.Start < 0 || sp.End > len(src) || sp.End <= sp.Start {
		return nil
	}
	text := string(src[sp.Start:sp.End])
	var out []syntax.Span
	for _, m := range pieceRe.FindAllStringIndex(text, -1) {
		body := text[m[0]+2 : m[1]-1]
		for _, w := range wordRe.FindAllStringIndex(body, -1) {
			if w[0] > 0 && (body[w[0]-1] == '.' || body[w[0]-1] == ':') {
				continue
			}
			start := sp.Start + len([]rune(text[:m[0]+2+w[0]]))
			out = append(out, syntax.Span{Start: start, End: start + len([]rune(body[w[0]:w[1]]))})
		}
	}
	return out
}

// pieceAt finds the name inside `${...}` at off in a string or command.
func pieceAt(src []rune, sp syntax.Span, off int) (string, bool) {
	for _, p := range pieces(src, sp) {
		if within(p, off) {
			return string(src[p.Start:p.End]), true
		}
	}
	return "", false
}

// evalUses finds every place n is written in the project's files, its
// declaration included.
func (s *Session) evalUses(n evalName) []Location {
	seen := map[Location]bool{}
	var out []Location
	addLoc := func(path string, sp syntax.Span) {
		l := Location{Path: path, Span: sp}
		if !seen[l] {
			seen[l] = true
			out = append(out, l)
		}
	}
	if !strings.HasPrefix(n.Path, prelude.Prefix) {
		addLoc(n.Path, n.Def)
	}
	paths := make([]string, 0, len(s.R.Source))
	for p := range s.R.Source {
		paths = append(paths, p)
	}
	sort.Strings(paths)
	for _, path := range paths {
		src := s.R.Source[path]
		for _, it := range src.Items {
			bound := boundNames(it)
			value := n.Kind == nameConst || n.Kind == nameTemplate || n.Kind == nameBlueprint
			syntax.Walk(it, func(x syntax.Node) bool {
				switch x := x.(type) {
				case *syntax.Ident:
					if value && x.Name == n.Name && !bound[n.Name] {
						addLoc(path, x.Sp)
					}
				case *syntax.TypeName:
					if (n.Kind == nameType || n.Kind == nameBlueprint) && x.Name == n.Name {
						addLoc(path, typeNameSpan(x))
					}
				case *syntax.Decl:
					if x.Name != nil && declNameKind(x) == n.Kind && x.Name.Text == n.Name {
						addLoc(path, x.Name.Sp)
					}
				case *syntax.ConstStmt:
					if n.Kind == nameConst && x.Name != nil && x.Name.Text == n.Name {
						addLoc(path, x.Name.Sp)
					}
				case *syntax.String, *syntax.CommandStmt:
					if !value || bound[n.Name] {
						break
					}
					for _, p := range pieces(src.Src, x.Span()) {
						if string(src.Src[p.Start:p.End]) == n.Name {
							addLoc(path, p)
						}
					}
				}
				return true
			})
		}
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Path != out[j].Path {
			return out[i].Path < out[j].Path
		}
		return out[i].Span.Start < out[j].Span.Start
	})
	return out
}

var identRe = regexp.MustCompile(`^[A-Za-z_][A-Za-z0-9_]*$`)

// renameEval renames an evaluator name everywhere it is written.
func (s *Session) renameEval(n evalName, to string) ([]Edit, string) {
	if strings.HasPrefix(n.Path, prelude.Prefix) {
		return nil, "'" + n.Name + "' is built in and cannot be renamed"
	}
	if !identRe.MatchString(to) {
		return nil, "a name is letters, digits and underscores, not starting with a digit"
	}
	if to != n.Name {
		for _, k := range []nameKind{nameConst, nameTemplate, nameBlueprint, nameType} {
			if _, taken := s.lookupName(k, to); taken {
				return nil, "'" + to + "' is already declared"
			}
		}
	}
	var edits []Edit
	for _, l := range s.evalUses(n) {
		edits = append(edits, Edit{Path: l.Path, Span: l.Span, Text: to})
	}
	return edits, ""
}

// templateHead writes a Template's signature: `twice(x: int, y: int = 2): int`,
// and its parameters one by one.
func templateHead(d *syntax.Decl) (string, []string) {
	var params []string
	if d.Params != nil {
		for _, p := range d.Params.Params {
			if p.Name == nil {
				continue
			}
			t := p.Name.Text
			if p.Type != nil {
				t += ": " + syntax.TypeText(p.Type)
			}
			if p.Default != nil {
				t += " = " + format.Expr(p.Default)
			}
			params = append(params, t)
		}
	}
	label := d.NameText() + "(" + strings.Join(params, ", ") + ")"
	if d.Type != nil {
		label += ": " + syntax.TypeText(d.Type)
	}
	return label, params
}

// docText is a declaration's doc comment as Markdown, with its @param
// texts for the parameters given.
func docText(raw string) (string, syntax.Doc) {
	doc := syntax.ParseDoc(raw)
	text := doc.Text
	if doc.Deprecated != nil {
		text = strings.TrimSpace("**Deprecated.** " + *doc.Deprecated + "\n\n" + text)
	}
	return text, doc
}

// sourceTemplateCall finds the innermost call of a Template whose
// argument list holds off, in the source as written.
func (s *Session) sourceTemplateCall(path string, off int) (*syntax.File, *syntax.Call, *eval.Template) {
	src := s.R.Source[path]
	if src == nil || s.R.Eval == nil {
		return nil, nil, nil
	}
	var best *syntax.Call
	var bound map[string]bool
	for _, it := range src.Items {
		if !it.Span().Contains(off) {
			continue
		}
		syntax.Walk(it, func(n syntax.Node) bool {
			if c, ok := n.(*syntax.Call); ok && off > c.ArgsSp.Start && off < c.ArgsSp.End {
				best = c // the walk reaches inner calls later
			}
			return true
		})
		bound = boundNames(it)
	}
	if best == nil {
		return nil, nil, nil
	}
	id, ok := best.Fn.(*syntax.Ident)
	if !ok || bound[id.Name] {
		return nil, nil, nil
	}
	t := s.R.Eval.Templates[id.Name]
	if t == nil {
		return nil, nil, nil
	}
	return src, best, t
}

// templateSignature is signature help for a call of a Template, found in
// the source as written, since expansion replaces the call.
func (s *Session) templateSignature(path string, off int) (SignatureHelp, bool) {
	src, call, t := s.sourceTemplateCall(path, off)
	if call == nil {
		return SignatureHelp{}, false
	}
	h := SignatureHelp{}
	h.Label, h.Params = templateHead(t.D)
	var doc syntax.Doc
	h.Doc, doc = docText(t.D.Doc)
	for _, p := range h.Params {
		name, _, _ := strings.Cut(p, ":")
		name, _, _ = strings.Cut(name, " =")
		h.ParamDocs = append(h.ParamDocs, doc.Param(strings.TrimSpace(name)))
	}
	h.Active = activeParam(src, call, off, h.Params)
	return h, true
}

// templateArgCandidates lists what fits the Template parameter at off:
// values of its type, then the consts and value Templates in scope.
func (s *Session) templateArgCandidates(path string, off int, c *sema.Checker) ([]sema.Candidate, bool) {
	src, call, t := s.sourceTemplateCall(path, off)
	if call == nil || t.D.Params == nil {
		return nil, false
	}
	// Typing a parameter's name, `twice(x| = 1)`, is not a value.
	for _, a := range call.Args {
		if a.Name != nil && within(a.Name.Sp, off) {
			return nil, false
		}
	}
	_, params := templateHead(t.D)
	i := activeParam(src, call, off, params)
	if i < 0 || i >= len(t.D.Params.Params) {
		return nil, false
	}
	p := t.D.Params.Params[i]
	out := templateNamedArgs(t, call, i)
	if i == 0 {
		// A path the Template turns into an asset file: Piece("|").
		if helper, prefix, suffix, ok := s.templateAssetPath(t.D.Name.Text); ok {
			ext, what, _ := sema.AssetFileType(helper)
			for _, f := range s.P.Config.AssetFiles(s.P.Files, ext) {
				if strings.HasPrefix(f, prefix) && strings.HasSuffix(f, suffix) && len(f) >= len(prefix)+len(suffix) {
					f = f[len(prefix) : len(f)-len(suffix)]
					out = append(out, sema.Candidate{Label: `"` + f + `"`, Kind: sema.CandFile, Detail: what, Rank: -1})
				}
			}
		}
	}
	if p.Type != nil && s.R.TypeEnv != nil {
		if want := s.R.TypeEnv.Lookup(t.Path, p.Type); want != nil {
			// The values of its type, the consts and Templates that give
			// one, and the consts whose type is not written.
			out = append(out, c.TypeCandidates(want)...)
			for _, k := range s.evalValueCandidates() {
				if k.Kind == sema.CandID && s.R.Eval.Consts[k.Label] != nil && s.R.Eval.Consts[k.Label].S.Type == nil {
					out = append(out, k)
				}
			}
			return out, true
		}
	}
	return append(out, s.evalValueCandidates()...), true
}

// templateNamedArgs offers the parameters after the i-th that the call
// does not set by name yet, as named arguments: `title(@a, "x", subtitle = ...)`.
func templateNamedArgs(t *eval.Template, call *syntax.Call, i int) []sema.Candidate {
	present := map[string]bool{}
	for _, a := range call.Args {
		if a.Name != nil {
			present[a.Name.Text] = true
		}
	}
	var out []sema.Candidate
	for j, p := range t.D.Params.Params {
		if j < i || p.Name == nil || present[p.Name.Text] {
			continue
		}
		detail := "parameter"
		if p.Type != nil {
			detail = syntax.TypeText(p.Type)
		}
		out = append(out, sema.Candidate{Label: p.Name.Text, Kind: sema.CandField, Insert: p.Name.Text + " = $1", Detail: detail, Rank: 200 + j})
	}
	return out
}

// typedEval lists the consts and value Templates that have a written
// type, to be offered wherever that type fits.
func (s *Session) typedEval() []sema.TypedCandidate {
	if s.typed != nil || s.R.Eval == nil || s.R.TypeEnv == nil {
		return s.typed
	}
	s.typed = []sema.TypedCandidate{}
	for _, c := range s.evalValueCandidates() {
		var t *spec.Type
		if k := s.R.Eval.Consts[c.Label]; k != nil && c.Kind == sema.CandID {
			if k.S.Type != nil {
				t = s.R.TypeEnv.Lookup(k.Path, k.S.Type)
			}
		} else if tm := s.R.Eval.Templates[c.Label]; tm != nil && tm.D.Type != nil {
			t = s.R.TypeEnv.Lookup(tm.Path, tm.D.Type)
		}
		if t != nil {
			s.typed = append(s.typed, sema.TypedCandidate{Candidate: c, Type: t})
		}
	}
	return s.typed
}

// templateParams are the parameters of t, which its body can name, each
// offered where its type fits. A parameter without a type fits anywhere.
func (s *Session) templateParams(t *eval.Template) []sema.TypedCandidate {
	if t.D.Params == nil {
		return nil
	}
	var out []sema.TypedCandidate
	for _, p := range t.D.Params.Params {
		if p.Name == nil {
			continue
		}
		c := sema.Candidate{Label: p.Name.Text, Kind: sema.CandID, Detail: "parameter", Rank: -2}
		ty := &spec.Type{Kind: "any"}
		if p.Type != nil {
			c.Detail += ": " + syntax.TypeText(p.Type)
			ty = s.R.TypeEnv.Lookup(t.Path, p.Type)
		}
		out = append(out, sema.TypedCandidate{Candidate: c, Type: ty})
	}
	return out
}

// blueprintFields are the fields and lets of bp, which its emit block can name,
// each offered where its type fits.
func (s *Session) blueprintFields(bp *eval.Blueprint) []sema.TypedCandidate {
	var out []sema.TypedCandidate
	for _, f := range bp.Fields {
		if f.Name == nil {
			continue
		}
		c := sema.Candidate{Label: f.Name.Text, Kind: sema.CandID, Detail: "field", Rank: -2}
		ty := &spec.Type{Kind: "any"}
		if f.Type != nil {
			c.Detail += ": " + syntax.TypeText(f.Type)
			ty = s.R.TypeEnv.Lookup(bp.Path, f.Type)
		}
		out = append(out, sema.TypedCandidate{Candidate: c, Type: ty})
	}
	for _, l := range bp.Lets {
		if l.Name == nil {
			continue
		}
		c := sema.Candidate{Label: l.Name.Text, Kind: sema.CandID, Detail: "let", Rank: -2}
		ty := &spec.Type{Kind: "any"}
		if l.Type != nil {
			c.Detail += ": " + syntax.TypeText(l.Type)
			ty = s.R.TypeEnv.Lookup(bp.Path, l.Type)
		}
		out = append(out, sema.TypedCandidate{Candidate: c, Type: ty})
	}
	return out
}

// evalValueCandidates are the consts and value Templates a value can
// name, after what its type offers.
func (s *Session) evalValueCandidates() []sema.Candidate {
	e := s.R.Eval
	var out []sema.Candidate
	names := make([]string, 0, len(e.Consts))
	for n := range e.Consts {
		names = append(names, n)
	}
	sort.Strings(names)
	for _, n := range names {
		c := e.Consts[n]
		detail := "const"
		if c.S.Type != nil {
			detail += ": " + syntax.TypeText(c.S.Type)
		} else if t := eval.TypeOf(c.Value()); t != "" {
			detail += ": " + t
		}
		doc, _ := docText(c.S.Doc)
		out = append(out, sema.Candidate{Label: n, Kind: sema.CandID, Detail: detail, Doc: doc, Rank: 5})
	}
	names = names[:0]
	for n, t := range e.Templates {
		if t.D.Type != nil && syntax.TypeText(t.D.Type) == "Commands" || internal(t.D) {
			continue
		}
		names = append(names, n)
	}
	sort.Strings(names)
	for _, n := range names {
		t := e.Templates[n]
		label, _ := templateHead(t.D)
		doc, _ := docText(t.D.Doc)
		out = append(out, sema.Candidate{Label: n, Kind: sema.CandCallable, Detail: "Template", Doc: "```mms\nTemplate " + label + "\n```\n\n" + doc, Insert: n + "($1)", Rank: 6})
	}
	return out
}

// evalKeywords start the language's own top-level items.
var evalKeywords = []sema.Candidate{
	{Label: "const", Kind: sema.CandKeyword, Detail: "a named value", Insert: "const ${1:name} = $0",
		Doc: "A value with a name, copied in wherever it is named: `const max_wave = 10`."},
	{Label: "Template", Kind: sema.CandKeyword, Detail: "a reusable value or commands", Insert: "Template ${1:name}(${2:x: int}): ${3:int} = $0",
		Doc: "A value or commands with parameters, expanded where it is called: `Template twice(x: int): int = x * 2`."},
	{Label: "Type", Kind: sema.CandKeyword, Detail: "a named type", Insert: "Type ${1:Name} {\n\t$0\n}",
		Doc: "A record type with fields, or an alias such as `Type Mode = \"rigid\" | \"terrain_matching\"`."},
	{Label: "Blueprint", Kind: sema.CandKeyword, Detail: "a declaration kind", Insert: "Blueprint ${1:Name} {\n\t$0\n\temit {\n\t}\n}",
		Doc: "A new declaration kind: its fields, `check` blocks, an `as` conversion and the declarations it `emit`s."},
	{Label: "for", Kind: sema.CandKeyword, Detail: "declarations for each value", Insert: "for ${1:x} in ${2:[]} {\n\t$0\n}",
		Doc: "Repeats the declarations in its body for each value, with `${x}` in their names."},
}

// evalHover describes a const, Template, Blueprint or Type at off.
func (s *Session) evalHover(path string, off int) (string, syntax.Span, bool) {
	n, ok := s.evalNameAt(path, off)
	if !ok {
		return "", syntax.Span{}, false
	}
	var sp syntax.Span
	for _, l := range s.evalUses(n) {
		if l.Path == path && within(l.Span, off) {
			sp = l.Span
		}
	}
	if sp == (syntax.Span{}) {
		// A use in a prelude-only name: the word at off.
		sp = syntax.Span{Start: off, End: off}
	}
	e := s.R.Eval
	var head, doc string
	switch n.Kind {
	case nameConst:
		c := e.Consts[n.Name]
		head = "const " + n.Name
		if c.S.Type != nil {
			head += ": " + syntax.TypeText(c.S.Type)
		} else if t := eval.TypeOf(c.Value()); t != "" {
			head += ": " + t
		}
		if c.S.Value != nil {
			v := format.Expr(c.S.Value)
			if len(v) > 200 || strings.Contains(v, "\n") {
				v = strings.SplitN(v, "\n", 2)[0] + " …"
			}
			head += " = " + v
		}
		doc, _ = docText(c.S.Doc)
	case nameTemplate:
		t := e.Templates[n.Name]
		label, _ := templateHead(t.D)
		head = "Template " + label
		doc, _ = docText(t.D.Doc)
	case nameBlueprint:
		b := e.Blueprints[n.Name]
		head = "Blueprint " + n.Name
		doc, _ = docText(b.D.Doc)
	case nameType:
		t := e.Types[n.Name]
		head = "Type " + n.Name
		if t.D.TypeValue != nil {
			head += " = " + syntax.TypeText(t.D.TypeValue)
		}
		doc, _ = docText(t.D.Doc)
	}
	md := "```mms\n" + head + "\n```"
	if doc != "" {
		md += "\n\n" + doc
	}
	return md, sp, true
}
