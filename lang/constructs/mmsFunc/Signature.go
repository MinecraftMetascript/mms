package mmsFunc

import (
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
)

func init() {
	traversal.Register(
		func(ctx *grammar.FuncOfContext, ns string, scope *traversal.Scope) traversal.Construct {
			out := &Signature{
				params: make([]Param, 0),
			}
			for _, param := range ctx.AllParamDef() {
				name := param.Variable()
				if name == nil {
					scope.DiagnoseSemanticError("Found unnamed parameter", param)
				}
				kind := param.FuncParamType()
				if kind == nil {
					scope.DiagnoseSemanticError("Found untyped parameter", param)
				}
				out.params = append(out.params, Param{
					Kind: kind.GetText(),
					Name: name.GetText(),
				})
			}
			return out
		},
	)

}

type Param struct {
	Kind string
	Name string
}
type Signature struct {
	params []Param
}

func (s Signature) MarshalJSON() ([]byte, error) {
	return []byte("{\"_\": \"MMS Func cannot be serialized\"}"), nil
}

func (s Signature) ExportSymbol(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	return nil
}

func (s Signature) Params() []Param {
	return s.params
}
