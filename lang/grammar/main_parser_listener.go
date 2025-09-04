// Code generated from ./grammar/Main_Parser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // Main_Parser
import "github.com/antlr4-go/antlr/v4"

// Main_ParserListener is a complete listener for a parse tree produced by Main_Parser.
type Main_ParserListener interface {
	antlr.ParseTreeListener

	// EnterScript is called when entering the script production.
	EnterScript(c *ScriptContext)

	// EnterNamespaceDefinition is called when entering the namespaceDefinition production.
	EnterNamespaceDefinition(c *NamespaceDefinitionContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterDefinition is called when entering the definition production.
	EnterDefinition(c *DefinitionContext)

	// EnterSurfaceRuleDefinition is called when entering the surfaceRuleDefinition production.
	EnterSurfaceRuleDefinition(c *SurfaceRuleDefinitionContext)

	// EnterSurfaceRule is called when entering the surfaceRule production.
	EnterSurfaceRule(c *SurfaceRuleContext)

	// EnterSurfaceRule_Block is called when entering the surfaceRule_Block production.
	EnterSurfaceRule_Block(c *SurfaceRule_BlockContext)

	// EnterSurfaceRule_Bandlands is called when entering the surfaceRule_Bandlands production.
	EnterSurfaceRule_Bandlands(c *SurfaceRule_BandlandsContext)

	// EnterSurfaceRule_Condition is called when entering the surfaceRule_Condition production.
	EnterSurfaceRule_Condition(c *SurfaceRule_ConditionContext)

	// EnterSurfaceRule_Sequence is called when entering the surfaceRule_Sequence production.
	EnterSurfaceRule_Sequence(c *SurfaceRule_SequenceContext)

	// EnterSurfaceRule_Reference is called when entering the surfaceRule_Reference production.
	EnterSurfaceRule_Reference(c *SurfaceRule_ReferenceContext)

	// EnterSurfaceConditionDefinition is called when entering the surfaceConditionDefinition production.
	EnterSurfaceConditionDefinition(c *SurfaceConditionDefinitionContext)

	// EnterSurfaceCondition is called when entering the surfaceCondition production.
	EnterSurfaceCondition(c *SurfaceConditionContext)

	// EnterSurfaceCondition_Reference is called when entering the surfaceCondition_Reference production.
	EnterSurfaceCondition_Reference(c *SurfaceCondition_ReferenceContext)

	// EnterSurfaceCondition_And is called when entering the surfaceCondition_And production.
	EnterSurfaceCondition_And(c *SurfaceCondition_AndContext)

	// EnterSurfaceCondition_Or is called when entering the surfaceCondition_Or production.
	EnterSurfaceCondition_Or(c *SurfaceCondition_OrContext)

	// EnterSurfaceCondition_Not is called when entering the surfaceCondition_Not production.
	EnterSurfaceCondition_Not(c *SurfaceCondition_NotContext)

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

	// EnterResourceReference is called when entering the resourceReference production.
	EnterResourceReference(c *ResourceReferenceContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterVerticalAnchorDefinition is called when entering the verticalAnchorDefinition production.
	EnterVerticalAnchorDefinition(c *VerticalAnchorDefinitionContext)

	// EnterVerticalAnchor is called when entering the verticalAnchor production.
	EnterVerticalAnchor(c *VerticalAnchorContext)

	// EnterVerticalAnchor_Absolute is called when entering the verticalAnchor_Absolute production.
	EnterVerticalAnchor_Absolute(c *VerticalAnchor_AbsoluteContext)

	// EnterVerticalAnchor_AboveBottom is called when entering the verticalAnchor_AboveBottom production.
	EnterVerticalAnchor_AboveBottom(c *VerticalAnchor_AboveBottomContext)

	// EnterVerticalAnchor_BelowTop is called when entering the verticalAnchor_BelowTop production.
	EnterVerticalAnchor_BelowTop(c *VerticalAnchor_BelowTopContext)

	// EnterVerticalAnchor_Reference is called when entering the verticalAnchor_Reference production.
	EnterVerticalAnchor_Reference(c *VerticalAnchor_ReferenceContext)

	// ExitScript is called when exiting the script production.
	ExitScript(c *ScriptContext)

	// ExitNamespaceDefinition is called when exiting the namespaceDefinition production.
	ExitNamespaceDefinition(c *NamespaceDefinitionContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitDefinition is called when exiting the definition production.
	ExitDefinition(c *DefinitionContext)

	// ExitSurfaceRuleDefinition is called when exiting the surfaceRuleDefinition production.
	ExitSurfaceRuleDefinition(c *SurfaceRuleDefinitionContext)

	// ExitSurfaceRule is called when exiting the surfaceRule production.
	ExitSurfaceRule(c *SurfaceRuleContext)

	// ExitSurfaceRule_Block is called when exiting the surfaceRule_Block production.
	ExitSurfaceRule_Block(c *SurfaceRule_BlockContext)

	// ExitSurfaceRule_Bandlands is called when exiting the surfaceRule_Bandlands production.
	ExitSurfaceRule_Bandlands(c *SurfaceRule_BandlandsContext)

	// ExitSurfaceRule_Condition is called when exiting the surfaceRule_Condition production.
	ExitSurfaceRule_Condition(c *SurfaceRule_ConditionContext)

	// ExitSurfaceRule_Sequence is called when exiting the surfaceRule_Sequence production.
	ExitSurfaceRule_Sequence(c *SurfaceRule_SequenceContext)

	// ExitSurfaceRule_Reference is called when exiting the surfaceRule_Reference production.
	ExitSurfaceRule_Reference(c *SurfaceRule_ReferenceContext)

	// ExitSurfaceConditionDefinition is called when exiting the surfaceConditionDefinition production.
	ExitSurfaceConditionDefinition(c *SurfaceConditionDefinitionContext)

	// ExitSurfaceCondition is called when exiting the surfaceCondition production.
	ExitSurfaceCondition(c *SurfaceConditionContext)

	// ExitSurfaceCondition_Reference is called when exiting the surfaceCondition_Reference production.
	ExitSurfaceCondition_Reference(c *SurfaceCondition_ReferenceContext)

	// ExitSurfaceCondition_And is called when exiting the surfaceCondition_And production.
	ExitSurfaceCondition_And(c *SurfaceCondition_AndContext)

	// ExitSurfaceCondition_Or is called when exiting the surfaceCondition_Or production.
	ExitSurfaceCondition_Or(c *SurfaceCondition_OrContext)

	// ExitSurfaceCondition_Not is called when exiting the surfaceCondition_Not production.
	ExitSurfaceCondition_Not(c *SurfaceCondition_NotContext)

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

	// ExitResourceReference is called when exiting the resourceReference production.
	ExitResourceReference(c *ResourceReferenceContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitVerticalAnchorDefinition is called when exiting the verticalAnchorDefinition production.
	ExitVerticalAnchorDefinition(c *VerticalAnchorDefinitionContext)

	// ExitVerticalAnchor is called when exiting the verticalAnchor production.
	ExitVerticalAnchor(c *VerticalAnchorContext)

	// ExitVerticalAnchor_Absolute is called when exiting the verticalAnchor_Absolute production.
	ExitVerticalAnchor_Absolute(c *VerticalAnchor_AbsoluteContext)

	// ExitVerticalAnchor_AboveBottom is called when exiting the verticalAnchor_AboveBottom production.
	ExitVerticalAnchor_AboveBottom(c *VerticalAnchor_AboveBottomContext)

	// ExitVerticalAnchor_BelowTop is called when exiting the verticalAnchor_BelowTop production.
	ExitVerticalAnchor_BelowTop(c *VerticalAnchor_BelowTopContext)

	// ExitVerticalAnchor_Reference is called when exiting the verticalAnchor_Reference production.
	ExitVerticalAnchor_Reference(c *VerticalAnchor_ReferenceContext)
}
