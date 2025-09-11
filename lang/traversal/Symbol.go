package traversal

import (
	"encoding/json"

	"github.com/antlr4-go/antlr/v4"
)

type Symbol interface {
	GetNameLocation() TextLocation
	GetContentLocation() TextLocation
	GetValue() Construct
	GetReference() *Reference
	GetType() string
}

type BaseSymbol struct {
	nameLocation    TextLocation
	contentLocation TextLocation
	value           Construct
	ref             *Reference
	kind            string
}

func (s BaseSymbol) GetType() string {
	return s.kind
}

func NewSymbol(nameLocation TextLocation, contentLocation TextLocation, value Construct, ref *Reference, kind string) BaseSymbol {
	return BaseSymbol{
		nameLocation:    nameLocation,
		contentLocation: contentLocation,
		value:           value,
		ref:             ref,
		kind:            kind,
	}
}

type DeclarationContext interface {
	antlr.ParserRuleContext
	Identifier() antlr.TerminalNode
}

func ProcessDeclaration(ctx DeclarationContext, valueCtx antlr.ParserRuleContext, scope *Scope, namespace string, kind string) Symbol {
	out := &BaseSymbol{}
	out.kind = kind

	if id := ctx.Identifier(); id == nil {
		scope.DiagnoseSemanticError("Missing Identifier", ctx)
	} else {
		out.ref = NewReference(id.GetText(), namespace)
		out.nameLocation = TerminalNodeLocation(id, scope.CurrentFile)
	}

	out.value = ConstructRegistry.Construct(valueCtx, namespace, scope)
	if out.value == nil {
		scope.DiagnoseSemanticError("No Value found for Declaration", ctx)
	} else {
		out.contentLocation = RuleLocation(valueCtx, scope.CurrentFile)
	}

	err := scope.Register(*out)
	if err != nil {
		return nil
	}

	return out

}

func (s BaseSymbol) GetNameLocation() TextLocation {
	return s.nameLocation
}

func (s BaseSymbol) GetContentLocation() TextLocation {
	return s.contentLocation
}

func (s BaseSymbol) GetValue() Construct {
	return s.value
}

func (s BaseSymbol) GetReference() *Reference {
	return s.ref
}

func (s BaseSymbol) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		NameLocation    TextLocation `json:"nameLocation"`
		ContentLocation TextLocation `json:"contentLocation"`
		Value           Construct    `json:"value"`
		Ref             *Reference   `json:"ref"`
		Type            string       `json:"type"`
	}{
		NameLocation:    s.nameLocation,
		ContentLocation: s.contentLocation,
		Value:           s.GetValue(),
		Ref:             s.ref,
		Type:            s.GetType(),
	})
}
