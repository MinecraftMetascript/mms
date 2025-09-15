package primitives

import (
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
)

type DeclarationFactory struct{}

func (d DeclarationFactory) Create(ctx *grammar.DeclareContext, namespace string, scope *traversal.Scope) traversal.Symbol {
	idCtx := ctx.Identifier()
	if idCtx == nil {
		scope.DiagnoseSemanticError("Missing identifier", ctx)
		return nil
	}
	id := idCtx.GetText()

	return traversal.
		NewEmptySymbol(
			traversal.TerminalNodeLocation(idCtx, scope.CurrentFile),
			traversal.NewReference(id, namespace),
		)
}

func (d DeclarationFactory) GetHelp(node traversal.Symbol, symbol traversal.Symbol, location traversal.TextLocation) *traversal.Help {
	return nil
}

func (d DeclarationFactory) Export(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	return nil
}
