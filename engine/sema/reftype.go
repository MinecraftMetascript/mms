package sema

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/minecraftmetascript/mms/engine/spec"
)

// TypePart is one piece of a readable type, such as "list of " or
// "DensityFunction". Link, when set, is where the piece points to in the
// reference.
type TypePart struct {
	Text string   `json:"text"`
	Link *RefLink `json:"link,omitempty"`
}

// RefLink points at a helper group, a helper, or a type entry.
type RefLink struct {
	// Group is a dispatcher, with Helper "" for the whole group.
	Group  *string `json:"group,omitempty"`
	Helper string  `json:"helper,omitempty"`
	// Type is the ID of a RefType.
	Type string `json:"type,omitempty"`
}

// RefType is a named type that helpers alone do not explain, such as
// DensityFunctionRef (an id or a density function) or an enum.
type RefType struct {
	ID     string     `json:"id"`
	Name   string     `json:"name"`
	Def    []TypePart `json:"def,omitempty"`
	Values []RefValue `json:"values,omitempty"`
	Fields []RefParam `json:"fields,omitempty"`
}

type RefValue struct {
	Value string `json:"value"`
	Doc   string `json:"doc,omitempty"`
}

// refTypes renders types for the reference: readable words instead of
// mcdoc syntax, with named types linked.
type refTypes struct {
	c *Checker
	// helper reports whether the reference lists a helper, or with name
	// "" a helper group.
	helper func(group, name string) bool
	types  map[string]*RefType
	queue  []string
}

func typeID(path string) string {
	return strings.ReplaceAll(strings.TrimPrefix(path, "::java::"), "::", ".")
}

func plainParts(parts []TypePart) string {
	var sb strings.Builder
	for _, p := range parts {
		sb.WriteString(p.Text)
	}
	return sb.String()
}

// readableRegistry drops the worldgen/ prefix: "placed_feature" reads
// better than "worldgen/placed_feature".
func readableRegistry(r string) string {
	return strings.TrimPrefix(strings.TrimPrefix(r, "minecraft:"), "worldgen/")
}

func rangeWords(r *spec.Range) string {
	if r == nil {
		return ""
	}
	return " " + strings.TrimPrefix(spec.Format(&spec.Type{Kind: "int", ValueRange: r}, 0), "int @ ")
}

func lengthWords(r *spec.Range) string {
	if r == nil {
		return ""
	}
	f := func(x float64) string { return strconv.FormatFloat(x, 'f', -1, 64) }
	switch {
	case r.Min != nil && r.Max != nil && *r.Min == *r.Max:
		return " (exactly " + f(*r.Min) + ")"
	case r.Min != nil && r.Max != nil:
		return " (" + f(*r.Min) + " to " + f(*r.Max) + ")"
	case r.Min != nil:
		return " (at least " + f(*r.Min) + ")"
	case r.Max != nil:
		return " (at most " + f(*r.Max) + ")"
	}
	return ""
}

func text(s string) TypePart { return TypePart{Text: s} }

// format renders v. Named types become links; their definitions are
// rendered later into the type entries.
func (rt *refTypes) format(v tview, depth int) []TypePart {
	t := v.T
	if t == nil {
		return []TypePart{text("any")}
	}
	if depth > 8 {
		return []TypePart{text("…")}
	}
	switch t.Kind {
	case "reference":
		if b, ok := v.Env.lookup(t.Path); ok {
			return rt.format(b, depth+1)
		}
		if rt.c.Pack.Types[t.Path] == nil {
			return []TypePart{text(spec.ShortPath(t.Path))} // an unbound type parameter
		}
		return []TypePart{rt.named(t.Path, v.Env)}
	case "string":
		id := idOf(t)
		if id.Registry == "" {
			if id.Present {
				return []TypePart{text("id")}
			}
			return []TypePart{text("string" + lengthWords(t.LengthRange))}
		}
		r := readableRegistry(id.Registry)
		switch id.Tags {
		case "allowed", "implicit":
			return []TypePart{text(r + " id | #" + r + " tag")}
		case "required":
			return []TypePart{text("#" + r + " tag")}
		}
		return []TypePart{text(r + " id")}
	case "boolean":
		return []TypePart{text("bool")}
	case "byte", "short", "int", "long":
		return []TypePart{text("int" + rangeWords(t.ValueRange))}
	case "float", "double":
		return []TypePart{text("number" + rangeWords(t.ValueRange))}
	case "byte_array", "int_array", "long_array":
		return []TypePart{text("list of int" + lengthWords(t.LengthRange))}
	case "any":
		return []TypePart{text("any")}
	case "literal":
		if t.Literal != nil {
			return []TypePart{text(string(t.Literal.Value))}
		}
	case "enum":
		var vs []string
		for _, e := range t.Values {
			vs = append(vs, string(e.Value))
		}
		if len(vs) > 6 {
			vs = append(vs[:6], "…")
		}
		return []TypePart{text(strings.Join(vs, " | "))}
	case "list":
		item := rt.format(tview{T: t.Item, Env: v.Env}, depth+1)
		if len(item) == 1 && item[0].Link == nil && strings.HasSuffix(item[0].Text, " id") && !strings.Contains(item[0].Text, "|") {
			item[0].Text += "s"
		}
		if t.Item != nil && (t.Item.Kind == "union" && len(t.Item.Members) > 1 || t.Item.Kind == "list" && t.Item.LengthRange != nil) {
			item = append(append([]TypePart{text("(")}, item...), text(")"))
		}
		out := append([]TypePart{text("list of ")}, item...)
		return append(out, text(lengthWords(t.LengthRange)))
	case "tuple":
		out := []TypePart{text("[")}
		for i, it := range t.Items {
			if i > 0 {
				out = append(out, text(", "))
			}
			out = append(out, rt.format(tview{T: it, Env: v.Env}, depth+1)...)
		}
		return append(out, text("]"))
	case "union":
		return rt.union(v, depth)
	case "struct":
		_, disp := rt.c.fields(v, "")
		if disp != nil {
			g := disp.Dispatcher
			if !rt.helper(g, "") {
				return []TypePart{text(readableRegistry(g) + " object")}
			}
			return []TypePart{{Text: readableRegistry(g), Link: &RefLink{Group: &g}}}
		}
		// An inline struct such as `{ below_top: int }` that a helper
		// builds: name the helper.
		var out []TypePart
		for _, name := range rt.c.callableNames(v) {
			if rt.helper("", name) {
				if len(out) > 0 {
					out = append(out, text(" | "))
				}
				g := ""
				out = append(out, TypePart{Text: name + "(…)", Link: &RefLink{Group: &g, Helper: name}})
			}
		}
		if out != nil {
			return out
		}
		return rt.inlineStruct(v, depth)
	case "concrete":
		out := rt.format(tview{T: t.Child, Env: v.Env}, depth+1)
		var args []TypePart
		for i, a := range t.TypeArgs {
			if i > 0 {
				args = append(args, text(", "))
			}
			args = append(args, rt.format(tview{T: a, Env: v.Env}, depth+1)...)
		}
		if plain := plainParts(args); plain != "any" && plain != "" {
			out = append(out, text(" of "))
			out = append(out, args...)
		}
		return out
	case "template":
		return rt.format(tview{T: t.Child, Env: v.Env}, depth+1)
	case "dispatcher", "indexed":
		if n := rt.c.norm(v); n.T != nil && n.T != t {
			return rt.format(n, depth+1)
		}
		// A type picked by another field, as in
		// `gradient_density_coordinate[[axis]]`.
		f := spec.Format(t, 0)
		if i, j := strings.Index(f, "[["), strings.LastIndex(f, "]]"); i >= 0 && j > i+2 && !strings.ContainsAny(f[i+2:j], "%{") {
			return []TypePart{text("any (depends on " + f[i+2:j] + ")")}
		}
		return []TypePart{text("any")}
	}
	return []TypePart{text("any")}
}

// union joins members with " | ", dropping repeats and the parentheses
// mcdoc puts around single members.
func (rt *refTypes) union(v tview, depth int) []TypePart {
	ms := v.T.Members
	if len(ms) == 0 {
		return []TypePart{text("nothing")}
	}
	// Ids before lists, so `[block] | block or #tag` reads as
	// "block id | #block tag | list of block id".
	ordered := make([]*spec.Type, 0, len(ms))
	for _, m := range ms {
		if m.Kind != "list" {
			ordered = append(ordered, m)
		}
	}
	for _, m := range ms {
		if m.Kind == "list" {
			ordered = append(ordered, m)
		}
	}
	var out []TypePart
	seen := map[string]bool{}
	for _, m := range ordered {
		p := rt.format(tview{T: m, Env: v.Env}, depth+1)
		key := plainParts(p)
		if seen[key] {
			continue
		}
		seen[key] = true
		if len(out) > 0 {
			out = append(out, text(" | "))
		}
		out = append(out, p...)
	}
	return out
}

// named links a named type to what explains it best: the group of
// helpers that build it, its one helper, or its own type entry.
func (rt *refTypes) named(path string, e *env) TypePart {
	name := spec.ShortPath(path)
	v := rt.c.norm(tview{T: &spec.Type{Kind: "reference", Path: path}, Env: e})
	if v.T != nil && v.T.Kind == "struct" {
		_, disp := rt.c.fields(v, "")
		if disp != nil && rt.helper(disp.Dispatcher, "") {
			g := disp.Dispatcher
			return TypePart{Text: name, Link: &RefLink{Group: &g}}
		}
		if rt.helper("", name) {
			g := ""
			return TypePart{Text: name, Link: &RefLink{Group: &g, Helper: name}}
		}
	}
	id := typeID(path)
	if _, ok := rt.types[id]; !ok {
		rt.types[id] = &RefType{ID: id, Name: name}
		rt.queue = append(rt.queue, path)
	}
	return TypePart{Text: name, Link: &RefLink{Type: id}}
}

// define fills the type entries queued by named, and those their
// definitions reference in turn.
func (rt *refTypes) define() {
	for len(rt.queue) > 0 {
		path := rt.queue[0]
		rt.queue = rt.queue[1:]
		entry := rt.types[typeID(path)]
		t := rt.c.Pack.Types[path]
		if t == nil {
			entry.Def = []TypePart{text("any")}
			continue
		}
		v := tview{T: t}
		n := rt.c.norm(v)
		switch {
		case n.T != nil && n.T.Kind == "enum":
			for _, e := range n.T.Values {
				s := string(e.Value)
				var str string
				if json.Unmarshal(e.Value, &str) == nil {
					s = str
				}
				entry.Values = append(entry.Values, RefValue{Value: s, Doc: e.Desc})
			}
			entry.Def = []TypePart{text("one of the values below")}
			if len(entry.Values) == 0 {
				entry.Def = []TypePart{text("string")}
			}
		case n.T != nil && n.T.Kind == "struct":
			fs, _ := rt.c.fields(n, "")
			entry.Def = []TypePart{text("object")}
			for _, f := range fs {
				if f.KeyType != nil {
					key := plainParts(rt.format(*f.KeyType, 1))
					entry.Fields = append(entry.Fields, RefParam{Name: "[" + key + "]", Optional: true, Doc: f.Desc, TypeParts: rt.format(f.Type, 1)})
					continue
				}
				entry.Fields = append(entry.Fields, RefParam{Name: f.Key, Optional: f.Optional, Doc: f.Desc, TypeParts: rt.format(f.Type, 1)})
			}
		default:
			entry.Def = rt.format(v, 0)
		}
		for i := range entry.Fields {
			entry.Fields[i].Type = plainParts(entry.Fields[i].TypeParts)
		}
	}
}

// inlineStruct renders a small anonymous struct field by field, and a
// map-like one as "map of key to value".
func (rt *refTypes) inlineStruct(v tview, depth int) []TypePart {
	fs, _ := rt.c.fields(v, "")
	if len(fs) == 0 || depth > 4 {
		return []TypePart{text("object")}
	}
	if len(fs) == 1 && fs[0].KeyType != nil {
		out := append([]TypePart{text("map of ")}, rt.format(*fs[0].KeyType, depth+1)...)
		out = append(out, text(" to "))
		return append(out, rt.format(fs[0].Type, depth+1)...)
	}
	if len(fs) > 4 {
		return []TypePart{text("object")}
	}
	out := []TypePart{text("{ ")}
	for i, f := range fs {
		if i > 0 {
			out = append(out, text(", "))
		}
		key := f.Key
		if f.KeyType != nil {
			key = "[" + plainParts(rt.format(*f.KeyType, depth+1)) + "]"
		}
		if f.Optional {
			key += "?"
		}
		out = append(out, text(key+": "))
		out = append(out, rt.format(f.Type, depth+1)...)
	}
	return append(out, text(" }"))
}
