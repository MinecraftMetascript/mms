// Package command parses raw command lines against the Brigadier command
// tree in the spec pack. It reports where a line stops matching and what
// was expected there, and lists what can come next for completion.
//
// Arguments are checked for syntax, not meaning: a selector's brackets
// must balance and its options must be known, but whether an entity
// exists is the game's business. Vanilla ids are checked against the
// registries the pack lists; ids in other namespaces may be defined by the
// project and are left to the caller.
package command

import (
	"sort"
	"strings"

	"github.com/minecraftmetascript/mms/engine/spec"
)

// Problem is a syntax error in a command. Start and End are rune offsets
// into the command text.
type Problem struct {
	Start, End int
	Message    string
}

// ID is a resource id used by a command, for callers that check ids in
// their own namespaces (functions, objectives and so on).
type ID struct {
	Start, End int
	Registry   string // "function", "objective", "worldgen/biome", ...
	Value      string // as written, possibly with '#'
}

// Result is the outcome of checking one command line.
type Result struct {
	Problem *Problem
	IDs     []ID
}

// Checker checks commands against one spec pack.
type Checker struct {
	pack *spec.Pack
}

func New(p *spec.Pack) *Checker { return &Checker{pack: p} }

// interp is the start of an MMS interpolation, which matches any argument.
const interp = "${"

// Check parses one command line, without the leading slash.
func (c *Checker) Check(text string) Result { return c.CheckWith(text, nil) }

// CheckWith is Check with a lookup for ids the project declares, which
// may extend vanilla registries. Registry names are as in the spec pack:
// "function", "tag/item", "worldgen/biome".
func (c *Checker) CheckWith(text string, known func(registry, id string) bool) Result {
	s := &state{c: c, src: []rune(text), known: known}
	if strings.HasPrefix(text, "$") {
		return Result{} // a macro line: its arguments exist only at run time
	}
	ok := s.children(c.pack.Commands, 0, true)
	if ok {
		return Result{IDs: s.okIDs}
	}
	p := s.best
	if p == nil {
		p = &Problem{Start: 0, End: len(s.src), Message: "this command is not valid"}
	}
	return Result{Problem: p}
}

type state struct {
	c     *Checker
	src   []rune
	best  *Problem
	ids   []ID // ids along the current path
	okIDs []ID
	known func(registry, id string) bool
}

// fail records a problem, keeping the one furthest into the line: that is
// the branch the author most likely meant.
func (s *state) fail(start, end int, msg string) bool {
	if end < start {
		end = start
	}
	if s.best == nil || start > s.best.Start {
		s.best = &Problem{Start: start, End: end, Message: msg}
	}
	return false
}

// targets returns the nodes that may follow n: its children, or the
// children of the node it redirects to. A literal with neither, such as
// `execute run`, redirects to the root.
func (s *state) targets(n *spec.CommandNode) *spec.CommandNode {
	if n.Redirect != nil && len(n.Redirect) > 0 {
		t := s.c.pack.Commands
		for _, name := range n.Redirect {
			if t = t.Children[name]; t == nil {
				return &spec.CommandNode{}
			}
		}
		return t
	}
	if len(n.Children) == 0 && !n.Executable && n.Type == "literal" {
		return s.c.pack.Commands
	}
	return n
}

// names returns a node's children: literals first, then arguments, each
// sorted, so matching is deterministic.
func names(n *spec.CommandNode) (lits, args []string) {
	for name, ch := range n.Children {
		if ch.Type == "literal" {
			lits = append(lits, name)
		} else {
			args = append(args, name)
		}
	}
	sort.Strings(lits)
	sort.Strings(args)
	return
}

// word returns the end of the space-delimited word at pos.
func (s *state) word(pos int) int {
	end := pos
	for end < len(s.src) && s.src[end] != ' ' {
		end++
	}
	return end
}

// children matches the rest of the line, from pos, against n's children.
func (s *state) children(n *spec.CommandNode, pos int, root bool) bool {
	lits, args := names(n)
	end := s.word(pos)
	w := string(s.src[pos:end])
	if w == "" {
		if pos >= len(s.src) {
			return s.fail(pos, pos, "expected "+describe(n, lits, args))
		}
		return s.fail(pos, pos+1, "unexpected space")
	}
	wild := strings.HasPrefix(w, interp)
	for _, name := range lits {
		if name == w || wild {
			if s.node(n.Children[name], end) {
				return true
			}
		}
	}
	for _, name := range args {
		ch := n.Children[name]
		mark := len(s.ids)
		aend, ok := s.arg(name, ch, pos)
		if ok && (aend == len(s.src) || s.src[aend] == ' ') && s.node(ch, aend) {
			return true
		}
		if ok && aend < len(s.src) && s.src[aend] != ' ' {
			s.fail(aend, s.word(aend), "expected a space after "+argLabel(name, ch))
		}
		s.ids = s.ids[:mark]
	}
	switch {
	case root:
		return s.fail(pos, end, "unknown command '"+w+"'")
	case len(args) == 0:
		return s.fail(pos, end, "unknown option '"+w+"'; expected "+describe(n, lits, args))
	}
	return s.fail(pos, end, "expected "+describe(n, lits, args))
}

// node continues after a child that matched up to pos.
func (s *state) node(n *spec.CommandNode, pos int) bool {
	if pos >= len(s.src) {
		if n.Executable {
			s.okIDs = append([]ID(nil), s.ids...)
			return true
		}
		t := s.targets(n)
		lits, args := names(t)
		return s.fail(pos, pos, "the command is incomplete; expected "+describe(t, lits, args))
	}
	t := s.targets(n)
	if len(t.Children) == 0 {
		return s.fail(pos, len(s.src), "unexpected text after the end of the command")
	}
	return s.children(t, pos+1, t == s.c.pack.Commands)
}

// describe names what a node accepts, for messages.
func describe(n *spec.CommandNode, lits, args []string) string {
	var parts []string
	if len(lits) > 6 {
		parts = append(parts, strings.Join(lits[:6], ", ")+", ...")
	} else {
		parts = append(parts, lits...)
	}
	for _, a := range args {
		parts = append(parts, argLabel(a, n.Children[a]))
	}
	switch len(parts) {
	case 0:
		return "the end of the command"
	case 1:
		return parts[0]
	}
	return "one of " + strings.Join(parts, ", ")
}

func argLabel(name string, n *spec.CommandNode) string {
	if d, ok := parserNames[strings.TrimPrefix(n.Parser, "minecraft:")]; ok {
		return "<" + name + ": " + d + ">"
	}
	return "<" + name + ">"
}
