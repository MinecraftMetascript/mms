grammar MinecraftMetascript;

file: NL* (namedBlock NL*)*;
namedBlock: Identifier Identifier NL* '{' ((varDecl | block) NL*)* '}';
block: Identifier NL* '{' NL* ((varDecl) NL*)* '}';
varDecl: Identifier '=' /* A value? */ value?;

resourceReference: (Identifier ':')? Identifier;

// Allow trailing commas to ensure this parses properly when editing
fn: Identifier fnArgBody  ('.' fn)*?;
fnArgBody: '(' (value ',')* (value ','?)? ')';

value: number | String | fn | resourceReference | conditional | list;

condition
  : '!' condition                                       #condNegate
  | rootCondition                                       #condPrimary
  | '(' NL* condition NL* ')'                           #condGrouped
  | condition NL* '&&' NL*  condition                   #condAnd
  | condition NL* '||' NL*  condition                   #condOr
  ;

rootCondition: value;

// TODO: Should we have support for "Else"?
conditional: 'If'  NL* conditionalBody NL* value?;
conditionalBody: '(' NL* condition?  NL* ')';
list: '['  NL* (value NL* ','? NL*)* value?  NL* ']';

Int: '-'? [0-9]+;
Float: '-'? ([0-9]+ '.' [0-9]*) | ([0-9]* '.' [0-9]+);
number: Int | Float;

String: '"' ~[\r\n]* '"';

WS: [ \t]+ -> skip;
NL: [\n] -> channel(HIDDEN);

// Laziest
Identifier: [a-zA-Z_] [a-zA-Z0-9_/]*;

BlockComment: '/*' .*? '*/' -> channel(HIDDEN);
LineComment: '//' ~[\r\n]* -> channel(HIDDEN);
