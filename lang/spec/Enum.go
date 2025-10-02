package spec

import (
	"github.com/minecraftmetascript/mms/lang/ast"
	"github.com/minecraftmetascript/mms/lang/grammar"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type EnumSpec struct {
	Options []string
	Help    string
}

func NewEnumSpec(options ...string) EnumSpec {
	return EnumSpec{Options: options}
}

func (e EnumSpec) Match(valueCtx grammar.IValueContext) (ast.Node, []ast.Diagnostic) {
	// Enums will be parsed in as resource references
	if valueCtx.ResourceReference() == nil {
		return nil, nil
	}
	rr := valueCtx.ResourceReference()

	if len(rr.AllIdentifier()) != 1 {
		// Not a match
		return nil, nil
	}
	value := rr.Identifier(0).GetText()

	out := &EnumNode{spec: e, Value: value}

	valueValid := false
	for _, option := range e.Options {
		if option == value {
			valueValid = true
			break
		}
	}
	if !valueValid {
		// TODO: Diagnose
	}
	return out, nil
}

func (e EnumSpec) SetHelp(help string) EnumSpec {
	out := &e
	out.Help = help
	return *out
}

type EnumNode struct {
	ast.BaseSymbol
	Value string
	spec  EnumSpec
}

func (e EnumNode) GetHelp() string {
	return e.spec.Help
}

func (e EnumNode) Complete(fileSource string, position protocol.Position, triggerChar *string, symbols map[string]*ast.Namespace) []protocol.CompletionItem {
	/*
		Steps:
			- Identify if there is already some text
			- If there is, filter the enum options by that text
			- If there is not, return all enum options
	*/
	return make([]protocol.CompletionItem, 0)
}

func GetEnumNodeValue(n ast.Node) *string {
	if n == nil {
		return nil
	}
	if n, ok := n.(*EnumNode); ok {
		return &n.Value
	}
	return nil
}
