package sema

import (
	"embed"
	"io/fs"
	"regexp"
	"sort"
	"strings"
	"sync"

	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/spec"
)

// Hand-written documentation for helpers, filling what vanilla-mcdoc
// leaves out: mcdoc has no descriptions for types, and documents only
// about a quarter of the fields. The files in helpdocs/ are Markdown:
//
//	## minecraft:material_condition/stone_depth
//	One or two sentences on what the helper makes and when to use it.
//
//	- surface_type: What the parameter means.
//	- offset: Continuation lines are indented
//	  by two spaces.
//
//	```mms
//	If(StoneDepth(Floor)) { grass_block }
//	```
//
// A heading names a dispatched helper by dispatcher and case, and a
// struct helper by its type's short name (LootPool).

//go:embed helpdocs/*.md
var helpFiles embed.FS

// HelperDoc is the hand-written documentation of one helper.
type HelperDoc struct {
	Summary string            // Markdown
	Params  map[string]string // Markdown by field key
	Example string            // MMS source
}

// ReferenceURL is the docs site's helper reference, which hovers link to.
const ReferenceURL = "https://minecraftmetascript.github.io/mms/docs/reference"

var helpDocs = sync.OnceValue(func() map[string]*HelperDoc {
	out := map[string]*HelperDoc{}
	files, _ := fs.Glob(helpFiles, "helpdocs/*.md")
	for _, name := range files {
		b, _ := helpFiles.ReadFile(name)
		parseHelpDocs(string(b), out)
	}
	return out
})

var paramLine = regexp.MustCompile("^- `?([a-z0-9_]+)`?: (.*)$")

func parseHelpDocs(src string, out map[string]*HelperDoc) {
	var cur *HelperDoc
	var summary []string
	var param string
	inCode := false
	var code []string
	flush := func() {
		if cur != nil {
			cur.Summary = strings.TrimSpace(strings.Join(summary, "\n"))
		}
		summary, param = nil, ""
	}
	for _, line := range strings.Split(src, "\n") {
		line = strings.TrimRight(line, " \t\r")
		switch {
		case inCode:
			if line == "```" {
				inCode = false
				if cur != nil {
					cur.Example = strings.Join(code, "\n")
				}
				continue
			}
			code = append(code, line)
		case strings.HasPrefix(line, "## "):
			flush()
			cur = &HelperDoc{Params: map[string]string{}}
			out[strings.TrimSpace(line[3:])] = cur
		case cur == nil:
		case strings.HasPrefix(line, "```"):
			inCode, code, param = true, nil, ""
		case paramLine.MatchString(line):
			m := paramLine.FindStringSubmatch(line)
			param = m[1]
			cur.Params[param] = m[2]
		case param != "" && strings.HasPrefix(line, "  "):
			cur.Params[param] += " " + strings.TrimSpace(line)
		default:
			param = ""
			summary = append(summary, line)
		}
	}
	flush()
}

// docKey names a helper in helpdocs/.
func docKey(cl *Callable) string {
	if cl.Disp != nil && cl.Case != "" {
		return cl.Disp.Dispatcher + "/" + cl.Case
	}
	if cl.Struct.Name != "" {
		return spec.ShortPath(cl.Struct.Name)
	}
	return cl.Name
}

// HelpDoc returns the hand-written documentation of a helper, or nil.
func HelpDoc(cl *Callable) *HelperDoc {
	if cl == nil {
		return nil
	}
	return helpDocs()[docKey(cl)]
}

// withDocs adds the descriptions of a helper's fields from helpdocs/,
// ahead of mcdoc's own. A `## <dispatcher>/*` section documents fields
// that every helper of a dispatcher shares, such as `condition` on loot
// functions. fs is copied, never changed.
func withDocs(key, dispatcher string, fs []field) []field {
	d := helpDocs()[key]
	var shared *HelperDoc
	if dispatcher != "" {
		shared = helpDocs()[dispatcher+"/*"]
	}
	if (d == nil || len(d.Params) == 0) && shared == nil {
		return fs
	}
	out := make([]field, len(fs))
	copy(out, fs)
	for i := range out {
		doc, ok := "", false
		if d != nil {
			doc, ok = d.Params[out[i].Key]
		}
		if !ok && shared != nil {
			doc, ok = shared.Params[out[i].Key]
		}
		if !ok {
			continue
		}
		if mc := strings.TrimSpace(out[i].Desc); mc != "" && !strings.Contains(doc, mc) {
			doc += " " + mc
		}
		out[i].Desc = doc
	}
	return out
}

// RefAnchor is the id of a helper's entry on the reference page.
func RefAnchor(dispatcher, name string) string {
	if dispatcher == "" {
		return "objects-" + name
	}
	return nonWord.ReplaceAllString(strings.TrimPrefix(dispatcher, "minecraft:"), "-") + "-" + name
}

var nonWord = regexp.MustCompile(`\W`)

func refLink(cl *Callable) string {
	group := ""
	if cl.Disp != nil {
		group = cl.Disp.Dispatcher
	}
	return ReferenceURL + "#" + RefAnchor(group, cl.Name)
}

// groupNoun names what a dispatcher's helpers make, in words:
// "minecraft:material_condition" is "a material condition".
func groupNoun(dispatcher string) string {
	s := strings.TrimPrefix(dispatcher, "minecraft:")
	s = strings.TrimPrefix(s, "worldgen/")
	s = strings.ReplaceAll(s, "_", " ")
	if s == "" {
		return "an object"
	}
	if strings.ContainsRune("aeiou", rune(s[0])) {
		return "an " + s
	}
	return "a " + s
}

// makes is the sentence that says what a helper builds.
func (c *Checker) makes(cl *Callable) string {
	if cl.Case != "" && cl.Disp != nil {
		return "Makes `minecraft:" + cl.Case + "`, " + groupNoun(cl.Disp.Dispatcher) + "."
	}
	return "Makes " + describe(cl.Struct) + "."
}

// defaultOf is the value MMS writes for a parameter left out, or "".
func defaultOf(cl *Callable, key string) string {
	if cl.Ov == nil {
		return ""
	}
	if v, ok := cl.Ov.defaults[key]; ok {
		return jsonv.Compact(v)
	}
	return ""
}

// shortType is a parameter's type as a signature shows it: a word or two,
// so signatures fit on narrow screens.
func (c *Checker) shortType(v tview) string {
	if v.Name != "" {
		return spec.ShortPath(v.Name)
	}
	n := c.norm(v)
	if n.Name != "" {
		return spec.ShortPath(n.Name)
	}
	if n.T == nil {
		return "any"
	}
	switch n.T.Kind {
	case "string":
		if id := idOf(n.T); id.Registry != "" {
			return strings.TrimPrefix(id.Registry, "worldgen/") + " id"
		}
		return "string"
	case "boolean":
		return "bool"
	case "int", "long", "short", "byte":
		return "int"
	case "float", "double":
		return "float"
	case "list":
		return "[" + c.shortType(tview{T: n.T.Item, Env: n.Env}) + "]"
	case "struct":
		return "object"
	case "enum":
		return "enum"
	case "union":
		var ms []string
		seen := map[string]bool{}
		for _, m := range n.T.Members {
			s := c.shortType(tview{T: m, Env: n.Env})
			if !seen[s] {
				seen[s] = true
				ms = append(ms, s)
			}
		}
		if len(ms) <= 3 {
			return strings.Join(ms, " | ")
		}
	}
	s := spec.Format(n.T, 0)
	if len(s) > 24 {
		return "value"
	}
	return s
}

// enumValues lists the values of a small enum type, or nil.
func (c *Checker) enumValues(v tview) []string {
	n := c.norm(v)
	if n.T == nil || n.T.Kind != "enum" || len(n.T.Values) > 10 {
		return nil
	}
	var out []string
	for _, e := range n.T.Values {
		out = append(out, "`"+e.String()+"`")
	}
	return out
}

// paramDoc is the Markdown for one parameter: its type, default and
// description.
func (c *Checker) paramDoc(cl *Callable, f field) string {
	var meta []string
	meta = append(meta, describe(f.Type))
	if d := defaultOf(cl, f.Key); d != "" {
		meta = append(meta, "default `"+d+"`")
	} else if f.Optional {
		meta = append(meta, "optional")
	}
	s := "*" + strings.Join(meta, ", ") + "*"
	if f.Desc != "" {
		s += ". " + strings.TrimSpace(f.Desc)
	}
	if vs := c.enumValues(f.Type); len(vs) > 0 && !strings.Contains(f.Desc, vs[0]) {
		s += " One of " + strings.Join(vs, ", ") + "."
	}
	return s
}

// helperMethods lists the methods an overlay adds, such as `.Add()`.
func helperMethods(cl *Callable) []string {
	if cl.Ov == nil {
		return nil
	}
	var out []string
	for name, m := range cl.Ov.methods {
		if m.val != nil {
			out = append(out, "`."+name+"()` sets `"+m.field+"` to `"+jsonv.Compact(m.val)+"`")
		} else {
			out = append(out, "`."+name+"(x)` sets `"+m.field+"`")
		}
	}
	sort.Strings(out)
	return out
}

// helperHover is the full hover for a helper call.
func (c *Checker) helperHover(cl *Callable) string {
	doc := HelpDoc(cl)
	var sb strings.Builder
	sb.WriteString("```mms\n" + c.signature(cl) + "\n```\n\n")
	if doc != nil && doc.Summary != "" {
		sb.WriteString(doc.Summary + "\n\n")
	}
	sb.WriteString(c.makes(cl) + "\n")
	if ps := c.positional(cl); len(ps) > 0 {
		sb.WriteString("\n**Parameters**\n\n")
		for _, f := range ps {
			sb.WriteString("- `" + f.Key + "`: " + c.paramDoc(cl, f) + "\n")
		}
	}
	if ms := helperMethods(cl); len(ms) > 0 {
		sb.WriteString("\n**Methods**\n\n")
		for _, m := range ms {
			sb.WriteString("- " + m + "\n")
		}
	}
	if doc != nil && doc.Example != "" {
		sb.WriteString("\n**Example**\n\n```mms\n" + doc.Example + "\n```\n")
	}
	sb.WriteString("\n[Reference](" + refLink(cl) + ")")
	return sb.String()
}

// helperSignatureDoc is the short documentation signature help shows
// under the call: the summary's first paragraph, what it makes and a
// link. Signature help sits over the code, so it stays brief.
func (c *Checker) helperSignatureDoc(cl *Callable) string {
	var parts []string
	if doc := HelpDoc(cl); doc != nil && doc.Summary != "" {
		first, _, _ := strings.Cut(doc.Summary, "\n\n")
		parts = append(parts, first)
	}
	parts = append(parts, c.makes(cl)+" [Reference]("+refLink(cl)+")")
	return strings.Join(parts, "\n\n")
}

// HelpDocs returns every hand-written helper doc by key, for tests.
func HelpDocs() map[string]*HelperDoc { return helpDocs() }
