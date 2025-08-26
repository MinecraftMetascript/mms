package lib

import (
	"encoding/json"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lang/grammars"
)

type Reference struct {
	json.Marshaler
	Namespace string
	Name      string
}

func (r Reference) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s:%s\"", r.Namespace, r.Name)), nil
}

func (r Reference) String() string {
	return fmt.Sprintf("%s:%s", r.Namespace, r.Name)
}

type Referential interface {
	Reference() grammars.IReferenceContext
	Identifier() antlr.TerminalNode
}

func ParseReferential(defaultNamespace string, ctx Referential) Reference {
	result := Reference{
		Namespace: "",
		Name:      "",
	}
	// Reference has 2 parts (namespace:identifier)
	if ref := ctx.Reference(); ref != nil {
		// Get the namespace
		child := ref.GetChild(0)
		switch child.(type) {
		case antlr.TerminalNode:
			result.Namespace = child.(antlr.TerminalNode).GetText()
		default:
			result.Namespace = defaultNamespace
		}

		child = ref.GetChild(2)
		switch child.(type) {
		case antlr.TerminalNode:
			result.Name = child.(antlr.TerminalNode).GetText()
		case antlr.RuleNode:
			result.Name = child.(antlr.RuleNode).GetText()
		}

	} else if id := ctx.Identifier(); id != nil {
		result.Namespace = defaultNamespace
		result.Name = id.GetText()
	}
	return result
}
