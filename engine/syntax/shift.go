package syntax

// shiftFile copies a parsed file with every span moved by delta, recording
// old-to-new nodes in nodes when it is not nil.
func shiftFile(f *File, delta int, nodes map[Node]Node) *File {
	s := shifter{d: delta, m: nodes}
	out := &File{Src: f.Src, Decls: make([]*Decl, len(f.Decls)), Tokens: make([]Token, len(f.Tokens)), Diags: make([]Diagnostic, len(f.Diags))}
	for i, t := range f.Tokens {
		t.Sp = s.sp(t.Sp)
		out.Tokens[i] = t
	}
	for i, d := range f.Diags {
		d.Sp = s.sp(d.Sp)
		out.Diags[i] = d
	}
	for i, d := range f.Decls {
		out.Decls[i] = s.decl(d)
	}
	return out
}

type shifter struct {
	d int
	m map[Node]Node
}

func (s shifter) sp(sp Span) Span { return Span{sp.Start + s.d, sp.End + s.d} }

func (s shifter) rec(old, new Node) {
	if s.m != nil {
		s.m[old] = new
	}
}

func (s shifter) decl(d *Decl) *Decl {
	n := &Decl{Kind: s.ident(d.Kind), Name: s.name(d.Name), Type: s.name(d.Type), Value: s.expr(d.Value), Body: s.block(d.Body), Doc: d.Doc, Sp: s.sp(d.Sp)}
	if d.Params != nil {
		n.Params = &ParamList{Sp: s.sp(d.Params.Sp)}
		for _, p := range d.Params.Params {
			np := &Param{Name: s.name(p.Name), Type: s.name(p.Type), Sp: s.sp(p.Sp)}
			s.rec(p, np)
			n.Params.Params = append(n.Params.Params, np)
		}
		s.rec(d.Params, n.Params)
	}
	s.rec(d, n)
	return n
}

func (s shifter) ident(x *Ident) *Ident {
	if x == nil {
		return nil
	}
	n := &Ident{Name: x.Name, Sp: s.sp(x.Sp)}
	s.rec(x, n)
	return n
}

func (s shifter) name(x *Name) *Name {
	if x == nil {
		return nil
	}
	n := &Name{Text: x.Text, Sp: s.sp(x.Sp)}
	s.rec(x, n)
	return n
}

func (s shifter) block(b *Block) *Block {
	if b == nil {
		return nil
	}
	n := &Block{Closed: b.Closed, Sp: s.sp(b.Sp), Items: make([]Stmt, len(b.Items))}
	for i, it := range b.Items {
		n.Items[i] = s.stmt(it)
	}
	s.rec(b, n)
	return n
}

func (s shifter) args(as []*Arg) []*Arg {
	if as == nil {
		return nil
	}
	out := make([]*Arg, len(as))
	for i, a := range as {
		n := &Arg{Name: s.name(a.Name), Value: s.expr(a.Value), Sp: s.sp(a.Sp)}
		s.rec(a, n)
		out[i] = n
	}
	return out
}

func (s shifter) stmt(st Stmt) Stmt {
	var n Stmt
	switch x := st.(type) {
	case nil:
		return nil
	case *Block:
		return s.block(x)
	case *LetStmt:
		n = &LetStmt{Name: s.name(x.Name), Type: s.name(x.Type), Value: s.expr(x.Value), Doc: x.Doc, Sp: s.sp(x.Sp)}
	case *IfStmt:
		n = &IfStmt{Cond: s.expr(x.Cond), Then: s.block(x.Then), Else: s.stmt(x.Else), Sp: s.sp(x.Sp)}
	case *WhileStmt:
		n = &WhileStmt{Cond: s.expr(x.Cond), Body: s.block(x.Body), Sp: s.sp(x.Sp)}
	case *ForStmt:
		n = &ForStmt{Var: s.name(x.Var), Iter: s.expr(x.Iter), Body: s.block(x.Body), Sp: s.sp(x.Sp)}
	case *ReturnStmt:
		n = &ReturnStmt{Value: s.expr(x.Value), Sp: s.sp(x.Sp)}
	case *BreakStmt:
		n = &BreakStmt{Sp: s.sp(x.Sp)}
	case *ContinueStmt:
		n = &ContinueStmt{Sp: s.sp(x.Sp)}
	case *CommandStmt:
		n = &CommandStmt{Text: x.Text, Sp: s.sp(x.Sp)}
	case *FieldStmt:
		n = &FieldStmt{Key: s.name(x.Key), KeyQuoted: x.KeyQuoted, Value: s.expr(x.Value), Sp: s.sp(x.Sp)}
	case *AssignStmt:
		n = &AssignStmt{Target: s.expr(x.Target), Op: x.Op, Value: s.expr(x.Value), Sp: s.sp(x.Sp)}
	case *ExprStmt:
		n = &ExprStmt{X: s.expr(x.X), Sp: s.sp(x.Sp)}
	default:
		panic("syntax: shift: unknown statement")
	}
	s.rec(st, n)
	return n
}

func (s shifter) expr(e Expr) Expr {
	var n Expr
	switch x := e.(type) {
	case nil:
		return nil
	case *Ident:
		return s.ident(x)
	case *Number:
		n = &Number{Text: x.Text, Sp: s.sp(x.Sp)}
	case *String:
		n = &String{Value: x.Value, Raw: x.Raw, Sp: s.sp(x.Sp)}
	case *Bool:
		n = &Bool{Value: x.Value, Sp: s.sp(x.Sp)}
	case *ResLoc:
		n = &ResLoc{Text: x.Text, Sp: s.sp(x.Sp)}
	case *Tag:
		n = &Tag{Name: x.Name, Sp: s.sp(x.Sp)}
	case *Selector:
		n = &Selector{Name: x.Name, Sp: s.sp(x.Sp)}
	case *Coord:
		n = &Coord{Text: x.Text, Sp: s.sp(x.Sp)}
	case *Paren:
		n = &Paren{X: s.expr(x.X), Sp: s.sp(x.Sp)}
	case *List:
		l := &List{Sp: s.sp(x.Sp)}
		if x.Elems != nil {
			l.Elems = make([]Expr, len(x.Elems))
			for i, el := range x.Elems {
				l.Elems[i] = s.expr(el)
			}
		}
		n = l
	case *Object:
		n = &Object{Body: s.block(x.Body), Sp: s.sp(x.Sp)}
	case *Unary:
		n = &Unary{Op: x.Op, X: s.expr(x.X), Sp: s.sp(x.Sp)}
	case *Binary:
		n = &Binary{Op: x.Op, L: s.expr(x.L), R: s.expr(x.R), Sp: s.sp(x.Sp)}
	case *Range:
		n = &Range{Lo: s.expr(x.Lo), Hi: s.expr(x.Hi), Sp: s.sp(x.Sp)}
	case *Member:
		n = &Member{X: s.expr(x.X), Name: s.name(x.Name), Sp: s.sp(x.Sp)}
	case *Call:
		n = &Call{Fn: s.expr(x.Fn), Args: s.args(x.Args), ArgsSp: s.sp(x.ArgsSp), Sp: s.sp(x.Sp)}
	case *Index:
		n = &Index{X: s.expr(x.X), Args: s.args(x.Args), Sp: s.sp(x.Sp)}
	case *Trailing:
		n = &Trailing{X: s.expr(x.X), Block: s.block(x.Block), Sp: s.sp(x.Sp)}
	case *Bad:
		n = &Bad{Sp: s.sp(x.Sp)}
	default:
		panic("syntax: shift: unknown expression")
	}
	s.rec(e, n)
	return n
}

// ShiftLeaf returns a copy of a leaf name node (an identifier, resource
// location, tag or string) moved by delta, or nil for any other node. It
// moves nodes that are not part of a declaration's tree, such as those
// parsed from the inside of a command, along with the declaration.
func ShiftLeaf(n Node, delta int) Node {
	s := shifter{d: delta}
	switch n := n.(type) {
	case *Ident:
		return &Ident{Name: n.Name, Sp: s.sp(n.Sp)}
	case *ResLoc:
		return &ResLoc{Text: n.Text, Sp: s.sp(n.Sp)}
	case *Tag:
		return &Tag{Name: n.Name, Sp: s.sp(n.Sp)}
	case *String:
		return &String{Value: n.Value, Raw: n.Raw, Sp: s.sp(n.Sp)}
	}
	return nil
}
