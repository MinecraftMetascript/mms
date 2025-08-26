package primitives

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lang/grammars"
)

type Reference struct {
	json.Marshaler
	Namespace string
	Name      string
}

func (r Reference) String() string {
	return fmt.Sprintf("%s:%s", r.Namespace, r.Name)
}

func (r Reference) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s:%s\"", r.Namespace, r.Name)), nil
}

func FromResourceReference(ctx grammars.IResourceReferenceContext, defaultNamespace string) (*Reference, error) {
	if ref := ctx.Reference(); ref != nil {
		return &Reference{
			Namespace: ref.GetChild(0).(antlr.TerminalNode).GetText(),
			Name:      ref.GetChild(2).(antlr.TerminalNode).GetText(),
		}, nil
	}
	if id := ctx.Identifier(); id != nil {
		return &Reference{
			Namespace: defaultNamespace,
			Name:      id.GetText(),
		}, nil
	}
	if keyword := ctx.Keyword(); keyword != nil {
		return &Reference{
			Namespace: defaultNamespace,
			Name:      keyword.GetText(),
		}, nil
	}

	return nil, errors.New("invalid resource reference")

}
