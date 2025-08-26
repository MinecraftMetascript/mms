package surface_rules

import (
	"encoding/json"
	"fmt"

	"github.com/minecraftmetascript/mms/lang/grammars"
)

func NewSequenceRule(
	ctx *grammars.SurfaceRule_SequenceContext,
	factory RuleFactory,
) (*SequenceRule, []error) {
	rules := make([]SurfaceRule, 0)
	errors := make([]error, 0)
	for _, childRuleCtx := range ctx.AllSurfaceRule() {
		childRule, err := factory.ConstructSurfaceRule(childRuleCtx.(*grammars.SurfaceRuleContext))
		if err != nil {
			errors = append(errors, err...)
			continue
		}
		rules = append(rules, childRule)
	}

	errOut := errors
	if len(errors) == 0 {
		errOut = nil
	}
	return &SequenceRule{
		Rules: rules,
	}, errOut

}

type SequenceRule struct {
	BaseRule
	Rules []SurfaceRule
}

func (r SequenceRule) Type() SurfaceRuleKind {
	return SequenceRuleKind
}

func (r SequenceRule) String() string {
	return fmt.Sprintf("surfaceRule(%s)", r.Type())
}

func (r SequenceRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type     SurfaceRuleKind `json:"type"`
		Sequence []SurfaceRule   `json:"sequence"`
	}{
		Type:     r.Type(),
		Sequence: r.Rules,
	})
}
