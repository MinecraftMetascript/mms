package lsp

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lang/grammar"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type FoldingRangeVisitor struct {
	grammar.BaseMinecraftMetascriptListener

	foldingRanges []protocol.FoldingRange
}

func (frv *FoldingRangeVisitor) mkFoldingRange(ctx antlr.ParserRuleContext) {
	startLine := protocol.UInteger(ctx.GetStart().GetLine())
	startCol := protocol.UInteger(ctx.GetStart().GetColumn())
	endLine := protocol.UInteger(ctx.GetStop().GetLine())
	endCol := protocol.UInteger(ctx.GetStop().GetColumn())
	kind := "namespace"
	fmt.Printf("Making folding range %d:%d-%d:%d\n", startLine, startCol, endLine, endCol)
	frv.foldingRanges = append(frv.foldingRanges, protocol.FoldingRange{
		StartLine:      startLine,
		StartCharacter: &startCol,
		EndLine:        endLine,
		EndCharacter:   &endCol,
		Kind:           &kind,
	})
}

func (frv *FoldingRangeVisitor) EnterNamespaceDeclaration(ctx *grammar.NamespaceDeclarationContext) {
	frv.mkFoldingRange(ctx)
}
func (frv *FoldingRangeVisitor) EnterDensityFnBlock(ctx *grammar.DensityFnBlockContext) {
	frv.mkFoldingRange(ctx)
}
func (frv *FoldingRangeVisitor) EnterSurfaceBlock(ctx *grammar.SurfaceBlockContext) {
	frv.mkFoldingRange(ctx)
}
func (frv *FoldingRangeVisitor) EnterNoiseBlock(ctx *grammar.NoiseBlockContext) {
	frv.mkFoldingRange(ctx)
}
