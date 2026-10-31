// Package fn is the function backend: a small structured IR for function
// bodies and its lowering to .mcfunction files.
//
// Values live on one scoreboard objective per namespace, "mms.<ns>":
//
//	$name        a Global
//	#<fn>.<x>    a parameter, local or temporary of function <fn>
//	#c<n>        the constant n, set by the load function
//	#cf, #rv     control flow code and return value, used to leave
//	             generated sub-functions early
//
// Structured control flow and execute blocks become generated
// sub-functions next to the function: <fn>/_if0, <fn>/_while1,
// <fn>/_exec2 and so on.
package fn

import (
	"strconv"
)

// Score is a scoreboard slot: a holder on an objective.
type Score struct {
	Holder, Obj string
}

func (s Score) String() string { return s.Holder + " " + s.Obj }

// Operand is a constant or a score.
type Operand struct {
	Const   int32
	IsConst bool
	S       Score
}

func Const(v int32) Operand   { return Operand{Const: v, IsConst: true} }
func Of(s Score) Operand      { return Operand{S: s} }
func (o Operand) str() string { return strconv.Itoa(int(o.Const)) }

// Block is a list of operations run in order.
type Block struct {
	Ops []Op
}

func (b *Block) Add(ops ...Op) { b.Ops = append(b.Ops, ops...) }

// Op is one operation.
type Op interface{ op() }

type (
	// Cmd is a command line, written as it is.
	Cmd struct{ Text string }
	// Set is `Dst = Src`.
	Set struct {
		Dst Score
		Src Operand
	}
	// Arith is `Dst op= Src` for + - * / % and min (<) / max (>).
	Arith struct {
		Dst Score
		Op  string // "+=", "-=", "*=", "/=", "%=", "<", ">"
		Src Operand
	}
	// Call runs a function, storing its result in Dst when set. Args
	// pass the arguments (into the callee's parameter scores and storage)
	// just before the call. A recursive call saves the scores in Frame
	// that are still needed afterwards, and the storage at Strings, on a
	// stack around the call.
	Call struct {
		Fn        string
		Dst       *Score
		Args      []Op
		Recursive bool
		Frame     []Score
		Strings   string
	}
	// If runs Then when Cond holds, otherwise Else.
	If struct {
		Cond       Cond
		Then, Else *Block
	}
	// Loop runs Body while Cond holds, running Post after each pass
	// (and on continue).
	Loop struct {
		Cond       Cond
		Body, Post *Block
	}
	// Macro runs one command line holding $(key) placeholders through a
	// macro sub-function, with the arguments copied into storage first.
	Macro struct {
		Line string
		Args []MacroArg
	}
	// Exec runs Body under an execute chain ("as @a", "at @s"), which
	// may run it once per entity or not at all. Body never leaves early:
	// the checker rejects break, continue and return inside it.
	Exec struct {
		Parts []string
		Body  *Block
	}
	// Return leaves the function, with a value when Val is set.
	Return   struct{ Val *Operand }
	Break    struct{}
	Continue struct{}
)

func (Cmd) op()      {}
func (Macro) op()    {}
func (Set) op()      {}
func (Arith) op()    {}
func (Call) op()     {}
func (If) op()       {}
func (Loop) op()     {}
func (Exec) op()     {}
func (Return) op()   {}
func (Break) op()    {}
func (Continue) op() {}

// MacroArg is a macro argument: a score, a storage location ("<storage>
// <path>"), or a fixed text.
type MacroArg struct {
	Key   string
	Score *Score
	Path  string
	Text  *string
}

// Storage is the command storage for a namespace's non-score values.
func Storage(ns string) string { return ns + ":mms" }

// LocalPath is the storage path of a function's string local.
func LocalPath(fn, name string) string { return "l.\"" + fn + "\"." + name }

// GlobalPath is the storage path of a string Global.
func GlobalPath(name string) string { return "g." + name }

// Cond is a condition: Pre runs first, then every part of the execute
// chain must hold ("if score #a mms.x < #b mms.x", "unless entity @s").
// A Cond with no parts always holds; Never marks one that never does.
type Cond struct {
	Pre   []Op
	Parts []string
	Never bool
}

// Always reports whether the condition is known to hold.
func (c Cond) Always() bool { return !c.Never && len(c.Parts) == 0 }

// Func is a function to generate.
type Func struct {
	ID   string // "ns:path"
	Body *Block
}
