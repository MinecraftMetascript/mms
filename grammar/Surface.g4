grammar Surface;
import Core_Lang;

surfaceBlock: 'Surface' NL* '{' NL* (surfaceStatement NL*)* NL* '}';
surfaceStatement: verticalAnchorDeclaration | surfaceConditionDeclaration | surfaceRuleDeclaration;

verticalAnchor: ('~'? Int) | Identifier;
verticalAnchorDeclaration: Identifier '=' NL* verticalAnchor;

surfaceCondition:
    (
    surfaceCondition_Not
        | surfaceCondition_AboveSurface
        | surfaceCondition_Biome
        | surfaceCondition_Hole
        | surfaceCondition_Steep
        | surfaceCondition_Freezing
        | surfaceCondition_NoiseThreshold
        | surfaceCondition_StoneDepth
        | surfaceCondition_AboveWater
        | surfaceCondition_YAbove
        | surfaceCondition_Reference
        | surfaceCondition_And
        | surfaceCondition_Or
        | surfaceCondition_VerticalGradient
    )
;
surfaceConditionDeclaration: Identifier '=' NL* surfaceCondition;

surfaceCondition_Not: '!' surfaceCondition;
surfaceCondition_And: 'And' NL* '(' NL* (surfaceCondition NL*)* surfaceCondition NL* ')';
surfaceCondition_Or: 'Or' NL* '(' NL* (surfaceCondition NL*)* surfaceCondition NL* ')';


surfaceCondition_Reference: resourceReference;
surfaceCondition_AboveSurface: 'AboveSurface()';
surfaceCondition_Biome: 'Biome' '(' NL* (resourceReference ',' NL*)* NL* resourceReference NL* ')';

surfaceCondition_Hole: 'Hole' '(' ')';
surfaceCondition_Steep: 'Steep' '(' ')';
surfaceCondition_Freezing: 'Freezing' '(' ')';


// These are split out to make them easier to differentiate in the go code
surfaceCondition_NoiseThresholdBuilder: builder_Max | builder_Min;
surfaceCondition_NoiseThreshold: 'NoiseThreshold' '(' resourceReference ')' NL* (surfaceCondition_NoiseThresholdBuilder NL*)*;

StoneDepthMode: 'Floor' | 'Ceiling';
surfaceCondition_StoneDepth:
    'StoneDepth(' StoneDepthMode ')' NL*
    (surfaceCondition_StoneDepthBuilder NL*)*;

surfaceCondition_StoneDepthBuilder:
    builder_Offset
    | builder_Add
    | surfaceCondition_StoneDepthBuilder_SecondaryDepthRange
    ;

surfaceCondition_StoneDepthBuilder_SecondaryDepthRange:'.SecondaryDepthRange(' Int ')';

surfaceCondition_VerticalGradient: 'VerticalGradient' '(' String ')' NL* (surfaceCondition_VerticalGradientBuilder NL*)*;
surfaceCondition_VerticalGradientBuilder:
    builder_Top
    | builder_Bottom
    ;



surfaceCondition_AboveWater: 'AboveWater' '(' ')' NL* (surfaceCondition_AboveWaterBuilder NL*)*;
surfaceCondition_AboveWaterBuilder:
    builder_Offset
    | builder_Add
    | builder_Mul;

surfaceCondition_YAbove: 'YAbove' '(' verticalAnchor ')' NL* (surfaceCondition_YAboveBuilder NL*)* ;
surfaceCondition_YAboveBuilder: builder_MulInt | builder_Add;

surfaceRuleDeclaration: Identifier '=' NL*  surfaceRule;
surfaceRule: surfaceRule_Block
            | surfaceRule_Sequence
            | surfaceRule_Reference
            | surfaceRule_If
            | surfaceRule_Bandlands;

surfaceRule_Reference: resourceReference;
surfaceRule_Block: 'Block' '(' resourceReference ')';
surfaceRule_Sequence: '[' NL* (surfaceRule NL*)* NL* ']';
surfaceRule_Bandlands: 'Bandlands' '(' ')';
surfaceRule_If: 'If' NL* '(' NL* surfaceCondition NL* ')' NL* surfaceRule;