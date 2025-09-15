package traversal

import (
	"encoding/json"

	"github.com/antlr4-go/antlr/v4"
)

type SymbolKind string

const (
	VerticalAnchor = "VerticalAnchor"
	Noise          = "Noise"
)

type Symbol interface {
	GetNameLocation() TextLocation
	GetContentLocation() TextLocation
	GetValue() Node
	GetReference() *Reference
	GetKind() SymbolKind
}

type BaseSymbol struct {
	nameLocation    TextLocation
	contentLocation TextLocation
	value           Node
	ref             *Reference
	kind            SymbolKind
}

func (s BaseSymbol) GetKind() SymbolKind {
	return s.kind
}

func NewEmptySymbol(nameLocation TextLocation, ref *Reference) BaseSymbol {
	return BaseSymbol{
		nameLocation: nameLocation,
		ref:          ref,
	}
}

func NewSymbol(nameLocation TextLocation, contentLocation TextLocation, value Node, ref *Reference, kind SymbolKind) BaseSymbol {
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

func (s BaseSymbol) GetNameLocation() TextLocation {
	return s.nameLocation
}

func (s BaseSymbol) GetContentLocation() TextLocation {
	return s.contentLocation
}

func (s BaseSymbol) GetValue() Node {
	return s.value
}

func (s BaseSymbol) GetReference() *Reference {
	return s.ref
}

func (s BaseSymbol) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		NameLocation    TextLocation `json:"nameLocation"`
		ContentLocation TextLocation `json:"contentLocation"`
		Value           Node         `json:"value"`
		Ref             *Reference   `json:"ref"`
		Type            SymbolKind   `json:"type"`
	}{
		NameLocation:    s.nameLocation,
		ContentLocation: s.contentLocation,
		Value:           s.GetValue(),
		Ref:             s.ref,
		Type:            s.GetKind(),
	})
}
