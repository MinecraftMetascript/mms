package command

import (
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/spec"
)

func checker(t *testing.T) *Checker {
	t.Helper()
	p, err := spec.Default()
	if err != nil {
		t.Fatal(err)
	}
	return New(p)
}

func TestValidCommands(t *testing.T) {
	c := checker(t)
	for _, cmd := range []string{
		"say hello world",
		"tellraw @a {\"text\":\"hi\",\"color\":\"red\"}",
		"tellraw @a[tag=x,scores={k=1..}] [\"a\", {text: 'b b'}]",
		"scoreboard objectives add kills deathCount",
		"scoreboard players operation #a mms.ns += #b mms.ns",
		"scoreboard players set @s kills 0",
		"execute as @a at @s if score @s kills matches 3.. run give @s minecraft:diamond 2",
		"execute if entity @e[type=zombie,distance=..5,limit=1] run say near",
		"execute store result score #x obj run data get storage ns:mms vars.x",
		"execute unless block ~ ~-1 ~ #minecraft:logs run setblock ~ ~ ~ stone[]",
		"setblock 0 64 0 minecraft:oak_stairs[facing=north]{}",
		"tp @s ^ ^ ^1",
		"tp @s ~ ~10 ~ 90 0",
		"function ns:foo/bar",
		"function #minecraft:load",
		"function ns:macro with storage ns:mms args",
		"data modify storage ns:mms stack append value {x: 1}",
		"effect give @s minecraft:speed 10 1 true",
		"gamemode creative @a",
		"return 1",
		"return run function ns:x",
		"give @s diamond_sword[enchantments={sharpness:5}] 1",
		"clear @s #minecraft:logs",
		"particle minecraft:dust{color:[1.0,0.0,0.0],scale:1} ~ ~ ~ 0 0 0 0 1",
		"summon zombie ~ ~ ~ {NoAI:1b}",
		"kill @e[type=!player]",
		"say ${name} joined",
		"scoreboard players set ${who} kills ${n}",
		"execute positioned ^ ^ ^${speed} run summon marker",
		"setblock ~ ~${dy} ~ stone",
		"weather clear 10s",
		"time set 1d",
		"execute positioned as @s run tp @s ~ ~1 ~",
		"item replace entity @s weapon.mainhand with stone",
		"attribute @s minecraft:max_health base set 40",
		"execute if predicate ns:is_night run say night",
		"scoreboard objectives setdisplay sidebar.team.red kills",
	} {
		if r := c.Check(cmd); r.Problem != nil {
			t.Errorf("%q: %s at %d", cmd, r.Problem.Message, r.Problem.Start)
		}
	}
}

func TestInvalidCommands(t *testing.T) {
	c := checker(t)
	for _, tc := range []struct{ cmd, want string }{
		{"sya hi", "unknown command 'sya'"},
		{"gamemode creatve", "unknown game mode 'creatve'"},
		{"scoreboard players set @s kills", "incomplete"},
		{"scoreboard players set @s kills x", "expected an integer"},
		{"give @s minecraft:diamnd", "unknown item 'minecraft:diamnd'"},
		{"tp @s ^ ~ ^", "cannot be mixed"},
		{"kill @x", "unknown selector"},
		{"kill @e[colour=red]", "unknown selector option 'colour'"},
		{"execute as @a run", "incomplete"},
		{"execute as @a rn say hi", "expected"},
		{"tellraw @a {\"text\":\"hi\"", "not closed"},
		{"effect give @s minecraft:speed 10 300", "at most 255"},
		{"setblock 0 64.5 0 stone", "whole block"},
		{"say", "incomplete"},
		{"time set 1d extra", "unexpected text"},
		{"execute if score @s kills matches abc run say", "range"},
		{"spectate @e", "only one entity"},
	} {
		r := c.Check(tc.cmd)
		if r.Problem == nil {
			t.Errorf("%q: no problem, want %q", tc.cmd, tc.want)
			continue
		}
		if !strings.Contains(r.Problem.Message, tc.want) {
			t.Errorf("%q: got %q at %d, want %q", tc.cmd, r.Problem.Message, r.Problem.Start, tc.want)
		}
	}
}

// A space inside an NBT path makes the game read the rest as the next
// argument; the problem says so and the fix joins the path again.
func TestNBTPathSpace(t *testing.T) {
	c := checker(t)
	for _, tc := range []struct{ cmd, want string }{
		{"data get entity @s SelectedItem components.minecraft:custom_data.current_mode_idx", "data get entity @s SelectedItem.components.minecraft:custom_data.current_mode_idx"},
		{"data get entity @s SelectedItem .components", "data get entity @s SelectedItem.components"},
		{"data get entity @s Inventory [0].id", "data get entity @s Inventory[0].id"},
		{"execute store result entity @s SelectedItem components.minecraft:custom_data.x int 1 run say hi", "execute store result entity @s SelectedItem.components.minecraft:custom_data.x int 1 run say hi"},
	} {
		r := c.Check(tc.cmd)
		if r.Problem == nil || r.Problem.Fix == nil {
			t.Errorf("%q: got %+v, want a fix", tc.cmd, r.Problem)
			continue
		}
		if !strings.Contains(r.Problem.Message, "cannot contain spaces") {
			t.Errorf("%q: message %q", tc.cmd, r.Problem.Message)
		}
		src := []rune(tc.cmd)
		f := r.Problem.Fix
		got := string(src[:f.Start]) + f.Text + string(src[f.End:])
		if got != tc.want {
			t.Errorf("%q: fixed to %q, want %q", tc.cmd, got, tc.want)
		}
		if fixed := c.Check(got); fixed.Problem != nil {
			t.Errorf("%q: still invalid: %s", got, fixed.Problem.Message)
		}
	}
	// Words that are not the rest of a path keep their usual messages.
	for _, cmd := range []string{"data get entity @s SelectedItem components 2", "execute store result entity @s Pos[0] intt 1 run say hi"} {
		if r := c.Check(cmd); r.Problem == nil || r.Problem.Fix != nil {
			t.Errorf("%q: got %+v, want a problem without a fix", cmd, r.Problem)
		}
	}
	if r := c.Check("data get entity @s SelectedItem.components.minecraft:custom_data.current_mode_idx"); r.Problem != nil {
		t.Errorf("dotted path: %s", r.Problem.Message)
	}
}
