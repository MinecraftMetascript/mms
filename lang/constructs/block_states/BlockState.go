package block_states

import (
	"encoding/json"

	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
)

func init() {
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
