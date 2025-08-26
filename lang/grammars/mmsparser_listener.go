// Code generated from ./grammars/MMSParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammars // MMSParser
import "github.com/antlr4-go/antlr/v4"

// MMSParserListener is a complete listener for a parse tree produced by MMSParser.
type MMSParserListener interface {
	antlr.ParseTreeListener

	// EnterNamespaceDeclaration is called when entering the namespaceDeclaration production.
	EnterNamespaceDeclaration(c *NamespaceDeclarationContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterMmsFile is called when entering the mmsFile production.
	EnterMmsFile(c *MmsFileContext)

	// EnterSurfaceDeclaration is called when entering the surfaceDeclaration production.
	EnterSurfaceDeclaration(c *SurfaceDeclarationContext)

	// EnterSurfaceDefinition is called when entering the surfaceDefinition production.
	EnterSurfaceDefinition(c *SurfaceDefinitionContext)

	// EnterSurfaceRuleReference is called when entering the surfaceRuleReference production.
	EnterSurfaceRuleReference(c *SurfaceRuleReferenceContext)

	// EnterSurfaceRuleDeclaration is called when entering the surfaceRuleDeclaration production.
	EnterSurfaceRuleDeclaration(c *SurfaceRuleDeclarationContext)

	// EnterSurfaceRule is called when entering the surfaceRule production.
	EnterSurfaceRule(c *SurfaceRuleContext)

	// EnterSurfaceRule_Conditional is called when entering the surfaceRule_Conditional production.
	EnterSurfaceRule_Conditional(c *SurfaceRule_ConditionalContext)

	// EnterSurfaceRule_Bandlands is called when entering the surfaceRule_Bandlands production.
	EnterSurfaceRule_Bandlands(c *SurfaceRule_BandlandsContext)

	// EnterSurfaceRule_Block is called when entering the surfaceRule_Block production.
	EnterSurfaceRule_Block(c *SurfaceRule_BlockContext)

	// EnterSurfaceRule_Sequence is called when entering the surfaceRule_Sequence production.
	EnterSurfaceRule_Sequence(c *SurfaceRule_SequenceContext)

	// EnterSurfaceConditionReference is called when entering the surfaceConditionReference production.
	EnterSurfaceConditionReference(c *SurfaceConditionReferenceContext)

	// EnterSurfaceConditionDeclaration is called when entering the surfaceConditionDeclaration production.
	EnterSurfaceConditionDeclaration(c *SurfaceConditionDeclarationContext)

	// EnterSurfaceCondition is called when entering the surfaceCondition production.
	EnterSurfaceCondition(c *SurfaceConditionContext)

	// EnterSurfaceCondition_AboveSurface is called when entering the surfaceCondition_AboveSurface production.
	EnterSurfaceCondition_AboveSurface(c *SurfaceCondition_AboveSurfaceContext)

	// EnterSurfaceCondition_Biome is called when entering the surfaceCondition_Biome production.
	EnterSurfaceCondition_Biome(c *SurfaceCondition_BiomeContext)

	// EnterSurfaceCondition_Hole is called when entering the surfaceCondition_Hole production.
	EnterSurfaceCondition_Hole(c *SurfaceCondition_HoleContext)

	// EnterSurfaceCondition_Noise is called when entering the surfaceCondition_Noise production.
	EnterSurfaceCondition_Noise(c *SurfaceCondition_NoiseContext)

	// EnterSurfaceCondition_Steep is called when entering the surfaceCondition_Steep production.
	EnterSurfaceCondition_Steep(c *SurfaceCondition_SteepContext)

	// EnterSurfaceCondition_StoneDepth is called when entering the surfaceCondition_StoneDepth production.
	EnterSurfaceCondition_StoneDepth(c *SurfaceCondition_StoneDepthContext)

	// EnterSurfaceCondition_Freezing is called when entering the surfaceCondition_Freezing production.
	EnterSurfaceCondition_Freezing(c *SurfaceCondition_FreezingContext)

	// EnterSurfaceCondition_VerticalGradient is called when entering the surfaceCondition_VerticalGradient production.
	EnterSurfaceCondition_VerticalGradient(c *SurfaceCondition_VerticalGradientContext)

	// EnterSurfaceCondition_AboveWater is called when entering the surfaceCondition_AboveWater production.
	EnterSurfaceCondition_AboveWater(c *SurfaceCondition_AboveWaterContext)

	// EnterSurfaceCondition_YAbove is called when entering the surfaceCondition_YAbove production.
	EnterSurfaceCondition_YAbove(c *SurfaceCondition_YAboveContext)

	// EnterSurfaceCondition_Compound__Item is called when entering the surfaceCondition_Compound__Item production.
	EnterSurfaceCondition_Compound__Item(c *SurfaceCondition_Compound__ItemContext)

	// EnterSurfaceCondition_Compound is called when entering the surfaceCondition_Compound production.
	EnterSurfaceCondition_Compound(c *SurfaceCondition_CompoundContext)

	// EnterReference is called when entering the reference production.
	EnterReference(c *ReferenceContext)

	// EnterResourceReference is called when entering the resourceReference production.
	EnterResourceReference(c *ResourceReferenceContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterKeyword is called when entering the keyword production.
	EnterKeyword(c *KeywordContext)

	// EnterVerticalAnchor is called when entering the verticalAnchor production.
	EnterVerticalAnchor(c *VerticalAnchorContext)

	// EnterVerticalAnchor_Absolute is called when entering the verticalAnchor_Absolute production.
	EnterVerticalAnchor_Absolute(c *VerticalAnchor_AbsoluteContext)

	// EnterVerticalAnchor_AboveBottom is called when entering the verticalAnchor_AboveBottom production.
	EnterVerticalAnchor_AboveBottom(c *VerticalAnchor_AboveBottomContext)

	// EnterVerticalAnchor_BelowTop is called when entering the verticalAnchor_BelowTop production.
	EnterVerticalAnchor_BelowTop(c *VerticalAnchor_BelowTopContext)

	// EnterWorldgenDeclaration is called when entering the worldgenDeclaration production.
	EnterWorldgenDeclaration(c *WorldgenDeclarationContext)

	// EnterNoiseDeclaration is called when entering the noiseDeclaration production.
	EnterNoiseDeclaration(c *NoiseDeclarationContext)

	// EnterNoiseDefinition is called when entering the noiseDefinition production.
	EnterNoiseDefinition(c *NoiseDefinitionContext)

	// ExitNamespaceDeclaration is called when exiting the namespaceDeclaration production.
	ExitNamespaceDeclaration(c *NamespaceDeclarationContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitMmsFile is called when exiting the mmsFile production.
	ExitMmsFile(c *MmsFileContext)

	// ExitSurfaceDeclaration is called when exiting the surfaceDeclaration production.
	ExitSurfaceDeclaration(c *SurfaceDeclarationContext)

	// ExitSurfaceDefinition is called when exiting the surfaceDefinition production.
	ExitSurfaceDefinition(c *SurfaceDefinitionContext)

	// ExitSurfaceRuleReference is called when exiting the surfaceRuleReference production.
	ExitSurfaceRuleReference(c *SurfaceRuleReferenceContext)

	// ExitSurfaceRuleDeclaration is called when exiting the surfaceRuleDeclaration production.
	ExitSurfaceRuleDeclaration(c *SurfaceRuleDeclarationContext)

	// ExitSurfaceRule is called when exiting the surfaceRule production.
	ExitSurfaceRule(c *SurfaceRuleContext)

	// ExitSurfaceRule_Conditional is called when exiting the surfaceRule_Conditional production.
	ExitSurfaceRule_Conditional(c *SurfaceRule_ConditionalContext)

	// ExitSurfaceRule_Bandlands is called when exiting the surfaceRule_Bandlands production.
	ExitSurfaceRule_Bandlands(c *SurfaceRule_BandlandsContext)

	// ExitSurfaceRule_Block is called when exiting the surfaceRule_Block production.
	ExitSurfaceRule_Block(c *SurfaceRule_BlockContext)

	// ExitSurfaceRule_Sequence is called when exiting the surfaceRule_Sequence production.
	ExitSurfaceRule_Sequence(c *SurfaceRule_SequenceContext)

	// ExitSurfaceConditionReference is called when exiting the surfaceConditionReference production.
	ExitSurfaceConditionReference(c *SurfaceConditionReferenceContext)

	// ExitSurfaceConditionDeclaration is called when exiting the surfaceConditionDeclaration production.
	ExitSurfaceConditionDeclaration(c *SurfaceConditionDeclarationContext)

	// ExitSurfaceCondition is called when exiting the surfaceCondition production.
	ExitSurfaceCondition(c *SurfaceConditionContext)

	// ExitSurfaceCondition_AboveSurface is called when exiting the surfaceCondition_AboveSurface production.
	ExitSurfaceCondition_AboveSurface(c *SurfaceCondition_AboveSurfaceContext)

	// ExitSurfaceCondition_Biome is called when exiting the surfaceCondition_Biome production.
	ExitSurfaceCondition_Biome(c *SurfaceCondition_BiomeContext)

	// ExitSurfaceCondition_Hole is called when exiting the surfaceCondition_Hole production.
	ExitSurfaceCondition_Hole(c *SurfaceCondition_HoleContext)

	// ExitSurfaceCondition_Noise is called when exiting the surfaceCondition_Noise production.
	ExitSurfaceCondition_Noise(c *SurfaceCondition_NoiseContext)

	// ExitSurfaceCondition_Steep is called when exiting the surfaceCondition_Steep production.
	ExitSurfaceCondition_Steep(c *SurfaceCondition_SteepContext)

	// ExitSurfaceCondition_StoneDepth is called when exiting the surfaceCondition_StoneDepth production.
	ExitSurfaceCondition_StoneDepth(c *SurfaceCondition_StoneDepthContext)

	// ExitSurfaceCondition_Freezing is called when exiting the surfaceCondition_Freezing production.
	ExitSurfaceCondition_Freezing(c *SurfaceCondition_FreezingContext)

	// ExitSurfaceCondition_VerticalGradient is called when exiting the surfaceCondition_VerticalGradient production.
	ExitSurfaceCondition_VerticalGradient(c *SurfaceCondition_VerticalGradientContext)

	// ExitSurfaceCondition_AboveWater is called when exiting the surfaceCondition_AboveWater production.
	ExitSurfaceCondition_AboveWater(c *SurfaceCondition_AboveWaterContext)

	// ExitSurfaceCondition_YAbove is called when exiting the surfaceCondition_YAbove production.
	ExitSurfaceCondition_YAbove(c *SurfaceCondition_YAboveContext)

	// ExitSurfaceCondition_Compound__Item is called when exiting the surfaceCondition_Compound__Item production.
	ExitSurfaceCondition_Compound__Item(c *SurfaceCondition_Compound__ItemContext)

	// ExitSurfaceCondition_Compound is called when exiting the surfaceCondition_Compound production.
	ExitSurfaceCondition_Compound(c *SurfaceCondition_CompoundContext)

	// ExitReference is called when exiting the reference production.
	ExitReference(c *ReferenceContext)

	// ExitResourceReference is called when exiting the resourceReference production.
	ExitResourceReference(c *ResourceReferenceContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitKeyword is called when exiting the keyword production.
	ExitKeyword(c *KeywordContext)

	// ExitVerticalAnchor is called when exiting the verticalAnchor production.
	ExitVerticalAnchor(c *VerticalAnchorContext)

	// ExitVerticalAnchor_Absolute is called when exiting the verticalAnchor_Absolute production.
	ExitVerticalAnchor_Absolute(c *VerticalAnchor_AbsoluteContext)

	// ExitVerticalAnchor_AboveBottom is called when exiting the verticalAnchor_AboveBottom production.
	ExitVerticalAnchor_AboveBottom(c *VerticalAnchor_AboveBottomContext)

	// ExitVerticalAnchor_BelowTop is called when exiting the verticalAnchor_BelowTop production.
	ExitVerticalAnchor_BelowTop(c *VerticalAnchor_BelowTopContext)

	// ExitWorldgenDeclaration is called when exiting the worldgenDeclaration production.
	ExitWorldgenDeclaration(c *WorldgenDeclarationContext)

	// ExitNoiseDeclaration is called when exiting the noiseDeclaration production.
	ExitNoiseDeclaration(c *NoiseDeclarationContext)

	// ExitNoiseDefinition is called when exiting the noiseDefinition production.
	ExitNoiseDefinition(c *NoiseDefinitionContext)
}
