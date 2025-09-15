package block_states

import (
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
)

type BlockStateFactory struct{}

func (b BlockStateFactory) CreateDeclaration(ctx traversal.DeclarableContext, namespace string, scope *traversal.Scope) (traversal.Symbol, bool) {
	return nil, false
}

func (b BlockStateFactory) Create(ctx *grammar.BlockStateContext, namespace string, scope *traversal.Scope) *BlockState {
	nameCtx := ctx.ResourceReference()
	if nameCtx == nil {
		return nil
	}
	name := traversal.ConstructRegistry.Construct(nameCtx, namespace, scope).(*traversal.Reference)
	return &BlockState{
		Name:     *name,
		location: traversal.RuleLocation(ctx, scope.CurrentFile),
	}
}

func (b BlockStateFactory) GetHelp(node *BlockState, symbol traversal.Symbol, location traversal.TextLocation) *traversal.Help {
	return &traversal.Help{
		Content:  "Defines a Block State",
		Position: node.GetLocation(),
	}
}

func (b BlockStateFactory) Export(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	if rootDir == nil {
		return nil
	}
	rootDir.
		MkDir("_debug", nil).
		MkDir("block_states", nil).
		MkFile(symbol.GetReference().GetName()+".json", "", nil)
	return nil
}

func init() {
	traversal.RegisterNodeFactory(BlockStateFactory{}, false)
}

func BlockStateRef(r traversal.Reference) *BlockState {
	return &BlockState{Name: r}
}

type BlockState struct {
	Name     traversal.Reference `json:"Name"`
	location traversal.TextLocation
}

func (bs BlockState) GetLocation() traversal.TextLocation {
	return bs.location
}
