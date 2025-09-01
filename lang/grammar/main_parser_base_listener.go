// Code generated from ./grammar/Main_Parser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // Main_Parser
import "github.com/antlr4-go/antlr/v4"

// BaseMain_ParserListener is a complete listener for a parse tree produced by Main_Parser.
type BaseMain_ParserListener struct{}

var _ Main_ParserListener = &BaseMain_ParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMain_ParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMain_ParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMain_ParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMain_ParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterScript is called when production script is entered.
func (s *BaseMain_ParserListener) EnterScript(ctx *ScriptContext) {}

// ExitScript is called when production script is exited.
func (s *BaseMain_ParserListener) ExitScript(ctx *ScriptContext) {}

// EnterNamespaceDefinition is called when production namespaceDefinition is entered.
func (s *BaseMain_ParserListener) EnterNamespaceDefinition(ctx *NamespaceDefinitionContext) {}

// ExitNamespaceDefinition is called when production namespaceDefinition is exited.
func (s *BaseMain_ParserListener) ExitNamespaceDefinition(ctx *NamespaceDefinitionContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseMain_ParserListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseMain_ParserListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterDefinition is called when production definition is entered.
func (s *BaseMain_ParserListener) EnterDefinition(ctx *DefinitionContext) {}

// ExitDefinition is called when production definition is exited.
func (s *BaseMain_ParserListener) ExitDefinition(ctx *DefinitionContext) {}

// EnterSurfaceRuleDefinition is called when production surfaceRuleDefinition is entered.
func (s *BaseMain_ParserListener) EnterSurfaceRuleDefinition(ctx *SurfaceRuleDefinitionContext) {}

// ExitSurfaceRuleDefinition is called when production surfaceRuleDefinition is exited.
func (s *BaseMain_ParserListener) ExitSurfaceRuleDefinition(ctx *SurfaceRuleDefinitionContext) {}

// EnterSurfaceRule is called when production surfaceRule is entered.
func (s *BaseMain_ParserListener) EnterSurfaceRule(ctx *SurfaceRuleContext) {}

// ExitSurfaceRule is called when production surfaceRule is exited.
func (s *BaseMain_ParserListener) ExitSurfaceRule(ctx *SurfaceRuleContext) {}

// EnterSurfaceRule_Block is called when production surfaceRule_Block is entered.
func (s *BaseMain_ParserListener) EnterSurfaceRule_Block(ctx *SurfaceRule_BlockContext) {}

// ExitSurfaceRule_Block is called when production surfaceRule_Block is exited.
func (s *BaseMain_ParserListener) ExitSurfaceRule_Block(ctx *SurfaceRule_BlockContext) {}

// EnterSurfaceRule_Bandlands is called when production surfaceRule_Bandlands is entered.
func (s *BaseMain_ParserListener) EnterSurfaceRule_Bandlands(ctx *SurfaceRule_BandlandsContext) {}

// ExitSurfaceRule_Bandlands is called when production surfaceRule_Bandlands is exited.
func (s *BaseMain_ParserListener) ExitSurfaceRule_Bandlands(ctx *SurfaceRule_BandlandsContext) {}

// EnterSurfaceRule_Condition is called when production surfaceRule_Condition is entered.
func (s *BaseMain_ParserListener) EnterSurfaceRule_Condition(ctx *SurfaceRule_ConditionContext) {}

// ExitSurfaceRule_Condition is called when production surfaceRule_Condition is exited.
func (s *BaseMain_ParserListener) ExitSurfaceRule_Condition(ctx *SurfaceRule_ConditionContext) {}

// EnterSurfaceRule_Sequence is called when production surfaceRule_Sequence is entered.
func (s *BaseMain_ParserListener) EnterSurfaceRule_Sequence(ctx *SurfaceRule_SequenceContext) {}

// ExitSurfaceRule_Sequence is called when production surfaceRule_Sequence is exited.
func (s *BaseMain_ParserListener) ExitSurfaceRule_Sequence(ctx *SurfaceRule_SequenceContext) {}

// EnterSurfaceRule_Reference is called when production surfaceRule_Reference is entered.
func (s *BaseMain_ParserListener) EnterSurfaceRule_Reference(ctx *SurfaceRule_ReferenceContext) {}

// ExitSurfaceRule_Reference is called when production surfaceRule_Reference is exited.
func (s *BaseMain_ParserListener) ExitSurfaceRule_Reference(ctx *SurfaceRule_ReferenceContext) {}

// EnterSurfaceConditionDefinition is called when production surfaceConditionDefinition is entered.
func (s *BaseMain_ParserListener) EnterSurfaceConditionDefinition(ctx *SurfaceConditionDefinitionContext) {
}

// ExitSurfaceConditionDefinition is called when production surfaceConditionDefinition is exited.
func (s *BaseMain_ParserListener) ExitSurfaceConditionDefinition(ctx *SurfaceConditionDefinitionContext) {
}

// EnterSurfaceCondition is called when production surfaceCondition is entered.
func (s *BaseMain_ParserListener) EnterSurfaceCondition(ctx *SurfaceConditionContext) {}

// ExitSurfaceCondition is called when production surfaceCondition is exited.
func (s *BaseMain_ParserListener) ExitSurfaceCondition(ctx *SurfaceConditionContext) {}

// EnterSurfaceCondition_Reference is called when production surfaceCondition_Reference is entered.
func (s *BaseMain_ParserListener) EnterSurfaceCondition_Reference(ctx *SurfaceCondition_ReferenceContext) {
}

// ExitSurfaceCondition_Reference is called when production surfaceCondition_Reference is exited.
func (s *BaseMain_ParserListener) ExitSurfaceCondition_Reference(ctx *SurfaceCondition_ReferenceContext) {
}

// EnterSurfaceCondition_And is called when production surfaceCondition_And is entered.
func (s *BaseMain_ParserListener) EnterSurfaceCondition_And(ctx *SurfaceCondition_AndContext) {}

// ExitSurfaceCondition_And is called when production surfaceCondition_And is exited.
func (s *BaseMain_ParserListener) ExitSurfaceCondition_And(ctx *SurfaceCondition_AndContext) {}

// EnterSurfaceCondition_Or is called when production surfaceCondition_Or is entered.
func (s *BaseMain_ParserListener) EnterSurfaceCondition_Or(ctx *SurfaceCondition_OrContext) {}

// ExitSurfaceCondition_Or is called when production surfaceCondition_Or is exited.
func (s *BaseMain_ParserListener) ExitSurfaceCondition_Or(ctx *SurfaceCondition_OrContext) {}

// EnterSurfaceCondition_AboveSurface is called when production surfaceCondition_AboveSurface is entered.
func (s *BaseMain_ParserListener) EnterSurfaceCondition_AboveSurface(ctx *SurfaceCondition_AboveSurfaceContext) {
}

// ExitSurfaceCondition_AboveSurface is called when production surfaceCondition_AboveSurface is exited.
func (s *BaseMain_ParserListener) ExitSurfaceCondition_AboveSurface(ctx *SurfaceCondition_AboveSurfaceContext) {
}

// EnterSurfaceCondition_Biome is called when production surfaceCondition_Biome is entered.
func (s *BaseMain_ParserListener) EnterSurfaceCondition_Biome(ctx *SurfaceCondition_BiomeContext) {}

// ExitSurfaceCondition_Biome is called when production surfaceCondition_Biome is exited.
func (s *BaseMain_ParserListener) ExitSurfaceCondition_Biome(ctx *SurfaceCondition_BiomeContext) {}

// EnterSurfaceCondition_Hole is called when production surfaceCondition_Hole is entered.
func (s *BaseMain_ParserListener) EnterSurfaceCondition_Hole(ctx *SurfaceCondition_HoleContext) {}

// ExitSurfaceCondition_Hole is called when production surfaceCondition_Hole is exited.
func (s *BaseMain_ParserListener) ExitSurfaceCondition_Hole(ctx *SurfaceCondition_HoleContext) {}

// EnterSurfaceCondition_Noise is called when production surfaceCondition_Noise is entered.
func (s *BaseMain_ParserListener) EnterSurfaceCondition_Noise(ctx *SurfaceCondition_NoiseContext) {}

// ExitSurfaceCondition_Noise is called when production surfaceCondition_Noise is exited.
func (s *BaseMain_ParserListener) ExitSurfaceCondition_Noise(ctx *SurfaceCondition_NoiseContext) {}

// EnterSurfaceCondition_Steep is called when production surfaceCondition_Steep is entered.
func (s *BaseMain_ParserListener) EnterSurfaceCondition_Steep(ctx *SurfaceCondition_SteepContext) {}

// ExitSurfaceCondition_Steep is called when production surfaceCondition_Steep is exited.
func (s *BaseMain_ParserListener) ExitSurfaceCondition_Steep(ctx *SurfaceCondition_SteepContext) {}

// EnterSurfaceCondition_StoneDepth is called when production surfaceCondition_StoneDepth is entered.
func (s *BaseMain_ParserListener) EnterSurfaceCondition_StoneDepth(ctx *SurfaceCondition_StoneDepthContext) {
}

// ExitSurfaceCondition_StoneDepth is called when production surfaceCondition_StoneDepth is exited.
func (s *BaseMain_ParserListener) ExitSurfaceCondition_StoneDepth(ctx *SurfaceCondition_StoneDepthContext) {
}

// EnterSurfaceCondition_Freezing is called when production surfaceCondition_Freezing is entered.
func (s *BaseMain_ParserListener) EnterSurfaceCondition_Freezing(ctx *SurfaceCondition_FreezingContext) {
}

// ExitSurfaceCondition_Freezing is called when production surfaceCondition_Freezing is exited.
func (s *BaseMain_ParserListener) ExitSurfaceCondition_Freezing(ctx *SurfaceCondition_FreezingContext) {
}

// EnterSurfaceCondition_VerticalGradient is called when production surfaceCondition_VerticalGradient is entered.
func (s *BaseMain_ParserListener) EnterSurfaceCondition_VerticalGradient(ctx *SurfaceCondition_VerticalGradientContext) {
}

// ExitSurfaceCondition_VerticalGradient is called when production surfaceCondition_VerticalGradient is exited.
func (s *BaseMain_ParserListener) ExitSurfaceCondition_VerticalGradient(ctx *SurfaceCondition_VerticalGradientContext) {
}

// EnterSurfaceCondition_AboveWater is called when production surfaceCondition_AboveWater is entered.
func (s *BaseMain_ParserListener) EnterSurfaceCondition_AboveWater(ctx *SurfaceCondition_AboveWaterContext) {
}

// ExitSurfaceCondition_AboveWater is called when production surfaceCondition_AboveWater is exited.
func (s *BaseMain_ParserListener) ExitSurfaceCondition_AboveWater(ctx *SurfaceCondition_AboveWaterContext) {
}

// EnterSurfaceCondition_YAbove is called when production surfaceCondition_YAbove is entered.
func (s *BaseMain_ParserListener) EnterSurfaceCondition_YAbove(ctx *SurfaceCondition_YAboveContext) {}

// ExitSurfaceCondition_YAbove is called when production surfaceCondition_YAbove is exited.
func (s *BaseMain_ParserListener) ExitSurfaceCondition_YAbove(ctx *SurfaceCondition_YAboveContext) {}

// EnterResourceReference is called when production resourceReference is entered.
func (s *BaseMain_ParserListener) EnterResourceReference(ctx *ResourceReferenceContext) {}

// ExitResourceReference is called when production resourceReference is exited.
func (s *BaseMain_ParserListener) ExitResourceReference(ctx *ResourceReferenceContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseMain_ParserListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseMain_ParserListener) ExitNumber(ctx *NumberContext) {}

// EnterVerticalAnchorDefinition is called when production verticalAnchorDefinition is entered.
func (s *BaseMain_ParserListener) EnterVerticalAnchorDefinition(ctx *VerticalAnchorDefinitionContext) {
}

// ExitVerticalAnchorDefinition is called when production verticalAnchorDefinition is exited.
func (s *BaseMain_ParserListener) ExitVerticalAnchorDefinition(ctx *VerticalAnchorDefinitionContext) {
}

// EnterVerticalAnchor_Absolute is called when production verticalAnchor_Absolute is entered.
func (s *BaseMain_ParserListener) EnterVerticalAnchor_Absolute(ctx *VerticalAnchor_AbsoluteContext) {}

// ExitVerticalAnchor_Absolute is called when production verticalAnchor_Absolute is exited.
func (s *BaseMain_ParserListener) ExitVerticalAnchor_Absolute(ctx *VerticalAnchor_AbsoluteContext) {}

// EnterVerticalAnchor_AboveBottom is called when production verticalAnchor_AboveBottom is entered.
func (s *BaseMain_ParserListener) EnterVerticalAnchor_AboveBottom(ctx *VerticalAnchor_AboveBottomContext) {
}

// ExitVerticalAnchor_AboveBottom is called when production verticalAnchor_AboveBottom is exited.
func (s *BaseMain_ParserListener) ExitVerticalAnchor_AboveBottom(ctx *VerticalAnchor_AboveBottomContext) {
}

// EnterVerticalAnchor_BelowTop is called when production verticalAnchor_BelowTop is entered.
func (s *BaseMain_ParserListener) EnterVerticalAnchor_BelowTop(ctx *VerticalAnchor_BelowTopContext) {}

// ExitVerticalAnchor_BelowTop is called when production verticalAnchor_BelowTop is exited.
func (s *BaseMain_ParserListener) ExitVerticalAnchor_BelowTop(ctx *VerticalAnchor_BelowTopContext) {}

// EnterVerticalAnchor_Reference is called when production verticalAnchor_Reference is entered.
func (s *BaseMain_ParserListener) EnterVerticalAnchor_Reference(ctx *VerticalAnchor_ReferenceContext) {
}

// ExitVerticalAnchor_Reference is called when production verticalAnchor_Reference is exited.
func (s *BaseMain_ParserListener) ExitVerticalAnchor_Reference(ctx *VerticalAnchor_ReferenceContext) {
}
