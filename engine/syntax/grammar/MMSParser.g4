// Parser for Minecraft Metascript. The grammar knows shapes, not Minecraft:
// a declaration is `Kind name ...` and Kind is resolved later. Optional
// tails (args?, chains with a missing name) let half-typed code such as
// `Noise(-7).` parse, which completion relies on.
parser grammar MMSParser;

options { tokenVocab = MMSLexer; }

file : topItem* EOF;

// Entry point for the expression inside a `${...}` piece of a name.
exprOnly : expr EOF;

// Namespace example
// Noise ridges = Noise(-7)
// LootTable blocks/ruby_ore { pools = [ ... ] }
// const spacing: int = 32
// for w in woods { Recipe ${w}_crate = ... }
// extern Template font_width(text: string): int
topItem
    : EXTERN? declaration                                         # declItem
    | constDecl                                                   # constItem
    | letDecl                                                     # letItem
    | FOR IDENT (COMMA IDENT)? IN expr block                      # forItem
    ;

// Global deaths: int = 0
// Function add(a: int, b: int): int { return a + b }
// OnLoad { ... }
// Type LangText = string | {string: string}
// JigsawHole room_exit, hall_entry
declaration
    : kind=TYPE name=declName ASSIGN typeExpr
    | kind=TYPE name=declName block
    | kind=IDENT name=declName (COMMA declName)+
    | kind=IDENT name=declName (typeAnnotation | params typeAnnotation?)? declBody?
    | kind=IDENT params? block
    ;

// The same shapes nested in a block (an emit block or a top-level for). A
// rule of its own keeps what may follow a top-level declaration narrow,
// which keeps error recovery at the top level sharp. `Kind { }` without a
// name parses as an expression statement; the builder turns it into a
// declaration where declarations belong.
innerDecl
    : kind=IDENT name=declName (COMMA declName)+
    | kind=IDENT name=declName (typeAnnotation | params typeAnnotation?)? innerBody?
    ;

innerBody : ASSIGN value | block;

declName : word | RESLOC | TNAME;

// An identifier, or a body keyword used as a plain name.
word : IDENT | TYPE | PRIMARY | AS | CHECK | EMIT | SETUP | EXPECT | RUN;

// When the value after '=' is left out and the next line starts a new
// declaration, recovery inserts a MISSING token as the value (see
// markMissingValues), so the value never swallows that declaration.
declBody
    : ASSIGN value
    | block
    ;

constDecl : CONST IDENT typeAnnotation? (ASSIGN value)?;
letDecl   : LET IDENT typeAnnotation? (ASSIGN value)?;

typeAnnotation : COLON typeExpr;

// int, Id<item>, [T], {string: T}, T?, A | B, "rigid" | "terrain_matching",
// and a record written in place, { x: int, y: int = 0 }, whose fields are
// separated by commas or line breaks as in a Type body.
typeExpr : typeArm (PIPE typeArm)*;
typeArm  : typeAtom QUESTION?;
typeAtom
    : typeName (LT typeExpr (COMMA typeExpr)* GT)?               # namedType
    | STRING                                                      # literalType
    | LBRACK typeExpr RBRACK                                      # listType
    | LBRACE typeField ((COMMA | SEP)? typeField)* (COMMA | SEP)? RBRACE # recordType
    | LBRACE mapKey COLON typeExpr RBRACE                         # mapType
    | LPAREN typeExpr RPAREN                                      # parenType
    ;
// A record's field. `{string: T}` reads as a record with one field too;
// the builder turns a lone field named `string`, or one with a PascalCase
// name, into the map type it is.
typeField : IDENT typeAnnotation? (ASSIGN value)?;
// A map's key type other than a plain name, which reads as a record's
// field: `{Id<item>: int}`. The two alternatives part at the second token,
// so a half-typed field still recovers inside the record.
mapKey : typeName LT typeExpr (COMMA typeExpr)* GT | RESLOC | LPAREN typeExpr RPAREN;
typeName : IDENT (DOT IDENT)* | RESLOC;

params : LPAREN (param (COMMA param)* COMMA?)? RPAREN;
param  : IDENT typeAnnotation? (ASSIGN value)?;

// A block holds items. In a resource body they are fields; in a function
// body they are statements; in a Type or Blueprint body they are field
// declarations and sections. The checker decides what fits by context.
block : LBRACE blockItem* RBRACE;

blockItem
    : statement (COMMA | SEP)?
    ;

statement
    : letDecl                                                     # letStmt
    | constDecl                                                   # constStmt
    | IF LPAREN expr RPAREN thenPart (ELSE elsePart)?             # ifStmt
    | WHILE LPAREN expr RPAREN block                              # whileStmt
    | FOR IDENT (COMMA IDENT)? IN expr block                      # forStmt
    | RETURN value?                                               # returnStmt
    | BREAK                                                       # breakStmt
    | CONTINUE                                                    # continueStmt
    | COMMAND                                                     # commandStmt
    | PRIMARY innerDecl                                           # primaryStmt
    | innerDecl                                                   # declStmt
    | AS typeExpr ASSIGN value hint?                              # asStmt
    | section=(CHECK | EMIT | SETUP) block                        # sectionStmt
    | EXPECT expr                                                 # expectStmt
    | RUN bareCall                                                # runStmt
    | IDENT typeAnnotation (ASSIGN value)?                        # fieldDeclStmt
    | fieldKey ASSIGN value                                       # fieldStmt
    | IDENT (DOT IDENT)+ ASSIGN value                             # pathAssignStmt
    | target=(IDENT | TNAME) LBRACK arg RBRACK ASSIGN value       # indexAssignStmt
    | expr op=(PLUS_EQ | MINUS_EQ | STAR_EQ | SLASH_EQ | PCT_EQ) expr # compoundAssignStmt
    | expr                                                        # exprStmt
    ;

// A statement-level if takes a block, or a single call as in
// `if (cooldown < 1) error(cooldown, "...")`. The call must start with a
// name, so the choice is made on the first token and a broken block still
// recovers as a block.
thenPart : block | bareCall;
elsePart : ifTail | block | bareCall;
bareCall : IDENT postfix*;

// `hint("rule", "message")` after an `as` conversion.
hint : {p.word("hint")}? IDENT args;

ifTail : IF LPAREN expr RPAREN thenPart (ELSE elsePart)?;

fieldKey : word | RESLOC | STRING;

// A value is an expression, or an if expression such as
// `item_name = if (name != none) Translate(...)`. If expressions sit only
// where a value is expected, so a statement that starts with `if` is always
// an if statement.
// A command as a value, `let n = /data get block ~ ~ ~ Items`, is its
// result (or, for a bool, whether it succeeded).
value
    : IF LPAREN expr RPAREN value (ELSE value)?
    | COMMAND
    | expr
    ;

expr
    : expr postfix                                       # postfixExpr
    | op=(BANG | MINUS | PLUS) expr                      # unaryExpr
    | expr op=(STAR | SLASH | PERCENT) expr              # mulExpr
    | expr op=(PLUS | MINUS) expr                        # addExpr
    | expr DOTDOT expr?                                  # rangeExpr
    | DOTDOT expr                                        # rangeToExpr
    | expr op=(LT | LE | GT | GE) expr                   # compareExpr
    | expr op=(EQEQ | NEQ) expr                          # equalExpr
    | expr ANDAND expr                                   # andExpr
    | expr OROR expr                                     # orExpr
    | expr QQ expr                                       # coalesceExpr
    | ELLIPSIS expr                                      # spreadExpr
    | primary                                            # primaryExpr
    ;

postfix
    : DOT name=word?                                     # memberPostfix
    | args                                               # callPostfix
    | LBRACK (arg (COMMA arg)* COMMA?)? RBRACK           # indexPostfix
    | block                                              # blockPostfix
    ;

args : LPAREN (arg (COMMA arg)* COMMA?)? RPAREN;
arg  : (word ASSIGN)? value;

primary
    : NUMBER                                             # numberLit
    | STRING                                             # stringLit
    | (TRUE | FALSE)                                     # boolLit
    | RUN block                                          # runExpr
    | IDENT                                              # identRef
    | RESLOC                                             # reslocRef
    | TNAME                                              # templatedRef
    | HASH (IDENT | RESLOC | TNAME)                      # tagRef
    | AT IDENT                                           # selector
    | ((TILDE | CARET) (MINUS? NUMBER)? | REL_NAME)+     # coordinate
    | LPAREN value RPAREN                                # parenExpr
    | LBRACK FOR IDENT (COMMA IDENT)? IN expr (IF expr)? COLON value RBRACK # comprehension
    | LBRACK (value (COMMA? value)* COMMA?)? RBRACK            # listLit
    | block                                              # objectLit
    | MISSING                                            # missingLit
    ;
