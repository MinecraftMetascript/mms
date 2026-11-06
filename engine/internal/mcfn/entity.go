package mcfn

import (
	"fmt"
	"strconv"
	"strings"
)

// Entity is a summoned entity: its type, tags and position. The VM has
// no physics, no players and no world, so entities only exist to be
// selected, counted, tagged and killed.
type Entity struct {
	Type string // with its namespace: "minecraft:zombie"
	Tags map[string]bool
	Pos  [3]float64
	NBT  map[string]any // the rest of the summon data, as given
}

func fullID(s string) string {
	if strings.Contains(s, ":") {
		return s
	}
	return "minecraft:" + s
}

// summon <type> [x y z] [nbt]
func (vm *VM) summon(w []string) (result, error) {
	if len(w) == 0 {
		return result{}, fmt.Errorf("summon needs an entity type")
	}
	e := &Entity{Type: fullID(w[0]), Tags: map[string]bool{}, NBT: map[string]any{}}
	rest := w[1:]
	if len(rest) >= 3 {
		pos, err := vm.position(rest[:3])
		if err != nil {
			return result{}, err
		}
		e.Pos, rest = pos, rest[3:]
	} else if len(rest) != 0 && !strings.HasPrefix(rest[0], "{") {
		return result{}, fmt.Errorf("summon needs three coordinates")
	} else {
		e.Pos = vm.pos
	}
	if len(rest) > 0 {
		v, err := ParseSNBT(strings.Join(rest, " "))
		if err != nil {
			return result{}, err
		}
		m, ok := v.(map[string]any)
		if !ok {
			return result{}, fmt.Errorf("summon data must be a compound")
		}
		for k, v := range m {
			if k != "Tags" {
				e.NBT[k] = v
				continue
			}
			l, ok := v.([]any)
			if !ok {
				return result{}, fmt.Errorf("Tags must be a list of strings")
			}
			for _, t := range l {
				s, ok := t.(string)
				if !ok {
					return result{}, fmt.Errorf("Tags must be a list of strings")
				}
				e.Tags[s] = true
			}
		}
	}
	vm.Entities = append(vm.Entities, e)
	return result{value: 1, ok: true}, nil
}

// position reads x y z, absolute or relative (~) to the executor.
func (vm *VM) position(w []string) ([3]float64, error) {
	var p [3]float64
	for i, s := range w {
		base := 0.0
		if rel, ok := strings.CutPrefix(s, "~"); ok {
			base, s = vm.pos[i], rel
			if s == "" {
				p[i] = base
				continue
			}
		} else if strings.HasPrefix(s, "^") {
			return p, fmt.Errorf("the test interpreter does not support local coordinates (%s)", s)
		}
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return p, fmt.Errorf("expected a coordinate, got %q", w[i])
		}
		p[i] = base + f
	}
	return p, nil
}

// kill [targets]
func (vm *VM) kill(w []string) (result, error) {
	sel := "@s"
	if len(w) > 0 {
		sel = w[0]
	}
	es, err := vm.selectEntities(sel)
	if err != nil {
		return result{}, err
	}
	if len(es) == 0 {
		return result{failed: true}, nil
	}
	dead := map[*Entity]bool{}
	for _, e := range es {
		dead[e] = true
	}
	kept := vm.Entities[:0]
	for _, e := range vm.Entities {
		if !dead[e] {
			kept = append(kept, e)
		}
	}
	vm.Entities = kept
	return result{value: int32(len(es)), ok: true}, nil
}

// tag <targets> add|remove <name>
func (vm *VM) tag(w []string) (result, error) {
	if len(w) != 3 || (w[1] != "add" && w[1] != "remove") {
		return result{}, fmt.Errorf("the test interpreter supports tag <targets> add|remove <name>")
	}
	es, err := vm.selectEntities(w[0])
	if err != nil {
		return result{}, err
	}
	n := 0
	for _, e := range es {
		if e.Tags[w[2]] != (w[1] == "add") {
			n++
			if w[1] == "add" {
				e.Tags[w[2]] = true
			} else {
				delete(e.Tags, w[2])
			}
		}
	}
	if n == 0 {
		return result{failed: true}, nil
	}
	return result{value: int32(n), ok: true}, nil
}

// selectEntities resolves @e and @s with the type, tag and limit
// arguments. Other selectors and arguments fail, so a test never passes
// by ignoring a filter. @a selects nobody: the VM has no players.
func (vm *VM) selectEntities(sel string) ([]*Entity, error) {
	head, args := sel, ""
	if i := strings.IndexByte(sel, '['); i >= 0 {
		if !strings.HasSuffix(sel, "]") {
			return nil, fmt.Errorf("unclosed selector %q", sel)
		}
		head, args = sel[:i], sel[i+1:len(sel)-1]
	}
	var pool []*Entity
	switch head {
	case "@e":
		pool = vm.Entities
	case "@s":
		if vm.self != nil {
			pool = []*Entity{vm.self}
		}
	case "@a":
	default:
		return nil, fmt.Errorf("the test interpreter does not support the selector %q", sel)
	}
	limit := -1
	var tests []func(*Entity) bool
	for _, a := range splitArgs(args) {
		k, v, ok := strings.Cut(a, "=")
		if !ok {
			return nil, fmt.Errorf("selector argument %q has no value", a)
		}
		k = strings.TrimSpace(k)
		v = strings.TrimSpace(v)
		neg := false
		if rest, ok := strings.CutPrefix(v, "!"); ok {
			neg, v = true, rest
		}
		switch k {
		case "type":
			if strings.HasPrefix(v, "#") {
				return nil, fmt.Errorf("the test interpreter does not support entity type tags (%s)", v)
			}
			t := fullID(v)
			tests = append(tests, func(e *Entity) bool { return (e.Type == t) != neg })
		case "tag":
			name := v
			tests = append(tests, func(e *Entity) bool {
				if name == "" { // tag= has no tags, tag=! has some
					return (len(e.Tags) == 0) != neg
				}
				return e.Tags[name] != neg
			})
		case "limit":
			n, err := strconv.Atoi(v)
			if err != nil || n < 1 || neg {
				return nil, fmt.Errorf("limit must be a positive integer, got %q", v)
			}
			limit = n
		case "sort":
			if v != "arbitrary" || neg {
				return nil, fmt.Errorf("the test interpreter only supports sort=arbitrary")
			}
		default:
			return nil, fmt.Errorf("the test interpreter does not support the selector argument %q", k)
		}
	}
	var out []*Entity
	for _, e := range pool {
		keep := true
		for _, t := range tests {
			keep = keep && t(e)
		}
		if keep {
			out = append(out, e)
			if len(out) == limit {
				break
			}
		}
	}
	return out, nil
}

// splitArgs splits selector arguments at top-level commas.
func splitArgs(s string) []string {
	if strings.TrimSpace(s) == "" {
		return nil
	}
	var out []string
	depth, start := 0, 0
	var quote byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case quote != 0:
			if c == quote {
				quote = 0
			}
		case c == '"' || c == '\'':
			quote = c
		case c == '{' || c == '[':
			depth++
		case c == '}' || c == ']':
			depth--
		case c == ',' && depth == 0:
			out = append(out, s[start:i])
			start = i + 1
		}
	}
	return append(out, s[start:])
}
