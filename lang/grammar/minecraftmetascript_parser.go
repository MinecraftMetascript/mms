// Code generated from ./grammar/MinecraftMetascript.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // MinecraftMetascript
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

type MinecraftMetascriptParser struct {
	*antlr.BaseParser
}

var MinecraftMetascriptParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func minecraftmetascriptParserInit() {
	staticData := &MinecraftMetascriptParserStaticData
	staticData.LiteralNames = []string{
		"", "'{'", "'}'", "'='", "':'", "'('", "','", "')'", "'.'", "'!'", "'&&'",
		"'||'", "'If'", "'['", "']'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "Int", "Float",
		"String", "WS", "NL", "Identifier", "BlockComment", "LineComment",
	}
	staticData.RuleNames = []string{
		"file", "namedBlock", "block", "varDecl", "resourceReference", "fn",
		"value", "condition", "rootCondition", "conditional", "conditionalBody",
		"list", "number",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 22, 276, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 5, 0, 28, 8, 0, 10, 0, 12, 0, 31,
		9, 0, 1, 0, 1, 0, 5, 0, 35, 8, 0, 10, 0, 12, 0, 38, 9, 0, 5, 0, 40, 8,
		0, 10, 0, 12, 0, 43, 9, 0, 1, 1, 1, 1, 1, 1, 5, 1, 48, 8, 1, 10, 1, 12,
		1, 51, 9, 1, 1, 1, 1, 1, 1, 1, 3, 1, 56, 8, 1, 1, 1, 5, 1, 59, 8, 1, 10,
		1, 12, 1, 62, 9, 1, 5, 1, 64, 8, 1, 10, 1, 12, 1, 67, 9, 1, 1, 1, 1, 1,
		1, 2, 1, 2, 5, 2, 73, 8, 2, 10, 2, 12, 2, 76, 9, 2, 1, 2, 1, 2, 5, 2, 80,
		8, 2, 10, 2, 12, 2, 83, 9, 2, 1, 2, 1, 2, 5, 2, 87, 8, 2, 10, 2, 12, 2,
		90, 9, 2, 5, 2, 92, 8, 2, 10, 2, 12, 2, 95, 9, 2, 1, 2, 1, 2, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 4, 1, 4, 3, 4, 105, 8, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 5, 5, 114, 8, 5, 10, 5, 12, 5, 117, 9, 5, 1, 5, 1, 5, 3,
		5, 121, 8, 5, 3, 5, 123, 8, 5, 1, 5, 1, 5, 1, 5, 5, 5, 128, 8, 5, 10, 5,
		12, 5, 131, 9, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 139, 8, 6,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 147, 8, 7, 10, 7, 12, 7, 150,
		9, 7, 1, 7, 1, 7, 5, 7, 154, 8, 7, 10, 7, 12, 7, 157, 9, 7, 1, 7, 1, 7,
		3, 7, 161, 8, 7, 1, 7, 1, 7, 5, 7, 165, 8, 7, 10, 7, 12, 7, 168, 9, 7,
		1, 7, 1, 7, 5, 7, 172, 8, 7, 10, 7, 12, 7, 175, 9, 7, 1, 7, 1, 7, 1, 7,
		5, 7, 180, 8, 7, 10, 7, 12, 7, 183, 9, 7, 1, 7, 1, 7, 5, 7, 187, 8, 7,
		10, 7, 12, 7, 190, 9, 7, 1, 7, 5, 7, 193, 8, 7, 10, 7, 12, 7, 196, 9, 7,
		1, 8, 1, 8, 1, 9, 1, 9, 5, 9, 202, 8, 9, 10, 9, 12, 9, 205, 9, 9, 1, 9,
		1, 9, 5, 9, 209, 8, 9, 10, 9, 12, 9, 212, 9, 9, 1, 9, 3, 9, 215, 8, 9,
		1, 10, 1, 10, 5, 10, 219, 8, 10, 10, 10, 12, 10, 222, 9, 10, 1, 10, 3,
		10, 225, 8, 10, 1, 10, 5, 10, 228, 8, 10, 10, 10, 12, 10, 231, 9, 10, 1,
		10, 1, 10, 1, 11, 1, 11, 5, 11, 237, 8, 11, 10, 11, 12, 11, 240, 9, 11,
		1, 11, 1, 11, 5, 11, 244, 8, 11, 10, 11, 12, 11, 247, 9, 11, 1, 11, 3,
		11, 250, 8, 11, 1, 11, 5, 11, 253, 8, 11, 10, 11, 12, 11, 256, 9, 11, 5,
		11, 258, 8, 11, 10, 11, 12, 11, 261, 9, 11, 1, 11, 3, 11, 264, 8, 11, 1,
		11, 5, 11, 267, 8, 11, 10, 11, 12, 11, 270, 9, 11, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 12, 1, 129, 1, 14, 13, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
		22, 24, 0, 1, 1, 0, 15, 16, 306, 0, 29, 1, 0, 0, 0, 2, 44, 1, 0, 0, 0,
		4, 70, 1, 0, 0, 0, 6, 98, 1, 0, 0, 0, 8, 104, 1, 0, 0, 0, 10, 108, 1, 0,
		0, 0, 12, 138, 1, 0, 0, 0, 14, 160, 1, 0, 0, 0, 16, 197, 1, 0, 0, 0, 18,
		199, 1, 0, 0, 0, 20, 216, 1, 0, 0, 0, 22, 234, 1, 0, 0, 0, 24, 273, 1,
		0, 0, 0, 26, 28, 5, 19, 0, 0, 27, 26, 1, 0, 0, 0, 28, 31, 1, 0, 0, 0, 29,
		27, 1, 0, 0, 0, 29, 30, 1, 0, 0, 0, 30, 41, 1, 0, 0, 0, 31, 29, 1, 0, 0,
		0, 32, 36, 3, 2, 1, 0, 33, 35, 5, 19, 0, 0, 34, 33, 1, 0, 0, 0, 35, 38,
		1, 0, 0, 0, 36, 34, 1, 0, 0, 0, 36, 37, 1, 0, 0, 0, 37, 40, 1, 0, 0, 0,
		38, 36, 1, 0, 0, 0, 39, 32, 1, 0, 0, 0, 40, 43, 1, 0, 0, 0, 41, 39, 1,
		0, 0, 0, 41, 42, 1, 0, 0, 0, 42, 1, 1, 0, 0, 0, 43, 41, 1, 0, 0, 0, 44,
		45, 5, 20, 0, 0, 45, 49, 5, 20, 0, 0, 46, 48, 5, 19, 0, 0, 47, 46, 1, 0,
		0, 0, 48, 51, 1, 0, 0, 0, 49, 47, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 52,
		1, 0, 0, 0, 51, 49, 1, 0, 0, 0, 52, 65, 5, 1, 0, 0, 53, 56, 3, 6, 3, 0,
		54, 56, 3, 4, 2, 0, 55, 53, 1, 0, 0, 0, 55, 54, 1, 0, 0, 0, 56, 60, 1,
		0, 0, 0, 57, 59, 5, 19, 0, 0, 58, 57, 1, 0, 0, 0, 59, 62, 1, 0, 0, 0, 60,
		58, 1, 0, 0, 0, 60, 61, 1, 0, 0, 0, 61, 64, 1, 0, 0, 0, 62, 60, 1, 0, 0,
		0, 63, 55, 1, 0, 0, 0, 64, 67, 1, 0, 0, 0, 65, 63, 1, 0, 0, 0, 65, 66,
		1, 0, 0, 0, 66, 68, 1, 0, 0, 0, 67, 65, 1, 0, 0, 0, 68, 69, 5, 2, 0, 0,
		69, 3, 1, 0, 0, 0, 70, 74, 5, 20, 0, 0, 71, 73, 5, 19, 0, 0, 72, 71, 1,
		0, 0, 0, 73, 76, 1, 0, 0, 0, 74, 72, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75,
		77, 1, 0, 0, 0, 76, 74, 1, 0, 0, 0, 77, 81, 5, 1, 0, 0, 78, 80, 5, 19,
		0, 0, 79, 78, 1, 0, 0, 0, 80, 83, 1, 0, 0, 0, 81, 79, 1, 0, 0, 0, 81, 82,
		1, 0, 0, 0, 82, 93, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 84, 88, 3, 6, 3, 0,
		85, 87, 5, 19, 0, 0, 86, 85, 1, 0, 0, 0, 87, 90, 1, 0, 0, 0, 88, 86, 1,
		0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 92, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 91,
		84, 1, 0, 0, 0, 92, 95, 1, 0, 0, 0, 93, 91, 1, 0, 0, 0, 93, 94, 1, 0, 0,
		0, 94, 96, 1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 96, 97, 5, 2, 0, 0, 97, 5, 1,
		0, 0, 0, 98, 99, 5, 20, 0, 0, 99, 100, 5, 3, 0, 0, 100, 101, 3, 12, 6,
		0, 101, 7, 1, 0, 0, 0, 102, 103, 5, 20, 0, 0, 103, 105, 5, 4, 0, 0, 104,
		102, 1, 0, 0, 0, 104, 105, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 107,
		5, 20, 0, 0, 107, 9, 1, 0, 0, 0, 108, 109, 5, 20, 0, 0, 109, 115, 5, 5,
		0, 0, 110, 111, 3, 12, 6, 0, 111, 112, 5, 6, 0, 0, 112, 114, 1, 0, 0, 0,
		113, 110, 1, 0, 0, 0, 114, 117, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 115,
		116, 1, 0, 0, 0, 116, 122, 1, 0, 0, 0, 117, 115, 1, 0, 0, 0, 118, 120,
		3, 12, 6, 0, 119, 121, 5, 6, 0, 0, 120, 119, 1, 0, 0, 0, 120, 121, 1, 0,
		0, 0, 121, 123, 1, 0, 0, 0, 122, 118, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0,
		123, 124, 1, 0, 0, 0, 124, 129, 5, 7, 0, 0, 125, 126, 5, 8, 0, 0, 126,
		128, 3, 10, 5, 0, 127, 125, 1, 0, 0, 0, 128, 131, 1, 0, 0, 0, 129, 130,
		1, 0, 0, 0, 129, 127, 1, 0, 0, 0, 130, 11, 1, 0, 0, 0, 131, 129, 1, 0,
		0, 0, 132, 139, 3, 24, 12, 0, 133, 139, 5, 17, 0, 0, 134, 139, 3, 10, 5,
		0, 135, 139, 3, 8, 4, 0, 136, 139, 3, 18, 9, 0, 137, 139, 3, 22, 11, 0,
		138, 132, 1, 0, 0, 0, 138, 133, 1, 0, 0, 0, 138, 134, 1, 0, 0, 0, 138,
		135, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 138, 137, 1, 0, 0, 0, 139, 13, 1,
		0, 0, 0, 140, 141, 6, 7, -1, 0, 141, 142, 5, 9, 0, 0, 142, 161, 3, 14,
		7, 5, 143, 161, 3, 16, 8, 0, 144, 148, 5, 5, 0, 0, 145, 147, 5, 19, 0,
		0, 146, 145, 1, 0, 0, 0, 147, 150, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 148,
		149, 1, 0, 0, 0, 149, 151, 1, 0, 0, 0, 150, 148, 1, 0, 0, 0, 151, 155,
		3, 14, 7, 0, 152, 154, 5, 19, 0, 0, 153, 152, 1, 0, 0, 0, 154, 157, 1,
		0, 0, 0, 155, 153, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156, 158, 1, 0, 0,
		0, 157, 155, 1, 0, 0, 0, 158, 159, 5, 7, 0, 0, 159, 161, 1, 0, 0, 0, 160,
		140, 1, 0, 0, 0, 160, 143, 1, 0, 0, 0, 160, 144, 1, 0, 0, 0, 161, 194,
		1, 0, 0, 0, 162, 166, 10, 2, 0, 0, 163, 165, 5, 19, 0, 0, 164, 163, 1,
		0, 0, 0, 165, 168, 1, 0, 0, 0, 166, 164, 1, 0, 0, 0, 166, 167, 1, 0, 0,
		0, 167, 169, 1, 0, 0, 0, 168, 166, 1, 0, 0, 0, 169, 173, 5, 10, 0, 0, 170,
		172, 5, 19, 0, 0, 171, 170, 1, 0, 0, 0, 172, 175, 1, 0, 0, 0, 173, 171,
		1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 176, 1, 0, 0, 0, 175, 173, 1, 0,
		0, 0, 176, 193, 3, 14, 7, 3, 177, 181, 10, 1, 0, 0, 178, 180, 5, 19, 0,
		0, 179, 178, 1, 0, 0, 0, 180, 183, 1, 0, 0, 0, 181, 179, 1, 0, 0, 0, 181,
		182, 1, 0, 0, 0, 182, 184, 1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 184, 188,
		5, 11, 0, 0, 185, 187, 5, 19, 0, 0, 186, 185, 1, 0, 0, 0, 187, 190, 1,
		0, 0, 0, 188, 186, 1, 0, 0, 0, 188, 189, 1, 0, 0, 0, 189, 191, 1, 0, 0,
		0, 190, 188, 1, 0, 0, 0, 191, 193, 3, 14, 7, 2, 192, 162, 1, 0, 0, 0, 192,
		177, 1, 0, 0, 0, 193, 196, 1, 0, 0, 0, 194, 192, 1, 0, 0, 0, 194, 195,
		1, 0, 0, 0, 195, 15, 1, 0, 0, 0, 196, 194, 1, 0, 0, 0, 197, 198, 3, 12,
		6, 0, 198, 17, 1, 0, 0, 0, 199, 203, 5, 12, 0, 0, 200, 202, 5, 19, 0, 0,
		201, 200, 1, 0, 0, 0, 202, 205, 1, 0, 0, 0, 203, 201, 1, 0, 0, 0, 203,
		204, 1, 0, 0, 0, 204, 206, 1, 0, 0, 0, 205, 203, 1, 0, 0, 0, 206, 210,
		3, 20, 10, 0, 207, 209, 5, 19, 0, 0, 208, 207, 1, 0, 0, 0, 209, 212, 1,
		0, 0, 0, 210, 208, 1, 0, 0, 0, 210, 211, 1, 0, 0, 0, 211, 214, 1, 0, 0,
		0, 212, 210, 1, 0, 0, 0, 213, 215, 3, 12, 6, 0, 214, 213, 1, 0, 0, 0, 214,
		215, 1, 0, 0, 0, 215, 19, 1, 0, 0, 0, 216, 220, 5, 5, 0, 0, 217, 219, 5,
		19, 0, 0, 218, 217, 1, 0, 0, 0, 219, 222, 1, 0, 0, 0, 220, 218, 1, 0, 0,
		0, 220, 221, 1, 0, 0, 0, 221, 224, 1, 0, 0, 0, 222, 220, 1, 0, 0, 0, 223,
		225, 3, 14, 7, 0, 224, 223, 1, 0, 0, 0, 224, 225, 1, 0, 0, 0, 225, 229,
		1, 0, 0, 0, 226, 228, 5, 19, 0, 0, 227, 226, 1, 0, 0, 0, 228, 231, 1, 0,
		0, 0, 229, 227, 1, 0, 0, 0, 229, 230, 1, 0, 0, 0, 230, 232, 1, 0, 0, 0,
		231, 229, 1, 0, 0, 0, 232, 233, 5, 7, 0, 0, 233, 21, 1, 0, 0, 0, 234, 238,
		5, 13, 0, 0, 235, 237, 5, 19, 0, 0, 236, 235, 1, 0, 0, 0, 237, 240, 1,
		0, 0, 0, 238, 236, 1, 0, 0, 0, 238, 239, 1, 0, 0, 0, 239, 259, 1, 0, 0,
		0, 240, 238, 1, 0, 0, 0, 241, 245, 3, 12, 6, 0, 242, 244, 5, 19, 0, 0,
		243, 242, 1, 0, 0, 0, 244, 247, 1, 0, 0, 0, 245, 243, 1, 0, 0, 0, 245,
		246, 1, 0, 0, 0, 246, 249, 1, 0, 0, 0, 247, 245, 1, 0, 0, 0, 248, 250,
		5, 6, 0, 0, 249, 248, 1, 0, 0, 0, 249, 250, 1, 0, 0, 0, 250, 254, 1, 0,
		0, 0, 251, 253, 5, 19, 0, 0, 252, 251, 1, 0, 0, 0, 253, 256, 1, 0, 0, 0,
		254, 252, 1, 0, 0, 0, 254, 255, 1, 0, 0, 0, 255, 258, 1, 0, 0, 0, 256,
		254, 1, 0, 0, 0, 257, 241, 1, 0, 0, 0, 258, 261, 1, 0, 0, 0, 259, 257,
		1, 0, 0, 0, 259, 260, 1, 0, 0, 0, 260, 263, 1, 0, 0, 0, 261, 259, 1, 0,
		0, 0, 262, 264, 3, 12, 6, 0, 263, 262, 1, 0, 0, 0, 263, 264, 1, 0, 0, 0,
		264, 268, 1, 0, 0, 0, 265, 267, 5, 19, 0, 0, 266, 265, 1, 0, 0, 0, 267,
		270, 1, 0, 0, 0, 268, 266, 1, 0, 0, 0, 268, 269, 1, 0, 0, 0, 269, 271,
		1, 0, 0, 0, 270, 268, 1, 0, 0, 0, 271, 272, 5, 14, 0, 0, 272, 23, 1, 0,
		0, 0, 273, 274, 7, 0, 0, 0, 274, 25, 1, 0, 0, 0, 39, 29, 36, 41, 49, 55,
		60, 65, 74, 81, 88, 93, 104, 115, 120, 122, 129, 138, 148, 155, 160, 166,
		173, 181, 188, 192, 194, 203, 210, 214, 220, 224, 229, 238, 245, 249, 254,
		259, 263, 268,
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

// MinecraftMetascriptParserInit initializes any static state used to implement MinecraftMetascriptParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMinecraftMetascriptParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MinecraftMetascriptParserInit() {
	staticData := &MinecraftMetascriptParserStaticData
	staticData.once.Do(minecraftmetascriptParserInit)
}

// NewMinecraftMetascriptParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMinecraftMetascriptParser(input antlr.TokenStream) *MinecraftMetascriptParser {
	MinecraftMetascriptParserInit()
	this := new(MinecraftMetascriptParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &MinecraftMetascriptParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "MinecraftMetascript.g4"

	return this
}

// MinecraftMetascriptParser tokens.
const (
	MinecraftMetascriptParserEOF          = antlr.TokenEOF
	MinecraftMetascriptParserT__0         = 1
	MinecraftMetascriptParserT__1         = 2
	MinecraftMetascriptParserT__2         = 3
	MinecraftMetascriptParserT__3         = 4
	MinecraftMetascriptParserT__4         = 5
	MinecraftMetascriptParserT__5         = 6
	MinecraftMetascriptParserT__6         = 7
	MinecraftMetascriptParserT__7         = 8
	MinecraftMetascriptParserT__8         = 9
	MinecraftMetascriptParserT__9         = 10
	MinecraftMetascriptParserT__10        = 11
	MinecraftMetascriptParserT__11        = 12
	MinecraftMetascriptParserT__12        = 13
	MinecraftMetascriptParserT__13        = 14
	MinecraftMetascriptParserInt          = 15
	MinecraftMetascriptParserFloat        = 16
	MinecraftMetascriptParserString_      = 17
	MinecraftMetascriptParserWS           = 18
	MinecraftMetascriptParserNL           = 19
	MinecraftMetascriptParserIdentifier   = 20
	MinecraftMetascriptParserBlockComment = 21
	MinecraftMetascriptParserLineComment  = 22
)

// MinecraftMetascriptParser rules.
const (
	MinecraftMetascriptParserRULE_file              = 0
	MinecraftMetascriptParserRULE_namedBlock        = 1
	MinecraftMetascriptParserRULE_block             = 2
	MinecraftMetascriptParserRULE_varDecl           = 3
	MinecraftMetascriptParserRULE_resourceReference = 4
	MinecraftMetascriptParserRULE_fn                = 5
	MinecraftMetascriptParserRULE_value             = 6
	MinecraftMetascriptParserRULE_condition         = 7
	MinecraftMetascriptParserRULE_rootCondition     = 8
	MinecraftMetascriptParserRULE_conditional       = 9
	MinecraftMetascriptParserRULE_conditionalBody   = 10
	MinecraftMetascriptParserRULE_list              = 11
	MinecraftMetascriptParserRULE_number            = 12
)

// IFileContext is an interface to support dynamic dispatch.
type IFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode
	AllNamedBlock() []INamedBlockContext
	NamedBlock(i int) INamedBlockContext

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
	p.RuleIndex = MinecraftMetascriptParserRULE_file
	return p
}

func InitEmptyFileContext(p *FileContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_file
}

func (*FileContext) IsFileContext() {}

func NewFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileContext {
	var p = new(FileContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_file

	return p
}

func (s *FileContext) GetParser() antlr.Parser { return s.parser }

func (s *FileContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(MinecraftMetascriptParserNL)
}

func (s *FileContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserNL, i)
}

func (s *FileContext) AllNamedBlock() []INamedBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INamedBlockContext); ok {
			len++
		}
	}

	tst := make([]INamedBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INamedBlockContext); ok {
			tst[i] = t.(INamedBlockContext)
			i++
		}
	}

	return tst
}

func (s *FileContext) NamedBlock(i int) INamedBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamedBlockContext); ok {
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

	return t.(INamedBlockContext)
}

func (s *FileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterFile(s)
	}
}

func (s *FileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitFile(s)
	}
}

func (p *MinecraftMetascriptParser) File() (localctx IFileContext) {
	localctx = NewFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MinecraftMetascriptParserRULE_file)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(26)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserIdentifier {
		{
			p.SetState(32)
			p.NamedBlock()
		}
		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MinecraftMetascriptParserNL {
			{
				p.SetState(33)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(38)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(43)
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

// INamedBlockContext is an interface to support dynamic dispatch.
type INamedBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentifier() []antlr.TerminalNode
	Identifier(i int) antlr.TerminalNode
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode
	AllVarDecl() []IVarDeclContext
	VarDecl(i int) IVarDeclContext
	AllBlock() []IBlockContext
	Block(i int) IBlockContext

	// IsNamedBlockContext differentiates from other interfaces.
	IsNamedBlockContext()
}

type NamedBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamedBlockContext() *NamedBlockContext {
	var p = new(NamedBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_namedBlock
	return p
}

func InitEmptyNamedBlockContext(p *NamedBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_namedBlock
}

func (*NamedBlockContext) IsNamedBlockContext() {}

func NewNamedBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamedBlockContext {
	var p = new(NamedBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_namedBlock

	return p
}

func (s *NamedBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *NamedBlockContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(MinecraftMetascriptParserIdentifier)
}

func (s *NamedBlockContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserIdentifier, i)
}

func (s *NamedBlockContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(MinecraftMetascriptParserNL)
}

func (s *NamedBlockContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserNL, i)
}

func (s *NamedBlockContext) AllVarDecl() []IVarDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarDeclContext); ok {
			len++
		}
	}

	tst := make([]IVarDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarDeclContext); ok {
			tst[i] = t.(IVarDeclContext)
			i++
		}
	}

	return tst
}

func (s *NamedBlockContext) VarDecl(i int) IVarDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclContext); ok {
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

	return t.(IVarDeclContext)
}

func (s *NamedBlockContext) AllBlock() []IBlockContext {
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

func (s *NamedBlockContext) Block(i int) IBlockContext {
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

func (s *NamedBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamedBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamedBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterNamedBlock(s)
	}
}

func (s *NamedBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitNamedBlock(s)
	}
}

func (p *MinecraftMetascriptParser) NamedBlock() (localctx INamedBlockContext) {
	localctx = NewNamedBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MinecraftMetascriptParserRULE_namedBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		p.Match(MinecraftMetascriptParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(45)
		p.Match(MinecraftMetascriptParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(46)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(52)
		p.Match(MinecraftMetascriptParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserIdentifier {
		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(53)
				p.VarDecl()
			}

		case 2:
			{
				p.SetState(54)
				p.Block()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}
		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MinecraftMetascriptParserNL {
			{
				p.SetState(57)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(62)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(68)
		p.Match(MinecraftMetascriptParserT__1)
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

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode
	AllVarDecl() []IVarDeclContext
	VarDecl(i int) IVarDeclContext

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
	p.RuleIndex = MinecraftMetascriptParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) Identifier() antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserIdentifier, 0)
}

func (s *BlockContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(MinecraftMetascriptParserNL)
}

func (s *BlockContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserNL, i)
}

func (s *BlockContext) AllVarDecl() []IVarDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarDeclContext); ok {
			len++
		}
	}

	tst := make([]IVarDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarDeclContext); ok {
			tst[i] = t.(IVarDeclContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) VarDecl(i int) IVarDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclContext); ok {
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

	return t.(IVarDeclContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *MinecraftMetascriptParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MinecraftMetascriptParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(70)
		p.Match(MinecraftMetascriptParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(71)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(77)
		p.Match(MinecraftMetascriptParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(78)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserIdentifier {
		{
			p.SetState(84)
			p.VarDecl()
		}

		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MinecraftMetascriptParserNL {
			{
				p.SetState(85)
				p.Match(MinecraftMetascriptParserNL)
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
		}

		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(96)
		p.Match(MinecraftMetascriptParserT__1)
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

// IVarDeclContext is an interface to support dynamic dispatch.
type IVarDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	Value() IValueContext

	// IsVarDeclContext differentiates from other interfaces.
	IsVarDeclContext()
}

type VarDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDeclContext() *VarDeclContext {
	var p = new(VarDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_varDecl
	return p
}

func InitEmptyVarDeclContext(p *VarDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_varDecl
}

func (*VarDeclContext) IsVarDeclContext() {}

func NewVarDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclContext {
	var p = new(VarDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_varDecl

	return p
}

func (s *VarDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDeclContext) Identifier() antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserIdentifier, 0)
}

func (s *VarDeclContext) Value() IValueContext {
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

func (s *VarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterVarDecl(s)
	}
}

func (s *VarDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitVarDecl(s)
	}
}

func (p *MinecraftMetascriptParser) VarDecl() (localctx IVarDeclContext) {
	localctx = NewVarDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MinecraftMetascriptParserRULE_varDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(98)
		p.Match(MinecraftMetascriptParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(99)
		p.Match(MinecraftMetascriptParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(100)
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

// IResourceReferenceContext is an interface to support dynamic dispatch.
type IResourceReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentifier() []antlr.TerminalNode
	Identifier(i int) antlr.TerminalNode

	// IsResourceReferenceContext differentiates from other interfaces.
	IsResourceReferenceContext()
}

type ResourceReferenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResourceReferenceContext() *ResourceReferenceContext {
	var p = new(ResourceReferenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_resourceReference
	return p
}

func InitEmptyResourceReferenceContext(p *ResourceReferenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_resourceReference
}

func (*ResourceReferenceContext) IsResourceReferenceContext() {}

func NewResourceReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResourceReferenceContext {
	var p = new(ResourceReferenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_resourceReference

	return p
}

func (s *ResourceReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *ResourceReferenceContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(MinecraftMetascriptParserIdentifier)
}

func (s *ResourceReferenceContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserIdentifier, i)
}

func (s *ResourceReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResourceReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResourceReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterResourceReference(s)
	}
}

func (s *ResourceReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitResourceReference(s)
	}
}

func (p *MinecraftMetascriptParser) ResourceReference() (localctx IResourceReferenceContext) {
	localctx = NewResourceReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MinecraftMetascriptParserRULE_resourceReference)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(104)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(102)
			p.Match(MinecraftMetascriptParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(103)
			p.Match(MinecraftMetascriptParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(106)
		p.Match(MinecraftMetascriptParserIdentifier)
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

// IFnContext is an interface to support dynamic dispatch.
type IFnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	AllValue() []IValueContext
	Value(i int) IValueContext
	AllFn() []IFnContext
	Fn(i int) IFnContext

	// IsFnContext differentiates from other interfaces.
	IsFnContext()
}

type FnContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFnContext() *FnContext {
	var p = new(FnContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_fn
	return p
}

func InitEmptyFnContext(p *FnContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_fn
}

func (*FnContext) IsFnContext() {}

func NewFnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FnContext {
	var p = new(FnContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_fn

	return p
}

func (s *FnContext) GetParser() antlr.Parser { return s.parser }

func (s *FnContext) Identifier() antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserIdentifier, 0)
}

func (s *FnContext) AllValue() []IValueContext {
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

func (s *FnContext) Value(i int) IValueContext {
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

func (s *FnContext) AllFn() []IFnContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFnContext); ok {
			len++
		}
	}

	tst := make([]IFnContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFnContext); ok {
			tst[i] = t.(IFnContext)
			i++
		}
	}

	return tst
}

func (s *FnContext) Fn(i int) IFnContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFnContext); ok {
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

	return t.(IFnContext)
}

func (s *FnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterFn(s)
	}
}

func (s *FnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitFn(s)
	}
}

func (p *MinecraftMetascriptParser) Fn() (localctx IFnContext) {
	localctx = NewFnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MinecraftMetascriptParserRULE_fn)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		p.Match(MinecraftMetascriptParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(109)
		p.Match(MinecraftMetascriptParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(110)
				p.Value()
			}
			{
				p.SetState(111)
				p.Match(MinecraftMetascriptParserT__5)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1290240) != 0 {
		{
			p.SetState(118)
			p.Value()
		}
		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MinecraftMetascriptParserT__5 {
			{
				p.SetState(119)
				p.Match(MinecraftMetascriptParserT__5)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(124)
		p.Match(MinecraftMetascriptParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			{
				p.SetState(125)
				p.Match(MinecraftMetascriptParserT__7)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(126)
				p.Fn()
			}

		}
		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
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

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Number() INumberContext
	String_() antlr.TerminalNode
	Fn() IFnContext
	ResourceReference() IResourceReferenceContext
	Conditional() IConditionalContext
	List() IListContext

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
	p.RuleIndex = MinecraftMetascriptParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *ValueContext) String_() antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserString_, 0)
}

func (s *ValueContext) Fn() IFnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFnContext)
}

func (s *ValueContext) ResourceReference() IResourceReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IResourceReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IResourceReferenceContext)
}

func (s *ValueContext) Conditional() IConditionalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalContext)
}

func (s *ValueContext) List() IListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *MinecraftMetascriptParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MinecraftMetascriptParserRULE_value)
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(132)
			p.Number()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(133)
			p.Match(MinecraftMetascriptParserString_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(134)
			p.Fn()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(135)
			p.ResourceReference()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(136)
			p.Conditional()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(137)
			p.List()
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

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_condition
	return p
}

func InitEmptyConditionContext(p *ConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_condition
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) CopyAll(ctx *ConditionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CondGroupedContext struct {
	ConditionContext
}

func NewCondGroupedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CondGroupedContext {
	var p = new(CondGroupedContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *CondGroupedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CondGroupedContext) Condition() IConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *CondGroupedContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(MinecraftMetascriptParserNL)
}

func (s *CondGroupedContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserNL, i)
}

func (s *CondGroupedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterCondGrouped(s)
	}
}

func (s *CondGroupedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitCondGrouped(s)
	}
}

type CondNegateContext struct {
	ConditionContext
}

func NewCondNegateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CondNegateContext {
	var p = new(CondNegateContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *CondNegateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CondNegateContext) Condition() IConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *CondNegateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterCondNegate(s)
	}
}

func (s *CondNegateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitCondNegate(s)
	}
}

type CondAndContext struct {
	ConditionContext
}

func NewCondAndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CondAndContext {
	var p = new(CondAndContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *CondAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CondAndContext) AllCondition() []IConditionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConditionContext); ok {
			len++
		}
	}

	tst := make([]IConditionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConditionContext); ok {
			tst[i] = t.(IConditionContext)
			i++
		}
	}

	return tst
}

func (s *CondAndContext) Condition(i int) IConditionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
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

	return t.(IConditionContext)
}

func (s *CondAndContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(MinecraftMetascriptParserNL)
}

func (s *CondAndContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserNL, i)
}

func (s *CondAndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterCondAnd(s)
	}
}

func (s *CondAndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitCondAnd(s)
	}
}

type CondOrContext struct {
	ConditionContext
}

func NewCondOrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CondOrContext {
	var p = new(CondOrContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *CondOrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CondOrContext) AllCondition() []IConditionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConditionContext); ok {
			len++
		}
	}

	tst := make([]IConditionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConditionContext); ok {
			tst[i] = t.(IConditionContext)
			i++
		}
	}

	return tst
}

func (s *CondOrContext) Condition(i int) IConditionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
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

	return t.(IConditionContext)
}

func (s *CondOrContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(MinecraftMetascriptParserNL)
}

func (s *CondOrContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserNL, i)
}

func (s *CondOrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterCondOr(s)
	}
}

func (s *CondOrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitCondOr(s)
	}
}

type CondPrimaryContext struct {
	ConditionContext
}

func NewCondPrimaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CondPrimaryContext {
	var p = new(CondPrimaryContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *CondPrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CondPrimaryContext) RootCondition() IRootConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRootConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRootConditionContext)
}

func (s *CondPrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterCondPrimary(s)
	}
}

func (s *CondPrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitCondPrimary(s)
	}
}

func (p *MinecraftMetascriptParser) Condition() (localctx IConditionContext) {
	return p.condition(0)
}

func (p *MinecraftMetascriptParser) condition(_p int) (localctx IConditionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewConditionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IConditionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 14
	p.EnterRecursionRule(localctx, 14, MinecraftMetascriptParserRULE_condition, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinecraftMetascriptParserT__8:
		localctx = NewCondNegateContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(141)
			p.Match(MinecraftMetascriptParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(142)
			p.condition(5)
		}

	case MinecraftMetascriptParserT__11, MinecraftMetascriptParserT__12, MinecraftMetascriptParserInt, MinecraftMetascriptParserFloat, MinecraftMetascriptParserString_, MinecraftMetascriptParserIdentifier:
		localctx = NewCondPrimaryContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(143)
			p.RootCondition()
		}

	case MinecraftMetascriptParserT__4:
		localctx = NewCondGroupedContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(144)
			p.Match(MinecraftMetascriptParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MinecraftMetascriptParserNL {
			{
				p.SetState(145)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(150)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(151)
			p.condition(0)
		}
		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MinecraftMetascriptParserNL {
			{
				p.SetState(152)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(157)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(158)
			p.Match(MinecraftMetascriptParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(192)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) {
			case 1:
				localctx = NewCondAndContext(p, NewConditionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MinecraftMetascriptParserRULE_condition)
				p.SetState(162)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				p.SetState(166)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				for _la == MinecraftMetascriptParserNL {
					{
						p.SetState(163)
						p.Match(MinecraftMetascriptParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

					p.SetState(168)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(169)
					p.Match(MinecraftMetascriptParserT__9)
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

				for _la == MinecraftMetascriptParserNL {
					{
						p.SetState(170)
						p.Match(MinecraftMetascriptParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

					p.SetState(175)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(176)
					p.condition(3)
				}

			case 2:
				localctx = NewCondOrContext(p, NewConditionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MinecraftMetascriptParserRULE_condition)
				p.SetState(177)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				p.SetState(181)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				for _la == MinecraftMetascriptParserNL {
					{
						p.SetState(178)
						p.Match(MinecraftMetascriptParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

					p.SetState(183)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(184)
					p.Match(MinecraftMetascriptParserT__10)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(188)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				for _la == MinecraftMetascriptParserNL {
					{
						p.SetState(185)
						p.Match(MinecraftMetascriptParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

					p.SetState(190)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(191)
					p.condition(2)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(196)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext())
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

// IRootConditionContext is an interface to support dynamic dispatch.
type IRootConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Value() IValueContext

	// IsRootConditionContext differentiates from other interfaces.
	IsRootConditionContext()
}

type RootConditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRootConditionContext() *RootConditionContext {
	var p = new(RootConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_rootCondition
	return p
}

func InitEmptyRootConditionContext(p *RootConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_rootCondition
}

func (*RootConditionContext) IsRootConditionContext() {}

func NewRootConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootConditionContext {
	var p = new(RootConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_rootCondition

	return p
}

func (s *RootConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *RootConditionContext) Value() IValueContext {
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

func (s *RootConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RootConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterRootCondition(s)
	}
}

func (s *RootConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitRootCondition(s)
	}
}

func (p *MinecraftMetascriptParser) RootCondition() (localctx IRootConditionContext) {
	localctx = NewRootConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MinecraftMetascriptParserRULE_rootCondition)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(197)
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

// IConditionalContext is an interface to support dynamic dispatch.
type IConditionalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ConditionalBody() IConditionalBodyContext
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode
	Value() IValueContext

	// IsConditionalContext differentiates from other interfaces.
	IsConditionalContext()
}

type ConditionalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionalContext() *ConditionalContext {
	var p = new(ConditionalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_conditional
	return p
}

func InitEmptyConditionalContext(p *ConditionalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_conditional
}

func (*ConditionalContext) IsConditionalContext() {}

func NewConditionalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalContext {
	var p = new(ConditionalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_conditional

	return p
}

func (s *ConditionalContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalContext) ConditionalBody() IConditionalBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalBodyContext)
}

func (s *ConditionalContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(MinecraftMetascriptParserNL)
}

func (s *ConditionalContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserNL, i)
}

func (s *ConditionalContext) Value() IValueContext {
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

func (s *ConditionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterConditional(s)
	}
}

func (s *ConditionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitConditional(s)
	}
}

func (p *MinecraftMetascriptParser) Conditional() (localctx IConditionalContext) {
	localctx = NewConditionalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MinecraftMetascriptParserRULE_conditional)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(199)
		p.Match(MinecraftMetascriptParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(200)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(205)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(206)
		p.ConditionalBody()
	}
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(207)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(212)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(214)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(213)
			p.Value()
		}

	} else if p.HasError() { // JIM
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

// IConditionalBodyContext is an interface to support dynamic dispatch.
type IConditionalBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode
	Condition() IConditionContext

	// IsConditionalBodyContext differentiates from other interfaces.
	IsConditionalBodyContext()
}

type ConditionalBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionalBodyContext() *ConditionalBodyContext {
	var p = new(ConditionalBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_conditionalBody
	return p
}

func InitEmptyConditionalBodyContext(p *ConditionalBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_conditionalBody
}

func (*ConditionalBodyContext) IsConditionalBodyContext() {}

func NewConditionalBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalBodyContext {
	var p = new(ConditionalBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_conditionalBody

	return p
}

func (s *ConditionalBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalBodyContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(MinecraftMetascriptParserNL)
}

func (s *ConditionalBodyContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserNL, i)
}

func (s *ConditionalBodyContext) Condition() IConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *ConditionalBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionalBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterConditionalBody(s)
	}
}

func (s *ConditionalBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitConditionalBody(s)
	}
}

func (p *MinecraftMetascriptParser) ConditionalBody() (localctx IConditionalBodyContext) {
	localctx = NewConditionalBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MinecraftMetascriptParserRULE_conditionalBody)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(216)
		p.Match(MinecraftMetascriptParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(220)
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
				p.SetState(217)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(222)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1290784) != 0 {
		{
			p.SetState(223)
			p.condition(0)
		}

	}
	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(226)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(231)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(232)
		p.Match(MinecraftMetascriptParserT__6)
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

// IListContext is an interface to support dynamic dispatch.
type IListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode
	AllValue() []IValueContext
	Value(i int) IValueContext

	// IsListContext differentiates from other interfaces.
	IsListContext()
}

type ListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListContext() *ListContext {
	var p = new(ListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_list
	return p
}

func InitEmptyListContext(p *ListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_list
}

func (*ListContext) IsListContext() {}

func NewListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListContext {
	var p = new(ListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_list

	return p
}

func (s *ListContext) GetParser() antlr.Parser { return s.parser }

func (s *ListContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(MinecraftMetascriptParserNL)
}

func (s *ListContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserNL, i)
}

func (s *ListContext) AllValue() []IValueContext {
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

func (s *ListContext) Value(i int) IValueContext {
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

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterList(s)
	}
}

func (s *ListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitList(s)
	}
}

func (p *MinecraftMetascriptParser) List() (localctx IListContext) {
	localctx = NewListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, MinecraftMetascriptParserRULE_list)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(234)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(238)
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
				p.SetState(235)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(240)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(241)
				p.Value()
			}
			p.SetState(245)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
			for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				if _alt == 1 {
					{
						p.SetState(242)
						p.Match(MinecraftMetascriptParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				p.SetState(247)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext())
				if p.HasError() {
					goto errorExit
				}
			}
			p.SetState(249)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == MinecraftMetascriptParserT__5 {
				{
					p.SetState(248)
					p.Match(MinecraftMetascriptParserT__5)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(254)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
			for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				if _alt == 1 {
					{
						p.SetState(251)
						p.Match(MinecraftMetascriptParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				p.SetState(256)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext())
				if p.HasError() {
					goto errorExit
				}
			}

		}
		p.SetState(261)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(263)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1290240) != 0 {
		{
			p.SetState(262)
			p.Value()
		}

	}
	p.SetState(268)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(265)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(270)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(271)
		p.Match(MinecraftMetascriptParserT__13)
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

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Int() antlr.TerminalNode
	Float() antlr.TerminalNode

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_number
	return p
}

func InitEmptyNumberContext(p *NumberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_number
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) Int() antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserInt, 0)
}

func (s *NumberContext) Float() antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserFloat, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *MinecraftMetascriptParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, MinecraftMetascriptParserRULE_number)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(273)
		_la = p.GetTokenStream().LA(1)

		if !(_la == MinecraftMetascriptParserInt || _la == MinecraftMetascriptParserFloat) {
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

func (p *MinecraftMetascriptParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 7:
		var t *ConditionContext = nil
		if localctx != nil {
			t = localctx.(*ConditionContext)
		}
		return p.Condition_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *MinecraftMetascriptParser) Condition_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
