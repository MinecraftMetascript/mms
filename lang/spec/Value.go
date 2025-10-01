package spec

import (
	"github.com/minecraftmetascript/mms/lang/ast"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lib"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

var SnippetFormat = protocol.InsertTextFormatSnippet
var SnippetKind = protocol.CompletionItemKindSnippet
var MethodKind = protocol.CompletionItemKindMethod
var ReferenceKind = protocol.CompletionItemKindReference

type ValueSpec interface {
	Match(valueCtx grammar.IValueContext) (ast.Node, []ast.Diagnostic)
}

type ValueSpecList struct {
	specs []ValueSpec
}

func (bsl *ValueSpecList) Add(bs ValueSpec) {
	bsl.specs = append(bsl.specs, bs)
}

func (bsl *ValueSpecList) Match(ctx grammar.IValueContext) (ast.Node, []ast.Diagnostic) {
	if ctx == nil {
		return nil, nil
	}
	for _, bs := range bsl.specs {
		val, diags := bs.Match(ctx)

		if !lib.IsNilInterface(val) || (diags != nil && len(diags) > 0) {
			return val, diags
		}
	}
	return nil, nil
}

func NewValueSpecList(
	specs ...ValueSpec,
) ValueSpecList {
	out := ValueSpecList{
		specs: specs,
	}
	return out
}
