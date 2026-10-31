package sema

import (
	"sort"
	"strings"

	"github.com/minecraftmetascript/mms/engine/syntax"
)

// Block state definitions (BlockState, assets/<ns>/blockstates/) select
// models by the block's state properties: variant keys such as
// "facing=north,lit=true" and multipart `when` objects. The schema only
// says they are strings, so the checker reads them against the block's
// properties from the spec pack (decision D77).

// blockState checks a BlockState declaration's variant keys and
// multipart conditions against the properties of the block it is named
// after.
func (c *Checker) blockState(d *syntax.Decl, body *syntax.Block) {
	if d.Name == nil || body == nil {
		return
	}
	ns, path := splitID(d.Name.Text)
	if ns == "" {
		ns = c.NS
	}
	id := ns + ":" + path
	var props map[string][]string
	if ns == "minecraft" {
		if bs := c.Pack.Blocks["minecraft:"+path]; bs != nil {
			props = bs.Properties
		} else if bs := c.Pack.Blocks[path]; bs != nil {
			props = bs.Properties
		}
	}
	if props == nil {
		if ns == "minecraft" {
			msg := "the game has no block '" + id + "', so it never reads this block state"
			if s := closest(path, blockNames(c)); s != "" {
				msg += "; did you mean '" + s + "'?"
			}
			c.warnf("mms/unknown-block", d.Name.Sp, "%s", msg)
		} else {
			msg := "the game only reads block states of its own blocks, and '" + id + "' is not one; name it after a block"
			if s := closest(path, blockNames(c)); s != "" {
				msg += ", such as minecraft:" + s
			}
			c.warnf("mms/unknown-block", d.Name.Sp, "%s", msg)
		}
		return
	}
	for _, it := range body.Items {
		f, ok := it.(*syntax.FieldStmt)
		if !ok {
			continue
		}
		switch f.Key.Text {
		case "variants":
			if b := objectBody(f.Value); b != nil {
				for _, it := range b.Items {
					if v, ok := it.(*syntax.FieldStmt); ok {
						c.variantKey(v.Key, props, id)
					}
				}
			}
		case "multipart":
			l, ok := f.Value.(*syntax.List)
			if !ok {
				continue
			}
			for _, e := range l.Elems {
				b := objectBody(e)
				if b == nil {
					continue
				}
				for _, it := range b.Items {
					if w, ok := it.(*syntax.FieldStmt); ok && w.Key.Text == "when" {
						c.whenCondition(w.Value, props, id)
					}
				}
			}
		}
	}
}

// variantKey checks "prop=value,prop=value"; "" matches every state.
func (c *Checker) variantKey(k *syntax.Name, props map[string][]string, block string) {
	if k.Text == "" {
		return
	}
	seen := map[string]bool{}
	for _, part := range strings.Split(k.Text, ",") {
		name, value, ok := strings.Cut(part, "=")
		if !ok {
			c.errorf("mms/block-state", k.Sp, "a variant key is \"\" or property=value pairs separated by ',', such as \"%s\"", exampleKey(props))
			return
		}
		if seen[name] {
			c.errorf("mms/block-state", k.Sp, "property '%s' is set twice", name)
			return
		}
		seen[name] = true
		if !c.stateValue(k.Sp, name, value, props, block) {
			return
		}
	}
}

// whenCondition checks a multipart `when`: { prop = "a|b" }, or OR / AND
// lists of such objects.
func (c *Checker) whenCondition(e syntax.Expr, props map[string][]string, block string) {
	if l, ok := e.(*syntax.List); ok {
		for _, x := range l.Elems {
			c.whenCondition(x, props, block)
		}
		return
	}
	b := objectBody(e)
	if b == nil {
		return
	}
	c.info(b).Props = props
	for _, it := range b.Items {
		f, ok := it.(*syntax.FieldStmt)
		if !ok {
			continue
		}
		if f.Key.Text == "OR" || f.Key.Text == "AND" {
			c.whenCondition(f.Value, props, block)
			continue
		}
		if !c.stateValue(f.Key.Sp, f.Key.Text, "", props, block) {
			continue
		}
		s, ok := f.Value.(*syntax.String)
		if !ok {
			continue // the schema reports non-strings
		}
		for _, v := range strings.Split(s.Value, "|") {
			if !c.stateValue(s.Sp, f.Key.Text, v, props, block) {
				break
			}
		}
	}
}

// stateValue reports an unknown property, or a value the property cannot
// take. value "" checks only the name.
func (c *Checker) stateValue(sp syntax.Span, name, value string, props map[string][]string, block string) bool {
	values, ok := props[name]
	if !ok {
		names := sortedKeys(props)
		if len(names) == 0 {
			c.errorf("mms/block-state", sp, "%s has no state properties, so its only variant is \"\"", block)
			return false
		}
		msg := "'" + block + "' has no property '" + name + "'"
		if s := closest(name, names); s != "" {
			msg += "; did you mean '" + s + "'?"
		} else {
			msg += "; it has " + strings.Join(names, ", ")
		}
		c.errorf("mms/block-state", sp, "%s", msg)
		return false
	}
	if value == "" {
		return true
	}
	for _, v := range values {
		if v == value {
			return true
		}
	}
	c.errorf("mms/block-state", sp, "'%s' is not a value of %s's '%s'; it is one of %s", value, block, name, strings.Join(values, ", "))
	return false
}

func objectBody(e syntax.Expr) *syntax.Block {
	for {
		p, ok := e.(*syntax.Paren)
		if !ok {
			break
		}
		e = p.X
	}
	if o, ok := e.(*syntax.Object); ok {
		return o.Body
	}
	return nil
}

func exampleKey(props map[string][]string) string {
	var parts []string
	for _, n := range sortedKeys(props) {
		if len(parts) == 2 {
			break
		}
		parts = append(parts, n+"="+props[n][0])
	}
	if len(parts) == 0 {
		return ""
	}
	return strings.Join(parts, ",")
}

func sortedKeys(m map[string][]string) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

func blockNames(c *Checker) []string {
	out := make([]string, 0, len(c.Pack.Blocks))
	for id := range c.Pack.Blocks {
		out = append(out, strings.TrimPrefix(id, "minecraft:"))
	}
	sort.Strings(out)
	return out
}

// blockPropCandidates offers a block's state properties in a multipart
// `when` object, with their values as detail.
func (c *Checker) blockPropCandidates(in *Info, present map[string]bool) []Candidate {
	var out []Candidate
	for _, n := range sortedKeys(in.Props) {
		if present[n] {
			continue
		}
		out = append(out, Candidate{Label: n, Kind: CandField, Insert: n + " = \"$1\"", Detail: strings.Join(in.Props[n], " | "), Rank: 1})
	}
	return out
}
