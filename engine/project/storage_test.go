package project

import (
	"context"
	"fmt"
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
		{"Type S { tier: 1 | 2 | 3 }\nStorage s: S\nFunction f {\n  s.tier = 4\n}", "expected one of 1, 2 or 3, but this is 4"},
		{"Type S { tier: 1 | 2 | 3 }\nStorage s: S\nFunction f {\n  s.tier = \"x\"\n}", "expected int, but this is string"},
		{"Type Tier = 1 | 2\nFunction f {\n  let t: Tier = 3\n}", "expected one of 1 or 2, but this is 3"},
		{"Function f {\n  let t: 0 | 1 = 1\n  t = -1\n}", "expected one of 0 or 1, but this is -1"},
		{"Function f {\n  let m: \"a\" | \"b\" = \"c\"\n}", `expected one of "a" or "b", but this is "c"`},
		{"Type S { speed: 1 | 2.5 }", "a number in a type must be a whole number"},
		{"Type S { big: 3000000000 | 1 }", "3000000000 does not fit an int"},
		{"Type S {\n  blah: { x: int, y: int }\n}\nStorage s: S\nFunction f {\n  s.blah.z = 1\n}", "s.blah has no field 'z'"},
		{"Type S {\n  a: int\n}\nStorage s: S\nFunction f {\n  s.a.b = 1\n}", "s.a is an integer, not a record"},
		{"Type S {\n  blah: { x: int, y }\n}", "a record's field needs a type: y: Type"},
		{"const modes: [string] = [\"a\", \"b\"]\nFunction f {\n  let m = modes[2]\n}", "index 2 is outside the list of 2"},
		{"Global g: int = 0\nconst xs: [int] = [1, g]\nFunction f {\n  let m = xs[g]\n}", "holds int or string constants"},
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

// A field of a union of whole numbers is an int functions read and
// write, and it starts at the first of them. A let can have such a type
// too, written in place or named.
func TestStorageIntChoices(t *testing.T) {
	src := `Type Tier = 1 | 2 | 3
Type S {
  tier: Tier
  step: -1 | 0 | 1 = 0
}
Storage s: S
Function f(): int {
  let next: 1 | 2 | 3
  if (s.tier == 1) {
    next = 2
  }
  s.tier = next
  s.step = -1
  return s.tier + s.step
}
`
	_, r := bookProject(t, src, nil)
	if d := diagText(r); d != "" {
		t.Fatal(d)
	}
	ld := emitted(t, r, "data/demo/function/mms_gen/load.mcfunction")
	for _, want := range []string{
		`execute unless data storage demo:mms s.tier run data modify storage demo:mms s.tier set value 1`,
		`execute unless data storage demo:mms s.step run data modify storage demo:mms s.step set value 0`,
	} {
		if !strings.Contains(ld, want) {
			t.Errorf("load lacks %s:\n%s", want, ld)
		}
	}
	vm := load(t, r, "demo")
	if got := call(t, vm, "demo:f"); got != 1 {
		t.Errorf("f = %d, want 1", got)
	}
	if got := call(t, vm, "demo:f"); got != 0 {
		t.Errorf("f again = %d, want 0", got)
	}
}

// A field can hold a record written in place, or a Type; functions reach
// its fields by path, and the load function sets each one's default.
func TestStorageNested(t *testing.T) {
	src := `Type Pos {
  x: int
  y: int = 64
}
Type S {
  blah: {
    x: int
    y: int = 2
    /** The way it faces. */
    facing: "north" | "south"
  }
  home: Pos
}
Storage s: S
Function f(): int {
  s.blah.x += 3
  s.home.x = s.blah.x + s.blah.y
  if (s.blah.facing == "north") {
    s.blah.facing = "south"
  }
  return s.home.x + s.home.y
}
`
	_, r := bookProject(t, src, nil)
	if d := diagText(r); d != "" {
		t.Fatal(d)
	}
	ld := emitted(t, r, "data/demo/function/mms_gen/load.mcfunction")
	for _, want := range []string{
		"execute unless data storage demo:mms s.blah.x run data modify storage demo:mms s.blah.x set value 0",
		"execute unless data storage demo:mms s.blah.y run data modify storage demo:mms s.blah.y set value 2",
		`execute unless data storage demo:mms s.blah.facing run data modify storage demo:mms s.blah.facing set value "north"`,
		"execute unless data storage demo:mms s.home.y run data modify storage demo:mms s.home.y set value 64",
	} {
		if !strings.Contains(ld, want) {
			t.Errorf("load lacks %s:\n%s", want, ld)
		}
	}
	vm := load(t, r, "demo")
	if got := call(t, vm, "demo:f"); got != 69 {
		t.Errorf("f = %d, want 69", got)
	}
	if got := call(t, vm, "demo:f"); got != 72 {
		t.Errorf("f again = %d, want 72", got)
	}
}

// A const list read at an index known only when the function runs picks
// its item then: cycling through modes kept in storage.
func TestConstListIndex(t *testing.T) {
	src := `Type S {
  mode: int
}
Storage st: S
const modes: [string] = ["Measure", "Replace", "Copy"]
const costs: [int] = [1, 5, 9]
Function next(): int {
  st.mode = (st.mode + 1) % len(modes)
  let name = modes[st.mode]
  tell(@a, "Mode: ${modes[st.mode]}")
  if (name == "Copy") {
    return 100 + costs[st.mode]
  }
  return costs[st.mode] + ${len(costs)} * 10
}
`
	_, r := bookProject(t, src, nil)
	if d := diagText(r); d != "" {
		t.Fatal(d)
	}
	vm := load(t, r, "demo")
	for i, want := range []int32{35, 109, 31, 35} {
		if got := call(t, vm, "demo:next"); got != want {
			t.Errorf("call %d: next = %d, want %d", i+1, got, want)
		}
	}
}

// TestConstListIndexConst reads a list at an index the function compiler
// knows, such as `${len(xs)} - 1`, which the evaluator does not fold.
func TestConstListIndexConst(t *testing.T) {
	src := `const costs: [int] = [1, 5, 9]
Function last(): int {
  return costs[${len(costs)} - 1]
}
`
	_, r := bookProject(t, src, nil)
	if d := diagText(r); d != "" {
		t.Fatal(d)
	}
	vm := load(t, r, "demo")
	if got := call(t, vm, "demo:last"); got != 9 {
		t.Errorf("last = %d, want 9", got)
	}

	p := New(spec.MustDefault(), Config{Namespace: "demo"})
	p.SetFile("main.mms", "const costs: [int] = [1, 5, 9]\nFunction f(): int {\n  return costs[${len(costs)}]\n}")
	r2, _ := p.Analyze(context.Background())
	var msgs []string
	for _, d := range r2.Diagnostics("main.mms") {
		msgs = append(msgs, d.Message)
	}
	if want := "index 3 is outside the list, which has 3 items (0 to 2)"; !strings.Contains(strings.Join(msgs, "\n"), want) {
		t.Errorf("diagnostics %q lack %q", msgs, want)
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

// TestStorageOptionalRecordDefaults keeps the inner defaults of a record
// field whose type is wrapped in parentheses or marked optional.
func TestStorageOptionalRecordDefaults(t *testing.T) {
	src := `Type S {
  a: ({ x: int = 5 })
  b: { y: int = 7 }?
}
Storage s: S
`
	_, r := bookProject(t, src, nil)
	if d := diagText(r); d != "" {
		t.Fatal(d)
	}
	ld := emitted(t, r, "data/demo/function/mms_gen/load.mcfunction")
	for _, want := range []string{
		"s.a.x set value 5",
		"s.b.y set value 7",
	} {
		if !strings.Contains(ld, want) {
			t.Errorf("load lacks %s:\n%s", want, ld)
		}
	}
}

// TestStorageIncremental checks that a change to a nested field or to a
// field's choices reaches functions that use the Storage.
func TestStorageIncremental(t *testing.T) {
	src := `Namespace demo
Type S {
  blah: { x: int }
  mode: "a" | "b"
}
Storage s: S
Function f {
  s.blah.x = 1
  s.mode = "b"
}
`
	edits := []func(string) string{
		func(s string) string { return strings.Replace(s, "blah: { x: int }", "blah: { x: string }", 1) },
		func(s string) string { return strings.Replace(s, `"a" | "b"`, `"a" | "c"`, 1) },
		func(s string) string { return strings.Replace(s, "blah: { x: string }", "blah: { x: int }", 1) },
		func(s string) string { return strings.Replace(s, `"a" | "c"`, `"a" | "b"`, 1) },
	}
	describe := func(r *Result) string {
		var sb strings.Builder
		for _, d := range r.Diagnostics("main.mms") {
			fmt.Fprintf(&sb, "%v %s %s\n", d.Sp, d.Code, d.Message)
		}
		sb.WriteString(dumpFiles(r.Emit()))
		return sb.String()
	}
	p := New(spec.MustDefault(), DefaultConfig())
	p.SetFile("main.mms", src)
	p.Analyze(context.Background())
	for i, edit := range edits {
		src = edit(src)
		p.SetFile("main.mms", src)
		got, _ := p.Analyze(context.Background())
		fresh := New(spec.MustDefault(), DefaultConfig())
		fresh.SetFile("main.mms", src)
		want, _ := fresh.Analyze(context.Background())
		if g, w := describe(got), describe(want); g != w {
			t.Errorf("edit %d: incremental result differs\n--- incremental\n%s\n--- fresh\n%s", i, g, w)
		}
	}
}

// A string shown with ${...} reads as plain text, not as SNBT: the game
// quotes and colours a string tag green unless its nbt component
// interprets it.
func TestShownStringIsPlain(t *testing.T) {
	src := `Type S {
  mode: int
  name: string = "none"
}
Storage st: S
const modes: [string] = ["Measure", "Replace", "Copy"]
Function next {
  st.mode = (st.mode + 1) % len(modes)
  let current = modes[st.mode]
  st.name = current
  actionbar(@s, "Set mode to ${current}")
  tell(@a, "Stored ${st.name}, ${st.mode}")
}
`
	_, r := bookProject(t, src, nil)
	if d := diagText(r); d != "" {
		t.Fatal(d)
	}
	vm := load(t, r, "demo")
	call(t, vm, "demo:next")
	if got, want := strings.Join(vm.Said, "|"), "Set mode to Replace|Stored Replace, 1"; got != want {
		t.Errorf("said %q, want %q", got, want)
	}
}
