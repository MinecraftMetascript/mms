// Package syntax parses MMS source into an AST with exact spans and keeps
// every token, including whitespace and comments, for the formatter and
// editor features. Parsing never fails: broken input yields Bad nodes and
// diagnostics, and the rest of the tree is still usable.
package syntax

import "strings"

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
	// Items holds every top-level item in source order: *Decl, *ConstStmt,
	// *LetStmt and *ForStmt.
	Items []Stmt
	// Decls holds the declarations among Items, in order. Declarations
	// nested in a top-level for are not included.
	Decls  []*Decl
	Tokens []Token // every token, hidden channel included, in order
	Diags  []Diagnostic
	Src    []rune
}

// Decl is a declaration: `Kind name: Type (params) = value` or
// `Kind name { ... }`. Every part but Kind is optional. Declarations sit at
// the top level, in a top-level for, and in a Blueprint's emit block.
type Decl struct {
	Kind *Ident
	Name *Name
	// More holds the further names of a comma list, as in
	// `JigsawHole room_exit, hall_entry`.
	More   []*Name
	Type   TypeExpr   // type annotation, e.g. `int` in `Global x: int = 0`
	Params *ParamList // nil when the declaration has no parameter list
	Value  Expr       // after '='
	// TypeValue is the right-hand side of a type alias, as in
	// `Type LangText = string | {string: string}`.
	TypeValue TypeExpr
	Body      *Block // a `{ ... }` body
	// Extern marks `extern Kind name ...`: a declaration whose meaning is
	// supplied by the engine.
	Extern bool
	// Primary marks `primary Kind name ...` in an emit block.
	Primary bool
	Doc     string // text of the /** */ comment right above, if any
	Sp      Span
}

func (d *Decl) Span() Span { return d.Sp }

// Names returns the declaration's names: Name, then More.
func (d *Decl) Names() []*Name {
	if d.Name == nil {
		return nil
	}
	return append([]*Name{d.Name}, d.More...)
}

// NameText returns the declared name or "" when it is missing.
func (d *Decl) NameText() string {
	if d.Name == nil {
		return ""
	}
	return d.Name.Text
}

// Name is an identifier or a resource location used as a name or key.
type Name struct {
	Text string
	// Parts is set when the name is a template such as `${w}_crate`.
	Parts []NamePart
	Sp    Span
}

func (n *Name) Span() Span { return n.Sp }

// Templated reports whether the name has `${...}` pieces.
func (n *Name) Templated() bool { return n != nil && len(n.Parts) > 0 }

// NamePart is one piece of a templated name: literal text, or the
// expression inside a `${...}`.
type NamePart struct {
	Lit string // set when X is nil
	X   Expr
	Sp  Span // the literal text, or the whole `${...}`
}

type ParamList struct {
	Params []*Param
	Sp     Span
}

func (p *ParamList) Span() Span { return p.Sp }

// Param is `name: Type = default`; Type and Default are optional.
type Param struct {
	Name    *Name
	Type    TypeExpr
	Default Expr
	Sp      Span
}

func (p *Param) Span() Span { return p.Sp }

// Block is `{ items }`. Closed is false when the closing brace is missing.
type Block struct {
	Items  []Stmt
	Closed bool
	// Bare marks the branch of an if statement written without braces, as
	// in `if (x) error(...)`. It holds one ExprStmt.
	Bare bool
	Sp   Span
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
		Type  TypeExpr
		Value Expr
		Doc   string // text of the /** */ comment right above, if any
		Sp    Span
	}
	// ConstStmt is `const name: Type = value`, at the top level or in a
	// body.
	ConstStmt struct {
		Name  *Name
		Type  TypeExpr
		Value Expr
		Doc   string
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
	// ForStmt is `for x in iter { }` or `for k, v in iter { }`, in a body
	// or at the top level, where its body holds declarations.
	ForStmt struct {
		Var  *Name
		Var2 *Name // the second variable of `for k, v in ...`, or nil
		Iter Expr
		Body *Block
		Sp   Span
	}
	ReturnStmt struct {
		Value Expr
		Sp    Span
	}
	BreakStmt    struct{ Sp Span }
	ContinueStmt struct{ Sp Span }
	// CommandStmt is a raw command line, without the leading '/'. It is
	// also an expression as the value of an assignment,
	// `let n = /data get block ~ ~ ~ Items`, where it stands for the
	// command's result.
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
	// FieldDecl is `name: Type = default` in a Type or Blueprint body.
	FieldDecl struct {
		Name    *Name
		Type    TypeExpr
		Default Expr
		Doc     string
		Sp      Span
	}
	// SectionStmt is a named block in a Blueprint or Test body:
	// `check { }`, `emit { }` or `setup { }`.
	SectionStmt struct {
		Word *Ident
		Body *Block
		Sp   Span
	}
	// AsStmt is `as Type = value hint("rule", "message")` in a Blueprint.
	AsStmt struct {
		Type  TypeExpr
		Value Expr
		Hint  *Call // nil without a hint
		Doc   string
		Sp    Span
	}
	// ExpectStmt is `expect cond` in a Test.
	ExpectStmt struct {
		X  Expr
		Sp Span
	}
	// RunStmt is `run call()` in a Test.
	RunStmt struct {
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
func (s *ConstStmt) Span() Span    { return s.Sp }
func (s *FieldDecl) Span() Span    { return s.Sp }
func (s *SectionStmt) Span() Span  { return s.Sp }
func (s *AsStmt) Span() Span       { return s.Sp }
func (s *ExpectStmt) Span() Span   { return s.Sp }
func (s *RunStmt) Span() Span      { return s.Sp }
func (*Decl) stmt()                {}
func (*ConstStmt) stmt()           {}
func (*FieldDecl) stmt()           {}
func (*SectionStmt) stmt()         {}
func (*AsStmt) stmt()              {}
func (*ExpectStmt) stmt()          {}
func (*RunStmt) stmt()             {}
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
	// A coordinate a variable offsets, `~x` or `^-spot.y`, is written
	// as is in Text; Names holds each one's name (an Ident, or a Member
	// for a field), in order.
	Coord struct {
		Text  string
		Names []Expr
		Sp    Span
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
	// TemplatedName is a name with `${...}` pieces used as a value, as in
	// `item/${self.path}`.
	TemplatedName struct {
		Text  string // as written
		Parts []NamePart
		Sp    Span
	}
	// Spread is `...x` in a list, an object or arguments.
	Spread struct {
		X  Expr
		Sp Span
	}
	// IfExpr is `if (cond) then else other`; Else may be nil, which
	// means none.
	IfExpr struct {
		Cond, Then, Else Expr
		Sp               Span
	}
	// RunExpr is an inline function body, `run { ... }`.
	RunExpr struct {
		Body *Block
		Sp   Span
	}
	// Comprehension is `[for x in iter if cond: body]`.
	Comprehension struct {
		Var  *Name
		Var2 *Name // `for k, v in ...`, or nil
		Iter Expr
		Cond Expr // nil without a filter
		Body Expr
		Sp   Span
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

func (e *Number) Span() Span        { return e.Sp }
func (e *String) Span() Span        { return e.Sp }
func (e *Bool) Span() Span          { return e.Sp }
func (e *Ident) Span() Span         { return e.Sp }
func (e *ResLoc) Span() Span        { return e.Sp }
func (e *Tag) Span() Span           { return e.Sp }
func (e *Selector) Span() Span      { return e.Sp }
func (e *Coord) Span() Span         { return e.Sp }
func (e *Paren) Span() Span         { return e.Sp }
func (e *List) Span() Span          { return e.Sp }
func (e *Object) Span() Span        { return e.Sp }
func (e *Unary) Span() Span         { return e.Sp }
func (e *Binary) Span() Span        { return e.Sp }
func (e *Range) Span() Span         { return e.Sp }
func (e *Member) Span() Span        { return e.Sp }
func (e *Call) Span() Span          { return e.Sp }
func (e *Index) Span() Span         { return e.Sp }
func (e *Trailing) Span() Span      { return e.Sp }
func (e *Bad) Span() Span           { return e.Sp }
func (e *TemplatedName) Span() Span { return e.Sp }
func (e *Spread) Span() Span        { return e.Sp }
func (e *IfExpr) Span() Span        { return e.Sp }
func (e *RunExpr) Span() Span       { return e.Sp }
func (e *Comprehension) Span() Span { return e.Sp }
func (*TemplatedName) expr()        {}
func (*Spread) expr()               {}
func (*IfExpr) expr()               {}
func (*RunExpr) expr()              {}
func (*Comprehension) expr()        {}
func (*Number) expr()               {}
func (*String) expr()               {}
func (*Bool) expr()                 {}
func (*Ident) expr()                {}
func (*ResLoc) expr()               {}
func (*Tag) expr()                  {}
func (*Selector) expr()             {}
func (*Coord) expr()                {}
func (*Paren) expr()                {}
func (*List) expr()                 {}
func (*Object) expr()               {}
func (*Unary) expr()                {}
func (*Binary) expr()               {}
func (*Range) expr()                {}
func (*Member) expr()               {}
func (*Call) expr()                 {}
func (*Index) expr()                {}
func (*Trailing) expr()             {}
func (*Bad) expr()                  {}
func (*CommandStmt) expr()          {}

// TypeExpr is a type written in source: `int`, `Id<item>`, `[T]`,
// `{string: T}`, `T?`, `A | B` or a literal word such as `"rigid"`.
type TypeExpr interface {
	Node
	typeExpr()
}

type (
	// TypeName is a named type with optional arguments: `int`,
	// `Id<item>`, or a dotted or namespaced word such as an objective
	// criterion.
	TypeName struct {
		Name string
		Args []TypeExpr
		Sp   Span
	}
	TypeOptional struct {
		X  TypeExpr
		Sp Span
	}
	TypeList struct {
		Elem TypeExpr
		Sp   Span
	}
	TypeMap struct {
		Key, Value TypeExpr
		Sp         Span
	}
	TypeUnion struct {
		Arms []TypeExpr
		Sp   Span
	}
	// TypeRecord is a record type written in place:
	// `{ x: int, y: int = 0 }`, with fields as in a Type body.
	TypeRecord struct {
		Fields []*FieldDecl
		Sp     Span
	}
	// TypeString is a literal word in a type, as in `"rigid" | "beard"`.
	TypeString struct {
		Value string
		Raw   string
		Sp    Span
	}
	// TypeNumber is a whole number in a type, as in `0 | 1 | 2`. Text
	// is the number as written, with its sign.
	TypeNumber struct {
		Text string
		Sp   Span
	}
	TypeParen struct {
		X  TypeExpr
		Sp Span
	}
)

func (t *TypeName) Span() Span     { return t.Sp }
func (t *TypeOptional) Span() Span { return t.Sp }
func (t *TypeList) Span() Span     { return t.Sp }
func (t *TypeMap) Span() Span      { return t.Sp }
func (t *TypeUnion) Span() Span    { return t.Sp }
func (t *TypeRecord) Span() Span   { return t.Sp }
func (t *TypeString) Span() Span   { return t.Sp }
func (t *TypeParen) Span() Span    { return t.Sp }
func (t *TypeNumber) Span() Span   { return t.Sp }
func (*TypeName) typeExpr()        {}
func (*TypeOptional) typeExpr()    {}
func (*TypeList) typeExpr()        {}
func (*TypeMap) typeExpr()         {}
func (*TypeUnion) typeExpr()       {}
func (*TypeRecord) typeExpr()      {}
func (*TypeString) typeExpr()      {}
func (*TypeParen) typeExpr()       {}
func (*TypeNumber) typeExpr()      {}

// TypeText renders a type in canonical form, or "" for nil.
func TypeText(t TypeExpr) string {
	switch t := t.(type) {
	case *TypeName:
		if len(t.Args) == 0 {
			return t.Name
		}
		args := make([]string, len(t.Args))
		for i, a := range t.Args {
			args[i] = TypeText(a)
		}
		return t.Name + "<" + strings.Join(args, ", ") + ">"
	case *TypeOptional:
		return TypeText(t.X) + "?"
	case *TypeList:
		return "[" + TypeText(t.Elem) + "]"
	case *TypeMap:
		return "{" + TypeText(t.Key) + ": " + TypeText(t.Value) + "}"
	case *TypeRecord:
		fs := make([]string, len(t.Fields))
		for i, f := range t.Fields {
			fs[i] = f.Name.Text
			if f.Type != nil {
				fs[i] += ": " + TypeText(f.Type)
			}
			if f.Default != nil {
				fs[i] += " = …"
			}
		}
		return "{ " + strings.Join(fs, ", ") + " }"
	case *TypeUnion:
		arms := make([]string, len(t.Arms))
		for i, a := range t.Arms {
			arms[i] = TypeText(a)
		}
		return strings.Join(arms, " | ")
	case *TypeString:
		return t.Raw
	case *TypeNumber:
		return t.Text
	case *TypeParen:
		return "(" + TypeText(t.X) + ")"
	}
	return ""
}

// CoordName reports whether w, one coordinate of a Coord's Text, is
// offset by a name, as `~x` or `^-y` are: it returns the `~` or `^` and
// whether the name is negated.
func CoordName(w string) (prefix string, neg, ok bool) {
	if len(w) < 2 || (w[0] != '~' && w[0] != '^') {
		return "", false, false
	}
	rest := w[1:]
	if rest[0] == '-' {
		neg, rest = true, rest[1:]
	}
	if rest == "" || !(rest[0] == '_' || rest[0] >= 'a' && rest[0] <= 'z' || rest[0] >= 'A' && rest[0] <= 'Z') {
		return "", false, false
	}
	return w[:1], neg, true
}
