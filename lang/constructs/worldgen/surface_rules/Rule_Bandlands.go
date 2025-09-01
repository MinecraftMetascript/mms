package surface_rules

import (
	"encoding/json"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"reflect"

	"github.com/antlr4-go/antlr/v4"
)

func init() {
	traversal.ConstructRegistry.Register(
		reflect.TypeFor[grammar.SurfaceRule_BandlandsContext](),
		func(_ antlr.ParserRuleContext, _ string, _ *traversal.Scope) traversal.Construct {
			return &Bandlands{}
		},
	)
}

type Bandlands struct {
	traversal.BaseConstruct
}

func (b Bandlands) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type SurfaceRuleKind `json:"type"`
	}{
		Type: SurfaceRule_Bandlands,
	})
}
