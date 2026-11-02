package project

import (
	"fmt"
	"io/fs"
	"sort"
	"strings"
	"sync"

	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/nbt"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// Jigsaw structures: a template pool lists structure pieces, and each
// piece holds jigsaw blocks. A jigsaw block names the pool it draws the
// next piece from, and the target: the name of a jigsaw block in that
// piece to join onto. Nothing checks this until the game generates the
// structure, where a missing pool or a target no piece has just stops
// the structure growing, without an error. These rules read the jigsaw
// blocks of the structure files Structure("...") includes and check them
// against the project's TemplatePool declarations.

// Jigsaw is one jigsaw block of a structure file.
type Jigsaw struct {
	Pos                [3]int
	Pool, Name, Target string
}

// structureCache holds parsed structure files by path, size and
// modification time, so analysing again after an edit does not reparse.
var structureCache = struct {
	sync.Mutex
	m map[string]structureFile
}{m: map[string]structureFile{}}

type structureFile struct {
	jigsaws []Jigsaw
	ok      bool // false when the file could not be read as a structure
}

// StructureJigsaws returns the jigsaw blocks of a structure file in the
// project; ok is false when it cannot be read as one.
func StructureJigsaws(fsys fs.FS, src string) ([]Jigsaw, bool) { return readJigsaws(fsys, src) }

// readJigsaws returns the jigsaw blocks of a structure file in fsys; ok
// is false when it cannot be read or parsed.
func readJigsaws(fsys fs.FS, src string) ([]Jigsaw, bool) {
	if fsys == nil {
		return nil, false
	}
	st, err := fs.Stat(fsys, src)
	if err != nil || st.IsDir() {
		return nil, false
	}
	key := fmt.Sprintf("%p|%s|%d|%d", fsys, src, st.Size(), st.ModTime().UnixNano())
	structureCache.Lock()
	f, hit := structureCache.m[key]
	structureCache.Unlock()
	if hit {
		return f.jigsaws, f.ok
	}
	if data, err := fs.ReadFile(fsys, src); err == nil {
		if root, err := nbt.Decode(data); err == nil {
			f = structureFile{jigsaws: Jigsaws(root), ok: true}
		}
	}
	structureCache.Lock()
	if len(structureCache.m) > 1024 {
		structureCache.m = map[string]structureFile{}
	}
	structureCache.m[key] = f
	structureCache.Unlock()
	return f.jigsaws, f.ok
}

// Jigsaws lists the jigsaw blocks of a decoded structure file, in the
// order the file lists its blocks.
func Jigsaws(root map[string]any) []Jigsaw {
	var palette []any
	if p, ok := root["palette"].([]any); ok {
		palette = p
	} else if ps, ok := root["palettes"].([]any); ok && len(ps) > 0 {
		palette, _ = ps[0].([]any)
	}
	blocks, _ := root["blocks"].([]any)
	var out []Jigsaw
	for _, b := range blocks {
		block, _ := b.(map[string]any)
		data, _ := block["nbt"].(map[string]any)
		if data == nil {
			continue
		}
		isJigsaw := data["id"] == "minecraft:jigsaw"
		if s, ok := block["state"].(int32); ok && int(s) < len(palette) && s >= 0 {
			if st, ok := palette[s].(map[string]any); ok && st["Name"] == "minecraft:jigsaw" {
				isJigsaw = true
			}
		}
		if !isJigsaw {
			continue
		}
		j := Jigsaw{Pool: nbtID(data["pool"]), Name: nbtID(data["name"]), Target: nbtID(data["target"])}
		if pos, ok := block["pos"].([]any); ok && len(pos) == 3 {
			for i, v := range pos {
				if n, ok := v.(int32); ok {
					j.Pos[i] = int(n)
				}
			}
		}
		out = append(out, j)
	}
	return out
}

// nbtID reads an id string, adding the minecraft namespace; absent or
// empty is minecraft:empty, the game's default.
func nbtID(v any) string {
	s, _ := v.(string)
	if s == "" {
		return "minecraft:empty"
	}
	if !strings.Contains(s, ":") {
		return "minecraft:" + s
	}
	return s
}

// structureUse is a structure file the project includes, at its first
// Structure("...") call.
type structureUse struct {
	assetUse
	jigsaws []Jigsaw
}

// jigsawModel is what the jigsaw rules check against.
type jigsawModel struct {
	uses []*structureUse
	// byID maps a structure id the project provides to its file.
	byID map[string]*structureUse
	// pools maps each declared template pool to the structure ids of its
	// pieces; known is false when a piece is not a project structure file
	// (vanilla, another pack, or a feature), so its jigsaws are unknown.
	pools map[string]*poolPieces
	// declared namespaces: a pool in one of these must be declared here.
	declared map[string]bool
}

type poolPieces struct {
	pieces []string
	known  bool
}

func (r *Result) jigsawModel() *jigsawModel {
	m := &jigsawModel{byID: map[string]*structureUse{}, pools: map[string]*poolPieces{}, declared: map[string]bool{r.cfg.Namespace: true}}
	seen := map[string]bool{}
	for _, u := range r.assets {
		if u.Helper != "Structure" || seen[u.Src] {
			continue
		}
		seen[u.Src] = true
		js, ok := readJigsaws(r.files, u.Src)
		if !ok {
			continue
		}
		s := &structureUse{assetUse: u, jigsaws: js}
		m.uses = append(m.uses, s)
		m.byID[u.ID] = s
	}
	for _, d := range r.Order {
		m.declared[d.NS] = true
		if d.Kind.Folder() != "worldgen/template_pool" {
			continue
		}
		p := &poolPieces{known: !hasErrors(d.Diags)}
		if o, ok := d.Value.(*jsonv.Object); ok {
			if els, ok := o.Get("elements").(jsonv.Array); ok {
				for _, e := range els {
					if eo, ok := e.(*jsonv.Object); ok {
						m.addPieces(p, eo.Get("element"))
					}
				}
			}
		}
		m.pools[d.ID()] = p
	}
	return m
}

// addPieces adds a pool element's structure ids; a list element adds
// each of its elements, and a feature or empty element adds none.
func (m *jigsawModel) addPieces(p *poolPieces, v jsonv.Value) {
	el, ok := v.(*jsonv.Object)
	if !ok {
		p.known = false
		return
	}
	kind, _ := el.Get("element_type").(jsonv.String)
	switch strings.TrimPrefix(string(kind), "minecraft:") {
	case "single_pool_element", "legacy_single_pool_element":
		loc, _ := el.Get("location").(jsonv.String)
		id := string(loc)
		if !strings.Contains(id, ":") {
			id = "minecraft:" + id
		}
		p.pieces = append(p.pieces, id)
		if m.byID[id] == nil {
			p.known = false
		}
	case "list_pool_element":
		if els, ok := el.Get("elements").(jsonv.Array); ok {
			for _, e := range els {
				m.addPieces(p, e)
			}
		}
	case "feature_pool_element", "empty_pool_element":
	default:
		p.known = false
	}
}

func (l *lintRun) jigsaws() *jigsawModel {
	if l.r.jigsaw == nil {
		l.r.jigsaw = l.r.jigsawModel()
	}
	return l.r.jigsaw
}

func posText(p [3]int) string { return fmt.Sprintf("%d %d %d", p[0], p[1], p[2]) }

// lintJigsawPool reports a jigsaw block whose pool is in one of the
// project's namespaces but that no TemplatePool declares.
func lintJigsawPool(l *lintRun) {
	m := l.jigsaws()
	for _, s := range m.uses {
		reported := map[string]bool{}
		for _, j := range s.jigsaws {
			ns, name, _ := strings.Cut(j.Pool, ":")
			if j.Pool == "minecraft:empty" || ns == "minecraft" || !m.declared[ns] || m.pools[j.Pool] != nil || reported[j.Pool] {
				continue
			}
			reported[j.Pool] = true
			d := syntax.Diagnostic{
				Code: "lint/" + l.rule.ID, Severity: l.sev, Sp: s.sp,
				Message: fmt.Sprintf("'%s' has a jigsaw block at %s that draws from the pool '%s', which no TemplatePool declares; the structure stops growing there", s.Src, posText(j.Pos), j.Pool),
			}
			if fix, ok := l.r.newPoolFix(s, ns, name); ok {
				d.Fixes = []syntax.Fix{fix}
			}
			l.out[s.decl.Unit.Path] = append(l.out[s.decl.Unit.Path], d)
		}
	}
}

// newPoolFix adds an empty TemplatePool at the end of the file of the
// Structure call, when that file ends in the pool's namespace.
func (r *Result) newPoolFix(s *structureUse, ns, name string) (syntax.Fix, bool) {
	var last *Decl
	for _, d := range r.Order {
		if d.Unit == s.decl.Unit {
			last = d
		}
	}
	if last == nil || last.NS != ns {
		return syntax.Fix{}, false
	}
	src := s.decl.Unit.File.Src
	text := fmt.Sprintf("\nTemplatePool %s = {\n  fallback = minecraft:empty\n  elements = []\n}\n", name)
	if len(src) > 0 && src[len(src)-1] != '\n' {
		text = "\n" + text
	}
	end := len(src)
	return syntax.Fix{Title: fmt.Sprintf("Declare TemplatePool %s", name), Edits: []syntax.TextEdit{{Sp: syntax.Span{Start: end, End: end}, Text: text}}}, true
}

// lintJigsawTarget reports a jigsaw block whose target no piece of its
// pool has: no jigsaw block in those pieces carries that name, so no
// piece can join it. Only pools whose pieces are all project structure
// files are checked.
func lintJigsawTarget(l *lintRun) {
	m := l.jigsaws()
	for _, s := range m.uses {
		reported := map[string]bool{}
		for _, j := range s.jigsaws {
			p := m.pools[j.Pool]
			if p == nil || !p.known || j.Target == "minecraft:empty" || reported[j.Pool+"|"+j.Target] {
				continue
			}
			names := map[string]bool{}
			for _, id := range p.pieces {
				for _, o := range m.byID[id].jigsaws {
					names[o.Name] = true
				}
			}
			if names[j.Target] {
				continue
			}
			reported[j.Pool+"|"+j.Target] = true
			var msg string
			if len(p.pieces) == 0 {
				msg = fmt.Sprintf("'%s' has a jigsaw block at %s that draws from the pool '%s', which has no structure pieces yet, so nothing joins there", s.Src, posText(j.Pos), j.Pool)
			} else {
				var have []string
				for n := range names {
					if n != "minecraft:empty" {
						have = append(have, n)
					}
				}
				sort.Strings(have)
				msg = fmt.Sprintf("'%s' has a jigsaw block at %s that targets '%s' in the pool '%s', but no piece of that pool has a jigsaw block named '%s', so the pieces never connect", s.Src, posText(j.Pos), j.Target, j.Pool, j.Target)
				if len(have) > 0 {
					msg += fmt.Sprintf(" (the pieces' jigsaw names: %s)", strings.Join(have, ", "))
				}
			}
			l.out[s.decl.Unit.Path] = append(l.out[s.decl.Unit.Path], syntax.Diagnostic{Code: "lint/" + l.rule.ID, Severity: l.sev, Sp: s.sp, Message: msg})
		}
	}
}
