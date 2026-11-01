package lsp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"testing/fstest"
	"time"

	"github.com/minecraftmetascript/mms/engine/format"
	"github.com/minecraftmetascript/mms/engine/ide"
	"github.com/minecraftmetascript/mms/engine/project"
	"github.com/minecraftmetascript/mms/engine/sema"
	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// Server is one language server connection. Project files are keyed by
// document URI. Each folder holding mms.json is its own project, and a
// file belongs to the nearest one above it; files under no mms.json share
// the default project.
type Server struct {
	tr   Transport
	pack *spec.Pack
	// Log receives internal errors; nil discards them.
	Log io.Writer

	roots []string // the editor's folders on disk; none in the browser
	// projects are keyed by the folder holding mms.json; "" is the
	// default project. nil until initialize.
	projects map[string]*workspace
	owners   map[string]*workspace // the project each file is in
	last     *workspace            // the project of the last edited file
	lastURI  string                // the last edited file
	open     map[string]bool
	// managed files were given by mms/setFiles: closing their editor
	// keeps them in the project (the browser has no disk to fall back on).
	managed   map[string]bool
	published map[string]bool
	// assets holds binary project files (textures, sounds) the browser
	// sent with mms/setAssets, since it has no disk for Texture() to read.
	assets   fstest.MapFS
	shutdown bool
}

// NewServer returns a server over tr checking against pack.
func NewServer(tr Transport, pack *spec.Pack) *Server {
	return &Server{tr: tr, pack: pack, open: map[string]bool{}, managed: map[string]bool{}, published: map[string]bool{}}
}

type message struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id,omitempty"`
	Method  string          `json:"method,omitempty"`
	Params  json.RawMessage `json:"params,omitempty"`
	Result  any             `json:"result,omitempty"`
	Error   *respError      `json:"error,omitempty"`
}

type respError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *respError) Error() string { return e.Message }

// JSON-RPC and LSP error codes.
const (
	codeParse          = -32700
	codeInvalidRequest = -32600
	codeMethodNotFound = -32601
	codeInvalidParams  = -32602
	codeInternal       = -32603
	codeNotInitialized = -32002
	codeRequestFailed  = -32803
)

// Run serves until the client sends exit or the transport closes. It
// returns nil after a clean shutdown.
func (s *Server) Run(ctx context.Context) error {
	for {
		raw, err := s.tr.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}
		var m message
		if err := json.Unmarshal(raw, &m); err != nil {
			s.reply(nil, nil, &respError{Code: codeParse, Message: err.Error()})
			continue
		}
		if m.Method == "exit" {
			if !s.shutdown {
				return errors.New("lsp: exit without shutdown")
			}
			return nil
		}
		if m.Method == "" {
			continue // a response to a request we sent
		}
		result, rerr := s.dispatch(ctx, m)
		if m.ID != nil {
			s.reply(m.ID, result, rerr)
		}
	}
}

// dispatch runs one handler; a panic becomes an internal error so the
// server stays up.
func (s *Server) dispatch(ctx context.Context, m message) (result any, rerr *respError) {
	defer func() {
		if r := recover(); r != nil {
			s.logf("panic in %s: %v", m.Method, r)
			result, rerr = nil, &respError{Code: codeInternal, Message: fmt.Sprintf("internal error: %v", r)}
		}
	}()
	if s.projects == nil && m.Method != "initialize" {
		if m.ID == nil {
			return nil, nil
		}
		return nil, &respError{Code: codeNotInitialized, Message: "server not initialized"}
	}
	h, ok := handlers[m.Method]
	if !ok {
		if m.ID == nil || strings.HasPrefix(m.Method, "$/") {
			return nil, nil
		}
		return nil, &respError{Code: codeMethodNotFound, Message: "method not found: " + m.Method}
	}
	return h(s, ctx, m.Params)
}

func (s *Server) reply(id json.RawMessage, result any, rerr *respError) {
	m := map[string]any{"jsonrpc": "2.0", "id": id}
	if id == nil {
		m["id"] = nil
	}
	if rerr != nil {
		m["error"] = rerr
	} else {
		m["result"] = result
	}
	s.send(m)
}

func (s *Server) notify(method string, params any) {
	s.send(map[string]any{"jsonrpc": "2.0", "method": method, "params": params})
}

func (s *Server) send(m any) {
	b, err := json.Marshal(m)
	if err != nil {
		s.logf("marshal: %v", err)
		return
	}
	if err := s.tr.Write(b); err != nil {
		s.logf("write: %v", err)
	}
}

func (s *Server) logf(format string, args ...any) {
	if s.Log != nil {
		fmt.Fprintf(s.Log, "mms lsp: "+format+"\n", args...)
	}
}

type handler func(s *Server, ctx context.Context, params json.RawMessage) (any, *respError)

var handlers map[string]handler

func init() {
	handlers = map[string]handler{
		"initialize":  typed((*Server).initialize),
		"initialized": func(*Server, context.Context, json.RawMessage) (any, *respError) { return nil, nil },
		"shutdown": func(s *Server, _ context.Context, _ json.RawMessage) (any, *respError) {
			s.shutdown = true
			return nil, nil
		},
		"textDocument/didOpen":             typed((*Server).didOpen),
		"textDocument/didChange":           typed((*Server).didChange),
		"textDocument/didClose":            typed((*Server).didClose),
		"textDocument/didSave":             func(*Server, context.Context, json.RawMessage) (any, *respError) { return nil, nil },
		"textDocument/completion":          typed((*Server).completion),
		"textDocument/hover":               typed((*Server).hover),
		"textDocument/signatureHelp":       typed((*Server).signatureHelp),
		"textDocument/definition":          typed((*Server).definition),
		"textDocument/references":          typed((*Server).references),
		"textDocument/prepareRename":       typed((*Server).prepareRename),
		"textDocument/rename":              typed((*Server).rename),
		"textDocument/formatting":          typed((*Server).formatting),
		"textDocument/documentSymbol":      typed((*Server).documentSymbol),
		"textDocument/semanticTokens/full": typed((*Server).semanticTokens),
		"workspace/didChangeWatchedFiles":  typed((*Server).didChangeWatched),
		"mms/output":                       typed((*Server).output),
		"textDocument/codeAction":          typed((*Server).codeAction),
		"mms/setConfig":                    typed((*Server).setConfig),
		"mms/setFiles":                     typed((*Server).setFiles),
		"mms/setAssets":                    typed((*Server).setAssets),
		"mms/importPack":                   typed((*Server).importPack),
		"mms/resourceKinds":                typed((*Server).resourceKinds),
		"mms/convertJSON":                  typed((*Server).convertJSON),
		"mms/ping":                         func(*Server, context.Context, json.RawMessage) (any, *respError) { return "pong", nil },
	}
}

// typed adapts a handler with decoded params.
func typed[P any](f func(*Server, context.Context, P) (any, *respError)) handler {
	return func(s *Server, ctx context.Context, raw json.RawMessage) (any, *respError) {
		var p P
		if len(raw) > 0 {
			if err := json.Unmarshal(raw, &p); err != nil {
				return nil, &respError{Code: codeInvalidParams, Message: err.Error()}
			}
		}
		return f(s, ctx, p)
	}
}

// ---- lifecycle

type initializeParams struct {
	RootURI          string `json:"rootUri"`
	RootPath         string `json:"rootPath"`
	WorkspaceFolders []struct {
		URI string `json:"uri"`
	} `json:"workspaceFolders"`
}

func (s *Server) initialize(ctx context.Context, p initializeParams) (any, *respError) {
	s.roots = nil
	for _, f := range p.WorkspaceFolders {
		if path := uriPath(f.URI); path != "" {
			s.roots = append(s.roots, path)
		}
	}
	if len(s.roots) == 0 {
		if path := uriPath(p.RootURI); path != "" {
			s.roots = []string{path}
		} else if p.RootPath != "" {
			s.roots = []string{p.RootPath}
		}
	}
	s.discover(project.DefaultConfig())
	s.analyzeAll(ctx)
	return map[string]any{
		"capabilities": map[string]any{
			"positionEncoding": "utf-16",
			"textDocumentSync": map[string]any{"openClose": true, "change": 2, "save": map[string]any{"includeText": false}},
			"completionProvider": map[string]any{
				"triggerCharacters": []string{".", "(", ",", "=", "[", "#", ":"},
			},
			"hoverProvider":              true,
			"signatureHelpProvider":      map[string]any{"triggerCharacters": []string{"(", ","}, "retriggerCharacters": []string{","}},
			"definitionProvider":         true,
			"referencesProvider":         true,
			"renameProvider":             map[string]any{"prepareProvider": true},
			"documentFormattingProvider": true,
			"documentSymbolProvider":     true,
			"codeActionProvider":         map[string]any{"codeActionKinds": []string{"quickfix", "source.fixAll.mms"}},
			"semanticTokensProvider": map[string]any{
				"legend": map[string]any{"tokenTypes": ide.TokenTypes, "tokenModifiers": ide.TokenModifiers},
				"full":   true,
			},
		},
		"serverInfo": map[string]any{"name": "mms"},
	}, nil
}

// ---- documents

type textDocumentItem struct {
	URI  string `json:"uri"`
	Text string `json:"text"`
}

func (d *textDocumentItem) UnmarshalJSON(b []byte) error {
	type plain textDocumentItem
	err := json.Unmarshal(b, (*plain)(d))
	d.URI = canonURI(d.URI)
	return err
}

type docID struct {
	URI string `json:"uri"`
}

func (d *docID) UnmarshalJSON(b []byte) error {
	type plain docID
	err := json.Unmarshal(b, (*plain)(d))
	d.URI = canonURI(d.URI)
	return err
}

func (s *Server) didOpen(ctx context.Context, p struct {
	TextDocument textDocumentItem `json:"textDocument"`
}) (any, *respError) {
	s.open[p.TextDocument.URI] = true
	w := s.setFile(p.TextDocument.URI, p.TextDocument.Text)
	s.last, s.lastURI = w, p.TextDocument.URI
	s.analyze(ctx, w)
	return nil, nil
}

func (s *Server) didChange(ctx context.Context, p struct {
	TextDocument   docID `json:"textDocument"`
	ContentChanges []struct {
		Range *Range `json:"range"`
		Text  string `json:"text"`
	} `json:"contentChanges"`
}) (any, *respError) {
	var text string
	if w := s.at(p.TextDocument.URI); w != nil {
		text, _ = w.proj.Text(p.TextDocument.URI)
	}
	for _, c := range p.ContentChanges {
		if c.Range == nil {
			text = c.Text
			continue
		}
		// Incremental edits are applied too, in case a client sends them.
		l := newLines(text)
		start, end := l.offset(c.Range.Start), l.offset(c.Range.End)
		src := []rune(text)
		text = string(src[:start]) + c.Text + string(src[end:])
	}
	w := s.setFile(p.TextDocument.URI, text)
	s.last, s.lastURI = w, p.TextDocument.URI
	s.analyze(ctx, w)
	return nil, nil
}

func (s *Server) didClose(ctx context.Context, p struct {
	TextDocument docID `json:"textDocument"`
}) (any, *respError) {
	uri := p.TextDocument.URI
	delete(s.open, uri)
	if s.managed[uri] {
		return nil, nil
	}
	w := s.at(uri)
	if w == nil {
		return nil, nil
	}
	// A file that belongs to the project on disk stays, with its saved
	// text; anything else leaves the project.
	if path := uriPath(uri); path != "" && w.contains(path) {
		if b, err := os.ReadFile(path); err == nil {
			w.proj.SetFile(uri, string(b))
			s.analyze(ctx, w)
			return nil, nil
		}
	}
	s.removeFile(uri)
	s.analyze(ctx, w)
	return nil, nil
}

func (s *Server) didChangeWatched(ctx context.Context, p struct {
	Changes []struct {
		URI  string `json:"uri"`
		Type int    `json:"type"` // 1 created, 2 changed, 3 deleted
	} `json:"changes"`
}) (any, *respError) {
	configChanged := false
	for _, c := range p.Changes {
		c.URI = canonURI(c.URI)
		path := uriPath(c.URI)
		if filepath.Base(path) == project.ConfigFile {
			configChanged = true
			continue
		}
		if s.open[c.URI] {
			continue
		}
		if c.Type == 3 {
			s.removeFile(c.URI)
			continue
		}
		if path != "" && strings.HasSuffix(path, ".mms") {
			if w := s.projectFor(c.URI); w.contains(path) {
				if b, err := os.ReadFile(path); err == nil {
					w.proj.SetFile(c.URI, string(b))
					s.owners[c.URI] = w
				}
			}
		}
	}
	if configChanged {
		// Files may move between projects, so start over.
		s.reload(ctx)
		return nil, nil
	}
	// Assets changed too, so every project is checked again.
	s.analyzeAll(ctx)
	return nil, nil
}

// setConfig replaces the default project's configuration; the browser
// editor sends it when mms.json changes, since it has no disk to read.
func (s *Server) setConfig(ctx context.Context, cfg project.Config) (any, *respError) {
	if cfg.Namespace == "" {
		cfg.Namespace = project.DefaultConfig().Namespace
	}
	if err := cfg.Validate(); err != nil {
		return nil, &respError{Code: codeInvalidParams, Message: err.Error()}
	}
	w := s.projects[""]
	np := project.New(s.pack, cfg)
	np.Files = w.proj.Files
	for _, path := range w.proj.Paths() {
		t, _ := w.proj.Text(path)
		np.SetFile(path, t)
	}
	w.proj = np
	s.analyze(ctx, w)
	return nil, nil
}

// setFiles adds, replaces and removes project files that are not open in
// an editor. The browser sends the whole project this way on start.
func (s *Server) setFiles(ctx context.Context, p struct {
	Files []struct {
		URI  string `json:"uri"`
		Text string `json:"text"`
	} `json:"files"`
	Remove []string `json:"remove"`
}) (any, *respError) {
	for _, f := range p.Files {
		s.managed[f.URI] = true
		if !s.open[f.URI] {
			s.setFile(f.URI, f.Text)
		}
	}
	for _, uri := range p.Remove {
		delete(s.managed, uri)
		delete(s.open, uri)
		s.removeFile(uri)
	}
	s.analyzeAll(ctx)
	return nil, nil
}

// setAssets adds, replaces and removes binary project files, by path
// relative to the project; data is base64 in JSON. The browser sends them
// because it has no disk; asset helpers such as Texture("a.png") read them.
func (s *Server) setAssets(ctx context.Context, p struct {
	Files []struct {
		Path string `json:"path"`
		Data []byte `json:"data"`
	} `json:"files"`
	Remove []string `json:"remove"`
}) (any, *respError) {
	if s.assets == nil {
		s.assets = fstest.MapFS{}
	}
	now := time.Now()
	for _, f := range p.Files {
		s.assets[path.Clean(f.Path)] = &fstest.MapFile{Data: f.Data, ModTime: now}
	}
	for _, rel := range p.Remove {
		delete(s.assets, path.Clean(rel))
	}
	// The browser's files are in the default project. Checks cached
	// against the old files see the new times and re-run.
	w := s.projects[""]
	w.proj.Files = s.assets
	s.analyze(ctx, w)
	return nil, nil
}

// analyze re-checks project w and publishes diagnostics for its files.
func (s *Server) analyze(ctx context.Context, w *workspace) {
	r, err := w.proj.Analyze(ctx)
	if err != nil {
		return
	}
	w.res = r
	seen := map[string]bool{}
	for _, u := range r.Units {
		seen[u.Path] = true
		ds := r.Diagnostics(u.Path)
		if len(ds) == 0 && !s.published[u.Path] {
			continue
		}
		s.publish(u, ds)
	}
	for uri := range s.published {
		// Files that left every project are cleared too.
		if o := s.owners[uri]; (o == w || o == nil) && !seen[uri] {
			s.notify("textDocument/publishDiagnostics", map[string]any{"uri": uri, "diagnostics": []any{}})
			delete(s.published, uri)
		}
	}
}

func (s *Server) publish(u *sema.Unit, ds []syntax.Diagnostic) {
	l := newLines(string(u.File.Src))
	out := make([]map[string]any, 0, len(ds))
	for _, d := range ds {
		m := map[string]any{"range": l.rng(d.Sp), "severity": int(d.Severity), "source": "mms", "message": d.Message}
		if d.Code != "" {
			m["code"] = d.Code
		}
		switch {
		case d.Unnecessary:
			m["tags"] = []int{1}
		case d.Deprecated:
			m["tags"] = []int{2}
		}
		if len(d.Fixes) > 0 {
			// The web editor applies fixes straight from the diagnostic.
			m["data"] = map[string]any{"fixes": fixesJSON(l, d.Fixes)}
		}
		out = append(out, m)
	}
	s.published[u.Path] = len(out) > 0
	s.notify("textDocument/publishDiagnostics", map[string]any{"uri": u.Path, "diagnostics": out})
}

// ---- requests

type posParams struct {
	TextDocument docID    `json:"textDocument"`
	Position     Position `json:"position"`
}

// locate returns the session of the file's project, the file's line index
// and the rune offset.
func (s *Server) locate(p posParams) (*ide.Session, *lines, int, bool) {
	w := s.at(p.TextDocument.URI)
	if w == nil || w.res == nil {
		return nil, nil, 0, false
	}
	u := w.res.Unit(p.TextDocument.URI)
	if u == nil {
		return nil, nil, 0, false
	}
	l := newLines(string(u.File.Src))
	return w.session(), l, l.offset(p.Position), true
}

func (w *workspace) session() *ide.Session { return &ide.Session{P: w.proj, R: w.res} }

// unit is the analyzed file uri, or nil.
func (s *Server) unit(uri string) (*workspace, *sema.Unit) {
	w := s.at(uri)
	if w == nil || w.res == nil {
		return nil, nil
	}
	return w, w.res.Unit(uri)
}

func (s *Server) lines(uri string) *lines {
	if _, u := s.unit(uri); u != nil {
		return newLines(string(u.File.Src))
	}
	return newLines("")
}

var completionKinds = map[sema.CandidateKind]int{
	sema.CandField: 10, sema.CandCallable: 3, sema.CandMethod: 2, sema.CandID: 18,
	sema.CandEnum: 20, sema.CandKeyword: 14, sema.CandKind: 7,
}

func (s *Server) completion(_ context.Context, p posParams) (any, *respError) {
	sess, l, off, ok := s.locate(p)
	if !ok {
		return nil, nil
	}
	items, replace := sess.Completion(p.TextDocument.URI, off)
	rng := l.rng(replace)
	out := make([]map[string]any, 0, len(items))
	for _, it := range items {
		m := map[string]any{"label": it.Label, "kind": completionKinds[it.Kind], "sortText": it.Sort}
		if it.Detail != "" {
			m["detail"] = it.Detail
		}
		if it.Doc != "" {
			m["documentation"] = map[string]any{"kind": "markdown", "value": it.Doc}
		}
		if it.Deprecated {
			m["tags"] = []int{1} // CompletionItemTag.Deprecated
		}
		text := it.Label
		if it.Snippet {
			text = it.Insert
			m["insertTextFormat"] = 2
		}
		m["textEdit"] = map[string]any{"range": rng, "newText": text}
		m["filterText"] = it.Label
		out = append(out, m)
	}
	return map[string]any{"isIncomplete": false, "items": out}, nil
}

func (s *Server) hover(_ context.Context, p posParams) (any, *respError) {
	sess, l, off, ok := s.locate(p)
	if !ok {
		return nil, nil
	}
	md, sp := sess.Hover(p.TextDocument.URI, off)
	if md == "" {
		return nil, nil
	}
	return map[string]any{"contents": map[string]any{"kind": "markdown", "value": md}, "range": l.rng(sp)}, nil
}

func (s *Server) signatureHelp(_ context.Context, p posParams) (any, *respError) {
	sess, _, off, ok := s.locate(p)
	if !ok {
		return nil, nil
	}
	h, ok := sess.Signature(p.TextDocument.URI, off)
	if !ok {
		return nil, nil
	}
	ps := make([]map[string]any, 0, len(h.Params))
	for i, pr := range h.Params {
		m := map[string]any{"label": pr}
		if i < len(h.ParamDocs) && h.ParamDocs[i] != "" {
			m["documentation"] = map[string]any{"kind": "markdown", "value": h.ParamDocs[i]}
		}
		ps = append(ps, m)
	}
	sig := map[string]any{"label": h.Label, "parameters": ps}
	if h.Doc != "" {
		sig["documentation"] = map[string]any{"kind": "markdown", "value": h.Doc}
	}
	return map[string]any{
		"signatures":      []any{sig},
		"activeSignature": 0, "activeParameter": h.Active,
	}, nil
}

type location struct {
	URI   string `json:"uri"`
	Range Range  `json:"range"`
}

func (s *Server) locations(locs []ide.Location) []location {
	out := make([]location, 0, len(locs))
	idx := map[string]*lines{}
	for _, loc := range locs {
		l := idx[loc.Path]
		if l == nil {
			l = s.lines(loc.Path)
			idx[loc.Path] = l
		}
		out = append(out, location{URI: loc.Path, Range: l.rng(loc.Span)})
	}
	return out
}

func (s *Server) definition(_ context.Context, p posParams) (any, *respError) {
	sess, _, off, ok := s.locate(p)
	if !ok {
		return nil, nil
	}
	return s.locations(sess.Definition(p.TextDocument.URI, off)), nil
}

func (s *Server) references(_ context.Context, p struct {
	posParams
	Context struct {
		IncludeDeclaration bool `json:"includeDeclaration"`
	} `json:"context"`
}) (any, *respError) {
	sess, _, off, ok := s.locate(p.posParams)
	if !ok {
		return nil, nil
	}
	return s.locations(sess.References(p.TextDocument.URI, off, p.Context.IncludeDeclaration)), nil
}

func (s *Server) prepareRename(_ context.Context, p posParams) (any, *respError) {
	sess, l, off, ok := s.locate(p)
	if !ok {
		return nil, nil
	}
	sp, ok := sess.PrepareRename(p.TextDocument.URI, off)
	if !ok {
		return nil, nil
	}
	return l.rng(sp), nil
}

func (s *Server) rename(_ context.Context, p struct {
	posParams
	NewName string `json:"newName"`
}) (any, *respError) {
	sess, _, off, ok := s.locate(p.posParams)
	if !ok {
		return nil, nil
	}
	edits, msg := sess.Rename(p.TextDocument.URI, off, p.NewName)
	if msg != "" {
		return nil, &respError{Code: codeRequestFailed, Message: msg}
	}
	changes := map[string][]map[string]any{}
	idx := map[string]*lines{}
	for _, e := range edits {
		l := idx[e.Path]
		if l == nil {
			l = s.lines(e.Path)
			idx[e.Path] = l
		}
		changes[e.Path] = append(changes[e.Path], map[string]any{"range": l.rng(e.Span), "newText": e.Text})
	}
	return map[string]any{"changes": changes}, nil
}

func (s *Server) formatting(_ context.Context, p struct {
	TextDocument docID `json:"textDocument"`
}) (any, *respError) {
	w := s.at(p.TextDocument.URI)
	if w == nil {
		return nil, nil
	}
	text, ok := w.proj.Text(p.TextDocument.URI)
	if !ok {
		return nil, nil
	}
	out, err := format.Source(text)
	if err != nil {
		// Files with syntax errors are left alone; the diagnostics say why.
		return nil, nil
	}
	if out == text {
		return []any{}, nil
	}
	l := newLines(text)
	return []any{map[string]any{"range": Range{Start: Position{}, End: l.pos(len(l.src))}, "newText": out}}, nil
}

var symbolKinds = map[string]int{"resource": 23, "function": 12, "namespace": 3, "tag": 18, "objective": 13}

func (s *Server) documentSymbol(_ context.Context, p struct {
	TextDocument docID `json:"textDocument"`
}) (any, *respError) {
	w := s.at(p.TextDocument.URI)
	if w == nil || w.res == nil {
		return []any{}, nil
	}
	l := s.lines(p.TextDocument.URI)
	syms := w.session().Symbols(p.TextDocument.URI)
	out := make([]map[string]any, 0, len(syms))
	for _, sym := range syms {
		out = append(out, map[string]any{
			"name": sym.Name, "detail": sym.Detail, "kind": symbolKinds[sym.Kind],
			"range": l.rng(sym.Span), "selectionRange": l.rng(sym.Sel),
		})
	}
	return out, nil
}

func (s *Server) semanticTokens(_ context.Context, p struct {
	TextDocument docID `json:"textDocument"`
}) (any, *respError) {
	w := s.at(p.TextDocument.URI)
	if w == nil || w.res == nil {
		return map[string]any{"data": []int{}}, nil
	}
	l := s.lines(p.TextDocument.URI)
	toks := w.session().SemanticTokens(p.TextDocument.URI)
	return map[string]any{"data": encodeTokens(l, toks)}, nil
}

// encodeTokens produces LSP's relative encoding, splitting tokens that
// span lines (clients need not support multi-line tokens).
func encodeTokens(l *lines, toks []ide.SemToken) []int {
	data := make([]int, 0, len(toks)*5)
	prevLine, prevChar := 0, 0
	emit := func(start, end Position, typ, mods int) {
		if end.Character <= start.Character {
			return
		}
		dl := start.Line - prevLine
		dc := start.Character
		if dl == 0 {
			dc -= prevChar
		}
		data = append(data, dl, dc, end.Character-start.Character, typ, mods)
		prevLine, prevChar = start.Line, start.Character
	}
	for _, t := range toks {
		start, end := l.pos(t.Span.Start), l.pos(t.Span.End)
		for line := start.Line; line <= end.Line; line++ {
			a, b := Position{Line: line}, Position{Line: line}
			if line == start.Line {
				a = start
			}
			if line == end.Line {
				b = end
			} else {
				b = l.pos(lineEnd(l, line))
			}
			emit(a, b, int(t.Type), t.Mods)
		}
	}
	return data
}

func lineEnd(l *lines, line int) int {
	if line+1 < len(l.starts) {
		e := l.starts[line+1] - 1
		if e > 0 && l.src[e-1] == '\r' {
			e--
		}
		return e
	}
	return len(l.src)
}

// output returns the generated pack for previews: every file's path and
// text, with the declaration it came from. Copied binary files have no
// text; they are listed in assets with the project file they copy, so the
// browser can put its own copy in a download. It is the pack of the project
// holding uri, or without one of the last edited file's project.
func (s *Server) output(_ context.Context, p struct {
	URI string `json:"uri"`
}) (any, *respError) {
	w := s.at(canonURI(p.URI))
	if w == nil {
		w = s.last
	}
	if w == nil || w.res == nil {
		w = s.projects[""]
	}
	if w.res == nil {
		return map[string]any{"files": []any{}, "assets": []any{}, "hasErrors": false}, nil
	}
	files := w.res.Emit()
	out := make([]map[string]any, 0, len(files))
	assets := []map[string]any{}
	idx := map[string]*lines{}
	for _, f := range files {
		if f.Asset != "" {
			a := map[string]any{"path": f.Path, "src": f.Asset}
			if f.Pack == sema.ResourcePack {
				a["pack"] = "resources"
			}
			assets = append(assets, a)
			continue
		}
		m := map[string]any{"path": f.Path, "text": string(f.Data)}
		if f.Pack == sema.ResourcePack {
			m["pack"] = "resources"
		}
		if f.Unit != "" {
			l := idx[f.Unit]
			if l == nil {
				l = s.lines(f.Unit)
				idx[f.Unit] = l
			}
			m["uri"], m["range"] = f.Unit, l.rng(f.Span)
		}
		out = append(out, m)
	}
	sort.Slice(out, func(i, j int) bool { return out[i]["path"].(string) < out[j]["path"].(string) })
	return map[string]any{"files": out, "assets": assets, "hasErrors": w.res.HasErrors()}, nil
}

// ---- URIs

// uriPath turns a file URI into a path. A Windows drive keeps its letter:
// file:///c%3A/x is c:/x, not /c:/x.
func uriPath(uri string) string {
	u, err := url.Parse(uri)
	if err != nil || u.Scheme != "file" {
		return ""
	}
	p := u.Path
	if len(p) >= 3 && p[0] == '/' && p[2] == ':' && isLetter(p[1]) {
		p = p[1:]
	}
	return filepath.FromSlash(p)
}

// pathURI turns a path into a file URI. A Windows path gets a leading
// slash, so its drive is part of the path rather than the host.
func pathURI(path string) string {
	p := filepath.ToSlash(path)
	if !strings.HasPrefix(p, "/") {
		p = "/" + p
	}
	return (&url.URL{Scheme: "file", Path: p}).String()
}

// canonURI spells a file URI the way pathURI does, so a file the client
// opens and the same file found on disk share one key. Other schemes,
// such as the browser editor's mem:, pass through.
func canonURI(uri string) string {
	if p := uriPath(uri); p != "" {
		return pathURI(p)
	}
	return uri
}

func isLetter(c byte) bool { return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' }
