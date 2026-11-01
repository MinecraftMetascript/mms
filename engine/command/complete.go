package command

import (
	"sort"
	"strings"

	"github.com/minecraftmetascript/mms/engine/spec"
)

// Suggestion is one completion inside a command.
type Suggestion struct {
	Label  string
	Detail string // the argument it fills, such as "<targets: selector>"
	ID     bool   // an id rather than a literal or keyword
}

// Complete lists what can be typed at off, a rune offset into the command
// text, and the offset where the word being typed starts. ids returns the
// project's own ids for a registry ("function", "objective", "tag/item").
func (c *Checker) Complete(text string, off int, ids func(registry string) []string) ([]Suggestion, int) {
	src := []rune(text)
	if off > len(src) {
		off = len(src)
	}
	ws := off
	for ws > 0 && src[ws-1] != ' ' {
		ws--
	}
	s := &state{c: c, src: src[:ws]}
	var nodes []*spec.CommandNode
	seen := map[*spec.CommandNode]bool{}
	var reach func(n *spec.CommandNode, pos int)
	reach = func(n *spec.CommandNode, pos int) {
		if pos == ws {
			if !seen[n] {
				seen[n] = true
				nodes = append(nodes, n)
			}
			return
		}
		end := s.word(pos)
		w := string(src[pos:end])
		lits, args := names(n)
		for _, name := range lits {
			if name == w && end < ws {
				reach(s.targets(n.Children[name]), end+1)
			}
		}
		for _, name := range args {
			ch := n.Children[name]
			if aend, ok := s.arg(name, ch, pos); ok && aend < ws && src[aend] == ' ' {
				reach(s.targets(ch), aend+1)
			}
		}
	}
	reach(c.pack.Commands, 0)

	var out []Suggestion
	dup := map[string]bool{}
	add := func(sg Suggestion) {
		if !dup[sg.Label] {
			dup[sg.Label] = true
			out = append(out, sg)
		}
	}
	for _, n := range nodes {
		lits, args := names(n)
		for _, l := range lits {
			add(Suggestion{Label: l})
		}
		for _, a := range args {
			ch := n.Children[a]
			detail := argLabel(a, ch)
			for _, v := range c.values(ch, ids) {
				v.Detail = detail
				add(v)
			}
		}
	}
	return out, ws
}

// values suggests values for an argument.
func (c *Checker) values(n *spec.CommandNode, ids func(string) []string) []Suggestion {
	p := strings.TrimPrefix(n.Parser, "minecraft:")
	words := func(ws ...string) []Suggestion {
		out := make([]Suggestion, len(ws))
		for i, w := range ws {
			out[i] = Suggestion{Label: w}
		}
		return out
	}
	registry := func(reg string, tags bool) []Suggestion {
		var out []Suggestion
		if ids != nil {
			for _, id := range ids(reg) {
				out = append(out, Suggestion{Label: id, ID: true})
			}
			if tags {
				for _, id := range ids("tag/" + reg) {
					out = append(out, Suggestion{Label: "#" + id, ID: true})
				}
			}
		}
		for _, id := range c.pack.RegistryIDs(reg) {
			out = append(out, Suggestion{Label: id, ID: true})
		}
		if tags {
			for _, id := range c.pack.RegistryIDs("tag/" + reg) {
				out = append(out, Suggestion{Label: "#" + id, ID: true})
			}
		}
		return out
	}
	switch p {
	case "entity", "game_profile":
		return words("@a", "@e", "@n", "@p", "@r", "@s")
	case "score_holder":
		return words("@a", "@e", "@n", "@p", "@r", "@s", "*")
	case "brigadier:bool":
		return words("false", "true")
	case "block_pos", "vec3":
		return words("~ ~ ~", "^ ^ ^")
	case "column_pos", "vec2", "rotation":
		return words("~ ~")
	case "objective_criteria":
		return words(Criteria...)
	case "scoreboard_slot":
		return words("below_name", "list", "sidebar")
	case "objective", "team":
		if ids == nil {
			return nil
		}
		return words(sorted(ids(p))...)
	case "function":
		return registry("function", true)
	case "resource", "resource_key", "resource_or_tag", "resource_or_tag_key":
		reg, _ := n.Properties["registry"].(string)
		return registry(strings.TrimPrefix(reg, "minecraft:"), strings.HasPrefix(p, "resource_or_tag"))
	}
	if reg, ok := parserRegistry[p]; ok {
		return registry(reg, p == "block_predicate" || p == "item_predicate")
	}
	if vals, ok := enums[p]; ok {
		return words(vals...)
	}
	return nil
}

func sorted(xs []string) []string {
	xs = append([]string(nil), xs...)
	sort.Strings(xs)
	return xs
}
