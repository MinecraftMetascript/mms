package project

import (
	"strings"
	"testing"
)

// TestFunctionTypeAliases checks that a project Type that is an int, bool
// or string, or a union of strings, works where a function value type
// goes: parameters, returns, let and Global.
func TestFunctionTypeAliases(t *testing.T) {
	out := emitText(t, `
Type Count = int
Type Mode = "easy" | "hard"
Global start: Count = 3
Function twice(n: Count): Count {
  return n * 2
}
Function pick(m: Mode) {
  /say ${m}
}
Function f {
  let k: Count = twice(start)
  pick("hard")
}
`)
	if got := out["data/demo/function/twice.mcfunction"]; !strings.Contains(got, "scoreboard") {
		t.Errorf("twice = %q", got)
	}
	if _, ok := out["data/demo/function/f.mcfunction"]; !ok {
		t.Errorf("no f.mcfunction in %v", keys(out))
	}
}
