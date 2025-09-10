package surface_rules

import (
	"encoding/json"
	"reflect"

	"github.com/minecraftmetascript/mms/lang/constructs/block_states"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"

	"github.com/antlr4-go/antlr/v4"
)

func init() {
	traversal.ConstructRegistry.Register(reflect.TypeFor[grammar.SurfaceRule_BlockContext](),
		func(baseCtx antlr.ParserRuleContext, _ string, scope *traversal.Scope) traversal.Construct {
			ctx, ok := baseCtx.(*grammar.SurfaceRule_BlockContext)
			if !ok {
				return nil
			}

			out := &Block{}
			if ref := ctx.ResourceReference(); ref != nil {
				cons := traversal.ConstructRegistry.Construct(
					ref, "minecraft", scope,
				)
				if r, ok := cons.(*traversal.Reference); ok && r != nil {

					out.BlockName = *block_states.BlockStateRef(*r)
					return out
				}
			}

			return nil
		},
	)
}

type Block struct {
	traversal.BaseConstruct

	BlockName block_states.BlockState
}

func (b Block) ExportSymbol(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	// TODO: Should we just accept at face value that symbol is a symbol for this value?
	data, err := json.Marshal(b)
	if err != nil {
		return err
	}
	rootDir.
		MkDir(symbol.GetReference().GetNamespace(), nil).
		MkDir("_debug", nil).
		MkDir("surface_rule", nil).
		MkFile(symbol.GetReference().GetName()+".json", string(data), nil)

	return nil

}

func (b Block) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(struct {
		Type  SurfaceRuleKind         `json:"type"`
		State block_states.BlockState `json:"result_state"`
	}{
		Type:  SurfaceRule_Block,
		State: b.BlockName,
	}, "", "  ")
}
