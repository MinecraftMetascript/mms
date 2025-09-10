package block_states

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
)

func init() {
	traversal.ConstructRegistry.Register(
		reflect.TypeFor[*grammar.BlockStateContext](),
		func(ctx antlr.ParserRuleContext, currentNamespace string, scope *traversal.Scope) traversal.Construct {
			fmt.Println("eeee")
			state := ctx.(*grammar.BlockStateContext)
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
