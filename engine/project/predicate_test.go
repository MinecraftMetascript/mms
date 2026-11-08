package project

import (
	"context"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/spec"
)

const predicateSrc = `Predicate sneaking = EntityProperties(entity = this, predicate = { flags = { is_sneaking = true } })
Global wave: int = 0
Function is_sneaking(): bool {
  return sneaking
}
Function f {
  if (sneaking && wave > 2) {
    /say both
  }
  if (!demo:sneaking) {
    /say standing
  }
}
Function shadowed {
  let sneaking = 1
  if (sneaking == 1) {
    /say local
  }
}
`

// A Predicate's name in function code is the condition that it holds.
func TestPredicateCondition(t *testing.T) {
	_, r := bookProject(t, predicateSrc, nil)
	if d := diagText(r); d != "" {
		t.Fatal(d)
	}
	for path, want := range map[string]string{
		"data/demo/function/is_sneaking.mcfunction": "execute if predicate demo:sneaking run scoreboard players set",
		"data/demo/function/f.mcfunction":           "execute if predicate demo:sneaking if score $wave mms.demo matches 3.. run say both\nexecute unless predicate demo:sneaking run say standing",
		"data/demo/function/shadowed.mcfunction":    "say local",
	} {
		if got := emitted(t, r, path); !strings.Contains(got, want) {
			t.Errorf("%s lacks %q:\n%s", path, want, got)
		}
	}
}

func TestPredicateConditionErrors(t *testing.T) {
	const pred = "Predicate sneaking = EntityProperties(entity = this, predicate = { flags = { is_sneaking = true } })\n"
	for _, tc := range []struct{ src, want string }{
		{pred + "Function f {\n  if (sneaking()) {\n  }\n}", "'demo:sneaking' is a Predicate, not a function; test it without parentheses, as in if (sneaking)"},
		{pred + "Function f {\n  let n: int = sneaking\n}", "expected int"},
		{pred + "Function f {\n  let n = sneaking + 1\n}", "needs ints"},
		{pred + "Global sneaking: int = 0\nFunction f {\n  if (sneaking) {\n  }\n}", "'demo:sneaking' is both a Global and a Predicate; rename one of them"},
		{"Function f {\n  if (sneaking) {\n  }\n}", "unknown variable 'sneaking'"},
	} {
		p := New(spec.MustDefault(), Config{Namespace: "demo"})
		p.SetFile("main.mms", tc.src)
		r, _ := p.Analyze(context.Background())
		var msgs []string
		for _, d := range r.Diagnostics("main.mms") {
			msgs = append(msgs, d.Message)
		}
		if got := strings.Join(msgs, "\n"); !strings.Contains(got, tc.want) {
			t.Errorf("%q:\ngot  %s\nwant %s", tc.src, got, tc.want)
		}
	}
}
