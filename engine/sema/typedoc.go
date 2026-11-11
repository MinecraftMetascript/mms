package sema

import (
	"strconv"
	"strings"
	"sync"

	"github.com/minecraftmetascript/mms/engine/spec"
)

// This file describes the schema's named types, such as ClimateParameter
// or IntProvider, the same way everywhere: in the reference's type
// entries, in hovers and in completion. Everything comes from mcdoc.

// TypeAnchor is the id of a type's entry on its reference page.
func TypeAnchor(id string) string { return "type-" + nonWord.ReplaceAllString(id, "-") }

// TypeURL is where the reference explains the named type at path.
func TypeURL(path string) string {
	return ReferenceURL + "/types/" + spec.TypeArea(path) + "#" + TypeAnchor(typeID(path))
}

// typeLink is the reference link of the named type at path.
func typeLink(path string) *RefLink {
	return &RefLink{Type: typeID(path), Area: spec.TypeArea(path)}
}

// SchemaTypePath is the path of the schema type MMS code names name, or
// "" when there is none.
func (c *Checker) SchemaTypePath(name string) string {
	return c.Pack.ShortNames()[name]
}

// nameableTypes lists the schema types that MMS code can name and that the
// reference gives an entry of their own: every name but a type parameter
// (T, V) and a declaration kind, which has a page of its own.
func nameableTypes(p *spec.Pack, ks *Kinds) []string {
	var out []string
	for name, path := range p.ShortNames() {
		if len(name) < 2 || strings.HasPrefix(name, "<") {
			continue
		}
		if _, ok := ks.Lookup(name); ok {
			continue
		}
		out = append(out, path)
	}
	return out
}

// HasTypeEntry reports whether the reference has an entry of its own for
// the schema type at path.
func (c *Checker) HasTypeEntry(path string) bool {
	name := spec.ShortPath(path)
	if len(name) < 2 || c.Pack.ShortNames()[name] != path {
		return false
	}
	if c.Kinds != nil {
		if _, ok := c.Kinds.Lookup(name); ok {
			return false
		}
	}
	return true
}

// DescribeType is the reference entry of the named schema type at path,
// worked out on its own for a hover or a completion item; nil when path
// names no type.
func (c *Checker) DescribeType(path string) *RefType {
	if c.typeAt(path) == nil {
		return nil
	}
	rt := &refTypes{c: c, types: map[string]*RefType{}, helper: func(group, name string) bool {
		if name == "" {
			return len(c.Pack.DispatchKeys(group)) > 0
		}
		return false
	}}
	e := rt.entry(path)
	rt.defineOne(path)
	return e
}

// TypeMarkdown renders a type entry for a hover: its shape, examples of
// values that fit, its fields and a link to the reference.
func TypeMarkdown(e *RefType, path string) string {
	if e == nil {
		return ""
	}
	var sb strings.Builder
	sb.WriteString(TypeSummary(e) + ".")
	writeExamples(&sb, e.Examples)
	if len(e.Values) > 0 {
		sb.WriteString("\n\n**Values**\n")
		for i, v := range e.Values {
			if i == 24 {
				sb.WriteString("\n- …")
				break
			}
			sb.WriteString("\n- `" + v.Value + "`")
			if v.Doc != "" {
				sb.WriteString(": " + oneLine(v.Doc))
			}
		}
	}
	if len(e.Fields) > 0 {
		sb.WriteString("\n\n**Fields**\n")
		for i, f := range e.Fields {
			if i == 24 {
				sb.WriteString("\n- …")
				break
			}
			opt := ""
			if f.Optional {
				opt = "?"
			}
			sb.WriteString("\n- `" + f.Name + opt + ": " + f.Type + "`")
			if f.Doc != "" {
				sb.WriteString(": " + oneLine(f.Doc))
			}
		}
	}
	if len(e.Includes) > 0 {
		var names []string
		for _, p := range e.Includes {
			names = append(names, "`"+p.Text+"`")
		}
		sb.WriteString("\n\nAlso has the fields of " + strings.Join(names, ", ") + ".")
	}
	if path != "" {
		sb.WriteString("\n\n[Reference](" + TypeURL(path) + ")")
	}
	return sb.String()
}

// typeNote describes the named type a value or field takes, for its
// hover: what it is, examples that fit and a link. "" for an unnamed
// type with no examples.
func (c *Checker) typeNote(v tview) string {
	n := c.norm(v)
	if n.Name != "" && c.HasTypeEntry(n.Name) {
		if e := c.DescribeType(n.Name); e != nil {
			summary := TypeSummary(e)
			if len(e.Values) == 0 && len(e.Fields) == 0 {
				// Written with the type arguments the value has, such as
				// IntProvider's number range.
				rt := &refTypes{c: c, types: map[string]*RefType{}, helper: func(string, string) bool { return true }}
				summary = plainParts(rt.format(n, 0))
			}
			var sb strings.Builder
			sb.WriteString("`" + e.Name + "` is " + summary + ".")
			writeExamples(&sb, c.examples(n))
			sb.WriteString(" [Reference](" + TypeURL(n.Name) + ")")
			return sb.String()
		}
	}
	var sb strings.Builder
	writeExamples(&sb, c.examples(n))
	return strings.TrimSpace(sb.String())
}

func writeExamples(sb *strings.Builder, xs []string) {
	if len(xs) == 0 {
		return
	}
	sb.WriteString(" For example ")
	for i, x := range xs {
		if i > 0 {
			if i == len(xs)-1 {
				sb.WriteString(" or ")
			} else {
				sb.WriteString(", ")
			}
		}
		sb.WriteString("`" + x + "`")
	}
	sb.WriteString(".")
}

// TypeSummary is a type entry in one line, such as
// "number -2..2 | list of number -2..2 (exactly 2)".
func TypeSummary(e *RefType) string {
	if e == nil {
		return ""
	}
	if len(e.Values) > 0 {
		var vs []string
		for _, v := range e.Values {
			vs = append(vs, v.Value)
		}
		if len(vs) > 8 {
			vs = append(vs[:8], "…")
		}
		return "one of " + strings.Join(vs, " | ")
	}
	if len(e.Fields) > 0 {
		var fs []string
		for _, f := range e.Fields {
			if !f.Optional {
				fs = append(fs, f.Name)
			}
		}
		if len(fs) == 0 {
			return "object with optional fields"
		}
		if len(fs) > 6 {
			fs = append(fs[:6], "…")
		}
		return "object { " + strings.Join(fs, ", ") + " }"
	}
	if len(e.Def) == 1 && e.Def[0].Link != nil && e.Def[0].Link.Group != nil && e.Def[0].Link.Helper == "" {
		return "made by one of the " + e.Def[0].Text + " helpers"
	}
	return plainParts(e.Def)
}

func oneLine(s string) string {
	s = strings.TrimSpace(s)
	if i := strings.Index(s, "\n\n"); i >= 0 {
		s = s[:i]
	}
	return strings.ReplaceAll(s, "\n", " ")
}

// valueArm is one way to write a value of a type that no helper, id or
// enum value already covers: a number, a list, a record or a string. A
// union has one per member.
type valueArm struct {
	Label   string // as completion lists it, such as [-2, 2]
	Snippet string // what completion inserts
	Detail  string // the arm's type in words
}

// valueArms lists the plain ways to write a value of v.
func (c *Checker) valueArms(v tview, depth int) []valueArm {
	v = c.norm(v)
	if v.T == nil || depth > 6 {
		return nil
	}
	rt := &refTypes{c: c, types: map[string]*RefType{}, helper: func(string, string) bool { return true }}
	detail := func() string { return plainParts(rt.format(v, 0)) }
	switch v.T.Kind {
	case "union":
		var out []valueArm
		seen := map[string]bool{}
		for _, m := range v.T.Members {
			for _, a := range c.valueArms(tview{T: m, Env: v.Env}, depth+1) {
				if !seen[a.Label] {
					seen[a.Label] = true
					out = append(out, a)
				}
			}
		}
		return out
	case "byte", "short", "int", "long", "float", "double":
		n := sampleNumber(v.T)
		return []valueArm{{Label: n, Snippet: "${1:" + n + "}", Detail: detail()}}
	case "list", "byte_array", "int_array", "long_array":
		item := v.T.Item
		if item == nil {
			item = &spec.Type{Kind: "int"}
		}
		if n, ok := exactLength(v.T.LengthRange); ok && n >= 1 && n <= 4 {
			iv := c.norm(tview{T: item, Env: v.Env})
			if iv.T != nil && isNumberKind(iv.T.Kind) {
				samples := numberRow(iv.T, n)
				var snip []string
				for i, s := range samples {
					snip = append(snip, "${"+strconv.Itoa(i+1)+":"+s+"}")
				}
				return []valueArm{{Label: "[" + strings.Join(samples, ", ") + "]", Snippet: "[" + strings.Join(snip, ", ") + "]", Detail: detail()}}
			}
		}
		return []valueArm{{Label: "[…]", Snippet: "[$1]", Detail: detail()}}
	case "tuple":
		var label, snip []string
		for i, it := range v.T.Items {
			s := "…"
			if iv := c.norm(tview{T: it, Env: v.Env}); iv.T != nil && isNumberKind(iv.T.Kind) {
				s = sampleNumber(iv.T)
			}
			label = append(label, s)
			snip = append(snip, "${"+strconv.Itoa(i+1)+":"+strings.TrimSuffix(s, "…")+"}")
		}
		return []valueArm{{Label: "[" + strings.Join(label, ", ") + "]", Snippet: "[" + strings.Join(snip, ", ") + "]", Detail: detail()}}
	case "struct":
		fs, disp := c.fields(v, "")
		if disp != nil {
			return nil // its helpers are listed
		}
		if len(fs) == 1 && fs[0].KeyType != nil {
			return []valueArm{{Label: "{ … }", Snippet: "{\n\t$1\n}", Detail: detail()}}
		}
		what := "record"
		if v.Name != "" {
			what = spec.ShortPath(v.Name) + " record"
		}
		var names, lines []string
		for _, f := range fs {
			if f.Optional || f.KeyType != nil {
				continue
			}
			names = append(names, f.Key)
			lines = append(lines, "\t"+f.Key+" = $"+strconv.Itoa(len(names)))
		}
		if len(names) == 0 {
			return []valueArm{{Label: "{ … }", Snippet: "{\n\t$1\n}", Detail: what}}
		}
		label := names
		if len(label) > 4 {
			label = append(label[:4:4], "…")
		}
		return []valueArm{{Label: "{ " + strings.Join(label, ", ") + " }", Snippet: "{\n" + strings.Join(lines, "\n") + "\n}", Detail: what}}
	case "string":
		if idOf(v.T).Present || isSelectorType(v.T) {
			return nil
		}
		return []valueArm{{Label: `"…"`, Snippet: `"$1"`, Detail: detail()}}
	}
	return nil
}

func isNumberKind(k string) bool {
	switch k {
	case "byte", "short", "int", "long", "float", "double":
		return true
	}
	return false
}

func exactLength(r *spec.Range) (int, bool) {
	if r == nil || r.Min == nil || r.Max == nil || *r.Min != *r.Max {
		return 0, false
	}
	return int(*r.Min), true
}

func formatNumber(x float64, integer bool) string {
	if integer {
		return strconv.FormatInt(int64(x), 10)
	}
	return strconv.FormatFloat(x, 'f', -1, 64)
}

func isInteger(k string) bool { return k == "byte" || k == "short" || k == "int" || k == "long" }

// sampleNumber is a value in t's range: 0 when it fits, else the bound
// nearest to it.
func sampleNumber(t *spec.Type) string {
	x := 0.0
	if r := t.ValueRange; r != nil {
		if r.Min != nil && x < *r.Min {
			x = *r.Min
		}
		if r.Max != nil && x > *r.Max {
			x = *r.Max
		}
		// An exclusive bound: step inside it.
		if r.Min != nil && x == *r.Min && r.Kind&1 != 0 || r.Max != nil && x == *r.Max && r.Kind&2 != 0 {
			switch {
			case r.Min != nil && r.Max != nil && !isInteger(t.Kind):
				x = (*r.Min + *r.Max) / 2
			case x == *r.Min:
				x++
			default:
				x--
			}
		}
	}
	return formatNumber(x, isInteger(t.Kind))
}

// numberRow is n sample values of t for a fixed-length list: a range's
// bounds for two, so `[min, max]` reads as what it is.
func numberRow(t *spec.Type, n int) []string {
	out := make([]string, n)
	for i := range out {
		out[i] = sampleNumber(t)
	}
	if n == 2 && t.ValueRange != nil && t.ValueRange.Min != nil && t.ValueRange.Max != nil {
		out[0] = formatNumber(*t.ValueRange.Min, isInteger(t.Kind))
		out[1] = formatNumber(*t.ValueRange.Max, isInteger(t.Kind))
	}
	return out
}

// examples are MMS values that fit v, one for each way of writing it.
func (c *Checker) examples(v tview) []string {
	v = c.norm(v)
	if v.T == nil {
		return nil
	}
	members := []*spec.Type{v.T}
	if v.T.Kind == "union" {
		members = v.T.Members
	}
	var out []string
	seen := map[string]bool{}
	for _, m := range members {
		x := c.sample(tview{T: m, Env: v.Env}, 0)
		if x == "" || seen[x] {
			continue
		}
		seen[x] = true
		out = append(out, x)
		if len(out) == 4 {
			break
		}
	}
	return out
}

// sample writes one value of v in MMS, or "" when none comes to mind.
func (c *Checker) sample(v tview, depth int) string {
	v = c.norm(v)
	if v.T == nil || depth > 3 {
		return ""
	}
	switch v.T.Kind {
	case "union":
		for _, m := range v.T.Members {
			if s := c.sample(tview{T: m, Env: v.Env}, depth+1); s != "" {
				return s
			}
		}
	case "byte", "short", "int", "long", "float", "double":
		return sampleNumber(v.T)
	case "boolean":
		return "true"
	case "literal":
		if v.T.Literal != nil {
			return v.T.Literal.String()
		}
	case "enum":
		if len(v.T.Values) > 0 {
			return v.T.Values[0].String()
		}
	case "string":
		if id := idOf(v.T); id.Registry != "" {
			if id.Tags == "required" {
				return "#minecraft:" + strings.TrimPrefix(readableRegistry(id.Registry), "tag/")
			}
			if ids := c.Pack.Registries[c.registryName(id.Registry)]; len(ids) > 0 {
				return ids[0]
			}
			return ""
		}
		if isSelectorType(v.T) {
			return "@s"
		}
		return `"text"`
	case "list", "byte_array", "int_array", "long_array":
		item := v.T.Item
		if item == nil {
			item = &spec.Type{Kind: "int"}
		}
		iv := c.norm(tview{T: item, Env: v.Env})
		if n, ok := exactLength(v.T.LengthRange); ok && n >= 1 && n <= 4 && iv.T != nil && isNumberKind(iv.T.Kind) {
			return "[" + strings.Join(numberRow(iv.T, n), ", ") + "]"
		}
		if s := c.sample(iv, depth+1); s != "" {
			return "[" + s + "]"
		}
		return ""
	case "tuple":
		var xs []string
		for _, it := range v.T.Items {
			s := c.sample(tview{T: it, Env: v.Env}, depth+1)
			if s == "" {
				return ""
			}
			xs = append(xs, s)
		}
		return "[" + strings.Join(xs, ", ") + "]"
	case "struct":
		fs, disp := c.fields(v, "")
		if names := c.callableNames(v); len(names) > 0 && disp != nil {
			if depth > 0 {
				return names[0] + "(…)"
			}
			name := names[0]
			if cl := c.callableByName(name, v); cl != nil && len(c.requiredSlots(cl)) == 0 {
				return name + "()"
			}
			return name + "(…)"
		}
		var parts []string
		for _, f := range fs {
			if f.Optional || f.KeyType != nil {
				continue
			}
			if depth > 0 {
				return "{ … }" // only the outer record is spelled out
			}
			s := c.sample(f.Type, depth+1)
			if s == "" {
				s = "…"
			}
			parts = append(parts, f.Key+" = "+s)
			if len(parts) == 4 {
				parts = append(parts, "…")
				break
			}
		}
		if len(parts) == 0 {
			return "{}"
		}
		return "{ " + strings.Join(parts, ", ") + " }"
	}
	return ""
}

// spreadDispatchers lists the dispatchers v's fields spread in, directly or
// through spread structs.
func (c *Checker) spreadDispatchers(v tview) []string {
	var out []string
	seen := map[string]bool{}
	var walk func(v tview, depth int)
	walk = func(v tview, depth int) {
		v = c.norm(v)
		if v.T == nil || v.T.Kind != "struct" || depth > 8 {
			return
		}
		for _, f := range v.T.Fields {
			if f.Kind != "spread" || f.Type == nil {
				continue
			}
			st := f.Type
			if st.Kind == "concrete" && st.Child != nil {
				st = st.Child
			}
			if st.Kind == "dispatcher" {
				if !seen[st.Registry] {
					seen[st.Registry] = true
					out = append(out, st.Registry)
				}
				continue
			}
			walk(tview{T: st, Env: v.Env}, depth+1)
		}
	}
	walk(v, 0)
	return out
}

// TypeBrief is what completion shows of a schema type's name.
type TypeBrief struct {
	Summary string // its shape in one line
	Doc     string // Markdown, as its hover shows it
	// Used is set when a data or resource pack file can hold it, so it
	// ranks before types only the game's own data uses.
	Used bool
}

var (
	briefMu sync.Mutex
	briefs  = map[*spec.Pack]map[string]TypeBrief{}
)

// TypeBriefs describes every schema type name, worked out once per pack.
func TypeBriefs(p *spec.Pack, ks *Kinds) map[string]TypeBrief {
	briefMu.Lock()
	defer briefMu.Unlock()
	if m, ok := briefs[p]; ok {
		return m
	}
	c := NewChecker(p, ks, &Table{}, "minecraft")
	used := usedTypes(p)
	m := map[string]TypeBrief{}
	for name, path := range p.ShortNames() {
		e := c.DescribeType(path)
		if e == nil {
			continue
		}
		link := ""
		if c.HasTypeEntry(path) {
			link = path
		}
		summary := TypeSummary(e)
		if r := []rune(summary); len(r) > 60 {
			summary = string(r[:59]) + "…"
		}
		m[name] = TypeBrief{Summary: summary, Doc: TypeMarkdown(e, link), Used: used[path]}
	}
	briefs[p] = m
	return m
}

// usedTypes are the named types a data or resource pack file can hold,
// found by walking every resource's schema through every dispatcher.
func usedTypes(p *spec.Pack) map[string]bool {
	out := map[string]bool{}
	seenT := map[*spec.Type]bool{}
	seenD := map[string]bool{}
	var walk func(t *spec.Type)
	walk = func(t *spec.Type) {
		if t == nil || seenT[t] {
			return
		}
		seenT[t] = true
		switch t.Kind {
		case "reference":
			if !out[t.Path] {
				out[t.Path] = true
				walk(p.Types[t.Path])
			}
		case "dispatcher":
			if !seenD[t.Registry] {
				seenD[t.Registry] = true
				for _, ct := range p.Dispatchers[t.Registry] {
					walk(ct)
				}
			}
		}
		for _, f := range t.Fields {
			walk(f.Type)
			walk(f.KeyType)
		}
		walk(t.Item)
		walk(t.Child)
		for _, ts := range [][]*spec.Type{t.Members, t.Items, t.TypeArgs} {
			for _, m := range ts {
				walk(m)
			}
		}
	}
	for _, r := range p.Resources {
		walk(r)
	}
	return out
}
