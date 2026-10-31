package sema

import (
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/spec"
)

// Every section in helpdocs/ names a real helper, and every parameter it
// documents is one of that helper's.
func TestHelpDocKeys(t *testing.T) {
	p := spec.MustDefault()
	ref := BuildReference(p, NewKinds(p))
	params := map[string]map[string]bool{}
	dispatchers := map[string]map[string]bool{}
	for _, g := range ref.Groups {
		for _, h := range g.Helpers {
			key := h.Name
			if h.Case != "" {
				key = g.Dispatcher + "/" + strings.TrimPrefix(h.Case, "minecraft:")
			}
			ps := map[string]bool{}
			for _, p := range h.Params {
				ps[p.Name] = true
			}
			params[key] = ps
			if dispatchers[g.Dispatcher] == nil {
				dispatchers[g.Dispatcher] = map[string]bool{}
			}
			for k := range ps {
				dispatchers[g.Dispatcher][k] = true
			}
		}
	}
	for key, d := range HelpDocs() {
		known, ok := params[key]
		if d, wild := strings.CutSuffix(key, "/*"); wild {
			known, ok = dispatchers[d]
		}
		if !ok {
			t.Errorf("helpdocs: %s is not a helper", key)
			continue
		}
		for name := range d.Params {
			if !known[name] {
				t.Errorf("helpdocs: %s has no parameter %s", key, name)
			}
		}
	}
}

func TestParseHelpDocs(t *testing.T) {
	out := map[string]*HelperDoc{}
	parseHelpDocs("# Title\n\nignored\n\n## a/b\nFirst.\n\nSecond.\n\n- x: One\n  two.\n- `y`: Three.\n\n```mms\nX()\n```\n", out)
	d := out["a/b"]
	if d == nil || d.Summary != "First.\n\nSecond." || d.Params["x"] != "One two." || d.Params["y"] != "Three." || d.Example != "X()" {
		t.Errorf("parsed %+v", d)
	}
}
