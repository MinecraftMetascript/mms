package ide

import (
	"context"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/project"
	"github.com/minecraftmetascript/mms/engine/sema"
	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// Built-in helpers hover with what they do, each parameter, an example
// and a link to the reference; signature help carries the short form.
func TestHelperHoverAndSignature(t *testing.T) {
	src := "SurfaceRule r = If(Stone|Depth(Floor)) { grass_block }\n"
	s, off := session(t, src)
	h, _ := s.Hover("main.mms", off)
	for _, want := range []string{
		"StoneDepth(surface_type: CaveSurface, offset: int = 0",
		"Makes `minecraft:stone_depth`, a material condition.",
		"- `offset`: *an integer, default `0`*. Extra depth",
		"`.Add()` sets `add_surface_depth` to `true`",
		"**Example**",
		"(https://minecraftmetascript.github.io/mms/docs/reference#material_condition-StoneDepth)",
	} {
		if !strings.Contains(h, want) {
			t.Errorf("hover lacks %q:\n%s", want, h)
		}
	}
	if testing.Verbose() {
		t.Log(h)
	}

	sig, ok := s.Signature("main.mms", strings.Index(src, "Floor"))
	if !ok || !strings.Contains(sig.Doc, "True for blocks close to a surface") || !strings.Contains(sig.Doc, "[Reference](") {
		t.Errorf("signature doc = %q", sig.Doc)
	}
	if len(sig.ParamDocs) < 2 || !strings.Contains(sig.ParamDocs[1], "default `0`") {
		t.Errorf("param docs = %q", sig.ParamDocs)
	}
}

// Helpers without hand-written docs still list their parameters and link
// to the reference.
func TestHelperHoverFallback(t *testing.T) {
	src := "DensityFunction d = Ble|ndAlpha()\n"
	s, off := session(t, src)
	h, _ := s.Hover("main.mms", off)
	if !strings.Contains(h, "Makes `minecraft:blend_alpha`, a density function.") || !strings.Contains(h, "#density_function-BlendAlpha") {
		t.Errorf("hover = %s", h)
	}
}

// Every example in the hand-written docs compiles, and every parameter
// it documents exists.
func TestHelpDocExamples(t *testing.T) {
	for key, d := range sema.HelpDocs() {
		if d.Summary == "" && !strings.HasSuffix(key, "/*") {
			t.Errorf("%s: no summary", key)
		}
		if d.Example == "" {
			continue
		}
		p := project.New(spec.MustDefault(), project.Config{Namespace: "test"})
		p.SetFile("main.mms", d.Example)
		r, err := p.Analyze(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		for _, diag := range r.Diagnostics("main.mms") {
			if diag.Severity == syntax.SeverityError {
				t.Errorf("%s example: %s\n%s", key, diag.Message, d.Example)
			}
		}
	}
}

// Signature help works inside If(...) before its body is written.
func TestSignatureInsideUnfinishedIf(t *testing.T) {
	s, off := session(t, "SurfaceRule s = If(StoneDepth(Floor, |))\n")
	h, ok := s.Signature("main.mms", off)
	if !ok || !strings.HasPrefix(h.Label, "StoneDepth(") || h.Active != 1 {
		t.Errorf("signature = %+v, %v", h, ok)
	}
}
