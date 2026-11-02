package sema

import (
	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// Right click: a custom item's on_right_click names a Function that
// runs when a player right-clicks with the item.
//
//	CustomItem ember_wand = { base = stick, name = "Ember Wand", on_right_click = zap }
//	Function zap {
//	  /particle flame ^ ^1.5 ^1 0.2 0.2 0.2 0.05 20
//	}
//
// This follows the consumable method of datapack.wiki/guide/right-click.
// The item gets a consumable component that takes days to finish, with
// no animation, so holding right click "uses" it without eating it, and
// a use_effects component so the player is not slowed while doing so.
// A using_item advancement matches the item by a custom_data marker and
// calls the function from its reward function.
//
// using_item fires every tick the button is held, so the reward function
// keeps a per-player score: it is set to 2 on every tick of use and the
// tick function counts it down. The function runs only when the score was 0,
// which is once per press.

// RightClickID names both the advancement and the reward function that
// run the on_right_click of the custom item id.
func RightClickID(id string) string {
	ns, path := splitID(id)
	return ns + ":mms_gen/right_click/" + path
}

// ClickObjective is the score on_right_click uses to run once per press.
func ClickObjective(ns string) string { return "mms_click." + ns }

// RightClickMarker is the custom_data key that marks a custom item with
// an on_right_click; its value is the item's id.
const RightClickMarker = "mms_item"

// rightClickComponents adds what a custom item with an on_right_click
// needs to its components: consumable and use_effects unless the item
// sets them, and the marker in custom_data. It reports a custom_data the
// marker cannot join.
func (c *Checker) rightClickComponents(comps *jsonv.Object, id string, at syntax.Span) {
	if !hasComponent(comps, "consumable") {
		cons := &jsonv.Object{}
		// About eleven days: long enough to never finish, and to never
		// reach the point where eating sounds and particles start.
		cons.Set("consume_seconds", jsonv.Number("1000000"))
		cons.Set("animation", jsonv.String("none"))
		cons.Set("has_consume_particles", jsonv.Bool(false))
		comps.Set("consumable", cons)
	}
	if !hasComponent(comps, "use_effects") {
		eff := &jsonv.Object{}
		eff.Set("can_sprint", jsonv.Bool(true))
		eff.Set("speed_multiplier", jsonv.Number("1"))
		eff.Set("interact_vibrations", jsonv.Bool(false))
		comps.Set("use_effects", eff)
	}
	key := "custom_data"
	if comps.Get(key) == nil && comps.Get("minecraft:"+key) != nil {
		key = "minecraft:" + key
	}
	data := &jsonv.Object{}
	switch v := comps.Get(key).(type) {
	case nil:
	case *jsonv.Object:
		data.Members = append(data.Members, v.Members...)
	default:
		c.errorf("mms/custom-item", at, "on_right_click marks the item in its custom_data, which must be written as an object such as { kind = \"wand\" }")
		return
	}
	data.Set(RightClickMarker, jsonv.String(id))
	comps.Set(key, data)
}

// RightClickAdvancement is the using_item advancement that runs a
// custom item's on_right_click: it matches the base item with the
// marker, and its reward is the generated function at RightClickID.
func RightClickAdvancement(id string, st *jsonv.Object) jsonv.Value {
	base, _ := st.Get("id").(jsonv.String)
	marker := &jsonv.Object{}
	marker.Set(RightClickMarker, jsonv.String(id))
	preds := &jsonv.Object{}
	preds.Set("minecraft:custom_data", marker)
	pred := &jsonv.Object{}
	pred.Set("items", base)
	pred.Set("predicates", preds)
	conds := &jsonv.Object{}
	conds.Set("item", pred)
	crit := &jsonv.Object{}
	crit.Set("trigger", jsonv.String("minecraft:using_item"))
	crit.Set("conditions", conds)
	criteria := &jsonv.Object{}
	criteria.Set("click", crit)
	rewards := &jsonv.Object{}
	rewards.Set("function", jsonv.String(RightClickID(id)))
	adv := &jsonv.Object{}
	adv.Set("criteria", criteria)
	adv.Set("rewards", rewards)
	return adv
}

// RightClickLines is the advancement's reward function: it revokes the
// advancement, so it fires again next tick, and calls fnID once per
// press.
func RightClickLines(id, fnID string) []string {
	ns, _ := splitID(id)
	obj := ClickObjective(ns)
	return []string{
		"advancement revoke @s only " + RightClickID(id),
		// Still held since the last press: keep the score up and stop.
		"execute if score @s " + obj + " matches 1.. run return run scoreboard players set @s " + obj + " 2",
		"scoreboard players set @s " + obj + " 2",
		"function " + fnID,
	}
}

// checkRightClick checks a custom item's on_right_click, which names a
// Function without parameters, and returns the function's id.
func (c *Checker) checkRightClick(e syntax.Expr, v jsonv.Value) string {
	s, _ := v.(jsonv.String)
	if s == "" || e == nil {
		return ""
	}
	fn := c.find("function", string(s))
	if fn == nil || fn.Kind.MMS != "Function" {
		_, path := splitID(string(s))
		c.errorf("mms/unknown-id", e.Span(), "no Function '%s'; on_right_click names a function to run, such as Function %s { ... }", s, path)
		return ""
	}
	if len(SigOf(fn.Decl).Params) > 0 {
		c.errorf("mms/custom-item", e.Span(), "Function '%s' takes parameters, but on_right_click calls it with none", s)
		return ""
	}
	return string(s)
}

// RightClickSetup is the load and tick lines a namespace with an
// on_right_click needs: the press objective, and counting it down.
func RightClickSetup(ns string) (load, tick []string) {
	obj := ClickObjective(ns)
	return []string{"scoreboard objectives add " + obj + " dummy"},
		[]string{"scoreboard players remove @a[scores={" + obj + "=1..}] " + obj + " 1"}
}
