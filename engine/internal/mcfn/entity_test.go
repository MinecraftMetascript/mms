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
