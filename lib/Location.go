package lib

import "github.com/antlr4-go/antlr/v4"

type Location struct {
	Column int
	Line   int
}

type TextLocation struct {
	Start Location
	End   Location
}

func GetRuleLocation(ctx antlr.ParserRuleContext) TextLocation {
	start := ctx.GetStart()
	end := ctx.GetStop()

	return TextLocation{
		Start: Location{Column: start.GetColumn(), Line: start.GetLine()},
		End:   Location{Column: end.GetColumn(), Line: end.GetLine()},
	}
}

type LocatedText interface {
	GetLocation() TextLocation
	SetLocation(TextLocation)
}
