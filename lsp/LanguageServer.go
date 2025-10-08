package lsp

import (
	"errors"
	"fmt"
	"io"
	"log"
	"slices"
	"strings"

	"github.com/minecraftmetascript/mms/lang/ast"
	"github.com/minecraftmetascript/mms/project"
	"github.com/tliron/commonlog"
	_ "github.com/tliron/commonlog/simple"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"github.com/tliron/glsp/server"
)

type LanguageServer struct {
	log commonlog.Logger

	name    string
	version string
	handler *protocol.Handler
	project *project.Project
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type logTranslator struct {
	l commonlog.Logger
}

func (l logTranslator) Write(p []byte) (n int, err error) {
	out := strings.Trim(string(p), "\n")
	l.l.Info(out)
	return len(out), nil
}

var ls *LanguageServer

func init() {

	ls = &LanguageServer{
		log:     commonlog.NewBackendLogger("MMS"),
		name:    "Minecraft Metascript",
		version: "0.4.0",
		project: project.NewProject(),
	}
	ls.handler = &protocol.Handler{
		Initialize:                 ls.Initialize,
		Initialized:                ls.Initialized,
		Shutdown:                   ls.Shutdown,
		TextDocumentDidOpen:        ls.TextDocumentDidOpen,
		TextDocumentDidChange:      ls.TextDocumentDidChange,
		TextDocumentDocumentSymbol: ls.TextDocumentDocumentSymbol,
		TextDocumentHover:          ls.TextDocumentHover,
		TextDocumentCompletion:     ls.TextDocumentCompletion,
	}

}

func StartStreaming(stream io.ReadWriteCloser) error {
	serve := server.NewServer(ls.handler, ls.name, false)
	serve.ServeStream(stream, serve.Log)
	return nil //?
}

func Start() error {
	commonlog.Configure(1, nil)

	serve := server.NewServer(ls.handler, ls.name, false)
	ls.log = serve.Log
	ls.log.SetMaxLevel(commonlog.Level(6))
	ls.log.Info("Language Server Starting...")
	log.SetOutput(logTranslator{ls.log})
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	return serve.RunStdio()
}
func (ls *LanguageServer) TextDocumentDocumentSymbol(context *glsp.Context, params *protocol.DocumentSymbolParams) (any, error) {
	path := params.TextDocument.URI

	f := ls.project.File(path)
	if f == nil {
		ls.log.Warningf("Tried to open %s -- could not find file in project", path)
		return nil, nil
	}

	out := make([]protocol.DocumentSymbol, 0)
	for ns, decls := range ls.project.Symbols() {
		for name, decl := range decls.AllDecls() {
			if decl.GetLocation().Filename == path {
				out = append(out, protocol.DocumentSymbol{
					Name:           fmt.Sprintf("%s:%s", ns, name),
					Detail:         nil,
					Kind:           protocol.SymbolKindVariable,
					Range:          decl.GetLocation().ToLspRange(),
					SelectionRange: decl.GetNameLocation().ToLspRange(),
					Children:       nil,
				})
			}
		}
	}

	return out, nil
}

func (ls *LanguageServer) TextDocumentDidOpen(ctx *glsp.Context, params *protocol.DidOpenTextDocumentParams) error {
	initialContent := params.TextDocument.Text
	path := params.TextDocument.URI

	f, err := ls.project.AddFile(path, initialContent)
	ls.PublishDiagnostics(ctx, f)
	return err
}

func (ls *LanguageServer) Initialize(context *glsp.Context, params *protocol.InitializeParams) (any, error) {
	capabilities := ls.handler.CreateServerCapabilities()
	capabilities.CompletionProvider.TriggerCharacters = []string{".", "(", ")", ":", "=", " ", ","}

	return protocol.InitializeResult{
		Capabilities: capabilities,
		ServerInfo: &protocol.InitializeResultServerInfo{
			Name:    ls.name,
			Version: &ls.version,
		},
	}, nil
}

func (ls *LanguageServer) Shutdown(context *glsp.Context) error {
	ls.log.Info("Stopping...")
	return nil
}

func (ls *LanguageServer) Initialized(context *glsp.Context, params *protocol.InitializedParams) error {
	return nil
}

func (ls *LanguageServer) TextDocumentDidChange(context *glsp.Context, params *protocol.DidChangeTextDocumentParams) error {
	doc := ls.project.File(params.TextDocument.URI)
	if doc == nil {
		ls.log.Error("no document for URI")
		return errors.New("no document for URI")
	}
	content := doc.Content()
	for _, change := range params.ContentChanges {
		switch c := change.(type) {
		case protocol.TextDocumentContentChangeEvent:
			startIndex, endIndex := c.Range.IndexesIn(content)
			content = content[:startIndex] + c.Text + content[endIndex:]
		case protocol.TextDocumentContentChangeEventWhole:
			content = c.Text
		}
	}

	f, err := ls.project.AddFile(params.TextDocument.URI, content)
	ls.PublishDiagnostics(context, f)
	return err
}

func (ls *LanguageServer) TextDocumentHover(context *glsp.Context, params *protocol.HoverParams) (*protocol.Hover, error) {
	file := ls.project.File(params.TextDocument.URI)
	candidates := positionToCandidateNodes(params.Position, file)

	if len(candidates) > 0 {
		for _, candidate := range candidates {
			log.Println(candidate.GetLocation())
			if helpful, ok := candidate.(ast.HelpfulNode); ok && helpful.GetHelp() != "" {
				return &protocol.Hover{
					Contents: protocol.MarkupContent{
						Kind:  protocol.MarkupKindMarkdown,
						Value: helpful.GetHelp(),
					},
				}, nil
			}
		}
	}

	return nil, nil
}

func positionToCandidateNodes(position protocol.Position, file *project.File) []ast.Node {
	rangeStart := ast.Location{
		Line:   int(position.Line + 1),
		Column: int(position.Character - 1),
	}
	rangeStart.Index = rangeStart.IndexIn(file.Content())

	rangeEnd := ast.Location{
		Line:   int(position.Line + 1),
		Column: int(position.Character),
	}
	rangeEnd.Index = rangeEnd.IndexIn(file.Content())

	candidates := file.NodesInRange(rangeStart, rangeEnd)

	slices.SortStableFunc(candidates, func(a, b ast.Node) int {
		aLoc := a.GetLocation()
		bLoc := b.GetLocation()

		if aLoc == nil || bLoc == nil {
			return 0
		}
		if aLoc.Contains(*bLoc) {
			return 1
		} else if bLoc.Contains(*aLoc) {
			return -1
		} else {
			return 0
		}
	})
	return candidates
}

func (ls *LanguageServer) TextDocumentCompletion(_ *glsp.Context, params *protocol.CompletionParams) (any, error) {
	file := ls.project.File(params.TextDocument.URI)
	candidates := positionToCandidateNodes(params.Position, file)

	if len(candidates) > 0 {
		for _, candidate := range candidates {
			if completable, ok := candidate.(ast.CompletableNode); ok {
				res := completable.Complete(file.Content(), params.Position, params.Context.TriggerCharacter, ls.project.Symbols())
				if res != nil {
					log.Println("Completion returned")
					return res, nil
				} else {
					log.Printf("Completion skip indicated with nil return from %T\n", completable)
				}
			} else {
				log.Println("Completion skip indicated with non-completable node")
			}
		}
	}

	return nil, nil
}
