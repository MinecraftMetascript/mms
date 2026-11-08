package lsp

import (
	"context"
	"strings"

	"github.com/minecraftmetascript/mms/engine/ide"
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
func (s *Server) codeAction(ctx context.Context, p struct {
	TextDocument docID `json:"textDocument"`
	Range        Range `json:"range"`
	Context      struct {
		Only []string `json:"only"`
	} `json:"context"`
}) (any, *respError) {
	uri := p.TextDocument.URI
	w, u := s.unit(uri)
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
	ds := w.res.Diagnostics(uri)
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
	if want("refactor.extract") {
		for _, x := range w.session().Extractions(uri, from, to) {
			// The name's position after the edits: the declaration is
			// inserted at a line start, ahead of the replaced value.
			at := l.pos(x.Edits[0].Span.Start)
			at.Character += len(x.Kind) + 1
			a := map[string]any{
				"title": x.Title,
				"kind":  "refactor.extract",
				"edit":  map[string]any{"changes": map[string]any{uri: editsJSON(l, toSyntaxEdits(x.Edits))}},
				"data":  map[string]any{"rename": map[string]any{"uri": uri, "position": at, "name": x.Name}},
			}
			if s.renameAfterExtract {
				a["command"] = map[string]any{"title": "Rename", "command": "mms.renameAt", "arguments": []any{uri, at}}
			}
			out = append(out, a)
		}
	}
	// Fix all analyses the project in rounds, so it is offered only when
	// asked for by kind (on save, or the Source Action menu), never for
	// the lightbulb request an editor sends on every cursor move.
	if len(p.Context.Only) > 0 && want("source.fixAll.mms") {
		// Fixes can expose more (an object inside one that became a call),
		// so fix in rounds, on a fork so the live project keeps its state.
		files, n, err := w.proj.Fork().FixAll(ctx, uri)
		if err != nil {
			return nil, &respError{Code: codeRequestFailed, Message: err.Error()}
		}
		if text := files[uri]; n > 0 {
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

func toSyntaxEdits(edits []ide.Edit) []syntax.TextEdit {
	out := make([]syntax.TextEdit, 0, len(edits))
	for _, e := range edits {
		out = append(out, syntax.TextEdit{Sp: e.Span, Text: e.Text})
	}
	return out
}
