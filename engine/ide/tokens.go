package ide

import (
	"sort"

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
)

// TokenTypes names the token types for the legend, in order.
var TokenTypes = []string{"keyword", "comment", "string", "number", "operator", "type", "variable", "function", "method", "property", "enumMember", "namespace", "macro"}

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
		case parser.MMSLexerLET, parser.MMSLexerIF, parser.MMSLexerELSE, parser.MMSLexerWHILE, parser.MMSLexerFOR,
			parser.MMSLexerIN, parser.MMSLexerRETURN, parser.MMSLexerBREAK, parser.MMSLexerCONTINUE,
			parser.MMSLexerTRUE, parser.MMSLexerFALSE:
			set(t.Sp, TokKeyword, 0)
		}
	}
	for _, d := range f.Decls {
		dr := s.R.Decls[d]
		info := func(n syntax.Node) bool {
			return dr != nil && dr.Info[n] != nil && dr.Info[n].Symbol != nil
		}
		syntax.Walk(d, func(n syntax.Node) bool {
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
					set(id.Sp, TokFunction, ModDefaultLibrary)
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
				if info(n) {
					set(n.Sp, TokVariable, 0)
				} else {
					set(n.Sp, TokEnumMember, 0)
				}
			case *syntax.ResLoc:
				if info(n) {
					set(n.Sp, TokVariable, 0)
				} else {
					set(n.Sp, TokEnumMember, 0)
				}
			case *syntax.Tag:
				set(n.Sp, TokVariable, 0)
			}
			return true
		})
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
