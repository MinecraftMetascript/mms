package ide

import (
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
