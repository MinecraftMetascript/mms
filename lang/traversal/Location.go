package traversal

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type Location struct {
	Line int
	Col  int
}

type TextLocation struct {
	Start    Location
	StartIdx int
	Stop     Location
	StopIdx  int
	Text     string
	Filename string
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
		Stop: Location{
			Line: ctx.GetSymbol().GetLine(),
			Col:  ctx.GetSymbol().GetColumn() + len(ctx.GetText()),
		},
		Text:     ctx.GetText(),
		Filename: filename,
	}
}

func RuleLocation(ctx antlr.ParserRuleContext, filename string) TextLocation {
	return TextLocation{
		Start: Location{
			Line: ctx.GetStart().GetLine(),
			Col:  ctx.GetStart().GetColumn(),
		},
		StartIdx: ctx.GetStart().GetStart(),
		Stop: Location{
			Line: ctx.GetStop().GetLine(),
			Col:  ctx.GetStop().GetColumn(),
		},
		StopIdx:  ctx.GetStop().GetStop(),
		Text:     ctx.GetText(),
		Filename: filename,
	}
}
