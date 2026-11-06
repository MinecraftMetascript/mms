package ide

import (
	"sort"
	"strings"

	"github.com/minecraftmetascript/mms/engine/prelude"
	"github.com/minecraftmetascript/mms/engine/syntax"
	"github.com/minecraftmetascript/mms/engine/syntax/parser"
)

// TokenType is a semantic token type. The order is the legend the
// language server advertises.
type TokenType int

const (
	TokKeyword TokenType = iota
	TokComment
	TokString
	TokNumber
	TokOperator
	TokType
	TokVariable
	TokFunction
	TokMethod
	TokProperty
	TokEnumMember
	TokNamespace
	TokMacro
	TokParameter
)

// TokenTypes names the token types for the legend, in order.
var TokenTypes = []string{"keyword", "comment", "string", "number", "operator", "type", "variable", "function", "method", "property", "enumMember", "namespace", "macro", "parameter"}

// Token modifiers, as bits.
const (
	ModDeclaration = 1 << iota
	ModDocumentation
	ModDefaultLibrary
)

// TokenModifiers names the modifier bits for the legend, in order.
var TokenModifiers = []string{"declaration", "documentation", "defaultLibrary"}

// SemToken is one classified span.
type SemToken struct {
	Span syntax.Span
	Type TokenType
	Mods int
}

// SemanticTokens classifies a file's tokens, using the syntax tree to tell
// kinds, helpers, methods, fields and references apart. Tokens are sorted
// and do not overlap.
func (s *Session) SemanticTokens(path string) []SemToken {
	u := s.R.Unit(path)
	if u == nil {
		return nil
	}
	f := u.File
	byStart := map[int]SemToken{}
	set := func(sp syntax.Span, t TokenType, mods int) {
		if sp.End > sp.Start {
			byStart[sp.Start] = SemToken{Span: sp, Type: t, Mods: mods}
		}
	}
	for _, t := range f.Tokens {
		if t.Synthetic {
			continue
		}
		switch t.Kind {
		case parser.MMSLexerDOC_COMMENT:
			set(t.Sp, TokComment, ModDocumentation)
		case parser.MMSLexerLINE_COMMENT, parser.MMSLexerBLOCK_COMMENT, parser.MMSLexerUNTERMINATED_COMMENT:
			set(t.Sp, TokComment, 0)
		case parser.MMSLexerSTRING:
			set(t.Sp, TokString, 0)
		case parser.MMSLexerNUMBER:
			set(t.Sp, TokNumber, 0)
		case parser.MMSLexerCOMMAND:
			set(t.Sp, TokMacro, 0)
		case parser.MMSLexerCONST, parser.MMSLexerEXTERN, parser.MMSLexerLET, parser.MMSLexerIF, parser.MMSLexerELSE,
			parser.MMSLexerWHILE, parser.MMSLexerFOR, parser.MMSLexerIN, parser.MMSLexerRETURN, parser.MMSLexerBREAK,
			parser.MMSLexerCONTINUE, parser.MMSLexerPRIMARY, parser.MMSLexerAS, parser.MMSLexerCHECK, parser.MMSLexerEMIT,
			parser.MMSLexerSETUP, parser.MMSLexerEXPECT, parser.MMSLexerRUN, parser.MMSLexerTRUE, parser.MMSLexerFALSE:
			set(t.Sp, TokKeyword, 0)
		}
	}
	// The checked declarations know which names are symbols. Nodes the
	// evaluator wrote at a call (a Template's result, a const's value)
	// sit at the call's span, so they are skipped: the source pass below
	// classifies what is written there.
	copied := func(n syntax.Node) bool {
		if s.R.Eval == nil {
			return false
		}
		_, ok := s.R.Eval.Nodes[n]
		return ok
	}
	for _, d := range f.Decls {
		dr := s.R.Decls[d]
		info := func(n syntax.Node) bool {
			return dr != nil && dr.Info[n] != nil && dr.Info[n].Symbol != nil
		}
		syntax.Walk(d, func(n syntax.Node) bool {
			if copied(n) {
				return false
			}
			s.classify(n, set, byStart, info)
			return true
		})
	}
	// The source as written: const, Template, Blueprint, Type and for
	// items, and the calls and names inside them, which expansion
	// replaced. It only fills spans the checked pass left.
	if src := s.R.Source[path]; src != nil {
		fill := func(sp syntax.Span, t TokenType, mods int) {
			if _, done := byStart[sp.Start]; !done {
				set(sp, t, mods)
			}
		}
		for _, it := range src.Items {
			s.sourceTokens(it, fill, byStart)
		}
	}
	out := make([]SemToken, 0, len(byStart))
	for _, t := range byStart {
		out = append(out, t)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Span.Start < out[j].Span.Start })
	// Drop overlaps (a Name inside a larger span), keeping the earlier one.
	kept := out[:0]
	end := -1
	for _, t := range out {
		if t.Span.Start >= end {
			kept = append(kept, t)
			end = t.Span.End
		}
	}
	return kept
}

// classify sets the token for one node of a checked declaration. info
// reports whether a name resolved to a symbol.
func (s *Session) classify(n syntax.Node, set func(syntax.Span, TokenType, int), byStart map[int]SemToken, info func(syntax.Node) bool) {
	switch n := n.(type) {
	case *syntax.Decl:
		if n.Kind != nil {
			set(n.Kind.Sp, TokType, 0)
		}
		if n.Name != nil {
			t := TokVariable
			if n.Kind != nil && n.Kind.Name == "Namespace" {
				t = TokNamespace
			}
			set(n.Name.Sp, t, ModDeclaration)
		}
	case *syntax.Call:
		if id, ok := n.Fn.(*syntax.Ident); ok {
			set(id.Sp, TokFunction, s.libraryMod(id.Name))
		}
	case *syntax.Member:
		if n.Name != nil {
			set(n.Name.Sp, TokMethod, 0)
		}
	case *syntax.FieldStmt:
		if !n.KeyQuoted {
			set(n.Key.Sp, TokProperty, 0)
		}
	case *syntax.Arg:
		if n.Name != nil {
			set(n.Name.Sp, TokProperty, 0)
		}
	case *syntax.Ident:
		if _, done := byStart[n.Sp.Start]; done {
			break
		}
		s.nameToken(n, set, info)
	case *syntax.ResLoc:
		s.nameToken(n, set, info)
	case *syntax.Tag:
		set(n.Sp, TokVariable, 0)
	case *syntax.TypeName:
		typeName(n, set)
	}
}

// nameToken marks a name that resolved to a symbol as a variable. One
// that did not is left to the source pass, which knows the consts and
// parameters a substituted name came from, unless there is no source.
func (s *Session) nameToken(n syntax.Node, set func(syntax.Span, TokenType, int), info func(syntax.Node) bool) {
	switch {
	case info(n):
		set(n.Span(), TokVariable, 0)
	case s.R.Source == nil:
		set(n.Span(), TokEnumMember, 0)
	}
}

// libraryMod marks a call of a built-in helper or prelude Template as
// part of the default library, and a call of the user's own Template not.
func (s *Session) libraryMod(name string) int {
	if s.R.Eval != nil {
		if t := s.R.Eval.Templates[name]; t != nil && !strings.HasPrefix(t.Path, prelude.Prefix) {
			return 0
		}
	}
	return ModDefaultLibrary
}

// typeName marks the name of a type, without its arguments.
func typeName(n *syntax.TypeName, set func(syntax.Span, TokenType, int)) {
	end := n.Sp.Start + len([]rune(n.Name))
	if end > n.Sp.End {
		end = n.Sp.End
	}
	set(syntax.Span{Start: n.Sp.Start, End: end}, TokType, 0)
}

// sourceTokens classifies one top-level item as written. Names bound in
// it (parameters, fields, loop and let variables, self) are variables or
// parameters; consts and Templates are found in the evaluation.
func (s *Session) sourceTokens(it syntax.Stmt, set func(syntax.Span, TokenType, int), byStart map[int]SemToken) {
	params := map[string]bool{}
	locals := map[string]bool{"self": true}
	syntax.Walk(it, func(n syntax.Node) bool {
		switch n := n.(type) {
		case *syntax.Param:
			if n.Name != nil {
				params[n.Name.Text] = true
			}
		case *syntax.FieldDecl:
			if n.Name != nil {
				params[n.Name.Text] = true
			}
		case *syntax.ForStmt:
			for _, v := range []*syntax.Name{n.Var, n.Var2} {
				if v != nil {
					locals[v.Text] = true
				}
			}
		case *syntax.Comprehension:
			for _, v := range []*syntax.Name{n.Var, n.Var2} {
				if v != nil {
					locals[v.Text] = true
				}
			}
		case *syntax.LetStmt:
			if n.Name != nil {
				locals[n.Name.Text] = true
			}
		}
		return true
	})
	named := func(name string) (TokenType, bool) {
		switch {
		case params[name]:
			return TokParameter, true
		case locals[name]:
			return TokVariable, true
		}
		if s.R.Eval != nil && s.R.Eval.Consts[name] != nil {
			return TokVariable, true
		}
		return 0, false
	}
	syntax.Walk(it, func(n syntax.Node) bool {
		switch n := n.(type) {
		case *syntax.Decl:
			if n.Kind != nil {
				set(n.Kind.Sp, TokType, 0)
			}
			if n.Name != nil && !n.Name.Templated() {
				t := TokVariable
				if n.Kind != nil {
					switch n.Kind.Name {
					case "Template":
						t = TokFunction
					case "Blueprint", "Type":
						t = TokType
					case "Namespace":
						t = TokNamespace
					}
				}
				set(n.Name.Sp, t, ModDeclaration)
			}
		case *syntax.ConstStmt:
			if n.Name != nil {
				set(n.Name.Sp, TokVariable, ModDeclaration)
			}
		case *syntax.LetStmt:
			if n.Name != nil {
				set(n.Name.Sp, TokVariable, ModDeclaration)
			}
		case *syntax.Param:
			if n.Name != nil {
				set(n.Name.Sp, TokParameter, ModDeclaration)
			}
		case *syntax.FieldDecl:
			if n.Name != nil {
				set(n.Name.Sp, TokProperty, ModDeclaration)
			}
		case *syntax.ForStmt:
			for _, v := range []*syntax.Name{n.Var, n.Var2} {
				if v != nil {
					set(v.Sp, TokVariable, ModDeclaration)
				}
			}
		case *syntax.Comprehension:
			for _, v := range []*syntax.Name{n.Var, n.Var2} {
				if v != nil {
					set(v.Sp, TokVariable, ModDeclaration)
				}
			}
		case *syntax.SectionStmt:
			if n.Word != nil {
				set(n.Word.Sp, TokKeyword, 0)
			}
		case *syntax.Call:
			if id, ok := n.Fn.(*syntax.Ident); ok {
				set(id.Sp, TokFunction, s.libraryMod(id.Name))
			}
		case *syntax.Member:
			if n.Name != nil {
				set(n.Name.Sp, TokProperty, 0)
			}
		case *syntax.FieldStmt:
			if !n.KeyQuoted && n.Key != nil {
				set(n.Key.Sp, TokProperty, 0)
			}
		case *syntax.Arg:
			if n.Name != nil {
				set(n.Name.Sp, TokProperty, 0)
			}
		case *syntax.TypeName:
			typeName(n, set)
		case *syntax.Ident:
			if t, ok := named(n.Name); ok {
				set(n.Sp, t, 0)
			} else {
				set(n.Sp, TokEnumMember, 0)
			}
		case *syntax.ResLoc:
			set(n.Sp, TokEnumMember, 0)
		case *syntax.Tag:
			set(n.Sp, TokVariable, 0)
		}
		return true
	})
}
