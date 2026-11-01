// Package syntax parses MMS source into an AST with exact spans and keeps
// every token, including whitespace and comments, for the formatter and
// editor features. Parsing never fails: broken input yields Bad nodes and
// diagnostics, and the rest of the tree is still usable.
package syntax

// Span is a half-open range of rune offsets into the source.
type Span struct {
	Start, End int
}

func (s Span) Contains(off int) bool { return s.Start <= off && off <= s.End }
func (s Span) Len() int              { return s.End - s.Start }

// Join returns the smallest span covering both.
func (s Span) Join(o Span) Span {
	if o.Start < s.Start {
		s.Start = o.Start
	}
	if o.End > s.End {
		s.End = o.End
	}
	return s
}

type Node interface{ Span() Span }

// File is one parsed source file.
type File struct {
	// Decls are the declarations at the top level, outside any
	// build-time `for` or `if`.
	Decls []*Decl
	// Items are every top-level item in order: the Decls, and build-time
	// statements (TopFor, TopIf, TopLet) with the declarations
	// inside them.
	Items  []Item
	Tokens []Token // every token, hidden channel included, in order
	Diags  []Diagnostic
	Src    []rune
}

// Decl is a top-level declaration: `Kind name: Type (params) = value` or
// `Kind name { ... }`. Every part but Kind is optional.
type Decl struct {
	Kind   *Ident
	Name   *Name
	Type   *Name      // type annotation, e.g. `int` in `Global x: int = 0`
	Params *ParamList // nil when the declaration has no parameter list
	Value  Expr       // after '='
	Body   *Block     // a `{ ... }` body
	Doc    string     // text of the /** */ comment right above, if any
	Sp     Span
}

func (d *Decl) Span() Span { return d.Sp }

// NameText returns the declared name or "" when it is missing.
func (d *Decl) NameText() string {
	if d.Name == nil {
		return ""
	}
	return d.Name.Text
}

// Name is an identifier or a resource location used as a name or type.
// A declaration name may also be a string, a template such as
// "planet_${i}" inside a build-time loop; Quoted is then set and Text is
// the unquoted value.
type Name struct {
	Text   string
	Quoted bool
	Sp     Span
}

// Item is a top-level item: a *Decl or a build-time statement.
type Item interface {
	Node
	item()
}

type (
	// TopFor is a build-time loop at the top level: its body is copied
	// once per value, with Var bound to the value and Index, if any, to
	// its position from 0.
	TopFor struct {
		Index *Name // `i` in `for i, x in xs`; nil without one
		Var   *Name
		Iter  Expr
		Body  *TopBlock
		Sp    Span
	}
	// TopIf keeps its Then or Else items depending on a build-time
	// condition.
	TopIf struct {
		Cond Expr
		Then *TopBlock
		Else Item // *TopIf, *TopBlock or nil
		Sp   Span
	}
	// TopLet binds a build-time value for the rest of its block.
	TopLet struct {
		Name  *Name
		Value Expr
		Doc   string
		Sp    Span
	}
	// TopBlock is the `{ items }` body of a TopFor or TopIf.
	TopBlock struct {
		Items  []Item
		Closed bool
		Sp     Span
	}
)

func (d *Decl) item()          {}
func (*TopFor) item()          {}
func (*TopIf) item()           {}
func (*TopLet) item()          {}
func (*TopBlock) item()        {}
func (s *TopFor) Span() Span   { return s.Sp }
func (s *TopIf) Span() Span    { return s.Sp }
func (s *TopLet) Span() Span   { return s.Sp }
func (s *TopBlock) Span() Span { return s.Sp }

func (n *Name) Span() Span { return n.Sp }

type ParamList struct {
	Params []*Param
	Sp     Span
}

func (p *ParamList) Span() Span { return p.Sp }

type Param struct {
	Name *Name
	Type *Name
	Sp   Span
}

func (p *Param) Span() Span { return p.Sp }

// Block is `{ items }`. Closed is false when the closing brace is missing.
type Block struct {
	Items  []Stmt
	Closed bool
	Sp     Span
}

func (b *Block) Span() Span { return b.Sp }

// Stmt is an item in a block: a field, a statement or a bare expression.
type Stmt interface {
	Node
	stmt()
}

type (
	LetStmt struct {
		Name  *Name
		Type  *Name
		Value Expr
		Doc   string // text of the /** */ comment right above, if any
		Sp    Span
	}
	IfStmt struct {
		Cond Expr
		Then *Block
		Else Stmt // *IfStmt, *Block or nil
		Sp   Span
	}
	WhileStmt struct {
		Cond Expr
		Body *Block
		Sp   Span
	}
	ForStmt struct {
		Index *Name // `i` in `for i, x in xs`; nil without one
		Var   *Name
		Iter  Expr
		Body  *Block
		Sp    Span
	}
	ReturnStmt struct {
		Value Expr
		Sp    Span
	}
	BreakStmt    struct{ Sp Span }
	ContinueStmt struct{ Sp Span }
	// CommandStmt is a raw command line, without the leading '/'.
	CommandStmt struct {
		Text string
		Sp   Span
	}
	// FieldStmt is `key = value`: a field in a resource body, or an
	// assignment in a function body.
	FieldStmt struct {
		Key       *Name
		KeyQuoted bool // the key was written as a string
		Value     Expr
		Sp        Span
	}
	// AssignStmt is a compound assignment such as `x += 1`.
	AssignStmt struct {
		Target Expr
		Op     string
		Value  Expr
		Sp     Span
	}
	ExprStmt struct {
		X  Expr
		Sp Span
	}
)

func (s *LetStmt) Span() Span      { return s.Sp }
func (s *IfStmt) Span() Span       { return s.Sp }
func (s *WhileStmt) Span() Span    { return s.Sp }
func (s *ForStmt) Span() Span      { return s.Sp }
func (s *ReturnStmt) Span() Span   { return s.Sp }
func (s *BreakStmt) Span() Span    { return s.Sp }
func (s *ContinueStmt) Span() Span { return s.Sp }
func (s *CommandStmt) Span() Span  { return s.Sp }
func (s *FieldStmt) Span() Span    { return s.Sp }
func (s *AssignStmt) Span() Span   { return s.Sp }
func (s *ExprStmt) Span() Span     { return s.Sp }
func (b *Block) stmt()             {}
func (*LetStmt) stmt()             {}
func (*IfStmt) stmt()              {}
func (*WhileStmt) stmt()           {}
func (*ForStmt) stmt()             {}
func (*ReturnStmt) stmt()          {}
func (*BreakStmt) stmt()           {}
func (*ContinueStmt) stmt()        {}
func (*CommandStmt) stmt()         {}
func (*FieldStmt) stmt()           {}
func (*AssignStmt) stmt()          {}
func (*ExprStmt) stmt()            {}

// Expr is an expression node.
type Expr interface {
	Node
	expr()
}

type (
	Number struct {
		Text string
		Sp   Span
	}
	String struct {
		Value string // unquoted, escapes resolved
		Raw   string // as written, quotes included
		Sp    Span
	}
	Bool struct {
		Value bool
		Sp    Span
	}
	Ident struct {
		Name string
		Sp   Span
	}
	ResLoc struct {
		Text string
		Sp   Span
	}
	// Tag is `#ns:path`; Name excludes the '#'.
	Tag struct {
		Name string
		Sp   Span
	}
	// Selector is `@s`; Name excludes the '@'. Arguments are an Index.
	Selector struct {
		Name string
		Sp   Span
	}
	// Coord is a relative or local coordinate group such as `~ ~1 ~`.
	Coord struct {
		Text string
		Sp   Span
	}
	Paren struct {
		X  Expr
		Sp Span
	}
	List struct {
		Elems []Expr
		Sp    Span
	}
	// Object is a `{ ... }` literal value.
	Object struct {
		Body *Block
		Sp   Span
	}
	Unary struct {
		Op string
		X  Expr
		Sp Span
	}
	Binary struct {
		Op   string
		L, R Expr
		Sp   Span
	}
	// Range is `lo..hi`; either side may be nil.
	Range struct {
		Lo, Hi Expr
		Sp     Span
	}
	// Member is `x.name`; Name is nil while it is still being typed.
	Member struct {
		X    Expr
		Name *Name
		Sp   Span
	}
	Call struct {
		Fn   Expr
		Args []*Arg
		// ArgsSp covers the parentheses.
		ArgsSp Span
		Sp     Span
	}
	// Index is `x[args]`, used for block states and selector arguments.
	Index struct {
		X    Expr
		Args []*Arg
		Sp   Span
	}
	// Trailing is an expression followed by a block, as in
	// `If(cond) { ... }` or `As(@a) { ... }`.
	Trailing struct {
		X     Expr
		Block *Block
		Sp    Span
	}
	// Bad stands in for an expression that failed to parse.
	Bad struct{ Sp Span }
)

// Arg is a call or index argument; Name is set for `name = value`.
type Arg struct {
	Name  *Name
	Value Expr
	Sp    Span
}

func (a *Arg) Span() Span { return a.Sp }

func (e *Number) Span() Span   { return e.Sp }
func (e *String) Span() Span   { return e.Sp }
func (e *Bool) Span() Span     { return e.Sp }
func (e *Ident) Span() Span    { return e.Sp }
func (e *ResLoc) Span() Span   { return e.Sp }
func (e *Tag) Span() Span      { return e.Sp }
func (e *Selector) Span() Span { return e.Sp }
func (e *Coord) Span() Span    { return e.Sp }
func (e *Paren) Span() Span    { return e.Sp }
func (e *List) Span() Span     { return e.Sp }
func (e *Object) Span() Span   { return e.Sp }
func (e *Unary) Span() Span    { return e.Sp }
func (e *Binary) Span() Span   { return e.Sp }
func (e *Range) Span() Span    { return e.Sp }
func (e *Member) Span() Span   { return e.Sp }
func (e *Call) Span() Span     { return e.Sp }
func (e *Index) Span() Span    { return e.Sp }
func (e *Trailing) Span() Span { return e.Sp }
func (e *Bad) Span() Span      { return e.Sp }
func (*Number) expr()          {}
func (*String) expr()          {}
func (*Bool) expr()            {}
func (*Ident) expr()           {}
func (*ResLoc) expr()          {}
func (*Tag) expr()             {}
func (*Selector) expr()        {}
func (*Coord) expr()           {}
func (*Paren) expr()           {}
func (*List) expr()            {}
func (*Object) expr()          {}
func (*Unary) expr()           {}
func (*Binary) expr()          {}
func (*Range) expr()           {}
func (*Member) expr()          {}
func (*Call) expr()            {}
func (*Index) expr()           {}
func (*Trailing) expr()        {}
func (*Bad) expr()             {}
