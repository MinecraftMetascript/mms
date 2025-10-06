package ast

import (
	"fmt"

	"github.com/minecraftmetascript/mms/lib"
)

type SymbolKind string

const (
	SymbolNever            SymbolKind = "Never" // Used for things like block references that MMS doesn't resolve
	SymbolNoise            SymbolKind = "Noise"
	SymbolSurfaceRule      SymbolKind = "SurfaceRule"
	SymbolSurfaceCondition SymbolKind = "SurfaceCondition"
)

type Symbol interface {
	Node
	GetKind() SymbolKind
	ToFileTreeLike(name string) *lib.FileTreeLike
	ToSerializable() any

	GetNameLocation() *SourceLocation
	SetNameLocation(loc *SourceLocation)
	SetFilename(file string)
}

type BaseSymbol struct {
	BaseNode
	Location     *SourceLocation `json:"location"`
	NameLocation *SourceLocation `json:"nameLocation"`
	Kind         SymbolKind      `json:"kind"`
}

type EmptySymbol struct {
	BaseSymbol
}

func (n *BaseSymbol) GetLocation() *SourceLocation {
	return n.Location
}

func (n *BaseSymbol) GetNameLocation() *SourceLocation {
	return n.NameLocation
}
func (n *BaseSymbol) SetNameLocation(loc *SourceLocation) {
	n.NameLocation = loc
}
func (n *BaseSymbol) GetKind() SymbolKind {
	return n.Kind
}

func (n *BaseSymbol) SetFilename(file string) {
	if n.NameLocation != nil {
		n.NameLocation.Filename = file
	}
	if n.Location != nil {
		n.Location.Filename = file
	}
}
func (n *BaseSymbol) ToSerializable() any {
	return ""
}
func (n *BaseSymbol) ToFileTreeLike(name string) *lib.FileTreeLike {
	return nil
}

type SymbolExport struct {
	Value        any             `json:"value"`
	Kind         SymbolKind      `json:"kind"`
	NameLocation *SourceLocation `json:"nameLocation"`
	Location     *SourceLocation `json:"location"`
	Ref          string          `json:"ref"`
}

func MarshalSymbol(s Symbol, name, namespace string) SymbolExport {
	return SymbolExport{
		Value:        s.ToSerializable(),
		Kind:         s.GetKind(),
		NameLocation: s.GetNameLocation(),
		Location:     s.GetLocation(),
		Ref:          fmt.Sprintf("%s:%s", namespace, name),
	}

}
