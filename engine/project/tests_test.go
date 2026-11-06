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
		{"Test t {\n  expect 1 == 1\n}", "runs nothing"},
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
