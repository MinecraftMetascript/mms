package lsp

import (
	"context"
	"sort"
	"strings"
	"testing/fstest"

	"github.com/minecraftmetascript/mms/engine/importer"
	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/sema"
)

// importPack turns data pack files (paths relative to the pack root, such
// as data/ns/loot_table/x.json) into formatted MMS files. The browser
// reads the zip; the engine does the rest.
func (s *Server) importPack(ctx context.Context, p struct {
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
	if err := res.Tidy(ctx, s.pack); err != nil {
		return nil, &respError{Code: codeRequestFailed, Message: err.Error()}
	}
	files := make([]map[string]any, 0, len(res.Files))
	for _, f := range res.Files {
		files = append(files, map[string]any{"path": f.Path, "text": f.Text})
	}
	return map[string]any{
		"files":       files,
		"namespace":   res.Namespace,
		"description": res.Description,
		"skipped":     orEmpty(res.Skipped),
		"count":       res.Count,
	}, nil
}

// convertJSON turns one pasted JSON resource into a declaration for the
// file uri, tidied against the project so it means the same there.
func (s *Server) convertJSON(ctx context.Context, p struct {
	URI  string `json:"uri"`
	Kind string `json:"kind"`
	Name string `json:"name"`
	Text string `json:"text"`
}) (any, *respError) {
	kinds := sema.NewKinds(s.pack)
	k, ok := kinds.Lookup(p.Kind)
	if !ok || !k.IsResource() {
		return nil, &respError{Code: codeInvalidParams, Message: "unknown resource kind " + p.Kind}
	}
	v, err := jsonv.Parse([]byte(p.Text))
	if err != nil {
		return nil, &respError{Code: codeRequestFailed, Message: "not valid JSON: " + err.Error()}
	}
	ns := s.namespaceAtEnd(p.URI)
	id := p.Name
	if !strings.Contains(id, ":") {
		id = ns + ":" + id
	}
	var src string
	if k.TagOf != "" {
		src = importer.TagDecl(k, ns, id, v)
	} else {
		src = importer.Decl(k, ns, id, v)
	}
	text, err := importer.TidyIn(ctx, s.proj, ns, src)
	if err != nil {
		return nil, &respError{Code: codeRequestFailed, Message: err.Error()}
	}
	return map[string]any{"text": text}, nil
}

// namespaceAtEnd is the namespace a declaration appended to uri lives in.
func (s *Server) namespaceAtEnd(uri string) string {
	ns := s.proj.Config.Namespace
	if s.res == nil {
		return ns
	}
	if u := s.res.Unit(uri); u != nil {
		for _, d := range u.File.Decls {
			if d.Kind != nil && d.Kind.Name == "Namespace" && d.Name != nil {
				ns = d.Name.Text
			}
		}
	}
	return ns
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
