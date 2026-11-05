package project

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"slices"
	"sort"
	"strings"
	"sync"

	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/nbt"
	"github.com/minecraftmetascript/mms/engine/sema"
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
	// Orientation is the block's orientation state, such as north_up:
	// the way it faces, then the way its top points. Empty when unknown.
	Orientation string
	// Joint is rollable or aligned: whether a piece joined to an up or
	// down facing jigsaw block may turn. The game defaults it to aligned
	// for sideways facing blocks and rollable for up or down ones.
	Joint string
	// FinalState is the block the jigsaw block turns into once the
	// structure is placed, such as minecraft:air.
	FinalState string
	// Label is the name of the item the block was placed from (its
	// item_name or custom_name), which the block entity keeps among its
	// components; empty when it has none.
	Label string
}

// structureCache holds parsed structure files by path, size and
// modification time, so analysing again after an edit does not reparse.
var structureCache = struct {
	sync.Mutex
	m map[string]structureFile
}{m: map[string]structureFile{}}

type structureFile struct {
	jigsaws []Jigsaw
	size    [3]int // x, y, z; zero when the file has no size
	ok      bool   // false when the file could not be read as a structure
}

// StructureJigsaws returns the jigsaw blocks of a structure file in the
// project; ok is false when it cannot be read as one.
func StructureJigsaws(fsys fs.FS, src string) ([]Jigsaw, bool) { return readJigsaws(fsys, src) }

// StructureSize returns a structure file's size in blocks (x, y, z); ok
// is false when it cannot be read as one or has no size.
func StructureSize(fsys fs.FS, src string) ([3]int, bool) {
	f := readStructure(fsys, src)
	return f.size, f.ok && f.size != [3]int{}
}

// readJigsaws returns the jigsaw blocks of a structure file in fsys; ok
// is false when it cannot be read or parsed.
func readJigsaws(fsys fs.FS, src string) ([]Jigsaw, bool) {
	f := readStructure(fsys, src)
	return f.jigsaws, f.ok
}

// readStructure reads a structure file in fsys, cached.
func readStructure(fsys fs.FS, src string) structureFile {
	if fsys == nil {
		return structureFile{}
	}
	st, err := fs.Stat(fsys, src)
	if err != nil || st.IsDir() {
		return structureFile{}
	}
	key := fmt.Sprintf("%p|%s|%d|%d", fsys, src, st.Size(), st.ModTime().UnixNano())
	structureCache.Lock()
	f, hit := structureCache.m[key]
	structureCache.Unlock()
	if hit {
		return f
	}
	if data, err := fs.ReadFile(fsys, src); err == nil {
		if root, err := nbt.Decode(data); err == nil {
			f = structureFile{jigsaws: Jigsaws(root), ok: true}
			if size, ok := root["size"].([]any); ok && len(size) == 3 {
				for i, v := range size {
					if n, ok := v.(int32); ok {
						f.size[i] = int(n)
					}
				}
			}
		}
	}
	structureCache.Lock()
	if len(structureCache.m) > 1024 {
		structureCache.m = map[string]structureFile{}
	}
	structureCache.m[key] = f
	structureCache.Unlock()
	return f
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
		orientation := ""
		if s, ok := block["state"].(int32); ok && int(s) < len(palette) && s >= 0 {
			// Palette entries are {Name, Properties} up to 1.21 and
			// {id, properties} in 26.x structure files.
			if st, ok := palette[s].(map[string]any); ok && (st["Name"] == "minecraft:jigsaw" || st["id"] == "minecraft:jigsaw") {
				isJigsaw = true
				props, ok := st["Properties"].(map[string]any)
				if !ok {
					props, _ = st["properties"].(map[string]any)
				}
				orientation, _ = props["orientation"].(string)
			}
		}
		if !isJigsaw {
			continue
		}
		j := Jigsaw{Pool: nbtID(data["pool"]), Name: nbtID(data["name"]), Target: nbtID(data["target"]), Orientation: orientation}
		j.Joint, _ = data["joint"].(string)
		j.FinalState, _ = data["final_state"].(string)
		if comps, ok := data["components"].(map[string]any); ok {
			for _, k := range []string{"minecraft:custom_name", "minecraft:item_name"} {
				if j.Label = textOf(comps[k]); j.Label != "" {
					break
				}
			}
		}
		if j.Joint == "" && orientation != "" {
			if strings.HasPrefix(orientation, "up_") || strings.HasPrefix(orientation, "down_") {
				j.Joint = "rollable"
			} else {
				j.Joint = "aligned"
			}
		}
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

// textOf reads the plain text of a text component in NBT: a string, a
// compound with text (and extra), or a list of those; older files hold
// it as JSON in a string.
func textOf(v any) string {
	switch t := v.(type) {
	case string:
		if strings.HasPrefix(t, "{") || strings.HasPrefix(t, "[") || strings.HasPrefix(t, "\"") {
			var j any
			if json.Unmarshal([]byte(t), &j) == nil {
				return textOf(j)
			}
		}
		return t
	case map[string]any:
		out := textOf(t["text"])
		if extra, ok := t["extra"].([]any); ok {
			out += textOf(extra)
		}
		return out
	case []any:
		var sb strings.Builder
		for _, e := range t {
			sb.WriteString(textOf(e))
		}
		return sb.String()
	}
	return ""
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
	size    [3]int // zero when the file has no size
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
	// terrain holds the pieces placed with terrain_matching projection,
	// whose height follows the ground rather than the jigsaw block.
	terrain map[string]bool
	// empty is set when the pool has an empty or feature element, so
	// placing nothing there can be what the author meant.
	empty bool
	// fallback is the pool tried when none of its pieces fit.
	fallback string
}

func (r *Result) jigsawModel() *jigsawModel {
	m := &jigsawModel{byID: map[string]*structureUse{}, pools: map[string]*poolPieces{}, declared: map[string]bool{r.cfg.Namespace: true}}
	seen := map[string]bool{}
	for _, u := range r.assets {
		if u.Helper != "Structure" || seen[u.Src] {
			continue
		}
		seen[u.Src] = true
		f := readStructure(r.files, u.Src)
		if !f.ok {
			continue
		}
		s := &structureUse{assetUse: u, jigsaws: f.jigsaws, size: f.size}
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
			if fb, ok := o.Get("fallback").(jsonv.String); ok {
				p.fallback = string(fb)
				if !strings.Contains(p.fallback, ":") {
					p.fallback = "minecraft:" + p.fallback
				}
			}
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
		if proj, _ := el.Get("projection").(jsonv.String); strings.TrimPrefix(string(proj), "minecraft:") == "terrain_matching" {
			if p.terrain == nil {
				p.terrain = map[string]bool{}
			}
			p.terrain[id] = true
		}
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
		p.empty = true
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

// front is the way a jigsaw block faces, from its orientation; empty
// when unknown.
func (j Jigsaw) front() string {
	f, _, ok := strings.Cut(j.Orientation, "_")
	if !ok {
		return ""
	}
	return f
}

// canFace reports whether jigsaw blocks facing a and b can ever join. The
// game joins two blocks that face opposite ways after turning the new
// piece about the y axis, so any two sideways blocks can join, but one
// facing up joins only one facing down. Tops never block a join: a turn
// lines them up when the joint is aligned.
func canFace(a, b string) bool {
	vertical := func(f string) bool { return f == "up" || f == "down" }
	if !vertical(a) && !vertical(b) {
		return true
	}
	return a == "up" && b == "down" || a == "down" && b == "up"
}

func facingText(f string) string {
	if f == "up" || f == "down" {
		return f
	}
	return "sideways (" + f + ")"
}

// lintJigsawFacing reports a jigsaw block whose target blocks in its pool
// all face a way it can never join: up joins only down, down only up,
// and sideways only sideways.
func lintJigsawFacing(l *lintRun) {
	m := l.jigsaws()
	for _, s := range m.uses {
		reported := map[string]bool{}
		for _, j := range s.jigsaws {
			p := m.pools[j.Pool]
			from := j.front()
			if p == nil || !p.known || from == "" || j.Target == "minecraft:empty" {
				continue
			}
			key := j.Pool + "|" + j.Target + "|" + from
			if reported[key] {
				continue
			}
			var faces []string
			ok, any := false, false
			for _, id := range p.pieces {
				u := m.byID[id]
				if u == nil {
					continue
				}
				for _, o := range u.jigsaws {
					if o.Name != j.Target {
						continue
					}
					any = true
					to := o.front()
					if to == "" || canFace(from, to) {
						ok = true
						break
					}
					if !slices.Contains(faces, to) {
						faces = append(faces, to)
					}
				}
			}
			if ok || !any {
				continue
			}
			reported[key] = true
			want := "sideways"
			switch from {
			case "up":
				want = "down"
			case "down":
				want = "up"
			}
			for i, f := range faces {
				faces[i] = facingText(f)
			}
			msg := fmt.Sprintf("'%s' has a jigsaw block at %s facing %s that targets '%s' in the pool '%s', but every jigsaw block named '%s' there faces %s; it can only join one facing %s, so the pieces never connect", s.Src, posText(j.Pos), facingText(from), j.Target, j.Pool, j.Target, strings.Join(faces, " or "), want)
			l.out[s.decl.Unit.Path] = append(l.out[s.decl.Unit.Path], syntax.Diagnostic{Code: "lint/" + l.rule.ID, Severity: l.sev, Sp: s.sp, Message: msg})
		}
	}
}

var dirs = map[string][3]int{"north": {0, 0, -1}, "south": {0, 0, 1}, "west": {-1, 0, 0}, "east": {1, 0, 0}, "up": {0, 1, 0}, "down": {0, -1, 0}}

// rotate turns a position clockwise about the y axis r quarter turns, as
// the game turns a piece about its origin.
func rotate(p [3]int, r int) [3]int {
	switch r & 3 {
	case 1:
		return [3]int{-p[2], p[1], p[0]}
	case 2:
		return [3]int{-p[0], p[1], -p[2]}
	case 3:
		return [3]int{p[2], p[1], -p[0]}
	}
	return p
}

func add(a, b [3]int) [3]int { return [3]int{a[0] + b[0], a[1] + b[1], a[2] + b[2]} }
func sub(a, b [3]int) [3]int { return [3]int{a[0] - b[0], a[1] - b[1], a[2] - b[2]} }
func neg(a [3]int) [3]int    { return [3]int{-a[0], -a[1], -a[2]} }

// top is the way a jigsaw block's top points, from its orientation.
func (j Jigsaw) top() string {
	_, t, _ := strings.Cut(j.Orientation, "_")
	return t
}

// overflow is how far a child piece, joined at parent jigsaw j through
// its jigsaw o with r quarter turns, sticks out of the parent's box, per
// side (west, east, down, up, north, south). ok is false when that turn
// cannot join them. The y axis is skipped when skipY.
func overflow(parent [3]int, j Jigsaw, child [3]int, o Jigsaw, r int, skipY bool) (out [6]int, ok bool) {
	jf, of := dirs[j.front()], dirs[o.front()]
	if rotate(of, r) != neg(jf) {
		return out, false
	}
	if j.front() == "up" || j.front() == "down" {
		if j.Joint == "aligned" && rotate(dirs[o.top()], r) != dirs[j.top()] {
			return out, false
		}
	}
	at := sub(add(j.Pos, jf), rotate(o.Pos, r))
	a, b := add(rotate([3]int{0, 0, 0}, r), at), add(rotate(sub(child, [3]int{1, 1, 1}), r), at)
	for i := range 3 {
		if skipY && i == 1 {
			continue
		}
		lo, hi := min(a[i], b[i]), max(a[i], b[i])
		out[i*2] = max(0, -lo)
		out[i*2+1] = max(0, hi-(parent[i]-1))
	}
	return out, true
}

// lintJigsawFit reports pieces that can never join an internal jigsaw
// block (one whose front block is inside its own piece) because, turned
// any way that lines the jigsaw blocks up, they stick out of the piece:
// the game only places an internal piece wholly inside its parent.
func lintJigsawFit(l *lintRun) {
	m := l.jigsaws()
	sides := [6]string{"west", "east", "down", "up", "north", "south"}
	for _, s := range m.uses {
		if s.size == [3]int{} {
			continue
		}
		for _, j := range s.jigsaws {
			f := j.front()
			p := m.pools[j.Pool]
			if f == "" || p == nil || !p.known || j.Target == "minecraft:empty" {
				continue
			}
			front := add(j.Pos, dirs[f])
			inside := true
			for i := range 3 {
				inside = inside && front[i] >= 0 && front[i] < s.size[i]
			}
			if !inside {
				continue
			}
			var bad []string
			fits := false
			seen := map[string]bool{}
			for _, id := range p.pieces {
				u := m.byID[id]
				if u == nil || seen[id] || u.size == [3]int{} {
					continue
				}
				seen[id] = true
				best, joins := [6]int{}, false
				for _, o := range u.jigsaws {
					if o.Name != j.Target || o.front() == "" || !canFace(f, o.front()) {
						continue
					}
					for r := range 4 {
						out, ok := overflow(s.size, j, u.size, o, r, p.terrain[id])
						if !ok {
							continue
						}
						if !joins || sum(out) < sum(best) {
							best = out
						}
						joins = true
					}
				}
				if !joins {
					continue
				}
				if sum(best) == 0 {
					fits = true
					continue
				}
				var over []string
				for i, n := range best {
					if n > 0 {
						over = append(over, fmt.Sprintf("%d %s", n, sides[i]))
					}
				}
				bad = append(bad, fmt.Sprintf("'%s' (%d×%d×%d) sticks out %s", id, u.size[0], u.size[1], u.size[2], strings.Join(over, ", ")))
			}
			if len(bad) == 0 {
				continue
			}
			end := "so it is never placed there"
			if len(bad) > 1 {
				end = "so they are never placed there"
			}
			if !fits {
				end += "; nothing fits, so nothing joins there"
			}
			edge := map[string]string{"up": "top", "down": "bottom"}[f]
			if edge == "" {
				edge = f + " edge"
			}
			msg := fmt.Sprintf("'%s' has an internal jigsaw block at %s facing %s (it is not on the piece's %s, so the block in front of it is inside the %d×%d×%d piece), so a piece from '%s' that joins it must fit inside this piece: %s, %s. If it should join outward, move it to the %s", s.Src, posText(j.Pos), f, edge, s.size[0], s.size[1], s.size[2], j.Pool, strings.Join(bad, "; "), end, edge)
			l.out[s.decl.Unit.Path] = append(l.out[s.decl.Unit.Path], syntax.Diagnostic{Code: "lint/" + l.rule.ID, Severity: l.sev, Sp: s.sp, Message: msg})
		}
	}
}

func sum(a [6]int) int { return a[0] + a[1] + a[2] + a[3] + a[4] + a[5] }

// box is a block box, min and max corners inclusive.
type box [2][3]int

func (a box) overlaps(b box) bool {
	for i := range 3 {
		if a[1][i] < b[0][i] || b[1][i] < a[0][i] {
			return false
		}
	}
	return true
}

// placed is a piece placed in the frame of the piece the check starts
// from: turned r quarter turns, then moved by off.
type placed struct {
	size [3]int
	r    int
	off  [3]int
}

func (p placed) pos(q [3]int) [3]int { return add(rotate(q, p.r), p.off) }
func (p placed) dir(f string) [3]int { return rotate(dirs[f], p.r) }
func (p placed) box() box {
	a, b := p.pos([3]int{}), p.pos(sub(p.size, [3]int{1, 1, 1}))
	return box{{min(a[0], b[0]), min(a[1], b[1]), min(a[2], b[2])}, {max(a[0], b[0]), max(a[1], b[1]), max(a[2], b[2])}}
}

// join places a child piece of the given size onto jigsaw j of piece
// at, joining through the child's jigsaw o turned r quarter turns; ok is
// false when that turn does not line the two blocks up.
func join(at placed, j Jigsaw, size [3]int, o Jigsaw, r int) (placed, bool) {
	jf := at.dir(j.front())
	if rotate(dirs[o.front()], r) != neg(jf) {
		return placed{}, false
	}
	if (j.front() == "up" || j.front() == "down") && j.Joint == "aligned" && rotate(dirs[o.top()], r) != at.dir(j.top()) {
		return placed{}, false
	}
	return placed{size: size, r: r, off: sub(add(at.pos(j.Pos), jf), rotate(o.Pos, r))}, true
}

// internal reports whether the block in front of jigsaw j is inside the
// piece, so what joins there is placed inside it.
func internal(size [3]int, j Jigsaw) bool {
	f := add(j.Pos, dirs[j.front()])
	for i := range 3 {
		if f[i] < 0 || f[i] >= size[i] {
			return false
		}
	}
	return true
}

// lintJigsawOverlap reports a jigsaw block that can never grow when its
// piece joins a given parent: every piece its pool (and the pool's
// fallback) could place there, turned any way that lines the blocks up,
// overlaps the parent or the piece itself, and the game never places a
// piece that overlaps one already placed. A junction narrower than the
// roads around it is the usual cause. Only certain cases are reported:
// every piece involved is a rigid project structure file with a known
// size and facing, and no pool involved has an empty or feature element.
func lintJigsawOverlap(l *lintRun) {
	m := l.jigsaws()
	terrain := map[string]bool{}
	for _, p := range m.pools {
		for id := range p.terrain {
			terrain[id] = true
		}
	}
	// candidates lists a pool's pieces and its fallback's, or ok false
	// when placing nothing there may be meant or something is unknown.
	candidates := func(pool string) (out []*structureUse, ok bool) {
		seen := map[string]bool{}
		for pool != "" && pool != "minecraft:empty" && !seen[pool] {
			seen[pool] = true
			p := m.pools[pool]
			if p == nil || !p.known || p.empty {
				return nil, false
			}
			for _, id := range p.pieces {
				u := m.byID[id]
				if u == nil || u.size == [3]int{} || terrain[id] {
					return nil, false
				}
				out = append(out, u)
			}
			pool = p.fallback
		}
		return out, len(out) > 0
	}
	known := func(u *structureUse) bool {
		for _, j := range u.jigsaws {
			if j.front() == "" {
				return false
			}
		}
		return true
	}
	for _, c := range m.uses {
		if c.size == [3]int{} || terrain[c.ID] || !known(c) {
			continue
		}
		reported := map[int]bool{}
		// Each parent jigsaw block that can draw this piece.
		for _, pa := range m.uses {
			if pa.size == [3]int{} || terrain[pa.ID] || !known(pa) {
				continue
			}
			for _, jp := range pa.jigsaws {
				if jp.Target == "minecraft:empty" || internal(pa.size, jp) {
					continue
				}
				pieces, ok := candidates(jp.Pool)
				if !ok || !slices.Contains(pieces, c) {
					continue
				}
				parent := placed{size: pa.size}
				for ci, jc := range c.jigsaws {
					if jc.Name != jp.Target {
						continue
					}
					for r := range 4 {
						child, ok := join(parent, jp, c.size, jc, r)
						if !ok || child.box().overlaps(parent.box()) {
							continue
						}
						for oi, jo := range c.jigsaws {
							if oi == ci || reported[oi] || jo.Target == "minecraft:empty" || internal(c.size, jo) {
								continue
							}
							grow, ok := candidates(jo.Pool)
							if !ok {
								continue
							}
							fits, tried := false, false
							for _, g := range grow {
								for _, jg := range g.jigsaws {
									if jg.Name != jo.Target {
										continue
									}
									for rg := range 4 {
										gp, ok := join(child, jo, g.size, jg, rg)
										if !ok {
											continue
										}
										tried = true
										if !gp.box().overlaps(parent.box()) && !gp.box().overlaps(child.box()) {
											fits = true
										}
									}
								}
							}
							if !tried || fits {
								continue
							}
							reported[oi] = true
							msg := fmt.Sprintf("When '%s' joins '%s' by its jigsaw block at %s (onto the one at %s), its jigsaw block at %s ('%s', drawing from '%s') can never grow: every piece that pool or its fallback could place there overlaps '%s' or '%s', however it is turned, and the game never places overlapping pieces. A junction narrower than the pieces around it is the usual cause", c.Src, pa.Src, posText(jc.Pos), posText(jp.Pos), posText(jo.Pos), jo.Name, jo.Pool, pa.Src, c.Src)
							l.out[c.decl.Unit.Path] = append(l.out[c.decl.Unit.Path], syntax.Diagnostic{Code: "lint/" + l.rule.ID, Severity: l.sev, Sp: c.sp, Message: msg})
						}
					}
				}
			}
		}
	}
}

// lintJigsawHole reports a jigsaw block whose name or target is in a
// namespace that declares JigsawHoles but is not one of them. Holes
// are opt-in: a namespace without any is not checked.
func lintJigsawHole(l *lintRun) {
	holes, nsHas := map[string]bool{}, map[string]bool{}
	for _, s := range l.r.Syms.All() {
		if s.Kind.MMS == sema.JigsawHoleKind {
			holes[s.ID()] = true
			nsHas[s.NS] = true
		}
	}
	if len(holes) == 0 {
		return
	}
	for _, s := range l.jigsaws().uses {
		reported := map[string]bool{}
		for _, j := range s.jigsaws {
			for _, f := range []struct{ field, id string }{{"name", j.Name}, {"target", j.Target}} {
				ns, _, _ := strings.Cut(f.id, ":")
				if f.id == "minecraft:empty" || !nsHas[ns] || holes[f.id] || reported[f.id] {
					continue
				}
				reported[f.id] = true
				d := syntax.Diagnostic{
					Code: "lint/" + l.rule.ID, Severity: l.sev, Sp: s.sp,
					Message: fmt.Sprintf("'%s' has a jigsaw block at %s whose %s is '%s', which is not a declared JigsawHole%s", s.Src, posText(j.Pos), f.field, f.id, nearHole(f.id, holes)),
				}
				if fix, ok := l.r.newHoleFix(s, f.id); ok {
					d.Fixes = []syntax.Fix{fix}
				}
				l.out[s.decl.Unit.Path] = append(l.out[s.decl.Unit.Path], d)
			}
		}
	}
}

// nearHole suggests the declared hole closest to id, for a typo.
func nearHole(id string, holes map[string]bool) string {
	best, bestD := "", 3
	for h := range holes {
		if d := editDistance(id, h); d < bestD || (d == bestD && best != "" && h < best) {
			best, bestD = h, d
		}
	}
	if best == "" {
		return ""
	}
	return fmt.Sprintf(" (did you mean '%s'?)", best)
}

func editDistance(a, b string) int {
	prev := make([]int, len(b)+1)
	for j := range prev {
		prev[j] = j
	}
	for i := 1; i <= len(a); i++ {
		cur := make([]int, len(b)+1)
		cur[0] = i
		for j := 1; j <= len(b); j++ {
			cost := 1
			if a[i-1] == b[j-1] {
				cost = 0
			}
			cur[j] = min(prev[j]+1, cur[j-1]+1, prev[j-1]+cost)
		}
		prev = cur
	}
	return prev[len(b)]
}

// newHoleFix adds `JigsawHole name` at the end of the file of the
// Structure call: the bare name when the file ends in the hole's
// namespace, else the full id.
func (r *Result) newHoleFix(s *structureUse, id string) (syntax.Fix, bool) {
	ns, name, _ := strings.Cut(id, ":")
	// The namespace in force at the end of the file: the last Namespace
	// declaration's, or the project's.
	end := r.cfg.Namespace
	for _, d := range s.decl.Unit.File.Decls {
		if d.Kind != nil && d.Kind.Name == "Namespace" && d.Name != nil {
			end = d.Name.Text
		}
	}
	decl := name
	if end != ns {
		decl = id
	}
	src := s.decl.Unit.File.Src
	text := fmt.Sprintf("\nJigsawHole %s\n", decl)
	if len(src) > 0 && src[len(src)-1] != '\n' {
		text = "\n" + text
	}
	at := len(src)
	return syntax.Fix{Title: fmt.Sprintf("Declare JigsawHole %s", decl), Edits: []syntax.TextEdit{{Sp: syntax.Span{Start: at, End: at}, Text: text}}}, true
}
