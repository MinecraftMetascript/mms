package surface_rules

import (
	"encoding/json"
	"fmt"

	"github.com/minecraftmetascript/mms/lang/grammars"
	"github.com/minecraftmetascript/mms/lib"
)

type SurfaceRuleKind string

const (
	BandlandsRuleKind   SurfaceRuleKind = "minecraft:bandlands"
	BlockRuleKind       SurfaceRuleKind = "minecraft:block"
	SequenceRuleKind    SurfaceRuleKind = "minecraft:sequence"
	ConditionalRuleKind SurfaceRuleKind = "minecraft:condition"
	ReferenceRuleKind   SurfaceRuleKind = "__mms:reference"
)

type SurfaceRule interface {
	json.Marshaler
	fmt.Stringer
	lib.LocatedText

	Type() SurfaceRuleKind
}

type BaseRule struct {
	lib.LocatedText
	location lib.TextLocation
}

func (b *BaseRule) GetLocation() lib.TextLocation {
	return b.location
}
func (b *BaseRule) SetLocation(location lib.TextLocation) {
	b.location = location
}

type RuleFactory interface {
	ConstructSurfaceRule(*grammars.SurfaceRuleContext) (SurfaceRule, []error)
}
