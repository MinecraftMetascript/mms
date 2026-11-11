package project

import (
	"context"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

const testsSrc = `Type Siege { wave: int = 0 }
Storage game: Siege
Function next_wave {
  game.wave += 1
  for i in 1..game.wave * 2 {
    As(@e[tag=siege.spawn]).At(@s) {
      /summon zombie ~ ~ ~ {Tags:["siege.enemy"]}
    }
  }
}
Test second_wave_spawns_four {
  setup {
    game.wave = 1
    /summon marker 0 64 0 {Tags:["siege.spawn"]}
  }
  run next_wave()
  expect game.wave == 2
  expect count(@e[tag=siege.enemy]) == 4
}
Test wrong_count {
  run next_wave()
  expect exists(@e[tag=siege.enemy])
}
Function double(n: int): int {
  return n * 2
}
Test pure {
  expect double(21) == 42
}
Test off_by_one {
  setup {
    let start = 5
    game.wave = start
  }
  run next_wave()
  expect game.wave == start
}
`

func TestTests(t *testing.T) {
	_, r := bookProject(t, testsSrc, nil)
	if d := diagText(r); d != "" {
		t.Fatal(d)
	}
	for _, f := range r.Emit() {
		if strings.Contains(f.Path, "mms_test") {
			t.Errorf("the pack has a test function: %s", f.Path)
		}
	}
	res, err := r.RunTests(nil)
	if err != nil {
		t.Fatal(err)
	}
	got := map[string]TestResult{}
	for _, x := range res {
		got[x.Name] = x
	}
	if x := got["second_wave_spawns_four"]; !x.Pass {
		t.Errorf("second_wave_spawns_four failed: %s", x.Message)
	}
	if x := got["wrong_count"]; x.Pass || x.Expect != 1 || x.Message != "expected exists(@e[tag=siege.enemy])" {
		t.Errorf("wrong_count: %+v", x)
	}
	if x := got["pure"]; !x.Pass {
		t.Errorf("pure: %+v", x)
	}
	if x := got["off_by_one"]; x.Pass || x.Message != "expected game.wave == start, but game.wave is 6 and start is 5" {
		t.Errorf("off_by_one: %+v", x)
	}
	only, _ := r.RunTests(func(n string) bool { return n == "wrong_count" })
	if len(only) != 1 {
		t.Errorf("filtered run: %d results", len(only))
	}
}

func TestTestErrors(t *testing.T) {
	for _, tc := range []struct{ src, want string }{
		{"Function f {\n}\nTest t {\n  run f()\n}", "checks nothing"},
		{"Function f {\n}\nTest t {\n  run f()\n  setup {\n  }\n  expect true\n}", "setup comes first"},
		{"Function f {\n}\nTest t {\n  run f()\n  let x = 1\n  expect true\n}", "a Test holds a setup"},
		{"Function f {\n}\nTest t {\n  run f()\n  expect 1\n}", "expected a condition"},
		{"Function f {\n}\nTest t {\n  run g()\n  expect true\n}", "g"},
	} {
		p := New(spec.MustDefault(), Config{Namespace: "demo"})
		p.SetFile("main.mms", tc.src)
		r, _ := p.Analyze(context.Background())
		var msgs []string
		for _, d := range r.Diagnostics("main.mms") {
			if d.Severity == syntax.SeverityError {
				msgs = append(msgs, d.Message)
			}
		}
		if len(msgs) == 0 || !strings.Contains(strings.Join(msgs, "\n"), tc.want) {
			t.Errorf("%q: errors %q, want %q", tc.src, msgs, tc.want)
		}
	}
}

// A Global, a local or a Storage field given to a Commands Template in
// function code is its value when the code runs.
func TestTemplateRuntimeArgs(t *testing.T) {
	_, r := bookProject(t, `Global g: int = 5
Type S { wave: int = 7 }
Storage game: S
Global out: int = 0
Template add(x: int): Commands {
  out += x
}
Function f {
  let y = 3
  add(g)
  add(y)
  add(game.wave)
}
Test sums {
  run f()
  expect out == 15
}
`, nil)
	if d := diagText(r); d != "" {
		t.Fatal(d)
	}
	res, err := r.RunTests(nil)
	if err != nil || len(res) != 1 || !res[0].Pass {
		t.Errorf("%v %+v", err, res)
	}
}

// TestFunctionPitfalls runs code that used to need workarounds: a
// foldered function called without its namespace, a bool command value
// whose execute chain stops early, an assignment that reads its own
// target, and a test that runs a function tag.
func TestFunctionPitfalls(t *testing.T) {
	src := `Global x: bool = true
Global y: bool = true
Global ok: bool = true
Global calls: int = 0
Function players/update_tier {
  calls += 1
}
Function tiers {
  players/update_tier()
  demo:players/update_tier()
}
Function both {
  ok = /execute if entity @e[tag=none] if entity @e
}
Function and_self {
  x = x && y
}
Function flip {
  x = !x
}
FunctionTag hook = [players/update_tier, #inner]
FunctionTag inner = [players/update_tier, tiers]
Function fire {
  /function #demo:hook
}
Test foldered_call {
  run tiers()
  expect calls == 2
}
Test chain_stops_early {
  run both()
  expect !ok
}
Test and_keeps_true {
  run and_self()
  expect x
}
Test and_gives_false {
  setup { y = false }
  run and_self()
  expect !x
}
Test not_self {
  run flip()
  expect !x
}
Test function_tag {
  run fire()
  expect calls == 3
}
`
	_, r := bookProject(t, src, nil)
	if d := diagText(r); d != "" {
		t.Fatal(d)
	}
	res, err := r.RunTests(nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 6 {
		t.Fatalf("ran %d tests, want 6", len(res))
	}
	for _, tr := range res {
		if !tr.Pass {
			t.Errorf("%s: %s", tr.Name, tr.Message)
		}
	}
}
