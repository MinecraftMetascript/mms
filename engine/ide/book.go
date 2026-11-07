package ide

import (
	"strconv"
	"strings"

	"github.com/minecraftmetascript/mms/engine/sema"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// Text markup completion and hover, in .mmsbook and .mmstext files, in
// the strings of a Book's text field and in strings where a text
// component goes.

// bookText finds the book markup at off: in a .mmsbook file, the file;
// in a .mms file, the text field string the cursor is in. It returns the
// markup's runes and the file offset of its first rune.
func (s *Session) bookText(path string, off int) (text []rune, base int, ok bool) {
	u := s.R.Unit(path)
	if u == nil {
		return nil, 0, false
	}
	if sema.IsBookFile(path) {
		return u.File.Src, 0, true
	}
	start, end, ok := assetPathString(u.File, off)
	if !ok || !s.inBookText(u, start) {
		return nil, 0, false
	}
	return u.File.Src[start:end], start, true
}

// inBookText reports whether the string starting at off is text markup:
// a Book's text, or a string where a text component goes.
func (s *Session) inBookText(u *sema.Unit, off int) bool {
	a := s.at(u.Path, off)
	if a == nil || a.decl == nil {
		return false
	}
	for i := len(a.path) - 1; i >= 0; i-- {
		if str, ok := a.path[i].(*syntax.String); ok {
			in := a.info(str)
			return in != nil && in.Markup
		}
	}
	return false
}

// bookContext is what the cursor is in: a {...} tag, a link target or a
// ${...} expression, with the text typed so far.
type bookContext struct {
	kind  string // "tag", "target", "expr"
	typed string // from the opening to the cursor
	start int    // offset in the markup of the first typed rune
}

func bookContextAt(text []rune, off int) (bookContext, bool) {
	if off > len(text) {
		off = len(text)
	}
	for i := off - 1; i >= 0; i-- {
		switch text[i] {
		case '\n', '}', ')':
			return bookContext{}, false
		case '{':
			if i > 0 && text[i-1] == '\\' {
				return bookContext{}, false
			}
			if i > 0 && text[i-1] == '$' {
				return bookContext{kind: "expr", typed: string(text[i+1 : off]), start: i + 1}, true
			}
			return bookContext{kind: "tag", typed: string(text[i+1 : off]), start: i + 1}, true
		case '(':
			if i > 0 && text[i-1] == ']' {
				return bookContext{kind: "target", typed: string(text[i+1 : off]), start: i + 1}, true
			}
		}
	}
	return bookContext{}, false
}

var bookTagHelp = map[string]string{
	"/":             "Closes the innermost `{...}` style.",
	"key:":          "The key the reader has bound to an action, such as `{key:jump}`, which shows Space unless they changed it.",
	"sprite:":       "An inline icon from a texture atlas: `{sprite:item/diamond}`, or `{sprite:block/stone}`. Add `atlas:gui` for another atlas.",
	"head:":         "A player's head, inline: `{head:Notch}`.",
	"t:":            "A translated text, such as `{t:block.minecraft.stone}`; the reader's language decides the words.",
	"shadow:":       "The text's shadow color: `shadow:#3f1f00`, `shadow:#3f1f0080` with alpha, or `shadow:none`.",
	"font:":         "A font: `font:uniform`, `font:alt`, or one of the project's fonts.",
	"bold":          "Bold text, like `**...**`.",
	"italic":        "Italic text, like `*...*`.",
	"underlined":    "Underlined text, like `__...__`.",
	"strikethrough": "Struck-through text, like `~~...~~`.",
	"obfuscated":    "Scrambled, constantly changing text (the §k effect), like `||...||`.",
}

var bookTargetHelp = map[string]string{
	"page:":    "Turns to a page of this book: `[Next](page:2)`.",
	"run:":     "Runs a command as the reader when clicked: `[Start](run:/function demo:start)`. The game asks the reader to confirm commands that need more permissions than they have.",
	"suggest:": "Puts text in the reader's chat box: `[Teleport](suggest:/tp @s ~ ~10 ~)`.",
	"copy:":    "Copies text to the clipboard: `[Seed](copy:12345)`.",
	"hover:":   "Shows text when the reader points at the link, with no click: `[ruby](hover:A red gem)`.",
	"https://": "Opens a web page, after the game asks: `[Wiki](https://minecraft.wiki)`.",
}

// bookCompletion completes book markup at off in the file.
func (s *Session) bookCompletion(path string, off int) ([]Item, syntax.Span, bool) {
	text, base, ok := s.bookText(path, off)
	if !ok {
		return nil, syntax.Span{}, false
	}
	ctx, ok := bookContextAt(text, off-base)
	if !ok {
		return nil, syntax.Span{Start: off, End: off}, true
	}
	var items []Item
	add := func(label string, k sema.CandidateKind, detail, doc string) {
		items = append(items, Item{Label: label, Kind: k, Detail: detail, Doc: doc, Sort: sortKey(len(items))})
	}
	replace := func(from int) syntax.Span { return syntax.Span{Start: base + from, End: off} }
	switch ctx.kind {
	case "tag":
		// The word being typed: tags combine with spaces.
		w := ctx.typed[strings.LastIndexByte(ctx.typed, ' ')+1:]
		wStart := ctx.start + len([]rune(ctx.typed)) - len([]rune(w))
		pre, val, hasVal := strings.Cut(w, ":")
		if hasVal {
			vStart := wStart + len([]rune(pre)) + 1
			switch pre {
			case "key":
				for _, k := range s.specEnum("::java::util::text::Keybind") {
					add(strings.TrimPrefix(k, "key."), sema.CandEnum, k, "")
				}
			case "sprite":
				for _, t := range s.P.Pack.Registries["texture"] {
					if strings.HasPrefix(t, "item/") || strings.HasPrefix(t, "block/") {
						add(t, sema.CandID, "texture", "")
					}
				}
			case "font":
				for _, f := range s.P.Pack.Registries["font"] {
					if !strings.HasPrefix(f, "include/") {
						add(f, sema.CandID, "font", "")
					}
				}
			case "atlas":
				for _, a := range s.P.Pack.Registries["atlas"] {
					add(a, sema.CandID, "atlas", "")
				}
			default:
				return nil, replace(vStart), true
			}
			_ = val
			return items, replace(vStart), true
		}
		if strings.TrimSpace(ctx.typed) == w {
			// Only alone in the braces.
			for _, k := range []string{"/", "key:", "sprite:", "head:", "t:"} {
				add(k, sema.CandKeyword, "", bookTagHelp[k])
			}
		}
		for _, c := range s.specEnum("::java::util::text::TextColor") {
			add(c, sema.CandEnum, "color", "")
		}
		for _, w := range []string{"bold", "italic", "underlined", "strikethrough", "obfuscated", "shadow:", "font:"} {
			add(w, sema.CandKeyword, "style", bookTagHelp[w])
		}
		return items, replace(wStart), true
	case "target":
		kind, rest, hasKind := strings.Cut(ctx.typed, ":")
		if hasKind && kind == "run" {
			cmd := strings.TrimPrefix(rest, "/")
			cmdStart := ctx.start + len([]rune(ctx.typed)) - len([]rune(cmd))
			its, ws := s.commandItems(cmd, len([]rune(cmd)))
			return its, replace(cmdStart + ws), true
		}
		if hasKind && kind == "page" {
			pages := 1 + strings.Count("\n"+string(text)+"\n", "\n---\n")
			for i := 1; i <= pages; i++ {
				add(strconv.Itoa(i), sema.CandEnum, "page", "")
			}
			return items, replace(ctx.start + len([]rune(kind)) + 1), true
		}
		if hasKind {
			return nil, replace(off - base), true
		}
		for _, k := range []string{"page:", "run:/", "suggest:/", "copy:", "hover:", "https://"} {
			add(k, sema.CandKeyword, "link", bookTargetHelp[strings.TrimSuffix(k, "/")])
		}
		return items, replace(ctx.start), true
	case "expr":
		w := ctx.typed
		for _, sym := range s.R.Syms.All() {
			switch sym.Kind.MMS {
			case "Objective":
				add(sym.Path, sema.CandID, "objective: write "+sym.Path+"[@s]", sema.DocMarkdown(sym.Decl.Doc))
			case "Global":
				add(sym.Path, sema.CandID, "Global", sema.DocMarkdown(sym.Decl.Doc))
			}
		}
		return items, replace(ctx.start + len([]rune(w)) - len([]rune(lastIdent(w)))), true
	}
	return nil, syntax.Span{Start: off, End: off}, true
}

func lastIdent(s string) string {
	i := len(s)
	for i > 0 {
		c := s[i-1]
		if c == '_' || c == ':' || c == '.' || c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9' {
			i--
			continue
		}
		break
	}
	return s[i:]
}

func (s *Session) specEnum(path string) []string {
	t := s.P.Pack.Resolve(s.P.Pack.Types[path])
	if t == nil {
		return nil
	}
	var out []string
	for _, v := range t.Values {
		out = append(out, v.String())
	}
	return out
}

// bookHover describes the tag or link target at off in book markup.
func (s *Session) bookHover(path string, off int) (string, syntax.Span, bool) {
	text, base, ok := s.bookText(path, off)
	if !ok {
		return "", syntax.Span{}, false
	}
	o := off - base
	// The {...} or (...) around the cursor.
	lo, hi := -1, -1
	for i := min(o, len(text)-1); i >= 0 && text[i] != '\n'; i-- {
		if text[i] == '{' || (text[i] == '(' && i > 0 && text[i-1] == ']') {
			lo = i
			break
		}
		if (text[i] == '}' || text[i] == ')') && i != o {
			break
		}
	}
	if lo < 0 {
		return "", syntax.Span{}, true
	}
	closer := '}'
	if text[lo] == '(' {
		closer = ')'
	}
	for i := lo + 1; i < len(text) && text[i] != '\n'; i++ {
		if text[i] == closer {
			hi = i
			break
		}
	}
	if hi < 0 || o > hi {
		return "", syntax.Span{}, true
	}
	body := strings.TrimSpace(string(text[lo+1 : hi]))
	sp := syntax.Span{Start: base + lo, End: base + hi + 1}
	if closer == ')' {
		for k, doc := range bookTargetHelp {
			if strings.HasPrefix(body, k) || strings.HasPrefix(body, strings.TrimSuffix(k, "//")) {
				return "**Link** `" + body + "`\n\n" + doc, sp, true
			}
		}
		return "", syntax.Span{}, true
	}
	if lo > 0 && text[lo-1] == '$' {
		return s.markupExprHover(path, off, body), sp, true
	}
	if k, v, ok := strings.Cut(body, ":"); ok && bookTagHelp[k+":"] != "" && !strings.Contains(k, " ") {
		doc := bookTagHelp[k+":"]
		if k == "key" {
			doc = "Shows the key the reader has bound to `key." + strings.TrimPrefix(v, "key.") + "`.\n\n" + doc
		}
		return doc, sp, true
	}
	var parts []string
	for _, w := range strings.Fields(body) {
		key := w
		if i := strings.IndexByte(w, ':'); i > 0 {
			key = w[:i+1]
		}
		switch {
		case bookTagHelp[key] != "":
			parts = append(parts, "`"+w+"`: "+bookTagHelp[key])
		default:
			parts = append(parts, "`"+w+"`: a text color")
		}
	}
	if len(parts) == 0 {
		return "", syntax.Span{}, true
	}
	return strings.Join(parts, "\n\n") + "\n\nEnds at the next `{/}`.", sp, true
}

// bookTrigger reports whether a trigger character opens completion in
// book markup: '{', ':', '(' and '$' do there.
func (s *Session) bookTrigger(path string, off int) bool {
	text, base, ok := s.bookText(path, off)
	if !ok {
		return false
	}
	_, ok = bookContextAt(text, off-base)
	return ok
}

// markupExprHover describes `${body}` in text markup: the Global or
// Storage field it names, and when it is filled in, which depends on
// where the text is. A book fills it in when the reader first opens it,
// function code when the line runs, and other text when it is shown.
func (s *Session) markupExprHover(path string, off int, body string) string {
	when := "A score, a Global or a constant, filled in when the text is shown."
	a := s.at(path, off)
	switch {
	case sema.IsBookFile(path) || a != nil && a.decl != nil && a.decl.D != nil && a.decl.D.Kind != nil && a.decl.D.Kind.Name == "Book":
		when = "A score, a Global or a constant, filled in when the reader first opens the book."
	case a != nil && a.code():
		when = "A value the function computes, filled in when this line runs: a local, a parameter, a Global, a Storage field, a score or a constant."
	}
	if md := s.markupNameHover(body); md != "" {
		return strings.TrimSpace(md) + "\n\n" + when
	}
	return when
}

// markupNameHover describes a Global, `g`, or a Storage field,
// `game.wave`, by the name written in ${...}; "" for anything else.
func (s *Session) markupNameHover(body string) string {
	name, field, isField := strings.Cut(body, ".")
	for _, sym := range s.R.Syms.All() {
		if sym.Path != name && sym.ID() != name {
			continue
		}
		switch {
		case !isField && sym.Kind.MMS == "Global":
			if d := s.R.Decls[sym.Decl]; d != nil {
				return s.declHover(d)
			}
		case isField && sym.Kind.MMS == "Storage":
			tn, ok := sym.Decl.Type.(*syntax.TypeName)
			if !ok {
				return ""
			}
			td := s.R.Syms.TypeDecls[tn.Name]
			if td == nil || td.Body == nil {
				return ""
			}
			for _, it := range td.Body.Items {
				if f, ok := it.(*syntax.FieldDecl); ok && f.Name != nil && f.Name.Text == field && f.Type != nil {
					return "```mms\n" + body + ": " + syntax.TypeText(f.Type) + "\n```\n" + sema.DocMarkdown(f.Doc)
				}
			}
		}
	}
	return ""
}
