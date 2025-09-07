grammar Core_Lang;

resourceReference: (Identifier ':')? Identifier;


Int: '-'? [0-9]+;
Float: ('-'? [0-9]* '.' [0-9]+);
String: '"' ~[\r\n]* '"';
number: Int | Float;

NL: [\n];
WS: [ \t]+ -> skip;
Identifier: [a-zA-Z][a-zA-Z0-9_]*;

BlockComment: '/*' .*? '*/' -> channel(HIDDEN);
LineComment: '//' ~[\r\n]* -> channel(HIDDEN);
