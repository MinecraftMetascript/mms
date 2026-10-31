// Parser for Minecraft Metascript. The grammar knows shapes, not Minecraft:
// a declaration is `Kind name ...` and Kind is resolved against the spec
// pack later. Optional tails (args?, chains with a missing name) let
// half-typed code such as `Noise(-7).` parse, which completion relies on.
parser grammar MMSParser;

options { tokenVocab = MMSLexer; }

file : declaration* EOF;

// Namespace example
// Noise ridges = Noise(-7)
// LootTable blocks/ruby_ore { pools = [ ... ] }
// Global deaths: int = 0
// Function on_kill(target: selector) { ... }
// OnLoad { ... }
declaration
    : kind=IDENT name=declName typeAnnotation? params? declBody?
    | kind=IDENT params? block
    ;

declName : IDENT | RESLOC;

// The value after '=' is optional to the parser so that `Noise a =`
// followed by a new declaration on the next line recovers cleanly; the
// predicate stops the value from swallowing that declaration.
declBody
    : ASSIGN ({p.valueFollows()}? expr)?
    | block
    ;

typeAnnotation : COLON typeRef;
typeRef        : IDENT | RESLOC;

params : LPAREN (param (COMMA param)* COMMA?)? RPAREN;
param  : IDENT typeAnnotation?;

// A block holds items. In a resource body they are fields; in a function
// body they are statements; the checker decides which by context.
block : LBRACE blockItem* RBRACE;

blockItem
    : statement COMMA?
    ;

statement
    : LET IDENT typeAnnotation? (ASSIGN expr)?                    # letStmt
    | IF LPAREN expr RPAREN block (ELSE (ifTail | block))?        # ifStmt
    | WHILE LPAREN expr RPAREN block                              # whileStmt
    | FOR IDENT IN expr block                                     # forStmt
    | RETURN expr?                                                # returnStmt
    | BREAK                                                       # breakStmt
    | CONTINUE                                                    # continueStmt
    | COMMAND                                                     # commandStmt
    | fieldKey ASSIGN expr                                        # fieldStmt
    | expr op=(PLUS_EQ | MINUS_EQ | STAR_EQ | SLASH_EQ | PCT_EQ) expr # compoundAssignStmt
    | expr                                                        # exprStmt
    ;

ifTail : IF LPAREN expr RPAREN block (ELSE (ifTail | block))?;

fieldKey : IDENT | RESLOC | STRING;

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
    | <assoc=right> expr ARROW expr                      # arrowExpr
    | primary                                            # primaryExpr
    ;

postfix
    : DOT name=IDENT?                                    # memberPostfix
    | args                                               # callPostfix
    | LBRACK (arg (COMMA arg)* COMMA?)? RBRACK           # indexPostfix
    | block                                              # blockPostfix
    ;

args : LPAREN (arg (COMMA arg)* COMMA?)? RPAREN;
arg  : (IDENT ASSIGN)? expr;

primary
    : NUMBER                                             # numberLit
    | STRING                                             # stringLit
    | (TRUE | FALSE)                                     # boolLit
    | IDENT                                              # identRef
    | RESLOC                                             # reslocRef
    | HASH (IDENT | RESLOC)                              # tagRef
    | AT IDENT                                           # selector
    | ((TILDE | CARET) (MINUS? NUMBER)?)+                # coordinate
    | LPAREN expr RPAREN                                 # parenExpr
    | LBRACK (expr (COMMA expr)* COMMA?)? RBRACK         # listLit
    | block                                              # objectLit
    ;
