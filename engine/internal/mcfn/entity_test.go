package mcfn

import (
	"strings"
	"testing"
)

func TestEntities(t *testing.T) {
	vm := New(map[string][]string{
		"t:count": {
			"scoreboard objectives add o dummy",
			"summon marker 0 64 0 {Tags:[\"spawn\"]}",
			"summon marker 4 64 0 {Tags:[\"spawn\"]}",
			"execute as @e[tag=spawn] at @s run summon zombie ~ ~ ~ {Tags:[\"enemy\"]}",
			"execute at @e[tag=spawn,limit=1] run summon minecraft:zombie ~1 ~ ~",
			"tag @e[type=zombie,tag=!enemy] add enemy",
			"execute store result score n o if entity @e[type=zombie,tag=enemy]",
			"kill @e[type=marker]",
			"execute store result score m o if entity @e[type=marker]",
			"execute if entity @e[type=zombie] run return run scoreboard players get n o",
			"return 0",
		},
	})
	got, ok, err := vm.Run("t:count")
	if err != nil || !ok || got != 3 {
		t.Fatalf("count = %d %v %v, want 3", got, ok, err)
	}
	if m, _ := vm.Score("m", "o"); m != 0 {
		t.Errorf("markers after kill = %d", m)
	}
	if p := vm.Entities[1].Pos; p != [3]float64{4, 64, 0} {
		t.Errorf("zombie at %v, want the second marker's position", p)
	}
}

func TestEntityUnsupported(t *testing.T) {
	for _, line := range []string{
		"kill @e[distance=..3]",
		"kill @p",
		"execute if entity @e[type=#minecraft:undead] run say hi",
		"kill @e[sort=nearest]",
	} {
		vm := New(nil)
		if err := vm.Exec(line); err == nil || !strings.Contains(err.Error(), "does not support") && !strings.Contains(err.Error(), "only supports") {
			t.Errorf("%s: err %v", line, err)
		}
	}
}

// Each entity keeps its own scores, which selectors read with scores=
// and which go when it is killed.
func TestEntityScores(t *testing.T) {
	vm := New(map[string][]string{
		"t:run": {
			"scoreboard objectives add hp dummy",
			"scoreboard objectives add o dummy",
			"summon zombie 0 64 0 {Tags:[\"a\"]}",
			"summon zombie 0 64 0 {Tags:[\"b\"]}",
			"scoreboard players set @e[type=zombie] hp 10",
			"execute as @e[tag=b] run scoreboard players remove @s hp 4",
			"scoreboard players set #two o 2",
			"execute as @e[tag=a] run scoreboard players operation @s hp *= #two o",
			"execute store result score #hurt o if entity @e[scores={hp=..9}]",
			"execute as @e[tag=a] store result score @s o run scoreboard players get @s hp",
			"kill @e[scores={hp=..9}]",
			"execute if score @e[tag=a,limit=1] hp matches 20 run return run scoreboard players get #hurt o",
			"return 0",
		},
	})
	got, ok, err := vm.Run("t:run")
	if err != nil || !ok || got != 1 {
		t.Fatalf("run = %d %v %v, want 1", got, ok, err)
	}
	if len(vm.Entities) != 1 {
		t.Fatalf("%d entities left, want 1", len(vm.Entities))
	}
	if v, _ := vm.Score(vm.Entities[0].key, "o"); v != 20 {
		t.Errorf("stored score = %d, want 20", v)
	}
	for _, m := range vm.Scores {
		if _, ok := m["\x00entity2"]; ok {
			t.Errorf("the killed zombie still has a score")
		}
	}
	if err := vm.Exec("summon zombie"); err != nil {
		t.Fatal(err)
	}
	if err := vm.Exec("scoreboard players get @e hp"); err == nil {
		t.Error("reading one score of two entities did not fail")
	}
}
