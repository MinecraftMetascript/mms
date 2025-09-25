package traversal

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func FilterByLocation[T any](location TextLocation, m map[TextLocation]T) map[TextLocation]T {
	out := make(map[TextLocation]T, 0)
	for l, val := range m {
		if location.Contains(l) {
			out[l] = val
		}
	}
	return out
}

type Location struct {
	Line int
	Col  int
}

func (l Location) ToLspPosition() protocol.Position {
	return protocol.Position{
		Line:      protocol.UInteger(l.Line - 1), // Protocol positions are 0 indexed
		Character: protocol.UInteger(l.Col),
	}
}
func (l Location) OffsetLine(offset int) Location {
	return Location{
		Line: l.Line + offset,
		Col:  l.Col,
	}
}

func (l Location) OffsetColumn(offset int) Location {
	return Location{
		Line: l.Line,
		Col:  l.Col + offset,
	}
}

type TextLocation struct {
	Start    Location
	StartIdx int
	Stop     Location
	StopIdx  int
	Text     string
	Filename string
}

func (tl TextLocation) Contains(other TextLocation) bool {
	return tl.StartIdx <= other.StartIdx && tl.StopIdx >= other.StopIdx
}

func (tl TextLocation) ContainsLocation(location Location) bool {
	return tl.Start.Line <= location.Line && tl.Stop.Line >= location.Line
}

func (tl TextLocation) String() string {
	return fmt.Sprintf("(%d,%d)->(%d,%d)", tl.Start.Line, tl.Start.Col, tl.Stop.Line, tl.Stop.Col)
}

func TerminalNodeLocation(ctx antlr.TerminalNode, filename string) TextLocation {
	return TextLocation{
		Start: Location{
			Line: ctx.GetSymbol().GetLine(),
			Col:  ctx.GetSymbol().GetColumn(),
		},
		StartIdx: ctx.GetSymbol().GetStart(),
		Stop: Location{
			Line: ctx.GetSymbol().GetLine(),
			Col:  ctx.GetSymbol().GetColumn() + len(ctx.GetText()),
		},
		StopIdx:  ctx.GetSymbol().GetStop(),
		Text:     ctx.GetText(),
		Filename: filename,
	}
}

func RuleLocation(ctx antlr.ParserRuleContext, filename string) TextLocation {
	return TextLocation{
		Start:    RuleStart(ctx),
		StartIdx: ctx.GetStart().GetStart(),
		Stop:     RuleStop(ctx),
		StopIdx:  ctx.GetStop().GetStop(),
		Text:     ctx.GetText(),
		Filename: filename,
	}
}

func RuleStop(ctx antlr.ParserRuleContext) Location {
	return Location{
		Line: ctx.GetStop().GetLine(),
		Col:  ctx.GetStop().GetColumn(),
	}
}

func RuleStart(ctx antlr.ParserRuleContext) Location {
	return Location{
		Line: ctx.GetStart().GetLine(),
		Col:  ctx.GetStart().GetColumn(),
	}
}
