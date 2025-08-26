parser grammar MMS_SurfaceRules;
options {
	tokenVocab = MMSLexer;
}

import MMS_Lang_Parsers, MMS_VerticalAnchor;

surfaceDeclaration: Keyword_Surface surfaceDefinition;
surfaceDefinition:
	CurlyOpen NL* (
		(surfaceRuleDeclaration | surfaceConditionDeclaration) NL*
	)* CurlyClose;

//// Surface Rules
surfaceRuleReference: (Identifier | reference);
surfaceRuleDeclaration: Keyword_Rule Identifier NL* surfaceRule;
surfaceRule:
	surfaceRule_Conditional
	| surfaceRule_Bandlands
	| surfaceRule_Block
	| surfaceRule_Sequence
	| surfaceRuleReference;
surfaceRule_Conditional:
	Keyword_If NL* Bang? RoundOpen NL* surfaceCondition NL* RoundClose NL* (
		surfaceRule
	);

surfaceRule_Bandlands: Keyword_Bandlands;
surfaceRule_Block: Keyword_Block resourceReference;
surfaceRule_Sequence:
	Keyword_Sequence SquareOpen NL* ((surfaceRule) NL*)* SquareClose;

//// Surface Conditions
surfaceConditionReference: (Identifier | reference);
surfaceConditionDeclaration:
	Keyword_Condition Identifier NL* surfaceCondition;
surfaceCondition:
	surfaceCondition_AboveSurface
	| surfaceCondition_Biome
	| surfaceCondition_Hole
	| surfaceCondition_Noise
	| surfaceCondition_Steep
	| surfaceCondition_StoneDepth
	| surfaceCondition_Freezing
	| surfaceCondition_VerticalGradient
	| surfaceCondition_AboveWater
	| surfaceCondition_YAbove
	| surfaceCondition_Compound
	| surfaceConditionReference;

surfaceCondition_AboveSurface: Keyword_AbovePreliminarySurface;
surfaceCondition_Biome:
	Keyword_Biome SquareOpen NL* (resourceReference NL*)* SquareClose;
surfaceCondition_Hole: Keyword_Hole;
surfaceCondition_Noise:
	Keyword_Noise resourceReference SquareOpen number Comma number SquareClose;
surfaceCondition_Steep: Keyword_Steep;
surfaceCondition_StoneDepth:
	Keyword_StoneDepth (Keyword_Floor | Keyword_Ceiling) //
	Int // Offset
	(Keyword_Add | Keyword_Sub) // add surface depth
	Int; // secondary depth range

surfaceCondition_Freezing: Keyword_Freezing;

surfaceCondition_VerticalGradient:
	Keyword_VerticalGradient String verticalAnchor Comma verticalAnchor;

surfaceCondition_AboveWater:
	Keyword_AboveWater //
	Int // Offset
	number // Multiplier
	(Keyword_Add | Keyword_Sub);

surfaceCondition_YAbove:
	Keyword_YAbove verticalAnchor Int (Keyword_Add | Keyword_Sub);

surfaceCondition_Compound__Item: Bang? surfaceCondition;
surfaceCondition_Compound:
	Keyword_And NL* RoundOpen NL* (surfaceCondition_Compound__Item NL*)* RoundClose;