package primitives

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
)

func init() {
	traversal.RegisterHelp[*grammar.VerticalAnchorContext](
		func(construct traversal.Construct, symbol traversal.Symbol, location traversal.TextLocation) *string {
			out := "Defines a vertical anchor, which can be absolute, above the bottom, or below the top.<br/>Absolute: `100`<br/>Above bottom: `~-100`<br/>Below top: `~100`"
			return &out
		})
	traversal.Register(
		func(anchor *grammar.VerticalAnchorContext, _ string, _ *traversal.Scope) traversal.Construct {
			if anchor.Identifier() != nil {
				// Construct a reference
			} else {
				t := anchor.GetText()
				isRelative := strings.Contains(t, "~")
				if i := anchor.Int(); i != nil {
					var value int
					if v, err := strconv.Atoi(i.GetText()); err == nil {
						value = v
					}
					var anchorType VerticalAnchorType
					if isRelative {
						if value < 0 {
							anchorType = VerticalAnchorType_BelowTop
						} else {
							anchorType = VerticalAnchorType_AboveBottom
						}
					} else {
						anchorType = VerticalAnchorType_Absolute
					}

					return &VerticalAnchor{
						Type:  anchorType,
						Value: value,
					}
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
