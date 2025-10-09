package spec

import (
	"fmt"
	"log"

	"github.com/minecraftmetascript/mms/lang/ast"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lib"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// Block Spec List
type BlockSpecList struct {
	blocks map[string]*BlockSpec
}

func (bsl *BlockSpecList) Add(bs *BlockSpec) {
	bsl.blocks[bs.Kind] = bs
}

func (bsl *BlockSpecList) Get(kind string) *BlockSpec {
	return bsl.blocks[kind]
}

func (bsl *BlockSpecList) Remove(bs *BlockSpec) {
	delete(bsl.blocks, bs.Kind)
}

func NewBlockSpecList(
	blocks ...BlockSpec,
) *BlockSpecList {
	out := &BlockSpecList{
		blocks: make(map[string]*BlockSpec),
	}
	for _, bs := range blocks {
		out.Add(&bs)
	}
	return out
}

/// Block Spec

type BlockSpec struct {
	Kind          string
	AllowedValues ValueSpecList
}

func (b BlockNode) ExtractInlineSymbols() []ast.Symbol {
	return b.inlineSymbols
}

func NewBlockSpec(kind string, allowedValues []ValueSpec) BlockSpec {
	return BlockSpec{Kind: kind, AllowedValues: *NewValueSpecList(allowedValues...)}
}

func (b BlockSpec) Match(ctx grammar.IBlockContext) (*BlockNode, []ast.Diagnostic) {
	kind := ctx.Identifier()
	if kind == nil {
		return nil, []ast.Diagnostic{
			{
				Location: ast.RuleLocation(ctx),
				Message:  "Missing Block Kind",
				Severity: ast.Error,
			},
		}
	}
	if b.Kind != kind.GetText() {
		// Ctx isn't guaranteed to be invalid, but isn't a match for this spec
		return nil, nil
	}
	out := &BlockNode{
		Kind:          b.Kind,
		Declarations:  make(map[string]ast.Node),
		location:      ast.RuleLocation(ctx),
		spec:          b,
		inlineSymbols: []ast.Symbol{},
	}
	diags := make([]ast.Diagnostic, 0)
	for _, decl := range ctx.AllVarDecl() {
		idCtx := decl.Identifier()
		if idCtx == nil {
			diags = append(diags, ast.Diagnostic{
				Location: ast.RuleLocation(decl),
				Message:  "Missing variable name",
				Severity: ast.Error,
			})
			continue
		}
		id := idCtx.GetText()
		valueCtx := decl.Value()
		if valueCtx == nil {
			diags = append(diags, ast.Diagnostic{
				Location: ast.RuleLocation(decl),
				Message:  "Missing value",
				Severity: ast.Error,
			})
			continue
		} else if val, d := b.AllowedValues.Match(valueCtx); lib.IsNilInterface(val) {
			// Nil Case
			if d != nil && len(d) > 0 {
				// Use diagnostics from the value spec matching
				diags = append(diags, d...)
			} else {
				// No specific diagnostics - provide helpful generic message
				msg := fmt.Sprintf("Invalid value for '%s' in %s block. Value does not match any allowed type.", id, b.Kind)
				diags = append(diags, ast.Diagnostic{
					Location: ast.RuleLocation(decl),
					Message:  msg,
					Severity: ast.Error,
				})
			}
			declL := ast.RuleLocation(decl)
			idL := ast.TerminalLocation(idCtx)
			out.Declarations[id] = &ast.EmptySymbol{
				BaseSymbol: ast.BaseSymbol{
					BaseNode: ast.BaseNode{
						Location: &declL,
					},
					Location:     &declL,
					NameLocation: &idL,
					Kind:         "",
				},
			}
		} else {
			// At this point, val should not equal nil
			if s, ok := val.(ast.Symbol); ok {
				docString := decl.DocString()
				if docString != nil {
					// TODO: This will probably kill newlines -- not what we want
					s.SetDocstring(docString.GetText())
				}
				idL := ast.TerminalLocation(idCtx)
				s.SetNameLocation(&idL)
			} else {
				log.Println("Invalid value for block declaration:", val)
			}
			out.Declarations[id] = val

			if e, ok := val.(ast.ExtractableNode); ok {
				out.inlineSymbols = append(out.inlineSymbols, e.ExtractInlineSymbols()...)
			}

			// Include any warnings/info diagnostics from successful match
			if d != nil && len(d) > 0 {
				diags = append(diags, d...)
			}
		}

	}
	if len(diags) > 0 {
		return out, diags
	}
	return out, nil
}

// / Block Node
type BlockNode struct {
	Kind          string
	Declarations  map[string]ast.Node
	location      ast.SourceLocation
	spec          BlockSpec
	inlineSymbols []ast.Symbol
}

func (b BlockNode) Complete(fileSource string, position protocol.Position, triggerChar *string, symbols map[string]*ast.Namespace) []protocol.CompletionItem {
	for _, d := range b.Declarations {
		if d.GetLocation() == nil {
			// No location don't care
			return nil
		}
		if d.GetLocation().ContainsPosition(position) {
			if _, ok := d.(*ast.EmptySymbol); ok {
				break
			}
			return nil
		}
		if s, ok := d.(ast.Symbol); ok {
			if s.GetNameLocation().ContainsPosition(position) {
				return nil
			}
		}
	}

	out := make([]protocol.CompletionItem, 0)
	for _, v := range b.spec.AllowedValues.specs {
		if _, ok := v.(*ListSpec); ok {
			filterTxt := "Sequence"
			out = append(out, protocol.CompletionItem{
				Label:            "Sequence",
				FilterText:       &filterTxt,
				SortText:         &filterTxt,
				Kind:             &StructKind,
				InsertTextFormat: &SnippetFormat,
				TextEdit: protocol.TextEdit{
					Range: protocol.Range{
						Start: position,
						End:   position,
					},
					NewText: "[${1}]",
				},
			})
		} else if c, ok := v.(ast.CompletableNode); ok {
			out = append(out, c.Complete(fileSource, position, triggerChar, symbols)...)
		}
	}

	// Filter by prefix if user has typed something
	return out
}

func (b BlockNode) Children() []ast.Node {
	out := make([]ast.Node, 0)
	for _, node := range b.Declarations {
		out = append(out, node)
	}
	return out
}

func (b BlockNode) GetLocation() *ast.SourceLocation {
	return &b.location
}
