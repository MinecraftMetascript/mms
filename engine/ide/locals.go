package ide

import (
	"strings"

	"github.com/minecraftmetascript/mms/engine/format"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// local is a name an item binds itself: a Template parameter, a Blueprint
// field or let, a let in a body, or a loop variable.
type local struct {
	Name  *syntax.Name
	What  string // "parameter", "field", "let" or "loop variable"
	Type  syntax.TypeExpr
	Value syntax.Expr // the default or the let's value, or nil
	Doc   string      // Markdown
	Scope syntax.Span // where uses see it
}

// bindings lists the names item binds, outermost first.
func bindings(it syntax.Node) []local {
	var out []local
	var walk func(n syntax.Node, parent syntax.Node, decl *syntax.Decl)
	walk = func(n, parent syntax.Node, decl *syntax.Decl) {
		switch x := n.(type) {
		case *syntax.Decl:
			decl = x
		case *syntax.Param:
			if x.Name != nil && decl != nil {
				doc := ""
				if decl.Doc != "" {
					doc = syntax.ParseDoc(decl.Doc).Param(x.Name.Text)
				}
				out = append(out, local{Name: x.Name, What: "parameter", Type: x.Type, Value: x.Default, Doc: doc, Scope: decl.Sp})
			}
		case *syntax.FieldDecl:
			if x.Name != nil && decl != nil {
				doc, _ := docText(x.Doc)
				out = append(out, local{Name: x.Name, What: "field", Type: x.Type, Value: x.Default, Doc: doc, Scope: decl.Sp})
			}
		case *syntax.LetStmt:
			if x.Name != nil {
				doc, _ := docText(x.Doc)
				scope := x.Sp
				if parent != nil {
					scope = syntax.Span{Start: x.Sp.Start, End: parent.Span().End}
				}
				if _, ok := parent.(*syntax.Decl); ok && decl != nil && decl.Kind != nil && decl.Kind.Name == "Blueprint" {
					scope = decl.Sp // a Blueprint's lets are seen by its whole body
				}
				out = append(out, local{Name: x.Name, What: "let", Type: x.Type, Value: x.Value, Doc: doc, Scope: scope})
			}
		case *syntax.ForStmt:
			for _, v := range []*syntax.Name{x.Var, x.Var2} {
				if v != nil {
					out = append(out, local{Name: v, What: "loop variable", Scope: x.Sp})
				}
			}
		case *syntax.Comprehension:
			for _, v := range []*syntax.Name{x.Var, x.Var2} {
				if v != nil {
					out = append(out, local{Name: v, What: "loop variable", Scope: x.Sp})
				}
			}
		}
		for _, c := range syntax.Children(n) {
			walk(c, n, decl)
		}
	}
	walk(it, nil, nil)
	return out
}

// resolveLocal finds the binding a use of name at off sees: the one with
// the narrowest scope around it.
func resolveLocal(ls []local, name string, off int) (local, bool) {
	var best local
	ok := false
	for _, l := range ls {
		if l.Name.Text != name || !within(l.Scope, off) {
			continue
		}
		if !ok || l.Scope.End-l.Scope.Start <= best.Scope.End-best.Scope.Start {
			best, ok = l, true
		}
	}
	return best, ok
}

// localAt finds the local at off, at its binding or at a use, with the
// span of the word there.
func (s *Session) localAt(path string, off int) (local, syntax.Span, bool) {
	src := s.R.Source[path]
	if src == nil {
		return local{}, syntax.Span{}, false
	}
	for _, it := range src.Items {
		if !within(it.Span(), off) {
			continue
		}
		ls := bindings(it)
		if len(ls) == 0 {
			continue
		}
		for _, l := range ls {
			if within(l.Name.Sp, off) {
				return l, l.Name.Sp, true
			}
		}
		var found local
		var sp syntax.Span
		ok := false
		syntax.Walk(it, func(n syntax.Node) bool {
			if !within(n.Span(), off) {
				return true
			}
			switch n := n.(type) {
			case *syntax.Ident:
				if within(n.Sp, off) {
					if l, ok2 := resolveLocal(ls, n.Name, off); ok2 {
						found, sp, ok = l, n.Sp, true
					}
				}
			case *syntax.String, *syntax.CommandStmt:
				for _, p := range pieces(src.Src, n.Span()) {
					if within(p, off) {
						if l, ok2 := resolveLocal(ls, string(src.Src[p.Start:p.End]), off); ok2 {
							found, sp, ok = l, p, true
						}
					}
				}
			}
			return true
		})
		if ok {
			return found, sp, true
		}
	}
	return local{}, syntax.Span{}, false
}

// localUses lists the binding of l and every use of it in path.
func (s *Session) localUses(path string, l local) []Location {
	src := s.R.Source[path]
	out := []Location{{Path: path, Span: l.Name.Sp}}
	if src == nil {
		return out
	}
	for _, it := range src.Items {
		if !within(it.Span(), l.Name.Sp.Start) {
			continue
		}
		ls := bindings(it)
		same := func(name string, at int) bool {
			r, ok := resolveLocal(ls, name, at)
			return ok && r.Name == l.Name
		}
		syntax.Walk(it, func(n syntax.Node) bool {
			switch n := n.(type) {
			case *syntax.Ident:
				if n.Name == l.Name.Text && same(n.Name, n.Sp.Start) {
					out = append(out, Location{Path: path, Span: n.Sp})
				}
			case *syntax.String, *syntax.CommandStmt:
				for _, p := range pieces(src.Src, n.Span()) {
					if string(src.Src[p.Start:p.End]) == l.Name.Text && same(l.Name.Text, p.Start) {
						out = append(out, Location{Path: path, Span: p})
					}
				}
			}
			return true
		})
	}
	return out
}

// localHover describes a local.
func localHover(l local) string {
	head := l.What + " " + l.Name.Text
	if l.Type != nil {
		head += ": " + syntax.TypeText(l.Type)
	}
	if l.Value != nil {
		v := format.Expr(l.Value)
		if len(v) > 120 {
			v = v[:120] + " …"
		}
		if i := strings.IndexByte(v, '\n'); i >= 0 {
			v = v[:i] + " …"
		}
		head += " = " + v
	}
	md := "```mms\n" + head + "\n```"
	if l.Doc != "" {
		md += "\n\n" + l.Doc
	}
	return md
}
