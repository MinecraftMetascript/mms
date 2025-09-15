package getters

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lang/traversal"
)

func GetChildNode(ctx antlr.ParserRuleContext, namespace string, scope *traversal.Scope, missingMessage string) traversal.Node {
	if ctx != nil {
		val := traversal.ConstructNode(ctx, namespace, scope)
		if val == nil {
			scope.DiagnoseSemanticError(missingMessage, ctx)
		} else {
			return val
		}
	} else {
		scope.DiagnoseSemanticError(missingMessage, ctx)
	}
	return nil
}
