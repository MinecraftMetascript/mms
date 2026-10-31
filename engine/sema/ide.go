package sema

import (
	"fmt"
	"sort"
	"strings"

	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// This file is the checker's API for editor features: completion
// candidates, hover text and signatures, all derived from the expected
// types the checker recorded.

// CandidateKind says what a completion candidate inserts.
type CandidateKind int

const (
	CandField CandidateKind = iota
	CandCallable
	CandMethod
	CandID
	CandEnum
	CandKeyword
	CandKind
)

// Candidate is one completion item.
type Candidate struct {
	Label  string
	Kind   CandidateKind
	Detail string
	Doc    string
	// Deprecated is set for declarations whose doc comment says so.
	Deprecated bool
	// Insert is a snippet ($1 marks the cursor), or "" to insert Label.
	Insert string
	// Rank orders candidates: lower first.
	Rank int
}

// HasType reports whether the checker recorded an expected type.
func (in *Info) HasType() bool { return in != nil && in.typed }

// TypeName is the expected type in a few words.
func (in *Info) TypeName() string {
	if in == nil || !in.typed {
		return ""
	}
	return describe(in.Type)
}

// FieldName returns the field a key, named argument or method names.
func (in *Info) FieldName() (key, doc, typ string, ok bool) {
	if in == nil || in.Field == nil {
		return "", "", "", false
	}
	return in.Field.Key, in.Field.Desc, describe(in.Field.Type), true
}

// ValueCandidates lists what can be written where in expects a value.
func (c *Checker) ValueCandidates(in *Info) []Candidate {
	if in == nil || !in.typed {
		return nil
	}
	return dedupe(c.valueCandidates(in.Type, 0))
}

func (c *Checker) valueCandidates(v tview, depth int) []Candidate {
	v = c.norm(v)
	if v.T == nil || depth > 6 {
		return nil
	}
	var out []Candidate
	switch v.T.Kind {
	case "union":
		for _, m := range v.T.Members {
			out = append(out, c.valueCandidates(tview{T: m, Env: v.Env}, depth+1)...)
		}
	case "struct":
		for _, name := range c.callableNames(tview{T: v.T, Env: v.Env, Name: v.Name}) {
			cand := Candidate{Label: name, Kind: CandCallable, Insert: name + "($1)", Rank: 1}
			if cl := c.callableByName(name, v); cl != nil {
				if cl.Case != "" {
					cand.Detail = "minecraft:" + cl.Case
				}
				cand.Doc = c.signature(cl)
				if len(c.requiredSlots(cl)) == 0 {
					cand.Insert = name + "()"
				}
			} else if name == "If" {
				cand.Insert = "If($1) {\n\t$0\n}"
				cand.Detail = "condition"
			}
			out = append(out, cand)
		}
		for _, s := range sugarCallables[v.Name] {
			if s == "If" {
				out = append(out, Candidate{Label: "If", Kind: CandCallable, Insert: "If($1) {\n\t$0\n}", Detail: "condition", Rank: 0})
			}
		}
	case "string":
		id := idOf(v.T)
		if !id.Present {
			break
		}
		reg := c.registryName(id.Registry)
		if id.Tags != "required" {
			for _, s := range c.Syms.All() {
				if s.Kind.Folder() == reg {
					out = append(out, Candidate{Label: c.shortID(s.ID()), Kind: CandID, Detail: s.Kind.Name, Doc: DocMarkdown(s.Decl.Doc), Deprecated: Deprecated(s.Decl), Rank: 0})
				}
			}
			for _, x := range c.Pack.Registries[reg] {
				out = append(out, Candidate{Label: x, Kind: CandID, Detail: reg, Rank: 2})
			}
		}
		if id.Tags != "" {
			for _, s := range c.Syms.All() {
				if s.Kind.TagOf == reg {
					out = append(out, Candidate{Label: "#" + c.shortID(s.ID()), Kind: CandID, Detail: reg + " tag", Doc: DocMarkdown(s.Decl.Doc), Deprecated: Deprecated(s.Decl), Rank: 1})
				}
			}
			for _, x := range c.Pack.Registries["tag/"+reg] {
				out = append(out, Candidate{Label: "#" + x, Kind: CandID, Detail: reg + " tag", Rank: 3})
			}
		}
	case "enum":
		for _, e := range v.T.Values {
			out = append(out, Candidate{Label: e.String(), Kind: CandEnum, Doc: e.Desc, Rank: 0})
		}
	case "boolean":
		out = append(out, Candidate{Label: "true", Kind: CandKeyword}, Candidate{Label: "false", Kind: CandKeyword})
	case "literal":
		if v.T.Literal != nil {
			out = append(out, Candidate{Label: v.T.Literal.String(), Kind: CandEnum})
		}
	}
	return out
}

func (c *Checker) shortID(id string) string {
	if ns, path := splitID(id); ns == c.NS {
		return path
	}
	return id
}

func dedupe(cs []Candidate) []Candidate {
	seen := map[string]int{}
	var out []Candidate
	for _, c := range cs {
		if i, ok := seen[c.Label]; ok {
			if c.Rank < out[i].Rank {
				out[i] = c
			}
			continue
		}
		seen[c.Label] = len(out)
		out = append(out, c)
	}
	sort.SliceStable(out, func(i, j int) bool {
		if out[i].Rank != out[j].Rank {
			return out[i].Rank < out[j].Rank
		}
		return out[i].Label < out[j].Label
	})
	return out
}

// FieldCandidates lists the fields not yet set in an object body or call
// (in is the Info of the body block or the call). present holds the keys
// already written.
func (c *Checker) FieldCandidates(in *Info, present map[string]bool, call bool) []Candidate {
	if in == nil {
		return nil
	}
	fs := in.Fields
	if in.Disp == nil && fs == nil && in.typed {
		fs, _ = c.fields(in.Type, "")
	}
	var out []Candidate
	if in.Disp != nil && !present[in.Disp.KeyField] && !call {
		out = append(out, Candidate{Label: in.Disp.KeyField, Kind: CandField, Insert: in.Disp.KeyField + " = $1", Detail: "selects the type", Rank: 0})
	}
	for _, f := range fs {
		if present[f.Key] || f.KeyType != nil || (in.Disp != nil && f.Key == in.Disp.KeyField) {
			continue
		}
		rank := 1
		if f.Optional {
			rank = 2
		}
		out = append(out, Candidate{Label: f.Key, Kind: CandField, Insert: f.Key + " = $1", Detail: describe(f.Type), Doc: f.Desc, Rank: rank})
	}
	return dedupe(out)
}

// MethodCandidates lists builder methods for the call whose Info is in.
func (c *Checker) MethodCandidates(in *Info) []Candidate {
	if in == nil || in.Callable == nil {
		return nil
	}
	cl := in.Callable
	var out []Candidate
	for _, f := range cl.Fields {
		if f.KeyType != nil || (cl.Disp != nil && f.Key == cl.Disp.KeyField) {
			continue
		}
		name := Pascal(f.Key)
		ins := name + "($1)"
		if t := c.norm(f.Type); t.T != nil && t.T.Kind == "boolean" {
			ins = name + "()"
		}
		out = append(out, Candidate{Label: name, Kind: CandMethod, Insert: ins, Detail: f.Key + ": " + describe(f.Type), Doc: f.Desc, Rank: 1})
	}
	if cl.Ov != nil {
		for name, m := range cl.Ov.methods {
			ins := name + "($1)"
			if m.val != nil {
				ins = name + "()"
			}
			out = append(out, Candidate{Label: name, Kind: CandMethod, Insert: ins, Detail: "sets " + m.field, Rank: 0})
		}
	}
	return dedupe(out)
}

// ArgInfo returns an Info expecting the type of the positional argument
// at index (counting only positional arguments) of a call.
func (c *Checker) ArgInfo(in *Info, index int, named map[string]bool) *Info {
	if in == nil || in.Callable == nil {
		return nil
	}
	var free []field
	for _, f := range c.positional(in.Callable) {
		if !named[f.Key] {
			free = append(free, f)
		}
	}
	if index >= len(free) {
		if len(free) > 0 && c.acceptsList(free[len(free)-1].Type) {
			return c.listItemInfo(free[len(free)-1])
		}
		return nil
	}
	f := free[index]
	if c.acceptsList(f.Type) {
		return c.listItemInfo(f)
	}
	return &Info{Type: c.norm(f.Type), typed: true, Field: &f}
}

func (c *Checker) listItemInfo(f field) *Info {
	v := c.norm(f.Type)
	var item tview
	switch {
	case v.T != nil && v.T.Kind == "list":
		item = tview{T: v.T.Item, Env: v.Env}
	case v.T != nil && v.T.Kind == "union":
		// A list-or-single field: offer both.
		return &Info{Type: v, typed: true, Field: &f}
	}
	return &Info{Type: c.norm(item), typed: true, Field: &f}
}

// ListItemInfo returns an Info for the items of a list type.
func (c *Checker) ListItemInfo(in *Info) *Info {
	if in == nil || !in.typed {
		return nil
	}
	v := c.norm(in.Type)
	if v.T != nil && v.T.Kind == "union" {
		for _, m := range v.T.Members {
			mv := c.norm(tview{T: m, Env: v.Env})
			if mv.T != nil && mv.T.Kind == "list" {
				v = mv
				break
			}
		}
	}
	if v.T == nil || v.T.Kind != "list" {
		if v.Name == tMaterialRule {
			return &Info{Type: v, typed: true}
		}
		return nil
	}
	return &Info{Type: c.norm(tview{T: v.T.Item, Env: v.Env}), typed: true}
}

func (c *Checker) requiredSlots(cl *Callable) []field {
	var out []field
	for _, f := range c.positional(cl) {
		if f.Optional {
			continue
		}
		if cl.Ov != nil {
			if _, ok := cl.Ov.defaults[f.Key]; ok {
				continue
			}
		}
		out = append(out, f)
	}
	return out
}

// Signature returns a call's label and parameter labels for signature
// help, from the Info of the call.
func (c *Checker) Signature(in *Info) (label string, params []string, docs []string) {
	if in == nil || in.Callable == nil {
		return "", nil, nil
	}
	cl := in.Callable
	for _, f := range c.positional(cl) {
		p := f.Key
		if f.Optional {
			p += "?"
		} else if cl.Ov != nil {
			if _, ok := cl.Ov.defaults[f.Key]; ok {
				p += "?"
			}
		}
		params = append(params, p+": "+describe(f.Type))
		docs = append(docs, f.Desc)
	}
	return cl.Name + "(" + strings.Join(params, ", ") + ")", params, docs
}

func (c *Checker) signature(cl *Callable) string {
	label, _, _ := c.Signature(&Info{Callable: cl})
	return label
}

// Hover returns Markdown for a node's Info, or "".
func (c *Checker) Hover(in *Info) string {
	if in == nil {
		return ""
	}
	var parts []string
	switch {
	case in.Callable != nil:
		cl := in.Callable
		head := "```mms\n" + c.signature(cl) + "\n```"
		if cl.Case != "" {
			head += "\nMakes `minecraft:" + cl.Case + "` (" + describe(cl.Struct) + ")."
		} else {
			head += "\nMakes " + describe(cl.Struct) + "."
		}
		parts = append(parts, head)
	case in.Field != nil:
		parts = append(parts, fmt.Sprintf("```mms\n%s: %s\n```", in.Field.Key, describe(in.Field.Type)))
		if in.Field.Desc != "" {
			parts = append(parts, in.Field.Desc)
		}
	case in.Symbol != nil:
		s := in.Symbol
		parts = append(parts, "```mms\n"+SymbolHead(s.Kind.Name, s.ID(), s.Decl)+"\n```")
		if md := DocMarkdown(s.Decl.Doc); md != "" {
			parts = append(parts, md)
		}
	case in.Local != nil:
		l := in.Local
		head := "let " + l.Name + ": " + l.Type
		switch l.What {
		case "parameter":
			head = "(parameter) " + l.Name + ": " + l.Type
		case "for":
			head = "(loop variable) " + l.Name + ": " + l.Type
		}
		parts = append(parts, "```mms\n"+head+"\n```")
		if md := DocMarkdown(l.Doc); md != "" {
			parts = append(parts, md)
		}
	case in.ID != "":
		parts = append(parts, "`"+in.ID+"`")
		if in.typed {
			parts = append(parts, "Expected: "+describe(in.Type))
		}
	case in.typed:
		parts = append(parts, "Expected: "+describe(in.Type))
	}
	return strings.Join(parts, "\n\n")
}

// SymbolHead is the code line a hover shows for a declaration: its kind
// and id, plus the signature of a Function and the type of a Global.
func SymbolHead(kind, id string, d *syntax.Decl) string {
	head := kind + " " + id
	if d == nil {
		return head
	}
	switch kind {
	case "Function":
		if d.Params != nil {
			sig := SigOf(d)
			var ps []string
			for _, p := range sig.Params {
				ps = append(ps, p.Name+": "+p.Type)
			}
			head += "(" + strings.Join(ps, ", ") + ")"
		}
		if d.Type != nil {
			head += ": " + d.Type.Text
		}
	case "Global":
		if d.Type != nil {
			head += ": " + d.Type.Text
		}
	}
	return head
}

// DocMarkdown renders a declaration's doc comment, tags included, as
// Markdown; "" when there is none.
func DocMarkdown(doc string) string {
	if doc == "" {
		return ""
	}
	return syntax.ParseDoc(doc).Markdown()
}

// Deprecated reports whether a declaration's doc comment has @deprecated.
func Deprecated(d *syntax.Decl) bool {
	return d != nil && d.Doc != "" && syntax.ParseDoc(d.Doc).Deprecated != nil
}

// TypeDoc renders the expected type in mcdoc-like syntax, for hover.
func (in *Info) TypeDoc() string {
	if in == nil || !in.typed || in.Type.T == nil {
		return ""
	}
	return spec.Format(in.Type.T, 1)
}

// ChainBase returns the innermost call of a builder chain such as
// `Noise(x).XzScale(1).`, which is where the checker records the helper.
func ChainBase(e syntax.Expr) *syntax.Call {
	base, _, _, _ := unchain(e)
	return base
}

// MethodArgInfo returns an Info expecting the argument of `.Method(...)` on
// the call whose Info is in.
func (c *Checker) MethodArgInfo(in *Info, method string) *Info {
	if in == nil || in.Callable == nil {
		return nil
	}
	cl := in.Callable
	key := Snake(method)
	if cl.Ov != nil {
		if m, ok := cl.Ov.methods[method]; ok {
			key = m.field
		}
	}
	f := fieldByKey(cl.Fields, key)
	if f == nil {
		return nil
	}
	if c.acceptsList(f.Type) {
		return c.listItemInfo(*f)
	}
	return &Info{Type: c.norm(f.Type), typed: true, Field: f}
}

// DeclKinds lists declaration kinds for completion at the top level.
func (k *Kinds) DeclKinds() []Candidate {
	var out []Candidate
	for _, name := range k.Names() {
		kind, _ := k.Lookup(name)
		cand := Candidate{Label: name, Kind: CandKind, Detail: kind.Resource}
		switch {
		case name == "Namespace":
			cand.Insert, cand.Detail = "Namespace ${1:name}", "sets the namespace for the rest of the file"
		case name == "OnLoad" || name == "OnTick":
			cand.Insert = name + " {\n\t$0\n}"
			cand.Detail = "runs on " + strings.ToLower(strings.TrimPrefix(name, "On"))
		case name == "Function":
			cand.Insert = "Function ${1:name} {\n\t$0\n}"
		case name == "Objective":
			cand.Insert = "Objective ${1:name} = ${2:dummy}"
		case name == "Global":
			cand.Insert = "Global ${1:name}: int = ${2:0}"
		case kind.TagOf != "":
			cand.Insert = name + " ${1:name} = [$0]"
			cand.Detail = kind.TagOf + " tag"
		default:
			cand.Insert = name + " ${1:name} = $0"
		}
		out = append(out, cand)
	}
	return out
}

// CallInfo resolves the helper a call names against the type in expects,
// for chains the checker could not finish (such as `Noise(x).` while the
// method name is still being typed).
func (c *Checker) CallInfo(call *syntax.Call, in *Info) *Info {
	if in == nil || !in.typed {
		return nil
	}
	id, ok := call.Fn.(*syntax.Ident)
	if !ok {
		return nil
	}
	var found *Callable
	var walk func(v tview, depth int)
	walk = func(v tview, depth int) {
		v = c.norm(v)
		if found != nil || v.T == nil || depth > 8 {
			return
		}
		if v.T.Kind == "union" {
			for _, m := range v.T.Members {
				walk(tview{T: m, Env: v.Env}, depth+1)
			}
			return
		}
		found = c.callableByName(id.Name, v)
	}
	walk(in.Type, 0)
	if found == nil {
		return nil
	}
	return &Info{Type: found.Struct, Callable: found, Fields: found.Fields, Disp: found.Disp, typed: true}
}
