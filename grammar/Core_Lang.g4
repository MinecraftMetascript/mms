grammar Core_Lang;

import Noise, DensityFunctions, Surface;

builder_XZScale: '.XZScale(' number ')';
builder_YScale: '.YScale(' number ')';
builder_XZFactor: '.XZFactor(' number ')';
builder_YFactor: '.YFactor(' number ')';
builder_Noise: '.Noise' (noiseDefinition | ('(' resourceReference ')'));
builder_Smear: 'Smear' '(' number ')';
builder_Type1: '.Type1' '(' ')';
builder_Type2: '.Type2' '(' ')';
builder_ShiftX: '.ShiftX' '(' densityFn ')';
builder_ShiftY: '.ShiftY' '(' densityFn ')';
builder_ShiftZ: '.ShiftZ' '(' densityFn ')';
builder_Amplitudes: '.Amplitudes' '(' (number ',')* number  ')';
builder_Offset: '.Offset(' Int ')';
builder_Add:'.Add()';
builder_Mul: '.Mul(' number ')';
builder_MulInt: '.Mul(' Int ')';
builder_Min: '.Min(' number ')';
builder_Max: '.Max(' number ')';
builder_Top: '.Top' '(' verticalAnchor ')';
builder_TopLiteral: '.Top' '(' Int ')';
builder_Bottom: '.Bottom' '(' verticalAnchor ')';
builder_BottomLiteral: '.Bottom' '(' Int ')';
builder_InRange: '.InRange' '(' densityFn ')';
builder_OutRange: '.OutRange' '(' densityFn ')';


resourceReference: (Identifier ':')? Identifier;


Int: '-'? [0-9]+;
Float: ('-'? [0-9]* '.' [0-9]+);
String: '"' ~[\r\n]* '"';
number: Int | Float;

NL: [\n];
WS: [ \t]+ -> skip;
Identifier: [a-zA-Z] [a-zA-Z0-9_]*;

BlockComment: '/*' .*? '*/' -> channel(HIDDEN);
LineComment: '//' ~[\r\n]* -> channel(HIDDEN);
