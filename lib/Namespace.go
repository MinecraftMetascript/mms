package lib

import (
	"encoding/json"
	"fmt"
)

type SymbolKind string

const (
	SymbolKindNoise            = "WorldGen__Noise"
	SymbolKindSurfaceRule      = "WorldGen__SurfaceRule"
	SymbolKindSurfaceCondition = "WorldGen__SurfaceCondition"
)

type Symbol[K json.Marshaler] struct {
	fmt.Stringer
	Location TextLocation `json:"location,inline"`
	File     string       `json:"source"`
	Ref      Reference    `json:"reference"`
	Value    K            `json:"value"`
	Kind     SymbolKind   `json:"kind"`
}

func NewNamespace() *Namespace {
	return &Namespace{
		symbols: make(map[string]Symbol[json.Marshaler]),
	}
}

type Namespace struct {
	symbols map[string]Symbol[json.Marshaler]
}

func (n *Namespace) Get(name string) (Symbol[json.Marshaler], bool) {
	v, ok := n.symbols[name]
	return v, ok
}

func (n *Namespace) Set(name string, value Symbol[json.Marshaler]) error {
	if _, ok := n.symbols[name]; ok {
		return fmt.Errorf("symbol '%s' already exists", name)
	}
	n.symbols[name] = value
	return nil
}

type MergeIssue struct {
	Location TextLocation
	File     string
	Message  string
}

func (n *Namespace) Merge(other *Namespace) []MergeIssue {
	issues := make([]MergeIssue, 0)
	for name, value := range other.symbols {
		if _, ok := n.symbols[name]; ok {
			issues = append(issues, MergeIssue{
				Location: value.Location,
				File:     value.File,
				Message:  "Duplicate symbol",
			},
			)
		}
		n.symbols[name] = value
	}

	if len(issues) > 0 {
		return issues
	}
	return nil
}

func AllOf[T json.Marshaler](n *Namespace) map[string]Symbol[T] {
	out := make(map[string]Symbol[T])
	for name, sym := range n.symbols {
		// Check if the value in the symbol is of type T
		value, ok := sym.Value.(T)
		if !ok {
			continue
		}
		// Create a new symbol with the value cast as T
		newSymbol := Symbol[T]{
			File:     sym.File,
			Location: sym.Location,
			Ref:      sym.Ref,
			Value:    value,
			Kind:     sym.Kind,
		}
		out[name] = newSymbol
	}
	return out
}
