package project

import (
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// A negative number is a number wherever a number's shorthand applies:
// `YAbove(-128)` takes the height -128, as `YAbove(64)` takes 64.
func TestNegativeNumberForms(t *testing.T) {
	p := New(spec.MustDefault(), Config{Namespace: "test"})
	p.SetFile("main.mms", "SurfaceRule r = [\n  If(YAbove(-128)) { Block(deepslate) },\n  If(YAbove(64)) { Block(stone) },\n]\n")
	r := analyze(t, p)
	for _, d := range r.Diagnostics("main.mms") {
		if d.Severity == syntax.SeverityError {
			t.Errorf("%s", d.Message)
		}
	}
	out := string(jsonv.Compact(r.Order[0].Value))
	for _, want := range []string{`"absolute": -128`, `"absolute": 64`} {
		if !strings.Contains(out, want) {
			t.Errorf("output lacks %s: %s", want, out)
		}
	}
}
