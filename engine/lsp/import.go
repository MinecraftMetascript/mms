package lsp

import (
	"context"
	"sort"
	"strings"
	"testing/fstest"

	"github.com/minecraftmetascript/mms/engine/format"
	"github.com/minecraftmetascript/mms/engine/importer"
	"github.com/minecraftmetascript/mms/engine/sema"
)

// importPack turns data pack files (paths relative to the pack root, such
// as data/ns/loot_table/x.json) into formatted MMS files. The browser
// reads the zip; the engine does the rest.
func (s *Server) importPack(_ context.Context, p struct {
	Files []struct {
		Path string `json:"path"`
		Text string `json:"text"`
	} `json:"files"`
}) (any, *respError) {
	fsys := fstest.MapFS{}
	for _, f := range p.Files {
		fsys[strings.TrimPrefix(f.Path, "/")] = &fstest.MapFile{Data: []byte(f.Text)}
	}
	res, err := importer.Pack(fsys, sema.NewKinds(s.pack), s.pack)
	if err != nil {
		return nil, &respError{Code: codeRequestFailed, Message: err.Error()}
	}
	files := make([]map[string]any, 0, len(res.Files))
	for _, f := range res.Files {
		text := f.Text
		if out, err := format.Source(text); err == nil {
			text = out
		}
		files = append(files, map[string]any{"path": f.Path, "text": text})
	}
	return map[string]any{
		"files":       files,
		"namespace":   res.Namespace,
		"description": res.Description,
		"skipped":     orEmpty(res.Skipped),
		"count":       res.Count,
	}, nil
}

// resourceKinds lists the JSON resource kinds and their folders, for the
// editor's Paste JSON dialog.
func (s *Server) resourceKinds(_ context.Context, _ struct{}) (any, *respError) {
	ks := sema.NewKinds(s.pack)
	out := []map[string]string{}
	for _, name := range ks.Names() {
		k, _ := ks.Lookup(name)
		if k.Resource == "" {
			continue
		}
		out = append(out, map[string]string{"kind": name, "folder": k.Folder()})
	}
	sort.Slice(out, func(i, j int) bool { return out[i]["kind"] < out[j]["kind"] })
	return out, nil
}

func orEmpty(s []string) []string {
	if s == nil {
		return []string{}
	}
	return s
}
