package fn

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/internal/mcfn"
)

const obj = "mms.demo"

func sc(holder string) Score { return Score{Holder: holder, Obj: obj} }

// tmp is a compiler temporary of function f.
func tmp(n int) Score { return sc(GenHolder("f", fmt.Sprintf("t%d", n))) }

var (
	a = sc("$a")
	b = sc("$b")
	x = sc("$x")
)

// TestCoalesce checks that a temporary used only to build a value is
// replaced by the value's destination, in a block, in a condition's
// pre-ops and in call arguments, and kept when it is read elsewhere.
func TestCoalesce(t *testing.T) {
	chain := func(dst Score) []Op {
		return []Op{Set{Dst: tmp(0), Src: Of(a)}, Arith{Dst: tmp(0), Op: "+=", Src: Of(b)}, Set{Dst: dst, Src: Of(tmp(0))}}
	}
	want := []Op{Set{Dst: x, Src: Of(a)}, Arith{Dst: x, Op: "+=", Src: Of(b)}}

	f := &Func{ID: "demo:f", Body: &Block{Ops: chain(x)}}
	Optimize(f)
	if !reflect.DeepEqual(f.Body.Ops, want) {
		t.Errorf("block:\n got %v\nwant %v", f.Body.Ops, want)
	}

	f = &Func{ID: "demo:f", Body: &Block{Ops: []Op{If{Cond: Cond{Pre: chain(x), Parts: []string{"if score $x " + obj + " matches 1.."}}, Then: &Block{Ops: []Op{Cmd{Text: "say hi"}}}}}}}
	Optimize(f)
	if got := f.Body.Ops[0].(If).Cond.Pre; !reflect.DeepEqual(got, want) {
		t.Errorf("if condition:\n got %v\nwant %v", got, want)
	}

	f = &Func{ID: "demo:f", Body: &Block{Ops: []Op{Loop{Cond: Cond{Pre: chain(x), Parts: []string{"if score $x " + obj + " matches ..9"}}, Body: &Block{}}}}}
	Optimize(f)
	if got := f.Body.Ops[0].(Loop).Cond.Pre; !reflect.DeepEqual(got, want) {
		t.Errorf("loop condition:\n got %v\nwant %v", got, want)
	}

	// Read again later: the temporary stays.
	ops := append(chain(x), Cmd{Text: "scoreboard players operation $b " + obj + " = " + tmp(0).String()})
	f = &Func{ID: "demo:f", Body: &Block{Ops: append([]Op(nil), ops...)}}
	Optimize(f)
	if !reflect.DeepEqual(f.Body.Ops, ops) {
		t.Errorf("shared temporary was rewritten:\n got %v\nwant %v", f.Body.Ops, ops)
	}
}

// TestOptimizeKeepsBehaviour runs IR programs with and without the
// optimiser in the test interpreter and compares what they leave behind.
func TestOptimizeKeepsBehaviour(t *testing.T) {
	ifPos := func(s Score) Cond { return Cond{Parts: []string{"if score " + s.String() + " matches 1.."}} }
	programs := map[string][]Op{
		"arithmetic": {
			Set{Dst: a, Src: Const(7)}, Set{Dst: b, Src: Const(5)},
			Set{Dst: tmp(0), Src: Of(a)}, Arith{Dst: tmp(0), Op: "*=", Src: Of(b)}, Arith{Dst: tmp(0), Op: "-=", Src: Const(3)},
			Set{Dst: x, Src: Of(tmp(0))},
			Return{Val: &Operand{S: x}},
		},
		"branch on a computed value": {
			Set{Dst: a, Src: Const(2)},
			If{
				Cond: Cond{Pre: []Op{Set{Dst: tmp(0), Src: Of(a)}, Arith{Dst: tmp(0), Op: "-=", Src: Const(1)}, Set{Dst: x, Src: Of(tmp(0))}}, Parts: ifPos(x).Parts},
				Then: &Block{Ops: []Op{Set{Dst: b, Src: Const(10)}}},
				Else: &Block{Ops: []Op{Set{Dst: b, Src: Const(20)}}},
			},
			Return{Val: &Operand{S: b}},
		},
		"loop with break": {
			Set{Dst: a, Src: Const(0)}, Set{Dst: b, Src: Const(0)},
			Loop{
				Cond: Cond{Pre: []Op{Set{Dst: tmp(1), Src: Const(10)}, Arith{Dst: tmp(1), Op: "-=", Src: Of(a)}, Set{Dst: x, Src: Of(tmp(1))}}, Parts: ifPos(x).Parts},
				Body: &Block{Ops: []Op{
					Arith{Dst: b, Op: "+=", Src: Of(a)},
					If{Cond: Cond{Parts: []string{"if score " + b.String() + " matches 15.."}}, Then: &Block{Ops: []Op{Break{}}}},
				}},
				Post: &Block{Ops: []Op{Arith{Dst: a, Op: "+=", Src: Const(1)}}},
			},
			Return{Val: &Operand{S: b}},
		},
	}
	want := map[string]int32{"arithmetic": 32, "branch on a computed value": 10, "loop with break": 15}
	for name, ops := range programs {
		t.Run(name, func(t *testing.T) {
			run := func(optimize bool) (int32, map[string]int32) {
				out := generate(Func{ID: "demo:f", Body: &Block{Ops: cloneOps(ops)}}, optimize)
				files := out.Files
				files["demo:load"] = Setup("demo", out.Consts)
				files["demo:mms_gen/ret"] = RetHelper("demo")
				vm := mcfn.New(files)
				if _, _, err := vm.Run("demo:load"); err != nil {
					t.Fatal(err)
				}
				v, _, err := vm.Run("demo:f")
				if err != nil {
					t.Fatalf("optimize=%v: %v\n%s", optimize, err, dump(files))
				}
				globals := map[string]int32{}
				for h, s := range vm.Scores[obj] {
					if strings.HasPrefix(h, "$") {
						globals[h] = s
					}
				}
				return v, globals
			}
			v0, g0 := run(false)
			v1, g1 := run(true)
			if v0 != want[name] {
				t.Errorf("unoptimised returned %d, want %d", v0, want[name])
			}
			if v0 != v1 || !reflect.DeepEqual(g0, g1) {
				t.Errorf("unoptimised returned %d with %v, optimised %d with %v", v0, g0, v1, g1)
			}
		})
	}
}

// cloneOps copies a program's ops and blocks, since Optimize rewrites in
// place.
func cloneOps(ops []Op) []Op {
	out := make([]Op, len(ops))
	for i, op := range ops {
		switch o := op.(type) {
		case If:
			o.Cond.Pre = cloneOps(o.Cond.Pre)
			o.Then, o.Else = cloneBlock(o.Then), cloneBlock(o.Else)
			op = o
		case Loop:
			o.Cond.Pre = cloneOps(o.Cond.Pre)
			o.Body, o.Post = cloneBlock(o.Body), cloneBlock(o.Post)
			op = o
		}
		out[i] = op
	}
	return out
}

func cloneBlock(b *Block) *Block {
	if b == nil {
		return nil
	}
	return &Block{Ops: cloneOps(b.Ops)}
}

func dump(files map[string][]string) string {
	var sb strings.Builder
	for id, lines := range files {
		sb.WriteString("# " + id + "\n" + strings.Join(lines, "\n") + "\n")
	}
	return sb.String()
}
