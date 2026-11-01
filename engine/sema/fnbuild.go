package sema

import (
	"math"
	"strconv"

	"github.com/minecraftmetascript/mms/engine/fn"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// Build-time values in function bodies. A Const, a loop variable or a
// computation on them is known while compiling, so it becomes a constant
// in the commands: no scores, no macros. A `for` over a list is unrolled,
// one copy of the body per element; a `for` over a range stays a
// scoreboard loop. break and continue work the same in both.

// unrolled is the state of one unrolled loop while compiling its body.
type unrolled struct {
	flag    *fn.Score // set to 1 by break; nil when the body has no break
	wrapped bool      // each iteration runs in its own loop, for break and continue
}

// MaxUnrolled is the warning level for commands from one unrolled loop;
// twice it is an error.
const MaxUnrolled = 256

// evaluator returns an evaluator for the function's current build-time
// scope, in which runtime variables hide build-time names.
func (f *fc) evaluator() *Evaluator {
	base := f.c.evaluator()
	if base.counts == nil {
		base.counts = map[string]int{}
	}
	return &Evaluator{Env: f.env, ctx: base.ctx, counts: base.counts, report: base.report, Uses: base.Uses,
		Shadowed: func(name string) bool { _, ok := f.lookup(name); return ok }}
}

// buildValue computes e when it is a build-time expression. ok is false
// for runtime expressions; v is nil when evaluation failed (reported).
func (f *fc) buildValue(e syntax.Expr) (v *Value, ok bool) {
	if f.env == nil || e == nil {
		return nil, false
	}
	ev := f.evaluator()
	if _, isStr := e.(*syntax.String); isStr {
		return nil, false // strings interpolate at runtime too; see expr
	}
	if !mentionsBuildTime(e, ev) || !wholeBuildTime(e, ev) {
		return nil, false
	}
	return ev.Eval(e), true
}

// buildTime compiles a build-time expression to a constant.
func (f *fc) buildTime(e syntax.Expr) (*val, bool) {
	if s, isStr := e.(*syntax.String); isStr && f.env != nil {
		ev := f.evaluator()
		if text, ok := ev.interpolate(s.Value, s.Sp, false); ok && text != s.Value {
			return &val{typ: tString, str: &text}, true
		}
		return nil, false
	}
	v, ok := f.buildValue(e)
	if !ok {
		return nil, false
	}
	if v == nil {
		return nil, true
	}
	if in := f.c.info(e); in.Help == "" && in.Symbol == nil && in.Local == nil {
		in.Help = "Build-time value: `" + short(v.String()) + "`"
	}
	return f.constVal(v, e.Span())
}

func (f *fc) constVal(v *Value, sp syntax.Span) (*val, bool) {
	switch v.Kind {
	case VNum:
		if !v.Int || v.Num > math.MaxInt32 || v.Num < math.MinInt32 {
			f.c.errorf("mms/type-mismatch", sp, "this is %s; scores are whole numbers from -2147483648 to 2147483647", short(v.String()))
			return nil, true
		}
		return &val{typ: tInt, op: fn.Const(int32(v.Num))}, true
	case VBool:
		n := int32(0)
		if v.Bool {
			n = 1
		}
		return &val{typ: tBool, op: fn.Const(n)}, true
	case VStr, VWord, VTag:
		t := v.Text()
		return &val{typ: tString, str: &t}, true
	}
	f.c.errorf("mms/type-mismatch", sp, "this is %s, which only exists at build time; loop over it with for, or pick one element: xs[0]", v.TypeName())
	return nil, true
}

// buildLet binds `let x = value` at build time when the value is a list,
// record or generator; scalars stay runtime variables holding a constant.
func (f *fc) buildLet(s *syntax.LetStmt) bool {
	if s.Name == nil || s.Value == nil || s.Type != nil {
		return false
	}
	v, ok := f.buildValue(s.Value)
	if !ok || v == nil {
		return ok
	}
	switch v.Kind {
	case VList, VRecord, VRand:
		f.env = f.env.With(s.Name.Text, v)
		f.c.info(s.Name).Local = &Local{Name: s.Name.Text, Type: "build-time " + kindWord(v), What: "let", Doc: "`" + short(v.String()) + "`"}
		return true
	}
	return false
}

func kindWord(v *Value) string {
	switch v.Kind {
	case VList:
		return "list"
	case VRecord:
		return "record"
	case VRand:
		return "Rand"
	}
	return "value"
}

// unrollFor compiles `for x in list`: the body once per element, with x
// bound at build time. It reports false when the loop is not over a
// build-time list, leaving it to the scoreboard loop.
func (f *fc) unrollFor(out *fn.Block, s *syntax.ForStmt) bool {
	if _, isRange := s.Iter.(*syntax.Range); isRange {
		return false
	}
	v, ok := f.buildValue(s.Iter)
	if l, isList := s.Iter.(*syntax.List); isList && !ok && f.env != nil {
		v, ok = f.evaluator().Eval(l), true
	}
	if !ok {
		return false
	}
	if v == nil || s.Var == nil {
		return true
	}
	if v.Kind != VList {
		f.c.errorf("mms/type-mismatch", s.Iter.Span(), "a for loop goes over a list or a range, but this is %s", v.TypeName())
		return true
	}
	if len(v.Elems) > 2*MaxIterations {
		f.c.errorf("mms/expansion-limit", s.Iter.Span(), "this loop runs %d times; the limit is %d", len(v.Elems), 2*MaxIterations)
		return true
	}
	f.c.info(s.Var).Local = &Local{Name: s.Var.Text, Type: "build-time " + elemKinds(v), What: "for", Doc: "Values: `" + short(v.String()) + "`"}
	f.c.info(s.Iter).Help = "Build-time loop: runs " + count(len(v.Elems), "time") + ", unrolled while compiling."
	jumps := hasJump(s.Body)
	u := &unrolled{wrapped: jumps.brk || jumps.cont}
	if jumps.brk {
		t := f.temp()
		u.flag = &t
		out.Add(fn.Set{Dst: t, Src: fn.Const(0)})
	}
	saved := f.env
	f.loops++
	defer func() { f.loops-- }()
	f.unroll = append(f.unroll, u)
	before := countOps(out)
	var iters []*fn.Block
	for i, el := range v.Elems {
		f.env = saved
		if s.Index != nil {
			f.env = f.env.With(s.Index.Text, intValue(i))
		}
		f.env = f.env.With(s.Var.Text, el)
		body := &fn.Block{}
		f.block(body, s.Body)
		iters = append(iters, body)
	}
	f.unroll = f.unroll[:len(f.unroll)-1]
	f.env = saved
	if !u.wrapped {
		for _, b := range iters {
			out.Ops = append(out.Ops, b.Ops...)
		}
	} else {
		// Each iteration is a loop that runs once: continue leaves it
		// for the next iteration, and break also sets the flag that
		// stops the rest.
		all := &fn.Block{}
		for i, b := range iters {
			b.Add(fn.Break{})
			all.Add(fn.Loop{Body: b})
			if u.flag != nil && i < len(iters)-1 {
				all.Add(fn.If{Cond: fn.Cond{Parts: []string{"if score " + u.flag.String() + " matches 1"}}, Then: &fn.Block{Ops: []fn.Op{fn.Break{}}}})
			}
		}
		all.Add(fn.Break{})
		out.Add(fn.Loop{Body: all})
	}
	if n := countOps(out) - before; n > 2*MaxUnrolled {
		f.c.errorf("mms/large-unroll", s.Iter.Span(), "unrolling this loop makes %d commands; the limit is %d. Loop over a range instead, so it runs as a scoreboard loop", n, 2*MaxUnrolled)
	} else if n > MaxUnrolled {
		f.c.warnf("mms/large-unroll", s.Iter.Span(), "unrolling this loop makes %d commands; consider a range loop, which runs as a scoreboard loop", n)
	}
	return true
}

func elemKinds(v *Value) string {
	if len(v.Elems) == 0 {
		return "value"
	}
	k := v.Elems[0].Kind
	for _, e := range v.Elems {
		if e.Kind != k {
			return "value"
		}
	}
	switch k {
	case VNum:
		return "number"
	case VStr:
		return "string"
	case VWord:
		return "id"
	case VRecord:
		return "record"
	case VList:
		return "list"
	case VBool:
		return "bool"
	}
	return "value"
}

func count(n int, word string) string { return strconv.Itoa(n) + " " + plural(n, word) }

// innerUnrolled returns the innermost loop when it is unrolled and has
// a break flag.
func (f *fc) innerUnrolled() *unrolled {
	if len(f.unroll) == 0 {
		return nil
	}
	if u := f.unroll[len(f.unroll)-1]; u != nil && u.flag != nil {
		return u
	}
	return nil
}

type jumps struct{ brk, cont bool }

// hasJump finds break and continue that belong to this loop, not to a
// loop nested in it.
func hasJump(b *syntax.Block) jumps {
	var j jumps
	var walk func(n syntax.Node)
	walk = func(n syntax.Node) {
		switch n.(type) {
		case *syntax.BreakStmt:
			j.brk = true
			return
		case *syntax.ContinueStmt:
			j.cont = true
			return
		case *syntax.WhileStmt, *syntax.ForStmt:
			return
		}
		for _, c := range syntax.Children(n) {
			walk(c)
		}
	}
	walk(b)
	return j
}

func countOps(b *fn.Block) int {
	n := 0
	for _, op := range b.Ops {
		n++
		switch o := op.(type) {
		case fn.If:
			if o.Then != nil {
				n += countOps(o.Then)
			}
			if o.Else != nil {
				n += countOps(o.Else)
			}
		case fn.Loop:
			if o.Body != nil {
				n += countOps(o.Body)
			}
		}
	}
	return n
}
