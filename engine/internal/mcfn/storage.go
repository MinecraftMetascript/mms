package mcfn

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

// Storage values are map[string]any (compounds), []any (lists), string,
// int32 (and other ints as int32), float64.

func (vm *VM) root(id string) map[string]any {
	m, ok := vm.Storage[id]
	if !ok {
		m = map[string]any{}
		vm.Storage[id] = m
	}
	return m
}

type step struct {
	key   string
	index *int
}

func parsePath(p string) ([]step, error) {
	var out []step
	for p != "" {
		switch {
		case p[0] == '.':
			p = p[1:]
		case p[0] == '[':
			j := strings.IndexByte(p, ']')
			if j < 0 {
				return nil, fmt.Errorf("unclosed [ in path")
			}
			n, err := strconv.Atoi(p[1:j])
			if err != nil {
				return nil, fmt.Errorf("unsupported path index %q", p[1:j])
			}
			out = append(out, step{index: &n})
			p = p[j+1:]
		case p[0] == '"':
			j := strings.IndexByte(p[1:], '"')
			if j < 0 {
				return nil, fmt.Errorf("unclosed quote in path")
			}
			out = append(out, step{key: p[1 : j+1]})
			p = p[j+2:]
		default:
			j := strings.IndexAny(p, ".[")
			if j < 0 {
				j = len(p)
			}
			out = append(out, step{key: p[:j]})
			p = p[j:]
		}
	}
	return out, nil
}

func (vm *VM) get(id, path string) (any, error) {
	steps, err := parsePath(path)
	if err != nil {
		return nil, err
	}
	var cur any = vm.root(id)
	for _, s := range steps {
		switch c := cur.(type) {
		case map[string]any:
			if s.index != nil {
				return nil, fmt.Errorf("index on a compound")
			}
			v, ok := c[s.key]
			if !ok {
				return nil, fmt.Errorf("nothing at %s", path)
			}
			cur = v
		case []any:
			if s.index == nil {
				return nil, fmt.Errorf("key on a list")
			}
			i := *s.index
			if i < 0 {
				i += len(c)
			}
			if i < 0 || i >= len(c) {
				return nil, fmt.Errorf("nothing at %s", path)
			}
			cur = c[i]
		default:
			return nil, fmt.Errorf("nothing at %s", path)
		}
	}
	return cur, nil
}

// put sets the value at path, creating compounds on the way. With remove
// set, it deletes it instead.
func (vm *VM) put(id, path string, v any, remove bool) error {
	steps, err := parsePath(path)
	if err != nil {
		return err
	}
	if len(steps) == 0 {
		m, ok := v.(map[string]any)
		if !ok {
			return fmt.Errorf("storage root must be a compound")
		}
		vm.Storage[id] = m
		return nil
	}
	var cur any = vm.root(id)
	for i, s := range steps {
		last := i == len(steps)-1
		switch c := cur.(type) {
		case map[string]any:
			if last {
				if remove {
					delete(c, s.key)
				} else {
					c[s.key] = v
				}
				return nil
			}
			next, ok := c[s.key]
			if !ok {
				if remove {
					return nil
				}
				next = map[string]any{}
				c[s.key] = next
			}
			cur = next
		case []any:
			if s.index == nil {
				return fmt.Errorf("key on a list")
			}
			j := *s.index
			if j < 0 {
				j += len(c)
			}
			if j < 0 || j >= len(c) {
				return fmt.Errorf("index out of range in %s", path)
			}
			if last {
				if remove {
					return vm.setList(id, steps[:i], append(append([]any(nil), c[:j]...), c[j+1:]...))
				}
				c[j] = v
				return nil
			}
			cur = c[j]
		default:
			return fmt.Errorf("cannot go into a %T at %s", cur, path)
		}
	}
	return nil
}

func (vm *VM) setList(id string, steps []step, l []any) error {
	return vm.put(id, joinPath(steps), l, false)
}

func joinPath(steps []step) string {
	var sb strings.Builder
	for i, s := range steps {
		if s.index != nil {
			fmt.Fprintf(&sb, "[%d]", *s.index)
			continue
		}
		if i > 0 {
			sb.WriteByte('.')
		}
		sb.WriteString(s.key)
	}
	return sb.String()
}

func (vm *VM) storeTo(spec []string, v int32) error {
	// <id> <path> <type> <scale>
	scale, err := strconv.ParseFloat(spec[3], 64)
	if err != nil {
		return err
	}
	x := float64(v) * scale
	var val any
	switch spec[2] {
	case "int", "short", "byte", "long":
		val = int32(math.Floor(x))
	default:
		val = x
	}
	return vm.put(spec[0], spec[1], val, false)
}

// data runs `data get|modify|remove storage ...`.
func (vm *VM) data(w []string) (result, error) {
	if len(w) < 3 || w[1] != "storage" {
		return result{}, fmt.Errorf("only data on storage is supported")
	}
	id := w[2]
	switch w[0] {
	case "get":
		path := ""
		if len(w) > 3 {
			path = w[3]
		}
		v, err := vm.get(id, path)
		if err != nil {
			return result{}, err
		}
		scale := 1.0
		if len(w) > 4 {
			if scale, err = strconv.ParseFloat(w[4], 64); err != nil {
				return result{}, err
			}
		}
		switch x := v.(type) {
		case int32:
			return result{value: int32(math.Floor(float64(x) * scale)), ok: true}, nil
		case float64:
			return result{value: int32(math.Floor(x * scale)), ok: true}, nil
		case string:
			return result{value: int32(len(x)), ok: true}, nil
		case []any:
			return result{value: int32(len(x)), ok: true}, nil
		case map[string]any:
			return result{value: int32(len(x)), ok: true}, nil
		}
		return result{}, fmt.Errorf("cannot get %T", v)
	case "remove":
		return result{value: 1, ok: true}, vm.put(id, w[3], nil, true)
	case "modify":
		// data modify storage <id> <path> <set|append|prepend|merge> (value <snbt> | from storage <id> <path> | string storage <id> <path>)
		if len(w) < 6 {
			return result{}, fmt.Errorf("incomplete data modify")
		}
		path, mode := w[3], w[4]
		var v any
		switch w[5] {
		case "value":
			var err error
			if v, err = ParseSNBT(strings.Join(w[6:], " ")); err != nil {
				return result{}, err
			}
		case "from":
			src, err := vm.get(w[7], w[8])
			if err != nil {
				return result{}, err
			}
			v = clone(src)
		case "string":
			src, err := vm.get(w[7], w[8])
			if err != nil {
				return result{}, err
			}
			v = toText(src)
		default:
			return result{}, fmt.Errorf("unsupported data source %q", w[5])
		}
		switch mode {
		case "set":
			if old, err := vm.get(id, path); err == nil && SNBT(old) == SNBT(v) {
				return result{failed: true}, nil // nothing changed: the command fails
			}
			return result{value: 1, ok: true}, vm.put(id, path, v, false)
		case "append", "prepend":
			cur, err := vm.get(id, path)
			if err != nil {
				cur = []any{}
			}
			l, ok := cur.([]any)
			if !ok {
				return result{}, fmt.Errorf("%s is not a list", path)
			}
			if mode == "append" {
				l = append(l, v)
			} else {
				l = append([]any{v}, l...)
			}
			return result{value: int32(len(l)), ok: true}, vm.put(id, path, l, false)
		case "merge":
			cur, err := vm.get(id, path)
			m, ok := cur.(map[string]any)
			if err != nil || !ok {
				m = map[string]any{}
			}
			add, ok := v.(map[string]any)
			if !ok {
				return result{}, fmt.Errorf("merge needs a compound")
			}
			for k, x := range add {
				m[k] = x
			}
			return result{value: 1, ok: true}, vm.put(id, path, m, false)
		}
		return result{}, fmt.Errorf("unsupported data modify %q", mode)
	}
	return result{}, fmt.Errorf("unsupported data %q", w[0])
}

func clone(v any) any {
	switch x := v.(type) {
	case map[string]any:
		m := make(map[string]any, len(x))
		for k, y := range x {
			m[k] = clone(y)
		}
		return m
	case []any:
		l := make([]any, len(x))
		for i, y := range x {
			l[i] = clone(y)
		}
		return l
	}
	return v
}

func toText(v any) string {
	switch x := v.(type) {
	case string:
		return x
	case int32:
		return strconv.Itoa(int(x))
	case float64:
		return strconv.FormatFloat(x, 'f', -1, 64)
	}
	return SNBT(v)
}

// SNBT writes a value as SNBT, with sorted keys.
func SNBT(v any) string {
	switch x := v.(type) {
	case map[string]any:
		keys := make([]string, 0, len(x))
		for k := range x {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		parts := make([]string, len(keys))
		for i, k := range keys {
			parts[i] = k + ":" + SNBT(x[k])
		}
		return "{" + strings.Join(parts, ",") + "}"
	case []any:
		parts := make([]string, len(x))
		for i, y := range x {
			parts[i] = SNBT(y)
		}
		return "[" + strings.Join(parts, ",") + "]"
	case string:
		return strconv.Quote(x)
	case int32:
		return strconv.Itoa(int(x))
	case float64:
		return strconv.FormatFloat(x, 'f', -1, 64) + "d"
	}
	return fmt.Sprint(v)
}

// snbtMacro is how a macro argument is substituted: strings bare,
// numbers as written.
func snbtMacro(v any) string {
	if s, ok := v.(string); ok {
		return s
	}
	if f, ok := v.(float64); ok {
		return strconv.FormatFloat(f, 'f', -1, 64)
	}
	return SNBT(v)
}

// ParseSNBT reads the SNBT subset the backend writes.
func ParseSNBT(s string) (any, error) {
	p := &snbt{s: strings.TrimSpace(s)}
	v, err := p.value()
	if err != nil {
		return nil, err
	}
	p.ws()
	if p.i != len(p.s) {
		return nil, fmt.Errorf("unexpected %q in SNBT", p.s[p.i:])
	}
	return v, nil
}

type snbt struct {
	s string
	i int
}

func (p *snbt) ws() {
	for p.i < len(p.s) && p.s[p.i] == ' ' {
		p.i++
	}
}

func (p *snbt) value() (any, error) {
	p.ws()
	if p.i >= len(p.s) {
		return nil, fmt.Errorf("unexpected end of SNBT")
	}
	switch c := p.s[p.i]; {
	case c == '{':
		p.i++
		m := map[string]any{}
		for {
			p.ws()
			if p.i < len(p.s) && p.s[p.i] == '}' {
				p.i++
				return m, nil
			}
			k, err := p.word()
			if err != nil {
				return nil, err
			}
			p.ws()
			if p.i >= len(p.s) || p.s[p.i] != ':' {
				return nil, fmt.Errorf("expected ':' in SNBT")
			}
			p.i++
			v, err := p.value()
			if err != nil {
				return nil, err
			}
			m[k] = v
			p.ws()
			if p.i < len(p.s) && p.s[p.i] == ',' {
				p.i++
			}
		}
	case c == '[':
		p.i++
		var l []any
		for {
			p.ws()
			if p.i < len(p.s) && p.s[p.i] == ']' {
				p.i++
				if l == nil {
					l = []any{}
				}
				return l, nil
			}
			v, err := p.value()
			if err != nil {
				return nil, err
			}
			l = append(l, v)
			p.ws()
			if p.i < len(p.s) && p.s[p.i] == ',' {
				p.i++
			}
		}
	case c == '"' || c == '\'':
		return p.word()
	}
	w, err := p.word()
	if err != nil {
		return nil, err
	}
	switch w {
	case "true":
		return int32(1), nil
	case "false":
		return int32(0), nil
	}
	num := strings.TrimRight(w, "bslBSL")
	if n, err := strconv.ParseInt(num, 10, 64); err == nil {
		return int32(n), nil
	}
	if f, err := strconv.ParseFloat(strings.TrimRight(w, "fdFD"), 64); err == nil {
		return f, nil
	}
	return w, nil
}

func (p *snbt) word() (string, error) {
	if p.i < len(p.s) && (p.s[p.i] == '"' || p.s[p.i] == '\'') {
		q := p.s[p.i]
		var sb strings.Builder
		for j := p.i + 1; j < len(p.s); j++ {
			switch p.s[j] {
			case '\\':
				j++
				if j < len(p.s) {
					sb.WriteByte(p.s[j])
				}
			case q:
				p.i = j + 1
				return sb.String(), nil
			default:
				sb.WriteByte(p.s[j])
			}
		}
		return "", fmt.Errorf("unclosed string in SNBT")
	}
	start := p.i
	for p.i < len(p.s) && strings.IndexByte(" ,:]}", p.s[p.i]) < 0 {
		p.i++
	}
	if p.i == start {
		return "", fmt.Errorf("expected a value in SNBT")
	}
	return p.s[start:p.i], nil
}
