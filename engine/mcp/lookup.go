package mcp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/minecraftmetascript/mms/engine/prelude"
	"github.com/minecraftmetascript/mms/engine/project"
	"github.com/minecraftmetascript/mms/engine/sema"
	"github.com/minecraftmetascript/mms/engine/spec"
)

// lookupCategories are the parts of the reference lookup can be kept to.
var lookupCategories = []string{"kind", "helper", "template", "blueprint", "type", "rule"}

// entry is one thing the reference describes.
type entry struct {
	cat, name string
	// line is the entry in one line, for a list of matches.
	line string
	// full is the whole entry.
	full string
	// text is what a search reads beyond the name, lower case.
	text string
}

// lazyReference is the language reference as entries, built on first
// use: it is the same data `mms reference` prints.
type lazyReference struct {
	once    sync.Once
	entries []entry
}

func (s *Server) reference() []entry {
	s.ref.once.Do(func() { s.ref.entries = buildEntries(s.pack) })
	return s.ref.entries
}

func firstSentence(doc string) string {
	doc = strings.Join(strings.Fields(doc), " ")
	if i := strings.Index(doc, ". "); i >= 0 {
		doc = doc[:i+1]
	}
	if len(doc) > 160 {
		doc = doc[:157] + "…"
	}
	return doc
}

func indent(text, by string) string {
	return by + strings.ReplaceAll(strings.TrimRight(text, "\n"), "\n", "\n"+by)
}

func buildEntries(p *spec.Pack) []entry {
	ref := sema.BuildReference(p, sema.NewKinds(p))
	var out []entry
	add := func(cat, name, line, full, more string) {
		out = append(out, entry{cat: cat, name: name, line: line, full: strings.TrimRight(full, "\n") + "\n", text: strings.ToLower(line + " " + more)})
	}
	for _, k := range ref.Kinds {
		where := "a kind MMS adds; it is not a file of its own"
		if k.Folder != "" {
			root := "data"
			if k.Resources {
				root = "assets"
			}
			where = fmt.Sprintf("writes %s/<namespace>/%s/<name>.json", root, k.Folder)
		}
		var b strings.Builder
		fmt.Fprintf(&b, "kind %s: %s\n", k.Name, where)
		if k.Doc != "" {
			b.WriteString(indent(k.Doc, "  ") + "\n")
		}
		add("kind", k.Name, strings.TrimSpace(fmt.Sprintf("kind %s: %s. %s", k.Name, where, firstSentence(k.Doc))), b.String(), k.Folder+" "+k.Doc)
	}
	for _, g := range ref.Groups {
		for _, h := range g.Helpers {
			makes := "an object"
			if g.Dispatcher != "" {
				makes = "a " + g.Dispatcher
			}
			var b strings.Builder
			fmt.Fprintf(&b, "helper %s (builds %s)\n", h.Signature, makes)
			if h.Doc != "" {
				b.WriteString(indent(h.Doc, "  ") + "\n")
			}
			writeParams(&b, "parameters", h.Params)
			if h.Example != "" {
				fmt.Fprintf(&b, "  example:\n%s\n", indent(h.Example, "    "))
			}
			line := strings.TrimSpace(fmt.Sprintf("helper %s: builds %s. %s", h.Signature, makes, firstSentence(h.Doc)))
			add("helper", h.Name, line, b.String(), g.Dispatcher+" "+h.Case+" "+h.Doc)
		}
	}
	for _, d := range prelude.Reference() {
		cat := strings.ToLower(d.Kind)
		var b strings.Builder
		fmt.Fprintf(&b, "%s %s\n", d.Kind, d.Signature)
		if d.Doc != "" {
			b.WriteString(indent(d.Doc, "  ") + "\n")
		}
		if d.As != "" {
			fmt.Fprintf(&b, "  its name stands for: %s\n", d.As)
		}
		if len(d.Fields) > 0 {
			label := "fields"
			if d.Kind == "Template" {
				label = "parameters"
			}
			fmt.Fprintf(&b, "  %s:\n", label)
			for _, f := range d.Fields {
				t := f.Type
				if t == "" {
					t = "any"
				}
				fmt.Fprintf(&b, "    %s: %s%s%s\n", f.Name, t, optional(f.Optional, f.Default), docTail(f.Doc))
			}
		}
		for _, u := range d.Uses {
			fmt.Fprintf(&b, "  where %s goes%s\n", u.Type, docTail(u.Doc))
		}
		add(cat, d.Name, strings.TrimSpace(fmt.Sprintf("%s %s. %s", d.Kind, d.Signature, firstSentence(d.Doc))), b.String(), d.Doc)
	}
	for _, t := range ref.Types {
		var def strings.Builder
		for _, part := range t.Def {
			def.WriteString(part.Text)
		}
		var b strings.Builder
		fmt.Fprintf(&b, "type %s", t.Name)
		if def.Len() > 0 {
			fmt.Fprintf(&b, " = %s", def.String())
		}
		b.WriteByte('\n')
		var vals []string
		for _, v := range t.Values {
			vals = append(vals, v.Value)
			fmt.Fprintf(&b, "  %s%s\n", v.Value, docTail(v.Doc))
		}
		writeParams(&b, "fields", t.Fields)
		line := fmt.Sprintf("type %s", t.Name)
		switch {
		case def.Len() > 0:
			line += " = " + firstSentence(def.String())
		case len(vals) > 0:
			line += ": one of " + firstSentence(strings.Join(vals, ", "))
		case len(t.Fields) > 0:
			line += fmt.Sprintf(": an object with %s", count(len(t.Fields), "field"))
		}
		add("type", t.Name, line, b.String(), t.ID)
	}
	for _, r := range project.Rules {
		sev := project.SeverityName(r.Severity)
		full := fmt.Sprintf("lint rule %s (default: %s)\n%s\n  Set its severity in mms.json: \"lint\": {%q: \"off\"}\n", r.ID, sev, indent(r.Doc, "  "), r.ID)
		add("rule", r.ID, fmt.Sprintf("lint rule %s (%s): %s", r.ID, sev, firstSentence(r.Doc)), full, r.Doc)
	}
	return out
}

func optional(opt bool, def string) string {
	switch {
	case def != "":
		return " (default " + def + ")"
	case opt:
		return " (optional)"
	}
	return ""
}

func docTail(doc string) string {
	if doc == "" {
		return ""
	}
	return ". " + strings.Join(strings.Fields(doc), " ")
}

func writeParams(b *strings.Builder, label string, ps []sema.RefParam) {
	if len(ps) == 0 {
		return
	}
	fmt.Fprintf(b, "  %s:\n", label)
	for _, p := range ps {
		fmt.Fprintf(b, "    %s: %s%s%s\n", p.Name, p.Type, optional(p.Optional, p.Default), docTail(p.Doc))
	}
}

func (s *Server) lookup(_ context.Context, args json.RawMessage) (string, error) {
	var a struct {
		Query string `json:"query"`
		In    string `json:"in"`
		Limit int    `json:"limit"`
	}
	if err := parse(args, &a); err != nil {
		return "", err
	}
	q := strings.ToLower(strings.TrimSpace(a.Query))
	if q == "" {
		return "", errors.New("query is empty")
	}
	if a.In != "" && !contains(lookupCategories, a.In) {
		return "", fmt.Errorf("in must be one of %s", strings.Join(lookupCategories, ", "))
	}
	if a.Limit <= 0 {
		a.Limit = 30
	}
	// Rank: the name itself, a name that starts with the query, a name
	// that contains it, then a description that does.
	var ranked [4][]entry
	for _, e := range s.reference() {
		if a.In != "" && e.cat != a.In {
			continue
		}
		name := strings.ToLower(e.name)
		switch {
		case name == q:
			ranked[0] = append(ranked[0], e)
		case strings.HasPrefix(name, q):
			ranked[1] = append(ranked[1], e)
		case strings.Contains(name, q):
			ranked[2] = append(ranked[2], e)
		case strings.Contains(e.text, q):
			ranked[3] = append(ranked[3], e)
		}
	}
	rest := append(append(ranked[1], ranked[2]...), ranked[3]...)
	var b strings.Builder
	exact := ranked[0]
	if len(exact) == 0 && len(rest) <= 3 {
		exact, rest = rest, nil
	}
	if len(exact) == 0 && len(rest) == 0 {
		return fmt.Sprintf("Nothing in the reference matches %q. Try a shorter word, or the docs tool's search for prose. Game ids such as blocks and items are in the vanilla tool.\n", a.Query), nil
	}
	for _, e := range exact {
		b.WriteString(e.full + "\n")
	}
	if len(rest) > 0 {
		if len(exact) > 0 {
			b.WriteString("Also matching:\n")
		}
		shown := rest
		if len(shown) > a.Limit {
			shown = shown[:a.Limit]
		}
		for _, e := range shown {
			fmt.Fprintf(&b, "- %s\n", e.line)
		}
		if len(rest) > len(shown) {
			fmt.Fprintf(&b, "… and %d more; narrow the query, or raise limit\n", len(rest)-len(shown))
		}
		if len(exact) == 0 {
			b.WriteString("\nLook up one by its exact name for its parameters and an example.\n")
		}
	}
	return b.String(), nil
}

func contains(list []string, s string) bool {
	for _, x := range list {
		if x == s {
			return true
		}
	}
	return false
}

func (s *Server) vanilla(_ context.Context, args json.RawMessage) (string, error) {
	var a struct {
		Registry string `json:"registry"`
		Query    string `json:"query"`
		File     string `json:"file"`
		Limit    int    `json:"limit"`
	}
	if err := parse(args, &a); err != nil {
		return "", err
	}
	if a.Limit <= 0 {
		a.Limit = 100
	}
	van := s.pack.Vanilla()
	var b strings.Builder
	if a.File != "" {
		return s.vanillaFile(van, a.File)
	}
	bare := func(id string) string { return strings.TrimPrefix(strings.TrimPrefix(id, "#"), "minecraft:") }
	q := strings.ToLower(bare(strings.TrimSpace(a.Query)))
	if a.Registry == "" && q == "" {
		fmt.Fprintf(&b, "Minecraft %s. Registries (name: number of ids):\n", s.pack.Version)
		for _, name := range sortedKeys(s.pack.Registries) {
			fmt.Fprintf(&b, "  %s: %d\n", name, len(s.pack.Registries[name]))
		}
		if len(van) > 0 {
			b.WriteString("\nFolders of vanilla data pack files (name: number of files), read with file as folder/id:\n")
			for _, name := range sortedKeys(van) {
				fmt.Fprintf(&b, "  %s: %d\n", name, len(van[name]))
			}
		}
		return b.String(), nil
	}
	if a.Registry != "" {
		reg := strings.Trim(bare(a.Registry), "/")
		ids, ok := s.pack.Registries[reg]
		if !ok {
			var near []string
			for _, name := range sortedKeys(s.pack.Registries) {
				if strings.Contains(name, reg) || strings.Contains(reg, name) {
					near = append(near, name)
				}
			}
			msg := fmt.Sprintf("no registry %q", a.Registry)
			if len(near) > 0 {
				msg += "; did you mean " + strings.Join(near, ", ") + "?"
			}
			return "", errors.New(msg + " Call vanilla with no arguments for the list.")
		}
		var hit []string
		for _, id := range ids {
			if q == "" || strings.Contains(id, q) {
				hit = append(hit, id)
			}
		}
		sort.Strings(hit)
		fmt.Fprintf(&b, "%s: %d of %s match\n", reg, len(hit), count(len(ids), "id"))
		writeIDs(&b, hit, a.Limit, "")
		if reg == "block" {
			if st := s.pack.Blocks[q]; st != nil && len(st.Properties) > 0 {
				fmt.Fprintf(&b, "\nBlock states of minecraft:%s:\n", q)
				for _, prop := range sortedKeys(st.Properties) {
					fmt.Fprintf(&b, "  %s: %s (default %s)\n", prop, strings.Join(st.Properties[prop], ", "), st.Defaults[prop])
				}
			}
		}
		if _, ok := van[reg]; ok && q != "" && contains(ids, q) {
			fmt.Fprintf(&b, "\nIts vanilla file: call vanilla with file %q\n", reg+"/"+q)
		}
		return b.String(), nil
	}
	total := 0
	for _, reg := range sortedKeys(s.pack.Registries) {
		var hit []string
		for _, id := range s.pack.Registries[reg] {
			if strings.Contains(id, q) {
				hit = append(hit, id)
			}
		}
		if len(hit) == 0 {
			continue
		}
		sort.Strings(hit)
		total += len(hit)
		fmt.Fprintf(&b, "%s (%d):\n", reg, len(hit))
		writeIDs(&b, hit, 12, "  ")
	}
	if total == 0 {
		return fmt.Sprintf("No vanilla id contains %q.\n", a.Query), nil
	}
	b.WriteString("\nGive registry as well for the whole list of one registry.\n")
	return b.String(), nil
}

func writeIDs(b *strings.Builder, ids []string, limit int, pad string) {
	shown := ids
	if len(shown) > limit {
		shown = shown[:limit]
	}
	for _, id := range shown {
		fmt.Fprintf(b, "%s  minecraft:%s\n", pad, id)
	}
	if len(ids) > len(shown) {
		fmt.Fprintf(b, "%s  … and %d more\n", pad, len(ids)-len(shown))
	}
}

func (s *Server) vanillaFile(van spec.Vanilla, file string) (string, error) {
	rest := strings.TrimPrefix(file, spec.VanillaPrefix)
	rest = strings.TrimPrefix(strings.TrimPrefix(rest, "/"), "data/minecraft/")
	rest = strings.TrimSuffix(strings.ReplaceAll(rest, "minecraft:", ""), ".json")
	rest = strings.ReplaceAll(rest, "#", "")
	if text, ok := van.File(spec.VanillaPrefix + "data/minecraft/" + rest + ".json"); ok {
		return text, nil
	}
	// A resource pack file: assets/models/block/stone (a folder as the
	// listing names it) or assets/minecraft/models/block/stone.
	if asset, ok := strings.CutPrefix(rest, "assets/"); ok {
		asset = strings.TrimPrefix(asset, "minecraft/")
		if text, ok := van.File(spec.VanillaPrefix + "assets/minecraft/" + asset + ".json"); ok {
			return text, nil
		}
	}
	// Say what is near: the files of the folder whose ids share the
	// last part of the path.
	last := rest[strings.LastIndex(rest, "/")+1:]
	var near []string
	for _, folder := range sortedKeys(van) {
		if !strings.HasPrefix(rest, folder+"/") && !strings.Contains(folder, strings.SplitN(rest, "/", 2)[0]) {
			continue
		}
		for id := range van[folder] {
			if last != "" && strings.Contains(id, last) {
				near = append(near, folder+"/"+id)
			}
		}
	}
	sort.Strings(near)
	if len(near) > 20 {
		near = near[:20]
	}
	msg := fmt.Sprintf("no vanilla file %q", file)
	if len(near) > 0 {
		msg += "; near it:\n  " + strings.Join(near, "\n  ")
	} else {
		msg += ". Files are named folder/id, such as loot_table/chests/simple_dungeon; call vanilla with no arguments for the folders."
	}
	return "", errors.New(msg)
}

func sortedKeys[V any](m map[string]V) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}
