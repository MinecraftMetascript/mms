package surface_rules

import (
	"encoding/json"

	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"

	"github.com/antlr4-go/antlr/v4"
)

func init() {
	traversal.Register(
		func(sequence *grammar.SurfaceRule_SequenceContext, ns string, scope *traversal.Scope) traversal.Construct {
			children := make([]traversal.Construct, 0)
			for _, childCtx := range sequence.AllSurfaceRule() {
				if childCtx.GetChildCount() > 0 {
					if prc, ok := childCtx.GetChild(0).(antlr.ParserRuleContext); ok {
						childRule := traversal.ConstructRegistry.Construct(prc, ns, scope)
						children = append(children, childRule)
					}
				}
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
	return json.MarshalIndent(struct {
		Type     SurfaceRuleKind       `json:"type"`
		Sequence []traversal.Construct `json:"sequence"`
	}{
		Type:     SurfaceRule_Sequence,
		Sequence: r.Rules,
	}, "", "  ")
}
