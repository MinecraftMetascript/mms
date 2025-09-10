// Code generated from ./grammar/MinecraftMetascript.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // MinecraftMetascript
import "github.com/antlr4-go/antlr/v4"

// BaseMinecraftMetascriptListener is a complete listener for a parse tree produced by MinecraftMetascriptParser.
type BaseMinecraftMetascriptListener struct{}

var _ MinecraftMetascriptListener = &BaseMinecraftMetascriptListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMinecraftMetascriptListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMinecraftMetascriptListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMinecraftMetascriptListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMinecraftMetascriptListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterScript is called when production script is entered.
func (s *BaseMinecraftMetascriptListener) EnterScript(ctx *ScriptContext) {}

// ExitScript is called when production script is exited.
func (s *BaseMinecraftMetascriptListener) ExitScript(ctx *ScriptContext) {}

// EnterNamespaceDeclaration is called when production namespaceDeclaration is entered.
func (s *BaseMinecraftMetascriptListener) EnterNamespaceDeclaration(ctx *NamespaceDeclarationContext) {
}

// ExitNamespaceDeclaration is called when production namespaceDeclaration is exited.
func (s *BaseMinecraftMetascriptListener) ExitNamespaceDeclaration(ctx *NamespaceDeclarationContext) {
}

// EnterNamespace is called when production namespace is entered.
func (s *BaseMinecraftMetascriptListener) EnterNamespace(ctx *NamespaceContext) {}

// ExitNamespace is called when production namespace is exited.
func (s *BaseMinecraftMetascriptListener) ExitNamespace(ctx *NamespaceContext) {}

// EnterContentBlocks is called when production contentBlocks is entered.
func (s *BaseMinecraftMetascriptListener) EnterContentBlocks(ctx *ContentBlocksContext) {}

// ExitContentBlocks is called when production contentBlocks is exited.
func (s *BaseMinecraftMetascriptListener) ExitContentBlocks(ctx *ContentBlocksContext) {}

// EnterSurfaceBlock is called when production surfaceBlock is entered.
func (s *BaseMinecraftMetascriptListener) EnterSurfaceBlock(ctx *SurfaceBlockContext) {}

// ExitSurfaceBlock is called when production surfaceBlock is exited.
func (s *BaseMinecraftMetascriptListener) ExitSurfaceBlock(ctx *SurfaceBlockContext) {}

// EnterSurfaceStatement is called when production surfaceStatement is entered.
func (s *BaseMinecraftMetascriptListener) EnterSurfaceStatement(ctx *SurfaceStatementContext) {}

// ExitSurfaceStatement is called when production surfaceStatement is exited.
func (s *BaseMinecraftMetascriptListener) ExitSurfaceStatement(ctx *SurfaceStatementContext) {}

// EnterVerticalAnchor is called when production verticalAnchor is entered.
func (s *BaseMinecraftMetascriptListener) EnterVerticalAnchor(ctx *VerticalAnchorContext) {}

// ExitVerticalAnchor is called when production verticalAnchor is exited.
func (s *BaseMinecraftMetascriptListener) ExitVerticalAnchor(ctx *VerticalAnchorContext) {}

// EnterVerticalAnchorDeclaration is called when production verticalAnchorDeclaration is entered.
func (s *BaseMinecraftMetascriptListener) EnterVerticalAnchorDeclaration(ctx *VerticalAnchorDeclarationContext) {
}

// ExitVerticalAnchorDeclaration is called when production verticalAnchorDeclaration is exited.
func (s *BaseMinecraftMetascriptListener) ExitVerticalAnchorDeclaration(ctx *VerticalAnchorDeclarationContext) {
}

// EnterSurfaceCondition is called when production surfaceCondition is entered.
func (s *BaseMinecraftMetascriptListener) EnterSurfaceCondition(ctx *SurfaceConditionContext) {}

// ExitSurfaceCondition is called when production surfaceCondition is exited.
func (s *BaseMinecraftMetascriptListener) ExitSurfaceCondition(ctx *SurfaceConditionContext) {}

// EnterSurfaceConditionDeclaration is called when production surfaceConditionDeclaration is entered.
func (s *BaseMinecraftMetascriptListener) EnterSurfaceConditionDeclaration(ctx *SurfaceConditionDeclarationContext) {
}

// ExitSurfaceConditionDeclaration is called when production surfaceConditionDeclaration is exited.
func (s *BaseMinecraftMetascriptListener) ExitSurfaceConditionDeclaration(ctx *SurfaceConditionDeclarationContext) {
}

// EnterSurfaceCondition_Not is called when production surfaceCondition_Not is entered.
func (s *BaseMinecraftMetascriptListener) EnterSurfaceCondition_Not(ctx *SurfaceCondition_NotContext) {
}

// ExitSurfaceCondition_Not is called when production surfaceCondition_Not is exited.
func (s *BaseMinecraftMetascriptListener) ExitSurfaceCondition_Not(ctx *SurfaceCondition_NotContext) {
}

// EnterSurfaceCondition_And is called when production surfaceCondition_And is entered.
func (s *BaseMinecraftMetascriptListener) EnterSurfaceCondition_And(ctx *SurfaceCondition_AndContext) {
}

// ExitSurfaceCondition_And is called when production surfaceCondition_And is exited.
func (s *BaseMinecraftMetascriptListener) ExitSurfaceCondition_And(ctx *SurfaceCondition_AndContext) {
}

// EnterSurfaceCondition_Or is called when production surfaceCondition_Or is entered.
func (s *BaseMinecraftMetascriptListener) EnterSurfaceCondition_Or(ctx *SurfaceCondition_OrContext) {}

// ExitSurfaceCondition_Or is called when production surfaceCondition_Or is exited.
func (s *BaseMinecraftMetascriptListener) ExitSurfaceCondition_Or(ctx *SurfaceCondition_OrContext) {}

// EnterSurfaceCondition_Reference is called when production surfaceCondition_Reference is entered.
func (s *BaseMinecraftMetascriptListener) EnterSurfaceCondition_Reference(ctx *SurfaceCondition_ReferenceContext) {
}

// ExitSurfaceCondition_Reference is called when production surfaceCondition_Reference is exited.
func (s *BaseMinecraftMetascriptListener) ExitSurfaceCondition_Reference(ctx *SurfaceCondition_ReferenceContext) {
}

// EnterSurfaceCondition_AboveSurface is called when production surfaceCondition_AboveSurface is entered.
func (s *BaseMinecraftMetascriptListener) EnterSurfaceCondition_AboveSurface(ctx *SurfaceCondition_AboveSurfaceContext) {
}

// ExitSurfaceCondition_AboveSurface is called when production surfaceCondition_AboveSurface is exited.
func (s *BaseMinecraftMetascriptListener) ExitSurfaceCondition_AboveSurface(ctx *SurfaceCondition_AboveSurfaceContext) {
}

// EnterSurfaceCondition_Biome is called when production surfaceCondition_Biome is entered.
func (s *BaseMinecraftMetascriptListener) EnterSurfaceCondition_Biome(ctx *SurfaceCondition_BiomeContext) {
}

// ExitSurfaceCondition_Biome is called when production surfaceCondition_Biome is exited.
func (s *BaseMinecraftMetascriptListener) ExitSurfaceCondition_Biome(ctx *SurfaceCondition_BiomeContext) {
}

// EnterSurfaceCondition_Hole is called when production surfaceCondition_Hole is entered.
func (s *BaseMinecraftMetascriptListener) EnterSurfaceCondition_Hole(ctx *SurfaceCondition_HoleContext) {
}

// ExitSurfaceCondition_Hole is called when production surfaceCondition_Hole is exited.
func (s *BaseMinecraftMetascriptListener) ExitSurfaceCondition_Hole(ctx *SurfaceCondition_HoleContext) {
}

// EnterSurfaceCondition_Steep is called when production surfaceCondition_Steep is entered.
func (s *BaseMinecraftMetascriptListener) EnterSurfaceCondition_Steep(ctx *SurfaceCondition_SteepContext) {
}

// ExitSurfaceCondition_Steep is called when production surfaceCondition_Steep is exited.
func (s *BaseMinecraftMetascriptListener) ExitSurfaceCondition_Steep(ctx *SurfaceCondition_SteepContext) {
}

// EnterSurfaceCondition_Freezing is called when production surfaceCondition_Freezing is entered.
func (s *BaseMinecraftMetascriptListener) EnterSurfaceCondition_Freezing(ctx *SurfaceCondition_FreezingContext) {
}

// ExitSurfaceCondition_Freezing is called when production surfaceCondition_Freezing is exited.
func (s *BaseMinecraftMetascriptListener) ExitSurfaceCondition_Freezing(ctx *SurfaceCondition_FreezingContext) {
}

// EnterSurfaceCondition_NoiseThresholdBuilder is called when production surfaceCondition_NoiseThresholdBuilder is entered.
func (s *BaseMinecraftMetascriptListener) EnterSurfaceCondition_NoiseThresholdBuilder(ctx *SurfaceCondition_NoiseThresholdBuilderContext) {
}

// ExitSurfaceCondition_NoiseThresholdBuilder is called when production surfaceCondition_NoiseThresholdBuilder is exited.
func (s *BaseMinecraftMetascriptListener) ExitSurfaceCondition_NoiseThresholdBuilder(ctx *SurfaceCondition_NoiseThresholdBuilderContext) {
}

// EnterSurfaceCondition_NoiseThreshold is called when production surfaceCondition_NoiseThreshold is entered.
func (s *BaseMinecraftMetascriptListener) EnterSurfaceCondition_NoiseThreshold(ctx *SurfaceCondition_NoiseThresholdContext) {
}

// ExitSurfaceCondition_NoiseThreshold is called when production surfaceCondition_NoiseThreshold is exited.
func (s *BaseMinecraftMetascriptListener) ExitSurfaceCondition_NoiseThreshold(ctx *SurfaceCondition_NoiseThresholdContext) {
}

// EnterSurfaceCondition_StoneDepth is called when production surfaceCondition_StoneDepth is entered.
func (s *BaseMinecraftMetascriptListener) EnterSurfaceCondition_StoneDepth(ctx *SurfaceCondition_StoneDepthContext) {
}

// ExitSurfaceCondition_StoneDepth is called when production surfaceCondition_StoneDepth is exited.
func (s *BaseMinecraftMetascriptListener) ExitSurfaceCondition_StoneDepth(ctx *SurfaceCondition_StoneDepthContext) {
}

// EnterSurfaceCondition_StoneDepthBuilder is called when production surfaceCondition_StoneDepthBuilder is entered.
func (s *BaseMinecraftMetascriptListener) EnterSurfaceCondition_StoneDepthBuilder(ctx *SurfaceCondition_StoneDepthBuilderContext) {
}

// ExitSurfaceCondition_StoneDepthBuilder is called when production surfaceCondition_StoneDepthBuilder is exited.
func (s *BaseMinecraftMetascriptListener) ExitSurfaceCondition_StoneDepthBuilder(ctx *SurfaceCondition_StoneDepthBuilderContext) {
}

// EnterSurfaceCondition_StoneDepthBuilder_SecondaryDepthRange is called when production surfaceCondition_StoneDepthBuilder_SecondaryDepthRange is entered.
func (s *BaseMinecraftMetascriptListener) EnterSurfaceCondition_StoneDepthBuilder_SecondaryDepthRange(ctx *SurfaceCondition_StoneDepthBuilder_SecondaryDepthRangeContext) {
}

// ExitSurfaceCondition_StoneDepthBuilder_SecondaryDepthRange is called when production surfaceCondition_StoneDepthBuilder_SecondaryDepthRange is exited.
func (s *BaseMinecraftMetascriptListener) ExitSurfaceCondition_StoneDepthBuilder_SecondaryDepthRange(ctx *SurfaceCondition_StoneDepthBuilder_SecondaryDepthRangeContext) {
}

// EnterSurfaceCondition_VerticalGradient is called when production surfaceCondition_VerticalGradient is entered.
func (s *BaseMinecraftMetascriptListener) EnterSurfaceCondition_VerticalGradient(ctx *SurfaceCondition_VerticalGradientContext) {
}

// ExitSurfaceCondition_VerticalGradient is called when production surfaceCondition_VerticalGradient is exited.
func (s *BaseMinecraftMetascriptListener) ExitSurfaceCondition_VerticalGradient(ctx *SurfaceCondition_VerticalGradientContext) {
}

// EnterSurfaceCondition_VerticalGradientBuilder is called when production surfaceCondition_VerticalGradientBuilder is entered.
func (s *BaseMinecraftMetascriptListener) EnterSurfaceCondition_VerticalGradientBuilder(ctx *SurfaceCondition_VerticalGradientBuilderContext) {
}

// ExitSurfaceCondition_VerticalGradientBuilder is called when production surfaceCondition_VerticalGradientBuilder is exited.
func (s *BaseMinecraftMetascriptListener) ExitSurfaceCondition_VerticalGradientBuilder(ctx *SurfaceCondition_VerticalGradientBuilderContext) {
}

// EnterSurfaceCondition_AboveWater is called when production surfaceCondition_AboveWater is entered.
func (s *BaseMinecraftMetascriptListener) EnterSurfaceCondition_AboveWater(ctx *SurfaceCondition_AboveWaterContext) {
}

// ExitSurfaceCondition_AboveWater is called when production surfaceCondition_AboveWater is exited.
func (s *BaseMinecraftMetascriptListener) ExitSurfaceCondition_AboveWater(ctx *SurfaceCondition_AboveWaterContext) {
}

// EnterSurfaceCondition_AboveWaterBuilder is called when production surfaceCondition_AboveWaterBuilder is entered.
func (s *BaseMinecraftMetascriptListener) EnterSurfaceCondition_AboveWaterBuilder(ctx *SurfaceCondition_AboveWaterBuilderContext) {
}

// ExitSurfaceCondition_AboveWaterBuilder is called when production surfaceCondition_AboveWaterBuilder is exited.
func (s *BaseMinecraftMetascriptListener) ExitSurfaceCondition_AboveWaterBuilder(ctx *SurfaceCondition_AboveWaterBuilderContext) {
}

// EnterSurfaceCondition_YAbove is called when production surfaceCondition_YAbove is entered.
func (s *BaseMinecraftMetascriptListener) EnterSurfaceCondition_YAbove(ctx *SurfaceCondition_YAboveContext) {
}

// ExitSurfaceCondition_YAbove is called when production surfaceCondition_YAbove is exited.
func (s *BaseMinecraftMetascriptListener) ExitSurfaceCondition_YAbove(ctx *SurfaceCondition_YAboveContext) {
}

// EnterSurfaceCondition_YAboveBuilder is called when production surfaceCondition_YAboveBuilder is entered.
func (s *BaseMinecraftMetascriptListener) EnterSurfaceCondition_YAboveBuilder(ctx *SurfaceCondition_YAboveBuilderContext) {
}

// ExitSurfaceCondition_YAboveBuilder is called when production surfaceCondition_YAboveBuilder is exited.
func (s *BaseMinecraftMetascriptListener) ExitSurfaceCondition_YAboveBuilder(ctx *SurfaceCondition_YAboveBuilderContext) {
}

// EnterSurfaceRuleDeclaration is called when production surfaceRuleDeclaration is entered.
func (s *BaseMinecraftMetascriptListener) EnterSurfaceRuleDeclaration(ctx *SurfaceRuleDeclarationContext) {
}

// ExitSurfaceRuleDeclaration is called when production surfaceRuleDeclaration is exited.
func (s *BaseMinecraftMetascriptListener) ExitSurfaceRuleDeclaration(ctx *SurfaceRuleDeclarationContext) {
}

// EnterSurfaceRule is called when production surfaceRule is entered.
func (s *BaseMinecraftMetascriptListener) EnterSurfaceRule(ctx *SurfaceRuleContext) {}

// ExitSurfaceRule is called when production surfaceRule is exited.
func (s *BaseMinecraftMetascriptListener) ExitSurfaceRule(ctx *SurfaceRuleContext) {}

// EnterSurfaceRule_Reference is called when production surfaceRule_Reference is entered.
func (s *BaseMinecraftMetascriptListener) EnterSurfaceRule_Reference(ctx *SurfaceRule_ReferenceContext) {
}

// ExitSurfaceRule_Reference is called when production surfaceRule_Reference is exited.
func (s *BaseMinecraftMetascriptListener) ExitSurfaceRule_Reference(ctx *SurfaceRule_ReferenceContext) {
}

// EnterSurfaceRule_Block is called when production surfaceRule_Block is entered.
func (s *BaseMinecraftMetascriptListener) EnterSurfaceRule_Block(ctx *SurfaceRule_BlockContext) {}

// ExitSurfaceRule_Block is called when production surfaceRule_Block is exited.
func (s *BaseMinecraftMetascriptListener) ExitSurfaceRule_Block(ctx *SurfaceRule_BlockContext) {}

// EnterSurfaceRule_Sequence is called when production surfaceRule_Sequence is entered.
func (s *BaseMinecraftMetascriptListener) EnterSurfaceRule_Sequence(ctx *SurfaceRule_SequenceContext) {
}

// ExitSurfaceRule_Sequence is called when production surfaceRule_Sequence is exited.
func (s *BaseMinecraftMetascriptListener) ExitSurfaceRule_Sequence(ctx *SurfaceRule_SequenceContext) {
}

// EnterSurfaceRule_Bandlands is called when production surfaceRule_Bandlands is entered.
func (s *BaseMinecraftMetascriptListener) EnterSurfaceRule_Bandlands(ctx *SurfaceRule_BandlandsContext) {
}

// ExitSurfaceRule_Bandlands is called when production surfaceRule_Bandlands is exited.
func (s *BaseMinecraftMetascriptListener) ExitSurfaceRule_Bandlands(ctx *SurfaceRule_BandlandsContext) {
}

// EnterSurfaceRule_If is called when production surfaceRule_If is entered.
func (s *BaseMinecraftMetascriptListener) EnterSurfaceRule_If(ctx *SurfaceRule_IfContext) {}

// ExitSurfaceRule_If is called when production surfaceRule_If is exited.
func (s *BaseMinecraftMetascriptListener) ExitSurfaceRule_If(ctx *SurfaceRule_IfContext) {}

// EnterDeclare is called when production declare is entered.
func (s *BaseMinecraftMetascriptListener) EnterDeclare(ctx *DeclareContext) {}

// ExitDeclare is called when production declare is exited.
func (s *BaseMinecraftMetascriptListener) ExitDeclare(ctx *DeclareContext) {}

// EnterBuilder_XZScale is called when production builder_XZScale is entered.
func (s *BaseMinecraftMetascriptListener) EnterBuilder_XZScale(ctx *Builder_XZScaleContext) {}

// ExitBuilder_XZScale is called when production builder_XZScale is exited.
func (s *BaseMinecraftMetascriptListener) ExitBuilder_XZScale(ctx *Builder_XZScaleContext) {}

// EnterBuilder_YScale is called when production builder_YScale is entered.
func (s *BaseMinecraftMetascriptListener) EnterBuilder_YScale(ctx *Builder_YScaleContext) {}

// ExitBuilder_YScale is called when production builder_YScale is exited.
func (s *BaseMinecraftMetascriptListener) ExitBuilder_YScale(ctx *Builder_YScaleContext) {}

// EnterBuilder_XZFactor is called when production builder_XZFactor is entered.
func (s *BaseMinecraftMetascriptListener) EnterBuilder_XZFactor(ctx *Builder_XZFactorContext) {}

// ExitBuilder_XZFactor is called when production builder_XZFactor is exited.
func (s *BaseMinecraftMetascriptListener) ExitBuilder_XZFactor(ctx *Builder_XZFactorContext) {}

// EnterBuilder_YFactor is called when production builder_YFactor is entered.
func (s *BaseMinecraftMetascriptListener) EnterBuilder_YFactor(ctx *Builder_YFactorContext) {}

// ExitBuilder_YFactor is called when production builder_YFactor is exited.
func (s *BaseMinecraftMetascriptListener) ExitBuilder_YFactor(ctx *Builder_YFactorContext) {}

// EnterBuilder_Noise is called when production builder_Noise is entered.
func (s *BaseMinecraftMetascriptListener) EnterBuilder_Noise(ctx *Builder_NoiseContext) {}

// ExitBuilder_Noise is called when production builder_Noise is exited.
func (s *BaseMinecraftMetascriptListener) ExitBuilder_Noise(ctx *Builder_NoiseContext) {}

// EnterBuilder_Smear is called when production builder_Smear is entered.
func (s *BaseMinecraftMetascriptListener) EnterBuilder_Smear(ctx *Builder_SmearContext) {}

// ExitBuilder_Smear is called when production builder_Smear is exited.
func (s *BaseMinecraftMetascriptListener) ExitBuilder_Smear(ctx *Builder_SmearContext) {}

// EnterBuilder_Type1 is called when production builder_Type1 is entered.
func (s *BaseMinecraftMetascriptListener) EnterBuilder_Type1(ctx *Builder_Type1Context) {}

// ExitBuilder_Type1 is called when production builder_Type1 is exited.
func (s *BaseMinecraftMetascriptListener) ExitBuilder_Type1(ctx *Builder_Type1Context) {}

// EnterBuilder_Type2 is called when production builder_Type2 is entered.
func (s *BaseMinecraftMetascriptListener) EnterBuilder_Type2(ctx *Builder_Type2Context) {}

// ExitBuilder_Type2 is called when production builder_Type2 is exited.
func (s *BaseMinecraftMetascriptListener) ExitBuilder_Type2(ctx *Builder_Type2Context) {}

// EnterBuilder_ShiftX is called when production builder_ShiftX is entered.
func (s *BaseMinecraftMetascriptListener) EnterBuilder_ShiftX(ctx *Builder_ShiftXContext) {}

// ExitBuilder_ShiftX is called when production builder_ShiftX is exited.
func (s *BaseMinecraftMetascriptListener) ExitBuilder_ShiftX(ctx *Builder_ShiftXContext) {}

// EnterBuilder_ShiftY is called when production builder_ShiftY is entered.
func (s *BaseMinecraftMetascriptListener) EnterBuilder_ShiftY(ctx *Builder_ShiftYContext) {}

// ExitBuilder_ShiftY is called when production builder_ShiftY is exited.
func (s *BaseMinecraftMetascriptListener) ExitBuilder_ShiftY(ctx *Builder_ShiftYContext) {}

// EnterBuilder_ShiftZ is called when production builder_ShiftZ is entered.
func (s *BaseMinecraftMetascriptListener) EnterBuilder_ShiftZ(ctx *Builder_ShiftZContext) {}

// ExitBuilder_ShiftZ is called when production builder_ShiftZ is exited.
func (s *BaseMinecraftMetascriptListener) ExitBuilder_ShiftZ(ctx *Builder_ShiftZContext) {}

// EnterBuilder_Amplitudes is called when production builder_Amplitudes is entered.
func (s *BaseMinecraftMetascriptListener) EnterBuilder_Amplitudes(ctx *Builder_AmplitudesContext) {}

// ExitBuilder_Amplitudes is called when production builder_Amplitudes is exited.
func (s *BaseMinecraftMetascriptListener) ExitBuilder_Amplitudes(ctx *Builder_AmplitudesContext) {}

// EnterBuilder_Offset is called when production builder_Offset is entered.
func (s *BaseMinecraftMetascriptListener) EnterBuilder_Offset(ctx *Builder_OffsetContext) {}

// ExitBuilder_Offset is called when production builder_Offset is exited.
func (s *BaseMinecraftMetascriptListener) ExitBuilder_Offset(ctx *Builder_OffsetContext) {}

// EnterBuilder_Add is called when production builder_Add is entered.
func (s *BaseMinecraftMetascriptListener) EnterBuilder_Add(ctx *Builder_AddContext) {}

// ExitBuilder_Add is called when production builder_Add is exited.
func (s *BaseMinecraftMetascriptListener) ExitBuilder_Add(ctx *Builder_AddContext) {}

// EnterBuilder_Mul is called when production builder_Mul is entered.
func (s *BaseMinecraftMetascriptListener) EnterBuilder_Mul(ctx *Builder_MulContext) {}

// ExitBuilder_Mul is called when production builder_Mul is exited.
func (s *BaseMinecraftMetascriptListener) ExitBuilder_Mul(ctx *Builder_MulContext) {}

// EnterBuilder_MulInt is called when production builder_MulInt is entered.
func (s *BaseMinecraftMetascriptListener) EnterBuilder_MulInt(ctx *Builder_MulIntContext) {}

// ExitBuilder_MulInt is called when production builder_MulInt is exited.
func (s *BaseMinecraftMetascriptListener) ExitBuilder_MulInt(ctx *Builder_MulIntContext) {}

// EnterBuilder_Min is called when production builder_Min is entered.
func (s *BaseMinecraftMetascriptListener) EnterBuilder_Min(ctx *Builder_MinContext) {}

// ExitBuilder_Min is called when production builder_Min is exited.
func (s *BaseMinecraftMetascriptListener) ExitBuilder_Min(ctx *Builder_MinContext) {}

// EnterBuilder_Max is called when production builder_Max is entered.
func (s *BaseMinecraftMetascriptListener) EnterBuilder_Max(ctx *Builder_MaxContext) {}

// ExitBuilder_Max is called when production builder_Max is exited.
func (s *BaseMinecraftMetascriptListener) ExitBuilder_Max(ctx *Builder_MaxContext) {}

// EnterBuilder_Top is called when production builder_Top is entered.
func (s *BaseMinecraftMetascriptListener) EnterBuilder_Top(ctx *Builder_TopContext) {}

// ExitBuilder_Top is called when production builder_Top is exited.
func (s *BaseMinecraftMetascriptListener) ExitBuilder_Top(ctx *Builder_TopContext) {}

// EnterBuilder_TopLiteral is called when production builder_TopLiteral is entered.
func (s *BaseMinecraftMetascriptListener) EnterBuilder_TopLiteral(ctx *Builder_TopLiteralContext) {}

// ExitBuilder_TopLiteral is called when production builder_TopLiteral is exited.
func (s *BaseMinecraftMetascriptListener) ExitBuilder_TopLiteral(ctx *Builder_TopLiteralContext) {}

// EnterBuilder_Bottom is called when production builder_Bottom is entered.
func (s *BaseMinecraftMetascriptListener) EnterBuilder_Bottom(ctx *Builder_BottomContext) {}

// ExitBuilder_Bottom is called when production builder_Bottom is exited.
func (s *BaseMinecraftMetascriptListener) ExitBuilder_Bottom(ctx *Builder_BottomContext) {}

// EnterBuilder_BottomLiteral is called when production builder_BottomLiteral is entered.
func (s *BaseMinecraftMetascriptListener) EnterBuilder_BottomLiteral(ctx *Builder_BottomLiteralContext) {
}

// ExitBuilder_BottomLiteral is called when production builder_BottomLiteral is exited.
func (s *BaseMinecraftMetascriptListener) ExitBuilder_BottomLiteral(ctx *Builder_BottomLiteralContext) {
}

// EnterBuilder_InRange is called when production builder_InRange is entered.
func (s *BaseMinecraftMetascriptListener) EnterBuilder_InRange(ctx *Builder_InRangeContext) {}

// ExitBuilder_InRange is called when production builder_InRange is exited.
func (s *BaseMinecraftMetascriptListener) ExitBuilder_InRange(ctx *Builder_InRangeContext) {}

// EnterBuilder_OutRange is called when production builder_OutRange is entered.
func (s *BaseMinecraftMetascriptListener) EnterBuilder_OutRange(ctx *Builder_OutRangeContext) {}

// ExitBuilder_OutRange is called when production builder_OutRange is exited.
func (s *BaseMinecraftMetascriptListener) ExitBuilder_OutRange(ctx *Builder_OutRangeContext) {}

// EnterResourceReference is called when production resourceReference is entered.
func (s *BaseMinecraftMetascriptListener) EnterResourceReference(ctx *ResourceReferenceContext) {}

// ExitResourceReference is called when production resourceReference is exited.
func (s *BaseMinecraftMetascriptListener) ExitResourceReference(ctx *ResourceReferenceContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseMinecraftMetascriptListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseMinecraftMetascriptListener) ExitNumber(ctx *NumberContext) {}

// EnterNoiseBlock is called when production noiseBlock is entered.
func (s *BaseMinecraftMetascriptListener) EnterNoiseBlock(ctx *NoiseBlockContext) {}

// ExitNoiseBlock is called when production noiseBlock is exited.
func (s *BaseMinecraftMetascriptListener) ExitNoiseBlock(ctx *NoiseBlockContext) {}

// EnterNoiseDeclaration is called when production noiseDeclaration is entered.
func (s *BaseMinecraftMetascriptListener) EnterNoiseDeclaration(ctx *NoiseDeclarationContext) {}

// ExitNoiseDeclaration is called when production noiseDeclaration is exited.
func (s *BaseMinecraftMetascriptListener) ExitNoiseDeclaration(ctx *NoiseDeclarationContext) {}

// EnterNoise is called when production noise is entered.
func (s *BaseMinecraftMetascriptListener) EnterNoise(ctx *NoiseContext) {}

// ExitNoise is called when production noise is exited.
func (s *BaseMinecraftMetascriptListener) ExitNoise(ctx *NoiseContext) {}

// EnterNoiseDefinition is called when production noiseDefinition is entered.
func (s *BaseMinecraftMetascriptListener) EnterNoiseDefinition(ctx *NoiseDefinitionContext) {}

// ExitNoiseDefinition is called when production noiseDefinition is exited.
func (s *BaseMinecraftMetascriptListener) ExitNoiseDefinition(ctx *NoiseDefinitionContext) {}

// EnterNoise_Builder is called when production noise_Builder is entered.
func (s *BaseMinecraftMetascriptListener) EnterNoise_Builder(ctx *Noise_BuilderContext) {}

// ExitNoise_Builder is called when production noise_Builder is exited.
func (s *BaseMinecraftMetascriptListener) ExitNoise_Builder(ctx *Noise_BuilderContext) {}

// EnterDensityFnBlock is called when production densityFnBlock is entered.
func (s *BaseMinecraftMetascriptListener) EnterDensityFnBlock(ctx *DensityFnBlockContext) {}

// ExitDensityFnBlock is called when production densityFnBlock is exited.
func (s *BaseMinecraftMetascriptListener) ExitDensityFnBlock(ctx *DensityFnBlockContext) {}

// EnterDensityFnDeclaration is called when production densityFnDeclaration is entered.
func (s *BaseMinecraftMetascriptListener) EnterDensityFnDeclaration(ctx *DensityFnDeclarationContext) {
}

// ExitDensityFnDeclaration is called when production densityFnDeclaration is exited.
func (s *BaseMinecraftMetascriptListener) ExitDensityFnDeclaration(ctx *DensityFnDeclarationContext) {
}

// EnterDensityFn is called when production densityFn is entered.
func (s *BaseMinecraftMetascriptListener) EnterDensityFn(ctx *DensityFnContext) {}

// ExitDensityFn is called when production densityFn is exited.
func (s *BaseMinecraftMetascriptListener) ExitDensityFn(ctx *DensityFnContext) {}

// EnterDensityFn_NoInput is called when production densityFn_NoInput is entered.
func (s *BaseMinecraftMetascriptListener) EnterDensityFn_NoInput(ctx *DensityFn_NoInputContext) {}

// ExitDensityFn_NoInput is called when production densityFn_NoInput is exited.
func (s *BaseMinecraftMetascriptListener) ExitDensityFn_NoInput(ctx *DensityFn_NoInputContext) {}

// EnterDensityFn_SingleInput is called when production densityFn_SingleInput is entered.
func (s *BaseMinecraftMetascriptListener) EnterDensityFn_SingleInput(ctx *DensityFn_SingleInputContext) {
}

// ExitDensityFn_SingleInput is called when production densityFn_SingleInput is exited.
func (s *BaseMinecraftMetascriptListener) ExitDensityFn_SingleInput(ctx *DensityFn_SingleInputContext) {
}

// EnterDensityFn_InlineNoise is called when production densityFn_InlineNoise is entered.
func (s *BaseMinecraftMetascriptListener) EnterDensityFn_InlineNoise(ctx *DensityFn_InlineNoiseContext) {
}

// ExitDensityFn_InlineNoise is called when production densityFn_InlineNoise is exited.
func (s *BaseMinecraftMetascriptListener) ExitDensityFn_InlineNoise(ctx *DensityFn_InlineNoiseContext) {
}

// EnterDensityFn_Noise is called when production densityFn_Noise is entered.
func (s *BaseMinecraftMetascriptListener) EnterDensityFn_Noise(ctx *DensityFn_NoiseContext) {}

// ExitDensityFn_Noise is called when production densityFn_Noise is exited.
func (s *BaseMinecraftMetascriptListener) ExitDensityFn_Noise(ctx *DensityFn_NoiseContext) {}

// EnterDensityFn_NoiseBuilder is called when production densityFn_NoiseBuilder is entered.
func (s *BaseMinecraftMetascriptListener) EnterDensityFn_NoiseBuilder(ctx *DensityFn_NoiseBuilderContext) {
}

// ExitDensityFn_NoiseBuilder is called when production densityFn_NoiseBuilder is exited.
func (s *BaseMinecraftMetascriptListener) ExitDensityFn_NoiseBuilder(ctx *DensityFn_NoiseBuilderContext) {
}

// EnterDensityFn_Cache is called when production densityFn_Cache is entered.
func (s *BaseMinecraftMetascriptListener) EnterDensityFn_Cache(ctx *DensityFn_CacheContext) {}

// ExitDensityFn_Cache is called when production densityFn_Cache is exited.
func (s *BaseMinecraftMetascriptListener) ExitDensityFn_Cache(ctx *DensityFn_CacheContext) {}

// EnterDensityFn_DualInput is called when production densityFn_DualInput is entered.
func (s *BaseMinecraftMetascriptListener) EnterDensityFn_DualInput(ctx *DensityFn_DualInputContext) {}

// ExitDensityFn_DualInput is called when production densityFn_DualInput is exited.
func (s *BaseMinecraftMetascriptListener) ExitDensityFn_DualInput(ctx *DensityFn_DualInputContext) {}

// EnterDensityFn_OldBlendedNoise is called when production densityFn_OldBlendedNoise is entered.
func (s *BaseMinecraftMetascriptListener) EnterDensityFn_OldBlendedNoise(ctx *DensityFn_OldBlendedNoiseContext) {
}

// ExitDensityFn_OldBlendedNoise is called when production densityFn_OldBlendedNoise is exited.
func (s *BaseMinecraftMetascriptListener) ExitDensityFn_OldBlendedNoise(ctx *DensityFn_OldBlendedNoiseContext) {
}

// EnterDensityFn_OldBlendedNoiseBuilder is called when production densityFn_OldBlendedNoiseBuilder is entered.
func (s *BaseMinecraftMetascriptListener) EnterDensityFn_OldBlendedNoiseBuilder(ctx *DensityFn_OldBlendedNoiseBuilderContext) {
}

// ExitDensityFn_OldBlendedNoiseBuilder is called when production densityFn_OldBlendedNoiseBuilder is exited.
func (s *BaseMinecraftMetascriptListener) ExitDensityFn_OldBlendedNoiseBuilder(ctx *DensityFn_OldBlendedNoiseBuilderContext) {
}

// EnterDensityFn_WierdScaledSampler is called when production densityFn_WierdScaledSampler is entered.
func (s *BaseMinecraftMetascriptListener) EnterDensityFn_WierdScaledSampler(ctx *DensityFn_WierdScaledSamplerContext) {
}

// ExitDensityFn_WierdScaledSampler is called when production densityFn_WierdScaledSampler is exited.
func (s *BaseMinecraftMetascriptListener) ExitDensityFn_WierdScaledSampler(ctx *DensityFn_WierdScaledSamplerContext) {
}

// EnterDensityFn_WierdScaledSamplerBuilder is called when production densityFn_WierdScaledSamplerBuilder is entered.
func (s *BaseMinecraftMetascriptListener) EnterDensityFn_WierdScaledSamplerBuilder(ctx *DensityFn_WierdScaledSamplerBuilderContext) {
}

// ExitDensityFn_WierdScaledSamplerBuilder is called when production densityFn_WierdScaledSamplerBuilder is exited.
func (s *BaseMinecraftMetascriptListener) ExitDensityFn_WierdScaledSamplerBuilder(ctx *DensityFn_WierdScaledSamplerBuilderContext) {
}

// EnterDensityFn_ShiftedNoise is called when production densityFn_ShiftedNoise is entered.
func (s *BaseMinecraftMetascriptListener) EnterDensityFn_ShiftedNoise(ctx *DensityFn_ShiftedNoiseContext) {
}

// ExitDensityFn_ShiftedNoise is called when production densityFn_ShiftedNoise is exited.
func (s *BaseMinecraftMetascriptListener) ExitDensityFn_ShiftedNoise(ctx *DensityFn_ShiftedNoiseContext) {
}

// EnterDensityFn_ShiftedNoiseBuilder is called when production densityFn_ShiftedNoiseBuilder is entered.
func (s *BaseMinecraftMetascriptListener) EnterDensityFn_ShiftedNoiseBuilder(ctx *DensityFn_ShiftedNoiseBuilderContext) {
}

// ExitDensityFn_ShiftedNoiseBuilder is called when production densityFn_ShiftedNoiseBuilder is exited.
func (s *BaseMinecraftMetascriptListener) ExitDensityFn_ShiftedNoiseBuilder(ctx *DensityFn_ShiftedNoiseBuilderContext) {
}

// EnterDensityFn_RangeChoice is called when production densityFn_RangeChoice is entered.
func (s *BaseMinecraftMetascriptListener) EnterDensityFn_RangeChoice(ctx *DensityFn_RangeChoiceContext) {
}

// ExitDensityFn_RangeChoice is called when production densityFn_RangeChoice is exited.
func (s *BaseMinecraftMetascriptListener) ExitDensityFn_RangeChoice(ctx *DensityFn_RangeChoiceContext) {
}

// EnterDensityFn_RangeChoiceBuilder is called when production densityFn_RangeChoiceBuilder is entered.
func (s *BaseMinecraftMetascriptListener) EnterDensityFn_RangeChoiceBuilder(ctx *DensityFn_RangeChoiceBuilderContext) {
}

// ExitDensityFn_RangeChoiceBuilder is called when production densityFn_RangeChoiceBuilder is exited.
func (s *BaseMinecraftMetascriptListener) ExitDensityFn_RangeChoiceBuilder(ctx *DensityFn_RangeChoiceBuilderContext) {
}

// EnterDensityFn_Clamp is called when production densityFn_Clamp is entered.
func (s *BaseMinecraftMetascriptListener) EnterDensityFn_Clamp(ctx *DensityFn_ClampContext) {}

// ExitDensityFn_Clamp is called when production densityFn_Clamp is exited.
func (s *BaseMinecraftMetascriptListener) ExitDensityFn_Clamp(ctx *DensityFn_ClampContext) {}

// EnterDensityFn_ClampBuilder is called when production densityFn_ClampBuilder is entered.
func (s *BaseMinecraftMetascriptListener) EnterDensityFn_ClampBuilder(ctx *DensityFn_ClampBuilderContext) {
}

// ExitDensityFn_ClampBuilder is called when production densityFn_ClampBuilder is exited.
func (s *BaseMinecraftMetascriptListener) ExitDensityFn_ClampBuilder(ctx *DensityFn_ClampBuilderContext) {
}

// EnterDensityFn_YClampedGradient is called when production densityFn_YClampedGradient is entered.
func (s *BaseMinecraftMetascriptListener) EnterDensityFn_YClampedGradient(ctx *DensityFn_YClampedGradientContext) {
}

// ExitDensityFn_YClampedGradient is called when production densityFn_YClampedGradient is exited.
func (s *BaseMinecraftMetascriptListener) ExitDensityFn_YClampedGradient(ctx *DensityFn_YClampedGradientContext) {
}

// EnterDensityFn_YClampedGradientBuilder is called when production densityFn_YClampedGradientBuilder is entered.
func (s *BaseMinecraftMetascriptListener) EnterDensityFn_YClampedGradientBuilder(ctx *DensityFn_YClampedGradientBuilderContext) {
}

// ExitDensityFn_YClampedGradientBuilder is called when production densityFn_YClampedGradientBuilder is exited.
func (s *BaseMinecraftMetascriptListener) ExitDensityFn_YClampedGradientBuilder(ctx *DensityFn_YClampedGradientBuilderContext) {
}

// EnterDensityFn_SplineFn is called when production densityFn_SplineFn is entered.
func (s *BaseMinecraftMetascriptListener) EnterDensityFn_SplineFn(ctx *DensityFn_SplineFnContext) {}

// ExitDensityFn_SplineFn is called when production densityFn_SplineFn is exited.
func (s *BaseMinecraftMetascriptListener) ExitDensityFn_SplineFn(ctx *DensityFn_SplineFnContext) {}

// EnterDensityFn_Spline is called when production densityFn_Spline is entered.
func (s *BaseMinecraftMetascriptListener) EnterDensityFn_Spline(ctx *DensityFn_SplineContext) {}

// ExitDensityFn_Spline is called when production densityFn_Spline is exited.
func (s *BaseMinecraftMetascriptListener) ExitDensityFn_Spline(ctx *DensityFn_SplineContext) {}

// EnterDensityFn_SplinePoint is called when production densityFn_SplinePoint is entered.
func (s *BaseMinecraftMetascriptListener) EnterDensityFn_SplinePoint(ctx *DensityFn_SplinePointContext) {
}

// ExitDensityFn_SplinePoint is called when production densityFn_SplinePoint is exited.
func (s *BaseMinecraftMetascriptListener) ExitDensityFn_SplinePoint(ctx *DensityFn_SplinePointContext) {
}

// EnterDensityFn_Constant is called when production densityFn_Constant is entered.
func (s *BaseMinecraftMetascriptListener) EnterDensityFn_Constant(ctx *DensityFn_ConstantContext) {}

// ExitDensityFn_Constant is called when production densityFn_Constant is exited.
func (s *BaseMinecraftMetascriptListener) ExitDensityFn_Constant(ctx *DensityFn_ConstantContext) {}

// EnterDensityFn_Reference is called when production densityFn_Reference is entered.
func (s *BaseMinecraftMetascriptListener) EnterDensityFn_Reference(ctx *DensityFn_ReferenceContext) {}

// ExitDensityFn_Reference is called when production densityFn_Reference is exited.
func (s *BaseMinecraftMetascriptListener) ExitDensityFn_Reference(ctx *DensityFn_ReferenceContext) {}

// EnterDensityFn_Math is called when production densityFn_Math is entered.
func (s *BaseMinecraftMetascriptListener) EnterDensityFn_Math(ctx *DensityFn_MathContext) {}

// ExitDensityFn_Math is called when production densityFn_Math is exited.
func (s *BaseMinecraftMetascriptListener) ExitDensityFn_Math(ctx *DensityFn_MathContext) {}

// EnterNoiseRouterBlock is called when production noiseRouterBlock is entered.
func (s *BaseMinecraftMetascriptListener) EnterNoiseRouterBlock(ctx *NoiseRouterBlockContext) {}

// ExitNoiseRouterBlock is called when production noiseRouterBlock is exited.
func (s *BaseMinecraftMetascriptListener) ExitNoiseRouterBlock(ctx *NoiseRouterBlockContext) {}

// EnterNoiseRouterDeclaration is called when production noiseRouterDeclaration is entered.
func (s *BaseMinecraftMetascriptListener) EnterNoiseRouterDeclaration(ctx *NoiseRouterDeclarationContext) {
}

// ExitNoiseRouterDeclaration is called when production noiseRouterDeclaration is exited.
func (s *BaseMinecraftMetascriptListener) ExitNoiseRouterDeclaration(ctx *NoiseRouterDeclarationContext) {
}

// EnterNoiseRouter is called when production noiseRouter is entered.
func (s *BaseMinecraftMetascriptListener) EnterNoiseRouter(ctx *NoiseRouterContext) {}

// ExitNoiseRouter is called when production noiseRouter is exited.
func (s *BaseMinecraftMetascriptListener) ExitNoiseRouter(ctx *NoiseRouterContext) {}

// EnterNoiseRouter_Builder is called when production noiseRouter_Builder is entered.
func (s *BaseMinecraftMetascriptListener) EnterNoiseRouter_Builder(ctx *NoiseRouter_BuilderContext) {}

// ExitNoiseRouter_Builder is called when production noiseRouter_Builder is exited.
func (s *BaseMinecraftMetascriptListener) ExitNoiseRouter_Builder(ctx *NoiseRouter_BuilderContext) {}

// EnterNoiseSettingsBlock is called when production noiseSettingsBlock is entered.
func (s *BaseMinecraftMetascriptListener) EnterNoiseSettingsBlock(ctx *NoiseSettingsBlockContext) {}

// ExitNoiseSettingsBlock is called when production noiseSettingsBlock is exited.
func (s *BaseMinecraftMetascriptListener) ExitNoiseSettingsBlock(ctx *NoiseSettingsBlockContext) {}

// EnterNoiseSettingsDeclaration is called when production noiseSettingsDeclaration is entered.
func (s *BaseMinecraftMetascriptListener) EnterNoiseSettingsDeclaration(ctx *NoiseSettingsDeclarationContext) {
}

// ExitNoiseSettingsDeclaration is called when production noiseSettingsDeclaration is exited.
func (s *BaseMinecraftMetascriptListener) ExitNoiseSettingsDeclaration(ctx *NoiseSettingsDeclarationContext) {
}

// EnterNoiseSettings is called when production noiseSettings is entered.
func (s *BaseMinecraftMetascriptListener) EnterNoiseSettings(ctx *NoiseSettingsContext) {}

// ExitNoiseSettings is called when production noiseSettings is exited.
func (s *BaseMinecraftMetascriptListener) ExitNoiseSettings(ctx *NoiseSettingsContext) {}

// EnterNoiseSettings_Builder is called when production noiseSettings_Builder is entered.
func (s *BaseMinecraftMetascriptListener) EnterNoiseSettings_Builder(ctx *NoiseSettings_BuilderContext) {
}

// ExitNoiseSettings_Builder is called when production noiseSettings_Builder is exited.
func (s *BaseMinecraftMetascriptListener) ExitNoiseSettings_Builder(ctx *NoiseSettings_BuilderContext) {
}

// EnterBuilder_NoiseSize is called when production builder_NoiseSize is entered.
func (s *BaseMinecraftMetascriptListener) EnterBuilder_NoiseSize(ctx *Builder_NoiseSizeContext) {}

// ExitBuilder_NoiseSize is called when production builder_NoiseSize is exited.
func (s *BaseMinecraftMetascriptListener) ExitBuilder_NoiseSize(ctx *Builder_NoiseSizeContext) {}

// EnterBuilder_NoiseRouter is called when production builder_NoiseRouter is entered.
func (s *BaseMinecraftMetascriptListener) EnterBuilder_NoiseRouter(ctx *Builder_NoiseRouterContext) {}

// ExitBuilder_NoiseRouter is called when production builder_NoiseRouter is exited.
func (s *BaseMinecraftMetascriptListener) ExitBuilder_NoiseRouter(ctx *Builder_NoiseRouterContext) {}

// EnterBuilder_SeaLevel is called when production builder_SeaLevel is entered.
func (s *BaseMinecraftMetascriptListener) EnterBuilder_SeaLevel(ctx *Builder_SeaLevelContext) {}

// ExitBuilder_SeaLevel is called when production builder_SeaLevel is exited.
func (s *BaseMinecraftMetascriptListener) ExitBuilder_SeaLevel(ctx *Builder_SeaLevelContext) {}

// EnterBuilder_DisableCreatures is called when production builder_DisableCreatures is entered.
func (s *BaseMinecraftMetascriptListener) EnterBuilder_DisableCreatures(ctx *Builder_DisableCreaturesContext) {
}

// ExitBuilder_DisableCreatures is called when production builder_DisableCreatures is exited.
func (s *BaseMinecraftMetascriptListener) ExitBuilder_DisableCreatures(ctx *Builder_DisableCreaturesContext) {
}

// EnterBuilder_DisableVeins is called when production builder_DisableVeins is entered.
func (s *BaseMinecraftMetascriptListener) EnterBuilder_DisableVeins(ctx *Builder_DisableVeinsContext) {
}

// ExitBuilder_DisableVeins is called when production builder_DisableVeins is exited.
func (s *BaseMinecraftMetascriptListener) ExitBuilder_DisableVeins(ctx *Builder_DisableVeinsContext) {
}

// EnterBuilder_DisableAquifers is called when production builder_DisableAquifers is entered.
func (s *BaseMinecraftMetascriptListener) EnterBuilder_DisableAquifers(ctx *Builder_DisableAquifersContext) {
}

// ExitBuilder_DisableAquifers is called when production builder_DisableAquifers is exited.
func (s *BaseMinecraftMetascriptListener) ExitBuilder_DisableAquifers(ctx *Builder_DisableAquifersContext) {
}

// EnterBuilder_LegacyRandomSource is called when production builder_LegacyRandomSource is entered.
func (s *BaseMinecraftMetascriptListener) EnterBuilder_LegacyRandomSource(ctx *Builder_LegacyRandomSourceContext) {
}

// ExitBuilder_LegacyRandomSource is called when production builder_LegacyRandomSource is exited.
func (s *BaseMinecraftMetascriptListener) ExitBuilder_LegacyRandomSource(ctx *Builder_LegacyRandomSourceContext) {
}

// EnterBuilder_DefaultBlock is called when production builder_DefaultBlock is entered.
func (s *BaseMinecraftMetascriptListener) EnterBuilder_DefaultBlock(ctx *Builder_DefaultBlockContext) {
}

// ExitBuilder_DefaultBlock is called when production builder_DefaultBlock is exited.
func (s *BaseMinecraftMetascriptListener) ExitBuilder_DefaultBlock(ctx *Builder_DefaultBlockContext) {
}

// EnterBuilder_DefaultFluid is called when production builder_DefaultFluid is entered.
func (s *BaseMinecraftMetascriptListener) EnterBuilder_DefaultFluid(ctx *Builder_DefaultFluidContext) {
}

// ExitBuilder_DefaultFluid is called when production builder_DefaultFluid is exited.
func (s *BaseMinecraftMetascriptListener) ExitBuilder_DefaultFluid(ctx *Builder_DefaultFluidContext) {
}

// EnterBuilder_SpawnTarget is called when production builder_SpawnTarget is entered.
func (s *BaseMinecraftMetascriptListener) EnterBuilder_SpawnTarget(ctx *Builder_SpawnTargetContext) {}

// ExitBuilder_SpawnTarget is called when production builder_SpawnTarget is exited.
func (s *BaseMinecraftMetascriptListener) ExitBuilder_SpawnTarget(ctx *Builder_SpawnTargetContext) {}

// EnterBuilder_MinY is called when production builder_MinY is entered.
func (s *BaseMinecraftMetascriptListener) EnterBuilder_MinY(ctx *Builder_MinYContext) {}

// ExitBuilder_MinY is called when production builder_MinY is exited.
func (s *BaseMinecraftMetascriptListener) ExitBuilder_MinY(ctx *Builder_MinYContext) {}

// EnterBuilder_Height is called when production builder_Height is entered.
func (s *BaseMinecraftMetascriptListener) EnterBuilder_Height(ctx *Builder_HeightContext) {}

// ExitBuilder_Height is called when production builder_Height is exited.
func (s *BaseMinecraftMetascriptListener) ExitBuilder_Height(ctx *Builder_HeightContext) {}

// EnterBuilder_SurfaceRule is called when production builder_SurfaceRule is entered.
func (s *BaseMinecraftMetascriptListener) EnterBuilder_SurfaceRule(ctx *Builder_SurfaceRuleContext) {}

// ExitBuilder_SurfaceRule is called when production builder_SurfaceRule is exited.
func (s *BaseMinecraftMetascriptListener) ExitBuilder_SurfaceRule(ctx *Builder_SurfaceRuleContext) {}

// EnterBlockState is called when production blockState is entered.
func (s *BaseMinecraftMetascriptListener) EnterBlockState(ctx *BlockStateContext) {}

// ExitBlockState is called when production blockState is exited.
func (s *BaseMinecraftMetascriptListener) ExitBlockState(ctx *BlockStateContext) {}

// EnterBlockState_Builder is called when production blockState_Builder is entered.
func (s *BaseMinecraftMetascriptListener) EnterBlockState_Builder(ctx *BlockState_BuilderContext) {}

// ExitBlockState_Builder is called when production blockState_Builder is exited.
func (s *BaseMinecraftMetascriptListener) ExitBlockState_Builder(ctx *BlockState_BuilderContext) {}
