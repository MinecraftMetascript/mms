package density_functions

import (
	"encoding/json"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lang/builder_chain"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
	"github.com/minecraftmetascript/mms/lsp/completions"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type NoiseFnFactory struct{}

func (n NoiseFnFactory) Create(ctx *grammar.DensityFn_NoiseContext, namespace string, scope *traversal.Scope) *NoiseDensityFn {
	noiseFnBuilder := builder_chain.NewBuilderChain(
		builder_chain.Build(
			func(ctx *grammar.Builder_XZScaleContext, target *NoiseDensityFn, scope *traversal.Scope, namespace string) {
				builder_chain.Builder_GetFloat(ctx, func(v float64) { target.XzScale = v }, scope, "XzScale")
			},
		), builder_chain.Build(
			func(ctx *grammar.Builder_YScaleContext, target *NoiseDensityFn, scope *traversal.Scope, namespace string) {
				builder_chain.Builder_GetFloat(ctx, func(v float64) { target.YScale = v }, scope, "YScale")
			},
		),
	)

	out := &NoiseDensityFn{
		XzScale:  1,
		YScale:   1,
		location: traversal.RuleLocation(ctx, scope.CurrentFile),
		ctx:      ctx,
		scope:    scope,
	}

	if noiseInlineCtx := ctx.DensityFn_InlineNoise(); noiseInlineCtx != nil {
		noiseRef := traversal.ConstructRegistry.Construct(noiseInlineCtx, namespace, scope)
		if noiseRef == nil {
			// TODO: Diagnose
		} else if ref, ok := noiseRef.(*traversal.Reference); ok {
			out.Noise = *ref
		}
	}

	if refCtx := ctx.ResourceReference(); refCtx != nil {
		cons := traversal.ConstructRegistry.Construct(refCtx, namespace, scope)
		if r, ok := cons.(*traversal.Reference); ok {
			out.Noise = *r
		}
	}

	for _, r := range ctx.AllDensityFn_NoiseBuilder() {
		child := r.GetChild(0)
		if child == nil {
			continue
		}
		builder_chain.Invoke(noiseFnBuilder, child.(antlr.ParserRuleContext), out, scope, namespace)
	}

	return out

}

func (n NoiseFnFactory) GetHelp(node *NoiseDensityFn, symbol traversal.Symbol, location traversal.TextLocation) *traversal.Help {
	return &traversal.Help{
		Content:  "Uses a noise function to determine density.<br/>Uses .XzScale( densityFn ) and .YScale( densityFn ) to scale the noise function.",
		Position: node.GetLocation(),
	}
}

func init() {
	traversal.RegisterNodeFactory(NoiseFnFactory{})
}

type NoiseDensityFn struct {
	Noise    traversal.Reference
	XzScale  float64
	YScale   float64
	location traversal.TextLocation
	ctx      *grammar.DensityFn_NoiseContext
	scope    *traversal.Scope
}

func (c NoiseDensityFn) GetLocation() traversal.TextLocation {
	return c.location
}

func (c NoiseDensityFn) GetCompletions(cursorPosition protocol.Position) []protocol.CompletionItem {
	items := []protocol.CompletionItem{}

	// Provide reference completions only when cursor is between the parentheses of Noise(...)
	if c.ctx != nil && c.scope != nil {
		cursorPos := completions.CursorWithin(c.ctx, cursorPosition, "(", ")")
		allPossible := completions.ReferenceCompletions(c.scope, traversal.Noise, cursorPosition)
		switch cursorPos {
		case completions.CursorWithinEmpty:
			items = append(items, allPossible...)
		case completions.CursorAtEnd:
			ref := c.ctx.ResourceReference()
			filteredCompletions := completions.FilterByReference(c.scope, traversal.Noise, cursorPosition, allPossible, ref)
			items = append(items, filteredCompletions...)
		case completions.CursorNotWithin:
			// Cursor is not within the parens of the root, so we want to provide builder functions

			// We do want to check for children though, as we don't want to suggest builders if the user is currently trying to edit one
			if cursorPos := completions.CursorWithinRecursive(c.ctx, cursorPosition, "(", ")"); cursorPos != completions.CursorNotWithin {
				break
			}
			if len(lib.GetAntlrChildren[*grammar.Builder_XZScaleContext](c.ctx)) == 0 {
				items = append(items,
					completions.BuilderFnCompletion(
						"XZScale",
						"XZScale(${1})",
						c.location.Stop.ToLspPosition(),
					),
				)
			}
			if len(lib.GetAntlrChildren[*grammar.Builder_YScaleContext](c.ctx)) == 0 {
				items = append(items,
					completions.BuilderFnCompletion(
						"YScale",
						"YScale(${1})",
						c.location.Stop.ToLspPosition(),
					),
				)
			}
		}
	}

	return items
}

func (c NoiseDensityFn) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(struct {
		Type    string  `json:"type"`
		Noise   string  `json:"noise"`
		XzScale float64 `json:"xz_scale"`
		YScale  float64 `json:"y_scale"`
	}{
		Type:    "minecraft:noise",
		Noise:   c.Noise.String(),
		XzScale: c.XzScale,
		YScale:  c.YScale,
	}, "", "  ")
}
