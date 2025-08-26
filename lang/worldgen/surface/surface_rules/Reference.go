package surface_rules

import (
	"encoding/json"
	"fmt"

	"github.com/minecraftmetascript/mms/lang/grammars"
	"github.com/minecraftmetascript/mms/lib"
)

type ReferenceRule struct {
	BaseRule
	lib.Reference
}

func NewReferenceRule(ctx *grammars.SurfaceRuleReferenceContext) (*ReferenceRule, error) {
	return &ReferenceRule{
		Reference: lib.Reference{
			Name: ctx.GetText(),
		},
	}, nil
}

func (r ReferenceRule) Type() SurfaceRuleKind {
	return ReferenceRuleKind
}

func (r ReferenceRule) String() string {
	return fmt.Sprintf("surfaceRule(%s)", r.Type())
}

func (r ReferenceRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type SurfaceRuleKind `json:"type"`
		Name string          `json:"name"`
	}{
		Type: r.Type(),
		Name: r.Name,
	})
}
