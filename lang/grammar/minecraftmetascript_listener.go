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

	// EnterSharedBuilder_Offset is called when entering the sharedBuilder_Offset production.
	EnterSharedBuilder_Offset(c *SharedBuilder_OffsetContext)

	// EnterSharedBuilder_Add is called when entering the sharedBuilder_Add production.
	EnterSharedBuilder_Add(c *SharedBuilder_AddContext)

	// EnterSharedBuilder_Mul is called when entering the sharedBuilder_Mul production.
	EnterSharedBuilder_Mul(c *SharedBuilder_MulContext)

	// EnterSharedBuilder_MulInt is called when entering the sharedBuilder_MulInt production.
	EnterSharedBuilder_MulInt(c *SharedBuilder_MulIntContext)

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

	// EnterSurfaceCondition_NoiseThresholdBuilder_Min is called when entering the surfaceCondition_NoiseThresholdBuilder_Min production.
	EnterSurfaceCondition_NoiseThresholdBuilder_Min(c *SurfaceCondition_NoiseThresholdBuilder_MinContext)

	// EnterSurfaceCondition_NoiseThresholdBuilder_Max is called when entering the surfaceCondition_NoiseThresholdBuilder_Max production.
	EnterSurfaceCondition_NoiseThresholdBuilder_Max(c *SurfaceCondition_NoiseThresholdBuilder_MaxContext)

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

	// EnterSurfaceCondition_VerticalGradientBuilder_Top is called when entering the surfaceCondition_VerticalGradientBuilder_Top production.
	EnterSurfaceCondition_VerticalGradientBuilder_Top(c *SurfaceCondition_VerticalGradientBuilder_TopContext)

	// EnterSurfaceCondition_VerticalGradientBuilder_Bottom is called when entering the surfaceCondition_VerticalGradientBuilder_Bottom production.
	EnterSurfaceCondition_VerticalGradientBuilder_Bottom(c *SurfaceCondition_VerticalGradientBuilder_BottomContext)

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

	// EnterResourceReference is called when entering the resourceReference production.
	EnterResourceReference(c *ResourceReferenceContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterSharedBuilder_XzScale is called when entering the sharedBuilder_XzScale production.
	EnterSharedBuilder_XzScale(c *SharedBuilder_XzScaleContext)

	// EnterSharedBuilder_YScale is called when entering the sharedBuilder_YScale production.
	EnterSharedBuilder_YScale(c *SharedBuilder_YScaleContext)

	// EnterSharedBuilder_XzFactor is called when entering the sharedBuilder_XzFactor production.
	EnterSharedBuilder_XzFactor(c *SharedBuilder_XzFactorContext)

	// EnterSharedBuilder_YFactor is called when entering the sharedBuilder_YFactor production.
	EnterSharedBuilder_YFactor(c *SharedBuilder_YFactorContext)

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

	// EnterDensityFn_OldBlendedNoiseBuilder_Smear is called when entering the densityFn_OldBlendedNoiseBuilder_Smear production.
	EnterDensityFn_OldBlendedNoiseBuilder_Smear(c *DensityFn_OldBlendedNoiseBuilder_SmearContext)

	// EnterDensityFn_Constant is called when entering the densityFn_Constant production.
	EnterDensityFn_Constant(c *DensityFn_ConstantContext)

	// EnterDensityFn_Reference is called when entering the densityFn_Reference production.
	EnterDensityFn_Reference(c *DensityFn_ReferenceContext)

	// EnterDensityFn_Math is called when entering the densityFn_Math production.
	EnterDensityFn_Math(c *DensityFn_MathContext)

	// EnterNoiseBlock is called when entering the noiseBlock production.
	EnterNoiseBlock(c *NoiseBlockContext)

	// EnterNoiseDeclaration is called when entering the noiseDeclaration production.
	EnterNoiseDeclaration(c *NoiseDeclarationContext)

	// EnterNoise is called when entering the noise production.
	EnterNoise(c *NoiseContext)

	// EnterNoise_Builder is called when entering the noise_Builder production.
	EnterNoise_Builder(c *Noise_BuilderContext)

	// EnterNoise_Builder_Amplitudes is called when entering the noise_Builder_Amplitudes production.
	EnterNoise_Builder_Amplitudes(c *Noise_Builder_AmplitudesContext)

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

	// ExitSharedBuilder_Offset is called when exiting the sharedBuilder_Offset production.
	ExitSharedBuilder_Offset(c *SharedBuilder_OffsetContext)

	// ExitSharedBuilder_Add is called when exiting the sharedBuilder_Add production.
	ExitSharedBuilder_Add(c *SharedBuilder_AddContext)

	// ExitSharedBuilder_Mul is called when exiting the sharedBuilder_Mul production.
	ExitSharedBuilder_Mul(c *SharedBuilder_MulContext)

	// ExitSharedBuilder_MulInt is called when exiting the sharedBuilder_MulInt production.
	ExitSharedBuilder_MulInt(c *SharedBuilder_MulIntContext)

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

	// ExitSurfaceCondition_NoiseThresholdBuilder_Min is called when exiting the surfaceCondition_NoiseThresholdBuilder_Min production.
	ExitSurfaceCondition_NoiseThresholdBuilder_Min(c *SurfaceCondition_NoiseThresholdBuilder_MinContext)

	// ExitSurfaceCondition_NoiseThresholdBuilder_Max is called when exiting the surfaceCondition_NoiseThresholdBuilder_Max production.
	ExitSurfaceCondition_NoiseThresholdBuilder_Max(c *SurfaceCondition_NoiseThresholdBuilder_MaxContext)

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

	// ExitSurfaceCondition_VerticalGradientBuilder_Top is called when exiting the surfaceCondition_VerticalGradientBuilder_Top production.
	ExitSurfaceCondition_VerticalGradientBuilder_Top(c *SurfaceCondition_VerticalGradientBuilder_TopContext)

	// ExitSurfaceCondition_VerticalGradientBuilder_Bottom is called when exiting the surfaceCondition_VerticalGradientBuilder_Bottom production.
	ExitSurfaceCondition_VerticalGradientBuilder_Bottom(c *SurfaceCondition_VerticalGradientBuilder_BottomContext)

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

	// ExitResourceReference is called when exiting the resourceReference production.
	ExitResourceReference(c *ResourceReferenceContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitSharedBuilder_XzScale is called when exiting the sharedBuilder_XzScale production.
	ExitSharedBuilder_XzScale(c *SharedBuilder_XzScaleContext)

	// ExitSharedBuilder_YScale is called when exiting the sharedBuilder_YScale production.
	ExitSharedBuilder_YScale(c *SharedBuilder_YScaleContext)

	// ExitSharedBuilder_XzFactor is called when exiting the sharedBuilder_XzFactor production.
	ExitSharedBuilder_XzFactor(c *SharedBuilder_XzFactorContext)

	// ExitSharedBuilder_YFactor is called when exiting the sharedBuilder_YFactor production.
	ExitSharedBuilder_YFactor(c *SharedBuilder_YFactorContext)

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

	// ExitDensityFn_OldBlendedNoiseBuilder_Smear is called when exiting the densityFn_OldBlendedNoiseBuilder_Smear production.
	ExitDensityFn_OldBlendedNoiseBuilder_Smear(c *DensityFn_OldBlendedNoiseBuilder_SmearContext)

	// ExitDensityFn_Constant is called when exiting the densityFn_Constant production.
	ExitDensityFn_Constant(c *DensityFn_ConstantContext)

	// ExitDensityFn_Reference is called when exiting the densityFn_Reference production.
	ExitDensityFn_Reference(c *DensityFn_ReferenceContext)

	// ExitDensityFn_Math is called when exiting the densityFn_Math production.
	ExitDensityFn_Math(c *DensityFn_MathContext)

	// ExitNoiseBlock is called when exiting the noiseBlock production.
	ExitNoiseBlock(c *NoiseBlockContext)

	// ExitNoiseDeclaration is called when exiting the noiseDeclaration production.
	ExitNoiseDeclaration(c *NoiseDeclarationContext)

	// ExitNoise is called when exiting the noise production.
	ExitNoise(c *NoiseContext)

	// ExitNoise_Builder is called when exiting the noise_Builder production.
	ExitNoise_Builder(c *Noise_BuilderContext)

	// ExitNoise_Builder_Amplitudes is called when exiting the noise_Builder_Amplitudes production.
	ExitNoise_Builder_Amplitudes(c *Noise_Builder_AmplitudesContext)
}
