package surface_rules

import (
	"encoding/json"
	"mms2/lang/grammar"
	"mms2/lang/traversal"
	"mms2/lib"
	"reflect"

	"github.com/antlr4-go/antlr/v4"
)

func init() {
	traversal.ConstructRegistry.Register(
		reflect.TypeFor[grammar.SurfaceRule_SequenceContext](),
		func(ctx_ antlr.ParserRuleContext, ns string, scope *traversal.Scope) traversal.Construct {
			sequence := ctx_.(*grammar.SurfaceRule_SequenceContext)

			children := make([]traversal.Construct, 0)
			for _, childCtx := range sequence.AllSurfaceRule() {
				childRule := traversal.ConstructRegistry.Construct(childCtx.GetChild(0).(antlr.ParserRuleContext), ns, scope)
				children = append(children, childRule)
			}

			return &Sequence{
				Rules: children,
			}
		},
	)
}

type Sequence struct {
	traversal.Construct
	Rules []traversal.Construct
}

func (s Sequence) ExportSymbol(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	return exportSurfaceRule(symbol, rootDir, s)
}

func (r Sequence) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type     SurfaceRuleKind       `json:"type"`
		Sequence []traversal.Construct `json:"sequence"`
	}{
		Type:     SurfaceRule_Sequence,
		Sequence: r.Rules,
	})
}
