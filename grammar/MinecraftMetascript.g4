grammar MinecraftMetascript;

file: NL* (namedBlock NL*)*;
namedBlock: Identifier Identifier NL* '{' ((varDecl | block) NL*)* '}';
block: Identifier NL* '{' NL* ((varDecl) NL*)* '}';
varDecl: Identifier '=' /* A value? */ value;

resourceReference: (Identifier ':')? Identifier;

fn: Identifier '(' (value ',')* value? ')' ('.' fn)*;

value: number | String | fn | resourceReference | conditional | list;

condition
  : rootCondition                                  #condPrimary
  | '(' condition '&&' condition ')'               #condGroupedAnd
  | '(' condition '||' condition ')'               #condGroupedOr
  | condition '&&' condition                       #condAnd
  | condition '||' condition                       #condOr
  ;

rootCondition: '!'? value;

conditional: 'If'  NL* '(' NL* condition  NL* ')' NL* value;
list: '['  NL* (value  NL*)* value?  NL* ']';

Int: '-'? [0-9]+;
Float: '-'? [0-9]* '.' [0-9]+;
number: Int | Float;

String: '"' ~[\r\n]* '"';

WS: [ \t]+ -> skip;
NL: [\n] -> channel(HIDDEN);

// Laziest
Identifier: [a-zA-Z] [a-zA-Z0-9_/]*;

BlockComment: '/*' .*? '*/' -> channel(HIDDEN);
LineComment: '//' ~[\r\n]* -> channel(HIDDEN);
