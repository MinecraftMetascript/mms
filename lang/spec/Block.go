package spec

import (
	"github.com/minecraftmetascript/mms/lang/ast"
	"github.com/minecraftmetascript/mms/lang/grammar"
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

func NewBlockSpec(kind string, allowedValues []ValueSpec) BlockSpec {
	return BlockSpec{Kind: kind, AllowedValues: NewValueSpecList(allowedValues...)}
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
		Kind:         b.Kind,
		Declarations: make(map[string]ast.Node),
		location:     ast.RuleLocation(ctx),
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

		if val, d := b.AllowedValues.Match(valueCtx); val == nil {
			if d != nil && len(d) > 0 {
				diags = append(diags, d...)
			} else {
				diags = append(diags, ast.Diagnostic{
					Location: ast.RuleLocation(decl),
					Message:  "Invalid value",
					Severity: ast.Error,
				})
				continue
			}
		} else {
			out.Declarations[id] = val
		}

	}
	if len(diags) > 0 {
		return out, diags
	}
	return out, nil
}

// / Block Node
type BlockNode struct {
	Kind         string
	Declarations map[string]ast.Node
	location     ast.SourceLocation
}
