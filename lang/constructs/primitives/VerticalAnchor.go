package primitives

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
)

type VerticalAnchorFactory struct{}

func (v VerticalAnchorFactory) CreateDeclaration(ctx *grammar.VerticalAnchorDeclarationContext, namespace string, scope *traversal.Scope) (traversal.Symbol, bool) {
	if ctx.Declare() != nil {
		declaration := DeclarationFactory{}.Create(ctx.Declare().(*grammar.DeclareContext), namespace, scope)
		if declaration == nil {
			scope.DiagnoseSemanticError("Missing identifier", ctx.Declare())
		}
		anchor := v.Create(ctx.VerticalAnchor().(*grammar.VerticalAnchorContext), namespace, scope)
		if anchor == nil {
			scope.DiagnoseSemanticError("Missing or invalid vertical anchor", ctx.VerticalAnchor())
		}
		return traversal.NewSymbol(
			declaration.GetNameLocation(),
			traversal.RuleLocation(ctx.VerticalAnchor(), scope.CurrentFile),
			anchor,
			declaration.GetReference(),
			traversal.VerticalAnchor,
		), true

	} else {
		return nil, false
	}
}

func (v VerticalAnchorFactory) Create(ctx *grammar.VerticalAnchorContext, namespace string, scope *traversal.Scope) *VerticalAnchor {
	t := ctx.GetText()
	isRelative := strings.Contains(t, "~")
	if i := ctx.Int(); i != nil {
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
			Type:     anchorType,
			Value:    value,
			location: traversal.RuleLocation(ctx, scope.CurrentFile),
		}
	}
	return nil
}

func (v VerticalAnchorFactory) GetHelp(node *VerticalAnchor, symbol traversal.Symbol, location traversal.TextLocation) *traversal.Help {
	return &traversal.Help{
		Content:  "Defines a vertical anchor, which can be absolute, above the bottom, or below the top.<br/>Absolute: `100`<br/>Above bottom: `~-100`<br/>Below top: `~100`",
		Position: node.GetLocation(),
	}
}

func (v VerticalAnchorFactory) Export(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	if rootDir == nil {
		return nil
	}
	rootDir.
		MkDir("_debug", nil).
		MkDir("vertical_anchor", nil).
		MkFile(symbol.GetReference().GetName()+".json", "", nil)
	return nil
}

func init() {
	traversal.RegisterNodeFactory(VerticalAnchorFactory{}, true)
}

type VerticalAnchorType string

const (
	VerticalAnchorType_Absolute    VerticalAnchorType = "absolute"
	VerticalAnchorType_AboveBottom VerticalAnchorType = "above_bottom"
	VerticalAnchorType_BelowTop    VerticalAnchorType = "below_top"
)

type VerticalAnchor struct {
	location traversal.TextLocation
	Type     VerticalAnchorType
	Value    int
}

func (v VerticalAnchor) GetLocation() traversal.TextLocation {
	return v.location
}

func (v VerticalAnchor) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("{\"%s\": %d}", v.Type, v.Value)), nil
}
