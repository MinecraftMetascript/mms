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
		"'false'", "", "", "", "", "", "", "", "", "", "", "", "", "'...'",
		"'..'", "'??'", "'+='", "'-='", "'*='", "'/='", "'%='", "'=='", "'!='",
		"'<='", "'>='", "'&&'", "'||'", "'|'", "'?'", "'{'", "'}'", "'('", "')'",
		"'['", "']'", "','", "':'", "'.'", "'='", "'<'", "'>'", "'+'", "'-'",
		"'*'", "'/'", "'%'", "'!'", "'#'", "'@'", "", "'~'", "'^'",
	}
	staticData.SymbolicNames = []string{
		"", "MISSING", "SEP", "CONST", "TYPE", "EXTERN", "LET", "IF", "ELSE",
		"WHILE", "FOR", "IN", "RETURN", "BREAK", "CONTINUE", "PRIMARY", "AS",
		"CHECK", "EMIT", "SETUP", "EXPECT", "RUN", "TRUE", "FALSE", "COMMAND",
		"DOC_COMMENT", "BLOCK_COMMENT", "UNTERMINATED_COMMENT", "LINE_COMMENT",
		"NEWLINE", "WS", "RESLOC", "TNAME", "IDENT", "NUMBER", "STRING", "ELLIPSIS",
		"DOTDOT", "QQ", "PLUS_EQ", "MINUS_EQ", "STAR_EQ", "SLASH_EQ", "PCT_EQ",
		"EQEQ", "NEQ", "LE", "GE", "ANDAND", "OROR", "PIPE", "QUESTION", "LBRACE",
		"RBRACE", "LPAREN", "RPAREN", "LBRACK", "RBRACK", "COMMA", "COLON",
		"DOT", "ASSIGN", "LT", "GT", "PLUS", "MINUS", "STAR", "SLASH", "PERCENT",
		"BANG", "HASH", "AT", "REL_NAME", "TILDE", "CARET", "ERROR_CHAR",
	}
	staticData.RuleNames = []string{
		"file", "exprOnly", "topItem", "declaration", "innerDecl", "innerBody",
		"declName", "word", "declBody", "constDecl", "letDecl", "typeAnnotation",
		"typeExpr", "typeArm", "typeAtom", "typeField", "mapKey", "typeName",
		"params", "param", "block", "blockItem", "statement", "thenPart", "elsePart",
		"bareCall", "hint", "ifTail", "fieldKey", "value", "expr", "postfix",
		"args", "arg", "primary",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 75, 619, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 1, 0, 5, 0, 72, 8, 0, 10,
		0, 12, 0, 75, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 3, 2, 83, 8, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 92, 8, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 3, 2, 98, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 4, 3, 113, 8, 3, 11, 3, 12, 3, 114, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 122, 8, 3, 3, 3, 124, 8, 3, 1, 3, 3, 3, 127,
		8, 3, 1, 3, 1, 3, 3, 3, 131, 8, 3, 1, 3, 3, 3, 134, 8, 3, 1, 4, 1, 4, 1,
		4, 1, 4, 4, 4, 140, 8, 4, 11, 4, 12, 4, 141, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 3, 4, 149, 8, 4, 3, 4, 151, 8, 4, 1, 4, 3, 4, 154, 8, 4, 3, 4, 156,
		8, 4, 1, 5, 1, 5, 1, 5, 3, 5, 161, 8, 5, 1, 6, 1, 6, 1, 6, 3, 6, 166, 8,
		6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 3, 8, 173, 8, 8, 1, 9, 1, 9, 1, 9, 3,
		9, 178, 8, 9, 1, 9, 1, 9, 3, 9, 182, 8, 9, 1, 10, 1, 10, 1, 10, 3, 10,
		187, 8, 10, 1, 10, 1, 10, 3, 10, 191, 8, 10, 1, 11, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 12, 5, 12, 199, 8, 12, 10, 12, 12, 12, 202, 9, 12, 1, 13, 1,
		13, 3, 13, 206, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 213, 8,
		14, 10, 14, 12, 14, 216, 9, 14, 1, 14, 1, 14, 3, 14, 220, 8, 14, 1, 14,
		1, 14, 3, 14, 224, 8, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 3, 14, 234, 8, 14, 1, 14, 5, 14, 237, 8, 14, 10, 14, 12, 14,
		240, 9, 14, 1, 14, 3, 14, 243, 8, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 257, 8, 14, 1,
		15, 1, 15, 3, 15, 261, 8, 15, 1, 15, 1, 15, 3, 15, 265, 8, 15, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 5, 16, 272, 8, 16, 10, 16, 12, 16, 275, 9, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 284, 8, 16, 1,
		17, 1, 17, 1, 17, 5, 17, 289, 8, 17, 10, 17, 12, 17, 292, 9, 17, 1, 17,
		3, 17, 295, 8, 17, 1, 18, 1, 18, 1, 18, 1, 18, 5, 18, 301, 8, 18, 10, 18,
		12, 18, 304, 9, 18, 1, 18, 3, 18, 307, 8, 18, 3, 18, 309, 8, 18, 1, 18,
		1, 18, 1, 19, 1, 19, 3, 19, 315, 8, 19, 1, 19, 1, 19, 3, 19, 319, 8, 19,
		1, 20, 1, 20, 5, 20, 323, 8, 20, 10, 20, 12, 20, 326, 9, 20, 1, 20, 1,
		20, 1, 21, 1, 21, 3, 21, 332, 8, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 343, 8, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 355, 8, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 363, 8, 22, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 376,
		8, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 3, 22, 388, 8, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		4, 22, 397, 8, 22, 11, 22, 12, 22, 398, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22,
		415, 8, 22, 1, 23, 1, 23, 3, 23, 419, 8, 23, 1, 24, 1, 24, 1, 24, 3, 24,
		424, 8, 24, 1, 25, 1, 25, 5, 25, 428, 8, 25, 10, 25, 12, 25, 431, 9, 25,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		27, 3, 27, 444, 8, 27, 1, 28, 1, 28, 1, 28, 3, 28, 449, 8, 28, 1, 29, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 3, 29, 458, 8, 29, 1, 29, 1, 29,
		3, 29, 462, 8, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1,
		30, 3, 30, 472, 8, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1,
		30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 3, 30, 500,
		8, 30, 5, 30, 502, 8, 30, 10, 30, 12, 30, 505, 9, 30, 1, 31, 1, 31, 3,
		31, 509, 8, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 5, 31, 516, 8, 31, 10,
		31, 12, 31, 519, 9, 31, 1, 31, 3, 31, 522, 8, 31, 3, 31, 524, 8, 31, 1,
		31, 1, 31, 3, 31, 528, 8, 31, 1, 32, 1, 32, 1, 32, 1, 32, 5, 32, 534, 8,
		32, 10, 32, 12, 32, 537, 9, 32, 1, 32, 3, 32, 540, 8, 32, 3, 32, 542, 8,
		32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 3, 33, 549, 8, 33, 1, 33, 1, 33,
		1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 34, 3, 34, 567, 8, 34, 1, 34, 3, 34, 570, 8, 34, 1,
		34, 4, 34, 573, 8, 34, 11, 34, 12, 34, 574, 1, 34, 1, 34, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 3, 34, 586, 8, 34, 1, 34, 1, 34, 1,
		34, 1, 34, 3, 34, 592, 8, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34,
		1, 34, 3, 34, 601, 8, 34, 1, 34, 5, 34, 604, 8, 34, 10, 34, 12, 34, 607,
		9, 34, 1, 34, 3, 34, 610, 8, 34, 3, 34, 612, 8, 34, 1, 34, 1, 34, 1, 34,
		3, 34, 617, 8, 34, 1, 34, 0, 1, 60, 35, 0, 2, 4, 6, 8, 10, 12, 14, 16,
		18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52,
		54, 56, 58, 60, 62, 64, 66, 68, 0, 13, 3, 0, 4, 4, 15, 21, 33, 33, 2, 0,
		2, 2, 58, 58, 1, 0, 17, 19, 1, 0, 32, 33, 1, 0, 39, 43, 2, 0, 64, 65, 69,
		69, 1, 0, 66, 68, 1, 0, 64, 65, 2, 0, 46, 47, 62, 63, 1, 0, 44, 45, 1,
		0, 22, 23, 1, 0, 31, 33, 1, 0, 73, 74, 725, 0, 73, 1, 0, 0, 0, 2, 78, 1,
		0, 0, 0, 4, 97, 1, 0, 0, 0, 6, 133, 1, 0, 0, 0, 8, 155, 1, 0, 0, 0, 10,
		160, 1, 0, 0, 0, 12, 165, 1, 0, 0, 0, 14, 167, 1, 0, 0, 0, 16, 172, 1,
		0, 0, 0, 18, 174, 1, 0, 0, 0, 20, 183, 1, 0, 0, 0, 22, 192, 1, 0, 0, 0,
		24, 195, 1, 0, 0, 0, 26, 203, 1, 0, 0, 0, 28, 256, 1, 0, 0, 0, 30, 258,
		1, 0, 0, 0, 32, 283, 1, 0, 0, 0, 34, 294, 1, 0, 0, 0, 36, 296, 1, 0, 0,
		0, 38, 312, 1, 0, 0, 0, 40, 320, 1, 0, 0, 0, 42, 329, 1, 0, 0, 0, 44, 414,
		1, 0, 0, 0, 46, 418, 1, 0, 0, 0, 48, 423, 1, 0, 0, 0, 50, 425, 1, 0, 0,
		0, 52, 432, 1, 0, 0, 0, 54, 436, 1, 0, 0, 0, 56, 448, 1, 0, 0, 0, 58, 461,
		1, 0, 0, 0, 60, 471, 1, 0, 0, 0, 62, 527, 1, 0, 0, 0, 64, 529, 1, 0, 0,
		0, 66, 548, 1, 0, 0, 0, 68, 616, 1, 0, 0, 0, 70, 72, 3, 4, 2, 0, 71, 70,
		1, 0, 0, 0, 72, 75, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0,
		74, 76, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 76, 77, 5, 0, 0, 1, 77, 1, 1, 0,
		0, 0, 78, 79, 3, 60, 30, 0, 79, 80, 5, 0, 0, 1, 80, 3, 1, 0, 0, 0, 81,
		83, 5, 5, 0, 0, 82, 81, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 84, 1, 0, 0,
		0, 84, 98, 3, 6, 3, 0, 85, 98, 3, 18, 9, 0, 86, 98, 3, 20, 10, 0, 87, 88,
		5, 10, 0, 0, 88, 91, 5, 33, 0, 0, 89, 90, 5, 58, 0, 0, 90, 92, 5, 33, 0,
		0, 91, 89, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 94,
		5, 11, 0, 0, 94, 95, 3, 60, 30, 0, 95, 96, 3, 40, 20, 0, 96, 98, 1, 0,
		0, 0, 97, 82, 1, 0, 0, 0, 97, 85, 1, 0, 0, 0, 97, 86, 1, 0, 0, 0, 97, 87,
		1, 0, 0, 0, 98, 5, 1, 0, 0, 0, 99, 100, 5, 4, 0, 0, 100, 101, 3, 12, 6,
		0, 101, 102, 5, 61, 0, 0, 102, 103, 3, 24, 12, 0, 103, 134, 1, 0, 0, 0,
		104, 105, 5, 4, 0, 0, 105, 106, 3, 12, 6, 0, 106, 107, 3, 40, 20, 0, 107,
		134, 1, 0, 0, 0, 108, 109, 5, 33, 0, 0, 109, 112, 3, 12, 6, 0, 110, 111,
		5, 58, 0, 0, 111, 113, 3, 12, 6, 0, 112, 110, 1, 0, 0, 0, 113, 114, 1,
		0, 0, 0, 114, 112, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 134, 1, 0, 0,
		0, 116, 117, 5, 33, 0, 0, 117, 123, 3, 12, 6, 0, 118, 124, 3, 22, 11, 0,
		119, 121, 3, 36, 18, 0, 120, 122, 3, 22, 11, 0, 121, 120, 1, 0, 0, 0, 121,
		122, 1, 0, 0, 0, 122, 124, 1, 0, 0, 0, 123, 118, 1, 0, 0, 0, 123, 119,
		1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 126, 1, 0, 0, 0, 125, 127, 3, 16,
		8, 0, 126, 125, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 134, 1, 0, 0, 0,
		128, 130, 5, 33, 0, 0, 129, 131, 3, 36, 18, 0, 130, 129, 1, 0, 0, 0, 130,
		131, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 134, 3, 40, 20, 0, 133, 99,
		1, 0, 0, 0, 133, 104, 1, 0, 0, 0, 133, 108, 1, 0, 0, 0, 133, 116, 1, 0,
		0, 0, 133, 128, 1, 0, 0, 0, 134, 7, 1, 0, 0, 0, 135, 136, 5, 33, 0, 0,
		136, 139, 3, 12, 6, 0, 137, 138, 5, 58, 0, 0, 138, 140, 3, 12, 6, 0, 139,
		137, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 141, 142,
		1, 0, 0, 0, 142, 156, 1, 0, 0, 0, 143, 144, 5, 33, 0, 0, 144, 150, 3, 12,
		6, 0, 145, 151, 3, 22, 11, 0, 146, 148, 3, 36, 18, 0, 147, 149, 3, 22,
		11, 0, 148, 147, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 151, 1, 0, 0, 0,
		150, 145, 1, 0, 0, 0, 150, 146, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151,
		153, 1, 0, 0, 0, 152, 154, 3, 10, 5, 0, 153, 152, 1, 0, 0, 0, 153, 154,
		1, 0, 0, 0, 154, 156, 1, 0, 0, 0, 155, 135, 1, 0, 0, 0, 155, 143, 1, 0,
		0, 0, 156, 9, 1, 0, 0, 0, 157, 158, 5, 61, 0, 0, 158, 161, 3, 58, 29, 0,
		159, 161, 3, 40, 20, 0, 160, 157, 1, 0, 0, 0, 160, 159, 1, 0, 0, 0, 161,
		11, 1, 0, 0, 0, 162, 166, 3, 14, 7, 0, 163, 166, 5, 31, 0, 0, 164, 166,
		5, 32, 0, 0, 165, 162, 1, 0, 0, 0, 165, 163, 1, 0, 0, 0, 165, 164, 1, 0,
		0, 0, 166, 13, 1, 0, 0, 0, 167, 168, 7, 0, 0, 0, 168, 15, 1, 0, 0, 0, 169,
		170, 5, 61, 0, 0, 170, 173, 3, 58, 29, 0, 171, 173, 3, 40, 20, 0, 172,
		169, 1, 0, 0, 0, 172, 171, 1, 0, 0, 0, 173, 17, 1, 0, 0, 0, 174, 175, 5,
		3, 0, 0, 175, 177, 5, 33, 0, 0, 176, 178, 3, 22, 11, 0, 177, 176, 1, 0,
		0, 0, 177, 178, 1, 0, 0, 0, 178, 181, 1, 0, 0, 0, 179, 180, 5, 61, 0, 0,
		180, 182, 3, 58, 29, 0, 181, 179, 1, 0, 0, 0, 181, 182, 1, 0, 0, 0, 182,
		19, 1, 0, 0, 0, 183, 184, 5, 6, 0, 0, 184, 186, 5, 33, 0, 0, 185, 187,
		3, 22, 11, 0, 186, 185, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 190, 1,
		0, 0, 0, 188, 189, 5, 61, 0, 0, 189, 191, 3, 58, 29, 0, 190, 188, 1, 0,
		0, 0, 190, 191, 1, 0, 0, 0, 191, 21, 1, 0, 0, 0, 192, 193, 5, 59, 0, 0,
		193, 194, 3, 24, 12, 0, 194, 23, 1, 0, 0, 0, 195, 200, 3, 26, 13, 0, 196,
		197, 5, 50, 0, 0, 197, 199, 3, 26, 13, 0, 198, 196, 1, 0, 0, 0, 199, 202,
		1, 0, 0, 0, 200, 198, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201, 25, 1, 0,
		0, 0, 202, 200, 1, 0, 0, 0, 203, 205, 3, 28, 14, 0, 204, 206, 5, 51, 0,
		0, 205, 204, 1, 0, 0, 0, 205, 206, 1, 0, 0, 0, 206, 27, 1, 0, 0, 0, 207,
		219, 3, 34, 17, 0, 208, 209, 5, 62, 0, 0, 209, 214, 3, 24, 12, 0, 210,
		211, 5, 58, 0, 0, 211, 213, 3, 24, 12, 0, 212, 210, 1, 0, 0, 0, 213, 216,
		1, 0, 0, 0, 214, 212, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 217, 1, 0,
		0, 0, 216, 214, 1, 0, 0, 0, 217, 218, 5, 63, 0, 0, 218, 220, 1, 0, 0, 0,
		219, 208, 1, 0, 0, 0, 219, 220, 1, 0, 0, 0, 220, 257, 1, 0, 0, 0, 221,
		257, 5, 35, 0, 0, 222, 224, 5, 65, 0, 0, 223, 222, 1, 0, 0, 0, 223, 224,
		1, 0, 0, 0, 224, 225, 1, 0, 0, 0, 225, 257, 5, 34, 0, 0, 226, 227, 5, 56,
		0, 0, 227, 228, 3, 24, 12, 0, 228, 229, 5, 57, 0, 0, 229, 257, 1, 0, 0,
		0, 230, 231, 5, 52, 0, 0, 231, 238, 3, 30, 15, 0, 232, 234, 7, 1, 0, 0,
		233, 232, 1, 0, 0, 0, 233, 234, 1, 0, 0, 0, 234, 235, 1, 0, 0, 0, 235,
		237, 3, 30, 15, 0, 236, 233, 1, 0, 0, 0, 237, 240, 1, 0, 0, 0, 238, 236,
		1, 0, 0, 0, 238, 239, 1, 0, 0, 0, 239, 242, 1, 0, 0, 0, 240, 238, 1, 0,
		0, 0, 241, 243, 7, 1, 0, 0, 242, 241, 1, 0, 0, 0, 242, 243, 1, 0, 0, 0,
		243, 244, 1, 0, 0, 0, 244, 245, 5, 53, 0, 0, 245, 257, 1, 0, 0, 0, 246,
		247, 5, 52, 0, 0, 247, 248, 3, 32, 16, 0, 248, 249, 5, 59, 0, 0, 249, 250,
		3, 24, 12, 0, 250, 251, 5, 53, 0, 0, 251, 257, 1, 0, 0, 0, 252, 253, 5,
		54, 0, 0, 253, 254, 3, 24, 12, 0, 254, 255, 5, 55, 0, 0, 255, 257, 1, 0,
		0, 0, 256, 207, 1, 0, 0, 0, 256, 221, 1, 0, 0, 0, 256, 223, 1, 0, 0, 0,
		256, 226, 1, 0, 0, 0, 256, 230, 1, 0, 0, 0, 256, 246, 1, 0, 0, 0, 256,
		252, 1, 0, 0, 0, 257, 29, 1, 0, 0, 0, 258, 260, 5, 33, 0, 0, 259, 261,
		3, 22, 11, 0, 260, 259, 1, 0, 0, 0, 260, 261, 1, 0, 0, 0, 261, 264, 1,
		0, 0, 0, 262, 263, 5, 61, 0, 0, 263, 265, 3, 58, 29, 0, 264, 262, 1, 0,
		0, 0, 264, 265, 1, 0, 0, 0, 265, 31, 1, 0, 0, 0, 266, 267, 3, 34, 17, 0,
		267, 268, 5, 62, 0, 0, 268, 273, 3, 24, 12, 0, 269, 270, 5, 58, 0, 0, 270,
		272, 3, 24, 12, 0, 271, 269, 1, 0, 0, 0, 272, 275, 1, 0, 0, 0, 273, 271,
		1, 0, 0, 0, 273, 274, 1, 0, 0, 0, 274, 276, 1, 0, 0, 0, 275, 273, 1, 0,
		0, 0, 276, 277, 5, 63, 0, 0, 277, 284, 1, 0, 0, 0, 278, 284, 5, 31, 0,
		0, 279, 280, 5, 54, 0, 0, 280, 281, 3, 24, 12, 0, 281, 282, 5, 55, 0, 0,
		282, 284, 1, 0, 0, 0, 283, 266, 1, 0, 0, 0, 283, 278, 1, 0, 0, 0, 283,
		279, 1, 0, 0, 0, 284, 33, 1, 0, 0, 0, 285, 290, 5, 33, 0, 0, 286, 287,
		5, 60, 0, 0, 287, 289, 5, 33, 0, 0, 288, 286, 1, 0, 0, 0, 289, 292, 1,
		0, 0, 0, 290, 288, 1, 0, 0, 0, 290, 291, 1, 0, 0, 0, 291, 295, 1, 0, 0,
		0, 292, 290, 1, 0, 0, 0, 293, 295, 5, 31, 0, 0, 294, 285, 1, 0, 0, 0, 294,
		293, 1, 0, 0, 0, 295, 35, 1, 0, 0, 0, 296, 308, 5, 54, 0, 0, 297, 302,
		3, 38, 19, 0, 298, 299, 5, 58, 0, 0, 299, 301, 3, 38, 19, 0, 300, 298,
		1, 0, 0, 0, 301, 304, 1, 0, 0, 0, 302, 300, 1, 0, 0, 0, 302, 303, 1, 0,
		0, 0, 303, 306, 1, 0, 0, 0, 304, 302, 1, 0, 0, 0, 305, 307, 5, 58, 0, 0,
		306, 305, 1, 0, 0, 0, 306, 307, 1, 0, 0, 0, 307, 309, 1, 0, 0, 0, 308,
		297, 1, 0, 0, 0, 308, 309, 1, 0, 0, 0, 309, 310, 1, 0, 0, 0, 310, 311,
		5, 55, 0, 0, 311, 37, 1, 0, 0, 0, 312, 314, 5, 33, 0, 0, 313, 315, 3, 22,
		11, 0, 314, 313, 1, 0, 0, 0, 314, 315, 1, 0, 0, 0, 315, 318, 1, 0, 0, 0,
		316, 317, 5, 61, 0, 0, 317, 319, 3, 58, 29, 0, 318, 316, 1, 0, 0, 0, 318,
		319, 1, 0, 0, 0, 319, 39, 1, 0, 0, 0, 320, 324, 5, 52, 0, 0, 321, 323,
		3, 42, 21, 0, 322, 321, 1, 0, 0, 0, 323, 326, 1, 0, 0, 0, 324, 322, 1,
		0, 0, 0, 324, 325, 1, 0, 0, 0, 325, 327, 1, 0, 0, 0, 326, 324, 1, 0, 0,
		0, 327, 328, 5, 53, 0, 0, 328, 41, 1, 0, 0, 0, 329, 331, 3, 44, 22, 0,
		330, 332, 7, 1, 0, 0, 331, 330, 1, 0, 0, 0, 331, 332, 1, 0, 0, 0, 332,
		43, 1, 0, 0, 0, 333, 415, 3, 20, 10, 0, 334, 415, 3, 18, 9, 0, 335, 336,
		5, 7, 0, 0, 336, 337, 5, 54, 0, 0, 337, 338, 3, 60, 30, 0, 338, 339, 5,
		55, 0, 0, 339, 342, 3, 46, 23, 0, 340, 341, 5, 8, 0, 0, 341, 343, 3, 48,
		24, 0, 342, 340, 1, 0, 0, 0, 342, 343, 1, 0, 0, 0, 343, 415, 1, 0, 0, 0,
		344, 345, 5, 9, 0, 0, 345, 346, 5, 54, 0, 0, 346, 347, 3, 60, 30, 0, 347,
		348, 5, 55, 0, 0, 348, 349, 3, 40, 20, 0, 349, 415, 1, 0, 0, 0, 350, 351,
		5, 10, 0, 0, 351, 354, 5, 33, 0, 0, 352, 353, 5, 58, 0, 0, 353, 355, 5,
		33, 0, 0, 354, 352, 1, 0, 0, 0, 354, 355, 1, 0, 0, 0, 355, 356, 1, 0, 0,
		0, 356, 357, 5, 11, 0, 0, 357, 358, 3, 60, 30, 0, 358, 359, 3, 40, 20,
		0, 359, 415, 1, 0, 0, 0, 360, 362, 5, 12, 0, 0, 361, 363, 3, 58, 29, 0,
		362, 361, 1, 0, 0, 0, 362, 363, 1, 0, 0, 0, 363, 415, 1, 0, 0, 0, 364,
		415, 5, 13, 0, 0, 365, 415, 5, 14, 0, 0, 366, 415, 5, 24, 0, 0, 367, 368,
		5, 15, 0, 0, 368, 415, 3, 8, 4, 0, 369, 415, 3, 8, 4, 0, 370, 371, 5, 16,
		0, 0, 371, 372, 3, 24, 12, 0, 372, 373, 5, 61, 0, 0, 373, 375, 3, 58, 29,
		0, 374, 376, 3, 52, 26, 0, 375, 374, 1, 0, 0, 0, 375, 376, 1, 0, 0, 0,
		376, 415, 1, 0, 0, 0, 377, 378, 7, 2, 0, 0, 378, 415, 3, 40, 20, 0, 379,
		380, 5, 20, 0, 0, 380, 415, 3, 60, 30, 0, 381, 382, 5, 21, 0, 0, 382, 415,
		3, 50, 25, 0, 383, 384, 5, 33, 0, 0, 384, 387, 3, 22, 11, 0, 385, 386,
		5, 61, 0, 0, 386, 388, 3, 58, 29, 0, 387, 385, 1, 0, 0, 0, 387, 388, 1,
		0, 0, 0, 388, 415, 1, 0, 0, 0, 389, 390, 3, 56, 28, 0, 390, 391, 5, 61,
		0, 0, 391, 392, 3, 58, 29, 0, 392, 415, 1, 0, 0, 0, 393, 396, 5, 33, 0,
		0, 394, 395, 5, 60, 0, 0, 395, 397, 5, 33, 0, 0, 396, 394, 1, 0, 0, 0,
		397, 398, 1, 0, 0, 0, 398, 396, 1, 0, 0, 0, 398, 399, 1, 0, 0, 0, 399,
		400, 1, 0, 0, 0, 400, 401, 5, 61, 0, 0, 401, 415, 3, 58, 29, 0, 402, 403,
		7, 3, 0, 0, 403, 404, 5, 56, 0, 0, 404, 405, 3, 66, 33, 0, 405, 406, 5,
		57, 0, 0, 406, 407, 5, 61, 0, 0, 407, 408, 3, 58, 29, 0, 408, 415, 1, 0,
		0, 0, 409, 410, 3, 60, 30, 0, 410, 411, 7, 4, 0, 0, 411, 412, 3, 60, 30,
		0, 412, 415, 1, 0, 0, 0, 413, 415, 3, 60, 30, 0, 414, 333, 1, 0, 0, 0,
		414, 334, 1, 0, 0, 0, 414, 335, 1, 0, 0, 0, 414, 344, 1, 0, 0, 0, 414,
		350, 1, 0, 0, 0, 414, 360, 1, 0, 0, 0, 414, 364, 1, 0, 0, 0, 414, 365,
		1, 0, 0, 0, 414, 366, 1, 0, 0, 0, 414, 367, 1, 0, 0, 0, 414, 369, 1, 0,
		0, 0, 414, 370, 1, 0, 0, 0, 414, 377, 1, 0, 0, 0, 414, 379, 1, 0, 0, 0,
		414, 381, 1, 0, 0, 0, 414, 383, 1, 0, 0, 0, 414, 389, 1, 0, 0, 0, 414,
		393, 1, 0, 0, 0, 414, 402, 1, 0, 0, 0, 414, 409, 1, 0, 0, 0, 414, 413,
		1, 0, 0, 0, 415, 45, 1, 0, 0, 0, 416, 419, 3, 40, 20, 0, 417, 419, 3, 50,
		25, 0, 418, 416, 1, 0, 0, 0, 418, 417, 1, 0, 0, 0, 419, 47, 1, 0, 0, 0,
		420, 424, 3, 54, 27, 0, 421, 424, 3, 40, 20, 0, 422, 424, 3, 50, 25, 0,
		423, 420, 1, 0, 0, 0, 423, 421, 1, 0, 0, 0, 423, 422, 1, 0, 0, 0, 424,
		49, 1, 0, 0, 0, 425, 429, 5, 33, 0, 0, 426, 428, 3, 62, 31, 0, 427, 426,
		1, 0, 0, 0, 428, 431, 1, 0, 0, 0, 429, 427, 1, 0, 0, 0, 429, 430, 1, 0,
		0, 0, 430, 51, 1, 0, 0, 0, 431, 429, 1, 0, 0, 0, 432, 433, 4, 26, 0, 0,
		433, 434, 5, 33, 0, 0, 434, 435, 3, 64, 32, 0, 435, 53, 1, 0, 0, 0, 436,
		437, 5, 7, 0, 0, 437, 438, 5, 54, 0, 0, 438, 439, 3, 60, 30, 0, 439, 440,
		5, 55, 0, 0, 440, 443, 3, 46, 23, 0, 441, 442, 5, 8, 0, 0, 442, 444, 3,
		48, 24, 0, 443, 441, 1, 0, 0, 0, 443, 444, 1, 0, 0, 0, 444, 55, 1, 0, 0,
		0, 445, 449, 3, 14, 7, 0, 446, 449, 5, 31, 0, 0, 447, 449, 5, 35, 0, 0,
		448, 445, 1, 0, 0, 0, 448, 446, 1, 0, 0, 0, 448, 447, 1, 0, 0, 0, 449,
		57, 1, 0, 0, 0, 450, 451, 5, 7, 0, 0, 451, 452, 5, 54, 0, 0, 452, 453,
		3, 60, 30, 0, 453, 454, 5, 55, 0, 0, 454, 457, 3, 58, 29, 0, 455, 456,
		5, 8, 0, 0, 456, 458, 3, 58, 29, 0, 457, 455, 1, 0, 0, 0, 457, 458, 1,
		0, 0, 0, 458, 462, 1, 0, 0, 0, 459, 462, 5, 24, 0, 0, 460, 462, 3, 60,
		30, 0, 461, 450, 1, 0, 0, 0, 461, 459, 1, 0, 0, 0, 461, 460, 1, 0, 0, 0,
		462, 59, 1, 0, 0, 0, 463, 464, 6, 30, -1, 0, 464, 465, 7, 5, 0, 0, 465,
		472, 3, 60, 30, 12, 466, 467, 5, 37, 0, 0, 467, 472, 3, 60, 30, 8, 468,
		469, 5, 36, 0, 0, 469, 472, 3, 60, 30, 2, 470, 472, 3, 68, 34, 0, 471,
		463, 1, 0, 0, 0, 471, 466, 1, 0, 0, 0, 471, 468, 1, 0, 0, 0, 471, 470,
		1, 0, 0, 0, 472, 503, 1, 0, 0, 0, 473, 474, 10, 11, 0, 0, 474, 475, 7,
		6, 0, 0, 475, 502, 3, 60, 30, 12, 476, 477, 10, 10, 0, 0, 477, 478, 7,
		7, 0, 0, 478, 502, 3, 60, 30, 11, 479, 480, 10, 7, 0, 0, 480, 481, 7, 8,
		0, 0, 481, 502, 3, 60, 30, 8, 482, 483, 10, 6, 0, 0, 483, 484, 7, 9, 0,
		0, 484, 502, 3, 60, 30, 7, 485, 486, 10, 5, 0, 0, 486, 487, 5, 48, 0, 0,
		487, 502, 3, 60, 30, 6, 488, 489, 10, 4, 0, 0, 489, 490, 5, 49, 0, 0, 490,
		502, 3, 60, 30, 5, 491, 492, 10, 3, 0, 0, 492, 493, 5, 38, 0, 0, 493, 502,
		3, 60, 30, 4, 494, 495, 10, 13, 0, 0, 495, 502, 3, 62, 31, 0, 496, 497,
		10, 9, 0, 0, 497, 499, 5, 37, 0, 0, 498, 500, 3, 60, 30, 0, 499, 498, 1,
		0, 0, 0, 499, 500, 1, 0, 0, 0, 500, 502, 1, 0, 0, 0, 501, 473, 1, 0, 0,
		0, 501, 476, 1, 0, 0, 0, 501, 479, 1, 0, 0, 0, 501, 482, 1, 0, 0, 0, 501,
		485, 1, 0, 0, 0, 501, 488, 1, 0, 0, 0, 501, 491, 1, 0, 0, 0, 501, 494,
		1, 0, 0, 0, 501, 496, 1, 0, 0, 0, 502, 505, 1, 0, 0, 0, 503, 501, 1, 0,
		0, 0, 503, 504, 1, 0, 0, 0, 504, 61, 1, 0, 0, 0, 505, 503, 1, 0, 0, 0,
		506, 508, 5, 60, 0, 0, 507, 509, 3, 14, 7, 0, 508, 507, 1, 0, 0, 0, 508,
		509, 1, 0, 0, 0, 509, 528, 1, 0, 0, 0, 510, 528, 3, 64, 32, 0, 511, 523,
		5, 56, 0, 0, 512, 517, 3, 66, 33, 0, 513, 514, 5, 58, 0, 0, 514, 516, 3,
		66, 33, 0, 515, 513, 1, 0, 0, 0, 516, 519, 1, 0, 0, 0, 517, 515, 1, 0,
		0, 0, 517, 518, 1, 0, 0, 0, 518, 521, 1, 0, 0, 0, 519, 517, 1, 0, 0, 0,
		520, 522, 5, 58, 0, 0, 521, 520, 1, 0, 0, 0, 521, 522, 1, 0, 0, 0, 522,
		524, 1, 0, 0, 0, 523, 512, 1, 0, 0, 0, 523, 524, 1, 0, 0, 0, 524, 525,
		1, 0, 0, 0, 525, 528, 5, 57, 0, 0, 526, 528, 3, 40, 20, 0, 527, 506, 1,
		0, 0, 0, 527, 510, 1, 0, 0, 0, 527, 511, 1, 0, 0, 0, 527, 526, 1, 0, 0,
		0, 528, 63, 1, 0, 0, 0, 529, 541, 5, 54, 0, 0, 530, 535, 3, 66, 33, 0,
		531, 532, 5, 58, 0, 0, 532, 534, 3, 66, 33, 0, 533, 531, 1, 0, 0, 0, 534,
		537, 1, 0, 0, 0, 535, 533, 1, 0, 0, 0, 535, 536, 1, 0, 0, 0, 536, 539,
		1, 0, 0, 0, 537, 535, 1, 0, 0, 0, 538, 540, 5, 58, 0, 0, 539, 538, 1, 0,
		0, 0, 539, 540, 1, 0, 0, 0, 540, 542, 1, 0, 0, 0, 541, 530, 1, 0, 0, 0,
		541, 542, 1, 0, 0, 0, 542, 543, 1, 0, 0, 0, 543, 544, 5, 55, 0, 0, 544,
		65, 1, 0, 0, 0, 545, 546, 3, 14, 7, 0, 546, 547, 5, 61, 0, 0, 547, 549,
		1, 0, 0, 0, 548, 545, 1, 0, 0, 0, 548, 549, 1, 0, 0, 0, 549, 550, 1, 0,
		0, 0, 550, 551, 3, 58, 29, 0, 551, 67, 1, 0, 0, 0, 552, 617, 5, 34, 0,
		0, 553, 617, 5, 35, 0, 0, 554, 617, 7, 10, 0, 0, 555, 556, 5, 21, 0, 0,
		556, 617, 3, 40, 20, 0, 557, 617, 5, 33, 0, 0, 558, 617, 5, 31, 0, 0, 559,
		617, 5, 32, 0, 0, 560, 561, 5, 70, 0, 0, 561, 617, 7, 11, 0, 0, 562, 563,
		5, 71, 0, 0, 563, 617, 5, 33, 0, 0, 564, 569, 7, 12, 0, 0, 565, 567, 5,
		65, 0, 0, 566, 565, 1, 0, 0, 0, 566, 567, 1, 0, 0, 0, 567, 568, 1, 0, 0,
		0, 568, 570, 5, 34, 0, 0, 569, 566, 1, 0, 0, 0, 569, 570, 1, 0, 0, 0, 570,
		573, 1, 0, 0, 0, 571, 573, 5, 72, 0, 0, 572, 564, 1, 0, 0, 0, 572, 571,
		1, 0, 0, 0, 573, 574, 1, 0, 0, 0, 574, 572, 1, 0, 0, 0, 574, 575, 1, 0,
		0, 0, 575, 617, 1, 0, 0, 0, 576, 577, 5, 54, 0, 0, 577, 578, 3, 58, 29,
		0, 578, 579, 5, 55, 0, 0, 579, 617, 1, 0, 0, 0, 580, 581, 5, 56, 0, 0,
		581, 582, 5, 10, 0, 0, 582, 585, 5, 33, 0, 0, 583, 584, 5, 58, 0, 0, 584,
		586, 5, 33, 0, 0, 585, 583, 1, 0, 0, 0, 585, 586, 1, 0, 0, 0, 586, 587,
		1, 0, 0, 0, 587, 588, 5, 11, 0, 0, 588, 591, 3, 60, 30, 0, 589, 590, 5,
		7, 0, 0, 590, 592, 3, 60, 30, 0, 591, 589, 1, 0, 0, 0, 591, 592, 1, 0,
		0, 0, 592, 593, 1, 0, 0, 0, 593, 594, 5, 59, 0, 0, 594, 595, 3, 58, 29,
		0, 595, 596, 5, 57, 0, 0, 596, 617, 1, 0, 0, 0, 597, 611, 5, 56, 0, 0,
		598, 605, 3, 58, 29, 0, 599, 601, 5, 58, 0, 0, 600, 599, 1, 0, 0, 0, 600,
		601, 1, 0, 0, 0, 601, 602, 1, 0, 0, 0, 602, 604, 3, 58, 29, 0, 603, 600,
		1, 0, 0, 0, 604, 607, 1, 0, 0, 0, 605, 603, 1, 0, 0, 0, 605, 606, 1, 0,
		0, 0, 606, 609, 1, 0, 0, 0, 607, 605, 1, 0, 0, 0, 608, 610, 5, 58, 0, 0,
		609, 608, 1, 0, 0, 0, 609, 610, 1, 0, 0, 0, 610, 612, 1, 0, 0, 0, 611,
		598, 1, 0, 0, 0, 611, 612, 1, 0, 0, 0, 612, 613, 1, 0, 0, 0, 613, 617,
		5, 57, 0, 0, 614, 617, 3, 40, 20, 0, 615, 617, 5, 1, 0, 0, 616, 552, 1,
		0, 0, 0, 616, 553, 1, 0, 0, 0, 616, 554, 1, 0, 0, 0, 616, 555, 1, 0, 0,
		0, 616, 557, 1, 0, 0, 0, 616, 558, 1, 0, 0, 0, 616, 559, 1, 0, 0, 0, 616,
		560, 1, 0, 0, 0, 616, 562, 1, 0, 0, 0, 616, 572, 1, 0, 0, 0, 616, 576,
		1, 0, 0, 0, 616, 580, 1, 0, 0, 0, 616, 597, 1, 0, 0, 0, 616, 614, 1, 0,
		0, 0, 616, 615, 1, 0, 0, 0, 617, 69, 1, 0, 0, 0, 82, 73, 82, 91, 97, 114,
		121, 123, 126, 130, 133, 141, 148, 150, 153, 155, 160, 165, 172, 177, 181,
		186, 190, 200, 205, 214, 219, 223, 233, 238, 242, 256, 260, 264, 273, 283,
		290, 294, 302, 306, 308, 314, 318, 324, 331, 342, 354, 362, 375, 387, 398,
		414, 418, 423, 429, 443, 448, 457, 461, 471, 499, 501, 503, 508, 517, 521,
		523, 527, 535, 539, 541, 548, 566, 569, 572, 574, 585, 591, 600, 605, 609,
		611, 616,
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
	MMSParserELLIPSIS             = 36
	MMSParserDOTDOT               = 37
	MMSParserQQ                   = 38
	MMSParserPLUS_EQ              = 39
	MMSParserMINUS_EQ             = 40
	MMSParserSTAR_EQ              = 41
	MMSParserSLASH_EQ             = 42
	MMSParserPCT_EQ               = 43
	MMSParserEQEQ                 = 44
	MMSParserNEQ                  = 45
	MMSParserLE                   = 46
	MMSParserGE                   = 47
	MMSParserANDAND               = 48
	MMSParserOROR                 = 49
	MMSParserPIPE                 = 50
	MMSParserQUESTION             = 51
	MMSParserLBRACE               = 52
	MMSParserRBRACE               = 53
	MMSParserLPAREN               = 54
	MMSParserRPAREN               = 55
	MMSParserLBRACK               = 56
	MMSParserRBRACK               = 57
	MMSParserCOMMA                = 58
	MMSParserCOLON                = 59
	MMSParserDOT                  = 60
	MMSParserASSIGN               = 61
	MMSParserLT                   = 62
	MMSParserGT                   = 63
	MMSParserPLUS                 = 64
	MMSParserMINUS                = 65
	MMSParserSTAR                 = 66
	MMSParserSLASH                = 67
	MMSParserPERCENT              = 68
	MMSParserBANG                 = 69
	MMSParserHASH                 = 70
	MMSParserAT                   = 71
	MMSParserREL_NAME             = 72
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
	MMSParserRULE_typeField      = 15
	MMSParserRULE_mapKey         = 16
	MMSParserRULE_typeName       = 17
	MMSParserRULE_params         = 18
	MMSParserRULE_param          = 19
	MMSParserRULE_block          = 20
	MMSParserRULE_blockItem      = 21
	MMSParserRULE_statement      = 22
	MMSParserRULE_thenPart       = 23
	MMSParserRULE_elsePart       = 24
	MMSParserRULE_bareCall       = 25
	MMSParserRULE_hint           = 26
	MMSParserRULE_ifTail         = 27
	MMSParserRULE_fieldKey       = 28
	MMSParserRULE_value          = 29
	MMSParserRULE_expr           = 30
	MMSParserRULE_postfix        = 31
	MMSParserRULE_args           = 32
	MMSParserRULE_arg            = 33
	MMSParserRULE_primary        = 34
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
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8589935736) != 0 {
		{
			p.SetState(70)
			p.TopItem()
		}

		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(76)
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
		p.SetState(78)
		p.expr(0)
	}
	{
		p.SetState(79)
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

	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MMSParserTYPE, MMSParserEXTERN, MMSParserIDENT:
		localctx = NewDeclItemContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(82)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MMSParserEXTERN {
			{
				p.SetState(81)
				p.Match(MMSParserEXTERN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(84)
			p.Declaration()
		}

	case MMSParserCONST:
		localctx = NewConstItemContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(85)
			p.ConstDecl()
		}

	case MMSParserLET:
		localctx = NewLetItemContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(86)
			p.LetDecl()
		}

	case MMSParserFOR:
		localctx = NewForItemContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(87)
			p.Match(MMSParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(88)
			p.Match(MMSParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MMSParserCOMMA {
			{
				p.SetState(89)
				p.Match(MMSParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(90)
				p.Match(MMSParserIDENT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(93)
			p.Match(MMSParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(94)
			p.expr(0)
		}
		{
			p.SetState(95)
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

	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(99)

			var _m = p.Match(MMSParserTYPE)

			localctx.(*DeclarationContext).kind = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(100)

			var _x = p.DeclName()

			localctx.(*DeclarationContext).name = _x
		}
		{
			p.SetState(101)
			p.Match(MMSParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(102)
			p.TypeExpr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(104)

			var _m = p.Match(MMSParserTYPE)

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
		{
			p.SetState(106)
			p.Block()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(108)

			var _m = p.Match(MMSParserIDENT)

			localctx.(*DeclarationContext).kind = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(109)

			var _x = p.DeclName()

			localctx.(*DeclarationContext).name = _x
		}
		p.SetState(112)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == MMSParserCOMMA {
			{
				p.SetState(110)
				p.Match(MMSParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(111)
				p.DeclName()
			}

			p.SetState(114)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(116)

			var _m = p.Match(MMSParserIDENT)

			localctx.(*DeclarationContext).kind = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(117)

			var _x = p.DeclName()

			localctx.(*DeclarationContext).name = _x
		}
		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		switch p.GetTokenStream().LA(1) {
		case MMSParserCOLON:
			{
				p.SetState(118)
				p.TypeAnnotation()
			}

		case MMSParserLPAREN:
			{
				p.SetState(119)
				p.Params()
			}
			p.SetState(121)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == MMSParserCOLON {
				{
					p.SetState(120)
					p.TypeAnnotation()
				}

			}

		case MMSParserEOF, MMSParserCONST, MMSParserTYPE, MMSParserEXTERN, MMSParserLET, MMSParserFOR, MMSParserIDENT, MMSParserLBRACE, MMSParserASSIGN:

		default:
		}
		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MMSParserLBRACE || _la == MMSParserASSIGN {
			{
				p.SetState(125)
				p.DeclBody()
			}

		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(128)

			var _m = p.Match(MMSParserIDENT)

			localctx.(*DeclarationContext).kind = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MMSParserLPAREN {
			{
				p.SetState(129)
				p.Params()
			}

		}
		{
			p.SetState(132)
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

	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(135)

			var _m = p.Match(MMSParserIDENT)

			localctx.(*InnerDeclContext).kind = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(136)

			var _x = p.DeclName()

			localctx.(*InnerDeclContext).name = _x
		}
		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(137)
					p.Match(MMSParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(138)
					p.DeclName()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(141)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(143)

			var _m = p.Match(MMSParserIDENT)

			localctx.(*InnerDeclContext).kind = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(144)

			var _x = p.DeclName()

			localctx.(*InnerDeclContext).name = _x
		}
		p.SetState(150)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(145)
				p.TypeAnnotation()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) == 2 {
			{
				p.SetState(146)
				p.Params()
			}
			p.SetState(148)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == MMSParserCOLON {
				{
					p.SetState(147)
					p.TypeAnnotation()
				}

			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(153)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(152)
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
	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MMSParserASSIGN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(157)
			p.Match(MMSParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(158)
			p.Value()
		}

	case MMSParserLBRACE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(159)
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
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MMSParserTYPE, MMSParserPRIMARY, MMSParserAS, MMSParserCHECK, MMSParserEMIT, MMSParserSETUP, MMSParserEXPECT, MMSParserRUN, MMSParserIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(162)
			p.Word()
		}

	case MMSParserRESLOC:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(163)
			p.Match(MMSParserRESLOC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MMSParserTNAME:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(164)
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
		p.SetState(167)
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
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MMSParserASSIGN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(169)
			p.Match(MMSParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(170)
			p.Value()
		}

	case MMSParserLBRACE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(171)
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
		p.SetState(174)
		p.Match(MMSParserCONST)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(175)
		p.Match(MMSParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MMSParserCOLON {
		{
			p.SetState(176)
			p.TypeAnnotation()
		}

	}
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MMSParserASSIGN {
		{
			p.SetState(179)
			p.Match(MMSParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(180)
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
		p.SetState(183)
		p.Match(MMSParserLET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(184)
		p.Match(MMSParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MMSParserCOLON {
		{
			p.SetState(185)
			p.TypeAnnotation()
		}

	}
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MMSParserASSIGN {
		{
			p.SetState(188)
			p.Match(MMSParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(189)
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
		p.SetState(192)
		p.Match(MMSParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(193)
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
		p.SetState(195)
		p.TypeArm()
	}
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MMSParserPIPE {
		{
			p.SetState(196)
			p.Match(MMSParserPIPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(197)
			p.TypeArm()
		}

		p.SetState(202)
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
		p.SetState(203)
		p.TypeAtom()
	}
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MMSParserQUESTION {
		{
			p.SetState(204)
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

type NumberTypeContext struct {
	TypeAtomContext
}

func NewNumberTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberTypeContext {
	var p = new(NumberTypeContext)

	InitEmptyTypeAtomContext(&p.TypeAtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeAtomContext))

	return p
}

func (s *NumberTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberTypeContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(MMSParserNUMBER, 0)
}

func (s *NumberTypeContext) MINUS() antlr.TerminalNode {
	return s.GetToken(MMSParserMINUS, 0)
}

type RecordTypeContext struct {
	TypeAtomContext
}

func NewRecordTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RecordTypeContext {
	var p = new(RecordTypeContext)

	InitEmptyTypeAtomContext(&p.TypeAtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeAtomContext))

	return p
}

func (s *RecordTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecordTypeContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(MMSParserLBRACE, 0)
}

func (s *RecordTypeContext) AllTypeField() []ITypeFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeFieldContext); ok {
			len++
		}
	}

	tst := make([]ITypeFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeFieldContext); ok {
			tst[i] = t.(ITypeFieldContext)
			i++
		}
	}

	return tst
}

func (s *RecordTypeContext) TypeField(i int) ITypeFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeFieldContext); ok {
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

	return t.(ITypeFieldContext)
}

func (s *RecordTypeContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(MMSParserRBRACE, 0)
}

func (s *RecordTypeContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MMSParserCOMMA)
}

func (s *RecordTypeContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MMSParserCOMMA, i)
}

func (s *RecordTypeContext) AllSEP() []antlr.TerminalNode {
	return s.GetTokens(MMSParserSEP)
}

func (s *RecordTypeContext) SEP(i int) antlr.TerminalNode {
	return s.GetToken(MMSParserSEP, i)
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

func (s *MapTypeContext) MapKey() IMapKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapKeyContext)
}

func (s *MapTypeContext) COLON() antlr.TerminalNode {
	return s.GetToken(MMSParserCOLON, 0)
}

func (s *MapTypeContext) TypeExpr() ITypeExprContext {
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

	var _alt int

	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNamedTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(207)
			p.TypeName()
		}
		p.SetState(219)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MMSParserLT {
			{
				p.SetState(208)
				p.Match(MMSParserLT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(209)
				p.TypeExpr()
			}
			p.SetState(214)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == MMSParserCOMMA {
				{
					p.SetState(210)
					p.Match(MMSParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(211)
					p.TypeExpr()
				}

				p.SetState(216)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(217)
				p.Match(MMSParserGT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 2:
		localctx = NewLiteralTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(221)
			p.Match(MMSParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewNumberTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		p.SetState(223)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MMSParserMINUS {
			{
				p.SetState(222)
				p.Match(MMSParserMINUS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(225)
			p.Match(MMSParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewListTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(226)
			p.Match(MMSParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(227)
			p.TypeExpr()
		}
		{
			p.SetState(228)
			p.Match(MMSParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewRecordTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(230)
			p.Match(MMSParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(231)
			p.TypeField()
		}
		p.SetState(238)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				p.SetState(233)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == MMSParserSEP || _la == MMSParserCOMMA {
					{
						p.SetState(232)
						_la = p.GetTokenStream().LA(1)

						if !(_la == MMSParserSEP || _la == MMSParserCOMMA) {
							p.GetErrorHandler().RecoverInline(p)
						} else {
							p.GetErrorHandler().ReportMatch(p)
							p.Consume()
						}
					}

				}
				{
					p.SetState(235)
					p.TypeField()
				}

			}
			p.SetState(240)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(242)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MMSParserSEP || _la == MMSParserCOMMA {
			{
				p.SetState(241)
				_la = p.GetTokenStream().LA(1)

				if !(_la == MMSParserSEP || _la == MMSParserCOMMA) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(244)
			p.Match(MMSParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewMapTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(246)
			p.Match(MMSParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(247)
			p.MapKey()
		}
		{
			p.SetState(248)
			p.Match(MMSParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(249)
			p.TypeExpr()
		}
		{
			p.SetState(250)
			p.Match(MMSParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewParenTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(252)
			p.Match(MMSParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(253)
			p.TypeExpr()
		}
		{
			p.SetState(254)
			p.Match(MMSParserRPAREN)
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

// ITypeFieldContext is an interface to support dynamic dispatch.
type ITypeFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	TypeAnnotation() ITypeAnnotationContext
	ASSIGN() antlr.TerminalNode
	Value() IValueContext

	// IsTypeFieldContext differentiates from other interfaces.
	IsTypeFieldContext()
}

type TypeFieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeFieldContext() *TypeFieldContext {
	var p = new(TypeFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_typeField
	return p
}

func InitEmptyTypeFieldContext(p *TypeFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_typeField
}

func (*TypeFieldContext) IsTypeFieldContext() {}

func NewTypeFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeFieldContext {
	var p = new(TypeFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_typeField

	return p
}

func (s *TypeFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeFieldContext) IDENT() antlr.TerminalNode {
	return s.GetToken(MMSParserIDENT, 0)
}

func (s *TypeFieldContext) TypeAnnotation() ITypeAnnotationContext {
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

func (s *TypeFieldContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(MMSParserASSIGN, 0)
}

func (s *TypeFieldContext) Value() IValueContext {
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

func (s *TypeFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MMSParser) TypeField() (localctx ITypeFieldContext) {
	localctx = NewTypeFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, MMSParserRULE_typeField)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(258)
		p.Match(MMSParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MMSParserCOLON {
		{
			p.SetState(259)
			p.TypeAnnotation()
		}

	}
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MMSParserASSIGN {
		{
			p.SetState(262)
			p.Match(MMSParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(263)
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

// IMapKeyContext is an interface to support dynamic dispatch.
type IMapKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TypeName() ITypeNameContext
	LT() antlr.TerminalNode
	AllTypeExpr() []ITypeExprContext
	TypeExpr(i int) ITypeExprContext
	GT() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	RESLOC() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode

	// IsMapKeyContext differentiates from other interfaces.
	IsMapKeyContext()
}

type MapKeyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapKeyContext() *MapKeyContext {
	var p = new(MapKeyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_mapKey
	return p
}

func InitEmptyMapKeyContext(p *MapKeyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_mapKey
}

func (*MapKeyContext) IsMapKeyContext() {}

func NewMapKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapKeyContext {
	var p = new(MapKeyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_mapKey

	return p
}

func (s *MapKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *MapKeyContext) TypeName() ITypeNameContext {
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

func (s *MapKeyContext) LT() antlr.TerminalNode {
	return s.GetToken(MMSParserLT, 0)
}

func (s *MapKeyContext) AllTypeExpr() []ITypeExprContext {
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

func (s *MapKeyContext) TypeExpr(i int) ITypeExprContext {
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

func (s *MapKeyContext) GT() antlr.TerminalNode {
	return s.GetToken(MMSParserGT, 0)
}

func (s *MapKeyContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MMSParserCOMMA)
}

func (s *MapKeyContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MMSParserCOMMA, i)
}

func (s *MapKeyContext) RESLOC() antlr.TerminalNode {
	return s.GetToken(MMSParserRESLOC, 0)
}

func (s *MapKeyContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MMSParserLPAREN, 0)
}

func (s *MapKeyContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MMSParserRPAREN, 0)
}

func (s *MapKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MMSParser) MapKey() (localctx IMapKeyContext) {
	localctx = NewMapKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, MMSParserRULE_mapKey)
	var _la int

	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(266)
			p.TypeName()
		}
		{
			p.SetState(267)
			p.Match(MMSParserLT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(268)
			p.TypeExpr()
		}
		p.SetState(273)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MMSParserCOMMA {
			{
				p.SetState(269)
				p.Match(MMSParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(270)
				p.TypeExpr()
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
			p.Match(MMSParserGT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(278)
			p.Match(MMSParserRESLOC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(279)
			p.Match(MMSParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(280)
			p.TypeExpr()
		}
		{
			p.SetState(281)
			p.Match(MMSParserRPAREN)
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
	p.EnterRule(localctx, 34, MMSParserRULE_typeName)
	var _la int

	p.SetState(294)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MMSParserIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(285)
			p.Match(MMSParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(290)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MMSParserDOT {
			{
				p.SetState(286)
				p.Match(MMSParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(287)
				p.Match(MMSParserIDENT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(292)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case MMSParserRESLOC:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(293)
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
	p.EnterRule(localctx, 36, MMSParserRULE_params)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(296)
		p.Match(MMSParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(308)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MMSParserIDENT {
		{
			p.SetState(297)
			p.Param()
		}
		p.SetState(302)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(298)
					p.Match(MMSParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(299)
					p.Param()
				}

			}
			p.SetState(304)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(306)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MMSParserCOMMA {
			{
				p.SetState(305)
				p.Match(MMSParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(310)
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
	p.EnterRule(localctx, 38, MMSParserRULE_param)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(312)
		p.Match(MMSParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(314)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MMSParserCOLON {
		{
			p.SetState(313)
			p.TypeAnnotation()
		}

	}
	p.SetState(318)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MMSParserASSIGN {
		{
			p.SetState(316)
			p.Match(MMSParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(317)
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
	p.EnterRule(localctx, 40, MMSParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(320)
		p.Match(MMSParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(324)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&94575864938755802) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&2019) != 0) {
		{
			p.SetState(321)
			p.BlockItem()
		}

		p.SetState(326)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(327)
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
	p.EnterRule(localctx, 42, MMSParserRULE_blockItem)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(329)
		p.Statement()
	}
	p.SetState(331)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MMSParserSEP || _la == MMSParserCOMMA {
		{
			p.SetState(330)
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
	target antlr.Token
}

func NewIndexAssignStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexAssignStmtContext {
	var p = new(IndexAssignStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *IndexAssignStmtContext) GetTarget() antlr.Token { return s.target }

func (s *IndexAssignStmtContext) SetTarget(v antlr.Token) { s.target = v }

func (s *IndexAssignStmtContext) GetRuleContext() antlr.RuleContext {
	return s
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

func (s *IndexAssignStmtContext) IDENT() antlr.TerminalNode {
	return s.GetToken(MMSParserIDENT, 0)
}

func (s *IndexAssignStmtContext) TNAME() antlr.TerminalNode {
	return s.GetToken(MMSParserTNAME, 0)
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
	p.EnterRule(localctx, 44, MMSParserRULE_statement)
	var _la int

	p.SetState(414)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 50, p.GetParserRuleContext()) {
	case 1:
		localctx = NewLetStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(333)
			p.LetDecl()
		}

	case 2:
		localctx = NewConstStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(334)
			p.ConstDecl()
		}

	case 3:
		localctx = NewIfStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(335)
			p.Match(MMSParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(336)
			p.Match(MMSParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(337)
			p.expr(0)
		}
		{
			p.SetState(338)
			p.Match(MMSParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(339)
			p.ThenPart()
		}
		p.SetState(342)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MMSParserELSE {
			{
				p.SetState(340)
				p.Match(MMSParserELSE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(341)
				p.ElsePart()
			}

		}

	case 4:
		localctx = NewWhileStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(344)
			p.Match(MMSParserWHILE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(345)
			p.Match(MMSParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(346)
			p.expr(0)
		}
		{
			p.SetState(347)
			p.Match(MMSParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(348)
			p.Block()
		}

	case 5:
		localctx = NewForStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(350)
			p.Match(MMSParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(351)
			p.Match(MMSParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(354)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MMSParserCOMMA {
			{
				p.SetState(352)
				p.Match(MMSParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(353)
				p.Match(MMSParserIDENT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(356)
			p.Match(MMSParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(357)
			p.expr(0)
		}
		{
			p.SetState(358)
			p.Block()
		}

	case 6:
		localctx = NewReturnStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(360)
			p.Match(MMSParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(362)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 46, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(361)
				p.Value()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 7:
		localctx = NewBreakStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(364)
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
			p.SetState(365)
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
			p.SetState(366)
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
			p.SetState(367)
			p.Match(MMSParserPRIMARY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(368)
			p.InnerDecl()
		}

	case 11:
		localctx = NewDeclStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(369)
			p.InnerDecl()
		}

	case 12:
		localctx = NewAsStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(370)
			p.Match(MMSParserAS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(371)
			p.TypeExpr()
		}
		{
			p.SetState(372)
			p.Match(MMSParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(373)
			p.Value()
		}
		p.SetState(375)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 47, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(374)
				p.Hint()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 13:
		localctx = NewSectionStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(377)

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
			p.SetState(378)
			p.Block()
		}

	case 14:
		localctx = NewExpectStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(379)
			p.Match(MMSParserEXPECT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(380)
			p.expr(0)
		}

	case 15:
		localctx = NewRunStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(381)
			p.Match(MMSParserRUN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(382)
			p.BareCall()
		}

	case 16:
		localctx = NewFieldDeclStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(383)
			p.Match(MMSParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(384)
			p.TypeAnnotation()
		}
		p.SetState(387)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MMSParserASSIGN {
			{
				p.SetState(385)
				p.Match(MMSParserASSIGN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(386)
				p.Value()
			}

		}

	case 17:
		localctx = NewFieldStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(389)
			p.FieldKey()
		}
		{
			p.SetState(390)
			p.Match(MMSParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(391)
			p.Value()
		}

	case 18:
		localctx = NewPathAssignStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(393)
			p.Match(MMSParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(396)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == MMSParserDOT {
			{
				p.SetState(394)
				p.Match(MMSParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(395)
				p.Match(MMSParserIDENT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(398)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(400)
			p.Match(MMSParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(401)
			p.Value()
		}

	case 19:
		localctx = NewIndexAssignStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(402)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*IndexAssignStmtContext).target = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == MMSParserTNAME || _la == MMSParserIDENT) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*IndexAssignStmtContext).target = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(403)
			p.Match(MMSParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(404)
			p.Arg()
		}
		{
			p.SetState(405)
			p.Match(MMSParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(406)
			p.Match(MMSParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(407)
			p.Value()
		}

	case 20:
		localctx = NewCompoundAssignStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(409)
			p.expr(0)
		}
		{
			p.SetState(410)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*CompoundAssignStmtContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17042430230528) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*CompoundAssignStmtContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(411)
			p.expr(0)
		}

	case 21:
		localctx = NewExprStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(413)
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
	p.EnterRule(localctx, 46, MMSParserRULE_thenPart)
	p.SetState(418)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MMSParserLBRACE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(416)
			p.Block()
		}

	case MMSParserIDENT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(417)
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
	p.EnterRule(localctx, 48, MMSParserRULE_elsePart)
	p.SetState(423)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MMSParserIF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(420)
			p.IfTail()
		}

	case MMSParserLBRACE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(421)
			p.Block()
		}

	case MMSParserIDENT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(422)
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
	p.EnterRule(localctx, 50, MMSParserRULE_bareCall)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(425)
		p.Match(MMSParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(429)
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
			{
				p.SetState(426)
				p.Postfix()
			}

		}
		p.SetState(431)
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
	p.EnterRule(localctx, 52, MMSParserRULE_hint)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(432)

	if !(p.word("hint")) {
		p.SetError(antlr.NewFailedPredicateException(p, "p.word(\"hint\")", ""))
		goto errorExit
	}
	{
		p.SetState(433)
		p.Match(MMSParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(434)
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
	p.EnterRule(localctx, 54, MMSParserRULE_ifTail)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(436)
		p.Match(MMSParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(437)
		p.Match(MMSParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(438)
		p.expr(0)
	}
	{
		p.SetState(439)
		p.Match(MMSParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(440)
		p.ThenPart()
	}
	p.SetState(443)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MMSParserELSE {
		{
			p.SetState(441)
			p.Match(MMSParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(442)
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
	p.EnterRule(localctx, 56, MMSParserRULE_fieldKey)
	p.SetState(448)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MMSParserTYPE, MMSParserPRIMARY, MMSParserAS, MMSParserCHECK, MMSParserEMIT, MMSParserSETUP, MMSParserEXPECT, MMSParserRUN, MMSParserIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(445)
			p.Word()
		}

	case MMSParserRESLOC:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(446)
			p.Match(MMSParserRESLOC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MMSParserSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(447)
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
	COMMAND() antlr.TerminalNode

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

func (s *ValueContext) COMMAND() antlr.TerminalNode {
	return s.GetToken(MMSParserCOMMAND, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MMSParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, MMSParserRULE_value)
	p.SetState(461)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MMSParserIF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(450)
			p.Match(MMSParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(451)
			p.Match(MMSParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(452)
			p.expr(0)
		}
		{
			p.SetState(453)
			p.Match(MMSParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(454)
			p.Value()
		}
		p.SetState(457)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 56, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(455)
				p.Match(MMSParserELSE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(456)
				p.Value()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case MMSParserCOMMAND:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(459)
			p.Match(MMSParserCOMMAND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MMSParserMISSING, MMSParserRUN, MMSParserTRUE, MMSParserFALSE, MMSParserRESLOC, MMSParserTNAME, MMSParserIDENT, MMSParserNUMBER, MMSParserSTRING, MMSParserELLIPSIS, MMSParserDOTDOT, MMSParserLBRACE, MMSParserLPAREN, MMSParserLBRACK, MMSParserPLUS, MMSParserMINUS, MMSParserBANG, MMSParserHASH, MMSParserAT, MMSParserREL_NAME, MMSParserTILDE, MMSParserCARET:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(460)
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
	_startState := 60
	p.EnterRecursionRule(localctx, 60, MMSParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(471)
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
			p.SetState(464)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*UnaryExprContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&35) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*UnaryExprContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(465)
			p.expr(12)
		}

	case MMSParserDOTDOT:
		localctx = NewRangeToExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(466)
			p.Match(MMSParserDOTDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(467)
			p.expr(8)
		}

	case MMSParserELLIPSIS:
		localctx = NewSpreadExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(468)
			p.Match(MMSParserELLIPSIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(469)
			p.expr(2)
		}

	case MMSParserMISSING, MMSParserRUN, MMSParserTRUE, MMSParserFALSE, MMSParserRESLOC, MMSParserTNAME, MMSParserIDENT, MMSParserNUMBER, MMSParserSTRING, MMSParserLBRACE, MMSParserLPAREN, MMSParserLBRACK, MMSParserHASH, MMSParserAT, MMSParserREL_NAME, MMSParserTILDE, MMSParserCARET:
		localctx = NewPrimaryExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(470)
			p.Primary()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(503)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 61, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(501)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 60, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MMSParserRULE_expr)
				p.SetState(473)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(474)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64((_la-66)) & ^0x3f) == 0 && ((int64(1)<<(_la-66))&7) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(475)
					p.expr(12)
				}

			case 2:
				localctx = NewAddExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MMSParserRULE_expr)
				p.SetState(476)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(477)

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
					p.SetState(478)
					p.expr(11)
				}

			case 3:
				localctx = NewCompareExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MMSParserRULE_expr)
				p.SetState(479)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(480)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*CompareExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-4611474912194854912) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*CompareExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(481)
					p.expr(8)
				}

			case 4:
				localctx = NewEqualExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MMSParserRULE_expr)
				p.SetState(482)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(483)

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
					p.SetState(484)
					p.expr(7)
				}

			case 5:
				localctx = NewAndExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MMSParserRULE_expr)
				p.SetState(485)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(486)
					p.Match(MMSParserANDAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(487)
					p.expr(6)
				}

			case 6:
				localctx = NewOrExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MMSParserRULE_expr)
				p.SetState(488)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(489)
					p.Match(MMSParserOROR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(490)
					p.expr(5)
				}

			case 7:
				localctx = NewCoalesceExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MMSParserRULE_expr)
				p.SetState(491)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(492)
					p.Match(MMSParserQQ)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(493)
					p.expr(4)
				}

			case 8:
				localctx = NewPostfixExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MMSParserRULE_expr)
				p.SetState(494)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(495)
					p.Postfix()
				}

			case 9:
				localctx = NewRangeExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MMSParserRULE_expr)
				p.SetState(496)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(497)
					p.Match(MMSParserDOTDOT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(499)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 59, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(498)
						p.expr(0)
					}

				} else if p.HasError() { // JIM
					goto errorExit
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(505)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 61, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 62, MMSParserRULE_postfix)
	var _la int

	var _alt int

	p.SetState(527)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MMSParserDOT:
		localctx = NewMemberPostfixContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(506)
			p.Match(MMSParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(508)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 62, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(507)

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
			p.SetState(510)
			p.Args()
		}

	case MMSParserLBRACK:
		localctx = NewIndexPostfixContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(511)
			p.Match(MMSParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(523)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&94575864938725522) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&2019) != 0) {
			{
				p.SetState(512)
				p.Arg()
			}
			p.SetState(517)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 63, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
			for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				if _alt == 1 {
					{
						p.SetState(513)
						p.Match(MMSParserCOMMA)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}
					{
						p.SetState(514)
						p.Arg()
					}

				}
				p.SetState(519)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 63, p.GetParserRuleContext())
				if p.HasError() {
					goto errorExit
				}
			}
			p.SetState(521)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == MMSParserCOMMA {
				{
					p.SetState(520)
					p.Match(MMSParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}

		}
		{
			p.SetState(525)
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
			p.SetState(526)
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
	p.EnterRule(localctx, 64, MMSParserRULE_args)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(529)
		p.Match(MMSParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(541)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&94575864938725522) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&2019) != 0) {
		{
			p.SetState(530)
			p.Arg()
		}
		p.SetState(535)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 67, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(531)
					p.Match(MMSParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(532)
					p.Arg()
				}

			}
			p.SetState(537)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 67, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(539)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MMSParserCOMMA {
			{
				p.SetState(538)
				p.Match(MMSParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(543)
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
	p.EnterRule(localctx, 66, MMSParserRULE_arg)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(548)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 70, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(545)
			p.Word()
		}
		{
			p.SetState(546)
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
		p.SetState(550)
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

func (s *CoordinateContext) AllREL_NAME() []antlr.TerminalNode {
	return s.GetTokens(MMSParserREL_NAME)
}

func (s *CoordinateContext) REL_NAME(i int) antlr.TerminalNode {
	return s.GetToken(MMSParserREL_NAME, i)
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
	p.EnterRule(localctx, 68, MMSParserRULE_primary)
	var _la int

	var _alt int

	p.SetState(616)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 81, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNumberLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(552)
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
			p.SetState(553)
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
			p.SetState(554)
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
			p.SetState(555)
			p.Match(MMSParserRUN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(556)
			p.Block()
		}

	case 5:
		localctx = NewIdentRefContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(557)
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
			p.SetState(558)
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
			p.SetState(559)
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
			p.SetState(560)
			p.Match(MMSParserHASH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(561)
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
			p.SetState(562)
			p.Match(MMSParserAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(563)
			p.Match(MMSParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		localctx = NewCoordinateContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		p.SetState(572)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				p.SetState(572)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}

				switch p.GetTokenStream().LA(1) {
				case MMSParserTILDE, MMSParserCARET:
					{
						p.SetState(564)
						_la = p.GetTokenStream().LA(1)

						if !(_la == MMSParserTILDE || _la == MMSParserCARET) {
							p.GetErrorHandler().RecoverInline(p)
						} else {
							p.GetErrorHandler().ReportMatch(p)
							p.Consume()
						}
					}
					p.SetState(569)
					p.GetErrorHandler().Sync(p)

					if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 72, p.GetParserRuleContext()) == 1 {
						p.SetState(566)
						p.GetErrorHandler().Sync(p)
						if p.HasError() {
							goto errorExit
						}
						_la = p.GetTokenStream().LA(1)

						if _la == MMSParserMINUS {
							{
								p.SetState(565)
								p.Match(MMSParserMINUS)
								if p.HasError() {
									// Recognition error - abort rule
									goto errorExit
								}
							}

						}
						{
							p.SetState(568)
							p.Match(MMSParserNUMBER)
							if p.HasError() {
								// Recognition error - abort rule
								goto errorExit
							}
						}

					} else if p.HasError() { // JIM
						goto errorExit
					}

				case MMSParserREL_NAME:
					{
						p.SetState(571)
						p.Match(MMSParserREL_NAME)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				default:
					p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
					goto errorExit
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(574)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 74, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case 11:
		localctx = NewParenExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(576)
			p.Match(MMSParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(577)
			p.Value()
		}
		{
			p.SetState(578)
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
			p.SetState(580)
			p.Match(MMSParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(581)
			p.Match(MMSParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(582)
			p.Match(MMSParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(585)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MMSParserCOMMA {
			{
				p.SetState(583)
				p.Match(MMSParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(584)
				p.Match(MMSParserIDENT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(587)
			p.Match(MMSParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(588)
			p.expr(0)
		}
		p.SetState(591)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MMSParserIF {
			{
				p.SetState(589)
				p.Match(MMSParserIF)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(590)
				p.expr(0)
			}

		}
		{
			p.SetState(593)
			p.Match(MMSParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(594)
			p.Value()
		}
		{
			p.SetState(595)
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
			p.SetState(597)
			p.Match(MMSParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(611)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&94575864936661122) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&2019) != 0) {
			{
				p.SetState(598)
				p.Value()
			}
			p.SetState(605)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 78, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
			for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				if _alt == 1 {
					p.SetState(600)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_la = p.GetTokenStream().LA(1)

					if _la == MMSParserCOMMA {
						{
							p.SetState(599)
							p.Match(MMSParserCOMMA)
							if p.HasError() {
								// Recognition error - abort rule
								goto errorExit
							}
						}

					}
					{
						p.SetState(602)
						p.Value()
					}

				}
				p.SetState(607)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 78, p.GetParserRuleContext())
				if p.HasError() {
					goto errorExit
				}
			}
			p.SetState(609)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == MMSParserCOMMA {
				{
					p.SetState(608)
					p.Match(MMSParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}

		}
		{
			p.SetState(613)
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
			p.SetState(614)
			p.Block()
		}

	case 15:
		localctx = NewMissingLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(615)
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
	case 26:
		var t *HintContext = nil
		if localctx != nil {
			t = localctx.(*HintContext)
		}
		return p.Hint_Sempred(t, predIndex)

	case 30:
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
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 9)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
