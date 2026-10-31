package lsp

import (
	"context"
	"strings"

	"github.com/minecraftmetascript/mms/engine/project"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

func fixesJSON(l *lines, fixes []syntax.Fix) []map[string]any {
	out := make([]map[string]any, 0, len(fixes))
	for _, f := range fixes {
		out = append(out, map[string]any{"title": f.Title, "edits": editsJSON(l, f.Edits), "safe": f.Safe})
	}
	return out
}

func editsJSON(l *lines, edits []syntax.TextEdit) []map[string]any {
	out := make([]map[string]any, 0, len(edits))
	for _, e := range edits {
		out = append(out, map[string]any{"range": l.rng(e.Sp), "newText": e.Text})
	}
	return out
}

// codeAction offers each fix of the diagnostics touching the range as a
// quick fix, and "fix all" as a source action.
func (s *Server) codeAction(_ context.Context, p struct {
	TextDocument docID `json:"textDocument"`
	Range        Range `json:"range"`
	Context      struct {
		Only []string `json:"only"`
	} `json:"context"`
}) (any, *respError) {
	uri := p.TextDocument.URI
	u := s.res.Unit(uri)
	if u == nil {
		return []any{}, nil
	}
	l := newLines(string(u.File.Src))
	from, to := l.offset(p.Range.Start), l.offset(p.Range.End)
	want := func(kind string) bool {
		if len(p.Context.Only) == 0 {
			return true
		}
		for _, k := range p.Context.Only {
			if kind == k || strings.HasPrefix(kind, k+".") {
				return true
			}
		}
		return false
	}
	ds := s.res.Diagnostics(uri)
	out := []any{}
	if want("quickfix") {
		for _, d := range ds {
			if d.Sp.End < from || d.Sp.Start > to {
				continue
			}
			for i, f := range d.Fixes {
				out = append(out, map[string]any{
					"title":       f.Title,
					"kind":        "quickfix",
					"isPreferred": i == 0 && f.Safe,
					"diagnostics": []any{map[string]any{"range": l.rng(d.Sp), "message": d.Message, "code": d.Code}},
					"edit":        map[string]any{"changes": map[string]any{uri: editsJSON(l, f.Edits)}},
				})
			}
		}
	}
	if want("source.fixAll.mms") {
		text, n := project.ApplyFixes(u.File.Src, ds)
		if n > 0 {
			end := l.pos(len(u.File.Src))
			out = append(out, map[string]any{
				"title": "Fix all MMS problems",
				"kind":  "source.fixAll.mms",
				"edit": map[string]any{"changes": map[string]any{uri: []any{
					map[string]any{"range": Range{Start: Position{}, End: end}, "newText": text},
				}}},
			})
		}
	}
	return out, nil
}
