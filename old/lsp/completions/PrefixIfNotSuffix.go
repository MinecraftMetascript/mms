package completions

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

func PrefixIfNotSuffix(ctx antlr.ParserRuleContext, s string, trail string) string {
	if CtxHasTrailingToken(ctx, trail) {
		return s
	}
	return fmt.Sprintf("%s%s", trail, s)
}
