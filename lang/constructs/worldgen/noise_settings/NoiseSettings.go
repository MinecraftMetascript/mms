package noise_settings

import (
	"encoding/json"
	"reflect"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lang/builder_chain"
	"github.com/minecraftmetascript/mms/lang/constructs/block_states"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
)

func getBuilder() *builder_chain.BuilderChain[NoiseSettings] {
	return builder_chain.NewBuilderChain[NoiseSettings](
		builder_chain.Build(
			func(ctx *grammar.Builder_SeaLevelContext, target *NoiseSettings, scope *traversal.Scope, namespace string) {
				builder_chain.Builder_GetInt(ctx, func(v int) { target.SeaLevel = v }, scope, "SeaLevel")
			},
		),
		builder_chain.Build(
			func(ctx *grammar.Builder_DisableCreaturesContext, target *NoiseSettings, scope *traversal.Scope, namespace string) {
				target.DisableCreatures = true
			},
		),
		builder_chain.Build(
			func(ctx *grammar.Builder_DisableVeinsContext, target *NoiseSettings, scope *traversal.Scope, namespace string) {
				target.DisableVeins = true
			},
		),
		builder_chain.Build(
			func(ctx *grammar.Builder_DisableAquifersContext, target *NoiseSettings, scope *traversal.Scope, namespace string) {
				target.DisableAquifers = true
			},
		),
		builder_chain.Build(
			func(ctx *grammar.Builder_LegacyRandomSourceContext, target *NoiseSettings, scope *traversal.Scope, namespace string) {
				target.LegacyRandomSource = true
			},
		),
		builder_chain.Build(
			func(ctx *grammar.Builder_DefaultBlockContext, target *NoiseSettings, scope *traversal.Scope, namespace string) {

				state := traversal.ConstructRegistry.Construct(ctx.BlockState(), namespace, scope)
				if state == nil {
					state = traversal.ConstructRegistry.Construct(ctx.ResourceReference(), namespace, scope)
					if state == nil {
						return
					}
					target.DefaultBlock = block_states.BlockStateRef(*state.(*traversal.Reference))
					return
				}
				target.DefaultBlock = state
			},
		),
		builder_chain.Build(
			func(ctx *grammar.Builder_DefaultFluidContext, target *NoiseSettings, scope *traversal.Scope, namespace string) {
				state := traversal.ConstructRegistry.Construct(ctx.BlockState(), namespace, scope)
				if state == nil {
					return
				}
				target.DefaultFluid = state
			},
		),
		builder_chain.Build(
			func(ctx *grammar.Builder_MinYContext, target *NoiseSettings, scope *traversal.Scope, namespace string) {
				builder_chain.Builder_GetInt(ctx, func(v int) { target.MinY = v }, scope, "MinY")
			},
		),
		builder_chain.Build(
			func(ctx *grammar.Builder_HeightContext, target *NoiseSettings, scope *traversal.Scope, namespace string) {
				builder_chain.Builder_GetInt(ctx, func(v int) { target.Height = v }, scope, "MinY")
			},
		),
		builder_chain.Build(
			func(ctx *grammar.Builder_NoiseRouterContext, target *NoiseSettings, scope *traversal.Scope, namespace string) {
				if sym, _ := traversal.ExtractInlineConstruct(ctx.NoiseRouter(), namespace, scope, "NoiseRouter"); sym != nil {
					target.NoiseRouter = sym.GetValue()
				}
			},
		),
		builder_chain.Build(
			func(ctx *grammar.Builder_SurfaceRuleContext, target *NoiseSettings, scope *traversal.Scope, namespace string) {
				if sym, _ :=
					traversal.ExtractInlineConstruct(ctx.SurfaceRule(), namespace, scope, "SurfaceRule"); sym != nil {
					target.NoiseRouter = sym.GetValue()
				}
			},
		),
		builder_chain.Build(
			func(ctx *grammar.Builder_NoiseSizeContext, target *NoiseSettings, scope *traversal.Scope, namespace string) {
				horizontalCtx := ctx.Int(0)
				if horizontalCtx == nil {
					// todo: diagnose
					return
				}
				horizontalVal, err := strconv.Atoi(horizontalCtx.GetText())
				if err != nil {
					//todo: diagnose
				} else {
					target.NoiseSizeX = horizontalVal
				}
				verticalCtx := ctx.Int(1)
				if verticalCtx == nil {
					// todo: diagnose
					return
				}
				verticalVal, err := strconv.Atoi(verticalCtx.GetText())
				if err != nil {
					//todo: diagnose
				} else {
					target.NoiseSizeY = verticalVal
				}

			},
		),
	)
}

func init() {
	traversal.Register(
		func(block *grammar.NoiseSettingsBlockContext, currentNamespace string, scope *traversal.Scope) traversal.Construct {
			for _, child := range block.GetChildren() {
				if rule, ok := child.(antlr.ParserRuleContext); ok {
					traversal.ConstructRegistry.Construct(rule, currentNamespace, scope)
				}
			}
			return nil
		},
	)
	traversal.Register(
		func(declaration *grammar.NoiseSettingsDeclarationContext, currentNamespace string, scope *traversal.Scope) traversal.Construct {
			if v := traversal.ProcessDeclaration(declaration.Declare(), declaration.NoiseSettings(), scope, currentNamespace, "NoiseSettings"); v != nil {
				return v.GetValue()
			}
			return nil
		},
	)
	traversal.Register(
		func(settings *grammar.NoiseSettingsContext, currentNamespace string, scope *traversal.Scope) traversal.Construct {
			chain := getBuilder()

			out := &NoiseSettings{
				SeaLevel:           63,
				DisableAquifers:    false,
				DisableCreatures:   false,
				DisableVeins:       false,
				LegacyRandomSource: false,
				SpawnTarget:        make([]json.Marshaler, 0),
				NoiseSizeX:         1,
				NoiseSizeY:         1,
				// Overworld Defaults
				MinY:   -64,
				Height: 384,
			}
			for _, builder := range settings.AllNoiseSettings_Builder() {
				builder_chain.Invoke(chain, builder.GetChild(0).(antlr.ParserRuleContext), out, scope, currentNamespace)
			}

			builder_chain.Require(
				chain,
				settings,
				scope,
				reflect.TypeFor[*grammar.Builder_DefaultBlockContext](),
				".DefaultBlock",
			)
			builder_chain.Require(
				chain,
				settings,
				scope,
				reflect.TypeFor[*grammar.Builder_DefaultFluidContext](),
				".DefaultFluid",
			)
			builder_chain.Require(
				chain,
				settings,
				scope,
				reflect.TypeFor[*grammar.Builder_NoiseRouterContext](),
				".NoiseRouter",
			)
			builder_chain.Require(
				chain,
				settings,
				scope,
				reflect.TypeFor[*grammar.Builder_SurfaceRuleContext](),
				".SurfaceRule",
			)
			return out
		},
	)
}

type NoiseSettings struct {
	SeaLevel           int
	DisableCreatures   bool
	DisableVeins       bool
	DisableAquifers    bool
	LegacyRandomSource bool
	DefaultBlock       traversal.Construct
	DefaultFluid       traversal.Construct
	SpawnTarget        []json.Marshaler

	MinY       int
	Height     int
	NoiseSizeX int
	NoiseSizeY int

	NoiseRouter traversal.Construct
	SurfaceRule traversal.Construct
}

func (n *NoiseSettings) MarshalJSON() ([]byte, error) {
	noise := struct {
		MinY       int `json:"min_y"`
		Height     int `json:"height"`
		NoiseSizeX int `json:"size_horizontal"`
		NoiseSizeY int `json:"size_vertical"`
	}{
		MinY:       n.MinY,
		Height:     n.Height,
		NoiseSizeX: n.NoiseSizeX,
		NoiseSizeY: n.NoiseSizeY,
	}

	return json.MarshalIndent(struct {
		SeaLevel           int                 `json:"sea_level"`
		CreaturesEnabled   bool                `json:"disable_mob_generation"`
		VeinsEnabled       bool                `json:"ore_veins_enabled"`
		AquifersEnabled    bool                `json:"aquifers_enabled"`
		LegacyRandomSource bool                `json:"legacy_random_source"`
		DefaultBlock       traversal.Construct `json:"default_block"`
		DefaultFluid       traversal.Construct `json:"default_fluid"`
		Noise              interface{}         `json:"noise"`
	}{
		SeaLevel:           n.SeaLevel,
		CreaturesEnabled:   n.DisableCreatures,
		VeinsEnabled:       !n.DisableVeins,
		AquifersEnabled:    !n.DisableAquifers,
		LegacyRandomSource: n.LegacyRandomSource,
		DefaultBlock:       n.DefaultBlock,
		DefaultFluid:       n.DefaultFluid,
		Noise:              noise,
	}, "", "  ")
}

func (n *NoiseSettings) ExportSymbol(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	data, err := json.MarshalIndent(n, "", "  ")
	if err != nil {
		return err
	}
	rootDir.
		MkDir(symbol.GetReference().GetNamespace(), nil).
		MkDir("worldgen", nil).
		MkDir("noise_settings", nil).
		MkFile(symbol.GetReference().GetName()+".json", string(data), nil)
	return nil
}
