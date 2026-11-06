// Code generated from MMSParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // MMSParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type MMSParser struct {
	*antlr.BaseParser
}

var MMSParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func mmsparserParserInit() {
	staticData := &MMSParserParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "'const'", "'Type'", "'extern'", "'let'", "'if'", "'else'",
		"'while'", "'for'", "'in'", "'return'", "'break'", "'continue'", "'primary'",
		"'as'", "'check'", "'emit'", "'setup'", "'expect'", "'run'", "'true'",
		"'false'", "", "", "", "", "", "", "", "", "", "", "", "", "'->'", "'...'",
		"'..'", "'??'", "'+='", "'-='", "'*='", "'/='", "'%='", "'=='", "'!='",
		"'<='", "'>='", "'&&'", "'||'", "'|'", "'?'", "'{'", "'}'", "'('", "')'",
		"'['", "']'", "','", "':'", "'.'", "'='", "'<'", "'>'", "'+'", "'-'",
		"'*'", "'/'", "'%'", "'!'", "'#'", "'@'", "'~'", "'^'",
	}
	staticData.SymbolicNames = []string{
		"", "MISSING", "SEP", "CONST", "TYPE", "EXTERN", "LET", "IF", "ELSE",
		"WHILE", "FOR", "IN", "RETURN", "BREAK", "CONTINUE", "PRIMARY", "AS",
		"CHECK", "EMIT", "SETUP", "EXPECT", "RUN", "TRUE", "FALSE", "COMMAND",
		"DOC_COMMENT", "BLOCK_COMMENT", "UNTERMINATED_COMMENT", "LINE_COMMENT",
		"NEWLINE", "WS", "RESLOC", "TNAME", "IDENT", "NUMBER", "STRING", "ARROW",
		"ELLIPSIS", "DOTDOT", "QQ", "PLUS_EQ", "MINUS_EQ", "STAR_EQ", "SLASH_EQ",
		"PCT_EQ", "EQEQ", "NEQ", "LE", "GE", "ANDAND", "OROR", "PIPE", "QUESTION",
		"LBRACE", "RBRACE", "LPAREN", "RPAREN", "LBRACK", "RBRACK", "COMMA",
		"COLON", "DOT", "ASSIGN", "LT", "GT", "PLUS", "MINUS", "STAR", "SLASH",
		"PERCENT", "BANG", "HASH", "AT", "TILDE", "CARET", "ERROR_CHAR",
	}
	staticData.RuleNames = []string{
		"file", "exprOnly", "topItem", "declaration", "innerDecl", "innerBody",
		"declName", "word", "declBody", "constDecl", "letDecl", "typeAnnotation",
		"typeExpr", "typeArm", "typeAtom", "typeName", "params", "param", "block",
		"blockItem", "statement", "thenPart", "elsePart", "bareCall", "hint",
		"ifTail", "fieldKey", "value", "expr", "postfix", "args", "arg", "primary",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 75, 569, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 1, 0, 5, 0, 68, 8, 0, 10, 0, 12, 0, 71, 9, 0, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 1, 1, 2, 3, 2, 79, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 3, 2, 88, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 94, 8, 2, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 4, 3, 109, 8, 3, 11, 3, 12, 3, 110, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3,
		3, 118, 8, 3, 3, 3, 120, 8, 3, 1, 3, 3, 3, 123, 8, 3, 1, 3, 1, 3, 3, 3,
		127, 8, 3, 1, 3, 3, 3, 130, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 4, 4, 136, 8,
		4, 11, 4, 12, 4, 137, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 145, 8, 4, 3,
		4, 147, 8, 4, 1, 4, 3, 4, 150, 8, 4, 3, 4, 152, 8, 4, 1, 5, 1, 5, 1, 5,
		3, 5, 157, 8, 5, 1, 6, 1, 6, 1, 6, 3, 6, 162, 8, 6, 1, 7, 1, 7, 1, 8, 1,
		8, 1, 8, 3, 8, 169, 8, 8, 1, 9, 1, 9, 1, 9, 3, 9, 174, 8, 9, 1, 9, 1, 9,
		3, 9, 178, 8, 9, 1, 10, 1, 10, 1, 10, 3, 10, 183, 8, 10, 1, 10, 1, 10,
		3, 10, 187, 8, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 5, 12, 195,
		8, 12, 10, 12, 12, 12, 198, 9, 12, 1, 13, 1, 13, 3, 13, 202, 8, 13, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 209, 8, 14, 10, 14, 12, 14, 212,
		9, 14, 1, 14, 1, 14, 3, 14, 216, 8, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		3, 14, 233, 8, 14, 1, 15, 1, 15, 1, 15, 5, 15, 238, 8, 15, 10, 15, 12,
		15, 241, 9, 15, 1, 15, 3, 15, 244, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 5,
		16, 250, 8, 16, 10, 16, 12, 16, 253, 9, 16, 1, 16, 3, 16, 256, 8, 16, 3,
		16, 258, 8, 16, 1, 16, 1, 16, 1, 17, 1, 17, 3, 17, 264, 8, 17, 1, 17, 1,
		17, 3, 17, 268, 8, 17, 1, 18, 1, 18, 5, 18, 272, 8, 18, 10, 18, 12, 18,
		275, 9, 18, 1, 18, 1, 18, 1, 19, 1, 19, 3, 19, 281, 8, 19, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 292, 8, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20,
		304, 8, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 312, 8, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 3, 20, 325, 8, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 3, 20, 337, 8, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 4, 20, 346, 8, 20, 11, 20, 12, 20, 347, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 3, 20, 364, 8, 20, 1, 21, 1, 21, 3, 21, 368, 8, 21, 1, 22, 1,
		22, 1, 22, 3, 22, 373, 8, 22, 1, 23, 1, 23, 5, 23, 377, 8, 23, 10, 23,
		12, 23, 380, 9, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 3, 25, 393, 8, 25, 1, 26, 1, 26, 1, 26, 3, 26,
		398, 8, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 3, 27, 407,
		8, 27, 1, 27, 3, 27, 410, 8, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 28, 3, 28, 420, 8, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28,
		1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28,
		1, 28, 1, 28, 1, 28, 3, 28, 451, 8, 28, 5, 28, 453, 8, 28, 10, 28, 12,
		28, 456, 9, 28, 1, 29, 1, 29, 3, 29, 460, 8, 29, 1, 29, 1, 29, 1, 29, 1,
		29, 1, 29, 5, 29, 467, 8, 29, 10, 29, 12, 29, 470, 9, 29, 1, 29, 3, 29,
		473, 8, 29, 3, 29, 475, 8, 29, 1, 29, 1, 29, 3, 29, 479, 8, 29, 1, 30,
		1, 30, 1, 30, 1, 30, 5, 30, 485, 8, 30, 10, 30, 12, 30, 488, 9, 30, 1,
		30, 3, 30, 491, 8, 30, 3, 30, 493, 8, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1,
		31, 3, 31, 500, 8, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32,
		1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 3, 32, 518,
		8, 32, 1, 32, 3, 32, 521, 8, 32, 4, 32, 523, 8, 32, 11, 32, 12, 32, 524,
		1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 3, 32, 536,
		8, 32, 1, 32, 1, 32, 1, 32, 1, 32, 3, 32, 542, 8, 32, 1, 32, 1, 32, 1,
		32, 1, 32, 1, 32, 1, 32, 1, 32, 3, 32, 551, 8, 32, 1, 32, 5, 32, 554, 8,
		32, 10, 32, 12, 32, 557, 9, 32, 1, 32, 3, 32, 560, 8, 32, 3, 32, 562, 8,
		32, 1, 32, 1, 32, 1, 32, 3, 32, 567, 8, 32, 1, 32, 0, 1, 56, 33, 0, 2,
		4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40,
		42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 0, 12, 3, 0, 4, 4, 15,
		21, 33, 33, 2, 0, 2, 2, 59, 59, 1, 0, 17, 19, 1, 0, 40, 44, 2, 0, 65, 66,
		70, 70, 1, 0, 67, 69, 1, 0, 65, 66, 2, 0, 47, 48, 63, 64, 1, 0, 45, 46,
		1, 0, 22, 23, 1, 0, 31, 33, 1, 0, 73, 74, 665, 0, 69, 1, 0, 0, 0, 2, 74,
		1, 0, 0, 0, 4, 93, 1, 0, 0, 0, 6, 129, 1, 0, 0, 0, 8, 151, 1, 0, 0, 0,
		10, 156, 1, 0, 0, 0, 12, 161, 1, 0, 0, 0, 14, 163, 1, 0, 0, 0, 16, 168,
		1, 0, 0, 0, 18, 170, 1, 0, 0, 0, 20, 179, 1, 0, 0, 0, 22, 188, 1, 0, 0,
		0, 24, 191, 1, 0, 0, 0, 26, 199, 1, 0, 0, 0, 28, 232, 1, 0, 0, 0, 30, 243,
		1, 0, 0, 0, 32, 245, 1, 0, 0, 0, 34, 261, 1, 0, 0, 0, 36, 269, 1, 0, 0,
		0, 38, 278, 1, 0, 0, 0, 40, 363, 1, 0, 0, 0, 42, 367, 1, 0, 0, 0, 44, 372,
		1, 0, 0, 0, 46, 374, 1, 0, 0, 0, 48, 381, 1, 0, 0, 0, 50, 385, 1, 0, 0,
		0, 52, 397, 1, 0, 0, 0, 54, 409, 1, 0, 0, 0, 56, 419, 1, 0, 0, 0, 58, 478,
		1, 0, 0, 0, 60, 480, 1, 0, 0, 0, 62, 499, 1, 0, 0, 0, 64, 566, 1, 0, 0,
		0, 66, 68, 3, 4, 2, 0, 67, 66, 1, 0, 0, 0, 68, 71, 1, 0, 0, 0, 69, 67,
		1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 72, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0,
		72, 73, 5, 0, 0, 1, 73, 1, 1, 0, 0, 0, 74, 75, 3, 56, 28, 0, 75, 76, 5,
		0, 0, 1, 76, 3, 1, 0, 0, 0, 77, 79, 5, 5, 0, 0, 78, 77, 1, 0, 0, 0, 78,
		79, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 94, 3, 6, 3, 0, 81, 94, 3, 18,
		9, 0, 82, 94, 3, 20, 10, 0, 83, 84, 5, 10, 0, 0, 84, 87, 5, 33, 0, 0, 85,
		86, 5, 59, 0, 0, 86, 88, 5, 33, 0, 0, 87, 85, 1, 0, 0, 0, 87, 88, 1, 0,
		0, 0, 88, 89, 1, 0, 0, 0, 89, 90, 5, 11, 0, 0, 90, 91, 3, 56, 28, 0, 91,
		92, 3, 36, 18, 0, 92, 94, 1, 0, 0, 0, 93, 78, 1, 0, 0, 0, 93, 81, 1, 0,
		0, 0, 93, 82, 1, 0, 0, 0, 93, 83, 1, 0, 0, 0, 94, 5, 1, 0, 0, 0, 95, 96,
		5, 4, 0, 0, 96, 97, 3, 12, 6, 0, 97, 98, 5, 62, 0, 0, 98, 99, 3, 24, 12,
		0, 99, 130, 1, 0, 0, 0, 100, 101, 5, 4, 0, 0, 101, 102, 3, 12, 6, 0, 102,
		103, 3, 36, 18, 0, 103, 130, 1, 0, 0, 0, 104, 105, 5, 33, 0, 0, 105, 108,
		3, 12, 6, 0, 106, 107, 5, 59, 0, 0, 107, 109, 3, 12, 6, 0, 108, 106, 1,
		0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 108, 1, 0, 0, 0, 110, 111, 1, 0, 0,
		0, 111, 130, 1, 0, 0, 0, 112, 113, 5, 33, 0, 0, 113, 119, 3, 12, 6, 0,
		114, 120, 3, 22, 11, 0, 115, 117, 3, 32, 16, 0, 116, 118, 3, 22, 11, 0,
		117, 116, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118, 120, 1, 0, 0, 0, 119,
		114, 1, 0, 0, 0, 119, 115, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120, 122,
		1, 0, 0, 0, 121, 123, 3, 16, 8, 0, 122, 121, 1, 0, 0, 0, 122, 123, 1, 0,
		0, 0, 123, 130, 1, 0, 0, 0, 124, 126, 5, 33, 0, 0, 125, 127, 3, 32, 16,
		0, 126, 125, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128,
		130, 3, 36, 18, 0, 129, 95, 1, 0, 0, 0, 129, 100, 1, 0, 0, 0, 129, 104,
		1, 0, 0, 0, 129, 112, 1, 0, 0, 0, 129, 124, 1, 0, 0, 0, 130, 7, 1, 0, 0,
		0, 131, 132, 5, 33, 0, 0, 132, 135, 3, 12, 6, 0, 133, 134, 5, 59, 0, 0,
		134, 136, 3, 12, 6, 0, 135, 133, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137,
		135, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 152, 1, 0, 0, 0, 139, 140,
		5, 33, 0, 0, 140, 146, 3, 12, 6, 0, 141, 147, 3, 22, 11, 0, 142, 144, 3,
		32, 16, 0, 143, 145, 3, 22, 11, 0, 144, 143, 1, 0, 0, 0, 144, 145, 1, 0,
		0, 0, 145, 147, 1, 0, 0, 0, 146, 141, 1, 0, 0, 0, 146, 142, 1, 0, 0, 0,
		146, 147, 1, 0, 0, 0, 147, 149, 1, 0, 0, 0, 148, 150, 3, 10, 5, 0, 149,
		148, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 152, 1, 0, 0, 0, 151, 131,
		1, 0, 0, 0, 151, 139, 1, 0, 0, 0, 152, 9, 1, 0, 0, 0, 153, 154, 5, 62,
		0, 0, 154, 157, 3, 54, 27, 0, 155, 157, 3, 36, 18, 0, 156, 153, 1, 0, 0,
		0, 156, 155, 1, 0, 0, 0, 157, 11, 1, 0, 0, 0, 158, 162, 3, 14, 7, 0, 159,
		162, 5, 31, 0, 0, 160, 162, 5, 32, 0, 0, 161, 158, 1, 0, 0, 0, 161, 159,
		1, 0, 0, 0, 161, 160, 1, 0, 0, 0, 162, 13, 1, 0, 0, 0, 163, 164, 7, 0,
		0, 0, 164, 15, 1, 0, 0, 0, 165, 166, 5, 62, 0, 0, 166, 169, 3, 54, 27,
		0, 167, 169, 3, 36, 18, 0, 168, 165, 1, 0, 0, 0, 168, 167, 1, 0, 0, 0,
		169, 17, 1, 0, 0, 0, 170, 171, 5, 3, 0, 0, 171, 173, 5, 33, 0, 0, 172,
		174, 3, 22, 11, 0, 173, 172, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 177,
		1, 0, 0, 0, 175, 176, 5, 62, 0, 0, 176, 178, 3, 54, 27, 0, 177, 175, 1,
		0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 19, 1, 0, 0, 0, 179, 180, 5, 6, 0,
		0, 180, 182, 5, 33, 0, 0, 181, 183, 3, 22, 11, 0, 182, 181, 1, 0, 0, 0,
		182, 183, 1, 0, 0, 0, 183, 186, 1, 0, 0, 0, 184, 185, 5, 62, 0, 0, 185,
		187, 3, 54, 27, 0, 186, 184, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 21,
		1, 0, 0, 0, 188, 189, 5, 60, 0, 0, 189, 190, 3, 24, 12, 0, 190, 23, 1,
		0, 0, 0, 191, 196, 3, 26, 13, 0, 192, 193, 5, 51, 0, 0, 193, 195, 3, 26,
		13, 0, 194, 192, 1, 0, 0, 0, 195, 198, 1, 0, 0, 0, 196, 194, 1, 0, 0, 0,
		196, 197, 1, 0, 0, 0, 197, 25, 1, 0, 0, 0, 198, 196, 1, 0, 0, 0, 199, 201,
		3, 28, 14, 0, 200, 202, 5, 52, 0, 0, 201, 200, 1, 0, 0, 0, 201, 202, 1,
		0, 0, 0, 202, 27, 1, 0, 0, 0, 203, 215, 3, 30, 15, 0, 204, 205, 5, 63,
		0, 0, 205, 210, 3, 24, 12, 0, 206, 207, 5, 59, 0, 0, 207, 209, 3, 24, 12,
		0, 208, 206, 1, 0, 0, 0, 209, 212, 1, 0, 0, 0, 210, 208, 1, 0, 0, 0, 210,
		211, 1, 0, 0, 0, 211, 213, 1, 0, 0, 0, 212, 210, 1, 0, 0, 0, 213, 214,
		5, 64, 0, 0, 214, 216, 1, 0, 0, 0, 215, 204, 1, 0, 0, 0, 215, 216, 1, 0,
		0, 0, 216, 233, 1, 0, 0, 0, 217, 233, 5, 35, 0, 0, 218, 219, 5, 57, 0,
		0, 219, 220, 3, 24, 12, 0, 220, 221, 5, 58, 0, 0, 221, 233, 1, 0, 0, 0,
		222, 223, 5, 53, 0, 0, 223, 224, 3, 24, 12, 0, 224, 225, 5, 60, 0, 0, 225,
		226, 3, 24, 12, 0, 226, 227, 5, 54, 0, 0, 227, 233, 1, 0, 0, 0, 228, 229,
		5, 55, 0, 0, 229, 230, 3, 24, 12, 0, 230, 231, 5, 56, 0, 0, 231, 233, 1,
		0, 0, 0, 232, 203, 1, 0, 0, 0, 232, 217, 1, 0, 0, 0, 232, 218, 1, 0, 0,
		0, 232, 222, 1, 0, 0, 0, 232, 228, 1, 0, 0, 0, 233, 29, 1, 0, 0, 0, 234,
		239, 5, 33, 0, 0, 235, 236, 5, 61, 0, 0, 236, 238, 5, 33, 0, 0, 237, 235,
		1, 0, 0, 0, 238, 241, 1, 0, 0, 0, 239, 237, 1, 0, 0, 0, 239, 240, 1, 0,
		0, 0, 240, 244, 1, 0, 0, 0, 241, 239, 1, 0, 0, 0, 242, 244, 5, 31, 0, 0,
		243, 234, 1, 0, 0, 0, 243, 242, 1, 0, 0, 0, 244, 31, 1, 0, 0, 0, 245, 257,
		5, 55, 0, 0, 246, 251, 3, 34, 17, 0, 247, 248, 5, 59, 0, 0, 248, 250, 3,
		34, 17, 0, 249, 247, 1, 0, 0, 0, 250, 253, 1, 0, 0, 0, 251, 249, 1, 0,
		0, 0, 251, 252, 1, 0, 0, 0, 252, 255, 1, 0, 0, 0, 253, 251, 1, 0, 0, 0,
		254, 256, 5, 59, 0, 0, 255, 254, 1, 0, 0, 0, 255, 256, 1, 0, 0, 0, 256,
		258, 1, 0, 0, 0, 257, 246, 1, 0, 0, 0, 257, 258, 1, 0, 0, 0, 258, 259,
		1, 0, 0, 0, 259, 260, 5, 56, 0, 0, 260, 33, 1, 0, 0, 0, 261, 263, 5, 33,
		0, 0, 262, 264, 3, 22, 11, 0, 263, 262, 1, 0, 0, 0, 263, 264, 1, 0, 0,
		0, 264, 267, 1, 0, 0, 0, 265, 266, 5, 62, 0, 0, 266, 268, 3, 54, 27, 0,
		267, 265, 1, 0, 0, 0, 267, 268, 1, 0, 0, 0, 268, 35, 1, 0, 0, 0, 269, 273,
		5, 53, 0, 0, 270, 272, 3, 38, 19, 0, 271, 270, 1, 0, 0, 0, 272, 275, 1,
		0, 0, 0, 273, 271, 1, 0, 0, 0, 273, 274, 1, 0, 0, 0, 274, 276, 1, 0, 0,
		0, 275, 273, 1, 0, 0, 0, 276, 277, 5, 54, 0, 0, 277, 37, 1, 0, 0, 0, 278,
		280, 3, 40, 20, 0, 279, 281, 7, 1, 0, 0, 280, 279, 1, 0, 0, 0, 280, 281,
		1, 0, 0, 0, 281, 39, 1, 0, 0, 0, 282, 364, 3, 20, 10, 0, 283, 364, 3, 18,
		9, 0, 284, 285, 5, 7, 0, 0, 285, 286, 5, 55, 0, 0, 286, 287, 3, 56, 28,
		0, 287, 288, 5, 56, 0, 0, 288, 291, 3, 42, 21, 0, 289, 290, 5, 8, 0, 0,
		290, 292, 3, 44, 22, 0, 291, 289, 1, 0, 0, 0, 291, 292, 1, 0, 0, 0, 292,
		364, 1, 0, 0, 0, 293, 294, 5, 9, 0, 0, 294, 295, 5, 55, 0, 0, 295, 296,
		3, 56, 28, 0, 296, 297, 5, 56, 0, 0, 297, 298, 3, 36, 18, 0, 298, 364,
		1, 0, 0, 0, 299, 300, 5, 10, 0, 0, 300, 303, 5, 33, 0, 0, 301, 302, 5,
		59, 0, 0, 302, 304, 5, 33, 0, 0, 303, 301, 1, 0, 0, 0, 303, 304, 1, 0,
		0, 0, 304, 305, 1, 0, 0, 0, 305, 306, 5, 11, 0, 0, 306, 307, 3, 56, 28,
		0, 307, 308, 3, 36, 18, 0, 308, 364, 1, 0, 0, 0, 309, 311, 5, 12, 0, 0,
		310, 312, 3, 54, 27, 0, 311, 310, 1, 0, 0, 0, 311, 312, 1, 0, 0, 0, 312,
		364, 1, 0, 0, 0, 313, 364, 5, 13, 0, 0, 314, 364, 5, 14, 0, 0, 315, 364,
		5, 24, 0, 0, 316, 317, 5, 15, 0, 0, 317, 364, 3, 8, 4, 0, 318, 364, 3,
		8, 4, 0, 319, 320, 5, 16, 0, 0, 320, 321, 3, 24, 12, 0, 321, 322, 5, 62,
		0, 0, 322, 324, 3, 54, 27, 0, 323, 325, 3, 48, 24, 0, 324, 323, 1, 0, 0,
		0, 324, 325, 1, 0, 0, 0, 325, 364, 1, 0, 0, 0, 326, 327, 7, 2, 0, 0, 327,
		364, 3, 36, 18, 0, 328, 329, 5, 20, 0, 0, 329, 364, 3, 56, 28, 0, 330,
		331, 5, 21, 0, 0, 331, 364, 3, 46, 23, 0, 332, 333, 5, 33, 0, 0, 333, 336,
		3, 22, 11, 0, 334, 335, 5, 62, 0, 0, 335, 337, 3, 54, 27, 0, 336, 334,
		1, 0, 0, 0, 336, 337, 1, 0, 0, 0, 337, 364, 1, 0, 0, 0, 338, 339, 3, 52,
		26, 0, 339, 340, 5, 62, 0, 0, 340, 341, 3, 54, 27, 0, 341, 364, 1, 0, 0,
		0, 342, 345, 5, 33, 0, 0, 343, 344, 5, 61, 0, 0, 344, 346, 5, 33, 0, 0,
		345, 343, 1, 0, 0, 0, 346, 347, 1, 0, 0, 0, 347, 345, 1, 0, 0, 0, 347,
		348, 1, 0, 0, 0, 348, 349, 1, 0, 0, 0, 349, 350, 5, 62, 0, 0, 350, 364,
		3, 54, 27, 0, 351, 352, 5, 33, 0, 0, 352, 353, 5, 57, 0, 0, 353, 354, 3,
		62, 31, 0, 354, 355, 5, 58, 0, 0, 355, 356, 5, 62, 0, 0, 356, 357, 3, 54,
		27, 0, 357, 364, 1, 0, 0, 0, 358, 359, 3, 56, 28, 0, 359, 360, 7, 3, 0,
		0, 360, 361, 3, 56, 28, 0, 361, 364, 1, 0, 0, 0, 362, 364, 3, 56, 28, 0,
		363, 282, 1, 0, 0, 0, 363, 283, 1, 0, 0, 0, 363, 284, 1, 0, 0, 0, 363,
		293, 1, 0, 0, 0, 363, 299, 1, 0, 0, 0, 363, 309, 1, 0, 0, 0, 363, 313,
		1, 0, 0, 0, 363, 314, 1, 0, 0, 0, 363, 315, 1, 0, 0, 0, 363, 316, 1, 0,
		0, 0, 363, 318, 1, 0, 0, 0, 363, 319, 1, 0, 0, 0, 363, 326, 1, 0, 0, 0,
		363, 328, 1, 0, 0, 0, 363, 330, 1, 0, 0, 0, 363, 332, 1, 0, 0, 0, 363,
		338, 1, 0, 0, 0, 363, 342, 1, 0, 0, 0, 363, 351, 1, 0, 0, 0, 363, 358,
		1, 0, 0, 0, 363, 362, 1, 0, 0, 0, 364, 41, 1, 0, 0, 0, 365, 368, 3, 36,
		18, 0, 366, 368, 3, 46, 23, 0, 367, 365, 1, 0, 0, 0, 367, 366, 1, 0, 0,
		0, 368, 43, 1, 0, 0, 0, 369, 373, 3, 50, 25, 0, 370, 373, 3, 36, 18, 0,
		371, 373, 3, 46, 23, 0, 372, 369, 1, 0, 0, 0, 372, 370, 1, 0, 0, 0, 372,
		371, 1, 0, 0, 0, 373, 45, 1, 0, 0, 0, 374, 378, 5, 33, 0, 0, 375, 377,
		3, 58, 29, 0, 376, 375, 1, 0, 0, 0, 377, 380, 1, 0, 0, 0, 378, 376, 1,
		0, 0, 0, 378, 379, 1, 0, 0, 0, 379, 47, 1, 0, 0, 0, 380, 378, 1, 0, 0,
		0, 381, 382, 4, 24, 0, 0, 382, 383, 5, 33, 0, 0, 383, 384, 3, 60, 30, 0,
		384, 49, 1, 0, 0, 0, 385, 386, 5, 7, 0, 0, 386, 387, 5, 55, 0, 0, 387,
		388, 3, 56, 28, 0, 388, 389, 5, 56, 0, 0, 389, 392, 3, 42, 21, 0, 390,
		391, 5, 8, 0, 0, 391, 393, 3, 44, 22, 0, 392, 390, 1, 0, 0, 0, 392, 393,
		1, 0, 0, 0, 393, 51, 1, 0, 0, 0, 394, 398, 3, 14, 7, 0, 395, 398, 5, 31,
		0, 0, 396, 398, 5, 35, 0, 0, 397, 394, 1, 0, 0, 0, 397, 395, 1, 0, 0, 0,
		397, 396, 1, 0, 0, 0, 398, 53, 1, 0, 0, 0, 399, 400, 5, 7, 0, 0, 400, 401,
		5, 55, 0, 0, 401, 402, 3, 56, 28, 0, 402, 403, 5, 56, 0, 0, 403, 406, 3,
		54, 27, 0, 404, 405, 5, 8, 0, 0, 405, 407, 3, 54, 27, 0, 406, 404, 1, 0,
		0, 0, 406, 407, 1, 0, 0, 0, 407, 410, 1, 0, 0, 0, 408, 410, 3, 56, 28,
		0, 409, 399, 1, 0, 0, 0, 409, 408, 1, 0, 0, 0, 410, 55, 1, 0, 0, 0, 411,
		412, 6, 28, -1, 0, 412, 413, 7, 4, 0, 0, 413, 420, 3, 56, 28, 13, 414,
		415, 5, 38, 0, 0, 415, 420, 3, 56, 28, 9, 416, 417, 5, 37, 0, 0, 417, 420,
		3, 56, 28, 2, 418, 420, 3, 64, 32, 0, 419, 411, 1, 0, 0, 0, 419, 414, 1,
		0, 0, 0, 419, 416, 1, 0, 0, 0, 419, 418, 1, 0, 0, 0, 420, 454, 1, 0, 0,
		0, 421, 422, 10, 12, 0, 0, 422, 423, 7, 5, 0, 0, 423, 453, 3, 56, 28, 13,
		424, 425, 10, 11, 0, 0, 425, 426, 7, 6, 0, 0, 426, 453, 3, 56, 28, 12,
		427, 428, 10, 8, 0, 0, 428, 429, 7, 7, 0, 0, 429, 453, 3, 56, 28, 9, 430,
		431, 10, 7, 0, 0, 431, 432, 7, 8, 0, 0, 432, 453, 3, 56, 28, 8, 433, 434,
		10, 6, 0, 0, 434, 435, 5, 49, 0, 0, 435, 453, 3, 56, 28, 7, 436, 437, 10,
		5, 0, 0, 437, 438, 5, 50, 0, 0, 438, 453, 3, 56, 28, 6, 439, 440, 10, 4,
		0, 0, 440, 441, 5, 39, 0, 0, 441, 453, 3, 56, 28, 5, 442, 443, 10, 3, 0,
		0, 443, 444, 5, 36, 0, 0, 444, 453, 3, 56, 28, 3, 445, 446, 10, 14, 0,
		0, 446, 453, 3, 58, 29, 0, 447, 448, 10, 10, 0, 0, 448, 450, 5, 38, 0,
		0, 449, 451, 3, 56, 28, 0, 450, 449, 1, 0, 0, 0, 450, 451, 1, 0, 0, 0,
		451, 453, 1, 0, 0, 0, 452, 421, 1, 0, 0, 0, 452, 424, 1, 0, 0, 0, 452,
		427, 1, 0, 0, 0, 452, 430, 1, 0, 0, 0, 452, 433, 1, 0, 0, 0, 452, 436,
		1, 0, 0, 0, 452, 439, 1, 0, 0, 0, 452, 442, 1, 0, 0, 0, 452, 445, 1, 0,
		0, 0, 452, 447, 1, 0, 0, 0, 453, 456, 1, 0, 0, 0, 454, 452, 1, 0, 0, 0,
		454, 455, 1, 0, 0, 0, 455, 57, 1, 0, 0, 0, 456, 454, 1, 0, 0, 0, 457, 459,
		5, 61, 0, 0, 458, 460, 3, 14, 7, 0, 459, 458, 1, 0, 0, 0, 459, 460, 1,
		0, 0, 0, 460, 479, 1, 0, 0, 0, 461, 479, 3, 60, 30, 0, 462, 474, 5, 57,
		0, 0, 463, 468, 3, 62, 31, 0, 464, 465, 5, 59, 0, 0, 465, 467, 3, 62, 31,
		0, 466, 464, 1, 0, 0, 0, 467, 470, 1, 0, 0, 0, 468, 466, 1, 0, 0, 0, 468,
		469, 1, 0, 0, 0, 469, 472, 1, 0, 0, 0, 470, 468, 1, 0, 0, 0, 471, 473,
		5, 59, 0, 0, 472, 471, 1, 0, 0, 0, 472, 473, 1, 0, 0, 0, 473, 475, 1, 0,
		0, 0, 474, 463, 1, 0, 0, 0, 474, 475, 1, 0, 0, 0, 475, 476, 1, 0, 0, 0,
		476, 479, 5, 58, 0, 0, 477, 479, 3, 36, 18, 0, 478, 457, 1, 0, 0, 0, 478,
		461, 1, 0, 0, 0, 478, 462, 1, 0, 0, 0, 478, 477, 1, 0, 0, 0, 479, 59, 1,
		0, 0, 0, 480, 492, 5, 55, 0, 0, 481, 486, 3, 62, 31, 0, 482, 483, 5, 59,
		0, 0, 483, 485, 3, 62, 31, 0, 484, 482, 1, 0, 0, 0, 485, 488, 1, 0, 0,
		0, 486, 484, 1, 0, 0, 0, 486, 487, 1, 0, 0, 0, 487, 490, 1, 0, 0, 0, 488,
		486, 1, 0, 0, 0, 489, 491, 5, 59, 0, 0, 490, 489, 1, 0, 0, 0, 490, 491,
		1, 0, 0, 0, 491, 493, 1, 0, 0, 0, 492, 481, 1, 0, 0, 0, 492, 493, 1, 0,
		0, 0, 493, 494, 1, 0, 0, 0, 494, 495, 5, 56, 0, 0, 495, 61, 1, 0, 0, 0,
		496, 497, 3, 14, 7, 0, 497, 498, 5, 62, 0, 0, 498, 500, 1, 0, 0, 0, 499,
		496, 1, 0, 0, 0, 499, 500, 1, 0, 0, 0, 500, 501, 1, 0, 0, 0, 501, 502,
		3, 54, 27, 0, 502, 63, 1, 0, 0, 0, 503, 567, 5, 34, 0, 0, 504, 567, 5,
		35, 0, 0, 505, 567, 7, 9, 0, 0, 506, 507, 5, 21, 0, 0, 507, 567, 3, 36,
		18, 0, 508, 567, 5, 33, 0, 0, 509, 567, 5, 31, 0, 0, 510, 567, 5, 32, 0,
		0, 511, 512, 5, 71, 0, 0, 512, 567, 7, 10, 0, 0, 513, 514, 5, 72, 0, 0,
		514, 567, 5, 33, 0, 0, 515, 520, 7, 11, 0, 0, 516, 518, 5, 66, 0, 0, 517,
		516, 1, 0, 0, 0, 517, 518, 1, 0, 0, 0, 518, 519, 1, 0, 0, 0, 519, 521,
		5, 34, 0, 0, 520, 517, 1, 0, 0, 0, 520, 521, 1, 0, 0, 0, 521, 523, 1, 0,
		0, 0, 522, 515, 1, 0, 0, 0, 523, 524, 1, 0, 0, 0, 524, 522, 1, 0, 0, 0,
		524, 525, 1, 0, 0, 0, 525, 567, 1, 0, 0, 0, 526, 527, 5, 55, 0, 0, 527,
		528, 3, 54, 27, 0, 528, 529, 5, 56, 0, 0, 529, 567, 1, 0, 0, 0, 530, 531,
		5, 57, 0, 0, 531, 532, 5, 10, 0, 0, 532, 535, 5, 33, 0, 0, 533, 534, 5,
		59, 0, 0, 534, 536, 5, 33, 0, 0, 535, 533, 1, 0, 0, 0, 535, 536, 1, 0,
		0, 0, 536, 537, 1, 0, 0, 0, 537, 538, 5, 11, 0, 0, 538, 541, 3, 56, 28,
		0, 539, 540, 5, 7, 0, 0, 540, 542, 3, 56, 28, 0, 541, 539, 1, 0, 0, 0,
		541, 542, 1, 0, 0, 0, 542, 543, 1, 0, 0, 0, 543, 544, 5, 60, 0, 0, 544,
		545, 3, 54, 27, 0, 545, 546, 5, 58, 0, 0, 546, 567, 1, 0, 0, 0, 547, 561,
		5, 57, 0, 0, 548, 555, 3, 54, 27, 0, 549, 551, 5, 59, 0, 0, 550, 549, 1,
		0, 0, 0, 550, 551, 1, 0, 0, 0, 551, 552, 1, 0, 0, 0, 552, 554, 3, 54, 27,
		0, 553, 550, 1, 0, 0, 0, 554, 557, 1, 0, 0, 0, 555, 553, 1, 0, 0, 0, 555,
		556, 1, 0, 0, 0, 556, 559, 1, 0, 0, 0, 557, 555, 1, 0, 0, 0, 558, 560,
		5, 59, 0, 0, 559, 558, 1, 0, 0, 0, 559, 560, 1, 0, 0, 0, 560, 562, 1, 0,
		0, 0, 561, 548, 1, 0, 0, 0, 561, 562, 1, 0, 0, 0, 562, 563, 1, 0, 0, 0,
		563, 567, 5, 58, 0, 0, 564, 567, 3, 36, 18, 0, 565, 567, 5, 1, 0, 0, 566,
		503, 1, 0, 0, 0, 566, 504, 1, 0, 0, 0, 566, 505, 1, 0, 0, 0, 566, 506,
		1, 0, 0, 0, 566, 508, 1, 0, 0, 0, 566, 509, 1, 0, 0, 0, 566, 510, 1, 0,
		0, 0, 566, 511, 1, 0, 0, 0, 566, 513, 1, 0, 0, 0, 566, 522, 1, 0, 0, 0,
		566, 526, 1, 0, 0, 0, 566, 530, 1, 0, 0, 0, 566, 547, 1, 0, 0, 0, 566,
		564, 1, 0, 0, 0, 566, 565, 1, 0, 0, 0, 567, 65, 1, 0, 0, 0, 73, 69, 78,
		87, 93, 110, 117, 119, 122, 126, 129, 137, 144, 146, 149, 151, 156, 161,
		168, 173, 177, 182, 186, 196, 201, 210, 215, 232, 239, 243, 251, 255, 257,
		263, 267, 273, 280, 291, 303, 311, 324, 336, 347, 363, 367, 372, 378, 392,
		397, 406, 409, 419, 450, 452, 454, 459, 468, 472, 474, 478, 486, 490, 492,
		499, 517, 520, 524, 535, 541, 550, 555, 559, 561, 566,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// MMSParserInit initializes any static state used to implement MMSParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMMSParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MMSParserInit() {
	staticData := &MMSParserParserStaticData
	staticData.once.Do(mmsparserParserInit)
}

// NewMMSParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMMSParser(input antlr.TokenStream) *MMSParser {
	MMSParserInit()
	this := new(MMSParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &MMSParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "MMSParser.g4"

	return this
}

// MMSParser tokens.
const (
	MMSParserEOF                  = antlr.TokenEOF
	MMSParserMISSING              = 1
	MMSParserSEP                  = 2
	MMSParserCONST                = 3
	MMSParserTYPE                 = 4
	MMSParserEXTERN               = 5
	MMSParserLET                  = 6
	MMSParserIF                   = 7
	MMSParserELSE                 = 8
	MMSParserWHILE                = 9
	MMSParserFOR                  = 10
	MMSParserIN                   = 11
	MMSParserRETURN               = 12
	MMSParserBREAK                = 13
	MMSParserCONTINUE             = 14
	MMSParserPRIMARY              = 15
	MMSParserAS                   = 16
	MMSParserCHECK                = 17
	MMSParserEMIT                 = 18
	MMSParserSETUP                = 19
	MMSParserEXPECT               = 20
	MMSParserRUN                  = 21
	MMSParserTRUE                 = 22
	MMSParserFALSE                = 23
	MMSParserCOMMAND              = 24
	MMSParserDOC_COMMENT          = 25
	MMSParserBLOCK_COMMENT        = 26
	MMSParserUNTERMINATED_COMMENT = 27
	MMSParserLINE_COMMENT         = 28
	MMSParserNEWLINE              = 29
	MMSParserWS                   = 30
	MMSParserRESLOC               = 31
	MMSParserTNAME                = 32
	MMSParserIDENT                = 33
	MMSParserNUMBER               = 34
	MMSParserSTRING               = 35
	MMSParserARROW                = 36
	MMSParserELLIPSIS             = 37
	MMSParserDOTDOT               = 38
	MMSParserQQ                   = 39
	MMSParserPLUS_EQ              = 40
	MMSParserMINUS_EQ             = 41
	MMSParserSTAR_EQ              = 42
	MMSParserSLASH_EQ             = 43
	MMSParserPCT_EQ               = 44
	MMSParserEQEQ                 = 45
	MMSParserNEQ                  = 46
	MMSParserLE                   = 47
	MMSParserGE                   = 48
	MMSParserANDAND               = 49
	MMSParserOROR                 = 50
	MMSParserPIPE                 = 51
	MMSParserQUESTION             = 52
	MMSParserLBRACE               = 53
	MMSParserRBRACE               = 54
	MMSParserLPAREN               = 55
	MMSParserRPAREN               = 56
	MMSParserLBRACK               = 57
	MMSParserRBRACK               = 58
	MMSParserCOMMA                = 59
	MMSParserCOLON                = 60
	MMSParserDOT                  = 61
	MMSParserASSIGN               = 62
	MMSParserLT                   = 63
	MMSParserGT                   = 64
	MMSParserPLUS                 = 65
	MMSParserMINUS                = 66
	MMSParserSTAR                 = 67
	MMSParserSLASH                = 68
	MMSParserPERCENT              = 69
	MMSParserBANG                 = 70
	MMSParserHASH                 = 71
	MMSParserAT                   = 72
	MMSParserTILDE                = 73
	MMSParserCARET                = 74
	MMSParserERROR_CHAR           = 75
)

// MMSParser rules.
const (
	MMSParserRULE_file           = 0
	MMSParserRULE_exprOnly       = 1
	MMSParserRULE_topItem        = 2
	MMSParserRULE_declaration    = 3
	MMSParserRULE_innerDecl      = 4
	MMSParserRULE_innerBody      = 5
	MMSParserRULE_declName       = 6
	MMSParserRULE_word           = 7
	MMSParserRULE_declBody       = 8
	MMSParserRULE_constDecl      = 9
	MMSParserRULE_letDecl        = 10
	MMSParserRULE_typeAnnotation = 11
	MMSParserRULE_typeExpr       = 12
	MMSParserRULE_typeArm        = 13
	MMSParserRULE_typeAtom       = 14
	MMSParserRULE_typeName       = 15
	MMSParserRULE_params         = 16
	MMSParserRULE_param          = 17
	MMSParserRULE_block          = 18
	MMSParserRULE_blockItem      = 19
	MMSParserRULE_statement      = 20
	MMSParserRULE_thenPart       = 21
	MMSParserRULE_elsePart       = 22
	MMSParserRULE_bareCall       = 23
	MMSParserRULE_hint           = 24
	MMSParserRULE_ifTail         = 25
	MMSParserRULE_fieldKey       = 26
	MMSParserRULE_value          = 27
	MMSParserRULE_expr           = 28
	MMSParserRULE_postfix        = 29
	MMSParserRULE_args           = 30
	MMSParserRULE_arg            = 31
	MMSParserRULE_primary        = 32
)

// IFileContext is an interface to support dynamic dispatch.
type IFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllTopItem() []ITopItemContext
	TopItem(i int) ITopItemContext

	// IsFileContext differentiates from other interfaces.
	IsFileContext()
}

type FileContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFileContext() *FileContext {
	var p = new(FileContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_file
	return p
}

func InitEmptyFileContext(p *FileContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_file
}

func (*FileContext) IsFileContext() {}

func NewFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileContext {
	var p = new(FileContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_file

	return p
}

func (s *FileContext) GetParser() antlr.Parser { return s.parser }

func (s *FileContext) EOF() antlr.TerminalNode {
	return s.GetToken(MMSParserEOF, 0)
}

func (s *FileContext) AllTopItem() []ITopItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITopItemContext); ok {
			len++
		}
	}

	tst := make([]ITopItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITopItemContext); ok {
			tst[i] = t.(ITopItemContext)
			i++
		}
	}

	return tst
}

func (s *FileContext) TopItem(i int) ITopItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITopItemContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITopItemContext)
}

func (s *FileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MMSParser) File() (localctx IFileContext) {
	localctx = NewFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MMSParserRULE_file)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8589935736) != 0 {
		{
			p.SetState(66)
			p.TopItem()
		}

		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(72)
		p.Match(MMSParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprOnlyContext is an interface to support dynamic dispatch.
type IExprOnlyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	EOF() antlr.TerminalNode

	// IsExprOnlyContext differentiates from other interfaces.
	IsExprOnlyContext()
}

type ExprOnlyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprOnlyContext() *ExprOnlyContext {
	var p = new(ExprOnlyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_exprOnly
	return p
}

func InitEmptyExprOnlyContext(p *ExprOnlyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_exprOnly
}

func (*ExprOnlyContext) IsExprOnlyContext() {}

func NewExprOnlyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprOnlyContext {
	var p = new(ExprOnlyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_exprOnly

	return p
}

func (s *ExprOnlyContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprOnlyContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprOnlyContext) EOF() antlr.TerminalNode {
	return s.GetToken(MMSParserEOF, 0)
}

func (s *ExprOnlyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprOnlyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MMSParser) ExprOnly() (localctx IExprOnlyContext) {
	localctx = NewExprOnlyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MMSParserRULE_exprOnly)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)
		p.expr(0)
	}
	{
		p.SetState(75)
		p.Match(MMSParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITopItemContext is an interface to support dynamic dispatch.
type ITopItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTopItemContext differentiates from other interfaces.
	IsTopItemContext()
}

type TopItemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopItemContext() *TopItemContext {
	var p = new(TopItemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_topItem
	return p
}

func InitEmptyTopItemContext(p *TopItemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_topItem
}

func (*TopItemContext) IsTopItemContext() {}

func NewTopItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopItemContext {
	var p = new(TopItemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_topItem

	return p
}

func (s *TopItemContext) GetParser() antlr.Parser { return s.parser }

func (s *TopItemContext) CopyAll(ctx *TopItemContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TopItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DeclItemContext struct {
	TopItemContext
}

func NewDeclItemContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclItemContext {
	var p = new(DeclItemContext)

	InitEmptyTopItemContext(&p.TopItemContext)
	p.parser = parser
	p.CopyAll(ctx.(*TopItemContext))

	return p
}

func (s *DeclItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclItemContext) Declaration() IDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *DeclItemContext) EXTERN() antlr.TerminalNode {
	return s.GetToken(MMSParserEXTERN, 0)
}

type ForItemContext struct {
	TopItemContext
}

func NewForItemContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForItemContext {
	var p = new(ForItemContext)

	InitEmptyTopItemContext(&p.TopItemContext)
	p.parser = parser
	p.CopyAll(ctx.(*TopItemContext))

	return p
}

func (s *ForItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForItemContext) FOR() antlr.TerminalNode {
	return s.GetToken(MMSParserFOR, 0)
}

func (s *ForItemContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(MMSParserIDENT)
}

func (s *ForItemContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(MMSParserIDENT, i)
}

func (s *ForItemContext) IN() antlr.TerminalNode {
	return s.GetToken(MMSParserIN, 0)
}

func (s *ForItemContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForItemContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForItemContext) COMMA() antlr.TerminalNode {
	return s.GetToken(MMSParserCOMMA, 0)
}

type LetItemContext struct {
	TopItemContext
}

func NewLetItemContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LetItemContext {
	var p = new(LetItemContext)

	InitEmptyTopItemContext(&p.TopItemContext)
	p.parser = parser
	p.CopyAll(ctx.(*TopItemContext))

	return p
}

func (s *LetItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetItemContext) LetDecl() ILetDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILetDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILetDeclContext)
}

type ConstItemContext struct {
	TopItemContext
}

func NewConstItemContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConstItemContext {
	var p = new(ConstItemContext)

	InitEmptyTopItemContext(&p.TopItemContext)
	p.parser = parser
	p.CopyAll(ctx.(*TopItemContext))

	return p
}

func (s *ConstItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstItemContext) ConstDecl() IConstDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstDeclContext)
}

func (p *MMSParser) TopItem() (localctx ITopItemContext) {
	localctx = NewTopItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MMSParserRULE_topItem)
	var _la int

	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MMSParserTYPE, MMSParserEXTERN, MMSParserIDENT:
		localctx = NewDeclItemContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MMSParserEXTERN {
			{
				p.SetState(77)
				p.Match(MMSParserEXTERN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(80)
			p.Declaration()
		}

	case MMSParserCONST:
		localctx = NewConstItemContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(81)
			p.ConstDecl()
		}

	case MMSParserLET:
		localctx = NewLetItemContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(82)
			p.LetDecl()
		}

	case MMSParserFOR:
		localctx = NewForItemContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(83)
			p.Match(MMSParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(84)
			p.Match(MMSParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MMSParserCOMMA {
			{
				p.SetState(85)
				p.Match(MMSParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(86)
				p.Match(MMSParserIDENT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(89)
			p.Match(MMSParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(90)
			p.expr(0)
		}
		{
			p.SetState(91)
			p.Block()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetKind returns the kind token.
	GetKind() antlr.Token

	// SetKind sets the kind token.
	SetKind(antlr.Token)

	// GetName returns the name rule contexts.
	GetName() IDeclNameContext

	// SetName sets the name rule contexts.
	SetName(IDeclNameContext)

	// Getter signatures
	ASSIGN() antlr.TerminalNode
	TypeExpr() ITypeExprContext
	TYPE() antlr.TerminalNode
	AllDeclName() []IDeclNameContext
	DeclName(i int) IDeclNameContext
	Block() IBlockContext
	IDENT() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	TypeAnnotation() ITypeAnnotationContext
	Params() IParamsContext
	DeclBody() IDeclBodyContext

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	kind   antlr.Token
	name   IDeclNameContext
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_declaration
	return p
}

func InitEmptyDeclarationContext(p *DeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_declaration
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) GetKind() antlr.Token { return s.kind }

func (s *DeclarationContext) SetKind(v antlr.Token) { s.kind = v }

func (s *DeclarationContext) GetName() IDeclNameContext { return s.name }

func (s *DeclarationContext) SetName(v IDeclNameContext) { s.name = v }

func (s *DeclarationContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(MMSParserASSIGN, 0)
}

func (s *DeclarationContext) TypeExpr() ITypeExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeExprContext)
}

func (s *DeclarationContext) TYPE() antlr.TerminalNode {
	return s.GetToken(MMSParserTYPE, 0)
}

func (s *DeclarationContext) AllDeclName() []IDeclNameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclNameContext); ok {
			len++
		}
	}

	tst := make([]IDeclNameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclNameContext); ok {
			tst[i] = t.(IDeclNameContext)
			i++
		}
	}

	return tst
}

func (s *DeclarationContext) DeclName(i int) IDeclNameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclNameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclNameContext)
}

func (s *DeclarationContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *DeclarationContext) IDENT() antlr.TerminalNode {
	return s.GetToken(MMSParserIDENT, 0)
}

func (s *DeclarationContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MMSParserCOMMA)
}

func (s *DeclarationContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MMSParserCOMMA, i)
}

func (s *DeclarationContext) TypeAnnotation() ITypeAnnotationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeAnnotationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeAnnotationContext)
}

func (s *DeclarationContext) Params() IParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamsContext)
}

func (s *DeclarationContext) DeclBody() IDeclBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclBodyContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MMSParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MMSParserRULE_declaration)
	var _la int

	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(95)

			var _m = p.Match(MMSParserTYPE)

			localctx.(*DeclarationContext).kind = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(96)

			var _x = p.DeclName()

			localctx.(*DeclarationContext).name = _x
		}
		{
			p.SetState(97)
			p.Match(MMSParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(98)
			p.TypeExpr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(100)

			var _m = p.Match(MMSParserTYPE)

			localctx.(*DeclarationContext).kind = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(101)

			var _x = p.DeclName()

			localctx.(*DeclarationContext).name = _x
		}
		{
			p.SetState(102)
			p.Block()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(104)

			var _m = p.Match(MMSParserIDENT)

			localctx.(*DeclarationContext).kind = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(105)

			var _x = p.DeclName()

			localctx.(*DeclarationContext).name = _x
		}
		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == MMSParserCOMMA {
			{
				p.SetState(106)
				p.Match(MMSParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(107)
				p.DeclName()
			}

			p.SetState(110)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(112)

			var _m = p.Match(MMSParserIDENT)

			localctx.(*DeclarationContext).kind = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(113)

			var _x = p.DeclName()

			localctx.(*DeclarationContext).name = _x
		}
		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		switch p.GetTokenStream().LA(1) {
		case MMSParserCOLON:
			{
				p.SetState(114)
				p.TypeAnnotation()
			}

		case MMSParserLPAREN:
			{
				p.SetState(115)
				p.Params()
			}
			p.SetState(117)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == MMSParserCOLON {
				{
					p.SetState(116)
					p.TypeAnnotation()
				}

			}

		case MMSParserEOF, MMSParserCONST, MMSParserTYPE, MMSParserEXTERN, MMSParserLET, MMSParserFOR, MMSParserIDENT, MMSParserLBRACE, MMSParserASSIGN:

		default:
		}
		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MMSParserLBRACE || _la == MMSParserASSIGN {
			{
				p.SetState(121)
				p.DeclBody()
			}

		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(124)

			var _m = p.Match(MMSParserIDENT)

			localctx.(*DeclarationContext).kind = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MMSParserLPAREN {
			{
				p.SetState(125)
				p.Params()
			}

		}
		{
			p.SetState(128)
			p.Block()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInnerDeclContext is an interface to support dynamic dispatch.
type IInnerDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetKind returns the kind token.
	GetKind() antlr.Token

	// SetKind sets the kind token.
	SetKind(antlr.Token)

	// GetName returns the name rule contexts.
	GetName() IDeclNameContext

	// SetName sets the name rule contexts.
	SetName(IDeclNameContext)

	// Getter signatures
	IDENT() antlr.TerminalNode
	AllDeclName() []IDeclNameContext
	DeclName(i int) IDeclNameContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	TypeAnnotation() ITypeAnnotationContext
	Params() IParamsContext
	InnerBody() IInnerBodyContext

	// IsInnerDeclContext differentiates from other interfaces.
	IsInnerDeclContext()
}

type InnerDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	kind   antlr.Token
	name   IDeclNameContext
}

func NewEmptyInnerDeclContext() *InnerDeclContext {
	var p = new(InnerDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_innerDecl
	return p
}

func InitEmptyInnerDeclContext(p *InnerDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_innerDecl
}

func (*InnerDeclContext) IsInnerDeclContext() {}

func NewInnerDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InnerDeclContext {
	var p = new(InnerDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_innerDecl

	return p
}

func (s *InnerDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *InnerDeclContext) GetKind() antlr.Token { return s.kind }

func (s *InnerDeclContext) SetKind(v antlr.Token) { s.kind = v }

func (s *InnerDeclContext) GetName() IDeclNameContext { return s.name }

func (s *InnerDeclContext) SetName(v IDeclNameContext) { s.name = v }

func (s *InnerDeclContext) IDENT() antlr.TerminalNode {
	return s.GetToken(MMSParserIDENT, 0)
}

func (s *InnerDeclContext) AllDeclName() []IDeclNameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclNameContext); ok {
			len++
		}
	}

	tst := make([]IDeclNameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclNameContext); ok {
			tst[i] = t.(IDeclNameContext)
			i++
		}
	}

	return tst
}

func (s *InnerDeclContext) DeclName(i int) IDeclNameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclNameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclNameContext)
}

func (s *InnerDeclContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MMSParserCOMMA)
}

func (s *InnerDeclContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MMSParserCOMMA, i)
}

func (s *InnerDeclContext) TypeAnnotation() ITypeAnnotationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeAnnotationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeAnnotationContext)
}

func (s *InnerDeclContext) Params() IParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamsContext)
}

func (s *InnerDeclContext) InnerBody() IInnerBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInnerBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInnerBodyContext)
}

func (s *InnerDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InnerDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MMSParser) InnerDecl() (localctx IInnerDeclContext) {
	localctx = NewInnerDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MMSParserRULE_innerDecl)
	var _la int

	var _alt int

	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(131)

			var _m = p.Match(MMSParserIDENT)

			localctx.(*InnerDeclContext).kind = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(132)

			var _x = p.DeclName()

			localctx.(*InnerDeclContext).name = _x
		}
		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(133)
					p.Match(MMSParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(134)
					p.DeclName()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(137)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(139)

			var _m = p.Match(MMSParserIDENT)

			localctx.(*InnerDeclContext).kind = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(140)

			var _x = p.DeclName()

			localctx.(*InnerDeclContext).name = _x
		}
		p.SetState(146)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(141)
				p.TypeAnnotation()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) == 2 {
			{
				p.SetState(142)
				p.Params()
			}
			p.SetState(144)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == MMSParserCOLON {
				{
					p.SetState(143)
					p.TypeAnnotation()
				}

			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(149)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(148)
				p.InnerBody()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInnerBodyContext is an interface to support dynamic dispatch.
type IInnerBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ASSIGN() antlr.TerminalNode
	Value() IValueContext
	Block() IBlockContext

	// IsInnerBodyContext differentiates from other interfaces.
	IsInnerBodyContext()
}

type InnerBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInnerBodyContext() *InnerBodyContext {
	var p = new(InnerBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_innerBody
	return p
}

func InitEmptyInnerBodyContext(p *InnerBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_innerBody
}

func (*InnerBodyContext) IsInnerBodyContext() {}

func NewInnerBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InnerBodyContext {
	var p = new(InnerBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_innerBody

	return p
}

func (s *InnerBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *InnerBodyContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(MMSParserASSIGN, 0)
}

func (s *InnerBodyContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *InnerBodyContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *InnerBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InnerBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MMSParser) InnerBody() (localctx IInnerBodyContext) {
	localctx = NewInnerBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MMSParserRULE_innerBody)
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MMSParserASSIGN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(153)
			p.Match(MMSParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(154)
			p.Value()
		}

	case MMSParserLBRACE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(155)
			p.Block()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclNameContext is an interface to support dynamic dispatch.
type IDeclNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Word() IWordContext
	RESLOC() antlr.TerminalNode
	TNAME() antlr.TerminalNode

	// IsDeclNameContext differentiates from other interfaces.
	IsDeclNameContext()
}

type DeclNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclNameContext() *DeclNameContext {
	var p = new(DeclNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_declName
	return p
}

func InitEmptyDeclNameContext(p *DeclNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_declName
}

func (*DeclNameContext) IsDeclNameContext() {}

func NewDeclNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclNameContext {
	var p = new(DeclNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_declName

	return p
}

func (s *DeclNameContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclNameContext) Word() IWordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWordContext)
}

func (s *DeclNameContext) RESLOC() antlr.TerminalNode {
	return s.GetToken(MMSParserRESLOC, 0)
}

func (s *DeclNameContext) TNAME() antlr.TerminalNode {
	return s.GetToken(MMSParserTNAME, 0)
}

func (s *DeclNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MMSParser) DeclName() (localctx IDeclNameContext) {
	localctx = NewDeclNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MMSParserRULE_declName)
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MMSParserTYPE, MMSParserPRIMARY, MMSParserAS, MMSParserCHECK, MMSParserEMIT, MMSParserSETUP, MMSParserEXPECT, MMSParserRUN, MMSParserIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(158)
			p.Word()
		}

	case MMSParserRESLOC:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(159)
			p.Match(MMSParserRESLOC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MMSParserTNAME:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(160)
			p.Match(MMSParserTNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWordContext is an interface to support dynamic dispatch.
type IWordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	TYPE() antlr.TerminalNode
	PRIMARY() antlr.TerminalNode
	AS() antlr.TerminalNode
	CHECK() antlr.TerminalNode
	EMIT() antlr.TerminalNode
	SETUP() antlr.TerminalNode
	EXPECT() antlr.TerminalNode
	RUN() antlr.TerminalNode

	// IsWordContext differentiates from other interfaces.
	IsWordContext()
}

type WordContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWordContext() *WordContext {
	var p = new(WordContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_word
	return p
}

func InitEmptyWordContext(p *WordContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_word
}

func (*WordContext) IsWordContext() {}

func NewWordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WordContext {
	var p = new(WordContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_word

	return p
}

func (s *WordContext) GetParser() antlr.Parser { return s.parser }

func (s *WordContext) IDENT() antlr.TerminalNode {
	return s.GetToken(MMSParserIDENT, 0)
}

func (s *WordContext) TYPE() antlr.TerminalNode {
	return s.GetToken(MMSParserTYPE, 0)
}

func (s *WordContext) PRIMARY() antlr.TerminalNode {
	return s.GetToken(MMSParserPRIMARY, 0)
}

func (s *WordContext) AS() antlr.TerminalNode {
	return s.GetToken(MMSParserAS, 0)
}

func (s *WordContext) CHECK() antlr.TerminalNode {
	return s.GetToken(MMSParserCHECK, 0)
}

func (s *WordContext) EMIT() antlr.TerminalNode {
	return s.GetToken(MMSParserEMIT, 0)
}

func (s *WordContext) SETUP() antlr.TerminalNode {
	return s.GetToken(MMSParserSETUP, 0)
}

func (s *WordContext) EXPECT() antlr.TerminalNode {
	return s.GetToken(MMSParserEXPECT, 0)
}

func (s *WordContext) RUN() antlr.TerminalNode {
	return s.GetToken(MMSParserRUN, 0)
}

func (s *WordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MMSParser) Word() (localctx IWordContext) {
	localctx = NewWordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MMSParserRULE_word)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8594096144) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclBodyContext is an interface to support dynamic dispatch.
type IDeclBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ASSIGN() antlr.TerminalNode
	Value() IValueContext
	Block() IBlockContext

	// IsDeclBodyContext differentiates from other interfaces.
	IsDeclBodyContext()
}

type DeclBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclBodyContext() *DeclBodyContext {
	var p = new(DeclBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_declBody
	return p
}

func InitEmptyDeclBodyContext(p *DeclBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_declBody
}

func (*DeclBodyContext) IsDeclBodyContext() {}

func NewDeclBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclBodyContext {
	var p = new(DeclBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_declBody

	return p
}

func (s *DeclBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclBodyContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(MMSParserASSIGN, 0)
}

func (s *DeclBodyContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *DeclBodyContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *DeclBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MMSParser) DeclBody() (localctx IDeclBodyContext) {
	localctx = NewDeclBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MMSParserRULE_declBody)
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MMSParserASSIGN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(165)
			p.Match(MMSParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(166)
			p.Value()
		}

	case MMSParserLBRACE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(167)
			p.Block()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstDeclContext is an interface to support dynamic dispatch.
type IConstDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONST() antlr.TerminalNode
	IDENT() antlr.TerminalNode
	TypeAnnotation() ITypeAnnotationContext
	ASSIGN() antlr.TerminalNode
	Value() IValueContext

	// IsConstDeclContext differentiates from other interfaces.
	IsConstDeclContext()
}

type ConstDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstDeclContext() *ConstDeclContext {
	var p = new(ConstDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_constDecl
	return p
}

func InitEmptyConstDeclContext(p *ConstDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_constDecl
}

func (*ConstDeclContext) IsConstDeclContext() {}

func NewConstDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstDeclContext {
	var p = new(ConstDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_constDecl

	return p
}

func (s *ConstDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstDeclContext) CONST() antlr.TerminalNode {
	return s.GetToken(MMSParserCONST, 0)
}

func (s *ConstDeclContext) IDENT() antlr.TerminalNode {
	return s.GetToken(MMSParserIDENT, 0)
}

func (s *ConstDeclContext) TypeAnnotation() ITypeAnnotationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeAnnotationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeAnnotationContext)
}

func (s *ConstDeclContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(MMSParserASSIGN, 0)
}

func (s *ConstDeclContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ConstDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MMSParser) ConstDecl() (localctx IConstDeclContext) {
	localctx = NewConstDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MMSParserRULE_constDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(170)
		p.Match(MMSParserCONST)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(171)
		p.Match(MMSParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MMSParserCOLON {
		{
			p.SetState(172)
			p.TypeAnnotation()
		}

	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MMSParserASSIGN {
		{
			p.SetState(175)
			p.Match(MMSParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(176)
			p.Value()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILetDeclContext is an interface to support dynamic dispatch.
type ILetDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LET() antlr.TerminalNode
	IDENT() antlr.TerminalNode
	TypeAnnotation() ITypeAnnotationContext
	ASSIGN() antlr.TerminalNode
	Value() IValueContext

	// IsLetDeclContext differentiates from other interfaces.
	IsLetDeclContext()
}

type LetDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLetDeclContext() *LetDeclContext {
	var p = new(LetDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_letDecl
	return p
}

func InitEmptyLetDeclContext(p *LetDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_letDecl
}

func (*LetDeclContext) IsLetDeclContext() {}

func NewLetDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LetDeclContext {
	var p = new(LetDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_letDecl

	return p
}

func (s *LetDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *LetDeclContext) LET() antlr.TerminalNode {
	return s.GetToken(MMSParserLET, 0)
}

func (s *LetDeclContext) IDENT() antlr.TerminalNode {
	return s.GetToken(MMSParserIDENT, 0)
}

func (s *LetDeclContext) TypeAnnotation() ITypeAnnotationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeAnnotationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeAnnotationContext)
}

func (s *LetDeclContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(MMSParserASSIGN, 0)
}

func (s *LetDeclContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *LetDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MMSParser) LetDecl() (localctx ILetDeclContext) {
	localctx = NewLetDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MMSParserRULE_letDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)
		p.Match(MMSParserLET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(180)
		p.Match(MMSParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MMSParserCOLON {
		{
			p.SetState(181)
			p.TypeAnnotation()
		}

	}
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MMSParserASSIGN {
		{
			p.SetState(184)
			p.Match(MMSParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(185)
			p.Value()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeAnnotationContext is an interface to support dynamic dispatch.
type ITypeAnnotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COLON() antlr.TerminalNode
	TypeExpr() ITypeExprContext

	// IsTypeAnnotationContext differentiates from other interfaces.
	IsTypeAnnotationContext()
}

type TypeAnnotationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeAnnotationContext() *TypeAnnotationContext {
	var p = new(TypeAnnotationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_typeAnnotation
	return p
}

func InitEmptyTypeAnnotationContext(p *TypeAnnotationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_typeAnnotation
}

func (*TypeAnnotationContext) IsTypeAnnotationContext() {}

func NewTypeAnnotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeAnnotationContext {
	var p = new(TypeAnnotationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_typeAnnotation

	return p
}

func (s *TypeAnnotationContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeAnnotationContext) COLON() antlr.TerminalNode {
	return s.GetToken(MMSParserCOLON, 0)
}

func (s *TypeAnnotationContext) TypeExpr() ITypeExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeExprContext)
}

func (s *TypeAnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeAnnotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MMSParser) TypeAnnotation() (localctx ITypeAnnotationContext) {
	localctx = NewTypeAnnotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, MMSParserRULE_typeAnnotation)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(188)
		p.Match(MMSParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(189)
		p.TypeExpr()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeExprContext is an interface to support dynamic dispatch.
type ITypeExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTypeArm() []ITypeArmContext
	TypeArm(i int) ITypeArmContext
	AllPIPE() []antlr.TerminalNode
	PIPE(i int) antlr.TerminalNode

	// IsTypeExprContext differentiates from other interfaces.
	IsTypeExprContext()
}

type TypeExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeExprContext() *TypeExprContext {
	var p = new(TypeExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_typeExpr
	return p
}

func InitEmptyTypeExprContext(p *TypeExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_typeExpr
}

func (*TypeExprContext) IsTypeExprContext() {}

func NewTypeExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeExprContext {
	var p = new(TypeExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_typeExpr

	return p
}

func (s *TypeExprContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeExprContext) AllTypeArm() []ITypeArmContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeArmContext); ok {
			len++
		}
	}

	tst := make([]ITypeArmContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeArmContext); ok {
			tst[i] = t.(ITypeArmContext)
			i++
		}
	}

	return tst
}

func (s *TypeExprContext) TypeArm(i int) ITypeArmContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeArmContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeArmContext)
}

func (s *TypeExprContext) AllPIPE() []antlr.TerminalNode {
	return s.GetTokens(MMSParserPIPE)
}

func (s *TypeExprContext) PIPE(i int) antlr.TerminalNode {
	return s.GetToken(MMSParserPIPE, i)
}

func (s *TypeExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MMSParser) TypeExpr() (localctx ITypeExprContext) {
	localctx = NewTypeExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, MMSParserRULE_typeExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(191)
		p.TypeArm()
	}
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MMSParserPIPE {
		{
			p.SetState(192)
			p.Match(MMSParserPIPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(193)
			p.TypeArm()
		}

		p.SetState(198)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeArmContext is an interface to support dynamic dispatch.
type ITypeArmContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TypeAtom() ITypeAtomContext
	QUESTION() antlr.TerminalNode

	// IsTypeArmContext differentiates from other interfaces.
	IsTypeArmContext()
}

type TypeArmContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeArmContext() *TypeArmContext {
	var p = new(TypeArmContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_typeArm
	return p
}

func InitEmptyTypeArmContext(p *TypeArmContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_typeArm
}

func (*TypeArmContext) IsTypeArmContext() {}

func NewTypeArmContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeArmContext {
	var p = new(TypeArmContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_typeArm

	return p
}

func (s *TypeArmContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeArmContext) TypeAtom() ITypeAtomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeAtomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeAtomContext)
}

func (s *TypeArmContext) QUESTION() antlr.TerminalNode {
	return s.GetToken(MMSParserQUESTION, 0)
}

func (s *TypeArmContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeArmContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MMSParser) TypeArm() (localctx ITypeArmContext) {
	localctx = NewTypeArmContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, MMSParserRULE_typeArm)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(199)
		p.TypeAtom()
	}
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MMSParserQUESTION {
		{
			p.SetState(200)
			p.Match(MMSParserQUESTION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeAtomContext is an interface to support dynamic dispatch.
type ITypeAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTypeAtomContext differentiates from other interfaces.
	IsTypeAtomContext()
}

type TypeAtomContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeAtomContext() *TypeAtomContext {
	var p = new(TypeAtomContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_typeAtom
	return p
}

func InitEmptyTypeAtomContext(p *TypeAtomContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_typeAtom
}

func (*TypeAtomContext) IsTypeAtomContext() {}

func NewTypeAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeAtomContext {
	var p = new(TypeAtomContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_typeAtom

	return p
}

func (s *TypeAtomContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeAtomContext) CopyAll(ctx *TypeAtomContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TypeAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeAtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NamedTypeContext struct {
	TypeAtomContext
}

func NewNamedTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NamedTypeContext {
	var p = new(NamedTypeContext)

	InitEmptyTypeAtomContext(&p.TypeAtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeAtomContext))

	return p
}

func (s *NamedTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamedTypeContext) TypeName() ITypeNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *NamedTypeContext) LT() antlr.TerminalNode {
	return s.GetToken(MMSParserLT, 0)
}

func (s *NamedTypeContext) AllTypeExpr() []ITypeExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeExprContext); ok {
			len++
		}
	}

	tst := make([]ITypeExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeExprContext); ok {
			tst[i] = t.(ITypeExprContext)
			i++
		}
	}

	return tst
}

func (s *NamedTypeContext) TypeExpr(i int) ITypeExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeExprContext)
}

func (s *NamedTypeContext) GT() antlr.TerminalNode {
	return s.GetToken(MMSParserGT, 0)
}

func (s *NamedTypeContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MMSParserCOMMA)
}

func (s *NamedTypeContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MMSParserCOMMA, i)
}

type LiteralTypeContext struct {
	TypeAtomContext
}

func NewLiteralTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralTypeContext {
	var p = new(LiteralTypeContext)

	InitEmptyTypeAtomContext(&p.TypeAtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeAtomContext))

	return p
}

func (s *LiteralTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralTypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(MMSParserSTRING, 0)
}

type MapTypeContext struct {
	TypeAtomContext
}

func NewMapTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MapTypeContext {
	var p = new(MapTypeContext)

	InitEmptyTypeAtomContext(&p.TypeAtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeAtomContext))

	return p
}

func (s *MapTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapTypeContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(MMSParserLBRACE, 0)
}

func (s *MapTypeContext) AllTypeExpr() []ITypeExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeExprContext); ok {
			len++
		}
	}

	tst := make([]ITypeExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeExprContext); ok {
			tst[i] = t.(ITypeExprContext)
			i++
		}
	}

	return tst
}

func (s *MapTypeContext) TypeExpr(i int) ITypeExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeExprContext)
}

func (s *MapTypeContext) COLON() antlr.TerminalNode {
	return s.GetToken(MMSParserCOLON, 0)
}

func (s *MapTypeContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(MMSParserRBRACE, 0)
}

type ParenTypeContext struct {
	TypeAtomContext
}

func NewParenTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenTypeContext {
	var p = new(ParenTypeContext)

	InitEmptyTypeAtomContext(&p.TypeAtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeAtomContext))

	return p
}

func (s *ParenTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenTypeContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MMSParserLPAREN, 0)
}

func (s *ParenTypeContext) TypeExpr() ITypeExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeExprContext)
}

func (s *ParenTypeContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MMSParserRPAREN, 0)
}

type ListTypeContext struct {
	TypeAtomContext
}

func NewListTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListTypeContext {
	var p = new(ListTypeContext)

	InitEmptyTypeAtomContext(&p.TypeAtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeAtomContext))

	return p
}

func (s *ListTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListTypeContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(MMSParserLBRACK, 0)
}

func (s *ListTypeContext) TypeExpr() ITypeExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeExprContext)
}

func (s *ListTypeContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(MMSParserRBRACK, 0)
}

func (p *MMSParser) TypeAtom() (localctx ITypeAtomContext) {
	localctx = NewTypeAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, MMSParserRULE_typeAtom)
	var _la int

	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MMSParserRESLOC, MMSParserIDENT:
		localctx = NewNamedTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(203)
			p.TypeName()
		}
		p.SetState(215)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MMSParserLT {
			{
				p.SetState(204)
				p.Match(MMSParserLT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(205)
				p.TypeExpr()
			}
			p.SetState(210)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == MMSParserCOMMA {
				{
					p.SetState(206)
					p.Match(MMSParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(207)
					p.TypeExpr()
				}

				p.SetState(212)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(213)
				p.Match(MMSParserGT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case MMSParserSTRING:
		localctx = NewLiteralTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(217)
			p.Match(MMSParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MMSParserLBRACK:
		localctx = NewListTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(218)
			p.Match(MMSParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(219)
			p.TypeExpr()
		}
		{
			p.SetState(220)
			p.Match(MMSParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MMSParserLBRACE:
		localctx = NewMapTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(222)
			p.Match(MMSParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(223)
			p.TypeExpr()
		}
		{
			p.SetState(224)
			p.Match(MMSParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(225)
			p.TypeExpr()
		}
		{
			p.SetState(226)
			p.Match(MMSParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MMSParserLPAREN:
		localctx = NewParenTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(228)
			p.Match(MMSParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(229)
			p.TypeExpr()
		}
		{
			p.SetState(230)
			p.Match(MMSParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeNameContext is an interface to support dynamic dispatch.
type ITypeNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENT() []antlr.TerminalNode
	IDENT(i int) antlr.TerminalNode
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode
	RESLOC() antlr.TerminalNode

	// IsTypeNameContext differentiates from other interfaces.
	IsTypeNameContext()
}

type TypeNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeNameContext() *TypeNameContext {
	var p = new(TypeNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_typeName
	return p
}

func InitEmptyTypeNameContext(p *TypeNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_typeName
}

func (*TypeNameContext) IsTypeNameContext() {}

func NewTypeNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeNameContext {
	var p = new(TypeNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_typeName

	return p
}

func (s *TypeNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeNameContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(MMSParserIDENT)
}

func (s *TypeNameContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(MMSParserIDENT, i)
}

func (s *TypeNameContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(MMSParserDOT)
}

func (s *TypeNameContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(MMSParserDOT, i)
}

func (s *TypeNameContext) RESLOC() antlr.TerminalNode {
	return s.GetToken(MMSParserRESLOC, 0)
}

func (s *TypeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MMSParser) TypeName() (localctx ITypeNameContext) {
	localctx = NewTypeNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, MMSParserRULE_typeName)
	var _la int

	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MMSParserIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(234)
			p.Match(MMSParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(239)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MMSParserDOT {
			{
				p.SetState(235)
				p.Match(MMSParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(236)
				p.Match(MMSParserIDENT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(241)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case MMSParserRESLOC:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(242)
			p.Match(MMSParserRESLOC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParamsContext is an interface to support dynamic dispatch.
type IParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllParam() []IParamContext
	Param(i int) IParamContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsParamsContext differentiates from other interfaces.
	IsParamsContext()
}

type ParamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamsContext() *ParamsContext {
	var p = new(ParamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_params
	return p
}

func InitEmptyParamsContext(p *ParamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_params
}

func (*ParamsContext) IsParamsContext() {}

func NewParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamsContext {
	var p = new(ParamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_params

	return p
}

func (s *ParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamsContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MMSParserLPAREN, 0)
}

func (s *ParamsContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MMSParserRPAREN, 0)
}

func (s *ParamsContext) AllParam() []IParamContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParamContext); ok {
			len++
		}
	}

	tst := make([]IParamContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParamContext); ok {
			tst[i] = t.(IParamContext)
			i++
		}
	}

	return tst
}

func (s *ParamsContext) Param(i int) IParamContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamContext)
}

func (s *ParamsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MMSParserCOMMA)
}

func (s *ParamsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MMSParserCOMMA, i)
}

func (s *ParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MMSParser) Params() (localctx IParamsContext) {
	localctx = NewParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, MMSParserRULE_params)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(245)
		p.Match(MMSParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MMSParserIDENT {
		{
			p.SetState(246)
			p.Param()
		}
		p.SetState(251)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(247)
					p.Match(MMSParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(248)
					p.Param()
				}

			}
			p.SetState(253)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(255)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MMSParserCOMMA {
			{
				p.SetState(254)
				p.Match(MMSParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(259)
		p.Match(MMSParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParamContext is an interface to support dynamic dispatch.
type IParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	TypeAnnotation() ITypeAnnotationContext
	ASSIGN() antlr.TerminalNode
	Value() IValueContext

	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamContext() *ParamContext {
	var p = new(ParamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_param
	return p
}

func InitEmptyParamContext(p *ParamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_param
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) IDENT() antlr.TerminalNode {
	return s.GetToken(MMSParserIDENT, 0)
}

func (s *ParamContext) TypeAnnotation() ITypeAnnotationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeAnnotationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeAnnotationContext)
}

func (s *ParamContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(MMSParserASSIGN, 0)
}

func (s *ParamContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MMSParser) Param() (localctx IParamContext) {
	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, MMSParserRULE_param)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(261)
		p.Match(MMSParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(263)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MMSParserCOLON {
		{
			p.SetState(262)
			p.TypeAnnotation()
		}

	}
	p.SetState(267)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MMSParserASSIGN {
		{
			p.SetState(265)
			p.Match(MMSParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(266)
			p.Value()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllBlockItem() []IBlockItemContext
	BlockItem(i int) IBlockItemContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(MMSParserLBRACE, 0)
}

func (s *BlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(MMSParserRBRACE, 0)
}

func (s *BlockContext) AllBlockItem() []IBlockItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockItemContext); ok {
			len++
		}
	}

	tst := make([]IBlockItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockItemContext); ok {
			tst[i] = t.(IBlockItemContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) BlockItem(i int) IBlockItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockItemContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockItemContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MMSParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, MMSParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(269)
		p.Match(MMSParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&189151663271966426) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&995) != 0) {
		{
			p.SetState(270)
			p.BlockItem()
		}

		p.SetState(275)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(276)
		p.Match(MMSParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockItemContext is an interface to support dynamic dispatch.
type IBlockItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Statement() IStatementContext
	COMMA() antlr.TerminalNode
	SEP() antlr.TerminalNode

	// IsBlockItemContext differentiates from other interfaces.
	IsBlockItemContext()
}

type BlockItemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockItemContext() *BlockItemContext {
	var p = new(BlockItemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_blockItem
	return p
}

func InitEmptyBlockItemContext(p *BlockItemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_blockItem
}

func (*BlockItemContext) IsBlockItemContext() {}

func NewBlockItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockItemContext {
	var p = new(BlockItemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_blockItem

	return p
}

func (s *BlockItemContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockItemContext) Statement() IStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockItemContext) COMMA() antlr.TerminalNode {
	return s.GetToken(MMSParserCOMMA, 0)
}

func (s *BlockItemContext) SEP() antlr.TerminalNode {
	return s.GetToken(MMSParserSEP, 0)
}

func (s *BlockItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MMSParser) BlockItem() (localctx IBlockItemContext) {
	localctx = NewBlockItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, MMSParserRULE_blockItem)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(278)
		p.Statement()
	}
	p.SetState(280)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MMSParserSEP || _la == MMSParserCOMMA {
		{
			p.SetState(279)
			_la = p.GetTokenStream().LA(1)

			if !(_la == MMSParserSEP || _la == MMSParserCOMMA) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) CopyAll(ctx *StatementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExpectStmtContext struct {
	StatementContext
}

func NewExpectStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpectStmtContext {
	var p = new(ExpectStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *ExpectStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpectStmtContext) EXPECT() antlr.TerminalNode {
	return s.GetToken(MMSParserEXPECT, 0)
}

func (s *ExpectStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

type ForStmtContext struct {
	StatementContext
}

func NewForStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForStmtContext {
	var p = new(ForStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *ForStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStmtContext) FOR() antlr.TerminalNode {
	return s.GetToken(MMSParserFOR, 0)
}

func (s *ForStmtContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(MMSParserIDENT)
}

func (s *ForStmtContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(MMSParserIDENT, i)
}

func (s *ForStmtContext) IN() antlr.TerminalNode {
	return s.GetToken(MMSParserIN, 0)
}

func (s *ForStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForStmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForStmtContext) COMMA() antlr.TerminalNode {
	return s.GetToken(MMSParserCOMMA, 0)
}

type WhileStmtContext struct {
	StatementContext
}

func NewWhileStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WhileStmtContext {
	var p = new(WhileStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *WhileStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStmtContext) WHILE() antlr.TerminalNode {
	return s.GetToken(MMSParserWHILE, 0)
}

func (s *WhileStmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MMSParserLPAREN, 0)
}

func (s *WhileStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *WhileStmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MMSParserRPAREN, 0)
}

func (s *WhileStmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

type PrimaryStmtContext struct {
	StatementContext
}

func NewPrimaryStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrimaryStmtContext {
	var p = new(PrimaryStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *PrimaryStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryStmtContext) PRIMARY() antlr.TerminalNode {
	return s.GetToken(MMSParserPRIMARY, 0)
}

func (s *PrimaryStmtContext) InnerDecl() IInnerDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInnerDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInnerDeclContext)
}

type ConstStmtContext struct {
	StatementContext
}

func NewConstStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConstStmtContext {
	var p = new(ConstStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *ConstStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstStmtContext) ConstDecl() IConstDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstDeclContext)
}

type DeclStmtContext struct {
	StatementContext
}

func NewDeclStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclStmtContext {
	var p = new(DeclStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *DeclStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclStmtContext) InnerDecl() IInnerDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInnerDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInnerDeclContext)
}

type LetStmtContext struct {
	StatementContext
}

func NewLetStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LetStmtContext {
	var p = new(LetStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *LetStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetStmtContext) LetDecl() ILetDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILetDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILetDeclContext)
}

type ReturnStmtContext struct {
	StatementContext
}

func NewReturnStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnStmtContext {
	var p = new(ReturnStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *ReturnStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStmtContext) RETURN() antlr.TerminalNode {
	return s.GetToken(MMSParserRETURN, 0)
}

func (s *ReturnStmtContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

type FieldStmtContext struct {
	StatementContext
}

func NewFieldStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldStmtContext {
	var p = new(FieldStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *FieldStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldStmtContext) FieldKey() IFieldKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldKeyContext)
}

func (s *FieldStmtContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(MMSParserASSIGN, 0)
}

func (s *FieldStmtContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

type SectionStmtContext struct {
	StatementContext
	section antlr.Token
}

func NewSectionStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SectionStmtContext {
	var p = new(SectionStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *SectionStmtContext) GetSection() antlr.Token { return s.section }

func (s *SectionStmtContext) SetSection(v antlr.Token) { s.section = v }

func (s *SectionStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SectionStmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *SectionStmtContext) CHECK() antlr.TerminalNode {
	return s.GetToken(MMSParserCHECK, 0)
}

func (s *SectionStmtContext) EMIT() antlr.TerminalNode {
	return s.GetToken(MMSParserEMIT, 0)
}

func (s *SectionStmtContext) SETUP() antlr.TerminalNode {
	return s.GetToken(MMSParserSETUP, 0)
}

type RunStmtContext struct {
	StatementContext
}

func NewRunStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RunStmtContext {
	var p = new(RunStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *RunStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RunStmtContext) RUN() antlr.TerminalNode {
	return s.GetToken(MMSParserRUN, 0)
}

func (s *RunStmtContext) BareCall() IBareCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBareCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBareCallContext)
}

type ExprStmtContext struct {
	StatementContext
}

func NewExprStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprStmtContext {
	var p = new(ExprStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *ExprStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

type IfStmtContext struct {
	StatementContext
}

func NewIfStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfStmtContext {
	var p = new(IfStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) IF() antlr.TerminalNode {
	return s.GetToken(MMSParserIF, 0)
}

func (s *IfStmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MMSParserLPAREN, 0)
}

func (s *IfStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfStmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MMSParserRPAREN, 0)
}

func (s *IfStmtContext) ThenPart() IThenPartContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IThenPartContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IThenPartContext)
}

func (s *IfStmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(MMSParserELSE, 0)
}

func (s *IfStmtContext) ElsePart() IElsePartContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElsePartContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElsePartContext)
}

type CommandStmtContext struct {
	StatementContext
}

func NewCommandStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CommandStmtContext {
	var p = new(CommandStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *CommandStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandStmtContext) COMMAND() antlr.TerminalNode {
	return s.GetToken(MMSParserCOMMAND, 0)
}

type CompoundAssignStmtContext struct {
	StatementContext
	op antlr.Token
}

func NewCompoundAssignStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompoundAssignStmtContext {
	var p = new(CompoundAssignStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *CompoundAssignStmtContext) GetOp() antlr.Token { return s.op }

func (s *CompoundAssignStmtContext) SetOp(v antlr.Token) { s.op = v }

func (s *CompoundAssignStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompoundAssignStmtContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *CompoundAssignStmtContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CompoundAssignStmtContext) PLUS_EQ() antlr.TerminalNode {
	return s.GetToken(MMSParserPLUS_EQ, 0)
}

func (s *CompoundAssignStmtContext) MINUS_EQ() antlr.TerminalNode {
	return s.GetToken(MMSParserMINUS_EQ, 0)
}

func (s *CompoundAssignStmtContext) STAR_EQ() antlr.TerminalNode {
	return s.GetToken(MMSParserSTAR_EQ, 0)
}

func (s *CompoundAssignStmtContext) SLASH_EQ() antlr.TerminalNode {
	return s.GetToken(MMSParserSLASH_EQ, 0)
}

func (s *CompoundAssignStmtContext) PCT_EQ() antlr.TerminalNode {
	return s.GetToken(MMSParserPCT_EQ, 0)
}

type BreakStmtContext struct {
	StatementContext
}

func NewBreakStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BreakStmtContext {
	var p = new(BreakStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *BreakStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStmtContext) BREAK() antlr.TerminalNode {
	return s.GetToken(MMSParserBREAK, 0)
}

type AsStmtContext struct {
	StatementContext
}

func NewAsStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AsStmtContext {
	var p = new(AsStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *AsStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsStmtContext) AS() antlr.TerminalNode {
	return s.GetToken(MMSParserAS, 0)
}

func (s *AsStmtContext) TypeExpr() ITypeExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeExprContext)
}

func (s *AsStmtContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(MMSParserASSIGN, 0)
}

func (s *AsStmtContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *AsStmtContext) Hint() IHintContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHintContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHintContext)
}

type PathAssignStmtContext struct {
	StatementContext
}

func NewPathAssignStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PathAssignStmtContext {
	var p = new(PathAssignStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *PathAssignStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PathAssignStmtContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(MMSParserIDENT)
}

func (s *PathAssignStmtContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(MMSParserIDENT, i)
}

func (s *PathAssignStmtContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(MMSParserASSIGN, 0)
}

func (s *PathAssignStmtContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *PathAssignStmtContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(MMSParserDOT)
}

func (s *PathAssignStmtContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(MMSParserDOT, i)
}

type ContinueStmtContext struct {
	StatementContext
}

func NewContinueStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContinueStmtContext {
	var p = new(ContinueStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *ContinueStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStmtContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(MMSParserCONTINUE, 0)
}

type IndexAssignStmtContext struct {
	StatementContext
}

func NewIndexAssignStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexAssignStmtContext {
	var p = new(IndexAssignStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *IndexAssignStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexAssignStmtContext) IDENT() antlr.TerminalNode {
	return s.GetToken(MMSParserIDENT, 0)
}

func (s *IndexAssignStmtContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(MMSParserLBRACK, 0)
}

func (s *IndexAssignStmtContext) Arg() IArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgContext)
}

func (s *IndexAssignStmtContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(MMSParserRBRACK, 0)
}

func (s *IndexAssignStmtContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(MMSParserASSIGN, 0)
}

func (s *IndexAssignStmtContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

type FieldDeclStmtContext struct {
	StatementContext
}

func NewFieldDeclStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldDeclStmtContext {
	var p = new(FieldDeclStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *FieldDeclStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldDeclStmtContext) IDENT() antlr.TerminalNode {
	return s.GetToken(MMSParserIDENT, 0)
}

func (s *FieldDeclStmtContext) TypeAnnotation() ITypeAnnotationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeAnnotationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeAnnotationContext)
}

func (s *FieldDeclStmtContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(MMSParserASSIGN, 0)
}

func (s *FieldDeclStmtContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (p *MMSParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, MMSParserRULE_statement)
	var _la int

	p.SetState(363)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext()) {
	case 1:
		localctx = NewLetStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(282)
			p.LetDecl()
		}

	case 2:
		localctx = NewConstStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(283)
			p.ConstDecl()
		}

	case 3:
		localctx = NewIfStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(284)
			p.Match(MMSParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(285)
			p.Match(MMSParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(286)
			p.expr(0)
		}
		{
			p.SetState(287)
			p.Match(MMSParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(288)
			p.ThenPart()
		}
		p.SetState(291)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MMSParserELSE {
			{
				p.SetState(289)
				p.Match(MMSParserELSE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(290)
				p.ElsePart()
			}

		}

	case 4:
		localctx = NewWhileStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(293)
			p.Match(MMSParserWHILE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(294)
			p.Match(MMSParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(295)
			p.expr(0)
		}
		{
			p.SetState(296)
			p.Match(MMSParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(297)
			p.Block()
		}

	case 5:
		localctx = NewForStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(299)
			p.Match(MMSParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(300)
			p.Match(MMSParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(303)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MMSParserCOMMA {
			{
				p.SetState(301)
				p.Match(MMSParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(302)
				p.Match(MMSParserIDENT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(305)
			p.Match(MMSParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(306)
			p.expr(0)
		}
		{
			p.SetState(307)
			p.Block()
		}

	case 6:
		localctx = NewReturnStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(309)
			p.Match(MMSParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(311)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(310)
				p.Value()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 7:
		localctx = NewBreakStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(313)
			p.Match(MMSParserBREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewContinueStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(314)
			p.Match(MMSParserCONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		localctx = NewCommandStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(315)
			p.Match(MMSParserCOMMAND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		localctx = NewPrimaryStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(316)
			p.Match(MMSParserPRIMARY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(317)
			p.InnerDecl()
		}

	case 11:
		localctx = NewDeclStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(318)
			p.InnerDecl()
		}

	case 12:
		localctx = NewAsStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(319)
			p.Match(MMSParserAS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(320)
			p.TypeExpr()
		}
		{
			p.SetState(321)
			p.Match(MMSParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(322)
			p.Value()
		}
		p.SetState(324)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(323)
				p.Hint()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 13:
		localctx = NewSectionStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(326)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*SectionStmtContext).section = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&917504) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*SectionStmtContext).section = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(327)
			p.Block()
		}

	case 14:
		localctx = NewExpectStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(328)
			p.Match(MMSParserEXPECT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(329)
			p.expr(0)
		}

	case 15:
		localctx = NewRunStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(330)
			p.Match(MMSParserRUN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(331)
			p.BareCall()
		}

	case 16:
		localctx = NewFieldDeclStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(332)
			p.Match(MMSParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(333)
			p.TypeAnnotation()
		}
		p.SetState(336)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MMSParserASSIGN {
			{
				p.SetState(334)
				p.Match(MMSParserASSIGN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(335)
				p.Value()
			}

		}

	case 17:
		localctx = NewFieldStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(338)
			p.FieldKey()
		}
		{
			p.SetState(339)
			p.Match(MMSParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(340)
			p.Value()
		}

	case 18:
		localctx = NewPathAssignStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(342)
			p.Match(MMSParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(345)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == MMSParserDOT {
			{
				p.SetState(343)
				p.Match(MMSParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(344)
				p.Match(MMSParserIDENT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(347)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(349)
			p.Match(MMSParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(350)
			p.Value()
		}

	case 19:
		localctx = NewIndexAssignStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(351)
			p.Match(MMSParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(352)
			p.Match(MMSParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(353)
			p.Arg()
		}
		{
			p.SetState(354)
			p.Match(MMSParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(355)
			p.Match(MMSParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(356)
			p.Value()
		}

	case 20:
		localctx = NewCompoundAssignStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(358)
			p.expr(0)
		}
		{
			p.SetState(359)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*CompoundAssignStmtContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&34084860461056) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*CompoundAssignStmtContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(360)
			p.expr(0)
		}

	case 21:
		localctx = NewExprStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(362)
			p.expr(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IThenPartContext is an interface to support dynamic dispatch.
type IThenPartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext
	BareCall() IBareCallContext

	// IsThenPartContext differentiates from other interfaces.
	IsThenPartContext()
}

type ThenPartContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyThenPartContext() *ThenPartContext {
	var p = new(ThenPartContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_thenPart
	return p
}

func InitEmptyThenPartContext(p *ThenPartContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_thenPart
}

func (*ThenPartContext) IsThenPartContext() {}

func NewThenPartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ThenPartContext {
	var p = new(ThenPartContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_thenPart

	return p
}

func (s *ThenPartContext) GetParser() antlr.Parser { return s.parser }

func (s *ThenPartContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ThenPartContext) BareCall() IBareCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBareCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBareCallContext)
}

func (s *ThenPartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThenPartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MMSParser) ThenPart() (localctx IThenPartContext) {
	localctx = NewThenPartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, MMSParserRULE_thenPart)
	p.SetState(367)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MMSParserLBRACE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(365)
			p.Block()
		}

	case MMSParserIDENT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(366)
			p.BareCall()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElsePartContext is an interface to support dynamic dispatch.
type IElsePartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IfTail() IIfTailContext
	Block() IBlockContext
	BareCall() IBareCallContext

	// IsElsePartContext differentiates from other interfaces.
	IsElsePartContext()
}

type ElsePartContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElsePartContext() *ElsePartContext {
	var p = new(ElsePartContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_elsePart
	return p
}

func InitEmptyElsePartContext(p *ElsePartContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_elsePart
}

func (*ElsePartContext) IsElsePartContext() {}

func NewElsePartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElsePartContext {
	var p = new(ElsePartContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_elsePart

	return p
}

func (s *ElsePartContext) GetParser() antlr.Parser { return s.parser }

func (s *ElsePartContext) IfTail() IIfTailContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfTailContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfTailContext)
}

func (s *ElsePartContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ElsePartContext) BareCall() IBareCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBareCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBareCallContext)
}

func (s *ElsePartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElsePartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MMSParser) ElsePart() (localctx IElsePartContext) {
	localctx = NewElsePartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, MMSParserRULE_elsePart)
	p.SetState(372)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MMSParserIF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(369)
			p.IfTail()
		}

	case MMSParserLBRACE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(370)
			p.Block()
		}

	case MMSParserIDENT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(371)
			p.BareCall()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBareCallContext is an interface to support dynamic dispatch.
type IBareCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	AllPostfix() []IPostfixContext
	Postfix(i int) IPostfixContext

	// IsBareCallContext differentiates from other interfaces.
	IsBareCallContext()
}

type BareCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBareCallContext() *BareCallContext {
	var p = new(BareCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_bareCall
	return p
}

func InitEmptyBareCallContext(p *BareCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_bareCall
}

func (*BareCallContext) IsBareCallContext() {}

func NewBareCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BareCallContext {
	var p = new(BareCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_bareCall

	return p
}

func (s *BareCallContext) GetParser() antlr.Parser { return s.parser }

func (s *BareCallContext) IDENT() antlr.TerminalNode {
	return s.GetToken(MMSParserIDENT, 0)
}

func (s *BareCallContext) AllPostfix() []IPostfixContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPostfixContext); ok {
			len++
		}
	}

	tst := make([]IPostfixContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPostfixContext); ok {
			tst[i] = t.(IPostfixContext)
			i++
		}
	}

	return tst
}

func (s *BareCallContext) Postfix(i int) IPostfixContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPostfixContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPostfixContext)
}

func (s *BareCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BareCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MMSParser) BareCall() (localctx IBareCallContext) {
	localctx = NewBareCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, MMSParserRULE_bareCall)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(374)
		p.Match(MMSParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(378)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(375)
				p.Postfix()
			}

		}
		p.SetState(380)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHintContext is an interface to support dynamic dispatch.
type IHintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	Args() IArgsContext

	// IsHintContext differentiates from other interfaces.
	IsHintContext()
}

type HintContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHintContext() *HintContext {
	var p = new(HintContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_hint
	return p
}

func InitEmptyHintContext(p *HintContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_hint
}

func (*HintContext) IsHintContext() {}

func NewHintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HintContext {
	var p = new(HintContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_hint

	return p
}

func (s *HintContext) GetParser() antlr.Parser { return s.parser }

func (s *HintContext) IDENT() antlr.TerminalNode {
	return s.GetToken(MMSParserIDENT, 0)
}

func (s *HintContext) Args() IArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
}

func (s *HintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MMSParser) Hint() (localctx IHintContext) {
	localctx = NewHintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, MMSParserRULE_hint)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(381)

	if !(p.word("hint")) {
		p.SetError(antlr.NewFailedPredicateException(p, "p.word(\"hint\")", ""))
		goto errorExit
	}
	{
		p.SetState(382)
		p.Match(MMSParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(383)
		p.Args()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfTailContext is an interface to support dynamic dispatch.
type IIfTailContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expr() IExprContext
	RPAREN() antlr.TerminalNode
	ThenPart() IThenPartContext
	ELSE() antlr.TerminalNode
	ElsePart() IElsePartContext

	// IsIfTailContext differentiates from other interfaces.
	IsIfTailContext()
}

type IfTailContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfTailContext() *IfTailContext {
	var p = new(IfTailContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_ifTail
	return p
}

func InitEmptyIfTailContext(p *IfTailContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_ifTail
}

func (*IfTailContext) IsIfTailContext() {}

func NewIfTailContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfTailContext {
	var p = new(IfTailContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_ifTail

	return p
}

func (s *IfTailContext) GetParser() antlr.Parser { return s.parser }

func (s *IfTailContext) IF() antlr.TerminalNode {
	return s.GetToken(MMSParserIF, 0)
}

func (s *IfTailContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MMSParserLPAREN, 0)
}

func (s *IfTailContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfTailContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MMSParserRPAREN, 0)
}

func (s *IfTailContext) ThenPart() IThenPartContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IThenPartContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IThenPartContext)
}

func (s *IfTailContext) ELSE() antlr.TerminalNode {
	return s.GetToken(MMSParserELSE, 0)
}

func (s *IfTailContext) ElsePart() IElsePartContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElsePartContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElsePartContext)
}

func (s *IfTailContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfTailContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MMSParser) IfTail() (localctx IIfTailContext) {
	localctx = NewIfTailContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, MMSParserRULE_ifTail)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(385)
		p.Match(MMSParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(386)
		p.Match(MMSParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(387)
		p.expr(0)
	}
	{
		p.SetState(388)
		p.Match(MMSParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(389)
		p.ThenPart()
	}
	p.SetState(392)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MMSParserELSE {
		{
			p.SetState(390)
			p.Match(MMSParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(391)
			p.ElsePart()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldKeyContext is an interface to support dynamic dispatch.
type IFieldKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Word() IWordContext
	RESLOC() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsFieldKeyContext differentiates from other interfaces.
	IsFieldKeyContext()
}

type FieldKeyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldKeyContext() *FieldKeyContext {
	var p = new(FieldKeyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_fieldKey
	return p
}

func InitEmptyFieldKeyContext(p *FieldKeyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_fieldKey
}

func (*FieldKeyContext) IsFieldKeyContext() {}

func NewFieldKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldKeyContext {
	var p = new(FieldKeyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_fieldKey

	return p
}

func (s *FieldKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldKeyContext) Word() IWordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWordContext)
}

func (s *FieldKeyContext) RESLOC() antlr.TerminalNode {
	return s.GetToken(MMSParserRESLOC, 0)
}

func (s *FieldKeyContext) STRING() antlr.TerminalNode {
	return s.GetToken(MMSParserSTRING, 0)
}

func (s *FieldKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MMSParser) FieldKey() (localctx IFieldKeyContext) {
	localctx = NewFieldKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, MMSParserRULE_fieldKey)
	p.SetState(397)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MMSParserTYPE, MMSParserPRIMARY, MMSParserAS, MMSParserCHECK, MMSParserEMIT, MMSParserSETUP, MMSParserEXPECT, MMSParserRUN, MMSParserIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(394)
			p.Word()
		}

	case MMSParserRESLOC:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(395)
			p.Match(MMSParserRESLOC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MMSParserSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(396)
			p.Match(MMSParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expr() IExprContext
	RPAREN() antlr.TerminalNode
	AllValue() []IValueContext
	Value(i int) IValueContext
	ELSE() antlr.TerminalNode

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) IF() antlr.TerminalNode {
	return s.GetToken(MMSParserIF, 0)
}

func (s *ValueContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MMSParserLPAREN, 0)
}

func (s *ValueContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ValueContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MMSParserRPAREN, 0)
}

func (s *ValueContext) AllValue() []IValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueContext); ok {
			len++
		}
	}

	tst := make([]IValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueContext); ok {
			tst[i] = t.(IValueContext)
			i++
		}
	}

	return tst
}

func (s *ValueContext) Value(i int) IValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ValueContext) ELSE() antlr.TerminalNode {
	return s.GetToken(MMSParserELSE, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MMSParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, MMSParserRULE_value)
	p.SetState(409)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MMSParserIF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(399)
			p.Match(MMSParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(400)
			p.Match(MMSParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(401)
			p.expr(0)
		}
		{
			p.SetState(402)
			p.Match(MMSParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(403)
			p.Value()
		}
		p.SetState(406)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 48, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(404)
				p.Match(MMSParserELSE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(405)
				p.Value()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case MMSParserMISSING, MMSParserRUN, MMSParserTRUE, MMSParserFALSE, MMSParserRESLOC, MMSParserTNAME, MMSParserIDENT, MMSParserNUMBER, MMSParserSTRING, MMSParserELLIPSIS, MMSParserDOTDOT, MMSParserLBRACE, MMSParserLPAREN, MMSParserLBRACK, MMSParserPLUS, MMSParserMINUS, MMSParserBANG, MMSParserHASH, MMSParserAT, MMSParserTILDE, MMSParserCARET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(408)
			p.expr(0)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type OrExprContext struct {
	ExprContext
}

func NewOrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrExprContext {
	var p = new(OrExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *OrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *OrExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *OrExprContext) OROR() antlr.TerminalNode {
	return s.GetToken(MMSParserOROR, 0)
}

type ArrowExprContext struct {
	ExprContext
}

func NewArrowExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrowExprContext {
	var p = new(ArrowExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ArrowExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrowExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ArrowExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ArrowExprContext) ARROW() antlr.TerminalNode {
	return s.GetToken(MMSParserARROW, 0)
}

type CoalesceExprContext struct {
	ExprContext
}

func NewCoalesceExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CoalesceExprContext {
	var p = new(CoalesceExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *CoalesceExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CoalesceExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *CoalesceExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CoalesceExprContext) QQ() antlr.TerminalNode {
	return s.GetToken(MMSParserQQ, 0)
}

type UnaryExprContext struct {
	ExprContext
	op antlr.Token
}

func NewUnaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryExprContext {
	var p = new(UnaryExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *UnaryExprContext) GetOp() antlr.Token { return s.op }

func (s *UnaryExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *UnaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *UnaryExprContext) BANG() antlr.TerminalNode {
	return s.GetToken(MMSParserBANG, 0)
}

func (s *UnaryExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(MMSParserMINUS, 0)
}

func (s *UnaryExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(MMSParserPLUS, 0)
}

type PrimaryExprContext struct {
	ExprContext
}

func NewPrimaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrimaryExprContext {
	var p = new(PrimaryExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *PrimaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExprContext) Primary() IPrimaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

type AddExprContext struct {
	ExprContext
	op antlr.Token
}

func NewAddExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddExprContext {
	var p = new(AddExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AddExprContext) GetOp() antlr.Token { return s.op }

func (s *AddExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *AddExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AddExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(MMSParserPLUS, 0)
}

func (s *AddExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(MMSParserMINUS, 0)
}

type MulExprContext struct {
	ExprContext
	op antlr.Token
}

func NewMulExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulExprContext {
	var p = new(MulExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *MulExprContext) GetOp() antlr.Token { return s.op }

func (s *MulExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *MulExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MulExprContext) STAR() antlr.TerminalNode {
	return s.GetToken(MMSParserSTAR, 0)
}

func (s *MulExprContext) SLASH() antlr.TerminalNode {
	return s.GetToken(MMSParserSLASH, 0)
}

func (s *MulExprContext) PERCENT() antlr.TerminalNode {
	return s.GetToken(MMSParserPERCENT, 0)
}

type SpreadExprContext struct {
	ExprContext
}

func NewSpreadExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SpreadExprContext {
	var p = new(SpreadExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *SpreadExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpreadExprContext) ELLIPSIS() antlr.TerminalNode {
	return s.GetToken(MMSParserELLIPSIS, 0)
}

func (s *SpreadExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

type RangeToExprContext struct {
	ExprContext
}

func NewRangeToExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RangeToExprContext {
	var p = new(RangeToExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *RangeToExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeToExprContext) DOTDOT() antlr.TerminalNode {
	return s.GetToken(MMSParserDOTDOT, 0)
}

func (s *RangeToExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

type PostfixExprContext struct {
	ExprContext
}

func NewPostfixExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PostfixExprContext {
	var p = new(PostfixExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *PostfixExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostfixExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PostfixExprContext) Postfix() IPostfixContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPostfixContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPostfixContext)
}

type RangeExprContext struct {
	ExprContext
}

func NewRangeExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RangeExprContext {
	var p = new(RangeExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *RangeExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *RangeExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RangeExprContext) DOTDOT() antlr.TerminalNode {
	return s.GetToken(MMSParserDOTDOT, 0)
}

type CompareExprContext struct {
	ExprContext
	op antlr.Token
}

func NewCompareExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompareExprContext {
	var p = new(CompareExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *CompareExprContext) GetOp() antlr.Token { return s.op }

func (s *CompareExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *CompareExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *CompareExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CompareExprContext) LT() antlr.TerminalNode {
	return s.GetToken(MMSParserLT, 0)
}

func (s *CompareExprContext) LE() antlr.TerminalNode {
	return s.GetToken(MMSParserLE, 0)
}

func (s *CompareExprContext) GT() antlr.TerminalNode {
	return s.GetToken(MMSParserGT, 0)
}

func (s *CompareExprContext) GE() antlr.TerminalNode {
	return s.GetToken(MMSParserGE, 0)
}

type EqualExprContext struct {
	ExprContext
	op antlr.Token
}

func NewEqualExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualExprContext {
	var p = new(EqualExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *EqualExprContext) GetOp() antlr.Token { return s.op }

func (s *EqualExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *EqualExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *EqualExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EqualExprContext) EQEQ() antlr.TerminalNode {
	return s.GetToken(MMSParserEQEQ, 0)
}

func (s *EqualExprContext) NEQ() antlr.TerminalNode {
	return s.GetToken(MMSParserNEQ, 0)
}

type AndExprContext struct {
	ExprContext
}

func NewAndExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndExprContext {
	var p = new(AndExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *AndExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AndExprContext) ANDAND() antlr.TerminalNode {
	return s.GetToken(MMSParserANDAND, 0)
}

func (p *MMSParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *MMSParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 56
	p.EnterRecursionRule(localctx, 56, MMSParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(419)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MMSParserPLUS, MMSParserMINUS, MMSParserBANG:
		localctx = NewUnaryExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(412)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*UnaryExprContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&35) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*UnaryExprContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(413)
			p.expr(13)
		}

	case MMSParserDOTDOT:
		localctx = NewRangeToExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(414)
			p.Match(MMSParserDOTDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(415)
			p.expr(9)
		}

	case MMSParserELLIPSIS:
		localctx = NewSpreadExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(416)
			p.Match(MMSParserELLIPSIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(417)
			p.expr(2)
		}

	case MMSParserMISSING, MMSParserRUN, MMSParserTRUE, MMSParserFALSE, MMSParserRESLOC, MMSParserTNAME, MMSParserIDENT, MMSParserNUMBER, MMSParserSTRING, MMSParserLBRACE, MMSParserLPAREN, MMSParserLBRACK, MMSParserHASH, MMSParserAT, MMSParserTILDE, MMSParserCARET:
		localctx = NewPrimaryExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(418)
			p.Primary()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(454)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 53, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(452)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 52, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MMSParserRULE_expr)
				p.SetState(421)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(422)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64((_la-67)) & ^0x3f) == 0 && ((int64(1)<<(_la-67))&7) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(423)
					p.expr(13)
				}

			case 2:
				localctx = NewAddExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MMSParserRULE_expr)
				p.SetState(424)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(425)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == MMSParserPLUS || _la == MMSParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(426)
					p.expr(12)
				}

			case 3:
				localctx = NewCompareExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MMSParserRULE_expr)
				p.SetState(427)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(428)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*CompareExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64((_la-47)) & ^0x3f) == 0 && ((int64(1)<<(_la-47))&196611) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*CompareExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(429)
					p.expr(9)
				}

			case 4:
				localctx = NewEqualExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MMSParserRULE_expr)
				p.SetState(430)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(431)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*EqualExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == MMSParserEQEQ || _la == MMSParserNEQ) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*EqualExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(432)
					p.expr(8)
				}

			case 5:
				localctx = NewAndExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MMSParserRULE_expr)
				p.SetState(433)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(434)
					p.Match(MMSParserANDAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(435)
					p.expr(7)
				}

			case 6:
				localctx = NewOrExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MMSParserRULE_expr)
				p.SetState(436)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(437)
					p.Match(MMSParserOROR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(438)
					p.expr(6)
				}

			case 7:
				localctx = NewCoalesceExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MMSParserRULE_expr)
				p.SetState(439)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(440)
					p.Match(MMSParserQQ)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(441)
					p.expr(5)
				}

			case 8:
				localctx = NewArrowExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MMSParserRULE_expr)
				p.SetState(442)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(443)
					p.Match(MMSParserARROW)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(444)
					p.expr(3)
				}

			case 9:
				localctx = NewPostfixExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MMSParserRULE_expr)
				p.SetState(445)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(446)
					p.Postfix()
				}

			case 10:
				localctx = NewRangeExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MMSParserRULE_expr)
				p.SetState(447)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(448)
					p.Match(MMSParserDOTDOT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(450)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 51, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(449)
						p.expr(0)
					}

				} else if p.HasError() { // JIM
					goto errorExit
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(456)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 53, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPostfixContext is an interface to support dynamic dispatch.
type IPostfixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPostfixContext differentiates from other interfaces.
	IsPostfixContext()
}

type PostfixContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPostfixContext() *PostfixContext {
	var p = new(PostfixContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_postfix
	return p
}

func InitEmptyPostfixContext(p *PostfixContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_postfix
}

func (*PostfixContext) IsPostfixContext() {}

func NewPostfixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PostfixContext {
	var p = new(PostfixContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_postfix

	return p
}

func (s *PostfixContext) GetParser() antlr.Parser { return s.parser }

func (s *PostfixContext) CopyAll(ctx *PostfixContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *PostfixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostfixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MemberPostfixContext struct {
	PostfixContext
	name IWordContext
}

func NewMemberPostfixContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MemberPostfixContext {
	var p = new(MemberPostfixContext)

	InitEmptyPostfixContext(&p.PostfixContext)
	p.parser = parser
	p.CopyAll(ctx.(*PostfixContext))

	return p
}

func (s *MemberPostfixContext) GetName() IWordContext { return s.name }

func (s *MemberPostfixContext) SetName(v IWordContext) { s.name = v }

func (s *MemberPostfixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberPostfixContext) DOT() antlr.TerminalNode {
	return s.GetToken(MMSParserDOT, 0)
}

func (s *MemberPostfixContext) Word() IWordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWordContext)
}

type IndexPostfixContext struct {
	PostfixContext
}

func NewIndexPostfixContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexPostfixContext {
	var p = new(IndexPostfixContext)

	InitEmptyPostfixContext(&p.PostfixContext)
	p.parser = parser
	p.CopyAll(ctx.(*PostfixContext))

	return p
}

func (s *IndexPostfixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexPostfixContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(MMSParserLBRACK, 0)
}

func (s *IndexPostfixContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(MMSParserRBRACK, 0)
}

func (s *IndexPostfixContext) AllArg() []IArgContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArgContext); ok {
			len++
		}
	}

	tst := make([]IArgContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArgContext); ok {
			tst[i] = t.(IArgContext)
			i++
		}
	}

	return tst
}

func (s *IndexPostfixContext) Arg(i int) IArgContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgContext)
}

func (s *IndexPostfixContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MMSParserCOMMA)
}

func (s *IndexPostfixContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MMSParserCOMMA, i)
}

type BlockPostfixContext struct {
	PostfixContext
}

func NewBlockPostfixContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BlockPostfixContext {
	var p = new(BlockPostfixContext)

	InitEmptyPostfixContext(&p.PostfixContext)
	p.parser = parser
	p.CopyAll(ctx.(*PostfixContext))

	return p
}

func (s *BlockPostfixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockPostfixContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

type CallPostfixContext struct {
	PostfixContext
}

func NewCallPostfixContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallPostfixContext {
	var p = new(CallPostfixContext)

	InitEmptyPostfixContext(&p.PostfixContext)
	p.parser = parser
	p.CopyAll(ctx.(*PostfixContext))

	return p
}

func (s *CallPostfixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallPostfixContext) Args() IArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
}

func (p *MMSParser) Postfix() (localctx IPostfixContext) {
	localctx = NewPostfixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, MMSParserRULE_postfix)
	var _la int

	var _alt int

	p.SetState(478)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MMSParserDOT:
		localctx = NewMemberPostfixContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(457)
			p.Match(MMSParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(459)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 54, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(458)

				var _x = p.Word()

				localctx.(*MemberPostfixContext).name = _x
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case MMSParserLPAREN:
		localctx = NewCallPostfixContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(461)
			p.Args()
		}

	case MMSParserLBRACK:
		localctx = NewIndexPostfixContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(462)
			p.Match(MMSParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(474)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&189151663255158930) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&995) != 0) {
			{
				p.SetState(463)
				p.Arg()
			}
			p.SetState(468)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 55, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
			for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				if _alt == 1 {
					{
						p.SetState(464)
						p.Match(MMSParserCOMMA)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}
					{
						p.SetState(465)
						p.Arg()
					}

				}
				p.SetState(470)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 55, p.GetParserRuleContext())
				if p.HasError() {
					goto errorExit
				}
			}
			p.SetState(472)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == MMSParserCOMMA {
				{
					p.SetState(471)
					p.Match(MMSParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}

		}
		{
			p.SetState(476)
			p.Match(MMSParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MMSParserLBRACE:
		localctx = NewBlockPostfixContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(477)
			p.Block()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgsContext is an interface to support dynamic dispatch.
type IArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllArg() []IArgContext
	Arg(i int) IArgContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArgsContext differentiates from other interfaces.
	IsArgsContext()
}

type ArgsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgsContext() *ArgsContext {
	var p = new(ArgsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_args
	return p
}

func InitEmptyArgsContext(p *ArgsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_args
}

func (*ArgsContext) IsArgsContext() {}

func NewArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgsContext {
	var p = new(ArgsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_args

	return p
}

func (s *ArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgsContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MMSParserLPAREN, 0)
}

func (s *ArgsContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MMSParserRPAREN, 0)
}

func (s *ArgsContext) AllArg() []IArgContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArgContext); ok {
			len++
		}
	}

	tst := make([]IArgContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArgContext); ok {
			tst[i] = t.(IArgContext)
			i++
		}
	}

	return tst
}

func (s *ArgsContext) Arg(i int) IArgContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgContext)
}

func (s *ArgsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MMSParserCOMMA)
}

func (s *ArgsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MMSParserCOMMA, i)
}

func (s *ArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MMSParser) Args() (localctx IArgsContext) {
	localctx = NewArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, MMSParserRULE_args)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(480)
		p.Match(MMSParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(492)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&189151663255158930) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&995) != 0) {
		{
			p.SetState(481)
			p.Arg()
		}
		p.SetState(486)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 59, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(482)
					p.Match(MMSParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(483)
					p.Arg()
				}

			}
			p.SetState(488)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 59, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(490)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MMSParserCOMMA {
			{
				p.SetState(489)
				p.Match(MMSParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(494)
		p.Match(MMSParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgContext is an interface to support dynamic dispatch.
type IArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Value() IValueContext
	Word() IWordContext
	ASSIGN() antlr.TerminalNode

	// IsArgContext differentiates from other interfaces.
	IsArgContext()
}

type ArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgContext() *ArgContext {
	var p = new(ArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_arg
	return p
}

func InitEmptyArgContext(p *ArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_arg
}

func (*ArgContext) IsArgContext() {}

func NewArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgContext {
	var p = new(ArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_arg

	return p
}

func (s *ArgContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ArgContext) Word() IWordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWordContext)
}

func (s *ArgContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(MMSParserASSIGN, 0)
}

func (s *ArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MMSParser) Arg() (localctx IArgContext) {
	localctx = NewArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, MMSParserRULE_arg)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(499)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 62, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(496)
			p.Word()
		}
		{
			p.SetState(497)
			p.Match(MMSParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(501)
		p.Value()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_primary
	return p
}

func InitEmptyPrimaryContext(p *PrimaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_primary
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_primary

	return p
}

func (s *PrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryContext) CopyAll(ctx *PrimaryContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TagRefContext struct {
	PrimaryContext
}

func NewTagRefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TagRefContext {
	var p = new(TagRefContext)

	InitEmptyPrimaryContext(&p.PrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryContext))

	return p
}

func (s *TagRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagRefContext) HASH() antlr.TerminalNode {
	return s.GetToken(MMSParserHASH, 0)
}

func (s *TagRefContext) IDENT() antlr.TerminalNode {
	return s.GetToken(MMSParserIDENT, 0)
}

func (s *TagRefContext) RESLOC() antlr.TerminalNode {
	return s.GetToken(MMSParserRESLOC, 0)
}

func (s *TagRefContext) TNAME() antlr.TerminalNode {
	return s.GetToken(MMSParserTNAME, 0)
}

type ReslocRefContext struct {
	PrimaryContext
}

func NewReslocRefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReslocRefContext {
	var p = new(ReslocRefContext)

	InitEmptyPrimaryContext(&p.PrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryContext))

	return p
}

func (s *ReslocRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReslocRefContext) RESLOC() antlr.TerminalNode {
	return s.GetToken(MMSParserRESLOC, 0)
}

type ListLitContext struct {
	PrimaryContext
}

func NewListLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListLitContext {
	var p = new(ListLitContext)

	InitEmptyPrimaryContext(&p.PrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryContext))

	return p
}

func (s *ListLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListLitContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(MMSParserLBRACK, 0)
}

func (s *ListLitContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(MMSParserRBRACK, 0)
}

func (s *ListLitContext) AllValue() []IValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueContext); ok {
			len++
		}
	}

	tst := make([]IValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueContext); ok {
			tst[i] = t.(IValueContext)
			i++
		}
	}

	return tst
}

func (s *ListLitContext) Value(i int) IValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ListLitContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MMSParserCOMMA)
}

func (s *ListLitContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MMSParserCOMMA, i)
}

type CoordinateContext struct {
	PrimaryContext
}

func NewCoordinateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CoordinateContext {
	var p = new(CoordinateContext)

	InitEmptyPrimaryContext(&p.PrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryContext))

	return p
}

func (s *CoordinateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CoordinateContext) AllTILDE() []antlr.TerminalNode {
	return s.GetTokens(MMSParserTILDE)
}

func (s *CoordinateContext) TILDE(i int) antlr.TerminalNode {
	return s.GetToken(MMSParserTILDE, i)
}

func (s *CoordinateContext) AllCARET() []antlr.TerminalNode {
	return s.GetTokens(MMSParserCARET)
}

func (s *CoordinateContext) CARET(i int) antlr.TerminalNode {
	return s.GetToken(MMSParserCARET, i)
}

func (s *CoordinateContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(MMSParserNUMBER)
}

func (s *CoordinateContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(MMSParserNUMBER, i)
}

func (s *CoordinateContext) AllMINUS() []antlr.TerminalNode {
	return s.GetTokens(MMSParserMINUS)
}

func (s *CoordinateContext) MINUS(i int) antlr.TerminalNode {
	return s.GetToken(MMSParserMINUS, i)
}

type MissingLitContext struct {
	PrimaryContext
}

func NewMissingLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MissingLitContext {
	var p = new(MissingLitContext)

	InitEmptyPrimaryContext(&p.PrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryContext))

	return p
}

func (s *MissingLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MissingLitContext) MISSING() antlr.TerminalNode {
	return s.GetToken(MMSParserMISSING, 0)
}

type IdentRefContext struct {
	PrimaryContext
}

func NewIdentRefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentRefContext {
	var p = new(IdentRefContext)

	InitEmptyPrimaryContext(&p.PrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryContext))

	return p
}

func (s *IdentRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentRefContext) IDENT() antlr.TerminalNode {
	return s.GetToken(MMSParserIDENT, 0)
}

type ComprehensionContext struct {
	PrimaryContext
}

func NewComprehensionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComprehensionContext {
	var p = new(ComprehensionContext)

	InitEmptyPrimaryContext(&p.PrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryContext))

	return p
}

func (s *ComprehensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComprehensionContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(MMSParserLBRACK, 0)
}

func (s *ComprehensionContext) FOR() antlr.TerminalNode {
	return s.GetToken(MMSParserFOR, 0)
}

func (s *ComprehensionContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(MMSParserIDENT)
}

func (s *ComprehensionContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(MMSParserIDENT, i)
}

func (s *ComprehensionContext) IN() antlr.TerminalNode {
	return s.GetToken(MMSParserIN, 0)
}

func (s *ComprehensionContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ComprehensionContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ComprehensionContext) COLON() antlr.TerminalNode {
	return s.GetToken(MMSParserCOLON, 0)
}

func (s *ComprehensionContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ComprehensionContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(MMSParserRBRACK, 0)
}

func (s *ComprehensionContext) COMMA() antlr.TerminalNode {
	return s.GetToken(MMSParserCOMMA, 0)
}

func (s *ComprehensionContext) IF() antlr.TerminalNode {
	return s.GetToken(MMSParserIF, 0)
}

type ParenExprContext struct {
	PrimaryContext
}

func NewParenExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExprContext {
	var p = new(ParenExprContext)

	InitEmptyPrimaryContext(&p.PrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryContext))

	return p
}

func (s *ParenExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MMSParserLPAREN, 0)
}

func (s *ParenExprContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ParenExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MMSParserRPAREN, 0)
}

type ObjectLitContext struct {
	PrimaryContext
}

func NewObjectLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ObjectLitContext {
	var p = new(ObjectLitContext)

	InitEmptyPrimaryContext(&p.PrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryContext))

	return p
}

func (s *ObjectLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectLitContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

type StringLitContext struct {
	PrimaryContext
}

func NewStringLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringLitContext {
	var p = new(StringLitContext)

	InitEmptyPrimaryContext(&p.PrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryContext))

	return p
}

func (s *StringLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLitContext) STRING() antlr.TerminalNode {
	return s.GetToken(MMSParserSTRING, 0)
}

type TemplatedRefContext struct {
	PrimaryContext
}

func NewTemplatedRefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TemplatedRefContext {
	var p = new(TemplatedRefContext)

	InitEmptyPrimaryContext(&p.PrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryContext))

	return p
}

func (s *TemplatedRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TemplatedRefContext) TNAME() antlr.TerminalNode {
	return s.GetToken(MMSParserTNAME, 0)
}

type BoolLitContext struct {
	PrimaryContext
}

func NewBoolLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolLitContext {
	var p = new(BoolLitContext)

	InitEmptyPrimaryContext(&p.PrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryContext))

	return p
}

func (s *BoolLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolLitContext) TRUE() antlr.TerminalNode {
	return s.GetToken(MMSParserTRUE, 0)
}

func (s *BoolLitContext) FALSE() antlr.TerminalNode {
	return s.GetToken(MMSParserFALSE, 0)
}

type NumberLitContext struct {
	PrimaryContext
}

func NewNumberLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberLitContext {
	var p = new(NumberLitContext)

	InitEmptyPrimaryContext(&p.PrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryContext))

	return p
}

func (s *NumberLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberLitContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(MMSParserNUMBER, 0)
}

type SelectorContext struct {
	PrimaryContext
}

func NewSelectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelectorContext {
	var p = new(SelectorContext)

	InitEmptyPrimaryContext(&p.PrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryContext))

	return p
}

func (s *SelectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorContext) AT() antlr.TerminalNode {
	return s.GetToken(MMSParserAT, 0)
}

func (s *SelectorContext) IDENT() antlr.TerminalNode {
	return s.GetToken(MMSParserIDENT, 0)
}

type RunExprContext struct {
	PrimaryContext
}

func NewRunExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RunExprContext {
	var p = new(RunExprContext)

	InitEmptyPrimaryContext(&p.PrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryContext))

	return p
}

func (s *RunExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RunExprContext) RUN() antlr.TerminalNode {
	return s.GetToken(MMSParserRUN, 0)
}

func (s *RunExprContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (p *MMSParser) Primary() (localctx IPrimaryContext) {
	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, MMSParserRULE_primary)
	var _la int

	var _alt int

	p.SetState(566)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 72, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNumberLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(503)
			p.Match(MMSParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewStringLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(504)
			p.Match(MMSParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewBoolLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(505)
			_la = p.GetTokenStream().LA(1)

			if !(_la == MMSParserTRUE || _la == MMSParserFALSE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 4:
		localctx = NewRunExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(506)
			p.Match(MMSParserRUN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(507)
			p.Block()
		}

	case 5:
		localctx = NewIdentRefContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(508)
			p.Match(MMSParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewReslocRefContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(509)
			p.Match(MMSParserRESLOC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewTemplatedRefContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(510)
			p.Match(MMSParserTNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewTagRefContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(511)
			p.Match(MMSParserHASH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(512)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&15032385536) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 9:
		localctx = NewSelectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(513)
			p.Match(MMSParserAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(514)
			p.Match(MMSParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		localctx = NewCoordinateContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		p.SetState(522)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(515)
					_la = p.GetTokenStream().LA(1)

					if !(_la == MMSParserTILDE || _la == MMSParserCARET) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				p.SetState(520)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 64, p.GetParserRuleContext()) == 1 {
					p.SetState(517)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_la = p.GetTokenStream().LA(1)

					if _la == MMSParserMINUS {
						{
							p.SetState(516)
							p.Match(MMSParserMINUS)
							if p.HasError() {
								// Recognition error - abort rule
								goto errorExit
							}
						}

					}
					{
						p.SetState(519)
						p.Match(MMSParserNUMBER)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				} else if p.HasError() { // JIM
					goto errorExit
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(524)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 65, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case 11:
		localctx = NewParenExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(526)
			p.Match(MMSParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(527)
			p.Value()
		}
		{
			p.SetState(528)
			p.Match(MMSParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 12:
		localctx = NewComprehensionContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(530)
			p.Match(MMSParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(531)
			p.Match(MMSParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(532)
			p.Match(MMSParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(535)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MMSParserCOMMA {
			{
				p.SetState(533)
				p.Match(MMSParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(534)
				p.Match(MMSParserIDENT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(537)
			p.Match(MMSParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(538)
			p.expr(0)
		}
		p.SetState(541)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MMSParserIF {
			{
				p.SetState(539)
				p.Match(MMSParserIF)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(540)
				p.expr(0)
			}

		}
		{
			p.SetState(543)
			p.Match(MMSParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(544)
			p.Value()
		}
		{
			p.SetState(545)
			p.Match(MMSParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 13:
		localctx = NewListLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(547)
			p.Match(MMSParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(561)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&189151663253094530) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&995) != 0) {
			{
				p.SetState(548)
				p.Value()
			}
			p.SetState(555)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 69, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
			for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				if _alt == 1 {
					p.SetState(550)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_la = p.GetTokenStream().LA(1)

					if _la == MMSParserCOMMA {
						{
							p.SetState(549)
							p.Match(MMSParserCOMMA)
							if p.HasError() {
								// Recognition error - abort rule
								goto errorExit
							}
						}

					}
					{
						p.SetState(552)
						p.Value()
					}

				}
				p.SetState(557)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 69, p.GetParserRuleContext())
				if p.HasError() {
					goto errorExit
				}
			}
			p.SetState(559)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == MMSParserCOMMA {
				{
					p.SetState(558)
					p.Match(MMSParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}

		}
		{
			p.SetState(563)
			p.Match(MMSParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 14:
		localctx = NewObjectLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(564)
			p.Block()
		}

	case 15:
		localctx = NewMissingLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(565)
			p.Match(MMSParserMISSING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *MMSParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 24:
		var t *HintContext = nil
		if localctx != nil {
			t = localctx.(*HintContext)
		}
		return p.Hint_Sempred(t, predIndex)

	case 28:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *MMSParser) Hint_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.word("hint")

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *MMSParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 10)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
