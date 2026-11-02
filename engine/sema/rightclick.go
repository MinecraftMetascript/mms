package sema

import (
	"strings"

	"github.com/minecraftmetascript/mms/engine/fn"
	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// Right-click hooks: code that runs when a player right-clicks with a
// custom item.
//
//	CustomItem ember_wand = { base = stick, name = "Ember Wand" }
//	OnRightClick ember_wand {
//	  /summon small_fireball ^ ^1.5 ^1
//	}
//
// This follows the consumable method of datapack.wiki/guide/right-click.
// The item gets a consumable component that takes days to finish, with
// no animation, so holding right click "uses" it without eating it, and
// a use_effects component so the player is not slowed while doing so.
// A using_item advancement matches the item by a custom_data marker and
// runs the body from its reward function.
//
// using_item fires every tick the button is held, so the reward function
// keeps a per-player score: it is set to 2 on every tick of use and the
// tick function counts it down. The body runs only when the score was 0,
// which is once per press.

const rightClickKind = "OnRightClick"

// rightClickKindFolder is the symbol folder OnRightClick hooks are in.
const rightClickKindFolder = "on_right_click"

// RightClickID names both the advancement and the reward function of
// the OnRightClick of the custom item id.
func RightClickID(id string) string {
	ns, path := splitID(id)
	return ns + ":mms_gen/right_click/" + path
}

// ClickObjective is the score OnRightClick uses to run once per press.
func ClickObjective(ns string) string { return "mms_click." + ns }

// RightClickMarker is the custom_data key that marks a custom item with
// an OnRightClick; its value is the item's id.
const RightClickMarker = "mms_item"

// rightClickComponents adds what a custom item with an OnRightClick
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
		c.errorf("mms/custom-item", at, "OnRightClick marks the item in its custom_data, which must be written as an object such as { kind = \"wand\" }")
		return
	}
	data.Set(RightClickMarker, jsonv.String(id))
	comps.Set(key, data)
}

// customItemBase is the base item id a custom item declaration names,
// namespaced, or "" when it has none.
func customItemBase(d *syntax.Decl) string {
	var text string
	switch b := customItemFields(d)["base"].(type) {
	case *syntax.Ident:
		text = b.Name
	case *syntax.ResLoc:
		text = b.Text
	default:
		return ""
	}
	if !strings.Contains(text, ":") {
		text = "minecraft:" + text
	}
	return text
}

// OnRightClick checks `OnRightClick item { ... }` and returns the
// using_item advancement and the compiled reward function. id is the
// custom item's "ns:path".
func (c *Checker) OnRightClick(d *syntax.Decl, id string, src []rune) (jsonv.Value, *fn.Func) {
	const usage = "OnRightClick needs a custom item and a body: OnRightClick ember_wand { ... }"
	if d.Name == nil || d.Body == nil {
		sp := d.Kind.Sp
		if d.Name != nil {
			sp = d.Name.Sp
		}
		c.errorf("mms/missing-value", sp, "%s", usage)
		return nil, nil
	}
	item := c.find("item", id)
	if item == nil || item.Kind.MMS != customItemKind {
		_, path := splitID(id)
		msg := "'" + id + "' is not a CustomItem; OnRightClick works on custom items, so declare one such as CustomItem " + path + " = { base = stick }"
		if item == nil && contains(c.Pack.Registries["item"], path) {
			msg = "'" + path + "' is a vanilla item; OnRightClick works on custom items, so declare one such as CustomItem my_" + path + " = { base = " + path + " } and use its name here"
		}
		c.errorf("mms/unknown-id", d.Name.Sp, "%s", msg)
		return nil, nil
	}
	in := c.info(d.Name)
	in.Symbol, in.ID = item, item.ID()
	in.Bare, in.Folder = BareID, "item"

	fid := RightClickID(id)
	f := c.CompileFunction(&syntax.Decl{Kind: d.Kind, Body: d.Body, Doc: d.Doc, Sp: d.Sp}, fid, src)
	base := customItemBase(item.Decl)
	if f == nil || base == "" {
		return nil, nil
	}
	marker := &jsonv.Object{}
	marker.Set(RightClickMarker, jsonv.String(id))
	preds := &jsonv.Object{}
	preds.Set("minecraft:custom_data", marker)
	pred := &jsonv.Object{}
	pred.Set("items", jsonv.String(base))
	pred.Set("predicates", preds)
	conds := &jsonv.Object{}
	conds.Set("item", pred)
	crit := &jsonv.Object{}
	crit.Set("trigger", jsonv.String("minecraft:using_item"))
	crit.Set("conditions", conds)
	criteria := &jsonv.Object{}
	criteria.Set("click", crit)
	rewards := &jsonv.Object{}
	rewards.Set("function", jsonv.String(fid))
	adv := &jsonv.Object{}
	adv.Set("criteria", criteria)
	adv.Set("rewards", rewards)

	ns, _ := splitID(id)
	obj := ClickObjective(ns)
	f.Body.Ops = append([]fn.Op{
		// Revoke first, so the advancement fires again next tick.
		fn.Cmd{Text: "advancement revoke @s only " + fid},
		// Still held since the last press: keep the score up and stop.
		fn.Cmd{Text: "execute if score @s " + obj + " matches 1.. run return run scoreboard players set @s " + obj + " 2"},
		fn.Cmd{Text: "scoreboard players set @s " + obj + " 2"},
	}, f.Body.Ops...)
	return adv, f
}

// RightClickSetup is the load and tick lines a namespace with an
// OnRightClick needs: the press objective, and counting it down.
func RightClickSetup(ns string) (load, tick []string) {
	obj := ClickObjective(ns)
	return []string{"scoreboard objectives add " + obj + " dummy"},
		[]string{"scoreboard players remove @a[scores={" + obj + "=1..}] " + obj + " 1"}
}
