package sema

import (
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
