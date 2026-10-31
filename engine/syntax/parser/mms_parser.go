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
		"", "'let'", "'if'", "'else'", "'while'", "'for'", "'in'", "'return'",
		"'break'", "'continue'", "'true'", "'false'", "", "", "", "", "", "",
		"", "", "", "", "", "'->'", "'..'", "'+='", "'-='", "'*='", "'/='",
		"'%='", "'=='", "'!='", "'<='", "'>='", "'&&'", "'||'", "'{'", "'}'",
		"'('", "')'", "'['", "']'", "','", "':'", "'.'", "'='", "'<'", "'>'",
		"'+'", "'-'", "'*'", "'/'", "'%'", "'!'", "'#'", "'@'", "'~'", "'^'",
	}
	staticData.SymbolicNames = []string{
		"", "LET", "IF", "ELSE", "WHILE", "FOR", "IN", "RETURN", "BREAK", "CONTINUE",
		"TRUE", "FALSE", "COMMAND", "DOC_COMMENT", "BLOCK_COMMENT", "UNTERMINATED_COMMENT",
		"LINE_COMMENT", "NEWLINE", "WS", "RESLOC", "IDENT", "NUMBER", "STRING",
		"ARROW", "DOTDOT", "PLUS_EQ", "MINUS_EQ", "STAR_EQ", "SLASH_EQ", "PCT_EQ",
		"EQEQ", "NEQ", "LE", "GE", "ANDAND", "OROR", "LBRACE", "RBRACE", "LPAREN",
		"RPAREN", "LBRACK", "RBRACK", "COMMA", "COLON", "DOT", "ASSIGN", "LT",
		"GT", "PLUS", "MINUS", "STAR", "SLASH", "PERCENT", "BANG", "HASH", "AT",
		"TILDE", "CARET", "ERROR_CHAR",
	}
	staticData.RuleNames = []string{
		"file", "declaration", "declName", "declBody", "typeAnnotation", "typeRef",
		"params", "param", "block", "blockItem", "statement", "ifTail", "fieldKey",
		"expr", "postfix", "args", "arg", "primary",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 58, 324, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 1, 0, 5, 0, 38, 8, 0, 10, 0, 12, 0, 41, 9,
		0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 50, 8, 1, 3, 1, 52,
		8, 1, 1, 1, 3, 1, 55, 8, 1, 1, 1, 1, 1, 3, 1, 59, 8, 1, 1, 1, 3, 1, 62,
		8, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 3, 3, 69, 8, 3, 1, 3, 3, 3, 72, 8,
		3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 83, 8, 6,
		10, 6, 12, 6, 86, 9, 6, 1, 6, 3, 6, 89, 8, 6, 3, 6, 91, 8, 6, 1, 6, 1,
		6, 1, 7, 1, 7, 3, 7, 97, 8, 7, 1, 8, 1, 8, 5, 8, 101, 8, 8, 10, 8, 12,
		8, 104, 9, 8, 1, 8, 1, 8, 1, 9, 1, 9, 3, 9, 110, 8, 9, 1, 10, 1, 10, 1,
		10, 3, 10, 115, 8, 10, 1, 10, 1, 10, 3, 10, 119, 8, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 129, 8, 10, 3, 10, 131, 8,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 147, 8, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 4, 10, 159, 8, 10, 11, 10,
		12, 10, 160, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 177, 8, 10, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 187, 8, 11, 3, 11, 189,
		8, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 199,
		8, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 227, 8, 13, 5, 13, 229,
		8, 13, 10, 13, 12, 13, 232, 9, 13, 1, 14, 1, 14, 3, 14, 236, 8, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 243, 8, 14, 10, 14, 12, 14, 246,
		9, 14, 1, 14, 3, 14, 249, 8, 14, 3, 14, 251, 8, 14, 1, 14, 1, 14, 3, 14,
		255, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 261, 8, 15, 10, 15, 12,
		15, 264, 9, 15, 1, 15, 3, 15, 267, 8, 15, 3, 15, 269, 8, 15, 1, 15, 1,
		15, 1, 16, 1, 16, 3, 16, 275, 8, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 290, 8,
		17, 1, 17, 3, 17, 293, 8, 17, 4, 17, 295, 8, 17, 11, 17, 12, 17, 296, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 307, 8, 17,
		1, 17, 5, 17, 310, 8, 17, 10, 17, 12, 17, 313, 9, 17, 1, 17, 3, 17, 316,
		8, 17, 3, 17, 318, 8, 17, 1, 17, 1, 17, 3, 17, 322, 8, 17, 1, 17, 0, 1,
		26, 18, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
		34, 0, 10, 1, 0, 19, 20, 1, 0, 25, 29, 2, 0, 19, 20, 22, 22, 2, 0, 48,
		49, 53, 53, 1, 0, 50, 52, 1, 0, 48, 49, 2, 0, 32, 33, 46, 47, 1, 0, 30,
		31, 1, 0, 10, 11, 1, 0, 56, 57, 380, 0, 39, 1, 0, 0, 0, 2, 61, 1, 0, 0,
		0, 4, 63, 1, 0, 0, 0, 6, 71, 1, 0, 0, 0, 8, 73, 1, 0, 0, 0, 10, 76, 1,
		0, 0, 0, 12, 78, 1, 0, 0, 0, 14, 94, 1, 0, 0, 0, 16, 98, 1, 0, 0, 0, 18,
		107, 1, 0, 0, 0, 20, 176, 1, 0, 0, 0, 22, 178, 1, 0, 0, 0, 24, 190, 1,
		0, 0, 0, 26, 198, 1, 0, 0, 0, 28, 254, 1, 0, 0, 0, 30, 256, 1, 0, 0, 0,
		32, 274, 1, 0, 0, 0, 34, 321, 1, 0, 0, 0, 36, 38, 3, 2, 1, 0, 37, 36, 1,
		0, 0, 0, 38, 41, 1, 0, 0, 0, 39, 37, 1, 0, 0, 0, 39, 40, 1, 0, 0, 0, 40,
		42, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0, 42, 43, 5, 0, 0, 1, 43, 1, 1, 0, 0,
		0, 44, 45, 5, 20, 0, 0, 45, 51, 3, 4, 2, 0, 46, 52, 3, 8, 4, 0, 47, 49,
		3, 12, 6, 0, 48, 50, 3, 8, 4, 0, 49, 48, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0,
		50, 52, 1, 0, 0, 0, 51, 46, 1, 0, 0, 0, 51, 47, 1, 0, 0, 0, 51, 52, 1,
		0, 0, 0, 52, 54, 1, 0, 0, 0, 53, 55, 3, 6, 3, 0, 54, 53, 1, 0, 0, 0, 54,
		55, 1, 0, 0, 0, 55, 62, 1, 0, 0, 0, 56, 58, 5, 20, 0, 0, 57, 59, 3, 12,
		6, 0, 58, 57, 1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 62,
		3, 16, 8, 0, 61, 44, 1, 0, 0, 0, 61, 56, 1, 0, 0, 0, 62, 3, 1, 0, 0, 0,
		63, 64, 7, 0, 0, 0, 64, 5, 1, 0, 0, 0, 65, 68, 5, 45, 0, 0, 66, 67, 4,
		3, 0, 0, 67, 69, 3, 26, 13, 0, 68, 66, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0,
		69, 72, 1, 0, 0, 0, 70, 72, 3, 16, 8, 0, 71, 65, 1, 0, 0, 0, 71, 70, 1,
		0, 0, 0, 72, 7, 1, 0, 0, 0, 73, 74, 5, 43, 0, 0, 74, 75, 3, 10, 5, 0, 75,
		9, 1, 0, 0, 0, 76, 77, 7, 0, 0, 0, 77, 11, 1, 0, 0, 0, 78, 90, 5, 38, 0,
		0, 79, 84, 3, 14, 7, 0, 80, 81, 5, 42, 0, 0, 81, 83, 3, 14, 7, 0, 82, 80,
		1, 0, 0, 0, 83, 86, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0,
		85, 88, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0, 87, 89, 5, 42, 0, 0, 88, 87, 1,
		0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 91, 1, 0, 0, 0, 90, 79, 1, 0, 0, 0, 90,
		91, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 93, 5, 39, 0, 0, 93, 13, 1, 0,
		0, 0, 94, 96, 5, 20, 0, 0, 95, 97, 3, 8, 4, 0, 96, 95, 1, 0, 0, 0, 96,
		97, 1, 0, 0, 0, 97, 15, 1, 0, 0, 0, 98, 102, 5, 36, 0, 0, 99, 101, 3, 18,
		9, 0, 100, 99, 1, 0, 0, 0, 101, 104, 1, 0, 0, 0, 102, 100, 1, 0, 0, 0,
		102, 103, 1, 0, 0, 0, 103, 105, 1, 0, 0, 0, 104, 102, 1, 0, 0, 0, 105,
		106, 5, 37, 0, 0, 106, 17, 1, 0, 0, 0, 107, 109, 3, 20, 10, 0, 108, 110,
		5, 42, 0, 0, 109, 108, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 19, 1, 0,
		0, 0, 111, 112, 5, 1, 0, 0, 112, 114, 5, 20, 0, 0, 113, 115, 3, 8, 4, 0,
		114, 113, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 118, 1, 0, 0, 0, 116,
		117, 5, 45, 0, 0, 117, 119, 3, 26, 13, 0, 118, 116, 1, 0, 0, 0, 118, 119,
		1, 0, 0, 0, 119, 177, 1, 0, 0, 0, 120, 121, 5, 2, 0, 0, 121, 122, 5, 38,
		0, 0, 122, 123, 3, 26, 13, 0, 123, 124, 5, 39, 0, 0, 124, 130, 3, 16, 8,
		0, 125, 128, 5, 3, 0, 0, 126, 129, 3, 22, 11, 0, 127, 129, 3, 16, 8, 0,
		128, 126, 1, 0, 0, 0, 128, 127, 1, 0, 0, 0, 129, 131, 1, 0, 0, 0, 130,
		125, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131, 177, 1, 0, 0, 0, 132, 133,
		5, 4, 0, 0, 133, 134, 5, 38, 0, 0, 134, 135, 3, 26, 13, 0, 135, 136, 5,
		39, 0, 0, 136, 137, 3, 16, 8, 0, 137, 177, 1, 0, 0, 0, 138, 139, 5, 5,
		0, 0, 139, 140, 5, 20, 0, 0, 140, 141, 5, 6, 0, 0, 141, 142, 3, 26, 13,
		0, 142, 143, 3, 16, 8, 0, 143, 177, 1, 0, 0, 0, 144, 146, 5, 7, 0, 0, 145,
		147, 3, 26, 13, 0, 146, 145, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 177,
		1, 0, 0, 0, 148, 177, 5, 8, 0, 0, 149, 177, 5, 9, 0, 0, 150, 177, 5, 12,
		0, 0, 151, 152, 3, 24, 12, 0, 152, 153, 5, 45, 0, 0, 153, 154, 3, 26, 13,
		0, 154, 177, 1, 0, 0, 0, 155, 158, 5, 20, 0, 0, 156, 157, 5, 44, 0, 0,
		157, 159, 5, 20, 0, 0, 158, 156, 1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160,
		158, 1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162, 163,
		5, 45, 0, 0, 163, 177, 3, 26, 13, 0, 164, 165, 5, 20, 0, 0, 165, 166, 5,
		40, 0, 0, 166, 167, 3, 32, 16, 0, 167, 168, 5, 41, 0, 0, 168, 169, 5, 45,
		0, 0, 169, 170, 3, 26, 13, 0, 170, 177, 1, 0, 0, 0, 171, 172, 3, 26, 13,
		0, 172, 173, 7, 1, 0, 0, 173, 174, 3, 26, 13, 0, 174, 177, 1, 0, 0, 0,
		175, 177, 3, 26, 13, 0, 176, 111, 1, 0, 0, 0, 176, 120, 1, 0, 0, 0, 176,
		132, 1, 0, 0, 0, 176, 138, 1, 0, 0, 0, 176, 144, 1, 0, 0, 0, 176, 148,
		1, 0, 0, 0, 176, 149, 1, 0, 0, 0, 176, 150, 1, 0, 0, 0, 176, 151, 1, 0,
		0, 0, 176, 155, 1, 0, 0, 0, 176, 164, 1, 0, 0, 0, 176, 171, 1, 0, 0, 0,
		176, 175, 1, 0, 0, 0, 177, 21, 1, 0, 0, 0, 178, 179, 5, 2, 0, 0, 179, 180,
		5, 38, 0, 0, 180, 181, 3, 26, 13, 0, 181, 182, 5, 39, 0, 0, 182, 188, 3,
		16, 8, 0, 183, 186, 5, 3, 0, 0, 184, 187, 3, 22, 11, 0, 185, 187, 3, 16,
		8, 0, 186, 184, 1, 0, 0, 0, 186, 185, 1, 0, 0, 0, 187, 189, 1, 0, 0, 0,
		188, 183, 1, 0, 0, 0, 188, 189, 1, 0, 0, 0, 189, 23, 1, 0, 0, 0, 190, 191,
		7, 2, 0, 0, 191, 25, 1, 0, 0, 0, 192, 193, 6, 13, -1, 0, 193, 194, 7, 3,
		0, 0, 194, 199, 3, 26, 13, 11, 195, 196, 5, 24, 0, 0, 196, 199, 3, 26,
		13, 7, 197, 199, 3, 34, 17, 0, 198, 192, 1, 0, 0, 0, 198, 195, 1, 0, 0,
		0, 198, 197, 1, 0, 0, 0, 199, 230, 1, 0, 0, 0, 200, 201, 10, 10, 0, 0,
		201, 202, 7, 4, 0, 0, 202, 229, 3, 26, 13, 11, 203, 204, 10, 9, 0, 0, 204,
		205, 7, 5, 0, 0, 205, 229, 3, 26, 13, 10, 206, 207, 10, 6, 0, 0, 207, 208,
		7, 6, 0, 0, 208, 229, 3, 26, 13, 7, 209, 210, 10, 5, 0, 0, 210, 211, 7,
		7, 0, 0, 211, 229, 3, 26, 13, 6, 212, 213, 10, 4, 0, 0, 213, 214, 5, 34,
		0, 0, 214, 229, 3, 26, 13, 5, 215, 216, 10, 3, 0, 0, 216, 217, 5, 35, 0,
		0, 217, 229, 3, 26, 13, 4, 218, 219, 10, 2, 0, 0, 219, 220, 5, 23, 0, 0,
		220, 229, 3, 26, 13, 2, 221, 222, 10, 12, 0, 0, 222, 229, 3, 28, 14, 0,
		223, 224, 10, 8, 0, 0, 224, 226, 5, 24, 0, 0, 225, 227, 3, 26, 13, 0, 226,
		225, 1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227, 229, 1, 0, 0, 0, 228, 200,
		1, 0, 0, 0, 228, 203, 1, 0, 0, 0, 228, 206, 1, 0, 0, 0, 228, 209, 1, 0,
		0, 0, 228, 212, 1, 0, 0, 0, 228, 215, 1, 0, 0, 0, 228, 218, 1, 0, 0, 0,
		228, 221, 1, 0, 0, 0, 228, 223, 1, 0, 0, 0, 229, 232, 1, 0, 0, 0, 230,
		228, 1, 0, 0, 0, 230, 231, 1, 0, 0, 0, 231, 27, 1, 0, 0, 0, 232, 230, 1,
		0, 0, 0, 233, 235, 5, 44, 0, 0, 234, 236, 5, 20, 0, 0, 235, 234, 1, 0,
		0, 0, 235, 236, 1, 0, 0, 0, 236, 255, 1, 0, 0, 0, 237, 255, 3, 30, 15,
		0, 238, 250, 5, 40, 0, 0, 239, 244, 3, 32, 16, 0, 240, 241, 5, 42, 0, 0,
		241, 243, 3, 32, 16, 0, 242, 240, 1, 0, 0, 0, 243, 246, 1, 0, 0, 0, 244,
		242, 1, 0, 0, 0, 244, 245, 1, 0, 0, 0, 245, 248, 1, 0, 0, 0, 246, 244,
		1, 0, 0, 0, 247, 249, 5, 42, 0, 0, 248, 247, 1, 0, 0, 0, 248, 249, 1, 0,
		0, 0, 249, 251, 1, 0, 0, 0, 250, 239, 1, 0, 0, 0, 250, 251, 1, 0, 0, 0,
		251, 252, 1, 0, 0, 0, 252, 255, 5, 41, 0, 0, 253, 255, 3, 16, 8, 0, 254,
		233, 1, 0, 0, 0, 254, 237, 1, 0, 0, 0, 254, 238, 1, 0, 0, 0, 254, 253,
		1, 0, 0, 0, 255, 29, 1, 0, 0, 0, 256, 268, 5, 38, 0, 0, 257, 262, 3, 32,
		16, 0, 258, 259, 5, 42, 0, 0, 259, 261, 3, 32, 16, 0, 260, 258, 1, 0, 0,
		0, 261, 264, 1, 0, 0, 0, 262, 260, 1, 0, 0, 0, 262, 263, 1, 0, 0, 0, 263,
		266, 1, 0, 0, 0, 264, 262, 1, 0, 0, 0, 265, 267, 5, 42, 0, 0, 266, 265,
		1, 0, 0, 0, 266, 267, 1, 0, 0, 0, 267, 269, 1, 0, 0, 0, 268, 257, 1, 0,
		0, 0, 268, 269, 1, 0, 0, 0, 269, 270, 1, 0, 0, 0, 270, 271, 5, 39, 0, 0,
		271, 31, 1, 0, 0, 0, 272, 273, 5, 20, 0, 0, 273, 275, 5, 45, 0, 0, 274,
		272, 1, 0, 0, 0, 274, 275, 1, 0, 0, 0, 275, 276, 1, 0, 0, 0, 276, 277,
		3, 26, 13, 0, 277, 33, 1, 0, 0, 0, 278, 322, 5, 21, 0, 0, 279, 322, 5,
		22, 0, 0, 280, 322, 7, 8, 0, 0, 281, 322, 5, 20, 0, 0, 282, 322, 5, 19,
		0, 0, 283, 284, 5, 54, 0, 0, 284, 322, 7, 0, 0, 0, 285, 286, 5, 55, 0,
		0, 286, 322, 5, 20, 0, 0, 287, 292, 7, 9, 0, 0, 288, 290, 5, 49, 0, 0,
		289, 288, 1, 0, 0, 0, 289, 290, 1, 0, 0, 0, 290, 291, 1, 0, 0, 0, 291,
		293, 5, 21, 0, 0, 292, 289, 1, 0, 0, 0, 292, 293, 1, 0, 0, 0, 293, 295,
		1, 0, 0, 0, 294, 287, 1, 0, 0, 0, 295, 296, 1, 0, 0, 0, 296, 294, 1, 0,
		0, 0, 296, 297, 1, 0, 0, 0, 297, 322, 1, 0, 0, 0, 298, 299, 5, 38, 0, 0,
		299, 300, 3, 26, 13, 0, 300, 301, 5, 39, 0, 0, 301, 322, 1, 0, 0, 0, 302,
		317, 5, 40, 0, 0, 303, 311, 3, 26, 13, 0, 304, 307, 5, 42, 0, 0, 305, 307,
		4, 17, 10, 0, 306, 304, 1, 0, 0, 0, 306, 305, 1, 0, 0, 0, 307, 308, 1,
		0, 0, 0, 308, 310, 3, 26, 13, 0, 309, 306, 1, 0, 0, 0, 310, 313, 1, 0,
		0, 0, 311, 309, 1, 0, 0, 0, 311, 312, 1, 0, 0, 0, 312, 315, 1, 0, 0, 0,
		313, 311, 1, 0, 0, 0, 314, 316, 5, 42, 0, 0, 315, 314, 1, 0, 0, 0, 315,
		316, 1, 0, 0, 0, 316, 318, 1, 0, 0, 0, 317, 303, 1, 0, 0, 0, 317, 318,
		1, 0, 0, 0, 318, 319, 1, 0, 0, 0, 319, 322, 5, 41, 0, 0, 320, 322, 3, 16,
		8, 0, 321, 278, 1, 0, 0, 0, 321, 279, 1, 0, 0, 0, 321, 280, 1, 0, 0, 0,
		321, 281, 1, 0, 0, 0, 321, 282, 1, 0, 0, 0, 321, 283, 1, 0, 0, 0, 321,
		285, 1, 0, 0, 0, 321, 294, 1, 0, 0, 0, 321, 298, 1, 0, 0, 0, 321, 302,
		1, 0, 0, 0, 321, 320, 1, 0, 0, 0, 322, 35, 1, 0, 0, 0, 44, 39, 49, 51,
		54, 58, 61, 68, 71, 84, 88, 90, 96, 102, 109, 114, 118, 128, 130, 146,
		160, 176, 186, 188, 198, 226, 228, 230, 235, 244, 248, 250, 254, 262, 266,
		268, 274, 289, 292, 296, 306, 311, 315, 317, 321,
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
	MMSParserLET                  = 1
	MMSParserIF                   = 2
	MMSParserELSE                 = 3
	MMSParserWHILE                = 4
	MMSParserFOR                  = 5
	MMSParserIN                   = 6
	MMSParserRETURN               = 7
	MMSParserBREAK                = 8
	MMSParserCONTINUE             = 9
	MMSParserTRUE                 = 10
	MMSParserFALSE                = 11
	MMSParserCOMMAND              = 12
	MMSParserDOC_COMMENT          = 13
	MMSParserBLOCK_COMMENT        = 14
	MMSParserUNTERMINATED_COMMENT = 15
	MMSParserLINE_COMMENT         = 16
	MMSParserNEWLINE              = 17
	MMSParserWS                   = 18
	MMSParserRESLOC               = 19
	MMSParserIDENT                = 20
	MMSParserNUMBER               = 21
	MMSParserSTRING               = 22
	MMSParserARROW                = 23
	MMSParserDOTDOT               = 24
	MMSParserPLUS_EQ              = 25
	MMSParserMINUS_EQ             = 26
	MMSParserSTAR_EQ              = 27
	MMSParserSLASH_EQ             = 28
	MMSParserPCT_EQ               = 29
	MMSParserEQEQ                 = 30
	MMSParserNEQ                  = 31
	MMSParserLE                   = 32
	MMSParserGE                   = 33
	MMSParserANDAND               = 34
	MMSParserOROR                 = 35
	MMSParserLBRACE               = 36
	MMSParserRBRACE               = 37
	MMSParserLPAREN               = 38
	MMSParserRPAREN               = 39
	MMSParserLBRACK               = 40
	MMSParserRBRACK               = 41
	MMSParserCOMMA                = 42
	MMSParserCOLON                = 43
	MMSParserDOT                  = 44
	MMSParserASSIGN               = 45
	MMSParserLT                   = 46
	MMSParserGT                   = 47
	MMSParserPLUS                 = 48
	MMSParserMINUS                = 49
	MMSParserSTAR                 = 50
	MMSParserSLASH                = 51
	MMSParserPERCENT              = 52
	MMSParserBANG                 = 53
	MMSParserHASH                 = 54
	MMSParserAT                   = 55
	MMSParserTILDE                = 56
	MMSParserCARET                = 57
	MMSParserERROR_CHAR           = 58
)

// MMSParser rules.
const (
	MMSParserRULE_file           = 0
	MMSParserRULE_declaration    = 1
	MMSParserRULE_declName       = 2
	MMSParserRULE_declBody       = 3
	MMSParserRULE_typeAnnotation = 4
	MMSParserRULE_typeRef        = 5
	MMSParserRULE_params         = 6
	MMSParserRULE_param          = 7
	MMSParserRULE_block          = 8
	MMSParserRULE_blockItem      = 9
	MMSParserRULE_statement      = 10
	MMSParserRULE_ifTail         = 11
	MMSParserRULE_fieldKey       = 12
	MMSParserRULE_expr           = 13
	MMSParserRULE_postfix        = 14
	MMSParserRULE_args           = 15
	MMSParserRULE_arg            = 16
	MMSParserRULE_primary        = 17
)

// IFileContext is an interface to support dynamic dispatch.
type IFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllDeclaration() []IDeclarationContext
	Declaration(i int) IDeclarationContext

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

func (s *FileContext) AllDeclaration() []IDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclarationContext); ok {
			tst[i] = t.(IDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *FileContext) Declaration(i int) IDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
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

	return t.(IDeclarationContext)
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
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MMSParserIDENT {
		{
			p.SetState(36)
			p.Declaration()
		}

		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(42)
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
	IDENT() antlr.TerminalNode
	DeclName() IDeclNameContext
	TypeAnnotation() ITypeAnnotationContext
	Params() IParamsContext
	DeclBody() IDeclBodyContext
	Block() IBlockContext

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

func (s *DeclarationContext) IDENT() antlr.TerminalNode {
	return s.GetToken(MMSParserIDENT, 0)
}

func (s *DeclarationContext) DeclName() IDeclNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclNameContext)
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

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MMSParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MMSParserRULE_declaration)
	var _la int

	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(44)

			var _m = p.Match(MMSParserIDENT)

			localctx.(*DeclarationContext).kind = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(45)

			var _x = p.DeclName()

			localctx.(*DeclarationContext).name = _x
		}
		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		switch p.GetTokenStream().LA(1) {
		case MMSParserCOLON:
			{
				p.SetState(46)
				p.TypeAnnotation()
			}

		case MMSParserLPAREN:
			{
				p.SetState(47)
				p.Params()
			}
			p.SetState(49)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == MMSParserCOLON {
				{
					p.SetState(48)
					p.TypeAnnotation()
				}

			}

		case MMSParserEOF, MMSParserIDENT, MMSParserLBRACE, MMSParserASSIGN:

		default:
		}
		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MMSParserLBRACE || _la == MMSParserASSIGN {
			{
				p.SetState(53)
				p.DeclBody()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(56)

			var _m = p.Match(MMSParserIDENT)

			localctx.(*DeclarationContext).kind = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MMSParserLPAREN {
			{
				p.SetState(57)
				p.Params()
			}

		}
		{
			p.SetState(60)
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

// IDeclNameContext is an interface to support dynamic dispatch.
type IDeclNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	RESLOC() antlr.TerminalNode

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

func (s *DeclNameContext) IDENT() antlr.TerminalNode {
	return s.GetToken(MMSParserIDENT, 0)
}

func (s *DeclNameContext) RESLOC() antlr.TerminalNode {
	return s.GetToken(MMSParserRESLOC, 0)
}

func (s *DeclNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MMSParser) DeclName() (localctx IDeclNameContext) {
	localctx = NewDeclNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MMSParserRULE_declName)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(63)
		_la = p.GetTokenStream().LA(1)

		if !(_la == MMSParserRESLOC || _la == MMSParserIDENT) {
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
	Expr() IExprContext
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

func (s *DeclBodyContext) Expr() IExprContext {
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
	p.EnterRule(localctx, 6, MMSParserRULE_declBody)
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MMSParserASSIGN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(65)
			p.Match(MMSParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(68)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
			p.SetState(66)

			if !(p.valueFollows()) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.valueFollows()", ""))
				goto errorExit
			}
			{
				p.SetState(67)
				p.expr(0)
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case MMSParserLBRACE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(70)
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

// ITypeAnnotationContext is an interface to support dynamic dispatch.
type ITypeAnnotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COLON() antlr.TerminalNode
	TypeRef() ITypeRefContext

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

func (s *TypeAnnotationContext) TypeRef() ITypeRefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeRefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeRefContext)
}

func (s *TypeAnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeAnnotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MMSParser) TypeAnnotation() (localctx ITypeAnnotationContext) {
	localctx = NewTypeAnnotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MMSParserRULE_typeAnnotation)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(73)
		p.Match(MMSParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(74)
		p.TypeRef()
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

// ITypeRefContext is an interface to support dynamic dispatch.
type ITypeRefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	RESLOC() antlr.TerminalNode

	// IsTypeRefContext differentiates from other interfaces.
	IsTypeRefContext()
}

type TypeRefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeRefContext() *TypeRefContext {
	var p = new(TypeRefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_typeRef
	return p
}

func InitEmptyTypeRefContext(p *TypeRefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_typeRef
}

func (*TypeRefContext) IsTypeRefContext() {}

func NewTypeRefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeRefContext {
	var p = new(TypeRefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_typeRef

	return p
}

func (s *TypeRefContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeRefContext) IDENT() antlr.TerminalNode {
	return s.GetToken(MMSParserIDENT, 0)
}

func (s *TypeRefContext) RESLOC() antlr.TerminalNode {
	return s.GetToken(MMSParserRESLOC, 0)
}

func (s *TypeRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeRefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MMSParser) TypeRef() (localctx ITypeRefContext) {
	localctx = NewTypeRefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MMSParserRULE_typeRef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(76)
		_la = p.GetTokenStream().LA(1)

		if !(_la == MMSParserRESLOC || _la == MMSParserIDENT) {
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
	p.EnterRule(localctx, 12, MMSParserRULE_params)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Match(MMSParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MMSParserIDENT {
		{
			p.SetState(79)
			p.Param()
		}
		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(80)
					p.Match(MMSParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(81)
					p.Param()
				}

			}
			p.SetState(86)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MMSParserCOMMA {
			{
				p.SetState(87)
				p.Match(MMSParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(92)
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

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MMSParser) Param() (localctx IParamContext) {
	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MMSParserRULE_param)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(94)
		p.Match(MMSParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MMSParserCOLON {
		{
			p.SetState(95)
			p.TypeAnnotation()
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
	p.EnterRule(localctx, 16, MMSParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(98)
		p.Match(MMSParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&280069044960763830) != 0 {
		{
			p.SetState(99)
			p.BlockItem()
		}

		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(105)
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

func (s *BlockItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MMSParser) BlockItem() (localctx IBlockItemContext) {
	localctx = NewBlockItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MMSParserRULE_blockItem)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(107)
		p.Statement()
	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MMSParserCOMMA {
		{
			p.SetState(108)
			p.Match(MMSParserCOMMA)
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

func (s *ForStmtContext) IDENT() antlr.TerminalNode {
	return s.GetToken(MMSParserIDENT, 0)
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

func (s *LetStmtContext) LET() antlr.TerminalNode {
	return s.GetToken(MMSParserLET, 0)
}

func (s *LetStmtContext) IDENT() antlr.TerminalNode {
	return s.GetToken(MMSParserIDENT, 0)
}

func (s *LetStmtContext) TypeAnnotation() ITypeAnnotationContext {
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

func (s *LetStmtContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(MMSParserASSIGN, 0)
}

func (s *LetStmtContext) Expr() IExprContext {
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

func (s *ReturnStmtContext) Expr() IExprContext {
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

func (s *FieldStmtContext) Expr() IExprContext {
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

func (s *IfStmtContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfStmtContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
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

	return t.(IBlockContext)
}

func (s *IfStmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(MMSParserELSE, 0)
}

func (s *IfStmtContext) IfTail() IIfTailContext {
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

func (s *PathAssignStmtContext) Expr() IExprContext {
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

func (s *IndexAssignStmtContext) Expr() IExprContext {
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

func (p *MMSParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MMSParserRULE_statement)
	var _la int

	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		localctx = NewLetStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(111)
			p.Match(MMSParserLET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(112)
			p.Match(MMSParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MMSParserCOLON {
			{
				p.SetState(113)
				p.TypeAnnotation()
			}

		}
		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MMSParserASSIGN {
			{
				p.SetState(116)
				p.Match(MMSParserASSIGN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(117)
				p.expr(0)
			}

		}

	case 2:
		localctx = NewIfStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(120)
			p.Match(MMSParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(121)
			p.Match(MMSParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(122)
			p.expr(0)
		}
		{
			p.SetState(123)
			p.Match(MMSParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(124)
			p.Block()
		}
		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MMSParserELSE {
			{
				p.SetState(125)
				p.Match(MMSParserELSE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(128)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case MMSParserIF:
				{
					p.SetState(126)
					p.IfTail()
				}

			case MMSParserLBRACE:
				{
					p.SetState(127)
					p.Block()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

		}

	case 3:
		localctx = NewWhileStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(132)
			p.Match(MMSParserWHILE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(133)
			p.Match(MMSParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(134)
			p.expr(0)
		}
		{
			p.SetState(135)
			p.Match(MMSParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(136)
			p.Block()
		}

	case 4:
		localctx = NewForStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(138)
			p.Match(MMSParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(139)
			p.Match(MMSParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(140)
			p.Match(MMSParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(141)
			p.expr(0)
		}
		{
			p.SetState(142)
			p.Block()
		}

	case 5:
		localctx = NewReturnStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(144)
			p.Match(MMSParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(146)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(145)
				p.expr(0)
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 6:
		localctx = NewBreakStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(148)
			p.Match(MMSParserBREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewContinueStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(149)
			p.Match(MMSParserCONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewCommandStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(150)
			p.Match(MMSParserCOMMAND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		localctx = NewFieldStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(151)
			p.FieldKey()
		}
		{
			p.SetState(152)
			p.Match(MMSParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(153)
			p.expr(0)
		}

	case 10:
		localctx = NewPathAssignStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(155)
			p.Match(MMSParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(158)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == MMSParserDOT {
			{
				p.SetState(156)
				p.Match(MMSParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(157)
				p.Match(MMSParserIDENT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(160)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(162)
			p.Match(MMSParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(163)
			p.expr(0)
		}

	case 11:
		localctx = NewIndexAssignStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(164)
			p.Match(MMSParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(165)
			p.Match(MMSParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(166)
			p.Arg()
		}
		{
			p.SetState(167)
			p.Match(MMSParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(168)
			p.Match(MMSParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(169)
			p.expr(0)
		}

	case 12:
		localctx = NewCompoundAssignStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(171)
			p.expr(0)
		}
		{
			p.SetState(172)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*CompoundAssignStmtContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1040187392) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*CompoundAssignStmtContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(173)
			p.expr(0)
		}

	case 13:
		localctx = NewExprStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(175)
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
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	ELSE() antlr.TerminalNode
	IfTail() IIfTailContext

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

func (s *IfTailContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfTailContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
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

	return t.(IBlockContext)
}

func (s *IfTailContext) ELSE() antlr.TerminalNode {
	return s.GetToken(MMSParserELSE, 0)
}

func (s *IfTailContext) IfTail() IIfTailContext {
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

func (s *IfTailContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfTailContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *MMSParser) IfTail() (localctx IIfTailContext) {
	localctx = NewIfTailContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, MMSParserRULE_ifTail)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(178)
		p.Match(MMSParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(179)
		p.Match(MMSParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(180)
		p.expr(0)
	}
	{
		p.SetState(181)
		p.Match(MMSParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(182)
		p.Block()
	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MMSParserELSE {
		{
			p.SetState(183)
			p.Match(MMSParserELSE)
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

		switch p.GetTokenStream().LA(1) {
		case MMSParserIF:
			{
				p.SetState(184)
				p.IfTail()
			}

		case MMSParserLBRACE:
			{
				p.SetState(185)
				p.Block()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

// IFieldKeyContext is an interface to support dynamic dispatch.
type IFieldKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
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

func (s *FieldKeyContext) IDENT() antlr.TerminalNode {
	return s.GetToken(MMSParserIDENT, 0)
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
	p.EnterRule(localctx, 24, MMSParserRULE_fieldKey)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(190)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&5767168) != 0) {
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
	_startState := 26
	p.EnterRecursionRule(localctx, 26, MMSParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(198)
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
			p.SetState(193)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*UnaryExprContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9851624184872960) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*UnaryExprContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(194)
			p.expr(11)
		}

	case MMSParserDOTDOT:
		localctx = NewRangeToExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(195)
			p.Match(MMSParserDOTDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(196)
			p.expr(7)
		}

	case MMSParserTRUE, MMSParserFALSE, MMSParserRESLOC, MMSParserIDENT, MMSParserNUMBER, MMSParserSTRING, MMSParserLBRACE, MMSParserLPAREN, MMSParserLBRACK, MMSParserHASH, MMSParserAT, MMSParserTILDE, MMSParserCARET:
		localctx = NewPrimaryExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(197)
			p.Primary()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(228)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MMSParserRULE_expr)
				p.SetState(200)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(201)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7881299347898368) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(202)
					p.expr(11)
				}

			case 2:
				localctx = NewAddExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MMSParserRULE_expr)
				p.SetState(203)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(204)

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
					p.SetState(205)
					p.expr(10)
				}

			case 3:
				localctx = NewCompareExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MMSParserRULE_expr)
				p.SetState(206)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(207)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*CompareExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&211119117434880) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*CompareExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(208)
					p.expr(7)
				}

			case 4:
				localctx = NewEqualExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MMSParserRULE_expr)
				p.SetState(209)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(210)

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
					p.SetState(211)
					p.expr(6)
				}

			case 5:
				localctx = NewAndExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MMSParserRULE_expr)
				p.SetState(212)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(213)
					p.Match(MMSParserANDAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(214)
					p.expr(5)
				}

			case 6:
				localctx = NewOrExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MMSParserRULE_expr)
				p.SetState(215)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(216)
					p.Match(MMSParserOROR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(217)
					p.expr(4)
				}

			case 7:
				localctx = NewArrowExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MMSParserRULE_expr)
				p.SetState(218)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(219)
					p.Match(MMSParserARROW)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(220)
					p.expr(2)
				}

			case 8:
				localctx = NewPostfixExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MMSParserRULE_expr)
				p.SetState(221)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(222)
					p.Postfix()
				}

			case 9:
				localctx = NewRangeExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MMSParserRULE_expr)
				p.SetState(223)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(224)
					p.Match(MMSParserDOTDOT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(226)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(225)
						p.expr(0)
					}

				} else if p.HasError() { // JIM
					goto errorExit
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(232)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
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
	name antlr.Token
}

func NewMemberPostfixContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MemberPostfixContext {
	var p = new(MemberPostfixContext)

	InitEmptyPostfixContext(&p.PostfixContext)
	p.parser = parser
	p.CopyAll(ctx.(*PostfixContext))

	return p
}

func (s *MemberPostfixContext) GetName() antlr.Token { return s.name }

func (s *MemberPostfixContext) SetName(v antlr.Token) { s.name = v }

func (s *MemberPostfixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberPostfixContext) DOT() antlr.TerminalNode {
	return s.GetToken(MMSParserDOT, 0)
}

func (s *MemberPostfixContext) IDENT() antlr.TerminalNode {
	return s.GetToken(MMSParserIDENT, 0)
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
	p.EnterRule(localctx, 28, MMSParserRULE_postfix)
	var _la int

	var _alt int

	p.SetState(254)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MMSParserDOT:
		localctx = NewMemberPostfixContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(233)
			p.Match(MMSParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(235)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(234)

				var _m = p.Match(MMSParserIDENT)

				localctx.(*MemberPostfixContext).name = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case MMSParserLPAREN:
		localctx = NewCallPostfixContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(237)
			p.Args()
		}

	case MMSParserLBRACK:
		localctx = NewIndexPostfixContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(238)
			p.Match(MMSParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(250)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&280069044960758784) != 0 {
			{
				p.SetState(239)
				p.Arg()
			}
			p.SetState(244)
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
					{
						p.SetState(240)
						p.Match(MMSParserCOMMA)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}
					{
						p.SetState(241)
						p.Arg()
					}

				}
				p.SetState(246)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext())
				if p.HasError() {
					goto errorExit
				}
			}
			p.SetState(248)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == MMSParserCOMMA {
				{
					p.SetState(247)
					p.Match(MMSParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}

		}
		{
			p.SetState(252)
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
			p.SetState(253)
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
	p.EnterRule(localctx, 30, MMSParserRULE_args)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(256)
		p.Match(MMSParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(268)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&280069044960758784) != 0 {
		{
			p.SetState(257)
			p.Arg()
		}
		p.SetState(262)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(258)
					p.Match(MMSParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(259)
					p.Arg()
				}

			}
			p.SetState(264)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(266)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MMSParserCOMMA {
			{
				p.SetState(265)
				p.Match(MMSParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(270)
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
	Expr() IExprContext
	IDENT() antlr.TerminalNode
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

func (s *ArgContext) Expr() IExprContext {
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

func (s *ArgContext) IDENT() antlr.TerminalNode {
	return s.GetToken(MMSParserIDENT, 0)
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
	p.EnterRule(localctx, 32, MMSParserRULE_arg)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(274)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(272)
			p.Match(MMSParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(273)
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
		p.SetState(276)
		p.expr(0)
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

func (s *ListLitContext) AllExpr() []IExprContext {
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

func (s *ListLitContext) Expr(i int) IExprContext {
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

func (s *ParenExprContext) Expr() IExprContext {
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

func (p *MMSParser) Primary() (localctx IPrimaryContext) {
	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, MMSParserRULE_primary)
	var _la int

	var _alt int

	p.SetState(321)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MMSParserNUMBER:
		localctx = NewNumberLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(278)
			p.Match(MMSParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MMSParserSTRING:
		localctx = NewStringLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(279)
			p.Match(MMSParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MMSParserTRUE, MMSParserFALSE:
		localctx = NewBoolLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(280)
			_la = p.GetTokenStream().LA(1)

			if !(_la == MMSParserTRUE || _la == MMSParserFALSE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case MMSParserIDENT:
		localctx = NewIdentRefContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(281)
			p.Match(MMSParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MMSParserRESLOC:
		localctx = NewReslocRefContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(282)
			p.Match(MMSParserRESLOC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MMSParserHASH:
		localctx = NewTagRefContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(283)
			p.Match(MMSParserHASH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(284)
			_la = p.GetTokenStream().LA(1)

			if !(_la == MMSParserRESLOC || _la == MMSParserIDENT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case MMSParserAT:
		localctx = NewSelectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(285)
			p.Match(MMSParserAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(286)
			p.Match(MMSParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MMSParserTILDE, MMSParserCARET:
		localctx = NewCoordinateContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		p.SetState(294)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(287)
					_la = p.GetTokenStream().LA(1)

					if !(_la == MMSParserTILDE || _la == MMSParserCARET) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				p.SetState(292)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext()) == 1 {
					p.SetState(289)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_la = p.GetTokenStream().LA(1)

					if _la == MMSParserMINUS {
						{
							p.SetState(288)
							p.Match(MMSParserMINUS)
							if p.HasError() {
								// Recognition error - abort rule
								goto errorExit
							}
						}

					}
					{
						p.SetState(291)
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

			p.SetState(296)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case MMSParserLPAREN:
		localctx = NewParenExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(298)
			p.Match(MMSParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(299)
			p.expr(0)
		}
		{
			p.SetState(300)
			p.Match(MMSParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MMSParserLBRACK:
		localctx = NewListLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(302)
			p.Match(MMSParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(317)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&280069044960758784) != 0 {
			{
				p.SetState(303)
				p.expr(0)
			}
			p.SetState(311)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
			for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				if _alt == 1 {
					p.SetState(306)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}

					switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext()) {
					case 1:
						{
							p.SetState(304)
							p.Match(MMSParserCOMMA)
							if p.HasError() {
								// Recognition error - abort rule
								goto errorExit
							}
						}

					case 2:
						p.SetState(305)

						if !(p.onNewLine()) {
							p.SetError(antlr.NewFailedPredicateException(p, "p.onNewLine()", ""))
							goto errorExit
						}

					case antlr.ATNInvalidAltNumber:
						goto errorExit
					}
					{
						p.SetState(308)
						p.expr(0)
					}

				}
				p.SetState(313)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext())
				if p.HasError() {
					goto errorExit
				}
			}
			p.SetState(315)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == MMSParserCOMMA {
				{
					p.SetState(314)
					p.Match(MMSParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}

		}
		{
			p.SetState(319)
			p.Match(MMSParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MMSParserLBRACE:
		localctx = NewObjectLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(320)
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

func (p *MMSParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 3:
		var t *DeclBodyContext = nil
		if localctx != nil {
			t = localctx.(*DeclBodyContext)
		}
		return p.DeclBody_Sempred(t, predIndex)

	case 13:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	case 17:
		var t *PrimaryContext = nil
		if localctx != nil {
			t = localctx.(*PrimaryContext)
		}
		return p.Primary_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *MMSParser) DeclBody_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.valueFollows()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *MMSParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 8)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *MMSParser) Primary_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 10:
		return p.onNewLine()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
