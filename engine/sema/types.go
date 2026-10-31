package sema

import (
	"sort"
	"strings"

	"github.com/minecraftmetascript/mms/engine/spec"
)

// tview is a type seen from a position: the type, the template bindings in
// scope, and the path of the last reference followed to reach it (which
// names it for messages, sugar rules and callables).
type tview struct {
	T    *spec.Type
	Env  *env
	Name string
	// RefID is set when a reference with #[id] led here, as in
	// `#[id] LootContextParamSets`: the enum's values are then ids.
	RefID bool
}

type env struct {
	params map[string]tview
	parent *env
}

func (e *env) lookup(path string) (tview, bool) {
	for ; e != nil; e = e.parent {
		if v, ok := e.params[path]; ok {
			return v, true
		}
	}
	return tview{}, false
}

// Short returns a readable name for the type.
func (v tview) Short() string {
	if v.Name != "" {
		return spec.ShortPath(v.Name)
	}
	if v.T == nil {
		return "any"
	}
	return spec.Format(v.T, 0)
}

// norm follows references, applies templates and resolves static
// dispatches until it reaches a structural type. An unresolvable type
// becomes nil, which the checker treats as `any`.
func (c *Checker) norm(v tview) tview {
	for i := 0; i < 64 && v.T != nil; i++ {
		t := v.T
		switch t.Kind {
		case "reference":
			if b, ok := v.Env.lookup(t.Path); ok {
				v = b
				continue
			}
			_, hasID := t.Attr("id")
			v = tview{T: c.Pack.Types[t.Path], Env: v.Env, Name: t.Path, RefID: v.RefID || hasID}
		case "concrete":
			// Follow references only: norm would strip the template
			// before its parameters are bound.
			child := tview{T: t.Child, Env: v.Env}
			for j := 0; j < 64 && child.T != nil && child.T.Kind == "reference"; j++ {
				if b, ok := child.Env.lookup(child.T.Path); ok {
					child = b
					continue
				}
				child = tview{T: c.Pack.Types[child.T.Path], Env: child.Env, Name: child.T.Path}
			}
			if child.T == nil || child.T.Kind != "template" {
				v = child // type arguments on a non-template are ignored
				continue
			}
			e := &env{params: map[string]tview{}, parent: v.Env}
			for j, p := range child.T.TypeParams {
				if j < len(t.TypeArgs) {
					e.params[p.Path] = tview{T: t.TypeArgs[j], Env: v.Env}
				}
			}
			v = tview{T: child.T.Child, Env: e, Name: child.Name}
		case "template":
			v = tview{T: t.Child, Env: v.Env, Name: v.Name}
		case "dispatcher":
			if key, ok := staticIndex(t.ParallelIndices); ok {
				cs, _ := c.Pack.Dispatch(t.Registry, key)
				v = tview{T: cs, Env: v.Env}
				continue
			}
			return v
		case "indexed":
			child := t.Child
			if child != nil && child.Kind == "dispatcher" {
				if key, ok := staticIndex(t.ParallelIndices); ok {
					cs, _ := c.Pack.Dispatch(child.Registry, key)
					v = tview{T: cs, Env: v.Env}
					continue
				}
			}
			return tview{Env: v.Env, Name: v.Name}
		default:
			return v
		}
	}
	return v
}

func staticIndex(idx []*spec.Index) (string, bool) {
	if len(idx) == 1 && idx[0].Kind == "static" {
		return idx[0].Value, true
	}
	return "", false
}

// field is a struct field after spreads are flattened.
type field struct {
	Key      string
	KeyType  *tview // set for map-like fields: `[#[id="block"] string]: T`
	Type     tview
	Optional bool
	Desc     string
}

// dispInfo describes a struct whose remaining fields depend on the value
// of one key field, like `type` in `{ type: ..., ...dispatcher[[type]] }`.
type dispInfo struct {
	KeyField   string
	Dispatcher string
	// KeyRegistry is the registry the key field's ids come from, which can
	// list keys that have no dispatcher case (types without fields).
	KeyRegistry string
	// NoCase is set when the selected key has no dispatcher case: the
	// mcdoc does not describe its fields, so extra fields are allowed.
	NoCase bool
}

// fields flattens a struct's fields. caseKey picks the dispatcher case for
// a dispatched struct; with "" only the common fields are returned.
func (c *Checker) fields(v tview, caseKey string) ([]field, *dispInfo) {
	var primary string
	return c.fieldsWith(v, func(keyField string) string {
		if primary == "" {
			primary = keyField
		}
		if keyField == primary {
			return caseKey
		}
		return ""
	})
}

// fieldsWith flattens a struct's fields; keyOf gives the case for each
// dispatcher spread by its key field (a struct can have several, such as
// `type` and `formula`). The first dispatcher found is the primary one.
func (c *Checker) fieldsWith(v tview, keyOf func(keyField string) string) ([]field, *dispInfo) {
	var out []field
	var disp *dispInfo
	var walk func(v tview, depth int)
	walk = func(v tview, depth int) {
		v = c.norm(v)
		if v.T != nil && v.T.Kind == "union" && len(v.T.Members) == 1 {
			walk(tview{T: v.T.Members[0], Env: v.Env}, depth+1)
			return
		}
		if v.T == nil || v.T.Kind != "struct" || depth > 8 {
			return
		}
		for _, f := range v.T.Fields {
			if f.Kind == "spread" {
				st := f.Type
				// `...dispatcher[[type]]<T>`: dispatch, then apply T to
				// the case.
				var targs []*spec.Type
				if st != nil && st.Kind == "concrete" && st.Child != nil && st.Child.Kind == "dispatcher" {
					targs = st.TypeArgs
					st = st.Child
				}
				if st != nil && st.Kind == "dispatcher" {
					if acc, ok := dynamicKey(st.ParallelIndices); ok {
						if disp == nil {
							disp = &dispInfo{KeyField: acc, Dispatcher: st.Registry}
						}
						if caseKey := keyOf(acc); caseKey != "" {
							cs, ok := c.Pack.Dispatch(st.Registry, caseKey)
							if !ok && disp.KeyField == acc {
								disp.NoCase = true
							}
							if ok {
								if targs != nil {
									cs = &spec.Type{Kind: "concrete", Child: cs, TypeArgs: targs}
								}
								walk(tview{T: cs, Env: v.Env}, depth+1)
							}
						}
						continue
					}
				}
				walk(tview{T: st, Env: v.Env}, depth+1)
				continue
			}
			fd := field{Type: tview{T: f.Type, Env: v.Env}, Optional: f.Optional, Desc: f.Desc, Key: f.Key}
			if f.KeyType != nil {
				fd.KeyType = &tview{T: f.KeyType, Env: v.Env}
			}
			out = append(out, fd)
		}
	}
	walk(v, 0)
	if disp != nil {
		for _, f := range out {
			if f.Key == disp.KeyField {
				disp.KeyRegistry = c.idRegistry(f.Type)
			}
		}
	}
	return out, disp
}

func dynamicKey(idx []*spec.Index) (string, bool) {
	if len(idx) != 1 || idx[0].Kind != "dynamic" || len(idx[0].Accessor) != 1 {
		return "", false
	}
	s := strings.Trim(string(idx[0].Accessor[0]), `"`)
	if strings.ContainsAny(s, "{}") {
		return "", false
	}
	return s, true
}

// idInfo is what a string type's #[id] attribute says.
type idInfo struct {
	Registry string // "" for any id
	Tags     string // "", "allowed", "implicit" or "required"
	Present  bool
}

func idOf(t *spec.Type) idInfo {
	a, ok := t.Attr("id")
	if !ok {
		return idInfo{}
	}
	info := idInfo{Present: true}
	if s, ok := a.LiteralString(); ok {
		info.Registry = s
		return info
	}
	if a.Value != nil && a.Value.Kind == "tree" {
		if r := a.Value.Tree["registry"]; r != nil && r.Literal != nil {
			info.Registry = r.Literal.String()
		}
		if tg := a.Value.Tree["tags"]; tg != nil && tg.Literal != nil {
			info.Tags = tg.Literal.String()
		}
	}
	return info
}

// idRegistry finds the id registry of a string type, looking through
// single-member unions.
func (c *Checker) idRegistry(v tview) string {
	v = c.norm(v)
	if v.T == nil {
		return ""
	}
	switch v.T.Kind {
	case "string":
		return c.registryName(idOf(v.T).Registry)
	case "union":
		for _, m := range v.T.Members {
			if r := c.idRegistry(tview{T: m, Env: v.Env}); r != "" {
				return r
			}
		}
	}
	return ""
}

// registryName maps an mcdoc registry name onto the pack's registries;
// mcdoc writes some worldgen registries without their prefix.
func (c *Checker) registryName(r string) string {
	if r == "" {
		return ""
	}
	if _, ok := c.Pack.Registries[r]; ok {
		return r
	}
	if _, ok := c.Pack.Registries["worldgen/"+r]; ok {
		return "worldgen/" + r
	}
	return r
}

// dispatchKeys lists the case keys of a dispatched struct: dispatcher cases
// plus key-registry ids, which covers field-less types the mcdoc omits.
func (c *Checker) dispatchKeys(d *dispInfo) []string {
	set := map[string]bool{}
	for _, k := range c.Pack.DispatchKeys(d.Dispatcher) {
		set[k] = true
	}
	for _, id := range c.Pack.Registries[d.KeyRegistry] {
		set[id] = true
	}
	keys := make([]string, 0, len(set))
	for k := range set {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (c *Checker) hasDispatchKey(d *dispInfo, key string) bool {
	if _, ok := c.Pack.Dispatch(d.Dispatcher, key); ok {
		return true
	}
	return d.KeyRegistry != "" && c.Pack.HasID(d.KeyRegistry, key)
}

// mapFieldFor picks the map-like field whose key type accepts key, falling
// back to the first map field.
func (c *Checker) mapFieldFor(fs []field, key string) *field {
	var first *field
	for i := range fs {
		f := &fs[i]
		if f.KeyType == nil {
			continue
		}
		if first == nil {
			first = f
		}
		if c.keyAccepts(*f.KeyType, key) {
			return f
		}
	}
	return first
}

func (c *Checker) keyAccepts(kt tview, key string) bool {
	kt = c.norm(kt)
	if kt.T == nil {
		return true
	}
	switch kt.T.Kind {
	case "string":
		id := idOf(kt.T)
		if !id.Present || id.Registry == "" {
			return !strings.HasPrefix(key, "!") || id.Present
		}
		if prefix, ok := idPrefix(kt.T); ok {
			if !strings.HasPrefix(key, prefix) {
				return false
			}
			key = key[len(prefix):]
		} else if strings.HasPrefix(key, "!") {
			return false
		}
		reg := c.registryName(id.Registry)
		if _, known := c.Pack.Registries[reg]; !known {
			return true
		}
		return c.Pack.HasID(reg, key)
	case "enum":
		for _, v := range kt.T.Values {
			if v.String() == key || "minecraft:"+v.String() == key {
				return true
			}
		}
		return false
	case "literal":
		return kt.T.Literal != nil && kt.T.Literal.String() == key
	case "union":
		for _, m := range kt.T.Members {
			if c.keyAccepts(tview{T: m, Env: kt.Env}, key) {
				return true
			}
		}
		return false
	}
	return true
}

func idPrefix(t *spec.Type) (string, bool) {
	a, ok := t.Attr("id")
	if !ok || a.Value == nil || a.Value.Kind != "tree" {
		return "", false
	}
	if p := a.Value.Tree["prefix"]; p != nil && p.Literal != nil {
		return p.Literal.String(), true
	}
	return "", false
}

// mapValueType resolves a map field's value type for one key: values that
// dispatch on the key itself (`dispatcher[[%key]]`) pick the key's case.
func (c *Checker) mapValueType(f *field, key string) tview {
	t := f.Type.T
	if t != nil && t.Kind == "dispatcher" && len(t.ParallelIndices) == 1 && t.ParallelIndices[0].Kind == "dynamic" {
		acc := t.ParallelIndices[0].Accessor
		if len(acc) == 1 && strings.Contains(string(acc[0]), `"key"`) {
			k := strings.TrimPrefix(strings.TrimPrefix(key, "!"), "minecraft:")
			if cs, ok := c.Pack.Dispatch(t.Registry, k); ok {
				return tview{T: cs, Env: f.Type.Env}
			}
			return tview{Env: f.Type.Env}
		}
	}
	return f.Type
}
