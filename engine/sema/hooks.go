package sema

import (
	"strings"

	"github.com/minecraftmetascript/mms/engine/fn"
	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// Event hooks: declarations whose body runs when something happens to a
// player. Each compiles to a generated function under mms_gen/ and the
// wiring that calls it.
//
//	OnAdvancement name = Trigger(...) { ... }  an advancement with one
//	                                           criterion, a reward function
//	                                           that runs the body and
//	                                           revokes it, so it fires again
//	OnFirstJoin { ... }                        once per player, from a
//	                                           tick advancement never revoked
//	OnRespawn { ... }                          after a player respawns, from
//	                                           a deathCount objective

// HookKinds are the event hook kinds that, like OnLoad and OnTick, have
// an unnamed body collected into one generated function per namespace.
var HookKinds = map[string]string{
	"OnLoad":      "load",
	"OnTick":      "tick",
	"OnFirstJoin": "first_join",
	"OnRespawn":   "respawn",
}

// IsCodeKind reports whether a kind's body is function code.
func IsCodeKind(k Kind) bool {
	return k.MMS == "Function" || k.MMS == "OnAdvancement" || HookKinds[k.MMS] != ""
}

// AdvancementFunction names the reward function of an OnAdvancement.
func AdvancementFunction(id string) string {
	ns, path, _ := strings.Cut(id, ":")
	return ns + ":mms_gen/advancement/" + path
}

// RespawnObjective is the deathCount objective OnRespawn watches.
func RespawnObjective(ns string) string { return "mms_respawn." + ns }

// OnAdvancement checks `OnAdvancement name = Trigger(...) { ... }` and
// returns the advancement JSON and the compiled reward function. id is
// the advancement's "ns:path".
func (c *Checker) OnAdvancement(d *syntax.Decl, id string, src []rune) (jsonv.Value, *fn.Func) {
	const usage = "OnAdvancement needs a trigger and a body: OnAdvancement placed = PlacedBlock(...) { ... }"
	if d.Name == nil {
		c.errorf("mms/missing-value", d.Kind.Sp, "%s", usage)
		return nil, nil
	}
	root := tview{T: c.Pack.Resources["advancement"]}
	if root.T != nil && root.T.Kind == "reference" {
		root.Name = root.T.Path
	}
	criterion := func(e syntax.Expr) *jsonv.Object {
		sp := e.Span()
		o, _ := c.check(mkObject(sp, "criteria", mkObject(sp, "event", e)), root).(*jsonv.Object)
		return o
	}
	t, ok := d.Value.(*syntax.Trailing)
	if !ok {
		if d.Value != nil {
			criterion(d.Value) // for completion while the trigger is typed
		}
		sp := d.Kind.Sp
		if d.Value != nil {
			sp = d.Value.Span()
		} else if d.Body != nil {
			sp = d.Body.Sp
		}
		c.errorf("mms/missing-value", sp, "%s", usage)
		return nil, nil
	}
	adv := criterion(t.X)
	fid := AdvancementFunction(id)
	f := c.CompileFunction(&syntax.Decl{Kind: d.Kind, Body: t.Block, Doc: d.Doc, Sp: d.Sp}, fid, src)
	if adv == nil || f == nil {
		return nil, nil
	}
	rewards := &jsonv.Object{}
	rewards.Set("function", jsonv.String(fid))
	adv.Set("rewards", rewards)
	// Revoke first, so the advancement can fire again even if the body
	// stops early.
	f.Body.Ops = append([]fn.Op{fn.Cmd{Text: "advancement revoke @s only " + id}}, f.Body.Ops...)
	return adv, f
}

// FirstJoinAdvancement is the advancement that runs a namespace's
// OnFirstJoin bodies: a tick trigger, granted on a player's first tick
// and never revoked.
func FirstJoinAdvancement(ns string) jsonv.Value {
	crit := &jsonv.Object{}
	crit.Set("trigger", jsonv.String("minecraft:tick"))
	criteria := &jsonv.Object{}
	criteria.Set("join", crit)
	rewards := &jsonv.Object{}
	rewards.Set("function", jsonv.String(ns+":mms_gen/first_join"))
	o := &jsonv.Object{}
	o.Set("criteria", criteria)
	o.Set("rewards", rewards)
	return o
}
