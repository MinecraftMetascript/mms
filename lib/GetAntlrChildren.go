package lib

import "github.com/antlr4-go/antlr/v4"

func GetAntlrChildren[T antlr.ParserRuleContext](root antlr.ParserRuleContext) []T {
	children := make([]T, 0)
	for _, child := range root.GetChildren() {
		if child == nil {
			continue
		}
		if r, ok := child.(T); ok {
			children = append(children, r)
		} else if r, ok := child.(antlr.ParserRuleContext); ok {
			children = append(children, GetAntlrChildren[T](r)...)
		}
	}
	return children
}

func HasAntlrChild[T antlr.ParserRuleContext](root antlr.ParserRuleContext) bool {
	return len(GetAntlrChildren[T](root)) > 0
}
