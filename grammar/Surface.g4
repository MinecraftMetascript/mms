grammar Surface;
import Core_Lang;

surface: 'Surface' NL* '{' NL* (surfaceStatement NL*)* NL* '}';
surfaceStatement: verticalAnchorDeclaration | surfaceConditionDeclaration | surfaceRuleDeclaration;

verticalAnchor: ('~'? Int) | Identifier;
verticalAnchorDeclaration: Identifier '=' verticalAnchor;

sharedBuilder_Offset: '.Offset(' Int ')';
sharedBuilder_Add:'.Add()';
sharedBuilder_Mul: '.Mul(' number ')';
sharedBuilder_MulInt: '.Mul(' Int ')';

surfaceCondition:
    (
    surfaceCondition_Not
        | surfaceCondition_AboveSurface
        | surfaceCondition_Biome
        | surfaceCondition_Hole
        | surfaceCondition_Steep
        | surfaceCondition_Freezing
        | surfaceCondition_Noise
        | surfaceCondition_StoneDepth
        | surfaceCondition_AboveWater
        | surfaceCondition_YAbove
        | surfaceCondition_Reference
        | surfaceCondition_And
        | surfaceCondition_Or
        | surfaceCondition_VerticalGradient
    )
;
surfaceConditionDeclaration: Identifier '=' surfaceCondition;

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
surfaceCondition_NoiseBuilder_Min: '.Min(' number ')';
surfaceCondition_NoiseBuilder_Max: '.Max(' number ')';
surfaceCondition_NoiseBuilder: surfaceCondition_NoiseBuilder_Max | surfaceCondition_NoiseBuilder_Min;
surfaceCondition_Noise: 'Noise' '(' resourceReference ')' NL* (surfaceCondition_NoiseBuilder NL*)*;

StoneDepthMode: 'Floor' | 'Ceiling';
surfaceCondition_StoneDepth:
    'StoneDepth(' StoneDepthMode ')' NL*
    (surfaceCondition_StoneDepthBuilder NL*)*;

surfaceCondition_StoneDepthBuilder:
    sharedBuilder_Offset
    | sharedBuilder_Add
    | surfaceCondition_StoneDepthBuilder_SecondaryDepthRange
    ;

surfaceCondition_StoneDepthBuilder_SecondaryDepthRange:'.secondaryDepthRange(' Int ')';

surfaceCondition_VerticalGradient: 'VerticalGradient' '(' String ')' NL* (surfaceCondition_VerticalGradientBuilder NL*)*;
surfaceCondition_VerticalGradientBuilder:
    surfaceCondition_VerticalGradientBuilder_Top
    | surfaceCondition_VerticalGradientBuilder_Bottom
    ;

surfaceCondition_VerticalGradientBuilder_Top: '.Top' '(' verticalAnchor ')';
surfaceCondition_VerticalGradientBuilder_Bottom: '.Bottom' '(' verticalAnchor ')';


surfaceCondition_AboveWater: 'AboveWater' '(' ')' NL* (surfaceCondition_AboveWaterBuilder NL*)*;
surfaceCondition_AboveWaterBuilder:
    sharedBuilder_Offset
    | sharedBuilder_Add
    | sharedBuilder_Mul;

surfaceCondition_YAbove: 'YAbove' '(' verticalAnchor ')' NL* (surfaceCondition_YAboveBuilder NL*)* ;
surfaceCondition_YAboveBuilder: sharedBuilder_MulInt | sharedBuilder_Add;

surfaceRuleDeclaration: Identifier '=' surfaceRule;
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