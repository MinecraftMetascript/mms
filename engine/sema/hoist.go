package sema

import (
	"strconv"

	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// Inline values where an id is required. Some fields only take an id,
// such as a surface rule's noise_threshold `noise` or a biome source's
// biome, and some take an inline value that the game cannot use, such as
// a density function's `noise`: the game seeds a noise from its id and
// fails to build the noise router for an inline one ("Cannot create
// noise without key"), and so does Deepslate. An inline value written
// there is checked as its own declaration of the registry's kind, written
// to its own file under mms_gen/, and the field gets its id:
//
//	DensityFunction hills = Noise({ base_octave = -6, amplitude_modifiers = [1] }, 1, 1)
//
// writes worldgen/noise/mms_gen/hills/noise.json and refers to it as
// <ns>:mms_gen/hills/noise.

// Hoisted is an inline value written out as its own resource.
type Hoisted struct {
	Kind  Kind
	ID    string // "ns:mms_gen/<owner>/<name>"
	Value jsonv.Value
	Sp    syntax.Span // the inline value's span
}

// keyRequired lists registries whose inline form the schema allows but
// the game rejects at runtime, so an inline value is hoisted anyway.
var keyRequired = map[string]bool{
	"worldgen/noise": true,
}

// hoistKind returns the declaration kind of an id's registry when an
// inline value can be hoisted into it.
func (c *Checker) hoistKind(id idInfo) (Kind, bool) {
	if !id.Present || id.Registry == "" || id.Slot || id.Tags == "required" || id.Tags == "implicit" || c.Kinds == nil {
		return Kind{}, false
	}
	k, ok := c.Kinds.ForFolder(c.registryName(id.Registry))
	if !ok || !k.IsResource() || k.TagOf != "" || k.Assets != "" {
		return Kind{}, false
	}
	return k, true
}

// hoistFits is fits for an id position: an inline value hoisted there.
func (c *Checker) hoistFits(e syntax.Expr, id idInfo) bool {
	k, ok := c.hoistKind(id)
	return ok && c.hoistable(e, k)
}

// hoistable reports whether e is written as an inline value (not an id)
// that could be a value of k.
func (c *Checker) hoistable(e syntax.Expr, k Kind) bool {
	switch x := e.(type) {
	case *syntax.Paren:
		return c.hoistable(x.X, k)
	case *syntax.String, *syntax.Ident, *syntax.ResLoc, *syntax.Tag, *syntax.Bad, nil:
		return false
	}
	return c.fits(e, c.kindRoot(k))
}

func (c *Checker) kindRoot(k Kind) tview {
	root := tview{T: c.Pack.Resources[k.Resource]}
	if root.T != nil && root.T.Kind == "reference" {
		root.Name = root.T.Path
	}
	return root
}

// hoist checks e as a value of the id's kind and returns the id of the
// file it is written to; ok is false when e is not an inline value of it.
func (c *Checker) hoist(e syntax.Expr, id idInfo) (jsonv.Value, bool) {
	k, ok := c.hoistKind(id)
	if !ok || !c.hoistable(e, k) {
		return nil, false
	}
	v := c.check(e, c.kindRoot(k))
	if c.trial {
		return jsonv.String(c.NS + ":mms_gen/inline"), true
	}
	if v == nil {
		return nil, true
	}
	hid := c.hoistID(k)
	c.Hoisted = append(c.Hoisted, Hoisted{Kind: k, ID: hid, Value: v, Sp: e.Span()})
	in := c.info(e)
	in.ID = hid
	if in.Help == "" && in.Callable == nil {
		in.Help = "```mms\n" + k.Name + " " + hid + "\n```\n\nAn inline " + k.Name + " where an id is needed: it is written to its own file and referenced by this id."
	}
	return jsonv.String(hid), true
}

// hoistID names the next hoisted value of kind k: under mms_gen/, by the
// declaration that holds it and the kind's folder, numbered from the
// second one on.
func (c *Checker) hoistID(k Kind) string {
	ns, owner := c.NS, "inline"
	if c.Owner != "" {
		ns, owner = splitID(c.Owner)
		if ns == "" {
			ns = c.NS
		}
	}
	base := ns + ":mms_gen/" + owner + "/" + lastSeg(k.Folder())
	if c.hoistN == nil {
		c.hoistN = map[string]int{}
	}
	c.hoistN[base]++
	if n := c.hoistN[base]; n > 1 {
		return base + "_" + strconv.Itoa(n)
	}
	return base
}

// keyRequiredID returns the id member of a union whose registry needs an
// id even though an inline value is allowed.
func (c *Checker) keyRequiredID(v tview) (idInfo, bool) {
	for _, m := range v.T.Members {
		if m.Kind != "string" {
			continue
		}
		if id := idOf(m); id.Present && keyRequired[c.registryName(id.Registry)] {
			return id, true
		}
	}
	return idInfo{}, false
}

// preferInline drops union members that e fits only by being hoisted,
// when another member takes it inline: an inline density function stays
// inline where a DensityFunctionRef allows one.
func (c *Checker) preferInline(e syntax.Expr, cands []tview) []tview {
	if len(cands) < 2 {
		return cands
	}
	var keep []tview
	for _, m := range cands {
		if m.T != nil && m.T.Kind == "string" && c.hoistFits(e, idOf(m.T)) {
			if _, ok := e.(*syntax.Call); !ok || !c.assetFits(e, idOf(m.T)) {
				continue
			}
		}
		keep = append(keep, m)
	}
	if len(keep) == 0 {
		return cands
	}
	return keep
}
