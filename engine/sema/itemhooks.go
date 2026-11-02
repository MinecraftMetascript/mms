package sema

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
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
	// projectile hooks run from an enchantment on the arrow's weapon.
	projectile bool
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
	{field: "on_projectile_hit_block", projectile: true,
		desc: "A `Function` without parameters to run when an arrow hits a block: an arrow shot from the item, for a bow or crossbow base, or the arrow that is the item, for an arrow, tipped_arrow or spectral_arrow base. It runs as the arrow, at the point it hit; `execute on origin` is the shooter."},
	{field: "on_projectile_hit_entity", projectile: true,
		desc: "A `Function` without parameters to run when an arrow hurts an entity: an arrow shot from the item, for a bow or crossbow base, or the arrow that is the item, for an arrow, tipped_arrow or spectral_arrow base. It runs as and at the entity hit; `execute on attacker` is the shooter."},
	{field: "on_projectile_land", projectile: true,
		desc: "A `Function` without parameters to run where a thrown splash or lingering potion that is the item lands, whatever it hits. It runs at the spot, as a marker that is removed afterwards. Needs a splash_potion or lingering_potion base."},
}

// ProjectileWeapons are the bases whose projectile hooks follow the
// arrows they shoot; ProjectileAmmo are the bases that are the arrow.
var (
	ProjectileWeapons = []string{"minecraft:bow", "minecraft:crossbow"}
	ProjectileAmmo    = []string{"minecraft:arrow", "minecraft:tipped_arrow", "minecraft:spectral_arrow"}
	// ProjectileThrown are the bases on_projectile_land follows: thrown
	// items that keep no weapon, so no enchantment runs for them.
	ProjectileThrown = []string{"minecraft:splash_potion", "minecraft:lingering_potion"}
)

// ProjectileLandTag marks the markers riding thrown items with
// on_projectile_land, per namespace.
func ProjectileLandTag(ns string) string { return "mms_land." + ns }

// ProjectileTag marks an arrow the tick function has looked at, per
// namespace.
func ProjectileTag(ns string) string { return "mms_shot." + ns }

// ProjectileSpotted is the function the tick function runs as each new
// arrow.
func ProjectileSpotted(ns string) string { return ns + ":mms_gen/projectile/spotted" }

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
func (c *Checker) checkItemHooks(d *syntax.Decl, o *jsonv.Object, comps *jsonv.Object, id string, out *CustomItemOut) map[string]string {
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
	c.checkUseCost(fields, o, comps, press, out)
	if len(hooks) == 0 {
		return nil
	}
	for _, f := range []string{"on_projectile_hit_block", "on_projectile_hit_entity"} {
		if hooks[f] == "" {
			continue
		}
		b, _ := o.Get("base").(jsonv.String)
		if !slices.Contains(ProjectileWeapons, string(b)) && !slices.Contains(ProjectileAmmo, string(b)) {
			c.errorf("mms/custom-item", fields[f].Span(), "%s follows arrows, so the base must be a bow or crossbow (for the arrows it shoots) or an arrow, tipped_arrow or spectral_arrow (for itself)", f)
			delete(hooks, f)
		}
	}
	if f := "on_projectile_land"; hooks[f] != "" {
		if b, _ := o.Get("base").(jsonv.String); !slices.Contains(ProjectileThrown, string(b)) {
			msg := "on_projectile_land follows a thrown potion, so the base must be splash_potion or lingering_potion"
			if slices.Contains(ProjectileWeapons, string(b)) || slices.Contains(ProjectileAmmo, string(b)) {
				msg += "; for arrows use on_projectile_hit_block and on_projectile_hit_entity"
			}
			c.errorf("mms/custom-item", fields[f].Span(), "%s", msg)
			delete(hooks, f)
		}
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

// checkUseCost checks consume_on_use and damage_on_use, which need a
// right-click hook to pay them, and damage_on_use a max_damage component.
// It gives a damageable item a stack size of 1, as the game requires.
func (c *Checker) checkUseCost(fields map[string]syntax.Expr, o, comps *jsonv.Object, press string, out *CustomItemOut) {
	num := func(key string) int {
		n, _ := o.Get(key).(jsonv.Number)
		v, _ := strconv.Atoi(string(n))
		return v
	}
	count, damage := num("consume_on_use"), num("damage_on_use")
	for _, k := range []string{"consume_on_use", "damage_on_use"} {
		if num(k) > 0 && press == "" {
			c.errorf("mms/custom-item", fields[k].Span(), "%s is paid on each right-click use; add on_right_click, on_use_on_block or on_right_click_entity", k)
			return
		}
	}
	if damage > 0 {
		max := 0
		for _, k := range []string{"max_damage", "minecraft:max_damage"} {
			if n, ok := comps.Get(k).(jsonv.Number); ok {
				max, _ = strconv.Atoi(string(n))
			}
		}
		if max <= 0 {
			c.errorf("mms/custom-item", fields["damage_on_use"].Span(), "damage_on_use needs the item's durability: add max_damage to its components, such as components = { max_damage = 32 }")
			return
		}
		switch n := componentValue(comps, "max_stack_size").(type) {
		case nil:
			comps.Set("max_stack_size", jsonv.Number("1"))
		case jsonv.Number:
			if n != "1" {
				c.errorf("mms/custom-item", fields["components"].Span(), "an item with durability cannot stack: the game requires max_stack_size = 1 with max_damage")
				return
			}
		}
		out.Use.MaxDamage = max
	}
	out.Use.Count, out.Use.Damage = count, damage
}

func componentValue(comps *jsonv.Object, name string) jsonv.Value {
	if v := comps.Get(name); v != nil {
		return v
	}
	return comps.Get("minecraft:" + name)
}

// checkHookFunction checks an on_* field, which names a Function
// without parameters or holds an inline { ... } block that hoisting has
// compiled as one, and returns the function's id.
func (c *Checker) checkHookFunction(field string, e syntax.Expr, v jsonv.Value) string {
	s, _ := v.(jsonv.String)
	if s == "" || e == nil {
		return ""
	}
	for _, h := range c.Hoisted {
		if h.ID == string(s) && h.Fn != nil {
			return string(s) // an inline { ... } block, compiled as a Function
		}
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
	// Enchantments are the projectile hooks' enchantments by id, and
	// Spotted the lines they add to the namespace's ProjectileSpotted.
	Enchantments map[string]jsonv.Value
	Spotted      []string
	// Thrown is set when the item has on_projectile_land, so the
	// namespace needs the thrown potions' tick lines.
	Thrown bool
}

// ItemHookOutput builds the files for a custom item's hooks.
func ItemHookOutput(id string, it *CustomItemOut) ItemHookFiles {
	out := ItemHookFiles{Advancements: map[string]jsonv.Value{}, Functions: map[string][]string{}, Enchantments: map[string]jsonv.Value{}}
	if it == nil || len(it.Hooks) == 0 {
		return out
	}
	st, hooks := it.Stack, it.Hooks
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
		if h == nil || h.press || h.projectile {
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

	projectileHooks(id, string(base), hooks, &out)

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
	// The use cost is paid once per press: after on_right_click, which
	// runs on every press, or else after the block or entity hook.
	cost := ""
	if it.Use.Count > 0 || it.Use.Damage > 0 {
		cost = "function " + press + "/cost"
		out.Functions[press+"/cost"] = costLines(id, string(base), press+"/cost")
		for _, hand := range []string{"mainhand", "offhand"} {
			out.Functions[press+"/cost/"+hand] = costHandLines(hand, press+"/cost", it.Use)
			if it.Use.Damage > 0 {
				out.Functions[press+"/cost/"+hand+"/break"] = breakLines(hand, hooks["on_break"])
			}
		}
		if click != "" {
			body = append(body, cost)
			cost = ""
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
		blockFn := block
		if block != "" && cost != "" {
			blockFn = ray + "/block"
			out.Functions[blockFn] = []string{"function " + block, cost}
		}
		out.Functions[ray] = rayLines(obj, ray, blockFn, entity)
		if entity != "" {
			ls := []string{
				"tag @s add " + ItemHookTarget,
				"execute at @s as @p[tag=mms_clicker] run function " + entity,
			}
			if cost != "" {
				ls = append(ls, "execute as @p[tag=mms_clicker] run "+cost)
			}
			out.Functions[ray+"/entity"] = append(ls, "tag @s remove "+ItemHookTarget)
		}
	}
	conds := &jsonv.Object{}
	conds.Set("item", itemPred())
	add("press", "using_item", conds, body)
	return out
}

// projectileHooks writes the projectile hooks. An arrow keeps the bow
// or crossbow it was shot from as its weapon, and when it hits, the game
// runs the weapon's enchantment effects at the exact spot: hit_block as
// the arrow, post_attack as the entity it hurt. So each item with
// projectile hooks gets an enchantment that runs them, and the tick
// function adds it to the weapon of each new arrow that is the item or
// was shot from it. It is never on the item itself, so it shows in no
// tooltip and adds no glint. The tick function runs before the arrow
// first moves, so even a point-blank hit is caught.
func projectileHooks(id, base string, hooks map[string]string, out *ItemHookFiles) {
	if land := hooks["on_projectile_land"]; land != "" {
		landHook(id, land, out)
	}
	hitBlock, hitEntity := hooks["on_projectile_hit_block"], hooks["on_projectile_hit_entity"]
	if hitBlock == "" && hitEntity == "" {
		return
	}
	ench := ItemHookID(id, "projectile")
	run := func(fn string) *jsonv.Object {
		e := &jsonv.Object{}
		e.Set("type", jsonv.String("minecraft:run_function"))
		e.Set("function", jsonv.String(fn))
		return e
	}
	effects := &jsonv.Object{}
	if hitBlock != "" {
		eff := &jsonv.Object{}
		eff.Set("effect", run(hitBlock))
		effects.Set("minecraft:hit_block", jsonv.Array([]jsonv.Value{eff}))
	}
	if hitEntity != "" {
		eff := &jsonv.Object{}
		eff.Set("enchanted", jsonv.String("attacker"))
		eff.Set("affected", jsonv.String("victim"))
		eff.Set("effect", run(hitEntity))
		effects.Set("minecraft:post_attack", jsonv.Array([]jsonv.Value{eff}))
	}
	cost := func(n string) *jsonv.Object {
		c := &jsonv.Object{}
		c.Set("base", jsonv.Number(n))
		c.Set("per_level_above_first", jsonv.Number("0"))
		return c
	}
	desc := &jsonv.Object{}
	desc.Set("text", jsonv.String(id))
	e := &jsonv.Object{}
	e.Set("description", desc)
	var items []jsonv.Value
	for _, w := range ProjectileWeapons {
		items = append(items, jsonv.String(w))
	}
	e.Set("supported_items", jsonv.Array(items))
	e.Set("weight", jsonv.Number("1"))
	e.Set("max_level", jsonv.Number("1"))
	e.Set("min_cost", cost("1"))
	e.Set("max_cost", cost("1"))
	e.Set("anvil_cost", jsonv.Number("0"))
	e.Set("slots", jsonv.Array([]jsonv.Value{jsonv.String("any")}))
	e.Set("effects", effects)
	out.Enchantments[ench] = e

	// An arrow from a dispenser has no weapon, so it gets a plain bow.
	out.Functions[ench] = []string{
		"execute unless data entity @s weapon run data modify entity @s weapon set value {id:\"minecraft:bow\",count:1}",
		"data modify entity @s weapon.components.\"minecraft:enchantments\"." + quoteSNBT(ench) + " set value 1",
	}
	slot := "weapon"
	if slices.Contains(ProjectileAmmo, base) {
		slot = "item"
	}
	out.Spotted = append(out.Spotted, "execute if data entity @s "+slot+".components.\"minecraft:custom_data\"{"+ItemHookMarker+":"+quoteSNBT(id)+"} run function "+ench)
}

// landHook writes on_projectile_land. A thrown potion keeps no weapon,
// so no enchantment runs when it lands. Instead the tick function puts a
// marker on each new one, riding it. The game removes the potion when it
// lands and drops its passengers where they are, so a marker that is no
// longer riding anything marks the spot, give or take the potion's last
// tick of flight. The marker holds the function to run in its data.
func landHook(id, fn string, out *ItemHookFiles) {
	ns, _ := splitID(id)
	arm := ItemHookID(id, "land")
	out.Thrown = true
	out.Functions[arm] = []string{
		"tag @s add mms_new",
		"execute summon marker run function " + arm + "/mount",
		"tag @s remove mms_new",
	}
	out.Functions[arm+"/mount"] = []string{
		"tag @s add " + ProjectileLandTag(ns),
		"data modify entity @s data.mms_land set value " + quoteSNBT(fn),
		"ride @s mount @n[tag=mms_new]",
	}
	out.Spotted = append(out.Spotted, "execute if data entity @s Item.components.\"minecraft:custom_data\"{"+ItemHookMarker+":"+quoteSNBT(id)+"} run function "+arm)
}

// ProjectileSetup is what a namespace with projectile hooks needs, given
// the lines items added to its spotted function, and whether any item
// has on_projectile_land: tick lines, functions and predicates by id.
func ProjectileSetup(ns string, spotted []string, thrown bool) (tick []string, fns map[string][]string, preds map[string]jsonv.Value) {
	sort.Strings(spotted)
	tag := ProjectileTag(ns)
	spot := ProjectileSpotted(ns)
	fns = map[string][]string{spot: append([]string{"tag @s add " + tag}, spotted...)}
	preds = map[string]jsonv.Value{}
	tick = []string{"execute as @e[type=#minecraft:arrows,tag=!" + tag + "] run function " + spot}
	if !thrown {
		return tick, fns, preds
	}
	for _, t := range ProjectileThrown {
		tick = append(tick, "execute as @e[type="+strings.TrimPrefix(t, "minecraft:")+",tag=!"+tag+"] run function "+spot)
	}
	riding := ns + ":mms_gen/projectile/riding"
	vehicle := &jsonv.Object{}
	vehicle.Set("vehicle", &jsonv.Object{})
	pred := &jsonv.Object{}
	pred.Set("type", jsonv.String("minecraft:entity_properties"))
	pred.Set("entity", jsonv.String("this"))
	pred.Set("predicate", vehicle)
	preds[riding] = pred
	land := ns + ":mms_gen/projectile/land"
	tick = append(tick, "execute as @e[type=marker,tag="+ProjectileLandTag(ns)+",predicate=!"+riding+"] at @s run function "+land)
	fns[land] = []string{"function " + land + "/run with entity @s data", "kill @s"}
	fns[land+"/run"] = []string{"$function $(mms_land)"}
	return tick, fns, preds
}

// UseCost is what each use of a custom item with right-click hooks
// takes from it: items from the stack, and durability.
type UseCost struct {
	Count, Damage int
	// MaxDamage is the item's max_damage component, which turns Damage
	// into the fraction set_damage takes.
	MaxDamage int
}

// costLines pays the use cost from the hand that holds the item: the
// main hand when it does, otherwise the off hand.
func costLines(id, base, cost string) []string {
	pred := strings.TrimPrefix(base, "minecraft:") + "[custom_data~{" + ItemHookMarker + ":" + quoteSNBT(id) + "}]"
	return []string{
		"execute if items entity @s weapon.mainhand " + pred + " run return run function " + cost + "/mainhand",
		"execute if items entity @s weapon.offhand " + pred + " run function " + cost + "/offhand",
	}
}

// costHandLines pays the use cost from one hand.
func costHandLines(hand, cost string, u UseCost) []string {
	slot := "weapon." + hand
	var ls []string
	if u.Count > 0 {
		ls = append(ls, fmt.Sprintf("item modify entity @s %s {type:\"minecraft:set_count\",count:%d,add:true}", slot, -u.Count))
	}
	if u.Damage > 0 && u.MaxDamage > 0 {
		// set_damage works in fractions of max_damage and rounds the new
		// damage down, so half a point more keeps float error from
		// costing one point less.
		f := (float64(u.Damage) + 0.5) / float64(u.MaxDamage)
		if f > 1 {
			f = 1
		}
		ls = append(ls,
			fmt.Sprintf("item modify entity @s %s {type:\"minecraft:set_damage\",damage:%s,add:true}", slot, strconv.FormatFloat(-f, 'f', -1, 64)),
			"execute if items entity @s "+slot+" *[damage~{durability:{max:0}}] run function "+cost+"/"+hand+"/break")
	}
	return ls
}

// breakLines breaks an item whose durability a use took to 0, as the
// game does: it is gone, with the break sound, and on_break runs.
func breakLines(hand, onBreak string) []string {
	ls := []string{
		"item replace entity @s weapon." + hand + " with air",
		"playsound entity.item.break player @a",
	}
	if onBreak != "" {
		ls = append(ls, "function "+onBreak)
	}
	return ls
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
