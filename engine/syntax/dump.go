package syntax

import (
	"fmt"
	"strconv"
	"strings"
)

// Dump renders the AST as an indented outline, for golden tests and
// debugging.
func Dump(f *File) string {
	var d dumper
	for _, it := range f.Items {
		d.item(it)
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

func (d *dumper) item(it Item) {
	switch x := it.(type) {
	case *Decl:
		d.decl(x)
	case *TopFor:
		name := "?"
		if x.Var != nil {
			name = x.Var.Text
		}
		if x.Index != nil {
			name = x.Index.Text + ", " + name
		}
		d.line("TopFor %s", name)
		d.in(func() { d.expr(x.Iter); d.item(x.Body) })
	case *TopIf:
		d.line("TopIf")
		d.in(func() {
			d.expr(x.Cond)
			d.item(x.Then)
			if x.Else != nil {
				d.line("Else")
				d.in(func() { d.item(x.Else) })
			}
		})
	case *TopLet:
		name := "?"
		if x.Name != nil {
			name = x.Name.Text
		}
		d.line("TopLet %s", name)
		if x.Value != nil {
			d.in(func() { d.expr(x.Value) })
		}
	case *TopBlock:
		d.line("TopBlock")
		d.in(func() {
			for _, c := range x.Items {
				d.item(c)
			}
		})
	}
}

func (d *dumper) decl(decl *Decl) {
	head := decl.Kind.Name
	if decl.Name != nil {
		if decl.Name.Quoted {
			head += " " + strconv.Quote(decl.Name.Text)
		} else {
			head += " " + decl.Name.Text
		}
	}
	if decl.Type != nil {
		head += ": " + decl.Type.Text
	}
	if decl.Params != nil {
		var ps []string
		for _, p := range decl.Params.Params {
			s := "?"
			if p.Name != nil {
				s = p.Name.Text
			}
			if p.Type != nil {
				s += ": " + p.Type.Text
			}
			ps = append(ps, s)
		}
		head += "(" + strings.Join(ps, ", ") + ")"
	}
	d.line("Decl %s", head)
	d.in(func() {
		if decl.Value != nil {
			d.expr(decl.Value)
		}
		if decl.Body != nil {
			d.block(decl.Body)
		}
	})
}

func (d *dumper) block(b *Block) {
	if b.Closed {
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
	case *LetStmt:
		name := "?"
		if s.Name != nil {
			name = s.Name.Text
		}
		d.line("Let %s", name)
		if s.Value != nil {
			d.in(func() { d.expr(s.Value) })
		}
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
