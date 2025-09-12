package mmsFunc

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type MmsFunc interface {
	GetSignature() Signature
	GetContent() antlr.ParserRuleContext
}

var funcRegistry map[string]MmsFunc = make(map[string]MmsFunc)

func Register(name string, fn MmsFunc) {
	funcRegistry[name] = fn
	fmt.Println("Registered function: " + name)
}

func Get(name string) MmsFunc {
	return funcRegistry[name]
}
