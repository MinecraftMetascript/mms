package lsp

import (
	"errors"
	"fmt"
	"io"
	"log"

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
	log.SetOutput(commonlog.GetWriter())
	log.Println("Output from log package!")

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
	ls.log.Infof("%s", ls.project.Symbols())
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

	ls.PublishDiagnostics(context, f)
	return out, nil
}

func (ls *LanguageServer) TextDocumentDidOpen(ctx *glsp.Context, params *protocol.DidOpenTextDocumentParams) error {
	initialContent := params.TextDocument.Text
	path := string(params.TextDocument.URI)

	_, err := ls.project.AddFile(path, initialContent)
	return err
}

func (ls *LanguageServer) Initialize(context *glsp.Context, params *protocol.InitializeParams) (any, error) {
	capabilities := ls.handler.CreateServerCapabilities()
	//capabilities.CompletionProvider.TriggerCharacters = []string{".", "(", ":"}

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

	_, err := ls.project.AddFile(params.TextDocument.URI, content)
	return err
}

func (ls *LanguageServer) TextDocumentHover(context *glsp.Context, params *protocol.HoverParams) (*protocol.Hover, error) {
	// TODO: Use AST() (map of location -> node) to discover nodes at the cursor. iterate from least to most specific and return the first one that matches.
	file := ls.project.File(params.TextDocument.URI)
	candidates := file.NodesAtPosition(
		ast.Location{
			Line:   int(params.Position.Line + 1),
			Column: int(params.Position.Character),
			Index:  params.Position.IndexIn(file.Content()),
		},
	)

	if len(candidates) > 0 {
		for _, candidate := range candidates {
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
