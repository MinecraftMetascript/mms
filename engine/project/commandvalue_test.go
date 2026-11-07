package project

import (
	"context"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// TestCommandValue checks `/command` as the value of an assignment: its
// result goes through execute store result, a bool target stores
// success, an execute command carries the store without "run execute",
// ${...} works in it, and Blueprints fill it in.
func TestCommandValue(t *testing.T) {
	p := New(spec.MustDefault(), Config{Namespace: "demo"})
	p.SetFile("main.mms", `
Objective kills
Global g = 0
Type Probe {
  full: bool
}
Storage probe: Probe
Function f(k: int) {
  let items = /data get block ~ ~ ~ Items
  let chest: bool = /execute if block ~ ~ ~ chest
  g = /time query gametime
  kills[@s] = /data get entity @s Health
  probe.full = /execute if data block ~ ~ ~ Items[26]
  let n = /data get storage demo:x list[${k}]
  if (chest && items > n) {
    /say ok
  }
}
Blueprint Reader {
  path: string
  emit {
    Function ${self.path} {
      let n = /data get storage demo:x ${path}
      /say ${n}
    }
  }
}
Reader read_a { path = "a.b" }
`)
	r, _ := p.Analyze(context.Background())
	if d := diagText(r); d != "" {
		t.Fatalf("unexpected diagnostics:\n%s", d)
	}
	files := map[string]string{}
	for _, f := range r.Emit() {
		files[f.Path] = string(f.Data)
	}
	f := files["data/demo/function/f.mcfunction"]
	for _, want := range []string{
		"execute store result score #f.-t0 mms.demo run data get block ~ ~ ~ Items\n",
		"execute store success score #f.-t1 mms.demo if block ~ ~ ~ chest\n",
		"execute store result score #f.-t2 mms.demo run time query gametime\nscoreboard players operation $g mms.demo = #f.-t2 mms.demo\n",
		"execute store result score #f.-t3 mms.demo run data get entity @s Health\nscoreboard players operation @s kills = #f.-t3 mms.demo\n",
		"execute store success score #f.-t4 mms.demo if data block ~ ~ ~ Items[26]\nexecute store result storage demo:mms probe.full byte 1 run scoreboard players get #f.-t4 mms.demo\n",
	} {
		if !strings.Contains(f, want) {
			t.Errorf("f lacks %q:\n%s", want, f)
		}
	}
	if m := files["data/demo/function/f/_macro0.mcfunction"]; m != "$execute store result score #f.-t5 mms.demo run data get storage demo:x list[$(a0)]\n" {
		t.Errorf("macro = %q", m)
	}
	if got := files["data/demo/function/read_a.mcfunction"]; !strings.Contains(got, "run data get storage demo:x a.b\n") {
		t.Errorf("the Blueprint's ${path} is not filled in:\n%s", got)
	}

	errs := map[string]string{
		"Function f {\n  let s: string = /say hi\n}":                                 "mms/type-mismatch",
		"Type T {\n  s: string\n}\nStorage st: T\nFunction f {\n  st.s = /say hi\n}": "mms/type-mismatch",
		"Global x = /time query gametime":                                            "mms/not-constant",
		"Function f {\n  let n = /scoreboard nonsense\n}":                            "mms/command",
	}
	for src, want := range errs {
		p := New(spec.MustDefault(), Config{Namespace: "demo"})
		p.SetFile("main.mms", src)
		r, _ := p.Analyze(context.Background())
		var codes []string
		for _, d := range r.Diagnostics("main.mms") {
			if d.Severity == syntax.SeverityError {
				codes = append(codes, d.Code)
			}
		}
		if len(codes) == 0 || codes[0] != want {
			t.Errorf("%q: errors %v, want %s", src, codes, want)
		}
	}
}

// TestCommandValueRuns runs command values in the interpreter.
func TestCommandValueRuns(t *testing.T) {
	vm := world(t, `
Function read(): int {
  /data modify storage demo:x a set value 7
  let a = /data get storage demo:x a
  let has: bool = /execute if data storage demo:x a
  let missing: bool = /execute if data storage demo:x b
  if (has && !missing) {
    return a
  }
  return -1
}
`)
	if got := call(t, vm, "demo:read"); got != 7 {
		t.Errorf("read() = %d, want 7", got)
	}
}
