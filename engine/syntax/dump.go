package syntax

import (
	"fmt"
	"strings"
)

// Dump renders the AST as an indented outline, for golden tests and
// debugging.
func Dump(f *File) string {
	var d dumper
	for _, it := range f.Items {
		d.stmt(it)
	}
	for _, diag := range f.Diags {
		d.line("! %d-%d %s", diag.Sp.Start, diag.Sp.End, diag.Message)
	}
	return d.String()
}

type dumper struct {
	strings.Builder
	depth int
}

func (d *dumper) line(format string, args ...any) {
	d.WriteString(strings.Repeat("  ", d.depth))
	fmt.Fprintf(d, format, args...)
	d.WriteByte('\n')
}

func (d *dumper) in(f func()) { d.depth++; f(); d.depth-- }

func (d *dumper) decl(decl *Decl) {
	head := decl.Kind.Name
	if decl.Extern {
		head = "extern " + head
	}
	if decl.Primary {
		head = "primary " + head
	}
	for i, n := range decl.Names() {
		if i > 0 {
			head += ","
		}
		head += " " + d.name(n)
	}
	if decl.Type != nil {
		head += ": " + TypeText(decl.Type)
	}
	if decl.Params != nil {
		var ps []string
		for _, p := range decl.Params.Params {
			s := "?"
			if p.Name != nil {
				s = p.Name.Text
			}
			if p.Type != nil {
				s += ": " + TypeText(p.Type)
			}
			if p.Default != nil {
				s += " = ..."
			}
			ps = append(ps, s)
		}
		head += "(" + strings.Join(ps, ", ") + ")"
	}
	if decl.TypeValue != nil {
		head += " = type " + TypeText(decl.TypeValue)
	}
	d.line("Decl %s", head)
	d.in(func() {
		if decl.Params != nil {
			for _, p := range decl.Params.Params {
				if p.Default != nil {
					d.line("Default %s", p.Name.Text)
					d.in(func() { d.expr(p.Default) })
				}
			}
		}
		for _, n := range decl.Names() {
			d.parts(n.Parts)
		}
		if decl.Value != nil {
			d.expr(decl.Value)
		}
		if decl.Body != nil {
			d.block(decl.Body)
		}
	})
}

// name renders a name; a templated one is marked so its pieces follow.
func (d *dumper) name(n *Name) string {
	if n.Templated() {
		return n.Text + " (templated)"
	}
	return n.Text
}

func (d *dumper) parts(ps []NamePart) {
	for _, p := range ps {
		if p.X == nil {
			d.line("Lit %q", p.Lit)
			continue
		}
		d.line("Interp")
		d.in(func() { d.expr(p.X) })
	}
}

func (d *dumper) typed(label string, t TypeExpr) {
	if t != nil {
		d.line("%s: %s", label, TypeText(t))
	}
}

func (d *dumper) block(b *Block) {
	if b.Bare {
		d.line("Block (bare)")
	} else if b.Closed {
		d.line("Block")
	} else {
		d.line("Block (unclosed)")
	}
	d.in(func() {
		for _, s := range b.Items {
			d.stmt(s)
		}
	})
}

func (d *dumper) stmt(s Stmt) {
	switch s := s.(type) {
	case *Decl:
		d.decl(s)
	case *LetStmt:
		name := "?"
		if s.Name != nil {
			name = s.Name.Text
		}
		d.line("Let %s", name)
		d.in(func() {
			d.typed("Type", s.Type)
			if s.Value != nil {
				d.expr(s.Value)
			}
		})
	case *ConstStmt:
		name := "?"
		if s.Name != nil {
			name = s.Name.Text
		}
		d.line("Const %s", name)
		d.in(func() {
			d.typed("Type", s.Type)
			if s.Value != nil {
				d.expr(s.Value)
			}
		})
	case *FieldDecl:
		d.line("FieldDecl %s: %s", s.Name.Text, TypeText(s.Type))
		if s.Default != nil {
			d.in(func() { d.expr(s.Default) })
		}
	case *SectionStmt:
		d.line("Section %s", s.Word.Name)
		d.in(func() { d.block(s.Body) })
	case *AsStmt:
		d.line("As %s", TypeText(s.Type))
		d.in(func() {
			d.expr(s.Value)
			if s.Hint != nil {
				d.line("Hint")
				d.in(func() { d.args(s.Hint.Args) })
			}
		})
	case *ExpectStmt:
		d.line("Expect")
		d.in(func() { d.expr(s.X) })
	case *RunStmt:
		d.line("Run")
		d.in(func() { d.expr(s.X) })
	case *IfStmt:
		d.line("If")
		d.in(func() {
			d.expr(s.Cond)
			d.block(s.Then)
			if s.Else != nil {
				d.line("Else")
				d.in(func() { d.stmt(s.Else) })
			}
		})
	case *WhileStmt:
		d.line("While")
		d.in(func() { d.expr(s.Cond); d.block(s.Body) })
	case *ForStmt:
		name := "?"
		if s.Var != nil {
			name = s.Var.Text
		}
		if s.Var2 != nil {
			name += ", " + s.Var2.Text
		}
		d.line("For %s", name)
		d.in(func() { d.expr(s.Iter); d.block(s.Body) })
	case *ReturnStmt:
		d.line("Return")
		if s.Value != nil {
			d.in(func() { d.expr(s.Value) })
		}
	case *BreakStmt:
		d.line("Break")
	case *ContinueStmt:
		d.line("Continue")
	case *CommandStmt:
		d.line("Command %q", s.Text)
	case *FieldStmt:
		d.line("Field %s", s.Key.Text)
		d.in(func() { d.expr(s.Value) })
	case *AssignStmt:
		d.line("Assign %s", s.Op)
		d.in(func() { d.expr(s.Target); d.expr(s.Value) })
	case *ExprStmt:
		d.expr(s.X)
	case *Block:
		d.block(s)
	}
}

func (d *dumper) expr(e Expr) {
	switch e := e.(type) {
	case *Number:
		d.line("Number %s", e.Text)
	case *String:
		d.line("String %q", e.Value)
	case *Bool:
		d.line("Bool %v", e.Value)
	case *Ident:
		d.line("Ident %s", e.Name)
	case *ResLoc:
		d.line("ResLoc %s", e.Text)
	case *Tag:
		d.line("Tag %s", e.Name)
	case *Selector:
		d.line("Selector %s", e.Name)
	case *Coord:
		d.line("Coord %s", e.Text)
	case *Paren:
		d.line("Paren")
		d.in(func() { d.expr(e.X) })
	case *List:
		d.line("List")
		d.in(func() {
			for _, x := range e.Elems {
				d.expr(x)
			}
		})
	case *Object:
		d.line("Object")
		d.in(func() { d.block(e.Body) })
	case *Unary:
		d.line("Unary %s", e.Op)
		d.in(func() { d.expr(e.X) })
	case *Binary:
		d.line("Binary %s", e.Op)
		d.in(func() { d.expr(e.L); d.expr(e.R) })
	case *Range:
		d.line("Range")
		d.in(func() {
			if e.Lo != nil {
				d.expr(e.Lo)
			} else {
				d.line("-")
			}
			if e.Hi != nil {
				d.expr(e.Hi)
			} else {
				d.line("-")
			}
		})
	case *Member:
		name := "<missing>"
		if e.Name != nil {
			name = e.Name.Text
		}
		d.line("Member .%s", name)
		d.in(func() { d.expr(e.X) })
	case *Call:
		d.line("Call")
		d.in(func() {
			d.expr(e.Fn)
			d.args(e.Args)
		})
	case *Index:
		d.line("Index")
		d.in(func() {
			d.expr(e.X)
			d.args(e.Args)
		})
	case *Trailing:
		d.line("Trailing")
		d.in(func() { d.expr(e.X); d.block(e.Block) })
	case *TemplatedName:
		d.line("TemplatedName %s", e.Text)
		d.in(func() { d.parts(e.Parts) })
	case *Spread:
		d.line("Spread")
		d.in(func() { d.expr(e.X) })
	case *IfExpr:
		d.line("IfExpr")
		d.in(func() {
			d.expr(e.Cond)
			d.expr(e.Then)
			if e.Else != nil {
				d.line("Else")
				d.in(func() { d.expr(e.Else) })
			}
		})
	case *RunExpr:
		d.line("Run")
		d.in(func() { d.block(e.Body) })
	case *Comprehension:
		name := "?"
		if e.Var != nil {
			name = e.Var.Text
		}
		if e.Var2 != nil {
			name += ", " + e.Var2.Text
		}
		d.line("Comprehension %s", name)
		d.in(func() {
			d.expr(e.Iter)
			if e.Cond != nil {
				d.line("If")
				d.in(func() { d.expr(e.Cond) })
			}
			d.expr(e.Body)
		})
	case *Bad:
		d.line("Bad %d-%d", e.Sp.Start, e.Sp.End)
	case nil:
		d.line("<nil>")
	}
}

func (d *dumper) args(args []*Arg) {
	for _, a := range args {
		if a.Name != nil {
			d.line("Arg %s =", a.Name.Text)
			d.in(func() { d.expr(a.Value) })
		} else {
			d.expr(a.Value)
		}
	}
}
