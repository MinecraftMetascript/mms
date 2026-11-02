package sema

import (
	"sort"
	"strings"

	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// Item hooks: a custom item's on_* fields each name a Function that runs
// when something happens with the item, as and at the player.
//
//	CustomItem ember_wand = { base = stick, name = "Ember Wand", on_right_click = zap }
//	Function zap {
//	  /particle flame ^ ^1.5 ^1 0.2 0.2 0.2 0.05 20
//	}
//
// Each hook is an advancement whose trigger matches the item by a
// custom_data marker, and a reward function that revokes it and calls
// the hook's function.
//
// Right click follows the consumable method of
// datapack.wiki/guide/right-click: the item gets a consumable component
// that takes days to finish, with no animation, so holding right click
// "uses" it without eating it, and a use_effects component so the player
// is not slowed. A using_item advancement fires every tick the button is
// held. on_hold_right_click runs on each of those ticks; for the others,
// a per-player score is set to 2 on every tick of use and the tick
// function counts it down, so they run only when it was 0: once per
// press. on_use_on_block and on_right_click_entity then cast a ray from
// the player's eyes, as far as their interaction ranges, and run at the
// first block or entity it meets.

// ItemHookMarker is the custom_data key that marks a custom item with
// hooks; its value is the item's id.
const ItemHookMarker = "mms_item"

// ItemHookTarget is the tag the clicked entity has while an
// on_right_click_entity function runs.
const ItemHookTarget = "mms_target"

type itemHook struct {
	field, desc string
	// press hooks share the using_item advancement; the others have
	// their own trigger.
	press   bool
	trigger string
}

var itemHooks = []itemHook{
	{field: "on_right_click", press: true,
		desc: "A `Function` without parameters to run, as and at the player, when they right-click with the item: once per press, not again while the button is held."},
	{field: "on_hold_right_click", press: true,
		desc: "A `Function` without parameters to run, as and at the player, on every tick they hold right click with the item."},
	{field: "on_use_on_block", press: true,
		desc: "A `Function` without parameters to run when a player right-clicks a block with the item, once per press: as the player, positioned at the centre of the block."},
	{field: "on_right_click_entity", press: true,
		desc: "A `Function` without parameters to run when a player right-clicks an entity with the item, once per press: as the player, at the entity, which has the tag `mms_target` while it runs."},
	{field: "on_hit", trigger: "player_hurt_entity",
		desc: "A `Function` without parameters to run, as and at the player, when they hurt an entity while holding the item in their main hand."},
	{field: "on_kill", trigger: "player_killed_entity",
		desc: "A `Function` without parameters to run, as and at the player, when they kill an entity while holding the item in their main hand."},
	{field: "on_consume", trigger: "consume_item",
		desc: "A `Function` without parameters to run, as and at the player, when they finish eating or drinking the item. The item needs a `consumable` component, or a base item that is food or a potion. It cannot be combined with the right-click hooks."},
	{field: "on_pickup", trigger: "thrown_item_picked_up_by_player",
		desc: "A `Function` without parameters to run, as and at the player, when they pick the item up from the ground."},
	{field: "on_break", trigger: "item_durability_changed",
		desc: "A `Function` without parameters to run, as and at the player, when the item runs out of durability and breaks. The item needs durability: a tool, weapon or armor base, or a `max_damage` component."},
}

func itemHookByField(f string) *itemHook {
	for i := range itemHooks {
		if itemHooks[i].field == f {
			return &itemHooks[i]
		}
	}
	return nil
}

// ItemHookID names the advancement and reward function of one of a
// custom item's hooks: "press" for the right-click ones, or the hook
// name without "on_".
func ItemHookID(id, hook string) string {
	ns, path := splitID(id)
	return ns + ":mms_gen/item_hook/" + path + "/" + hook
}

// ClickObjective is the score the press hooks use to run once per press,
// and to cast their ray.
func ClickObjective(ns string) string { return "mms_click." + ns }

// checkItemHooks checks a custom item's on_* fields, which name
// Functions without parameters, and adds the components they need. It
// returns the hooks by field.
func (c *Checker) checkItemHooks(d *syntax.Decl, o *jsonv.Object, comps *jsonv.Object, id string) map[string]string {
	fields := customItemFields(d)
	hooks := map[string]string{}
	press := ""
	for _, h := range itemHooks {
		if fid := c.checkHookFunction(h.field, fields[h.field], o.Get(h.field)); fid != "" {
			hooks[h.field] = fid
			if h.press && press == "" {
				press = h.field
			}
		}
	}
	if len(hooks) == 0 {
		return nil
	}
	if hooks["on_consume"] != "" && press != "" {
		c.errorf("mms/custom-item", fields["on_consume"].Span(), "on_consume needs the item to be eaten, but %s makes right click use it without ever finishing; leave out one or the other", press)
	}
	sp := d.Name.Sp
	if e := fields["components"]; e != nil {
		sp = e.Span()
	}
	if press != "" {
		if !hasComponent(comps, "consumable") {
			cons := &jsonv.Object{}
			// About eleven days: long enough to never finish, and to
			// never reach the point where eating sounds and particles
			// start.
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
		c.errorf("mms/custom-item", sp, "the on_* fields mark the item in its custom_data, which must be written as an object such as { kind = \"wand\" }")
		return hooks
	}
	data.Set(ItemHookMarker, jsonv.String(id))
	comps.Set(key, data)
	return hooks
}

// checkHookFunction checks an on_* field, which names a Function
// without parameters, and returns the function's id.
func (c *Checker) checkHookFunction(field string, e syntax.Expr, v jsonv.Value) string {
	s, _ := v.(jsonv.String)
	if s == "" || e == nil {
		return ""
	}
	fn := c.find("function", string(s))
	if fn == nil || fn.Kind.MMS != "Function" {
		_, path := splitID(string(s))
		c.errorf("mms/unknown-id", e.Span(), "no Function '%s'; %s names a function to run, such as Function %s { ... }", s, field, path)
		return ""
	}
	if len(SigOf(fn.Decl).Params) > 0 {
		c.errorf("mms/custom-item", e.Span(), "Function '%s' takes parameters, but %s calls it with none", s, field)
		return ""
	}
	return string(s)
}

// ItemHookFiles is what a custom item's hooks write: advancements and
// functions by id, and whether the namespace needs ClickSetup.
type ItemHookFiles struct {
	Advancements map[string]jsonv.Value
	Functions    map[string][]string
	Click        bool
}

// ItemHookOutput builds the files for a custom item's hooks. st is its
// item stack; hooks maps each on_* field to its function.
func ItemHookOutput(id string, st *jsonv.Object, hooks map[string]string) ItemHookFiles {
	out := ItemHookFiles{Advancements: map[string]jsonv.Value{}, Functions: map[string][]string{}}
	if len(hooks) == 0 {
		return out
	}
	base, _ := st.Get("id").(jsonv.String)
	itemPred := func() *jsonv.Object {
		marker := &jsonv.Object{}
		marker.Set(ItemHookMarker, jsonv.String(id))
		preds := &jsonv.Object{}
		preds.Set("minecraft:custom_data", marker)
		pred := &jsonv.Object{}
		pred.Set("items", base)
		pred.Set("predicates", preds)
		return pred
	}
	add := func(name, trigger string, conds *jsonv.Object, body []string) {
		hid := ItemHookID(id, name)
		crit := &jsonv.Object{}
		crit.Set("trigger", jsonv.String("minecraft:"+trigger))
		if conds != nil {
			crit.Set("conditions", conds)
		}
		criteria := &jsonv.Object{}
		criteria.Set(name, crit)
		rewards := &jsonv.Object{}
		rewards.Set("function", jsonv.String(hid))
		adv := &jsonv.Object{}
		adv.Set("criteria", criteria)
		adv.Set("rewards", rewards)
		out.Advancements[hid] = adv
		// Revoke first, so the advancement fires again next time.
		out.Functions[hid] = append([]string{"advancement revoke @s only " + hid}, body...)
	}
	mainhand := func() *jsonv.Object {
		eq := &jsonv.Object{}
		eq.Set("mainhand", itemPred())
		pred := &jsonv.Object{}
		pred.Set("equipment", eq)
		player := &jsonv.Object{}
		player.Set("type", jsonv.String("minecraft:entity_properties"))
		player.Set("entity", jsonv.String("this"))
		player.Set("predicate", pred)
		conds := &jsonv.Object{}
		conds.Set("player", player)
		return conds
	}

	var fields []string
	for f := range hooks {
		fields = append(fields, f)
	}
	sort.Strings(fields)
	for _, f := range fields {
		h := itemHookByField(f)
		if h == nil || h.press {
			continue
		}
		name := strings.TrimPrefix(f, "on_")
		call := []string{"function " + hooks[f]}
		switch f {
		case "on_hit", "on_kill":
			add(name, h.trigger, mainhand(), call)
		case "on_break":
			conds := &jsonv.Object{}
			conds.Set("item", itemPred())
			dur := &jsonv.Object{}
			dur.Set("max", jsonv.Number("0"))
			conds.Set("durability", dur)
			add(name, h.trigger, conds, call)
		default:
			conds := &jsonv.Object{}
			conds.Set("item", itemPred())
			add(name, h.trigger, conds, call)
		}
	}

	hold, click, block, entity := hooks["on_hold_right_click"], hooks["on_right_click"], hooks["on_use_on_block"], hooks["on_right_click_entity"]
	if hold == "" && click == "" && block == "" && entity == "" {
		return out
	}
	ns, _ := splitID(id)
	obj := ClickObjective(ns)
	press := ItemHookID(id, "press")
	var body []string
	if hold != "" {
		body = append(body, "function "+hold)
	}
	if click != "" || block != "" || entity != "" {
		out.Click = true
		body = append(body,
			// Still held since the last press: keep the score up and stop.
			"execute if score @s "+obj+" matches 1.. run return run scoreboard players set @s "+obj+" 2",
			"scoreboard players set @s "+obj+" 2")
		if click != "" {
			body = append(body, "function "+click)
		}
	}
	if block != "" || entity != "" {
		ray := press + "/ray"
		body = append(body,
			"tag @s add mms_clicker",
			"execute store result score #block "+obj+" run attribute @s block_interaction_range get 10",
			"execute store result score #entity "+obj+" run attribute @s entity_interaction_range get 10",
			"scoreboard players set #ray "+obj+" 0",
			// Back to feet anchoring, or each step would add the eye height.
			"execute anchored eyes positioned ^ ^ ^ anchored feet run function "+ray,
			"tag @s remove mms_clicker")
		out.Functions[ray] = rayLines(obj, ray, block, entity)
		if entity != "" {
			out.Functions[ray+"/entity"] = []string{
				"tag @s add " + ItemHookTarget,
				"execute at @s as @p[tag=mms_clicker] run function " + entity,
				"tag @s remove " + ItemHookTarget,
			}
		}
	}
	conds := &jsonv.Object{}
	conds.Set("item", itemPred())
	add("press", "using_item", conds, body)
	return out
}

// notTargeted lists the blocks a player's crosshair passes through.
var notTargeted = []string{"#minecraft:air", "water", "lava", "light", "structure_void", "bubble_column"}

// unpickable lists the entities a player's crosshair passes through.
var unpickable = []string{"item", "experience_orb", "marker", "area_effect_cloud", "arrow", "spectral_arrow", "trident", "block_display", "item_display", "text_display", "lightning_bolt"}

// rayLines is one step of the press hooks' ray: 0.1 blocks, as the
// player, from the player's eyes. It stops at the first entity within
// entity_interaction_range or block within block_interaction_range.
func rayLines(obj, ray, block, entity string) []string {
	ls := []string{"scoreboard players add #ray " + obj + " 1"}
	// An entity whose hitbox holds this point: its box meets both the
	// unit box below the point and the one above it.
	sel := "@e[dx=0,dy=0,dz=0,tag=!mms_clicker"
	for _, t := range unpickable {
		sel += ",type=!" + t
	}
	sel += "]"
	run := "return 1"
	if entity != "" {
		run = "return run function " + ray + "/entity"
	}
	ls = append(ls, "execute if score #ray "+obj+" <= #entity "+obj+" positioned ~-0.99 ~-0.99 ~-0.99 as "+sel+" positioned ~0.99 ~0.99 ~0.99 if entity @s[dx=0,dy=0,dz=0] run "+run)
	cond := "execute"
	for _, b := range notTargeted {
		cond += " unless block ~ ~ ~ " + b
	}
	run = "return 1"
	if block != "" {
		run = "return run execute align xyz positioned ~0.5 ~0.5 ~0.5 run function " + block
	}
	ls = append(ls, cond+" run "+run)
	ls = append(ls, "execute if score #ray "+obj+" < #block "+obj+" positioned ^ ^ ^0.1 run function "+ray)
	return ls
}

// ClickSetup is the load and tick lines a namespace with a press hook
// needs: the press objective, and counting it down.
func ClickSetup(ns string) (load, tick []string) {
	obj := ClickObjective(ns)
	return []string{"scoreboard objectives add " + obj + " dummy"},
		[]string{"scoreboard players remove @a[scores={" + obj + "=1..}] " + obj + " 1"}
}

// itemHookFields are the on_* fields of the CustomItem struct.
func itemHookFields() []*specField {
	var out []*specField
	for _, h := range itemHooks {
		out = append(out, &specField{h.field, h.desc})
	}
	return out
}

type specField struct{ key, desc string }
