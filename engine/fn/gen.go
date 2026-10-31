package fn

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Output is what generating functions produces.
type Output struct {
	// Main is the id of the function itself.
	Main string
	// Files maps function ids ("ns:path") to their lines, the function
	// itself and its generated sub-functions.
	Files map[string][]string
	// Consts lists the constant scores (#c<n>) the code reads.
	Consts map[int32]bool
	// Ret is set when the code uses the shared <ns>:mms_gen/ret helper.
	Ret bool
}

// Objective is the scoreboard objective for a namespace's values.
func Objective(ns string) string { return "mms." + ns }

// Holder returns the score holder of a function's parameter, local or
// temporary, named name, for the function at path.
func Holder(path, name string) string { return "#" + path + "." + name }

// Control-flow codes in #cf.
const (
	cfBreak    = 1
	cfContinue = 2
	cfReturn   = 3
)

type frameKind int

const (
	topFrame  frameKind = iota // the function itself
	loopFrame                  // a loop's own function
	subFrame                   // a branch's function
)

type loopInfo struct {
	id   string
	post *Block
	cont string // the continue helper, made when first needed
}

type frame struct {
	id   string
	kind frameKind
	loop *loopInfo
}

type gen struct {
	ns, path, obj string
	out           *Output
	n             int
}

// Generate lowers f and returns its files. Generated sub-function names
// and flags are numbered in source order, so output is deterministic.
func Generate(f Func) *Output {
	ns, path, _ := strings.Cut(f.ID, ":")
	g := &gen{ns: ns, path: path, obj: Objective(ns), out: &Output{Main: f.ID, Files: map[string][]string{}, Consts: map[int32]bool{}}}
	fr := &frame{id: f.ID, kind: topFrame}
	g.out.Files[f.ID] = g.block(fr, f.Body)
	return g.out
}

func (g *gen) score(holder string) Score { return Score{Holder: holder, Obj: g.obj} }

func (g *gen) cf() Score { return g.score("#cf") }
func (g *gen) rv() Score { return g.score("#rv") }

// constant returns the score holding v, registering it for the load
// function.
func (g *gen) constant(v int32) Score {
	g.out.Consts[v] = true
	return g.score("#c" + strconv.Itoa(int(v)))
}

func (g *gen) next(kind string) string {
	id := fmt.Sprintf("%s:%s/_%s%d", g.ns, g.path, kind, g.n)
	g.n++
	return id
}

func (g *gen) flag() Score {
	s := g.score(Holder(g.path, "_f"+strconv.Itoa(g.n)))
	g.n++
	return s
}

func (g *gen) block(fr *frame, b *Block) []string {
	var lines []string
	if b == nil {
		return nil
	}
	for _, op := range b.Ops {
		lines = append(lines, g.op(fr, op)...)
	}
	return lines
}

func (g *gen) op(fr *frame, op Op) []string {
	switch o := op.(type) {
	case Cmd:
		return []string{o.Text}
	case Set:
		if o.Src.IsConst {
			return []string{fmt.Sprintf("scoreboard players set %s %d", o.Dst, o.Src.Const)}
		}
		if o.Src.S == o.Dst {
			return nil
		}
		return []string{fmt.Sprintf("scoreboard players operation %s = %s", o.Dst, o.Src.S)}
	case Arith:
		src := o.Src
		if src.IsConst && (o.Op == "+=" || o.Op == "-=") && src.Const != -1<<31 {
			v := src.Const
			if o.Op == "-=" {
				v = -v
			}
			if v >= 0 {
				return []string{fmt.Sprintf("scoreboard players add %s %d", o.Dst, v)}
			}
			return []string{fmt.Sprintf("scoreboard players remove %s %d", o.Dst, -v)}
		}
		if src.IsConst {
			src = Of(g.constant(src.Const))
		}
		return []string{fmt.Sprintf("scoreboard players operation %s %s %s", o.Dst, o.Op, src.S)}
	case Call:
		if o.Dst != nil {
			return []string{fmt.Sprintf("execute store result score %s run function %s", *o.Dst, o.Fn)}
		}
		return []string{"function " + o.Fn}
	case Macro:
		return g.macro(o)
	case Return:
		return g.ret(fr, o)
	case Break:
		if fr.kind == loopFrame {
			return []string{"return 0"}
		}
		return []string{fmt.Sprintf("scoreboard players set %s %d", g.cf(), cfBreak), "return 0"}
	case Continue:
		if fr.kind == loopFrame {
			return append(g.block(fr, fr.loop.post), "return run function "+fr.loop.id)
		}
		return []string{fmt.Sprintf("scoreboard players set %s %d", g.cf(), cfContinue), "return 0"}
	case If:
		return g.ifOp(fr, o)
	case Loop:
		return g.loop(fr, o)
	}
	panic(fmt.Sprintf("fn: unknown op %T", op))
}

func (g *gen) ret(fr *frame, o Return) []string {
	if fr.kind == topFrame {
		switch {
		case o.Val == nil:
			return []string{"return 0"}
		case o.Val.IsConst:
			return []string{"return " + o.Val.str()}
		}
		return []string{"return run scoreboard players get " + o.Val.S.String()}
	}
	var lines []string
	if o.Val != nil {
		lines = g.op(fr, Set{Dst: g.rv(), Src: *o.Val})
	}
	return append(lines, fmt.Sprintf("scoreboard players set %s %d", g.cf(), cfReturn), "return 0")
}

// single renders a block as one command when it is one simple operation,
// so `if` can run it inline rather than through a sub-function.
func (g *gen) single(fr *frame, b *Block) (string, bool) {
	if b == nil || len(b.Ops) != 1 {
		return "", false
	}
	switch o := b.Ops[0].(type) {
	case Cmd, Set, Arith, Call:
	case Return:
		if fr.kind != topFrame {
			return "", false
		}
	case Break:
		if fr.kind != loopFrame {
			return "", false
		}
	case Continue:
		if fr.kind != loopFrame || (fr.loop.post != nil && len(fr.loop.post.Ops) > 0) {
			return "", false
		}
	default:
		_ = o
		return "", false
	}
	lines := g.op(fr, b.Ops[0])
	if len(lines) != 1 {
		return "", false
	}
	return lines[0], true
}

// run returns the command that runs b from fr, creating a sub-function
// when needed, and the lines that pass an early exit on afterwards.
func (g *gen) run(fr *frame, b *Block, kind string) (string, []string) {
	if line, ok := g.single(fr, b); ok {
		return line, nil
	}
	id := g.next(kind)
	sub := &frame{id: id, kind: subFrame, loop: fr.loop}
	g.out.Files[id] = g.block(sub, b)
	return "function " + id, g.propagate(fr, exits(b))
}

type exitSet struct{ brk, cont, ret bool }

func (e exitSet) any() bool { return e.brk || e.cont || e.ret }

// exits reports which early exits can leave b when it runs as its own
// function.
func exits(b *Block) exitSet {
	var e exitSet
	if b == nil {
		return e
	}
	for _, op := range b.Ops {
		switch o := op.(type) {
		case Return:
			e.ret = true
		case Break:
			e.brk = true
		case Continue:
			e.cont = true
		case If:
			for _, x := range []exitSet{exits(o.Then), exits(o.Else)} {
				e.brk, e.cont, e.ret = e.brk || x.brk, e.cont || x.cont, e.ret || x.ret
			}
		case Loop:
			e.ret = e.ret || exits(o.Body).ret
		}
	}
	return e
}

// propagate passes early exits from a sub-function on to fr.
func (g *gen) propagate(fr *frame, e exitSet) []string {
	if !e.any() {
		return nil
	}
	cf := g.cf()
	switch fr.kind {
	case topFrame:
		if e.ret {
			g.out.Ret = true
			return []string{fmt.Sprintf("execute if score %s matches %d run return run function %s:mms_gen/ret", cf, cfReturn, g.ns)}
		}
		return nil
	case loopFrame:
		var lines []string
		if e.brk {
			lines = append(lines, fmt.Sprintf("execute if score %s matches %d run return run scoreboard players set %s 0", cf, cfBreak, cf))
		}
		if e.cont {
			lines = append(lines, fmt.Sprintf("execute if score %s matches %d run return run function %s", cf, cfContinue, g.contHelper(fr)))
		}
		if e.ret {
			lines = append(lines, fmt.Sprintf("execute if score %s matches %d run return 0", cf, cfReturn))
		}
		return lines
	}
	return []string{fmt.Sprintf("execute if score %s matches 1.. run return 0", cf)}
}

// contHelper makes the function that finishes a continue that came from
// a sub-function: it clears #cf, runs the loop's post step and loops.
func (g *gen) contHelper(fr *frame) string {
	l := fr.loop
	if l.cont == "" {
		l.cont = g.next("continue")
		lines := []string{fmt.Sprintf("scoreboard players set %s 0", g.cf())}
		lines = append(lines, g.block(fr, l.post)...)
		g.out.Files[l.cont] = append(lines, "return run function "+l.id)
	}
	return l.cont
}

// chain joins execute parts and a command, merging nested executes.
func chain(parts []string, cmd string) string {
	if len(parts) == 0 {
		return cmd
	}
	head := "execute " + strings.Join(parts, " ")
	if rest, ok := strings.CutPrefix(cmd, "execute "); ok {
		return head + " " + rest
	}
	return head + " run " + cmd
}

func (g *gen) ifOp(fr *frame, o If) []string {
	lines := g.block(fr, &Block{Ops: o.Cond.Pre})
	switch {
	case o.Cond.Never:
		return append(lines, g.block(fr, o.Else)...)
	case o.Cond.Always():
		return append(lines, g.block(fr, o.Then)...)
	}
	if o.Else == nil || len(o.Else.Ops) == 0 {
		if o.Then == nil || len(o.Then.Ops) == 0 {
			return lines
		}
		cmd, after := g.run(fr, o.Then, "if")
		return append(append(lines, chain(o.Cond.Parts, cmd)), after...)
	}
	// The condition is evaluated once, before either branch runs.
	f := g.flag()
	lines = append(lines,
		fmt.Sprintf("scoreboard players set %s 0", f),
		chain(o.Cond.Parts, fmt.Sprintf("scoreboard players set %s 1", f)))
	if o.Then != nil && len(o.Then.Ops) > 0 {
		cmd, after := g.run(fr, o.Then, "if")
		lines = append(append(lines, chain([]string{fmt.Sprintf("if score %s matches 1", f)}, cmd)), after...)
	}
	cmd, after := g.run(fr, o.Else, "else")
	return append(append(lines, chain([]string{fmt.Sprintf("unless score %s matches 1", f)}, cmd)), after...)
}

// Negate flips a one-part condition. It returns false for chains, whose
// negation is not a chain.
func Negate(parts []string) ([]string, bool) {
	if len(parts) != 1 {
		return nil, false
	}
	p := parts[0]
	if rest, ok := strings.CutPrefix(p, "if "); ok {
		return []string{"unless " + rest}, true
	}
	if rest, ok := strings.CutPrefix(p, "unless "); ok {
		return []string{"if " + rest}, true
	}
	return nil, false
}

func (g *gen) loop(fr *frame, o Loop) []string {
	if o.Cond.Never {
		return g.block(fr, &Block{Ops: o.Cond.Pre})
	}
	id := g.next("while")
	l := &frame{id: id, kind: loopFrame, loop: &loopInfo{id: id, post: o.Post}}
	lines := g.block(l, &Block{Ops: o.Cond.Pre})
	if !o.Cond.Always() {
		if neg, ok := Negate(o.Cond.Parts); ok {
			lines = append(lines, chain(neg, "return 0"))
		} else {
			f := g.flag()
			lines = append(lines,
				fmt.Sprintf("scoreboard players set %s 0", f),
				chain(o.Cond.Parts, fmt.Sprintf("scoreboard players set %s 1", f)),
				fmt.Sprintf("execute unless score %s matches 1 run return 0", f))
		}
	}
	lines = append(lines, g.block(l, o.Body)...)
	lines = append(lines, g.block(l, o.Post)...)
	// A tail call that returns, so the loop function always has a result
	// for `return run function` on the continue path.
	lines = append(lines, "return run function "+id)
	g.out.Files[id] = lines
	return append([]string{"function " + id}, g.propagate(fr, exitSet{ret: exits(o.Body).ret})...)
}

// Setup returns the load-function lines a namespace's generated code
// needs: its objective, constants and cleared control flow state.
func Setup(ns string, consts map[int32]bool) []string {
	obj := Objective(ns)
	lines := []string{
		"scoreboard objectives add " + obj + " dummy",
		"scoreboard players set #cf " + obj + " 0",
	}
	vs := make([]int, 0, len(consts))
	for v := range consts {
		vs = append(vs, int(v))
	}
	sort.Ints(vs)
	for _, v := range vs {
		lines = append(lines, fmt.Sprintf("scoreboard players set #c%d %s %d", v, obj, v))
	}
	return lines
}

// RetHelper is the body of <ns>:mms_gen/ret, which finishes a return that
// came from a sub-function.
func RetHelper(ns string) []string {
	obj := Objective(ns)
	return []string{
		"scoreboard players set #cf " + obj + " 0",
		"return run scoreboard players get #rv " + obj,
	}
}

// ID returns the id of the function itself.
func (o *Output) ID() string { return o.Main }

// macro copies the arguments into storage and calls a one-line macro
// function.
func (g *gen) macro(o Macro) []string {
	st := Storage(g.ns)
	var lines []string
	for _, a := range o.Args {
		path := "m." + a.Key
		switch {
		case a.Score != nil:
			lines = append(lines, fmt.Sprintf("execute store result storage %s %s int 1 run scoreboard players get %s", st, path, *a.Score))
		case a.Text != nil:
			lines = append(lines, fmt.Sprintf("data modify storage %s %s set value %s", st, path, QuoteSNBT(*a.Text)))
		default:
			lines = append(lines, fmt.Sprintf("data modify storage %s %s set from storage %s", st, path, a.Path))
		}
	}
	id := g.next("macro")
	g.out.Files[id] = []string{"$" + o.Line}
	return append(lines, fmt.Sprintf("function %s with storage %s m", id, st))
}

// QuoteSNBT quotes a string for SNBT.
func QuoteSNBT(s string) string {
	return `"` + strings.NewReplacer(`\`, `\\`, `"`, `\"`).Replace(s) + `"`
}
