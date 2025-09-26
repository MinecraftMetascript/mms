package completions

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lib"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func AppendIfOccurances[T antlr.ParserRuleContext](
	root antlr.ParserRuleContext,
	items []protocol.CompletionItem,
	item protocol.CompletionItem,
	n int,
) []protocol.CompletionItem {
	if n > 1 && len(lib.GetAntlrChildren[T](root)) > n {
		// Don't append if we have more than n occurances
		return items
	}
	if n == 1 && lib.HasAntlrChild[T](root) {
		// Don't append if we have exactly one occurrence
		return items
	}
	return append(items, item)
}
