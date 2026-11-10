package mcfn

import (
	"strings"
	"testing"

	"github.com/MinecraftMetascript/nbt-go"
)

// These pin how the interpreter treats storage, so a change to how it
// holds NBT cannot change what a test sees.

// run runs lines as one function and returns what it said; each line's
// result score lands in r.<n> when the line is an execute store.
func run(t *testing.T, lines ...string) *VM {
	t.Helper()
	vm := New(map[string][]string{"t:f": append([]string{"scoreboard objectives add r dummy"}, lines...)})
	if _, _, err := vm.Run("t:f"); err != nil {
		t.Fatal(err)
	}
	return vm
}

func score(t *testing.T, vm *VM, holder string) int32 {
	t.Helper()
	v, ok := vm.Score(holder, "r")
	if !ok {
		t.Fatalf("%s has no score", holder)
	}
	return v
}

func snbtAt(t *testing.T, vm *VM, id, path string) string {
	t.Helper()
	v, err := vm.Get(id, path)
	if err != nil {
		t.Fatalf("%s %s: %v", id, path, err)
	}
	return v.String()
}

func TestStorageSetAndGet(t *testing.T) {
	vm := run(t,
		`data modify storage d:s a set value 5`,
		`data modify storage d:s b set value 1.5d`,
		`data modify storage d:s c set value "hello"`,
		`data modify storage d:s d set value [1,2,3]`,
		`data modify storage d:s e set value {x:1,y:{z:"q"}}`,
		`data modify storage d:s f set value []`,
		`data modify storage d:s g.h.i set value 7`,
		`data modify storage d:s "quoted key" set value 8`,
		`execute store result score a r run data get storage d:s a`,
		`execute store result score a10 r run data get storage d:s a 10`,
		`execute store result score b r run data get storage d:s b`,
		`execute store result score b100 r run data get storage d:s b 100`,
		`execute store result score c r run data get storage d:s c`,
		`execute store result score d r run data get storage d:s d`,
		`execute store result score e r run data get storage d:s e`,
		`execute store result score d1 r run data get storage d:s d[1]`,
		`execute store result score dlast r run data get storage d:s d[-1]`,
		`execute store result score z r run data get storage d:s e.y`,
		`execute store result score q r run data get storage d:s "quoted key"`,
	)
	for holder, want := range map[string]int32{"a": 5, "a10": 50, "b": 1, "b100": 150, "c": 5, "d": 3, "e": 2, "d1": 2, "dlast": 3, "z": 1, "q": 8} {
		if got := score(t, vm, holder); got != want {
			t.Errorf("%s = %d, want %d", holder, got, want)
		}
	}
	for path, want := range map[string]string{
		"a": "5", "b": "1.5d", "c": `"hello"`, "d": "[1,2,3]", "e": `{x:1,y:{z:"q"}}`, "f": "[]", "g": "{h:{i:7}}", "e.y.z": `"q"`,
	} {
		if got := snbtAt(t, vm, "d:s", path); got != want {
			t.Errorf("%s = %s, want %s", path, got, want)
		}
	}
}

func TestStorageMissing(t *testing.T) {
	vm := New(nil)
	for _, path := range []string{"nope", "a.b", "l[3]", "l[-4]", "s.x", "l.x"} {
		vm.Exec(`data modify storage d:s l set value [1,2,3]`)
		vm.Exec(`data modify storage d:s s set value "x"`)
		if _, err := vm.Get("d:s", path); err == nil {
			t.Errorf("%s: no error", path)
		}
	}
	if err := vm.Exec(`execute store result score x r run data get storage d:s nope`); err == nil {
		t.Errorf("data get of nothing: no error")
	}
}

// data modify set fails when it changes nothing, as in the game.
func TestStorageSetUnchanged(t *testing.T) {
	vm := run(t,
		`data modify storage d:s a set value {x:1,y:"z"}`,
		`execute store success score same r run data modify storage d:s a set value {y:"z",x:1}`,
		`execute store success score other r run data modify storage d:s a set value {x:2,y:"z"}`,
		`execute store success score copy r run data modify storage d:s b set from storage d:s a`,
		`execute store success score missing r run data modify storage d:s c set from storage d:s nothing`,
	)
	for holder, want := range map[string]int32{"same": 0, "other": 1, "copy": 1, "missing": 0} {
		if got := score(t, vm, holder); got != want {
			t.Errorf("%s = %d, want %d", holder, got, want)
		}
	}
	// set from copies: changing the copy leaves the source alone.
	vm.Exec(`data modify storage d:s b.x set value 9`)
	if got := snbtAt(t, vm, "d:s", "a"); got != `{x:2,y:"z"}` {
		t.Errorf("source changed with its copy: %s", got)
	}
}

func TestStorageLists(t *testing.T) {
	vm := run(t,
		`data modify storage d:s l append value 2`,
		`data modify storage d:s l append value 3`,
		`execute store result score n r run data modify storage d:s l prepend value 1`,
		`data modify storage d:s l[1] set value 20`,
		`data modify storage d:s l[-1] set value 30`,
		`data modify storage d:s c append value {a:1}`,
		`data modify storage d:s c[0].a set value 2`,
		`data remove storage d:s l[0]`,
	)
	if got := score(t, vm, "n"); got != 3 {
		t.Errorf("prepend result %d, want the new length 3", got)
	}
	if got := snbtAt(t, vm, "d:s", "l"); got != "[20,30]" {
		t.Errorf("l = %s", got)
	}
	if got := snbtAt(t, vm, "d:s", "c"); got != "[{a:2}]" {
		t.Errorf("c = %s", got)
	}
	vm.Exec(`data modify storage d:s c set value 1`)
	if err := vm.Exec(`data modify storage d:s c append value 1`); err == nil {
		t.Errorf("append to a number: no error")
	}
	if err := vm.Exec(`data modify storage d:s l[5] set value 1`); err == nil {
		t.Errorf("set past the end of a list: no error")
	}
}

func TestStorageMergeAndRemove(t *testing.T) {
	vm := run(t,
		`data modify storage d:s m set value {a:1,b:2}`,
		`data modify storage d:s m merge value {b:3,c:"x"}`,
		`data modify storage d:s n merge value {a:1}`,
		`data remove storage d:s m.a`,
		`data remove storage d:s nothing.here`,
	)
	if got := snbtAt(t, vm, "d:s", "m"); got != `{b:3,c:"x"}` {
		t.Errorf("m = %s", got)
	}
	if got := snbtAt(t, vm, "d:s", "n"); got != "{a:1}" {
		t.Errorf("n = %s", got)
	}
	if err := vm.Exec(`data modify storage d:s m merge value 1`); err == nil {
		t.Errorf("merge of a number: no error")
	}
}

// data modify ... string reads a value as text: a string as itself, a
// number as written, anything else as SNBT.
func TestStorageString(t *testing.T) {
	vm := run(t,
		`data modify storage d:s src set value {s:"word",i:42,d:2.5d,l:[1],c:{k:"v"}}`,
		`data modify storage d:s s set string storage d:s src.s`,
		`data modify storage d:s i set string storage d:s src.i`,
		`data modify storage d:s d set string storage d:s src.d`,
		`data modify storage d:s l set string storage d:s src.l`,
		`data modify storage d:s c set string storage d:s src.c`,
	)
	for path, want := range map[string]string{"s": "word", "i": "42", "d": "2.5", "l": "[1]", "c": `{k:"v"}`} {
		v, err := vm.Get("d:s", path)
		if err != nil {
			t.Fatal(err)
		}
		if s, ok := v.(nbt.String); !ok || string(s) != want {
			t.Errorf("%s = %#v, want the string %q", path, v, want)
		}
	}
}

// execute store result storage writes the scaled result, floored for
// whole-number types, as a tag of the type it names.
func TestStoreResultStorage(t *testing.T) {
	vm := run(t,
		`scoreboard players set x r 7`,
		`execute store result storage d:s i int 1 run scoreboard players get x r`,
		`execute store result storage d:s half int 0.5 run scoreboard players get x r`,
		`execute store result storage d:s d double 0.5 run scoreboard players get x r`,
		`execute store success storage d:s ok byte 1 run scoreboard players get x r`,
		`execute store result storage d:s s short 1 run scoreboard players get x r`,
		`execute store result storage d:s l long 2 run scoreboard players get x r`,
		`execute store result storage d:s f float 0.5 run scoreboard players get x r`,
		`execute store result storage d:s g.h int 1 run scoreboard players get x r`,
	)
	for path, want := range map[string]string{"i": "7", "half": "3", "d": "3.5d", "ok": "1b", "s": "7s", "l": "14L", "f": "3.5f", "g.h": "7"} {
		if got := snbtAt(t, vm, "d:s", path); got != want {
			t.Errorf("%s = %s, want %s", path, got, want)
		}
	}
}

func TestStorageCondition(t *testing.T) {
	vm := run(t,
		`data modify storage d:s a set value {b:[1]}`,
		`execute store success score has r if data storage d:s a.b[0]`,
		`execute store success score not r if data storage d:s a.c`,
		`execute store success score unless r unless data storage d:s a.c`,
	)
	for holder, want := range map[string]int32{"has": 1, "not": 0, "unless": 1} {
		if got := score(t, vm, holder); got != want {
			t.Errorf("%s = %d, want %d", holder, got, want)
		}
	}
}

// Macro arguments are substituted as the game does: strings bare,
// numbers as written, anything else as SNBT.
func TestMacroArguments(t *testing.T) {
	vm := New(map[string][]string{
		"t:f": {
			`data modify storage d:s m set value {s:"hi there",i:3,d:0.25d,c:{a:1}}`,
			`function t:m with storage d:s m`,
		},
		"t:m": {`$say $(s) $(i) $(d) $(c)`},
	})
	if _, _, err := vm.Run("t:f"); err != nil {
		t.Fatal(err)
	}
	if got, want := strings.Join(vm.Said, "|"), "hi there 3 0.25 {a:1}"; got != want {
		t.Errorf("said %q, want %q", got, want)
	}
	vm = New(map[string][]string{"t:f": {`data modify storage d:s m set value 1`, `function t:m with storage d:s m`}, "t:m": {`$say $(x)`}})
	if _, _, err := vm.Run("t:f"); err == nil {
		t.Errorf("macro arguments from a number: no error")
	}
}

// tellraw shows what a player reads: an nbt component's string tag is
// quoted SNBT unless it interprets it.
func TestTextComponents(t *testing.T) {
	vm := run(t,
		`scoreboard players set p r 12`,
		`data modify storage d:s v set value {s:"word",q:'say "hi"',i:4,d:1.5d,l:[1,2],c:{a:"b"}}`,
		`tellraw @a {text:"plain"}`,
		`tellraw @a "bare"`,
		`tellraw @a ["a",{text:"b",color:"gold"},{text:"c",extra:["d",{text:"e"}]}]`,
		`tellraw @a {score:{name:"p",objective:"r"}}`,
		`tellraw @a {score:{name:"nobody",objective:"r"}}`,
		`tellraw @a {nbt:"v.s",storage:"d:s"}`,
		`tellraw @a {nbt:"v.s",storage:"d:s",interpret:true}`,
		`tellraw @a {nbt:"v.s",storage:"d:s",interpret:1b}`,
		`tellraw @a {nbt:"v.i",storage:"d:s"}`,
		`tellraw @a {nbt:"v.d",storage:"d:s"}`,
		`tellraw @a {nbt:"v.l",storage:"d:s"}`,
		`tellraw @a {nbt:"v.c",storage:"d:s"}`,
		`tellraw @a {nbt:"v.q",storage:"d:s"}`,
		`data modify storage d:s v.n set value "two\nlines"`,
		`tellraw @a {nbt:"v.n",storage:"d:s"}`,
		`tellraw @a {selector:"@s"}`,
		`tellraw @a {keybind:"key.jump"}`,
		`tellraw @a {text:'it\'s',extra:[{text:"\"q\" \u00e9\x21\s"}]}`,
		`tellraw @a {text:"tab\tnew\nline"}`,
		`tellraw @a {text:7}`,
		`title @a actionbar {text:"bar"}`,
	)
	want := []string{
		"plain", "bare", "abcde", "12", "",
		`"word"`, "word", "word", "4", "1.5", "[1,2]", `{a:"b"}`,
		// Quoted as the game quotes: the other quote from the first one
		// inside, and a newline escaped.
		`'say "hi"'`, `"two\nlines"`,
		"@s", "[key.jump]", `it's"q" é! `, "tab\tnew\nline", "7", "bar",
	}
	if strings.Join(vm.Said, "|") != strings.Join(want, "|") {
		t.Errorf("said\n%q\nwant\n%q", vm.Said, want)
	}
	for _, line := range []string{
		`tellraw @a {nbt:"nothing",storage:"d:s"}`,
		`tellraw @a {nbt:"x",entity:"@s"}`,
		`tellraw @a {text:"unclosed}`,
	} {
		if err := vm.Exec(line); err == nil {
			t.Errorf("%s: no error", line)
		}
	}
}

// Summon data keeps its tags apart from the rest of the data.
func TestSummonData(t *testing.T) {
	vm := run(t, `summon marker 0 0 0 {Tags:["a","b"],CustomName:"x",data:{n:1}}`)
	e := vm.Entities[0]
	if !e.Tags["a"] || !e.Tags["b"] || len(e.Tags) != 2 {
		t.Errorf("tags %v", e.Tags)
	}
	if got := e.NBT.String(); got != `{CustomName:"x",data:{n:1}}` {
		t.Errorf("nbt %s", got)
	}
	for _, line := range []string{`summon marker 0 0 0 {Tags:"a"}`, `summon marker 0 0 0 {Tags:[1]}`, `summon marker 0 0 0 [1]`} {
		if err := vm.Exec(line); err == nil {
			t.Errorf("%s: no error", line)
		}
	}
}
