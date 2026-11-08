package sema

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/minecraftmetascript/mms/engine/fn"
	"github.com/minecraftmetascript/mms/engine/internal/near"
	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// Storage: `Storage game: SiegeState` keeps a record of the Type
// SiegeState in command storage, at <ns>:mms game. Functions read and
// write its int, bool and string fields as game.wave; the load function
// sets each field that is missing to its default, so values survive a
// reload and a field added later starts at its default.

// typePrefix is the path the type checker gives a project's Type
// declarations in Table.Types (types.Prefix, which sema cannot import).
const typePrefix = "::mms::"

// reservedStorage are the storage paths the function compiler uses for
// globals (g), string locals (l), macro arguments (m) and the call
// stack (stack), and where a failing Test keeps its values (expect).
var reservedStorage = map[string]bool{"g": true, "l": true, "m": true, "stack": true, ExpectStorage: true}

// storageType returns the record type a Storage declaration names, and
// its Type declaration, or nil.
func (c *Checker) storageType(d *syntax.Decl) (*spec.Type, *syntax.Decl) {
	tn, ok := d.Type.(*syntax.TypeName)
	if !ok || c.Syms == nil {
		return nil, nil
	}
	t := c.Syms.Types[typePrefix+tn.Name]
	if t == nil || t.Kind != "struct" {
		return nil, nil
	}
	return t, c.Syms.TypeDecls[tn.Name]
}

// Storage checks `Storage name: Record` and returns the load-function
// lines that set its missing fields to their defaults.
func (c *Checker) Storage(d *syntax.Decl) []string {
	if d.Name == nil {
		c.errorf("mms/syntax", d.Kind.Sp, "Storage needs a name and a record type: `Storage game: GameState`")
		return nil
	}
	if d.Type == nil {
		c.errorf("mms/missing-value", d.Name.Sp, "Storage needs a record type, a Type with fields: `Storage %s: GameState`", d.Name.Text)
		return nil
	}
	if d.Value != nil || d.Body != nil || d.Params != nil {
		var extra syntax.Node = d.Body
		if d.Value != nil {
			extra = d.Value
		} else if d.Params != nil {
			extra = d.Params
		}
		c.errorf("mms/syntax", extra.Span(), "a Storage starts from its type's defaults; set them in the Type: `wave: int = 0`")
		return nil
	}
	ns, path := c.NS, d.Name.Text
	if i := strings.IndexByte(path, ':'); i >= 0 {
		ns, path = path[:i], path[i+1:]
	}
	if reservedStorage[path] {
		c.errorf("mms/reserved-name", d.Name.Sp, "the compiler keeps its own values in storage %s %s; choose another name", fn.Storage(ns), path)
		return nil
	}
	t, td := c.storageType(d)
	if t == nil {
		c.errorf("mms/type-mismatch", d.Type.Span(), "a Storage holds a record: a Type with fields, such as `Type GameState { wave: int = 0 }`")
		return nil
	}
	loc := fn.Storage(ns) + " " + path
	var decls []*syntax.FieldDecl
	if td != nil && td.Body != nil {
		decls = fieldDecls(td.Body.Items)
	}
	return c.storageDefaults(ns, loc, t, decls, 0)
}

// fieldDecls are the field declarations among a Type body's items.
func fieldDecls(items []syntax.Stmt) []*syntax.FieldDecl {
	var out []*syntax.FieldDecl
	for _, it := range items {
		if f, ok := it.(*syntax.FieldDecl); ok && f.Name != nil {
			out = append(out, f)
		}
	}
	return out
}

// storageDefaults are the load-function lines that set each missing field
// of record t, kept at loc, to its default. A field holding a record
// without a default of its own gets its fields' defaults, so a field
// added to it later starts at its default too.
func (c *Checker) storageDefaults(ns, loc string, t *spec.Type, decls []*syntax.FieldDecl, depth int) []string {
	byName := map[string]*syntax.FieldDecl{}
	for _, d := range decls {
		byName[d.Name.Text] = d
	}
	var lines []string
	for _, f := range t.Fields {
		if f.Key == "" {
			continue
		}
		at := loc + "." + snbtKey(f.Key)
		d := byName[f.Key]
		var v jsonv.Value
		switch {
		case d != nil && d.Default != nil:
			// The default was checked against the field's type with the
			// Type; this only turns it into a value.
			sub := NewChecker(c.Pack, c.Kinds, c.Syms, ns)
			sub.Files, sub.Texts = c.Files, c.Texts
			v = sub.check(d.Default, tview{T: f.Type})
		case c.record(f.Type) != nil && depth < 16:
			var inner []*syntax.FieldDecl
			if d != nil {
				dt := d.Type
				for {
					if o, ok := dt.(*syntax.TypeOptional); ok {
						dt = o.X
					} else if p, ok := dt.(*syntax.TypeParen); ok {
						dt = p.X
					} else {
						break
					}
				}
				switch dt := dt.(type) {
				case *syntax.TypeRecord:
					inner = dt.Fields
				case *syntax.TypeName:
					if c.Syms != nil {
						if nd := c.Syms.TypeDecls[dt.Name]; nd != nil && nd.Body != nil {
							inner = fieldDecls(nd.Body.Items)
						}
					}
				}
			}
			lines = append(lines, c.storageDefaults(ns, at, c.record(f.Type), inner, depth+1)...)
			continue
		default:
			v = zeroOf(f.Type)
		}
		if v == nil {
			continue
		}
		var sb strings.Builder
		writeSNBT(&sb, v)
		lines = append(lines, fmt.Sprintf("execute unless data storage %s run data modify storage %s set value %s", at, at, sb.String()))
	}
	return lines
}

// zeroOf is the value a field without a default starts at: 0, false or
// "" for the types functions use, nil for others.
func zeroOf(t *spec.Type) jsonv.Value {
	switch storageKind(t) {
	case tInt:
		return jsonv.Number("0")
	case tBool:
		return jsonv.Bool(false)
	case tString:
		if cs := storageChoices(t); len(cs) > 0 {
			return jsonv.String(cs[0])
		}
		return jsonv.String("")
	}
	return nil
}

// storageKind is the function type of a field: int, bool or string, or
// "" for a field functions cannot read.
// A union of string literals, `mode: "x" | "y"`, is a string that holds
// one of them.
func storageKind(t *spec.Type) string {
	if t == nil {
		return ""
	}
	switch t.Kind {
	case "int", "byte", "short", "long":
		return tInt
	case "boolean":
		return tBool
	case "string":
		return tString
	case "literal":
		if t.Literal != nil && t.Literal.Kind == "string" {
			return tString
		}
	case "enum":
		if t.EnumKind == "string" {
			return tString
		}
	case "union":
		if storageChoices(t) != nil {
			return tString
		}
	}
	return ""
}

// storageChoices are the strings a field of a union of string literals,
// or a string enum, can hold, in order, or nil for any other field.
func storageChoices(t *spec.Type) []string {
	if t == nil {
		return nil
	}
	var out []string
	switch t.Kind {
	case "literal":
		if t.Literal != nil && t.Literal.Kind == "string" {
			out = append(out, t.Literal.String())
		}
	case "enum":
		if t.EnumKind == "string" {
			for _, v := range t.Values {
				var s string
				if json.Unmarshal(v.Value, &s) == nil {
					out = append(out, s)
				}
			}
		}
	case "union":
		for _, m := range t.Members {
			cs := storageChoices(m)
			if cs == nil {
				return nil
			}
			out = append(out, cs...)
		}
	}
	return out
}

func snbtKey(k string) string {
	if bareSNBTKey.MatchString(k) {
		return k
	}
	return fn.QuoteSNBT(k)
}

// storageSig fingerprints a Storage declaration's fields, which functions
// that use it depend on.
func storageSig(s *Symbol, t *Table) string {
	if s.Decl == nil || s.Decl.Type == nil {
		return ""
	}
	tn, ok := s.Decl.Type.(*syntax.TypeName)
	if !ok {
		return syntax.TypeText(s.Decl.Type)
	}
	rt := t.Types[typePrefix+tn.Name]
	if rt == nil {
		return tn.Name
	}
	return tn.Name + fieldSig(rt, t, 0)
}

// fieldSig fingerprints a record's fields, following nested records and
// listing the strings a choice field allows.
func fieldSig(rt *spec.Type, t *Table, depth int) string {
	var fs []string
	for _, f := range rt.Fields {
		ft := f.Type
		for i := 0; ft != nil && ft.Kind == "reference" && i < 8; i++ {
			ft = t.Types[ft.Path]
		}
		sig := f.Key + ":" + storageKind(f.Type)
		if cs := storageChoices(ft); cs != nil {
			sig += "(" + strings.Join(cs, "|") + ")"
		}
		if ft != nil && ft.Kind == "struct" && depth < 16 {
			sig += fieldSig(ft, t, depth+1)
		}
		fs = append(fs, sig)
	}
	sort.Strings(fs)
	return "{" + strings.Join(fs, ",") + "}"
}

// storageField resolves `game.wave` in a function to the field's storage
// location and function type.
func (f *fc) storageField(e *syntax.Member) (*local, bool) {
	base, names := storageChain(e)
	var id, name string
	switch x := base.(type) {
	case *syntax.Ident:
		if _, ok := f.lookup(x.Name); ok {
			return nil, false
		}
		id, name = f.ns+":"+x.Name, x.Name
	case *syntax.ResLoc:
		id = x.Text
		_, name, _ = strings.Cut(x.Text, ":")
	default:
		return nil, false
	}
	sym := f.c.findSig("Storage", id)
	if sym == nil {
		return nil, false
	}
	in := f.c.info(base)
	in.Symbol = sym
	in.ID = id
	for _, n := range names {
		if n == nil {
			return nil, true
		}
	}
	t, _ := f.c.storageType(sym.Decl)
	if t == nil {
		f.c.errorf("mms/type-mismatch", base.Span(), "'%s' has no record type", name)
		return nil, true
	}
	path, loc := name, fn.Storage(sym.NS)+" "+sym.Path
	for i, n := range names {
		var fl *spec.Field
		for _, x := range t.Fields {
			if x.Key == n.Text {
				fl = x
			}
		}
		if fl == nil {
			var keys []string
			for _, x := range t.Fields {
				keys = append(keys, x.Key)
			}
			f.c.errorf("mms/unknown-name", n.Sp, "%s has no field '%s'%s", path, n.Text, suggestOf(n.Text, keys))
			return nil, true
		}
		path += "." + fl.Key
		loc += "." + snbtKey(fl.Key)
		f.c.info(n).Help = "```mms\n" + path + ": " + describe(tview{T: fl.Type}) + "\n```\n" + DocMarkdown(fl.Desc)
		if i < len(names)-1 {
			if t = f.c.record(fl.Type); t == nil {
				f.c.errorf("mms/type-mismatch", names[i+1].Sp, "%s is %s, not a record with fields", path, describe(tview{T: fl.Type}))
				return nil, true
			}
			continue
		}
		k := storageKind(fl.Type)
		if k == "" {
			f.c.errorf("mms/unsupported", n.Sp, "a function reads int, bool and string fields; %s is another type", path)
			return nil, true
		}
		if k == tString {
			return &local{typ: tString, path: loc, choices: storageChoices(fl.Type)}, true
		}
		return &local{typ: k, nbt: loc}, true
	}
	return nil, true
}

// storageChain splits `game.a.b` into the expression before the first
// dot and the names after it.
func storageChain(e syntax.Expr) (base syntax.Expr, names []*syntax.Name) {
	for {
		m, ok := e.(*syntax.Member)
		if !ok {
			return e, names
		}
		names = append([]*syntax.Name{m.Name}, names...)
		e = m.X
	}
}

// record is t as a record with fields, following a project Type's name,
// or nil when t is not one.
func (c *Checker) record(t *spec.Type) *spec.Type {
	for i := 0; t != nil && t.Kind == "reference" && i < 8; i++ {
		if c.Syms == nil {
			return nil
		}
		t = c.Syms.Types[t.Path]
	}
	if t == nil || t.Kind != "struct" {
		return nil
	}
	for _, f := range t.Fields {
		if f.Key == "" {
			return nil // a map
		}
	}
	return t
}

// suggestOf offers the closest of names, as "; did you mean 'x'?".
func suggestOf(name string, names []string) string {
	best := near.Closest(name, names)
	if best == "" {
		return ""
	}
	return "; did you mean '" + best + "'?"
}

// load reads a storage field into a fresh score.
func (f *fc) load(pre *[]fn.Op, l local) fn.Score {
	t := f.temp()
	*pre = append(*pre, fn.Cmd{Text: fmt.Sprintf("execute store result score %s run data get storage %s", t, l.nbt)})
	return t
}

// save writes an operand to a storage field.
func (f *fc) save(pre *[]fn.Op, l local, op fn.Operand) {
	kind := "int"
	if l.typ == tBool {
		kind = "byte"
	}
	if op.IsConst {
		v := fmt.Sprint(op.Const)
		if l.typ == tBool {
			v = "false"
			if op.Const != 0 {
				v = "true"
			}
		}
		*pre = append(*pre, fn.Cmd{Text: "data modify storage " + l.nbt + " set value " + v})
		return
	}
	*pre = append(*pre, fn.Cmd{Text: fmt.Sprintf("execute store result storage %s %s 1 run scoreboard players get %s", l.nbt, kind, op.S)})
}

// selectorTest lowers exists(target) to a condition and count(target) to
// the number of entities it selects.
func (f *fc) selectorTest(pre *[]fn.Op, name string, e *syntax.Call) *val {
	c := f.c
	if len(e.Args) != 1 || e.Args[0].Name != nil {
		c.errorf(argCountCode(len(e.Args), 1), e.ArgsSp, "%s takes one selector: %s(@e[tag=boss])", name, name)
		return nil
	}
	sel := selectorText(e.Args[0].Value)
	if sel == "" {
		c.errorf("mms/type-mismatch", e.Args[0].Value.Span(), "expected a selector, such as @e[tag=boss]")
		return nil
	}
	if name == "exists" {
		return &val{typ: tBool, cond: &fn.Cond{Parts: []string{"if entity " + sel}}}
	}
	t := f.temp()
	*pre = append(*pre, fn.Cmd{Text: fmt.Sprintf("execute store result score %s if entity %s", t, sel)})
	return &val{typ: tInt, op: fn.Of(t)}
}

// StorageFieldCandidates lists the fields of the Storage that x names in
// function code, for completion after `game.`; nil when x names none.
func (c *Checker) StorageFieldCandidates(x syntax.Expr) []Candidate {
	base, names := storageChain(x)
	sym := c.storageOf(base)
	if sym == nil {
		return nil
	}
	t, _ := c.storageType(sym.Decl)
	// After `game.a.`, the fields of the record a holds.
	for _, n := range names {
		if t == nil || n == nil {
			return nil
		}
		var next *spec.Type
		for _, f := range t.Fields {
			if f.Key == n.Text {
				next = c.record(f.Type)
			}
		}
		t = next
	}
	if t == nil {
		return nil
	}
	var out []Candidate
	for i, f := range t.Fields {
		detail := describe(tview{T: f.Type})
		if storageKind(f.Type) == "" && c.record(f.Type) == nil {
			detail += " (not readable in functions)"
		}
		out = append(out, Candidate{Label: f.Key, Kind: CandField, Detail: detail, Doc: DocMarkdown(f.Desc), Rank: i})
	}
	return out
}

// storageOf returns the Storage a bare name or id in function code names.
func (c *Checker) storageOf(x syntax.Expr) *Symbol {
	switch x := x.(type) {
	case *syntax.Ident:
		return c.find("Storage", c.NS+":"+x.Name)
	case *syntax.ResLoc:
		return c.find("Storage", x.Text)
	}
	return nil
}
