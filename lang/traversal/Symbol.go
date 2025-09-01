package traversal

import "encoding/json"

type Symbol interface {
	GetNameLocation() TextLocation
	GetContentLocation() TextLocation
	GetValue() Construct
	GetReference() *Reference
}

type BaseSymbol struct {
	nameLocation    TextLocation
	contentLocation TextLocation
	value           Construct
	ref             *Reference
}

func NewSymbol(nameLocation TextLocation, contentLocation TextLocation, value Construct, ref *Reference) BaseSymbol {
	return BaseSymbol{
		nameLocation:    nameLocation,
		contentLocation: contentLocation,
		value:           value,
		ref:             ref,
	}
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
		Ref             Reference    `json:"ref"`
	}{
		NameLocation:    s.nameLocation,
		ContentLocation: s.contentLocation,
		Value:           s.GetValue(),
		Ref:             *s.ref,
	})
}
