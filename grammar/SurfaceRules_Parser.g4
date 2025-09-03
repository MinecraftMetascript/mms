parser grammar SurfaceRules_Parser;

import Lang_Parser, VerticalAnchor_Parser;

options {
    tokenVocab = Main_Lexer;
}

surfaceRuleDefinition: Keyword_SurfaceRule NL* CurlyOpen NL* (surfaceRule) NL* CurlyClose;

surfaceRule: surfaceRule_Block | surfaceRule_Bandlands | surfaceRule_Condition | surfaceRule_Sequence | surfaceRule_Reference;

surfaceRule_Block: Keyword_Block resourceReference;
surfaceRule_Bandlands: Keyword_Bandlands;
surfaceRule_Condition: Keyword_If RoundOpen NL* surfaceCondition NL* RoundClose NL* surfaceRule;
surfaceRule_Sequence: Keyword_Sequence SquareOpen NL* (surfaceRule NL*)* NL* SquareClose;
surfaceRule_Reference: resourceReference;

surfaceConditionDefinition: Keyword_SurfaceCondition NL* CurlyOpen NL* surfaceCondition NL* CurlyClose;

surfaceCondition:
    Bang?
    (surfaceCondition_AboveSurface
    | surfaceCondition_Biome
    | surfaceCondition_Hole
    | surfaceCondition_Noise
    | surfaceCondition_Steep
    | surfaceCondition_StoneDepth
    | surfaceCondition_Freezing
    | surfaceCondition_VerticalGradient
    | surfaceCondition_AboveWater
    | surfaceCondition_YAbove
    | surfaceCondition_Reference
    | surfaceCondition_And
    | surfaceCondition_Or);

surfaceCondition_Reference: resourceReference;

surfaceCondition_And: Keyword_And NL* RoundOpen NL* (surfaceCondition NL*)* surfaceCondition NL* RoundClose;
surfaceCondition_Or: Keyword_Or NL* RoundOpen NL* (surfaceCondition NL*)* surfaceCondition NL* RoundClose;


surfaceCondition_AboveSurface: Keyword_AboveSurface;
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

