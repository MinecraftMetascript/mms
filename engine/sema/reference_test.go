package sema

import (
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/spec"
)

func TestReference(t *testing.T) {
	p := spec.MustDefault()
	ref := BuildReference(p, NewKinds(p))
	kinds := map[string]RefKind{}
	for _, k := range ref.Kinds {
		kinds[k.Name] = k
	}
	if k := kinds["SurfaceRule"]; k.Folder != "worldgen/material_rule" {
		t.Errorf("SurfaceRule folder = %q", k.Folder)
	}
	if k := kinds["Function"]; k.Folder != "" || k.Doc == "" {
		t.Errorf("Function = %+v, want an MMS kind with a doc", k)
	}
	helpers := map[string]RefHelper{}
	for _, g := range ref.Groups {
		for _, h := range g.Helpers {
			helpers[g.Dispatcher+" "+h.Name] = h
		}
	}
	for key, sig := range map[string]string{
		"minecraft:density_function Gradient":        "Gradient(axis: ",
		"minecraft:loot_condition SurvivesExplosion": "SurvivesExplosion()",
		"minecraft:biome_source Fixed":               "Fixed(biome: ",
	} {
		h, ok := helpers[key]
		if !ok {
			t.Errorf("no helper %s", key)
			continue
		}
		if len(h.Signature) < len(sig) || h.Signature[:len(sig)] != sig {
			t.Errorf("%s signature = %q, want prefix %q", key, h.Signature, sig)
		}
	}
}

// Reference types read as words, and every link in them lands somewhere.
func TestReferenceTypes(t *testing.T) {
	p := spec.MustDefault()
	ref := BuildReference(p, NewKinds(p))
	helpers := map[string]map[string]bool{}
	for _, g := range ref.Groups {
		helpers[g.Dispatcher] = map[string]bool{}
		for _, h := range g.Helpers {
			helpers[g.Dispatcher][h.Name] = true
		}
	}
	types := map[string]bool{}
	for _, ty := range ref.Types {
		types[ty.ID] = true
	}
	check := func(where string, parts []TypePart) {
		for _, part := range parts {
			for _, raw := range []string{"#[", " @ ", "[[", "struct {"} {
				if strings.Contains(part.Text, raw) {
					t.Errorf("%s: type %q has mcdoc syntax %q", where, plainParts(parts), raw)
				}
			}
			l := part.Link
			switch {
			case l == nil:
			case l.Type != "":
				if !types[l.Type] {
					t.Errorf("%s: link to missing type %s", where, l.Type)
				}
			case l.Group != nil:
				if hs, ok := helpers[*l.Group]; !ok || (l.Helper != "" && !hs[l.Helper]) {
					t.Errorf("%s: link to missing helper %s %s", where, *l.Group, l.Helper)
				}
			}
		}
	}
	var biome RefHelper
	for _, g := range ref.Groups {
		for _, h := range g.Helpers {
			if g.Dispatcher == "" && h.Name == "Biome" {
				biome = h
			}
			for _, prm := range h.Params {
				check(g.Dispatcher+" "+h.Name+"."+prm.Name, prm.TypeParts)
			}
		}
	}
	for _, ty := range ref.Types {
		check(ty.ID, ty.Def)
		for _, f := range ty.Fields {
			check(ty.ID+"."+f.Name, f.TypeParts)
		}
	}
	got := map[string]string{}
	for _, prm := range biome.Params {
		got[prm.Name] = prm.Type
	}
	if got["temperature"] != "number" || got["has_precipitation"] != "bool" || got["carvers"] != "CarverListRef" {
		t.Errorf("Biome params = %v", got)
	}
}
