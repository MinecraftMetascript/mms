package ide

import (
	"strings"

	"github.com/minecraftmetascript/mms/engine/syntax"
)

// evalExtractions offers the language's own extractions for a selection,
// read from the source as written: a value to a const, and statements of
// a function to a Commands Template. Neither is offered for text that
// reads a parameter or local, which a top-level declaration cannot see.
func (s *Session) evalExtractions(path string, from, to int) []Extraction {
	src := s.R.Source[path]
	if src == nil || s.R.Eval == nil || from >= to {
		return nil
	}
	var item syntax.Stmt
	for _, it := range src.Items {
		if sp := it.Span(); sp.Start <= from && to <= sp.End {
			item = it
		}
	}
	if item == nil {
		return nil
	}
	bound := boundNames(item)
	at := declStart(src.Src, item.Span().Start)
	var out []Extraction
	if x, parent := selectedExpr(item, from, to); x != nil && !readsBound(src.Src, x, bound) {
		name := s.newEvalName(constName(parent))
		head := "const " + name + " = "
		out = append(out, Extraction{
			Title:  "Extract " + preview(x, src.Src) + " to a const",
			Kind:   "const",
			Name:   name,
			NameAt: at + len("const "),
			Edits: []Edit{
				{Path: path, Span: syntax.Span{Start: at, End: at}, Text: head + reindent(src.Src, x.Span()) + "\n\n"},
				{Path: path, Span: x.Span(), Text: name},
			},
		})
	}
	if d, ok := item.(*syntax.Decl); ok && d.Kind != nil && syntax.IsCodeKind(d.Kind.Name) {
		if stmts := selectedStmts(d, from, to); len(stmts) > 0 && !stmtsReadBound(src.Src, stmts, bound) {
			name := s.newEvalName("commands")
			sp := syntax.Span{Start: stmts[0].Span().Start, End: stmts[len(stmts)-1].Span().End}
			var body []string
			for _, st := range stmts {
				body = append(body, "  "+reindent(src.Src, st.Span()))
			}
			out = append(out, Extraction{
				Title:  "Extract " + statementsLabel(len(stmts)) + " to a Template",
				Kind:   "Template",
				Name:   name,
				NameAt: at + len("Template "),
				Edits: []Edit{
					{Path: path, Span: syntax.Span{Start: at, End: at}, Text: "Template " + name + "(): Commands {\n" + strings.Join(body, "\n") + "\n}\n\n"},
					{Path: path, Span: sp, Text: name + "()"},
				},
			})
		}
	}
	return out
}

func statementsLabel(n int) string {
	if n == 1 {
		return "this statement"
	}
	return "these statements"
}

// selectedExpr finds the expression the selection covers: the smallest
// one that holds it. A bare name is already a name, and a const's whole
// value is already a const.
func selectedExpr(item syntax.Stmt, from, to int) (syntax.Expr, syntax.Node) {
	var best syntax.Expr
	var parent, last syntax.Node
	syntax.Walk(item, func(n syntax.Node) bool {
		sp := n.Span()
		if sp.Start > from || sp.End < to {
			return false
		}
		if x, ok := n.(syntax.Expr); ok {
			best, parent = x, last
		}
		last = n
		return true
	})
	switch best.(type) {
	case nil, *syntax.Ident, *syntax.Bad, *syntax.RunExpr:
		return nil, nil
	}
	if k, ok := item.(*syntax.ConstStmt); ok && k.Value == best {
		return nil, nil
	}
	// Only a selection that is the expression, give or take spaces.
	sp := best.Span()
	if from > sp.Start+1 || to < sp.End-1 {
		return nil, nil
	}
	return best, parent
}

// selectedStmts finds the statements of a function body the selection
// covers, whole, in one block.
func selectedStmts(d *syntax.Decl, from, to int) []syntax.Stmt {
	var out []syntax.Stmt
	syntax.Walk(d, func(n syntax.Node) bool {
		b, ok := n.(*syntax.Block)
		if !ok || len(out) > 0 {
			return len(out) == 0
		}
		var in []syntax.Stmt
		for _, st := range b.Items {
			sp := st.Span()
			if sp.Start >= from && sp.End <= to {
				in = append(in, st)
			} else if sp.Start < to && sp.End > from {
				in = nil // the selection cuts this statement
				break
			}
		}
		if len(in) > 0 {
			out = in
			return false
		}
		return true
	})
	for _, st := range out {
		switch st.(type) {
		case *syntax.LetStmt, *syntax.ReturnStmt, *syntax.BreakStmt, *syntax.ContinueStmt:
			return nil // they mean something only where they are
		}
	}
	return out
}

// readsBound reports whether x names a parameter or local of its item,
// in code or inside ${...}.
func readsBound(src []rune, x syntax.Node, bound map[string]bool) bool {
	found := false
	syntax.Walk(x, func(n syntax.Node) bool {
		switch n := n.(type) {
		case *syntax.Ident:
			found = found || bound[n.Name]
		case *syntax.String, *syntax.CommandStmt:
			for _, p := range pieces(src, n.Span()) {
				found = found || bound[string(src[p.Start:p.End])]
			}
		}
		return !found
	})
	return found
}

func stmtsReadBound(src []rune, stmts []syntax.Stmt, bound map[string]bool) bool {
	for _, st := range stmts {
		if readsBound(src, st, bound) {
			return true
		}
	}
	return false
}

// constName names an extracted value after the field or argument it
// fills, or `value`.
func constName(parent syntax.Node) string {
	name := ""
	switch p := parent.(type) {
	case *syntax.FieldStmt:
		if p.Key != nil {
			name = p.Key.Text
		}
	case *syntax.Arg:
		if p.Name != nil {
			name = p.Name.Text
		}
	}
	var b strings.Builder
	for _, r := range strings.ToLower(name) {
		if r >= 'a' && r <= 'z' || r >= '0' && r <= '9' || r == '_' {
			b.WriteRune(r)
		}
	}
	if n := b.String(); n != "" && (n[0] < '0' || n[0] > '9') {
		return n
	}
	return "value"
}

// newEvalName makes name unique among the consts, Templates, Types and
// Blueprints.
func (s *Session) newEvalName(name string) string {
	taken := map[string]bool{}
	e := s.R.Eval
	for _, k := range e.Consts {
		taken[k.Name()] = true
	}
	for _, t := range e.Templates {
		taken[t.D.Name.Text] = true
	}
	for _, b := range e.Blueprints {
		taken[b.D.Name.Text] = true
	}
	for _, t := range e.Types {
		taken[t.D.Name.Text] = true
	}
	return uniqueName(taken, name)
}
