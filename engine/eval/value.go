package eval

import (
	"math"
	"strconv"
	"strings"

	"github.com/minecraftmetascript/mms/engine/syntax"
)

// Origin is a place in the source: a file and a span in it.
type Origin struct {
	Path string
	Sp   syntax.Span
}

// FrameKind says what an expansion step was.
type FrameKind int

const (
	FrameFor       FrameKind = iota // one iteration of a top-level or emit `for`
	FrameTemplate                   // a Template call
	FrameBlueprint                  // a declaration of a Blueprint's kind
	FrameConst                      // a const copied in where it is named
)

func (k FrameKind) String() string {
	switch k {
	case FrameFor:
		return "for"
	case FrameTemplate:
		return "Template"
	case FrameBlueprint:
		return "Blueprint"
	case FrameConst:
		return "const"
	}
	return "?"
}

// Frame is one step of expansion: a loop iteration, a Template call, a
// Blueprint instance or a const use.
type Frame struct {
	Kind FrameKind
	// Name is what the editor shows: `w = oak`, `Piece`, `Wand`.
	Name string
	// Call is where the step was entered, which is user code for the
	// outermost frame.
	Call Origin
	// Def is the loop, Template, Blueprint or const that was expanded.
	Def Origin
}

// Prov is where a value or declaration came from. At is the first user
// span that caused it, the one diagnostics are reported at. Def, when
// set, is where its text was written, such as a line in a Template body.
// Frames lists the expansion steps it came through, outermost first.
type Prov struct {
	At     Origin
	Def    *Origin
	Frames []Frame
}

// Value is a compile-time value. Each kind carries the span it is lowered
// back to (pos.Sp) and, for text written in a Template, Blueprint or const,
// where that text is (pos.def).
type Value interface {
	at() *pos
}

type pos struct {
	Sp  syntax.Span
	def *Origin
	fr  []Frame
}

func (p *pos) at() *pos { return p }

type (
	// None is the absent value: an unset optional field, or an if
	// expression whose condition was false and that has no else. A
	// field whose value is none is left out. Written out as `none`
	// (lit), it compares as none but stays the word `none` in output,
	// which is also a value of some game enums (`skybox = none`).
	None struct {
		pos
		lit bool
	}
	Bool struct {
		pos
		V bool
	}
	// Num is a number. Int numbers keep integer arithmetic.
	Num struct {
		pos
		F   float64
		Int bool
	}
	Str struct {
		pos
		S string
	}
	// Word is an identifier, resource location or templated name used as
	// a value: `oak`, `minecraft:empty`, `item/${self.path}`.
	Word struct {
		pos
		Text string
	}
	List struct {
		pos
		Elems []Value
	}
	// Record is an object literal's fields, in order.
	Record struct {
		pos
		Keys []string
		Vals []Value
		// id, when set, is what the record means as a value: a
		// Blueprint's self, whose fields are id, namespace and path, is
		// its id.
		id *Word
	}
	// RangeV is `lo..hi`, both ends included; either may be nil.
	RangeV struct {
		pos
		Lo, Hi Value
	}
	// Code is anything the evaluator does not compute, such as a call of
	// a built-in helper, already expanded: `Pool(1..3, #gems)`.
	Code struct {
		pos
		X syntax.Expr
	}
	// Stmts are the statements a Template returning Commands expands to.
	Stmts struct {
		pos
		Items []syntax.Stmt
	}
)

func (r *Record) get(k string) (Value, bool) {
	for i, key := range r.Keys {
		if key == k {
			return r.Vals[i], true
		}
	}
	return nil, false
}

func (r *Record) set(k string, v Value) {
	for i, key := range r.Keys {
		if key == k {
			r.Vals[i] = v
			return
		}
	}
	r.Keys = append(r.Keys, k)
	r.Vals = append(r.Vals, v)
}

// known reports whether v is a fully computed scalar: what a condition, a
// name piece or a `${...}` piece needs.
func known(v Value) bool {
	switch v.(type) {
	case *None, *Bool, *Num, *Str, *Word:
		return true
	}
	return false
}

// text spells a known scalar the way it reads in a name or command.
func text(v Value) (string, bool) {
	switch v := v.(type) {
	case *Str:
		return v.S, true
	case *Word:
		return v.Text, true
	case *Num:
		return numText(v), true
	case *Bool:
		return strconv.FormatBool(v.V), true
	}
	return "", false
}

func numText(n *Num) string {
	if n.Int {
		return strconv.FormatInt(int64(n.F), 10)
	}
	// Six places is finer than any game coordinate, and hides float
	// noise such as cos(90) = 6e-17.
	f := math.Round(n.F*1e6) / 1e6
	if f == 0 {
		f = 0 // not -0
	}
	return strconv.FormatFloat(f, 'f', -1, 64)
}

// equal compares two known values. Words and strings compare by text, so
// `w != mangrove` and `w != "mangrove"` agree.
func equal(a, b Value) bool {
	switch a := a.(type) {
	case *None:
		_, ok := b.(*None)
		return ok
	case *Num:
		if b, ok := b.(*Num); ok {
			return a.F == b.F
		}
		return false
	case *Bool:
		if b, ok := b.(*Bool); ok {
			return a.V == b.V
		}
		return false
	}
	if _, ok := b.(*None); ok {
		return false
	}
	at, ok1 := text(a)
	bt, ok2 := text(b)
	return ok1 && ok2 && strings.TrimPrefix(at, "minecraft:") == strings.TrimPrefix(bt, "minecraft:")
}

// withPos returns v and everything in it moved to sp, recording where the
// text was written: a const copied in at a use.
func withPos(v Value, sp syntax.Span, fr []Frame, path string) Value {
	if v == nil {
		return nil
	}
	p := *v.at()
	def := p.def
	if def == nil {
		def = &Origin{Path: path, Sp: p.Sp}
	}
	np := pos{Sp: sp, def: def, fr: fr}
	switch v := v.(type) {
	case *None:
		return &None{np, v.lit}
	case *Bool:
		return &Bool{np, v.V}
	case *Num:
		return &Num{np, v.F, v.Int}
	case *Str:
		return &Str{np, v.S}
	case *Word:
		return &Word{np, v.Text}
	case *List:
		out := &List{pos: np}
		for _, e := range v.Elems {
			out.Elems = append(out.Elems, withPos(e, sp, fr, path))
		}
		return out
	case *Record:
		out := &Record{pos: np, Keys: append([]string(nil), v.Keys...)}
		if v.id != nil {
			out.id = &Word{np, v.id.Text}
		}
		for _, e := range v.Vals {
			out.Vals = append(out.Vals, withPos(e, sp, fr, path))
		}
		return out
	case *RangeV:
		return &RangeV{np, withPos(v.Lo, sp, fr, path), withPos(v.Hi, sp, fr, path)}
	case *Code:
		return &Code{np, syntax.CloneExprAt(v.X, sp)}
	case *Stmts:
		out := &Stmts{pos: np}
		for _, st := range v.Items {
			out.Items = append(out.Items, syntax.CloneStmtAt(st, sp))
		}
		return out
	}
	return v
}

// kindName names a value's kind in messages.
func kindName(v Value) string {
	switch v.(type) {
	case *None:
		return "none"
	case *Bool:
		return "a bool"
	case *Num:
		return "a number"
	case *Str:
		return "a string"
	case *Word:
		return "a name"
	case *List:
		return "a list"
	case *Record:
		return "an object"
	case *RangeV:
		return "a range"
	case *Stmts:
		return "commands"
	}
	return "a value only known when the pack runs"
}
