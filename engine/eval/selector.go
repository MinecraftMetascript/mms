package eval

import (
	"strings"

	"github.com/minecraftmetascript/mms/engine/command"
	"github.com/minecraftmetascript/mms/engine/format"
	"github.com/minecraftmetascript/mms/engine/internal/near"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// extendSelector is `guards[tag=awake]`: a selector value, from a const
// or a Template, with more options. Options add to the base, as the game
// reads repeated ones: a tag narrows further, a negated type excludes
// another, scores={...} merge. An option the base already sets that the
// game takes once is an error, since neither copy would be right to
// drop. It returns nil when base is not a selector or x has unnamed
// arguments, such as an objective's score hp[@s].
func (c *ctx) extendSelector(base syntax.Expr, x *syntax.Index) syntax.Expr {
	var head *syntax.Selector
	var have []*syntax.Arg
	switch b := base.(type) {
	case *syntax.Selector:
		head = b
	case *syntax.Index:
		s, ok := b.X.(*syntax.Selector)
		if !ok {
			return nil
		}
		head, have = s, b.Args
	default:
		return nil
	}
	if len(x.Args) == 0 {
		return nil
	}
	for _, a := range x.Args {
		if a.Name == nil {
			return nil
		}
	}
	name := format.Expr(x.X)
	out := append([]*syntax.Arg(nil), have...)
	for _, a := range c.selectorArgs(x.Args) {
		opt := a.Name.Text
		if !knownOption(opt) {
			msg := "unknown selector option '" + opt + "'"
			if best := near.Closest(opt, command.SelectorOptionNames()); best != "" {
				msg += "; did you mean '" + best + "'?"
			}
			c.errorf("mms/selector", a.Name.Sp, msg)
			continue
		}
		i := indexOpt(out, opt)
		if i < 0 {
			out = append(out, a)
			continue
		}
		switch command.SelectorRepeat(opt) {
		case command.SelectorRepeats:
			out = append(out, a)
			continue
		case command.SelectorNegatedRepeats:
			if negated(a.Value) && allNegated(out, opt) {
				out = append(out, a)
				continue
			}
			c.errorf("mms/selector", a.Sp, name+" already sets "+opt+"; only negated ones, such as "+opt+"=!x, can be added")
			continue
		case command.SelectorMap:
			if m := c.mergeMap(out[i], a, name); m != nil {
				out[i] = m
			}
			continue
		}
		c.errorf("mms/selector", a.Sp, name+" already sets "+opt+"; set it here or in "+name+", not both")
	}
	return &syntax.Index{X: head, Args: out, Sp: c.span(x.Sp)}
}

func knownOption(opt string) bool {
	for _, n := range command.SelectorOptionNames() {
		if n == opt {
			return true
		}
	}
	return strings.Contains(opt, "${")
}

func indexOpt(as []*syntax.Arg, opt string) int {
	for i, a := range as {
		if a.Name != nil && a.Name.Text == opt {
			return i
		}
	}
	return -1
}

func negated(e syntax.Expr) bool {
	if u, ok := e.(*syntax.Unary); ok && u.Op == "!" {
		return true
	}
	if s, ok := e.(*syntax.String); ok {
		return strings.HasPrefix(s.Value, "!")
	}
	return false
}

func allNegated(as []*syntax.Arg, opt string) bool {
	for _, a := range as {
		if a.Name != nil && a.Name.Text == opt && !negated(a.Value) {
			return false
		}
	}
	return true
}

// mergeMap merges scores={...} or advancements={...} into the base's.
func (c *ctx) mergeMap(base, add *syntax.Arg, name string) *syntax.Arg {
	bo, ok1 := base.Value.(*syntax.Object)
	ao, ok2 := add.Value.(*syntax.Object)
	if !ok1 || !ok2 || bo.Body == nil || ao.Body == nil {
		c.errorf("mms/selector", add.Sp, name+" already sets "+add.Name.Text+"; write both as {...} to merge them")
		return nil
	}
	keys := map[string]bool{}
	for _, it := range bo.Body.Items {
		if f, ok := it.(*syntax.FieldStmt); ok && f.Key != nil {
			keys[f.Key.Text] = true
		}
	}
	items := append([]syntax.Stmt(nil), bo.Body.Items...)
	for _, it := range ao.Body.Items {
		if f, ok := it.(*syntax.FieldStmt); ok && f.Key != nil && keys[f.Key.Text] {
			c.errorf("mms/selector", f.Key.Sp, name+" already sets "+add.Name.Text+" for "+f.Key.Text)
			continue
		}
		items = append(items, it)
	}
	body := *bo.Body
	body.Items = items
	return &syntax.Arg{Name: base.Name, Value: &syntax.Object{Body: &body, Sp: bo.Sp}, Sp: base.Sp}
}
