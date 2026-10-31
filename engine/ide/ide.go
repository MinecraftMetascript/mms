// Package ide implements editor features over an analysed project:
// completion, hover, signature help, navigation, rename, symbols and
// semantic tokens. Positions are rune offsets; the lsp package converts.
package ide

import (
	"sort"
	"strings"

	"github.com/minecraftmetascript/mms/engine/command"
	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/project"
	"github.com/minecraftmetascript/mms/engine/sema"
	"github.com/minecraftmetascript/mms/engine/syntax"
	"github.com/minecraftmetascript/mms/engine/syntax/parser"
)

// Session pairs a project with one analysis result.
type Session struct {
	P *project.Project
	R *project.Result
}

// Location is a span in a file.
type Location struct {
	Path string
	Span syntax.Span
}

// at gathers what is at an offset: the unit, the path of nodes and the
// declaration's checked result.
type at struct {
	u    *sema.Unit
	path []syntax.Node
	decl *project.Decl
	c    *sema.Checker
}

func (s *Session) at(path string, off int) *at {
	u := s.R.Unit(path)
	if u == nil {
		return nil
	}
	a := &at{u: u, path: syntax.PathAt(u.File, off)}
	// Past the end of an unfinished declaration (`X x = |` at the end of
	// the file), the cursor still belongs to it.
	if len(a.path) == 1 {
		if prev := prevToken(u.File, off); prev != nil && continues(u.File, prev, off) {
			a.path = syntax.PathAt(u.File, prev.Sp.End)
		}
	}
	if len(a.path) > 1 {
		if d, ok := a.path[1].(*syntax.Decl); ok {
			a.decl = s.R.Decls[d]
		}
	}
	if a.decl != nil {
		a.c = sema.NewChecker(s.P.Pack, s.P.Kinds, s.R.Syms, a.decl.NS)
	} else {
		a.c = sema.NewChecker(s.P.Pack, s.P.Kinds, s.R.Syms, s.P.Config.Namespace)
	}
	return a
}

func (a *at) info(n syntax.Node) *sema.Info {
	if a.decl == nil || a.decl.Info == nil {
		return nil
	}
	return a.decl.Info[n]
}

// Item is a completion item.
type Item struct {
	Label  string
	Kind   sema.CandidateKind
	Detail string
	Doc    string
	// Deprecated marks a declaration whose doc comment has @deprecated.
	Deprecated bool
	Insert     string // snippet when Snippet is set
	Snippet    bool
	Sort       string
}

// Completion lists completions at off, and the span the chosen item
// replaces (the word being typed).
func (s *Session) Completion(path string, off int) ([]Item, syntax.Span) {
	a := s.at(path, off)
	if a == nil {
		return nil, syntax.Span{Start: off, End: off}
	}
	f := a.u.File
	// No completion inside comments or strings (old issue #17).
	if i := f.TokenAt(off); i >= 0 {
		t := f.Tokens[i]
		inside := t.Sp.Start < off && off < t.Sp.End
		atEnd := off == t.Sp.End && (t.Kind == parser.MMSLexerLINE_COMMENT || t.Kind == parser.MMSLexerUNTERMINATED_COMMENT)
		if (t.IsComment() || t.Kind == parser.MMSLexerSTRING) && (inside || atEnd) {
			return nil, syntax.Span{Start: off, End: off}
		}
		if t.Kind == parser.MMSLexerCOMMAND && off > t.Sp.Start {
			return s.commandCompletion(a, t.Sp, off)
		}
	}
	replace := syntax.Span{Start: off, End: off}
	inner := a.path[len(a.path)-1]
	switch n := inner.(type) {
	case *syntax.Ident:
		replace = n.Sp
	case *syntax.Name:
		replace = n.Sp
	case *syntax.ResLoc:
		replace = n.Sp
	}
	cands := s.complete(a, off)
	items := make([]Item, 0, len(cands))
	for i, c := range cands {
		it := Item{Label: c.Label, Kind: c.Kind, Detail: c.Detail, Doc: c.Doc, Deprecated: c.Deprecated, Sort: sortKey(i)}
		if c.Insert != "" {
			it.Insert, it.Snippet = c.Insert, true
		}
		items = append(items, it)
	}
	return items, replace
}

func sortKey(i int) string {
	const digits = "0123456789"
	b := []byte{'0', '0', '0', '0'}
	for p := 3; p >= 0 && i > 0; p-- {
		b[p] = digits[i%10]
		i /= 10
	}
	return string(b)
}

func (s *Session) complete(a *at, off int) []sema.Candidate {
	f := a.u.File
	// Top level, or typing a declaration kind.
	if len(a.path) == 1 {
		return s.P.Kinds.DeclKinds()
	}
	d, _ := a.path[1].(*syntax.Decl)
	if d != nil && d.Kind.Sp.Contains(off) {
		return s.P.Kinds.DeclKinds()
	}
	if d != nil && d.Name != nil && d.Name.Sp.Contains(off) {
		return nil
	}
	if a.decl == nil {
		return nil
	}
	// After `=`, `(`, `,` or `[` with nothing typed yet, the value is a
	// Bad node or missing; find the node that expects it.
	if prev := prevToken(f, off); prev != nil && prev.Kind == parser.MMSLexerASSIGN {
		for i := len(a.path) - 1; i >= 0; i-- {
			if bad := badAfter(a.path[i], prev.Sp.End, off); bad != nil {
				if in := a.info(bad); in.HasType() {
					return a.c.ValueCandidates(in)
				}
			}
		}
	}
	for i := len(a.path) - 1; i >= 1; i-- {
		n := a.path[i]
		var parent syntax.Node
		if i > 0 {
			parent = a.path[i-1]
		}
		switch n := n.(type) {
		case *syntax.Ident, *syntax.ResLoc, *syntax.Bad:
			// A bare word in an object body is a field being typed.
			if _, ok := parent.(*syntax.ExprStmt); ok && i >= 2 {
				if b, ok := a.path[i-2].(*syntax.Block); ok {
					if cs := a.blockFields(b); cs != nil {
						return cs
					}
				}
			}
			if in := a.info(n); in.HasType() {
				return a.c.ValueCandidates(in)
			}
			// The helper name of a call: the expected type is recorded on
			// the outermost node of the chain.
			if call, ok := parent.(*syntax.Call); ok && call.Fn == n {
				for j := i - 1; j >= 1; j-- {
					x, ok := a.path[j].(syntax.Expr)
					if !ok {
						break
					}
					if in := a.info(x); in.HasType() {
						return a.c.ValueCandidates(in)
					}
				}
			}
		case *syntax.Name:
			switch p := parent.(type) {
			case *syntax.Member:
				return a.methods(p)
			case *syntax.FieldStmt:
				if i >= 2 {
					if b, ok := a.path[i-2].(*syntax.Block); ok {
						return a.blockFields(b)
					}
				}
			case *syntax.Arg:
				if i >= 2 {
					if call, ok := a.path[i-2].(*syntax.Call); ok {
						return a.namedArgs(call)
					}
				}
			}
			return nil
		case *syntax.Member:
			if n.Name == nil || !n.Name.Sp.Contains(off) {
				return a.methods(n)
			}
		case *syntax.Call:
			if off > n.ArgsSp.Start && off < n.ArgsSp.End {
				return a.callArgs(n, off)
			}
		case *syntax.Index:
			return nil
		case *syntax.List:
			if off > n.Sp.Start && off < n.Sp.End {
				if in := a.c.ListItemInfo(a.info(n)); in != nil {
					return a.c.ValueCandidates(in)
				}
			}
		case *syntax.Trailing:
			if n.Block.Sp.Contains(off) && off > n.Block.Sp.Start {
				if in := a.info(n); in.HasType() {
					return a.c.ValueCandidates(in)
				}
			}
		case *syntax.Block:
			if cs := a.blockFields(n); cs != nil {
				return cs
			}
		}
	}
	return nil
}

// badAfter finds a Bad node under n starting between from and to.
func badAfter(n syntax.Node, from, to int) syntax.Node {
	var found syntax.Node
	syntax.Walk(n, func(x syntax.Node) bool {
		if found != nil {
			return false
		}
		if b, ok := x.(*syntax.Bad); ok && b.Sp.Start >= from && b.Sp.Start <= to {
			found = b
		}
		return true
	})
	return found
}

// continues reports whether text typed at off extends the construct that
// ends with prev: an operator that needs a right-hand side, or anything on
// the same line.
func continues(f *syntax.File, prev *syntax.Token, off int) bool {
	switch prev.Kind {
	case parser.MMSLexerASSIGN, parser.MMSLexerCOMMA, parser.MMSLexerLPAREN, parser.MMSLexerLBRACK, parser.MMSLexerDOT:
		return true
	}
	l1, _ := f.LineCol(prev.Sp.End)
	l2, _ := f.LineCol(off)
	return l1 == l2
}

func prevToken(f *syntax.File, off int) *syntax.Token {
	for i := len(f.Tokens) - 1; i >= 0; i-- {
		t := &f.Tokens[i]
		if t.Sp.End <= off && !t.Hidden && !t.Synthetic {
			return t
		}
	}
	return nil
}

// blockFields offers the unset fields of an object body or declaration
// body.
func (a *at) blockFields(b *syntax.Block) []sema.Candidate {
	in := a.info(b)
	if in == nil {
		return nil
	}
	present := map[string]bool{}
	for _, it := range b.Items {
		if f, ok := it.(*syntax.FieldStmt); ok {
			present[f.Key.Text] = true
		}
	}
	return a.c.FieldCandidates(in, present, false)
}

func (a *at) methods(m *syntax.Member) []sema.Candidate {
	return a.c.MethodCandidates(a.baseInfo(m))
}

// baseInfo returns the Info of the call a chain starts from. When the
// checker stopped early, the call is resolved again against the type the
// outermost chain node expects.
func (a *at) baseInfo(e syntax.Expr) *sema.Info {
	base := sema.ChainBase(e)
	if base == nil {
		return nil
	}
	if in := a.info(base); in != nil && in.Callable != nil {
		return in
	}
	for i := len(a.path) - 1; i >= 0; i-- {
		x, ok := a.path[i].(syntax.Expr)
		if !ok {
			break
		}
		if in := a.info(x); in.HasType() {
			if ci := a.c.CallInfo(base, in); ci != nil {
				return ci
			}
		}
	}
	return nil
}

func (a *at) namedArgs(call *syntax.Call) []sema.Candidate {
	base := sema.ChainBase(call)
	in := a.info(base)
	if in == nil {
		return nil
	}
	present := map[string]bool{}
	for _, arg := range call.Args {
		if arg.Name != nil {
			present[arg.Name.Text] = true
		}
	}
	return a.c.FieldCandidates(in, present, true)
}

// callArgs offers values for the argument at off, plus named arguments.
func (a *at) callArgs(call *syntax.Call, off int) []sema.Candidate {
	// A method call `X(...).M(|)`: the argument is the method's field.
	if m, ok := call.Fn.(*syntax.Member); ok && m.Name != nil {
		if in := a.c.MethodArgInfo(a.baseInfo(m.X), m.Name.Text); in != nil {
			return a.c.ValueCandidates(in)
		}
		return nil
	}
	in := a.info(call)
	if in == nil || in.Callable == nil {
		// Sugar such as If(...) records the expected type on the call.
		if in != nil && in.HasType() && len(call.Args) == 0 {
			return nil
		}
		return nil
	}
	index := 0
	named := map[string]bool{}
	for _, arg := range call.Args {
		if arg.Name != nil {
			named[arg.Name.Text] = true
			continue
		}
		if arg.Sp.End < off && commaBetween(a.u.File, arg.Sp.End, off) {
			index++
		}
	}
	var out []sema.Candidate
	if ai := a.c.ArgInfo(in, index, named); ai != nil {
		out = a.c.ValueCandidates(ai)
	}
	for _, c := range a.c.FieldCandidates(in, named, true) {
		c.Rank += 10
		c.Insert = c.Label + " = $1"
		out = append(out, c)
	}
	return out
}

func commaBetween(f *syntax.File, from, to int) bool {
	for _, t := range f.Tokens {
		if t.Sp.Start >= from && t.Sp.End <= to && t.Kind == parser.MMSLexerCOMMA {
			return true
		}
	}
	return false
}

// Hover returns Markdown for the node at off and the span it describes.
func (s *Session) Hover(path string, off int) (string, syntax.Span) {
	a := s.at(path, off)
	if a == nil || len(a.path) < 2 {
		return "", syntax.Span{}
	}
	d := a.path[1].(*syntax.Decl)
	inner := a.path[len(a.path)-1]
	if a.decl != nil && d.Name != nil && d.Name.Sp.Contains(off) {
		return s.declHover(a.decl), d.Name.Sp
	}
	if d.Kind.Sp.Contains(off) {
		k := a.decl
		if k != nil && k.Kind.Resource != "" {
			where := "."
			if k.Kind.Assets != "" {
				where = ", in the resource pack."
			}
			return "```mms\n" + d.Kind.Name + "\n```\nWrites `" + k.Kind.Dir("<namespace>") + "/<name>.json`" + where, d.Kind.Sp
		}
		return "", syntax.Span{}
	}
	if n, in := a.inCommand(off); in != nil {
		if md := a.c.Hover(in); md != "" {
			return md, n.Span()
		}
		return "", syntax.Span{}
	}
	for i := len(a.path) - 1; i >= 2; i-- {
		n := a.path[i]
		if in := a.info(n); in != nil {
			if md := a.c.Hover(in); md != "" {
				return md, n.Span()
			}
		}
		// Only the innermost word gets type-only hovers.
		if n != inner {
			break
		}
	}
	return "", syntax.Span{}
}

func (s *Session) declHover(d *project.Decl) string {
	var sb strings.Builder
	sb.WriteString("```mms\n" + sema.SymbolHead(d.D.Kind.Name, d.ID(), d.D) + "\n```\n")
	if md := sema.DocMarkdown(d.D.Doc); md != "" {
		sb.WriteString(md + "\n\n")
	}
	if d.Value != nil {
		js := string(jsonv.Marshal(d.Value))
		lines := strings.Split(js, "\n")
		if len(lines) > 25 {
			js = strings.Join(lines[:25], "\n") + "\n…"
		}
		sb.WriteString("Lowers to:\n```json\n" + strings.TrimRight(js, "\n") + "\n```")
	}
	return sb.String()
}

// SignatureHelp describes the call around the cursor.
type SignatureHelp struct {
	Label     string
	Doc       string   // Markdown, or ""
	Params    []string // parameter labels
	ParamDocs []string // Markdown per parameter, "" when undocumented
	Active    int
}

// Signature returns signature help for the call around off.
func (s *Session) Signature(path string, off int) (SignatureHelp, bool) {
	a := s.at(path, off)
	if a == nil {
		return SignatureHelp{}, false
	}
	for i := len(a.path) - 1; i >= 0; i-- {
		call, isCall := a.path[i].(*syntax.Call)
		if !isCall || off <= call.ArgsSp.Start || off >= call.ArgsSp.End {
			continue
		}
		if _, isMethod := call.Fn.(*syntax.Member); isMethod {
			return SignatureHelp{}, false
		}
		var h SignatureHelp
		h.Label, h.Params, h.ParamDocs = a.c.Signature(a.info(call))
		h.Doc = a.c.SignatureDoc(a.info(call))
		if h.Label == "" {
			h = s.functionSignature(a.info(call.Fn))
		}
		if h.Label == "" {
			return SignatureHelp{}, false
		}
		active := 0
		for _, arg := range call.Args {
			if arg.Name == nil && arg.Sp.End < off && commaBetween(a.u.File, arg.Sp.End, off) {
				active++
			}
		}
		h.Active = min(active, max(len(h.Params)-1, 0))
		return h, true
	}
	return SignatureHelp{}, false
}

// functionSignature is signature help for a call to one of the project's
// Functions, with its doc comment and @param texts.
func (s *Session) functionSignature(in *sema.Info) SignatureHelp {
	if in == nil || in.Symbol == nil {
		return SignatureHelp{}
	}
	sym := s.current(in.Symbol)
	if sym.Kind.MMS != "Function" {
		return SignatureHelp{}
	}
	doc := syntax.ParseDoc(sym.Decl.Doc)
	h := SignatureHelp{Label: strings.TrimPrefix(sema.SymbolHead("Function", sym.Path, sym.Decl), "Function "), Doc: doc.Text}
	if doc.Deprecated != nil {
		h.Doc = strings.TrimSpace("**Deprecated.** " + *doc.Deprecated + "\n\n" + h.Doc)
	}
	for _, p := range sema.SigOf(sym.Decl).Params {
		h.Params = append(h.Params, p.Name+": "+p.Type)
		h.ParamDocs = append(h.ParamDocs, doc.Param(p.Name))
	}
	return h
}

// target finds the symbol at off: a reference or a declaration's name.
func (s *Session) target(path string, off int) *sema.Symbol {
	a := s.at(path, off)
	if a == nil || len(a.path) < 2 {
		return nil
	}
	d := a.path[1].(*syntax.Decl)
	if d.Name != nil && d.Name.Sp.Contains(off) {
		for _, sym := range s.R.Syms.All() {
			if sym.Decl == d {
				return sym
			}
		}
		return nil
	}
	if _, in := a.inCommand(off); in != nil {
		return s.current(in.Symbol) // nil for a local
	}
	for i := len(a.path) - 1; i >= 2; i-- {
		if in := a.info(a.path[i]); in != nil && in.Symbol != nil {
			return s.current(in.Symbol)
		}
	}
	return nil
}

// inCommand finds the reference at off inside a raw command: an id the
// command names, or a name inside ${...}. Commands are text in the syntax
// tree, so the checker records these with nodes of their own.
func (a *at) inCommand(off int) (syntax.Node, *sema.Info) {
	cmd, ok := a.path[len(a.path)-1].(*syntax.CommandStmt)
	if !ok || a.decl == nil {
		return nil, nil
	}
	var best syntax.Node
	var bestIn *sema.Info
	for n, in := range a.decl.Info {
		if in.Symbol == nil && in.Local == nil {
			continue
		}
		switch n.(type) {
		case *syntax.Ident, *syntax.ResLoc, *syntax.Tag, *syntax.String:
		default:
			continue
		}
		sp := n.Span()
		if !sp.Contains(off) || sp.Start < cmd.Sp.Start || sp.End > cmd.Sp.End {
			continue
		}
		if best == nil || sp.Len() < best.Span().Len() || sp.Len() == best.Span().Len() && sp.Start < best.Span().Start {
			best, bestIn = n, in
		}
	}
	return best, bestIn
}

// current maps a symbol recorded by an earlier analysis (a declaration
// whose result was reused) to this analysis's symbol, whose declaration
// node and spans are up to date.
func (s *Session) current(sym *sema.Symbol) *sema.Symbol {
	if sym == nil {
		return nil
	}
	if cur := s.R.Syms.Same(sym); cur != nil {
		return cur
	}
	return sym
}

// Definition returns where the symbol at off is declared.
func (s *Session) Definition(path string, off int) []Location {
	sym := s.target(path, off)
	if sym == nil || sym.Decl.Name == nil {
		return nil
	}
	return []Location{{Path: sym.Unit.Path, Span: sym.Decl.Name.Sp}}
}

// References returns every use of the symbol at off, with its declaration
// first when includeDecl is set.
func (s *Session) References(path string, off int, includeDecl bool) []Location {
	sym := s.target(path, off)
	if sym == nil {
		return nil
	}
	var out []Location
	if includeDecl && sym.Decl.Name != nil {
		out = append(out, Location{Path: sym.Unit.Path, Span: sym.Decl.Name.Sp})
	}
	out = append(out, s.uses(sym)...)
	return out
}

func (s *Session) uses(sym *sema.Symbol) []Location {
	seen := map[Location]bool{}
	var out []Location
	for _, d := range s.R.Order {
		for n, in := range d.Info {
			if in.Symbol == nil || (in.Symbol != sym && s.current(in.Symbol) != sym) {
				continue
			}
			switch n.(type) {
			case *syntax.Ident, *syntax.ResLoc, *syntax.Tag, *syntax.String:
			default:
				continue
			}
			loc := Location{Path: d.Unit.Path, Span: n.Span()}
			if !seen[loc] {
				seen[loc] = true
				out = append(out, loc)
			}
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

// Edit replaces a span with new text.
type Edit struct {
	Path string
	Span syntax.Span
	Text string
}

// PrepareRename returns the span of the renamable name at off.
func (s *Session) PrepareRename(path string, off int) (syntax.Span, bool) {
	sym := s.target(path, off)
	if sym == nil {
		return syntax.Span{}, false
	}
	a := s.at(path, off)
	inner := a.path[len(a.path)-1]
	return inner.Span(), true
}

// Rename renames the symbol at off to newPath (a resource path such as
// `hills` or `terrain/hills`), editing the declaration and every use.
func (s *Session) Rename(path string, off int, newPath string) ([]Edit, string) {
	sym := s.target(path, off)
	if sym == nil {
		return nil, "nothing to rename here"
	}
	newPath = strings.TrimPrefix(newPath, sym.NS+":")
	if strings.ContainsAny(newPath, ": ") || newPath == "" || newPath != strings.ToLower(newPath) {
		return nil, "a resource name must be a lowercase path"
	}
	var edits []Edit
	if n := sym.Decl.Name; n != nil {
		text := newPath
		if strings.Contains(n.Text, ":") {
			text = sym.NS + ":" + newPath
		}
		edits = append(edits, Edit{Path: sym.Unit.Path, Span: n.Sp, Text: text})
	}
	for _, loc := range s.uses(sym) {
		u := s.R.Unit(loc.Path)
		old := string(u.File.Src[loc.Span.Start:loc.Span.End])
		text := newPath
		prefix := ""
		if strings.HasPrefix(old, "#") {
			prefix = "#"
			old = old[1:]
		}
		if strings.HasPrefix(old, `"`) {
			edits = append(edits, Edit{Path: loc.Path, Span: loc.Span, Text: `"` + prefix + sym.NS + ":" + newPath + `"`})
			continue
		}
		if strings.Contains(old, ":") {
			text = sym.NS + ":" + newPath
		}
		edits = append(edits, Edit{Path: loc.Path, Span: loc.Span, Text: prefix + text})
	}
	return edits, ""
}

// Symbol is a document outline entry.
type Symbol struct {
	Name, Detail string
	Kind         string // "resource", "function", "namespace", "tag", "objective"
	Span, Sel    syntax.Span
}

// Symbols returns a file's outline.
func (s *Session) Symbols(path string) []Symbol {
	u := s.R.Unit(path)
	if u == nil {
		return nil
	}
	var out []Symbol
	for _, d := range u.File.Decls {
		sym := Symbol{Detail: d.Kind.Name, Span: d.Sp, Sel: d.Kind.Sp, Kind: "resource"}
		if d.Name != nil {
			sym.Name, sym.Sel = d.Name.Text, d.Name.Sp
		} else {
			sym.Name = d.Kind.Name
		}
		if dr := s.R.Decls[d]; dr != nil {
			switch {
			case dr.Kind.MMS == "Namespace":
				sym.Kind = "namespace"
			case dr.Kind.MMS != "" && dr.Kind.MMS != "Objective" && dr.Kind.MMS != "Global":
				sym.Kind = "function"
			case dr.Kind.TagOf != "":
				sym.Kind = "tag"
			case dr.Kind.MMS != "":
				sym.Kind = "objective"
			}
		} else if d.Kind.Name == "Namespace" {
			sym.Kind = "namespace"
		}
		out = append(out, sym)
	}
	return out
}

// commandCompletion completes inside a raw command: literals from the
// command tree and values for its arguments, including the project's own
// functions, objectives and resources.
func (s *Session) commandCompletion(a *at, sp syntax.Span, off int) ([]Item, syntax.Span) {
	base := sp.Start + 1
	text := string(a.u.File.Src[base:sp.End])
	docs := map[string]*sema.Symbol{} // suggested label -> declaration
	ids := func(reg string) []string {
		folder := reg
		if tg, ok := strings.CutPrefix(reg, "tag/"); ok {
			folder = "tags/" + tg
		}
		var out []string
		for _, sym := range s.R.Syms.All() {
			switch {
			case reg == "objective" && sym.Kind.MMS == "Objective":
				out = append(out, sym.Path)
				docs[sym.Path] = sym
			case reg != "objective" && sym.Kind.MMS != "Objective" && sym.Kind.MMS != "Global" && sym.Kind.Folder() == folder:
				out = append(out, sym.ID()) // commands need ids namespaced
				docs[sym.ID()], docs["#"+sym.ID()] = sym, sym
			}
		}
		return out
	}
	sugg, ws := command.New(s.P.Pack).Complete(text, off-base, ids)
	items := make([]Item, 0, len(sugg))
	for i, sg := range sugg {
		k := sema.CandKeyword
		if sg.ID {
			k = sema.CandID
		}
		it := Item{Label: sg.Label, Kind: k, Detail: sg.Detail, Sort: sortKey(i)}
		if sym := docs[sg.Label]; sg.ID && sym != nil {
			it.Doc, it.Deprecated = sema.DocMarkdown(sym.Decl.Doc), sema.Deprecated(sym.Decl)
		}
		items = append(items, it)
	}
	return items, syntax.Span{Start: base + ws, End: off}
}
