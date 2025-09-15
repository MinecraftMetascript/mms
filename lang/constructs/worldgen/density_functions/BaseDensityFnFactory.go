package density_functions

import (
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
)

type BaseDensityFnFactory struct{}

func (b BaseDensityFnFactory) CreateDeclaration(ctx traversal.DeclarableContext, namespace string, scope *traversal.Scope) (traversal.Symbol, bool) {
	return nil, false
}

func (b BaseDensityFnFactory) Export(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	return exportDensityFunction(symbol, rootDir, symbol.GetValue())
}
