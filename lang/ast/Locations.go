package ast

import (
	"fmt"
	"math"

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

func (l Location) ColOffset(off int) Location {
	out := &l
	out.Column += off
	return *out
}

func (l Location) String() string {
	return fmt.Sprintf("(%d,%d)", l.Line, l.Column)
}

func (l Location) IndexIn(content string) int {
	lines := 1
	index := 0
	for i := 0; i < len(content); i++ {
		if lines == l.Line && index == l.Column {
			return i
		}
		if content[i] == '\n' {
			lines++
			index = 0
		} else {
			index++
		}
	}
	return len(content)
}

type SourceLocation struct {
	Start    Location `json:"start"`
	Stop     Location `json:"stop"`
	Filename string   `json:"file"`
}

func (sl SourceLocation) String() string {
	return fmt.Sprintf("[(%s) - (%s)]", sl.Start, sl.Stop)
}

func (sl SourceLocation) Distance(idx int) int {
	startDist := math.Abs(float64(sl.Start.Index - idx))
	stopDist := math.Abs(float64(sl.Stop.Index - idx))

	return int(math.Min(startDist, stopDist))
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
	if !lineContained {
		return false
	}
	if sl.Start.Line == l.Line && sl.Start.Column >= l.Column {
		return false
	}
	if sl.Stop.Line == l.Line && sl.Stop.Column <= l.Column {
		return false
	}
	return true
}

func (sl SourceLocation) ContainsPosition(p protocol.Position) bool {
	return sl.ContainsLocation(Location{
		Line:   int(p.Line) + 1,
		Column: int(p.Character),
	})
}

func (sl SourceLocation) BeforePosition(p protocol.Position) bool {
	return sl.Stop.Line < int(p.Line)+1 || (sl.Stop.Line == int(p.Line)+1 && sl.Stop.Column < int(p.Character))
}

func (sl SourceLocation) AfterPosition(p protocol.Position) bool {
	return sl.Start.Line > int(p.Line)+1 || (sl.Start.Line == int(p.Line)+1 && sl.Start.Column > int(p.Character))
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
		Column: ctx.GetColumn() + (ctx.GetStop() - ctx.GetStart()),
		Index:  ctx.GetStop(),
	}
}

func RuleLocation(ctx antlr.ParserRuleContext) SourceLocation {
	return SourceLocation{
		Start: TokenStart(ctx.GetStart()),
		Stop:  TokenStop(ctx.GetStop()).ColOffset(1),
	}
}

func TerminalLocation(ctx antlr.TerminalNode) SourceLocation {
	return SourceLocation{
		Start: TokenStart(ctx.GetSymbol()),
		Stop:  TokenStop(ctx.GetSymbol()).ColOffset(1),
	}
}
