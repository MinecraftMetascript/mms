package project

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/spec"
)

// TestVanillaParity builds vanilla files rewritten in MMS and compares the
// output with the originals, ignoring key order, number formatting and the
// minecraft: prefix. Set MMS_VANILLA to an mcmeta 26.3-data checkout.
func TestVanillaParity(t *testing.T) {
	root := os.Getenv("MMS_VANILLA")
	if root == "" {
		t.Skip("set MMS_VANILLA to an mcmeta 26.3-data checkout")
	}
	src, err := os.ReadFile("testdata/vanilla/worldgen.mms")
	if err != nil {
		t.Fatal(err)
	}
	p := New(spec.MustDefault(), DefaultConfig())
	p.SetFile("worldgen.mms", string(src))
	r, _ := p.Analyze(context.Background())
	for _, d := range r.Diagnostics("worldgen.mms") {
		if !strings.HasPrefix(d.Code, "lint/") { // it overrides vanilla on purpose
			t.Errorf("%v %s %s", d.Sp, d.Code, d.Message)
		}
	}
	n := 0
	for _, f := range r.Emit() {
		if !strings.HasPrefix(f.Path, "data/") {
			continue
		}
		n++
		want, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(f.Path)))
		if err != nil {
			t.Errorf("%s: %v", f.Path, err)
			continue
		}
		if g, w := normalize(t, f.Data), normalize(t, want); !reflect.DeepEqual(g, w) {
			gb, _ := json.Marshal(g)
			wb, _ := json.Marshal(w)
			t.Errorf("%s differs:\n got %s\nwant %s", f.Path, gb, wb)
		}
	}
	if n < 6 {
		t.Errorf("only %d files built", n)
	}
}

func normalize(t *testing.T, b []byte) any {
	var v any
	if err := json.Unmarshal(b, &v); err != nil {
		t.Fatal(err)
	}
	var walk func(any) any
	walk = func(v any) any {
		switch x := v.(type) {
		case map[string]any:
			for k, e := range x {
				x[k] = walk(e)
			}
		case []any:
			for i, e := range x {
				x[i] = walk(e)
			}
		case string:
			return strings.TrimPrefix(x, "minecraft:")
		}
		return v
	}
	return walk(v)
}
