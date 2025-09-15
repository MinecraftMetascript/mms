package noise

import (
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type NoiseBlock struct {
	location     traversal.TextLocation
	Declarations []traversal.Symbol
}

func (n NoiseBlock) GetCompletions(cursorPosition protocol.Position) []protocol.CompletionItem {

	out := []protocol.CompletionItem{}

	// TODO: Make sure we aren't conflicting with any other declarations?
	out = append(out, protocol.CompletionItem{
		Label:            "New Noise",
		Kind:             lib.Ptr(protocol.CompletionItemKindSnippet),
		Detail:           lib.Ptr("Declare a new Noise type"),
		InsertText:       lib.Ptr("$0 = Noise($1).Amplitudes($2)"),
		InsertTextFormat: lib.Ptr(protocol.InsertTextFormatSnippet),
	})

	return out
}

func (n NoiseBlock) GetLocation() traversal.TextLocation {
	return n.location
}

func init() {
	traversal.RegisterNodeFactory(NoiseBlockFactory{}, false)
}

type NoiseBlockFactory struct{}

func (n NoiseBlockFactory) CreateDeclaration(ctx traversal.DeclarableContext, namespace string, scope *traversal.Scope) (traversal.Symbol, bool) {
	return nil, false
}

func (n NoiseBlockFactory) Create(ctx *grammar.NoiseBlockContext, namespace string, scope *traversal.Scope) *NoiseBlock {
	out := &NoiseBlock{
		location:     traversal.RuleLocation(ctx, scope.CurrentFile),
		Declarations: make([]traversal.Symbol, 0),
	}
	for _, fn := range ctx.AllNoiseDeclaration() {
		res, ok := traversal.DeclareNode(fn, namespace, scope)
		if ok {
			out.Declarations = append(out.Declarations, res)
		}
	}
	return out
}

func (n NoiseBlockFactory) GetHelp(node *NoiseBlock, symbol traversal.Symbol, location traversal.TextLocation) *traversal.Help {
	return nil
}

func (n NoiseBlockFactory) Export(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	return nil
}
