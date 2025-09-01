// Code generated from ./grammar/Main_Parser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // Main_Parser
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

type Main_Parser struct {
	*antlr.BaseParser
}

var Main_ParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func main_parserParserInit() {
	staticData := &Main_ParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'namespace'", "'Absolute'", "'AboveBottom'", "'BelowTop'", "'SurfaceRule'",
		"'SurfaceCondition'", "'Block'", "'If'", "'Sequence'", "'Bandlands'",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"'AboveSurface'", "'Biome'", "'Hole'", "'Noise'", "'Steep'", "'StoneDepth'",
		"'Freezing'", "'Temperature'", "'VerticalGradient'", "'AboveWater'",
		"'YAbove'", "'Floor'", "'Ceiling'", "'and'", "'or'", "'add'", "'sub'",
		"':='", "'{'", "'}'", "'['", "']'", "'('", "')'", "'.'", "'!'", "','",
		"':'", "';'", "'&'", "'&&'",
	}
	staticData.SymbolicNames = []string{
		"", "Keyword_Namespace", "Keyword_Absolute", "Keyword_AboveBottom",
		"Keyword_BelowTop", "Keyword_SurfaceRule", "Keyword_SurfaceCondition",
		"Keyword_Block", "Keyword_If", "Keyword_Sequence", "Keyword_Bandlands",
		"SurfaceRule_Head_Block", "SurfaceRule_Head_Bandlands", "SurfaceRule_Head_Condition",
		"SurfaceRule_Head_Sequence", "SurfaceCondition_Head_AboveSurface", "SurfaceCondition_Head_Biome",
		"SurfaceCondition_Head_Hole", "SurfaceCondition_Head_Noise", "SurfaceCondition_Head_Steep",
		"SurfaceCondition_Head_StoneDepth", "SurfaceCondition_Head_Freezing",
		"SurfaceCondition_Head_Temperature", "SurfaceCondition_Head_VerticalGradient",
		"SurfaceCondition_Head_AboveWater", "SurfaceCondition_Head_YAbove",
		"SurfaceCondition_Head_Floor", "SurfaceCondition_Head_Ceiling", "Keyword_AboveSurface",
		"Keyword_Biome", "Keyword_Hole", "Keyword_Noise", "Keyword_Steep", "Keyword_StoneDepth",
		"Keyword_Freezing", "Keyword_Temperature", "Keyword_VerticalGradient",
		"Keyword_AboveWater", "Keyword_YAbove", "Keyword_Floor", "Keyword_Ceiling",
		"Keyword_And", "Keyword_Or", "Keyword_Add", "Keyword_Sub", "Operator_Declare",
		"CurlyOpen", "CurlyClose", "SquareOpen", "SquareClose", "RoundOpen",
		"RoundClose", "Dot", "Bang", "Comma", "Colon", "SemiColon", "Amp", "DoubleAmp",
		"Int", "Float", "String", "NL", "WS", "Identifier", "BlockComment",
		"LineComment",
	}
	staticData.RuleNames = []string{
		"script", "namespaceDefinition", "declaration", "definition", "surfaceRuleDefinition",
		"surfaceRule", "surfaceRule_Block", "surfaceRule_Bandlands", "surfaceRule_Condition",
		"surfaceRule_Sequence", "surfaceRule_Reference", "surfaceConditionDefinition",
		"surfaceCondition", "surfaceCondition_Reference", "surfaceCondition_And",
		"surfaceCondition_Or", "surfaceCondition_AboveSurface", "surfaceCondition_Biome",
		"surfaceCondition_Hole", "surfaceCondition_Noise", "surfaceCondition_Steep",
		"surfaceCondition_StoneDepth", "surfaceCondition_Freezing", "surfaceCondition_VerticalGradient",
		"surfaceCondition_AboveWater", "surfaceCondition_YAbove", "resourceReference",
		"number", "verticalAnchorDefinition", "verticalAnchor_Absolute", "verticalAnchor_AboveBottom",
		"verticalAnchor_BelowTop", "verticalAnchor_Reference",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 66, 403, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 1, 0, 5, 0, 68, 8, 0, 10, 0, 12, 0, 71, 9, 0, 1, 0, 1,
		0, 5, 0, 75, 8, 0, 10, 0, 12, 0, 78, 9, 0, 1, 0, 1, 0, 5, 0, 82, 8, 0,
		10, 0, 12, 0, 85, 9, 0, 4, 0, 87, 8, 0, 11, 0, 12, 0, 88, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 2, 1, 2, 1, 2, 5, 2, 98, 8, 2, 10, 2, 12, 2, 101, 9, 2, 1,
		2, 1, 2, 1, 3, 1, 3, 1, 3, 3, 3, 108, 8, 3, 1, 4, 1, 4, 5, 4, 112, 8, 4,
		10, 4, 12, 4, 115, 9, 4, 1, 4, 1, 4, 5, 4, 119, 8, 4, 10, 4, 12, 4, 122,
		9, 4, 1, 4, 1, 4, 5, 4, 126, 8, 4, 10, 4, 12, 4, 129, 9, 4, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 138, 8, 5, 1, 6, 1, 6, 1, 6, 1, 7,
		1, 7, 1, 8, 1, 8, 1, 8, 5, 8, 148, 8, 8, 10, 8, 12, 8, 151, 9, 8, 1, 8,
		1, 8, 5, 8, 155, 8, 8, 10, 8, 12, 8, 158, 9, 8, 1, 8, 1, 8, 5, 8, 162,
		8, 8, 10, 8, 12, 8, 165, 9, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 5, 9, 172,
		8, 9, 10, 9, 12, 9, 175, 9, 9, 1, 9, 1, 9, 5, 9, 179, 8, 9, 10, 9, 12,
		9, 182, 9, 9, 5, 9, 184, 8, 9, 10, 9, 12, 9, 187, 9, 9, 1, 9, 5, 9, 190,
		8, 9, 10, 9, 12, 9, 193, 9, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11,
		5, 11, 201, 8, 11, 10, 11, 12, 11, 204, 9, 11, 1, 11, 1, 11, 5, 11, 208,
		8, 11, 10, 11, 12, 11, 211, 9, 11, 1, 11, 1, 11, 5, 11, 215, 8, 11, 10,
		11, 12, 11, 218, 9, 11, 1, 11, 1, 11, 1, 12, 3, 12, 223, 8, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 3, 12, 238, 8, 12, 1, 13, 1, 13, 1, 14, 1, 14, 5, 14, 244, 8,
		14, 10, 14, 12, 14, 247, 9, 14, 1, 14, 1, 14, 5, 14, 251, 8, 14, 10, 14,
		12, 14, 254, 9, 14, 1, 14, 1, 14, 5, 14, 258, 8, 14, 10, 14, 12, 14, 261,
		9, 14, 5, 14, 263, 8, 14, 10, 14, 12, 14, 266, 9, 14, 1, 14, 1, 14, 5,
		14, 270, 8, 14, 10, 14, 12, 14, 273, 9, 14, 1, 14, 1, 14, 1, 15, 1, 15,
		5, 15, 279, 8, 15, 10, 15, 12, 15, 282, 9, 15, 1, 15, 1, 15, 5, 15, 286,
		8, 15, 10, 15, 12, 15, 289, 9, 15, 1, 15, 1, 15, 5, 15, 293, 8, 15, 10,
		15, 12, 15, 296, 9, 15, 5, 15, 298, 8, 15, 10, 15, 12, 15, 301, 9, 15,
		1, 15, 1, 15, 5, 15, 305, 8, 15, 10, 15, 12, 15, 308, 9, 15, 1, 15, 1,
		15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 5, 17, 317, 8, 17, 10, 17, 12, 17,
		320, 9, 17, 1, 17, 1, 17, 5, 17, 324, 8, 17, 10, 17, 12, 17, 327, 9, 17,
		5, 17, 329, 8, 17, 10, 17, 12, 17, 332, 9, 17, 1, 17, 1, 17, 1, 18, 1,
		18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 3, 26, 374, 8, 26, 1, 26, 1, 26, 1,
		27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 3, 28, 384, 8, 28, 1, 29, 1, 29,
		1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 0, 0, 33, 0, 2, 4, 6, 8, 10, 12,
		14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48,
		50, 52, 54, 56, 58, 60, 62, 64, 0, 3, 1, 0, 39, 40, 1, 0, 43, 44, 1, 0,
		59, 60, 423, 0, 69, 1, 0, 0, 0, 2, 90, 1, 0, 0, 0, 4, 94, 1, 0, 0, 0, 6,
		107, 1, 0, 0, 0, 8, 109, 1, 0, 0, 0, 10, 137, 1, 0, 0, 0, 12, 139, 1, 0,
		0, 0, 14, 142, 1, 0, 0, 0, 16, 144, 1, 0, 0, 0, 18, 168, 1, 0, 0, 0, 20,
		196, 1, 0, 0, 0, 22, 198, 1, 0, 0, 0, 24, 222, 1, 0, 0, 0, 26, 239, 1,
		0, 0, 0, 28, 241, 1, 0, 0, 0, 30, 276, 1, 0, 0, 0, 32, 311, 1, 0, 0, 0,
		34, 313, 1, 0, 0, 0, 36, 335, 1, 0, 0, 0, 38, 337, 1, 0, 0, 0, 40, 345,
		1, 0, 0, 0, 42, 347, 1, 0, 0, 0, 44, 353, 1, 0, 0, 0, 46, 355, 1, 0, 0,
		0, 48, 361, 1, 0, 0, 0, 50, 366, 1, 0, 0, 0, 52, 373, 1, 0, 0, 0, 54, 377,
		1, 0, 0, 0, 56, 383, 1, 0, 0, 0, 58, 385, 1, 0, 0, 0, 60, 390, 1, 0, 0,
		0, 62, 395, 1, 0, 0, 0, 64, 400, 1, 0, 0, 0, 66, 68, 5, 62, 0, 0, 67, 66,
		1, 0, 0, 0, 68, 71, 1, 0, 0, 0, 69, 67, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0,
		70, 72, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 72, 76, 3, 2, 1, 0, 73, 75, 5,
		62, 0, 0, 74, 73, 1, 0, 0, 0, 75, 78, 1, 0, 0, 0, 76, 74, 1, 0, 0, 0, 76,
		77, 1, 0, 0, 0, 77, 86, 1, 0, 0, 0, 78, 76, 1, 0, 0, 0, 79, 83, 3, 4, 2,
		0, 80, 82, 5, 62, 0, 0, 81, 80, 1, 0, 0, 0, 82, 85, 1, 0, 0, 0, 83, 81,
		1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 87, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0,
		86, 79, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 88, 89, 1,
		0, 0, 0, 89, 1, 1, 0, 0, 0, 90, 91, 5, 1, 0, 0, 91, 92, 5, 64, 0, 0, 92,
		93, 5, 56, 0, 0, 93, 3, 1, 0, 0, 0, 94, 95, 5, 64, 0, 0, 95, 99, 5, 45,
		0, 0, 96, 98, 5, 62, 0, 0, 97, 96, 1, 0, 0, 0, 98, 101, 1, 0, 0, 0, 99,
		97, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 102, 1, 0, 0, 0, 101, 99, 1,
		0, 0, 0, 102, 103, 3, 6, 3, 0, 103, 5, 1, 0, 0, 0, 104, 108, 3, 8, 4, 0,
		105, 108, 3, 22, 11, 0, 106, 108, 3, 56, 28, 0, 107, 104, 1, 0, 0, 0, 107,
		105, 1, 0, 0, 0, 107, 106, 1, 0, 0, 0, 108, 7, 1, 0, 0, 0, 109, 113, 5,
		5, 0, 0, 110, 112, 5, 62, 0, 0, 111, 110, 1, 0, 0, 0, 112, 115, 1, 0, 0,
		0, 113, 111, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 116, 1, 0, 0, 0, 115,
		113, 1, 0, 0, 0, 116, 120, 5, 46, 0, 0, 117, 119, 5, 62, 0, 0, 118, 117,
		1, 0, 0, 0, 119, 122, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 120, 121, 1, 0,
		0, 0, 121, 123, 1, 0, 0, 0, 122, 120, 1, 0, 0, 0, 123, 127, 3, 10, 5, 0,
		124, 126, 5, 62, 0, 0, 125, 124, 1, 0, 0, 0, 126, 129, 1, 0, 0, 0, 127,
		125, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 130, 1, 0, 0, 0, 129, 127,
		1, 0, 0, 0, 130, 131, 5, 47, 0, 0, 131, 9, 1, 0, 0, 0, 132, 138, 3, 12,
		6, 0, 133, 138, 3, 14, 7, 0, 134, 138, 3, 16, 8, 0, 135, 138, 3, 18, 9,
		0, 136, 138, 3, 20, 10, 0, 137, 132, 1, 0, 0, 0, 137, 133, 1, 0, 0, 0,
		137, 134, 1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 137, 136, 1, 0, 0, 0, 138,
		11, 1, 0, 0, 0, 139, 140, 5, 7, 0, 0, 140, 141, 3, 52, 26, 0, 141, 13,
		1, 0, 0, 0, 142, 143, 5, 10, 0, 0, 143, 15, 1, 0, 0, 0, 144, 145, 5, 8,
		0, 0, 145, 149, 5, 50, 0, 0, 146, 148, 5, 62, 0, 0, 147, 146, 1, 0, 0,
		0, 148, 151, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150,
		152, 1, 0, 0, 0, 151, 149, 1, 0, 0, 0, 152, 156, 3, 24, 12, 0, 153, 155,
		5, 62, 0, 0, 154, 153, 1, 0, 0, 0, 155, 158, 1, 0, 0, 0, 156, 154, 1, 0,
		0, 0, 156, 157, 1, 0, 0, 0, 157, 159, 1, 0, 0, 0, 158, 156, 1, 0, 0, 0,
		159, 163, 5, 51, 0, 0, 160, 162, 5, 62, 0, 0, 161, 160, 1, 0, 0, 0, 162,
		165, 1, 0, 0, 0, 163, 161, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 166,
		1, 0, 0, 0, 165, 163, 1, 0, 0, 0, 166, 167, 3, 10, 5, 0, 167, 17, 1, 0,
		0, 0, 168, 169, 5, 9, 0, 0, 169, 173, 5, 48, 0, 0, 170, 172, 5, 62, 0,
		0, 171, 170, 1, 0, 0, 0, 172, 175, 1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 173,
		174, 1, 0, 0, 0, 174, 185, 1, 0, 0, 0, 175, 173, 1, 0, 0, 0, 176, 180,
		3, 10, 5, 0, 177, 179, 5, 62, 0, 0, 178, 177, 1, 0, 0, 0, 179, 182, 1,
		0, 0, 0, 180, 178, 1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181, 184, 1, 0, 0,
		0, 182, 180, 1, 0, 0, 0, 183, 176, 1, 0, 0, 0, 184, 187, 1, 0, 0, 0, 185,
		183, 1, 0, 0, 0, 185, 186, 1, 0, 0, 0, 186, 191, 1, 0, 0, 0, 187, 185,
		1, 0, 0, 0, 188, 190, 5, 62, 0, 0, 189, 188, 1, 0, 0, 0, 190, 193, 1, 0,
		0, 0, 191, 189, 1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 192, 194, 1, 0, 0, 0,
		193, 191, 1, 0, 0, 0, 194, 195, 5, 49, 0, 0, 195, 19, 1, 0, 0, 0, 196,
		197, 3, 52, 26, 0, 197, 21, 1, 0, 0, 0, 198, 202, 5, 6, 0, 0, 199, 201,
		5, 62, 0, 0, 200, 199, 1, 0, 0, 0, 201, 204, 1, 0, 0, 0, 202, 200, 1, 0,
		0, 0, 202, 203, 1, 0, 0, 0, 203, 205, 1, 0, 0, 0, 204, 202, 1, 0, 0, 0,
		205, 209, 5, 46, 0, 0, 206, 208, 5, 62, 0, 0, 207, 206, 1, 0, 0, 0, 208,
		211, 1, 0, 0, 0, 209, 207, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0, 210, 212,
		1, 0, 0, 0, 211, 209, 1, 0, 0, 0, 212, 216, 3, 24, 12, 0, 213, 215, 5,
		62, 0, 0, 214, 213, 1, 0, 0, 0, 215, 218, 1, 0, 0, 0, 216, 214, 1, 0, 0,
		0, 216, 217, 1, 0, 0, 0, 217, 219, 1, 0, 0, 0, 218, 216, 1, 0, 0, 0, 219,
		220, 5, 47, 0, 0, 220, 23, 1, 0, 0, 0, 221, 223, 5, 53, 0, 0, 222, 221,
		1, 0, 0, 0, 222, 223, 1, 0, 0, 0, 223, 237, 1, 0, 0, 0, 224, 238, 3, 32,
		16, 0, 225, 238, 3, 34, 17, 0, 226, 238, 3, 36, 18, 0, 227, 238, 3, 38,
		19, 0, 228, 238, 3, 40, 20, 0, 229, 238, 3, 42, 21, 0, 230, 238, 3, 44,
		22, 0, 231, 238, 3, 46, 23, 0, 232, 238, 3, 48, 24, 0, 233, 238, 3, 50,
		25, 0, 234, 238, 3, 26, 13, 0, 235, 238, 3, 28, 14, 0, 236, 238, 3, 30,
		15, 0, 237, 224, 1, 0, 0, 0, 237, 225, 1, 0, 0, 0, 237, 226, 1, 0, 0, 0,
		237, 227, 1, 0, 0, 0, 237, 228, 1, 0, 0, 0, 237, 229, 1, 0, 0, 0, 237,
		230, 1, 0, 0, 0, 237, 231, 1, 0, 0, 0, 237, 232, 1, 0, 0, 0, 237, 233,
		1, 0, 0, 0, 237, 234, 1, 0, 0, 0, 237, 235, 1, 0, 0, 0, 237, 236, 1, 0,
		0, 0, 238, 25, 1, 0, 0, 0, 239, 240, 3, 52, 26, 0, 240, 27, 1, 0, 0, 0,
		241, 245, 5, 41, 0, 0, 242, 244, 5, 62, 0, 0, 243, 242, 1, 0, 0, 0, 244,
		247, 1, 0, 0, 0, 245, 243, 1, 0, 0, 0, 245, 246, 1, 0, 0, 0, 246, 248,
		1, 0, 0, 0, 247, 245, 1, 0, 0, 0, 248, 252, 5, 50, 0, 0, 249, 251, 5, 62,
		0, 0, 250, 249, 1, 0, 0, 0, 251, 254, 1, 0, 0, 0, 252, 250, 1, 0, 0, 0,
		252, 253, 1, 0, 0, 0, 253, 264, 1, 0, 0, 0, 254, 252, 1, 0, 0, 0, 255,
		259, 3, 24, 12, 0, 256, 258, 5, 62, 0, 0, 257, 256, 1, 0, 0, 0, 258, 261,
		1, 0, 0, 0, 259, 257, 1, 0, 0, 0, 259, 260, 1, 0, 0, 0, 260, 263, 1, 0,
		0, 0, 261, 259, 1, 0, 0, 0, 262, 255, 1, 0, 0, 0, 263, 266, 1, 0, 0, 0,
		264, 262, 1, 0, 0, 0, 264, 265, 1, 0, 0, 0, 265, 267, 1, 0, 0, 0, 266,
		264, 1, 0, 0, 0, 267, 271, 3, 24, 12, 0, 268, 270, 5, 62, 0, 0, 269, 268,
		1, 0, 0, 0, 270, 273, 1, 0, 0, 0, 271, 269, 1, 0, 0, 0, 271, 272, 1, 0,
		0, 0, 272, 274, 1, 0, 0, 0, 273, 271, 1, 0, 0, 0, 274, 275, 5, 51, 0, 0,
		275, 29, 1, 0, 0, 0, 276, 280, 5, 42, 0, 0, 277, 279, 5, 62, 0, 0, 278,
		277, 1, 0, 0, 0, 279, 282, 1, 0, 0, 0, 280, 278, 1, 0, 0, 0, 280, 281,
		1, 0, 0, 0, 281, 283, 1, 0, 0, 0, 282, 280, 1, 0, 0, 0, 283, 287, 5, 50,
		0, 0, 284, 286, 5, 62, 0, 0, 285, 284, 1, 0, 0, 0, 286, 289, 1, 0, 0, 0,
		287, 285, 1, 0, 0, 0, 287, 288, 1, 0, 0, 0, 288, 299, 1, 0, 0, 0, 289,
		287, 1, 0, 0, 0, 290, 294, 3, 24, 12, 0, 291, 293, 5, 62, 0, 0, 292, 291,
		1, 0, 0, 0, 293, 296, 1, 0, 0, 0, 294, 292, 1, 0, 0, 0, 294, 295, 1, 0,
		0, 0, 295, 298, 1, 0, 0, 0, 296, 294, 1, 0, 0, 0, 297, 290, 1, 0, 0, 0,
		298, 301, 1, 0, 0, 0, 299, 297, 1, 0, 0, 0, 299, 300, 1, 0, 0, 0, 300,
		302, 1, 0, 0, 0, 301, 299, 1, 0, 0, 0, 302, 306, 3, 24, 12, 0, 303, 305,
		5, 62, 0, 0, 304, 303, 1, 0, 0, 0, 305, 308, 1, 0, 0, 0, 306, 304, 1, 0,
		0, 0, 306, 307, 1, 0, 0, 0, 307, 309, 1, 0, 0, 0, 308, 306, 1, 0, 0, 0,
		309, 310, 5, 51, 0, 0, 310, 31, 1, 0, 0, 0, 311, 312, 5, 28, 0, 0, 312,
		33, 1, 0, 0, 0, 313, 314, 5, 29, 0, 0, 314, 318, 5, 48, 0, 0, 315, 317,
		5, 62, 0, 0, 316, 315, 1, 0, 0, 0, 317, 320, 1, 0, 0, 0, 318, 316, 1, 0,
		0, 0, 318, 319, 1, 0, 0, 0, 319, 330, 1, 0, 0, 0, 320, 318, 1, 0, 0, 0,
		321, 325, 3, 52, 26, 0, 322, 324, 5, 62, 0, 0, 323, 322, 1, 0, 0, 0, 324,
		327, 1, 0, 0, 0, 325, 323, 1, 0, 0, 0, 325, 326, 1, 0, 0, 0, 326, 329,
		1, 0, 0, 0, 327, 325, 1, 0, 0, 0, 328, 321, 1, 0, 0, 0, 329, 332, 1, 0,
		0, 0, 330, 328, 1, 0, 0, 0, 330, 331, 1, 0, 0, 0, 331, 333, 1, 0, 0, 0,
		332, 330, 1, 0, 0, 0, 333, 334, 5, 49, 0, 0, 334, 35, 1, 0, 0, 0, 335,
		336, 5, 30, 0, 0, 336, 37, 1, 0, 0, 0, 337, 338, 5, 31, 0, 0, 338, 339,
		3, 52, 26, 0, 339, 340, 5, 48, 0, 0, 340, 341, 3, 54, 27, 0, 341, 342,
		5, 54, 0, 0, 342, 343, 3, 54, 27, 0, 343, 344, 5, 49, 0, 0, 344, 39, 1,
		0, 0, 0, 345, 346, 5, 32, 0, 0, 346, 41, 1, 0, 0, 0, 347, 348, 5, 33, 0,
		0, 348, 349, 7, 0, 0, 0, 349, 350, 5, 59, 0, 0, 350, 351, 7, 1, 0, 0, 351,
		352, 5, 59, 0, 0, 352, 43, 1, 0, 0, 0, 353, 354, 5, 34, 0, 0, 354, 45,
		1, 0, 0, 0, 355, 356, 5, 36, 0, 0, 356, 357, 5, 61, 0, 0, 357, 358, 3,
		56, 28, 0, 358, 359, 5, 54, 0, 0, 359, 360, 3, 56, 28, 0, 360, 47, 1, 0,
		0, 0, 361, 362, 5, 37, 0, 0, 362, 363, 5, 59, 0, 0, 363, 364, 3, 54, 27,
		0, 364, 365, 7, 1, 0, 0, 365, 49, 1, 0, 0, 0, 366, 367, 5, 38, 0, 0, 367,
		368, 3, 56, 28, 0, 368, 369, 5, 59, 0, 0, 369, 370, 7, 1, 0, 0, 370, 51,
		1, 0, 0, 0, 371, 372, 5, 64, 0, 0, 372, 374, 5, 55, 0, 0, 373, 371, 1,
		0, 0, 0, 373, 374, 1, 0, 0, 0, 374, 375, 1, 0, 0, 0, 375, 376, 5, 64, 0,
		0, 376, 53, 1, 0, 0, 0, 377, 378, 7, 2, 0, 0, 378, 55, 1, 0, 0, 0, 379,
		384, 3, 58, 29, 0, 380, 384, 3, 60, 30, 0, 381, 384, 3, 62, 31, 0, 382,
		384, 3, 64, 32, 0, 383, 379, 1, 0, 0, 0, 383, 380, 1, 0, 0, 0, 383, 381,
		1, 0, 0, 0, 383, 382, 1, 0, 0, 0, 384, 57, 1, 0, 0, 0, 385, 386, 5, 2,
		0, 0, 386, 387, 5, 50, 0, 0, 387, 388, 5, 59, 0, 0, 388, 389, 5, 51, 0,
		0, 389, 59, 1, 0, 0, 0, 390, 391, 5, 3, 0, 0, 391, 392, 5, 50, 0, 0, 392,
		393, 5, 59, 0, 0, 393, 394, 5, 51, 0, 0, 394, 61, 1, 0, 0, 0, 395, 396,
		5, 4, 0, 0, 396, 397, 5, 50, 0, 0, 397, 398, 5, 59, 0, 0, 398, 399, 5,
		51, 0, 0, 399, 63, 1, 0, 0, 0, 400, 401, 3, 52, 26, 0, 401, 65, 1, 0, 0,
		0, 37, 69, 76, 83, 88, 99, 107, 113, 120, 127, 137, 149, 156, 163, 173,
		180, 185, 191, 202, 209, 216, 222, 237, 245, 252, 259, 264, 271, 280, 287,
		294, 299, 306, 318, 325, 330, 373, 383,
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

// Main_ParserInit initializes any static state used to implement Main_Parser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMain_Parser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Main_ParserInit() {
	staticData := &Main_ParserParserStaticData
	staticData.once.Do(main_parserParserInit)
}

// NewMain_Parser produces a new parser instance for the optional input antlr.TokenStream.
func NewMain_Parser(input antlr.TokenStream) *Main_Parser {
	Main_ParserInit()
	this := new(Main_Parser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &Main_ParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Main_Parser.g4"

	return this
}

// Main_Parser tokens.
const (
	Main_ParserEOF                                    = antlr.TokenEOF
	Main_ParserKeyword_Namespace                      = 1
	Main_ParserKeyword_Absolute                       = 2
	Main_ParserKeyword_AboveBottom                    = 3
	Main_ParserKeyword_BelowTop                       = 4
	Main_ParserKeyword_SurfaceRule                    = 5
	Main_ParserKeyword_SurfaceCondition               = 6
	Main_ParserKeyword_Block                          = 7
	Main_ParserKeyword_If                             = 8
	Main_ParserKeyword_Sequence                       = 9
	Main_ParserKeyword_Bandlands                      = 10
	Main_ParserSurfaceRule_Head_Block                 = 11
	Main_ParserSurfaceRule_Head_Bandlands             = 12
	Main_ParserSurfaceRule_Head_Condition             = 13
	Main_ParserSurfaceRule_Head_Sequence              = 14
	Main_ParserSurfaceCondition_Head_AboveSurface     = 15
	Main_ParserSurfaceCondition_Head_Biome            = 16
	Main_ParserSurfaceCondition_Head_Hole             = 17
	Main_ParserSurfaceCondition_Head_Noise            = 18
	Main_ParserSurfaceCondition_Head_Steep            = 19
	Main_ParserSurfaceCondition_Head_StoneDepth       = 20
	Main_ParserSurfaceCondition_Head_Freezing         = 21
	Main_ParserSurfaceCondition_Head_Temperature      = 22
	Main_ParserSurfaceCondition_Head_VerticalGradient = 23
	Main_ParserSurfaceCondition_Head_AboveWater       = 24
	Main_ParserSurfaceCondition_Head_YAbove           = 25
	Main_ParserSurfaceCondition_Head_Floor            = 26
	Main_ParserSurfaceCondition_Head_Ceiling          = 27
	Main_ParserKeyword_AboveSurface                   = 28
	Main_ParserKeyword_Biome                          = 29
	Main_ParserKeyword_Hole                           = 30
	Main_ParserKeyword_Noise                          = 31
	Main_ParserKeyword_Steep                          = 32
	Main_ParserKeyword_StoneDepth                     = 33
	Main_ParserKeyword_Freezing                       = 34
	Main_ParserKeyword_Temperature                    = 35
	Main_ParserKeyword_VerticalGradient               = 36
	Main_ParserKeyword_AboveWater                     = 37
	Main_ParserKeyword_YAbove                         = 38
	Main_ParserKeyword_Floor                          = 39
	Main_ParserKeyword_Ceiling                        = 40
	Main_ParserKeyword_And                            = 41
	Main_ParserKeyword_Or                             = 42
	Main_ParserKeyword_Add                            = 43
	Main_ParserKeyword_Sub                            = 44
	Main_ParserOperator_Declare                       = 45
	Main_ParserCurlyOpen                              = 46
	Main_ParserCurlyClose                             = 47
	Main_ParserSquareOpen                             = 48
	Main_ParserSquareClose                            = 49
	Main_ParserRoundOpen                              = 50
	Main_ParserRoundClose                             = 51
	Main_ParserDot                                    = 52
	Main_ParserBang                                   = 53
	Main_ParserComma                                  = 54
	Main_ParserColon                                  = 55
	Main_ParserSemiColon                              = 56
	Main_ParserAmp                                    = 57
	Main_ParserDoubleAmp                              = 58
	Main_ParserInt                                    = 59
	Main_ParserFloat                                  = 60
	Main_ParserString_                                = 61
	Main_ParserNL                                     = 62
	Main_ParserWS                                     = 63
	Main_ParserIdentifier                             = 64
	Main_ParserBlockComment                           = 65
	Main_ParserLineComment                            = 66
)

// Main_Parser rules.
const (
	Main_ParserRULE_script                            = 0
	Main_ParserRULE_namespaceDefinition               = 1
	Main_ParserRULE_declaration                       = 2
	Main_ParserRULE_definition                        = 3
	Main_ParserRULE_surfaceRuleDefinition             = 4
	Main_ParserRULE_surfaceRule                       = 5
	Main_ParserRULE_surfaceRule_Block                 = 6
	Main_ParserRULE_surfaceRule_Bandlands             = 7
	Main_ParserRULE_surfaceRule_Condition             = 8
	Main_ParserRULE_surfaceRule_Sequence              = 9
	Main_ParserRULE_surfaceRule_Reference             = 10
	Main_ParserRULE_surfaceConditionDefinition        = 11
	Main_ParserRULE_surfaceCondition                  = 12
	Main_ParserRULE_surfaceCondition_Reference        = 13
	Main_ParserRULE_surfaceCondition_And              = 14
	Main_ParserRULE_surfaceCondition_Or               = 15
	Main_ParserRULE_surfaceCondition_AboveSurface     = 16
	Main_ParserRULE_surfaceCondition_Biome            = 17
	Main_ParserRULE_surfaceCondition_Hole             = 18
	Main_ParserRULE_surfaceCondition_Noise            = 19
	Main_ParserRULE_surfaceCondition_Steep            = 20
	Main_ParserRULE_surfaceCondition_StoneDepth       = 21
	Main_ParserRULE_surfaceCondition_Freezing         = 22
	Main_ParserRULE_surfaceCondition_VerticalGradient = 23
	Main_ParserRULE_surfaceCondition_AboveWater       = 24
	Main_ParserRULE_surfaceCondition_YAbove           = 25
	Main_ParserRULE_resourceReference                 = 26
	Main_ParserRULE_number                            = 27
	Main_ParserRULE_verticalAnchorDefinition          = 28
	Main_ParserRULE_verticalAnchor_Absolute           = 29
	Main_ParserRULE_verticalAnchor_AboveBottom        = 30
	Main_ParserRULE_verticalAnchor_BelowTop           = 31
	Main_ParserRULE_verticalAnchor_Reference          = 32
)

// IScriptContext is an interface to support dynamic dispatch.
type IScriptContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NamespaceDefinition() INamespaceDefinitionContext
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode
	AllDeclaration() []IDeclarationContext
	Declaration(i int) IDeclarationContext

	// IsScriptContext differentiates from other interfaces.
	IsScriptContext()
}

type ScriptContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScriptContext() *ScriptContext {
	var p = new(ScriptContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_script
	return p
}

func InitEmptyScriptContext(p *ScriptContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_script
}

func (*ScriptContext) IsScriptContext() {}

func NewScriptContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScriptContext {
	var p = new(ScriptContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Main_ParserRULE_script

	return p
}

func (s *ScriptContext) GetParser() antlr.Parser { return s.parser }

func (s *ScriptContext) NamespaceDefinition() INamespaceDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamespaceDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamespaceDefinitionContext)
}

func (s *ScriptContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(Main_ParserNL)
}

func (s *ScriptContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(Main_ParserNL, i)
}

func (s *ScriptContext) AllDeclaration() []IDeclarationContext {
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

func (s *ScriptContext) Declaration(i int) IDeclarationContext {
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

func (s *ScriptContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScriptContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScriptContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.EnterScript(s)
	}
}

func (s *ScriptContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.ExitScript(s)
	}
}

func (p *Main_Parser) Script() (localctx IScriptContext) {
	localctx = NewScriptContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Main_ParserRULE_script)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Main_ParserNL {
		{
			p.SetState(66)
			p.Match(Main_ParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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
		p.NamespaceDefinition()
	}
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Main_ParserNL {
		{
			p.SetState(73)
			p.Match(Main_ParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == Main_ParserIdentifier {
		{
			p.SetState(79)
			p.Declaration()
		}
		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == Main_ParserNL {
			{
				p.SetState(80)
				p.Match(Main_ParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(85)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(88)
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

// INamespaceDefinitionContext is an interface to support dynamic dispatch.
type INamespaceDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Keyword_Namespace() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	SemiColon() antlr.TerminalNode

	// IsNamespaceDefinitionContext differentiates from other interfaces.
	IsNamespaceDefinitionContext()
}

type NamespaceDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamespaceDefinitionContext() *NamespaceDefinitionContext {
	var p = new(NamespaceDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_namespaceDefinition
	return p
}

func InitEmptyNamespaceDefinitionContext(p *NamespaceDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_namespaceDefinition
}

func (*NamespaceDefinitionContext) IsNamespaceDefinitionContext() {}

func NewNamespaceDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamespaceDefinitionContext {
	var p = new(NamespaceDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Main_ParserRULE_namespaceDefinition

	return p
}

func (s *NamespaceDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *NamespaceDefinitionContext) Keyword_Namespace() antlr.TerminalNode {
	return s.GetToken(Main_ParserKeyword_Namespace, 0)
}

func (s *NamespaceDefinitionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(Main_ParserIdentifier, 0)
}

func (s *NamespaceDefinitionContext) SemiColon() antlr.TerminalNode {
	return s.GetToken(Main_ParserSemiColon, 0)
}

func (s *NamespaceDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamespaceDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamespaceDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.EnterNamespaceDefinition(s)
	}
}

func (s *NamespaceDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.ExitNamespaceDefinition(s)
	}
}

func (p *Main_Parser) NamespaceDefinition() (localctx INamespaceDefinitionContext) {
	localctx = NewNamespaceDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Main_ParserRULE_namespaceDefinition)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)
		p.Match(Main_ParserKeyword_Namespace)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(91)
		p.Match(Main_ParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(92)
		p.Match(Main_ParserSemiColon)
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

	// Getter signatures
	Identifier() antlr.TerminalNode
	Operator_Declare() antlr.TerminalNode
	Definition() IDefinitionContext
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_declaration
	return p
}

func InitEmptyDeclarationContext(p *DeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_declaration
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Main_ParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(Main_ParserIdentifier, 0)
}

func (s *DeclarationContext) Operator_Declare() antlr.TerminalNode {
	return s.GetToken(Main_ParserOperator_Declare, 0)
}

func (s *DeclarationContext) Definition() IDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefinitionContext)
}

func (s *DeclarationContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(Main_ParserNL)
}

func (s *DeclarationContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(Main_ParserNL, i)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (p *Main_Parser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, Main_ParserRULE_declaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(94)
		p.Match(Main_ParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(95)
		p.Match(Main_ParserOperator_Declare)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Main_ParserNL {
		{
			p.SetState(96)
			p.Match(Main_ParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(102)
		p.Definition()
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

// IDefinitionContext is an interface to support dynamic dispatch.
type IDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SurfaceRuleDefinition() ISurfaceRuleDefinitionContext
	SurfaceConditionDefinition() ISurfaceConditionDefinitionContext
	VerticalAnchorDefinition() IVerticalAnchorDefinitionContext

	// IsDefinitionContext differentiates from other interfaces.
	IsDefinitionContext()
}

type DefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefinitionContext() *DefinitionContext {
	var p = new(DefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_definition
	return p
}

func InitEmptyDefinitionContext(p *DefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_definition
}

func (*DefinitionContext) IsDefinitionContext() {}

func NewDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefinitionContext {
	var p = new(DefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Main_ParserRULE_definition

	return p
}

func (s *DefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *DefinitionContext) SurfaceRuleDefinition() ISurfaceRuleDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceRuleDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISurfaceRuleDefinitionContext)
}

func (s *DefinitionContext) SurfaceConditionDefinition() ISurfaceConditionDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceConditionDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISurfaceConditionDefinitionContext)
}

func (s *DefinitionContext) VerticalAnchorDefinition() IVerticalAnchorDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVerticalAnchorDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVerticalAnchorDefinitionContext)
}

func (s *DefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.EnterDefinition(s)
	}
}

func (s *DefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.ExitDefinition(s)
	}
}

func (p *Main_Parser) Definition() (localctx IDefinitionContext) {
	localctx = NewDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, Main_ParserRULE_definition)
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Main_ParserKeyword_SurfaceRule:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(104)
			p.SurfaceRuleDefinition()
		}

	case Main_ParserKeyword_SurfaceCondition:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(105)
			p.SurfaceConditionDefinition()
		}

	case Main_ParserKeyword_Absolute, Main_ParserKeyword_AboveBottom, Main_ParserKeyword_BelowTop, Main_ParserIdentifier:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(106)
			p.VerticalAnchorDefinition()
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

// ISurfaceRuleDefinitionContext is an interface to support dynamic dispatch.
type ISurfaceRuleDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Keyword_SurfaceRule() antlr.TerminalNode
	CurlyOpen() antlr.TerminalNode
	CurlyClose() antlr.TerminalNode
	SurfaceRule() ISurfaceRuleContext
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode

	// IsSurfaceRuleDefinitionContext differentiates from other interfaces.
	IsSurfaceRuleDefinitionContext()
}

type SurfaceRuleDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceRuleDefinitionContext() *SurfaceRuleDefinitionContext {
	var p = new(SurfaceRuleDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_surfaceRuleDefinition
	return p
}

func InitEmptySurfaceRuleDefinitionContext(p *SurfaceRuleDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_surfaceRuleDefinition
}

func (*SurfaceRuleDefinitionContext) IsSurfaceRuleDefinitionContext() {}

func NewSurfaceRuleDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceRuleDefinitionContext {
	var p = new(SurfaceRuleDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Main_ParserRULE_surfaceRuleDefinition

	return p
}

func (s *SurfaceRuleDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceRuleDefinitionContext) Keyword_SurfaceRule() antlr.TerminalNode {
	return s.GetToken(Main_ParserKeyword_SurfaceRule, 0)
}

func (s *SurfaceRuleDefinitionContext) CurlyOpen() antlr.TerminalNode {
	return s.GetToken(Main_ParserCurlyOpen, 0)
}

func (s *SurfaceRuleDefinitionContext) CurlyClose() antlr.TerminalNode {
	return s.GetToken(Main_ParserCurlyClose, 0)
}

func (s *SurfaceRuleDefinitionContext) SurfaceRule() ISurfaceRuleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceRuleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISurfaceRuleContext)
}

func (s *SurfaceRuleDefinitionContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(Main_ParserNL)
}

func (s *SurfaceRuleDefinitionContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(Main_ParserNL, i)
}

func (s *SurfaceRuleDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceRuleDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceRuleDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.EnterSurfaceRuleDefinition(s)
	}
}

func (s *SurfaceRuleDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.ExitSurfaceRuleDefinition(s)
	}
}

func (p *Main_Parser) SurfaceRuleDefinition() (localctx ISurfaceRuleDefinitionContext) {
	localctx = NewSurfaceRuleDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, Main_ParserRULE_surfaceRuleDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		p.Match(Main_ParserKeyword_SurfaceRule)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Main_ParserNL {
		{
			p.SetState(110)
			p.Match(Main_ParserNL)
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
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(116)
		p.Match(Main_ParserCurlyOpen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Main_ParserNL {
		{
			p.SetState(117)
			p.Match(Main_ParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

	{
		p.SetState(123)
		p.SurfaceRule()
	}

	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Main_ParserNL {
		{
			p.SetState(124)
			p.Match(Main_ParserNL)
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
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(130)
		p.Match(Main_ParserCurlyClose)
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

// ISurfaceRuleContext is an interface to support dynamic dispatch.
type ISurfaceRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SurfaceRule_Block() ISurfaceRule_BlockContext
	SurfaceRule_Bandlands() ISurfaceRule_BandlandsContext
	SurfaceRule_Condition() ISurfaceRule_ConditionContext
	SurfaceRule_Sequence() ISurfaceRule_SequenceContext
	SurfaceRule_Reference() ISurfaceRule_ReferenceContext

	// IsSurfaceRuleContext differentiates from other interfaces.
	IsSurfaceRuleContext()
}

type SurfaceRuleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceRuleContext() *SurfaceRuleContext {
	var p = new(SurfaceRuleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_surfaceRule
	return p
}

func InitEmptySurfaceRuleContext(p *SurfaceRuleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_surfaceRule
}

func (*SurfaceRuleContext) IsSurfaceRuleContext() {}

func NewSurfaceRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceRuleContext {
	var p = new(SurfaceRuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Main_ParserRULE_surfaceRule

	return p
}

func (s *SurfaceRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceRuleContext) SurfaceRule_Block() ISurfaceRule_BlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceRule_BlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISurfaceRule_BlockContext)
}

func (s *SurfaceRuleContext) SurfaceRule_Bandlands() ISurfaceRule_BandlandsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceRule_BandlandsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISurfaceRule_BandlandsContext)
}

func (s *SurfaceRuleContext) SurfaceRule_Condition() ISurfaceRule_ConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceRule_ConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISurfaceRule_ConditionContext)
}

func (s *SurfaceRuleContext) SurfaceRule_Sequence() ISurfaceRule_SequenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceRule_SequenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISurfaceRule_SequenceContext)
}

func (s *SurfaceRuleContext) SurfaceRule_Reference() ISurfaceRule_ReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceRule_ReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISurfaceRule_ReferenceContext)
}

func (s *SurfaceRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.EnterSurfaceRule(s)
	}
}

func (s *SurfaceRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.ExitSurfaceRule(s)
	}
}

func (p *Main_Parser) SurfaceRule() (localctx ISurfaceRuleContext) {
	localctx = NewSurfaceRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, Main_ParserRULE_surfaceRule)
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Main_ParserKeyword_Block:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(132)
			p.SurfaceRule_Block()
		}

	case Main_ParserKeyword_Bandlands:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(133)
			p.SurfaceRule_Bandlands()
		}

	case Main_ParserKeyword_If:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(134)
			p.SurfaceRule_Condition()
		}

	case Main_ParserKeyword_Sequence:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(135)
			p.SurfaceRule_Sequence()
		}

	case Main_ParserIdentifier:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(136)
			p.SurfaceRule_Reference()
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

// ISurfaceRule_BlockContext is an interface to support dynamic dispatch.
type ISurfaceRule_BlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Keyword_Block() antlr.TerminalNode
	ResourceReference() IResourceReferenceContext

	// IsSurfaceRule_BlockContext differentiates from other interfaces.
	IsSurfaceRule_BlockContext()
}

type SurfaceRule_BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceRule_BlockContext() *SurfaceRule_BlockContext {
	var p = new(SurfaceRule_BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_surfaceRule_Block
	return p
}

func InitEmptySurfaceRule_BlockContext(p *SurfaceRule_BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_surfaceRule_Block
}

func (*SurfaceRule_BlockContext) IsSurfaceRule_BlockContext() {}

func NewSurfaceRule_BlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceRule_BlockContext {
	var p = new(SurfaceRule_BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Main_ParserRULE_surfaceRule_Block

	return p
}

func (s *SurfaceRule_BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceRule_BlockContext) Keyword_Block() antlr.TerminalNode {
	return s.GetToken(Main_ParserKeyword_Block, 0)
}

func (s *SurfaceRule_BlockContext) ResourceReference() IResourceReferenceContext {
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

func (s *SurfaceRule_BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceRule_BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceRule_BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.EnterSurfaceRule_Block(s)
	}
}

func (s *SurfaceRule_BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.ExitSurfaceRule_Block(s)
	}
}

func (p *Main_Parser) SurfaceRule_Block() (localctx ISurfaceRule_BlockContext) {
	localctx = NewSurfaceRule_BlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, Main_ParserRULE_surfaceRule_Block)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(139)
		p.Match(Main_ParserKeyword_Block)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(140)
		p.ResourceReference()
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

// ISurfaceRule_BandlandsContext is an interface to support dynamic dispatch.
type ISurfaceRule_BandlandsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Keyword_Bandlands() antlr.TerminalNode

	// IsSurfaceRule_BandlandsContext differentiates from other interfaces.
	IsSurfaceRule_BandlandsContext()
}

type SurfaceRule_BandlandsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceRule_BandlandsContext() *SurfaceRule_BandlandsContext {
	var p = new(SurfaceRule_BandlandsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_surfaceRule_Bandlands
	return p
}

func InitEmptySurfaceRule_BandlandsContext(p *SurfaceRule_BandlandsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_surfaceRule_Bandlands
}

func (*SurfaceRule_BandlandsContext) IsSurfaceRule_BandlandsContext() {}

func NewSurfaceRule_BandlandsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceRule_BandlandsContext {
	var p = new(SurfaceRule_BandlandsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Main_ParserRULE_surfaceRule_Bandlands

	return p
}

func (s *SurfaceRule_BandlandsContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceRule_BandlandsContext) Keyword_Bandlands() antlr.TerminalNode {
	return s.GetToken(Main_ParserKeyword_Bandlands, 0)
}

func (s *SurfaceRule_BandlandsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceRule_BandlandsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceRule_BandlandsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.EnterSurfaceRule_Bandlands(s)
	}
}

func (s *SurfaceRule_BandlandsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.ExitSurfaceRule_Bandlands(s)
	}
}

func (p *Main_Parser) SurfaceRule_Bandlands() (localctx ISurfaceRule_BandlandsContext) {
	localctx = NewSurfaceRule_BandlandsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, Main_ParserRULE_surfaceRule_Bandlands)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(142)
		p.Match(Main_ParserKeyword_Bandlands)
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

// ISurfaceRule_ConditionContext is an interface to support dynamic dispatch.
type ISurfaceRule_ConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Keyword_If() antlr.TerminalNode
	RoundOpen() antlr.TerminalNode
	SurfaceCondition() ISurfaceConditionContext
	RoundClose() antlr.TerminalNode
	SurfaceRule() ISurfaceRuleContext
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode

	// IsSurfaceRule_ConditionContext differentiates from other interfaces.
	IsSurfaceRule_ConditionContext()
}

type SurfaceRule_ConditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceRule_ConditionContext() *SurfaceRule_ConditionContext {
	var p = new(SurfaceRule_ConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_surfaceRule_Condition
	return p
}

func InitEmptySurfaceRule_ConditionContext(p *SurfaceRule_ConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_surfaceRule_Condition
}

func (*SurfaceRule_ConditionContext) IsSurfaceRule_ConditionContext() {}

func NewSurfaceRule_ConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceRule_ConditionContext {
	var p = new(SurfaceRule_ConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Main_ParserRULE_surfaceRule_Condition

	return p
}

func (s *SurfaceRule_ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceRule_ConditionContext) Keyword_If() antlr.TerminalNode {
	return s.GetToken(Main_ParserKeyword_If, 0)
}

func (s *SurfaceRule_ConditionContext) RoundOpen() antlr.TerminalNode {
	return s.GetToken(Main_ParserRoundOpen, 0)
}

func (s *SurfaceRule_ConditionContext) SurfaceCondition() ISurfaceConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISurfaceConditionContext)
}

func (s *SurfaceRule_ConditionContext) RoundClose() antlr.TerminalNode {
	return s.GetToken(Main_ParserRoundClose, 0)
}

func (s *SurfaceRule_ConditionContext) SurfaceRule() ISurfaceRuleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceRuleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISurfaceRuleContext)
}

func (s *SurfaceRule_ConditionContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(Main_ParserNL)
}

func (s *SurfaceRule_ConditionContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(Main_ParserNL, i)
}

func (s *SurfaceRule_ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceRule_ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceRule_ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.EnterSurfaceRule_Condition(s)
	}
}

func (s *SurfaceRule_ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.ExitSurfaceRule_Condition(s)
	}
}

func (p *Main_Parser) SurfaceRule_Condition() (localctx ISurfaceRule_ConditionContext) {
	localctx = NewSurfaceRule_ConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, Main_ParserRULE_surfaceRule_Condition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)
		p.Match(Main_ParserKeyword_If)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(145)
		p.Match(Main_ParserRoundOpen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Main_ParserNL {
		{
			p.SetState(146)
			p.Match(Main_ParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(152)
		p.SurfaceCondition()
	}
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Main_ParserNL {
		{
			p.SetState(153)
			p.Match(Main_ParserNL)
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
	}
	{
		p.SetState(159)
		p.Match(Main_ParserRoundClose)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Main_ParserNL {
		{
			p.SetState(160)
			p.Match(Main_ParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(166)
		p.SurfaceRule()
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

// ISurfaceRule_SequenceContext is an interface to support dynamic dispatch.
type ISurfaceRule_SequenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Keyword_Sequence() antlr.TerminalNode
	SquareOpen() antlr.TerminalNode
	SquareClose() antlr.TerminalNode
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode
	AllSurfaceRule() []ISurfaceRuleContext
	SurfaceRule(i int) ISurfaceRuleContext

	// IsSurfaceRule_SequenceContext differentiates from other interfaces.
	IsSurfaceRule_SequenceContext()
}

type SurfaceRule_SequenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceRule_SequenceContext() *SurfaceRule_SequenceContext {
	var p = new(SurfaceRule_SequenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_surfaceRule_Sequence
	return p
}

func InitEmptySurfaceRule_SequenceContext(p *SurfaceRule_SequenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_surfaceRule_Sequence
}

func (*SurfaceRule_SequenceContext) IsSurfaceRule_SequenceContext() {}

func NewSurfaceRule_SequenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceRule_SequenceContext {
	var p = new(SurfaceRule_SequenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Main_ParserRULE_surfaceRule_Sequence

	return p
}

func (s *SurfaceRule_SequenceContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceRule_SequenceContext) Keyword_Sequence() antlr.TerminalNode {
	return s.GetToken(Main_ParserKeyword_Sequence, 0)
}

func (s *SurfaceRule_SequenceContext) SquareOpen() antlr.TerminalNode {
	return s.GetToken(Main_ParserSquareOpen, 0)
}

func (s *SurfaceRule_SequenceContext) SquareClose() antlr.TerminalNode {
	return s.GetToken(Main_ParserSquareClose, 0)
}

func (s *SurfaceRule_SequenceContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(Main_ParserNL)
}

func (s *SurfaceRule_SequenceContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(Main_ParserNL, i)
}

func (s *SurfaceRule_SequenceContext) AllSurfaceRule() []ISurfaceRuleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISurfaceRuleContext); ok {
			len++
		}
	}

	tst := make([]ISurfaceRuleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISurfaceRuleContext); ok {
			tst[i] = t.(ISurfaceRuleContext)
			i++
		}
	}

	return tst
}

func (s *SurfaceRule_SequenceContext) SurfaceRule(i int) ISurfaceRuleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceRuleContext); ok {
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

	return t.(ISurfaceRuleContext)
}

func (s *SurfaceRule_SequenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceRule_SequenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceRule_SequenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.EnterSurfaceRule_Sequence(s)
	}
}

func (s *SurfaceRule_SequenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.ExitSurfaceRule_Sequence(s)
	}
}

func (p *Main_Parser) SurfaceRule_Sequence() (localctx ISurfaceRule_SequenceContext) {
	localctx = NewSurfaceRule_SequenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, Main_ParserRULE_surfaceRule_Sequence)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		p.Match(Main_ParserKeyword_Sequence)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(169)
		p.Match(Main_ParserSquareOpen)
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
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(170)
				p.Match(Main_ParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(175)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-7)) & ^0x3f) == 0 && ((int64(1)<<(_la-7))&144115188075855887) != 0 {
		{
			p.SetState(176)
			p.SurfaceRule()
		}
		p.SetState(180)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(177)
					p.Match(Main_ParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(182)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

		p.SetState(187)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Main_ParserNL {
		{
			p.SetState(188)
			p.Match(Main_ParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(193)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(194)
		p.Match(Main_ParserSquareClose)
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

// ISurfaceRule_ReferenceContext is an interface to support dynamic dispatch.
type ISurfaceRule_ReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ResourceReference() IResourceReferenceContext

	// IsSurfaceRule_ReferenceContext differentiates from other interfaces.
	IsSurfaceRule_ReferenceContext()
}

type SurfaceRule_ReferenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceRule_ReferenceContext() *SurfaceRule_ReferenceContext {
	var p = new(SurfaceRule_ReferenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_surfaceRule_Reference
	return p
}

func InitEmptySurfaceRule_ReferenceContext(p *SurfaceRule_ReferenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_surfaceRule_Reference
}

func (*SurfaceRule_ReferenceContext) IsSurfaceRule_ReferenceContext() {}

func NewSurfaceRule_ReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceRule_ReferenceContext {
	var p = new(SurfaceRule_ReferenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Main_ParserRULE_surfaceRule_Reference

	return p
}

func (s *SurfaceRule_ReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceRule_ReferenceContext) ResourceReference() IResourceReferenceContext {
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

func (s *SurfaceRule_ReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceRule_ReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceRule_ReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.EnterSurfaceRule_Reference(s)
	}
}

func (s *SurfaceRule_ReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.ExitSurfaceRule_Reference(s)
	}
}

func (p *Main_Parser) SurfaceRule_Reference() (localctx ISurfaceRule_ReferenceContext) {
	localctx = NewSurfaceRule_ReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, Main_ParserRULE_surfaceRule_Reference)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.ResourceReference()
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

// ISurfaceConditionDefinitionContext is an interface to support dynamic dispatch.
type ISurfaceConditionDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Keyword_SurfaceCondition() antlr.TerminalNode
	CurlyOpen() antlr.TerminalNode
	SurfaceCondition() ISurfaceConditionContext
	CurlyClose() antlr.TerminalNode
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode

	// IsSurfaceConditionDefinitionContext differentiates from other interfaces.
	IsSurfaceConditionDefinitionContext()
}

type SurfaceConditionDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceConditionDefinitionContext() *SurfaceConditionDefinitionContext {
	var p = new(SurfaceConditionDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_surfaceConditionDefinition
	return p
}

func InitEmptySurfaceConditionDefinitionContext(p *SurfaceConditionDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_surfaceConditionDefinition
}

func (*SurfaceConditionDefinitionContext) IsSurfaceConditionDefinitionContext() {}

func NewSurfaceConditionDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceConditionDefinitionContext {
	var p = new(SurfaceConditionDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Main_ParserRULE_surfaceConditionDefinition

	return p
}

func (s *SurfaceConditionDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceConditionDefinitionContext) Keyword_SurfaceCondition() antlr.TerminalNode {
	return s.GetToken(Main_ParserKeyword_SurfaceCondition, 0)
}

func (s *SurfaceConditionDefinitionContext) CurlyOpen() antlr.TerminalNode {
	return s.GetToken(Main_ParserCurlyOpen, 0)
}

func (s *SurfaceConditionDefinitionContext) SurfaceCondition() ISurfaceConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISurfaceConditionContext)
}

func (s *SurfaceConditionDefinitionContext) CurlyClose() antlr.TerminalNode {
	return s.GetToken(Main_ParserCurlyClose, 0)
}

func (s *SurfaceConditionDefinitionContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(Main_ParserNL)
}

func (s *SurfaceConditionDefinitionContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(Main_ParserNL, i)
}

func (s *SurfaceConditionDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceConditionDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceConditionDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.EnterSurfaceConditionDefinition(s)
	}
}

func (s *SurfaceConditionDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.ExitSurfaceConditionDefinition(s)
	}
}

func (p *Main_Parser) SurfaceConditionDefinition() (localctx ISurfaceConditionDefinitionContext) {
	localctx = NewSurfaceConditionDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, Main_ParserRULE_surfaceConditionDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(198)
		p.Match(Main_ParserKeyword_SurfaceCondition)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Main_ParserNL {
		{
			p.SetState(199)
			p.Match(Main_ParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(204)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(205)
		p.Match(Main_ParserCurlyOpen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Main_ParserNL {
		{
			p.SetState(206)
			p.Match(Main_ParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(211)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(212)
		p.SurfaceCondition()
	}
	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Main_ParserNL {
		{
			p.SetState(213)
			p.Match(Main_ParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(218)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(219)
		p.Match(Main_ParserCurlyClose)
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

// ISurfaceConditionContext is an interface to support dynamic dispatch.
type ISurfaceConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SurfaceCondition_AboveSurface() ISurfaceCondition_AboveSurfaceContext
	SurfaceCondition_Biome() ISurfaceCondition_BiomeContext
	SurfaceCondition_Hole() ISurfaceCondition_HoleContext
	SurfaceCondition_Noise() ISurfaceCondition_NoiseContext
	SurfaceCondition_Steep() ISurfaceCondition_SteepContext
	SurfaceCondition_StoneDepth() ISurfaceCondition_StoneDepthContext
	SurfaceCondition_Freezing() ISurfaceCondition_FreezingContext
	SurfaceCondition_VerticalGradient() ISurfaceCondition_VerticalGradientContext
	SurfaceCondition_AboveWater() ISurfaceCondition_AboveWaterContext
	SurfaceCondition_YAbove() ISurfaceCondition_YAboveContext
	SurfaceCondition_Reference() ISurfaceCondition_ReferenceContext
	SurfaceCondition_And() ISurfaceCondition_AndContext
	SurfaceCondition_Or() ISurfaceCondition_OrContext
	Bang() antlr.TerminalNode

	// IsSurfaceConditionContext differentiates from other interfaces.
	IsSurfaceConditionContext()
}

type SurfaceConditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceConditionContext() *SurfaceConditionContext {
	var p = new(SurfaceConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_surfaceCondition
	return p
}

func InitEmptySurfaceConditionContext(p *SurfaceConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_surfaceCondition
}

func (*SurfaceConditionContext) IsSurfaceConditionContext() {}

func NewSurfaceConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceConditionContext {
	var p = new(SurfaceConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Main_ParserRULE_surfaceCondition

	return p
}

func (s *SurfaceConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceConditionContext) SurfaceCondition_AboveSurface() ISurfaceCondition_AboveSurfaceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceCondition_AboveSurfaceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISurfaceCondition_AboveSurfaceContext)
}

func (s *SurfaceConditionContext) SurfaceCondition_Biome() ISurfaceCondition_BiomeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceCondition_BiomeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISurfaceCondition_BiomeContext)
}

func (s *SurfaceConditionContext) SurfaceCondition_Hole() ISurfaceCondition_HoleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceCondition_HoleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISurfaceCondition_HoleContext)
}

func (s *SurfaceConditionContext) SurfaceCondition_Noise() ISurfaceCondition_NoiseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceCondition_NoiseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISurfaceCondition_NoiseContext)
}

func (s *SurfaceConditionContext) SurfaceCondition_Steep() ISurfaceCondition_SteepContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceCondition_SteepContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISurfaceCondition_SteepContext)
}

func (s *SurfaceConditionContext) SurfaceCondition_StoneDepth() ISurfaceCondition_StoneDepthContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceCondition_StoneDepthContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISurfaceCondition_StoneDepthContext)
}

func (s *SurfaceConditionContext) SurfaceCondition_Freezing() ISurfaceCondition_FreezingContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceCondition_FreezingContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISurfaceCondition_FreezingContext)
}

func (s *SurfaceConditionContext) SurfaceCondition_VerticalGradient() ISurfaceCondition_VerticalGradientContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceCondition_VerticalGradientContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISurfaceCondition_VerticalGradientContext)
}

func (s *SurfaceConditionContext) SurfaceCondition_AboveWater() ISurfaceCondition_AboveWaterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceCondition_AboveWaterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISurfaceCondition_AboveWaterContext)
}

func (s *SurfaceConditionContext) SurfaceCondition_YAbove() ISurfaceCondition_YAboveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceCondition_YAboveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISurfaceCondition_YAboveContext)
}

func (s *SurfaceConditionContext) SurfaceCondition_Reference() ISurfaceCondition_ReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceCondition_ReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISurfaceCondition_ReferenceContext)
}

func (s *SurfaceConditionContext) SurfaceCondition_And() ISurfaceCondition_AndContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceCondition_AndContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISurfaceCondition_AndContext)
}

func (s *SurfaceConditionContext) SurfaceCondition_Or() ISurfaceCondition_OrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceCondition_OrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISurfaceCondition_OrContext)
}

func (s *SurfaceConditionContext) Bang() antlr.TerminalNode {
	return s.GetToken(Main_ParserBang, 0)
}

func (s *SurfaceConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.EnterSurfaceCondition(s)
	}
}

func (s *SurfaceConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.ExitSurfaceCondition(s)
	}
}

func (p *Main_Parser) SurfaceCondition() (localctx ISurfaceConditionContext) {
	localctx = NewSurfaceConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, Main_ParserRULE_surfaceCondition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Main_ParserBang {
		{
			p.SetState(221)
			p.Match(Main_ParserBang)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Main_ParserKeyword_AboveSurface:
		{
			p.SetState(224)
			p.SurfaceCondition_AboveSurface()
		}

	case Main_ParserKeyword_Biome:
		{
			p.SetState(225)
			p.SurfaceCondition_Biome()
		}

	case Main_ParserKeyword_Hole:
		{
			p.SetState(226)
			p.SurfaceCondition_Hole()
		}

	case Main_ParserKeyword_Noise:
		{
			p.SetState(227)
			p.SurfaceCondition_Noise()
		}

	case Main_ParserKeyword_Steep:
		{
			p.SetState(228)
			p.SurfaceCondition_Steep()
		}

	case Main_ParserKeyword_StoneDepth:
		{
			p.SetState(229)
			p.SurfaceCondition_StoneDepth()
		}

	case Main_ParserKeyword_Freezing:
		{
			p.SetState(230)
			p.SurfaceCondition_Freezing()
		}

	case Main_ParserKeyword_VerticalGradient:
		{
			p.SetState(231)
			p.SurfaceCondition_VerticalGradient()
		}

	case Main_ParserKeyword_AboveWater:
		{
			p.SetState(232)
			p.SurfaceCondition_AboveWater()
		}

	case Main_ParserKeyword_YAbove:
		{
			p.SetState(233)
			p.SurfaceCondition_YAbove()
		}

	case Main_ParserIdentifier:
		{
			p.SetState(234)
			p.SurfaceCondition_Reference()
		}

	case Main_ParserKeyword_And:
		{
			p.SetState(235)
			p.SurfaceCondition_And()
		}

	case Main_ParserKeyword_Or:
		{
			p.SetState(236)
			p.SurfaceCondition_Or()
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

// ISurfaceCondition_ReferenceContext is an interface to support dynamic dispatch.
type ISurfaceCondition_ReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ResourceReference() IResourceReferenceContext

	// IsSurfaceCondition_ReferenceContext differentiates from other interfaces.
	IsSurfaceCondition_ReferenceContext()
}

type SurfaceCondition_ReferenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceCondition_ReferenceContext() *SurfaceCondition_ReferenceContext {
	var p = new(SurfaceCondition_ReferenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_surfaceCondition_Reference
	return p
}

func InitEmptySurfaceCondition_ReferenceContext(p *SurfaceCondition_ReferenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_surfaceCondition_Reference
}

func (*SurfaceCondition_ReferenceContext) IsSurfaceCondition_ReferenceContext() {}

func NewSurfaceCondition_ReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_ReferenceContext {
	var p = new(SurfaceCondition_ReferenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Main_ParserRULE_surfaceCondition_Reference

	return p
}

func (s *SurfaceCondition_ReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceCondition_ReferenceContext) ResourceReference() IResourceReferenceContext {
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

func (s *SurfaceCondition_ReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_ReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_ReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.EnterSurfaceCondition_Reference(s)
	}
}

func (s *SurfaceCondition_ReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.ExitSurfaceCondition_Reference(s)
	}
}

func (p *Main_Parser) SurfaceCondition_Reference() (localctx ISurfaceCondition_ReferenceContext) {
	localctx = NewSurfaceCondition_ReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, Main_ParserRULE_surfaceCondition_Reference)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(239)
		p.ResourceReference()
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

// ISurfaceCondition_AndContext is an interface to support dynamic dispatch.
type ISurfaceCondition_AndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Keyword_And() antlr.TerminalNode
	RoundOpen() antlr.TerminalNode
	AllSurfaceCondition() []ISurfaceConditionContext
	SurfaceCondition(i int) ISurfaceConditionContext
	RoundClose() antlr.TerminalNode
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode

	// IsSurfaceCondition_AndContext differentiates from other interfaces.
	IsSurfaceCondition_AndContext()
}

type SurfaceCondition_AndContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceCondition_AndContext() *SurfaceCondition_AndContext {
	var p = new(SurfaceCondition_AndContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_surfaceCondition_And
	return p
}

func InitEmptySurfaceCondition_AndContext(p *SurfaceCondition_AndContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_surfaceCondition_And
}

func (*SurfaceCondition_AndContext) IsSurfaceCondition_AndContext() {}

func NewSurfaceCondition_AndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_AndContext {
	var p = new(SurfaceCondition_AndContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Main_ParserRULE_surfaceCondition_And

	return p
}

func (s *SurfaceCondition_AndContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceCondition_AndContext) Keyword_And() antlr.TerminalNode {
	return s.GetToken(Main_ParserKeyword_And, 0)
}

func (s *SurfaceCondition_AndContext) RoundOpen() antlr.TerminalNode {
	return s.GetToken(Main_ParserRoundOpen, 0)
}

func (s *SurfaceCondition_AndContext) AllSurfaceCondition() []ISurfaceConditionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISurfaceConditionContext); ok {
			len++
		}
	}

	tst := make([]ISurfaceConditionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISurfaceConditionContext); ok {
			tst[i] = t.(ISurfaceConditionContext)
			i++
		}
	}

	return tst
}

func (s *SurfaceCondition_AndContext) SurfaceCondition(i int) ISurfaceConditionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceConditionContext); ok {
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

	return t.(ISurfaceConditionContext)
}

func (s *SurfaceCondition_AndContext) RoundClose() antlr.TerminalNode {
	return s.GetToken(Main_ParserRoundClose, 0)
}

func (s *SurfaceCondition_AndContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(Main_ParserNL)
}

func (s *SurfaceCondition_AndContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(Main_ParserNL, i)
}

func (s *SurfaceCondition_AndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_AndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_AndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.EnterSurfaceCondition_And(s)
	}
}

func (s *SurfaceCondition_AndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.ExitSurfaceCondition_And(s)
	}
}

func (p *Main_Parser) SurfaceCondition_And() (localctx ISurfaceCondition_AndContext) {
	localctx = NewSurfaceCondition_AndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, Main_ParserRULE_surfaceCondition_And)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(241)
		p.Match(Main_ParserKeyword_And)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Main_ParserNL {
		{
			p.SetState(242)
			p.Match(Main_ParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(247)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(248)
		p.Match(Main_ParserRoundOpen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Main_ParserNL {
		{
			p.SetState(249)
			p.Match(Main_ParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(254)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(264)
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
			{
				p.SetState(255)
				p.SurfaceCondition()
			}
			p.SetState(259)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == Main_ParserNL {
				{
					p.SetState(256)
					p.Match(Main_ParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				p.SetState(261)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		p.SetState(266)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(267)
		p.SurfaceCondition()
	}
	p.SetState(271)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Main_ParserNL {
		{
			p.SetState(268)
			p.Match(Main_ParserNL)
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
	}
	{
		p.SetState(274)
		p.Match(Main_ParserRoundClose)
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

// ISurfaceCondition_OrContext is an interface to support dynamic dispatch.
type ISurfaceCondition_OrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Keyword_Or() antlr.TerminalNode
	RoundOpen() antlr.TerminalNode
	AllSurfaceCondition() []ISurfaceConditionContext
	SurfaceCondition(i int) ISurfaceConditionContext
	RoundClose() antlr.TerminalNode
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode

	// IsSurfaceCondition_OrContext differentiates from other interfaces.
	IsSurfaceCondition_OrContext()
}

type SurfaceCondition_OrContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceCondition_OrContext() *SurfaceCondition_OrContext {
	var p = new(SurfaceCondition_OrContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_surfaceCondition_Or
	return p
}

func InitEmptySurfaceCondition_OrContext(p *SurfaceCondition_OrContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_surfaceCondition_Or
}

func (*SurfaceCondition_OrContext) IsSurfaceCondition_OrContext() {}

func NewSurfaceCondition_OrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_OrContext {
	var p = new(SurfaceCondition_OrContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Main_ParserRULE_surfaceCondition_Or

	return p
}

func (s *SurfaceCondition_OrContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceCondition_OrContext) Keyword_Or() antlr.TerminalNode {
	return s.GetToken(Main_ParserKeyword_Or, 0)
}

func (s *SurfaceCondition_OrContext) RoundOpen() antlr.TerminalNode {
	return s.GetToken(Main_ParserRoundOpen, 0)
}

func (s *SurfaceCondition_OrContext) AllSurfaceCondition() []ISurfaceConditionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISurfaceConditionContext); ok {
			len++
		}
	}

	tst := make([]ISurfaceConditionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISurfaceConditionContext); ok {
			tst[i] = t.(ISurfaceConditionContext)
			i++
		}
	}

	return tst
}

func (s *SurfaceCondition_OrContext) SurfaceCondition(i int) ISurfaceConditionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceConditionContext); ok {
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

	return t.(ISurfaceConditionContext)
}

func (s *SurfaceCondition_OrContext) RoundClose() antlr.TerminalNode {
	return s.GetToken(Main_ParserRoundClose, 0)
}

func (s *SurfaceCondition_OrContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(Main_ParserNL)
}

func (s *SurfaceCondition_OrContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(Main_ParserNL, i)
}

func (s *SurfaceCondition_OrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_OrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_OrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.EnterSurfaceCondition_Or(s)
	}
}

func (s *SurfaceCondition_OrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.ExitSurfaceCondition_Or(s)
	}
}

func (p *Main_Parser) SurfaceCondition_Or() (localctx ISurfaceCondition_OrContext) {
	localctx = NewSurfaceCondition_OrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, Main_ParserRULE_surfaceCondition_Or)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(276)
		p.Match(Main_ParserKeyword_Or)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(280)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Main_ParserNL {
		{
			p.SetState(277)
			p.Match(Main_ParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(282)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(283)
		p.Match(Main_ParserRoundOpen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(287)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Main_ParserNL {
		{
			p.SetState(284)
			p.Match(Main_ParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(289)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(290)
				p.SurfaceCondition()
			}
			p.SetState(294)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == Main_ParserNL {
				{
					p.SetState(291)
					p.Match(Main_ParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				p.SetState(296)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		p.SetState(301)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(302)
		p.SurfaceCondition()
	}
	p.SetState(306)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Main_ParserNL {
		{
			p.SetState(303)
			p.Match(Main_ParserNL)
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
	}
	{
		p.SetState(309)
		p.Match(Main_ParserRoundClose)
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

// ISurfaceCondition_AboveSurfaceContext is an interface to support dynamic dispatch.
type ISurfaceCondition_AboveSurfaceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Keyword_AboveSurface() antlr.TerminalNode

	// IsSurfaceCondition_AboveSurfaceContext differentiates from other interfaces.
	IsSurfaceCondition_AboveSurfaceContext()
}

type SurfaceCondition_AboveSurfaceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceCondition_AboveSurfaceContext() *SurfaceCondition_AboveSurfaceContext {
	var p = new(SurfaceCondition_AboveSurfaceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_surfaceCondition_AboveSurface
	return p
}

func InitEmptySurfaceCondition_AboveSurfaceContext(p *SurfaceCondition_AboveSurfaceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_surfaceCondition_AboveSurface
}

func (*SurfaceCondition_AboveSurfaceContext) IsSurfaceCondition_AboveSurfaceContext() {}

func NewSurfaceCondition_AboveSurfaceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_AboveSurfaceContext {
	var p = new(SurfaceCondition_AboveSurfaceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Main_ParserRULE_surfaceCondition_AboveSurface

	return p
}

func (s *SurfaceCondition_AboveSurfaceContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceCondition_AboveSurfaceContext) Keyword_AboveSurface() antlr.TerminalNode {
	return s.GetToken(Main_ParserKeyword_AboveSurface, 0)
}

func (s *SurfaceCondition_AboveSurfaceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_AboveSurfaceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_AboveSurfaceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.EnterSurfaceCondition_AboveSurface(s)
	}
}

func (s *SurfaceCondition_AboveSurfaceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.ExitSurfaceCondition_AboveSurface(s)
	}
}

func (p *Main_Parser) SurfaceCondition_AboveSurface() (localctx ISurfaceCondition_AboveSurfaceContext) {
	localctx = NewSurfaceCondition_AboveSurfaceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, Main_ParserRULE_surfaceCondition_AboveSurface)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(311)
		p.Match(Main_ParserKeyword_AboveSurface)
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

// ISurfaceCondition_BiomeContext is an interface to support dynamic dispatch.
type ISurfaceCondition_BiomeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Keyword_Biome() antlr.TerminalNode
	SquareOpen() antlr.TerminalNode
	SquareClose() antlr.TerminalNode
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode
	AllResourceReference() []IResourceReferenceContext
	ResourceReference(i int) IResourceReferenceContext

	// IsSurfaceCondition_BiomeContext differentiates from other interfaces.
	IsSurfaceCondition_BiomeContext()
}

type SurfaceCondition_BiomeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceCondition_BiomeContext() *SurfaceCondition_BiomeContext {
	var p = new(SurfaceCondition_BiomeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_surfaceCondition_Biome
	return p
}

func InitEmptySurfaceCondition_BiomeContext(p *SurfaceCondition_BiomeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_surfaceCondition_Biome
}

func (*SurfaceCondition_BiomeContext) IsSurfaceCondition_BiomeContext() {}

func NewSurfaceCondition_BiomeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_BiomeContext {
	var p = new(SurfaceCondition_BiomeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Main_ParserRULE_surfaceCondition_Biome

	return p
}

func (s *SurfaceCondition_BiomeContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceCondition_BiomeContext) Keyword_Biome() antlr.TerminalNode {
	return s.GetToken(Main_ParserKeyword_Biome, 0)
}

func (s *SurfaceCondition_BiomeContext) SquareOpen() antlr.TerminalNode {
	return s.GetToken(Main_ParserSquareOpen, 0)
}

func (s *SurfaceCondition_BiomeContext) SquareClose() antlr.TerminalNode {
	return s.GetToken(Main_ParserSquareClose, 0)
}

func (s *SurfaceCondition_BiomeContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(Main_ParserNL)
}

func (s *SurfaceCondition_BiomeContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(Main_ParserNL, i)
}

func (s *SurfaceCondition_BiomeContext) AllResourceReference() []IResourceReferenceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IResourceReferenceContext); ok {
			len++
		}
	}

	tst := make([]IResourceReferenceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IResourceReferenceContext); ok {
			tst[i] = t.(IResourceReferenceContext)
			i++
		}
	}

	return tst
}

func (s *SurfaceCondition_BiomeContext) ResourceReference(i int) IResourceReferenceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IResourceReferenceContext); ok {
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

	return t.(IResourceReferenceContext)
}

func (s *SurfaceCondition_BiomeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_BiomeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_BiomeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.EnterSurfaceCondition_Biome(s)
	}
}

func (s *SurfaceCondition_BiomeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.ExitSurfaceCondition_Biome(s)
	}
}

func (p *Main_Parser) SurfaceCondition_Biome() (localctx ISurfaceCondition_BiomeContext) {
	localctx = NewSurfaceCondition_BiomeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, Main_ParserRULE_surfaceCondition_Biome)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(313)
		p.Match(Main_ParserKeyword_Biome)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(314)
		p.Match(Main_ParserSquareOpen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(318)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Main_ParserNL {
		{
			p.SetState(315)
			p.Match(Main_ParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(320)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(330)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Main_ParserIdentifier {
		{
			p.SetState(321)
			p.ResourceReference()
		}
		p.SetState(325)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == Main_ParserNL {
			{
				p.SetState(322)
				p.Match(Main_ParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(327)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(332)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(333)
		p.Match(Main_ParserSquareClose)
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

// ISurfaceCondition_HoleContext is an interface to support dynamic dispatch.
type ISurfaceCondition_HoleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Keyword_Hole() antlr.TerminalNode

	// IsSurfaceCondition_HoleContext differentiates from other interfaces.
	IsSurfaceCondition_HoleContext()
}

type SurfaceCondition_HoleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceCondition_HoleContext() *SurfaceCondition_HoleContext {
	var p = new(SurfaceCondition_HoleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_surfaceCondition_Hole
	return p
}

func InitEmptySurfaceCondition_HoleContext(p *SurfaceCondition_HoleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_surfaceCondition_Hole
}

func (*SurfaceCondition_HoleContext) IsSurfaceCondition_HoleContext() {}

func NewSurfaceCondition_HoleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_HoleContext {
	var p = new(SurfaceCondition_HoleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Main_ParserRULE_surfaceCondition_Hole

	return p
}

func (s *SurfaceCondition_HoleContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceCondition_HoleContext) Keyword_Hole() antlr.TerminalNode {
	return s.GetToken(Main_ParserKeyword_Hole, 0)
}

func (s *SurfaceCondition_HoleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_HoleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_HoleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.EnterSurfaceCondition_Hole(s)
	}
}

func (s *SurfaceCondition_HoleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.ExitSurfaceCondition_Hole(s)
	}
}

func (p *Main_Parser) SurfaceCondition_Hole() (localctx ISurfaceCondition_HoleContext) {
	localctx = NewSurfaceCondition_HoleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, Main_ParserRULE_surfaceCondition_Hole)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(335)
		p.Match(Main_ParserKeyword_Hole)
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

// ISurfaceCondition_NoiseContext is an interface to support dynamic dispatch.
type ISurfaceCondition_NoiseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Keyword_Noise() antlr.TerminalNode
	ResourceReference() IResourceReferenceContext
	SquareOpen() antlr.TerminalNode
	AllNumber() []INumberContext
	Number(i int) INumberContext
	Comma() antlr.TerminalNode
	SquareClose() antlr.TerminalNode

	// IsSurfaceCondition_NoiseContext differentiates from other interfaces.
	IsSurfaceCondition_NoiseContext()
}

type SurfaceCondition_NoiseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceCondition_NoiseContext() *SurfaceCondition_NoiseContext {
	var p = new(SurfaceCondition_NoiseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_surfaceCondition_Noise
	return p
}

func InitEmptySurfaceCondition_NoiseContext(p *SurfaceCondition_NoiseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_surfaceCondition_Noise
}

func (*SurfaceCondition_NoiseContext) IsSurfaceCondition_NoiseContext() {}

func NewSurfaceCondition_NoiseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_NoiseContext {
	var p = new(SurfaceCondition_NoiseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Main_ParserRULE_surfaceCondition_Noise

	return p
}

func (s *SurfaceCondition_NoiseContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceCondition_NoiseContext) Keyword_Noise() antlr.TerminalNode {
	return s.GetToken(Main_ParserKeyword_Noise, 0)
}

func (s *SurfaceCondition_NoiseContext) ResourceReference() IResourceReferenceContext {
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

func (s *SurfaceCondition_NoiseContext) SquareOpen() antlr.TerminalNode {
	return s.GetToken(Main_ParserSquareOpen, 0)
}

func (s *SurfaceCondition_NoiseContext) AllNumber() []INumberContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INumberContext); ok {
			len++
		}
	}

	tst := make([]INumberContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INumberContext); ok {
			tst[i] = t.(INumberContext)
			i++
		}
	}

	return tst
}

func (s *SurfaceCondition_NoiseContext) Number(i int) INumberContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
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

	return t.(INumberContext)
}

func (s *SurfaceCondition_NoiseContext) Comma() antlr.TerminalNode {
	return s.GetToken(Main_ParserComma, 0)
}

func (s *SurfaceCondition_NoiseContext) SquareClose() antlr.TerminalNode {
	return s.GetToken(Main_ParserSquareClose, 0)
}

func (s *SurfaceCondition_NoiseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_NoiseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_NoiseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.EnterSurfaceCondition_Noise(s)
	}
}

func (s *SurfaceCondition_NoiseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.ExitSurfaceCondition_Noise(s)
	}
}

func (p *Main_Parser) SurfaceCondition_Noise() (localctx ISurfaceCondition_NoiseContext) {
	localctx = NewSurfaceCondition_NoiseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, Main_ParserRULE_surfaceCondition_Noise)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(337)
		p.Match(Main_ParserKeyword_Noise)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(338)
		p.ResourceReference()
	}
	{
		p.SetState(339)
		p.Match(Main_ParserSquareOpen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(340)
		p.Number()
	}
	{
		p.SetState(341)
		p.Match(Main_ParserComma)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(342)
		p.Number()
	}
	{
		p.SetState(343)
		p.Match(Main_ParserSquareClose)
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

// ISurfaceCondition_SteepContext is an interface to support dynamic dispatch.
type ISurfaceCondition_SteepContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Keyword_Steep() antlr.TerminalNode

	// IsSurfaceCondition_SteepContext differentiates from other interfaces.
	IsSurfaceCondition_SteepContext()
}

type SurfaceCondition_SteepContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceCondition_SteepContext() *SurfaceCondition_SteepContext {
	var p = new(SurfaceCondition_SteepContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_surfaceCondition_Steep
	return p
}

func InitEmptySurfaceCondition_SteepContext(p *SurfaceCondition_SteepContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_surfaceCondition_Steep
}

func (*SurfaceCondition_SteepContext) IsSurfaceCondition_SteepContext() {}

func NewSurfaceCondition_SteepContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_SteepContext {
	var p = new(SurfaceCondition_SteepContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Main_ParserRULE_surfaceCondition_Steep

	return p
}

func (s *SurfaceCondition_SteepContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceCondition_SteepContext) Keyword_Steep() antlr.TerminalNode {
	return s.GetToken(Main_ParserKeyword_Steep, 0)
}

func (s *SurfaceCondition_SteepContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_SteepContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_SteepContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.EnterSurfaceCondition_Steep(s)
	}
}

func (s *SurfaceCondition_SteepContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.ExitSurfaceCondition_Steep(s)
	}
}

func (p *Main_Parser) SurfaceCondition_Steep() (localctx ISurfaceCondition_SteepContext) {
	localctx = NewSurfaceCondition_SteepContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, Main_ParserRULE_surfaceCondition_Steep)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(345)
		p.Match(Main_ParserKeyword_Steep)
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

// ISurfaceCondition_StoneDepthContext is an interface to support dynamic dispatch.
type ISurfaceCondition_StoneDepthContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Keyword_StoneDepth() antlr.TerminalNode
	AllInt() []antlr.TerminalNode
	Int(i int) antlr.TerminalNode
	Keyword_Floor() antlr.TerminalNode
	Keyword_Ceiling() antlr.TerminalNode
	Keyword_Add() antlr.TerminalNode
	Keyword_Sub() antlr.TerminalNode

	// IsSurfaceCondition_StoneDepthContext differentiates from other interfaces.
	IsSurfaceCondition_StoneDepthContext()
}

type SurfaceCondition_StoneDepthContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceCondition_StoneDepthContext() *SurfaceCondition_StoneDepthContext {
	var p = new(SurfaceCondition_StoneDepthContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_surfaceCondition_StoneDepth
	return p
}

func InitEmptySurfaceCondition_StoneDepthContext(p *SurfaceCondition_StoneDepthContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_surfaceCondition_StoneDepth
}

func (*SurfaceCondition_StoneDepthContext) IsSurfaceCondition_StoneDepthContext() {}

func NewSurfaceCondition_StoneDepthContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_StoneDepthContext {
	var p = new(SurfaceCondition_StoneDepthContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Main_ParserRULE_surfaceCondition_StoneDepth

	return p
}

func (s *SurfaceCondition_StoneDepthContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceCondition_StoneDepthContext) Keyword_StoneDepth() antlr.TerminalNode {
	return s.GetToken(Main_ParserKeyword_StoneDepth, 0)
}

func (s *SurfaceCondition_StoneDepthContext) AllInt() []antlr.TerminalNode {
	return s.GetTokens(Main_ParserInt)
}

func (s *SurfaceCondition_StoneDepthContext) Int(i int) antlr.TerminalNode {
	return s.GetToken(Main_ParserInt, i)
}

func (s *SurfaceCondition_StoneDepthContext) Keyword_Floor() antlr.TerminalNode {
	return s.GetToken(Main_ParserKeyword_Floor, 0)
}

func (s *SurfaceCondition_StoneDepthContext) Keyword_Ceiling() antlr.TerminalNode {
	return s.GetToken(Main_ParserKeyword_Ceiling, 0)
}

func (s *SurfaceCondition_StoneDepthContext) Keyword_Add() antlr.TerminalNode {
	return s.GetToken(Main_ParserKeyword_Add, 0)
}

func (s *SurfaceCondition_StoneDepthContext) Keyword_Sub() antlr.TerminalNode {
	return s.GetToken(Main_ParserKeyword_Sub, 0)
}

func (s *SurfaceCondition_StoneDepthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_StoneDepthContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_StoneDepthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.EnterSurfaceCondition_StoneDepth(s)
	}
}

func (s *SurfaceCondition_StoneDepthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.ExitSurfaceCondition_StoneDepth(s)
	}
}

func (p *Main_Parser) SurfaceCondition_StoneDepth() (localctx ISurfaceCondition_StoneDepthContext) {
	localctx = NewSurfaceCondition_StoneDepthContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, Main_ParserRULE_surfaceCondition_StoneDepth)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(347)
		p.Match(Main_ParserKeyword_StoneDepth)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(348)
		_la = p.GetTokenStream().LA(1)

		if !(_la == Main_ParserKeyword_Floor || _la == Main_ParserKeyword_Ceiling) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(349)
		p.Match(Main_ParserInt)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(350)
		_la = p.GetTokenStream().LA(1)

		if !(_la == Main_ParserKeyword_Add || _la == Main_ParserKeyword_Sub) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(351)
		p.Match(Main_ParserInt)
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

// ISurfaceCondition_FreezingContext is an interface to support dynamic dispatch.
type ISurfaceCondition_FreezingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Keyword_Freezing() antlr.TerminalNode

	// IsSurfaceCondition_FreezingContext differentiates from other interfaces.
	IsSurfaceCondition_FreezingContext()
}

type SurfaceCondition_FreezingContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceCondition_FreezingContext() *SurfaceCondition_FreezingContext {
	var p = new(SurfaceCondition_FreezingContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_surfaceCondition_Freezing
	return p
}

func InitEmptySurfaceCondition_FreezingContext(p *SurfaceCondition_FreezingContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_surfaceCondition_Freezing
}

func (*SurfaceCondition_FreezingContext) IsSurfaceCondition_FreezingContext() {}

func NewSurfaceCondition_FreezingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_FreezingContext {
	var p = new(SurfaceCondition_FreezingContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Main_ParserRULE_surfaceCondition_Freezing

	return p
}

func (s *SurfaceCondition_FreezingContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceCondition_FreezingContext) Keyword_Freezing() antlr.TerminalNode {
	return s.GetToken(Main_ParserKeyword_Freezing, 0)
}

func (s *SurfaceCondition_FreezingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_FreezingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_FreezingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.EnterSurfaceCondition_Freezing(s)
	}
}

func (s *SurfaceCondition_FreezingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.ExitSurfaceCondition_Freezing(s)
	}
}

func (p *Main_Parser) SurfaceCondition_Freezing() (localctx ISurfaceCondition_FreezingContext) {
	localctx = NewSurfaceCondition_FreezingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, Main_ParserRULE_surfaceCondition_Freezing)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(353)
		p.Match(Main_ParserKeyword_Freezing)
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

// ISurfaceCondition_VerticalGradientContext is an interface to support dynamic dispatch.
type ISurfaceCondition_VerticalGradientContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Keyword_VerticalGradient() antlr.TerminalNode
	String_() antlr.TerminalNode
	AllVerticalAnchorDefinition() []IVerticalAnchorDefinitionContext
	VerticalAnchorDefinition(i int) IVerticalAnchorDefinitionContext
	Comma() antlr.TerminalNode

	// IsSurfaceCondition_VerticalGradientContext differentiates from other interfaces.
	IsSurfaceCondition_VerticalGradientContext()
}

type SurfaceCondition_VerticalGradientContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceCondition_VerticalGradientContext() *SurfaceCondition_VerticalGradientContext {
	var p = new(SurfaceCondition_VerticalGradientContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_surfaceCondition_VerticalGradient
	return p
}

func InitEmptySurfaceCondition_VerticalGradientContext(p *SurfaceCondition_VerticalGradientContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_surfaceCondition_VerticalGradient
}

func (*SurfaceCondition_VerticalGradientContext) IsSurfaceCondition_VerticalGradientContext() {}

func NewSurfaceCondition_VerticalGradientContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_VerticalGradientContext {
	var p = new(SurfaceCondition_VerticalGradientContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Main_ParserRULE_surfaceCondition_VerticalGradient

	return p
}

func (s *SurfaceCondition_VerticalGradientContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceCondition_VerticalGradientContext) Keyword_VerticalGradient() antlr.TerminalNode {
	return s.GetToken(Main_ParserKeyword_VerticalGradient, 0)
}

func (s *SurfaceCondition_VerticalGradientContext) String_() antlr.TerminalNode {
	return s.GetToken(Main_ParserString_, 0)
}

func (s *SurfaceCondition_VerticalGradientContext) AllVerticalAnchorDefinition() []IVerticalAnchorDefinitionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVerticalAnchorDefinitionContext); ok {
			len++
		}
	}

	tst := make([]IVerticalAnchorDefinitionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVerticalAnchorDefinitionContext); ok {
			tst[i] = t.(IVerticalAnchorDefinitionContext)
			i++
		}
	}

	return tst
}

func (s *SurfaceCondition_VerticalGradientContext) VerticalAnchorDefinition(i int) IVerticalAnchorDefinitionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVerticalAnchorDefinitionContext); ok {
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

	return t.(IVerticalAnchorDefinitionContext)
}

func (s *SurfaceCondition_VerticalGradientContext) Comma() antlr.TerminalNode {
	return s.GetToken(Main_ParserComma, 0)
}

func (s *SurfaceCondition_VerticalGradientContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_VerticalGradientContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_VerticalGradientContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.EnterSurfaceCondition_VerticalGradient(s)
	}
}

func (s *SurfaceCondition_VerticalGradientContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.ExitSurfaceCondition_VerticalGradient(s)
	}
}

func (p *Main_Parser) SurfaceCondition_VerticalGradient() (localctx ISurfaceCondition_VerticalGradientContext) {
	localctx = NewSurfaceCondition_VerticalGradientContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, Main_ParserRULE_surfaceCondition_VerticalGradient)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(355)
		p.Match(Main_ParserKeyword_VerticalGradient)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(356)
		p.Match(Main_ParserString_)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(357)
		p.VerticalAnchorDefinition()
	}
	{
		p.SetState(358)
		p.Match(Main_ParserComma)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(359)
		p.VerticalAnchorDefinition()
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

// ISurfaceCondition_AboveWaterContext is an interface to support dynamic dispatch.
type ISurfaceCondition_AboveWaterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Keyword_AboveWater() antlr.TerminalNode
	Int() antlr.TerminalNode
	Number() INumberContext
	Keyword_Add() antlr.TerminalNode
	Keyword_Sub() antlr.TerminalNode

	// IsSurfaceCondition_AboveWaterContext differentiates from other interfaces.
	IsSurfaceCondition_AboveWaterContext()
}

type SurfaceCondition_AboveWaterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceCondition_AboveWaterContext() *SurfaceCondition_AboveWaterContext {
	var p = new(SurfaceCondition_AboveWaterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_surfaceCondition_AboveWater
	return p
}

func InitEmptySurfaceCondition_AboveWaterContext(p *SurfaceCondition_AboveWaterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_surfaceCondition_AboveWater
}

func (*SurfaceCondition_AboveWaterContext) IsSurfaceCondition_AboveWaterContext() {}

func NewSurfaceCondition_AboveWaterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_AboveWaterContext {
	var p = new(SurfaceCondition_AboveWaterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Main_ParserRULE_surfaceCondition_AboveWater

	return p
}

func (s *SurfaceCondition_AboveWaterContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceCondition_AboveWaterContext) Keyword_AboveWater() antlr.TerminalNode {
	return s.GetToken(Main_ParserKeyword_AboveWater, 0)
}

func (s *SurfaceCondition_AboveWaterContext) Int() antlr.TerminalNode {
	return s.GetToken(Main_ParserInt, 0)
}

func (s *SurfaceCondition_AboveWaterContext) Number() INumberContext {
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

func (s *SurfaceCondition_AboveWaterContext) Keyword_Add() antlr.TerminalNode {
	return s.GetToken(Main_ParserKeyword_Add, 0)
}

func (s *SurfaceCondition_AboveWaterContext) Keyword_Sub() antlr.TerminalNode {
	return s.GetToken(Main_ParserKeyword_Sub, 0)
}

func (s *SurfaceCondition_AboveWaterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_AboveWaterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_AboveWaterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.EnterSurfaceCondition_AboveWater(s)
	}
}

func (s *SurfaceCondition_AboveWaterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.ExitSurfaceCondition_AboveWater(s)
	}
}

func (p *Main_Parser) SurfaceCondition_AboveWater() (localctx ISurfaceCondition_AboveWaterContext) {
	localctx = NewSurfaceCondition_AboveWaterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, Main_ParserRULE_surfaceCondition_AboveWater)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(361)
		p.Match(Main_ParserKeyword_AboveWater)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(362)
		p.Match(Main_ParserInt)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(363)
		p.Number()
	}
	{
		p.SetState(364)
		_la = p.GetTokenStream().LA(1)

		if !(_la == Main_ParserKeyword_Add || _la == Main_ParserKeyword_Sub) {
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

// ISurfaceCondition_YAboveContext is an interface to support dynamic dispatch.
type ISurfaceCondition_YAboveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Keyword_YAbove() antlr.TerminalNode
	VerticalAnchorDefinition() IVerticalAnchorDefinitionContext
	Int() antlr.TerminalNode
	Keyword_Add() antlr.TerminalNode
	Keyword_Sub() antlr.TerminalNode

	// IsSurfaceCondition_YAboveContext differentiates from other interfaces.
	IsSurfaceCondition_YAboveContext()
}

type SurfaceCondition_YAboveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceCondition_YAboveContext() *SurfaceCondition_YAboveContext {
	var p = new(SurfaceCondition_YAboveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_surfaceCondition_YAbove
	return p
}

func InitEmptySurfaceCondition_YAboveContext(p *SurfaceCondition_YAboveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_surfaceCondition_YAbove
}

func (*SurfaceCondition_YAboveContext) IsSurfaceCondition_YAboveContext() {}

func NewSurfaceCondition_YAboveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_YAboveContext {
	var p = new(SurfaceCondition_YAboveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Main_ParserRULE_surfaceCondition_YAbove

	return p
}

func (s *SurfaceCondition_YAboveContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceCondition_YAboveContext) Keyword_YAbove() antlr.TerminalNode {
	return s.GetToken(Main_ParserKeyword_YAbove, 0)
}

func (s *SurfaceCondition_YAboveContext) VerticalAnchorDefinition() IVerticalAnchorDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVerticalAnchorDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVerticalAnchorDefinitionContext)
}

func (s *SurfaceCondition_YAboveContext) Int() antlr.TerminalNode {
	return s.GetToken(Main_ParserInt, 0)
}

func (s *SurfaceCondition_YAboveContext) Keyword_Add() antlr.TerminalNode {
	return s.GetToken(Main_ParserKeyword_Add, 0)
}

func (s *SurfaceCondition_YAboveContext) Keyword_Sub() antlr.TerminalNode {
	return s.GetToken(Main_ParserKeyword_Sub, 0)
}

func (s *SurfaceCondition_YAboveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_YAboveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_YAboveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.EnterSurfaceCondition_YAbove(s)
	}
}

func (s *SurfaceCondition_YAboveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.ExitSurfaceCondition_YAbove(s)
	}
}

func (p *Main_Parser) SurfaceCondition_YAbove() (localctx ISurfaceCondition_YAboveContext) {
	localctx = NewSurfaceCondition_YAboveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, Main_ParserRULE_surfaceCondition_YAbove)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(366)
		p.Match(Main_ParserKeyword_YAbove)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(367)
		p.VerticalAnchorDefinition()
	}
	{
		p.SetState(368)
		p.Match(Main_ParserInt)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(369)
		_la = p.GetTokenStream().LA(1)

		if !(_la == Main_ParserKeyword_Add || _la == Main_ParserKeyword_Sub) {
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

// IResourceReferenceContext is an interface to support dynamic dispatch.
type IResourceReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentifier() []antlr.TerminalNode
	Identifier(i int) antlr.TerminalNode
	Colon() antlr.TerminalNode

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
	p.RuleIndex = Main_ParserRULE_resourceReference
	return p
}

func InitEmptyResourceReferenceContext(p *ResourceReferenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_resourceReference
}

func (*ResourceReferenceContext) IsResourceReferenceContext() {}

func NewResourceReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResourceReferenceContext {
	var p = new(ResourceReferenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Main_ParserRULE_resourceReference

	return p
}

func (s *ResourceReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *ResourceReferenceContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(Main_ParserIdentifier)
}

func (s *ResourceReferenceContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(Main_ParserIdentifier, i)
}

func (s *ResourceReferenceContext) Colon() antlr.TerminalNode {
	return s.GetToken(Main_ParserColon, 0)
}

func (s *ResourceReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResourceReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResourceReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.EnterResourceReference(s)
	}
}

func (s *ResourceReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.ExitResourceReference(s)
	}
}

func (p *Main_Parser) ResourceReference() (localctx IResourceReferenceContext) {
	localctx = NewResourceReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, Main_ParserRULE_resourceReference)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(373)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(371)
			p.Match(Main_ParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(372)
			p.Match(Main_ParserColon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(375)
		p.Match(Main_ParserIdentifier)
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
	p.RuleIndex = Main_ParserRULE_number
	return p
}

func InitEmptyNumberContext(p *NumberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_number
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Main_ParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) Int() antlr.TerminalNode {
	return s.GetToken(Main_ParserInt, 0)
}

func (s *NumberContext) Float() antlr.TerminalNode {
	return s.GetToken(Main_ParserFloat, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *Main_Parser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, Main_ParserRULE_number)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(377)
		_la = p.GetTokenStream().LA(1)

		if !(_la == Main_ParserInt || _la == Main_ParserFloat) {
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

// IVerticalAnchorDefinitionContext is an interface to support dynamic dispatch.
type IVerticalAnchorDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VerticalAnchor_Absolute() IVerticalAnchor_AbsoluteContext
	VerticalAnchor_AboveBottom() IVerticalAnchor_AboveBottomContext
	VerticalAnchor_BelowTop() IVerticalAnchor_BelowTopContext
	VerticalAnchor_Reference() IVerticalAnchor_ReferenceContext

	// IsVerticalAnchorDefinitionContext differentiates from other interfaces.
	IsVerticalAnchorDefinitionContext()
}

type VerticalAnchorDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVerticalAnchorDefinitionContext() *VerticalAnchorDefinitionContext {
	var p = new(VerticalAnchorDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_verticalAnchorDefinition
	return p
}

func InitEmptyVerticalAnchorDefinitionContext(p *VerticalAnchorDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_verticalAnchorDefinition
}

func (*VerticalAnchorDefinitionContext) IsVerticalAnchorDefinitionContext() {}

func NewVerticalAnchorDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VerticalAnchorDefinitionContext {
	var p = new(VerticalAnchorDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Main_ParserRULE_verticalAnchorDefinition

	return p
}

func (s *VerticalAnchorDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *VerticalAnchorDefinitionContext) VerticalAnchor_Absolute() IVerticalAnchor_AbsoluteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVerticalAnchor_AbsoluteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVerticalAnchor_AbsoluteContext)
}

func (s *VerticalAnchorDefinitionContext) VerticalAnchor_AboveBottom() IVerticalAnchor_AboveBottomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVerticalAnchor_AboveBottomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVerticalAnchor_AboveBottomContext)
}

func (s *VerticalAnchorDefinitionContext) VerticalAnchor_BelowTop() IVerticalAnchor_BelowTopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVerticalAnchor_BelowTopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVerticalAnchor_BelowTopContext)
}

func (s *VerticalAnchorDefinitionContext) VerticalAnchor_Reference() IVerticalAnchor_ReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVerticalAnchor_ReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVerticalAnchor_ReferenceContext)
}

func (s *VerticalAnchorDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VerticalAnchorDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VerticalAnchorDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.EnterVerticalAnchorDefinition(s)
	}
}

func (s *VerticalAnchorDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.ExitVerticalAnchorDefinition(s)
	}
}

func (p *Main_Parser) VerticalAnchorDefinition() (localctx IVerticalAnchorDefinitionContext) {
	localctx = NewVerticalAnchorDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, Main_ParserRULE_verticalAnchorDefinition)
	p.SetState(383)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Main_ParserKeyword_Absolute:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(379)
			p.VerticalAnchor_Absolute()
		}

	case Main_ParserKeyword_AboveBottom:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(380)
			p.VerticalAnchor_AboveBottom()
		}

	case Main_ParserKeyword_BelowTop:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(381)
			p.VerticalAnchor_BelowTop()
		}

	case Main_ParserIdentifier:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(382)
			p.VerticalAnchor_Reference()
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

// IVerticalAnchor_AbsoluteContext is an interface to support dynamic dispatch.
type IVerticalAnchor_AbsoluteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Keyword_Absolute() antlr.TerminalNode
	RoundOpen() antlr.TerminalNode
	Int() antlr.TerminalNode
	RoundClose() antlr.TerminalNode

	// IsVerticalAnchor_AbsoluteContext differentiates from other interfaces.
	IsVerticalAnchor_AbsoluteContext()
}

type VerticalAnchor_AbsoluteContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVerticalAnchor_AbsoluteContext() *VerticalAnchor_AbsoluteContext {
	var p = new(VerticalAnchor_AbsoluteContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_verticalAnchor_Absolute
	return p
}

func InitEmptyVerticalAnchor_AbsoluteContext(p *VerticalAnchor_AbsoluteContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_verticalAnchor_Absolute
}

func (*VerticalAnchor_AbsoluteContext) IsVerticalAnchor_AbsoluteContext() {}

func NewVerticalAnchor_AbsoluteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VerticalAnchor_AbsoluteContext {
	var p = new(VerticalAnchor_AbsoluteContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Main_ParserRULE_verticalAnchor_Absolute

	return p
}

func (s *VerticalAnchor_AbsoluteContext) GetParser() antlr.Parser { return s.parser }

func (s *VerticalAnchor_AbsoluteContext) Keyword_Absolute() antlr.TerminalNode {
	return s.GetToken(Main_ParserKeyword_Absolute, 0)
}

func (s *VerticalAnchor_AbsoluteContext) RoundOpen() antlr.TerminalNode {
	return s.GetToken(Main_ParserRoundOpen, 0)
}

func (s *VerticalAnchor_AbsoluteContext) Int() antlr.TerminalNode {
	return s.GetToken(Main_ParserInt, 0)
}

func (s *VerticalAnchor_AbsoluteContext) RoundClose() antlr.TerminalNode {
	return s.GetToken(Main_ParserRoundClose, 0)
}

func (s *VerticalAnchor_AbsoluteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VerticalAnchor_AbsoluteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VerticalAnchor_AbsoluteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.EnterVerticalAnchor_Absolute(s)
	}
}

func (s *VerticalAnchor_AbsoluteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.ExitVerticalAnchor_Absolute(s)
	}
}

func (p *Main_Parser) VerticalAnchor_Absolute() (localctx IVerticalAnchor_AbsoluteContext) {
	localctx = NewVerticalAnchor_AbsoluteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, Main_ParserRULE_verticalAnchor_Absolute)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(385)
		p.Match(Main_ParserKeyword_Absolute)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(386)
		p.Match(Main_ParserRoundOpen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(387)
		p.Match(Main_ParserInt)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(388)
		p.Match(Main_ParserRoundClose)
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

// IVerticalAnchor_AboveBottomContext is an interface to support dynamic dispatch.
type IVerticalAnchor_AboveBottomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Keyword_AboveBottom() antlr.TerminalNode
	RoundOpen() antlr.TerminalNode
	Int() antlr.TerminalNode
	RoundClose() antlr.TerminalNode

	// IsVerticalAnchor_AboveBottomContext differentiates from other interfaces.
	IsVerticalAnchor_AboveBottomContext()
}

type VerticalAnchor_AboveBottomContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVerticalAnchor_AboveBottomContext() *VerticalAnchor_AboveBottomContext {
	var p = new(VerticalAnchor_AboveBottomContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_verticalAnchor_AboveBottom
	return p
}

func InitEmptyVerticalAnchor_AboveBottomContext(p *VerticalAnchor_AboveBottomContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_verticalAnchor_AboveBottom
}

func (*VerticalAnchor_AboveBottomContext) IsVerticalAnchor_AboveBottomContext() {}

func NewVerticalAnchor_AboveBottomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VerticalAnchor_AboveBottomContext {
	var p = new(VerticalAnchor_AboveBottomContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Main_ParserRULE_verticalAnchor_AboveBottom

	return p
}

func (s *VerticalAnchor_AboveBottomContext) GetParser() antlr.Parser { return s.parser }

func (s *VerticalAnchor_AboveBottomContext) Keyword_AboveBottom() antlr.TerminalNode {
	return s.GetToken(Main_ParserKeyword_AboveBottom, 0)
}

func (s *VerticalAnchor_AboveBottomContext) RoundOpen() antlr.TerminalNode {
	return s.GetToken(Main_ParserRoundOpen, 0)
}

func (s *VerticalAnchor_AboveBottomContext) Int() antlr.TerminalNode {
	return s.GetToken(Main_ParserInt, 0)
}

func (s *VerticalAnchor_AboveBottomContext) RoundClose() antlr.TerminalNode {
	return s.GetToken(Main_ParserRoundClose, 0)
}

func (s *VerticalAnchor_AboveBottomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VerticalAnchor_AboveBottomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VerticalAnchor_AboveBottomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.EnterVerticalAnchor_AboveBottom(s)
	}
}

func (s *VerticalAnchor_AboveBottomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.ExitVerticalAnchor_AboveBottom(s)
	}
}

func (p *Main_Parser) VerticalAnchor_AboveBottom() (localctx IVerticalAnchor_AboveBottomContext) {
	localctx = NewVerticalAnchor_AboveBottomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, Main_ParserRULE_verticalAnchor_AboveBottom)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(390)
		p.Match(Main_ParserKeyword_AboveBottom)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(391)
		p.Match(Main_ParserRoundOpen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(392)
		p.Match(Main_ParserInt)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(393)
		p.Match(Main_ParserRoundClose)
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

// IVerticalAnchor_BelowTopContext is an interface to support dynamic dispatch.
type IVerticalAnchor_BelowTopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Keyword_BelowTop() antlr.TerminalNode
	RoundOpen() antlr.TerminalNode
	Int() antlr.TerminalNode
	RoundClose() antlr.TerminalNode

	// IsVerticalAnchor_BelowTopContext differentiates from other interfaces.
	IsVerticalAnchor_BelowTopContext()
}

type VerticalAnchor_BelowTopContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVerticalAnchor_BelowTopContext() *VerticalAnchor_BelowTopContext {
	var p = new(VerticalAnchor_BelowTopContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_verticalAnchor_BelowTop
	return p
}

func InitEmptyVerticalAnchor_BelowTopContext(p *VerticalAnchor_BelowTopContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_verticalAnchor_BelowTop
}

func (*VerticalAnchor_BelowTopContext) IsVerticalAnchor_BelowTopContext() {}

func NewVerticalAnchor_BelowTopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VerticalAnchor_BelowTopContext {
	var p = new(VerticalAnchor_BelowTopContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Main_ParserRULE_verticalAnchor_BelowTop

	return p
}

func (s *VerticalAnchor_BelowTopContext) GetParser() antlr.Parser { return s.parser }

func (s *VerticalAnchor_BelowTopContext) Keyword_BelowTop() antlr.TerminalNode {
	return s.GetToken(Main_ParserKeyword_BelowTop, 0)
}

func (s *VerticalAnchor_BelowTopContext) RoundOpen() antlr.TerminalNode {
	return s.GetToken(Main_ParserRoundOpen, 0)
}

func (s *VerticalAnchor_BelowTopContext) Int() antlr.TerminalNode {
	return s.GetToken(Main_ParserInt, 0)
}

func (s *VerticalAnchor_BelowTopContext) RoundClose() antlr.TerminalNode {
	return s.GetToken(Main_ParserRoundClose, 0)
}

func (s *VerticalAnchor_BelowTopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VerticalAnchor_BelowTopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VerticalAnchor_BelowTopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.EnterVerticalAnchor_BelowTop(s)
	}
}

func (s *VerticalAnchor_BelowTopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.ExitVerticalAnchor_BelowTop(s)
	}
}

func (p *Main_Parser) VerticalAnchor_BelowTop() (localctx IVerticalAnchor_BelowTopContext) {
	localctx = NewVerticalAnchor_BelowTopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, Main_ParserRULE_verticalAnchor_BelowTop)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(395)
		p.Match(Main_ParserKeyword_BelowTop)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(396)
		p.Match(Main_ParserRoundOpen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(397)
		p.Match(Main_ParserInt)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(398)
		p.Match(Main_ParserRoundClose)
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

// IVerticalAnchor_ReferenceContext is an interface to support dynamic dispatch.
type IVerticalAnchor_ReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ResourceReference() IResourceReferenceContext

	// IsVerticalAnchor_ReferenceContext differentiates from other interfaces.
	IsVerticalAnchor_ReferenceContext()
}

type VerticalAnchor_ReferenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVerticalAnchor_ReferenceContext() *VerticalAnchor_ReferenceContext {
	var p = new(VerticalAnchor_ReferenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_verticalAnchor_Reference
	return p
}

func InitEmptyVerticalAnchor_ReferenceContext(p *VerticalAnchor_ReferenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Main_ParserRULE_verticalAnchor_Reference
}

func (*VerticalAnchor_ReferenceContext) IsVerticalAnchor_ReferenceContext() {}

func NewVerticalAnchor_ReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VerticalAnchor_ReferenceContext {
	var p = new(VerticalAnchor_ReferenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Main_ParserRULE_verticalAnchor_Reference

	return p
}

func (s *VerticalAnchor_ReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *VerticalAnchor_ReferenceContext) ResourceReference() IResourceReferenceContext {
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

func (s *VerticalAnchor_ReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VerticalAnchor_ReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VerticalAnchor_ReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.EnterVerticalAnchor_Reference(s)
	}
}

func (s *VerticalAnchor_ReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Main_ParserListener); ok {
		listenerT.ExitVerticalAnchor_Reference(s)
	}
}

func (p *Main_Parser) VerticalAnchor_Reference() (localctx IVerticalAnchor_ReferenceContext) {
	localctx = NewVerticalAnchor_ReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, Main_ParserRULE_verticalAnchor_Reference)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(400)
		p.ResourceReference()
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
