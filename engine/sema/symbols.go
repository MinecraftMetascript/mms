package sema

import (
	"regexp"
	"strings"

	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// Unit is one parsed source file of a project.
type Unit struct {
	Path string
	File *syntax.File
}

// Symbol is a declared resource or MMS entity.
type Symbol struct {
	NS, Path string
	Kind     Kind
	Decl     *syntax.Decl
	Unit     *Unit
	table    *Table
}

// ID returns "ns:path".
func (s *Symbol) ID() string { return s.NS + ":" + s.Path }

// Table holds every declaration of a project, by kind and id.
type Table struct {
	byKey map[string]*Symbol // resource folder + "|" + id
	all   []*Symbol
	// DeclNS maps each declaration to the namespace it lives in.
	DeclNS map[*syntax.Decl]string
	// DeclKind maps each declaration to its resolved kind.
	DeclKind map[*syntax.Decl]Kind
	// Hidden are declarations the prelude's Blueprints emitted, which
	// the editor does not offer by name.
	Hidden map[*syntax.Decl]bool
	// Types are the types the project declares with `Type`, by the
	// path the type checker gives them ("::mms::Name").
	Types map[string]*spec.Type
	// TypeDecls are the project's Type declarations by name, for the
	// defaults of a record's fields.
	TypeDecls map[string]*syntax.Decl
	// Conversions are the `as` conversions of Blueprint instances, by
	// the instance's id.
	Conversions map[string][]*Conversion
}

func key(folder, id string) string { return folder + "|" + id }

// ownKey reports whether a kind's names live apart from every resource
// folder: an objective, a global and a jigsaw hole may share a name with
// a function or anything else.
func ownKey(k Kind) bool {
	return k.MMS == "Objective" || k.MMS == "Global" || k.MMS == "Storage" || k.MMS == "Test" || k.MMS == JigsawHoleKind
}

// Find returns the symbol of a resource folder ("worldgen/noise") and id.
func (t *Table) Find(folder, id string) *Symbol { return t.byKey[key(folder, id)] }

// Same returns this table's symbol for the declaration s names, which may
// come from an earlier table, or nil when it is no longer declared.
func (t *Table) Same(s *Symbol) *Symbol {
	if ownKey(s.Kind) {
		return t.byKey[key(s.Kind.MMS, s.ID())]
	}
	return t.byKey[key(s.Kind.Folder(), s.ID())]
}

func (t *Table) hasNamespace(ns string) bool {
	for _, s := range t.all {
		if s.NS == ns {
			return true
		}
	}
	return false
}

// OtherKind returns the first resource declaration with this id whose files
// live in a folder other than folder, or nil. It explains a bare reference
// that names a local declaration of the wrong kind.
func (t *Table) OtherKind(folder, id string) *Symbol {
	for _, s := range t.all {
		if s.ID() == id && s.Kind.MMS == "" && s.Kind.Folder() != "" && s.Kind.Folder() != folder {
			return s
		}
	}
	return nil
}

// All returns every symbol in declaration order.
func (t *Table) All() []*Symbol { return t.all }

var validPath = regexp.MustCompile(`^[a-z0-9_.\-/]+$`)
var validNSChars = regexp.MustCompile(`^[a-z0-9_.\-]+$`)

// ValidNamespace reports whether ns can name a namespace. "." and ".."
// would name a folder above data/, so they can't.
func ValidNamespace(ns string) bool {
	return validNSChars.MatchString(ns) && strings.Trim(ns, ".") != ""
}

// pathEscapes reports whether a resource path has an empty, "." or ".."
// segment, which would write outside the namespace's folder.
func pathEscapes(path string) bool {
	for _, part := range strings.Split(path, "/") {
		if part == "" || part == "." || part == ".." {
			return true
		}
	}
	return false
}

// Collect builds the symbol table. defaultNS applies until a file's first
// `Namespace x` declaration. It returns the table and its diagnostics,
// keyed by unit path.
func Collect(units []*Unit, kinds *Kinds, defaultNS string) (*Table, map[string][]syntax.Diagnostic) {
	t := &Table{byKey: map[string]*Symbol{}, DeclNS: map[*syntax.Decl]string{}, DeclKind: map[*syntax.Decl]Kind{}}
	diags := map[string][]syntax.Diagnostic{}
	report := func(u *Unit, code string, sp syntax.Span, msg string) {
		diags[u.Path] = append(diags[u.Path], syntax.Diagnostic{Code: code, Severity: syntax.SeverityError, Message: msg, Sp: sp})
	}
	for _, u := range units {
		ns := defaultNS
		for _, d := range u.File.Decls {
			k, ok := kinds.Lookup(d.Kind.Name)
			if !ok {
				if want, ok := Spelled(d.Kind.Name); ok {
					diags[u.Path] = append(diags[u.Path], respell("mms/unknown-kind", d.Kind.Sp, d.Kind.Name, want))
					continue
				}
				if kinds.IsValueType(d.Kind.Name) {
					diags[u.Path] = append(diags[u.Path], valueTypeDecl(d))
					continue
				}
				report(u, "mms/unknown-kind", d.Kind.Sp, "unknown declaration kind '"+d.Kind.Name+"'")
				continue
			}
			t.DeclKind[d] = k
			if k.MMS == "Namespace" {
				if d.Name == nil {
					continue
				}
				if !ValidNamespace(d.Name.Text) {
					report(u, "mms/invalid-name", d.Name.Sp, "namespace '"+d.Name.Text+"' must be lowercase letters, digits, '_', '-' or '.'")
					continue
				}
				ns = d.Name.Text
				continue
			}
			t.DeclNS[d] = ns
			if d.Name == nil {
				continue // OnLoad { } and friends
			}
			sns, path := splitID(d.Name.Text)
			if sns == "" {
				sns = ns
			}
			if !validPath.MatchString(path) || !ValidNamespace(sns) {
				report(u, "mms/invalid-name", d.Name.Sp, "resource names must be lowercase: use '"+Snake(d.Name.Text)+"'")
				continue
			}
			if pathEscapes(path) {
				report(u, "mms/invalid-name", d.Name.Sp, "a resource name's parts between '/' can't be empty, '.' or '..'")
				continue
			}
			if ext := fileExt(path); ext != "" {
				// The build adds the file's extension, so this would write
				// ruby.json.json, and references written without it miss.
				diags[u.Path] = append(diags[u.Path], syntax.Diagnostic{
					Code: "mms/invalid-name", Severity: syntax.SeverityError, Sp: d.Name.Sp,
					Message: "resource names leave out the file extension, which MMS adds itself: use '" + strings.TrimSuffix(d.Name.Text, ext) + "'",
					Fixes:   []syntax.Fix{{Title: "Remove '" + ext + "'", Safe: true, Edits: []syntax.TextEdit{{Sp: syntax.Span{Start: d.Name.Sp.End - len(ext), End: d.Name.Sp.End}}}}},
				})
				continue
			}
			sym := &Symbol{NS: sns, Path: path, Kind: k, Decl: d, Unit: u, table: t}
			kk := key(k.Folder(), sym.ID())
			if ownKey(k) {
				kk = key(k.MMS, sym.ID())
			}
			if prev, dup := t.byKey[kk]; dup {
				if k.TagOf != "" || k.Merges() {
					// Tags merge across declarations, like vanilla tag
					// files, and so do the entries of a language's lang
					// file.
					t.all = append(t.all, sym)
					continue
				}
				report(u, "mms/duplicate", d.Name.Sp, d.Kind.Name+" '"+sym.ID()+"' is already declared in "+prev.Unit.Path)
				continue
			}
			t.byKey[kk] = sym
			t.all = append(t.all, sym)
		}
	}
	return t, diags
}

// valueTypeDecl reports `BiomeSource src { ... }`: the name is a schema
// type that only appears inside other files (a dimension's biome
// source, a feature's placement), so it has no file of its own. A named
// one is a typed const, and a declaration with a plain body can be
// rewritten to one.
func valueTypeDecl(d *syntax.Decl) syntax.Diagnostic {
	t := d.Kind.Name
	diag := syntax.Diagnostic{
		Code: "mms/unknown-kind", Severity: syntax.SeverityError, Sp: d.Kind.Sp,
		Message: "'" + t + "' is a value inside other files, not a file of its own; name one with `const <name>: " + t + " = ...`",
	}
	if d.Name != nil && len(d.Name.Parts) == 0 && len(d.More) == 0 && d.Params == nil && d.Type == nil && d.Value == nil && d.Body != nil && !d.Extern && !d.Primary {
		sp := syntax.Span{Start: d.Kind.Sp.Start, End: d.Name.Sp.End}
		diag.Fixes = []syntax.Fix{{
			Title: "Write as a const",
			Edits: []syntax.TextEdit{{Sp: sp, Text: "const " + d.Name.Text + ": " + t + " ="}},
			Safe:  true,
		}}
	}
	return diag
}
