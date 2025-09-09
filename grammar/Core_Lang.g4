grammar Core_Lang;

import Noise, DensityFunctions, Surface;

builder_XZScale: '.XZScale(' NL* number NL* ')';
builder_YScale: '.YScale(' NL* number NL* ')';
builder_XZFactor: '.XZFactor(' NL* number NL* ')';
builder_YFactor: '.YFactor(' NL* number NL* ')';
builder_Noise: '.Noise' (noiseDefinition | ('(' NL* resourceReference NL* ')'));
builder_Smear: 'Smear' '(' NL* number NL* ')';
builder_Type1: '.Type1' '(' ')';
builder_Type2: '.Type2' '(' ')';
builder_ShiftX: '.ShiftX' '(' NL* densityFn NL* ')';
builder_ShiftY: '.ShiftY' '(' NL* densityFn NL* ')';
builder_ShiftZ: '.ShiftZ' '(' NL* densityFn NL* ')';
builder_Amplitudes: '.Amplitudes' '(' (number ',')* number  ')';
builder_Offset: '.Offset' '(' NL* Int NL* ')';
builder_Add:'.Add' '(' ')' ;
builder_Mul: '.Mul(' NL* number NL* ')';
builder_MulInt: '.Mul(' NL* Int NL* ')';
builder_Min: '.Min(' NL* number NL* ')';
builder_Max: '.Max(' NL* number NL* ')';
builder_Top: '.Top' '(' NL* verticalAnchor NL* ')';
builder_TopLiteral: '.Top' '(' Int ')';
builder_Bottom: '.Bottom' '(' NL* verticalAnchor NL* ')';
builder_BottomLiteral: '.Bottom' '('NL* Int NL* ')';
builder_InRange: '.InRange' '(' NL* densityFn NL* ')';
builder_OutRange: '.OutRange' '(' NL* densityFn NL* ')';


resourceReference: (Identifier ':')? Identifier;


Int: '-'? [0-9]+;
Float: ('-'? [0-9]* '.' [0-9]+);
String: '"' ~[\r\n]* '"';
number: Int | Float;

NL: [\n];
WS: [ \t]+ -> skip;
Identifier: [a-zA-Z] [a-zA-Z0-9_/]*;

BlockComment: '/*' .*? '*/' -> channel(HIDDEN);
LineComment: '//' ~[\r\n]* -> channel(HIDDEN);
