package traversal

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lang/grammar"
)

type Symbol interface {
	GetNameLocation() TextLocation
	GetContentLocation() TextLocation
	GetValue() Construct
	GetReference() *Reference
	GetType() string
}

type BaseSymbol struct {
	nameLocation    TextLocation
	contentLocation TextLocation
	value           Construct
	ref             *Reference
	_type           string
}

func (s BaseSymbol) GetType() string {
	return s._type
}

func NewSymbol(nameLocation TextLocation, contentLocation TextLocation, value Construct, ref *Reference) BaseSymbol {
	return BaseSymbol{
		nameLocation:    nameLocation,
		contentLocation: contentLocation,
		value:           value,
		ref:             ref,
	}
}

type DeclarationContext interface {
	antlr.ParserRuleContext
	Identifier() antlr.TerminalNode
}

func ProcessDeclaration(ctx DeclarationContext, valueCtx antlr.ParserRuleContext, scope *Scope, namespace string) Symbol {
	out := &BaseSymbol{}
	out._type = grammar.MinecraftMetascriptParserStaticData.RuleNames[valueCtx.GetRuleIndex()]
	out._type = fmt.Sprintf("%s%s", strings.ToUpper(string(out._type[0])), out._type[1:])

	if id := ctx.Identifier(); id == nil {
		scope.DiagnoseSemanticError("Missing Identifier", ctx)
	} else {
		out.ref = NewReference(id.GetText(), namespace)
		out.nameLocation = TerminalNodeLocation(id, scope.CurrentFile)
	}

	out.value = ConstructRegistry.Construct(valueCtx, namespace, scope)
	if out.value == nil {
		scope.DiagnoseSemanticError("Missing value", valueCtx)
	} else {
		out.contentLocation = RuleLocation(valueCtx, scope.CurrentFile)
	}

	err := scope.Register(*out)
	if err != nil {
		return nil
	}

	return out

}

func (s BaseSymbol) GetNameLocation() TextLocation {
	return s.nameLocation
}

func (s BaseSymbol) GetContentLocation() TextLocation {
	return s.contentLocation
}

func (s BaseSymbol) GetValue() Construct {
	return s.value
}

func (s BaseSymbol) GetReference() *Reference {
	return s.ref
}

func (s BaseSymbol) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		NameLocation    TextLocation `json:"nameLocation"`
		ContentLocation TextLocation `json:"contentLocation"`
		Value           Construct    `json:"value"`
		Ref             Reference    `json:"ref"`
	}{
		NameLocation:    s.nameLocation,
		ContentLocation: s.contentLocation,
		Value:           s.GetValue(),
		Ref:             *s.ref,
	})
}
