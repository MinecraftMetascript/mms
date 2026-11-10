package ide

import (
	"strings"

	"github.com/minecraftmetascript/mms/engine/eval"
	"github.com/minecraftmetascript/mms/engine/sema"
	"github.com/minecraftmetascript/mms/engine/syntax"
	"github.com/minecraftmetascript/mms/engine/syntax/parser"
)

// declKinds lists what can start a declaration: the kinds, then the
// Blueprints file path can name, except those whose doc says @internal.
func (s *Session) declKinds(path string) []sema.Candidate {
	out := append(s.P.Kinds.DeclKinds(), evalKeywords...)
	if s.R == nil || s.R.Eval == nil {
		return out
	}
	for _, bp := range s.R.Eval.BlueprintsIn(path) {
		if internal(bp.D) {
			continue
		}
		n := bp.D.Name.Text
		out = append(out, sema.Candidate{
			Label: n, Kind: sema.CandKind, Detail: "Blueprint",
			Doc: sema.DocMarkdown(bp.D.Doc), Deprecated: sema.Deprecated(bp.D),
			Insert: n + " ${1:name} {\n\t$0\n}",
		})
	}
	return out
}

// emitKeywords start the statements of an emit block other than a
// declaration.
var emitKeywords = []sema.Candidate{
	{Label: "primary", Kind: sema.CandKeyword, Detail: "the declaration the instance stands for", Insert: "primary ",
		Doc: "Marks the emitted declaration that the instance's name refers to: `primary CustomItem ${self.path} { ... }`."},
	{Label: "const", Kind: sema.CandKeyword, Detail: "a named value", Insert: "const ${1:name} = $0",
		Doc: "A value with a name, for the declarations after it: `const id = \"${self.path}\"`."},
	{Label: "if", Kind: sema.CandKeyword, Detail: "declarations when a condition holds", Insert: "if (${1:cond}) {\n\t$0\n}",
		Doc: "Emits the declarations in its body only when the condition, worked out when the pack is built, holds."},
	{Label: "for", Kind: sema.CandKeyword, Detail: "declarations for each value", Insert: "for ${1:x} in ${2:[]} {\n\t$0\n}",
		Doc: "Repeats the declarations in its body for each value, with `${x}` in their names."},
}

// emitCandidates lists what starts a statement in a Blueprint's emit
// block: the declaration kinds, the project's Blueprints and the emit
// keywords. Template, Type and Blueprint are file-level only.
func (s *Session) emitCandidates(path string) []sema.Candidate {
	var out []sema.Candidate
	for _, c := range s.declKinds(path) {
		switch c.Label {
		case "Template", "Type", "Blueprint", "const", "for":
			continue
		}
		out = append(out, c)
	}
	return append(out, emitKeywords...)
}

// emitStatementStart reports whether off starts a statement directly in
// a Blueprint's emit block, or in an if or for inside it: an empty line,
// or the first word of one, not inside an emitted declaration.
func emitStatementStart(f *syntax.File, path []syntax.Node, off int) bool {
	last := len(path) - 1
	if last < 0 {
		return false
	}
	if _, ok := path[last].(*syntax.Block); !ok {
		// The first word being typed: `Fu|`, read as an expression.
		if _, ok := path[last].(*syntax.Ident); !ok || last < 2 {
			return false
		}
		if _, ok := path[last-1].(*syntax.ExprStmt); !ok {
			return false
		}
		last -= 2
		if _, ok := path[last].(*syntax.Block); !ok {
			return false
		}
	} else if prev := prevToken(f, off); prev != nil && lineOf(f, prev.Sp.End) == lineOf(f, off) && prev.Kind != parser.MMSLexerLBRACE {
		// Not after something else on the same line.
		return false
	}
	for i := last - 1; i >= 1; i-- {
		switch n := path[i].(type) {
		case *syntax.SectionStmt:
			return n.Word != nil && n.Word.Name == "emit"
		case *syntax.IfStmt, *syntax.ForStmt, *syntax.Block:
		default:
			return false
		}
	}
	return false
}

func internal(d *syntax.Decl) bool {
	if d == nil || d.Doc == "" {
		return false
	}
	for _, t := range syntax.ParseDoc(d.Doc).Other {
		if t.Name == "internal" {
			return true
		}
	}
	return false
}

// blueprint returns the Blueprint a declaration's kind names, or nil.
// d is written in file path.
func (s *Session) blueprint(path string, d *syntax.Decl) *eval.Blueprint {
	if d == nil || d.Kind == nil || s.R == nil || s.R.Eval == nil {
		return nil
	}
	return s.R.Eval.Blueprint(path, d.Kind.Name)
}

// instanceHover describes a Blueprint instance at its name.
func instanceHover(d *syntax.Decl) string {
	md := "```mms\n" + d.Kind.Name + " " + d.NameText() + "\n```"
	if doc := sema.DocMarkdown(d.Doc); doc != "" {
		md += "\n" + doc
	}
	return md
}

// convertedAt returns the conversion of the Blueprint instance named at
// off, or nil: a use such as `ruby * 2` or `${ruby}`.
func (s *Session) convertedAt(path string, off int) *sema.Conversion {
	a := s.at(path, off)
	if a == nil || len(a.path) < 2 {
		return nil
	}
	if _, in := a.inCommand(off); in != nil {
		return in.Converted
	}
	for i := len(a.path) - 1; i >= 2; i-- {
		if in := a.info(a.path[i]); in != nil && in.Converted != nil {
			return in.Converted
		}
	}
	return nil
}

// blueprintKeywords start the items of a Blueprint's body other than a
// field, which is a name the author makes up.
var blueprintKeywords = []sema.Candidate{
	{Label: "let", Kind: sema.CandKeyword, Detail: "a value worked out from the fields", Insert: "let ${1:name} = $0",
		Doc: "A value the check and emit blocks can name, worked out once per instance: `let lang_key = \"item.${self.namespace}.${self.path}\"`."},
	{Label: "check", Kind: sema.CandKeyword, Detail: "problems with an instance", Insert: "check {\n\t$0\n}", Rank: 1,
		Doc: "Reports problems with an instance's fields where the instance is written: `if (cooldown < 1) { error(cooldown, \"...\") }`."},
	{Label: "emit", Kind: sema.CandKeyword, Detail: "the declarations an instance writes", Insert: "emit {\n\t$0\n}", Rank: 2,
		Doc: "The declarations each instance writes, with `primary` on the one its name stands for."},
	{Label: "as", Kind: sema.CandKeyword, Detail: "what the name means where a type is expected", Insert: "as ${1:Type} = $0", Rank: 3,
		Doc: "What an instance's name stands for where a value of the type goes: `as ItemStack = { id = base }`."},
}

// selfCandidates are the parts of `self`, the instance being expanded.
var selfCandidates = []sema.Candidate{
	{Label: "id", Kind: sema.CandField, Detail: "self: string", Doc: "The instance's id, `<namespace>:<path>`."},
	{Label: "namespace", Kind: sema.CandField, Detail: "self: string", Doc: "The namespace the instance is in.", Rank: 1},
	{Label: "path", Kind: sema.CandField, Detail: "self: string", Doc: "The instance's name, the path of its id.", Rank: 2},
}

// blueprintValues lists what an expression in a Blueprint can name,
// worked out when the pack is built: its fields and lets, self, then the
// consts and Templates.
func (s *Session) blueprintValues(bp *eval.Blueprint) []sema.Candidate {
	var out []sema.Candidate
	for _, tc := range s.blueprintFields(bp) {
		out = append(out, tc.Candidate)
	}
	out = append(out, sema.Candidate{Label: "self", Kind: sema.CandID, Detail: "the instance", Doc: "The instance being expanded: `self.id`, `self.namespace` and `self.path`.", Rank: 1})
	for _, c := range s.evalValueCandidates(bp.Path) {
		c.Rank += 10
		out = append(out, c)
	}
	for i, kw := range []string{"true", "false", "none"} {
		out = append(out, sema.Candidate{Label: kw, Kind: sema.CandKeyword, Rank: 200 + i})
	}
	return out
}

// blueprintComplete completes the parts of a Blueprint's own body that
// no declaration's type describes: the keywords starting its items and
// its check block's lines, the kind after `primary`, and the values its
// lets, defaults, checks and emit statements compute with. ok is false
// elsewhere, such as in the body of a declaration it emits.
func (s *Session) blueprintComplete(a *at, off int) ([]sema.Candidate, bool) {
	if a.decl == nil || a.decl.Emitter == nil || a.code() || len(a.path) < 3 {
		return nil, false
	}
	bp := a.decl.Emitter
	f := a.u.File
	start := off
	if sp, ok := wordAt(f, off); ok {
		start = sp.Start
	}
	prev := prevToken(f, start)
	// The kind of the declaration `primary` marks.
	if prev != nil && prev.Kind == parser.MMSLexerPRIMARY {
		var out []sema.Candidate
		for _, c := range s.emitCandidates(a.u.Path) {
			if c.Kind != sema.CandKeyword {
				out = append(out, c)
			}
		}
		return out, true
	}
	if a.decl.D != bp.D {
		return nil, false // inside an emitted declaration
	}
	// The type of an `as` conversion.
	if prev != nil && prev.Kind == parser.MMSLexerAS && s.R.TypeEnv != nil {
		return s.typeNameCandidates(a.u.Path), true
	}
	// The first word of a line: an item of the body, or a line of a
	// check block.
	if lineStart(f.Src, start) {
		b := innermostBlock(a.path)
		if b >= 2 && startsAt(a.path[b+1:], start) {
			if b == 2 && a.path[b] == bp.D.Body {
				return blueprintKeywords, true
			}
			if inCheck(a.path[:b+1]) {
				return s.checkCandidates(a), true
			}
		}
	}
	// `self.|`.
	if prev != nil && prev.Kind == parser.MMSLexerDOT {
		if recv := prevToken(f, prev.Sp.Start); recv != nil && string(f.Src[recv.Sp.Start:recv.Sp.End]) == "self" {
			return selfCandidates, true
		}
		return nil, false
	}
	if prev != nil && (exprStart(prev.Kind) || prev.Kind == parser.MMSLexerLPAREN || prev.Kind == parser.MMSLexerCOMMA || prev.Kind == parser.MMSLexerLBRACK) {
		return s.blueprintValuesAt(a, bp, off), true
	}
	return nil, false
}

// blueprintValuesAt is blueprintValues without the lets a value at off
// cannot use yet: in `let x = |`, x itself and the lets after it.
func (s *Session) blueprintValuesAt(a *at, bp *eval.Blueprint, off int) []sema.Candidate {
	var inLet *syntax.LetStmt
	for _, n := range a.path {
		if l, ok := n.(*syntax.LetStmt); ok {
			inLet = l
		}
	}
	if inLet == nil {
		for _, l := range bp.Lets {
			// A let without its value yet ends at its `=`.
			if l.Sp.Start <= off && (off <= l.Sp.End || lineOf(a.u.File, l.Sp.Start) == lineOf(a.u.File, off)) {
				inLet = l
			}
		}
	}
	all := s.blueprintValues(bp)
	if inLet == nil {
		return all
	}
	later := map[string]bool{}
	for _, l := range bp.Lets {
		if l.Name != nil && l.Sp.Start >= inLet.Sp.Start {
			later[l.Name.Text] = true
		}
	}
	out := all[:0:0]
	for _, c := range all {
		if !(later[c.Label] && c.Detail != "" && strings.HasPrefix(c.Detail, "let")) {
			out = append(out, c)
		}
	}
	return out
}

// blueprintInterpolation completes a `${...}` in a Blueprint's strings
// and templated names, `Objective ${self.|}_cd`, which the Blueprint
// fills in when the pack is built: its fields, lets and self, and after
// `self.` its parts. ok is false when off is not in an open `${` there.
func (s *Session) blueprintInterpolation(a *at, off int) ([]Item, syntax.Span, bool) {
	if a.decl == nil || a.decl.Emitter == nil {
		return nil, syntax.Span{}, false
	}
	f := a.u.File
	i := f.TokenAt(off)
	if i < 0 {
		return nil, syntax.Span{}, false
	}
	t := f.Tokens[i]
	switch {
	case t.Kind == parser.MMSLexerTNAME && t.Sp.Start < off && off <= t.Sp.End:
	case t.Kind == parser.MMSLexerSTRING && t.Sp.Start < off && off < t.Sp.End && !a.code():
	default:
		return nil, syntax.Span{}, false
	}
	src := f.Src
	open, depth := -1, 0
	for j := t.Sp.Start; j < off; j++ {
		switch {
		case open < 0 && src[j] == '$' && j+1 < off && src[j+1] == '{':
			open, depth = j+2, 1
			j++
		case open >= 0 && src[j] == '{':
			depth++
		case open >= 0 && src[j] == '}':
			if depth--; depth == 0 {
				open = -1
			}
		}
	}
	if open < 0 {
		return nil, syntax.Span{}, false
	}
	start := off
	for start > open && identRune(src[start-1]) {
		start--
	}
	var cands []sema.Candidate
	if start > open && src[start-1] == '.' {
		if strings.TrimSpace(string(src[receiverStart(src, open, start-1):start-1])) == "self" {
			cands = selfCandidates
		}
	} else {
		cands = s.blueprintValuesAt(a, a.decl.Emitter, off)
	}
	end := wordEnd(src, off)
	if end > t.Sp.End {
		end = t.Sp.End
	}
	return candidateItems(cands), syntax.Span{Start: start, End: end}, true
}

// innermostBlock is the index of the last Block on path, or -1.
func innermostBlock(path []syntax.Node) int {
	for i := len(path) - 1; i >= 0; i-- {
		if _, ok := path[i].(*syntax.Block); ok {
			return i
		}
	}
	return -1
}

// startsAt reports whether every node starts at off: the word being
// typed is the first thing in each, as in a statement of one word.
func startsAt(nodes []syntax.Node, off int) bool {
	for _, n := range nodes {
		if n.Span().Start != off {
			return false
		}
	}
	return true
}
