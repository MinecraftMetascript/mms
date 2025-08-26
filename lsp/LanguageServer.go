package lsp

import (
	"errors"
	"log"
	"os"

	"github.com/minecraftmetascript/mms/lang"
	"github.com/minecraftmetascript/mms/lang/worldgen/surface/surface_rules"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"github.com/tliron/glsp/server"
)

var lsName = "MMS Language Server"
var version = "0.1.0"
var handler *protocol.Handler
var ls *server.Server

type LanguageServerCtx struct {
	project *lang.MMSProject
}

func Start() error {
	context := &LanguageServerCtx{
		project: lang.NewMMSProject(),
	}
	handler = &protocol.Handler{
		Initialize:                         context.initialize,
		Initialized:                        context.initialized,
		Shutdown:                           context.shutdown,
		SetTrace:                           context.setTrace,
		TextDocumentCompletion:             context.textDocumentCompletion,
		WorkspaceDidChangeWorkspaceFolders: context.workspaceDidChangeWorkspaceFolders,
		WorkspaceDidChangeWatchedFiles:     context.workspaceDidChangeWatchedFiles,
		TextDocumentDidChange:              context.textDocumentDidChange,
		TextDocumentDidOpen:                context.textDocumentDidOpen,
		TextDocumentDidSave:                context.textDocumentDidSave,
		TextDocumentDidClose:               context.textDocumentDidClose,
		CompletionItemResolve:              context.completionItemResolve,
	}

	ls = server.NewServer(handler, lsName, true)

	log.Println("Server launched")

	return ls.RunStdio()
}

func (ctx *LanguageServerCtx) reparseFile(path protocol.URI) error {
	filepath := path[7:]
	content, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	err = ctx.project.AddFile(filepath, string(content))
	if err != nil {
		return err
	}
	return nil
}

func (ctx *LanguageServerCtx) textDocumentDidClose(context *glsp.Context, params *protocol.DidCloseTextDocumentParams) error {
	return ctx.reparseFile(params.TextDocument.URI)
}
func (ctx *LanguageServerCtx) textDocumentDidSave(context *glsp.Context, params *protocol.DidSaveTextDocumentParams) error {
	return ctx.reparseFile(params.TextDocument.URI)
}
func (ctx *LanguageServerCtx) textDocumentDidOpen(context *glsp.Context, params *protocol.DidOpenTextDocumentParams) error {
	return ctx.reparseFile(params.TextDocument.URI)
}
func (ctx *LanguageServerCtx) textDocumentDidChange(context *glsp.Context, params *protocol.DidChangeTextDocumentParams) error {
	return ctx.reparseFile(params.TextDocument.URI)
}

func (ctx *LanguageServerCtx) workspaceDidChangeWorkspaceFolders(context *glsp.Context, params *protocol.DidChangeWorkspaceFoldersParams) error {
	log.Println("WorkspaceDidChangeWorkspaceFolders", "added", params.Event.Added, "removed", params.Event.Removed)

	return nil
}

func (ctx *LanguageServerCtx) workspaceDidChangeWatchedFiles(context *glsp.Context, params *protocol.DidChangeWatchedFilesParams) error {
	errs := make([]error, 0)
	for _, change := range params.Changes {
		err := ctx.reparseFile(change.URI)
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		log.Println("[ERROR]: ", errs)
	}

	for namespace, rules := range lang.ProjectSymbols[surface_rules.SurfaceRule](ctx.project) {
		log.Println(namespace)
		for name, rule := range rules {
			log.Println(name, rule)
		}

	}
	return errors.Join(errs...)
}

func (ctx *LanguageServerCtx) initialize(context *glsp.Context, params *protocol.InitializeParams) (any, error) {
	capabilities := handler.CreateServerCapabilities()

	return protocol.InitializeResult{
		Capabilities: capabilities,
		ServerInfo: &protocol.InitializeResultServerInfo{
			Name:    lsName,
			Version: &version,
		},
	}, nil
}

func (ctx *LanguageServerCtx) initialized(context *glsp.Context, params *protocol.InitializedParams) error {
	log.Println("Initialized [Completion Testing 1]")
	return nil
}

func (ctx *LanguageServerCtx) shutdown(context *glsp.Context) error {
	protocol.SetTraceValue(protocol.TraceValueOff)
	return nil
}

func (ctx *LanguageServerCtx) setTrace(context *glsp.Context, params *protocol.SetTraceParams) error {
	protocol.SetTraceValue(params.Value)
	return nil
}
