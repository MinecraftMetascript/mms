package sema

import (
	"fmt"
	"math"
	"strings"

	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/spec"
)

// Validate checks a JSON value against a resource kind's schema and returns
// problems as "path: message" strings. The emitter runs it over every file
// so that a lowering bug shows up as an internal error, and tests run it
// over vanilla's own files to check the type model.
func Validate(p *spec.Pack, kinds *Kinds, resource string, v jsonv.Value) []string {
	c := &Checker{Pack: p, Kinds: kinds, Syms: &Table{byKey: map[string]*Symbol{}}, trial: true}
	root := tview{T: p.Resources[resource]}
	if root.T != nil && root.T.Kind == "reference" {
		root.Name = root.T.Path
	}
	var errs []string
	c.validate(v, root, "$", &errs, 0)
	return errs
}

func (c *Checker) validate(val jsonv.Value, v tview, path string, errs *[]string, depth int) {
	if depth > 100 {
		return
	}
	fail := func(format string, args ...any) {
		*errs = append(*errs, path+": "+fmt.Sprintf(format, args...))
	}
	v = c.norm(v)
	if v.T == nil {
		return
	}
	switch v.T.Kind {
	case "union":
		if len(v.T.Members) == 0 {
			return
		}
		best := []string(nil)
		for _, m := range v.T.Members {
			var e []string
			c.validate(val, tview{T: m, Env: v.Env}, path, &e, depth+1)
			if len(e) == 0 {
				return
			}
			if best == nil || len(e) < len(best) {
				best = e
			}
		}
		*errs = append(*errs, best...)
	case "struct":
		o, ok := val.(*jsonv.Object)
		if !ok {
			fail("expected an object (%s)", v.Short())
			return
		}
		keyOf := func(k string) string {
			if s, ok := o.Get(k).(jsonv.String); ok {
				return strings.TrimPrefix(string(s), "minecraft:")
			}
			return ""
		}
		fs, disp := c.fieldsWith(v, func(k string) string {
			if o.Get(k) == nil {
				return noneKey
			}
			return keyOf(k)
		})
		if disp != nil {
			if key := keyOf(disp.KeyField); key != "" && !c.hasDispatchKey(disp, key) {
				fail("unknown %s '%s'", disp.KeyField, key)
				return
			}
		}
		known := map[string]bool{}
		for i := range fs {
			f := &fs[i]
			if f.KeyType != nil {
				continue
			}
			known[f.Key] = true
			fv := o.Get(f.Key)
			if fv == nil {
				if !f.Optional {
					fail("missing field '%s'", f.Key)
				}
				continue
			}
			c.validate(fv, f.Type, path+"."+f.Key, errs, depth+1)
		}
		for _, m := range o.Members {
			if known[m.Key] {
				continue
			}
			if mf := c.mapFieldFor(fs, m.Key); mf != nil {
				c.validate(m.Value, c.mapValueType(mf, m.Key), path+"."+m.Key, errs, depth+1)
				continue
			}
			if disp != nil && disp.NoCase {
				continue
			}
			fail("unknown field '%s'", m.Key)
		}
	case "list", "byte_array", "int_array", "long_array":
		a, ok := val.(jsonv.Array)
		if !ok {
			fail("expected a list")
			return
		}
		item := tview{T: v.T.Item, Env: v.Env}
		if v.T.Kind != "list" {
			item = tview{T: &spec.Type{Kind: strings.TrimSuffix(v.T.Kind, "_array")}}
		}
		if r := v.T.LengthRange; r != nil && !inRange(float64(len(a)), r) {
			fail("list length %d not in%s", len(a), rangeText(r))
		}
		for i, x := range a {
			c.validate(x, item, fmt.Sprintf("%s[%d]", path, i), errs, depth+1)
		}
	case "tuple":
		a, ok := val.(jsonv.Array)
		if !ok || len(a) != len(v.T.Items) {
			fail("expected a list of %d items", len(v.T.Items))
			return
		}
		for i, x := range a {
			c.validate(x, tview{T: v.T.Items[i], Env: v.Env}, fmt.Sprintf("%s[%d]", path, i), errs, depth+1)
		}
	case "string":
		if _, ok := val.(jsonv.String); !ok {
			fail("expected a string")
		}
	case "boolean":
		if _, ok := val.(jsonv.Bool); ok {
			return
		}
		// NBT stores booleans as bytes, so 0 and 1 are accepted too.
		if n, ok := val.(jsonv.Number); ok && (n == "0" || n == "1") {
			return
		}
		fail("expected true or false")
	case "byte", "short", "int", "long", "float", "double":
		n, ok := val.(jsonv.Number)
		if !ok {
			fail("expected a number")
			return
		}
		f, _ := n.Float()
		switch v.T.Kind {
		case "byte", "short", "int", "long":
			if f != math.Trunc(f) {
				fail("expected an integer, got %s", n)
			}
		}
		if !inRange(f, v.T.ValueRange) {
			fail("%s not in%s", n, rangeText(v.T.ValueRange))
		}
	case "literal":
		if v.T.Literal == nil {
			return
		}
		want := v.T.Literal.String()
		got := ""
		switch x := val.(type) {
		case jsonv.String:
			got = string(x)
		case jsonv.Number:
			got = string(x)
		case jsonv.Bool:
			got = fmt.Sprint(bool(x))
		}
		if got != want && strings.TrimPrefix(got, "minecraft:") != strings.TrimPrefix(want, "minecraft:") {
			fail("expected %s, got %s", want, jsonv.Compact(val))
		}
	case "enum":
		got := ""
		switch x := val.(type) {
		case jsonv.String:
			got = strings.TrimPrefix(string(x), "minecraft:")
		case jsonv.Number:
			got = string(x)
		}
		for _, ev := range v.T.Values {
			if strings.TrimPrefix(ev.String(), "minecraft:") == got {
				return
			}
			if f1, ok := (jsonv.Number(ev.String())).Float(); ok {
				if f2, ok := jsonv.Number(got).Float(); ok && f1 == f2 {
					return
				}
			}
		}
		fail("%s is not one of the allowed values", jsonv.Compact(val))
	}
}
