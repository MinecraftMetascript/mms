package command

import (
	"strings"

	"github.com/minecraftmetascript/mms/engine/spec"
)

// Usage is the signature of the command node a word of a line matched,
// in the game's /help syntax: `<name>` for an argument, `[...]` for what
// may be left out, `(a|b)` for a choice and `-> execute` for a redirect.
type Usage struct {
	// Start and End are the rune range of the word.
	Start, End int
	// Lines are the forms the command can take from that node on, each
	// starting with the words matched so far.
	Lines []string
	// Arg describes the word when it is an argument, such as
	// "<target>: selector"; it is empty for a literal.
	Arg string
}

// usageLines is how many lines a hover shows in full before forms with
// many variants are folded into one line each.
const usageLines = 12

// Usage returns the signature at rune offset off of a command line
// (without the leading slash): the forms the command can take from the
// word at off.
func (c *Checker) Usage(text string, off int) (Usage, bool) {
	if strings.HasPrefix(text, "$") {
		return Usage{}, false
	}
	s := &state{c: c, src: []rune(text)}
	ok := s.children(c.pack.Commands, 0, true)
	trail := s.bestTrail
	if ok {
		trail = s.okTrail
	}
	if len(trail) == 0 {
		return Usage{}, false
	}
	i := len(trail) - 1
	if off < trail[0].end {
		i = 0
	} else {
		for j, st := range trail {
			if st.start <= off && off <= st.end {
				i = j
				break
			}
		}
	}
	var prefix []string
	for _, st := range trail[:i+1] {
		prefix = append(prefix, label(st.name, st.n))
	}
	st := trail[i]
	u := Usage{Start: st.start, End: st.end, Lines: c.expand(strings.Join(prefix, " "), st.n)}
	if st.n.Type != "literal" {
		u.Arg = argLabel(st.name, st.n)
	}
	return u, true
}

func label(name string, n *spec.CommandNode) string {
	if n.Type == "literal" {
		return name
	}
	return "<" + name + ">"
}

// next returns what may follow n and whether it goes back to the root,
// as `execute run` does.
func (c *Checker) next(n *spec.CommandNode) (*spec.CommandNode, bool) {
	t := (&state{c: c}).targets(n)
	return t, t == c.pack.Commands
}

// expand lists the forms after a hovered node, every one in full when
// they fit and folding the larger branches otherwise.
func (c *Checker) expand(prefix string, n *spec.CommandNode) []string {
	var lines []string
	if len(n.Redirect) > 0 {
		lines = append(lines, prefix+" -> "+strings.Join(n.Redirect, " "))
	}
	t, root := c.next(n)
	if root {
		return append(lines, prefix+" ...")
	}
	lits, args := names(t)
	kids := append(lits, args...)
	if len(kids) == 0 {
		return append(lines, prefix)
	}
	full := make([][]string, len(kids))
	total := 0
	for i, k := range kids {
		full[i] = c.forms(t.Children[k], 0)
		total += len(full[i])
	}
	for i, k := range kids {
		ch := t.Children[k]
		var tails []string
		if total <= usageLines || len(full[i]) <= 2 {
			tails = full[i]
		} else {
			tails = []string{c.smart(ch)}
		}
		for _, tail := range tails {
			part := label(k, ch) + tail
			if n.Executable {
				part = "[" + part + "]"
			}
			lines = append(lines, prefix+" "+part)
		}
	}
	return lines
}

// forms lists every form after n (each starting with a space, or empty
// at the end of a command), with what may be left out in brackets.
// Redirects are not followed, so the tree is finite.
func (c *Checker) forms(n *spec.CommandNode, depth int) []string {
	if len(n.Redirect) > 0 {
		return []string{" -> " + strings.Join(n.Redirect, " ")}
	}
	if _, root := c.next(n); root {
		return []string{" ..."}
	}
	if depth > 16 {
		return []string{" ..."}
	}
	lits, args := names(n)
	var out []string
	for _, k := range append(lits, args...) {
		ch := n.Children[k]
		for _, tail := range c.forms(ch, depth+1) {
			if n.Executable {
				out = append(out, " ["+label(k, ch)+tail+"]")
			} else {
				out = append(out, " "+label(k, ch)+tail)
			}
		}
	}
	if len(out) == 0 {
		return []string{""}
	}
	return out
}

// smart folds the forms after n into one line, as the game's /help does:
// a single child is followed, several become a choice.
func (c *Checker) smart(n *spec.CommandNode) string {
	if len(n.Redirect) > 0 {
		return " -> " + strings.Join(n.Redirect, " ")
	}
	if _, root := c.next(n); root {
		return " ..."
	}
	lits, args := names(n)
	kids := append(lits, args...)
	switch len(kids) {
	case 0:
		return ""
	case 1:
		ch := n.Children[kids[0]]
		if n.Executable {
			return " [" + label(kids[0], ch) + "]"
		}
		return " " + label(kids[0], ch) + c.smart(ch)
	}
	var parts []string
	for _, k := range kids {
		parts = append(parts, label(k, n.Children[k]))
	}
	if n.Executable {
		return " [" + strings.Join(parts, "|") + "]"
	}
	return " (" + strings.Join(parts, "|") + ")"
}
