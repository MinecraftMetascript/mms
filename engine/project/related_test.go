package project

import (
	"context"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/spec"
)

// TestRelated checks that a problem with text a Blueprint or const wrote
// stays where the user's code put it in, and points at the line that
// wrote it.
func TestRelated(t *testing.T) {
	for _, c := range []struct{ src, at, want, msg string }{
		{
			src: "Blueprint Thing {\n  n: int\n  emit {\n    Noise ${self.path} { base_octave = n, amplitude_modifiers = \"oops\" }\n  }\n}\nThing a { n = 1 }\n",
			at:  "a", want: `"oops"`, msg: "written in Blueprint Thing",
		},
		{
			src: "const bad = { base_octave = 1, amplitude_modifiers = \"oops\" }\nNoise a = bad\n",
			at:  "bad\n", want: "const bad = {", msg: "written in const bad",
		},
	} {
		p := New(spec.MustDefault(), Config{Namespace: "demo"})
		p.SetFile("a.mms", c.src)
		r, _ := p.Analyze(context.Background())
		ds := r.Diagnostics("a.mms")
		if len(ds) != 1 || len(ds[0].Related) != 1 {
			t.Errorf("%q: diagnostics = %+v", c.src, ds)
			continue
		}
		src := []rune(c.src)
		d, rel := ds[0], ds[0].Related[0]
		if got := string(src[d.Sp.Start:]); !strings.HasPrefix(got, c.at) {
			t.Errorf("reported at %q, want %q", got, c.at)
		}
		if got := string(src[rel.Sp.Start:rel.Sp.End]); !strings.HasPrefix(got, c.want) || rel.Path != "a.mms" || rel.Message != c.msg {
			t.Errorf("related = %q in %s (%s), want %q (%s)", got, rel.Path, rel.Message, c.want, c.msg)
		}
	}
}
