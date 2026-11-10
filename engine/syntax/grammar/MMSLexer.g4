// Lexer for Minecraft Metascript. Every character ends up in some token:
// whitespace and comments go to the hidden channel so the formatter and
// doc-comment lookup can see them.
lexer grammar MMSLexer;

// MISSING and SEP are never lexed. Recovery inserts MISSING where a
// field's value was left out (`key =` followed by the next field), so the
// parser does not read the next field as the value. SEP ends a statement
// whose last word is followed by a word on a later line, so `x` then `y = 1`
// on the next line stay two statements instead of reading as a
// declaration `x y = 1` (see separateWords in recover.go).
tokens { MISSING, SEP }

// Keywords are lowercase words only. Declaration kinds and helpers are
// ordinary PascalCase identifiers resolved later.
CONST    : 'const';
// `Type` is the one PascalCase keyword: its body is types, not values.
TYPE     : 'Type';
EXTERN   : 'extern';
// `private` scopes a const, Template, Type or Blueprint to its file.
PRIVATE  : 'private';
LET      : 'let';
IF       : 'if';
ELSE     : 'else';
WHILE    : 'while';
FOR      : 'for';
IN       : 'in';
RETURN   : 'return';
BREAK    : 'break';
CONTINUE : 'continue';
// Words that start a part of a Blueprint or Test body. They are still
// accepted as field keys, argument names, member names and declaration
// names (the parser's `word` rule), since game data uses some of them
// (`setup`, and `run` in hook declarations).
PRIMARY  : 'primary';
AS       : 'as';
CHECK    : 'check';
EMIT     : 'emit';
SETUP    : 'setup';
EXPECT   : 'expect';
RUN      : 'run';
TRUE     : 'true';
FALSE    : 'false';

// A raw command: a line in a function body whose first non-blank character
// is '/' followed by a command name, or the value of an assignment, after
// '=' (`let n = /data get block ~ ~ ~ Items`). It runs to the end of the
// line and is one token; the command checker parses it against the
// command tree.
COMMAND : '/' {p.atCommandStart()}? [a-z] ~[\r\n]*;

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

// A templated name: a name or resource location with `${...}` pieces, as
// in `${w}_crate` or `item/${self.path}`. Only declaration names and name
// expressions use it; strings and commands have their own interpolation.
// It never ends in ':', so `Objective kills_${w}: ...` is a name and a
// colon.
TNAME : TNAME_CHAR* INTERP ((TNAME_CHAR | INTERP)* (TNAME_END | INTERP))?;

IDENT  : IDENT_START IDENT_CHAR*;
NUMBER : DIGIT+ ('.' DIGIT+)? ([eE] [+-]? DIGIT+)?;
STRING : '"' (ESC | ~["\\\r\n])* '"' | '\'' (ESC | ~['\\\r\n])* '\'';

ELLIPSIS: '...';
DOTDOT  : '..';
QQ      : '??';
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
PIPE    : '|';
QUESTION: '?';

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
// A relative or local coordinate that a variable offsets, `~x`, `^-y`,
// `~spot.y`: one word, so `~ x` stays a coordinate and then a name.
REL_NAME : [~^] '-'? IDENT_START IDENT_CHAR* ('.' IDENT_START IDENT_CHAR*)*;
TILDE   : '~';
CARET   : '^';

// Anything else is a single error token, so the lexer never fails.
ERROR_CHAR : .;

fragment IDENT_START : [a-zA-Z_];
fragment IDENT_CHAR  : [a-zA-Z0-9_];
fragment NS_CHAR     : [a-z0-9_.-];
fragment PATH_CHAR   : [a-z0-9_./-];
fragment DIGIT       : [0-9];
fragment TNAME_CHAR  : [a-zA-Z0-9_./:-];
fragment TNAME_END   : [a-zA-Z0-9_./-];
fragment INTERP      : '${' ~[}\r\n]* '}';
fragment ESC         : '\\' .;
