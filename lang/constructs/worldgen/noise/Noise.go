package noise

import (
	"encoding/json"
	"strconv"

	"github.com/minecraftmetascript/mms/lang/constructs/primitives"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lang/traversal/getters"
	"github.com/minecraftmetascript/mms/lib"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type NoiseFactory struct{}

func (n NoiseFactory) CreateDeclaration(ctx *grammar.NoiseDeclarationContext, namespace string, scope *traversal.Scope) (traversal.Symbol, bool) {
	if ctx.Declare() != nil {
		declaration := primitives.DeclarationFactory{}.Create(ctx.Declare().(*grammar.DeclareContext), namespace, scope)
		noise := ctx.Noise()
		if noise == nil {
			// TODO: Diagnose?
			return nil, false
		}

		value := traversal.ConstructNode(noise.(*grammar.NoiseContext), namespace, scope)

		return traversal.NewSymbol(
			declaration.GetNameLocation(),
			traversal.RuleLocation(noise, scope.CurrentFile),
			value,
			declaration.GetReference(),
			traversal.Noise,
		), true
	}
	return nil, false
}

func (n NoiseFactory) Create(ctx *grammar.NoiseContext, namespace string, scope *traversal.Scope) *Noise {
	out := &Noise{
		Amplitudes: make([]float64, 0),
		location:   traversal.RuleLocation(ctx, scope.CurrentFile),
		ctx:        ctx,
	}

	def := ctx.NoiseDefinition()
	if def == nil {
		scope.DiagnoseSemanticError("Missing noise definition", ctx)
		return out
	}

	getters.GetInt(
		def, func(i int) { out.FirstOctave = i }, scope, "FirstOctave",
	)

	for _, r := range def.AllNoise_Builder() {
		if amplitudes := r.Builder_Amplitudes(); amplitudes != nil {
			for _, amplitude := range amplitudes.AllNumber() {
				if amplitude != nil {
					val, err := strconv.ParseFloat(amplitude.GetText(), 64)
					if err != nil {
						scope.DiagnoseSemanticError(
							"Invalid amplitude value",
							amplitude,
						)
					} else {
						out.Amplitudes = append(out.Amplitudes, val)
					}
				}
			}
		} else {
			scope.DiagnoseSemanticError("Missing amplitudes value", r)
		}
	}

	return out
}

func (n NoiseFactory) GetHelp(node *Noise, _ traversal.Symbol, _ traversal.TextLocation) *traversal.Help {
	return &traversal.Help{
		Content:  "Defines a noise function.<br/>`FirstOctave(int)` is required.<br/>`Amplitudes(float)` is required.<br/>Example: `Noise(FirstOctave(100), Amplitudes(0.5, 0.5))`",
		Position: node.GetLocation(),
	}
}

func (n NoiseFactory) Export(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	data, err := json.MarshalIndent(symbol.GetValue(), "", "  ")
	if err != nil {
		return err
	}
	rootDir.
		MkDir(symbol.GetReference().GetNamespace(), nil).
		MkDir("worldgen", nil).
		MkDir("noise", nil).
		MkFile(symbol.GetReference().GetName()+".json", string(data), nil)
	return nil
}

func init() {
	traversal.RegisterNodeFactory(NoiseFactory{}, true)
}

type Noise struct {
	FirstOctave int       `json:"firstOctave"`
	Amplitudes  []float64 `json:"amplitudes"`
	location    traversal.TextLocation
	ctx         *grammar.NoiseContext
}

func (n Noise) GetLocation() traversal.TextLocation {
	return n.location
}

func (n Noise) GetCompletions(_ protocol.Position) []protocol.CompletionItem {

	out := make([]protocol.CompletionItem, 0)
	if n.Amplitudes == nil || len(n.Amplitudes) == 0 {
		out = append(out, protocol.CompletionItem{
			Label:  "Amplitude",
			Kind:   lib.Ptr(protocol.CompletionItemKindMethod),
			Detail: lib.Ptr("Define the amplitudes of this noise function"),
			TextEdit: protocol.TextEdit{
				Range: protocol.Range{
					Start: n.location.Stop.OffsetColumn(1).ToLspPosition(),
					End:   n.location.Stop.OffsetColumn(1).ToLspPosition(),
				},
				NewText: ".Amplitudes(${0})",
			},
			InsertTextFormat: lib.Ptr(protocol.InsertTextFormatSnippet),
		})
	}

	return out

}
