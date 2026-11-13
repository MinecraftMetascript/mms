package eval

import (
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// Use is what a check block's used() asks about an instance: one of its
// emit consts, read as name.const, or the declarations of a kind it
// emits, referenced by id.
type Use struct {
	Const string
	Kind  string
}

// later is a check block that calls used(). It runs once the whole
// project is read, since only then is it known what reads the
// instance's outputs.
type later struct {
	c   *ctx
	blk *syntax.Block
}

// callsUsed reports whether a check block calls used().
func callsUsed(blk *syntax.Block) bool {
	found := false
	syntax.Walk(blk, func(n syntax.Node) bool {
		if c, ok := n.(*syntax.Call); ok {
			if id, ok := c.Fn.(*syntax.Ident); ok && id.Name == "used" {
				found = true
			}
		}
		return !found
	})
	return found
}

// RunLater runs the instance's check blocks that call used(), answering
// each used() with used, and returns what they report.
func (in *Instance) RunLater(used func(Use) bool) []syntax.Diagnostic {
	var out []syntax.Diagnostic
	for _, l := range in.later {
		e := *l.c.e
		e.sink, e.steps, e.over = &out, 0, false
		c := l.c.with()
		c.e, c.used = &e, used
		var discard []syntax.Stmt
		c.stmts(l.blk.Items, &discard)
	}
	return out
}

// HasLater reports whether the instance has check blocks that call
// used().
func (in *Instance) HasLater() bool { return len(in.later) > 0 }

// Uses are what the instance's Blueprint asks used() about, as written
// in its check blocks: the outputs whose use it checks itself.
func (in *Instance) Uses() []Use {
	var out []Use
	for _, l := range in.later {
		syntax.Walk(l.blk, func(n syntax.Node) bool {
			if c, ok := n.(*syntax.Call); ok {
				if id, ok := c.Fn.(*syntax.Ident); ok && id.Name == "used" && len(c.Args) == 1 {
					if a, ok := c.Args[0].Value.(*syntax.Ident); ok {
						out = append(out, in.useOf(a.Name))
					}
				}
			}
			return true
		})
	}
	return out
}

// useOf reads the argument of used(name): an emit const of the
// Blueprint, or else a kind of declaration.
func (in *Instance) useOf(name string) Use {
	if emitsConst(in.Blueprint, name) {
		return Use{Const: name}
	}
	return Use{Kind: name}
}

// emitsConst reports whether a const named name is at the top of one of
// bp's emit blocks.
func emitsConst(bp *Blueprint, name string) bool {
	for _, blk := range bp.Emits {
		for _, it := range blk.Items {
			if st, ok := it.(*syntax.ConstStmt); ok && st.Name != nil && st.Name.Text == name {
				return true
			}
		}
	}
	return false
}

// emitsKind reports whether bp's emit blocks declare a kind of
// declaration, at any depth.
func emitsKind(bp *Blueprint, kind string) bool {
	found := false
	for _, blk := range bp.Emits {
		syntax.Walk(blk, func(n syntax.Node) bool {
			if d, ok := n.(*syntax.Decl); ok && d.Kind != nil && d.Kind.Name == kind {
				found = true
			}
			return !found
		})
	}
	return found
}

// usedCall is used(x) in a check block: whether the project reads the
// instance's emit const x, or references a declaration of kind x that
// it emits.
func usedCall(c *ctx, x *syntax.Call, _ []argValue) Value {
	p := c.pos(x.Sp)
	yes := &Bool{p, true}
	var name string
	if len(x.Args) == 1 && x.Args[0].Name == nil {
		if id, ok := x.Args[0].Value.(*syntax.Ident); ok {
			name = id.Name
		}
	}
	if c.inst == nil || name == "" {
		c.errorf("mms/intrinsic-args", x.Sp, "used takes a const of the emit block, such as used(noise_entries), or a kind of declaration it emits, such as used(SurfaceRule)")
		return yes
	}
	bp := c.inst.Blueprint
	if !emitsConst(bp, name) && !emitsKind(bp, name) {
		c.errorf("mms/intrinsic-args", x.Args[0].Sp, bp.D.NameText()+"'s emit block has no const or declaration kind '"+name+"'; used takes a const at the top of the emit block, or a kind of declaration it emits")
		return yes
	}
	if c.used == nil {
		return yes
	}
	return &Bool{p, c.used(c.inst.useOf(name))}
}
