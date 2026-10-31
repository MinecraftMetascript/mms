package sema

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/spec"
)

// knownMismatch lists vanilla files that disagree with the pinned mcdoc,
// with the reason. Each is an upstream mcdoc gap, not an MMS bug.
var knownMismatch = map[string]string{
	"recipe/map_extending.json": "mcdoc requires result.id; vanilla writes result: {}",
	"wolf_sound_variant/":       "mcdoc's WolfSounds lacks step_sound, which 26.3 adds",
}

func knownMismatchFor(rel string) (string, bool) {
	for k, reason := range knownMismatch {
		if rel == k || (strings.HasSuffix(k, "/") && strings.HasPrefix(rel, k)) {
			return reason, true
		}
	}
	return "", false
}

// TestVanillaValidates checks every vanilla 26.3 data file against the
// spec pack's types, which tests the type model (unions, dispatch,
// templates) against ground truth. Set MMS_VANILLA to a checkout of
// misode/mcmeta at the 26.3-data tag.
func TestVanillaValidates(t *testing.T) {
	root := os.Getenv("MMS_VANILLA")
	if root == "" {
		t.Skip("set MMS_VANILLA to an mcmeta 26.3-data checkout")
	}
	p := spec.MustDefault()
	kinds := NewKinds(p)
	data := filepath.Join(root, "data", "minecraft")
	type stat struct{ files, bad int }
	stats := map[string]*stat{}
	var samples []string
	perKind := map[string]int{}
	for _, kind := range p.ResourceKinds() {
		dir := filepath.Join(data, filepath.FromSlash(kind))
		filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
			if err != nil || d.IsDir() || !strings.HasSuffix(path, ".json") {
				return nil
			}
			// A nested kind's folder lives inside its parent's.
			rel, _ := filepath.Rel(data, path)
			for _, other := range p.ResourceKinds() {
				if other != kind && strings.HasPrefix(other, kind+"/") && strings.HasPrefix(filepath.ToSlash(rel), other+"/") {
					return nil
				}
			}
			b, _ := os.ReadFile(path)
			dec := json.NewDecoder(bytes.NewReader(b))
			dec.UseNumber()
			var x any
			if dec.Decode(&x) != nil {
				return nil
			}
			st := stats[kind]
			if st == nil {
				st = &stat{}
				stats[kind] = st
			}
			st.files++
			if reason, ok := knownMismatchFor(filepath.ToSlash(rel)); ok {
				_ = reason
				return nil
			}
			if errs := Validate(p, kinds, kind, jsonv.FromAny(x)); len(errs) > 0 {
				st.bad++
				if n := perKind[kind]; n < 3 {
					perKind[kind]++
					samples = append(samples, filepath.ToSlash(rel)+": "+errs[0])
				}
			}
			return nil
		})
	}
	var keys []string
	for k := range stats {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	total, bad := 0, 0
	for _, k := range keys {
		total += stats[k].files
		bad += stats[k].bad
		if stats[k].bad > 0 {
			t.Logf("%-45s %4d/%4d invalid", k, stats[k].bad, stats[k].files)
		}
	}
	t.Logf("%d of %d vanilla files invalid", bad, total)
	for _, s := range samples {
		t.Log(s)
	}
	if bad > 0 {
		t.Fail()
	}
}

// TestVanillaAssetsValidate checks vanilla's resource pack files (block
// states, models, item models, ...) against the spec pack's types. Set
// MMS_VANILLA_ASSETS to a checkout of misode/mcmeta at 26.3-assets-json.
func TestVanillaAssetsValidate(t *testing.T) {
	root := os.Getenv("MMS_VANILLA_ASSETS")
	if root == "" {
		t.Skip("set MMS_VANILLA_ASSETS to an mcmeta 26.3-assets-json checkout")
	}
	p := spec.MustDefault()
	kinds := NewKinds(p)
	total, bad := 0, 0
	for _, kind := range p.AssetKinds() {
		dir := filepath.Join(root, "assets", "minecraft", p.Assets[kind])
		files, shown := 0, 0
		filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
			if err != nil || d.IsDir() || !strings.HasSuffix(path, ".json") {
				return nil
			}
			b, _ := os.ReadFile(path)
			dec := json.NewDecoder(bytes.NewReader(b))
			dec.UseNumber()
			var x any
			if dec.Decode(&x) != nil {
				return nil
			}
			rel, _ := filepath.Rel(root, path)
			if filepath.ToSlash(rel) == "assets/minecraft/lang/deprecated.json" {
				return nil // the lang/deprecated kind, not a language
			}
			files++
			if errs := assetErrors(Validate(p, kinds, kind, jsonv.FromAny(x))); len(errs) > 0 {
				bad++
				if shown < 3 {
					shown++
					t.Errorf("%s: %s", filepath.ToSlash(rel), errs[0])
				}
			}
			return nil
		})
		if files == 0 {
			t.Errorf("%s: no files in %s", kind, dir)
		}
		total += files
	}
	t.Logf("%d of %d vanilla asset files invalid", bad, total)
}

// assetErrors drops what vanilla's resource pack writes that the game
// ignores and mcdoc leaves out: "__comment" style fields and the editor
// element names and groups Blockbench writes into models.
func assetErrors(errs []string) []string {
	var out []string
	for _, e := range errs {
		if strings.Contains(e, "unknown field '__") || e == "$: unknown field 'groups'" || e == "$: unknown field 'texture_size'" ||
			(strings.Contains(e, ".elements[") && strings.HasSuffix(e, "unknown field 'name'")) {
			continue
		}
		out = append(out, e)
	}
	return out
}
