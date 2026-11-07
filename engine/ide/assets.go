package ide

import (
	"github.com/minecraftmetascript/mms/engine/sema"
	"github.com/minecraftmetascript/mms/engine/syntax"
	"github.com/minecraftmetascript/mms/engine/syntax/parser"
	"strings"
)

// assetPathCompletion offers project files for the path argument of an
// asset helper, Texture("|"), Sound("|") or StructureFile("|"), listing the
// files with the extension the helper takes. It also works before the
// closing quote is typed. ok is false when off is not in such a string.
func (s *Session) assetPathCompletion(f *syntax.File, off int) (items []Item, replace syntax.Span, ok bool) {
	start, end, ok := assetPathString(f, off)
	if !ok {
		return nil, replace, false
	}
	helper := assetHelperBefore(f, start-1)
	if helper == "TextFile" {
		seen := map[string]bool{}
		for _, u := range s.R.Units {
			if !sema.IsBookFile(u.Path) {
				continue
			}
			// TextFile takes a path in the project; a host that names
			// files by URI, such as the language server, has them under
			// the project's Root.
			rel := u.Path
			if s.P.Root != "" && strings.HasPrefix(rel, s.P.Root) {
				rel = rel[len(s.P.Root):]
			} else if strings.Contains(rel, "://") {
				continue // not in this project
			}
			if seen[rel] {
				continue
			}
			seen[rel] = true
			items = append(items, Item{Label: rel, Kind: sema.CandFile, Detail: "book text", Sort: sortKey(len(items))})
		}
		return items, syntax.Span{Start: start, End: end}, true
	}
	prefix, suffix := "", ""
	ext, what, ok := sema.AssetFileType(helper)
	if !ok {
		// A Template that passes its first parameter on as part of an
		// asset path, as Piece("halls/hall") does with
		// StructureFile("structures/${path}.nbt").
		helper, prefix, suffix, ok = s.templateAssetPath(helper)
		if !ok {
			return nil, replace, false
		}
		ext, what, _ = sema.AssetFileType(helper)
	}
	for _, p := range s.P.Config.AssetFiles(s.P.Files, ext) {
		if !strings.HasPrefix(p, prefix) || !strings.HasSuffix(p, suffix) || len(p) < len(prefix)+len(suffix) {
			continue
		}
		p = p[len(prefix) : len(p)-len(suffix)]
		items = append(items, Item{Label: p, Kind: sema.CandFile, Detail: what, Sort: sortKey(len(items))})
	}
	return items, syntax.Span{Start: start, End: end}, true
}

// templateAssetPath finds, in the Template name, a call of an asset
// helper whose path is the Template's first parameter with text around
// it: StructureFile("structures/${path}.nbt") gives StructureFile,
// "structures/" and ".nbt".
func (s *Session) templateAssetPath(name string) (helper, prefix, suffix string, ok bool) {
	if s.R.Eval == nil {
		return "", "", "", false
	}
	t := s.R.Eval.Templates[name]
	if t == nil || t.D.Params == nil || len(t.D.Params.Params) == 0 || t.D.Params.Params[0].Name == nil || t.D.Value == nil {
		return "", "", "", false
	}
	hole := "${" + t.D.Params.Params[0].Name.Text + "}"
	syntax.Walk(t.D.Value, func(n syntax.Node) bool {
		c, isCall := n.(*syntax.Call)
		if ok || !isCall || len(c.Args) == 0 {
			return !ok
		}
		id, isIdent := c.Fn.(*syntax.Ident)
		str, isStr := c.Args[0].Value.(*syntax.String)
		if !isIdent || !isStr {
			return true
		}
		if _, _, known := sema.AssetFileType(id.Name); !known {
			return true
		}
		if before, after, found := strings.Cut(str.Value, hole); found && !strings.Contains(after, "${") && !strings.Contains(before, "${") {
			helper, prefix, suffix, ok = id.Name, before, after, true
		}
		return !ok
	})
	return helper, prefix, suffix, ok
}

// assetPathString finds the string literal the cursor is in and returns
// the span of its text between the quotes. An unterminated string (the
// opening quote lexes on its own) runs to the cursor.
func assetPathString(f *syntax.File, off int) (start, end int, ok bool) {
	if i := f.TokenAt(off); i >= 0 {
		t := f.Tokens[i]
		if t.Kind == parser.MMSLexerSTRING && t.Sp.Start < off && off < t.Sp.End {
			return t.Sp.Start + 1, t.Sp.End - 1, true
		}
	}
	// Look back along the line for an opening quote with no partner.
	for q := off - 1; q >= 0 && q < len(f.Src); q-- {
		switch f.Src[q] {
		case '\n', '\r':
			return 0, 0, false
		case '"', '\'':
			for _, t := range f.Tokens {
				if t.Sp.Start == q && t.Kind == parser.MMSLexerERROR_CHAR {
					return q + 1, off, true
				}
			}
			return 0, 0, false
		}
	}
	return 0, 0, false
}

// assetHelperBefore returns the helper name when the string whose quote
// is at q is the first argument of a plain call `Name(`, or "".
func assetHelperBefore(f *syntax.File, q int) string {
	paren := prevToken(f, q)
	if paren == nil || paren.Kind != parser.MMSLexerLPAREN {
		return ""
	}
	name := prevToken(f, paren.Sp.Start)
	if name == nil || name.Kind != parser.MMSLexerIDENT {
		return ""
	}
	// A method such as x.Texture( is not the helper.
	if dot := prevToken(f, name.Sp.Start); dot != nil && dot.Kind == parser.MMSLexerDOT && dot.Sp.End == name.Sp.Start {
		return ""
	}
	return name.Text
}
