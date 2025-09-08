// Code generated from ./grammar/MinecraftMetascript.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // MinecraftMetascript
import "github.com/antlr4-go/antlr/v4"

// MinecraftMetascriptListener is a complete listener for a parse tree produced by MinecraftMetascriptParser.
type MinecraftMetascriptListener interface {
	antlr.ParseTreeListener

	// EnterScript is called when entering the script production.
	EnterScript(c *ScriptContext)

	// EnterNamespaceDeclaration is called when entering the namespaceDeclaration production.
	EnterNamespaceDeclaration(c *NamespaceDeclarationContext)

	// EnterNamespace is called when entering the namespace production.
	EnterNamespace(c *NamespaceContext)

	// EnterContentBlocks is called when entering the contentBlocks production.
	EnterContentBlocks(c *ContentBlocksContext)

	// EnterSurfaceBlock is called when entering the surfaceBlock production.
	EnterSurfaceBlock(c *SurfaceBlockContext)

	// EnterSurfaceStatement is called when entering the surfaceStatement production.
	EnterSurfaceStatement(c *SurfaceStatementContext)

	// EnterVerticalAnchor is called when entering the verticalAnchor production.
	EnterVerticalAnchor(c *VerticalAnchorContext)

	// EnterVerticalAnchorDeclaration is called when entering the verticalAnchorDeclaration production.
	EnterVerticalAnchorDeclaration(c *VerticalAnchorDeclarationContext)

	// EnterSurfaceCondition is called when entering the surfaceCondition production.
	EnterSurfaceCondition(c *SurfaceConditionContext)

	// EnterSurfaceConditionDeclaration is called when entering the surfaceConditionDeclaration production.
	EnterSurfaceConditionDeclaration(c *SurfaceConditionDeclarationContext)

	// EnterSurfaceCondition_Not is called when entering the surfaceCondition_Not production.
	EnterSurfaceCondition_Not(c *SurfaceCondition_NotContext)

	// EnterSurfaceCondition_And is called when entering the surfaceCondition_And production.
	EnterSurfaceCondition_And(c *SurfaceCondition_AndContext)

	// EnterSurfaceCondition_Or is called when entering the surfaceCondition_Or production.
	EnterSurfaceCondition_Or(c *SurfaceCondition_OrContext)

	// EnterSurfaceCondition_Reference is called when entering the surfaceCondition_Reference production.
	EnterSurfaceCondition_Reference(c *SurfaceCondition_ReferenceContext)

	// EnterSurfaceCondition_AboveSurface is called when entering the surfaceCondition_AboveSurface production.
	EnterSurfaceCondition_AboveSurface(c *SurfaceCondition_AboveSurfaceContext)

	// EnterSurfaceCondition_Biome is called when entering the surfaceCondition_Biome production.
	EnterSurfaceCondition_Biome(c *SurfaceCondition_BiomeContext)

	// EnterSurfaceCondition_Hole is called when entering the surfaceCondition_Hole production.
	EnterSurfaceCondition_Hole(c *SurfaceCondition_HoleContext)

	// EnterSurfaceCondition_Steep is called when entering the surfaceCondition_Steep production.
	EnterSurfaceCondition_Steep(c *SurfaceCondition_SteepContext)

	// EnterSurfaceCondition_Freezing is called when entering the surfaceCondition_Freezing production.
	EnterSurfaceCondition_Freezing(c *SurfaceCondition_FreezingContext)

	// EnterSurfaceCondition_NoiseThresholdBuilder is called when entering the surfaceCondition_NoiseThresholdBuilder production.
	EnterSurfaceCondition_NoiseThresholdBuilder(c *SurfaceCondition_NoiseThresholdBuilderContext)

	// EnterSurfaceCondition_NoiseThreshold is called when entering the surfaceCondition_NoiseThreshold production.
	EnterSurfaceCondition_NoiseThreshold(c *SurfaceCondition_NoiseThresholdContext)

	// EnterSurfaceCondition_StoneDepth is called when entering the surfaceCondition_StoneDepth production.
	EnterSurfaceCondition_StoneDepth(c *SurfaceCondition_StoneDepthContext)

	// EnterSurfaceCondition_StoneDepthBuilder is called when entering the surfaceCondition_StoneDepthBuilder production.
	EnterSurfaceCondition_StoneDepthBuilder(c *SurfaceCondition_StoneDepthBuilderContext)

	// EnterSurfaceCondition_StoneDepthBuilder_SecondaryDepthRange is called when entering the surfaceCondition_StoneDepthBuilder_SecondaryDepthRange production.
	EnterSurfaceCondition_StoneDepthBuilder_SecondaryDepthRange(c *SurfaceCondition_StoneDepthBuilder_SecondaryDepthRangeContext)

	// EnterSurfaceCondition_VerticalGradient is called when entering the surfaceCondition_VerticalGradient production.
	EnterSurfaceCondition_VerticalGradient(c *SurfaceCondition_VerticalGradientContext)

	// EnterSurfaceCondition_VerticalGradientBuilder is called when entering the surfaceCondition_VerticalGradientBuilder production.
	EnterSurfaceCondition_VerticalGradientBuilder(c *SurfaceCondition_VerticalGradientBuilderContext)

	// EnterSurfaceCondition_AboveWater is called when entering the surfaceCondition_AboveWater production.
	EnterSurfaceCondition_AboveWater(c *SurfaceCondition_AboveWaterContext)

	// EnterSurfaceCondition_AboveWaterBuilder is called when entering the surfaceCondition_AboveWaterBuilder production.
	EnterSurfaceCondition_AboveWaterBuilder(c *SurfaceCondition_AboveWaterBuilderContext)

	// EnterSurfaceCondition_YAbove is called when entering the surfaceCondition_YAbove production.
	EnterSurfaceCondition_YAbove(c *SurfaceCondition_YAboveContext)

	// EnterSurfaceCondition_YAboveBuilder is called when entering the surfaceCondition_YAboveBuilder production.
	EnterSurfaceCondition_YAboveBuilder(c *SurfaceCondition_YAboveBuilderContext)

	// EnterSurfaceRuleDeclaration is called when entering the surfaceRuleDeclaration production.
	EnterSurfaceRuleDeclaration(c *SurfaceRuleDeclarationContext)

	// EnterSurfaceRule is called when entering the surfaceRule production.
	EnterSurfaceRule(c *SurfaceRuleContext)

	// EnterSurfaceRule_Reference is called when entering the surfaceRule_Reference production.
	EnterSurfaceRule_Reference(c *SurfaceRule_ReferenceContext)

	// EnterSurfaceRule_Block is called when entering the surfaceRule_Block production.
	EnterSurfaceRule_Block(c *SurfaceRule_BlockContext)

	// EnterSurfaceRule_Sequence is called when entering the surfaceRule_Sequence production.
	EnterSurfaceRule_Sequence(c *SurfaceRule_SequenceContext)

	// EnterSurfaceRule_Bandlands is called when entering the surfaceRule_Bandlands production.
	EnterSurfaceRule_Bandlands(c *SurfaceRule_BandlandsContext)

	// EnterSurfaceRule_If is called when entering the surfaceRule_If production.
	EnterSurfaceRule_If(c *SurfaceRule_IfContext)

	// EnterBuilder_XZScale is called when entering the builder_XZScale production.
	EnterBuilder_XZScale(c *Builder_XZScaleContext)

	// EnterBuilder_YScale is called when entering the builder_YScale production.
	EnterBuilder_YScale(c *Builder_YScaleContext)

	// EnterBuilder_XZFactor is called when entering the builder_XZFactor production.
	EnterBuilder_XZFactor(c *Builder_XZFactorContext)

	// EnterBuilder_YFactor is called when entering the builder_YFactor production.
	EnterBuilder_YFactor(c *Builder_YFactorContext)

	// EnterBuilder_Noise is called when entering the builder_Noise production.
	EnterBuilder_Noise(c *Builder_NoiseContext)

	// EnterBuilder_Smear is called when entering the builder_Smear production.
	EnterBuilder_Smear(c *Builder_SmearContext)

	// EnterBuilder_Type1 is called when entering the builder_Type1 production.
	EnterBuilder_Type1(c *Builder_Type1Context)

	// EnterBuilder_Type2 is called when entering the builder_Type2 production.
	EnterBuilder_Type2(c *Builder_Type2Context)

	// EnterBuilder_Shift is called when entering the builder_Shift production.
	EnterBuilder_Shift(c *Builder_ShiftContext)

	// EnterBuilder_Amplitudes is called when entering the builder_Amplitudes production.
	EnterBuilder_Amplitudes(c *Builder_AmplitudesContext)

	// EnterBuilder_Offset is called when entering the builder_Offset production.
	EnterBuilder_Offset(c *Builder_OffsetContext)

	// EnterBuilder_Add is called when entering the builder_Add production.
	EnterBuilder_Add(c *Builder_AddContext)

	// EnterBuilder_Mul is called when entering the builder_Mul production.
	EnterBuilder_Mul(c *Builder_MulContext)

	// EnterBuilder_MulInt is called when entering the builder_MulInt production.
	EnterBuilder_MulInt(c *Builder_MulIntContext)

	// EnterBuilder_Min is called when entering the builder_Min production.
	EnterBuilder_Min(c *Builder_MinContext)

	// EnterBuilder_Max is called when entering the builder_Max production.
	EnterBuilder_Max(c *Builder_MaxContext)

	// EnterBuilder_Top is called when entering the builder_Top production.
	EnterBuilder_Top(c *Builder_TopContext)

	// EnterBuilder_TopLiteral is called when entering the builder_TopLiteral production.
	EnterBuilder_TopLiteral(c *Builder_TopLiteralContext)

	// EnterBuilder_Bottom is called when entering the builder_Bottom production.
	EnterBuilder_Bottom(c *Builder_BottomContext)

	// EnterBuilder_BottomLiteral is called when entering the builder_BottomLiteral production.
	EnterBuilder_BottomLiteral(c *Builder_BottomLiteralContext)

	// EnterBuilder_InRange is called when entering the builder_InRange production.
	EnterBuilder_InRange(c *Builder_InRangeContext)

	// EnterBuilder_OutRange is called when entering the builder_OutRange production.
	EnterBuilder_OutRange(c *Builder_OutRangeContext)

	// EnterResourceReference is called when entering the resourceReference production.
	EnterResourceReference(c *ResourceReferenceContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterNoiseBlock is called when entering the noiseBlock production.
	EnterNoiseBlock(c *NoiseBlockContext)

	// EnterNoiseDeclaration is called when entering the noiseDeclaration production.
	EnterNoiseDeclaration(c *NoiseDeclarationContext)

	// EnterNoise is called when entering the noise production.
	EnterNoise(c *NoiseContext)

	// EnterNoiseDefinition is called when entering the noiseDefinition production.
	EnterNoiseDefinition(c *NoiseDefinitionContext)

	// EnterNoise_Builder is called when entering the noise_Builder production.
	EnterNoise_Builder(c *Noise_BuilderContext)

	// EnterDensityFnBlock is called when entering the densityFnBlock production.
	EnterDensityFnBlock(c *DensityFnBlockContext)

	// EnterDensityFnDeclaration is called when entering the densityFnDeclaration production.
	EnterDensityFnDeclaration(c *DensityFnDeclarationContext)

	// EnterDensityFn is called when entering the densityFn production.
	EnterDensityFn(c *DensityFnContext)

	// EnterDensityFn_NoInput is called when entering the densityFn_NoInput production.
	EnterDensityFn_NoInput(c *DensityFn_NoInputContext)

	// EnterDensityFn_SingleInput is called when entering the densityFn_SingleInput production.
	EnterDensityFn_SingleInput(c *DensityFn_SingleInputContext)

	// EnterDensityFn_InlineNoise is called when entering the densityFn_InlineNoise production.
	EnterDensityFn_InlineNoise(c *DensityFn_InlineNoiseContext)

	// EnterDensityFn_Noise is called when entering the densityFn_Noise production.
	EnterDensityFn_Noise(c *DensityFn_NoiseContext)

	// EnterDensityFn_NoiseBuilder is called when entering the densityFn_NoiseBuilder production.
	EnterDensityFn_NoiseBuilder(c *DensityFn_NoiseBuilderContext)

	// EnterDensityFn_Cache is called when entering the densityFn_Cache production.
	EnterDensityFn_Cache(c *DensityFn_CacheContext)

	// EnterDensityFn_DualInput is called when entering the densityFn_DualInput production.
	EnterDensityFn_DualInput(c *DensityFn_DualInputContext)

	// EnterDensityFn_OldBlendedNoise is called when entering the densityFn_OldBlendedNoise production.
	EnterDensityFn_OldBlendedNoise(c *DensityFn_OldBlendedNoiseContext)

	// EnterDensityFn_OldBlendedNoiseBuilder is called when entering the densityFn_OldBlendedNoiseBuilder production.
	EnterDensityFn_OldBlendedNoiseBuilder(c *DensityFn_OldBlendedNoiseBuilderContext)

	// EnterDensityFn_WierdScaledSampler is called when entering the densityFn_WierdScaledSampler production.
	EnterDensityFn_WierdScaledSampler(c *DensityFn_WierdScaledSamplerContext)

	// EnterDensityFn_WierdScaledSamplerBuilder is called when entering the densityFn_WierdScaledSamplerBuilder production.
	EnterDensityFn_WierdScaledSamplerBuilder(c *DensityFn_WierdScaledSamplerBuilderContext)

	// EnterDensityFn_ShiftedNoise is called when entering the densityFn_ShiftedNoise production.
	EnterDensityFn_ShiftedNoise(c *DensityFn_ShiftedNoiseContext)

	// EnterDensityFn_ShiftedNoiseBuilder is called when entering the densityFn_ShiftedNoiseBuilder production.
	EnterDensityFn_ShiftedNoiseBuilder(c *DensityFn_ShiftedNoiseBuilderContext)

	// EnterDensityFn_RangeChoice is called when entering the densityFn_RangeChoice production.
	EnterDensityFn_RangeChoice(c *DensityFn_RangeChoiceContext)

	// EnterDensityFn_RangeChoiceBuilder is called when entering the densityFn_RangeChoiceBuilder production.
	EnterDensityFn_RangeChoiceBuilder(c *DensityFn_RangeChoiceBuilderContext)

	// EnterDensityFn_Clamp is called when entering the densityFn_Clamp production.
	EnterDensityFn_Clamp(c *DensityFn_ClampContext)

	// EnterDensityFn_ClampBuilder is called when entering the densityFn_ClampBuilder production.
	EnterDensityFn_ClampBuilder(c *DensityFn_ClampBuilderContext)

	// EnterDensityFn_YClampedGradient is called when entering the densityFn_YClampedGradient production.
	EnterDensityFn_YClampedGradient(c *DensityFn_YClampedGradientContext)

	// EnterDensityFn_YClampedGradientBuilder is called when entering the densityFn_YClampedGradientBuilder production.
	EnterDensityFn_YClampedGradientBuilder(c *DensityFn_YClampedGradientBuilderContext)

	// EnterDensityFn_SplineFn is called when entering the densityFn_SplineFn production.
	EnterDensityFn_SplineFn(c *DensityFn_SplineFnContext)

	// EnterDensityFn_Spline is called when entering the densityFn_Spline production.
	EnterDensityFn_Spline(c *DensityFn_SplineContext)

	// EnterDensityFn_SplinePoint is called when entering the densityFn_SplinePoint production.
	EnterDensityFn_SplinePoint(c *DensityFn_SplinePointContext)

	// EnterDensityFn_Constant is called when entering the densityFn_Constant production.
	EnterDensityFn_Constant(c *DensityFn_ConstantContext)

	// EnterDensityFn_Reference is called when entering the densityFn_Reference production.
	EnterDensityFn_Reference(c *DensityFn_ReferenceContext)

	// EnterDensityFn_Math is called when entering the densityFn_Math production.
	EnterDensityFn_Math(c *DensityFn_MathContext)

	// ExitScript is called when exiting the script production.
	ExitScript(c *ScriptContext)

	// ExitNamespaceDeclaration is called when exiting the namespaceDeclaration production.
	ExitNamespaceDeclaration(c *NamespaceDeclarationContext)

	// ExitNamespace is called when exiting the namespace production.
	ExitNamespace(c *NamespaceContext)

	// ExitContentBlocks is called when exiting the contentBlocks production.
	ExitContentBlocks(c *ContentBlocksContext)

	// ExitSurfaceBlock is called when exiting the surfaceBlock production.
	ExitSurfaceBlock(c *SurfaceBlockContext)

	// ExitSurfaceStatement is called when exiting the surfaceStatement production.
	ExitSurfaceStatement(c *SurfaceStatementContext)

	// ExitVerticalAnchor is called when exiting the verticalAnchor production.
	ExitVerticalAnchor(c *VerticalAnchorContext)

	// ExitVerticalAnchorDeclaration is called when exiting the verticalAnchorDeclaration production.
	ExitVerticalAnchorDeclaration(c *VerticalAnchorDeclarationContext)

	// ExitSurfaceCondition is called when exiting the surfaceCondition production.
	ExitSurfaceCondition(c *SurfaceConditionContext)

	// ExitSurfaceConditionDeclaration is called when exiting the surfaceConditionDeclaration production.
	ExitSurfaceConditionDeclaration(c *SurfaceConditionDeclarationContext)

	// ExitSurfaceCondition_Not is called when exiting the surfaceCondition_Not production.
	ExitSurfaceCondition_Not(c *SurfaceCondition_NotContext)

	// ExitSurfaceCondition_And is called when exiting the surfaceCondition_And production.
	ExitSurfaceCondition_And(c *SurfaceCondition_AndContext)

	// ExitSurfaceCondition_Or is called when exiting the surfaceCondition_Or production.
	ExitSurfaceCondition_Or(c *SurfaceCondition_OrContext)

	// ExitSurfaceCondition_Reference is called when exiting the surfaceCondition_Reference production.
	ExitSurfaceCondition_Reference(c *SurfaceCondition_ReferenceContext)

	// ExitSurfaceCondition_AboveSurface is called when exiting the surfaceCondition_AboveSurface production.
	ExitSurfaceCondition_AboveSurface(c *SurfaceCondition_AboveSurfaceContext)

	// ExitSurfaceCondition_Biome is called when exiting the surfaceCondition_Biome production.
	ExitSurfaceCondition_Biome(c *SurfaceCondition_BiomeContext)

	// ExitSurfaceCondition_Hole is called when exiting the surfaceCondition_Hole production.
	ExitSurfaceCondition_Hole(c *SurfaceCondition_HoleContext)

	// ExitSurfaceCondition_Steep is called when exiting the surfaceCondition_Steep production.
	ExitSurfaceCondition_Steep(c *SurfaceCondition_SteepContext)

	// ExitSurfaceCondition_Freezing is called when exiting the surfaceCondition_Freezing production.
	ExitSurfaceCondition_Freezing(c *SurfaceCondition_FreezingContext)

	// ExitSurfaceCondition_NoiseThresholdBuilder is called when exiting the surfaceCondition_NoiseThresholdBuilder production.
	ExitSurfaceCondition_NoiseThresholdBuilder(c *SurfaceCondition_NoiseThresholdBuilderContext)

	// ExitSurfaceCondition_NoiseThreshold is called when exiting the surfaceCondition_NoiseThreshold production.
	ExitSurfaceCondition_NoiseThreshold(c *SurfaceCondition_NoiseThresholdContext)

	// ExitSurfaceCondition_StoneDepth is called when exiting the surfaceCondition_StoneDepth production.
	ExitSurfaceCondition_StoneDepth(c *SurfaceCondition_StoneDepthContext)

	// ExitSurfaceCondition_StoneDepthBuilder is called when exiting the surfaceCondition_StoneDepthBuilder production.
	ExitSurfaceCondition_StoneDepthBuilder(c *SurfaceCondition_StoneDepthBuilderContext)

	// ExitSurfaceCondition_StoneDepthBuilder_SecondaryDepthRange is called when exiting the surfaceCondition_StoneDepthBuilder_SecondaryDepthRange production.
	ExitSurfaceCondition_StoneDepthBuilder_SecondaryDepthRange(c *SurfaceCondition_StoneDepthBuilder_SecondaryDepthRangeContext)

	// ExitSurfaceCondition_VerticalGradient is called when exiting the surfaceCondition_VerticalGradient production.
	ExitSurfaceCondition_VerticalGradient(c *SurfaceCondition_VerticalGradientContext)

	// ExitSurfaceCondition_VerticalGradientBuilder is called when exiting the surfaceCondition_VerticalGradientBuilder production.
	ExitSurfaceCondition_VerticalGradientBuilder(c *SurfaceCondition_VerticalGradientBuilderContext)

	// ExitSurfaceCondition_AboveWater is called when exiting the surfaceCondition_AboveWater production.
	ExitSurfaceCondition_AboveWater(c *SurfaceCondition_AboveWaterContext)

	// ExitSurfaceCondition_AboveWaterBuilder is called when exiting the surfaceCondition_AboveWaterBuilder production.
	ExitSurfaceCondition_AboveWaterBuilder(c *SurfaceCondition_AboveWaterBuilderContext)

	// ExitSurfaceCondition_YAbove is called when exiting the surfaceCondition_YAbove production.
	ExitSurfaceCondition_YAbove(c *SurfaceCondition_YAboveContext)

	// ExitSurfaceCondition_YAboveBuilder is called when exiting the surfaceCondition_YAboveBuilder production.
	ExitSurfaceCondition_YAboveBuilder(c *SurfaceCondition_YAboveBuilderContext)

	// ExitSurfaceRuleDeclaration is called when exiting the surfaceRuleDeclaration production.
	ExitSurfaceRuleDeclaration(c *SurfaceRuleDeclarationContext)

	// ExitSurfaceRule is called when exiting the surfaceRule production.
	ExitSurfaceRule(c *SurfaceRuleContext)

	// ExitSurfaceRule_Reference is called when exiting the surfaceRule_Reference production.
	ExitSurfaceRule_Reference(c *SurfaceRule_ReferenceContext)

	// ExitSurfaceRule_Block is called when exiting the surfaceRule_Block production.
	ExitSurfaceRule_Block(c *SurfaceRule_BlockContext)

	// ExitSurfaceRule_Sequence is called when exiting the surfaceRule_Sequence production.
	ExitSurfaceRule_Sequence(c *SurfaceRule_SequenceContext)

	// ExitSurfaceRule_Bandlands is called when exiting the surfaceRule_Bandlands production.
	ExitSurfaceRule_Bandlands(c *SurfaceRule_BandlandsContext)

	// ExitSurfaceRule_If is called when exiting the surfaceRule_If production.
	ExitSurfaceRule_If(c *SurfaceRule_IfContext)

	// ExitBuilder_XZScale is called when exiting the builder_XZScale production.
	ExitBuilder_XZScale(c *Builder_XZScaleContext)

	// ExitBuilder_YScale is called when exiting the builder_YScale production.
	ExitBuilder_YScale(c *Builder_YScaleContext)

	// ExitBuilder_XZFactor is called when exiting the builder_XZFactor production.
	ExitBuilder_XZFactor(c *Builder_XZFactorContext)

	// ExitBuilder_YFactor is called when exiting the builder_YFactor production.
	ExitBuilder_YFactor(c *Builder_YFactorContext)

	// ExitBuilder_Noise is called when exiting the builder_Noise production.
	ExitBuilder_Noise(c *Builder_NoiseContext)

	// ExitBuilder_Smear is called when exiting the builder_Smear production.
	ExitBuilder_Smear(c *Builder_SmearContext)

	// ExitBuilder_Type1 is called when exiting the builder_Type1 production.
	ExitBuilder_Type1(c *Builder_Type1Context)

	// ExitBuilder_Type2 is called when exiting the builder_Type2 production.
	ExitBuilder_Type2(c *Builder_Type2Context)

	// ExitBuilder_Shift is called when exiting the builder_Shift production.
	ExitBuilder_Shift(c *Builder_ShiftContext)

	// ExitBuilder_Amplitudes is called when exiting the builder_Amplitudes production.
	ExitBuilder_Amplitudes(c *Builder_AmplitudesContext)

	// ExitBuilder_Offset is called when exiting the builder_Offset production.
	ExitBuilder_Offset(c *Builder_OffsetContext)

	// ExitBuilder_Add is called when exiting the builder_Add production.
	ExitBuilder_Add(c *Builder_AddContext)

	// ExitBuilder_Mul is called when exiting the builder_Mul production.
	ExitBuilder_Mul(c *Builder_MulContext)

	// ExitBuilder_MulInt is called when exiting the builder_MulInt production.
	ExitBuilder_MulInt(c *Builder_MulIntContext)

	// ExitBuilder_Min is called when exiting the builder_Min production.
	ExitBuilder_Min(c *Builder_MinContext)

	// ExitBuilder_Max is called when exiting the builder_Max production.
	ExitBuilder_Max(c *Builder_MaxContext)

	// ExitBuilder_Top is called when exiting the builder_Top production.
	ExitBuilder_Top(c *Builder_TopContext)

	// ExitBuilder_TopLiteral is called when exiting the builder_TopLiteral production.
	ExitBuilder_TopLiteral(c *Builder_TopLiteralContext)

	// ExitBuilder_Bottom is called when exiting the builder_Bottom production.
	ExitBuilder_Bottom(c *Builder_BottomContext)

	// ExitBuilder_BottomLiteral is called when exiting the builder_BottomLiteral production.
	ExitBuilder_BottomLiteral(c *Builder_BottomLiteralContext)

	// ExitBuilder_InRange is called when exiting the builder_InRange production.
	ExitBuilder_InRange(c *Builder_InRangeContext)

	// ExitBuilder_OutRange is called when exiting the builder_OutRange production.
	ExitBuilder_OutRange(c *Builder_OutRangeContext)

	// ExitResourceReference is called when exiting the resourceReference production.
	ExitResourceReference(c *ResourceReferenceContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitNoiseBlock is called when exiting the noiseBlock production.
	ExitNoiseBlock(c *NoiseBlockContext)

	// ExitNoiseDeclaration is called when exiting the noiseDeclaration production.
	ExitNoiseDeclaration(c *NoiseDeclarationContext)

	// ExitNoise is called when exiting the noise production.
	ExitNoise(c *NoiseContext)

	// ExitNoiseDefinition is called when exiting the noiseDefinition production.
	ExitNoiseDefinition(c *NoiseDefinitionContext)

	// ExitNoise_Builder is called when exiting the noise_Builder production.
	ExitNoise_Builder(c *Noise_BuilderContext)

	// ExitDensityFnBlock is called when exiting the densityFnBlock production.
	ExitDensityFnBlock(c *DensityFnBlockContext)

	// ExitDensityFnDeclaration is called when exiting the densityFnDeclaration production.
	ExitDensityFnDeclaration(c *DensityFnDeclarationContext)

	// ExitDensityFn is called when exiting the densityFn production.
	ExitDensityFn(c *DensityFnContext)

	// ExitDensityFn_NoInput is called when exiting the densityFn_NoInput production.
	ExitDensityFn_NoInput(c *DensityFn_NoInputContext)

	// ExitDensityFn_SingleInput is called when exiting the densityFn_SingleInput production.
	ExitDensityFn_SingleInput(c *DensityFn_SingleInputContext)

	// ExitDensityFn_InlineNoise is called when exiting the densityFn_InlineNoise production.
	ExitDensityFn_InlineNoise(c *DensityFn_InlineNoiseContext)

	// ExitDensityFn_Noise is called when exiting the densityFn_Noise production.
	ExitDensityFn_Noise(c *DensityFn_NoiseContext)

	// ExitDensityFn_NoiseBuilder is called when exiting the densityFn_NoiseBuilder production.
	ExitDensityFn_NoiseBuilder(c *DensityFn_NoiseBuilderContext)

	// ExitDensityFn_Cache is called when exiting the densityFn_Cache production.
	ExitDensityFn_Cache(c *DensityFn_CacheContext)

	// ExitDensityFn_DualInput is called when exiting the densityFn_DualInput production.
	ExitDensityFn_DualInput(c *DensityFn_DualInputContext)

	// ExitDensityFn_OldBlendedNoise is called when exiting the densityFn_OldBlendedNoise production.
	ExitDensityFn_OldBlendedNoise(c *DensityFn_OldBlendedNoiseContext)

	// ExitDensityFn_OldBlendedNoiseBuilder is called when exiting the densityFn_OldBlendedNoiseBuilder production.
	ExitDensityFn_OldBlendedNoiseBuilder(c *DensityFn_OldBlendedNoiseBuilderContext)

	// ExitDensityFn_WierdScaledSampler is called when exiting the densityFn_WierdScaledSampler production.
	ExitDensityFn_WierdScaledSampler(c *DensityFn_WierdScaledSamplerContext)

	// ExitDensityFn_WierdScaledSamplerBuilder is called when exiting the densityFn_WierdScaledSamplerBuilder production.
	ExitDensityFn_WierdScaledSamplerBuilder(c *DensityFn_WierdScaledSamplerBuilderContext)

	// ExitDensityFn_ShiftedNoise is called when exiting the densityFn_ShiftedNoise production.
	ExitDensityFn_ShiftedNoise(c *DensityFn_ShiftedNoiseContext)

	// ExitDensityFn_ShiftedNoiseBuilder is called when exiting the densityFn_ShiftedNoiseBuilder production.
	ExitDensityFn_ShiftedNoiseBuilder(c *DensityFn_ShiftedNoiseBuilderContext)

	// ExitDensityFn_RangeChoice is called when exiting the densityFn_RangeChoice production.
	ExitDensityFn_RangeChoice(c *DensityFn_RangeChoiceContext)

	// ExitDensityFn_RangeChoiceBuilder is called when exiting the densityFn_RangeChoiceBuilder production.
	ExitDensityFn_RangeChoiceBuilder(c *DensityFn_RangeChoiceBuilderContext)

	// ExitDensityFn_Clamp is called when exiting the densityFn_Clamp production.
	ExitDensityFn_Clamp(c *DensityFn_ClampContext)

	// ExitDensityFn_ClampBuilder is called when exiting the densityFn_ClampBuilder production.
	ExitDensityFn_ClampBuilder(c *DensityFn_ClampBuilderContext)

	// ExitDensityFn_YClampedGradient is called when exiting the densityFn_YClampedGradient production.
	ExitDensityFn_YClampedGradient(c *DensityFn_YClampedGradientContext)

	// ExitDensityFn_YClampedGradientBuilder is called when exiting the densityFn_YClampedGradientBuilder production.
	ExitDensityFn_YClampedGradientBuilder(c *DensityFn_YClampedGradientBuilderContext)

	// ExitDensityFn_SplineFn is called when exiting the densityFn_SplineFn production.
	ExitDensityFn_SplineFn(c *DensityFn_SplineFnContext)

	// ExitDensityFn_Spline is called when exiting the densityFn_Spline production.
	ExitDensityFn_Spline(c *DensityFn_SplineContext)

	// ExitDensityFn_SplinePoint is called when exiting the densityFn_SplinePoint production.
	ExitDensityFn_SplinePoint(c *DensityFn_SplinePointContext)

	// ExitDensityFn_Constant is called when exiting the densityFn_Constant production.
	ExitDensityFn_Constant(c *DensityFn_ConstantContext)

	// ExitDensityFn_Reference is called when exiting the densityFn_Reference production.
	ExitDensityFn_Reference(c *DensityFn_ReferenceContext)

	// ExitDensityFn_Math is called when exiting the densityFn_Math production.
	ExitDensityFn_Math(c *DensityFn_MathContext)
}
