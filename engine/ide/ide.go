// Package ide implements editor features over an analysed project:
// completion, hover, signature help, navigation, rename, symbols and
// semantic tokens. Positions are rune offsets; the lsp package converts.
package ide

import (
	"fmt"
	"regexp"
	"slices"
	"sort"
	"strings"

	"github.com/minecraftmetascript/mms/engine/command"
	"github.com/minecraftmetascript/mms/engine/eval"
	"github.com/minecraftmetascript/mms/engine/fn"
	"github.com/minecraftmetascript/mms/engine/format"
	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/prelude"
	"github.com/minecraftmetascript/mms/engine/project"
	"github.com/minecraftmetascript/mms/engine/sema"
	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
	"github.com/minecraftmetascript/mms/engine/syntax/parser"
)

// Session pairs a project with one analysis result.
type Session struct {
	P *project.Project
	R *project.Result

	typed  []sema.TypedCandidate // typedEval's result, once worked out
	repair bool                  // this session completes in a repaired copy
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
	// fn is the function code at the cursor: the declaration of a
	// Function or hook, or an inline `{ ... }` function block (as a Decl
	// with only a body) in any declaration. nil outside function code.
	fn *syntax.Decl
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
			// `x = 1 + |` can leave the dangling operator outside the
			// declaration; the operand before it still belongs.
			if len(a.path) == 1 && mathOp(prev.Kind) {
				if before := prevToken(u.File, prev.Sp.Start); before != nil {
					a.path = syntax.PathAt(u.File, before.Sp.End)
				}
			}
		}
	}
	if len(a.path) > 1 {
		switch d := a.path[1].(type) {
		case *syntax.Decl:
			a.decl = s.R.Decls[d]
		case *syntax.ConstStmt:
			a.decl = s.R.Consts[d]
		}
	}
	// A declaration in a Blueprint's emit block was checked on its own.
	for i := len(a.path) - 1; i > 1; i-- {
		if d, ok := a.path[i].(*syntax.Decl); ok {
			if dr := s.R.Decls[d]; dr != nil && dr.Emitter != nil {
				a.decl = dr
				break
			}
		}
	}
	a.fn = a.codeRoot()
	if a.decl != nil {
		a.c = sema.NewChecker(s.P.Pack, s.P.Kinds, s.R.Syms, a.decl.NS)
	} else {
		a.c = sema.NewChecker(s.P.Pack, s.P.Kinds, s.R.Syms, s.P.Config.Namespace)
	}
	a.c.Typed = s.typedEval()
	if a.decl != nil && a.decl.Template != nil {
		a.c.Typed = append(s.templateParams(a.decl.Template), a.c.Typed...)
	}
	if a.decl != nil && a.decl.Emitter != nil {
		a.c.Typed = append(s.blueprintFields(a.decl.Emitter), a.c.Typed...)
	}
	s.templateArgInfo(a)
	return a
}

// templateArgInfo checks the argument of a Template call that the cursor
// is in against the parameter's type, so that completion and hover deep
// inside it, as in `Piece("a", processors = [Rule(|)])`, read the types
// there. The checker sees a Template call as a value of its result type
// and does not look inside its arguments.
func (s *Session) templateArgInfo(a *at) {
	if a.decl == nil || s.R.Eval == nil || s.R.TypeEnv == nil {
		return
	}
	for i := len(a.path) - 2; i >= 1; i-- {
		arg, ok := a.path[i+1].(*syntax.Arg)
		if !ok || arg.Value == nil {
			continue
		}
		call, ok := a.path[i].(*syntax.Call)
		if !ok {
			continue
		}
		id, ok := call.Fn.(*syntax.Ident)
		if !ok {
			continue
		}
		t := s.R.Eval.Templates[id.Name]
		if t == nil || t.D.Params == nil {
			continue
		}
		p := templateParam(t.D.Params.Params, call, arg)
		if p == nil || p.Type == nil {
			return
		}
		want := s.R.TypeEnv.Lookup(t.Path, p.Type)
		if want == nil || want.Kind == "any" {
			return
		}
		c := sema.NewChecker(s.P.Pack, s.P.Kinds, s.R.Syms, a.decl.NS)
		c.CheckType(arg.Value, want, syntax.TypeText(p.Type))
		d := *a.decl
		d.Info = make(map[syntax.Node]*sema.Info, len(a.decl.Info)+len(c.Info))
		for n, in := range a.decl.Info {
			d.Info[n] = in
		}
		for n, in := range c.Info {
			d.Info[n] = in
		}
		a.decl = &d
		return
	}
}

// templateParam is the parameter that arg of call sets: the one it
// names, or the next one in order that no argument names.
func templateParam(params []*syntax.Param, call *syntax.Call, arg *syntax.Arg) *syntax.Param {
	if arg.Name != nil {
		for _, p := range params {
			if p.Name != nil && p.Name.Text == arg.Name.Text {
				return p
			}
		}
		return nil
	}
	named := map[string]bool{}
	for _, x := range call.Args {
		if x.Name != nil {
			named[x.Name.Text] = true
		}
	}
	var free []*syntax.Param
	for _, p := range params {
		if p.Name != nil && !named[p.Name.Text] {
			free = append(free, p)
		}
	}
	k := 0
	for _, x := range call.Args {
		if x == arg {
			break
		}
		if x.Name == nil {
			k++
		}
	}
	if k < len(free) {
		return free[k]
	}
	return nil
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
	if rs := s.repairedSession(path, off); rs != nil {
		return rs.Completion(path, off)
	}
	// Book markup, in a .mmsbook file or a Book's text strings.
	if items, replace, ok := s.bookCompletion(path, off); ok {
		return items, replace
	}
	a := s.at(path, off)
	if a == nil {
		return nil, syntax.Span{Start: off, End: off}
	}
	f := a.u.File
	// The path of Texture("..."), Sound("...") or StructureFile("...") lists
	// the project's files of that type.
	if items, replace, ok := s.assetPathCompletion(f, off); ok {
		return items, replace
	}
	// A `${...}` a Blueprint fills in, in its strings and names.
	if items, replace, ok := s.blueprintInterpolation(a, off); ok {
		return items, replace
	}
	// No completion inside comments or strings (old issue #17).
	if i := f.TokenAt(off); i >= 0 {
		t := f.Tokens[i]
		inside := t.Sp.Start < off && off < t.Sp.End
		atEnd := off == t.Sp.End && (t.Kind == parser.MMSLexerLINE_COMMENT || t.Kind == parser.MMSLexerUNTERMINATED_COMMENT)
		if (t.IsComment() || t.Kind == parser.MMSLexerSTRING) && (inside || atEnd) {
			return nil, syntax.Span{Start: off, End: off}
		}
		if sp, ok := a.command(t, off); ok {
			return s.commandCompletion(a, sp, off)
		}
	}
	replace := syntax.Span{Start: off, End: off}
	inner := a.path[len(a.path)-1]
	switch n := inner.(type) {
	case *syntax.Ident:
		replace = word(f, n.Sp, off)
	case *syntax.Name:
		replace = n.Sp
	case *syntax.ResLoc:
		replace = word(f, n.Sp, off)
	case *syntax.Tag:
		replace = word(f, n.Sp, off)
	}
	// A word the tree has no node for, such as a lone `Fe` at the top
	// level (the parser drops a declaration with no name), is still
	// what the chosen item replaces; an empty span would leave it in
	// place and the editor would filter by nothing.
	if replace.Len() == 0 {
		if sp, ok := wordAt(f, off); ok {
			replace = sp
		}
	}
	cands, ok := s.selectorCompletion(a, off)
	if !ok {
		cands, ok = s.typeCompletion(a, off)
	}
	if !ok {
		cands, ok = s.templateArgCandidates(path, off, a.c)
	}
	if !ok {
		cands = s.complete(a, off)
	}
	// What no type describes in a Blueprint's own body: its keywords,
	// the kind after `primary`, and the names its values can use.
	if len(cands) == 0 {
		cands, _ = s.blueprintComplete(a, off)
	}
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

// word trims sp, the span of the value at the cursor, to the word being
// typed. Parse recovery can stretch a missing value over what follows it:
// `ingredient = |, result = ...` reads `, result` as the value, and
// replacing that would delete the comma and the next argument's name.
func word(f *syntax.File, sp syntax.Span, off int) syntax.Span {
	end := sp.Start
	for _, t := range f.Tokens {
		if t.Sp.Start < sp.Start {
			continue
		}
		if t.Sp.End > sp.End || t.Hidden || t.Kind == parser.MMSLexerCOMMA {
			break
		}
		end = t.Sp.End
	}
	return syntax.Span{Start: sp.Start, End: max(end, min(off, sp.End))}
}

// wordAt returns the span of the identifier or keyword token that off is
// in or right after.
func wordAt(f *syntax.File, off int) (syntax.Span, bool) {
	i := f.TokenAt(off)
	if i < 0 {
		return syntax.Span{}, false
	}
	t := f.Tokens[i]
	if t.Sp.Start == off || t.Synthetic {
		return syntax.Span{}, false
	}
	switch {
	case t.Kind == parser.MMSLexerAT:
		// A selector being typed: `@` alone, or `@a` below.
		return t.Sp, true
	case t.Kind == parser.MMSLexerREL_NAME:
		// The name of `~dx`, after the `~` or `^` and a `-`.
		sp := t.Sp
		sp.Start++
		if strings.HasPrefix(t.Text[1:], "-") {
			sp.Start++
		}
		if off < sp.Start {
			return syntax.Span{}, false
		}
		return sp, true
	case t.Kind == parser.MMSLexerIDENT || (t.Kind >= parser.MMSLexerCONST && t.Kind <= parser.MMSLexerFALSE):
		sp := t.Sp
		if i > 0 && f.Tokens[i-1].Kind == parser.MMSLexerAT && f.Tokens[i-1].Sp.End == sp.Start {
			sp.Start = f.Tokens[i-1].Sp.Start
		}
		return sp, true
	}
	return syntax.Span{}, false
}

// command reports whether off, in token t, is in a raw command, and
// returns the command's span.
func (a *at) command(t syntax.Token, off int) (syntax.Span, bool) {
	if t.Kind == parser.MMSLexerCOMMAND && off > t.Sp.Start {
		return t.Sp, true
	}
	// A lone `/` starting a line of function code, or after the `=` of
	// an assignment (`let n = /`), is a command whose name is not typed
	// yet; it lexes as SLASH until a letter follows.
	if t.Kind == parser.MMSLexerSLASH && off == t.Sp.End && a.code() && commandStart(a.u.File.Src, t.Sp.Start) {
		return t.Sp, true
	}
	return syntax.Span{}, false
}

// TriggerApplies reports whether completion should open when ch, one of
// the trigger characters only some places use, was just typed at off:
// '"' and '/' complete asset file paths, and '/' also starts a command.
// Elsewhere, such as after the '/' of a division, they open nothing.
func (s *Session) TriggerApplies(path string, off int, ch string) bool {
	// A `${` a Blueprint fills in, in its strings and names.
	if ch == "{" && !sema.IsBookFile(path) {
		if a := s.at(path, off); a != nil {
			if _, _, ok := s.blueprintInterpolation(a, off); ok {
				return true
			}
		}
	}
	// '{' and '$' otherwise only complete book markup.
	if sema.IsBookFile(path) || ch == "{" || ch == "$" {
		return s.bookTrigger(path, off)
	}
	if ch != `"` && ch != "/" {
		return true
	}
	a := s.at(path, off)
	if a == nil {
		return false
	}
	f := a.u.File
	if start, _, ok := assetPathString(f, off); ok {
		h := assetHelperBefore(f, start-1)
		if _, _, ok := sema.AssetFileType(h); ok || h == "TextFile" {
			return true
		}
		if _, _, _, ok := s.templateAssetPath(h); ok {
			return true
		}
	}
	if i := f.TokenAt(off); ch == "/" && i >= 0 {
		_, ok := a.command(f.Tokens[i], off)
		return ok
	}
	return false
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
		// After a kind on the same line, `Feature |`, a name comes next.
		start := off
		if sp, ok := wordAt(f, off); ok {
			start = sp.Start
		}
		if prev := prevToken(f, start); prev != nil && prev.Kind == parser.MMSLexerIDENT && lineOf(f, prev.Sp.End) == lineOf(f, off) {
			return nil
		}
		return s.declKinds()
	}
	d, _ := a.path[1].(*syntax.Decl)
	if d != nil && d.Kind.Sp.Contains(off) {
		return s.declKinds()
	}
	if d != nil && d.Name != nil && d.Name.Sp.Contains(off) {
		// A name on a later line than its kind is a new declaration's
		// kind: in `X x = foo` + newline + `Fe|`, the value `foo` reads
		// as the kind of a declaration named `Fe`.
		if l1, _ := f.LineCol(d.Kind.Sp.End); l1 != lineOf(f, off) {
			return s.declKinds()
		}
		return nil
	}
	if emitStatementStart(f, a.path, off) {
		return s.emitCandidates()
	}
	if a.decl == nil {
		return nil
	}
	if cs, ok := s.mmsComplete(a, off); ok {
		return cs
	}
	// After `=`, `(`, `,` or `[` with nothing typed yet, the value is a
	// Bad node or missing; find the node that expects it.
	if prev := prevToken(f, off); prev != nil && (prev.Kind == parser.MMSLexerASSIGN || mathOp(prev.Kind)) {
		// The missing value's Bad node can start at the next token, such
		// as the `}` closing `{ A = | }`.
		to := off
		if next := nextToken(f, off); next != nil {
			to = next.Sp.Start
		}
		for i := len(a.path) - 1; i >= 0; i-- {
			if bad := badAfter(a.path[i], prev.Sp.End, to); bad != nil {
				if in := a.info(bad); in.HasType() {
					return a.c.ValueCandidates(in)
				}
			}
		}
		// The right operand of `a + |` has the type of the left one (or
		// of the whole sum).
		if mathOp(prev.Kind) {
			for i := len(a.path) - 1; i >= 2; i-- {
				if x, ok := a.path[i].(syntax.Expr); ok {
					if in := a.info(x); in.HasType() {
						return a.c.ValueCandidates(in)
					}
				}
			}
		}
	}
	// `co| = 1` in function code assigns a Global or a local.
	if a.code() && assignTarget(a.path) {
		return s.statementCandidates(a, off)
	}
	for i := len(a.path) - 1; i >= 1; i-- {
		n := a.path[i]
		var parent syntax.Node
		if i > 0 {
			parent = a.path[i-1]
		}
		switch n := n.(type) {
		case *syntax.Ident, *syntax.ResLoc, *syntax.Tag, *syntax.Bad:
			// A word starting a statement in function code, or a helper
			// of an execute chain.
			if a.code() && startsStatement(a.path, i) {
				return s.statementCandidates(a, off)
			}
			// A line of a Blueprint's check block.
			if inCheck(a.path) && startsStatement(a.path, i) {
				return s.checkCandidates(a)
			}
			// A bare word in an object body is a field being typed.
			if _, ok := parent.(*syntax.ExprStmt); ok && i >= 2 {
				if b, ok := a.path[i-2].(*syntax.Block); ok {
					if cs := a.blockFields(b); cs != nil {
						return cs
					}
				}
			}
			if in := a.info(n); in.HasType() {
				out := a.c.ValueCandidates(in)
				// A word typed as an argument may be the start of the
				// argument's name: `Item(na|`.
				if arg, ok := parent.(*syntax.Arg); ok && arg.Name == nil && i >= 2 {
					if call, ok := a.path[i-2].(*syntax.Call); ok {
						out = append(out, a.namedArgs(call)...)
					}
				}
				return out
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
				if a.code() && sema.IsExecChain(p.X) {
					return sema.ExecCandidates(false)
				}
				if fs := a.c.StorageFieldCandidates(p.X); a.code() && fs != nil {
					return fs
				}
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
				if a.code() && sema.IsExecChain(n.X) {
					return sema.ExecCandidates(false)
				}
				if fs := a.c.StorageFieldCandidates(n.X); a.code() && fs != nil {
					return fs
				}
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
				// Inside a list, the enclosing object's fields never fit.
				return nil
			}
		case *syntax.Trailing:
			if n.Block.Sp.Contains(off) && off > n.Block.Sp.Start {
				// A surface rule's If(cond) { rule }: the block is the rule,
				// which may also name a declared or vanilla rule.
				if c, ok := n.X.(*syntax.Call); ok && isIdent(c.Fn, "If") {
					if ai := a.c.IfArgInfo(a.info(n), 1); ai != nil {
						return a.c.ValueCandidates(ai)
					}
				}
				if in := a.info(n); in.HasType() {
					return a.c.ValueCandidates(in)
				}
			}
		case *syntax.Block:
			if cs := a.blockFields(n); cs != nil {
				return cs
			}
			// An empty statement position in function code, such as a
			// blank line asked for with Ctrl+Space.
			if a.code() && i == len(a.path)-1 && off > n.Sp.Start && off < n.Sp.End {
				return s.statementCandidates(a, off)
			}
			// An object literal where no object fits, such as `height = {`:
			// the enclosing object's fields do not belong inside it.
			if _, ok := parent.(*syntax.Object); ok && !a.code() {
				return nil
			}
		}
	}
	return nil
}

// statementCandidates lists what can start a statement in function code:
// the variables in scope, the execute helpers, the statement keywords,
// then the Globals, objectives and functions.
func (s *Session) statementCandidates(a *at, off int) []sema.Candidate {
	if a.decl != nil && a.decl.Kind.MMS == "Test" && !inSection(a.path) {
		// A Test's own lines: setup first, then run and expect.
		var out []sema.Candidate
		for i, kw := range []struct{ word, doc string }{
			{"setup", "Commands and code that run first, in a fresh world after the pack loads: `setup { game.wave = 1 }`."},
			{"run", "Calls the function under test: `run next_wave()`."},
			{"expect", "A condition that must hold, checked in order: `expect game.wave == 2`."},
		} {
			out = append(out, sema.Candidate{Label: kw.word, Kind: sema.CandKeyword, Doc: kw.doc, Rank: i})
		}
		return out
	}
	out := locals(a.u.File, a.fn, off)
	out = append(out, sema.ExecCandidates(true)...)
	out = append(out, sema.TextCommandCandidates()...)
	for i, kw := range []string{"let", "if", "while", "for", "return"} {
		out = append(out, sema.Candidate{Label: kw, Kind: sema.CandKeyword, Rank: 100 + i})
	}
	for _, c := range s.projectNames(a.decl.NS) {
		c.Rank += 100
		out = append(out, c)
	}
	// In a Blueprint's emit block, its Function and Commands fields are
	// called: `cast()`.
	if a.decl != nil && a.decl.Emitter != nil {
		for _, f := range a.decl.Emitter.Fields {
			if f.Name != nil && f.Type != nil && (syntax.TypeText(f.Type) == "Function" || syntax.TypeText(f.Type) == "Commands") {
				out = append(out, sema.Candidate{Label: f.Name.Text, Kind: sema.CandCallable, Detail: "field: " + syntax.TypeText(f.Type), Insert: f.Name.Text + "()", Rank: 50})
			}
		}
	}
	// Templates that give commands, called as a statement.
	if e := s.R.Eval; e != nil {
		names := make([]string, 0, len(e.Templates))
		for name := range e.Templates {
			names = append(names, name)
		}
		sort.Strings(names)
		for _, name := range names {
			t := e.Templates[name]
			if t.D.Type == nil || syntax.TypeText(t.D.Type) != "Commands" || internal(t.D) {
				continue
			}
			label, _ := templateHead(t.D)
			doc, _ := docText(t.D.Doc)
			out = append(out, sema.Candidate{Label: name, Kind: sema.CandCallable, Detail: "Template",
				Doc: "```mms\nTemplate " + label + "\n```\n\n" + doc, Insert: name + "($1)", Rank: 180})
		}
	}
	return out
}

// assignTarget reports whether path ends at the key of a `key = value`
// line in a block of statements, which function code reads as an
// assignment, not an object field.
func assignTarget(path []syntax.Node) bool {
	n := len(path)
	if n < 4 {
		return false
	}
	name, ok := path[n-1].(*syntax.Name)
	if !ok {
		return false
	}
	fs, ok := path[n-2].(*syntax.FieldStmt)
	if !ok || fs.Key != name {
		return false
	}
	if _, ok := path[n-3].(*syntax.Block); !ok {
		return false
	}
	switch path[n-4].(type) {
	case *syntax.Decl, *syntax.IfStmt, *syntax.WhileStmt, *syntax.ForStmt, *syntax.Trailing, *syntax.RunExpr, *syntax.Block, *syntax.SectionStmt:
		return true
	}
	return false
}

// inCheck reports whether path is in a Blueprint's check { } block.
func inCheck(path []syntax.Node) bool {
	for _, n := range path {
		if sec, ok := n.(*syntax.SectionStmt); ok && sec.Word != nil && sec.Word.Name == "check" {
			return true
		}
	}
	return false
}

// checkCandidates lists what can start a line of a Blueprint's check
// block: error and warn, the statement keywords, and the fields.
func (s *Session) checkCandidates(a *at) []sema.Candidate {
	out := []sema.Candidate{
		{Label: "error", Kind: sema.CandCallable, Detail: "check", Doc: "Reports an error at a field: `error(cooldown, \"a cooldown is at least one tick\")`.", Insert: "error($1)"},
		{Label: "warn", Kind: sema.CandCallable, Detail: "check", Doc: "Reports a warning with a lint rule's name: `warn(cooldown, \"slow-wand\", \"...\")`.", Insert: "warn($1)", Rank: 1},
	}
	for i, kw := range []string{"if", "for", "let", "const"} {
		out = append(out, sema.Candidate{Label: kw, Kind: sema.CandKeyword, Rank: 10 + i})
	}
	for _, tc := range a.c.Typed {
		c := tc.Candidate
		c.Rank += 20
		out = append(out, c)
	}
	return out
}

// inSection reports whether a path is inside a section, such as a Test's
// setup { } block.
func inSection(path []syntax.Node) bool {
	for _, n := range path {
		if _, ok := n.(*syntax.SectionStmt); ok {
			return true
		}
	}
	return false
}

func lineOf(f *syntax.File, off int) int {
	l, _ := f.LineCol(off)
	return l
}

// lineStart reports whether only blanks precede off on its line.
func lineStart(src []rune, off int) bool {
	for i := off - 1; i >= 0 && src[i] != '\n'; i-- {
		if src[i] != ' ' && src[i] != '\t' {
			return false
		}
	}
	return true
}

// commandStart reports whether a '/' at off starts a command, as the
// lexer decides: at the start of a line, or after a lone '='.
func commandStart(src []rune, off int) bool {
	if lineStart(src, off) {
		return true
	}
	i := off - 1
	for i >= 0 && (src[i] == ' ' || src[i] == '\t') {
		i--
	}
	if i < 0 || src[i] != '=' {
		return false
	}
	return i == 0 || !strings.ContainsRune("=!<>+-*/%", src[i-1])
}

// code reports whether the cursor is in function code.
func (a *at) code() bool { return a.fn != nil }

// codeRoot finds the function code holding the cursor: the innermost
// inline function block on the path, such as CustomItem's
// `on_right_click = run { ... }` or an advancement's `function = run { ... }`,
// else the declaration when its kind holds code.
func (a *at) codeRoot() *syntax.Decl {
	if a.decl == nil {
		return nil
	}
	for i := len(a.path) - 1; i >= 2; i-- {
		if r, ok := a.path[i].(*syntax.RunExpr); ok && r.Body != nil {
			return &syntax.Decl{Body: r.Body, Sp: r.Sp}
		}
	}
	if sema.IsCodeKind(a.decl.Kind) {
		// OnAdvancement name = Trigger(...) { ... }: the code is the
		// trigger's block.
		if t, ok := a.decl.D.Value.(*syntax.Trailing); ok && a.decl.D.Body == nil && t.Block != nil {
			return &syntax.Decl{Kind: a.decl.D.Kind, Body: t.Block, Sp: a.decl.D.Sp}
		}
		return a.decl.D
	}
	return nil
}

// startsStatement reports whether path[i] is the head word of a statement:
// the word itself, or the first helper name of a call chain such as
// `As(@a).At(@s) { ... }`.
func startsStatement(path []syntax.Node, i int) bool {
	child := path[i]
	for j := i - 1; j >= 0; j-- {
		switch p := path[j].(type) {
		case *syntax.ExprStmt:
			return true
		case *syntax.AssignStmt:
			// The target of `x += 1`, `score[@s] = 2`.
			return p.Target == child
		case *syntax.Index:
			if p.X != child {
				return false
			}
		case *syntax.Call:
			if p.Fn != child {
				return false
			}
		case *syntax.Member:
			if p.X != child {
				return false
			}
		case *syntax.Trailing:
			if p.X != child {
				return false
			}
		default:
			return false
		}
		child = path[j]
	}
	return false
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
	if mathOp(prev.Kind) {
		return true
	}
	l1, _ := f.LineCol(prev.Sp.End)
	l2, _ := f.LineCol(off)
	return l1 == l2
}

// mathOp reports whether k is an arithmetic operator, which needs a
// right-hand operand.
func mathOp(k syntax.TokenKind) bool {
	switch k {
	case parser.MMSLexerPLUS, parser.MMSLexerMINUS, parser.MMSLexerSTAR, parser.MMSLexerSLASH, parser.MMSLexerPERCENT:
		return true
	}
	return false
}

// nextToken returns the first visible token starting at or after off.
func nextToken(f *syntax.File, off int) *syntax.Token {
	for i := range f.Tokens {
		t := &f.Tokens[i]
		if t.Sp.Start >= off && !t.Hidden && !t.Synthetic {
			return t
		}
	}
	return nil
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
	// A method's arguments are not its receiver's fields.
	if _, method := call.Fn.(*syntax.Member); method {
		return nil
	}
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
	// A surface rule's If(cond) { rule } takes a condition first.
	if id, ok := call.Fn.(*syntax.Ident); ok && id.Name == "If" {
		if ai := a.c.IfArgInfo(in, argIndex(a.u.File, call, off)); ai != nil {
			return a.c.ValueCandidates(ai)
		}
	}
	if in == nil || in.Callable == nil {
		// Shorthand such as a criterion helper records the fields its named
		// arguments set.
		if in != nil && in.Fields != nil {
			named := map[string]bool{}
			for _, arg := range call.Args {
				if arg.Name != nil {
					named[arg.Name.Text] = true
				}
			}
			var out []sema.Candidate
			for _, c := range a.c.FieldCandidates(in, named, true) {
				c.Insert = c.Label + " = $1"
				out = append(out, c)
			}
			return out
		}
		return nil
	}
	index := argIndex(a.u.File, call, off)
	named := map[string]bool{}
	for _, arg := range call.Args {
		if arg.Name != nil {
			named[arg.Name.Text] = true
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

// argIndex counts the positional arguments before off.
func argIndex(f *syntax.File, call *syntax.Call, off int) int {
	index := 0
	for _, arg := range call.Args {
		if arg.Name == nil && arg.Sp.End < off && commaBetween(f, arg.Sp.End, off) {
			index++
		}
	}
	return index
}

// activeParam is the parameter signature help highlights at off: the one
// a named argument there sets, or else the next positional slot, which
// skips parameters already given by name.
func activeParam(f *syntax.File, call *syntax.Call, off int, params []string) int {
	keys := make([]string, len(params))
	for i, p := range params {
		k, _, _ := strings.Cut(p, ":")
		keys[i] = strings.TrimSuffix(strings.TrimSpace(k), "?")
	}
	named := map[string]bool{}
	pos := 0
	for _, arg := range call.Args {
		before := arg.Sp.End < off && commaBetween(f, arg.Sp.End, off)
		if arg.Name == nil {
			if before {
				pos++
			}
			continue
		}
		key := arg.Name.Text
		if !slices.Contains(keys, key) {
			key = sema.Snake(key)
		}
		if !before && arg.Sp.Start <= off {
			if i := slices.Index(keys, key); i >= 0 {
				return i
			}
		}
		named[key] = true
	}
	last := 0
	for i, k := range keys {
		if named[k] {
			continue
		}
		if pos == 0 {
			return i
		}
		pos--
		last = i
	}
	return last
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
	md, sp := s.hover(path, off)
	// A bound name wins over what its word means elsewhere: an id
	// (`minecraft:rolls`) or a book's generic `${...}` hover.
	if l, lsp, ok := s.localAt(path, off); ok && (md == "" || strings.HasPrefix(md, "Expected:") || strings.Contains(md, "`minecraft:") || lsp != sp && sp.Contains(lsp.Start)) {
		lmd := localHover(l)
		if strings.HasPrefix(md, "Expected:") {
			lmd += "\n\n" + md
		} else if i := strings.Index(md, "Expected:"); i >= 0 {
			lmd += "\n\n" + md[i:]
		}
		md, sp = lmd, lsp
	} else if md == "" || strings.HasPrefix(md, "Expected:") {
		if emd, esp, ok := s.evalHover(path, off); ok {
			md, sp = emd, esp
		}
	}
	if md == "" {
		if tmd, tsp, ok := s.typeHover(path, off); ok {
			md, sp = tmd, tsp
		}
	}
	if md == "" {
		if cmd, csp, ok := s.commandHover(path, off); ok {
			md, sp = cmd, csp
		}
	}
	for _, c := range s.commandCalls(path) {
		if fsp := c.X.Fn.Span(); fsp.Contains(off) {
			if md != "" {
				md += "\n\n---\n\n"
			}
			return md + inlinedText(c, true), fsp
		}
	}
	return md, sp
}

func (s *Session) hover(path string, off int) (string, syntax.Span) {
	if md, sp, ok := s.bookHover(path, off); ok {
		return md, sp
	}
	a := s.at(path, off)
	if a == nil || len(a.path) < 2 {
		return "", syntax.Span{}
	}
	d, ok := a.path[1].(*syntax.Decl)
	if !ok || d.Kind == nil {
		return "", syntax.Span{} // a const or a top-level for
	}
	inner := a.path[len(a.path)-1]
	if a.decl != nil && a.decl.Blueprint != nil && d.Name != nil && d.Name.Sp.Contains(off) {
		return instanceHover(d), d.Name.Sp
	}
	if a.decl != nil && d.Name != nil && d.Name.Sp.Contains(off) {
		return s.declHover(a.decl), d.Name.Sp
	}
	if d.Kind.Sp.Contains(off) {
		if bp := s.blueprint(d); bp != nil {
			return "```mms\nBlueprint " + d.Kind.Name + "\n```\n" + sema.DocMarkdown(bp.D.Doc), d.Kind.Sp
		}
		k := a.decl
		if k != nil && k.Kind.Resource != "" {
			where := "."
			if k.Kind.Assets != "" {
				where = ", in the resource pack."
			}
			return "```mms\n" + d.Kind.Name + "\n```\nWrites `" + k.Kind.Dir("<namespace>") + "/<name>.json`" + where, d.Kind.Sp
		}
		if k != nil && k.Kind.MMS != "" {
			if doc := sema.KindDoc(k.Kind.MMS); doc != "" {
				return "```mms\n" + d.Kind.Name + "\n```\n" + doc, d.Kind.Sp
			}
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
				if in.Asset != nil && in.Asset.Helper == "StructureFile" {
					md += s.jigsawHover(in.Asset.Src)
				}
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

// jigsawHover lists a structure file's jigsaw blocks, for the hover of
// its StructureFile("...") call.
func (s *Session) jigsawHover(src string) string {
	js, ok := project.StructureJigsaws(s.P.Files, src)
	if !ok || len(js) == 0 {
		return ""
	}
	var sb strings.Builder
	sb.WriteString("\n\n**Jigsaw blocks**")
	if size, ok := project.StructureSize(s.P.Files, src); ok {
		fmt.Fprintf(&sb, " · size %d×%d×%d", size[0], size[1], size[2])
	}
	sb.WriteString("\n\n| At | Facing | Name | Pool | Target |\n|---|---|---|---|---|\n")
	for _, j := range js {
		facing := strings.ReplaceAll(j.Orientation, "_", " ")
		if j.Joint == "rollable" && (strings.HasPrefix(j.Orientation, "up_") || strings.HasPrefix(j.Orientation, "down_")) {
			facing += ", rollable"
		}
		fmt.Fprintf(&sb, "| %d %d %d | %s | `%s` | `%s` | `%s` |\n", j.Pos[0], j.Pos[1], j.Pos[2], facing, j.Name, j.Pool, j.Target)
	}
	return sb.String()
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
	if h, ok := s.templateSignature(path, off); ok {
		return h, true
	}
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
		if id, ok := call.Fn.(*syntax.Ident); ok && h.Label == "" && a.code() {
			h.Label, h.Params, h.Doc = sema.TextCommandSignature(id.Name)
			h.ParamDocs = make([]string, len(h.Params))
		}
		if h.Label == "" {
			return SignatureHelp{}, false
		}
		h.Active = activeParam(a.u.File, call, off, h.Params)
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
	d, ok := a.path[1].(*syntax.Decl)
	if !ok {
		return nil
	}
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
		if in.Symbol == nil && in.Local == nil && in.Converted == nil {
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
	if l, _, ok := s.localAt(path, off); ok {
		return []Location{{Path: path, Span: l.Name.Sp}}
	}
	if n, ok := s.evalNameAt(path, off); ok {
		// A prelude name's Path is its prelude file, which editors open
		// read-only.
		return []Location{{Path: n.Path, Span: n.Def}}
	}
	if cv := s.convertedAt(path, off); cv != nil && cv.Decl != nil && cv.Decl.Name != nil {
		return []Location{{Path: cv.Path, Span: cv.Decl.Name.Sp}}
	}
	sym := s.target(path, off)
	if sym == nil {
		if v := s.vanillaAt(path, off); v != "" {
			return []Location{{Path: v}}
		}
		return nil
	}
	if sym.Decl.Name == nil {
		return nil
	}
	return []Location{{Path: sym.Unit.Path, Span: sym.Decl.Name.Sp}}
}

// vanillaAt returns the path of the vanilla file that a minecraft id at off
// names, such as @vanilla/data/minecraft/worldgen/noise/ridge.json, or "".
func (s *Session) vanillaAt(path string, off int) string {
	v := s.P.Pack.Vanilla()
	a := s.at(path, off)
	if v == nil || a == nil || len(a.path) < 3 {
		return ""
	}
	in := a.info(a.path[len(a.path)-1])
	if _, ok := a.path[len(a.path)-1].(*syntax.CommandStmt); ok && a.decl != nil {
		// An id inside a command has a node of its own.
		in = nil
		for n, cin := range a.decl.Info {
			if cin.ID != "" && cin.Symbol == nil && n.Span().Contains(off) {
				in = cin
				break
			}
		}
	}
	if in == nil || in.ID == "" || in.Symbol != nil {
		return ""
	}
	return v.VanillaPath(in.Folder, in.ID)
}

// References returns every use of the symbol at off, with its declaration
// first when includeDecl is set.
func (s *Session) References(path string, off int, includeDecl bool) []Location {
	if l, _, ok := s.localAt(path, off); ok {
		out := s.localUses(path, l)
		if !includeDecl {
			out = out[1:]
		}
		return out
	}
	if n, ok := s.evalNameAt(path, off); ok {
		var out []Location
		for _, l := range s.evalUses(n) {
			if includeDecl || l != (Location{Path: n.Path, Span: n.Def}) {
				out = append(out, l)
			}
		}
		return out
	}
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
	if l, sp, ok := s.localAt(path, off); ok {
		// Callers and instances name parameters and fields too.
		return sp, l.What != "parameter" && l.What != "field"
	}
	if n, ok := s.evalNameAt(path, off); ok {
		if strings.HasPrefix(n.Path, prelude.Prefix) {
			return syntax.Span{}, false
		}
		for _, l := range s.evalUses(n) {
			if l.Path == path && within(l.Span, off) {
				return l.Span, true
			}
		}
		return syntax.Span{}, false
	}
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
	if l, _, ok := s.localAt(path, off); ok {
		if l.What == "parameter" || l.What == "field" {
			return nil, "a " + l.What + " is named by its callers too, so it can't be renamed here yet"
		}
		if !identRe.MatchString(newPath) {
			return nil, "a name is letters, digits and underscores"
		}
		var edits []Edit
		for _, u := range s.localUses(path, l) {
			edits = append(edits, Edit{Path: u.Path, Span: u.Span, Text: newPath})
		}
		return edits, ""
	}
	if n, ok := s.evalNameAt(path, off); ok {
		return s.renameEval(n, newPath)
	}
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
	// Kind is "resource", "function", "namespace", "tag", "objective",
	// "constant", "template", "type" or "blueprint".
	Kind      string
	Span, Sel syntax.Span
}

// Symbols returns a file's outline.
func (s *Session) Symbols(path string) []Symbol {
	u := s.R.Unit(path)
	if u == nil {
		return nil
	}
	var out []Symbol
	// What a Blueprint instance emits is listed as the instance, as
	// written, below.
	emitted := map[*syntax.Decl]bool{}
	if s.R.Eval != nil {
		if fo := s.R.Eval.Files[path]; fo != nil {
			for _, od := range fo.Decls {
				if od.Instance != nil {
					emitted[od.D] = true
				}
			}
		}
	}
	for _, d := range u.File.Decls {
		if emitted[d] {
			continue
		}
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
	// The source's consts, Templates, Types and Blueprints, which
	// expansion leaves out of the declarations.
	if src := s.R.Source[path]; src != nil {
		for _, it := range src.Items {
			switch it := it.(type) {
			case *syntax.ConstStmt:
				if it.Name != nil {
					out = append(out, Symbol{Name: it.Name.Text, Detail: "const", Kind: "constant", Span: it.Sp, Sel: it.Name.Sp})
				}
			case *syntax.Decl:
				if it.Name == nil || it.Kind == nil {
					continue
				}
				kind := map[string]string{"Template": "template", "Type": "type", "Blueprint": "blueprint"}[it.Kind.Name]
				if kind == "" && s.R.Eval.Blueprints[it.Kind.Name] != nil && !it.Name.Templated() {
					if bp := s.R.Eval.Blueprints[it.Kind.Name]; !bp.D.Extern {
						kind = "resource" // a Blueprint instance, such as a CustomItem
					}
				}
				if kind != "" {
					out = append(out, Symbol{Name: it.Name.Text, Detail: it.Kind.Name, Kind: kind, Span: it.Sp, Sel: it.Name.Sp})
				}
			}
		}
		sort.SliceStable(out, func(i, j int) bool { return out[i].Span.Start < out[j].Span.Start })
	}
	return out
}

// commandHover shows the signature of a raw command at off, from the
// word under the cursor on, in the game's /help syntax.
func (s *Session) commandHover(path string, off int) (string, syntax.Span, bool) {
	a := s.at(path, off)
	if a == nil || len(a.path) == 0 {
		return "", syntax.Span{}, false
	}
	cmd, ok := a.path[len(a.path)-1].(*syntax.CommandStmt)
	if !ok {
		return "", syntax.Span{}, false
	}
	base := cmd.Sp.Start + 1 // after the '/'
	u, ok := command.New(s.P.Pack).Usage(cmd.Text, max(off-base, 0))
	if !ok {
		return "", syntax.Span{}, false
	}
	md := "```\n/" + strings.Join(u.Lines, "\n/") + "\n```"
	if u.Arg != "" {
		md += "\n`" + u.Arg + "`"
	}
	sp := syntax.Span{Start: base + u.Start, End: base + u.End}
	if u.Start == 0 {
		sp.Start = cmd.Sp.Start // the '/' belongs to the command word
	}
	return md, sp, true
}

// commandCompletion completes inside a raw command: literals from the
// command tree and values for its arguments, including the project's own
// functions, objectives and resources.
func (s *Session) commandCompletion(a *at, sp syntax.Span, off int) ([]Item, syntax.Span) {
	base := sp.Start + 1
	if items, replace, ok := s.interpolationCompletion(a, base, off); ok {
		return items, replace
	}
	items, ws := s.commandItems(string(a.u.File.Src[base:sp.End]), off-base)
	return items, syntax.Span{Start: base + ws, End: off}
}

// commandItems completes command text at rune offset off, and returns
// where the word being completed starts.
func (s *Session) commandItems(text string, off int) ([]Item, int) {
	docs := map[string]*sema.Symbol{} // suggested label -> declaration
	ids := func(reg string) []string {
		folder := reg
		if tg, ok := strings.CutPrefix(reg, "tag/"); ok {
			folder = "tags/" + tg
		}
		if reg == "storage" {
			return s.storageIDs()
		}
		if reg == "sound_event" {
			return s.R.SoundIDs()
		}
		if id, ok := strings.CutPrefix(reg, "storage "); ok {
			return s.storagePaths(id)
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
	sugg, ws := command.New(s.P.Pack).Complete(text, off, ids)
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
	return items, ws
}

// storageRef finds the storage ids that commands in the project use.
var storageRef = regexp.MustCompile(`\bstorage +([a-z0-9_.\-]+:[a-z0-9_.\-/]+)`)

// storageIDs are the command storages a `storage` argument can name: the
// one the project's Storage declarations keep their records in, then the
// ones the project's commands already use.
func (s *Session) storageIDs() []string {
	seen := map[string]bool{}
	var out []string
	add := func(id string) {
		if !seen[id] {
			seen[id] = true
			out = append(out, id)
		}
	}
	for _, sym := range s.R.Syms.All() {
		if sym.Kind.MMS == "Storage" {
			add(fn.Storage(sym.NS))
		}
	}
	var used []string
	for _, u := range s.R.Units {
		if u.File == nil || prelude.Is(u.Path) {
			continue
		}
		for _, t := range u.File.Tokens {
			if t.Kind != parser.MMSLexerCOMMAND {
				continue
			}
			for _, m := range storageRef.FindAllStringSubmatch(string(u.File.Src[t.Sp.Start:t.Sp.End]), -1) {
				used = append(used, m[1])
			}
		}
	}
	sort.Strings(used)
	for _, id := range used {
		add(id)
	}
	return out
}

// storagePaths are the paths a command can write after storage id: the
// Storage declarations kept there and their fields, `game` and
// `game.wave`.
func (s *Session) storagePaths(id string) []string {
	var out []string
	for _, sym := range s.R.Syms.All() {
		if sym.Kind.MMS != "Storage" || fn.Storage(sym.NS) != id {
			continue
		}
		name := sym.Path
		out = append(out, name)
		tn, ok := sym.Decl.Type.(*syntax.TypeName)
		if !ok {
			continue
		}
		out = s.recordPaths(out, name, s.R.Syms.Types["::mms::"+tn.Name], 0)
	}
	return out
}

// recordPaths adds the dotted path of each field of t under name, and of
// the fields of nested records.
func (s *Session) recordPaths(out []string, name string, t *spec.Type, depth int) []string {
	for i := 0; t != nil && t.Kind == "reference" && i < 8; i++ {
		t = s.R.Syms.Types[t.Path]
	}
	if t == nil || t.Kind != "struct" || depth > 16 {
		return out
	}
	for _, f := range t.Fields {
		if f.Kind == "pair" && f.Key != "" {
			out = append(out, name+"."+f.Key)
			out = s.recordPaths(out, name+"."+f.Key, f.Type, depth+1)
		}
	}
	return out
}

// InlayHint is a label shown after a span, such as the base item a custom
// item stands for.
type InlayHint struct {
	Off     int // where the label goes
	Label   string
	Tooltip string
	// Type marks an inferred type, `: int`, which sits right after the
	// name.
	Type bool
}

// InlayHints lists the inlay hints in a file, in order: the base item
// after each custom item used as a plain item id.
func (s *Session) InlayHints(path string) []InlayHint {
	u := s.R.Unit(path)
	if u == nil {
		return nil
	}
	var out []InlayHint
	for _, d := range u.File.Decls {
		dr := s.R.Decls[d]
		if dr == nil {
			continue
		}
		for n, in := range dr.Info {
			if cv := in.Converted; cv != nil && cv.Hint != "" {
				// A Blueprint instance read as one of its `as`
				// conversions, which says why with a hint.
				to := format.Expr(cv.X)
				if r := []rune(to); len(r) > 30 {
					to = string(r[:29]) + "…"
				}
				out = append(out, InlayHint{Off: n.Span().End, Label: "→ " + to, Tooltip: cv.Hint})
			}
		}
	}
	out = append(out, s.typeHints(u)...)
	// A Commands Template writes its code where it is called.
	for _, c := range s.commandCalls(path) {
		out = append(out, InlayHint{Off: c.X.Sp.End, Label: "inlined", Tooltip: inlinedText(c, false)})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Off < out[j].Off })
	return out
}

// typeHints shows the inferred type after a `const` or `let` name that
// has none written: `let n = 0` reads `let n: int = 0`.
func (s *Session) typeHints(u *sema.Unit) []InlayHint {
	var out []InlayHint
	hint := func(n *syntax.Name, t string) {
		if n != nil && t != "" {
			out = append(out, InlayHint{Off: n.Sp.End, Label: ": " + t, Tooltip: "Inferred type", Type: true})
		}
	}
	for _, it := range u.File.Items {
		if k, ok := it.(*syntax.ConstStmt); ok && k.Name != nil && k.Type == nil && s.R.Eval != nil {
			if c := s.R.Eval.Consts[k.Name.Text]; c != nil && c.S == k {
				hint(k.Name, eval.TypeOf(c.Value()))
			}
		}
	}
	for _, d := range u.File.Decls {
		dr := s.R.Decls[d]
		if dr == nil || d.Body == nil {
			continue
		}
		syntax.Walk(d.Body, func(n syntax.Node) bool {
			if l, ok := n.(*syntax.LetStmt); ok && l.Name != nil && l.Type == nil {
				if in := dr.Info[l.Name]; in != nil && in.Local != nil {
					hint(l.Name, in.Local.Type)
				}
			}
			return true
		})
	}
	return out
}

// commandCalls lists the calls in a file of Templates that return
// Commands, which write their statements in place.
func (s *Session) commandCalls(path string) []*eval.Call {
	if s.R.Eval == nil {
		return nil
	}
	var out []*eval.Call
	for _, c := range s.R.Eval.Calls {
		if c.Path == path && c.X != nil && c.Template != nil && c.Template.D.Body != nil {
			out = append(out, c)
		}
	}
	return out
}

// maxInlined is how many lines of a Template's expansion hover shows.
const maxInlined = 20

// inlinedText says what a Commands Template call writes, as markdown for
// hover or as plain text for an inlay's tooltip.
func inlinedText(c *eval.Call, md bool) string {
	code := format.Stmts(c.Expanded)
	if lines := strings.Split(code, "\n"); len(lines) > maxInlined {
		code = strings.Join(lines[:maxInlined], "\n") + fmt.Sprintf("\n// … %d more lines", len(lines)-maxInlined)
	}
	name := c.Template.D.NameText()
	if code == "" {
		return "Template " + name + " is inlined here at compile time, and writes nothing for these arguments."
	}
	if md {
		return "Inlined at compile time: Template `" + name + "` writes\n\n```mms\n" + code + "\n```"
	}
	return "Template " + name + " is inlined here at compile time:\n" + code
}

func isIdent(e syntax.Expr, name string) bool {
	id, ok := e.(*syntax.Ident)
	return ok && id.Name == name
}
