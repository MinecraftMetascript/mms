package noise

import (
	"encoding/json"
	"log"

	"github.com/minecraftmetascript/mms/lang/grammars"
	"github.com/minecraftmetascript/mms/lang/mms_errors"
	"github.com/minecraftmetascript/mms/lib"
)

func NewNoiseVisitor(
	reportError func(mmsError mms_errors.MMSError),
	filename string,
) *Visitor {
	return &Visitor{
		filename:          filename,
		NoiseDeclarations: make(map[string]lib.Symbol[Noise]),
		reportError:       reportError,
	}
}

type Visitor struct {
	grammars.BaseMMSParserListener
	reportError       func(mmsError mms_errors.MMSError)
	namespace         string
	filename          string
	NoiseDeclarations map[string]lib.Symbol[Noise]
}

func (v *Visitor) DumpDeclarations(ns *lib.Namespace) {
	for name, rule := range v.NoiseDeclarations {
		ns.Set(name, lib.Symbol[json.Marshaler]{
			Value: rule.Value,
			File:  rule.File,
			Ref:   rule.Ref,
			Line:  rule.Line,
			Col:   rule.Col,
			Kind:  rule.Kind,
		})
	}
}

func (v *Visitor) Namespace() string {
	return v.namespace
}
func (v *Visitor) SetNamespace(namespace string) {
	v.namespace = namespace
}

func (v *Visitor) ExitNoiseDeclaration(ctx *grammars.NoiseDeclarationContext) {
	n, err := NewNoise(ctx.NoiseDefinition().(*grammars.NoiseDefinitionContext))
	if err != nil {
		log.Println("Error:", err)
		return
	}

	id := ctx.Identifier().GetText()

	v.NoiseDeclarations[id] = lib.Symbol[Noise]{
		Value: *n,
		File:  v.filename,
		Ref: lib.Reference{
			Namespace: v.namespace,
			Name:      id,
		},
		Line: ctx.GetStart().GetLine(),
		Col:  ctx.GetStart().GetColumn(),
		Kind: lib.SymbolKindNoise,
	}
}
