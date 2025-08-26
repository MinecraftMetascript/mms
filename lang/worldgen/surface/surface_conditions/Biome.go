package surface_conditions

import (
	"encoding/json"
	"fmt"

	"github.com/minecraftmetascript/mms/lang/grammars"
	"github.com/minecraftmetascript/mms/lang/primitives"
)

func NewBiomeCondition(ctx *grammars.SurfaceCondition_BiomeContext) (*BiomeCondition, error) {
	refs := ctx.AllResourceReference()
	biomes := make([]string, len(refs))
	for i, ref := range refs {
		ref, err := primitives.FromResourceReference(ref, "minecraft")
		if err != nil {
			return nil, err
		}
		biomes[i] = ref.String()
	}

	return &BiomeCondition{
		biomes: biomes,
	}, nil
}

type BiomeCondition struct {
	BaseCondition
	biomes []string
}

func (b BiomeCondition) Type() ConditionType { return BiomeConditionType }

func (b BiomeCondition) String() string {
	return fmt.Sprintf("surfaceCondition(%s)", b.Type())
}

func (b BiomeCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type    ConditionType `json:"type"`
		Biomes  []string      `json:"biomes"`
		Comment *string       `json:"__comment,omitempty"`
	}{
		Type:   b.Type(),
		Biomes: b.biomes,
	})
}
