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
	if sg, start, ok := c.selectorInside(src, ws, off, ids); ok {
		return sg, start
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
			for _, v := range c.storageValues(ch, src[:ws], ids) {
				v.Detail = detail
				add(v)
			}
			for _, v := range c.values(ch, ids) {
				v.Detail = detail
				add(v)
			}
		}
	}
	return out, ws
}

// storageValues suggests the storages after `storage`, as in
// `data modify storage |`, from ids("storage"), and the paths in one
// after its id, `storage demo:mms |`, from ids("storage " + id). Other
// ids have no registry to come from.
func (c *Checker) storageValues(n *spec.CommandNode, before []rune, ids func(string) []string) []Suggestion {
	if ids == nil {
		return nil
	}
	words := strings.Fields(string(before))
	var list []string
	switch strings.TrimPrefix(n.Parser, "minecraft:") {
	case "resource_location":
		if len(words) >= 1 && words[len(words)-1] == "storage" {
			list = ids("storage")
		}
	case "nbt_path":
		if len(words) >= 2 && words[len(words)-2] == "storage" {
			list = ids("storage " + words[len(words)-1])
		}
	}
	out := make([]Suggestion, 0, len(list))
	for _, id := range list {
		out = append(out, Suggestion{Label: id, ID: true})
	}
	return out
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

// selectorInside completes inside the brackets of a selector being typed,
// `@e[type=zombie,ta|`: an option name, or the value of an option that
// takes a known set. ok is false outside open brackets.
func (c *Checker) selectorInside(src []rune, ws, off int, ids func(string) []string) ([]Suggestion, int, bool) {
	word := string(src[ws:off])
	open := strings.LastIndex(word, "[")
	if !strings.HasPrefix(word, "@") || open < 0 || strings.Contains(word[open:], "]") {
		return nil, 0, false
	}
	// The option being typed starts after the last '[' or ',' outside
	// braces, as in scores={a=1,b=2}.
	start, depth := open+1, 0
	for i, r := range word {
		if i <= open {
			continue
		}
		switch r {
		case '{':
			depth++
		case '}':
			depth--
		case ',':
			if depth == 0 {
				start = i + 1
			}
		}
	}
	seg := word[start:]
	at := ws + len([]rune(word[:start]))
	key, val, hasVal := strings.Cut(seg, "=")
	if !hasVal {
		var out []Suggestion
		for _, o := range SelectorOptionNames() {
			out = append(out, Suggestion{Label: o + "=", Detail: "selector option"})
		}
		return out, at, true
	}
	at += len([]rune(key)) + 1
	neg := strings.HasPrefix(val, "!")
	if neg {
		at++
	}
	var vals []string
	id := false
	switch key {
	case "type":
		vals, id = c.pack.RegistryIDs("entity_type"), true
	case "gamemode":
		vals = enums["gamemode"]
	case "sort":
		vals = []string{"nearest", "furthest", "random", "arbitrary"}
	case "predicate":
		if ids != nil {
			vals, id = sorted(ids("predicate")), true
		}
	case "team":
		if ids != nil {
			vals = sorted(ids("team"))
		}
	case "scores":
		// Inside the braces: the objectives, `scores={kills=1..}`.
		inner := val
		if b := strings.LastIndexAny(inner, "{,"); b >= 0 && strings.HasPrefix(val, "{") && !strings.Contains(inner[b:], "=") {
			at += len([]rune(inner[:b+1]))
			if neg {
				at--
			}
			if ids != nil {
				vals = append(vals, sorted(ids("objective"))...)
			}
			break
		}
		return nil, at, true
	}
	out := make([]Suggestion, 0, len(vals))
	for _, v := range vals {
		out = append(out, Suggestion{Label: v, ID: id})
	}
	return out, at, true
}
