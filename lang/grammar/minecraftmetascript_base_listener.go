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

// EnterSharedBuilder_Offset is called when production sharedBuilder_Offset is entered.
func (s *BaseMinecraftMetascriptListener) EnterSharedBuilder_Offset(ctx *SharedBuilder_OffsetContext) {
}

// ExitSharedBuilder_Offset is called when production sharedBuilder_Offset is exited.
func (s *BaseMinecraftMetascriptListener) ExitSharedBuilder_Offset(ctx *SharedBuilder_OffsetContext) {
}

// EnterSharedBuilder_Add is called when production sharedBuilder_Add is entered.
func (s *BaseMinecraftMetascriptListener) EnterSharedBuilder_Add(ctx *SharedBuilder_AddContext) {}

// ExitSharedBuilder_Add is called when production sharedBuilder_Add is exited.
func (s *BaseMinecraftMetascriptListener) ExitSharedBuilder_Add(ctx *SharedBuilder_AddContext) {}

// EnterSharedBuilder_Mul is called when production sharedBuilder_Mul is entered.
func (s *BaseMinecraftMetascriptListener) EnterSharedBuilder_Mul(ctx *SharedBuilder_MulContext) {}

// ExitSharedBuilder_Mul is called when production sharedBuilder_Mul is exited.
func (s *BaseMinecraftMetascriptListener) ExitSharedBuilder_Mul(ctx *SharedBuilder_MulContext) {}

// EnterSharedBuilder_MulInt is called when production sharedBuilder_MulInt is entered.
func (s *BaseMinecraftMetascriptListener) EnterSharedBuilder_MulInt(ctx *SharedBuilder_MulIntContext) {
}

// ExitSharedBuilder_MulInt is called when production sharedBuilder_MulInt is exited.
func (s *BaseMinecraftMetascriptListener) ExitSharedBuilder_MulInt(ctx *SharedBuilder_MulIntContext) {
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

// EnterSurfaceCondition_NoiseThresholdBuilder_Min is called when production surfaceCondition_NoiseThresholdBuilder_Min is entered.
func (s *BaseMinecraftMetascriptListener) EnterSurfaceCondition_NoiseThresholdBuilder_Min(ctx *SurfaceCondition_NoiseThresholdBuilder_MinContext) {
}

// ExitSurfaceCondition_NoiseThresholdBuilder_Min is called when production surfaceCondition_NoiseThresholdBuilder_Min is exited.
func (s *BaseMinecraftMetascriptListener) ExitSurfaceCondition_NoiseThresholdBuilder_Min(ctx *SurfaceCondition_NoiseThresholdBuilder_MinContext) {
}

// EnterSurfaceCondition_NoiseThresholdBuilder_Max is called when production surfaceCondition_NoiseThresholdBuilder_Max is entered.
func (s *BaseMinecraftMetascriptListener) EnterSurfaceCondition_NoiseThresholdBuilder_Max(ctx *SurfaceCondition_NoiseThresholdBuilder_MaxContext) {
}

// ExitSurfaceCondition_NoiseThresholdBuilder_Max is called when production surfaceCondition_NoiseThresholdBuilder_Max is exited.
func (s *BaseMinecraftMetascriptListener) ExitSurfaceCondition_NoiseThresholdBuilder_Max(ctx *SurfaceCondition_NoiseThresholdBuilder_MaxContext) {
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

// EnterSurfaceCondition_VerticalGradientBuilder_Top is called when production surfaceCondition_VerticalGradientBuilder_Top is entered.
func (s *BaseMinecraftMetascriptListener) EnterSurfaceCondition_VerticalGradientBuilder_Top(ctx *SurfaceCondition_VerticalGradientBuilder_TopContext) {
}

// ExitSurfaceCondition_VerticalGradientBuilder_Top is called when production surfaceCondition_VerticalGradientBuilder_Top is exited.
func (s *BaseMinecraftMetascriptListener) ExitSurfaceCondition_VerticalGradientBuilder_Top(ctx *SurfaceCondition_VerticalGradientBuilder_TopContext) {
}

// EnterSurfaceCondition_VerticalGradientBuilder_Bottom is called when production surfaceCondition_VerticalGradientBuilder_Bottom is entered.
func (s *BaseMinecraftMetascriptListener) EnterSurfaceCondition_VerticalGradientBuilder_Bottom(ctx *SurfaceCondition_VerticalGradientBuilder_BottomContext) {
}

// ExitSurfaceCondition_VerticalGradientBuilder_Bottom is called when production surfaceCondition_VerticalGradientBuilder_Bottom is exited.
func (s *BaseMinecraftMetascriptListener) ExitSurfaceCondition_VerticalGradientBuilder_Bottom(ctx *SurfaceCondition_VerticalGradientBuilder_BottomContext) {
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

// EnterResourceReference is called when production resourceReference is entered.
func (s *BaseMinecraftMetascriptListener) EnterResourceReference(ctx *ResourceReferenceContext) {}

// ExitResourceReference is called when production resourceReference is exited.
func (s *BaseMinecraftMetascriptListener) ExitResourceReference(ctx *ResourceReferenceContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseMinecraftMetascriptListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseMinecraftMetascriptListener) ExitNumber(ctx *NumberContext) {}

// EnterSharedBuilder_XzScale is called when production sharedBuilder_XzScale is entered.
func (s *BaseMinecraftMetascriptListener) EnterSharedBuilder_XzScale(ctx *SharedBuilder_XzScaleContext) {
}

// ExitSharedBuilder_XzScale is called when production sharedBuilder_XzScale is exited.
func (s *BaseMinecraftMetascriptListener) ExitSharedBuilder_XzScale(ctx *SharedBuilder_XzScaleContext) {
}

// EnterSharedBuilder_YScale is called when production sharedBuilder_YScale is entered.
func (s *BaseMinecraftMetascriptListener) EnterSharedBuilder_YScale(ctx *SharedBuilder_YScaleContext) {
}

// ExitSharedBuilder_YScale is called when production sharedBuilder_YScale is exited.
func (s *BaseMinecraftMetascriptListener) ExitSharedBuilder_YScale(ctx *SharedBuilder_YScaleContext) {
}

// EnterSharedBuilder_XzFactor is called when production sharedBuilder_XzFactor is entered.
func (s *BaseMinecraftMetascriptListener) EnterSharedBuilder_XzFactor(ctx *SharedBuilder_XzFactorContext) {
}

// ExitSharedBuilder_XzFactor is called when production sharedBuilder_XzFactor is exited.
func (s *BaseMinecraftMetascriptListener) ExitSharedBuilder_XzFactor(ctx *SharedBuilder_XzFactorContext) {
}

// EnterSharedBuilder_YFactor is called when production sharedBuilder_YFactor is entered.
func (s *BaseMinecraftMetascriptListener) EnterSharedBuilder_YFactor(ctx *SharedBuilder_YFactorContext) {
}

// ExitSharedBuilder_YFactor is called when production sharedBuilder_YFactor is exited.
func (s *BaseMinecraftMetascriptListener) ExitSharedBuilder_YFactor(ctx *SharedBuilder_YFactorContext) {
}

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

// EnterDensityFn_OldBlendedNoiseBuilder_Smear is called when production densityFn_OldBlendedNoiseBuilder_Smear is entered.
func (s *BaseMinecraftMetascriptListener) EnterDensityFn_OldBlendedNoiseBuilder_Smear(ctx *DensityFn_OldBlendedNoiseBuilder_SmearContext) {
}

// ExitDensityFn_OldBlendedNoiseBuilder_Smear is called when production densityFn_OldBlendedNoiseBuilder_Smear is exited.
func (s *BaseMinecraftMetascriptListener) ExitDensityFn_OldBlendedNoiseBuilder_Smear(ctx *DensityFn_OldBlendedNoiseBuilder_SmearContext) {
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

// EnterNoise_Builder is called when production noise_Builder is entered.
func (s *BaseMinecraftMetascriptListener) EnterNoise_Builder(ctx *Noise_BuilderContext) {}

// ExitNoise_Builder is called when production noise_Builder is exited.
func (s *BaseMinecraftMetascriptListener) ExitNoise_Builder(ctx *Noise_BuilderContext) {}

// EnterNoise_Builder_Amplitudes is called when production noise_Builder_Amplitudes is entered.
func (s *BaseMinecraftMetascriptListener) EnterNoise_Builder_Amplitudes(ctx *Noise_Builder_AmplitudesContext) {
}

// ExitNoise_Builder_Amplitudes is called when production noise_Builder_Amplitudes is exited.
func (s *BaseMinecraftMetascriptListener) ExitNoise_Builder_Amplitudes(ctx *Noise_Builder_AmplitudesContext) {
}
