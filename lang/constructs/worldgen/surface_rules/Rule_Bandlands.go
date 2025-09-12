package surface_rules

import (
	"encoding/json"

	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
)

func init() {
	traversal.RegisterHelp[*grammar.SurfaceRule_BandlandsContext](
		func(construct traversal.Construct, symbol traversal.Symbol, location traversal.TextLocation) *string {
			out := "Used in badlands to place terracotta."
			return &out
		},
	)
	traversal.Register(
		func(_ *grammar.SurfaceRule_BandlandsContext, _ string, _ *traversal.Scope) traversal.Construct {
			return &Bandlands{}
		},
	)
}

type Bandlands struct {
	traversal.BaseConstruct
}

func (b Bandlands) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(struct {
		Type SurfaceRuleKind `json:"type"`
	}{
		Type: SurfaceRule_Bandlands,
	}, "", "  ")
}
