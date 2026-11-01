package ide

import (
	"github.com/minecraftmetascript/mms/engine/sema"
	"github.com/minecraftmetascript/mms/engine/syntax"
	"github.com/minecraftmetascript/mms/engine/syntax/parser"
)

// assetPathCompletion offers project files for the path argument of an
// asset helper, Texture("|"), Sound("|") or Structure("|"), listing the
// files with the extension the helper takes. It also works before the
// closing quote is typed. ok is false when off is not in such a string.
func (s *Session) assetPathCompletion(f *syntax.File, off int) (items []Item, replace syntax.Span, ok bool) {
	start, end, ok := assetPathString(f, off)
	if !ok {
		return nil, replace, false
	}
	helper := assetHelperBefore(f, start-1)
	ext, what, ok := sema.AssetFileType(helper)
	if !ok {
		return nil, replace, false
	}
	for i, p := range s.P.Config.AssetFiles(s.P.Files, ext) {
		items = append(items, Item{Label: p, Kind: sema.CandFile, Detail: what, Sort: sortKey(i)})
	}
	return items, syntax.Span{Start: start, End: end}, true
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
