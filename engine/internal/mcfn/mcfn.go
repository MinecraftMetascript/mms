// Package mcfn runs .mcfunction files for tests. It implements the
// commands the function backend emits (scoreboard players, execute with
// score conditions and store, function, return, data on storage) with the
// game's semantics: 32-bit wrapping arithmetic, floor division, `return`
// leaving the current function, and stores that only happen when a
// command produces a result. Anything else fails loudly, so a test never
// passes by ignoring a command.
package mcfn

import (
	"fmt"
	"strconv"
	"strings"
)

// VM holds a world's scores, storage and functions.
type VM struct {
	Functions map[string][]string // id -> lines
	Scores    map[string]map[string]int32
	Storage   map[string]map[string]any // storage id -> top-level compound
	Said      []string                  // `say` and `tellraw` output
	Steps     int                       // commands run
	MaxSteps  int
	depth     int
}

func New(files map[string][]string) *VM {
	return &VM{Functions: files, Scores: map[string]map[string]int32{}, Storage: map[string]map[string]any{}, MaxSteps: 1_000_000}
}

// Score returns a score and whether it is set.
func (vm *VM) Score(holder, obj string) (int32, bool) {
	v, ok := vm.Scores[obj][holder]
	return v, ok
}

func (vm *VM) set(holder, obj string, v int32) error {
	m, ok := vm.Scores[obj]
	if !ok {
		return fmt.Errorf("unknown objective %q", obj)
	}
	m[holder] = v
	return nil
}

// result is what a command produced.
type result struct {
	value    int32
	ok       bool // the command produced a result
	failed   bool // the command failed, as when data modify changes nothing
	returned bool // a return ran: leave the function
}

// Run runs a function and returns its return value, if it returned.
func (vm *VM) Run(id string) (int32, bool, error) {
	r, err := vm.call(id, nil)
	return r.value, r.ok, err
}

func (vm *VM) call(id string, macro map[string]any) (result, error) {
	lines, ok := vm.Functions[id]
	if !ok {
		return result{}, fmt.Errorf("unknown function %s", id)
	}
	vm.depth++
	defer func() { vm.depth-- }()
	if vm.depth > 10000 {
		return result{}, fmt.Errorf("function calls nested too deeply in %s", id)
	}
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if strings.HasPrefix(line, "$") {
			var err error
			if line, err = substitute(line[1:], macro); err != nil {
				return result{}, fmt.Errorf("%s:%d: %w", id, i+1, err)
			}
		}
		r, err := vm.exec(line)
		if err != nil {
			return result{}, fmt.Errorf("%s:%d: %s: %w", id, i+1, line, err)
		}
		if r.returned {
			return result{value: r.value, ok: r.ok}, nil
		}
	}
	return result{}, nil
}

func substitute(line string, macro map[string]any) (string, error) {
	var sb strings.Builder
	for {
		i := strings.Index(line, "$(")
		if i < 0 {
			sb.WriteString(line)
			return sb.String(), nil
		}
		j := strings.IndexByte(line[i:], ')')
		if j < 0 {
			return "", fmt.Errorf("unclosed macro argument")
		}
		key := line[i+2 : i+j]
		v, ok := macro[key]
		if !ok {
			return "", fmt.Errorf("missing macro argument %q", key)
		}
		sb.WriteString(line[:i])
		sb.WriteString(snbtMacro(v))
		line = line[i+j+1:]
	}
}

// Exec runs one command outside any function.
func (vm *VM) Exec(line string) error {
	_, err := vm.exec(line)
	return err
}

func (vm *VM) exec(line string) (result, error) {
	vm.Steps++
	if vm.Steps > vm.MaxSteps {
		return result{}, fmt.Errorf("more than %d commands: an endless loop?", vm.MaxSteps)
	}
	w := fields(line)
	if len(w) == 0 {
		return result{}, nil
	}
	switch w[0] {
	case "say":
		vm.Said = append(vm.Said, strings.Join(w[1:], " "))
		return result{value: 1, ok: true}, nil
	case "tellraw":
		vm.Said = append(vm.Said, strings.Join(w[2:], " "))
		return result{value: 1, ok: true}, nil
	case "scoreboard":
		return vm.scoreboard(w)
	case "execute":
		return vm.execute(w[1:])
	case "function":
		return vm.function(w[1:])
	case "return":
		if len(w) < 2 {
			return result{}, fmt.Errorf("return needs a value")
		}
		if w[1] == "run" {
			r, err := vm.exec(strings.Join(w[2:], " "))
			if err != nil {
				return result{}, err
			}
			if !r.ok {
				return result{}, fmt.Errorf("return run: the command produced no result or failed")
			}
			return result{value: r.value, ok: true, returned: true}, nil
		}
		if w[1] == "fail" {
			return result{returned: true}, nil
		}
		v, err := atoi(w[1])
		return result{value: v, ok: true, returned: true}, err
	case "data":
		return vm.data(w[1:])
	}
	return result{}, fmt.Errorf("the test interpreter does not support %q", w[0])
}

// fields splits a command into words, keeping bracketed and quoted
// parts together.
func fields(line string) []string {
	var out []string
	depth := 0
	var quote rune
	start := -1
	for i, r := range line {
		switch {
		case quote != 0:
			if r == quote {
				quote = 0
			}
		case r == '"' || r == '\'':
			quote = r
		case r == '{' || r == '[':
			depth++
		case r == '}' || r == ']':
			depth--
		case r == ' ' && depth == 0:
			if start >= 0 {
				out = append(out, line[start:i])
				start = -1
			}
			continue
		}
		if start < 0 {
			start = i
		}
	}
	if start >= 0 {
		out = append(out, line[start:])
	}
	return out
}

func atoi(s string) (int32, error) {
	v, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("expected an integer, got %q", s)
	}
	return int32(v), nil
}

func (vm *VM) function(w []string) (result, error) {
	if len(w) == 0 {
		return result{}, fmt.Errorf("function needs an id")
	}
	var macro map[string]any
	if len(w) >= 4 && w[1] == "with" && w[2] == "storage" {
		v, err := vm.get(w[3], strings.Join(w[4:], " "))
		if err != nil {
			return result{}, err
		}
		m, ok := v.(map[string]any)
		if !ok {
			return result{}, fmt.Errorf("macro arguments must be a compound")
		}
		macro = m
	} else if len(w) > 1 {
		return result{}, fmt.Errorf("unsupported function arguments %v", w[1:])
	}
	r, err := vm.call(w[0], macro)
	r.returned = false // a return inside the callee only leaves the callee
	return r, err
}

func (vm *VM) scoreboard(w []string) (result, error) {
	if len(w) >= 5 && w[1] == "objectives" && w[2] == "add" {
		if _, ok := vm.Scores[w[3]]; !ok {
			vm.Scores[w[3]] = map[string]int32{}
		}
		return result{value: 1, ok: true}, nil
	}
	if len(w) < 4 || w[1] != "players" {
		return result{}, fmt.Errorf("unsupported scoreboard command")
	}
	holder, obj := w[3], ""
	if len(w) > 4 {
		obj = w[4]
	}
	switch w[2] {
	case "set", "add", "remove":
		if len(w) != 6 {
			return result{}, fmt.Errorf("scoreboard players %s needs a value", w[2])
		}
		v, err := atoi(w[5])
		if err != nil {
			return result{}, err
		}
		cur, _ := vm.Score(holder, obj)
		switch w[2] {
		case "add":
			v = cur + v
		case "remove":
			v = cur - v
		}
		return result{value: v, ok: true}, vm.set(holder, obj, v)
	case "get":
		v, ok := vm.Score(holder, obj)
		if !ok {
			return result{failed: true}, nil // the game fails the command
		}
		return result{value: v, ok: true}, nil
	case "reset":
		delete(vm.Scores[obj], holder)
		return result{value: 1, ok: true}, nil
	case "operation":
		if len(w) != 8 {
			return result{}, fmt.Errorf("scoreboard players operation needs 5 arguments")
		}
		a, _ := vm.Score(holder, obj)
		b, ok := vm.Score(w[6], w[7])
		if !ok {
			return result{}, fmt.Errorf("%s has no score on %s", w[6], w[7])
		}
		var v int32
		switch w[5] {
		case "=":
			v = b
		case "+=":
			v = a + b
		case "-=":
			v = a - b
		case "*=":
			v = a * b
		case "/=":
			v = a
			if b != 0 {
				v = floorDiv(a, b)
			}
		case "%=":
			v = a
			if b != 0 {
				v = a - floorDiv(a, b)*b
			}
		case "<":
			v = min(a, b)
		case ">":
			v = max(a, b)
		case "><":
			if err := vm.set(w[6], w[7], a); err != nil {
				return result{}, err
			}
			v = b
		default:
			return result{}, fmt.Errorf("unknown operation %q", w[5])
		}
		return result{value: v, ok: true}, vm.set(holder, obj, v)
	}
	return result{}, fmt.Errorf("unsupported scoreboard players %s", w[2])
}

func floorDiv(a, b int32) int32 {
	if a == -1<<31 && b == -1 {
		return a
	}
	q := a / b
	if a%b != 0 && (a < 0) != (b < 0) {
		q--
	}
	return q
}

// execute runs an execute chain. A failed condition ends the chain with
// no result.
func (vm *VM) execute(w []string) (result, error) {
	type store struct {
		holder, obj string
		success     bool
		storage     []string
	}
	var stores []store
	for i := 0; i < len(w); {
		switch w[i] {
		case "run":
			r, err := vm.exec(strings.Join(w[i+1:], " "))
			if err != nil {
				return result{}, err
			}
			if r.ok || r.failed {
				for _, s := range stores {
					v := r.value
					if s.success {
						v = 1
					}
					if r.failed {
						v = 0
					}
					var err error
					if s.storage != nil {
						err = vm.storeTo(s.storage, v)
					} else {
						err = vm.set(s.holder, s.obj, v)
					}
					if err != nil {
						return result{}, err
					}
				}
			}
			return r, nil
		case "if", "unless":
			ok, n, err := vm.condition(w[i+1:])
			if err != nil {
				return result{}, err
			}
			if ok == (w[i] == "unless") {
				return result{}, nil
			}
			i += 1 + n
		case "store":
			if i+2 >= len(w) {
				return result{}, fmt.Errorf("incomplete store")
			}
			s := store{success: w[i+1] == "success"}
			switch w[i+2] {
			case "score":
				s.holder, s.obj = w[i+3], w[i+4]
				i += 5
			case "storage":
				// store result storage <id> <path> <type> <scale>
				s.storage = w[i+3 : i+7]
				i += 7
			default:
				return result{}, fmt.Errorf("unsupported store target %q", w[i+2])
			}
			stores = append(stores, s)
		default:
			return result{}, fmt.Errorf("unsupported execute subcommand %q", w[i])
		}
	}
	return result{}, fmt.Errorf("execute without run")
}

// condition evaluates `score ...` or `data storage ...` and returns how
// many words it used.
func (vm *VM) condition(w []string) (bool, int, error) {
	if len(w) >= 4 && w[0] == "score" {
		a, aok := vm.Score(w[1], w[2])
		if w[3] == "matches" {
			lo, hi, err := parseRange(w[4])
			if err != nil {
				return false, 0, err
			}
			return aok && a >= lo && a <= hi, 5, nil
		}
		b, bok := vm.Score(w[4], w[5])
		if !aok || !bok {
			return false, 6, nil
		}
		var ok bool
		switch w[3] {
		case "<":
			ok = a < b
		case "<=":
			ok = a <= b
		case "=":
			ok = a == b
		case ">=":
			ok = a >= b
		case ">":
			ok = a > b
		default:
			return false, 0, fmt.Errorf("unknown comparison %q", w[3])
		}
		return ok, 6, nil
	}
	if len(w) >= 4 && w[0] == "data" && w[1] == "storage" {
		_, err := vm.get(w[2], w[3])
		return err == nil, 4, nil
	}
	return false, 0, fmt.Errorf("unsupported condition %v", w)
}

func parseRange(s string) (int32, int32, error) {
	lo, hi := int32(-1<<31), int32(1<<31-1)
	a, b, isRange := strings.Cut(s, "..")
	var err error
	if !isRange {
		lo, err = atoi(s)
		return lo, lo, err
	}
	if a != "" {
		if lo, err = atoi(a); err != nil {
			return 0, 0, err
		}
	}
	if b != "" {
		if hi, err = atoi(b); err != nil {
			return 0, 0, err
		}
	}
	return lo, hi, nil
}
