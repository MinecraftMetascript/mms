package syntax

// Children returns a node's direct children in source order.
func Children(n Node) []Node {
	var out []Node
	add := func(ns ...Node) {
		for _, c := range ns {
			if c != nil && !isNilNode(c) {
				out = append(out, c)
			}
		}
	}
	switch n := n.(type) {
	case *File:
	case *Decl:
		add(n.Kind)
		if n.Name != nil {
			add(n.Name)
		}
		if n.Type != nil {
			add(n.Type)
		}
		if n.Params != nil {
			add(n.Params)
		}
		if n.Value != nil {
			add(n.Value)
		}
		if n.Body != nil {
			add(n.Body)
		}
	case *ParamList:
		for _, p := range n.Params {
			add(p)
		}
	case *Param:
		add(n.Name)
		if n.Type != nil {
			add(n.Type)
		}
	case *Block:
		for _, it := range n.Items {
			add(it)
		}
	case *LetStmt:
		add(n.Name)
		if n.Type != nil {
			add(n.Type)
		}
		if n.Value != nil {
			add(n.Value)
		}
	case *IfStmt:
		add(n.Cond, n.Then)
		if n.Else != nil {
			add(n.Else)
		}
	case *WhileStmt:
		add(n.Cond, n.Body)
	case *ForStmt:
		add(n.Var, n.Iter, n.Body)
	case *ReturnStmt:
		if n.Value != nil {
			add(n.Value)
		}
	case *FieldStmt:
		add(n.Key, n.Value)
	case *AssignStmt:
		add(n.Target, n.Value)
	case *ExprStmt:
		add(n.X)
	case *Paren:
		add(n.X)
	case *List:
		for _, e := range n.Elems {
			add(e)
		}
	case *Object:
		add(n.Body)
	case *Unary:
		add(n.X)
	case *Binary:
		add(n.L, n.R)
	case *Range:
		if n.Lo != nil {
			add(n.Lo)
		}
		if n.Hi != nil {
			add(n.Hi)
		}
	case *Member:
		add(n.X)
		if n.Name != nil {
			add(n.Name)
		}
	case *Call:
		add(n.Fn)
		for _, a := range n.Args {
			add(a)
		}
	case *Index:
		add(n.X)
		for _, a := range n.Args {
			add(a)
		}
	case *Trailing:
		add(n.X, n.Block)
	case *Arg:
		if n.Name != nil {
			add(n.Name)
		}
		add(n.Value)
	}
	return out
}

func isNilNode(n Node) bool {
	switch n := n.(type) {
	case *Name:
		return n == nil
	case *Ident:
		return n == nil
	case *Block:
		return n == nil
	}
	return false
}

// Span of a File covers all its source.
func (f *File) Span() Span { return Span{0, len(f.Src)} }

// Walk calls fn for n and its descendants, depth first, in source order.
// Returning false skips a node's children.
func Walk(n Node, fn func(Node) bool) {
	if !fn(n) {
		return
	}
	if f, ok := n.(*File); ok {
		for _, d := range f.Decls {
			Walk(d, fn)
		}
		return
	}
	for _, c := range Children(n) {
		Walk(c, fn)
	}
}

// PathAt returns the nodes containing offset off, outermost first. A node
// contains the offsets from its start through its end, so a cursor just
// after an identifier is still on it.
func PathAt(f *File, off int) []Node {
	var path []Node
	var visit func(n Node)
	visit = func(n Node) {
		path = append(path, n)
		var kids []Node
		if file, ok := n.(*File); ok {
			for _, d := range file.Decls {
				kids = append(kids, d)
			}
		} else {
			kids = Children(n)
		}
		// Prefer the last child containing off, so `a|.b` picks the
		// innermost node that ends at the cursor.
		for i := len(kids) - 1; i >= 0; i-- {
			if kids[i].Span().Contains(off) {
				visit(kids[i])
				return
			}
		}
	}
	visit(f)
	return path
}
