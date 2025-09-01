package surface_rules

import (
	"encoding/json"
	"mms2/lang/constructs/block_states"
	"mms2/lang/grammar"
	"mms2/lang/traversal"
	"mms2/lib"
	"reflect"

	"github.com/antlr4-go/antlr/v4"
)

func init() {
	traversal.ConstructRegistry.Register(reflect.TypeFor[grammar.SurfaceRule_BlockContext](),
		func(baseCtx antlr.ParserRuleContext, _ string, scope *traversal.Scope) traversal.Construct {
			ctx, ok := baseCtx.(*grammar.SurfaceRule_BlockContext)
			if !ok {
				return nil
			}

			var blockName string
			if ref := ctx.ResourceReference(); ref != nil {
				blockName = traversal.ConstructRegistry.Construct(
					ref, "minecraft", scope,
				).(*traversal.Reference).String()
			}

			return &Block{
				BlockName: blockName,
			}
		},
	)
}

type Block struct {
	traversal.BaseConstruct

	BlockName string
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
		State: block_states.BlockState{Name: b.BlockName},
	}, "", "  ")
}
