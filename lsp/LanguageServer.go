package lsp

import (
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/minecraftmetascript/mms/lang"
	"github.com/tliron/commonlog"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"github.com/tliron/glsp/server"
)

type LanguageServer struct {
	log commonlog.Logger

	name    string
	version string
	handler *protocol.Handler
	project *lang.Project

	documents      map[string]*document
	activeDocument *document
}

var ls *LanguageServer

func init() {
	ls = &LanguageServer{
		log:       commonlog.NewBackendLogger(),
		name:      "Minecraft Metascript",
		version:   "0.3.1",
		project:   lang.NewProject(),
		documents: make(map[string]*document),
	}
	ls.handler = &protocol.Handler{
		Initialize:                 ls.Initialize,
		Initialized:                ls.Initialized,
		Shutdown:                   ls.Shutdown,
		LogTrace:                   ls.LogTrace,
		SetTrace:                   ls.setTrace,
		TextDocumentDidOpen:        ls.TextDocumentDidOpen,
		TextDocumentDidChange:      ls.TextDocumentDidChange,
		TextDocumentDocumentSymbol: ls.DocumentSymbols,
		TextDocumentFoldingRange:   ls.TextDocumentFoldingRange,
		TextDocumentHover:          ls.TextDocumentHover,
	}
}

func StartStreaming(stream io.ReadWriteCloser) error {
	serve := server.NewServer(ls.handler, ls.name, true)
	serve.ServeStream(stream, serve.Log)
	return nil //?
}

func Start() error {
	serve := server.NewServer(ls.handler, ls.name, true)
	ls.log = serve.Log

	err := serve.RunStdio()
	if err != nil {
		serve.Log.Error(fmt.Sprintf("Failed to start language server: %s", err.Error()))
	} else {
		serve.Log.Info("Language server started")
	}
	return err
}

func (ls *LanguageServer) Initialize(context *glsp.Context, params *protocol.InitializeParams) (any, error) {
	capabilities := ls.handler.CreateServerCapabilities()

	return protocol.InitializeResult{
		Capabilities: capabilities,
		ServerInfo: &protocol.InitializeResultServerInfo{
			Name:    ls.name,
			Version: &ls.version,
		},
	}, nil
}

func (ls *LanguageServer) Initialized(context *glsp.Context, params *protocol.InitializedParams) error {
	ls.log.Info("Initialized")
	return nil
}

func (ls *LanguageServer) Shutdown(context *glsp.Context) error {
	ls.log.Info("Stopping...")
	protocol.SetTraceValue(protocol.TraceValueOff)
	return nil
}

func (ls *LanguageServer) setTrace(context *glsp.Context, params *protocol.SetTraceParams) error {
	protocol.SetTraceValue(params.Value)
	return nil
}

func (ls *LanguageServer) TextDocumentDidChange(context *glsp.Context, params *protocol.DidChangeTextDocumentParams) error {
	ls.log.Info("TextDocumentDidChange", params.TextDocument.URI)
	doc, ok := ls.documents[params.TextDocument.URI]
	if !ok {
		ls.log.Error("no document for URI")
		return errors.New("no document for URI")
	}
	if doc.ApplyChanges(params.ContentChanges) != nil {
		return errors.New("failed to apply changes")
	}
	doc.PublishDiagnostics(context)
	return nil
}

func (ls *LanguageServer) TextDocumentDidOpen(context *glsp.Context, params *protocol.DidOpenTextDocumentParams) error {
	ls.log.Info("TextDocumentDidOpen", params.TextDocument.Text)
	doc, err := newDocument(
		params.TextDocument.Text,
		params.TextDocument.URI,
		ls.project,
	)
	if err != nil {
		return err
	}
	ls.documents[params.TextDocument.URI] = doc
	doc.PublishDiagnostics(context)
	ls.activeDocument = doc
	return nil
}

func (ls *LanguageServer) TextDocumentFoldingRange(context *glsp.Context, params *protocol.FoldingRangeParams) ([]protocol.FoldingRange, error) {
	visitor := &FoldingRangeVisitor{
		foldingRanges: make([]protocol.FoldingRange, 0),
	}

	fmt.Println("Attempting to make folding ranges...")
	visitor.EnterEveryRule(ls.activeDocument.file.Script)
	ls.activeDocument.file.Script.EnterRule(visitor)
	return visitor.foldingRanges, nil

}

func (ls *LanguageServer) LogTrace(context *glsp.Context, params *protocol.LogTraceParams) error {
	log.Println(params.Message)
	return nil
}

func (ls *LanguageServer) TextDocumentHover(context *glsp.Context, params *protocol.HoverParams) (*protocol.Hover, error) {
	out := &protocol.Hover{}
	help := ls.activeDocument.HelpAtPosition(params.TextDocumentPositionParams.Position, params.TextDocument.URI)
	if help != nil {
		out.Contents = help.Content
		out.Range = &protocol.Range{
			Start: protocol.Position{
				Line:      protocol.UInteger(help.Position.Start.Line), // LSP is 0-indexed
				Character: protocol.UInteger(help.Position.Start.Col),
			},
			End: protocol.Position{
				Line:      protocol.UInteger(help.Position.Stop.Line), // LSP is 0-indexed
				Character: protocol.UInteger(help.Position.Stop.Col),
			},
		}
		return out, nil
	}

	return nil, nil
}
