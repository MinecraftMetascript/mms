package sema

import (
	"fmt"
	"strings"

	"github.com/minecraftmetascript/mms/engine/fn"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// A Test checks functions in the test interpreter:
//
//	Test second_wave_spawns_four {
//	  setup {
//	    game.wave = 1
//	    /summon marker 0 64 0 {Tags:["siege.spawn"]}
//	  }
//	  run next_wave()
//	  expect game.wave == 2
//	}
//
// It compiles to a function, <ns>:mms_test/<name>, that is never written
// to the pack. The function runs setup, then each run and expect in
// order, and returns the number of the first expect that fails (from 1),
// or 0. Before it returns a failure, it keeps both sides of a comparison
// of ints in storage, at <ns>:mms expect.left and expect.right, so the
// runner can say what the values were.

// TestPath is the folder of the functions Tests compile to.
const TestPath = "mms_test"

// ExpectStorage is the storage path a failing expect keeps its values at.
const ExpectStorage = "expect"

// TestExpect is one `expect` of a compiled Test.
type TestExpect struct {
	Sp   syntax.Span
	Text string // as written, without `expect`
	// Values is set when the expect compares two ints, whose values the
	// test function keeps when it fails; Left and Right are their text.
	Values      bool
	Left, Right string
}

// TestID is the id of the function a Test compiles to.
func TestID(ns, name string) string { return ns + ":" + TestPath + "/" + name }

// CompileTest checks a Test and lowers it to a function.
func (c *Checker) CompileTest(d *syntax.Decl, id string, src []rune) (*fn.Func, []TestExpect) {
	if d.Name == nil {
		c.errorf("mms/syntax", d.Kind.Sp, "a Test needs a name: `Test spawns_four { ... }`")
		return nil, nil
	}
	if d.Body == nil {
		c.errorf("mms/missing-value", d.Name.Sp, "a Test needs a body: `Test %s { run f()\n expect ... }`", d.Name.Text)
		return nil, nil
	}
	if d.Params != nil || d.Type != nil || d.Value != nil {
		var extra syntax.Node = d.Params
		if d.Type != nil {
			extra = d.Type
		} else if d.Value != nil {
			extra = d.Value
		}
		c.errorf("mms/syntax", extra.Span(), "a Test takes no parameters, type or value; set things up in its setup { } block")
		return nil, nil
	}
	ns, path, _ := strings.Cut(id, ":")
	f := &fc{c: c, src: src, ns: ns, path: path, obj: fn.Objective(ns), sig: FuncSig{Ret: tInt}, names: map[string]int{}, selfID: id}
	f.push()
	body := &fn.Block{}
	var expects []TestExpect
	for i, it := range d.Body.Items {
		switch s := it.(type) {
		case *syntax.SectionStmt:
			if s.Word == nil || s.Word.Name != "setup" {
				word := ""
				if s.Word != nil {
					word = s.Word.Name
				}
				c.errorf("mms/syntax", s.Span(), "a Test has a setup { } block, not %s { }", word)
				continue
			}
			if i != 0 {
				c.errorf("mms/misplaced", s.Word.Sp, "setup comes first in a Test, before run and expect")
			}
			// Its locals stay in scope for the expects.
			if s.Body != nil {
				for _, st := range s.Body.Items {
					f.stmt(body, st)
				}
			}
		case *syntax.RunStmt:
			f.stmt(body, &syntax.ExprStmt{X: s.X, Sp: s.Sp})
		case *syntax.ExpectStmt:
			if s.X == nil {
				continue
			}
			e := TestExpect{Sp: s.Sp, Text: strings.TrimSpace(string(src[s.X.Span().Start:s.X.Span().End]))}
			e.Values = f.expect(body, s.X, len(expects)+1)
			if b, ok := s.X.(*syntax.Binary); ok && e.Values {
				e.Left = strings.TrimSpace(string(src[b.L.Span().Start:b.L.Span().End]))
				e.Right = strings.TrimSpace(string(src[b.R.Span().Start:b.R.Span().End]))
			}
			expects = append(expects, e)
		default:
			c.errorf("mms/misplaced", it.Span(), "a Test holds a setup { } block, then run and expect lines: `run next_wave()`, `expect game.wave == 2`")
		}
	}
	if len(expects) == 0 {
		c.errorf("mms/missing-value", d.Name.Sp, "Test %s checks nothing; add `expect` lines, as in `expect game.wave == 2`", d.Name.Text)
	}
	zero := fn.Const(0)
	body.Add(fn.Return{Val: &zero})
	return &fn.Func{ID: id, Body: body}, expects
}

// expect lowers `expect x` to a check that returns n when x is false. It
// reports whether the check keeps the two sides of a comparison.
func (f *fc) expect(out *fn.Block, x syntax.Expr, n int) bool {
	fail := func(extra ...fn.Op) *fn.Block {
		v := fn.Const(int32(n))
		return &fn.Block{Ops: append(extra, fn.Return{Val: &v})}
	}
	b, ok := x.(*syntax.Binary)
	if !ok || !compareOps[b.Op] {
		out.Add(fn.If{Cond: *f.not(f.cond(x)), Then: fail()})
		return false
	}
	var p []fn.Op
	l := f.expr(&p, b.L)
	r := f.expr(&p, b.R)
	if l == nil || r == nil {
		return false
	}
	var keeps []fn.Op
	ints := l.typ == tInt && r.typ == tInt
	if ints {
		lo, ro := f.operand(&p, l), f.operand(&p, r)
		l, r = &val{typ: tInt, op: lo}, &val{typ: tInt, op: ro}
		at := fn.Storage(f.ns) + " " + ExpectStorage
		keeps = []fn.Op{keep(at+".left", lo), keep(at+".right", ro)}
	}
	v := f.compareVals(p, b, l, r)
	if v == nil {
		return false
	}
	out.Add(fn.If{Cond: *f.not(f.asCond(nil, v)), Then: fail(keeps...)})
	return ints
}

var compareOps = map[string]bool{"<": true, "<=": true, ">": true, ">=": true, "==": true, "!=": true}

// keep writes an operand to storage.
func keep(at string, o fn.Operand) fn.Op {
	if o.IsConst {
		return fn.Cmd{Text: fmt.Sprintf("data modify storage %s set value %d", at, o.Const)}
	}
	return fn.Cmd{Text: fmt.Sprintf("execute store result storage %s int 1 run scoreboard players get %s", at, o.S)}
}
