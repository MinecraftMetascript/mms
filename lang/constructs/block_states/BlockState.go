package block_states

import (
	"encoding/json"

	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
)

func init() {
	traversal.RegisterHelp[*grammar.BlockStateContext](
		func(state traversal.Construct, symbol traversal.Symbol, location traversal.TextLocation) *string {
			out := "Defines a Minecraft [block state](https://minecraft.wiki/w/Block_states).<br/>Currently *only has support for block name*.<br/>Example: `Block(minecraft:stone)`"
			return &out
		},
	)

	traversal.Register(
		func(state *grammar.BlockStateContext, currentNamespace string, scope *traversal.Scope) traversal.Construct {
			nameCtx := state.ResourceReference()
			if nameCtx == nil {
				return nil
			}
			name := traversal.ConstructRegistry.Construct(nameCtx, currentNamespace, scope).(*traversal.Reference)
			return &BlockState{
				Name: *name,
			}
		},
	)
}

func BlockStateRef(r traversal.Reference) *BlockState {
	return &BlockState{Name: r}
}

type BlockState struct {
	traversal.BaseConstruct
	Name traversal.Reference
}

func (bs BlockState) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(struct {
		Name traversal.Reference `json:"Name"`
	}{
		Name: bs.Name,
	}, "", "  ")
}
