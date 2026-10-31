package sema

import (
	"regexp"
	"sort"

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
}

func key(folder, id string) string { return folder + "|" + id }

// Find returns the symbol of a resource folder ("worldgen/noise") and id.
func (t *Table) Find(folder, id string) *Symbol { return t.byKey[key(folder, id)] }

// Same returns this table's symbol for the declaration s names, which may
// come from an earlier table, or nil when it is no longer declared.
func (t *Table) Same(s *Symbol) *Symbol {
	if s.Kind.MMS == "Objective" || s.Kind.MMS == "Global" {
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

// All returns every symbol in declaration order.
func (t *Table) All() []*Symbol { return t.all }

// InNamespace returns the symbols of one folder in a namespace, sorted.
func (t *Table) InNamespace(folder, ns string) []*Symbol {
	var out []*Symbol
	for _, s := range t.all {
		if s.NS == ns && s.Kind.Folder() == folder {
			out = append(out, s)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Path < out[j].Path })
	return out
}

var validPath = regexp.MustCompile(`^[a-z0-9_.\-/]+$`)
var validNS = regexp.MustCompile(`^[a-z0-9_.\-]+$`)

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
				report(u, "mms/unknown-kind", d.Kind.Sp, "unknown declaration kind '"+d.Kind.Name+"'")
				continue
			}
			t.DeclKind[d] = k
			if k.MMS == "Namespace" {
				if d.Name == nil {
					continue
				}
				if !validNS.MatchString(d.Name.Text) {
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
			if !validPath.MatchString(path) || !validNS.MatchString(sns) {
				report(u, "mms/invalid-name", d.Name.Sp, "resource names must be lowercase: use '"+Snake(d.Name.Text)+"'")
				continue
			}
			sym := &Symbol{NS: sns, Path: path, Kind: k, Decl: d, Unit: u}
			kk := key(k.Folder(), sym.ID())
			if k.MMS == "Objective" || k.MMS == "Global" {
				kk = key(k.MMS, sym.ID())
			}
			if prev, dup := t.byKey[kk]; dup {
				if k.TagOf != "" {
					// Tags merge across declarations, like vanilla tag files.
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
