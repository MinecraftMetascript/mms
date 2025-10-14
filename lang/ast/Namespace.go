package ast

import (
	"encoding/json"
	"fmt"
)

type Namespace struct {
	Declarations map[string]Symbol `json:"declarations"`
	Name         string            `json:"name"`
}

func (ns *Namespace) Size() int {
	if ns == nil {
		return 0
	}
	if ns.Declarations == nil {
		return 0
	}
	return len(ns.Declarations)
}

func (ns *Namespace) MarshalJSON() ([]byte, error) {
	out := map[string]any{}
	for n, d := range ns.Declarations {
		res := MarshalSymbol(d, n, ns.Name)
		out[n] = res

	}

	return json.Marshal(out)
}

func NewNamespace(name string) *Namespace {
	return &Namespace{Declarations: make(map[string]Symbol), Name: name}
}

func (ns *Namespace) GetDecl(name string) Symbol {
	return ns.Declarations[name]
}

func (ns *Namespace) Declare(name string, symbol Symbol) *Diagnostic {
	if _, ok := ns.Declarations[name]; ok {

		return &Diagnostic{
			Location: *symbol.GetLocation(),
			Message:  fmt.Sprintf("Duplicate identifier %s:%s", ns.Name, name),
			Severity: Error,
		}

	} else {

		ns.Declarations[name] = symbol
		return nil
	}
}

func (ns *Namespace) Merge(other *Namespace) {
	for name, symbol := range other.Declarations {
		ns.Declare(name, symbol)
	}
}

func (ns *Namespace) AllDecls() map[string]Symbol {
	return ns.Declarations
}

func (ns *Namespace) Delete(name string) {
	delete(ns.Declarations, name)
}
