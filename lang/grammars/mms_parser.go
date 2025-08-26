// Code generated from ./grammars/MMSParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammars // MMSParser
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
		"", "", "", "'surface'", "'rule'", "'condition'", "'sequence'", "'block'",
		"'bandlands'", "'above_preliminary_surface'", "'biome'", "'hole'", "'noise'",
		"'steep'", "'stone_depth'", "'freezing'", "'temperature'", "'vertical_gradient'",
		"'above_water'", "'y_above'", "'floor'", "'ceiling'", "'and'", "'add'",
		"'sub'", "'worldgen'", "'absolute'", "'above_bottom'", "'below_top'",
		"'namespace'", "'if'", "'else'", "'in'", "", "", "'['", "']'", "'{'",
		"'}'", "'('", "')'", "'!'", "','", "':'", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "Int", "Float", "Keyword_Surface", "Keyword_Rule", "Keyword_Condition",
		"Keyword_Sequence", "Keyword_Block", "Keyword_Bandlands", "Keyword_AbovePreliminarySurface",
		"Keyword_Biome", "Keyword_Hole", "Keyword_Noise", "Keyword_Steep", "Keyword_StoneDepth",
		"Keyword_Freezing", "Keyword_Temperature", "Keyword_VerticalGradient",
		"Keyword_AboveWater", "Keyword_YAbove", "Keyword_Floor", "Keyword_Ceiling",
		"Keyword_And", "Keyword_Add", "Keyword_Sub", "Keyword_WorldGen", "Keyword_Absolute",
		"Keyword_AboveBottom", "Keyword_BelowTop", "Keyword_Namespace", "Keyword_If",
		"Keyword_Else", "Keyword_In", "WS", "NL", "SquareOpen", "SquareClose",
		"CurlyOpen", "CurlyClose", "RoundOpen", "RoundClose", "Bang", "Comma",
		"Colon", "SemiColon", "String", "Identifier", "LineComment", "BlockComment",
	}
	staticData.RuleNames = []string{
		"namespaceDeclaration", "statement", "mmsFile", "surfaceDeclaration",
		"surfaceDefinition", "surfaceRuleReference", "surfaceRuleDeclaration",
		"surfaceRule", "surfaceRule_Conditional", "surfaceRule_Bandlands", "surfaceRule_Block",
		"surfaceRule_Sequence", "surfaceConditionReference", "surfaceConditionDeclaration",
		"surfaceCondition", "surfaceCondition_AboveSurface", "surfaceCondition_Biome",
		"surfaceCondition_Hole", "surfaceCondition_Noise", "surfaceCondition_Steep",
		"surfaceCondition_StoneDepth", "surfaceCondition_Freezing", "surfaceCondition_VerticalGradient",
		"surfaceCondition_AboveWater", "surfaceCondition_YAbove", "surfaceCondition_Compound__Item",
		"surfaceCondition_Compound", "reference", "resourceReference", "number",
		"keyword", "verticalAnchor", "verticalAnchor_Absolute", "verticalAnchor_AboveBottom",
		"verticalAnchor_BelowTop", "worldgenDeclaration", "noiseDeclaration",
		"noiseDefinition",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 48, 451, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 3, 1, 83, 8, 1, 1, 2,
		5, 2, 86, 8, 2, 10, 2, 12, 2, 89, 9, 2, 1, 2, 1, 2, 5, 2, 93, 8, 2, 10,
		2, 12, 2, 96, 9, 2, 1, 2, 1, 2, 4, 2, 100, 8, 2, 11, 2, 12, 2, 101, 5,
		2, 104, 8, 2, 10, 2, 12, 2, 107, 9, 2, 1, 2, 3, 2, 110, 8, 2, 1, 2, 5,
		2, 113, 8, 2, 10, 2, 12, 2, 116, 9, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1,
		4, 1, 4, 5, 4, 125, 8, 4, 10, 4, 12, 4, 128, 9, 4, 1, 4, 1, 4, 3, 4, 132,
		8, 4, 1, 4, 5, 4, 135, 8, 4, 10, 4, 12, 4, 138, 9, 4, 5, 4, 140, 8, 4,
		10, 4, 12, 4, 143, 9, 4, 1, 4, 1, 4, 1, 5, 1, 5, 3, 5, 149, 8, 5, 1, 6,
		1, 6, 1, 6, 5, 6, 154, 8, 6, 10, 6, 12, 6, 157, 9, 6, 1, 6, 1, 6, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 166, 8, 7, 1, 8, 1, 8, 5, 8, 170, 8, 8, 10,
		8, 12, 8, 173, 9, 8, 1, 8, 3, 8, 176, 8, 8, 1, 8, 1, 8, 5, 8, 180, 8, 8,
		10, 8, 12, 8, 183, 9, 8, 1, 8, 1, 8, 5, 8, 187, 8, 8, 10, 8, 12, 8, 190,
		9, 8, 1, 8, 1, 8, 5, 8, 194, 8, 8, 10, 8, 12, 8, 197, 9, 8, 1, 8, 1, 8,
		1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 5, 11, 209, 8, 11,
		10, 11, 12, 11, 212, 9, 11, 1, 11, 1, 11, 5, 11, 216, 8, 11, 10, 11, 12,
		11, 219, 9, 11, 5, 11, 221, 8, 11, 10, 11, 12, 11, 224, 9, 11, 1, 11, 1,
		11, 1, 12, 1, 12, 3, 12, 230, 8, 12, 1, 13, 1, 13, 1, 13, 5, 13, 235, 8,
		13, 10, 13, 12, 13, 238, 9, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 254, 8,
		14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 5, 16, 261, 8, 16, 10, 16, 12, 16,
		264, 9, 16, 1, 16, 1, 16, 5, 16, 268, 8, 16, 10, 16, 12, 16, 271, 9, 16,
		5, 16, 273, 8, 16, 10, 16, 12, 16, 276, 9, 16, 1, 16, 1, 16, 1, 17, 1,
		17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 24, 1, 25, 3, 25, 317, 8, 25, 1, 25, 1, 25, 1, 26, 1,
		26, 5, 26, 323, 8, 26, 10, 26, 12, 26, 326, 9, 26, 1, 26, 1, 26, 5, 26,
		330, 8, 26, 10, 26, 12, 26, 333, 9, 26, 1, 26, 1, 26, 5, 26, 337, 8, 26,
		10, 26, 12, 26, 340, 9, 26, 5, 26, 342, 8, 26, 10, 26, 12, 26, 345, 9,
		26, 1, 26, 1, 26, 1, 27, 1, 27, 3, 27, 351, 8, 27, 1, 27, 1, 27, 1, 27,
		3, 27, 356, 8, 27, 1, 28, 1, 28, 1, 28, 3, 28, 361, 8, 28, 3, 28, 363,
		8, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 3, 31, 372, 8,
		31, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 35,
		1, 35, 1, 35, 5, 35, 386, 8, 35, 10, 35, 12, 35, 389, 9, 35, 1, 35, 1,
		35, 5, 35, 393, 8, 35, 10, 35, 12, 35, 396, 9, 35, 5, 35, 398, 8, 35, 10,
		35, 12, 35, 401, 9, 35, 1, 35, 1, 35, 5, 35, 405, 8, 35, 10, 35, 12, 35,
		408, 9, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 5, 36, 415, 8, 36, 10, 36,
		12, 36, 418, 9, 36, 1, 36, 1, 36, 1, 37, 1, 37, 5, 37, 424, 8, 37, 10,
		37, 12, 37, 427, 9, 37, 1, 37, 1, 37, 5, 37, 431, 8, 37, 10, 37, 12, 37,
		434, 9, 37, 1, 37, 1, 37, 5, 37, 438, 8, 37, 10, 37, 12, 37, 441, 9, 37,
		4, 37, 443, 8, 37, 11, 37, 12, 37, 444, 1, 37, 1, 37, 3, 37, 449, 8, 37,
		1, 37, 0, 0, 38, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
		30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64,
		66, 68, 70, 72, 74, 0, 4, 1, 0, 20, 21, 1, 0, 23, 24, 1, 0, 1, 2, 2, 0,
		3, 24, 26, 32, 474, 0, 76, 1, 0, 0, 0, 2, 82, 1, 0, 0, 0, 4, 87, 1, 0,
		0, 0, 6, 119, 1, 0, 0, 0, 8, 122, 1, 0, 0, 0, 10, 148, 1, 0, 0, 0, 12,
		150, 1, 0, 0, 0, 14, 165, 1, 0, 0, 0, 16, 167, 1, 0, 0, 0, 18, 200, 1,
		0, 0, 0, 20, 202, 1, 0, 0, 0, 22, 205, 1, 0, 0, 0, 24, 229, 1, 0, 0, 0,
		26, 231, 1, 0, 0, 0, 28, 253, 1, 0, 0, 0, 30, 255, 1, 0, 0, 0, 32, 257,
		1, 0, 0, 0, 34, 279, 1, 0, 0, 0, 36, 281, 1, 0, 0, 0, 38, 289, 1, 0, 0,
		0, 40, 291, 1, 0, 0, 0, 42, 297, 1, 0, 0, 0, 44, 299, 1, 0, 0, 0, 46, 305,
		1, 0, 0, 0, 48, 310, 1, 0, 0, 0, 50, 316, 1, 0, 0, 0, 52, 320, 1, 0, 0,
		0, 54, 350, 1, 0, 0, 0, 56, 362, 1, 0, 0, 0, 58, 364, 1, 0, 0, 0, 60, 366,
		1, 0, 0, 0, 62, 371, 1, 0, 0, 0, 64, 373, 1, 0, 0, 0, 66, 376, 1, 0, 0,
		0, 68, 379, 1, 0, 0, 0, 70, 382, 1, 0, 0, 0, 72, 411, 1, 0, 0, 0, 74, 421,
		1, 0, 0, 0, 76, 77, 5, 29, 0, 0, 77, 78, 3, 56, 28, 0, 78, 79, 5, 44, 0,
		0, 79, 1, 1, 0, 0, 0, 80, 83, 3, 6, 3, 0, 81, 83, 3, 70, 35, 0, 82, 80,
		1, 0, 0, 0, 82, 81, 1, 0, 0, 0, 83, 3, 1, 0, 0, 0, 84, 86, 5, 34, 0, 0,
		85, 84, 1, 0, 0, 0, 86, 89, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 87, 88, 1,
		0, 0, 0, 88, 90, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 90, 94, 3, 0, 0, 0, 91,
		93, 5, 34, 0, 0, 92, 91, 1, 0, 0, 0, 93, 96, 1, 0, 0, 0, 94, 92, 1, 0,
		0, 0, 94, 95, 1, 0, 0, 0, 95, 105, 1, 0, 0, 0, 96, 94, 1, 0, 0, 0, 97,
		99, 3, 2, 1, 0, 98, 100, 5, 34, 0, 0, 99, 98, 1, 0, 0, 0, 100, 101, 1,
		0, 0, 0, 101, 99, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 104, 1, 0, 0,
		0, 103, 97, 1, 0, 0, 0, 104, 107, 1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 105,
		106, 1, 0, 0, 0, 106, 109, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 108, 110,
		3, 2, 1, 0, 109, 108, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 114, 1, 0,
		0, 0, 111, 113, 5, 34, 0, 0, 112, 111, 1, 0, 0, 0, 113, 116, 1, 0, 0, 0,
		114, 112, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 117, 1, 0, 0, 0, 116,
		114, 1, 0, 0, 0, 117, 118, 5, 0, 0, 1, 118, 5, 1, 0, 0, 0, 119, 120, 5,
		3, 0, 0, 120, 121, 3, 8, 4, 0, 121, 7, 1, 0, 0, 0, 122, 126, 5, 37, 0,
		0, 123, 125, 5, 34, 0, 0, 124, 123, 1, 0, 0, 0, 125, 128, 1, 0, 0, 0, 126,
		124, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 141, 1, 0, 0, 0, 128, 126,
		1, 0, 0, 0, 129, 132, 3, 12, 6, 0, 130, 132, 3, 26, 13, 0, 131, 129, 1,
		0, 0, 0, 131, 130, 1, 0, 0, 0, 132, 136, 1, 0, 0, 0, 133, 135, 5, 34, 0,
		0, 134, 133, 1, 0, 0, 0, 135, 138, 1, 0, 0, 0, 136, 134, 1, 0, 0, 0, 136,
		137, 1, 0, 0, 0, 137, 140, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 139, 131,
		1, 0, 0, 0, 140, 143, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 141, 142, 1, 0,
		0, 0, 142, 144, 1, 0, 0, 0, 143, 141, 1, 0, 0, 0, 144, 145, 5, 38, 0, 0,
		145, 9, 1, 0, 0, 0, 146, 149, 5, 46, 0, 0, 147, 149, 3, 54, 27, 0, 148,
		146, 1, 0, 0, 0, 148, 147, 1, 0, 0, 0, 149, 11, 1, 0, 0, 0, 150, 151, 5,
		4, 0, 0, 151, 155, 5, 46, 0, 0, 152, 154, 5, 34, 0, 0, 153, 152, 1, 0,
		0, 0, 154, 157, 1, 0, 0, 0, 155, 153, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0,
		156, 158, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0, 158, 159, 3, 14, 7, 0, 159,
		13, 1, 0, 0, 0, 160, 166, 3, 16, 8, 0, 161, 166, 3, 18, 9, 0, 162, 166,
		3, 20, 10, 0, 163, 166, 3, 22, 11, 0, 164, 166, 3, 10, 5, 0, 165, 160,
		1, 0, 0, 0, 165, 161, 1, 0, 0, 0, 165, 162, 1, 0, 0, 0, 165, 163, 1, 0,
		0, 0, 165, 164, 1, 0, 0, 0, 166, 15, 1, 0, 0, 0, 167, 171, 5, 30, 0, 0,
		168, 170, 5, 34, 0, 0, 169, 168, 1, 0, 0, 0, 170, 173, 1, 0, 0, 0, 171,
		169, 1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172, 175, 1, 0, 0, 0, 173, 171,
		1, 0, 0, 0, 174, 176, 5, 41, 0, 0, 175, 174, 1, 0, 0, 0, 175, 176, 1, 0,
		0, 0, 176, 177, 1, 0, 0, 0, 177, 181, 5, 39, 0, 0, 178, 180, 5, 34, 0,
		0, 179, 178, 1, 0, 0, 0, 180, 183, 1, 0, 0, 0, 181, 179, 1, 0, 0, 0, 181,
		182, 1, 0, 0, 0, 182, 184, 1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 184, 188,
		3, 28, 14, 0, 185, 187, 5, 34, 0, 0, 186, 185, 1, 0, 0, 0, 187, 190, 1,
		0, 0, 0, 188, 186, 1, 0, 0, 0, 188, 189, 1, 0, 0, 0, 189, 191, 1, 0, 0,
		0, 190, 188, 1, 0, 0, 0, 191, 195, 5, 40, 0, 0, 192, 194, 5, 34, 0, 0,
		193, 192, 1, 0, 0, 0, 194, 197, 1, 0, 0, 0, 195, 193, 1, 0, 0, 0, 195,
		196, 1, 0, 0, 0, 196, 198, 1, 0, 0, 0, 197, 195, 1, 0, 0, 0, 198, 199,
		3, 14, 7, 0, 199, 17, 1, 0, 0, 0, 200, 201, 5, 8, 0, 0, 201, 19, 1, 0,
		0, 0, 202, 203, 5, 7, 0, 0, 203, 204, 3, 56, 28, 0, 204, 21, 1, 0, 0, 0,
		205, 206, 5, 6, 0, 0, 206, 210, 5, 35, 0, 0, 207, 209, 5, 34, 0, 0, 208,
		207, 1, 0, 0, 0, 209, 212, 1, 0, 0, 0, 210, 208, 1, 0, 0, 0, 210, 211,
		1, 0, 0, 0, 211, 222, 1, 0, 0, 0, 212, 210, 1, 0, 0, 0, 213, 217, 3, 14,
		7, 0, 214, 216, 5, 34, 0, 0, 215, 214, 1, 0, 0, 0, 216, 219, 1, 0, 0, 0,
		217, 215, 1, 0, 0, 0, 217, 218, 1, 0, 0, 0, 218, 221, 1, 0, 0, 0, 219,
		217, 1, 0, 0, 0, 220, 213, 1, 0, 0, 0, 221, 224, 1, 0, 0, 0, 222, 220,
		1, 0, 0, 0, 222, 223, 1, 0, 0, 0, 223, 225, 1, 0, 0, 0, 224, 222, 1, 0,
		0, 0, 225, 226, 5, 36, 0, 0, 226, 23, 1, 0, 0, 0, 227, 230, 5, 46, 0, 0,
		228, 230, 3, 54, 27, 0, 229, 227, 1, 0, 0, 0, 229, 228, 1, 0, 0, 0, 230,
		25, 1, 0, 0, 0, 231, 232, 5, 5, 0, 0, 232, 236, 5, 46, 0, 0, 233, 235,
		5, 34, 0, 0, 234, 233, 1, 0, 0, 0, 235, 238, 1, 0, 0, 0, 236, 234, 1, 0,
		0, 0, 236, 237, 1, 0, 0, 0, 237, 239, 1, 0, 0, 0, 238, 236, 1, 0, 0, 0,
		239, 240, 3, 28, 14, 0, 240, 27, 1, 0, 0, 0, 241, 254, 3, 30, 15, 0, 242,
		254, 3, 32, 16, 0, 243, 254, 3, 34, 17, 0, 244, 254, 3, 36, 18, 0, 245,
		254, 3, 38, 19, 0, 246, 254, 3, 40, 20, 0, 247, 254, 3, 42, 21, 0, 248,
		254, 3, 44, 22, 0, 249, 254, 3, 46, 23, 0, 250, 254, 3, 48, 24, 0, 251,
		254, 3, 52, 26, 0, 252, 254, 3, 24, 12, 0, 253, 241, 1, 0, 0, 0, 253, 242,
		1, 0, 0, 0, 253, 243, 1, 0, 0, 0, 253, 244, 1, 0, 0, 0, 253, 245, 1, 0,
		0, 0, 253, 246, 1, 0, 0, 0, 253, 247, 1, 0, 0, 0, 253, 248, 1, 0, 0, 0,
		253, 249, 1, 0, 0, 0, 253, 250, 1, 0, 0, 0, 253, 251, 1, 0, 0, 0, 253,
		252, 1, 0, 0, 0, 254, 29, 1, 0, 0, 0, 255, 256, 5, 9, 0, 0, 256, 31, 1,
		0, 0, 0, 257, 258, 5, 10, 0, 0, 258, 262, 5, 35, 0, 0, 259, 261, 5, 34,
		0, 0, 260, 259, 1, 0, 0, 0, 261, 264, 1, 0, 0, 0, 262, 260, 1, 0, 0, 0,
		262, 263, 1, 0, 0, 0, 263, 274, 1, 0, 0, 0, 264, 262, 1, 0, 0, 0, 265,
		269, 3, 56, 28, 0, 266, 268, 5, 34, 0, 0, 267, 266, 1, 0, 0, 0, 268, 271,
		1, 0, 0, 0, 269, 267, 1, 0, 0, 0, 269, 270, 1, 0, 0, 0, 270, 273, 1, 0,
		0, 0, 271, 269, 1, 0, 0, 0, 272, 265, 1, 0, 0, 0, 273, 276, 1, 0, 0, 0,
		274, 272, 1, 0, 0, 0, 274, 275, 1, 0, 0, 0, 275, 277, 1, 0, 0, 0, 276,
		274, 1, 0, 0, 0, 277, 278, 5, 36, 0, 0, 278, 33, 1, 0, 0, 0, 279, 280,
		5, 11, 0, 0, 280, 35, 1, 0, 0, 0, 281, 282, 5, 12, 0, 0, 282, 283, 3, 56,
		28, 0, 283, 284, 5, 35, 0, 0, 284, 285, 3, 58, 29, 0, 285, 286, 5, 42,
		0, 0, 286, 287, 3, 58, 29, 0, 287, 288, 5, 36, 0, 0, 288, 37, 1, 0, 0,
		0, 289, 290, 5, 13, 0, 0, 290, 39, 1, 0, 0, 0, 291, 292, 5, 14, 0, 0, 292,
		293, 7, 0, 0, 0, 293, 294, 5, 1, 0, 0, 294, 295, 7, 1, 0, 0, 295, 296,
		5, 1, 0, 0, 296, 41, 1, 0, 0, 0, 297, 298, 5, 15, 0, 0, 298, 43, 1, 0,
		0, 0, 299, 300, 5, 17, 0, 0, 300, 301, 5, 45, 0, 0, 301, 302, 3, 62, 31,
		0, 302, 303, 5, 42, 0, 0, 303, 304, 3, 62, 31, 0, 304, 45, 1, 0, 0, 0,
		305, 306, 5, 18, 0, 0, 306, 307, 5, 1, 0, 0, 307, 308, 3, 58, 29, 0, 308,
		309, 7, 1, 0, 0, 309, 47, 1, 0, 0, 0, 310, 311, 5, 19, 0, 0, 311, 312,
		3, 62, 31, 0, 312, 313, 5, 1, 0, 0, 313, 314, 7, 1, 0, 0, 314, 49, 1, 0,
		0, 0, 315, 317, 5, 41, 0, 0, 316, 315, 1, 0, 0, 0, 316, 317, 1, 0, 0, 0,
		317, 318, 1, 0, 0, 0, 318, 319, 3, 28, 14, 0, 319, 51, 1, 0, 0, 0, 320,
		324, 5, 22, 0, 0, 321, 323, 5, 34, 0, 0, 322, 321, 1, 0, 0, 0, 323, 326,
		1, 0, 0, 0, 324, 322, 1, 0, 0, 0, 324, 325, 1, 0, 0, 0, 325, 327, 1, 0,
		0, 0, 326, 324, 1, 0, 0, 0, 327, 331, 5, 39, 0, 0, 328, 330, 5, 34, 0,
		0, 329, 328, 1, 0, 0, 0, 330, 333, 1, 0, 0, 0, 331, 329, 1, 0, 0, 0, 331,
		332, 1, 0, 0, 0, 332, 343, 1, 0, 0, 0, 333, 331, 1, 0, 0, 0, 334, 338,
		3, 50, 25, 0, 335, 337, 5, 34, 0, 0, 336, 335, 1, 0, 0, 0, 337, 340, 1,
		0, 0, 0, 338, 336, 1, 0, 0, 0, 338, 339, 1, 0, 0, 0, 339, 342, 1, 0, 0,
		0, 340, 338, 1, 0, 0, 0, 341, 334, 1, 0, 0, 0, 342, 345, 1, 0, 0, 0, 343,
		341, 1, 0, 0, 0, 343, 344, 1, 0, 0, 0, 344, 346, 1, 0, 0, 0, 345, 343,
		1, 0, 0, 0, 346, 347, 5, 40, 0, 0, 347, 53, 1, 0, 0, 0, 348, 351, 3, 60,
		30, 0, 349, 351, 5, 46, 0, 0, 350, 348, 1, 0, 0, 0, 350, 349, 1, 0, 0,
		0, 351, 352, 1, 0, 0, 0, 352, 355, 5, 43, 0, 0, 353, 356, 3, 60, 30, 0,
		354, 356, 5, 46, 0, 0, 355, 353, 1, 0, 0, 0, 355, 354, 1, 0, 0, 0, 356,
		55, 1, 0, 0, 0, 357, 363, 3, 54, 27, 0, 358, 361, 3, 60, 30, 0, 359, 361,
		5, 46, 0, 0, 360, 358, 1, 0, 0, 0, 360, 359, 1, 0, 0, 0, 361, 363, 1, 0,
		0, 0, 362, 357, 1, 0, 0, 0, 362, 360, 1, 0, 0, 0, 363, 57, 1, 0, 0, 0,
		364, 365, 7, 2, 0, 0, 365, 59, 1, 0, 0, 0, 366, 367, 7, 3, 0, 0, 367, 61,
		1, 0, 0, 0, 368, 372, 3, 64, 32, 0, 369, 372, 3, 66, 33, 0, 370, 372, 3,
		68, 34, 0, 371, 368, 1, 0, 0, 0, 371, 369, 1, 0, 0, 0, 371, 370, 1, 0,
		0, 0, 372, 63, 1, 0, 0, 0, 373, 374, 5, 26, 0, 0, 374, 375, 5, 1, 0, 0,
		375, 65, 1, 0, 0, 0, 376, 377, 5, 27, 0, 0, 377, 378, 5, 1, 0, 0, 378,
		67, 1, 0, 0, 0, 379, 380, 5, 28, 0, 0, 380, 381, 5, 1, 0, 0, 381, 69, 1,
		0, 0, 0, 382, 383, 5, 25, 0, 0, 383, 387, 5, 37, 0, 0, 384, 386, 5, 34,
		0, 0, 385, 384, 1, 0, 0, 0, 386, 389, 1, 0, 0, 0, 387, 385, 1, 0, 0, 0,
		387, 388, 1, 0, 0, 0, 388, 399, 1, 0, 0, 0, 389, 387, 1, 0, 0, 0, 390,
		394, 3, 72, 36, 0, 391, 393, 5, 34, 0, 0, 392, 391, 1, 0, 0, 0, 393, 396,
		1, 0, 0, 0, 394, 392, 1, 0, 0, 0, 394, 395, 1, 0, 0, 0, 395, 398, 1, 0,
		0, 0, 396, 394, 1, 0, 0, 0, 397, 390, 1, 0, 0, 0, 398, 401, 1, 0, 0, 0,
		399, 397, 1, 0, 0, 0, 399, 400, 1, 0, 0, 0, 400, 402, 1, 0, 0, 0, 401,
		399, 1, 0, 0, 0, 402, 406, 3, 72, 36, 0, 403, 405, 5, 34, 0, 0, 404, 403,
		1, 0, 0, 0, 405, 408, 1, 0, 0, 0, 406, 404, 1, 0, 0, 0, 406, 407, 1, 0,
		0, 0, 407, 409, 1, 0, 0, 0, 408, 406, 1, 0, 0, 0, 409, 410, 5, 38, 0, 0,
		410, 71, 1, 0, 0, 0, 411, 412, 5, 12, 0, 0, 412, 416, 5, 46, 0, 0, 413,
		415, 5, 34, 0, 0, 414, 413, 1, 0, 0, 0, 415, 418, 1, 0, 0, 0, 416, 414,
		1, 0, 0, 0, 416, 417, 1, 0, 0, 0, 417, 419, 1, 0, 0, 0, 418, 416, 1, 0,
		0, 0, 419, 420, 3, 74, 37, 0, 420, 73, 1, 0, 0, 0, 421, 425, 3, 58, 29,
		0, 422, 424, 5, 34, 0, 0, 423, 422, 1, 0, 0, 0, 424, 427, 1, 0, 0, 0, 425,
		423, 1, 0, 0, 0, 425, 426, 1, 0, 0, 0, 426, 448, 1, 0, 0, 0, 427, 425,
		1, 0, 0, 0, 428, 432, 5, 35, 0, 0, 429, 431, 5, 34, 0, 0, 430, 429, 1,
		0, 0, 0, 431, 434, 1, 0, 0, 0, 432, 430, 1, 0, 0, 0, 432, 433, 1, 0, 0,
		0, 433, 442, 1, 0, 0, 0, 434, 432, 1, 0, 0, 0, 435, 439, 3, 58, 29, 0,
		436, 438, 5, 34, 0, 0, 437, 436, 1, 0, 0, 0, 438, 441, 1, 0, 0, 0, 439,
		437, 1, 0, 0, 0, 439, 440, 1, 0, 0, 0, 440, 443, 1, 0, 0, 0, 441, 439,
		1, 0, 0, 0, 442, 435, 1, 0, 0, 0, 443, 444, 1, 0, 0, 0, 444, 442, 1, 0,
		0, 0, 444, 445, 1, 0, 0, 0, 445, 446, 1, 0, 0, 0, 446, 447, 5, 36, 0, 0,
		447, 449, 1, 0, 0, 0, 448, 428, 1, 0, 0, 0, 448, 449, 1, 0, 0, 0, 449,
		75, 1, 0, 0, 0, 48, 82, 87, 94, 101, 105, 109, 114, 126, 131, 136, 141,
		148, 155, 165, 171, 175, 181, 188, 195, 210, 217, 222, 229, 236, 253, 262,
		269, 274, 316, 324, 331, 338, 343, 350, 355, 360, 362, 371, 387, 394, 399,
		406, 416, 425, 432, 439, 444, 448,
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
	MMSParserEOF                             = antlr.TokenEOF
	MMSParserInt                             = 1
	MMSParserFloat                           = 2
	MMSParserKeyword_Surface                 = 3
	MMSParserKeyword_Rule                    = 4
	MMSParserKeyword_Condition               = 5
	MMSParserKeyword_Sequence                = 6
	MMSParserKeyword_Block                   = 7
	MMSParserKeyword_Bandlands               = 8
	MMSParserKeyword_AbovePreliminarySurface = 9
	MMSParserKeyword_Biome                   = 10
	MMSParserKeyword_Hole                    = 11
	MMSParserKeyword_Noise                   = 12
	MMSParserKeyword_Steep                   = 13
	MMSParserKeyword_StoneDepth              = 14
	MMSParserKeyword_Freezing                = 15
	MMSParserKeyword_Temperature             = 16
	MMSParserKeyword_VerticalGradient        = 17
	MMSParserKeyword_AboveWater              = 18
	MMSParserKeyword_YAbove                  = 19
	MMSParserKeyword_Floor                   = 20
	MMSParserKeyword_Ceiling                 = 21
	MMSParserKeyword_And                     = 22
	MMSParserKeyword_Add                     = 23
	MMSParserKeyword_Sub                     = 24
	MMSParserKeyword_WorldGen                = 25
	MMSParserKeyword_Absolute                = 26
	MMSParserKeyword_AboveBottom             = 27
	MMSParserKeyword_BelowTop                = 28
	MMSParserKeyword_Namespace               = 29
	MMSParserKeyword_If                      = 30
	MMSParserKeyword_Else                    = 31
	MMSParserKeyword_In                      = 32
	MMSParserWS                              = 33
	MMSParserNL                              = 34
	MMSParserSquareOpen                      = 35
	MMSParserSquareClose                     = 36
	MMSParserCurlyOpen                       = 37
	MMSParserCurlyClose                      = 38
	MMSParserRoundOpen                       = 39
	MMSParserRoundClose                      = 40
	MMSParserBang                            = 41
	MMSParserComma                           = 42
	MMSParserColon                           = 43
	MMSParserSemiColon                       = 44
	MMSParserString_                         = 45
	MMSParserIdentifier                      = 46
	MMSParserLineComment                     = 47
	MMSParserBlockComment                    = 48
)

// MMSParser rules.
const (
	MMSParserRULE_namespaceDeclaration              = 0
	MMSParserRULE_statement                         = 1
	MMSParserRULE_mmsFile                           = 2
	MMSParserRULE_surfaceDeclaration                = 3
	MMSParserRULE_surfaceDefinition                 = 4
	MMSParserRULE_surfaceRuleReference              = 5
	MMSParserRULE_surfaceRuleDeclaration            = 6
	MMSParserRULE_surfaceRule                       = 7
	MMSParserRULE_surfaceRule_Conditional           = 8
	MMSParserRULE_surfaceRule_Bandlands             = 9
	MMSParserRULE_surfaceRule_Block                 = 10
	MMSParserRULE_surfaceRule_Sequence              = 11
	MMSParserRULE_surfaceConditionReference         = 12
	MMSParserRULE_surfaceConditionDeclaration       = 13
	MMSParserRULE_surfaceCondition                  = 14
	MMSParserRULE_surfaceCondition_AboveSurface     = 15
	MMSParserRULE_surfaceCondition_Biome            = 16
	MMSParserRULE_surfaceCondition_Hole             = 17
	MMSParserRULE_surfaceCondition_Noise            = 18
	MMSParserRULE_surfaceCondition_Steep            = 19
	MMSParserRULE_surfaceCondition_StoneDepth       = 20
	MMSParserRULE_surfaceCondition_Freezing         = 21
	MMSParserRULE_surfaceCondition_VerticalGradient = 22
	MMSParserRULE_surfaceCondition_AboveWater       = 23
	MMSParserRULE_surfaceCondition_YAbove           = 24
	MMSParserRULE_surfaceCondition_Compound__Item   = 25
	MMSParserRULE_surfaceCondition_Compound         = 26
	MMSParserRULE_reference                         = 27
	MMSParserRULE_resourceReference                 = 28
	MMSParserRULE_number                            = 29
	MMSParserRULE_keyword                           = 30
	MMSParserRULE_verticalAnchor                    = 31
	MMSParserRULE_verticalAnchor_Absolute           = 32
	MMSParserRULE_verticalAnchor_AboveBottom        = 33
	MMSParserRULE_verticalAnchor_BelowTop           = 34
	MMSParserRULE_worldgenDeclaration               = 35
	MMSParserRULE_noiseDeclaration                  = 36
	MMSParserRULE_noiseDefinition                   = 37
)

// INamespaceDeclarationContext is an interface to support dynamic dispatch.
type INamespaceDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Keyword_Namespace() antlr.TerminalNode
	ResourceReference() IResourceReferenceContext
	SemiColon() antlr.TerminalNode

	// IsNamespaceDeclarationContext differentiates from other interfaces.
	IsNamespaceDeclarationContext()
}

type NamespaceDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamespaceDeclarationContext() *NamespaceDeclarationContext {
	var p = new(NamespaceDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_namespaceDeclaration
	return p
}

func InitEmptyNamespaceDeclarationContext(p *NamespaceDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_namespaceDeclaration
}

func (*NamespaceDeclarationContext) IsNamespaceDeclarationContext() {}

func NewNamespaceDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamespaceDeclarationContext {
	var p = new(NamespaceDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_namespaceDeclaration

	return p
}

func (s *NamespaceDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *NamespaceDeclarationContext) Keyword_Namespace() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_Namespace, 0)
}

func (s *NamespaceDeclarationContext) ResourceReference() IResourceReferenceContext {
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

func (s *NamespaceDeclarationContext) SemiColon() antlr.TerminalNode {
	return s.GetToken(MMSParserSemiColon, 0)
}

func (s *NamespaceDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamespaceDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamespaceDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.EnterNamespaceDeclaration(s)
	}
}

func (s *NamespaceDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.ExitNamespaceDeclaration(s)
	}
}

func (p *MMSParser) NamespaceDeclaration() (localctx INamespaceDeclarationContext) {
	localctx = NewNamespaceDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MMSParserRULE_namespaceDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(76)
		p.Match(MMSParserKeyword_Namespace)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(77)
		p.ResourceReference()
	}
	{
		p.SetState(78)
		p.Match(MMSParserSemiColon)
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

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SurfaceDeclaration() ISurfaceDeclarationContext
	WorldgenDeclaration() IWorldgenDeclarationContext

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

func (s *StatementContext) SurfaceDeclaration() ISurfaceDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISurfaceDeclarationContext)
}

func (s *StatementContext) WorldgenDeclaration() IWorldgenDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWorldgenDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWorldgenDeclarationContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *MMSParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MMSParserRULE_statement)
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MMSParserKeyword_Surface:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(80)
			p.SurfaceDeclaration()
		}

	case MMSParserKeyword_WorldGen:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(81)
			p.WorldgenDeclaration()
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

// IMmsFileContext is an interface to support dynamic dispatch.
type IMmsFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NamespaceDeclaration() INamespaceDeclarationContext
	EOF() antlr.TerminalNode
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsMmsFileContext differentiates from other interfaces.
	IsMmsFileContext()
}

type MmsFileContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMmsFileContext() *MmsFileContext {
	var p = new(MmsFileContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_mmsFile
	return p
}

func InitEmptyMmsFileContext(p *MmsFileContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_mmsFile
}

func (*MmsFileContext) IsMmsFileContext() {}

func NewMmsFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MmsFileContext {
	var p = new(MmsFileContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_mmsFile

	return p
}

func (s *MmsFileContext) GetParser() antlr.Parser { return s.parser }

func (s *MmsFileContext) NamespaceDeclaration() INamespaceDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamespaceDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamespaceDeclarationContext)
}

func (s *MmsFileContext) EOF() antlr.TerminalNode {
	return s.GetToken(MMSParserEOF, 0)
}

func (s *MmsFileContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(MMSParserNL)
}

func (s *MmsFileContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(MMSParserNL, i)
}

func (s *MmsFileContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *MmsFileContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *MmsFileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MmsFileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MmsFileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.EnterMmsFile(s)
	}
}

func (s *MmsFileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.ExitMmsFile(s)
	}
}

func (p *MMSParser) MmsFile() (localctx IMmsFileContext) {
	localctx = NewMmsFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MMSParserRULE_mmsFile)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MMSParserNL {
		{
			p.SetState(84)
			p.Match(MMSParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(90)
		p.NamespaceDeclaration()
	}
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(91)
				p.Match(MMSParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(97)
				p.Statement()
			}
			p.SetState(99)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
					{
						p.SetState(98)
						p.Match(MMSParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				default:
					p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
					goto errorExit
				}

				p.SetState(101)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
				if p.HasError() {
					goto errorExit
				}
			}

		}
		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MMSParserKeyword_Surface || _la == MMSParserKeyword_WorldGen {
		{
			p.SetState(108)
			p.Statement()
		}

	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MMSParserNL {
		{
			p.SetState(111)
			p.Match(MMSParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(117)
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

// ISurfaceDeclarationContext is an interface to support dynamic dispatch.
type ISurfaceDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Keyword_Surface() antlr.TerminalNode
	SurfaceDefinition() ISurfaceDefinitionContext

	// IsSurfaceDeclarationContext differentiates from other interfaces.
	IsSurfaceDeclarationContext()
}

type SurfaceDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceDeclarationContext() *SurfaceDeclarationContext {
	var p = new(SurfaceDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_surfaceDeclaration
	return p
}

func InitEmptySurfaceDeclarationContext(p *SurfaceDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_surfaceDeclaration
}

func (*SurfaceDeclarationContext) IsSurfaceDeclarationContext() {}

func NewSurfaceDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceDeclarationContext {
	var p = new(SurfaceDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_surfaceDeclaration

	return p
}

func (s *SurfaceDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceDeclarationContext) Keyword_Surface() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_Surface, 0)
}

func (s *SurfaceDeclarationContext) SurfaceDefinition() ISurfaceDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISurfaceDefinitionContext)
}

func (s *SurfaceDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.EnterSurfaceDeclaration(s)
	}
}

func (s *SurfaceDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.ExitSurfaceDeclaration(s)
	}
}

func (p *MMSParser) SurfaceDeclaration() (localctx ISurfaceDeclarationContext) {
	localctx = NewSurfaceDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MMSParserRULE_surfaceDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.Match(MMSParserKeyword_Surface)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(120)
		p.SurfaceDefinition()
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

// ISurfaceDefinitionContext is an interface to support dynamic dispatch.
type ISurfaceDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CurlyOpen() antlr.TerminalNode
	CurlyClose() antlr.TerminalNode
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode
	AllSurfaceRuleDeclaration() []ISurfaceRuleDeclarationContext
	SurfaceRuleDeclaration(i int) ISurfaceRuleDeclarationContext
	AllSurfaceConditionDeclaration() []ISurfaceConditionDeclarationContext
	SurfaceConditionDeclaration(i int) ISurfaceConditionDeclarationContext

	// IsSurfaceDefinitionContext differentiates from other interfaces.
	IsSurfaceDefinitionContext()
}

type SurfaceDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceDefinitionContext() *SurfaceDefinitionContext {
	var p = new(SurfaceDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_surfaceDefinition
	return p
}

func InitEmptySurfaceDefinitionContext(p *SurfaceDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_surfaceDefinition
}

func (*SurfaceDefinitionContext) IsSurfaceDefinitionContext() {}

func NewSurfaceDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceDefinitionContext {
	var p = new(SurfaceDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_surfaceDefinition

	return p
}

func (s *SurfaceDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceDefinitionContext) CurlyOpen() antlr.TerminalNode {
	return s.GetToken(MMSParserCurlyOpen, 0)
}

func (s *SurfaceDefinitionContext) CurlyClose() antlr.TerminalNode {
	return s.GetToken(MMSParserCurlyClose, 0)
}

func (s *SurfaceDefinitionContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(MMSParserNL)
}

func (s *SurfaceDefinitionContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(MMSParserNL, i)
}

func (s *SurfaceDefinitionContext) AllSurfaceRuleDeclaration() []ISurfaceRuleDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISurfaceRuleDeclarationContext); ok {
			len++
		}
	}

	tst := make([]ISurfaceRuleDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISurfaceRuleDeclarationContext); ok {
			tst[i] = t.(ISurfaceRuleDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *SurfaceDefinitionContext) SurfaceRuleDeclaration(i int) ISurfaceRuleDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceRuleDeclarationContext); ok {
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

	return t.(ISurfaceRuleDeclarationContext)
}

func (s *SurfaceDefinitionContext) AllSurfaceConditionDeclaration() []ISurfaceConditionDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISurfaceConditionDeclarationContext); ok {
			len++
		}
	}

	tst := make([]ISurfaceConditionDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISurfaceConditionDeclarationContext); ok {
			tst[i] = t.(ISurfaceConditionDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *SurfaceDefinitionContext) SurfaceConditionDeclaration(i int) ISurfaceConditionDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceConditionDeclarationContext); ok {
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

	return t.(ISurfaceConditionDeclarationContext)
}

func (s *SurfaceDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.EnterSurfaceDefinition(s)
	}
}

func (s *SurfaceDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.ExitSurfaceDefinition(s)
	}
}

func (p *MMSParser) SurfaceDefinition() (localctx ISurfaceDefinitionContext) {
	localctx = NewSurfaceDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MMSParserRULE_surfaceDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(122)
		p.Match(MMSParserCurlyOpen)
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

	for _la == MMSParserNL {
		{
			p.SetState(123)
			p.Match(MMSParserNL)
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
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MMSParserKeyword_Rule || _la == MMSParserKeyword_Condition {
		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case MMSParserKeyword_Rule:
			{
				p.SetState(129)
				p.SurfaceRuleDeclaration()
			}

		case MMSParserKeyword_Condition:
			{
				p.SetState(130)
				p.SurfaceConditionDeclaration()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}
		p.SetState(136)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MMSParserNL {
			{
				p.SetState(133)
				p.Match(MMSParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(138)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(143)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(144)
		p.Match(MMSParserCurlyClose)
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

// ISurfaceRuleReferenceContext is an interface to support dynamic dispatch.
type ISurfaceRuleReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	Reference() IReferenceContext

	// IsSurfaceRuleReferenceContext differentiates from other interfaces.
	IsSurfaceRuleReferenceContext()
}

type SurfaceRuleReferenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceRuleReferenceContext() *SurfaceRuleReferenceContext {
	var p = new(SurfaceRuleReferenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_surfaceRuleReference
	return p
}

func InitEmptySurfaceRuleReferenceContext(p *SurfaceRuleReferenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_surfaceRuleReference
}

func (*SurfaceRuleReferenceContext) IsSurfaceRuleReferenceContext() {}

func NewSurfaceRuleReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceRuleReferenceContext {
	var p = new(SurfaceRuleReferenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_surfaceRuleReference

	return p
}

func (s *SurfaceRuleReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceRuleReferenceContext) Identifier() antlr.TerminalNode {
	return s.GetToken(MMSParserIdentifier, 0)
}

func (s *SurfaceRuleReferenceContext) Reference() IReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReferenceContext)
}

func (s *SurfaceRuleReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceRuleReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceRuleReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.EnterSurfaceRuleReference(s)
	}
}

func (s *SurfaceRuleReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.ExitSurfaceRuleReference(s)
	}
}

func (p *MMSParser) SurfaceRuleReference() (localctx ISurfaceRuleReferenceContext) {
	localctx = NewSurfaceRuleReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MMSParserRULE_surfaceRuleReference)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(146)
			p.Match(MMSParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(147)
			p.Reference()
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

// ISurfaceRuleDeclarationContext is an interface to support dynamic dispatch.
type ISurfaceRuleDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Keyword_Rule() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	SurfaceRule() ISurfaceRuleContext
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode

	// IsSurfaceRuleDeclarationContext differentiates from other interfaces.
	IsSurfaceRuleDeclarationContext()
}

type SurfaceRuleDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceRuleDeclarationContext() *SurfaceRuleDeclarationContext {
	var p = new(SurfaceRuleDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_surfaceRuleDeclaration
	return p
}

func InitEmptySurfaceRuleDeclarationContext(p *SurfaceRuleDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_surfaceRuleDeclaration
}

func (*SurfaceRuleDeclarationContext) IsSurfaceRuleDeclarationContext() {}

func NewSurfaceRuleDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceRuleDeclarationContext {
	var p = new(SurfaceRuleDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_surfaceRuleDeclaration

	return p
}

func (s *SurfaceRuleDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceRuleDeclarationContext) Keyword_Rule() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_Rule, 0)
}

func (s *SurfaceRuleDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(MMSParserIdentifier, 0)
}

func (s *SurfaceRuleDeclarationContext) SurfaceRule() ISurfaceRuleContext {
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

func (s *SurfaceRuleDeclarationContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(MMSParserNL)
}

func (s *SurfaceRuleDeclarationContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(MMSParserNL, i)
}

func (s *SurfaceRuleDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceRuleDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceRuleDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.EnterSurfaceRuleDeclaration(s)
	}
}

func (s *SurfaceRuleDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.ExitSurfaceRuleDeclaration(s)
	}
}

func (p *MMSParser) SurfaceRuleDeclaration() (localctx ISurfaceRuleDeclarationContext) {
	localctx = NewSurfaceRuleDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MMSParserRULE_surfaceRuleDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(150)
		p.Match(MMSParserKeyword_Rule)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(151)
		p.Match(MMSParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MMSParserNL {
		{
			p.SetState(152)
			p.Match(MMSParserNL)
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

// ISurfaceRuleContext is an interface to support dynamic dispatch.
type ISurfaceRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SurfaceRule_Conditional() ISurfaceRule_ConditionalContext
	SurfaceRule_Bandlands() ISurfaceRule_BandlandsContext
	SurfaceRule_Block() ISurfaceRule_BlockContext
	SurfaceRule_Sequence() ISurfaceRule_SequenceContext
	SurfaceRuleReference() ISurfaceRuleReferenceContext

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
	p.RuleIndex = MMSParserRULE_surfaceRule
	return p
}

func InitEmptySurfaceRuleContext(p *SurfaceRuleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_surfaceRule
}

func (*SurfaceRuleContext) IsSurfaceRuleContext() {}

func NewSurfaceRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceRuleContext {
	var p = new(SurfaceRuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_surfaceRule

	return p
}

func (s *SurfaceRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceRuleContext) SurfaceRule_Conditional() ISurfaceRule_ConditionalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceRule_ConditionalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISurfaceRule_ConditionalContext)
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

func (s *SurfaceRuleContext) SurfaceRuleReference() ISurfaceRuleReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceRuleReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISurfaceRuleReferenceContext)
}

func (s *SurfaceRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.EnterSurfaceRule(s)
	}
}

func (s *SurfaceRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.ExitSurfaceRule(s)
	}
}

func (p *MMSParser) SurfaceRule() (localctx ISurfaceRuleContext) {
	localctx = NewSurfaceRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MMSParserRULE_surfaceRule)
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(160)
			p.SurfaceRule_Conditional()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(161)
			p.SurfaceRule_Bandlands()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(162)
			p.SurfaceRule_Block()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(163)
			p.SurfaceRule_Sequence()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(164)
			p.SurfaceRuleReference()
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

// ISurfaceRule_ConditionalContext is an interface to support dynamic dispatch.
type ISurfaceRule_ConditionalContext interface {
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
	Bang() antlr.TerminalNode

	// IsSurfaceRule_ConditionalContext differentiates from other interfaces.
	IsSurfaceRule_ConditionalContext()
}

type SurfaceRule_ConditionalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceRule_ConditionalContext() *SurfaceRule_ConditionalContext {
	var p = new(SurfaceRule_ConditionalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_surfaceRule_Conditional
	return p
}

func InitEmptySurfaceRule_ConditionalContext(p *SurfaceRule_ConditionalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_surfaceRule_Conditional
}

func (*SurfaceRule_ConditionalContext) IsSurfaceRule_ConditionalContext() {}

func NewSurfaceRule_ConditionalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceRule_ConditionalContext {
	var p = new(SurfaceRule_ConditionalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_surfaceRule_Conditional

	return p
}

func (s *SurfaceRule_ConditionalContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceRule_ConditionalContext) Keyword_If() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_If, 0)
}

func (s *SurfaceRule_ConditionalContext) RoundOpen() antlr.TerminalNode {
	return s.GetToken(MMSParserRoundOpen, 0)
}

func (s *SurfaceRule_ConditionalContext) SurfaceCondition() ISurfaceConditionContext {
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

func (s *SurfaceRule_ConditionalContext) RoundClose() antlr.TerminalNode {
	return s.GetToken(MMSParserRoundClose, 0)
}

func (s *SurfaceRule_ConditionalContext) SurfaceRule() ISurfaceRuleContext {
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

func (s *SurfaceRule_ConditionalContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(MMSParserNL)
}

func (s *SurfaceRule_ConditionalContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(MMSParserNL, i)
}

func (s *SurfaceRule_ConditionalContext) Bang() antlr.TerminalNode {
	return s.GetToken(MMSParserBang, 0)
}

func (s *SurfaceRule_ConditionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceRule_ConditionalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceRule_ConditionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.EnterSurfaceRule_Conditional(s)
	}
}

func (s *SurfaceRule_ConditionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.ExitSurfaceRule_Conditional(s)
	}
}

func (p *MMSParser) SurfaceRule_Conditional() (localctx ISurfaceRule_ConditionalContext) {
	localctx = NewSurfaceRule_ConditionalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MMSParserRULE_surfaceRule_Conditional)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)
		p.Match(MMSParserKeyword_If)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MMSParserNL {
		{
			p.SetState(168)
			p.Match(MMSParserNL)
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
	}
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MMSParserBang {
		{
			p.SetState(174)
			p.Match(MMSParserBang)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(177)
		p.Match(MMSParserRoundOpen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MMSParserNL {
		{
			p.SetState(178)
			p.Match(MMSParserNL)
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
		p.SurfaceCondition()
	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MMSParserNL {
		{
			p.SetState(185)
			p.Match(MMSParserNL)
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
		p.Match(MMSParserRoundClose)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MMSParserNL {
		{
			p.SetState(192)
			p.Match(MMSParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(197)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

	{
		p.SetState(198)
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
	p.RuleIndex = MMSParserRULE_surfaceRule_Bandlands
	return p
}

func InitEmptySurfaceRule_BandlandsContext(p *SurfaceRule_BandlandsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_surfaceRule_Bandlands
}

func (*SurfaceRule_BandlandsContext) IsSurfaceRule_BandlandsContext() {}

func NewSurfaceRule_BandlandsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceRule_BandlandsContext {
	var p = new(SurfaceRule_BandlandsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_surfaceRule_Bandlands

	return p
}

func (s *SurfaceRule_BandlandsContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceRule_BandlandsContext) Keyword_Bandlands() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_Bandlands, 0)
}

func (s *SurfaceRule_BandlandsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceRule_BandlandsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceRule_BandlandsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.EnterSurfaceRule_Bandlands(s)
	}
}

func (s *SurfaceRule_BandlandsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.ExitSurfaceRule_Bandlands(s)
	}
}

func (p *MMSParser) SurfaceRule_Bandlands() (localctx ISurfaceRule_BandlandsContext) {
	localctx = NewSurfaceRule_BandlandsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MMSParserRULE_surfaceRule_Bandlands)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(200)
		p.Match(MMSParserKeyword_Bandlands)
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
	p.RuleIndex = MMSParserRULE_surfaceRule_Block
	return p
}

func InitEmptySurfaceRule_BlockContext(p *SurfaceRule_BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_surfaceRule_Block
}

func (*SurfaceRule_BlockContext) IsSurfaceRule_BlockContext() {}

func NewSurfaceRule_BlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceRule_BlockContext {
	var p = new(SurfaceRule_BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_surfaceRule_Block

	return p
}

func (s *SurfaceRule_BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceRule_BlockContext) Keyword_Block() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_Block, 0)
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
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.EnterSurfaceRule_Block(s)
	}
}

func (s *SurfaceRule_BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.ExitSurfaceRule_Block(s)
	}
}

func (p *MMSParser) SurfaceRule_Block() (localctx ISurfaceRule_BlockContext) {
	localctx = NewSurfaceRule_BlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MMSParserRULE_surfaceRule_Block)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(202)
		p.Match(MMSParserKeyword_Block)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(203)
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
	p.RuleIndex = MMSParserRULE_surfaceRule_Sequence
	return p
}

func InitEmptySurfaceRule_SequenceContext(p *SurfaceRule_SequenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_surfaceRule_Sequence
}

func (*SurfaceRule_SequenceContext) IsSurfaceRule_SequenceContext() {}

func NewSurfaceRule_SequenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceRule_SequenceContext {
	var p = new(SurfaceRule_SequenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_surfaceRule_Sequence

	return p
}

func (s *SurfaceRule_SequenceContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceRule_SequenceContext) Keyword_Sequence() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_Sequence, 0)
}

func (s *SurfaceRule_SequenceContext) SquareOpen() antlr.TerminalNode {
	return s.GetToken(MMSParserSquareOpen, 0)
}

func (s *SurfaceRule_SequenceContext) SquareClose() antlr.TerminalNode {
	return s.GetToken(MMSParserSquareClose, 0)
}

func (s *SurfaceRule_SequenceContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(MMSParserNL)
}

func (s *SurfaceRule_SequenceContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(MMSParserNL, i)
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
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.EnterSurfaceRule_Sequence(s)
	}
}

func (s *SurfaceRule_SequenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.ExitSurfaceRule_Sequence(s)
	}
}

func (p *MMSParser) SurfaceRule_Sequence() (localctx ISurfaceRule_SequenceContext) {
	localctx = NewSurfaceRule_SequenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, MMSParserRULE_surfaceRule_Sequence)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)
		p.Match(MMSParserKeyword_Sequence)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(206)
		p.Match(MMSParserSquareOpen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MMSParserNL {
		{
			p.SetState(207)
			p.Match(MMSParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(212)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&70377300557816) != 0 {
		{
			p.SetState(213)
			p.SurfaceRule()
		}

		p.SetState(217)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MMSParserNL {
			{
				p.SetState(214)
				p.Match(MMSParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(219)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(224)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(225)
		p.Match(MMSParserSquareClose)
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

// ISurfaceConditionReferenceContext is an interface to support dynamic dispatch.
type ISurfaceConditionReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	Reference() IReferenceContext

	// IsSurfaceConditionReferenceContext differentiates from other interfaces.
	IsSurfaceConditionReferenceContext()
}

type SurfaceConditionReferenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceConditionReferenceContext() *SurfaceConditionReferenceContext {
	var p = new(SurfaceConditionReferenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_surfaceConditionReference
	return p
}

func InitEmptySurfaceConditionReferenceContext(p *SurfaceConditionReferenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_surfaceConditionReference
}

func (*SurfaceConditionReferenceContext) IsSurfaceConditionReferenceContext() {}

func NewSurfaceConditionReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceConditionReferenceContext {
	var p = new(SurfaceConditionReferenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_surfaceConditionReference

	return p
}

func (s *SurfaceConditionReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceConditionReferenceContext) Identifier() antlr.TerminalNode {
	return s.GetToken(MMSParserIdentifier, 0)
}

func (s *SurfaceConditionReferenceContext) Reference() IReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReferenceContext)
}

func (s *SurfaceConditionReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceConditionReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceConditionReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.EnterSurfaceConditionReference(s)
	}
}

func (s *SurfaceConditionReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.ExitSurfaceConditionReference(s)
	}
}

func (p *MMSParser) SurfaceConditionReference() (localctx ISurfaceConditionReferenceContext) {
	localctx = NewSurfaceConditionReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, MMSParserRULE_surfaceConditionReference)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(227)
			p.Match(MMSParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(228)
			p.Reference()
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

// ISurfaceConditionDeclarationContext is an interface to support dynamic dispatch.
type ISurfaceConditionDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Keyword_Condition() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	SurfaceCondition() ISurfaceConditionContext
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode

	// IsSurfaceConditionDeclarationContext differentiates from other interfaces.
	IsSurfaceConditionDeclarationContext()
}

type SurfaceConditionDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceConditionDeclarationContext() *SurfaceConditionDeclarationContext {
	var p = new(SurfaceConditionDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_surfaceConditionDeclaration
	return p
}

func InitEmptySurfaceConditionDeclarationContext(p *SurfaceConditionDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_surfaceConditionDeclaration
}

func (*SurfaceConditionDeclarationContext) IsSurfaceConditionDeclarationContext() {}

func NewSurfaceConditionDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceConditionDeclarationContext {
	var p = new(SurfaceConditionDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_surfaceConditionDeclaration

	return p
}

func (s *SurfaceConditionDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceConditionDeclarationContext) Keyword_Condition() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_Condition, 0)
}

func (s *SurfaceConditionDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(MMSParserIdentifier, 0)
}

func (s *SurfaceConditionDeclarationContext) SurfaceCondition() ISurfaceConditionContext {
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

func (s *SurfaceConditionDeclarationContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(MMSParserNL)
}

func (s *SurfaceConditionDeclarationContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(MMSParserNL, i)
}

func (s *SurfaceConditionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceConditionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceConditionDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.EnterSurfaceConditionDeclaration(s)
	}
}

func (s *SurfaceConditionDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.ExitSurfaceConditionDeclaration(s)
	}
}

func (p *MMSParser) SurfaceConditionDeclaration() (localctx ISurfaceConditionDeclarationContext) {
	localctx = NewSurfaceConditionDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, MMSParserRULE_surfaceConditionDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(231)
		p.Match(MMSParserKeyword_Condition)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(232)
		p.Match(MMSParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(236)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MMSParserNL {
		{
			p.SetState(233)
			p.Match(MMSParserNL)
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
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(239)
		p.SurfaceCondition()
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
	SurfaceCondition_Compound() ISurfaceCondition_CompoundContext
	SurfaceConditionReference() ISurfaceConditionReferenceContext

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
	p.RuleIndex = MMSParserRULE_surfaceCondition
	return p
}

func InitEmptySurfaceConditionContext(p *SurfaceConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_surfaceCondition
}

func (*SurfaceConditionContext) IsSurfaceConditionContext() {}

func NewSurfaceConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceConditionContext {
	var p = new(SurfaceConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_surfaceCondition

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

func (s *SurfaceConditionContext) SurfaceCondition_Compound() ISurfaceCondition_CompoundContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceCondition_CompoundContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISurfaceCondition_CompoundContext)
}

func (s *SurfaceConditionContext) SurfaceConditionReference() ISurfaceConditionReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceConditionReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISurfaceConditionReferenceContext)
}

func (s *SurfaceConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.EnterSurfaceCondition(s)
	}
}

func (s *SurfaceConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.ExitSurfaceCondition(s)
	}
}

func (p *MMSParser) SurfaceCondition() (localctx ISurfaceConditionContext) {
	localctx = NewSurfaceConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, MMSParserRULE_surfaceCondition)
	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(241)
			p.SurfaceCondition_AboveSurface()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(242)
			p.SurfaceCondition_Biome()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(243)
			p.SurfaceCondition_Hole()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(244)
			p.SurfaceCondition_Noise()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(245)
			p.SurfaceCondition_Steep()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(246)
			p.SurfaceCondition_StoneDepth()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(247)
			p.SurfaceCondition_Freezing()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(248)
			p.SurfaceCondition_VerticalGradient()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(249)
			p.SurfaceCondition_AboveWater()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(250)
			p.SurfaceCondition_YAbove()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(251)
			p.SurfaceCondition_Compound()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(252)
			p.SurfaceConditionReference()
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

// ISurfaceCondition_AboveSurfaceContext is an interface to support dynamic dispatch.
type ISurfaceCondition_AboveSurfaceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Keyword_AbovePreliminarySurface() antlr.TerminalNode

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
	p.RuleIndex = MMSParserRULE_surfaceCondition_AboveSurface
	return p
}

func InitEmptySurfaceCondition_AboveSurfaceContext(p *SurfaceCondition_AboveSurfaceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_surfaceCondition_AboveSurface
}

func (*SurfaceCondition_AboveSurfaceContext) IsSurfaceCondition_AboveSurfaceContext() {}

func NewSurfaceCondition_AboveSurfaceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_AboveSurfaceContext {
	var p = new(SurfaceCondition_AboveSurfaceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_surfaceCondition_AboveSurface

	return p
}

func (s *SurfaceCondition_AboveSurfaceContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceCondition_AboveSurfaceContext) Keyword_AbovePreliminarySurface() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_AbovePreliminarySurface, 0)
}

func (s *SurfaceCondition_AboveSurfaceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_AboveSurfaceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_AboveSurfaceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.EnterSurfaceCondition_AboveSurface(s)
	}
}

func (s *SurfaceCondition_AboveSurfaceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.ExitSurfaceCondition_AboveSurface(s)
	}
}

func (p *MMSParser) SurfaceCondition_AboveSurface() (localctx ISurfaceCondition_AboveSurfaceContext) {
	localctx = NewSurfaceCondition_AboveSurfaceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, MMSParserRULE_surfaceCondition_AboveSurface)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(255)
		p.Match(MMSParserKeyword_AbovePreliminarySurface)
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
	p.RuleIndex = MMSParserRULE_surfaceCondition_Biome
	return p
}

func InitEmptySurfaceCondition_BiomeContext(p *SurfaceCondition_BiomeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_surfaceCondition_Biome
}

func (*SurfaceCondition_BiomeContext) IsSurfaceCondition_BiomeContext() {}

func NewSurfaceCondition_BiomeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_BiomeContext {
	var p = new(SurfaceCondition_BiomeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_surfaceCondition_Biome

	return p
}

func (s *SurfaceCondition_BiomeContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceCondition_BiomeContext) Keyword_Biome() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_Biome, 0)
}

func (s *SurfaceCondition_BiomeContext) SquareOpen() antlr.TerminalNode {
	return s.GetToken(MMSParserSquareOpen, 0)
}

func (s *SurfaceCondition_BiomeContext) SquareClose() antlr.TerminalNode {
	return s.GetToken(MMSParserSquareClose, 0)
}

func (s *SurfaceCondition_BiomeContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(MMSParserNL)
}

func (s *SurfaceCondition_BiomeContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(MMSParserNL, i)
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
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.EnterSurfaceCondition_Biome(s)
	}
}

func (s *SurfaceCondition_BiomeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.ExitSurfaceCondition_Biome(s)
	}
}

func (p *MMSParser) SurfaceCondition_Biome() (localctx ISurfaceCondition_BiomeContext) {
	localctx = NewSurfaceCondition_BiomeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, MMSParserRULE_surfaceCondition_Biome)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(257)
		p.Match(MMSParserKeyword_Biome)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(258)
		p.Match(MMSParserSquareOpen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MMSParserNL {
		{
			p.SetState(259)
			p.Match(MMSParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(264)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(274)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&70377300557816) != 0 {
		{
			p.SetState(265)
			p.ResourceReference()
		}
		p.SetState(269)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MMSParserNL {
			{
				p.SetState(266)
				p.Match(MMSParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(271)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(276)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(277)
		p.Match(MMSParserSquareClose)
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
	p.RuleIndex = MMSParserRULE_surfaceCondition_Hole
	return p
}

func InitEmptySurfaceCondition_HoleContext(p *SurfaceCondition_HoleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_surfaceCondition_Hole
}

func (*SurfaceCondition_HoleContext) IsSurfaceCondition_HoleContext() {}

func NewSurfaceCondition_HoleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_HoleContext {
	var p = new(SurfaceCondition_HoleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_surfaceCondition_Hole

	return p
}

func (s *SurfaceCondition_HoleContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceCondition_HoleContext) Keyword_Hole() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_Hole, 0)
}

func (s *SurfaceCondition_HoleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_HoleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_HoleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.EnterSurfaceCondition_Hole(s)
	}
}

func (s *SurfaceCondition_HoleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.ExitSurfaceCondition_Hole(s)
	}
}

func (p *MMSParser) SurfaceCondition_Hole() (localctx ISurfaceCondition_HoleContext) {
	localctx = NewSurfaceCondition_HoleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, MMSParserRULE_surfaceCondition_Hole)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(279)
		p.Match(MMSParserKeyword_Hole)
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
	p.RuleIndex = MMSParserRULE_surfaceCondition_Noise
	return p
}

func InitEmptySurfaceCondition_NoiseContext(p *SurfaceCondition_NoiseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_surfaceCondition_Noise
}

func (*SurfaceCondition_NoiseContext) IsSurfaceCondition_NoiseContext() {}

func NewSurfaceCondition_NoiseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_NoiseContext {
	var p = new(SurfaceCondition_NoiseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_surfaceCondition_Noise

	return p
}

func (s *SurfaceCondition_NoiseContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceCondition_NoiseContext) Keyword_Noise() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_Noise, 0)
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
	return s.GetToken(MMSParserSquareOpen, 0)
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
	return s.GetToken(MMSParserComma, 0)
}

func (s *SurfaceCondition_NoiseContext) SquareClose() antlr.TerminalNode {
	return s.GetToken(MMSParserSquareClose, 0)
}

func (s *SurfaceCondition_NoiseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_NoiseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_NoiseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.EnterSurfaceCondition_Noise(s)
	}
}

func (s *SurfaceCondition_NoiseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.ExitSurfaceCondition_Noise(s)
	}
}

func (p *MMSParser) SurfaceCondition_Noise() (localctx ISurfaceCondition_NoiseContext) {
	localctx = NewSurfaceCondition_NoiseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, MMSParserRULE_surfaceCondition_Noise)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(281)
		p.Match(MMSParserKeyword_Noise)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(282)
		p.ResourceReference()
	}
	{
		p.SetState(283)
		p.Match(MMSParserSquareOpen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(284)
		p.Number()
	}
	{
		p.SetState(285)
		p.Match(MMSParserComma)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(286)
		p.Number()
	}
	{
		p.SetState(287)
		p.Match(MMSParserSquareClose)
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
	p.RuleIndex = MMSParserRULE_surfaceCondition_Steep
	return p
}

func InitEmptySurfaceCondition_SteepContext(p *SurfaceCondition_SteepContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_surfaceCondition_Steep
}

func (*SurfaceCondition_SteepContext) IsSurfaceCondition_SteepContext() {}

func NewSurfaceCondition_SteepContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_SteepContext {
	var p = new(SurfaceCondition_SteepContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_surfaceCondition_Steep

	return p
}

func (s *SurfaceCondition_SteepContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceCondition_SteepContext) Keyword_Steep() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_Steep, 0)
}

func (s *SurfaceCondition_SteepContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_SteepContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_SteepContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.EnterSurfaceCondition_Steep(s)
	}
}

func (s *SurfaceCondition_SteepContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.ExitSurfaceCondition_Steep(s)
	}
}

func (p *MMSParser) SurfaceCondition_Steep() (localctx ISurfaceCondition_SteepContext) {
	localctx = NewSurfaceCondition_SteepContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, MMSParserRULE_surfaceCondition_Steep)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(289)
		p.Match(MMSParserKeyword_Steep)
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
	p.RuleIndex = MMSParserRULE_surfaceCondition_StoneDepth
	return p
}

func InitEmptySurfaceCondition_StoneDepthContext(p *SurfaceCondition_StoneDepthContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_surfaceCondition_StoneDepth
}

func (*SurfaceCondition_StoneDepthContext) IsSurfaceCondition_StoneDepthContext() {}

func NewSurfaceCondition_StoneDepthContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_StoneDepthContext {
	var p = new(SurfaceCondition_StoneDepthContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_surfaceCondition_StoneDepth

	return p
}

func (s *SurfaceCondition_StoneDepthContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceCondition_StoneDepthContext) Keyword_StoneDepth() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_StoneDepth, 0)
}

func (s *SurfaceCondition_StoneDepthContext) AllInt() []antlr.TerminalNode {
	return s.GetTokens(MMSParserInt)
}

func (s *SurfaceCondition_StoneDepthContext) Int(i int) antlr.TerminalNode {
	return s.GetToken(MMSParserInt, i)
}

func (s *SurfaceCondition_StoneDepthContext) Keyword_Floor() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_Floor, 0)
}

func (s *SurfaceCondition_StoneDepthContext) Keyword_Ceiling() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_Ceiling, 0)
}

func (s *SurfaceCondition_StoneDepthContext) Keyword_Add() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_Add, 0)
}

func (s *SurfaceCondition_StoneDepthContext) Keyword_Sub() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_Sub, 0)
}

func (s *SurfaceCondition_StoneDepthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_StoneDepthContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_StoneDepthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.EnterSurfaceCondition_StoneDepth(s)
	}
}

func (s *SurfaceCondition_StoneDepthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.ExitSurfaceCondition_StoneDepth(s)
	}
}

func (p *MMSParser) SurfaceCondition_StoneDepth() (localctx ISurfaceCondition_StoneDepthContext) {
	localctx = NewSurfaceCondition_StoneDepthContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, MMSParserRULE_surfaceCondition_StoneDepth)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(291)
		p.Match(MMSParserKeyword_StoneDepth)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(292)
		_la = p.GetTokenStream().LA(1)

		if !(_la == MMSParserKeyword_Floor || _la == MMSParserKeyword_Ceiling) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(293)
		p.Match(MMSParserInt)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(294)
		_la = p.GetTokenStream().LA(1)

		if !(_la == MMSParserKeyword_Add || _la == MMSParserKeyword_Sub) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(295)
		p.Match(MMSParserInt)
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
	p.RuleIndex = MMSParserRULE_surfaceCondition_Freezing
	return p
}

func InitEmptySurfaceCondition_FreezingContext(p *SurfaceCondition_FreezingContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_surfaceCondition_Freezing
}

func (*SurfaceCondition_FreezingContext) IsSurfaceCondition_FreezingContext() {}

func NewSurfaceCondition_FreezingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_FreezingContext {
	var p = new(SurfaceCondition_FreezingContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_surfaceCondition_Freezing

	return p
}

func (s *SurfaceCondition_FreezingContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceCondition_FreezingContext) Keyword_Freezing() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_Freezing, 0)
}

func (s *SurfaceCondition_FreezingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_FreezingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_FreezingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.EnterSurfaceCondition_Freezing(s)
	}
}

func (s *SurfaceCondition_FreezingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.ExitSurfaceCondition_Freezing(s)
	}
}

func (p *MMSParser) SurfaceCondition_Freezing() (localctx ISurfaceCondition_FreezingContext) {
	localctx = NewSurfaceCondition_FreezingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, MMSParserRULE_surfaceCondition_Freezing)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(297)
		p.Match(MMSParserKeyword_Freezing)
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
	AllVerticalAnchor() []IVerticalAnchorContext
	VerticalAnchor(i int) IVerticalAnchorContext
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
	p.RuleIndex = MMSParserRULE_surfaceCondition_VerticalGradient
	return p
}

func InitEmptySurfaceCondition_VerticalGradientContext(p *SurfaceCondition_VerticalGradientContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_surfaceCondition_VerticalGradient
}

func (*SurfaceCondition_VerticalGradientContext) IsSurfaceCondition_VerticalGradientContext() {}

func NewSurfaceCondition_VerticalGradientContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_VerticalGradientContext {
	var p = new(SurfaceCondition_VerticalGradientContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_surfaceCondition_VerticalGradient

	return p
}

func (s *SurfaceCondition_VerticalGradientContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceCondition_VerticalGradientContext) Keyword_VerticalGradient() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_VerticalGradient, 0)
}

func (s *SurfaceCondition_VerticalGradientContext) String_() antlr.TerminalNode {
	return s.GetToken(MMSParserString_, 0)
}

func (s *SurfaceCondition_VerticalGradientContext) AllVerticalAnchor() []IVerticalAnchorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVerticalAnchorContext); ok {
			len++
		}
	}

	tst := make([]IVerticalAnchorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVerticalAnchorContext); ok {
			tst[i] = t.(IVerticalAnchorContext)
			i++
		}
	}

	return tst
}

func (s *SurfaceCondition_VerticalGradientContext) VerticalAnchor(i int) IVerticalAnchorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVerticalAnchorContext); ok {
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

	return t.(IVerticalAnchorContext)
}

func (s *SurfaceCondition_VerticalGradientContext) Comma() antlr.TerminalNode {
	return s.GetToken(MMSParserComma, 0)
}

func (s *SurfaceCondition_VerticalGradientContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_VerticalGradientContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_VerticalGradientContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.EnterSurfaceCondition_VerticalGradient(s)
	}
}

func (s *SurfaceCondition_VerticalGradientContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.ExitSurfaceCondition_VerticalGradient(s)
	}
}

func (p *MMSParser) SurfaceCondition_VerticalGradient() (localctx ISurfaceCondition_VerticalGradientContext) {
	localctx = NewSurfaceCondition_VerticalGradientContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, MMSParserRULE_surfaceCondition_VerticalGradient)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(299)
		p.Match(MMSParserKeyword_VerticalGradient)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(300)
		p.Match(MMSParserString_)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(301)
		p.VerticalAnchor()
	}
	{
		p.SetState(302)
		p.Match(MMSParserComma)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(303)
		p.VerticalAnchor()
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
	p.RuleIndex = MMSParserRULE_surfaceCondition_AboveWater
	return p
}

func InitEmptySurfaceCondition_AboveWaterContext(p *SurfaceCondition_AboveWaterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_surfaceCondition_AboveWater
}

func (*SurfaceCondition_AboveWaterContext) IsSurfaceCondition_AboveWaterContext() {}

func NewSurfaceCondition_AboveWaterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_AboveWaterContext {
	var p = new(SurfaceCondition_AboveWaterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_surfaceCondition_AboveWater

	return p
}

func (s *SurfaceCondition_AboveWaterContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceCondition_AboveWaterContext) Keyword_AboveWater() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_AboveWater, 0)
}

func (s *SurfaceCondition_AboveWaterContext) Int() antlr.TerminalNode {
	return s.GetToken(MMSParserInt, 0)
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
	return s.GetToken(MMSParserKeyword_Add, 0)
}

func (s *SurfaceCondition_AboveWaterContext) Keyword_Sub() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_Sub, 0)
}

func (s *SurfaceCondition_AboveWaterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_AboveWaterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_AboveWaterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.EnterSurfaceCondition_AboveWater(s)
	}
}

func (s *SurfaceCondition_AboveWaterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.ExitSurfaceCondition_AboveWater(s)
	}
}

func (p *MMSParser) SurfaceCondition_AboveWater() (localctx ISurfaceCondition_AboveWaterContext) {
	localctx = NewSurfaceCondition_AboveWaterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, MMSParserRULE_surfaceCondition_AboveWater)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(305)
		p.Match(MMSParserKeyword_AboveWater)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(306)
		p.Match(MMSParserInt)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(307)
		p.Number()
	}
	{
		p.SetState(308)
		_la = p.GetTokenStream().LA(1)

		if !(_la == MMSParserKeyword_Add || _la == MMSParserKeyword_Sub) {
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
	VerticalAnchor() IVerticalAnchorContext
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
	p.RuleIndex = MMSParserRULE_surfaceCondition_YAbove
	return p
}

func InitEmptySurfaceCondition_YAboveContext(p *SurfaceCondition_YAboveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_surfaceCondition_YAbove
}

func (*SurfaceCondition_YAboveContext) IsSurfaceCondition_YAboveContext() {}

func NewSurfaceCondition_YAboveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_YAboveContext {
	var p = new(SurfaceCondition_YAboveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_surfaceCondition_YAbove

	return p
}

func (s *SurfaceCondition_YAboveContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceCondition_YAboveContext) Keyword_YAbove() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_YAbove, 0)
}

func (s *SurfaceCondition_YAboveContext) VerticalAnchor() IVerticalAnchorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVerticalAnchorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVerticalAnchorContext)
}

func (s *SurfaceCondition_YAboveContext) Int() antlr.TerminalNode {
	return s.GetToken(MMSParserInt, 0)
}

func (s *SurfaceCondition_YAboveContext) Keyword_Add() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_Add, 0)
}

func (s *SurfaceCondition_YAboveContext) Keyword_Sub() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_Sub, 0)
}

func (s *SurfaceCondition_YAboveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_YAboveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_YAboveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.EnterSurfaceCondition_YAbove(s)
	}
}

func (s *SurfaceCondition_YAboveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.ExitSurfaceCondition_YAbove(s)
	}
}

func (p *MMSParser) SurfaceCondition_YAbove() (localctx ISurfaceCondition_YAboveContext) {
	localctx = NewSurfaceCondition_YAboveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, MMSParserRULE_surfaceCondition_YAbove)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(310)
		p.Match(MMSParserKeyword_YAbove)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(311)
		p.VerticalAnchor()
	}
	{
		p.SetState(312)
		p.Match(MMSParserInt)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(313)
		_la = p.GetTokenStream().LA(1)

		if !(_la == MMSParserKeyword_Add || _la == MMSParserKeyword_Sub) {
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

// ISurfaceCondition_Compound__ItemContext is an interface to support dynamic dispatch.
type ISurfaceCondition_Compound__ItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SurfaceCondition() ISurfaceConditionContext
	Bang() antlr.TerminalNode

	// IsSurfaceCondition_Compound__ItemContext differentiates from other interfaces.
	IsSurfaceCondition_Compound__ItemContext()
}

type SurfaceCondition_Compound__ItemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceCondition_Compound__ItemContext() *SurfaceCondition_Compound__ItemContext {
	var p = new(SurfaceCondition_Compound__ItemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_surfaceCondition_Compound__Item
	return p
}

func InitEmptySurfaceCondition_Compound__ItemContext(p *SurfaceCondition_Compound__ItemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_surfaceCondition_Compound__Item
}

func (*SurfaceCondition_Compound__ItemContext) IsSurfaceCondition_Compound__ItemContext() {}

func NewSurfaceCondition_Compound__ItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_Compound__ItemContext {
	var p = new(SurfaceCondition_Compound__ItemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_surfaceCondition_Compound__Item

	return p
}

func (s *SurfaceCondition_Compound__ItemContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceCondition_Compound__ItemContext) SurfaceCondition() ISurfaceConditionContext {
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

func (s *SurfaceCondition_Compound__ItemContext) Bang() antlr.TerminalNode {
	return s.GetToken(MMSParserBang, 0)
}

func (s *SurfaceCondition_Compound__ItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_Compound__ItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_Compound__ItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.EnterSurfaceCondition_Compound__Item(s)
	}
}

func (s *SurfaceCondition_Compound__ItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.ExitSurfaceCondition_Compound__Item(s)
	}
}

func (p *MMSParser) SurfaceCondition_Compound__Item() (localctx ISurfaceCondition_Compound__ItemContext) {
	localctx = NewSurfaceCondition_Compound__ItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, MMSParserRULE_surfaceCondition_Compound__Item)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(316)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MMSParserBang {
		{
			p.SetState(315)
			p.Match(MMSParserBang)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(318)
		p.SurfaceCondition()
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

// ISurfaceCondition_CompoundContext is an interface to support dynamic dispatch.
type ISurfaceCondition_CompoundContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Keyword_And() antlr.TerminalNode
	RoundOpen() antlr.TerminalNode
	RoundClose() antlr.TerminalNode
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode
	AllSurfaceCondition_Compound__Item() []ISurfaceCondition_Compound__ItemContext
	SurfaceCondition_Compound__Item(i int) ISurfaceCondition_Compound__ItemContext

	// IsSurfaceCondition_CompoundContext differentiates from other interfaces.
	IsSurfaceCondition_CompoundContext()
}

type SurfaceCondition_CompoundContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceCondition_CompoundContext() *SurfaceCondition_CompoundContext {
	var p = new(SurfaceCondition_CompoundContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_surfaceCondition_Compound
	return p
}

func InitEmptySurfaceCondition_CompoundContext(p *SurfaceCondition_CompoundContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_surfaceCondition_Compound
}

func (*SurfaceCondition_CompoundContext) IsSurfaceCondition_CompoundContext() {}

func NewSurfaceCondition_CompoundContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_CompoundContext {
	var p = new(SurfaceCondition_CompoundContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_surfaceCondition_Compound

	return p
}

func (s *SurfaceCondition_CompoundContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceCondition_CompoundContext) Keyword_And() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_And, 0)
}

func (s *SurfaceCondition_CompoundContext) RoundOpen() antlr.TerminalNode {
	return s.GetToken(MMSParserRoundOpen, 0)
}

func (s *SurfaceCondition_CompoundContext) RoundClose() antlr.TerminalNode {
	return s.GetToken(MMSParserRoundClose, 0)
}

func (s *SurfaceCondition_CompoundContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(MMSParserNL)
}

func (s *SurfaceCondition_CompoundContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(MMSParserNL, i)
}

func (s *SurfaceCondition_CompoundContext) AllSurfaceCondition_Compound__Item() []ISurfaceCondition_Compound__ItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISurfaceCondition_Compound__ItemContext); ok {
			len++
		}
	}

	tst := make([]ISurfaceCondition_Compound__ItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISurfaceCondition_Compound__ItemContext); ok {
			tst[i] = t.(ISurfaceCondition_Compound__ItemContext)
			i++
		}
	}

	return tst
}

func (s *SurfaceCondition_CompoundContext) SurfaceCondition_Compound__Item(i int) ISurfaceCondition_Compound__ItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceCondition_Compound__ItemContext); ok {
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

	return t.(ISurfaceCondition_Compound__ItemContext)
}

func (s *SurfaceCondition_CompoundContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_CompoundContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_CompoundContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.EnterSurfaceCondition_Compound(s)
	}
}

func (s *SurfaceCondition_CompoundContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.ExitSurfaceCondition_Compound(s)
	}
}

func (p *MMSParser) SurfaceCondition_Compound() (localctx ISurfaceCondition_CompoundContext) {
	localctx = NewSurfaceCondition_CompoundContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, MMSParserRULE_surfaceCondition_Compound)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(320)
		p.Match(MMSParserKeyword_And)
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

	for _la == MMSParserNL {
		{
			p.SetState(321)
			p.Match(MMSParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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
		p.Match(MMSParserRoundOpen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(331)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MMSParserNL {
		{
			p.SetState(328)
			p.Match(MMSParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(333)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(343)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&72576323813368) != 0 {
		{
			p.SetState(334)
			p.SurfaceCondition_Compound__Item()
		}
		p.SetState(338)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MMSParserNL {
			{
				p.SetState(335)
				p.Match(MMSParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(340)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(345)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(346)
		p.Match(MMSParserRoundClose)
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

// IReferenceContext is an interface to support dynamic dispatch.
type IReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Colon() antlr.TerminalNode
	AllKeyword() []IKeywordContext
	Keyword(i int) IKeywordContext
	AllIdentifier() []antlr.TerminalNode
	Identifier(i int) antlr.TerminalNode

	// IsReferenceContext differentiates from other interfaces.
	IsReferenceContext()
}

type ReferenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReferenceContext() *ReferenceContext {
	var p = new(ReferenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_reference
	return p
}

func InitEmptyReferenceContext(p *ReferenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_reference
}

func (*ReferenceContext) IsReferenceContext() {}

func NewReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReferenceContext {
	var p = new(ReferenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_reference

	return p
}

func (s *ReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *ReferenceContext) Colon() antlr.TerminalNode {
	return s.GetToken(MMSParserColon, 0)
}

func (s *ReferenceContext) AllKeyword() []IKeywordContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IKeywordContext); ok {
			len++
		}
	}

	tst := make([]IKeywordContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IKeywordContext); ok {
			tst[i] = t.(IKeywordContext)
			i++
		}
	}

	return tst
}

func (s *ReferenceContext) Keyword(i int) IKeywordContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeywordContext); ok {
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

	return t.(IKeywordContext)
}

func (s *ReferenceContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(MMSParserIdentifier)
}

func (s *ReferenceContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(MMSParserIdentifier, i)
}

func (s *ReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.EnterReference(s)
	}
}

func (s *ReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.ExitReference(s)
	}
}

func (p *MMSParser) Reference() (localctx IReferenceContext) {
	localctx = NewReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, MMSParserRULE_reference)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(350)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MMSParserKeyword_Surface, MMSParserKeyword_Rule, MMSParserKeyword_Condition, MMSParserKeyword_Sequence, MMSParserKeyword_Block, MMSParserKeyword_Bandlands, MMSParserKeyword_AbovePreliminarySurface, MMSParserKeyword_Biome, MMSParserKeyword_Hole, MMSParserKeyword_Noise, MMSParserKeyword_Steep, MMSParserKeyword_StoneDepth, MMSParserKeyword_Freezing, MMSParserKeyword_Temperature, MMSParserKeyword_VerticalGradient, MMSParserKeyword_AboveWater, MMSParserKeyword_YAbove, MMSParserKeyword_Floor, MMSParserKeyword_Ceiling, MMSParserKeyword_And, MMSParserKeyword_Add, MMSParserKeyword_Sub, MMSParserKeyword_Absolute, MMSParserKeyword_AboveBottom, MMSParserKeyword_BelowTop, MMSParserKeyword_Namespace, MMSParserKeyword_If, MMSParserKeyword_Else, MMSParserKeyword_In:
		{
			p.SetState(348)
			p.Keyword()
		}

	case MMSParserIdentifier:
		{
			p.SetState(349)
			p.Match(MMSParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(352)
		p.Match(MMSParserColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(355)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MMSParserKeyword_Surface, MMSParserKeyword_Rule, MMSParserKeyword_Condition, MMSParserKeyword_Sequence, MMSParserKeyword_Block, MMSParserKeyword_Bandlands, MMSParserKeyword_AbovePreliminarySurface, MMSParserKeyword_Biome, MMSParserKeyword_Hole, MMSParserKeyword_Noise, MMSParserKeyword_Steep, MMSParserKeyword_StoneDepth, MMSParserKeyword_Freezing, MMSParserKeyword_Temperature, MMSParserKeyword_VerticalGradient, MMSParserKeyword_AboveWater, MMSParserKeyword_YAbove, MMSParserKeyword_Floor, MMSParserKeyword_Ceiling, MMSParserKeyword_And, MMSParserKeyword_Add, MMSParserKeyword_Sub, MMSParserKeyword_Absolute, MMSParserKeyword_AboveBottom, MMSParserKeyword_BelowTop, MMSParserKeyword_Namespace, MMSParserKeyword_If, MMSParserKeyword_Else, MMSParserKeyword_In:
		{
			p.SetState(353)
			p.Keyword()
		}

	case MMSParserIdentifier:
		{
			p.SetState(354)
			p.Match(MMSParserIdentifier)
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

// IResourceReferenceContext is an interface to support dynamic dispatch.
type IResourceReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Reference() IReferenceContext
	Keyword() IKeywordContext
	Identifier() antlr.TerminalNode

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
	p.RuleIndex = MMSParserRULE_resourceReference
	return p
}

func InitEmptyResourceReferenceContext(p *ResourceReferenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_resourceReference
}

func (*ResourceReferenceContext) IsResourceReferenceContext() {}

func NewResourceReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResourceReferenceContext {
	var p = new(ResourceReferenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_resourceReference

	return p
}

func (s *ResourceReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *ResourceReferenceContext) Reference() IReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReferenceContext)
}

func (s *ResourceReferenceContext) Keyword() IKeywordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeywordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeywordContext)
}

func (s *ResourceReferenceContext) Identifier() antlr.TerminalNode {
	return s.GetToken(MMSParserIdentifier, 0)
}

func (s *ResourceReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResourceReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResourceReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.EnterResourceReference(s)
	}
}

func (s *ResourceReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.ExitResourceReference(s)
	}
}

func (p *MMSParser) ResourceReference() (localctx IResourceReferenceContext) {
	localctx = NewResourceReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, MMSParserRULE_resourceReference)
	p.SetState(362)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(357)
			p.Reference()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(360)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case MMSParserKeyword_Surface, MMSParserKeyword_Rule, MMSParserKeyword_Condition, MMSParserKeyword_Sequence, MMSParserKeyword_Block, MMSParserKeyword_Bandlands, MMSParserKeyword_AbovePreliminarySurface, MMSParserKeyword_Biome, MMSParserKeyword_Hole, MMSParserKeyword_Noise, MMSParserKeyword_Steep, MMSParserKeyword_StoneDepth, MMSParserKeyword_Freezing, MMSParserKeyword_Temperature, MMSParserKeyword_VerticalGradient, MMSParserKeyword_AboveWater, MMSParserKeyword_YAbove, MMSParserKeyword_Floor, MMSParserKeyword_Ceiling, MMSParserKeyword_And, MMSParserKeyword_Add, MMSParserKeyword_Sub, MMSParserKeyword_Absolute, MMSParserKeyword_AboveBottom, MMSParserKeyword_BelowTop, MMSParserKeyword_Namespace, MMSParserKeyword_If, MMSParserKeyword_Else, MMSParserKeyword_In:
			{
				p.SetState(358)
				p.Keyword()
			}

		case MMSParserIdentifier:
			{
				p.SetState(359)
				p.Match(MMSParserIdentifier)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = MMSParserRULE_number
	return p
}

func InitEmptyNumberContext(p *NumberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_number
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) Int() antlr.TerminalNode {
	return s.GetToken(MMSParserInt, 0)
}

func (s *NumberContext) Float() antlr.TerminalNode {
	return s.GetToken(MMSParserFloat, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *MMSParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, MMSParserRULE_number)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(364)
		_la = p.GetTokenStream().LA(1)

		if !(_la == MMSParserInt || _la == MMSParserFloat) {
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

// IKeywordContext is an interface to support dynamic dispatch.
type IKeywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Keyword_Surface() antlr.TerminalNode
	Keyword_Rule() antlr.TerminalNode
	Keyword_Condition() antlr.TerminalNode
	Keyword_Sequence() antlr.TerminalNode
	Keyword_Block() antlr.TerminalNode
	Keyword_Bandlands() antlr.TerminalNode
	Keyword_AbovePreliminarySurface() antlr.TerminalNode
	Keyword_Biome() antlr.TerminalNode
	Keyword_Hole() antlr.TerminalNode
	Keyword_Noise() antlr.TerminalNode
	Keyword_Steep() antlr.TerminalNode
	Keyword_StoneDepth() antlr.TerminalNode
	Keyword_Freezing() antlr.TerminalNode
	Keyword_Temperature() antlr.TerminalNode
	Keyword_VerticalGradient() antlr.TerminalNode
	Keyword_AboveWater() antlr.TerminalNode
	Keyword_YAbove() antlr.TerminalNode
	Keyword_Floor() antlr.TerminalNode
	Keyword_Ceiling() antlr.TerminalNode
	Keyword_And() antlr.TerminalNode
	Keyword_Add() antlr.TerminalNode
	Keyword_Sub() antlr.TerminalNode
	Keyword_Absolute() antlr.TerminalNode
	Keyword_AboveBottom() antlr.TerminalNode
	Keyword_BelowTop() antlr.TerminalNode
	Keyword_Namespace() antlr.TerminalNode
	Keyword_If() antlr.TerminalNode
	Keyword_Else() antlr.TerminalNode
	Keyword_In() antlr.TerminalNode

	// IsKeywordContext differentiates from other interfaces.
	IsKeywordContext()
}

type KeywordContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeywordContext() *KeywordContext {
	var p = new(KeywordContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_keyword
	return p
}

func InitEmptyKeywordContext(p *KeywordContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_keyword
}

func (*KeywordContext) IsKeywordContext() {}

func NewKeywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordContext {
	var p = new(KeywordContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_keyword

	return p
}

func (s *KeywordContext) GetParser() antlr.Parser { return s.parser }

func (s *KeywordContext) Keyword_Surface() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_Surface, 0)
}

func (s *KeywordContext) Keyword_Rule() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_Rule, 0)
}

func (s *KeywordContext) Keyword_Condition() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_Condition, 0)
}

func (s *KeywordContext) Keyword_Sequence() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_Sequence, 0)
}

func (s *KeywordContext) Keyword_Block() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_Block, 0)
}

func (s *KeywordContext) Keyword_Bandlands() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_Bandlands, 0)
}

func (s *KeywordContext) Keyword_AbovePreliminarySurface() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_AbovePreliminarySurface, 0)
}

func (s *KeywordContext) Keyword_Biome() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_Biome, 0)
}

func (s *KeywordContext) Keyword_Hole() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_Hole, 0)
}

func (s *KeywordContext) Keyword_Noise() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_Noise, 0)
}

func (s *KeywordContext) Keyword_Steep() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_Steep, 0)
}

func (s *KeywordContext) Keyword_StoneDepth() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_StoneDepth, 0)
}

func (s *KeywordContext) Keyword_Freezing() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_Freezing, 0)
}

func (s *KeywordContext) Keyword_Temperature() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_Temperature, 0)
}

func (s *KeywordContext) Keyword_VerticalGradient() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_VerticalGradient, 0)
}

func (s *KeywordContext) Keyword_AboveWater() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_AboveWater, 0)
}

func (s *KeywordContext) Keyword_YAbove() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_YAbove, 0)
}

func (s *KeywordContext) Keyword_Floor() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_Floor, 0)
}

func (s *KeywordContext) Keyword_Ceiling() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_Ceiling, 0)
}

func (s *KeywordContext) Keyword_And() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_And, 0)
}

func (s *KeywordContext) Keyword_Add() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_Add, 0)
}

func (s *KeywordContext) Keyword_Sub() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_Sub, 0)
}

func (s *KeywordContext) Keyword_Absolute() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_Absolute, 0)
}

func (s *KeywordContext) Keyword_AboveBottom() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_AboveBottom, 0)
}

func (s *KeywordContext) Keyword_BelowTop() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_BelowTop, 0)
}

func (s *KeywordContext) Keyword_Namespace() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_Namespace, 0)
}

func (s *KeywordContext) Keyword_If() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_If, 0)
}

func (s *KeywordContext) Keyword_Else() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_Else, 0)
}

func (s *KeywordContext) Keyword_In() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_In, 0)
}

func (s *KeywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.EnterKeyword(s)
	}
}

func (s *KeywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.ExitKeyword(s)
	}
}

func (p *MMSParser) Keyword() (localctx IKeywordContext) {
	localctx = NewKeywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, MMSParserRULE_keyword)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(366)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8556380152) != 0) {
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

// IVerticalAnchorContext is an interface to support dynamic dispatch.
type IVerticalAnchorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VerticalAnchor_Absolute() IVerticalAnchor_AbsoluteContext
	VerticalAnchor_AboveBottom() IVerticalAnchor_AboveBottomContext
	VerticalAnchor_BelowTop() IVerticalAnchor_BelowTopContext

	// IsVerticalAnchorContext differentiates from other interfaces.
	IsVerticalAnchorContext()
}

type VerticalAnchorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVerticalAnchorContext() *VerticalAnchorContext {
	var p = new(VerticalAnchorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_verticalAnchor
	return p
}

func InitEmptyVerticalAnchorContext(p *VerticalAnchorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_verticalAnchor
}

func (*VerticalAnchorContext) IsVerticalAnchorContext() {}

func NewVerticalAnchorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VerticalAnchorContext {
	var p = new(VerticalAnchorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_verticalAnchor

	return p
}

func (s *VerticalAnchorContext) GetParser() antlr.Parser { return s.parser }

func (s *VerticalAnchorContext) VerticalAnchor_Absolute() IVerticalAnchor_AbsoluteContext {
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

func (s *VerticalAnchorContext) VerticalAnchor_AboveBottom() IVerticalAnchor_AboveBottomContext {
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

func (s *VerticalAnchorContext) VerticalAnchor_BelowTop() IVerticalAnchor_BelowTopContext {
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

func (s *VerticalAnchorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VerticalAnchorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VerticalAnchorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.EnterVerticalAnchor(s)
	}
}

func (s *VerticalAnchorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.ExitVerticalAnchor(s)
	}
}

func (p *MMSParser) VerticalAnchor() (localctx IVerticalAnchorContext) {
	localctx = NewVerticalAnchorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, MMSParserRULE_verticalAnchor)
	p.SetState(371)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MMSParserKeyword_Absolute:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(368)
			p.VerticalAnchor_Absolute()
		}

	case MMSParserKeyword_AboveBottom:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(369)
			p.VerticalAnchor_AboveBottom()
		}

	case MMSParserKeyword_BelowTop:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(370)
			p.VerticalAnchor_BelowTop()
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
	Int() antlr.TerminalNode

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
	p.RuleIndex = MMSParserRULE_verticalAnchor_Absolute
	return p
}

func InitEmptyVerticalAnchor_AbsoluteContext(p *VerticalAnchor_AbsoluteContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_verticalAnchor_Absolute
}

func (*VerticalAnchor_AbsoluteContext) IsVerticalAnchor_AbsoluteContext() {}

func NewVerticalAnchor_AbsoluteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VerticalAnchor_AbsoluteContext {
	var p = new(VerticalAnchor_AbsoluteContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_verticalAnchor_Absolute

	return p
}

func (s *VerticalAnchor_AbsoluteContext) GetParser() antlr.Parser { return s.parser }

func (s *VerticalAnchor_AbsoluteContext) Keyword_Absolute() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_Absolute, 0)
}

func (s *VerticalAnchor_AbsoluteContext) Int() antlr.TerminalNode {
	return s.GetToken(MMSParserInt, 0)
}

func (s *VerticalAnchor_AbsoluteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VerticalAnchor_AbsoluteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VerticalAnchor_AbsoluteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.EnterVerticalAnchor_Absolute(s)
	}
}

func (s *VerticalAnchor_AbsoluteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.ExitVerticalAnchor_Absolute(s)
	}
}

func (p *MMSParser) VerticalAnchor_Absolute() (localctx IVerticalAnchor_AbsoluteContext) {
	localctx = NewVerticalAnchor_AbsoluteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, MMSParserRULE_verticalAnchor_Absolute)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(373)
		p.Match(MMSParserKeyword_Absolute)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(374)
		p.Match(MMSParserInt)
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
	Int() antlr.TerminalNode

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
	p.RuleIndex = MMSParserRULE_verticalAnchor_AboveBottom
	return p
}

func InitEmptyVerticalAnchor_AboveBottomContext(p *VerticalAnchor_AboveBottomContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_verticalAnchor_AboveBottom
}

func (*VerticalAnchor_AboveBottomContext) IsVerticalAnchor_AboveBottomContext() {}

func NewVerticalAnchor_AboveBottomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VerticalAnchor_AboveBottomContext {
	var p = new(VerticalAnchor_AboveBottomContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_verticalAnchor_AboveBottom

	return p
}

func (s *VerticalAnchor_AboveBottomContext) GetParser() antlr.Parser { return s.parser }

func (s *VerticalAnchor_AboveBottomContext) Keyword_AboveBottom() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_AboveBottom, 0)
}

func (s *VerticalAnchor_AboveBottomContext) Int() antlr.TerminalNode {
	return s.GetToken(MMSParserInt, 0)
}

func (s *VerticalAnchor_AboveBottomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VerticalAnchor_AboveBottomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VerticalAnchor_AboveBottomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.EnterVerticalAnchor_AboveBottom(s)
	}
}

func (s *VerticalAnchor_AboveBottomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.ExitVerticalAnchor_AboveBottom(s)
	}
}

func (p *MMSParser) VerticalAnchor_AboveBottom() (localctx IVerticalAnchor_AboveBottomContext) {
	localctx = NewVerticalAnchor_AboveBottomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, MMSParserRULE_verticalAnchor_AboveBottom)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(376)
		p.Match(MMSParserKeyword_AboveBottom)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(377)
		p.Match(MMSParserInt)
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
	Int() antlr.TerminalNode

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
	p.RuleIndex = MMSParserRULE_verticalAnchor_BelowTop
	return p
}

func InitEmptyVerticalAnchor_BelowTopContext(p *VerticalAnchor_BelowTopContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_verticalAnchor_BelowTop
}

func (*VerticalAnchor_BelowTopContext) IsVerticalAnchor_BelowTopContext() {}

func NewVerticalAnchor_BelowTopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VerticalAnchor_BelowTopContext {
	var p = new(VerticalAnchor_BelowTopContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_verticalAnchor_BelowTop

	return p
}

func (s *VerticalAnchor_BelowTopContext) GetParser() antlr.Parser { return s.parser }

func (s *VerticalAnchor_BelowTopContext) Keyword_BelowTop() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_BelowTop, 0)
}

func (s *VerticalAnchor_BelowTopContext) Int() antlr.TerminalNode {
	return s.GetToken(MMSParserInt, 0)
}

func (s *VerticalAnchor_BelowTopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VerticalAnchor_BelowTopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VerticalAnchor_BelowTopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.EnterVerticalAnchor_BelowTop(s)
	}
}

func (s *VerticalAnchor_BelowTopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.ExitVerticalAnchor_BelowTop(s)
	}
}

func (p *MMSParser) VerticalAnchor_BelowTop() (localctx IVerticalAnchor_BelowTopContext) {
	localctx = NewVerticalAnchor_BelowTopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, MMSParserRULE_verticalAnchor_BelowTop)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(379)
		p.Match(MMSParserKeyword_BelowTop)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(380)
		p.Match(MMSParserInt)
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

// IWorldgenDeclarationContext is an interface to support dynamic dispatch.
type IWorldgenDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Keyword_WorldGen() antlr.TerminalNode
	CurlyOpen() antlr.TerminalNode
	AllNoiseDeclaration() []INoiseDeclarationContext
	NoiseDeclaration(i int) INoiseDeclarationContext
	CurlyClose() antlr.TerminalNode
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode

	// IsWorldgenDeclarationContext differentiates from other interfaces.
	IsWorldgenDeclarationContext()
}

type WorldgenDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWorldgenDeclarationContext() *WorldgenDeclarationContext {
	var p = new(WorldgenDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_worldgenDeclaration
	return p
}

func InitEmptyWorldgenDeclarationContext(p *WorldgenDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_worldgenDeclaration
}

func (*WorldgenDeclarationContext) IsWorldgenDeclarationContext() {}

func NewWorldgenDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WorldgenDeclarationContext {
	var p = new(WorldgenDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_worldgenDeclaration

	return p
}

func (s *WorldgenDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *WorldgenDeclarationContext) Keyword_WorldGen() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_WorldGen, 0)
}

func (s *WorldgenDeclarationContext) CurlyOpen() antlr.TerminalNode {
	return s.GetToken(MMSParserCurlyOpen, 0)
}

func (s *WorldgenDeclarationContext) AllNoiseDeclaration() []INoiseDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INoiseDeclarationContext); ok {
			len++
		}
	}

	tst := make([]INoiseDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INoiseDeclarationContext); ok {
			tst[i] = t.(INoiseDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *WorldgenDeclarationContext) NoiseDeclaration(i int) INoiseDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INoiseDeclarationContext); ok {
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

	return t.(INoiseDeclarationContext)
}

func (s *WorldgenDeclarationContext) CurlyClose() antlr.TerminalNode {
	return s.GetToken(MMSParserCurlyClose, 0)
}

func (s *WorldgenDeclarationContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(MMSParserNL)
}

func (s *WorldgenDeclarationContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(MMSParserNL, i)
}

func (s *WorldgenDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WorldgenDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WorldgenDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.EnterWorldgenDeclaration(s)
	}
}

func (s *WorldgenDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.ExitWorldgenDeclaration(s)
	}
}

func (p *MMSParser) WorldgenDeclaration() (localctx IWorldgenDeclarationContext) {
	localctx = NewWorldgenDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, MMSParserRULE_worldgenDeclaration)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(382)
		p.Match(MMSParserKeyword_WorldGen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(383)
		p.Match(MMSParserCurlyOpen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(387)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MMSParserNL {
		{
			p.SetState(384)
			p.Match(MMSParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(389)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

	p.SetState(399)
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
			{
				p.SetState(390)
				p.NoiseDeclaration()
			}
			p.SetState(394)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == MMSParserNL {
				{
					p.SetState(391)
					p.Match(MMSParserNL)
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
			}

		}
		p.SetState(401)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

	{
		p.SetState(402)
		p.NoiseDeclaration()
	}
	p.SetState(406)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MMSParserNL {
		{
			p.SetState(403)
			p.Match(MMSParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(408)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(409)
		p.Match(MMSParserCurlyClose)
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

// INoiseDeclarationContext is an interface to support dynamic dispatch.
type INoiseDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Keyword_Noise() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	NoiseDefinition() INoiseDefinitionContext
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode

	// IsNoiseDeclarationContext differentiates from other interfaces.
	IsNoiseDeclarationContext()
}

type NoiseDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNoiseDeclarationContext() *NoiseDeclarationContext {
	var p = new(NoiseDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_noiseDeclaration
	return p
}

func InitEmptyNoiseDeclarationContext(p *NoiseDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_noiseDeclaration
}

func (*NoiseDeclarationContext) IsNoiseDeclarationContext() {}

func NewNoiseDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NoiseDeclarationContext {
	var p = new(NoiseDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_noiseDeclaration

	return p
}

func (s *NoiseDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *NoiseDeclarationContext) Keyword_Noise() antlr.TerminalNode {
	return s.GetToken(MMSParserKeyword_Noise, 0)
}

func (s *NoiseDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(MMSParserIdentifier, 0)
}

func (s *NoiseDeclarationContext) NoiseDefinition() INoiseDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INoiseDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INoiseDefinitionContext)
}

func (s *NoiseDeclarationContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(MMSParserNL)
}

func (s *NoiseDeclarationContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(MMSParserNL, i)
}

func (s *NoiseDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NoiseDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NoiseDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.EnterNoiseDeclaration(s)
	}
}

func (s *NoiseDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.ExitNoiseDeclaration(s)
	}
}

func (p *MMSParser) NoiseDeclaration() (localctx INoiseDeclarationContext) {
	localctx = NewNoiseDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, MMSParserRULE_noiseDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(411)
		p.Match(MMSParserKeyword_Noise)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(412)
		p.Match(MMSParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(416)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MMSParserNL {
		{
			p.SetState(413)
			p.Match(MMSParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(418)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(419)
		p.NoiseDefinition()
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

// INoiseDefinitionContext is an interface to support dynamic dispatch.
type INoiseDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNumber() []INumberContext
	Number(i int) INumberContext
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode
	SquareOpen() antlr.TerminalNode
	SquareClose() antlr.TerminalNode

	// IsNoiseDefinitionContext differentiates from other interfaces.
	IsNoiseDefinitionContext()
}

type NoiseDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNoiseDefinitionContext() *NoiseDefinitionContext {
	var p = new(NoiseDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_noiseDefinition
	return p
}

func InitEmptyNoiseDefinitionContext(p *NoiseDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MMSParserRULE_noiseDefinition
}

func (*NoiseDefinitionContext) IsNoiseDefinitionContext() {}

func NewNoiseDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NoiseDefinitionContext {
	var p = new(NoiseDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MMSParserRULE_noiseDefinition

	return p
}

func (s *NoiseDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *NoiseDefinitionContext) AllNumber() []INumberContext {
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

func (s *NoiseDefinitionContext) Number(i int) INumberContext {
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

func (s *NoiseDefinitionContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(MMSParserNL)
}

func (s *NoiseDefinitionContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(MMSParserNL, i)
}

func (s *NoiseDefinitionContext) SquareOpen() antlr.TerminalNode {
	return s.GetToken(MMSParserSquareOpen, 0)
}

func (s *NoiseDefinitionContext) SquareClose() antlr.TerminalNode {
	return s.GetToken(MMSParserSquareClose, 0)
}

func (s *NoiseDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NoiseDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NoiseDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.EnterNoiseDefinition(s)
	}
}

func (s *NoiseDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MMSParserListener); ok {
		listenerT.ExitNoiseDefinition(s)
	}
}

func (p *MMSParser) NoiseDefinition() (localctx INoiseDefinitionContext) {
	localctx = NewNoiseDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, MMSParserRULE_noiseDefinition)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(421)
		p.Number()
	}
	p.SetState(425)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(422)
				p.Match(MMSParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(427)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(448)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MMSParserSquareOpen {
		{
			p.SetState(428)
			p.Match(MMSParserSquareOpen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(432)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MMSParserNL {
			{
				p.SetState(429)
				p.Match(MMSParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(434)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(442)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == MMSParserInt || _la == MMSParserFloat {
			{
				p.SetState(435)
				p.Number()
			}
			p.SetState(439)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == MMSParserNL {
				{
					p.SetState(436)
					p.Match(MMSParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				p.SetState(441)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

			p.SetState(444)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(446)
			p.Match(MMSParserSquareClose)
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
