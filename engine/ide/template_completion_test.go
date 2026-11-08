package ide

import (
	"strings"
	"testing"
)

// Templates complete like the declarations they stand for: their
// parameter and return types list types, their bodies are checked
// against the return type and name the parameters, and a Commands
// Template's block is function code. A Blueprint's field types, `as`
// conversions, check block and the declarations in its emit block
// complete the same way, as do typed consts and Storage types.
func TestTemplateCompletion(t *testing.T) {
	cases := []struct {
		src        string
		want, not_ []string
	}{
		{"Template single(item: Id<item>): Lo^", []string{"LootPool", "int", "Id", "Commands", "Text"}, nil},
		{"Template single(item: Id<item>): ^ = Pool(1, item)", []string{"LootPool", "string"}, nil},
		{"Template single(item: ^): LootPool = Pool(1, item)", []string{"Id", "int", "Item"}, nil},
		{"Template single(item: I^): LootPool = Pool(1, item)", []string{"Id", "Ids", "int"}, nil},
		{"Template single(item: Id<^>): LootPool = Pool(1, item)", []string{"item", "block", "particle_type"}, []string{"int"}},
		{"Template single(item: Id<it^>): LootPool = Pool(1, item)", []string{"item", "block"}, []string{"int"}},
		{"Template single(item: Id<item>, rolls: int = 1): LootPool = Po^", []string{"Pool"}, []string{"int", "Id"}},
		{"Template single(item: Id<item>, rolls: int = 1): LootPool = Pool(ro^)", []string{"rolls"}, []string{"item"}},
		{"Template single(item: Id<item>, rolls: int = 1): LootPool = Pool(rolls, ^)", []string{"Item", "Tag"}, nil},
		{"Template ring(p: Id<particle>, radius: float, points: int = 16): Commands {\n  A^\n}", []string{"As", "At", "p", "radius", "points"}, []string{"LootPool"}},
		{"Template ring(p: Id<particle>, radius: float, points: int = 16): Commands {\n  let a = ra^\n}", []string{"radius"}, nil},
		{"Template ring(p: Id<particle>, radius: float, points: int = 16): Commands {\n  /particle ${^}\n}", []string{"p", "radius", "points"}, nil},
		{"Template ring(p: Id<particle>): Commands {\n  let n: ^\n}", []string{"int"}, []string{"LootPool"}},
		{"Blueprint B {\n  name: st^\n}", []string{"string", "Id"}, nil},
		{"Type T = { a: in^ }", []string{"int"}, nil},
		{"const x: In^ = 3", []string{"int", "Id"}, nil},
		{"const x: int = ^", nil, []string{"Id", "LootPool"}},
		{"Blueprint B {\n  name: string\n  emit {\n    LootTable t { po^ }\n  }\n}", []string{"pools"}, nil},
		// A statement in an emit block: declaration kinds, Blueprints and
		// the emit keywords, on an empty line or the first word typed.
		{"Blueprint throttle(): none {\n  emit {\n    Objective write_cooldown: dummy\n    ^\n  }\n}", []string{"Objective", "Function", "LootTable", "primary", "const", "if", "for", "throttle"}, []string{"Template", "Type", "Blueprint", "dummy"}},
		{"Blueprint B {\n  emit {\n    ^\n  }\n}", []string{"Objective", "primary"}, nil},
		{"Blueprint B {\n  emit {\n    Objective o: dummy\n    Fu^\n  }\n}", []string{"Function"}, nil},
		{"Blueprint B {\n  on: bool\n  emit {\n    if (on) {\n      ^\n    }\n  }\n}", []string{"Objective", "primary"}, nil},
		{"Blueprint B {\n  emit {\n    Objective o: ^\n  }\n}", []string{"dummy"}, []string{"primary", "Function"}},
		{"Blueprint B {\n  emit {\n    Function f {\n      ^\n    }\n  }\n}", []string{"As", "let"}, []string{"primary", "Objective"}},
		// A Blueprint's parameter and result types.
		{"Blueprint throttle(label: ^): none {\n  emit {\n  }\n}", []string{"string", "int", "Id"}, nil},
		{"Blueprint throttle(label: st^) {\n  emit {\n  }\n}", []string{"string"}, nil},
		{"Blueprint W {\n  cooldown: int = 20\n  check {\n    if (cooldown < 1) {\n      er^\n    }\n  }\n}", []string{"error", "warn", "cooldown"}, nil},
		{"Blueprint W {\n  cast: Function\n  emit {\n    primary CustomItem ${self.path} {\n      base = stick\n      on_right_click = run {\n        ca^\n      }\n    }\n  }\n}", []string{"cast", "As"}, nil},
		{"Blueprint W {\n  title: string\n  emit {\n    primary CustomItem ${self.path} {\n      base = stick\n      na^\n    }\n  }\n}", []string{"name", "on_hit"}, nil},
		{"Blueprint T {\n  title: string\n  as ItemStack = { id = gold_ingot, co^ }\n}", []string{"components", "count"}, nil},
		{"Blueprint T {\n  title: string\n  as Item^ = { id = gold_ingot }\n}", []string{"ItemStack", "Id"}, nil},
		{"const stormy: Predicate = WeatherCheck(th^)", []string{"thundering", "raining"}, nil},
		{"Type SiegeState = { wave: int }\nStorage siege: Si^", []string{"SiegeState"}, nil},
		{"Function f {\n  title(@a, \"x\", su^)\n}", []string{"subtitle"}, nil},
		{"Blueprint B {\n  name: string\n  rolls: int\n  emit {\n    LootTable t { pools = [Pool(ro^, apple)] }\n  }\n}", []string{"rolls"}, []string{"name"}},
	}
	for _, c := range cases {
		got := labelsAt(t, c.src+"\n")
		all := " " + strings.Join(got, " ") + " "
		for _, w := range c.want {
			if !strings.Contains(all, " "+w+" ") {
				t.Errorf("%q lacks %s: %.300s", c.src, w, strings.TrimSpace(all))
			}
		}
		for _, w := range c.not_ {
			if strings.Contains(all, " "+w+" ") {
				t.Errorf("%q offers %s: %.300s", c.src, w, strings.TrimSpace(all))
			}
		}
	}
}
