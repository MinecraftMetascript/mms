package sema

import (
	"sort"
	"strings"

	"github.com/minecraftmetascript/mms/engine/spec"
)

// Reference describes the language for generated documentation: every
// declaration kind, and every helper grouped by what it builds.
type Reference struct {
	Version string     `json:"version"`
	Kinds   []RefKind  `json:"kinds"`
	Groups  []RefGroup `json:"groups"`
}

// RefKind is a declaration kind. Folder is the data pack folder of a
// resource kind, "" for the kinds MMS adds.
type RefKind struct {
	Name   string `json:"name"`
	Folder string `json:"folder,omitempty"`
	Doc    string `json:"doc,omitempty"`
}

// RefGroup is the helpers of one dispatcher, such as
// minecraft:density_function, or of plain structs (Dispatcher "").
type RefGroup struct {
	Dispatcher string      `json:"dispatcher"`
	Helpers    []RefHelper `json:"helpers"`
}

// RefHelper is one helper. Case is the type it makes, "" for a struct
// helper.
type RefHelper struct {
	Name      string     `json:"name"`
	Case      string     `json:"case,omitempty"`
	Signature string     `json:"signature"`
	Params    []RefParam `json:"params,omitempty"`
}

type RefParam struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Optional bool   `json:"optional,omitempty"`
	Doc      string `json:"doc,omitempty"`
}

var mmsKindDocs = map[string]string{
	"Namespace": "Sets the namespace for the declarations after it in the file.",
	"Function":  "A function: commands, variables and control flow, compiled to .mcfunction files.",
	"Objective": "A scoreboard objective, created by the generated load function.",
	"Global":    "A value for the whole pack, set when the pack first loads.",
	"OnLoad":    "Code that runs when the pack loads, from #minecraft:load.",
	"OnTick":    "Code that runs every tick, from #minecraft:tick.",
}

// BuildReference walks every type reachable from the pack's resources and
// lists the helpers that can build each one.
func BuildReference(p *spec.Pack, ks *Kinds) Reference {
	c := NewChecker(p, ks, &Table{}, "minecraft")
	ref := Reference{Version: p.Version}
	for _, name := range ks.Names() {
		k, _ := ks.Lookup(name)
		rk := RefKind{Name: name, Folder: k.Folder()}
		if k.MMS != "" {
			rk.Folder = ""
			rk.Doc = mmsKindDocs[name]
		} else if k.TagOf != "" {
			rk.Doc = "A tag of " + k.TagOf + " entries. Other tags of the same registry follow `Tag` the same way."
		}
		ref.Kinds = append(ref.Kinds, rk)
	}

	groups := map[string]map[string]RefHelper{}
	seen := map[*spec.Type]bool{}
	add := func(group string, cl *Callable) {
		if groups[group] == nil {
			groups[group] = map[string]RefHelper{}
		}
		if _, ok := groups[group][cl.Name]; ok {
			return
		}
		h := RefHelper{Name: cl.Name, Signature: c.signature(cl)}
		if cl.Case != "" {
			h.Case = "minecraft:" + strings.TrimPrefix(cl.Case, "minecraft:")
		}
		_, params, docs := c.Signature(&Info{Callable: cl})
		for i, p := range params {
			name, typ, _ := strings.Cut(p, ": ")
			rp := RefParam{Name: strings.TrimSuffix(name, "?"), Type: typ, Optional: strings.HasSuffix(name, "?"), Doc: docs[i]}
			h.Params = append(h.Params, rp)
		}
		groups[group][cl.Name] = h
	}

	var walk func(v tview, depth int)
	walkFields := func(fs []field, depth int) {
		for _, f := range fs {
			walk(f.Type, depth+1)
			if f.KeyType != nil {
				walk(*f.KeyType, depth+1)
			}
		}
	}
	walk = func(v tview, depth int) {
		v = c.norm(v)
		if v.T == nil || depth > 40 || seen[v.T] {
			return
		}
		seen[v.T] = true
		switch v.T.Kind {
		case "union":
			for _, m := range v.T.Members {
				walk(tview{T: m, Env: v.Env}, depth+1)
			}
		case "list":
			walk(tview{T: v.T.Item, Env: v.Env}, depth+1)
		case "tuple":
			for _, it := range v.T.Items {
				walk(tview{T: it, Env: v.Env}, depth+1)
			}
		case "struct":
			fs, disp := c.fields(v, "")
			for _, name := range c.callableNames(v) {
				if cl := c.callableByName(name, v); cl != nil {
					group := ""
					if disp != nil {
						group = disp.Dispatcher
					}
					add(group, cl)
				}
			}
			walkFields(fs, depth)
			if disp != nil {
				for _, key := range c.dispatchKeys(disp) {
					cfs, _ := c.fields(v, key)
					walkFields(cfs, depth)
				}
			}
		}
	}
	for _, r := range p.ResourceKinds() {
		walk(tview{T: p.Resources[r]}, 0)
	}

	names := make([]string, 0, len(groups))
	for g := range groups {
		names = append(names, g)
	}
	sort.Strings(names)
	for _, g := range names {
		rg := RefGroup{Dispatcher: g}
		for _, h := range groups[g] {
			rg.Helpers = append(rg.Helpers, h)
		}
		sort.Slice(rg.Helpers, func(i, j int) bool { return rg.Helpers[i].Name < rg.Helpers[j].Name })
		ref.Groups = append(ref.Groups, rg)
	}
	return ref
}
