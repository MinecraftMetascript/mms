// Lexer for Minecraft Metascript. Every character ends up in some token:
// whitespace and comments go to the hidden channel so the formatter and
// doc-comment lookup can see them (docs/plan.md, "ANTLR setup").
lexer grammar MMSLexer;

// Keywords are lowercase statement words only. Declaration kinds and
// helpers are ordinary PascalCase identifiers resolved against the spec.
LET      : 'let';
IF       : 'if';
ELSE     : 'else';
WHILE    : 'while';
FOR      : 'for';
IN       : 'in';
RETURN   : 'return';
BREAK    : 'break';
CONTINUE : 'continue';
TRUE     : 'true';
FALSE    : 'false';

// A raw command: a line in a function body whose first non-blank character
// is '/' followed by a command name. The whole line is one token; the
// command checker parses it against the command tree.
COMMAND : '/' {p.atLineStart()}? [a-z] ~[\r\n]*;

DOC_COMMENT   : '/**' .*? '*/' -> channel(HIDDEN);
BLOCK_COMMENT : '/*' .*? '*/' -> channel(HIDDEN);
// An unclosed block comment runs to the end of the file, as in Go and C.
UNTERMINATED_COMMENT : '/*' (~'*' | '*'+ ~[*/])* '*'* EOF -> channel(HIDDEN);
LINE_COMMENT  : '//' ~[\r\n]* -> channel(HIDDEN);
NEWLINE       : '\r'? '\n' -> channel(HIDDEN);
WS            : [ \t]+ -> channel(HIDDEN);

// Resource locations: `minecraft:stone`, `blocks/ruby_ore`,
// `example:terrain/noise_0`. They never contain spaces, so `a / b` is a
// division and `a/b` is a path (the formatter always spaces operators).
RESLOC
    : NS_CHAR+ ':' PATH_CHAR+
    | IDENT_START IDENT_CHAR* ('/' PATH_CHAR+)+
    ;

IDENT  : IDENT_START IDENT_CHAR*;
NUMBER : DIGIT+ ('.' DIGIT+)? ([eE] [+-]? DIGIT+)?;
STRING : '"' (ESC | ~["\\\r\n])* '"' | '\'' (ESC | ~['\\\r\n])* '\'';

ARROW   : '->';
DOTDOT  : '..';
PLUS_EQ : '+=';
MINUS_EQ: '-=';
STAR_EQ : '*=';
SLASH_EQ: '/=';
PCT_EQ  : '%=';
EQEQ    : '==';
NEQ     : '!=';
LE      : '<=';
GE      : '>=';
ANDAND  : '&&';
OROR    : '||';

LBRACE  : '{';
RBRACE  : '}';
LPAREN  : '(';
RPAREN  : ')';
LBRACK  : '[';
RBRACK  : ']';
COMMA   : ',';
COLON   : ':';
DOT     : '.';
ASSIGN  : '=';
LT      : '<';
GT      : '>';
PLUS    : '+';
MINUS   : '-';
STAR    : '*';
SLASH   : '/';
PERCENT : '%';
BANG    : '!';
HASH    : '#';
AT      : '@';
TILDE   : '~';
CARET   : '^';

// Anything else is a single error token, so the lexer never fails.
ERROR_CHAR : .;

fragment IDENT_START : [a-zA-Z_];
fragment IDENT_CHAR  : [a-zA-Z0-9_];
fragment NS_CHAR     : [a-z0-9_.-];
fragment PATH_CHAR   : [a-z0-9_./-];
fragment DIGIT       : [0-9];
fragment ESC         : '\\' .;
