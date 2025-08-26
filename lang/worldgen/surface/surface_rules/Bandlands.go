package surface_rules

import (
	"encoding/json"
	"fmt"

	"github.com/minecraftmetascript/mms/lang/grammars"
)

func NewBandlandsRule(ctx *grammars.SurfaceRule_BandlandsContext) (*BandlandsRule, error) {
	return &BandlandsRule{}, nil
}

type BandlandsRule struct {
	BaseRule
}

func (r BandlandsRule) Type() SurfaceRuleKind {
	return BandlandsRuleKind
}

func (r BandlandsRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type SurfaceRuleKind `json:"type"`
	}{
		Type: r.Type(),
	})
}

func (r BandlandsRule) String() string {
	return fmt.Sprintf("surfaceRule(%s)", r.Type())
}
