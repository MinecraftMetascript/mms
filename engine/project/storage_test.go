package project

import (
	"context"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

const storageSrc = `Type State {
  wave: int = 3
  boss_alive: bool
  name: string = "Siege"
  mobs: [Id<entity_type>] = [zombie]
}
Storage game: State
Function step(): int {
  game.wave += 2
  game.boss_alive = game.wave > 4
  game.name = "Wave " + game.wave
  if (game.boss_alive) {
    return game.wave * 10
  }
  return game.wave
}
Function zombies(): int {
  return count(@e[type=zombie])
}
Function any_zombie(): bool {
  return exists(@e[type=zombie])
}
`

// A Storage keeps its record in command storage: the load function sets
// missing fields to their defaults, and functions read and write them.
func TestStorage(t *testing.T) {
	_, r := bookProject(t, storageSrc, nil)
	if d := diagText(r); d != "" {
		t.Fatal(d)
	}
	ld := emitted(t, r, "data/demo/function/mms_gen/load.mcfunction")
	for _, want := range []string{
		"execute unless data storage demo:mms game.wave run data modify storage demo:mms game.wave set value 3",
		"execute unless data storage demo:mms game.boss_alive run data modify storage demo:mms game.boss_alive set value false",
		`execute unless data storage demo:mms game.name run data modify storage demo:mms game.name set value "Siege"`,
		`execute unless data storage demo:mms game.mobs run data modify storage demo:mms game.mobs set value ["minecraft:zombie"]`,
	} {
		if !strings.Contains(ld, want) {
			t.Errorf("load lacks %s:\n%s", want, ld)
		}
	}
	vm := load(t, r, "demo")
	if got := call(t, vm, "demo:step"); got != 50 {
		t.Errorf("step = %d, want 50", got)
	}
	if got := call(t, vm, "demo:step"); got != 70 {
		t.Errorf("step again = %d, want 70", got)
	}
	if got := emitted(t, r, "data/demo/function/zombies.mcfunction"); !strings.Contains(got, "if entity @e[type=zombie]") {
		t.Errorf("zombies:\n%s", got)
	}
	if got := emitted(t, r, "data/demo/function/any_zombie.mcfunction"); !strings.Contains(got, "if entity @e[type=zombie]") {
		t.Errorf("any_zombie:\n%s", got)
	}
	// Loading again keeps the values.
	if _, _, err := vm.Run("demo:mms_gen/load"); err != nil {
		t.Fatal(err)
	}
	if got := call(t, vm, "demo:step"); got != 90 {
		t.Errorf("step after a reload = %d, want 90", got)
	}
}

func TestStorageErrors(t *testing.T) {
	for _, tc := range []struct{ src, want string }{
		{"Type S { a: int = 0 }\nStorage s: S\nFunction f {\n  s.b = 1\n}", "s has no field 'b'"},
		{"Type S { wave: int = 0 }\nStorage s: S\nFunction f {\n  s.wav = 1\n}", "did you mean 'wave'?"},
		{"Type S { a: [int] = [] }\nStorage s: S\nFunction f {\n  let x = s.a\n}", "reads int, bool and string fields"},
		{"Storage s: int", "holds a record"},
		{"Type S { a: int = 0 }\nStorage s", "needs a record type"},
		{"Type S { a: int = 0 }\nStorage g: S", "keeps its own values"},
		{"Type S { a: int = 0 }\nStorage s: S = {}", "starts from its type's defaults"},
		{"Type S { a: int = 0 }\nStorage s: S\nFunction f {\n  s.a = true\n}", "expected int"},
		{"Type S { mode: \"x\" | \"y\" }\nStorage s: S\nFunction f {\n  s.mode = \"z\"\n}", `expected one of "x" or "y", but this is "z"`},
		{"Type S { mode: \"x\" | \"y\" }\nStorage s: S\nFunction f {\n  s.mode += \"x\"\n}", "set it with ="},
		{"Function f(): bool {\n  return exists(\"x\")\n}", "expected a selector"},
		{"Function f(): int {\n  return count()\n}", "takes one selector"},
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

// A field of a union of string literals is a string functions read and
// write, and it starts at the first of them.
func TestStorageChoices(t *testing.T) {
	src := `Type S {
  mode: "x" | "y"
}
Storage s: S
Function f(): int {
  tell(@a, "Mode ${s.mode}")
  if (s.mode == "x") {
    s.mode = "y"
    return 1
  }
  return 0
}
`
	_, r := bookProject(t, src, nil)
	if d := diagText(r); d != "" {
		t.Fatal(d)
	}
	ld := emitted(t, r, "data/demo/function/mms_gen/load.mcfunction")
	if want := `execute unless data storage demo:mms s.mode run data modify storage demo:mms s.mode set value "x"`; !strings.Contains(ld, want) {
		t.Errorf("load lacks %s:\n%s", want, ld)
	}
	vm := load(t, r, "demo")
	if got := call(t, vm, "demo:f"); got != 1 {
		t.Errorf("f = %d, want 1", got)
	}
	if got := call(t, vm, "demo:f"); got != 0 {
		t.Errorf("f again = %d, want 0", got)
	}
}

// Changing a field's type rechecks the functions that use it.
func TestStorageDeps(t *testing.T) {
	p := New(spec.MustDefault(), Config{Namespace: "demo"})
	p.SetFile("a.mms", "Type S { a: int = 0 }\nStorage s: S\n")
	p.SetFile("b.mms", "Function f {\n  s.a += 1\n}\n")
	if _, err := p.Analyze(context.Background()); err != nil {
		t.Fatal(err)
	}
	p.SetFile("a.mms", "Type S { a: [int] = [] }\nStorage s: S\n")
	r, _ := p.Analyze(context.Background())
	found := false
	for _, d := range r.Diagnostics("b.mms") {
		found = found || strings.Contains(d.Message, "reads int, bool and string fields")
	}
	if !found {
		t.Errorf("b.mms was not rechecked: %v", r.Diagnostics("b.mms"))
	}
}
