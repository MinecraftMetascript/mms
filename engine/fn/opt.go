package fn

import "strings"

// Optimize rewrites a function's IR in place:
//
//   - copy coalescing: `t = a; t op= b; x = t`, where the temporary t is
//     used nowhere else, becomes `x = a; x op= b`
//   - frame pruning: a recursive call only saves the scores that are
//     read after it returns
func Optimize(f *Func) {
	uses := map[Score]int{}
	countUses(f.Body, uses)
	coalesce(f.Body, uses)
	l := &liveness{prune: true}
	l.block(f.Body, nil)
}

func isTemp(s Score) bool { return strings.Contains(s.Holder, "._t") }

// countUses counts each score's appearances, in operands and in text.
func countUses(b *Block, uses map[Score]int) {
	if b == nil {
		return
	}
	add := func(s Score) { uses[s]++ }
	opnd := func(o Operand) {
		if !o.IsConst {
			add(o.S)
		}
	}
	text := func(t string) {
		for _, w := range scoreRefs(t) {
			add(w)
		}
	}
	for _, op := range b.Ops {
		switch o := op.(type) {
		case Cmd:
			text(o.Text)
		case Set:
			add(o.Dst)
			opnd(o.Src)
		case Arith:
			add(o.Dst)
			opnd(o.Src)
		case Macro:
			for _, a := range o.Args {
				if a.Score != nil {
					add(*a.Score)
				}
			}
		case Call:
			if o.Dst != nil {
				add(*o.Dst)
			}
			countUses(&Block{Ops: o.Args}, uses)
		case If:
			countUses(&Block{Ops: o.Cond.Pre}, uses)
			for _, p := range o.Cond.Parts {
				text(p)
			}
			countUses(o.Then, uses)
			countUses(o.Else, uses)
		case Loop:
			countUses(&Block{Ops: o.Cond.Pre}, uses)
			for _, p := range o.Cond.Parts {
				text(p)
			}
			countUses(o.Body, uses)
			countUses(o.Post, uses)
		case Exec:
			for _, p := range o.Parts {
				text(p)
			}
			countUses(o.Body, uses)
		case Return:
			if o.Val != nil {
				opnd(*o.Val)
			}
		}
	}
}

// scoreRefs finds "<holder> <objective>" pairs for backend objectives in
// command text.
func scoreRefs(t string) []Score {
	var out []Score
	ws := strings.Fields(t)
	for i := 1; i < len(ws); i++ {
		if strings.HasPrefix(ws[i], "mms.") {
			out = append(out, Score{Holder: ws[i-1], Obj: ws[i]})
		}
	}
	return out
}

func coalesce(b *Block, uses map[Score]int) {
	if b == nil {
		return
	}
	var out []Op
	ops := b.Ops
	for i := 0; i < len(ops); i++ {
		if set, ok := ops[i].(Set); ok && isTemp(set.Dst) {
			t := set.Dst
			j := i + 1
			for j < len(ops) {
				a, ok := ops[j].(Arith)
				if !ok || a.Dst != t {
					break
				}
				j++
			}
			if j < len(ops) {
				if fin, ok := ops[j].(Set); ok && !fin.Src.IsConst && fin.Src.S == t && uses[t] == j-i+1 && fin.Dst != t {
					x := fin.Dst
					safe := true
					for _, op := range ops[i+1 : j] {
						if a := op.(Arith); !a.Src.IsConst && a.Src.S == x {
							safe = false
						}
					}
					if safe {
						out = append(out, Set{Dst: x, Src: set.Src})
						for _, op := range ops[i+1 : j] {
							a := op.(Arith)
							out = append(out, Arith{Dst: x, Op: a.Op, Src: a.Src})
						}
						i = j
						continue
					}
				}
			}
		}
		switch o := ops[i].(type) {
		case If:
			coalesce(&Block{Ops: o.Cond.Pre}, uses)
			coalesce(o.Then, uses)
			coalesce(o.Else, uses)
		case Loop:
			coalesce(o.Body, uses)
			coalesce(o.Post, uses)
		case Exec:
			coalesce(o.Body, uses)
		case Call:
			args := &Block{Ops: o.Args}
			coalesce(args, uses)
			o.Args = args.Ops
			ops[i] = o
		}
		out = append(out, ops[i])
	}
	b.Ops = out
}

// set is a set of scores.
type set map[Score]bool

func (s set) clone() set {
	c := make(set, len(s))
	for k := range s {
		c[k] = true
	}
	return c
}

func (s set) union(o set) set {
	c := s.clone()
	for k := range o {
		c[k] = true
	}
	return c
}

func (s set) equal(o set) bool {
	if len(s) != len(o) {
		return false
	}
	for k := range s {
		if !o[k] {
			return false
		}
	}
	return true
}

// loopCtx is what is live where break and continue go.
type loopCtx struct {
	after, cont set
}

// liveness computes, backwards, which scores are read before being
// written again, and prunes recursive calls' frames to those.
type liveness struct {
	loops []loopCtx
	// prune is off while a loop's sets are still growing, so frames are
	// only pruned with the final sets.
	prune bool
}

func reads(o Operand, s set) {
	if !o.IsConst {
		s[o.S] = true
	}
}

func readText(t string, s set) {
	for _, r := range scoreRefs(t) {
		s[r] = true
	}
}

// block returns what is live before b, given what is live after it.
func (l *liveness) block(b *Block, after set) set {
	live := set{}
	if after != nil {
		live = after.clone()
	}
	if b == nil {
		return live
	}
	for i := len(b.Ops) - 1; i >= 0; i-- {
		live = l.op(b, i, live)
	}
	return live
}

func (l *liveness) op(b *Block, i int, live set) set {
	switch o := b.Ops[i].(type) {
	case Cmd:
		readText(o.Text, live)
	case Set:
		delete(live, o.Dst)
		reads(o.Src, live)
	case Arith:
		live[o.Dst] = true
		reads(o.Src, live)
	case Macro:
		for _, a := range o.Args {
			if a.Score != nil {
				live[*a.Score] = true
			}
		}
	case Call:
		if o.Dst != nil {
			delete(live, *o.Dst)
		}
		if o.Recursive && l.prune {
			var keep []Score
			for _, s := range o.Frame {
				if live[s] {
					keep = append(keep, s)
				}
			}
			o.Frame = keep
			b.Ops[i] = o
		}
		// Saved scores are read before the arguments are stored.
		args := &Block{Ops: o.Args}
		live = l.block(args, live)
		for _, s := range o.Frame {
			live[s] = true
		}
	case If:
		then := l.block(o.Then, live)
		els := l.block(o.Else, live)
		live = then.union(els)
		live = l.cond(o.Cond, live)
	case Loop:
		after := live
		head := after.clone()
		pass := func() set {
			post := l.block(o.Post, head)
			l.loops = append(l.loops, loopCtx{after: after, cont: post})
			body := l.block(o.Body, post)
			l.loops = l.loops[:len(l.loops)-1]
			return l.cond(o.Cond, body.union(after))
		}
		prune := l.prune
		l.prune = false
		for k := 0; k < 100; k++ {
			next := pass().union(head)
			if next.equal(head) {
				break
			}
			head = next
		}
		l.prune = prune
		if prune {
			pass()
		}
		live = head
	case Exec:
		// The body runs any number of times, like a loop that always
		// holds and ends after any pass.
		live = l.op(&Block{Ops: []Op{Loop{Body: o.Body}}}, 0, live)
		for _, p := range o.Parts {
			readText(p, live)
		}
	case Return:
		live = set{}
		if o.Val != nil {
			reads(*o.Val, live)
		}
	case Break:
		live = set{}
		if n := len(l.loops); n > 0 {
			live = l.loops[n-1].after.clone()
		}
	case Continue:
		live = set{}
		if n := len(l.loops); n > 0 {
			live = l.loops[n-1].cont.clone()
		}
	}
	return live
}

func (l *liveness) cond(c Cond, live set) set {
	live = live.clone()
	for _, p := range c.Parts {
		readText(p, live)
	}
	return l.block(&Block{Ops: c.Pre}, live)
}
