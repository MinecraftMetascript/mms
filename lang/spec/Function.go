package spec

import (
	"fmt"
	"log"
	"slices"

	"github.com/minecraftmetascript/mms/lang/ast"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lib"
	"github.com/samber/lo"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type OverloadSpec struct {
	Args     []ValueSpec
	RestArg  ValueSpec
	Builders []FunctionSpec
}

type Overload struct {
	ast.BaseNode
	Args     []ast.Node
	Builders []FunctionNode
}

func (o Overload) Children() []ast.Node {
	children := make([]ast.Node, 0)
	for _, a := range o.Args {
		children = append(children, a)
	}
	for _, b := range o.Builders {
		children = append(children, &b)
	}

	return children
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
	l := ast.RuleLocation(ctx)
	out := &Overload{
		Args:     make([]ast.Node, len(spec.Args)),
		Builders: make([]FunctionNode, 0),
		BaseNode: ast.BaseNode{
			Location: &l,
		},
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
		if argSpec == nil {
			diags = append(diags, ast.Diagnostic{
				Location: ast.RuleLocation(argCtx),
				Message:  "Unexpected argument",
				Severity: ast.Warning,
			})
			continue
		}
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

	builders := ctx.AllFn()
	for _, builderCtx := range builders {
		for _, builderSpec := range spec.Builders {
			builder, builderDiags := builderSpec.matchFn(builderCtx)

			if builderDiags != nil && len(builderDiags) > 0 {
				diags = append(diags, builderDiags...)
				if builder != nil {
					out.Builders = append(out.Builders, *builder)
				}
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
	Kind ast.SymbolKind
	// Arguments Slice of overloads
	Overloads []OverloadSpec
	Help      string

	export func(fn FunctionNode, name string) *lib.FileTreeLike
	output func(fn FunctionNode) any
}

func (f FunctionSpec) matchFn(fnCtx grammar.IFnContext) (*FunctionNode, []ast.Diagnostic) {
	l := ast.RuleLocation(fnCtx)

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

		out := &FunctionNode{
			Name:      f.Name,
			overload:  overload,
			Arguments: res.(*Overload).Args,
			Builders:  res.(*Overload).Builders,
			spec:      f,
			source:    fnCtx.GetStart().GetInputStream().GetText(fnCtx.GetStart().GetStart(), fnCtx.GetStop().GetStop()+1),
			BaseSymbol: ast.BaseSymbol{
				Kind:     f.Kind,
				Location: &l,
			},
		}

		builders := make([]FunctionNode, 0)
		for _, b := range res.(*Overload).Builders {
			b.parent = out
			builders = append(builders, b)
		}
		out.Builders = builders

		return out, diags
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
	if valueCtx == nil {
		return nil, nil
	}
	fnCtx := valueCtx.Fn()
	if fnCtx == nil {
		return nil, nil
	}
	return f.matchFn(fnCtx)
}

func (f FunctionSpec) SetFileExporter(exporter func(n FunctionNode, name string) *lib.FileTreeLike) FunctionSpec {
	out := &f
	out.export = exporter
	return *out
}
func (f FunctionSpec) SetOutputFn(outputFn func(n FunctionNode) any) FunctionSpec {
	out := &f
	out.output = outputFn
	return *out
}

func (f FunctionSpec) SetKind(kind ast.SymbolKind) FunctionSpec {
	out := &f
	out.Kind = kind
	return *out
}
func (f FunctionSpec) SetHelp(help string) FunctionSpec {
	out := &f
	out.Help = help
	return *out
}

func NewFunctionSpec(name string, overloads ...OverloadSpec) FunctionSpec {
	return FunctionSpec{
		Name:      name,
		Overloads: overloads,
	}
}

type FunctionNode struct {
	ast.BaseSymbol
	Name      string
	Arguments []ast.Node
	Builders  []FunctionNode

	source   string
	spec     FunctionSpec
	parent   *FunctionNode
	overload OverloadSpec
}

func (n FunctionNode) Children() []ast.Node {
	children := make([]ast.Node, 0)
	for _, a := range n.Arguments {
		children = append(children, a)
	}
	for _, b := range n.Builders {
		children = append(children, &b)
	}

	return children
}

func (n FunctionNode) ToFileTreeLike(name string) *lib.FileTreeLike {
	if n.spec.export == nil {
		return nil
	}
	return n.spec.export(n, name)
}

func (n FunctionNode) ToSerializable() any {
	if n.spec.output == nil {
		return nil
	}
	return n.spec.output(n)
}

func (n FunctionNode) GetHelp() string {
	// TODO: Include usage / overload information
	return n.spec.Help
}

func locate(source string, idx int) ast.Location {
	out := &ast.Location{}
	for i := 0; i < idx; i++ {
		out.Index++
		switch source[i] {
		case '\n':
			out.Column = 0
			out.Line++
		default:
			out.Column++
		}
	}

	return *out
}

func (n FunctionNode) builderCompletions(includeLeadingDot bool) []protocol.CompletionItem {
	items := make([]protocol.CompletionItem, 0)
	builderSnapPosition := getBuilderInsertPosition(n)
	for _, overload := range n.spec.Overloads {
		for _, b := range overload.Builders {
			builderExists := slices.ContainsFunc(n.Builders, func(node FunctionNode) bool {
				return node.Name == b.Name
			})
			if !builderExists {
				snip := "%s(${1})"
				offset := 2
				if includeLeadingDot {
					snip = "." + snip
					offset = 1
				}
				items = append(items, protocol.CompletionItem{
					Label:            b.Name,
					Detail:           &b.Help,
					Kind:             &MethodKind,
					InsertTextFormat: &SnippetFormat,
					TextEdit: protocol.TextEdit{
						Range: protocol.Range{
							Start: builderSnapPosition.ColOffset(offset).ToLspPosition(),
							End:   builderSnapPosition.ColOffset(offset).ToLspPosition(),
						},
						NewText: fmt.Sprintf(snip, b.Name),
					},
				})
			}
		}
	}
	return lo.UniqBy(items, func(item protocol.CompletionItem) string {
		// TODO: How do we want to handle overloads here?
		return item.Label
	})

}

func (n FunctionNode) argCompletions(fileSource string, p protocol.Position, triggerChar *string, symbols map[string]*ast.Namespace) []protocol.CompletionItem {
	argIdx := -1
	cursorIdx := p.IndexIn(fileSource)

	// Find which argument the cursor is in or after
	for i, a := range n.Arguments {
		if a == nil {
			// Nil argument means this position hasn't been filled yet
			// This is likely where we want to complete
			argIdx = i
			break
		}

		al := a.GetLocation()
		if al == nil {
			continue
		}

		// Check if cursor is within this argument
		if al.ContainsPosition(p) {
			argIdx = i
			break
		}

		// Check if cursor is after this argument
		// This handles the case where cursor is between args (e.g., after comma)
		argEndIdx := al.Stop.IndexIn(fileSource)
		if cursorIdx > argEndIdx {
			// Cursor is after this argument, so we might be completing the next one
			argIdx = i + 1
		}
	}

	// If we still haven't found an index, check if we're completing the first argument
	if argIdx == -1 {
		// Check if cursor is before the first argument (if any exist)
		if len(n.Arguments) > 0 && n.Arguments[0] != nil {
			firstArgLoc := n.Arguments[0].GetLocation()
			if firstArgLoc != nil {
				firstArgIdx := firstArgLoc.Start.IndexIn(fileSource)
				if cursorIdx < firstArgIdx {
					argIdx = 0
				}
			}
		} else {
			// No arguments yet, we're completing the first one
			argIdx = 0
		}
	}

	// Clamp to valid range - if we're past all defined args, check for rest args
	if argIdx >= len(n.overload.Args) {
		if n.overload.RestArg != nil {
			// Use rest arg spec for completions
			if c, ok := n.overload.RestArg.(ast.CompletableNode); ok {
				return c.Complete(fileSource, p, triggerChar, symbols)
			}
		}
		return make([]protocol.CompletionItem, 0)
	}

	// Get completions for the identified argument
	if c, ok := n.overload.Args[argIdx].(ast.CompletableNode); ok {
		return c.Complete(fileSource, p, triggerChar, symbols)
	}

	return make([]protocol.CompletionItem, 0)
}

type functionCompletionMode string

const (
	functionCompletionModeArgs     functionCompletionMode = "ARGS"
	functionCompletionModeBuilders functionCompletionMode = "BUILDERS"
)

func (n FunctionNode) Complete(fileSource string, position protocol.Position, triggerChar *string, symbols map[string]*ast.Namespace) []protocol.CompletionItem {
	log.Printf("Completing function %s", n.Name)
	items := make([]protocol.CompletionItem, 0)
	var mode functionCompletionMode
	if triggerChar == nil || *triggerChar == "" {
		// Infer mode based on cursor position
		// If position is after the args closing paren, suggest builders
		// Otherwise, suggest args
		if n.Location != nil {
			endOfArgs := getBuilderInsertPosition(n)
			l := int(position.Line)
			if l > endOfArgs.Line ||
				(l == endOfArgs.Line && position.Character > uint32(endOfArgs.Column)) {
				mode = functionCompletionModeBuilders
			} else {
				mode = functionCompletionModeArgs
			}
		} else {
			mode = functionCompletionModeArgs
		}
	} else {
		switch *triggerChar {
		case ".", ")":
			mode = functionCompletionModeBuilders
		case "(", ",":
			mode = functionCompletionModeArgs
		default:
			mode = functionCompletionModeArgs
		}
	}
	log.Println("mode: ", mode)
	if triggerChar != nil {
		log.Printf("triggerChar: '%s'", *triggerChar)
	} else {
		log.Println("triggerChar: nil")
	}

	switch mode {
	case functionCompletionModeBuilders:
		idx := position.IndexIn(fileSource)
		items = n.builderCompletions(fileSource[idx-1] == ')')
		if len(items) == 0 && n.parent != nil {
			return n.parent.Complete(fileSource, position, triggerChar, symbols)
		}
	case functionCompletionModeArgs:
		items = n.argCompletions(fileSource, position, triggerChar, symbols)
	}

	log.Println("Completions: ", items)
	return items
}

func getBuilderInsertPosition(n FunctionNode) ast.Location {
	endOffset := 0

	for i := len(n.source) - 1; i >= 0; i-- {
		t := n.source[i]
		if t == ')' {
			endOffset = i
			break
		}
		if t == '.' {
			endOffset = i - 1
			break
		}
	}

	builderSnapInner := locate(n.source, endOffset)

	return ast.Location{
		Line:   n.Location.Start.Line + builderSnapInner.Line,
		Column: n.Location.Start.Column + builderSnapInner.Column,
		Index:  n.Location.Start.Index + builderSnapInner.Index,
	}
}
