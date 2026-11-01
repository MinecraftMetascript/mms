package command

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/minecraftmetascript/mms/engine/spec"
)

// parserNames describes argument types in messages.
var parserNames = map[string]string{
	"brigadier:bool": "true or false", "brigadier:integer": "integer", "brigadier:float": "number",
	"brigadier:double": "number", "brigadier:string": "string",
	"block_pos": "x y z", "column_pos": "x z", "vec3": "x y z", "vec2": "x z", "rotation": "yaw pitch",
	"entity": "selector", "score_holder": "score holder", "game_profile": "player",
	"block_state": "block", "block_predicate": "block or #tag", "item_stack": "item", "item_predicate": "item or #tag",
	"resource_location": "id", "function": "function", "objective": "objective", "objective_criteria": "criterion",
	"team": "team", "nbt_path": "NBT path", "nbt_tag": "NBT", "nbt_compound_tag": "NBT compound",
	"component": "text", "style": "style", "message": "message", "int_range": "range", "float_range": "range",
	"time": "time", "gamemode": "game mode", "entity_anchor": "feet or eyes", "heightmap": "heightmap",
	"operation": "operation", "scoreboard_slot": "display slot", "swizzle": "axes", "particle": "particle",
	"dimension": "dimension", "uuid": "UUID", "hex_color": "hex colour", "item_slot": "slot",
	"resource": "id", "resource_key": "id", "resource_or_tag": "id or #tag", "resource_or_tag_key": "id or #tag",
	"resource_selector": "id pattern", "loot_table": "loot table", "loot_predicate": "predicate",
	"loot_modifier": "item modifier", "dialog": "dialog", "template_mirror": "mirror", "template_rotation": "rotation",
	"team_color": "colour", "slot_source": "slot source",
}

// Criteria are the single-word scoreboard criteria. Compound criteria
// such as `minecraft.mined:minecraft.stone` and `teamkill.red` are valid
// too, but are not listed.
var Criteria = []string{"dummy", "trigger", "deathCount", "playerKillCount", "totalKillCount", "health", "xp", "level", "food", "air", "armor"}

// Enumerated word arguments.
var enums = map[string][]string{
	"gamemode":          {"adventure", "creative", "spectator", "survival"},
	"entity_anchor":     {"eyes", "feet"},
	"heightmap":         {"motion_blocking", "motion_blocking_no_leaves", "ocean_floor", "ocean_floor_wg", "world_surface", "world_surface_wg"},
	"operation":         {"%=", "*=", "+=", "-=", "/=", "<", "=", "><", ">"},
	"template_mirror":   {"front_back", "left_right", "none"},
	"template_rotation": {"180", "clockwise_90", "counterclockwise_90", "none"},
	"team_color":        append(append([]string(nil), colors...), "reset"),
}

var colors = []string{"aqua", "black", "blue", "dark_aqua", "dark_blue", "dark_gray", "dark_green", "dark_purple", "dark_red", "gold", "gray", "green", "light_purple", "red", "white", "yellow"}

// Selector options, from the game's EntitySelectorOptions.
var selectorOptions = map[string]bool{
	"advancements": true, "distance": true, "dx": true, "dy": true, "dz": true, "gamemode": true,
	"level": true, "limit": true, "name": true, "nbt": true, "predicate": true, "scores": true,
	"sort": true, "tag": true, "team": true, "type": true, "x": true, "x_rotation": true,
	"y": true, "y_rotation": true, "z": true,
}

// parserRegistry maps id arguments to the registry their ids come from.
var parserRegistry = map[string]string{
	"function": "function", "dimension": "dimension", "loot_table": "loot_table",
	"loot_predicate": "predicate", "loot_modifier": "item_modifier", "dialog": "dialog",
	"block_state": "block", "block_predicate": "block", "item_stack": "item", "item_predicate": "item",
	"particle": "particle_type",
}

// arg parses one argument at pos and returns where it ends.
func (s *state) arg(name string, n *spec.CommandNode, pos int) (int, bool) {
	p := strings.TrimPrefix(n.Parser, "minecraft:")
	if p == "message" || p == "brigadier:string" && n.Properties["type"] == "greedy" {
		return len(s.src), true
	}
	if end, ok := s.interpolated(pos); ok {
		return end, true
	}
	switch p {
	case "brigadier:bool":
		end := s.word(pos)
		if w := string(s.src[pos:end]); w != "true" && w != "false" {
			return 0, s.fail(pos, end, "expected true or false")
		}
		return end, true
	case "brigadier:integer", "brigadier:float", "brigadier:double":
		return s.number(pos, p == "brigadier:integer", n.Properties)
	case "brigadier:string":
		switch n.Properties["type"] {
		case "greedy":
			return len(s.src), true
		case "phrase":
			if s.src[pos] == '"' || s.src[pos] == '\'' {
				return s.quoted(pos)
			}
		}
		end := s.unquoted(pos)
		if end == pos {
			return 0, s.fail(pos, s.word(pos), "expected a word")
		}
		return end, true
	case "message":
		return len(s.src), true
	case "entity", "game_profile":
		return s.entity(pos, n.Properties)
	case "score_holder":
		if s.src[pos] == '@' {
			return s.entity(pos, n.Properties)
		}
		return s.word(pos), true
	case "block_pos", "vec3":
		return s.coords(pos, 3, p == "block_pos")
	case "column_pos", "vec2", "rotation":
		return s.coords(pos, 2, p == "column_pos")
	case "resource_location":
		return s.resloc(pos, false, "")
	case "objective_criteria":
		return s.word(pos), true
	case "function", "dimension":
		return s.resloc(pos, p == "function", parserRegistry[p])
	case "loot_table", "loot_predicate", "loot_modifier", "dialog":
		if c := s.src[pos]; c == '{' || c == '[' {
			return s.balanced(pos)
		}
		return s.resloc(pos, false, parserRegistry[p])
	case "resource", "resource_key", "resource_or_tag", "resource_or_tag_key":
		reg, _ := n.Properties["registry"].(string)
		return s.resloc(pos, strings.HasPrefix(p, "resource_or_tag"), strings.TrimPrefix(reg, "minecraft:"))
	case "resource_selector":
		return s.balanced(pos)
	case "block_state", "block_predicate", "item_stack", "item_predicate", "particle":
		if p == "item_predicate" && s.src[pos] == '*' {
			return s.balanced(pos)
		}
		end, ok := s.resloc(pos, p == "block_predicate" || p == "item_predicate", parserRegistry[p])
		if !ok {
			return 0, false
		}
		if end < len(s.src) && s.src[end] != ' ' {
			return s.balanced(end)
		}
		return end, true
	case "objective", "team":
		end := s.unquoted(pos)
		if end == pos {
			return 0, s.fail(pos, s.word(pos), "expected "+parserNames[p]+" name")
		}
		if p == "objective" {
			s.ids = append(s.ids, ID{Start: pos, End: end, Registry: "objective", Value: string(s.src[pos:end])})
		}
		return end, true
	case "int_range", "float_range":
		return s.rangeArg(pos, p == "int_range")
	case "time":
		end := s.word(pos)
		w := strings.TrimRight(string(s.src[pos:end]), "dst")
		if _, err := strconv.ParseFloat(w, 64); err != nil {
			return 0, s.fail(pos, end, "expected a time such as 20, 5s or 1d")
		}
		return end, true
	case "scoreboard_slot":
		end := s.word(pos)
		w := string(s.src[pos:end])
		if w == "list" || w == "sidebar" || w == "below_name" || w == "belowName" {
			return end, true
		}
		if c, ok := strings.CutPrefix(w, "sidebar.team."); ok && contains(colors, c) {
			return end, true
		}
		return 0, s.fail(pos, end, "expected list, sidebar, below_name or sidebar.team.<colour>")
	case "swizzle":
		end := s.word(pos)
		w := string(s.src[pos:end])
		seen := map[rune]bool{}
		for _, r := range w {
			if !strings.ContainsRune("xyz", r) || seen[r] {
				return 0, s.fail(pos, end, "expected axes such as xz or xyz")
			}
			seen[r] = true
		}
		return end, w != "" || s.fail(pos, end, "expected axes such as xz or xyz")
	case "hex_color":
		end := s.word(pos)
		w := string(s.src[pos:end])
		if _, err := strconv.ParseUint(w, 16, 32); err != nil || len(w) != 6 {
			return 0, s.fail(pos, end, "expected a hex colour such as ff8800")
		}
		return end, true
	}
	if vals, ok := enums[p]; ok {
		end := s.word(pos)
		w := string(s.src[pos:end])
		if !contains(vals, w) {
			return 0, s.fail(pos, end, fmt.Sprintf("unknown %s '%s'; expected %s", parserNames[p], w, strings.Join(vals, ", ")))
		}
		return end, true
	}
	// NBT, text components, paths, providers, slots and anything newer:
	// one token with balanced brackets and quotes.
	return s.balanced(pos)
}

func contains(xs []string, x string) bool {
	for _, y := range xs {
		if y == x {
			return true
		}
	}
	return false
}

// interpolated accepts a token holding an MMS interpolation: its value is
// only known once lowered.
func (s *state) interpolated(pos int) (int, bool) {
	end, ok := s.balancedEnd(pos)
	if !ok || !strings.Contains(string(s.src[pos:end]), interp) {
		return 0, false
	}
	return end, true
}

func isUnquoted(r rune) bool {
	return r >= '0' && r <= '9' || r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r == '_' || r == '-' || r == '.' || r == '+'
}

func (s *state) unquoted(pos int) int {
	for pos < len(s.src) && isUnquoted(s.src[pos]) {
		pos++
	}
	return pos
}

func (s *state) quoted(pos int) (int, bool) {
	q := s.src[pos]
	for i := pos + 1; i < len(s.src); i++ {
		switch s.src[i] {
		case '\\':
			i++
		case q:
			return i + 1, true
		}
	}
	return 0, s.fail(pos, len(s.src), "this string is not closed")
}

// balancedEnd finds the end of a token: the next space outside brackets
// and quotes.
func (s *state) balancedEnd(pos int) (int, bool) {
	var stack []rune
	for i := pos; i < len(s.src); i++ {
		switch c := s.src[i]; c {
		case ' ':
			if len(stack) == 0 {
				return i, true
			}
		case '"', '\'':
			end, ok := s.quotedQuiet(i)
			if !ok {
				return len(s.src), false
			}
			i = end - 1
		case '{', '[', '(':
			stack = append(stack, c)
		case '}', ']', ')':
			open := map[rune]rune{'}': '{', ']': '[', ')': '('}[c]
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return i, false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(s.src), len(stack) == 0
}

func (s *state) quotedQuiet(pos int) (int, bool) {
	q := s.src[pos]
	for i := pos + 1; i < len(s.src); i++ {
		switch s.src[i] {
		case '\\':
			i++
		case q:
			return i + 1, true
		}
	}
	return 0, false
}

func (s *state) balanced(pos int) (int, bool) {
	end, ok := s.balancedEnd(pos)
	if !ok {
		if end < len(s.src) {
			return 0, s.fail(end, end+1, "unexpected '"+string(s.src[end])+"'")
		}
		return 0, s.fail(pos, len(s.src), "a bracket or quote is not closed")
	}
	if end == pos {
		return 0, s.fail(pos, pos, "expected a value")
	}
	return end, true
}

func (s *state) number(pos int, integer bool, props map[string]any) (int, bool) {
	end := pos
	for end < len(s.src) && (s.src[end] >= '0' && s.src[end] <= '9' || s.src[end] == '-' || s.src[end] == '.' || s.src[end] == 'e' || s.src[end] == 'E' || s.src[end] == '+') {
		end++
	}
	w := string(s.src[pos:end])
	var v float64
	var err error
	if integer {
		var i int64
		i, err = strconv.ParseInt(w, 10, 32)
		v = float64(i)
	} else {
		v, err = strconv.ParseFloat(w, 64)
	}
	if err != nil || w == "" {
		kind := "a number"
		if integer {
			kind = "an integer"
		}
		return 0, s.fail(pos, s.word(pos), "expected "+kind)
	}
	if lo, ok := props["min"].(float64); ok && v < lo {
		return 0, s.fail(pos, end, fmt.Sprintf("must be at least %v", fmtNum(lo)))
	}
	if hi, ok := props["max"].(float64); ok && v > hi {
		return 0, s.fail(pos, end, fmt.Sprintf("must be at most %v", fmtNum(hi)))
	}
	return end, true
}

func fmtNum(f float64) string { return strconv.FormatFloat(f, 'f', -1, 64) }

func (s *state) rangeArg(pos int, integer bool) (int, bool) {
	end := s.word(pos)
	w := string(s.src[pos:end])
	lo, hi, found := strings.Cut(w, "..")
	parse := func(x string) bool {
		if x == "" {
			return true
		}
		if integer {
			_, err := strconv.ParseInt(x, 10, 32)
			return err == nil
		}
		_, err := strconv.ParseFloat(x, 64)
		return err == nil
	}
	if w == "" || w == ".." || !parse(lo) || (found && !parse(hi)) {
		return 0, s.fail(pos, end, "expected a range such as 5, 1..10, ..3 or 2..")
	}
	return end, true
}

// coords parses n coordinates: absolute, ~relative or ^local.
func (s *state) coords(pos, n int, blocks bool) (int, bool) {
	i := pos
	local, mixed := 0, -1
	for k := 0; k < n; k++ {
		if k > 0 {
			if i >= len(s.src) || s.src[i] != ' ' {
				return 0, s.fail(i, i, fmt.Sprintf("expected %d coordinates", n))
			}
			i++
		}
		if i >= len(s.src) {
			return 0, s.fail(i, i, fmt.Sprintf("expected %d coordinates", n))
		}
		end := s.word(i)
		w := string(s.src[i:end])
		rest := w
		if (strings.HasPrefix(w, "^") != (local > 0)) && k > 0 && mixed < 0 {
			mixed = i
		}
		switch {
		case strings.HasPrefix(w, "^"):
			local++
			rest = w[1:]
		case strings.HasPrefix(w, "~"):
			rest = w[1:]
		default:
			if blocks {
				if _, err := strconv.ParseInt(w, 10, 32); err != nil {
					return 0, s.fail(i, end, "expected a whole block coordinate, ~ or ^")
				}
			}
		}
		if rest != "" {
			if _, err := strconv.ParseFloat(rest, 64); err != nil {
				return 0, s.fail(i, end, "expected a coordinate such as 10, ~ or ~-2.5")
			}
		}
		i = end
	}
	if local != 0 && local != n {
		return 0, s.fail(mixed, s.word(mixed), "local coordinates (^) cannot be mixed with others")
	}
	return i, true
}

// resloc parses a resource location, optionally a #tag, and checks
// vanilla ids against the registry when the pack lists it.
func (s *state) resloc(pos int, tags bool, registry string) (int, bool) {
	i := pos
	tag := false
	if tags && i < len(s.src) && s.src[i] == '#' {
		tag = true
		i++
	}
	start := i
	colon := -1
	for i < len(s.src) {
		c := s.src[i]
		switch {
		case c >= 'a' && c <= 'z', c >= '0' && c <= '9', c == '_', c == '-', c == '.':
		case c == '/' && colon >= 0, c == '/' && colon < 0:
		case c == ':' && colon < 0:
			colon = i
		default:
			goto done
		}
		i++
	}
done:
	id := string(s.src[start:i])
	if id == "" || strings.HasSuffix(id, ":") {
		return 0, s.fail(pos, s.word(pos), "expected an id such as minecraft:stone")
	}
	if colon >= 0 && strings.ContainsRune(string(s.src[start:colon]), '/') {
		return 0, s.fail(pos, i, "a namespace cannot contain '/'")
	}
	full := id
	if colon < 0 {
		full = "minecraft:" + id
	}
	if registry != "" {
		reg := registry
		if tag {
			reg = "tag/" + registry
		}
		s.ids = append(s.ids, ID{Start: pos, End: i, Registry: reg, Value: string(s.src[pos:i])})
		if s.known != nil && s.known(reg, full) {
			return i, true
		}
		if strings.HasPrefix(full, "minecraft:") && s.c.pack.HasRegistry(reg) && !s.c.pack.HasID(reg, full) {
			what := strings.TrimPrefix(reg, "tag/")
			if tag {
				return 0, s.fail(pos, i, fmt.Sprintf("unknown %s tag '#%s'", what, full))
			}
			return 0, s.fail(pos, i, fmt.Sprintf("unknown %s '%s'", what, full))
		}
	}
	return i, true
}

// entity parses a selector (@e[...]), a player name or a UUID.
func (s *state) entity(pos int, props map[string]any) (int, bool) {
	if s.src[pos] != '@' {
		end := s.word(pos)
		w := string(s.src[pos:end])
		if strings.Count(w, "-") == 4 {
			return end, true // a UUID
		}
		if end-pos > 16 || strings.IndexFunc(w, func(r rune) bool { return !isUnquoted(r) }) >= 0 {
			return 0, s.fail(pos, end, "expected a selector such as @s, or a player name")
		}
		return end, true
	}
	if pos+1 >= len(s.src) || !strings.ContainsRune("aenprs", s.src[pos+1]) {
		return 0, s.fail(pos, s.word(pos), "unknown selector; expected @a, @e, @n, @p, @r or @s")
	}
	kind := s.src[pos+1]
	end := pos + 2
	opts := map[string]string{}
	if end < len(s.src) && s.src[end] == '[' {
		close, ok := s.balancedEnd(end)
		if !ok || s.src[close-1] != ']' {
			return 0, s.fail(end, close, "selector options must be closed with ']'")
		}
		if !s.selectorOptions(end+1, close-1, opts) {
			return 0, false
		}
		end = close
	}
	if props["amount"] == "single" && (kind == 'a' || kind == 'e') && opts["limit"] != "1" {
		return 0, s.fail(pos, end, "only one entity is allowed here, but the selector allows more; add limit=1 or use @s, @p, @n or @r")
	}
	if props["type"] == "players" && (kind == 'e' || kind == 'n') {
		if t := strings.TrimPrefix(opts["type"], "minecraft:"); t != "player" {
			return 0, s.fail(pos, end, "only players are allowed here; use @a, @p, @r or @s, or add type=player")
		}
	}
	return end, true
}

// selectorOptions checks key=value pairs between from and to.
func (s *state) selectorOptions(from, to int, opts map[string]string) bool {
	i := from
	for i < to {
		for i < to && s.src[i] == ' ' {
			i++
		}
		if i >= to {
			break
		}
		k := i
		for i < to && s.src[i] != '=' && s.src[i] != ',' {
			i++
		}
		key := strings.TrimSpace(string(s.src[k:i]))
		if i >= to || s.src[i] != '=' {
			return s.fail(k, i, "expected '=' after selector option '"+key+"'")
		}
		if !selectorOptions[key] && !strings.Contains(key, interp) {
			return s.fail(k, i, "unknown selector option '"+key+"'")
		}
		i++
		v := i
		depth := 0
		for i < to {
			c := s.src[i]
			if c == '"' || c == '\'' {
				end, ok := s.quotedQuiet(i)
				if !ok {
					return s.fail(i, to, "this string is not closed")
				}
				i = end
				continue
			}
			if c == '{' || c == '[' {
				depth++
			} else if c == '}' || c == ']' {
				depth--
			} else if c == ',' && depth == 0 {
				break
			}
			i++
		}
		opts[key] = strings.TrimSpace(string(s.src[v:i]))
		if i < to {
			i++ // the comma
		}
	}
	return true
}
