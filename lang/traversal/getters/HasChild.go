package getters

import (
	"fmt"
	"reflect"

	"github.com/antlr4-go/antlr/v4"
)

func Has[ChildType antlr.ParserRuleContext, Parent antlr.ParserRuleContext](ctx Parent) bool {
	fmt.Println(reflect.TypeOf(ctx).Elem().Name())
	for _, child := range ctx.GetChildren() {
		fmt.Println("\t", reflect.TypeOf(child).Elem().Name())
		if _, ok := child.(ChildType); ok {
			fmt.Println("\t\t", "Found")
			return true
		}
	}

	return false
}

func AnyHas[ChildType antlr.ParserRuleContext, Parent antlr.ParserRuleContext](ctx []Parent) bool {
	for _, c := range ctx {
		for _, child := range c.GetChildren() {
			if _, ok := child.(ChildType); ok {
				return true
			}
		}
	}
	return false
}
