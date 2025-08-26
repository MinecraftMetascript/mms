// Code generated from ./grammars/MMSParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammars // MMSParser
import "github.com/antlr4-go/antlr/v4"

// BaseMMSParserListener is a complete listener for a parse tree produced by MMSParser.
type BaseMMSParserListener struct{}

var _ MMSParserListener = &BaseMMSParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMMSParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMMSParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMMSParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMMSParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterNamespaceDeclaration is called when production namespaceDeclaration is entered.
func (s *BaseMMSParserListener) EnterNamespaceDeclaration(ctx *NamespaceDeclarationContext) {}

// ExitNamespaceDeclaration is called when production namespaceDeclaration is exited.
func (s *BaseMMSParserListener) ExitNamespaceDeclaration(ctx *NamespaceDeclarationContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseMMSParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseMMSParserListener) ExitStatement(ctx *StatementContext) {}

// EnterMmsFile is called when production mmsFile is entered.
func (s *BaseMMSParserListener) EnterMmsFile(ctx *MmsFileContext) {}

// ExitMmsFile is called when production mmsFile is exited.
func (s *BaseMMSParserListener) ExitMmsFile(ctx *MmsFileContext) {}

// EnterSurfaceDeclaration is called when production surfaceDeclaration is entered.
func (s *BaseMMSParserListener) EnterSurfaceDeclaration(ctx *SurfaceDeclarationContext) {}

// ExitSurfaceDeclaration is called when production surfaceDeclaration is exited.
func (s *BaseMMSParserListener) ExitSurfaceDeclaration(ctx *SurfaceDeclarationContext) {}

// EnterSurfaceDefinition is called when production surfaceDefinition is entered.
func (s *BaseMMSParserListener) EnterSurfaceDefinition(ctx *SurfaceDefinitionContext) {}

// ExitSurfaceDefinition is called when production surfaceDefinition is exited.
func (s *BaseMMSParserListener) ExitSurfaceDefinition(ctx *SurfaceDefinitionContext) {}

// EnterSurfaceRuleReference is called when production surfaceRuleReference is entered.
func (s *BaseMMSParserListener) EnterSurfaceRuleReference(ctx *SurfaceRuleReferenceContext) {}

// ExitSurfaceRuleReference is called when production surfaceRuleReference is exited.
func (s *BaseMMSParserListener) ExitSurfaceRuleReference(ctx *SurfaceRuleReferenceContext) {}

// EnterSurfaceRuleDeclaration is called when production surfaceRuleDeclaration is entered.
func (s *BaseMMSParserListener) EnterSurfaceRuleDeclaration(ctx *SurfaceRuleDeclarationContext) {}

// ExitSurfaceRuleDeclaration is called when production surfaceRuleDeclaration is exited.
func (s *BaseMMSParserListener) ExitSurfaceRuleDeclaration(ctx *SurfaceRuleDeclarationContext) {}

// EnterSurfaceRule is called when production surfaceRule is entered.
func (s *BaseMMSParserListener) EnterSurfaceRule(ctx *SurfaceRuleContext) {}

// ExitSurfaceRule is called when production surfaceRule is exited.
func (s *BaseMMSParserListener) ExitSurfaceRule(ctx *SurfaceRuleContext) {}

// EnterSurfaceRule_Conditional is called when production surfaceRule_Conditional is entered.
func (s *BaseMMSParserListener) EnterSurfaceRule_Conditional(ctx *SurfaceRule_ConditionalContext) {}

// ExitSurfaceRule_Conditional is called when production surfaceRule_Conditional is exited.
func (s *BaseMMSParserListener) ExitSurfaceRule_Conditional(ctx *SurfaceRule_ConditionalContext) {}

// EnterSurfaceRule_Bandlands is called when production surfaceRule_Bandlands is entered.
func (s *BaseMMSParserListener) EnterSurfaceRule_Bandlands(ctx *SurfaceRule_BandlandsContext) {}

// ExitSurfaceRule_Bandlands is called when production surfaceRule_Bandlands is exited.
func (s *BaseMMSParserListener) ExitSurfaceRule_Bandlands(ctx *SurfaceRule_BandlandsContext) {}

// EnterSurfaceRule_Block is called when production surfaceRule_Block is entered.
func (s *BaseMMSParserListener) EnterSurfaceRule_Block(ctx *SurfaceRule_BlockContext) {}

// ExitSurfaceRule_Block is called when production surfaceRule_Block is exited.
func (s *BaseMMSParserListener) ExitSurfaceRule_Block(ctx *SurfaceRule_BlockContext) {}

// EnterSurfaceRule_Sequence is called when production surfaceRule_Sequence is entered.
func (s *BaseMMSParserListener) EnterSurfaceRule_Sequence(ctx *SurfaceRule_SequenceContext) {}

// ExitSurfaceRule_Sequence is called when production surfaceRule_Sequence is exited.
func (s *BaseMMSParserListener) ExitSurfaceRule_Sequence(ctx *SurfaceRule_SequenceContext) {}

// EnterSurfaceConditionReference is called when production surfaceConditionReference is entered.
func (s *BaseMMSParserListener) EnterSurfaceConditionReference(ctx *SurfaceConditionReferenceContext) {
}

// ExitSurfaceConditionReference is called when production surfaceConditionReference is exited.
func (s *BaseMMSParserListener) ExitSurfaceConditionReference(ctx *SurfaceConditionReferenceContext) {
}

// EnterSurfaceConditionDeclaration is called when production surfaceConditionDeclaration is entered.
func (s *BaseMMSParserListener) EnterSurfaceConditionDeclaration(ctx *SurfaceConditionDeclarationContext) {
}

// ExitSurfaceConditionDeclaration is called when production surfaceConditionDeclaration is exited.
func (s *BaseMMSParserListener) ExitSurfaceConditionDeclaration(ctx *SurfaceConditionDeclarationContext) {
}

// EnterSurfaceCondition is called when production surfaceCondition is entered.
func (s *BaseMMSParserListener) EnterSurfaceCondition(ctx *SurfaceConditionContext) {}

// ExitSurfaceCondition is called when production surfaceCondition is exited.
func (s *BaseMMSParserListener) ExitSurfaceCondition(ctx *SurfaceConditionContext) {}

// EnterSurfaceCondition_AboveSurface is called when production surfaceCondition_AboveSurface is entered.
func (s *BaseMMSParserListener) EnterSurfaceCondition_AboveSurface(ctx *SurfaceCondition_AboveSurfaceContext) {
}

// ExitSurfaceCondition_AboveSurface is called when production surfaceCondition_AboveSurface is exited.
func (s *BaseMMSParserListener) ExitSurfaceCondition_AboveSurface(ctx *SurfaceCondition_AboveSurfaceContext) {
}

// EnterSurfaceCondition_Biome is called when production surfaceCondition_Biome is entered.
func (s *BaseMMSParserListener) EnterSurfaceCondition_Biome(ctx *SurfaceCondition_BiomeContext) {}

// ExitSurfaceCondition_Biome is called when production surfaceCondition_Biome is exited.
func (s *BaseMMSParserListener) ExitSurfaceCondition_Biome(ctx *SurfaceCondition_BiomeContext) {}

// EnterSurfaceCondition_Hole is called when production surfaceCondition_Hole is entered.
func (s *BaseMMSParserListener) EnterSurfaceCondition_Hole(ctx *SurfaceCondition_HoleContext) {}

// ExitSurfaceCondition_Hole is called when production surfaceCondition_Hole is exited.
func (s *BaseMMSParserListener) ExitSurfaceCondition_Hole(ctx *SurfaceCondition_HoleContext) {}

// EnterSurfaceCondition_Noise is called when production surfaceCondition_Noise is entered.
func (s *BaseMMSParserListener) EnterSurfaceCondition_Noise(ctx *SurfaceCondition_NoiseContext) {}

// ExitSurfaceCondition_Noise is called when production surfaceCondition_Noise is exited.
func (s *BaseMMSParserListener) ExitSurfaceCondition_Noise(ctx *SurfaceCondition_NoiseContext) {}

// EnterSurfaceCondition_Steep is called when production surfaceCondition_Steep is entered.
func (s *BaseMMSParserListener) EnterSurfaceCondition_Steep(ctx *SurfaceCondition_SteepContext) {}

// ExitSurfaceCondition_Steep is called when production surfaceCondition_Steep is exited.
func (s *BaseMMSParserListener) ExitSurfaceCondition_Steep(ctx *SurfaceCondition_SteepContext) {}

// EnterSurfaceCondition_StoneDepth is called when production surfaceCondition_StoneDepth is entered.
func (s *BaseMMSParserListener) EnterSurfaceCondition_StoneDepth(ctx *SurfaceCondition_StoneDepthContext) {
}

// ExitSurfaceCondition_StoneDepth is called when production surfaceCondition_StoneDepth is exited.
func (s *BaseMMSParserListener) ExitSurfaceCondition_StoneDepth(ctx *SurfaceCondition_StoneDepthContext) {
}

// EnterSurfaceCondition_Freezing is called when production surfaceCondition_Freezing is entered.
func (s *BaseMMSParserListener) EnterSurfaceCondition_Freezing(ctx *SurfaceCondition_FreezingContext) {
}

// ExitSurfaceCondition_Freezing is called when production surfaceCondition_Freezing is exited.
func (s *BaseMMSParserListener) ExitSurfaceCondition_Freezing(ctx *SurfaceCondition_FreezingContext) {
}

// EnterSurfaceCondition_VerticalGradient is called when production surfaceCondition_VerticalGradient is entered.
func (s *BaseMMSParserListener) EnterSurfaceCondition_VerticalGradient(ctx *SurfaceCondition_VerticalGradientContext) {
}

// ExitSurfaceCondition_VerticalGradient is called when production surfaceCondition_VerticalGradient is exited.
func (s *BaseMMSParserListener) ExitSurfaceCondition_VerticalGradient(ctx *SurfaceCondition_VerticalGradientContext) {
}

// EnterSurfaceCondition_AboveWater is called when production surfaceCondition_AboveWater is entered.
func (s *BaseMMSParserListener) EnterSurfaceCondition_AboveWater(ctx *SurfaceCondition_AboveWaterContext) {
}

// ExitSurfaceCondition_AboveWater is called when production surfaceCondition_AboveWater is exited.
func (s *BaseMMSParserListener) ExitSurfaceCondition_AboveWater(ctx *SurfaceCondition_AboveWaterContext) {
}

// EnterSurfaceCondition_YAbove is called when production surfaceCondition_YAbove is entered.
func (s *BaseMMSParserListener) EnterSurfaceCondition_YAbove(ctx *SurfaceCondition_YAboveContext) {}

// ExitSurfaceCondition_YAbove is called when production surfaceCondition_YAbove is exited.
func (s *BaseMMSParserListener) ExitSurfaceCondition_YAbove(ctx *SurfaceCondition_YAboveContext) {}

// EnterSurfaceCondition_Compound__Item is called when production surfaceCondition_Compound__Item is entered.
func (s *BaseMMSParserListener) EnterSurfaceCondition_Compound__Item(ctx *SurfaceCondition_Compound__ItemContext) {
}

// ExitSurfaceCondition_Compound__Item is called when production surfaceCondition_Compound__Item is exited.
func (s *BaseMMSParserListener) ExitSurfaceCondition_Compound__Item(ctx *SurfaceCondition_Compound__ItemContext) {
}

// EnterSurfaceCondition_Compound is called when production surfaceCondition_Compound is entered.
func (s *BaseMMSParserListener) EnterSurfaceCondition_Compound(ctx *SurfaceCondition_CompoundContext) {
}

// ExitSurfaceCondition_Compound is called when production surfaceCondition_Compound is exited.
func (s *BaseMMSParserListener) ExitSurfaceCondition_Compound(ctx *SurfaceCondition_CompoundContext) {
}

// EnterReference is called when production reference is entered.
func (s *BaseMMSParserListener) EnterReference(ctx *ReferenceContext) {}

// ExitReference is called when production reference is exited.
func (s *BaseMMSParserListener) ExitReference(ctx *ReferenceContext) {}

// EnterResourceReference is called when production resourceReference is entered.
func (s *BaseMMSParserListener) EnterResourceReference(ctx *ResourceReferenceContext) {}

// ExitResourceReference is called when production resourceReference is exited.
func (s *BaseMMSParserListener) ExitResourceReference(ctx *ResourceReferenceContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseMMSParserListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseMMSParserListener) ExitNumber(ctx *NumberContext) {}

// EnterKeyword is called when production keyword is entered.
func (s *BaseMMSParserListener) EnterKeyword(ctx *KeywordContext) {}

// ExitKeyword is called when production keyword is exited.
func (s *BaseMMSParserListener) ExitKeyword(ctx *KeywordContext) {}

// EnterVerticalAnchor is called when production verticalAnchor is entered.
func (s *BaseMMSParserListener) EnterVerticalAnchor(ctx *VerticalAnchorContext) {}

// ExitVerticalAnchor is called when production verticalAnchor is exited.
func (s *BaseMMSParserListener) ExitVerticalAnchor(ctx *VerticalAnchorContext) {}

// EnterVerticalAnchor_Absolute is called when production verticalAnchor_Absolute is entered.
func (s *BaseMMSParserListener) EnterVerticalAnchor_Absolute(ctx *VerticalAnchor_AbsoluteContext) {}

// ExitVerticalAnchor_Absolute is called when production verticalAnchor_Absolute is exited.
func (s *BaseMMSParserListener) ExitVerticalAnchor_Absolute(ctx *VerticalAnchor_AbsoluteContext) {}

// EnterVerticalAnchor_AboveBottom is called when production verticalAnchor_AboveBottom is entered.
func (s *BaseMMSParserListener) EnterVerticalAnchor_AboveBottom(ctx *VerticalAnchor_AboveBottomContext) {
}

// ExitVerticalAnchor_AboveBottom is called when production verticalAnchor_AboveBottom is exited.
func (s *BaseMMSParserListener) ExitVerticalAnchor_AboveBottom(ctx *VerticalAnchor_AboveBottomContext) {
}

// EnterVerticalAnchor_BelowTop is called when production verticalAnchor_BelowTop is entered.
func (s *BaseMMSParserListener) EnterVerticalAnchor_BelowTop(ctx *VerticalAnchor_BelowTopContext) {}

// ExitVerticalAnchor_BelowTop is called when production verticalAnchor_BelowTop is exited.
func (s *BaseMMSParserListener) ExitVerticalAnchor_BelowTop(ctx *VerticalAnchor_BelowTopContext) {}

// EnterWorldgenDeclaration is called when production worldgenDeclaration is entered.
func (s *BaseMMSParserListener) EnterWorldgenDeclaration(ctx *WorldgenDeclarationContext) {}

// ExitWorldgenDeclaration is called when production worldgenDeclaration is exited.
func (s *BaseMMSParserListener) ExitWorldgenDeclaration(ctx *WorldgenDeclarationContext) {}

// EnterNoiseDeclaration is called when production noiseDeclaration is entered.
func (s *BaseMMSParserListener) EnterNoiseDeclaration(ctx *NoiseDeclarationContext) {}

// ExitNoiseDeclaration is called when production noiseDeclaration is exited.
func (s *BaseMMSParserListener) ExitNoiseDeclaration(ctx *NoiseDeclarationContext) {}

// EnterNoiseDefinition is called when production noiseDefinition is entered.
func (s *BaseMMSParserListener) EnterNoiseDefinition(ctx *NoiseDefinitionContext) {}

// ExitNoiseDefinition is called when production noiseDefinition is exited.
func (s *BaseMMSParserListener) ExitNoiseDefinition(ctx *NoiseDefinitionContext) {}
