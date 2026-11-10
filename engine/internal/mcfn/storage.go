package mcfn

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/MinecraftMetascript/nbt-go"
)

// Storage holds nbt-go tag trees, so values keep their NBT types and
// paths follow the game's rules.

func (vm *VM) root(id string) nbt.Compound {
	m, ok := vm.Storage[id]
	if !ok {
		m = nbt.Compound{}
		vm.Storage[id] = m
	}
	return m
}

// get reads the one tag at path, or the whole storage for "".
func (vm *VM) get(id, path string) (nbt.Tag, error) {
	if path == "" {
		return vm.root(id), nil
	}
	p, err := nbt.ParsePath(path)
	if err != nil {
		return nil, err
	}
	switch ts := p.Get(vm.root(id)); len(ts) {
	case 0:
		return nil, fmt.Errorf("nothing at %s", path)
	case 1:
		return ts[0], nil
	default:
		return nil, fmt.Errorf("%d tags at %s, not one", len(ts), path)
	}
}

// setData stores v at path, creating compounds on the way, and reports
// whether anything changed.
func (vm *VM) setData(id, path string, v nbt.Tag) (bool, error) {
	if path == "" {
		m, ok := v.(nbt.Compound)
		if !ok {
			return false, fmt.Errorf("storage root must be a compound")
		}
		changed := !nbt.Equal(vm.root(id), m)
		vm.Storage[id] = m
		return changed, nil
	}
	p, err := nbt.ParsePath(path)
	if err != nil {
		return false, err
	}
	root := vm.root(id)
	n, err := p.Set(root, v)
	if err != nil {
		return false, err
	}
	if n == 0 && p.Count(root) == 0 {
		return false, fmt.Errorf("nothing at %s", path)
	}
	return n > 0, nil
}

// modify calls f with the tag at path, created with create if it is
// missing, and stores what f returns.
func (vm *VM) modify(id, path string, create func() nbt.Tag, f func(nbt.Tag) (nbt.Tag, error)) error {
	p, err := nbt.ParsePath(path)
	if err != nil {
		return err
	}
	n, err := p.Modify(vm.root(id), create, f)
	if err == nil && n == 0 {
		err = fmt.Errorf("nothing at %s", path)
	}
	return err
}

func (vm *VM) remove(id, path string) error {
	p, err := nbt.ParsePath(path)
	if err != nil {
		return err
	}
	p.Remove(vm.root(id))
	return nil
}

func (vm *VM) storeTo(spec []string, v int32) error {
	// <id> <path> <type> <scale>
	scale, err := strconv.ParseFloat(spec[3], 64)
	if err != nil {
		return err
	}
	x := float64(v) * scale
	var val nbt.Tag
	switch spec[2] {
	case "int":
		val = nbt.Int(math.Floor(x))
	case "short":
		val = nbt.Short(math.Floor(x))
	case "byte":
		val = nbt.Byte(math.Floor(x))
	case "long":
		val = nbt.Long(math.Floor(x))
	case "float":
		val = nbt.Float(x)
	case "double":
		val = nbt.Double(x)
	default:
		return fmt.Errorf("unknown store type %q", spec[2])
	}
	_, err = vm.setData(spec[0], spec[1], val)
	return err
}

// number reads a numeric tag.
func number(t nbt.Tag) (float64, bool) {
	switch x := t.(type) {
	case nbt.Byte:
		return float64(x), true
	case nbt.Short:
		return float64(x), true
	case nbt.Int:
		return float64(x), true
	case nbt.Long:
		return float64(x), true
	case nbt.Float:
		return float64(x), true
	case nbt.Double:
		return float64(x), true
	}
	return 0, false
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
		if n, ok := number(v); ok {
			return result{value: int32(math.Floor(n * scale)), ok: true}, nil
		}
		switch x := v.(type) {
		case nbt.String:
			return result{value: int32(len(x)), ok: true}, nil
		case nbt.List:
			return result{value: int32(len(x)), ok: true}, nil
		case nbt.Compound:
			return result{value: int32(len(x)), ok: true}, nil
		case nbt.ByteArray:
			return result{value: int32(len(x)), ok: true}, nil
		case nbt.IntArray:
			return result{value: int32(len(x)), ok: true}, nil
		case nbt.LongArray:
			return result{value: int32(len(x)), ok: true}, nil
		}
		return result{}, fmt.Errorf("cannot get %v", v.Type())
	case "remove":
		return result{value: 1, ok: true}, vm.remove(id, w[3])
	case "modify":
		// data modify storage <id> <path> <set|append|prepend|merge> (value <snbt> | from storage <id> <path> | string storage <id> <path>)
		if len(w) < 6 {
			return result{}, fmt.Errorf("incomplete data modify")
		}
		path, mode := w[3], w[4]
		var v nbt.Tag
		switch w[5] {
		case "value":
			var err error
			if v, err = nbt.ParseSNBT(strings.Join(w[6:], " ")); err != nil {
				return result{}, err
			}
		case "from":
			src, err := vm.get(w[7], w[8])
			if err != nil {
				return result{failed: true}, nil // nothing to copy: the command fails
			}
			v = nbt.Clone(src)
		case "string":
			src, err := vm.get(w[7], w[8])
			if err != nil {
				return result{}, err
			}
			v = nbt.String(toText(src))
		default:
			return result{}, fmt.Errorf("unsupported data source %q", w[5])
		}
		switch mode {
		case "set":
			changed, err := vm.setData(id, path, v)
			if err == nil && !changed {
				return result{failed: true}, nil // nothing changed: the command fails
			}
			return result{value: 1, ok: true}, err
		case "append", "prepend":
			n := 0
			err := vm.modify(id, path, func() nbt.Tag { return nbt.List{} }, func(cur nbt.Tag) (nbt.Tag, error) {
				l, ok := cur.(nbt.List)
				if !ok {
					return nil, fmt.Errorf("%s is not a list", path)
				}
				if mode == "append" {
					l = append(l, v)
				} else {
					l = append(nbt.List{v}, l...)
				}
				n = len(l)
				return l, nil
			})
			return result{value: int32(n), ok: true}, err
		case "merge":
			add, ok := v.(nbt.Compound)
			if !ok {
				return result{}, fmt.Errorf("merge needs a compound")
			}
			err := vm.modify(id, path, func() nbt.Tag { return nbt.Compound{} }, func(cur nbt.Tag) (nbt.Tag, error) {
				m, ok := cur.(nbt.Compound)
				if !ok {
					m = nbt.Compound{}
				}
				for k, x := range add {
					m[k] = x
				}
				return m, nil
			})
			return result{value: 1, ok: true}, err
		}
		return result{}, fmt.Errorf("unsupported data modify %q", mode)
	}
	return result{}, fmt.Errorf("unsupported data %q", w[0])
}

// toText reads a tag as text, as data modify ... string and macro
// arguments do: a string as itself, a number as written without its
// suffix, anything else as SNBT.
func toText(v nbt.Tag) string {
	switch x := v.(type) {
	case nbt.String:
		return string(x)
	case nbt.Float, nbt.Double:
		n, _ := number(x)
		return strconv.FormatFloat(n, 'f', -1, 64)
	case nbt.Byte, nbt.Short, nbt.Int, nbt.Long:
		n, _ := number(x)
		return strconv.FormatInt(int64(n), 10)
	}
	return v.String()
}

// Get reads a storage value at a path, such as "game.wave".
func (vm *VM) Get(id, path string) (nbt.Tag, error) { return vm.get(id, path) }
