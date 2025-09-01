package primitives

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"

	"github.com/antlr4-go/antlr/v4"
)

func init() {
	traversal.ConstructRegistry.Register(
		reflect.TypeFor[*grammar.VerticalAnchorDefinitionContext](),
		func(ctx antlr.ParserRuleContext, _ string, _ *traversal.Scope) traversal.Construct {
			anchor := ctx.(*grammar.VerticalAnchorDefinitionContext)

			if rule := anchor.VerticalAnchor_Absolute(); rule != nil {
				value, err := strconv.Atoi(rule.Int().GetText())

				if err != nil {
					return nil
				}

				return &VerticalAnchor{
					Type:  VerticalAnchorType_Absolute,
					Value: value,
				}
			}
			if rule := anchor.VerticalAnchor_AboveBottom(); rule != nil {
				value, err := strconv.Atoi(rule.Int().GetText())
				fmt.Println("Found above bottom", value)
				if err != nil {
					return nil
				}
				return &VerticalAnchor{
					Type:  VerticalAnchorType_AboveBottom,
					Value: value,
				}
			}
			if rule := anchor.VerticalAnchor_BelowTop(); rule != nil {
				value, err := strconv.Atoi(rule.Int().GetText())
				fmt.Println("Found below top", value)
				if err != nil {
					return nil
				}
				return &VerticalAnchor{
					Type:  VerticalAnchorType_BelowTop,
					Value: value,
				}
			}
			return nil
		},
	)
}

type VerticalAnchorType string

const (
	VerticalAnchorType_Absolute    VerticalAnchorType = "absolute"
	VerticalAnchorType_AboveBottom VerticalAnchorType = "above_bottom"
	VerticalAnchorType_BelowTop    VerticalAnchorType = "below_top"
)

type VerticalAnchor struct {
	traversal.BaseConstruct
	Type  VerticalAnchorType
	Value int
}

func (v VerticalAnchor) ExportSymbol(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	return nil
}

func (v VerticalAnchor) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("{\"%s\": %d}", v.Type, v.Value)), nil
}
