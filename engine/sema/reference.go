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
	// Types explains the named types that signatures link to and that no
	// helper group covers, such as enums and id-or-object unions.
	Types []*RefType `json:"types"`
}

// RefKind is a declaration kind. Folder is the folder under
// data/<namespace>/ of a data pack kind, or under assets/<namespace>/
// when Resources is set; "" for the kinds MMS adds.
type RefKind struct {
	Name      string `json:"name"`
	Folder    string `json:"folder,omitempty"`
	Resources bool   `json:"resources,omitempty"`
	Doc       string `json:"doc,omitempty"`
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
	// Doc and Example are hand-written, from helpdocs/; Doc is Markdown.
	Doc     string `json:"doc,omitempty"`
	Example string `json:"example,omitempty"`
}

// RefParam is a helper parameter or a field of a type entry. Type is the
// readable type as plain text; TypeParts is the same with links.
type RefParam struct {
	Name      string     `json:"name"`
	Type      string     `json:"type"`
	TypeParts []TypePart `json:"typeParts,omitempty"`
	Optional  bool       `json:"optional,omitempty"`
	// Default is the value MMS writes when the parameter is left out.
	Default string `json:"default,omitempty"`
	Doc     string `json:"doc,omitempty"`
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
		if k.Assets != "" {
			rk.Folder, rk.Resources = k.Assets, true
		}
		if k.MMS != "" {
			rk.Folder = ""
			rk.Doc = mmsKindDocs[name]
		} else if k.TagOf != "" {
			rk.Doc = "A tag of " + k.TagOf + " entries. Other tags of the same registry follow `Tag` the same way."
		}
		ref.Kinds = append(ref.Kinds, rk)
	}

	groups := map[string]map[string]*Callable{}
	seen := map[*spec.Type]bool{}
	add := func(group string, cl *Callable) {
		if groups[group] == nil {
			groups[group] = map[string]*Callable{}
		}
		if _, ok := groups[group][cl.Name]; !ok {
			groups[group][cl.Name] = cl
		}
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
	for _, r := range append(p.ResourceKinds(), p.AssetKinds()...) {
		walk(tview{T: p.Resources[r]}, 0)
	}

	rt := &refTypes{c: c, types: map[string]*RefType{}, helper: func(group, name string) bool {
		hs, ok := groups[group]
		if name == "" {
			return ok
		}
		_, ok = hs[name]
		return ok
	}}
	names := make([]string, 0, len(groups))
	for g := range groups {
		names = append(names, g)
	}
	sort.Strings(names)
	for _, g := range names {
		rg := RefGroup{Dispatcher: g}
		for _, cl := range groups[g] {
			rg.Helpers = append(rg.Helpers, c.refHelper(rt, cl))
		}
		sort.Slice(rg.Helpers, func(i, j int) bool { return rg.Helpers[i].Name < rg.Helpers[j].Name })
		ref.Groups = append(ref.Groups, rg)
	}
	rt.define()
	for _, t := range rt.types {
		ref.Types = append(ref.Types, t)
	}
	sort.Slice(ref.Types, func(i, j int) bool {
		a, b := ref.Types[i], ref.Types[j]
		if a.Name != b.Name {
			return a.Name < b.Name
		}
		return a.ID < b.ID
	})
	return ref
}

func (c *Checker) refHelper(rt *refTypes, cl *Callable) RefHelper {
	h := RefHelper{Name: cl.Name, Signature: c.signature(cl)}
	if d := HelpDoc(cl); d != nil {
		h.Doc, h.Example = d.Summary, d.Example
	}
	if cl.Case != "" {
		h.Case = "minecraft:" + strings.TrimPrefix(cl.Case, "minecraft:")
	}
	for _, f := range c.positional(cl) {
		opt := f.Optional
		if cl.Ov != nil {
			if _, ok := cl.Ov.defaults[f.Key]; ok {
				opt = true
			}
		}
		parts := rt.format(f.Type, 0)
		h.Params = append(h.Params, RefParam{Name: f.Key, Type: plainParts(parts), TypeParts: parts, Optional: opt, Default: defaultOf(cl, f.Key), Doc: f.Desc})
	}
	return h
}
