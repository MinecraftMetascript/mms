package ast

import (
	"fmt"
	"strings"

	"github.com/minecraftmetascript/mms/lib"
	"github.com/samber/lo"
)

type TagKind string

const (
	TagBlock                    = "Block"
	TagItem                     = "Item"
	TagFunction                 = "Function"
	TagFluid                    = "Fluid"
	TagEntityType               = "EntityType"
	TagGameEvent                = "GameEvent"
	TagBiome                    = "Biome"
	TagFlatLevelGeneratorPreset = "FlatLevelGeneratorPreset"
	TagWorldPreset              = "WorldPreset"
	TagStructure                = "Structure"
	TagCatVariant               = "CatVariant"
	TagPOIType                  = "POIType"
	TagPaintingVariant          = "PaintingVariant"
	TagBannerPattern            = "BannerPattern"
	TagInstrument               = "Instrument"
	TagDamageType               = "DamageType"
	TagEnchantment              = "Enchantment"
	TagDialog                   = "Dialog"
)

type SymbolKind string

const (
	SymbolNever            SymbolKind = "Never" // Used for things like block references that MMS doesn't resolve
	SymbolNoise            SymbolKind = "Noise"
	SymbolDimensionType    SymbolKind = "DimensionType"
	SymbolNoiseRouter      SymbolKind = "NoiseRouter"
	SymbolNoiseSettings    SymbolKind = "NoiseSettings"
	SymbolDimension        SymbolKind = "Dimension"
	SymbolSurfaceRule      SymbolKind = "SurfaceRule"
	SymbolSurfaceCondition SymbolKind = "SurfaceCondition"
	SymbolDensityFunction  SymbolKind = "DensityFunction"
)

type Symbol interface {
	Node
	GetKind() SymbolKind
	ToFileTreeLike(name string) *lib.FileTreeLike
	ToSerializable() any

	GetNameLocation() *SourceLocation
	SetNameLocation(loc *SourceLocation)
	SetFilename(file string)

	SetDocstring(doc string)
	GetDocstring() string

	Ref() string
	SetRef(ref string)
}

type BaseSymbol struct {
	BaseNode
	Location     *SourceLocation `json:"location"`
	NameLocation *SourceLocation `json:"nameLocation"`
	Kind         SymbolKind      `json:"kind"`
	DocString    string          `json:"doc"`
	ref          string
}

type EmptySymbol struct {
	BaseSymbol
}

func (n *BaseSymbol) Ref() string {
	return n.ref
}
func (n *BaseSymbol) SetRef(ref string) {
	n.ref = ref
}

func (n *BaseSymbol) GetDocstring() string {
	return n.DocString
}
func (n *BaseSymbol) SetDocstring(doc string) {
	final := doc
	final = strings.TrimPrefix(final, "/**")
	final = strings.TrimSuffix(final, "*/")

	final = strings.Join(lo.Map(strings.Split(final, "\n"), func(line string, index int) string {
		return strings.TrimPrefix(strings.TrimLeft(line, " \t"), "* ")
	}), "\n")

	n.DocString = final
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
func (n *BaseSymbol) ToFileTreeLike(_ string) *lib.FileTreeLike {
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

func InlineSymbolId(s Symbol) string {
	l := s.GetLocation()

	fn := l.Filename
	fn = strings.TrimPrefix(fn, "file://")
	fn = strings.ReplaceAll(fn, "/", "_")
	fn = strings.TrimSuffix(fn, ".mms")

	return fmt.Sprintf("%s_%s__%d", s.GetKind(), fn, l.Start.Index)
}
