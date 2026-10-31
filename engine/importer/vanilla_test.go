package importer

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/format"
	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/project"
	"github.com/minecraftmetascript/mms/engine/sema"
	"github.com/minecraftmetascript/mms/engine/spec"
)

// TestVanillaRoundTrip imports every vanilla 26.3 data file into MMS,
// builds the result and compares it with the original. Set MMS_VANILLA
// to an mcmeta 26.3-data checkout.
func TestVanillaRoundTrip(t *testing.T) {
	root := os.Getenv("MMS_VANILLA")
	if root == "" {
		t.Skip("set MMS_VANILLA to an mcmeta 26.3-data checkout")
	}
	pack := spec.MustDefault()
	kinds := sema.NewKinds(pack)
	data := filepath.Join(root, "data", "minecraft")
	srcs := map[string]*strings.Builder{} // one MMS file per folder
	originals := map[string]jsonv.Value{}
	skipped := map[string]int{}
	filepath.WalkDir(data, func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(path, ".json") {
			return nil
		}
		rel, _ := filepath.Rel(data, path)
		rel = filepath.ToSlash(rel)
		folder, id := SplitFolder(kinds, pack, "minecraft", rel)
		if folder == "" {
			skipped[strings.SplitN(rel, "/", 2)[0]]++
			return nil
		}
		k, _ := kinds.ForFolder(folder)
		b, _ := os.ReadFile(path)
		v, err := jsonv.Parse(b)
		if err != nil {
			t.Errorf("%s: %v", rel, err)
			return nil
		}
		sb := srcs[folder]
		if sb == nil {
			sb = &strings.Builder{}
			srcs[folder] = sb
		}
		if k.TagOf != "" {
			sb.WriteString(TagDecl(k, "minecraft", id, v))
		} else {
			sb.WriteString(Decl(k, "minecraft", id, v))
		}
		originals["data/minecraft/"+rel] = v
		return nil
	})
	t.Logf("skipped folders: %v", skipped)

	p := project.New(pack, project.Config{Namespace: "minecraft"})
	for folder, sb := range srcs {
		src := sb.String()
		if out, err := format.Source(src); err == nil {
			src = out
		} else {
			t.Errorf("%s: formatting the import failed: %v", folder, err)
		}
		p.SetFile(strings.ReplaceAll(folder, "/", "_")+".mms", src)
	}
	r, err := p.Analyze(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	byCode := map[string]int{}
	var samples []string
	seenSample := map[string]bool{}
	for _, u := range r.Units {
		for _, d := range r.Diagnostics(u.Path) {
			if d.Severity != 1 {
				continue
			}
			if knownError(d.Message) {
				continue
			}
			byCode[d.Code]++
			key := d.Code + "|" + u.Path + "|" + firstWords(d.Message)
			if !seenSample[key] && len(samples) < 80 {
				seenSample[key] = true
				line, col := u.File.LineCol(d.Sp.Start)
				samples = append(samples, fmt.Sprintf("%s:%d:%d %s %s\n    %s", u.Path, line, col, d.Code, d.Message, strings.TrimSpace(u.File.LineText(line))))
			}
		}
	}
	built := map[string]bool{}
	mismatch := 0
	for _, f := range r.Emit() {
		want, ok := originals[f.Path]
		if !ok {
			continue
		}
		if strings.HasPrefix(f.Path, "data/minecraft/wolf_sound_variant/") {
			// The unknown step_sound field is dropped (see knownErrors).
			built[f.Path] = true
			continue
		}
		built[f.Path] = true
		got, _ := jsonv.Parse(f.Data)
		if !reflect.DeepEqual(normalize(got), normalize(want)) {
			mismatch++
			if mismatch <= 10 {
				t.Errorf("%s differs:\n got %s\nwant %s", f.Path, jsonv.Compact(got), jsonv.Compact(want))
			}
		}
	}
	// Fixing every lint finding (bare names, folded constants) must not
	// change the output either.
	fixed, fixes, err := p.FixAll(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if dir := os.Getenv("MMS_DUMP"); dir != "" {
		for path, text := range fixed {
			os.WriteFile(filepath.Join(dir, path), []byte(text), 0o644)
		}
	}
	r2, _ := p.Analyze(context.Background())
	refixed := 0
	for _, f := range r2.Emit() {
		want, ok := originals[f.Path]
		if !ok || strings.HasPrefix(f.Path, "data/minecraft/wolf_sound_variant/") {
			continue
		}
		got, _ := jsonv.Parse(f.Data)
		if !reflect.DeepEqual(normalize(got), normalize(want)) {
			refixed++
			if refixed <= 5 {
				t.Errorf("%s differs after fixes:\n got %s\nwant %s", f.Path, jsonv.Compact(got), jsonv.Compact(want))
			}
		}
	}
	t.Logf("applied %d fixes; %d files differ after fixing", fixes, refixed)
	missing := 0
	var missingSamples []string
	for path := range originals {
		if !built[path] {
			missing++
			if len(missingSamples) < 10 {
				missingSamples = append(missingSamples, path)
			}
		}
	}
	sort.Strings(missingSamples)
	t.Logf("%d files, %d built, %d mismatched, %d missing; errors by code %v", len(originals), len(built), mismatch, missing, byCode)
	if len(samples) > 0 || missing > 0 {
		t.Errorf("errors:\n%s\nmissing e.g. %v", strings.Join(samples, "\n"), missingSamples)
	}
}

// normalize makes values comparable: numbers by value, ids without the
// minecraft: prefix.
func normalize(v jsonv.Value) any {
	switch x := v.(type) {
	case *jsonv.Object:
		m := map[string]any{}
		for _, mem := range x.Members {
			m[strings.TrimPrefix(mem.Key, "minecraft:")] = normalize(mem.Value)
		}
		return m
	case jsonv.Array:
		out := make([]any, len(x))
		for i, e := range x {
			out[i] = normalize(e)
		}
		return out
	case jsonv.Number:
		f, _ := x.Float()
		return f
	case jsonv.String:
		return strings.TrimPrefix(string(x), "minecraft:")
	case jsonv.Bool:
		return bool(x)
	}
	return nil
}

func firstWords(s string) string {
	f := strings.Fields(s)
	if len(f) > 3 {
		f = f[:3]
	}
	return strings.Join(f, " ")
}

// knownErrors are upstream data problems the round trip reports correctly:
// vanilla files that disagree with the pinned mcdoc or registries.
var knownErrors = []string{
	// recipe/map_extending.json has result: {} (see D32).
	"missing field 'id'",
	// wolf_sound_variant: mcdoc's WolfSounds lacks 26.3's step_sound.
	"unknown field 'step_sound'",
	// A template pool element names a structure file the 26.3 registry
	// summary does not list.
	"unknown structure 'minecraft:ancient_city/walls/intact_horizontal_wall_stairs_5'",
}

func knownError(msg string) bool {
	for _, k := range knownErrors {
		if strings.HasPrefix(msg, k) {
			return true
		}
	}
	return false
}
