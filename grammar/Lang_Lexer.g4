lexer grammar Lang_Lexer;

Operator_Declare: ':=';


CurlyOpen: '{';
CurlyClose: '}';
SquareOpen: '[';
SquareClose: ']';
RoundOpen: '(';
RoundClose: ')';
Dot: '.';
Bang: '!';
Comma: ',';
Colon: ':';
SemiColon: ';';
Amp: '&';
DoubleAmp: '&&';
Tilde: '~';
Dash: '-';

Int: '-'? [0-9]+;
Float: ('-'? [0-9]+ '.' [0-9]+);
String: '"' ~[\r\n]* '"';


NL: [\n];
WS: [ \t]+ -> skip;
Identifier: [a-zA-Z][a-zA-Z0-9_]*;


BlockComment: '/*' .*? '*/' -> channel(HIDDEN);
LineComment: '//' ~[\r\n]* -> channel(HIDDEN);

