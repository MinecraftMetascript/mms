// Package types resolves the types written in MMS (`int`, `Id<item>`,
// `[T]`, `{string: T}`, `T?`, unions, literal words, schema types such as
// `LootPool`, and the project's own `Type` declarations) into mcdoc
// types, the form the checker already understands, and checks the
// compile-time inputs the evaluator saw against them: Template
// arguments, Blueprint fields and const values. Every problem is
// reported at a span the user wrote.
package types

import (
	"encoding/json"
	"sort"
	"strings"
	"sync"

	"github.com/minecraftmetascript/mms/engine/eval"
	"github.com/minecraftmetascript/mms/engine/sema"
	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// Prefix is the mcdoc path prefix of the project's own types.
const Prefix = "::mms::"

// Env resolves type expressions for one analysis.
type Env struct {
	Pack  *spec.Pack
	Kinds *sema.Kinds
	// Types are the project's Type declarations, resolved, by path.
	Types map[string]*spec.Type
	// Diags are problems in type expressions, by file.
	Diags map[string][]syntax.Diagnostic

	decls map[string]*eval.TypeDef
	busy  map[string]bool
}

// short indexes schema types by the last part of their path, built once
// per pack.
var (
	shortMu    sync.Mutex
	shortIndex = map[*spec.Pack]map[string]string{}
)

func schemaNames(p *spec.Pack) map[string]string {
	shortMu.Lock()
	defer shortMu.Unlock()
	if m, ok := shortIndex[p]; ok {
		return m
	}
	byName := map[string][]string{}
	for path := range p.Types {
		i := strings.LastIndex(path, "::")
		byName[path[i+2:]] = append(byName[path[i+2:]], path)
	}
	m := map[string]string{}
	for name, paths := range byName {
		// The shortest path wins, then the first in order: Text is
		// ::java::util::text::Text, not a nested one.
		sort.Slice(paths, func(i, j int) bool {
			if len(paths[i]) != len(paths[j]) {
				return len(paths[i]) < len(paths[j])
			}
			return paths[i] < paths[j]
		})
		m[name] = paths[0]
	}
	shortIndex[p] = m
	return m
}

// New resolves the project's Type declarations.
func New(p *spec.Pack, kinds *sema.Kinds, r *eval.Result) *Env {
	e := &Env{Pack: p, Kinds: kinds, Types: map[string]*spec.Type{}, Diags: map[string][]syntax.Diagnostic{},
		decls: r.Types, busy: map[string]bool{}}
	names := make([]string, 0, len(r.Types))
	for n := range r.Types {
		names = append(names, n)
	}
	sort.Strings(names)
	for _, n := range names {
		e.define(n)
	}
	return e
}

func (e *Env) errorf(path string, sp syntax.Span, msg string) {
	e.Diags[path] = append(e.Diags[path], syntax.Diagnostic{Code: "mms/unknown-type", Severity: syntax.SeverityError, Message: msg, Sp: sp})
}

// define resolves one Type declaration into Types.
func (e *Env) define(name string) {
	path := Prefix + name
	if _, ok := e.Types[path]; ok || e.busy[name] {
		return
	}
	e.busy[name] = true
	defer delete(e.busy, name)
	td := e.decls[name]
	d := td.D
	var t *spec.Type
	switch {
	case d.TypeValue != nil:
		t, _ = e.Resolve(td.Path, d.TypeValue)
	case d.Body != nil:
		t = &spec.Type{Kind: "struct"}
		for _, it := range d.Body.Items {
			if _, ok := it.(*syntax.AsStmt); ok && d.Extern {
				continue // a built-in form's meaning: see Forms
			}
			f, ok := it.(*syntax.FieldDecl)
			if !ok {
				e.Diags[td.Path] = append(e.Diags[td.Path], syntax.Diagnostic{Code: "mms/type-body", Severity: syntax.SeverityError,
					Message: "a Type body holds fields: name: Type = default", Sp: it.Span()})
				continue
			}
			if f.Name == nil {
				continue
			}
			ft, opt := e.Resolve(td.Path, f.Type)
			t.Fields = append(t.Fields, &spec.Field{Kind: "pair", Key: f.Name.Text, Type: ft, Desc: f.Doc,
				Optional: opt || f.Default != nil})
		}
	}
	if t == nil {
		t = &spec.Type{Kind: "any"}
	}
	e.Types[path] = t
}

// Resolve turns a type expression written in file path into an mcdoc
// type. optional reports a top-level `T?`. Unknown names are reported
// and resolve to any.
func (e *Env) Resolve(path string, t syntax.TypeExpr) (rt *spec.Type, optional bool) {
	switch t := t.(type) {
	case nil:
		return &spec.Type{Kind: "any"}, false
	case *syntax.TypeOptional:
		inner, _ := e.Resolve(path, t.X)
		return inner, true
	case *syntax.TypeParen:
		return e.Resolve(path, t.X)
	case *syntax.TypeList:
		item, _ := e.Resolve(path, t.Elem)
		return &spec.Type{Kind: "list", Item: item}, false
	case *syntax.TypeMap:
		k, _ := e.Resolve(path, t.Key)
		v, _ := e.Resolve(path, t.Value)
		return &spec.Type{Kind: "struct", Fields: []*spec.Field{{Kind: "pair", KeyType: k, Type: v, Optional: true}}}, false
	case *syntax.TypeUnion:
		// A union of words is an enum: `"rigid" | "terrain_matching"`.
		words := true
		for _, a := range t.Arms {
			if _, ok := a.(*syntax.TypeString); !ok {
				words = false
			}
		}
		if words {
			en := &spec.Type{Kind: "enum", EnumKind: "string"}
			for _, a := range t.Arms {
				raw, _ := json.Marshal(a.(*syntax.TypeString).Value)
				en.Values = append(en.Values, &spec.EnumValue{Identifier: a.(*syntax.TypeString).Value, Value: raw})
			}
			return en, false
		}
		u := &spec.Type{Kind: "union"}
		for _, a := range t.Arms {
			m, opt := e.Resolve(path, a)
			optional = optional || opt
			u.Members = append(u.Members, m)
		}
		return u, optional
	case *syntax.TypeString:
		raw, _ := json.Marshal(t.Value)
		return &spec.Type{Kind: "literal", Literal: &spec.Literal{Kind: "string", Value: raw}}, false
	case *syntax.TypeName:
		return e.named(path, t), false
	}
	return &spec.Type{Kind: "any"}, false
}

// Lookup resolves a type expression for an editor, after checking: it
// reports nothing and changes nothing, so it is safe to call while other
// editor requests read the environment. Every Type declaration was
// resolved by New, so the copy only reads e.Types.
func (e *Env) Lookup(path string, t syntax.TypeExpr) *spec.Type {
	q := *e
	q.Diags = map[string][]syntax.Diagnostic{}
	q.busy = map[string]bool{}
	rt, _ := q.Resolve(path, t)
	return rt
}

// primitives are the language's own value types.
var primitives = map[string]string{
	"int": "int", "float": "float", "double": "double", "bool": "boolean", "string": "string", "any": "any",
	"byte": "byte", "short": "short", "long": "long",
}

// opaque are language types whose values the evaluator and the function
// compiler check rather than the schema: functions, commands and asset
// files.
var opaque = map[string]bool{"Function": true, "Commands": true, "Texture": true, "Sound": true, "StructureFile": true, "Selector": true}

func (e *Env) named(path string, t *syntax.TypeName) *spec.Type {
	if k, ok := primitives[t.Name]; ok {
		return &spec.Type{Kind: k}
	}
	if opaque[t.Name] {
		return &spec.Type{Kind: "any"}
	}
	if t.Name == "BookText" {
		return sema.BookTextType()
	}
	if t.Name == "Id" || t.Name == "Ids" {
		if len(t.Args) != 1 {
			e.errorf(path, t.Sp, t.Name+" takes one registry: "+t.Name+"<item>")
			return &spec.Type{Kind: "string"}
		}
		reg := syntax.TypeText(t.Args[0])
		reg = strings.TrimPrefix(reg, "minecraft:")
		if !e.Pack.HasRegistry(reg) && !e.isResourceFolder(reg) {
			e.errorf(path, t.Args[0].Span(), "unknown registry '"+reg+"'")
		}
		if t.Name == "Ids" {
			return idsType(reg)
		}
		return idType(reg)
	}
	if _, ok := e.decls[t.Name]; ok {
		e.define(t.Name)
		return &spec.Type{Kind: "reference", Path: Prefix + t.Name}
	}
	if k, ok := e.Kinds.Lookup(t.Name); ok && k.Resource != "" && !strings.HasPrefix(k.Resource, "tags/") {
		// A kind's name: an id of one, or an inline value where the
		// game allows it.
		return &spec.Type{Kind: "union", Members: []*spec.Type{idType(k.Resource), e.Pack.Resources[k.Resource]}}
	}
	if p, ok := schemaNames(e.Pack)[t.Name]; ok {
		return &spec.Type{Kind: "reference", Path: p}
	}
	e.errorf(path, t.Sp, "unknown type '"+t.Name+"'"+e.suggest(t.Name))
	return &spec.Type{Kind: "any"}
}

func (e *Env) isResourceFolder(reg string) bool {
	_, ok := e.Pack.Resources[reg]
	return ok
}

// suggest offers a close type name.
func (e *Env) suggest(name string) string {
	best, bestD := "", 3
	try := func(c string) {
		if d := distance(strings.ToLower(name), strings.ToLower(c)); d < bestD {
			best, bestD = c, d
		}
	}
	for p := range primitives {
		try(p)
	}
	for n := range e.decls {
		try(n)
	}
	if best == "" {
		return ""
	}
	return "; did you mean '" + best + "'?"
}

func distance(a, b string) int {
	ra, rb := []rune(a), []rune(b)
	prev := make([]int, len(rb)+1)
	for j := range prev {
		prev[j] = j
	}
	for i := 1; i <= len(ra); i++ {
		cur := make([]int, len(rb)+1)
		cur[0] = i
		for j := 1; j <= len(rb); j++ {
			cost := 1
			if ra[i-1] == rb[j-1] {
				cost = 0
			}
			cur[j] = min(prev[j]+1, cur[j-1]+1, prev[j-1]+cost)
		}
		prev = cur
	}
	return prev[len(rb)]
}

// idType is a string that names an id in a registry.
// idsType is Ids<reg>: a list of ids, or one id or tag, the way the
// game writes a set of biomes or blocks.
func idsType(reg string) *spec.Type {
	raw, _ := json.Marshal(reg)
	allowed, _ := json.Marshal("allowed")
	tagged := &spec.Type{Kind: "string", Attributes: []spec.Attribute{{Name: "id", Value: &spec.Type{Kind: "tree", Tree: map[string]*spec.Type{
		"registry": {Kind: "literal", Literal: &spec.Literal{Kind: "string", Value: raw}},
		"tags":     {Kind: "literal", Literal: &spec.Literal{Kind: "string", Value: allowed}},
	}}}}}
	return &spec.Type{Kind: "union", Members: []*spec.Type{{Kind: "list", Item: idType(reg)}, tagged}}
}

func idType(reg string) *spec.Type {
	raw, _ := json.Marshal(reg)
	return &spec.Type{Kind: "string", Attributes: []spec.Attribute{{Name: "id",
		Value: &spec.Type{Kind: "literal", Literal: &spec.Literal{Kind: "string", Value: raw}}}}}
}
