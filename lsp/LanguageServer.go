package lsp

import (
	"errors"
	"fmt"
	"io"
	"log"
	"slices"
	"strings"

	"github.com/minecraftmetascript/mms/lang/ast"
	"github.com/minecraftmetascript/mms/lang/spec"
	"github.com/minecraftmetascript/mms/project"
	"github.com/samber/lo"
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
		Initialize:                     ls.Initialize,
		Initialized:                    ls.Initialized,
		Shutdown:                       ls.Shutdown,
		TextDocumentDidOpen:            ls.TextDocumentDidOpen,
		TextDocumentDidChange:          ls.TextDocumentDidChange,
		TextDocumentDocumentSymbol:     ls.TextDocumentDocumentSymbol,
		TextDocumentHover:              ls.TextDocumentHover,
		TextDocumentCompletion:         ls.TextDocumentCompletion,
		TextDocumentDefinition:         ls.TextDocumentDefinition,
		TextDocumentReferences:         ls.TextDocumentReferences,
		TextDocumentSemanticTokensFull: ls.TextDocumentSemanticTokensFull,
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
func (ls *LanguageServer) TextDocumentDocumentSymbol(_ *glsp.Context, params *protocol.DocumentSymbolParams) (any, error) {
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
				v := protocol.DocumentSymbol{
					Name:     fmt.Sprintf("%s:%s", ns, name),
					Detail:   nil,
					Kind:     protocol.SymbolKindVariable,
					Range:    decl.GetLocation().ToLspRange(),
					Children: nil,
				}
				if nl := decl.GetNameLocation(); nl != nil {
					v.SelectionRange = nl.ToLspRange()
				} else {
					v.SelectionRange = v.Range
				}
				out = append(out, v)
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

func (ls *LanguageServer) Initialize(_ *glsp.Context, _ *protocol.InitializeParams) (any, error) {
	f := false
	capabilities := ls.handler.CreateServerCapabilities()
	capabilities.CompletionProvider.TriggerCharacters = []string{".", "(", ")", ":", "=", " ", ","}
	capabilities.SemanticTokensProvider = &protocol.SemanticTokensOptions{
		Legend: protocol.SemanticTokensLegend{
			TokenTypes: []string{
				"keyword", "variable", "function", "number", "string",
				"comment", "operator", "namespace", "type", "parameter",
			},
			TokenModifiers: []string{},
		},
		Full: protocol.SemanticDelta{
			Delta: &f,
		},
		Range: nil, // optional, you can support range later
	}

	return protocol.InitializeResult{
		Capabilities: capabilities,
		ServerInfo: &protocol.InitializeResultServerInfo{
			Name:    ls.name,
			Version: &ls.version,
		},
	}, nil
}

func (ls *LanguageServer) Shutdown(*glsp.Context) error {
	ls.log.Info("Stopping...")
	return nil
}

func (ls *LanguageServer) Initialized(_ *glsp.Context, _ *protocol.InitializedParams) error {
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

func (ls *LanguageServer) TextDocumentHover(_ *glsp.Context, params *protocol.HoverParams) (*protocol.Hover, error) {
	file := ls.project.File(params.TextDocument.URI)
	candidates := positionToCandidateNodes(params.Position, file)

	if len(candidates) > 0 {
		for _, candidate := range candidates {
			if s, ok := candidate.(ast.Symbol); ok {
				docstring := s.GetDocstring()
				targetLocation := s.GetNameLocation()
				kind := s.GetKind()
				if r, ok := s.(*spec.ReferenceNode); ok {
					if ns, ok := ls.project.Symbols()[r.Namespace]; ok {
						if n := ns.GetDecl(r.Name); n != nil {
							if n.GetDocstring() != "" {
								docstring = n.GetDocstring()
								targetLocation = r.GetLocation()
								kind = n.GetKind()
							}
						}
					}
				}
				if docstring != "" {
					if targetLocation != nil && targetLocation.ContainsPosition(params.Position) {
						out := &protocol.Hover{
							Contents: protocol.MarkupContent{
								Kind:  protocol.MarkupKindMarkdown,
								Value: fmt.Sprintf("# %s\n> %s\n\n%s", s.Ref(), kind, docstring),
							},
						}

						r := targetLocation.ToLspRange()
						out.Range = &r

						return out, nil
					}
				}
			}
			if helpful, ok := candidate.(ast.HelpfulNode); ok && helpful.GetHelp() != "" {
				out := &protocol.Hover{
					Contents: protocol.MarkupContent{
						Kind:  protocol.MarkupKindMarkdown,
						Value: helpful.GetHelp(),
					},
				}
				l := helpful.GetLocation()
				if l != nil {
					r := l.ToLspRange()
					out.Range = &r
				}

				return out, nil
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
					return res, nil
				}
			}
		}
	}
	return make([]protocol.CompletionItem, 0), nil
}

func (ls *LanguageServer) TextDocumentDefinition(_ *glsp.Context, params *protocol.DefinitionParams) (any, error) {
	// Location | []Location | []LocationLink | nil
	file := ls.project.File(params.TextDocument.URI)
	candidates := positionToCandidateNodes(params.Position, file)

	if ref, ok := candidates[0].(*spec.ReferenceNode); ok {
		if decl := ls.project.Symbols()[ref.Namespace].GetDecl(ref.Name); decl != nil {
			return protocol.Location{
				Range: decl.GetLocation().ToLspRange(),
				URI:   decl.GetLocation().Filename,
			}, nil
		}
	}
	return nil, nil

}
func findRefs(node ast.Node, ns string, n string) []protocol.Location {
	out := make([]protocol.Location, 0)
	if ref, ok := node.(*spec.ReferenceNode); ok {
		if ref.Namespace == ns && ref.Name == n {
			out = append(out, protocol.Location{
				Range: ref.GetLocation().ToLspRange(),
				URI:   ref.GetLocation().Filename,
			})
		}
	} else {
		for _, c := range node.Children() {
			out = append(out, findRefs(c, ns, n)...)
		}
	}
	return out
}
func (ls *LanguageServer) TextDocumentReferences(_ *glsp.Context, params *protocol.ReferenceParams) ([]protocol.Location, error) {
	file := ls.project.File(params.TextDocument.URI)
	candidates := positionToCandidateNodes(params.Position, file)

	out := make([]protocol.Location, 0)
	if ref, ok := candidates[0].(*spec.ReferenceNode); ok {
		for _, n := range file.AST() {
			out = append(out, findRefs(n, ref.Namespace, ref.Name)...)
		}
	}
	if s, ok := candidates[0].(ast.Symbol); ok {
		for _, n := range file.AST() {
			refParts := strings.Split(s.Ref(), ":")
			out = append(out, findRefs(n, refParts[0], refParts[1])...)
		}
	}

	return lo.Uniq(out), nil
}

//
//func (ls *LanguageServer) TextDocumentSemanticTokensFull(context *glsp.Context, params *protocol.SemanticTokensParams) (*protocol.SemanticTokens, error) {
//	out := &protocol.SemanticTokens{}
//	out.Data = make([]uint32, 0)
//
//	file := ls.project.File(params.TextDocument.URI)
//	if file == nil {
//		return out, nil
//	}
//
//	tokenStream := file.Tokens()
//	if tokenStream == nil || tokenStream.Size() == 0 {
//		return out, nil
//	}
//
//	// Mapping your token kinds to LSP semantic token types (just example mapping)
//	// You must adjust based on your token.Kind or token.Type or similar field.
//	var tokenTypeMapping = map[string]uint32{
//		"Keyword":      0,
//		"Variable":     1,
//		"Function":     2,
//		"Number":       3,
//		"String":       4,
//		"Comment":      5,
//		"Operator":     6,
//		"Namespace":    7,
//		"Type":         8,
//		"Parameter":    9,
//		"Int":          3,
//		"Float":        3,
//		"DocString":    5,
//		"BlockComment": 5,
//		"LineComment":  5,
//		// Add your token types accordingly
//	}
//
//	// You need to track previous token position for relative encoding
//	prevLine := 0
//	prevChar := 0
//
//	for i := 0; i < tokenStream.Size(); i++ {
//		token := tokenStream.Get(i)
//
//		line := int(token.GetLine()) - 1       // Convert 1-based to 0-based index
//		char := int(token.GetColumn()) - 1     // Convert 1-based to 0-based index
//		length := uint32(len(token.GetText())) // Length of token text
//
//		deltaLine := uint32(line - prevLine)
//		var deltaChar uint32
//		if deltaLine == 0 {
//			deltaChar = uint32(char - prevChar)
//		} else {
//			deltaChar = uint32(char)
//		}
//
//		// Determine token type - You should get the token kind or type properly
//		var tokenType uint32 = 0 // default unknown
//		// Example: if token.Kind is available:
//		// tokenType = tokenTypeMapping[token.Kind]
//
//		// If your token has a Kind or Type string, map it here:
//		kindStr := grammar.MinecraftMetascriptLexerLexerStaticData.SymbolicNames[token.GetTokenType()] // or token.Kind if it's string; adjust accordingly
//		if t, ok := tokenTypeMapping[kindStr]; ok {
//			tokenType = t
//		}
//
//		// No modifiers for now
//		tokenModifiers := uint32(0)
//
//		out.Data = append(out.Data, deltaLine, deltaChar, length, tokenType, tokenModifiers)
//
//		prevLine = line
//		prevChar = char
//	}
//
//	return out, nil
//}
