package ast

import (
	"github.com/antlr4-go/antlr/v4"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type Location struct {
	// Line 0-indexed
	Line int `json:"line"`
	// Column 0-indexed
	Column int `json:"column"`
	// Index Character index from the beginning of the file
	Index int `json:"index"`
}

func (l Location) ToLspPosition() protocol.Position {
	return protocol.Position{
		Line:      protocol.UInteger(l.Line - 1),
		Character: protocol.UInteger(l.Column),
	}
}

type SourceLocation struct {
	Start    Location `json:"start"`
	Stop     Location `json:"stop"`
	Filename string   `json:"file"`
}

func (sl SourceLocation) ToLspRange() protocol.Range {
	return protocol.Range{
		Start: sl.Start.ToLspPosition(),
		End:   sl.Stop.ToLspPosition(),
	}
}

func (sl SourceLocation) Intersects(other SourceLocation) bool {
	return sl.ContainsLocation(other.Start) || sl.ContainsLocation(other.Stop)
}

func (sl SourceLocation) Contains(other SourceLocation) bool {
	return sl.ContainsLocation(other.Start) && sl.ContainsLocation(other.Stop)
}

func (sl SourceLocation) ContainsLocation(l Location) bool {
	lineContained := sl.Start.Line <= l.Line && sl.Stop.Line >= l.Line
	columnContained := sl.Start.Column <= l.Column && l.Column >= l.Column
	return lineContained && columnContained
}

func (sl SourceLocation) ContainsPosition(p protocol.Position) bool {
	return sl.ContainsLocation(Location{
		Line:   int(p.Line) + 1,
		Column: int(p.Character),
	})
}

func TokenStart(ctx antlr.Token) Location {
	return Location{
		Line:   ctx.GetLine(),
		Column: ctx.GetColumn(),
		Index:  ctx.GetStart(),
	}
}
func TokenStop(ctx antlr.Token) Location {
	return Location{
		Line:   ctx.GetLine(),
		Column: ctx.GetColumn(),
		Index:  ctx.GetStop(),
	}
}

func RuleLocation(ctx antlr.ParserRuleContext) SourceLocation {
	return SourceLocation{
		Start: TokenStart(ctx.GetStart()),
		Stop:  TokenStop(ctx.GetStop()),
	}
}

func TerminalLocation(ctx antlr.TerminalNode) SourceLocation {
	return SourceLocation{
		Start: TokenStart(ctx.GetSymbol()),
		Stop:  TokenStop(ctx.GetSymbol()),
	}
}
