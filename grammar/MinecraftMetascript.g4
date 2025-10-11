grammar MinecraftMetascript;

file: NL* (namedBlock NL*)*;
namedBlock: Identifier Identifier NL* '{' ((varDecl | block) NL*)* '}';
block: Identifier NL* '{' NL* ((varDecl) NL*)* '}';
varDecl: (docString)? Identifier '=' /* A value? */ value?;

resourceReference: ((Identifier ':')? Identifier) | (Identifier ':' Identifier?);
resourceTag: '#' resourceReference;

// We allow the trailing "." here to prevent diag spam when typing
fn: Identifier fnArgBody  ('.' fn)*? '.'?;
// Allow trailing commas to ensure this parses properly when editing
fnArgBody: '(' (value ',')* (value ','?)? ')';

value: number | String | fn | resourceReference | conditional | list | resourceTag;

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

docString: DocString NL*;
DocString: '/**' .*? '*/';
BlockComment: '/*' .*? '*/' -> channel(HIDDEN);
LineComment: '//' ~[\r\n]* -> channel(HIDDEN);
