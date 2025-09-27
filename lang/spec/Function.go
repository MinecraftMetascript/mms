package spec

import (
	"github.com/minecraftmetascript/mms/lang/ast"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lib"
)

type OverloadSpec struct {
	Args     []ValueSpec
	RestArg  ValueSpec
	Builders []FunctionSpec
}

type Overload struct {
	Args     []ast.Node
	Builders []FunctionNode
}

func NewOverloadSpec(args []ValueSpec, restArg ValueSpec, builders []FunctionSpec) OverloadSpec {
	return OverloadSpec{
		Args:     args,
		RestArg:  restArg,
		Builders: builders,
	}
}

func (spec *OverloadSpec) Match(ctx grammar.IFnContext) (ast.Node, []ast.Diagnostic) {
	diags := make([]ast.Diagnostic, 0)
	out := &Overload{
		Args:     make([]ast.Node, len(spec.Args)),
		Builders: make([]FunctionNode, 0),
	}

	for i, argCtx := range ctx.AllValue() {
		var argSpec ValueSpec
		if i >= len(spec.Args) {
			if spec.RestArg == nil {
				diags = append(diags, ast.Diagnostic{
					Location: ast.RuleLocation(argCtx),
					Message:  "Unexpected argument",
					Severity: ast.Warning,
				})
			}
			argSpec = spec.RestArg
		} else {

			argSpec = spec.Args[i]
		}

		// TODO: We need to modify this to allow for error messages, but also
		// 			fallthrough when something doesn't match?
		value, argDiags := argSpec.Match(argCtx)
		if argDiags != nil && len(argDiags) > 0 {
			diags = append(diags, argDiags...)
		} else if value != nil {
			if i >= len(spec.Args) {
				// RestArgs
				out.Args = append(out.Args, value)
			} else {
				out.Args[i] = value
			}
		}
	}

	for _, builderCtx := range ctx.AllFn() {
		for _, builderSpec := range spec.Builders {
			builder, builderDiags := builderSpec.matchFn(builderCtx)

			if builderDiags != nil && len(builderDiags) > 0 {
				diags = append(diags, builderDiags...)
			} else if builder != nil {
				out.Builders = append(out.Builders, *builder)
			}
		}
	}

	if len(diags) > 0 {
		return out, diags
	} else {
		return out, nil
	}
}

type FunctionSpec struct {
	Name string
	Kind string
	// Arguments Slice of overloads
	Overloads []OverloadSpec

	export func(fn FunctionNode, name string) *lib.FileTreeLike
}

func (f FunctionSpec) matchFn(fnCtx grammar.IFnContext) (*FunctionNode, []ast.Diagnostic) {
	fnKindCtx := fnCtx.Identifier()
	if fnKindCtx == nil {
		// TODO: Diagnose?
		return nil, nil
	}
	fnKind := fnKindCtx.GetText()
	if fnKind != f.Name {
		// Not a match
		return nil, nil
	}

	for _, overload := range f.Overloads {
		res, diags := overload.Match(fnCtx)
		if res == nil {
			return nil, diags
		}
		return &FunctionNode{
			Name:      f.Name,
			Arguments: res.(*Overload).Args,
			Builders:  res.(*Overload).Builders,
			spec:      f,
			kind:      f.Kind,
		}, diags
	}

	return nil, []ast.Diagnostic{
		{
			Location: ast.RuleLocation(fnCtx),
			Message:  "No matching overload found",
			Severity: ast.Warning,
		},
	}

}

func (f FunctionSpec) Match(valueCtx grammar.IValueContext) (ast.Node, []ast.Diagnostic) {
	fnCtx := valueCtx.Fn()
	if fnCtx == nil {
		return nil, nil
	}
	return f.matchFn(fnCtx)
}

func (f FunctionSpec) SetExporter(exporter func(n FunctionNode, name string) *lib.FileTreeLike) FunctionSpec {
	out := &f
	out.export = exporter
	return *out
}
func (f FunctionSpec) SetKind(kind string) FunctionSpec {
	out := &f
	out.Kind = kind
	return *out
}

func NewFunctionSpec(kind string, overloads ...OverloadSpec) *FunctionSpec {
	return &FunctionSpec{
		Name:      kind,
		Overloads: overloads,
	}
}

type FunctionNode struct {
	Name      string
	Arguments []ast.Node
	Builders  []FunctionNode

	kind string
	spec FunctionSpec
}

func (n FunctionNode) Kind() string {
	return n.kind
}
func (n FunctionNode) Export(name string) *lib.FileTreeLike {
	return n.spec.export(n, name)
}
