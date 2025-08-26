package surface_conditions

import (
	"errors"

	"github.com/minecraftmetascript/mms/lang/grammars"
	"github.com/minecraftmetascript/mms/lib"
)

func NewSurfaceConditionReference(ctx *grammars.SurfaceConditionReferenceContext, defaultNamespace string) (*ConditionReference, error) {
	return &ConditionReference{
		Reference: lib.ParseReferential(defaultNamespace, ctx),
	}, nil
}

type ConditionReference struct {
	lib.Reference
	BaseCondition
}

func (c ConditionReference) Type() ConditionType {
	return ReferenceConditionType
}

func (c ConditionReference) MarshalJSON() ([]byte, error) {
	return nil, errors.New("ConditionReference cannot be marshaled")
}
