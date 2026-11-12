package eval

import (
	"bytes"
	"encoding/json"
	"sort"
	"strings"

	"github.com/minecraftmetascript/mms/engine/syntax"
)

// join is the texts of a list's items with sep between them.
func join(c *ctx, x *syntax.Call, args []argValue) Value {
	p := c.pos(x.Sp)
	vs, ok := knownArgs(c, x, args, "join", 2)
	if !ok {
		return &Str{p, ""}
	}
	sep, ok := text(vs[1])
	if !ok {
		c.errorf("mms/intrinsic-args", args[1].x.Value.Span(), "join's separator is a string, not "+kindName(vs[1]))
		return &Str{p, ""}
	}
	var l *List
	switch v := vs[0].(type) {
	case *List:
		l = v
	case *None:
		return &Str{p, ""}
	default:
		c.errorf("mms/intrinsic-args", args[0].x.Value.Span(), "join takes a list, not "+kindName(vs[0]))
		return &Str{p, ""}
	}
	parts := make([]string, 0, len(l.Elems))
	for _, e := range l.Elems {
		t, ok := text(e)
		if !ok {
			c.errorf("mms/intrinsic-args", args[0].x.Value.Span(), "join takes a list of strings, numbers and names, not one holding "+kindName(e))
			return &Str{p, ""}
		}
		parts = append(parts, t)
	}
	return &Str{p, strings.Join(parts, sep)}
}

// vanillaBlock is how a vanilla block looks:
//
//   - models: the paths of the block models its block state shows, and
//     those its item shows, such as a fence's block/oak_fence_inventory
//   - slots: the texture slots their faces show
//   - other_slots: the other slots their parent chains fill, such as
//     particle
//   - state: its block state file
//   - item: the item definition of its item, when that shows block
//     models; none otherwise, as for a door's flat picture
//
// A block vanilla does not have has no models and no slots.
func vanillaBlock(c *ctx, x *syntax.Call, args []argValue) Value {
	p := c.pos(x.Sp)
	models, slots, other := &List{pos: p}, &List{pos: p}, &List{pos: p}
	out := &Record{pos: p, Keys: []string{"models", "slots", "other_slots", "state", "item"},
		Vals: []Value{models, slots, other, &None{pos: p}, &None{pos: p}}}
	vs, ok := knownArgs(c, x, args, "vanilla_block", 1)
	if !ok {
		return out
	}
	id, ok := text(vs[0])
	if !ok {
		c.errorf("mms/intrinsic-args", args[0].x.Value.Span(), "vanilla_block takes a block id, not "+kindName(vs[0]))
		return out
	}
	path, vanilla := vanillaPath(id)
	if !vanilla || c.e.vanilla == nil {
		return out
	}
	state, ok := c.e.vanilla("blockstates", path)
	if !ok {
		return out
	}
	refs := stateModels(state)
	item, hasItem := c.e.vanilla("items", path)
	if hasItem {
		n := len(refs)
		refs = append(refs, itemModels(item)...)
		if len(refs) == n {
			hasItem = false
		}
	}
	seen := map[string]bool{}
	shown, rest := map[string]bool{}, map[string]bool{}
	for _, m := range refs {
		mp, ok := vanillaPath(m)
		if !ok || seen[mp] || !strings.HasPrefix(mp, "block/") {
			continue
		}
		seen[mp] = true
		if _, ok := c.e.vanilla("models", mp); !ok {
			continue
		}
		models.Elems = append(models.Elems, &Str{p, mp})
		c.e.modelSlots(mp, shown, rest)
	}
	for _, s := range sortedKeys(shown) {
		slots.Elems = append(slots.Elems, &Str{p, s})
	}
	for _, s := range sortedKeys(rest) {
		if !shown[s] {
			other.Elems = append(other.Elems, &Str{p, s})
		}
	}
	out.Vals[3] = jsonValue(p, state)
	if hasItem {
		out.Vals[4] = jsonValue(p, item)
	}
	return out
}

// itemModels lists the block models an item definition shows: the
// "model" of each "minecraft:model" item model in it.
func itemModels(raw []byte) []string {
	var v any
	if json.Unmarshal(raw, &v) != nil {
		return nil
	}
	var out []string
	var walk func(any)
	walk = func(v any) {
		switch v := v.(type) {
		case map[string]any:
			if t, _ := v["type"].(string); t == "minecraft:model" || t == "model" {
				if m, ok := v["model"].(string); ok {
					if p, ok := vanillaPath(m); ok && strings.HasPrefix(p, "block/") {
						out = append(out, m)
					}
				}
			}
			for _, k := range sortedKeys(mapKeys(v)) {
				walk(v[k])
			}
		case []any:
			for _, e := range v {
				walk(e)
			}
		}
	}
	walk(v)
	return out
}

// rebase is a value with every string in it that starts with old
// starting with new instead: rebase(state, "minecraft:block/", "block/oak_log/")
// points a block state's models at the project's own.
func rebase(c *ctx, x *syntax.Call, args []argValue) Value {
	p := c.pos(x.Sp)
	vs, ok := knownArgs(c, x, args, "rebase", 3)
	if !ok {
		return &None{pos: p}
	}
	old, ok1 := text(vs[1])
	nw, ok2 := text(vs[2])
	if !ok1 || !ok2 {
		c.errorf("mms/intrinsic-args", x.Sp, "rebase's old and new starts are strings")
		return vs[0]
	}
	var walk func(Value) Value
	walk = func(v Value) Value {
		switch v := v.(type) {
		case *Str:
			if rest, ok := strings.CutPrefix(v.S, old); ok {
				return &Str{v.pos, nw + rest}
			}
		case *Word:
			if rest, ok := strings.CutPrefix(v.Text, old); ok {
				return &Str{v.pos, nw + rest}
			}
		case *List:
			l := &List{pos: v.pos}
			for _, e := range v.Elems {
				l.Elems = append(l.Elems, walk(e))
			}
			return l
		case *Record:
			r := &Record{pos: v.pos, Keys: v.Keys}
			for _, e := range v.Vals {
				r.Vals = append(r.Vals, walk(e))
			}
			return r
		}
		return v
	}
	return walk(vs[0])
}

func sortedKeys(m map[string]bool) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

// vanillaPath is the path of a minecraft id: "oak_log" or
// "minecraft:oak_log" is oak_log. ok is false for another namespace.
func vanillaPath(id string) (string, bool) {
	if ns, path, ok := strings.Cut(id, ":"); ok {
		return path, ns == "minecraft"
	}
	return id, true
}

// stateModels lists the model ids a block state file names, in order:
// those of each variant, and those each multipart case applies.
func stateModels(raw []byte) []string {
	var st struct {
		Variants  map[string]json.RawMessage `json:"variants"`
		Multipart []struct {
			Apply json.RawMessage `json:"apply"`
		} `json:"multipart"`
	}
	if json.Unmarshal(raw, &st) != nil {
		return nil
	}
	var out []string
	add := func(raw json.RawMessage) {
		var one struct{ Model string }
		var many []struct{ Model string }
		if json.Unmarshal(raw, &many) == nil {
			for _, m := range many {
				out = append(out, m.Model)
			}
		} else if json.Unmarshal(raw, &one) == nil {
			out = append(out, one.Model)
		}
	}
	keys := make([]string, 0, len(st.Variants))
	for k := range st.Variants {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		add(st.Variants[k])
	}
	for _, m := range st.Multipart {
		add(m.Apply)
	}
	return out
}

// modelSlots adds the texture slots of a vanilla model's parent chain
// to shown and rest. A slot is shown when a face of the nearest model in
// the chain with elements leads to it and it is the last slot on the
// way: "#down" is "#end" in cube_column, so end is shown. The other
// slots the chain's textures fill, such as particle and down, go in rest.
func (e *evaluator) modelSlots(path string, shown, rest map[string]bool) {
	seen := map[string]bool{}
	textures := map[string]string{}
	var uses []string
	elements := false
	for path != "" && !seen[path] {
		seen[path] = true
		raw, ok := e.vanilla("models", path)
		if !ok {
			break
		}
		var m struct {
			Parent   string                     `json:"parent"`
			Textures map[string]json.RawMessage `json:"textures"`
			Elements []struct {
				Faces map[string]struct {
					Texture string `json:"texture"`
				} `json:"faces"`
			} `json:"elements"`
		}
		if json.Unmarshal(raw, &m) != nil {
			break
		}
		// Children's textures win over their parents'.
		for k, v := range m.Textures {
			if _, ok := textures[k]; !ok {
				var s string
				json.Unmarshal(v, &s)
				textures[k] = s
			}
		}
		if !elements && m.Elements != nil {
			elements = true
			for _, el := range m.Elements {
				for _, f := range el.Faces {
					if s, ok := strings.CutPrefix(f.Texture, "#"); ok {
						uses = append(uses, s)
					}
				}
			}
		}
		next, ok := vanillaPath(m.Parent)
		if !ok || strings.HasPrefix(next, "builtin/") {
			break
		}
		path = next
	}
	for _, u := range uses {
		for i := 0; i < 32; i++ {
			next, ok := strings.CutPrefix(textures[u], "#")
			if !ok {
				break
			}
			u = next
		}
		shown[u] = true
	}
	for k := range textures {
		rest[k] = true
	}
}

// jsonValue is a JSON file as a value: objects are records, arrays lists.
func jsonValue(p pos, raw []byte) Value {
	d := json.NewDecoder(bytes.NewReader(raw))
	d.UseNumber()
	var v any
	if d.Decode(&v) != nil {
		return &None{pos: p}
	}
	return anyValue(p, v)
}

func anyValue(p pos, v any) Value {
	switch v := v.(type) {
	case map[string]any:
		r := &Record{pos: p}
		for _, k := range sortedKeys(mapKeys(v)) {
			r.Keys = append(r.Keys, k)
			r.Vals = append(r.Vals, anyValue(p, v[k]))
		}
		return r
	case []any:
		l := &List{pos: p}
		for _, e := range v {
			l.Elems = append(l.Elems, anyValue(p, e))
		}
		return l
	case string:
		return &Str{p, v}
	case bool:
		return &Bool{p, v}
	case json.Number:
		if i, err := v.Int64(); err == nil {
			return &Num{p, float64(i), true}
		}
		f, _ := v.Float64()
		return &Num{p, f, false}
	}
	return &None{pos: p}
}

func mapKeys(m map[string]any) map[string]bool {
	out := make(map[string]bool, len(m))
	for k := range m {
		out[k] = true
	}
	return out
}
