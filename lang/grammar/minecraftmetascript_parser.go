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
		"", "'namespace'", "'{'", "'}'", "'Surface'", "'~'", "'='", "'.Offset('",
		"')'", "'.Add()'", "'.Mul('", "'!'", "'And'", "'('", "'Or'", "'AboveSurface()'",
		"'Biome'", "','", "'Hole'", "'Steep'", "'Freezing'", "'.Min('", "'.Max('",
		"'Noise'", "'StoneDepth('", "'.SecondaryDepthRange('", "'VerticalGradient'",
		"'.Top'", "'.Bottom'", "'AboveWater'", "'YAbove'", "'Block'", "'['",
		"']'", "'Bandlands'", "'If'", "':'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "StoneDepthMode", "Int", "Float", "String", "NL", "WS",
		"Identifier", "BlockComment", "LineComment",
	}
	staticData.RuleNames = []string{
		"script", "namespaceDeclaration", "namespace", "contentBlocks", "surface",
		"surfaceStatement", "verticalAnchor", "verticalAnchorDeclaration", "sharedBuilder_Offset",
		"sharedBuilder_Add", "sharedBuilder_Mul", "sharedBuilder_MulInt", "surfaceCondition",
		"surfaceConditionDeclaration", "surfaceCondition_Not", "surfaceCondition_And",
		"surfaceCondition_Or", "surfaceCondition_Reference", "surfaceCondition_AboveSurface",
		"surfaceCondition_Biome", "surfaceCondition_Hole", "surfaceCondition_Steep",
		"surfaceCondition_Freezing", "surfaceCondition_NoiseBuilder_Min", "surfaceCondition_NoiseBuilder_Max",
		"surfaceCondition_NoiseBuilder", "surfaceCondition_Noise", "surfaceCondition_StoneDepth",
		"surfaceCondition_StoneDepthBuilder", "surfaceCondition_StoneDepthBuilder_SecondaryDepthRange",
		"surfaceCondition_VerticalGradient", "surfaceCondition_VerticalGradientBuilder",
		"surfaceCondition_VerticalGradientBuilder_Top", "surfaceCondition_VerticalGradientBuilder_Bottom",
		"surfaceCondition_AboveWater", "surfaceCondition_AboveWaterBuilder",
		"surfaceCondition_YAbove", "surfaceCondition_YAboveBuilder", "surfaceRuleDeclaration",
		"surfaceRule", "surfaceRule_Reference", "surfaceRule_Block", "surfaceRule_Sequence",
		"surfaceRule_Bandlands", "surfaceRule_If", "resourceReference", "number",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 45, 612, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 1, 0,
		1, 0, 5, 0, 97, 8, 0, 10, 0, 12, 0, 100, 9, 0, 5, 0, 102, 8, 0, 10, 0,
		12, 0, 105, 9, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 5, 2, 112, 8, 2, 10, 2,
		12, 2, 115, 9, 2, 1, 2, 1, 2, 5, 2, 119, 8, 2, 10, 2, 12, 2, 122, 9, 2,
		1, 2, 1, 2, 5, 2, 126, 8, 2, 10, 2, 12, 2, 129, 9, 2, 5, 2, 131, 8, 2,
		10, 2, 12, 2, 134, 9, 2, 1, 2, 5, 2, 137, 8, 2, 10, 2, 12, 2, 140, 9, 2,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 5, 4, 148, 8, 4, 10, 4, 12, 4, 151,
		9, 4, 1, 4, 1, 4, 5, 4, 155, 8, 4, 10, 4, 12, 4, 158, 9, 4, 1, 4, 1, 4,
		5, 4, 162, 8, 4, 10, 4, 12, 4, 165, 9, 4, 5, 4, 167, 8, 4, 10, 4, 12, 4,
		170, 9, 4, 1, 4, 5, 4, 173, 8, 4, 10, 4, 12, 4, 176, 9, 4, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 5, 3, 5, 183, 8, 5, 1, 6, 3, 6, 186, 8, 6, 1, 6, 1, 6, 3,
		6, 190, 8, 6, 1, 7, 1, 7, 1, 7, 5, 7, 195, 8, 7, 10, 7, 12, 7, 198, 9,
		7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 230,
		8, 12, 1, 13, 1, 13, 1, 13, 5, 13, 235, 8, 13, 10, 13, 12, 13, 238, 9,
		13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 5, 15, 247, 8, 15,
		10, 15, 12, 15, 250, 9, 15, 1, 15, 1, 15, 5, 15, 254, 8, 15, 10, 15, 12,
		15, 257, 9, 15, 1, 15, 1, 15, 5, 15, 261, 8, 15, 10, 15, 12, 15, 264, 9,
		15, 5, 15, 266, 8, 15, 10, 15, 12, 15, 269, 9, 15, 1, 15, 1, 15, 5, 15,
		273, 8, 15, 10, 15, 12, 15, 276, 9, 15, 1, 15, 1, 15, 1, 16, 1, 16, 5,
		16, 282, 8, 16, 10, 16, 12, 16, 285, 9, 16, 1, 16, 1, 16, 5, 16, 289, 8,
		16, 10, 16, 12, 16, 292, 9, 16, 1, 16, 1, 16, 5, 16, 296, 8, 16, 10, 16,
		12, 16, 299, 9, 16, 5, 16, 301, 8, 16, 10, 16, 12, 16, 304, 9, 16, 1, 16,
		1, 16, 5, 16, 308, 8, 16, 10, 16, 12, 16, 311, 9, 16, 1, 16, 1, 16, 1,
		17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 5, 19, 322, 8, 19, 10, 19,
		12, 19, 325, 9, 19, 1, 19, 1, 19, 1, 19, 5, 19, 330, 8, 19, 10, 19, 12,
		19, 333, 9, 19, 5, 19, 335, 8, 19, 10, 19, 12, 19, 338, 9, 19, 1, 19, 5,
		19, 341, 8, 19, 10, 19, 12, 19, 344, 9, 19, 1, 19, 1, 19, 5, 19, 348, 8,
		19, 10, 19, 12, 19, 351, 9, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 3, 25, 377, 8, 25,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 5, 26, 384, 8, 26, 10, 26, 12, 26, 387,
		9, 26, 1, 26, 1, 26, 5, 26, 391, 8, 26, 10, 26, 12, 26, 394, 9, 26, 5,
		26, 396, 8, 26, 10, 26, 12, 26, 399, 9, 26, 1, 27, 1, 27, 1, 27, 1, 27,
		5, 27, 405, 8, 27, 10, 27, 12, 27, 408, 9, 27, 1, 27, 1, 27, 5, 27, 412,
		8, 27, 10, 27, 12, 27, 415, 9, 27, 5, 27, 417, 8, 27, 10, 27, 12, 27, 420,
		9, 27, 1, 28, 1, 28, 1, 28, 3, 28, 425, 8, 28, 1, 29, 1, 29, 1, 29, 1,
		29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 5, 30, 436, 8, 30, 10, 30, 12, 30,
		439, 9, 30, 1, 30, 1, 30, 5, 30, 443, 8, 30, 10, 30, 12, 30, 446, 9, 30,
		5, 30, 448, 8, 30, 10, 30, 12, 30, 451, 9, 30, 1, 31, 1, 31, 3, 31, 455,
		8, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1,
		33, 1, 34, 1, 34, 1, 34, 1, 34, 5, 34, 471, 8, 34, 10, 34, 12, 34, 474,
		9, 34, 1, 34, 1, 34, 5, 34, 478, 8, 34, 10, 34, 12, 34, 481, 9, 34, 5,
		34, 483, 8, 34, 10, 34, 12, 34, 486, 9, 34, 1, 35, 1, 35, 1, 35, 3, 35,
		491, 8, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 5, 36, 498, 8, 36, 10, 36,
		12, 36, 501, 9, 36, 1, 36, 1, 36, 5, 36, 505, 8, 36, 10, 36, 12, 36, 508,
		9, 36, 5, 36, 510, 8, 36, 10, 36, 12, 36, 513, 9, 36, 1, 37, 1, 37, 3,
		37, 517, 8, 37, 1, 38, 1, 38, 1, 38, 5, 38, 522, 8, 38, 10, 38, 12, 38,
		525, 9, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 3, 39, 534,
		8, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 42, 1, 42, 5,
		42, 545, 8, 42, 10, 42, 12, 42, 548, 9, 42, 1, 42, 1, 42, 5, 42, 552, 8,
		42, 10, 42, 12, 42, 555, 9, 42, 5, 42, 557, 8, 42, 10, 42, 12, 42, 560,
		9, 42, 1, 42, 5, 42, 563, 8, 42, 10, 42, 12, 42, 566, 9, 42, 1, 42, 1,
		42, 1, 43, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44, 5, 44, 576, 8, 44, 10, 44,
		12, 44, 579, 9, 44, 1, 44, 1, 44, 5, 44, 583, 8, 44, 10, 44, 12, 44, 586,
		9, 44, 1, 44, 1, 44, 5, 44, 590, 8, 44, 10, 44, 12, 44, 593, 9, 44, 1,
		44, 1, 44, 5, 44, 597, 8, 44, 10, 44, 12, 44, 600, 9, 44, 1, 44, 1, 44,
		1, 45, 1, 45, 3, 45, 606, 8, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 0,
		0, 47, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
		36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70,
		72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 0, 1, 1, 0, 38, 39, 646, 0,
		103, 1, 0, 0, 0, 2, 106, 1, 0, 0, 0, 4, 109, 1, 0, 0, 0, 6, 143, 1, 0,
		0, 0, 8, 145, 1, 0, 0, 0, 10, 182, 1, 0, 0, 0, 12, 189, 1, 0, 0, 0, 14,
		191, 1, 0, 0, 0, 16, 201, 1, 0, 0, 0, 18, 205, 1, 0, 0, 0, 20, 207, 1,
		0, 0, 0, 22, 211, 1, 0, 0, 0, 24, 229, 1, 0, 0, 0, 26, 231, 1, 0, 0, 0,
		28, 241, 1, 0, 0, 0, 30, 244, 1, 0, 0, 0, 32, 279, 1, 0, 0, 0, 34, 314,
		1, 0, 0, 0, 36, 316, 1, 0, 0, 0, 38, 318, 1, 0, 0, 0, 40, 354, 1, 0, 0,
		0, 42, 358, 1, 0, 0, 0, 44, 362, 1, 0, 0, 0, 46, 366, 1, 0, 0, 0, 48, 370,
		1, 0, 0, 0, 50, 376, 1, 0, 0, 0, 52, 378, 1, 0, 0, 0, 54, 400, 1, 0, 0,
		0, 56, 424, 1, 0, 0, 0, 58, 426, 1, 0, 0, 0, 60, 430, 1, 0, 0, 0, 62, 454,
		1, 0, 0, 0, 64, 456, 1, 0, 0, 0, 66, 461, 1, 0, 0, 0, 68, 466, 1, 0, 0,
		0, 70, 490, 1, 0, 0, 0, 72, 492, 1, 0, 0, 0, 74, 516, 1, 0, 0, 0, 76, 518,
		1, 0, 0, 0, 78, 533, 1, 0, 0, 0, 80, 535, 1, 0, 0, 0, 82, 537, 1, 0, 0,
		0, 84, 542, 1, 0, 0, 0, 86, 569, 1, 0, 0, 0, 88, 573, 1, 0, 0, 0, 90, 605,
		1, 0, 0, 0, 92, 609, 1, 0, 0, 0, 94, 98, 3, 4, 2, 0, 95, 97, 5, 41, 0,
		0, 96, 95, 1, 0, 0, 0, 97, 100, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 98, 99,
		1, 0, 0, 0, 99, 102, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 101, 94, 1, 0, 0,
		0, 102, 105, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104,
		1, 1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 106, 107, 5, 1, 0, 0, 107, 108, 5,
		43, 0, 0, 108, 3, 1, 0, 0, 0, 109, 113, 3, 2, 1, 0, 110, 112, 5, 41, 0,
		0, 111, 110, 1, 0, 0, 0, 112, 115, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 113,
		114, 1, 0, 0, 0, 114, 116, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 116, 120,
		5, 2, 0, 0, 117, 119, 5, 41, 0, 0, 118, 117, 1, 0, 0, 0, 119, 122, 1, 0,
		0, 0, 120, 118, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 132, 1, 0, 0, 0,
		122, 120, 1, 0, 0, 0, 123, 127, 3, 6, 3, 0, 124, 126, 5, 41, 0, 0, 125,
		124, 1, 0, 0, 0, 126, 129, 1, 0, 0, 0, 127, 125, 1, 0, 0, 0, 127, 128,
		1, 0, 0, 0, 128, 131, 1, 0, 0, 0, 129, 127, 1, 0, 0, 0, 130, 123, 1, 0,
		0, 0, 131, 134, 1, 0, 0, 0, 132, 130, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0,
		133, 138, 1, 0, 0, 0, 134, 132, 1, 0, 0, 0, 135, 137, 5, 41, 0, 0, 136,
		135, 1, 0, 0, 0, 137, 140, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 138, 139,
		1, 0, 0, 0, 139, 141, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 141, 142, 5, 3,
		0, 0, 142, 5, 1, 0, 0, 0, 143, 144, 3, 8, 4, 0, 144, 7, 1, 0, 0, 0, 145,
		149, 5, 4, 0, 0, 146, 148, 5, 41, 0, 0, 147, 146, 1, 0, 0, 0, 148, 151,
		1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 152, 1, 0,
		0, 0, 151, 149, 1, 0, 0, 0, 152, 156, 5, 2, 0, 0, 153, 155, 5, 41, 0, 0,
		154, 153, 1, 0, 0, 0, 155, 158, 1, 0, 0, 0, 156, 154, 1, 0, 0, 0, 156,
		157, 1, 0, 0, 0, 157, 168, 1, 0, 0, 0, 158, 156, 1, 0, 0, 0, 159, 163,
		3, 10, 5, 0, 160, 162, 5, 41, 0, 0, 161, 160, 1, 0, 0, 0, 162, 165, 1,
		0, 0, 0, 163, 161, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 167, 1, 0, 0,
		0, 165, 163, 1, 0, 0, 0, 166, 159, 1, 0, 0, 0, 167, 170, 1, 0, 0, 0, 168,
		166, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169, 174, 1, 0, 0, 0, 170, 168,
		1, 0, 0, 0, 171, 173, 5, 41, 0, 0, 172, 171, 1, 0, 0, 0, 173, 176, 1, 0,
		0, 0, 174, 172, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 177, 1, 0, 0, 0,
		176, 174, 1, 0, 0, 0, 177, 178, 5, 3, 0, 0, 178, 9, 1, 0, 0, 0, 179, 183,
		3, 14, 7, 0, 180, 183, 3, 26, 13, 0, 181, 183, 3, 76, 38, 0, 182, 179,
		1, 0, 0, 0, 182, 180, 1, 0, 0, 0, 182, 181, 1, 0, 0, 0, 183, 11, 1, 0,
		0, 0, 184, 186, 5, 5, 0, 0, 185, 184, 1, 0, 0, 0, 185, 186, 1, 0, 0, 0,
		186, 187, 1, 0, 0, 0, 187, 190, 5, 38, 0, 0, 188, 190, 5, 43, 0, 0, 189,
		185, 1, 0, 0, 0, 189, 188, 1, 0, 0, 0, 190, 13, 1, 0, 0, 0, 191, 192, 5,
		43, 0, 0, 192, 196, 5, 6, 0, 0, 193, 195, 5, 41, 0, 0, 194, 193, 1, 0,
		0, 0, 195, 198, 1, 0, 0, 0, 196, 194, 1, 0, 0, 0, 196, 197, 1, 0, 0, 0,
		197, 199, 1, 0, 0, 0, 198, 196, 1, 0, 0, 0, 199, 200, 3, 12, 6, 0, 200,
		15, 1, 0, 0, 0, 201, 202, 5, 7, 0, 0, 202, 203, 5, 38, 0, 0, 203, 204,
		5, 8, 0, 0, 204, 17, 1, 0, 0, 0, 205, 206, 5, 9, 0, 0, 206, 19, 1, 0, 0,
		0, 207, 208, 5, 10, 0, 0, 208, 209, 3, 92, 46, 0, 209, 210, 5, 8, 0, 0,
		210, 21, 1, 0, 0, 0, 211, 212, 5, 10, 0, 0, 212, 213, 5, 38, 0, 0, 213,
		214, 5, 8, 0, 0, 214, 23, 1, 0, 0, 0, 215, 230, 3, 28, 14, 0, 216, 230,
		3, 36, 18, 0, 217, 230, 3, 38, 19, 0, 218, 230, 3, 40, 20, 0, 219, 230,
		3, 42, 21, 0, 220, 230, 3, 44, 22, 0, 221, 230, 3, 52, 26, 0, 222, 230,
		3, 54, 27, 0, 223, 230, 3, 68, 34, 0, 224, 230, 3, 72, 36, 0, 225, 230,
		3, 34, 17, 0, 226, 230, 3, 30, 15, 0, 227, 230, 3, 32, 16, 0, 228, 230,
		3, 60, 30, 0, 229, 215, 1, 0, 0, 0, 229, 216, 1, 0, 0, 0, 229, 217, 1,
		0, 0, 0, 229, 218, 1, 0, 0, 0, 229, 219, 1, 0, 0, 0, 229, 220, 1, 0, 0,
		0, 229, 221, 1, 0, 0, 0, 229, 222, 1, 0, 0, 0, 229, 223, 1, 0, 0, 0, 229,
		224, 1, 0, 0, 0, 229, 225, 1, 0, 0, 0, 229, 226, 1, 0, 0, 0, 229, 227,
		1, 0, 0, 0, 229, 228, 1, 0, 0, 0, 230, 25, 1, 0, 0, 0, 231, 232, 5, 43,
		0, 0, 232, 236, 5, 6, 0, 0, 233, 235, 5, 41, 0, 0, 234, 233, 1, 0, 0, 0,
		235, 238, 1, 0, 0, 0, 236, 234, 1, 0, 0, 0, 236, 237, 1, 0, 0, 0, 237,
		239, 1, 0, 0, 0, 238, 236, 1, 0, 0, 0, 239, 240, 3, 24, 12, 0, 240, 27,
		1, 0, 0, 0, 241, 242, 5, 11, 0, 0, 242, 243, 3, 24, 12, 0, 243, 29, 1,
		0, 0, 0, 244, 248, 5, 12, 0, 0, 245, 247, 5, 41, 0, 0, 246, 245, 1, 0,
		0, 0, 247, 250, 1, 0, 0, 0, 248, 246, 1, 0, 0, 0, 248, 249, 1, 0, 0, 0,
		249, 251, 1, 0, 0, 0, 250, 248, 1, 0, 0, 0, 251, 255, 5, 13, 0, 0, 252,
		254, 5, 41, 0, 0, 253, 252, 1, 0, 0, 0, 254, 257, 1, 0, 0, 0, 255, 253,
		1, 0, 0, 0, 255, 256, 1, 0, 0, 0, 256, 267, 1, 0, 0, 0, 257, 255, 1, 0,
		0, 0, 258, 262, 3, 24, 12, 0, 259, 261, 5, 41, 0, 0, 260, 259, 1, 0, 0,
		0, 261, 264, 1, 0, 0, 0, 262, 260, 1, 0, 0, 0, 262, 263, 1, 0, 0, 0, 263,
		266, 1, 0, 0, 0, 264, 262, 1, 0, 0, 0, 265, 258, 1, 0, 0, 0, 266, 269,
		1, 0, 0, 0, 267, 265, 1, 0, 0, 0, 267, 268, 1, 0, 0, 0, 268, 270, 1, 0,
		0, 0, 269, 267, 1, 0, 0, 0, 270, 274, 3, 24, 12, 0, 271, 273, 5, 41, 0,
		0, 272, 271, 1, 0, 0, 0, 273, 276, 1, 0, 0, 0, 274, 272, 1, 0, 0, 0, 274,
		275, 1, 0, 0, 0, 275, 277, 1, 0, 0, 0, 276, 274, 1, 0, 0, 0, 277, 278,
		5, 8, 0, 0, 278, 31, 1, 0, 0, 0, 279, 283, 5, 14, 0, 0, 280, 282, 5, 41,
		0, 0, 281, 280, 1, 0, 0, 0, 282, 285, 1, 0, 0, 0, 283, 281, 1, 0, 0, 0,
		283, 284, 1, 0, 0, 0, 284, 286, 1, 0, 0, 0, 285, 283, 1, 0, 0, 0, 286,
		290, 5, 13, 0, 0, 287, 289, 5, 41, 0, 0, 288, 287, 1, 0, 0, 0, 289, 292,
		1, 0, 0, 0, 290, 288, 1, 0, 0, 0, 290, 291, 1, 0, 0, 0, 291, 302, 1, 0,
		0, 0, 292, 290, 1, 0, 0, 0, 293, 297, 3, 24, 12, 0, 294, 296, 5, 41, 0,
		0, 295, 294, 1, 0, 0, 0, 296, 299, 1, 0, 0, 0, 297, 295, 1, 0, 0, 0, 297,
		298, 1, 0, 0, 0, 298, 301, 1, 0, 0, 0, 299, 297, 1, 0, 0, 0, 300, 293,
		1, 0, 0, 0, 301, 304, 1, 0, 0, 0, 302, 300, 1, 0, 0, 0, 302, 303, 1, 0,
		0, 0, 303, 305, 1, 0, 0, 0, 304, 302, 1, 0, 0, 0, 305, 309, 3, 24, 12,
		0, 306, 308, 5, 41, 0, 0, 307, 306, 1, 0, 0, 0, 308, 311, 1, 0, 0, 0, 309,
		307, 1, 0, 0, 0, 309, 310, 1, 0, 0, 0, 310, 312, 1, 0, 0, 0, 311, 309,
		1, 0, 0, 0, 312, 313, 5, 8, 0, 0, 313, 33, 1, 0, 0, 0, 314, 315, 3, 90,
		45, 0, 315, 35, 1, 0, 0, 0, 316, 317, 5, 15, 0, 0, 317, 37, 1, 0, 0, 0,
		318, 319, 5, 16, 0, 0, 319, 323, 5, 13, 0, 0, 320, 322, 5, 41, 0, 0, 321,
		320, 1, 0, 0, 0, 322, 325, 1, 0, 0, 0, 323, 321, 1, 0, 0, 0, 323, 324,
		1, 0, 0, 0, 324, 336, 1, 0, 0, 0, 325, 323, 1, 0, 0, 0, 326, 327, 3, 90,
		45, 0, 327, 331, 5, 17, 0, 0, 328, 330, 5, 41, 0, 0, 329, 328, 1, 0, 0,
		0, 330, 333, 1, 0, 0, 0, 331, 329, 1, 0, 0, 0, 331, 332, 1, 0, 0, 0, 332,
		335, 1, 0, 0, 0, 333, 331, 1, 0, 0, 0, 334, 326, 1, 0, 0, 0, 335, 338,
		1, 0, 0, 0, 336, 334, 1, 0, 0, 0, 336, 337, 1, 0, 0, 0, 337, 342, 1, 0,
		0, 0, 338, 336, 1, 0, 0, 0, 339, 341, 5, 41, 0, 0, 340, 339, 1, 0, 0, 0,
		341, 344, 1, 0, 0, 0, 342, 340, 1, 0, 0, 0, 342, 343, 1, 0, 0, 0, 343,
		345, 1, 0, 0, 0, 344, 342, 1, 0, 0, 0, 345, 349, 3, 90, 45, 0, 346, 348,
		5, 41, 0, 0, 347, 346, 1, 0, 0, 0, 348, 351, 1, 0, 0, 0, 349, 347, 1, 0,
		0, 0, 349, 350, 1, 0, 0, 0, 350, 352, 1, 0, 0, 0, 351, 349, 1, 0, 0, 0,
		352, 353, 5, 8, 0, 0, 353, 39, 1, 0, 0, 0, 354, 355, 5, 18, 0, 0, 355,
		356, 5, 13, 0, 0, 356, 357, 5, 8, 0, 0, 357, 41, 1, 0, 0, 0, 358, 359,
		5, 19, 0, 0, 359, 360, 5, 13, 0, 0, 360, 361, 5, 8, 0, 0, 361, 43, 1, 0,
		0, 0, 362, 363, 5, 20, 0, 0, 363, 364, 5, 13, 0, 0, 364, 365, 5, 8, 0,
		0, 365, 45, 1, 0, 0, 0, 366, 367, 5, 21, 0, 0, 367, 368, 3, 92, 46, 0,
		368, 369, 5, 8, 0, 0, 369, 47, 1, 0, 0, 0, 370, 371, 5, 22, 0, 0, 371,
		372, 3, 92, 46, 0, 372, 373, 5, 8, 0, 0, 373, 49, 1, 0, 0, 0, 374, 377,
		3, 48, 24, 0, 375, 377, 3, 46, 23, 0, 376, 374, 1, 0, 0, 0, 376, 375, 1,
		0, 0, 0, 377, 51, 1, 0, 0, 0, 378, 379, 5, 23, 0, 0, 379, 380, 5, 13, 0,
		0, 380, 381, 3, 90, 45, 0, 381, 385, 5, 8, 0, 0, 382, 384, 5, 41, 0, 0,
		383, 382, 1, 0, 0, 0, 384, 387, 1, 0, 0, 0, 385, 383, 1, 0, 0, 0, 385,
		386, 1, 0, 0, 0, 386, 397, 1, 0, 0, 0, 387, 385, 1, 0, 0, 0, 388, 392,
		3, 50, 25, 0, 389, 391, 5, 41, 0, 0, 390, 389, 1, 0, 0, 0, 391, 394, 1,
		0, 0, 0, 392, 390, 1, 0, 0, 0, 392, 393, 1, 0, 0, 0, 393, 396, 1, 0, 0,
		0, 394, 392, 1, 0, 0, 0, 395, 388, 1, 0, 0, 0, 396, 399, 1, 0, 0, 0, 397,
		395, 1, 0, 0, 0, 397, 398, 1, 0, 0, 0, 398, 53, 1, 0, 0, 0, 399, 397, 1,
		0, 0, 0, 400, 401, 5, 24, 0, 0, 401, 402, 5, 37, 0, 0, 402, 406, 5, 8,
		0, 0, 403, 405, 5, 41, 0, 0, 404, 403, 1, 0, 0, 0, 405, 408, 1, 0, 0, 0,
		406, 404, 1, 0, 0, 0, 406, 407, 1, 0, 0, 0, 407, 418, 1, 0, 0, 0, 408,
		406, 1, 0, 0, 0, 409, 413, 3, 56, 28, 0, 410, 412, 5, 41, 0, 0, 411, 410,
		1, 0, 0, 0, 412, 415, 1, 0, 0, 0, 413, 411, 1, 0, 0, 0, 413, 414, 1, 0,
		0, 0, 414, 417, 1, 0, 0, 0, 415, 413, 1, 0, 0, 0, 416, 409, 1, 0, 0, 0,
		417, 420, 1, 0, 0, 0, 418, 416, 1, 0, 0, 0, 418, 419, 1, 0, 0, 0, 419,
		55, 1, 0, 0, 0, 420, 418, 1, 0, 0, 0, 421, 425, 3, 16, 8, 0, 422, 425,
		3, 18, 9, 0, 423, 425, 3, 58, 29, 0, 424, 421, 1, 0, 0, 0, 424, 422, 1,
		0, 0, 0, 424, 423, 1, 0, 0, 0, 425, 57, 1, 0, 0, 0, 426, 427, 5, 25, 0,
		0, 427, 428, 5, 38, 0, 0, 428, 429, 5, 8, 0, 0, 429, 59, 1, 0, 0, 0, 430,
		431, 5, 26, 0, 0, 431, 432, 5, 13, 0, 0, 432, 433, 5, 40, 0, 0, 433, 437,
		5, 8, 0, 0, 434, 436, 5, 41, 0, 0, 435, 434, 1, 0, 0, 0, 436, 439, 1, 0,
		0, 0, 437, 435, 1, 0, 0, 0, 437, 438, 1, 0, 0, 0, 438, 449, 1, 0, 0, 0,
		439, 437, 1, 0, 0, 0, 440, 444, 3, 62, 31, 0, 441, 443, 5, 41, 0, 0, 442,
		441, 1, 0, 0, 0, 443, 446, 1, 0, 0, 0, 444, 442, 1, 0, 0, 0, 444, 445,
		1, 0, 0, 0, 445, 448, 1, 0, 0, 0, 446, 444, 1, 0, 0, 0, 447, 440, 1, 0,
		0, 0, 448, 451, 1, 0, 0, 0, 449, 447, 1, 0, 0, 0, 449, 450, 1, 0, 0, 0,
		450, 61, 1, 0, 0, 0, 451, 449, 1, 0, 0, 0, 452, 455, 3, 64, 32, 0, 453,
		455, 3, 66, 33, 0, 454, 452, 1, 0, 0, 0, 454, 453, 1, 0, 0, 0, 455, 63,
		1, 0, 0, 0, 456, 457, 5, 27, 0, 0, 457, 458, 5, 13, 0, 0, 458, 459, 3,
		12, 6, 0, 459, 460, 5, 8, 0, 0, 460, 65, 1, 0, 0, 0, 461, 462, 5, 28, 0,
		0, 462, 463, 5, 13, 0, 0, 463, 464, 3, 12, 6, 0, 464, 465, 5, 8, 0, 0,
		465, 67, 1, 0, 0, 0, 466, 467, 5, 29, 0, 0, 467, 468, 5, 13, 0, 0, 468,
		472, 5, 8, 0, 0, 469, 471, 5, 41, 0, 0, 470, 469, 1, 0, 0, 0, 471, 474,
		1, 0, 0, 0, 472, 470, 1, 0, 0, 0, 472, 473, 1, 0, 0, 0, 473, 484, 1, 0,
		0, 0, 474, 472, 1, 0, 0, 0, 475, 479, 3, 70, 35, 0, 476, 478, 5, 41, 0,
		0, 477, 476, 1, 0, 0, 0, 478, 481, 1, 0, 0, 0, 479, 477, 1, 0, 0, 0, 479,
		480, 1, 0, 0, 0, 480, 483, 1, 0, 0, 0, 481, 479, 1, 0, 0, 0, 482, 475,
		1, 0, 0, 0, 483, 486, 1, 0, 0, 0, 484, 482, 1, 0, 0, 0, 484, 485, 1, 0,
		0, 0, 485, 69, 1, 0, 0, 0, 486, 484, 1, 0, 0, 0, 487, 491, 3, 16, 8, 0,
		488, 491, 3, 18, 9, 0, 489, 491, 3, 20, 10, 0, 490, 487, 1, 0, 0, 0, 490,
		488, 1, 0, 0, 0, 490, 489, 1, 0, 0, 0, 491, 71, 1, 0, 0, 0, 492, 493, 5,
		30, 0, 0, 493, 494, 5, 13, 0, 0, 494, 495, 3, 12, 6, 0, 495, 499, 5, 8,
		0, 0, 496, 498, 5, 41, 0, 0, 497, 496, 1, 0, 0, 0, 498, 501, 1, 0, 0, 0,
		499, 497, 1, 0, 0, 0, 499, 500, 1, 0, 0, 0, 500, 511, 1, 0, 0, 0, 501,
		499, 1, 0, 0, 0, 502, 506, 3, 74, 37, 0, 503, 505, 5, 41, 0, 0, 504, 503,
		1, 0, 0, 0, 505, 508, 1, 0, 0, 0, 506, 504, 1, 0, 0, 0, 506, 507, 1, 0,
		0, 0, 507, 510, 1, 0, 0, 0, 508, 506, 1, 0, 0, 0, 509, 502, 1, 0, 0, 0,
		510, 513, 1, 0, 0, 0, 511, 509, 1, 0, 0, 0, 511, 512, 1, 0, 0, 0, 512,
		73, 1, 0, 0, 0, 513, 511, 1, 0, 0, 0, 514, 517, 3, 22, 11, 0, 515, 517,
		3, 18, 9, 0, 516, 514, 1, 0, 0, 0, 516, 515, 1, 0, 0, 0, 517, 75, 1, 0,
		0, 0, 518, 519, 5, 43, 0, 0, 519, 523, 5, 6, 0, 0, 520, 522, 5, 41, 0,
		0, 521, 520, 1, 0, 0, 0, 522, 525, 1, 0, 0, 0, 523, 521, 1, 0, 0, 0, 523,
		524, 1, 0, 0, 0, 524, 526, 1, 0, 0, 0, 525, 523, 1, 0, 0, 0, 526, 527,
		3, 78, 39, 0, 527, 77, 1, 0, 0, 0, 528, 534, 3, 82, 41, 0, 529, 534, 3,
		84, 42, 0, 530, 534, 3, 80, 40, 0, 531, 534, 3, 88, 44, 0, 532, 534, 3,
		86, 43, 0, 533, 528, 1, 0, 0, 0, 533, 529, 1, 0, 0, 0, 533, 530, 1, 0,
		0, 0, 533, 531, 1, 0, 0, 0, 533, 532, 1, 0, 0, 0, 534, 79, 1, 0, 0, 0,
		535, 536, 3, 90, 45, 0, 536, 81, 1, 0, 0, 0, 537, 538, 5, 31, 0, 0, 538,
		539, 5, 13, 0, 0, 539, 540, 3, 90, 45, 0, 540, 541, 5, 8, 0, 0, 541, 83,
		1, 0, 0, 0, 542, 546, 5, 32, 0, 0, 543, 545, 5, 41, 0, 0, 544, 543, 1,
		0, 0, 0, 545, 548, 1, 0, 0, 0, 546, 544, 1, 0, 0, 0, 546, 547, 1, 0, 0,
		0, 547, 558, 1, 0, 0, 0, 548, 546, 1, 0, 0, 0, 549, 553, 3, 78, 39, 0,
		550, 552, 5, 41, 0, 0, 551, 550, 1, 0, 0, 0, 552, 555, 1, 0, 0, 0, 553,
		551, 1, 0, 0, 0, 553, 554, 1, 0, 0, 0, 554, 557, 1, 0, 0, 0, 555, 553,
		1, 0, 0, 0, 556, 549, 1, 0, 0, 0, 557, 560, 1, 0, 0, 0, 558, 556, 1, 0,
		0, 0, 558, 559, 1, 0, 0, 0, 559, 564, 1, 0, 0, 0, 560, 558, 1, 0, 0, 0,
		561, 563, 5, 41, 0, 0, 562, 561, 1, 0, 0, 0, 563, 566, 1, 0, 0, 0, 564,
		562, 1, 0, 0, 0, 564, 565, 1, 0, 0, 0, 565, 567, 1, 0, 0, 0, 566, 564,
		1, 0, 0, 0, 567, 568, 5, 33, 0, 0, 568, 85, 1, 0, 0, 0, 569, 570, 5, 34,
		0, 0, 570, 571, 5, 13, 0, 0, 571, 572, 5, 8, 0, 0, 572, 87, 1, 0, 0, 0,
		573, 577, 5, 35, 0, 0, 574, 576, 5, 41, 0, 0, 575, 574, 1, 0, 0, 0, 576,
		579, 1, 0, 0, 0, 577, 575, 1, 0, 0, 0, 577, 578, 1, 0, 0, 0, 578, 580,
		1, 0, 0, 0, 579, 577, 1, 0, 0, 0, 580, 584, 5, 13, 0, 0, 581, 583, 5, 41,
		0, 0, 582, 581, 1, 0, 0, 0, 583, 586, 1, 0, 0, 0, 584, 582, 1, 0, 0, 0,
		584, 585, 1, 0, 0, 0, 585, 587, 1, 0, 0, 0, 586, 584, 1, 0, 0, 0, 587,
		591, 3, 24, 12, 0, 588, 590, 5, 41, 0, 0, 589, 588, 1, 0, 0, 0, 590, 593,
		1, 0, 0, 0, 591, 589, 1, 0, 0, 0, 591, 592, 1, 0, 0, 0, 592, 594, 1, 0,
		0, 0, 593, 591, 1, 0, 0, 0, 594, 598, 5, 8, 0, 0, 595, 597, 5, 41, 0, 0,
		596, 595, 1, 0, 0, 0, 597, 600, 1, 0, 0, 0, 598, 596, 1, 0, 0, 0, 598,
		599, 1, 0, 0, 0, 599, 601, 1, 0, 0, 0, 600, 598, 1, 0, 0, 0, 601, 602,
		3, 78, 39, 0, 602, 89, 1, 0, 0, 0, 603, 604, 5, 43, 0, 0, 604, 606, 5,
		36, 0, 0, 605, 603, 1, 0, 0, 0, 605, 606, 1, 0, 0, 0, 606, 607, 1, 0, 0,
		0, 607, 608, 5, 43, 0, 0, 608, 91, 1, 0, 0, 0, 609, 610, 7, 0, 0, 0, 610,
		93, 1, 0, 0, 0, 64, 98, 103, 113, 120, 127, 132, 138, 149, 156, 163, 168,
		174, 182, 185, 189, 196, 229, 236, 248, 255, 262, 267, 274, 283, 290, 297,
		302, 309, 323, 331, 336, 342, 349, 376, 385, 392, 397, 406, 413, 418, 424,
		437, 444, 449, 454, 472, 479, 484, 490, 499, 506, 511, 516, 523, 533, 546,
		553, 558, 564, 577, 584, 591, 598, 605,
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
	MinecraftMetascriptParserEOF            = antlr.TokenEOF
	MinecraftMetascriptParserT__0           = 1
	MinecraftMetascriptParserT__1           = 2
	MinecraftMetascriptParserT__2           = 3
	MinecraftMetascriptParserT__3           = 4
	MinecraftMetascriptParserT__4           = 5
	MinecraftMetascriptParserT__5           = 6
	MinecraftMetascriptParserT__6           = 7
	MinecraftMetascriptParserT__7           = 8
	MinecraftMetascriptParserT__8           = 9
	MinecraftMetascriptParserT__9           = 10
	MinecraftMetascriptParserT__10          = 11
	MinecraftMetascriptParserT__11          = 12
	MinecraftMetascriptParserT__12          = 13
	MinecraftMetascriptParserT__13          = 14
	MinecraftMetascriptParserT__14          = 15
	MinecraftMetascriptParserT__15          = 16
	MinecraftMetascriptParserT__16          = 17
	MinecraftMetascriptParserT__17          = 18
	MinecraftMetascriptParserT__18          = 19
	MinecraftMetascriptParserT__19          = 20
	MinecraftMetascriptParserT__20          = 21
	MinecraftMetascriptParserT__21          = 22
	MinecraftMetascriptParserT__22          = 23
	MinecraftMetascriptParserT__23          = 24
	MinecraftMetascriptParserT__24          = 25
	MinecraftMetascriptParserT__25          = 26
	MinecraftMetascriptParserT__26          = 27
	MinecraftMetascriptParserT__27          = 28
	MinecraftMetascriptParserT__28          = 29
	MinecraftMetascriptParserT__29          = 30
	MinecraftMetascriptParserT__30          = 31
	MinecraftMetascriptParserT__31          = 32
	MinecraftMetascriptParserT__32          = 33
	MinecraftMetascriptParserT__33          = 34
	MinecraftMetascriptParserT__34          = 35
	MinecraftMetascriptParserT__35          = 36
	MinecraftMetascriptParserStoneDepthMode = 37
	MinecraftMetascriptParserInt            = 38
	MinecraftMetascriptParserFloat          = 39
	MinecraftMetascriptParserString_        = 40
	MinecraftMetascriptParserNL             = 41
	MinecraftMetascriptParserWS             = 42
	MinecraftMetascriptParserIdentifier     = 43
	MinecraftMetascriptParserBlockComment   = 44
	MinecraftMetascriptParserLineComment    = 45
)

// MinecraftMetascriptParser rules.
const (
	MinecraftMetascriptParserRULE_script                                                 = 0
	MinecraftMetascriptParserRULE_namespaceDeclaration                                   = 1
	MinecraftMetascriptParserRULE_namespace                                              = 2
	MinecraftMetascriptParserRULE_contentBlocks                                          = 3
	MinecraftMetascriptParserRULE_surface                                                = 4
	MinecraftMetascriptParserRULE_surfaceStatement                                       = 5
	MinecraftMetascriptParserRULE_verticalAnchor                                         = 6
	MinecraftMetascriptParserRULE_verticalAnchorDeclaration                              = 7
	MinecraftMetascriptParserRULE_sharedBuilder_Offset                                   = 8
	MinecraftMetascriptParserRULE_sharedBuilder_Add                                      = 9
	MinecraftMetascriptParserRULE_sharedBuilder_Mul                                      = 10
	MinecraftMetascriptParserRULE_sharedBuilder_MulInt                                   = 11
	MinecraftMetascriptParserRULE_surfaceCondition                                       = 12
	MinecraftMetascriptParserRULE_surfaceConditionDeclaration                            = 13
	MinecraftMetascriptParserRULE_surfaceCondition_Not                                   = 14
	MinecraftMetascriptParserRULE_surfaceCondition_And                                   = 15
	MinecraftMetascriptParserRULE_surfaceCondition_Or                                    = 16
	MinecraftMetascriptParserRULE_surfaceCondition_Reference                             = 17
	MinecraftMetascriptParserRULE_surfaceCondition_AboveSurface                          = 18
	MinecraftMetascriptParserRULE_surfaceCondition_Biome                                 = 19
	MinecraftMetascriptParserRULE_surfaceCondition_Hole                                  = 20
	MinecraftMetascriptParserRULE_surfaceCondition_Steep                                 = 21
	MinecraftMetascriptParserRULE_surfaceCondition_Freezing                              = 22
	MinecraftMetascriptParserRULE_surfaceCondition_NoiseBuilder_Min                      = 23
	MinecraftMetascriptParserRULE_surfaceCondition_NoiseBuilder_Max                      = 24
	MinecraftMetascriptParserRULE_surfaceCondition_NoiseBuilder                          = 25
	MinecraftMetascriptParserRULE_surfaceCondition_Noise                                 = 26
	MinecraftMetascriptParserRULE_surfaceCondition_StoneDepth                            = 27
	MinecraftMetascriptParserRULE_surfaceCondition_StoneDepthBuilder                     = 28
	MinecraftMetascriptParserRULE_surfaceCondition_StoneDepthBuilder_SecondaryDepthRange = 29
	MinecraftMetascriptParserRULE_surfaceCondition_VerticalGradient                      = 30
	MinecraftMetascriptParserRULE_surfaceCondition_VerticalGradientBuilder               = 31
	MinecraftMetascriptParserRULE_surfaceCondition_VerticalGradientBuilder_Top           = 32
	MinecraftMetascriptParserRULE_surfaceCondition_VerticalGradientBuilder_Bottom        = 33
	MinecraftMetascriptParserRULE_surfaceCondition_AboveWater                            = 34
	MinecraftMetascriptParserRULE_surfaceCondition_AboveWaterBuilder                     = 35
	MinecraftMetascriptParserRULE_surfaceCondition_YAbove                                = 36
	MinecraftMetascriptParserRULE_surfaceCondition_YAboveBuilder                         = 37
	MinecraftMetascriptParserRULE_surfaceRuleDeclaration                                 = 38
	MinecraftMetascriptParserRULE_surfaceRule                                            = 39
	MinecraftMetascriptParserRULE_surfaceRule_Reference                                  = 40
	MinecraftMetascriptParserRULE_surfaceRule_Block                                      = 41
	MinecraftMetascriptParserRULE_surfaceRule_Sequence                                   = 42
	MinecraftMetascriptParserRULE_surfaceRule_Bandlands                                  = 43
	MinecraftMetascriptParserRULE_surfaceRule_If                                         = 44
	MinecraftMetascriptParserRULE_resourceReference                                      = 45
	MinecraftMetascriptParserRULE_number                                                 = 46
)

// IScriptContext is an interface to support dynamic dispatch.
type IScriptContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNamespace() []INamespaceContext
	Namespace(i int) INamespaceContext
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode

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
	p.RuleIndex = MinecraftMetascriptParserRULE_script
	return p
}

func InitEmptyScriptContext(p *ScriptContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_script
}

func (*ScriptContext) IsScriptContext() {}

func NewScriptContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScriptContext {
	var p = new(ScriptContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_script

	return p
}

func (s *ScriptContext) GetParser() antlr.Parser { return s.parser }

func (s *ScriptContext) AllNamespace() []INamespaceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INamespaceContext); ok {
			len++
		}
	}

	tst := make([]INamespaceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INamespaceContext); ok {
			tst[i] = t.(INamespaceContext)
			i++
		}
	}

	return tst
}

func (s *ScriptContext) Namespace(i int) INamespaceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamespaceContext); ok {
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

	return t.(INamespaceContext)
}

func (s *ScriptContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(MinecraftMetascriptParserNL)
}

func (s *ScriptContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserNL, i)
}

func (s *ScriptContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScriptContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScriptContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterScript(s)
	}
}

func (s *ScriptContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitScript(s)
	}
}

func (p *MinecraftMetascriptParser) Script() (localctx IScriptContext) {
	localctx = NewScriptContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MinecraftMetascriptParserRULE_script)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserT__0 {
		{
			p.SetState(94)
			p.Namespace()
		}
		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MinecraftMetascriptParserNL {
			{
				p.SetState(95)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(100)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(105)
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

// INamespaceDeclarationContext is an interface to support dynamic dispatch.
type INamespaceDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode

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
	p.RuleIndex = MinecraftMetascriptParserRULE_namespaceDeclaration
	return p
}

func InitEmptyNamespaceDeclarationContext(p *NamespaceDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_namespaceDeclaration
}

func (*NamespaceDeclarationContext) IsNamespaceDeclarationContext() {}

func NewNamespaceDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamespaceDeclarationContext {
	var p = new(NamespaceDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_namespaceDeclaration

	return p
}

func (s *NamespaceDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *NamespaceDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserIdentifier, 0)
}

func (s *NamespaceDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamespaceDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamespaceDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterNamespaceDeclaration(s)
	}
}

func (s *NamespaceDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitNamespaceDeclaration(s)
	}
}

func (p *MinecraftMetascriptParser) NamespaceDeclaration() (localctx INamespaceDeclarationContext) {
	localctx = NewNamespaceDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MinecraftMetascriptParserRULE_namespaceDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)
		p.Match(MinecraftMetascriptParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(107)
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

// INamespaceContext is an interface to support dynamic dispatch.
type INamespaceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NamespaceDeclaration() INamespaceDeclarationContext
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode
	AllContentBlocks() []IContentBlocksContext
	ContentBlocks(i int) IContentBlocksContext

	// IsNamespaceContext differentiates from other interfaces.
	IsNamespaceContext()
}

type NamespaceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamespaceContext() *NamespaceContext {
	var p = new(NamespaceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_namespace
	return p
}

func InitEmptyNamespaceContext(p *NamespaceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_namespace
}

func (*NamespaceContext) IsNamespaceContext() {}

func NewNamespaceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamespaceContext {
	var p = new(NamespaceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_namespace

	return p
}

func (s *NamespaceContext) GetParser() antlr.Parser { return s.parser }

func (s *NamespaceContext) NamespaceDeclaration() INamespaceDeclarationContext {
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

func (s *NamespaceContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(MinecraftMetascriptParserNL)
}

func (s *NamespaceContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserNL, i)
}

func (s *NamespaceContext) AllContentBlocks() []IContentBlocksContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IContentBlocksContext); ok {
			len++
		}
	}

	tst := make([]IContentBlocksContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IContentBlocksContext); ok {
			tst[i] = t.(IContentBlocksContext)
			i++
		}
	}

	return tst
}

func (s *NamespaceContext) ContentBlocks(i int) IContentBlocksContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContentBlocksContext); ok {
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

	return t.(IContentBlocksContext)
}

func (s *NamespaceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamespaceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamespaceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterNamespace(s)
	}
}

func (s *NamespaceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitNamespace(s)
	}
}

func (p *MinecraftMetascriptParser) Namespace() (localctx INamespaceContext) {
	localctx = NewNamespaceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MinecraftMetascriptParserRULE_namespace)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		p.NamespaceDeclaration()
	}
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(110)
			p.Match(MinecraftMetascriptParserNL)
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
		p.Match(MinecraftMetascriptParserT__1)
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
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(117)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserT__3 {
		{
			p.SetState(123)
			p.ContentBlocks()
		}
		p.SetState(127)
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
					p.SetState(124)
					p.Match(MinecraftMetascriptParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(129)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(135)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(140)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(141)
		p.Match(MinecraftMetascriptParserT__2)
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

// IContentBlocksContext is an interface to support dynamic dispatch.
type IContentBlocksContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Surface() ISurfaceContext

	// IsContentBlocksContext differentiates from other interfaces.
	IsContentBlocksContext()
}

type ContentBlocksContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContentBlocksContext() *ContentBlocksContext {
	var p = new(ContentBlocksContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_contentBlocks
	return p
}

func InitEmptyContentBlocksContext(p *ContentBlocksContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_contentBlocks
}

func (*ContentBlocksContext) IsContentBlocksContext() {}

func NewContentBlocksContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContentBlocksContext {
	var p = new(ContentBlocksContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_contentBlocks

	return p
}

func (s *ContentBlocksContext) GetParser() antlr.Parser { return s.parser }

func (s *ContentBlocksContext) Surface() ISurfaceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISurfaceContext)
}

func (s *ContentBlocksContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContentBlocksContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContentBlocksContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterContentBlocks(s)
	}
}

func (s *ContentBlocksContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitContentBlocks(s)
	}
}

func (p *MinecraftMetascriptParser) ContentBlocks() (localctx IContentBlocksContext) {
	localctx = NewContentBlocksContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MinecraftMetascriptParserRULE_contentBlocks)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Surface()
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

// ISurfaceContext is an interface to support dynamic dispatch.
type ISurfaceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode
	AllSurfaceStatement() []ISurfaceStatementContext
	SurfaceStatement(i int) ISurfaceStatementContext

	// IsSurfaceContext differentiates from other interfaces.
	IsSurfaceContext()
}

type SurfaceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceContext() *SurfaceContext {
	var p = new(SurfaceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surface
	return p
}

func InitEmptySurfaceContext(p *SurfaceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surface
}

func (*SurfaceContext) IsSurfaceContext() {}

func NewSurfaceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceContext {
	var p = new(SurfaceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_surface

	return p
}

func (s *SurfaceContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(MinecraftMetascriptParserNL)
}

func (s *SurfaceContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserNL, i)
}

func (s *SurfaceContext) AllSurfaceStatement() []ISurfaceStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISurfaceStatementContext); ok {
			len++
		}
	}

	tst := make([]ISurfaceStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISurfaceStatementContext); ok {
			tst[i] = t.(ISurfaceStatementContext)
			i++
		}
	}

	return tst
}

func (s *SurfaceContext) SurfaceStatement(i int) ISurfaceStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceStatementContext); ok {
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

	return t.(ISurfaceStatementContext)
}

func (s *SurfaceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterSurface(s)
	}
}

func (s *SurfaceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitSurface(s)
	}
}

func (p *MinecraftMetascriptParser) Surface() (localctx ISurfaceContext) {
	localctx = NewSurfaceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MinecraftMetascriptParserRULE_surface)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(145)
		p.Match(MinecraftMetascriptParserT__3)
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

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(146)
			p.Match(MinecraftMetascriptParserNL)
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
		p.Match(MinecraftMetascriptParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(156)
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
				p.SetState(153)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(158)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserIdentifier {
		{
			p.SetState(159)
			p.SurfaceStatement()
		}
		p.SetState(163)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(160)
					p.Match(MinecraftMetascriptParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(165)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

		p.SetState(170)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(171)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(177)
		p.Match(MinecraftMetascriptParserT__2)
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

// ISurfaceStatementContext is an interface to support dynamic dispatch.
type ISurfaceStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VerticalAnchorDeclaration() IVerticalAnchorDeclarationContext
	SurfaceConditionDeclaration() ISurfaceConditionDeclarationContext
	SurfaceRuleDeclaration() ISurfaceRuleDeclarationContext

	// IsSurfaceStatementContext differentiates from other interfaces.
	IsSurfaceStatementContext()
}

type SurfaceStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceStatementContext() *SurfaceStatementContext {
	var p = new(SurfaceStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceStatement
	return p
}

func InitEmptySurfaceStatementContext(p *SurfaceStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceStatement
}

func (*SurfaceStatementContext) IsSurfaceStatementContext() {}

func NewSurfaceStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceStatementContext {
	var p = new(SurfaceStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceStatement

	return p
}

func (s *SurfaceStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceStatementContext) VerticalAnchorDeclaration() IVerticalAnchorDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVerticalAnchorDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVerticalAnchorDeclarationContext)
}

func (s *SurfaceStatementContext) SurfaceConditionDeclaration() ISurfaceConditionDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceConditionDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISurfaceConditionDeclarationContext)
}

func (s *SurfaceStatementContext) SurfaceRuleDeclaration() ISurfaceRuleDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceRuleDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISurfaceRuleDeclarationContext)
}

func (s *SurfaceStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterSurfaceStatement(s)
	}
}

func (s *SurfaceStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitSurfaceStatement(s)
	}
}

func (p *MinecraftMetascriptParser) SurfaceStatement() (localctx ISurfaceStatementContext) {
	localctx = NewSurfaceStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MinecraftMetascriptParserRULE_surfaceStatement)
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(179)
			p.VerticalAnchorDeclaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(180)
			p.SurfaceConditionDeclaration()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(181)
			p.SurfaceRuleDeclaration()
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

// IVerticalAnchorContext is an interface to support dynamic dispatch.
type IVerticalAnchorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Int() antlr.TerminalNode
	Identifier() antlr.TerminalNode

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
	p.RuleIndex = MinecraftMetascriptParserRULE_verticalAnchor
	return p
}

func InitEmptyVerticalAnchorContext(p *VerticalAnchorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_verticalAnchor
}

func (*VerticalAnchorContext) IsVerticalAnchorContext() {}

func NewVerticalAnchorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VerticalAnchorContext {
	var p = new(VerticalAnchorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_verticalAnchor

	return p
}

func (s *VerticalAnchorContext) GetParser() antlr.Parser { return s.parser }

func (s *VerticalAnchorContext) Int() antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserInt, 0)
}

func (s *VerticalAnchorContext) Identifier() antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserIdentifier, 0)
}

func (s *VerticalAnchorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VerticalAnchorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VerticalAnchorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterVerticalAnchor(s)
	}
}

func (s *VerticalAnchorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitVerticalAnchor(s)
	}
}

func (p *MinecraftMetascriptParser) VerticalAnchor() (localctx IVerticalAnchorContext) {
	localctx = NewVerticalAnchorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MinecraftMetascriptParserRULE_verticalAnchor)
	var _la int

	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinecraftMetascriptParserT__4, MinecraftMetascriptParserInt:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(185)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MinecraftMetascriptParserT__4 {
			{
				p.SetState(184)
				p.Match(MinecraftMetascriptParserT__4)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(187)
			p.Match(MinecraftMetascriptParserInt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MinecraftMetascriptParserIdentifier:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(188)
			p.Match(MinecraftMetascriptParserIdentifier)
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

// IVerticalAnchorDeclarationContext is an interface to support dynamic dispatch.
type IVerticalAnchorDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	VerticalAnchor() IVerticalAnchorContext
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode

	// IsVerticalAnchorDeclarationContext differentiates from other interfaces.
	IsVerticalAnchorDeclarationContext()
}

type VerticalAnchorDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVerticalAnchorDeclarationContext() *VerticalAnchorDeclarationContext {
	var p = new(VerticalAnchorDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_verticalAnchorDeclaration
	return p
}

func InitEmptyVerticalAnchorDeclarationContext(p *VerticalAnchorDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_verticalAnchorDeclaration
}

func (*VerticalAnchorDeclarationContext) IsVerticalAnchorDeclarationContext() {}

func NewVerticalAnchorDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VerticalAnchorDeclarationContext {
	var p = new(VerticalAnchorDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_verticalAnchorDeclaration

	return p
}

func (s *VerticalAnchorDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VerticalAnchorDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserIdentifier, 0)
}

func (s *VerticalAnchorDeclarationContext) VerticalAnchor() IVerticalAnchorContext {
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

func (s *VerticalAnchorDeclarationContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(MinecraftMetascriptParserNL)
}

func (s *VerticalAnchorDeclarationContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserNL, i)
}

func (s *VerticalAnchorDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VerticalAnchorDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VerticalAnchorDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterVerticalAnchorDeclaration(s)
	}
}

func (s *VerticalAnchorDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitVerticalAnchorDeclaration(s)
	}
}

func (p *MinecraftMetascriptParser) VerticalAnchorDeclaration() (localctx IVerticalAnchorDeclarationContext) {
	localctx = NewVerticalAnchorDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MinecraftMetascriptParserRULE_verticalAnchorDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(191)
		p.Match(MinecraftMetascriptParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(192)
		p.Match(MinecraftMetascriptParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(193)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(198)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(199)
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

// ISharedBuilder_OffsetContext is an interface to support dynamic dispatch.
type ISharedBuilder_OffsetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Int() antlr.TerminalNode

	// IsSharedBuilder_OffsetContext differentiates from other interfaces.
	IsSharedBuilder_OffsetContext()
}

type SharedBuilder_OffsetContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySharedBuilder_OffsetContext() *SharedBuilder_OffsetContext {
	var p = new(SharedBuilder_OffsetContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_sharedBuilder_Offset
	return p
}

func InitEmptySharedBuilder_OffsetContext(p *SharedBuilder_OffsetContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_sharedBuilder_Offset
}

func (*SharedBuilder_OffsetContext) IsSharedBuilder_OffsetContext() {}

func NewSharedBuilder_OffsetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SharedBuilder_OffsetContext {
	var p = new(SharedBuilder_OffsetContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_sharedBuilder_Offset

	return p
}

func (s *SharedBuilder_OffsetContext) GetParser() antlr.Parser { return s.parser }

func (s *SharedBuilder_OffsetContext) Int() antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserInt, 0)
}

func (s *SharedBuilder_OffsetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SharedBuilder_OffsetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SharedBuilder_OffsetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterSharedBuilder_Offset(s)
	}
}

func (s *SharedBuilder_OffsetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitSharedBuilder_Offset(s)
	}
}

func (p *MinecraftMetascriptParser) SharedBuilder_Offset() (localctx ISharedBuilder_OffsetContext) {
	localctx = NewSharedBuilder_OffsetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MinecraftMetascriptParserRULE_sharedBuilder_Offset)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(201)
		p.Match(MinecraftMetascriptParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(202)
		p.Match(MinecraftMetascriptParserInt)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(203)
		p.Match(MinecraftMetascriptParserT__7)
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

// ISharedBuilder_AddContext is an interface to support dynamic dispatch.
type ISharedBuilder_AddContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSharedBuilder_AddContext differentiates from other interfaces.
	IsSharedBuilder_AddContext()
}

type SharedBuilder_AddContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySharedBuilder_AddContext() *SharedBuilder_AddContext {
	var p = new(SharedBuilder_AddContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_sharedBuilder_Add
	return p
}

func InitEmptySharedBuilder_AddContext(p *SharedBuilder_AddContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_sharedBuilder_Add
}

func (*SharedBuilder_AddContext) IsSharedBuilder_AddContext() {}

func NewSharedBuilder_AddContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SharedBuilder_AddContext {
	var p = new(SharedBuilder_AddContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_sharedBuilder_Add

	return p
}

func (s *SharedBuilder_AddContext) GetParser() antlr.Parser { return s.parser }
func (s *SharedBuilder_AddContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SharedBuilder_AddContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SharedBuilder_AddContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterSharedBuilder_Add(s)
	}
}

func (s *SharedBuilder_AddContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitSharedBuilder_Add(s)
	}
}

func (p *MinecraftMetascriptParser) SharedBuilder_Add() (localctx ISharedBuilder_AddContext) {
	localctx = NewSharedBuilder_AddContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MinecraftMetascriptParserRULE_sharedBuilder_Add)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)
		p.Match(MinecraftMetascriptParserT__8)
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

// ISharedBuilder_MulContext is an interface to support dynamic dispatch.
type ISharedBuilder_MulContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Number() INumberContext

	// IsSharedBuilder_MulContext differentiates from other interfaces.
	IsSharedBuilder_MulContext()
}

type SharedBuilder_MulContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySharedBuilder_MulContext() *SharedBuilder_MulContext {
	var p = new(SharedBuilder_MulContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_sharedBuilder_Mul
	return p
}

func InitEmptySharedBuilder_MulContext(p *SharedBuilder_MulContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_sharedBuilder_Mul
}

func (*SharedBuilder_MulContext) IsSharedBuilder_MulContext() {}

func NewSharedBuilder_MulContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SharedBuilder_MulContext {
	var p = new(SharedBuilder_MulContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_sharedBuilder_Mul

	return p
}

func (s *SharedBuilder_MulContext) GetParser() antlr.Parser { return s.parser }

func (s *SharedBuilder_MulContext) Number() INumberContext {
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

func (s *SharedBuilder_MulContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SharedBuilder_MulContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SharedBuilder_MulContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterSharedBuilder_Mul(s)
	}
}

func (s *SharedBuilder_MulContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitSharedBuilder_Mul(s)
	}
}

func (p *MinecraftMetascriptParser) SharedBuilder_Mul() (localctx ISharedBuilder_MulContext) {
	localctx = NewSharedBuilder_MulContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MinecraftMetascriptParserRULE_sharedBuilder_Mul)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(207)
		p.Match(MinecraftMetascriptParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(208)
		p.Number()
	}
	{
		p.SetState(209)
		p.Match(MinecraftMetascriptParserT__7)
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

// ISharedBuilder_MulIntContext is an interface to support dynamic dispatch.
type ISharedBuilder_MulIntContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Int() antlr.TerminalNode

	// IsSharedBuilder_MulIntContext differentiates from other interfaces.
	IsSharedBuilder_MulIntContext()
}

type SharedBuilder_MulIntContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySharedBuilder_MulIntContext() *SharedBuilder_MulIntContext {
	var p = new(SharedBuilder_MulIntContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_sharedBuilder_MulInt
	return p
}

func InitEmptySharedBuilder_MulIntContext(p *SharedBuilder_MulIntContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_sharedBuilder_MulInt
}

func (*SharedBuilder_MulIntContext) IsSharedBuilder_MulIntContext() {}

func NewSharedBuilder_MulIntContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SharedBuilder_MulIntContext {
	var p = new(SharedBuilder_MulIntContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_sharedBuilder_MulInt

	return p
}

func (s *SharedBuilder_MulIntContext) GetParser() antlr.Parser { return s.parser }

func (s *SharedBuilder_MulIntContext) Int() antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserInt, 0)
}

func (s *SharedBuilder_MulIntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SharedBuilder_MulIntContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SharedBuilder_MulIntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterSharedBuilder_MulInt(s)
	}
}

func (s *SharedBuilder_MulIntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitSharedBuilder_MulInt(s)
	}
}

func (p *MinecraftMetascriptParser) SharedBuilder_MulInt() (localctx ISharedBuilder_MulIntContext) {
	localctx = NewSharedBuilder_MulIntContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, MinecraftMetascriptParserRULE_sharedBuilder_MulInt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(211)
		p.Match(MinecraftMetascriptParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(212)
		p.Match(MinecraftMetascriptParserInt)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(213)
		p.Match(MinecraftMetascriptParserT__7)
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
	SurfaceCondition_Not() ISurfaceCondition_NotContext
	SurfaceCondition_AboveSurface() ISurfaceCondition_AboveSurfaceContext
	SurfaceCondition_Biome() ISurfaceCondition_BiomeContext
	SurfaceCondition_Hole() ISurfaceCondition_HoleContext
	SurfaceCondition_Steep() ISurfaceCondition_SteepContext
	SurfaceCondition_Freezing() ISurfaceCondition_FreezingContext
	SurfaceCondition_Noise() ISurfaceCondition_NoiseContext
	SurfaceCondition_StoneDepth() ISurfaceCondition_StoneDepthContext
	SurfaceCondition_AboveWater() ISurfaceCondition_AboveWaterContext
	SurfaceCondition_YAbove() ISurfaceCondition_YAboveContext
	SurfaceCondition_Reference() ISurfaceCondition_ReferenceContext
	SurfaceCondition_And() ISurfaceCondition_AndContext
	SurfaceCondition_Or() ISurfaceCondition_OrContext
	SurfaceCondition_VerticalGradient() ISurfaceCondition_VerticalGradientContext

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
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition
	return p
}

func InitEmptySurfaceConditionContext(p *SurfaceConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition
}

func (*SurfaceConditionContext) IsSurfaceConditionContext() {}

func NewSurfaceConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceConditionContext {
	var p = new(SurfaceConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition

	return p
}

func (s *SurfaceConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceConditionContext) SurfaceCondition_Not() ISurfaceCondition_NotContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceCondition_NotContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISurfaceCondition_NotContext)
}

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

func (s *SurfaceConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterSurfaceCondition(s)
	}
}

func (s *SurfaceConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitSurfaceCondition(s)
	}
}

func (p *MinecraftMetascriptParser) SurfaceCondition() (localctx ISurfaceConditionContext) {
	localctx = NewSurfaceConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, MinecraftMetascriptParserRULE_surfaceCondition)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinecraftMetascriptParserT__10:
		{
			p.SetState(215)
			p.SurfaceCondition_Not()
		}

	case MinecraftMetascriptParserT__14:
		{
			p.SetState(216)
			p.SurfaceCondition_AboveSurface()
		}

	case MinecraftMetascriptParserT__15:
		{
			p.SetState(217)
			p.SurfaceCondition_Biome()
		}

	case MinecraftMetascriptParserT__17:
		{
			p.SetState(218)
			p.SurfaceCondition_Hole()
		}

	case MinecraftMetascriptParserT__18:
		{
			p.SetState(219)
			p.SurfaceCondition_Steep()
		}

	case MinecraftMetascriptParserT__19:
		{
			p.SetState(220)
			p.SurfaceCondition_Freezing()
		}

	case MinecraftMetascriptParserT__22:
		{
			p.SetState(221)
			p.SurfaceCondition_Noise()
		}

	case MinecraftMetascriptParserT__23:
		{
			p.SetState(222)
			p.SurfaceCondition_StoneDepth()
		}

	case MinecraftMetascriptParserT__28:
		{
			p.SetState(223)
			p.SurfaceCondition_AboveWater()
		}

	case MinecraftMetascriptParserT__29:
		{
			p.SetState(224)
			p.SurfaceCondition_YAbove()
		}

	case MinecraftMetascriptParserIdentifier:
		{
			p.SetState(225)
			p.SurfaceCondition_Reference()
		}

	case MinecraftMetascriptParserT__11:
		{
			p.SetState(226)
			p.SurfaceCondition_And()
		}

	case MinecraftMetascriptParserT__13:
		{
			p.SetState(227)
			p.SurfaceCondition_Or()
		}

	case MinecraftMetascriptParserT__25:
		{
			p.SetState(228)
			p.SurfaceCondition_VerticalGradient()
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

// ISurfaceConditionDeclarationContext is an interface to support dynamic dispatch.
type ISurfaceConditionDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
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
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceConditionDeclaration
	return p
}

func InitEmptySurfaceConditionDeclarationContext(p *SurfaceConditionDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceConditionDeclaration
}

func (*SurfaceConditionDeclarationContext) IsSurfaceConditionDeclarationContext() {}

func NewSurfaceConditionDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceConditionDeclarationContext {
	var p = new(SurfaceConditionDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceConditionDeclaration

	return p
}

func (s *SurfaceConditionDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceConditionDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserIdentifier, 0)
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
	return s.GetTokens(MinecraftMetascriptParserNL)
}

func (s *SurfaceConditionDeclarationContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserNL, i)
}

func (s *SurfaceConditionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceConditionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceConditionDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterSurfaceConditionDeclaration(s)
	}
}

func (s *SurfaceConditionDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitSurfaceConditionDeclaration(s)
	}
}

func (p *MinecraftMetascriptParser) SurfaceConditionDeclaration() (localctx ISurfaceConditionDeclarationContext) {
	localctx = NewSurfaceConditionDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, MinecraftMetascriptParserRULE_surfaceConditionDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(231)
		p.Match(MinecraftMetascriptParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(232)
		p.Match(MinecraftMetascriptParserT__5)
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

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(233)
			p.Match(MinecraftMetascriptParserNL)
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

// ISurfaceCondition_NotContext is an interface to support dynamic dispatch.
type ISurfaceCondition_NotContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SurfaceCondition() ISurfaceConditionContext

	// IsSurfaceCondition_NotContext differentiates from other interfaces.
	IsSurfaceCondition_NotContext()
}

type SurfaceCondition_NotContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceCondition_NotContext() *SurfaceCondition_NotContext {
	var p = new(SurfaceCondition_NotContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_Not
	return p
}

func InitEmptySurfaceCondition_NotContext(p *SurfaceCondition_NotContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_Not
}

func (*SurfaceCondition_NotContext) IsSurfaceCondition_NotContext() {}

func NewSurfaceCondition_NotContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_NotContext {
	var p = new(SurfaceCondition_NotContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_Not

	return p
}

func (s *SurfaceCondition_NotContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceCondition_NotContext) SurfaceCondition() ISurfaceConditionContext {
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

func (s *SurfaceCondition_NotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_NotContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_NotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterSurfaceCondition_Not(s)
	}
}

func (s *SurfaceCondition_NotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitSurfaceCondition_Not(s)
	}
}

func (p *MinecraftMetascriptParser) SurfaceCondition_Not() (localctx ISurfaceCondition_NotContext) {
	localctx = NewSurfaceCondition_NotContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, MinecraftMetascriptParserRULE_surfaceCondition_Not)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(241)
		p.Match(MinecraftMetascriptParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(242)
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

// ISurfaceCondition_AndContext is an interface to support dynamic dispatch.
type ISurfaceCondition_AndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSurfaceCondition() []ISurfaceConditionContext
	SurfaceCondition(i int) ISurfaceConditionContext
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
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_And
	return p
}

func InitEmptySurfaceCondition_AndContext(p *SurfaceCondition_AndContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_And
}

func (*SurfaceCondition_AndContext) IsSurfaceCondition_AndContext() {}

func NewSurfaceCondition_AndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_AndContext {
	var p = new(SurfaceCondition_AndContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_And

	return p
}

func (s *SurfaceCondition_AndContext) GetParser() antlr.Parser { return s.parser }

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

func (s *SurfaceCondition_AndContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(MinecraftMetascriptParserNL)
}

func (s *SurfaceCondition_AndContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserNL, i)
}

func (s *SurfaceCondition_AndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_AndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_AndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterSurfaceCondition_And(s)
	}
}

func (s *SurfaceCondition_AndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitSurfaceCondition_And(s)
	}
}

func (p *MinecraftMetascriptParser) SurfaceCondition_And() (localctx ISurfaceCondition_AndContext) {
	localctx = NewSurfaceCondition_AndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, MinecraftMetascriptParserRULE_surfaceCondition_And)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(244)
		p.Match(MinecraftMetascriptParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(245)
			p.Match(MinecraftMetascriptParserNL)
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
	}
	{
		p.SetState(251)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(252)
			p.Match(MinecraftMetascriptParserNL)
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
	}
	p.SetState(267)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(258)
				p.SurfaceCondition()
			}
			p.SetState(262)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == MinecraftMetascriptParserNL {
				{
					p.SetState(259)
					p.Match(MinecraftMetascriptParserNL)
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

		}
		p.SetState(269)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(270)
		p.SurfaceCondition()
	}
	p.SetState(274)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(271)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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
		p.Match(MinecraftMetascriptParserT__7)
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
	AllSurfaceCondition() []ISurfaceConditionContext
	SurfaceCondition(i int) ISurfaceConditionContext
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
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_Or
	return p
}

func InitEmptySurfaceCondition_OrContext(p *SurfaceCondition_OrContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_Or
}

func (*SurfaceCondition_OrContext) IsSurfaceCondition_OrContext() {}

func NewSurfaceCondition_OrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_OrContext {
	var p = new(SurfaceCondition_OrContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_Or

	return p
}

func (s *SurfaceCondition_OrContext) GetParser() antlr.Parser { return s.parser }

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

func (s *SurfaceCondition_OrContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(MinecraftMetascriptParserNL)
}

func (s *SurfaceCondition_OrContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserNL, i)
}

func (s *SurfaceCondition_OrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_OrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_OrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterSurfaceCondition_Or(s)
	}
}

func (s *SurfaceCondition_OrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitSurfaceCondition_Or(s)
	}
}

func (p *MinecraftMetascriptParser) SurfaceCondition_Or() (localctx ISurfaceCondition_OrContext) {
	localctx = NewSurfaceCondition_OrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, MinecraftMetascriptParserRULE_surfaceCondition_Or)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(279)
		p.Match(MinecraftMetascriptParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(280)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(285)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(286)
		p.Match(MinecraftMetascriptParserT__12)
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

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(287)
			p.Match(MinecraftMetascriptParserNL)
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
	p.SetState(302)
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
			{
				p.SetState(293)
				p.SurfaceCondition()
			}
			p.SetState(297)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == MinecraftMetascriptParserNL {
				{
					p.SetState(294)
					p.Match(MinecraftMetascriptParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				p.SetState(299)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		p.SetState(304)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(305)
		p.SurfaceCondition()
	}
	p.SetState(309)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(306)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(311)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(312)
		p.Match(MinecraftMetascriptParserT__7)
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
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_Reference
	return p
}

func InitEmptySurfaceCondition_ReferenceContext(p *SurfaceCondition_ReferenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_Reference
}

func (*SurfaceCondition_ReferenceContext) IsSurfaceCondition_ReferenceContext() {}

func NewSurfaceCondition_ReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_ReferenceContext {
	var p = new(SurfaceCondition_ReferenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_Reference

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
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterSurfaceCondition_Reference(s)
	}
}

func (s *SurfaceCondition_ReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitSurfaceCondition_Reference(s)
	}
}

func (p *MinecraftMetascriptParser) SurfaceCondition_Reference() (localctx ISurfaceCondition_ReferenceContext) {
	localctx = NewSurfaceCondition_ReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, MinecraftMetascriptParserRULE_surfaceCondition_Reference)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(314)
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

// ISurfaceCondition_AboveSurfaceContext is an interface to support dynamic dispatch.
type ISurfaceCondition_AboveSurfaceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
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
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_AboveSurface
	return p
}

func InitEmptySurfaceCondition_AboveSurfaceContext(p *SurfaceCondition_AboveSurfaceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_AboveSurface
}

func (*SurfaceCondition_AboveSurfaceContext) IsSurfaceCondition_AboveSurfaceContext() {}

func NewSurfaceCondition_AboveSurfaceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_AboveSurfaceContext {
	var p = new(SurfaceCondition_AboveSurfaceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_AboveSurface

	return p
}

func (s *SurfaceCondition_AboveSurfaceContext) GetParser() antlr.Parser { return s.parser }
func (s *SurfaceCondition_AboveSurfaceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_AboveSurfaceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_AboveSurfaceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterSurfaceCondition_AboveSurface(s)
	}
}

func (s *SurfaceCondition_AboveSurfaceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitSurfaceCondition_AboveSurface(s)
	}
}

func (p *MinecraftMetascriptParser) SurfaceCondition_AboveSurface() (localctx ISurfaceCondition_AboveSurfaceContext) {
	localctx = NewSurfaceCondition_AboveSurfaceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, MinecraftMetascriptParserRULE_surfaceCondition_AboveSurface)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(316)
		p.Match(MinecraftMetascriptParserT__14)
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
	AllResourceReference() []IResourceReferenceContext
	ResourceReference(i int) IResourceReferenceContext
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode

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
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_Biome
	return p
}

func InitEmptySurfaceCondition_BiomeContext(p *SurfaceCondition_BiomeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_Biome
}

func (*SurfaceCondition_BiomeContext) IsSurfaceCondition_BiomeContext() {}

func NewSurfaceCondition_BiomeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_BiomeContext {
	var p = new(SurfaceCondition_BiomeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_Biome

	return p
}

func (s *SurfaceCondition_BiomeContext) GetParser() antlr.Parser { return s.parser }

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

func (s *SurfaceCondition_BiomeContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(MinecraftMetascriptParserNL)
}

func (s *SurfaceCondition_BiomeContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserNL, i)
}

func (s *SurfaceCondition_BiomeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_BiomeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_BiomeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterSurfaceCondition_Biome(s)
	}
}

func (s *SurfaceCondition_BiomeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitSurfaceCondition_Biome(s)
	}
}

func (p *MinecraftMetascriptParser) SurfaceCondition_Biome() (localctx ISurfaceCondition_BiomeContext) {
	localctx = NewSurfaceCondition_BiomeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, MinecraftMetascriptParserRULE_surfaceCondition_Biome)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(318)
		p.Match(MinecraftMetascriptParserT__15)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(319)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(323)
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
				p.SetState(320)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(325)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(336)
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
				p.SetState(326)
				p.ResourceReference()
			}
			{
				p.SetState(327)
				p.Match(MinecraftMetascriptParserT__16)
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
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
			for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				if _alt == 1 {
					{
						p.SetState(328)
						p.Match(MinecraftMetascriptParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				p.SetState(333)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext())
				if p.HasError() {
					goto errorExit
				}
			}

		}
		p.SetState(338)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(342)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(339)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(344)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(345)
		p.ResourceReference()
	}
	p.SetState(349)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(346)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(351)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(352)
		p.Match(MinecraftMetascriptParserT__7)
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
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_Hole
	return p
}

func InitEmptySurfaceCondition_HoleContext(p *SurfaceCondition_HoleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_Hole
}

func (*SurfaceCondition_HoleContext) IsSurfaceCondition_HoleContext() {}

func NewSurfaceCondition_HoleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_HoleContext {
	var p = new(SurfaceCondition_HoleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_Hole

	return p
}

func (s *SurfaceCondition_HoleContext) GetParser() antlr.Parser { return s.parser }
func (s *SurfaceCondition_HoleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_HoleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_HoleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterSurfaceCondition_Hole(s)
	}
}

func (s *SurfaceCondition_HoleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitSurfaceCondition_Hole(s)
	}
}

func (p *MinecraftMetascriptParser) SurfaceCondition_Hole() (localctx ISurfaceCondition_HoleContext) {
	localctx = NewSurfaceCondition_HoleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, MinecraftMetascriptParserRULE_surfaceCondition_Hole)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(354)
		p.Match(MinecraftMetascriptParserT__17)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(355)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(356)
		p.Match(MinecraftMetascriptParserT__7)
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
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_Steep
	return p
}

func InitEmptySurfaceCondition_SteepContext(p *SurfaceCondition_SteepContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_Steep
}

func (*SurfaceCondition_SteepContext) IsSurfaceCondition_SteepContext() {}

func NewSurfaceCondition_SteepContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_SteepContext {
	var p = new(SurfaceCondition_SteepContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_Steep

	return p
}

func (s *SurfaceCondition_SteepContext) GetParser() antlr.Parser { return s.parser }
func (s *SurfaceCondition_SteepContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_SteepContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_SteepContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterSurfaceCondition_Steep(s)
	}
}

func (s *SurfaceCondition_SteepContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitSurfaceCondition_Steep(s)
	}
}

func (p *MinecraftMetascriptParser) SurfaceCondition_Steep() (localctx ISurfaceCondition_SteepContext) {
	localctx = NewSurfaceCondition_SteepContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, MinecraftMetascriptParserRULE_surfaceCondition_Steep)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(358)
		p.Match(MinecraftMetascriptParserT__18)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(359)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(360)
		p.Match(MinecraftMetascriptParserT__7)
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
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_Freezing
	return p
}

func InitEmptySurfaceCondition_FreezingContext(p *SurfaceCondition_FreezingContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_Freezing
}

func (*SurfaceCondition_FreezingContext) IsSurfaceCondition_FreezingContext() {}

func NewSurfaceCondition_FreezingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_FreezingContext {
	var p = new(SurfaceCondition_FreezingContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_Freezing

	return p
}

func (s *SurfaceCondition_FreezingContext) GetParser() antlr.Parser { return s.parser }
func (s *SurfaceCondition_FreezingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_FreezingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_FreezingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterSurfaceCondition_Freezing(s)
	}
}

func (s *SurfaceCondition_FreezingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitSurfaceCondition_Freezing(s)
	}
}

func (p *MinecraftMetascriptParser) SurfaceCondition_Freezing() (localctx ISurfaceCondition_FreezingContext) {
	localctx = NewSurfaceCondition_FreezingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, MinecraftMetascriptParserRULE_surfaceCondition_Freezing)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(362)
		p.Match(MinecraftMetascriptParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(363)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(364)
		p.Match(MinecraftMetascriptParserT__7)
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

// ISurfaceCondition_NoiseBuilder_MinContext is an interface to support dynamic dispatch.
type ISurfaceCondition_NoiseBuilder_MinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Number() INumberContext

	// IsSurfaceCondition_NoiseBuilder_MinContext differentiates from other interfaces.
	IsSurfaceCondition_NoiseBuilder_MinContext()
}

type SurfaceCondition_NoiseBuilder_MinContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceCondition_NoiseBuilder_MinContext() *SurfaceCondition_NoiseBuilder_MinContext {
	var p = new(SurfaceCondition_NoiseBuilder_MinContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_NoiseBuilder_Min
	return p
}

func InitEmptySurfaceCondition_NoiseBuilder_MinContext(p *SurfaceCondition_NoiseBuilder_MinContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_NoiseBuilder_Min
}

func (*SurfaceCondition_NoiseBuilder_MinContext) IsSurfaceCondition_NoiseBuilder_MinContext() {}

func NewSurfaceCondition_NoiseBuilder_MinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_NoiseBuilder_MinContext {
	var p = new(SurfaceCondition_NoiseBuilder_MinContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_NoiseBuilder_Min

	return p
}

func (s *SurfaceCondition_NoiseBuilder_MinContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceCondition_NoiseBuilder_MinContext) Number() INumberContext {
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

func (s *SurfaceCondition_NoiseBuilder_MinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_NoiseBuilder_MinContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_NoiseBuilder_MinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterSurfaceCondition_NoiseBuilder_Min(s)
	}
}

func (s *SurfaceCondition_NoiseBuilder_MinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitSurfaceCondition_NoiseBuilder_Min(s)
	}
}

func (p *MinecraftMetascriptParser) SurfaceCondition_NoiseBuilder_Min() (localctx ISurfaceCondition_NoiseBuilder_MinContext) {
	localctx = NewSurfaceCondition_NoiseBuilder_MinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, MinecraftMetascriptParserRULE_surfaceCondition_NoiseBuilder_Min)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(366)
		p.Match(MinecraftMetascriptParserT__20)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(367)
		p.Number()
	}
	{
		p.SetState(368)
		p.Match(MinecraftMetascriptParserT__7)
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

// ISurfaceCondition_NoiseBuilder_MaxContext is an interface to support dynamic dispatch.
type ISurfaceCondition_NoiseBuilder_MaxContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Number() INumberContext

	// IsSurfaceCondition_NoiseBuilder_MaxContext differentiates from other interfaces.
	IsSurfaceCondition_NoiseBuilder_MaxContext()
}

type SurfaceCondition_NoiseBuilder_MaxContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceCondition_NoiseBuilder_MaxContext() *SurfaceCondition_NoiseBuilder_MaxContext {
	var p = new(SurfaceCondition_NoiseBuilder_MaxContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_NoiseBuilder_Max
	return p
}

func InitEmptySurfaceCondition_NoiseBuilder_MaxContext(p *SurfaceCondition_NoiseBuilder_MaxContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_NoiseBuilder_Max
}

func (*SurfaceCondition_NoiseBuilder_MaxContext) IsSurfaceCondition_NoiseBuilder_MaxContext() {}

func NewSurfaceCondition_NoiseBuilder_MaxContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_NoiseBuilder_MaxContext {
	var p = new(SurfaceCondition_NoiseBuilder_MaxContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_NoiseBuilder_Max

	return p
}

func (s *SurfaceCondition_NoiseBuilder_MaxContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceCondition_NoiseBuilder_MaxContext) Number() INumberContext {
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

func (s *SurfaceCondition_NoiseBuilder_MaxContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_NoiseBuilder_MaxContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_NoiseBuilder_MaxContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterSurfaceCondition_NoiseBuilder_Max(s)
	}
}

func (s *SurfaceCondition_NoiseBuilder_MaxContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitSurfaceCondition_NoiseBuilder_Max(s)
	}
}

func (p *MinecraftMetascriptParser) SurfaceCondition_NoiseBuilder_Max() (localctx ISurfaceCondition_NoiseBuilder_MaxContext) {
	localctx = NewSurfaceCondition_NoiseBuilder_MaxContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, MinecraftMetascriptParserRULE_surfaceCondition_NoiseBuilder_Max)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(370)
		p.Match(MinecraftMetascriptParserT__21)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(371)
		p.Number()
	}
	{
		p.SetState(372)
		p.Match(MinecraftMetascriptParserT__7)
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

// ISurfaceCondition_NoiseBuilderContext is an interface to support dynamic dispatch.
type ISurfaceCondition_NoiseBuilderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SurfaceCondition_NoiseBuilder_Max() ISurfaceCondition_NoiseBuilder_MaxContext
	SurfaceCondition_NoiseBuilder_Min() ISurfaceCondition_NoiseBuilder_MinContext

	// IsSurfaceCondition_NoiseBuilderContext differentiates from other interfaces.
	IsSurfaceCondition_NoiseBuilderContext()
}

type SurfaceCondition_NoiseBuilderContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceCondition_NoiseBuilderContext() *SurfaceCondition_NoiseBuilderContext {
	var p = new(SurfaceCondition_NoiseBuilderContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_NoiseBuilder
	return p
}

func InitEmptySurfaceCondition_NoiseBuilderContext(p *SurfaceCondition_NoiseBuilderContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_NoiseBuilder
}

func (*SurfaceCondition_NoiseBuilderContext) IsSurfaceCondition_NoiseBuilderContext() {}

func NewSurfaceCondition_NoiseBuilderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_NoiseBuilderContext {
	var p = new(SurfaceCondition_NoiseBuilderContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_NoiseBuilder

	return p
}

func (s *SurfaceCondition_NoiseBuilderContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceCondition_NoiseBuilderContext) SurfaceCondition_NoiseBuilder_Max() ISurfaceCondition_NoiseBuilder_MaxContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceCondition_NoiseBuilder_MaxContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISurfaceCondition_NoiseBuilder_MaxContext)
}

func (s *SurfaceCondition_NoiseBuilderContext) SurfaceCondition_NoiseBuilder_Min() ISurfaceCondition_NoiseBuilder_MinContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceCondition_NoiseBuilder_MinContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISurfaceCondition_NoiseBuilder_MinContext)
}

func (s *SurfaceCondition_NoiseBuilderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_NoiseBuilderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_NoiseBuilderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterSurfaceCondition_NoiseBuilder(s)
	}
}

func (s *SurfaceCondition_NoiseBuilderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitSurfaceCondition_NoiseBuilder(s)
	}
}

func (p *MinecraftMetascriptParser) SurfaceCondition_NoiseBuilder() (localctx ISurfaceCondition_NoiseBuilderContext) {
	localctx = NewSurfaceCondition_NoiseBuilderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, MinecraftMetascriptParserRULE_surfaceCondition_NoiseBuilder)
	p.SetState(376)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinecraftMetascriptParserT__21:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(374)
			p.SurfaceCondition_NoiseBuilder_Max()
		}

	case MinecraftMetascriptParserT__20:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(375)
			p.SurfaceCondition_NoiseBuilder_Min()
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

// ISurfaceCondition_NoiseContext is an interface to support dynamic dispatch.
type ISurfaceCondition_NoiseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ResourceReference() IResourceReferenceContext
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode
	AllSurfaceCondition_NoiseBuilder() []ISurfaceCondition_NoiseBuilderContext
	SurfaceCondition_NoiseBuilder(i int) ISurfaceCondition_NoiseBuilderContext

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
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_Noise
	return p
}

func InitEmptySurfaceCondition_NoiseContext(p *SurfaceCondition_NoiseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_Noise
}

func (*SurfaceCondition_NoiseContext) IsSurfaceCondition_NoiseContext() {}

func NewSurfaceCondition_NoiseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_NoiseContext {
	var p = new(SurfaceCondition_NoiseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_Noise

	return p
}

func (s *SurfaceCondition_NoiseContext) GetParser() antlr.Parser { return s.parser }

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

func (s *SurfaceCondition_NoiseContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(MinecraftMetascriptParserNL)
}

func (s *SurfaceCondition_NoiseContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserNL, i)
}

func (s *SurfaceCondition_NoiseContext) AllSurfaceCondition_NoiseBuilder() []ISurfaceCondition_NoiseBuilderContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISurfaceCondition_NoiseBuilderContext); ok {
			len++
		}
	}

	tst := make([]ISurfaceCondition_NoiseBuilderContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISurfaceCondition_NoiseBuilderContext); ok {
			tst[i] = t.(ISurfaceCondition_NoiseBuilderContext)
			i++
		}
	}

	return tst
}

func (s *SurfaceCondition_NoiseContext) SurfaceCondition_NoiseBuilder(i int) ISurfaceCondition_NoiseBuilderContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceCondition_NoiseBuilderContext); ok {
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

	return t.(ISurfaceCondition_NoiseBuilderContext)
}

func (s *SurfaceCondition_NoiseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_NoiseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_NoiseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterSurfaceCondition_Noise(s)
	}
}

func (s *SurfaceCondition_NoiseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitSurfaceCondition_Noise(s)
	}
}

func (p *MinecraftMetascriptParser) SurfaceCondition_Noise() (localctx ISurfaceCondition_NoiseContext) {
	localctx = NewSurfaceCondition_NoiseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, MinecraftMetascriptParserRULE_surfaceCondition_Noise)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(378)
		p.Match(MinecraftMetascriptParserT__22)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(379)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(380)
		p.ResourceReference()
	}
	{
		p.SetState(381)
		p.Match(MinecraftMetascriptParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(385)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(382)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(387)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(397)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserT__20 || _la == MinecraftMetascriptParserT__21 {
		{
			p.SetState(388)
			p.SurfaceCondition_NoiseBuilder()
		}
		p.SetState(392)
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
					p.SetState(389)
					p.Match(MinecraftMetascriptParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(394)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

		p.SetState(399)
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

// ISurfaceCondition_StoneDepthContext is an interface to support dynamic dispatch.
type ISurfaceCondition_StoneDepthContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	StoneDepthMode() antlr.TerminalNode
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode
	AllSurfaceCondition_StoneDepthBuilder() []ISurfaceCondition_StoneDepthBuilderContext
	SurfaceCondition_StoneDepthBuilder(i int) ISurfaceCondition_StoneDepthBuilderContext

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
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_StoneDepth
	return p
}

func InitEmptySurfaceCondition_StoneDepthContext(p *SurfaceCondition_StoneDepthContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_StoneDepth
}

func (*SurfaceCondition_StoneDepthContext) IsSurfaceCondition_StoneDepthContext() {}

func NewSurfaceCondition_StoneDepthContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_StoneDepthContext {
	var p = new(SurfaceCondition_StoneDepthContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_StoneDepth

	return p
}

func (s *SurfaceCondition_StoneDepthContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceCondition_StoneDepthContext) StoneDepthMode() antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserStoneDepthMode, 0)
}

func (s *SurfaceCondition_StoneDepthContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(MinecraftMetascriptParserNL)
}

func (s *SurfaceCondition_StoneDepthContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserNL, i)
}

func (s *SurfaceCondition_StoneDepthContext) AllSurfaceCondition_StoneDepthBuilder() []ISurfaceCondition_StoneDepthBuilderContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISurfaceCondition_StoneDepthBuilderContext); ok {
			len++
		}
	}

	tst := make([]ISurfaceCondition_StoneDepthBuilderContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISurfaceCondition_StoneDepthBuilderContext); ok {
			tst[i] = t.(ISurfaceCondition_StoneDepthBuilderContext)
			i++
		}
	}

	return tst
}

func (s *SurfaceCondition_StoneDepthContext) SurfaceCondition_StoneDepthBuilder(i int) ISurfaceCondition_StoneDepthBuilderContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceCondition_StoneDepthBuilderContext); ok {
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

	return t.(ISurfaceCondition_StoneDepthBuilderContext)
}

func (s *SurfaceCondition_StoneDepthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_StoneDepthContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_StoneDepthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterSurfaceCondition_StoneDepth(s)
	}
}

func (s *SurfaceCondition_StoneDepthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitSurfaceCondition_StoneDepth(s)
	}
}

func (p *MinecraftMetascriptParser) SurfaceCondition_StoneDepth() (localctx ISurfaceCondition_StoneDepthContext) {
	localctx = NewSurfaceCondition_StoneDepthContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, MinecraftMetascriptParserRULE_surfaceCondition_StoneDepth)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(400)
		p.Match(MinecraftMetascriptParserT__23)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(401)
		p.Match(MinecraftMetascriptParserStoneDepthMode)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(402)
		p.Match(MinecraftMetascriptParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(406)
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
				p.SetState(403)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(408)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(418)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&33555072) != 0 {
		{
			p.SetState(409)
			p.SurfaceCondition_StoneDepthBuilder()
		}
		p.SetState(413)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(410)
					p.Match(MinecraftMetascriptParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(415)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

		p.SetState(420)
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

// ISurfaceCondition_StoneDepthBuilderContext is an interface to support dynamic dispatch.
type ISurfaceCondition_StoneDepthBuilderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SharedBuilder_Offset() ISharedBuilder_OffsetContext
	SharedBuilder_Add() ISharedBuilder_AddContext
	SurfaceCondition_StoneDepthBuilder_SecondaryDepthRange() ISurfaceCondition_StoneDepthBuilder_SecondaryDepthRangeContext

	// IsSurfaceCondition_StoneDepthBuilderContext differentiates from other interfaces.
	IsSurfaceCondition_StoneDepthBuilderContext()
}

type SurfaceCondition_StoneDepthBuilderContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceCondition_StoneDepthBuilderContext() *SurfaceCondition_StoneDepthBuilderContext {
	var p = new(SurfaceCondition_StoneDepthBuilderContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_StoneDepthBuilder
	return p
}

func InitEmptySurfaceCondition_StoneDepthBuilderContext(p *SurfaceCondition_StoneDepthBuilderContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_StoneDepthBuilder
}

func (*SurfaceCondition_StoneDepthBuilderContext) IsSurfaceCondition_StoneDepthBuilderContext() {}

func NewSurfaceCondition_StoneDepthBuilderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_StoneDepthBuilderContext {
	var p = new(SurfaceCondition_StoneDepthBuilderContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_StoneDepthBuilder

	return p
}

func (s *SurfaceCondition_StoneDepthBuilderContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceCondition_StoneDepthBuilderContext) SharedBuilder_Offset() ISharedBuilder_OffsetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISharedBuilder_OffsetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISharedBuilder_OffsetContext)
}

func (s *SurfaceCondition_StoneDepthBuilderContext) SharedBuilder_Add() ISharedBuilder_AddContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISharedBuilder_AddContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISharedBuilder_AddContext)
}

func (s *SurfaceCondition_StoneDepthBuilderContext) SurfaceCondition_StoneDepthBuilder_SecondaryDepthRange() ISurfaceCondition_StoneDepthBuilder_SecondaryDepthRangeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceCondition_StoneDepthBuilder_SecondaryDepthRangeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISurfaceCondition_StoneDepthBuilder_SecondaryDepthRangeContext)
}

func (s *SurfaceCondition_StoneDepthBuilderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_StoneDepthBuilderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_StoneDepthBuilderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterSurfaceCondition_StoneDepthBuilder(s)
	}
}

func (s *SurfaceCondition_StoneDepthBuilderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitSurfaceCondition_StoneDepthBuilder(s)
	}
}

func (p *MinecraftMetascriptParser) SurfaceCondition_StoneDepthBuilder() (localctx ISurfaceCondition_StoneDepthBuilderContext) {
	localctx = NewSurfaceCondition_StoneDepthBuilderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, MinecraftMetascriptParserRULE_surfaceCondition_StoneDepthBuilder)
	p.SetState(424)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinecraftMetascriptParserT__6:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(421)
			p.SharedBuilder_Offset()
		}

	case MinecraftMetascriptParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(422)
			p.SharedBuilder_Add()
		}

	case MinecraftMetascriptParserT__24:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(423)
			p.SurfaceCondition_StoneDepthBuilder_SecondaryDepthRange()
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

// ISurfaceCondition_StoneDepthBuilder_SecondaryDepthRangeContext is an interface to support dynamic dispatch.
type ISurfaceCondition_StoneDepthBuilder_SecondaryDepthRangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Int() antlr.TerminalNode

	// IsSurfaceCondition_StoneDepthBuilder_SecondaryDepthRangeContext differentiates from other interfaces.
	IsSurfaceCondition_StoneDepthBuilder_SecondaryDepthRangeContext()
}

type SurfaceCondition_StoneDepthBuilder_SecondaryDepthRangeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceCondition_StoneDepthBuilder_SecondaryDepthRangeContext() *SurfaceCondition_StoneDepthBuilder_SecondaryDepthRangeContext {
	var p = new(SurfaceCondition_StoneDepthBuilder_SecondaryDepthRangeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_StoneDepthBuilder_SecondaryDepthRange
	return p
}

func InitEmptySurfaceCondition_StoneDepthBuilder_SecondaryDepthRangeContext(p *SurfaceCondition_StoneDepthBuilder_SecondaryDepthRangeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_StoneDepthBuilder_SecondaryDepthRange
}

func (*SurfaceCondition_StoneDepthBuilder_SecondaryDepthRangeContext) IsSurfaceCondition_StoneDepthBuilder_SecondaryDepthRangeContext() {
}

func NewSurfaceCondition_StoneDepthBuilder_SecondaryDepthRangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_StoneDepthBuilder_SecondaryDepthRangeContext {
	var p = new(SurfaceCondition_StoneDepthBuilder_SecondaryDepthRangeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_StoneDepthBuilder_SecondaryDepthRange

	return p
}

func (s *SurfaceCondition_StoneDepthBuilder_SecondaryDepthRangeContext) GetParser() antlr.Parser {
	return s.parser
}

func (s *SurfaceCondition_StoneDepthBuilder_SecondaryDepthRangeContext) Int() antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserInt, 0)
}

func (s *SurfaceCondition_StoneDepthBuilder_SecondaryDepthRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_StoneDepthBuilder_SecondaryDepthRangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_StoneDepthBuilder_SecondaryDepthRangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterSurfaceCondition_StoneDepthBuilder_SecondaryDepthRange(s)
	}
}

func (s *SurfaceCondition_StoneDepthBuilder_SecondaryDepthRangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitSurfaceCondition_StoneDepthBuilder_SecondaryDepthRange(s)
	}
}

func (p *MinecraftMetascriptParser) SurfaceCondition_StoneDepthBuilder_SecondaryDepthRange() (localctx ISurfaceCondition_StoneDepthBuilder_SecondaryDepthRangeContext) {
	localctx = NewSurfaceCondition_StoneDepthBuilder_SecondaryDepthRangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, MinecraftMetascriptParserRULE_surfaceCondition_StoneDepthBuilder_SecondaryDepthRange)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(426)
		p.Match(MinecraftMetascriptParserT__24)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(427)
		p.Match(MinecraftMetascriptParserInt)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(428)
		p.Match(MinecraftMetascriptParserT__7)
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
	String_() antlr.TerminalNode
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode
	AllSurfaceCondition_VerticalGradientBuilder() []ISurfaceCondition_VerticalGradientBuilderContext
	SurfaceCondition_VerticalGradientBuilder(i int) ISurfaceCondition_VerticalGradientBuilderContext

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
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_VerticalGradient
	return p
}

func InitEmptySurfaceCondition_VerticalGradientContext(p *SurfaceCondition_VerticalGradientContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_VerticalGradient
}

func (*SurfaceCondition_VerticalGradientContext) IsSurfaceCondition_VerticalGradientContext() {}

func NewSurfaceCondition_VerticalGradientContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_VerticalGradientContext {
	var p = new(SurfaceCondition_VerticalGradientContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_VerticalGradient

	return p
}

func (s *SurfaceCondition_VerticalGradientContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceCondition_VerticalGradientContext) String_() antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserString_, 0)
}

func (s *SurfaceCondition_VerticalGradientContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(MinecraftMetascriptParserNL)
}

func (s *SurfaceCondition_VerticalGradientContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserNL, i)
}

func (s *SurfaceCondition_VerticalGradientContext) AllSurfaceCondition_VerticalGradientBuilder() []ISurfaceCondition_VerticalGradientBuilderContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISurfaceCondition_VerticalGradientBuilderContext); ok {
			len++
		}
	}

	tst := make([]ISurfaceCondition_VerticalGradientBuilderContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISurfaceCondition_VerticalGradientBuilderContext); ok {
			tst[i] = t.(ISurfaceCondition_VerticalGradientBuilderContext)
			i++
		}
	}

	return tst
}

func (s *SurfaceCondition_VerticalGradientContext) SurfaceCondition_VerticalGradientBuilder(i int) ISurfaceCondition_VerticalGradientBuilderContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceCondition_VerticalGradientBuilderContext); ok {
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

	return t.(ISurfaceCondition_VerticalGradientBuilderContext)
}

func (s *SurfaceCondition_VerticalGradientContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_VerticalGradientContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_VerticalGradientContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterSurfaceCondition_VerticalGradient(s)
	}
}

func (s *SurfaceCondition_VerticalGradientContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitSurfaceCondition_VerticalGradient(s)
	}
}

func (p *MinecraftMetascriptParser) SurfaceCondition_VerticalGradient() (localctx ISurfaceCondition_VerticalGradientContext) {
	localctx = NewSurfaceCondition_VerticalGradientContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, MinecraftMetascriptParserRULE_surfaceCondition_VerticalGradient)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(430)
		p.Match(MinecraftMetascriptParserT__25)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(431)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(432)
		p.Match(MinecraftMetascriptParserString_)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(433)
		p.Match(MinecraftMetascriptParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(437)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(434)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(439)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(449)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserT__26 || _la == MinecraftMetascriptParserT__27 {
		{
			p.SetState(440)
			p.SurfaceCondition_VerticalGradientBuilder()
		}
		p.SetState(444)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(441)
					p.Match(MinecraftMetascriptParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(446)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

		p.SetState(451)
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

// ISurfaceCondition_VerticalGradientBuilderContext is an interface to support dynamic dispatch.
type ISurfaceCondition_VerticalGradientBuilderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SurfaceCondition_VerticalGradientBuilder_Top() ISurfaceCondition_VerticalGradientBuilder_TopContext
	SurfaceCondition_VerticalGradientBuilder_Bottom() ISurfaceCondition_VerticalGradientBuilder_BottomContext

	// IsSurfaceCondition_VerticalGradientBuilderContext differentiates from other interfaces.
	IsSurfaceCondition_VerticalGradientBuilderContext()
}

type SurfaceCondition_VerticalGradientBuilderContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceCondition_VerticalGradientBuilderContext() *SurfaceCondition_VerticalGradientBuilderContext {
	var p = new(SurfaceCondition_VerticalGradientBuilderContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_VerticalGradientBuilder
	return p
}

func InitEmptySurfaceCondition_VerticalGradientBuilderContext(p *SurfaceCondition_VerticalGradientBuilderContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_VerticalGradientBuilder
}

func (*SurfaceCondition_VerticalGradientBuilderContext) IsSurfaceCondition_VerticalGradientBuilderContext() {
}

func NewSurfaceCondition_VerticalGradientBuilderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_VerticalGradientBuilderContext {
	var p = new(SurfaceCondition_VerticalGradientBuilderContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_VerticalGradientBuilder

	return p
}

func (s *SurfaceCondition_VerticalGradientBuilderContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceCondition_VerticalGradientBuilderContext) SurfaceCondition_VerticalGradientBuilder_Top() ISurfaceCondition_VerticalGradientBuilder_TopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceCondition_VerticalGradientBuilder_TopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISurfaceCondition_VerticalGradientBuilder_TopContext)
}

func (s *SurfaceCondition_VerticalGradientBuilderContext) SurfaceCondition_VerticalGradientBuilder_Bottom() ISurfaceCondition_VerticalGradientBuilder_BottomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceCondition_VerticalGradientBuilder_BottomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISurfaceCondition_VerticalGradientBuilder_BottomContext)
}

func (s *SurfaceCondition_VerticalGradientBuilderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_VerticalGradientBuilderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_VerticalGradientBuilderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterSurfaceCondition_VerticalGradientBuilder(s)
	}
}

func (s *SurfaceCondition_VerticalGradientBuilderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitSurfaceCondition_VerticalGradientBuilder(s)
	}
}

func (p *MinecraftMetascriptParser) SurfaceCondition_VerticalGradientBuilder() (localctx ISurfaceCondition_VerticalGradientBuilderContext) {
	localctx = NewSurfaceCondition_VerticalGradientBuilderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, MinecraftMetascriptParserRULE_surfaceCondition_VerticalGradientBuilder)
	p.SetState(454)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinecraftMetascriptParserT__26:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(452)
			p.SurfaceCondition_VerticalGradientBuilder_Top()
		}

	case MinecraftMetascriptParserT__27:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(453)
			p.SurfaceCondition_VerticalGradientBuilder_Bottom()
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

// ISurfaceCondition_VerticalGradientBuilder_TopContext is an interface to support dynamic dispatch.
type ISurfaceCondition_VerticalGradientBuilder_TopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VerticalAnchor() IVerticalAnchorContext

	// IsSurfaceCondition_VerticalGradientBuilder_TopContext differentiates from other interfaces.
	IsSurfaceCondition_VerticalGradientBuilder_TopContext()
}

type SurfaceCondition_VerticalGradientBuilder_TopContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceCondition_VerticalGradientBuilder_TopContext() *SurfaceCondition_VerticalGradientBuilder_TopContext {
	var p = new(SurfaceCondition_VerticalGradientBuilder_TopContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_VerticalGradientBuilder_Top
	return p
}

func InitEmptySurfaceCondition_VerticalGradientBuilder_TopContext(p *SurfaceCondition_VerticalGradientBuilder_TopContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_VerticalGradientBuilder_Top
}

func (*SurfaceCondition_VerticalGradientBuilder_TopContext) IsSurfaceCondition_VerticalGradientBuilder_TopContext() {
}

func NewSurfaceCondition_VerticalGradientBuilder_TopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_VerticalGradientBuilder_TopContext {
	var p = new(SurfaceCondition_VerticalGradientBuilder_TopContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_VerticalGradientBuilder_Top

	return p
}

func (s *SurfaceCondition_VerticalGradientBuilder_TopContext) GetParser() antlr.Parser {
	return s.parser
}

func (s *SurfaceCondition_VerticalGradientBuilder_TopContext) VerticalAnchor() IVerticalAnchorContext {
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

func (s *SurfaceCondition_VerticalGradientBuilder_TopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_VerticalGradientBuilder_TopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_VerticalGradientBuilder_TopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterSurfaceCondition_VerticalGradientBuilder_Top(s)
	}
}

func (s *SurfaceCondition_VerticalGradientBuilder_TopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitSurfaceCondition_VerticalGradientBuilder_Top(s)
	}
}

func (p *MinecraftMetascriptParser) SurfaceCondition_VerticalGradientBuilder_Top() (localctx ISurfaceCondition_VerticalGradientBuilder_TopContext) {
	localctx = NewSurfaceCondition_VerticalGradientBuilder_TopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, MinecraftMetascriptParserRULE_surfaceCondition_VerticalGradientBuilder_Top)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(456)
		p.Match(MinecraftMetascriptParserT__26)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(457)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(458)
		p.VerticalAnchor()
	}
	{
		p.SetState(459)
		p.Match(MinecraftMetascriptParserT__7)
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

// ISurfaceCondition_VerticalGradientBuilder_BottomContext is an interface to support dynamic dispatch.
type ISurfaceCondition_VerticalGradientBuilder_BottomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VerticalAnchor() IVerticalAnchorContext

	// IsSurfaceCondition_VerticalGradientBuilder_BottomContext differentiates from other interfaces.
	IsSurfaceCondition_VerticalGradientBuilder_BottomContext()
}

type SurfaceCondition_VerticalGradientBuilder_BottomContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceCondition_VerticalGradientBuilder_BottomContext() *SurfaceCondition_VerticalGradientBuilder_BottomContext {
	var p = new(SurfaceCondition_VerticalGradientBuilder_BottomContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_VerticalGradientBuilder_Bottom
	return p
}

func InitEmptySurfaceCondition_VerticalGradientBuilder_BottomContext(p *SurfaceCondition_VerticalGradientBuilder_BottomContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_VerticalGradientBuilder_Bottom
}

func (*SurfaceCondition_VerticalGradientBuilder_BottomContext) IsSurfaceCondition_VerticalGradientBuilder_BottomContext() {
}

func NewSurfaceCondition_VerticalGradientBuilder_BottomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_VerticalGradientBuilder_BottomContext {
	var p = new(SurfaceCondition_VerticalGradientBuilder_BottomContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_VerticalGradientBuilder_Bottom

	return p
}

func (s *SurfaceCondition_VerticalGradientBuilder_BottomContext) GetParser() antlr.Parser {
	return s.parser
}

func (s *SurfaceCondition_VerticalGradientBuilder_BottomContext) VerticalAnchor() IVerticalAnchorContext {
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

func (s *SurfaceCondition_VerticalGradientBuilder_BottomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_VerticalGradientBuilder_BottomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_VerticalGradientBuilder_BottomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterSurfaceCondition_VerticalGradientBuilder_Bottom(s)
	}
}

func (s *SurfaceCondition_VerticalGradientBuilder_BottomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitSurfaceCondition_VerticalGradientBuilder_Bottom(s)
	}
}

func (p *MinecraftMetascriptParser) SurfaceCondition_VerticalGradientBuilder_Bottom() (localctx ISurfaceCondition_VerticalGradientBuilder_BottomContext) {
	localctx = NewSurfaceCondition_VerticalGradientBuilder_BottomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, MinecraftMetascriptParserRULE_surfaceCondition_VerticalGradientBuilder_Bottom)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(461)
		p.Match(MinecraftMetascriptParserT__27)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(462)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(463)
		p.VerticalAnchor()
	}
	{
		p.SetState(464)
		p.Match(MinecraftMetascriptParserT__7)
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

// ISurfaceCondition_AboveWaterContext is an interface to support dynamic dispatch.
type ISurfaceCondition_AboveWaterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode
	AllSurfaceCondition_AboveWaterBuilder() []ISurfaceCondition_AboveWaterBuilderContext
	SurfaceCondition_AboveWaterBuilder(i int) ISurfaceCondition_AboveWaterBuilderContext

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
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_AboveWater
	return p
}

func InitEmptySurfaceCondition_AboveWaterContext(p *SurfaceCondition_AboveWaterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_AboveWater
}

func (*SurfaceCondition_AboveWaterContext) IsSurfaceCondition_AboveWaterContext() {}

func NewSurfaceCondition_AboveWaterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_AboveWaterContext {
	var p = new(SurfaceCondition_AboveWaterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_AboveWater

	return p
}

func (s *SurfaceCondition_AboveWaterContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceCondition_AboveWaterContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(MinecraftMetascriptParserNL)
}

func (s *SurfaceCondition_AboveWaterContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserNL, i)
}

func (s *SurfaceCondition_AboveWaterContext) AllSurfaceCondition_AboveWaterBuilder() []ISurfaceCondition_AboveWaterBuilderContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISurfaceCondition_AboveWaterBuilderContext); ok {
			len++
		}
	}

	tst := make([]ISurfaceCondition_AboveWaterBuilderContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISurfaceCondition_AboveWaterBuilderContext); ok {
			tst[i] = t.(ISurfaceCondition_AboveWaterBuilderContext)
			i++
		}
	}

	return tst
}

func (s *SurfaceCondition_AboveWaterContext) SurfaceCondition_AboveWaterBuilder(i int) ISurfaceCondition_AboveWaterBuilderContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceCondition_AboveWaterBuilderContext); ok {
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

	return t.(ISurfaceCondition_AboveWaterBuilderContext)
}

func (s *SurfaceCondition_AboveWaterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_AboveWaterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_AboveWaterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterSurfaceCondition_AboveWater(s)
	}
}

func (s *SurfaceCondition_AboveWaterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitSurfaceCondition_AboveWater(s)
	}
}

func (p *MinecraftMetascriptParser) SurfaceCondition_AboveWater() (localctx ISurfaceCondition_AboveWaterContext) {
	localctx = NewSurfaceCondition_AboveWaterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, MinecraftMetascriptParserRULE_surfaceCondition_AboveWater)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(466)
		p.Match(MinecraftMetascriptParserT__28)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(467)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(468)
		p.Match(MinecraftMetascriptParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(472)
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
				p.SetState(469)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(474)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(484)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1664) != 0 {
		{
			p.SetState(475)
			p.SurfaceCondition_AboveWaterBuilder()
		}
		p.SetState(479)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 46, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(476)
					p.Match(MinecraftMetascriptParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(481)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 46, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

		p.SetState(486)
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

// ISurfaceCondition_AboveWaterBuilderContext is an interface to support dynamic dispatch.
type ISurfaceCondition_AboveWaterBuilderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SharedBuilder_Offset() ISharedBuilder_OffsetContext
	SharedBuilder_Add() ISharedBuilder_AddContext
	SharedBuilder_Mul() ISharedBuilder_MulContext

	// IsSurfaceCondition_AboveWaterBuilderContext differentiates from other interfaces.
	IsSurfaceCondition_AboveWaterBuilderContext()
}

type SurfaceCondition_AboveWaterBuilderContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceCondition_AboveWaterBuilderContext() *SurfaceCondition_AboveWaterBuilderContext {
	var p = new(SurfaceCondition_AboveWaterBuilderContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_AboveWaterBuilder
	return p
}

func InitEmptySurfaceCondition_AboveWaterBuilderContext(p *SurfaceCondition_AboveWaterBuilderContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_AboveWaterBuilder
}

func (*SurfaceCondition_AboveWaterBuilderContext) IsSurfaceCondition_AboveWaterBuilderContext() {}

func NewSurfaceCondition_AboveWaterBuilderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_AboveWaterBuilderContext {
	var p = new(SurfaceCondition_AboveWaterBuilderContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_AboveWaterBuilder

	return p
}

func (s *SurfaceCondition_AboveWaterBuilderContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceCondition_AboveWaterBuilderContext) SharedBuilder_Offset() ISharedBuilder_OffsetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISharedBuilder_OffsetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISharedBuilder_OffsetContext)
}

func (s *SurfaceCondition_AboveWaterBuilderContext) SharedBuilder_Add() ISharedBuilder_AddContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISharedBuilder_AddContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISharedBuilder_AddContext)
}

func (s *SurfaceCondition_AboveWaterBuilderContext) SharedBuilder_Mul() ISharedBuilder_MulContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISharedBuilder_MulContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISharedBuilder_MulContext)
}

func (s *SurfaceCondition_AboveWaterBuilderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_AboveWaterBuilderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_AboveWaterBuilderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterSurfaceCondition_AboveWaterBuilder(s)
	}
}

func (s *SurfaceCondition_AboveWaterBuilderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitSurfaceCondition_AboveWaterBuilder(s)
	}
}

func (p *MinecraftMetascriptParser) SurfaceCondition_AboveWaterBuilder() (localctx ISurfaceCondition_AboveWaterBuilderContext) {
	localctx = NewSurfaceCondition_AboveWaterBuilderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, MinecraftMetascriptParserRULE_surfaceCondition_AboveWaterBuilder)
	p.SetState(490)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinecraftMetascriptParserT__6:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(487)
			p.SharedBuilder_Offset()
		}

	case MinecraftMetascriptParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(488)
			p.SharedBuilder_Add()
		}

	case MinecraftMetascriptParserT__9:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(489)
			p.SharedBuilder_Mul()
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

// ISurfaceCondition_YAboveContext is an interface to support dynamic dispatch.
type ISurfaceCondition_YAboveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VerticalAnchor() IVerticalAnchorContext
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode
	AllSurfaceCondition_YAboveBuilder() []ISurfaceCondition_YAboveBuilderContext
	SurfaceCondition_YAboveBuilder(i int) ISurfaceCondition_YAboveBuilderContext

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
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_YAbove
	return p
}

func InitEmptySurfaceCondition_YAboveContext(p *SurfaceCondition_YAboveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_YAbove
}

func (*SurfaceCondition_YAboveContext) IsSurfaceCondition_YAboveContext() {}

func NewSurfaceCondition_YAboveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_YAboveContext {
	var p = new(SurfaceCondition_YAboveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_YAbove

	return p
}

func (s *SurfaceCondition_YAboveContext) GetParser() antlr.Parser { return s.parser }

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

func (s *SurfaceCondition_YAboveContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(MinecraftMetascriptParserNL)
}

func (s *SurfaceCondition_YAboveContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserNL, i)
}

func (s *SurfaceCondition_YAboveContext) AllSurfaceCondition_YAboveBuilder() []ISurfaceCondition_YAboveBuilderContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISurfaceCondition_YAboveBuilderContext); ok {
			len++
		}
	}

	tst := make([]ISurfaceCondition_YAboveBuilderContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISurfaceCondition_YAboveBuilderContext); ok {
			tst[i] = t.(ISurfaceCondition_YAboveBuilderContext)
			i++
		}
	}

	return tst
}

func (s *SurfaceCondition_YAboveContext) SurfaceCondition_YAboveBuilder(i int) ISurfaceCondition_YAboveBuilderContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceCondition_YAboveBuilderContext); ok {
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

	return t.(ISurfaceCondition_YAboveBuilderContext)
}

func (s *SurfaceCondition_YAboveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_YAboveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_YAboveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterSurfaceCondition_YAbove(s)
	}
}

func (s *SurfaceCondition_YAboveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitSurfaceCondition_YAbove(s)
	}
}

func (p *MinecraftMetascriptParser) SurfaceCondition_YAbove() (localctx ISurfaceCondition_YAboveContext) {
	localctx = NewSurfaceCondition_YAboveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, MinecraftMetascriptParserRULE_surfaceCondition_YAbove)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(492)
		p.Match(MinecraftMetascriptParserT__29)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(493)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(494)
		p.VerticalAnchor()
	}
	{
		p.SetState(495)
		p.Match(MinecraftMetascriptParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(499)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 49, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(496)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(501)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 49, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(511)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserT__8 || _la == MinecraftMetascriptParserT__9 {
		{
			p.SetState(502)
			p.SurfaceCondition_YAboveBuilder()
		}
		p.SetState(506)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 50, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(503)
					p.Match(MinecraftMetascriptParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(508)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 50, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

		p.SetState(513)
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

// ISurfaceCondition_YAboveBuilderContext is an interface to support dynamic dispatch.
type ISurfaceCondition_YAboveBuilderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SharedBuilder_MulInt() ISharedBuilder_MulIntContext
	SharedBuilder_Add() ISharedBuilder_AddContext

	// IsSurfaceCondition_YAboveBuilderContext differentiates from other interfaces.
	IsSurfaceCondition_YAboveBuilderContext()
}

type SurfaceCondition_YAboveBuilderContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceCondition_YAboveBuilderContext() *SurfaceCondition_YAboveBuilderContext {
	var p = new(SurfaceCondition_YAboveBuilderContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_YAboveBuilder
	return p
}

func InitEmptySurfaceCondition_YAboveBuilderContext(p *SurfaceCondition_YAboveBuilderContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_YAboveBuilder
}

func (*SurfaceCondition_YAboveBuilderContext) IsSurfaceCondition_YAboveBuilderContext() {}

func NewSurfaceCondition_YAboveBuilderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_YAboveBuilderContext {
	var p = new(SurfaceCondition_YAboveBuilderContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_YAboveBuilder

	return p
}

func (s *SurfaceCondition_YAboveBuilderContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceCondition_YAboveBuilderContext) SharedBuilder_MulInt() ISharedBuilder_MulIntContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISharedBuilder_MulIntContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISharedBuilder_MulIntContext)
}

func (s *SurfaceCondition_YAboveBuilderContext) SharedBuilder_Add() ISharedBuilder_AddContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISharedBuilder_AddContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISharedBuilder_AddContext)
}

func (s *SurfaceCondition_YAboveBuilderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_YAboveBuilderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_YAboveBuilderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterSurfaceCondition_YAboveBuilder(s)
	}
}

func (s *SurfaceCondition_YAboveBuilderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitSurfaceCondition_YAboveBuilder(s)
	}
}

func (p *MinecraftMetascriptParser) SurfaceCondition_YAboveBuilder() (localctx ISurfaceCondition_YAboveBuilderContext) {
	localctx = NewSurfaceCondition_YAboveBuilderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, MinecraftMetascriptParserRULE_surfaceCondition_YAboveBuilder)
	p.SetState(516)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinecraftMetascriptParserT__9:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(514)
			p.SharedBuilder_MulInt()
		}

	case MinecraftMetascriptParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(515)
			p.SharedBuilder_Add()
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

// ISurfaceRuleDeclarationContext is an interface to support dynamic dispatch.
type ISurfaceRuleDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
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
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceRuleDeclaration
	return p
}

func InitEmptySurfaceRuleDeclarationContext(p *SurfaceRuleDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceRuleDeclaration
}

func (*SurfaceRuleDeclarationContext) IsSurfaceRuleDeclarationContext() {}

func NewSurfaceRuleDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceRuleDeclarationContext {
	var p = new(SurfaceRuleDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceRuleDeclaration

	return p
}

func (s *SurfaceRuleDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceRuleDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserIdentifier, 0)
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
	return s.GetTokens(MinecraftMetascriptParserNL)
}

func (s *SurfaceRuleDeclarationContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserNL, i)
}

func (s *SurfaceRuleDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceRuleDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceRuleDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterSurfaceRuleDeclaration(s)
	}
}

func (s *SurfaceRuleDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitSurfaceRuleDeclaration(s)
	}
}

func (p *MinecraftMetascriptParser) SurfaceRuleDeclaration() (localctx ISurfaceRuleDeclarationContext) {
	localctx = NewSurfaceRuleDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, MinecraftMetascriptParserRULE_surfaceRuleDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(518)
		p.Match(MinecraftMetascriptParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(519)
		p.Match(MinecraftMetascriptParserT__5)
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

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(520)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(525)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(526)
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
	SurfaceRule_Block() ISurfaceRule_BlockContext
	SurfaceRule_Sequence() ISurfaceRule_SequenceContext
	SurfaceRule_Reference() ISurfaceRule_ReferenceContext
	SurfaceRule_If() ISurfaceRule_IfContext
	SurfaceRule_Bandlands() ISurfaceRule_BandlandsContext

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
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceRule
	return p
}

func InitEmptySurfaceRuleContext(p *SurfaceRuleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceRule
}

func (*SurfaceRuleContext) IsSurfaceRuleContext() {}

func NewSurfaceRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceRuleContext {
	var p = new(SurfaceRuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceRule

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

func (s *SurfaceRuleContext) SurfaceRule_If() ISurfaceRule_IfContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceRule_IfContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISurfaceRule_IfContext)
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

func (s *SurfaceRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterSurfaceRule(s)
	}
}

func (s *SurfaceRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitSurfaceRule(s)
	}
}

func (p *MinecraftMetascriptParser) SurfaceRule() (localctx ISurfaceRuleContext) {
	localctx = NewSurfaceRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, MinecraftMetascriptParserRULE_surfaceRule)
	p.SetState(533)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinecraftMetascriptParserT__30:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(528)
			p.SurfaceRule_Block()
		}

	case MinecraftMetascriptParserT__31:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(529)
			p.SurfaceRule_Sequence()
		}

	case MinecraftMetascriptParserIdentifier:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(530)
			p.SurfaceRule_Reference()
		}

	case MinecraftMetascriptParserT__34:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(531)
			p.SurfaceRule_If()
		}

	case MinecraftMetascriptParserT__33:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(532)
			p.SurfaceRule_Bandlands()
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
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceRule_Reference
	return p
}

func InitEmptySurfaceRule_ReferenceContext(p *SurfaceRule_ReferenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceRule_Reference
}

func (*SurfaceRule_ReferenceContext) IsSurfaceRule_ReferenceContext() {}

func NewSurfaceRule_ReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceRule_ReferenceContext {
	var p = new(SurfaceRule_ReferenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceRule_Reference

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
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterSurfaceRule_Reference(s)
	}
}

func (s *SurfaceRule_ReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitSurfaceRule_Reference(s)
	}
}

func (p *MinecraftMetascriptParser) SurfaceRule_Reference() (localctx ISurfaceRule_ReferenceContext) {
	localctx = NewSurfaceRule_ReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, MinecraftMetascriptParserRULE_surfaceRule_Reference)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(535)
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

// ISurfaceRule_BlockContext is an interface to support dynamic dispatch.
type ISurfaceRule_BlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
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
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceRule_Block
	return p
}

func InitEmptySurfaceRule_BlockContext(p *SurfaceRule_BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceRule_Block
}

func (*SurfaceRule_BlockContext) IsSurfaceRule_BlockContext() {}

func NewSurfaceRule_BlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceRule_BlockContext {
	var p = new(SurfaceRule_BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceRule_Block

	return p
}

func (s *SurfaceRule_BlockContext) GetParser() antlr.Parser { return s.parser }

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
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterSurfaceRule_Block(s)
	}
}

func (s *SurfaceRule_BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitSurfaceRule_Block(s)
	}
}

func (p *MinecraftMetascriptParser) SurfaceRule_Block() (localctx ISurfaceRule_BlockContext) {
	localctx = NewSurfaceRule_BlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, MinecraftMetascriptParserRULE_surfaceRule_Block)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(537)
		p.Match(MinecraftMetascriptParserT__30)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(538)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(539)
		p.ResourceReference()
	}
	{
		p.SetState(540)
		p.Match(MinecraftMetascriptParserT__7)
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

// ISurfaceRule_SequenceContext is an interface to support dynamic dispatch.
type ISurfaceRule_SequenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
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
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceRule_Sequence
	return p
}

func InitEmptySurfaceRule_SequenceContext(p *SurfaceRule_SequenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceRule_Sequence
}

func (*SurfaceRule_SequenceContext) IsSurfaceRule_SequenceContext() {}

func NewSurfaceRule_SequenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceRule_SequenceContext {
	var p = new(SurfaceRule_SequenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceRule_Sequence

	return p
}

func (s *SurfaceRule_SequenceContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceRule_SequenceContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(MinecraftMetascriptParserNL)
}

func (s *SurfaceRule_SequenceContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserNL, i)
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
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterSurfaceRule_Sequence(s)
	}
}

func (s *SurfaceRule_SequenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitSurfaceRule_Sequence(s)
	}
}

func (p *MinecraftMetascriptParser) SurfaceRule_Sequence() (localctx ISurfaceRule_SequenceContext) {
	localctx = NewSurfaceRule_SequenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, MinecraftMetascriptParserRULE_surfaceRule_Sequence)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(542)
		p.Match(MinecraftMetascriptParserT__31)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(546)
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
				p.SetState(543)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(548)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 55, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(558)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8854075080704) != 0 {
		{
			p.SetState(549)
			p.SurfaceRule()
		}
		p.SetState(553)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 56, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(550)
					p.Match(MinecraftMetascriptParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(555)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 56, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

		p.SetState(560)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(564)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(561)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(566)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(567)
		p.Match(MinecraftMetascriptParserT__32)
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

// ISurfaceRule_BandlandsContext is an interface to support dynamic dispatch.
type ISurfaceRule_BandlandsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
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
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceRule_Bandlands
	return p
}

func InitEmptySurfaceRule_BandlandsContext(p *SurfaceRule_BandlandsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceRule_Bandlands
}

func (*SurfaceRule_BandlandsContext) IsSurfaceRule_BandlandsContext() {}

func NewSurfaceRule_BandlandsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceRule_BandlandsContext {
	var p = new(SurfaceRule_BandlandsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceRule_Bandlands

	return p
}

func (s *SurfaceRule_BandlandsContext) GetParser() antlr.Parser { return s.parser }
func (s *SurfaceRule_BandlandsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceRule_BandlandsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceRule_BandlandsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterSurfaceRule_Bandlands(s)
	}
}

func (s *SurfaceRule_BandlandsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitSurfaceRule_Bandlands(s)
	}
}

func (p *MinecraftMetascriptParser) SurfaceRule_Bandlands() (localctx ISurfaceRule_BandlandsContext) {
	localctx = NewSurfaceRule_BandlandsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, MinecraftMetascriptParserRULE_surfaceRule_Bandlands)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(569)
		p.Match(MinecraftMetascriptParserT__33)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(570)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(571)
		p.Match(MinecraftMetascriptParserT__7)
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

// ISurfaceRule_IfContext is an interface to support dynamic dispatch.
type ISurfaceRule_IfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SurfaceCondition() ISurfaceConditionContext
	SurfaceRule() ISurfaceRuleContext
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode

	// IsSurfaceRule_IfContext differentiates from other interfaces.
	IsSurfaceRule_IfContext()
}

type SurfaceRule_IfContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceRule_IfContext() *SurfaceRule_IfContext {
	var p = new(SurfaceRule_IfContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceRule_If
	return p
}

func InitEmptySurfaceRule_IfContext(p *SurfaceRule_IfContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceRule_If
}

func (*SurfaceRule_IfContext) IsSurfaceRule_IfContext() {}

func NewSurfaceRule_IfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceRule_IfContext {
	var p = new(SurfaceRule_IfContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceRule_If

	return p
}

func (s *SurfaceRule_IfContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceRule_IfContext) SurfaceCondition() ISurfaceConditionContext {
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

func (s *SurfaceRule_IfContext) SurfaceRule() ISurfaceRuleContext {
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

func (s *SurfaceRule_IfContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(MinecraftMetascriptParserNL)
}

func (s *SurfaceRule_IfContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserNL, i)
}

func (s *SurfaceRule_IfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceRule_IfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceRule_IfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterSurfaceRule_If(s)
	}
}

func (s *SurfaceRule_IfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitSurfaceRule_If(s)
	}
}

func (p *MinecraftMetascriptParser) SurfaceRule_If() (localctx ISurfaceRule_IfContext) {
	localctx = NewSurfaceRule_IfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, MinecraftMetascriptParserRULE_surfaceRule_If)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(573)
		p.Match(MinecraftMetascriptParserT__34)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(577)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(574)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(579)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(580)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(584)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(581)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(586)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(587)
		p.SurfaceCondition()
	}
	p.SetState(591)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(588)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(593)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(594)
		p.Match(MinecraftMetascriptParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(598)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(595)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(600)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(601)
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
	p.EnterRule(localctx, 90, MinecraftMetascriptParserRULE_resourceReference)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(605)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 63, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(603)
			p.Match(MinecraftMetascriptParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(604)
			p.Match(MinecraftMetascriptParserT__35)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(607)
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
	p.EnterRule(localctx, 92, MinecraftMetascriptParserRULE_number)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(609)
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
