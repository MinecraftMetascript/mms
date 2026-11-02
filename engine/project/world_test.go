package project

import (
	"context"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// build compiles src as namespace demo and returns the pack as text,
// failing on any error.
func build(t *testing.T, src string) string {
	t.Helper()
	p := New(spec.MustDefault(), Config{Namespace: "demo"})
	p.SetFile("main.mms", src)
	r, err := p.Analyze(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	for _, d := range r.Diagnostics("main.mms") {
		if d.Severity == syntax.SeverityError {
			t.Fatalf("%s: %s", d.Code, d.Message)
		}
	}
	return dumpFiles(r.Emit())
}

func wantLines(t *testing.T, got string, wants ...string) {
	t.Helper()
	for _, w := range wants {
		if !strings.Contains(got, w) {
			t.Errorf("missing %q in:\n%s", w, got)
		}
	}
}

// TestWorldConditions checks the typed condition helpers as values.
func TestWorldConditions(t *testing.T) {
	got := build(t, `
Objective hp
Function f {
  let n = 0
  if (Block(~ ~-1 ~, stone) && !Entity(@a[distance=..8])) {
    /say ground
  }
  if (Entity(@e[tag=boss]) || n > 2) {
    /say boss
  }
  let near: bool = At(@p).Entity(@e[type=zombie, distance=..4])
  if (near) {
    /say near
  }
  if (Data(@s::Inventory[0]) && Data(demo:path, queue[0]) && Data(~ ~-1 ~, Items)) {
    /say data
  }
  if (Items(@s, weapon.mainhand, diamond_sword) && Predicate(is_night) && Dimension(the_nether) && Loaded(~ ~ ~) && Biome(~ ~ ~, plains)) {
    /say all
  }
  if (Execute("if block ~ ~ ~ air").Entity(@s) && hp[@s] > 0) {
    /say raw
  }
  Entity(@e[tag=boss]) {
    /say block form
  }
}
Predicate is_night = TimeCheck(clock = overworld, value = 13000..23000)
`)
	wantLines(t, got,
		"execute if block ~ ~-1 ~ stone unless entity @a[distance=..8] run say ground\n",
		"execute if entity @e[tag=boss] run scoreboard players set #f._t0 mms.demo 1\n",
		"scoreboard players set #f.near mms.demo 0\nexecute at @p if entity @e[type=zombie, distance=..4] run scoreboard players set #f.near mms.demo 1\n",
		"execute if score #f.near mms.demo matches 1 run say near\n",
		"execute if data entity @s Inventory[0] if data storage demo:path queue[0] if data block ~ ~-1 ~ Items run say data\n",
		"execute if items entity @s weapon.mainhand diamond_sword if predicate demo:is_night if dimension the_nether if loaded ~ ~ ~ if biome ~ ~ ~ plains run say all\n",
		"execute if block ~ ~ ~ air if entity @s if score @s hp matches 1.. run say raw\n",
		"execute if entity @e[tag=boss] run say block form\n",
	)
}

// TestNBT checks :: reads and writes, run through the interpreter where
// it can (storage), and as text otherwise.
func TestNBT(t *testing.T) {
	src := `
Global total: int
Function f {
  /data modify storage demo:path queue set value [{x: 3, y: 2.75}, {x: 5}]
  let x = demo:path::queue[0].x
  let y = path::queue[0].y * 100
  path::queue[1].x = x + 10
  path::queue[0].x += 4
  path::count = x
  path::name = "bob"
  path::flag = x > 1
  path::last = path::queue[-1]
  total = path::queue[1].x
}
Function g(y: int, x: int) {
  @s::Health = Float(20)
  @s::Pos[1] = Double(y)
  @e[tag=m, limit=1]::Motion[0] = Double(x)
  ~ ~-1 ~::Items[0].count = Byte(3)
  @s::Invulnerable = true
}
`
	vm := world(t, src)
	call(t, vm, "demo:f")
	if v := score(t, vm, "$total", "mms.demo"); v != 13 {
		t.Errorf("total = %d, want 13", v)
	}
	if v := score(t, vm, "#f.x", "mms.demo"); v != 3 {
		t.Errorf("x = %d, want 3", v)
	}
	if v := score(t, vm, "#f.y", "mms.demo"); v != 275 {
		t.Errorf("y = %d, want 275", v)
	}
	got := build(t, src)
	wantLines(t, got,
		"execute store result score #f.x mms.demo run data get storage demo:path queue[0].x\n",
		"execute store result score #f.y mms.demo run data get storage demo:path queue[0].y 100\n",
		"execute store result storage demo:path queue[1].x int 1 run scoreboard players get #f._t",
		"execute store result score #f._t",
		"execute store result storage demo:path queue[0].x int 1 run scoreboard players get #f._t",
		"execute store result storage demo:path count int 1 run scoreboard players get #f.x mms.demo\n",
		"data modify storage demo:path name set value \"bob\"\n",
		"execute store result storage demo:path flag byte 1 run scoreboard players get #f._t",
		"data modify storage demo:path last set from storage demo:path queue[-1]\n",
		"data modify entity @s Health set value 20f\n",
		"execute store result entity @s Pos[1] double 1 run scoreboard players get #g.y mms.demo\n",
		"execute store result entity @e[tag=m,limit=1] Motion[0] double 1 run scoreboard players get #g.x mms.demo\n",
		"data modify block ~ ~-1 ~ Items[0].count set value 3b\n",
		"data modify entity @s Invulnerable set value 1b\n",
	)
}

// TestCommandValues checks /command on the right of = and in conditions.
func TestCommandValues(t *testing.T) {
	src := `
Function f {
  /data modify storage demo:s list set value [1, 2, 3]
  let n = /data get storage demo:s list
  let ok: bool = /data get storage demo:s missing
  let seen = 0
  let there: bool = /data get storage demo:s list
  if (there) {
    seen = 1
  }
  if (!ok && n == 3) {
    seen += 10
  }
  seen += /data get storage demo:s list[1]
}
Function h(): int {
  return /data get storage demo:s list[2]
}
`
	vm := world(t, src)
	call(t, vm, "demo:f")
	if v := score(t, vm, "#f.n", "mms.demo"); v != 3 {
		t.Errorf("n = %d, want 3", v)
	}
	if v := score(t, vm, "#f.ok", "mms.demo"); v != 0 {
		t.Errorf("ok = %d, want 0", v)
	}
	if v := score(t, vm, "#f.seen", "mms.demo"); v != 13 {
		t.Errorf("seen = %d, want 13", v)
	}
	if v := call(t, vm, "demo:h"); v != 3 {
		t.Errorf("h() = %d, want 3", v)
	}
	got := build(t, src)
	wantLines(t, got,
		"execute store result score #f.n mms.demo run data get storage demo:s list\n",
		"execute store success score #f.ok mms.demo run data get storage demo:s missing\n",
		"execute store success score #f.there mms.demo run data get storage demo:s list\n",
		"=== data/demo/function/h.mcfunction\nreturn run data get storage demo:s list[2]\n",
	)
}

// TestBuiltins checks Count, Abs, Min and Max.
func TestBuiltins(t *testing.T) {
	src := `
Function f(a: int, b: int): int {
  let m = Min(a, b, 7)
  let k = Max(a, Abs(b))
  let c = Abs(-4) + Min(3, 9) + Max(1, 2)
  return m * 100 + k * 10 + c
}
Function g {
  let n = Count(@e[tag=x])
}
`
	vm := world(t, src)
	if err := vm.Exec("scoreboard players set #f.a mms.demo 5"); err != nil {
		t.Fatal(err)
	}
	if err := vm.Exec("scoreboard players set #f.b mms.demo -9"); err != nil {
		t.Fatal(err)
	}
	if v := call(t, vm, "demo:f"); v != -900+90+9 {
		t.Errorf("f(5, -9) = %d, want -801", v)
	}
	got := build(t, src)
	wantLines(t, got, "execute store result score #g.n mms.demo if entity @e[tag=x]\n")
}

// TestUnsetScores checks that an unset objective score compares as 0.
func TestUnsetScores(t *testing.T) {
	src := `
Objective hp
Function f {
  let r = 0
  if (hp[@s] == 0) {
    r += 1
  }
  if (hp[@s] <= 0) {
    r += 10
  }
  if (hp[@s] < 5) {
    r += 100
  }
  if (hp[@s] != 0) {
    r += 1000
  }
  if (hp[@s] >= 1) {
    r += 10000
  }
  if (hp[@s] > -3) {
    r += 100000
  }
  let copy = hp[@s]
}
`
	vm := world(t, src)
	if err := vm.Exec("scoreboard players set #f.copy mms.demo 42"); err != nil {
		t.Fatal(err)
	}
	call(t, vm, "demo:f")
	if v := score(t, vm, "#f.r", "mms.demo"); v != 100111 {
		t.Errorf("unset: r = %d, want 100111", v)
	}
	if v := score(t, vm, "#f.copy", "mms.demo"); v != 0 {
		t.Errorf("unset copy = %d, want 0", v)
	}
	if err := vm.Exec("scoreboard players set @s hp 3"); err != nil {
		t.Fatal(err)
	}
	call(t, vm, "demo:f")
	if v := score(t, vm, "#f.r", "mms.demo"); v != 111100 {
		t.Errorf("set to 3: r = %d, want 111100", v)
	}
	if v := score(t, vm, "#f.copy", "mms.demo"); v != 3 {
		t.Errorf("copy = %d, want 3", v)
	}
}

func TestWorldErrors(t *testing.T) {
	cases := map[string]string{
		"Function f {\n  let x = At(@p)\n}":                                          "mms/type-mismatch",
		"Function f {\n  let x: int = Entity(@a)\n}":                                 "mms/type-mismatch",
		"Function f {\n  Entity(@a)\n}":                                              "mms/missing-value",
		"Function f {\n  Count(@a)\n}":                                               "mms/unused-value",
		"Function f {\n  let x = Count(1)\n}":                                        "mms/type-mismatch",
		"Function f {\n  let x = Float(2)\n}":                                        "mms/misplaced",
		"Function f {\n  let x = 1\n  x::y = 2\n}":                                   "mms/type-mismatch",
		"Function f {\n  @s::Pos[\"a\"] = 2\n}":                                      "mms/unsupported",
		"Function f {\n  @s::Pos = [1, 2, 3]\n}":                                     "mms/unsupported",
		"Function f {\n  @s::Health = Float(\"x\")\n}":                               "mms/type-mismatch",
		"Function f {\n  @s::Health += Float(1)\n}":                                  "mms/unsupported",
		"Function f {\n  let x = /sayy hi\n}":                                        "mms/command",
		"Function f {\n  let x = /say ${x}\n}":                                       "mms/unsupported",
		"Function f {\n  if (Block(~ ~ ~)) {\n  }\n}":                                "mms/missing-value",
		"Function f {\n  if (Block(~ ~ ~, not_a_block)) {\n  }\n}":                   "mms/command",
		"Function f {\n  let x = Min(1)\n}":                                          "mms/arity",
		"Global g: int = Count(@a)":                                                  "mms/not-constant",
		"Global g: int = @s::Health":                                                 "mms/not-constant",
		"Function f {\n  @s::Pos[1] = demo:s::x\n  @s::Pos[1] = Float(demo:s::x)\n}": "mms/unsupported",
	}
	for src, want := range cases {
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
