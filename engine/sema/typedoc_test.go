package sema

import (
	"sort"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/spec"
)

// Every schema type MMS code can name is explained: the reference has an
// entry for it, its hover says what it is, and completion where a value
// of it goes offers something. A type that regresses fails here by name.
func TestTypeDiscoverabilityMatrix(t *testing.T) {
	p := spec.MustDefault()
	ks := NewKinds(p)
	c := NewChecker(p, ks, &Table{}, "minecraft")
	ref := BuildReference(p, ks)
	entries := map[string]*RefType{}
	for _, e := range ref.Types {
		entries[e.ID] = e
	}
	var noEntry, noShape, noValues []string
	paths := nameableTypes(p, ks)
	sort.Strings(paths)
	for _, path := range paths {
		name := spec.ShortPath(path)
		if _, ok := entries[typeID(path)]; !ok {
			noEntry = append(noEntry, name)
		}
		e := c.DescribeType(path)
		if e == nil || TypeSummary(e) == "" || TypeSummary(e) == "any" {
			noShape = append(noShape, name)
		}
		if len(c.TypeCandidates(&spec.Type{Kind: "reference", Path: path})) == 0 && !valuelessOK(c, path) {
			noValues = append(noValues, name)
		}
	}
	for what, names := range map[string][]string{"no reference entry": noEntry, "no shape in its hover": noShape, "no value completion": noValues} {
		if len(names) > 0 {
			t.Errorf("%d types have %s: %s", len(names), what, strings.Join(names, ", "))
		}
	}
}

// valuelessOK reports whether no completion fits where a value of the
// type at path goes, because the schema says nothing of its shape.
func valuelessOK(c *Checker, path string) bool {
	v := c.norm(tview{T: &spec.Type{Kind: "reference", Path: path}})
	return v.T == nil || v.T.Kind == "any" || v.T.Kind == "enum" && len(v.T.Values) == 0
}

// Every link in the reference leads to an entry or a helper on its pages.
func TestReferenceLinks(t *testing.T) {
	p := spec.MustDefault()
	ref := BuildReference(p, NewKinds(p))
	types := map[string]bool{}
	for _, e := range ref.Types {
		types[e.ID] = true
	}
	helpers := map[string]bool{}
	groups := map[string]bool{}
	for _, g := range ref.Groups {
		groups[g.Dispatcher] = true
		for _, h := range g.Helpers {
			helpers[g.Dispatcher+"/"+h.Name] = true
		}
	}
	var broken []string
	check := func(where string, parts []TypePart) {
		for _, p := range parts {
			l := p.Link
			switch {
			case l == nil:
			case l.Type != "":
				if !types[l.Type] {
					broken = append(broken, where+" -> type "+l.Type)
				}
			case l.Group != nil && l.Helper != "":
				if !helpers[*l.Group+"/"+l.Helper] {
					broken = append(broken, where+" -> helper "+*l.Group+"/"+l.Helper)
				}
			case l.Group != nil:
				if !groups[*l.Group] {
					broken = append(broken, where+" -> group "+*l.Group)
				}
			}
		}
	}
	for _, g := range ref.Groups {
		for _, h := range g.Helpers {
			for _, prm := range h.Params {
				check(h.Name+"("+prm.Name+")", prm.TypeParts)
			}
		}
	}
	for _, e := range ref.Types {
		check(e.Name, e.Def)
		check(e.Name+" used by", e.UsedBy)
		check(e.Name+" includes", e.Includes)
		for _, f := range e.Fields {
			check(e.Name+"."+f.Name, f.TypeParts)
		}
	}
	if len(broken) > 0 {
		sort.Strings(broken)
		t.Errorf("%d broken links, such as %s", len(broken), strings.Join(broken[:min(len(broken), 20)], "; "))
	}
}

// ClimateParameter, the type that started this: its entry, hover and
// completion say it is a number or a [min, max] pair.
func TestClimateParameterDiscoverable(t *testing.T) {
	p := spec.MustDefault()
	c := NewChecker(p, NewKinds(p), &Table{}, "minecraft")
	path := c.SchemaTypePath("ClimateParameter")
	md := TypeMarkdown(c.DescribeType(path), path)
	for _, want := range []string{"number -2..2 | list of number -2..2 (exactly 2)", "`0` or `[-2, 2]`", "/reference/types/worldgen#type-data-worldgen-dimension-biome_source-ClimateParameter"} {
		if !strings.Contains(md, want) {
			t.Errorf("hover lacks %q:\n%s", want, md)
		}
	}
	var labels []string
	for _, cand := range c.TypeCandidates(&spec.Type{Kind: "reference", Path: path}) {
		labels = append(labels, cand.Label)
		if cand.Label == "[-2, 2]" && cand.Insert != "[${1:-2}, ${2:2}]" {
			t.Errorf("[min, max] inserts %q", cand.Insert)
		}
	}
	if strings.Join(labels, " ") != "0 [-2, 2]" {
		t.Errorf("completion = %v, want 0 and [-2, 2]", labels)
	}
}
