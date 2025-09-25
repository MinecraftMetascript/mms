package completions

import (
	"github.com/antlr4-go/antlr/v4"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type CursorWithinType string

const (
	CursorWithinEmpty   CursorWithinType = "empty"
	CursorWithinContent CursorWithinType = "content"
	CursorAtStart       CursorWithinType = "start"
	CursorAtEnd         CursorWithinType = "end"
	CursorNotWithin     CursorWithinType = "not_within"
)

func CursorWithinRecursive(ctx antlr.ParserRuleContext, cursorPosition protocol.Position, start string, end string) CursorWithinType {
	rootResult := CursorWithin(ctx, cursorPosition, start, end)
	if rootResult != CursorNotWithin {
		return rootResult
	}
	for _, child := range ctx.GetChildren() {
		if ctx, ok := child.(antlr.ParserRuleContext); ok {
			result := CursorWithinRecursive(ctx, cursorPosition, start, end)
			if result != CursorNotWithin {
				return result
			}
		}
	}
	return CursorNotWithin
}

func CursorWithin(ctx antlr.ParserRuleContext, cursorPosition protocol.Position, start string, end string) CursorWithinType {
	children := ctx.GetChildren()
	for i, child := range children {
		if startNode, ok := child.(antlr.TerminalNode); ok && startNode.GetText() == start {
			for j, endChild := range children[i+1:] {
				if endNode, ok := endChild.(antlr.TerminalNode); ok && endNode.GetText() == end {
					cursorLine := int(cursorPosition.Line) + 1
					cursorChar := int(cursorPosition.Character)
					startLine := startNode.GetSymbol().GetLine()
					endLine := endNode.GetSymbol().GetLine()
					startChar := startNode.GetSymbol().GetColumn()
					endChar := endNode.GetSymbol().GetColumn()

					if cursorLine < startLine || cursorLine > endLine {
						return CursorNotWithin
					}
					if cursorChar < startChar || cursorChar > endChar {
						return CursorNotWithin
					}
					if j == 0 {
						// There is nothing between the start and end, but the cursor is between them
						return CursorWithinEmpty
					}
					if cursorLine == startLine && cursorChar == startChar {
						// Cursor is at the start
						return CursorAtStart
					}
					if cursorLine == endLine && cursorChar == endChar {
						// Cursor is at the end, and within the separaters
						return CursorAtEnd
					}
					return CursorWithinContent
				}
			}
		}
	}
	return CursorNotWithin
}
