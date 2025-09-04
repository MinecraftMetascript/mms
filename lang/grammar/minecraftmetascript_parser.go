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
		"'Noise'", "'StoneDepth('", "'.secondaryDepthRange('", "'VerticalGradient'",
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
		4, 1, 45, 594, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
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
		6, 190, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1,
		9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 3, 12, 224, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14,
		1, 14, 1, 15, 1, 15, 5, 15, 235, 8, 15, 10, 15, 12, 15, 238, 9, 15, 1,
		15, 1, 15, 5, 15, 242, 8, 15, 10, 15, 12, 15, 245, 9, 15, 1, 15, 1, 15,
		5, 15, 249, 8, 15, 10, 15, 12, 15, 252, 9, 15, 5, 15, 254, 8, 15, 10, 15,
		12, 15, 257, 9, 15, 1, 15, 1, 15, 5, 15, 261, 8, 15, 10, 15, 12, 15, 264,
		9, 15, 1, 15, 1, 15, 1, 16, 1, 16, 5, 16, 270, 8, 16, 10, 16, 12, 16, 273,
		9, 16, 1, 16, 1, 16, 5, 16, 277, 8, 16, 10, 16, 12, 16, 280, 9, 16, 1,
		16, 1, 16, 5, 16, 284, 8, 16, 10, 16, 12, 16, 287, 9, 16, 5, 16, 289, 8,
		16, 10, 16, 12, 16, 292, 9, 16, 1, 16, 1, 16, 5, 16, 296, 8, 16, 10, 16,
		12, 16, 299, 9, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1,
		19, 1, 19, 5, 19, 310, 8, 19, 10, 19, 12, 19, 313, 9, 19, 1, 19, 1, 19,
		1, 19, 5, 19, 318, 8, 19, 10, 19, 12, 19, 321, 9, 19, 5, 19, 323, 8, 19,
		10, 19, 12, 19, 326, 9, 19, 1, 19, 5, 19, 329, 8, 19, 10, 19, 12, 19, 332,
		9, 19, 1, 19, 1, 19, 5, 19, 336, 8, 19, 10, 19, 12, 19, 339, 9, 19, 1,
		19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 25, 1, 25, 3, 25, 365, 8, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		5, 26, 372, 8, 26, 10, 26, 12, 26, 375, 9, 26, 1, 26, 1, 26, 5, 26, 379,
		8, 26, 10, 26, 12, 26, 382, 9, 26, 5, 26, 384, 8, 26, 10, 26, 12, 26, 387,
		9, 26, 1, 27, 1, 27, 1, 27, 1, 27, 5, 27, 393, 8, 27, 10, 27, 12, 27, 396,
		9, 27, 1, 27, 1, 27, 5, 27, 400, 8, 27, 10, 27, 12, 27, 403, 9, 27, 5,
		27, 405, 8, 27, 10, 27, 12, 27, 408, 9, 27, 1, 28, 1, 28, 1, 28, 3, 28,
		413, 8, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1,
		30, 5, 30, 424, 8, 30, 10, 30, 12, 30, 427, 9, 30, 1, 30, 1, 30, 5, 30,
		431, 8, 30, 10, 30, 12, 30, 434, 9, 30, 5, 30, 436, 8, 30, 10, 30, 12,
		30, 439, 9, 30, 1, 31, 1, 31, 3, 31, 443, 8, 31, 1, 32, 1, 32, 1, 32, 1,
		32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34,
		5, 34, 459, 8, 34, 10, 34, 12, 34, 462, 9, 34, 1, 34, 1, 34, 5, 34, 466,
		8, 34, 10, 34, 12, 34, 469, 9, 34, 5, 34, 471, 8, 34, 10, 34, 12, 34, 474,
		9, 34, 1, 35, 1, 35, 1, 35, 3, 35, 479, 8, 35, 1, 36, 1, 36, 1, 36, 1,
		36, 1, 36, 5, 36, 486, 8, 36, 10, 36, 12, 36, 489, 9, 36, 1, 36, 1, 36,
		5, 36, 493, 8, 36, 10, 36, 12, 36, 496, 9, 36, 5, 36, 498, 8, 36, 10, 36,
		12, 36, 501, 9, 36, 1, 37, 1, 37, 3, 37, 505, 8, 37, 1, 38, 1, 38, 1, 38,
		1, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 3, 39, 516, 8, 39, 1, 40, 1,
		40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 42, 1, 42, 5, 42, 527, 8, 42,
		10, 42, 12, 42, 530, 9, 42, 1, 42, 1, 42, 5, 42, 534, 8, 42, 10, 42, 12,
		42, 537, 9, 42, 5, 42, 539, 8, 42, 10, 42, 12, 42, 542, 9, 42, 1, 42, 5,
		42, 545, 8, 42, 10, 42, 12, 42, 548, 9, 42, 1, 42, 1, 42, 1, 43, 1, 43,
		1, 43, 1, 43, 1, 44, 1, 44, 5, 44, 558, 8, 44, 10, 44, 12, 44, 561, 9,
		44, 1, 44, 1, 44, 5, 44, 565, 8, 44, 10, 44, 12, 44, 568, 9, 44, 1, 44,
		1, 44, 5, 44, 572, 8, 44, 10, 44, 12, 44, 575, 9, 44, 1, 44, 1, 44, 5,
		44, 579, 8, 44, 10, 44, 12, 44, 582, 9, 44, 1, 44, 1, 44, 1, 45, 1, 45,
		3, 45, 588, 8, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 0, 0, 47, 0, 2, 4,
		6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
		44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78,
		80, 82, 84, 86, 88, 90, 92, 0, 1, 1, 0, 38, 39, 625, 0, 103, 1, 0, 0, 0,
		2, 106, 1, 0, 0, 0, 4, 109, 1, 0, 0, 0, 6, 143, 1, 0, 0, 0, 8, 145, 1,
		0, 0, 0, 10, 182, 1, 0, 0, 0, 12, 189, 1, 0, 0, 0, 14, 191, 1, 0, 0, 0,
		16, 195, 1, 0, 0, 0, 18, 199, 1, 0, 0, 0, 20, 201, 1, 0, 0, 0, 22, 205,
		1, 0, 0, 0, 24, 223, 1, 0, 0, 0, 26, 225, 1, 0, 0, 0, 28, 229, 1, 0, 0,
		0, 30, 232, 1, 0, 0, 0, 32, 267, 1, 0, 0, 0, 34, 302, 1, 0, 0, 0, 36, 304,
		1, 0, 0, 0, 38, 306, 1, 0, 0, 0, 40, 342, 1, 0, 0, 0, 42, 346, 1, 0, 0,
		0, 44, 350, 1, 0, 0, 0, 46, 354, 1, 0, 0, 0, 48, 358, 1, 0, 0, 0, 50, 364,
		1, 0, 0, 0, 52, 366, 1, 0, 0, 0, 54, 388, 1, 0, 0, 0, 56, 412, 1, 0, 0,
		0, 58, 414, 1, 0, 0, 0, 60, 418, 1, 0, 0, 0, 62, 442, 1, 0, 0, 0, 64, 444,
		1, 0, 0, 0, 66, 449, 1, 0, 0, 0, 68, 454, 1, 0, 0, 0, 70, 478, 1, 0, 0,
		0, 72, 480, 1, 0, 0, 0, 74, 504, 1, 0, 0, 0, 76, 506, 1, 0, 0, 0, 78, 515,
		1, 0, 0, 0, 80, 517, 1, 0, 0, 0, 82, 519, 1, 0, 0, 0, 84, 524, 1, 0, 0,
		0, 86, 551, 1, 0, 0, 0, 88, 555, 1, 0, 0, 0, 90, 587, 1, 0, 0, 0, 92, 591,
		1, 0, 0, 0, 94, 98, 3, 4, 2, 0, 95, 97, 5, 41, 0, 0, 96, 95, 1, 0, 0, 0,
		97, 100, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 102, 1,
		0, 0, 0, 100, 98, 1, 0, 0, 0, 101, 94, 1, 0, 0, 0, 102, 105, 1, 0, 0, 0,
		103, 101, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 1, 1, 0, 0, 0, 105, 103,
		1, 0, 0, 0, 106, 107, 5, 1, 0, 0, 107, 108, 5, 43, 0, 0, 108, 3, 1, 0,
		0, 0, 109, 113, 3, 2, 1, 0, 110, 112, 5, 41, 0, 0, 111, 110, 1, 0, 0, 0,
		112, 115, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114,
		116, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 116, 120, 5, 2, 0, 0, 117, 119,
		5, 41, 0, 0, 118, 117, 1, 0, 0, 0, 119, 122, 1, 0, 0, 0, 120, 118, 1, 0,
		0, 0, 120, 121, 1, 0, 0, 0, 121, 132, 1, 0, 0, 0, 122, 120, 1, 0, 0, 0,
		123, 127, 3, 6, 3, 0, 124, 126, 5, 41, 0, 0, 125, 124, 1, 0, 0, 0, 126,
		129, 1, 0, 0, 0, 127, 125, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 131,
		1, 0, 0, 0, 129, 127, 1, 0, 0, 0, 130, 123, 1, 0, 0, 0, 131, 134, 1, 0,
		0, 0, 132, 130, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 138, 1, 0, 0, 0,
		134, 132, 1, 0, 0, 0, 135, 137, 5, 41, 0, 0, 136, 135, 1, 0, 0, 0, 137,
		140, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 141,
		1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 141, 142, 5, 3, 0, 0, 142, 5, 1, 0, 0,
		0, 143, 144, 3, 8, 4, 0, 144, 7, 1, 0, 0, 0, 145, 149, 5, 4, 0, 0, 146,
		148, 5, 41, 0, 0, 147, 146, 1, 0, 0, 0, 148, 151, 1, 0, 0, 0, 149, 147,
		1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 152, 1, 0, 0, 0, 151, 149, 1, 0,
		0, 0, 152, 156, 5, 2, 0, 0, 153, 155, 5, 41, 0, 0, 154, 153, 1, 0, 0, 0,
		155, 158, 1, 0, 0, 0, 156, 154, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157,
		168, 1, 0, 0, 0, 158, 156, 1, 0, 0, 0, 159, 163, 3, 10, 5, 0, 160, 162,
		5, 41, 0, 0, 161, 160, 1, 0, 0, 0, 162, 165, 1, 0, 0, 0, 163, 161, 1, 0,
		0, 0, 163, 164, 1, 0, 0, 0, 164, 167, 1, 0, 0, 0, 165, 163, 1, 0, 0, 0,
		166, 159, 1, 0, 0, 0, 167, 170, 1, 0, 0, 0, 168, 166, 1, 0, 0, 0, 168,
		169, 1, 0, 0, 0, 169, 174, 1, 0, 0, 0, 170, 168, 1, 0, 0, 0, 171, 173,
		5, 41, 0, 0, 172, 171, 1, 0, 0, 0, 173, 176, 1, 0, 0, 0, 174, 172, 1, 0,
		0, 0, 174, 175, 1, 0, 0, 0, 175, 177, 1, 0, 0, 0, 176, 174, 1, 0, 0, 0,
		177, 178, 5, 3, 0, 0, 178, 9, 1, 0, 0, 0, 179, 183, 3, 14, 7, 0, 180, 183,
		3, 26, 13, 0, 181, 183, 3, 76, 38, 0, 182, 179, 1, 0, 0, 0, 182, 180, 1,
		0, 0, 0, 182, 181, 1, 0, 0, 0, 183, 11, 1, 0, 0, 0, 184, 186, 5, 5, 0,
		0, 185, 184, 1, 0, 0, 0, 185, 186, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187,
		190, 5, 38, 0, 0, 188, 190, 5, 43, 0, 0, 189, 185, 1, 0, 0, 0, 189, 188,
		1, 0, 0, 0, 190, 13, 1, 0, 0, 0, 191, 192, 5, 43, 0, 0, 192, 193, 5, 6,
		0, 0, 193, 194, 3, 12, 6, 0, 194, 15, 1, 0, 0, 0, 195, 196, 5, 7, 0, 0,
		196, 197, 5, 38, 0, 0, 197, 198, 5, 8, 0, 0, 198, 17, 1, 0, 0, 0, 199,
		200, 5, 9, 0, 0, 200, 19, 1, 0, 0, 0, 201, 202, 5, 10, 0, 0, 202, 203,
		3, 92, 46, 0, 203, 204, 5, 8, 0, 0, 204, 21, 1, 0, 0, 0, 205, 206, 5, 10,
		0, 0, 206, 207, 5, 38, 0, 0, 207, 208, 5, 8, 0, 0, 208, 23, 1, 0, 0, 0,
		209, 224, 3, 28, 14, 0, 210, 224, 3, 36, 18, 0, 211, 224, 3, 38, 19, 0,
		212, 224, 3, 40, 20, 0, 213, 224, 3, 42, 21, 0, 214, 224, 3, 44, 22, 0,
		215, 224, 3, 52, 26, 0, 216, 224, 3, 54, 27, 0, 217, 224, 3, 68, 34, 0,
		218, 224, 3, 72, 36, 0, 219, 224, 3, 34, 17, 0, 220, 224, 3, 30, 15, 0,
		221, 224, 3, 32, 16, 0, 222, 224, 3, 60, 30, 0, 223, 209, 1, 0, 0, 0, 223,
		210, 1, 0, 0, 0, 223, 211, 1, 0, 0, 0, 223, 212, 1, 0, 0, 0, 223, 213,
		1, 0, 0, 0, 223, 214, 1, 0, 0, 0, 223, 215, 1, 0, 0, 0, 223, 216, 1, 0,
		0, 0, 223, 217, 1, 0, 0, 0, 223, 218, 1, 0, 0, 0, 223, 219, 1, 0, 0, 0,
		223, 220, 1, 0, 0, 0, 223, 221, 1, 0, 0, 0, 223, 222, 1, 0, 0, 0, 224,
		25, 1, 0, 0, 0, 225, 226, 5, 43, 0, 0, 226, 227, 5, 6, 0, 0, 227, 228,
		3, 24, 12, 0, 228, 27, 1, 0, 0, 0, 229, 230, 5, 11, 0, 0, 230, 231, 3,
		24, 12, 0, 231, 29, 1, 0, 0, 0, 232, 236, 5, 12, 0, 0, 233, 235, 5, 41,
		0, 0, 234, 233, 1, 0, 0, 0, 235, 238, 1, 0, 0, 0, 236, 234, 1, 0, 0, 0,
		236, 237, 1, 0, 0, 0, 237, 239, 1, 0, 0, 0, 238, 236, 1, 0, 0, 0, 239,
		243, 5, 13, 0, 0, 240, 242, 5, 41, 0, 0, 241, 240, 1, 0, 0, 0, 242, 245,
		1, 0, 0, 0, 243, 241, 1, 0, 0, 0, 243, 244, 1, 0, 0, 0, 244, 255, 1, 0,
		0, 0, 245, 243, 1, 0, 0, 0, 246, 250, 3, 24, 12, 0, 247, 249, 5, 41, 0,
		0, 248, 247, 1, 0, 0, 0, 249, 252, 1, 0, 0, 0, 250, 248, 1, 0, 0, 0, 250,
		251, 1, 0, 0, 0, 251, 254, 1, 0, 0, 0, 252, 250, 1, 0, 0, 0, 253, 246,
		1, 0, 0, 0, 254, 257, 1, 0, 0, 0, 255, 253, 1, 0, 0, 0, 255, 256, 1, 0,
		0, 0, 256, 258, 1, 0, 0, 0, 257, 255, 1, 0, 0, 0, 258, 262, 3, 24, 12,
		0, 259, 261, 5, 41, 0, 0, 260, 259, 1, 0, 0, 0, 261, 264, 1, 0, 0, 0, 262,
		260, 1, 0, 0, 0, 262, 263, 1, 0, 0, 0, 263, 265, 1, 0, 0, 0, 264, 262,
		1, 0, 0, 0, 265, 266, 5, 8, 0, 0, 266, 31, 1, 0, 0, 0, 267, 271, 5, 14,
		0, 0, 268, 270, 5, 41, 0, 0, 269, 268, 1, 0, 0, 0, 270, 273, 1, 0, 0, 0,
		271, 269, 1, 0, 0, 0, 271, 272, 1, 0, 0, 0, 272, 274, 1, 0, 0, 0, 273,
		271, 1, 0, 0, 0, 274, 278, 5, 13, 0, 0, 275, 277, 5, 41, 0, 0, 276, 275,
		1, 0, 0, 0, 277, 280, 1, 0, 0, 0, 278, 276, 1, 0, 0, 0, 278, 279, 1, 0,
		0, 0, 279, 290, 1, 0, 0, 0, 280, 278, 1, 0, 0, 0, 281, 285, 3, 24, 12,
		0, 282, 284, 5, 41, 0, 0, 283, 282, 1, 0, 0, 0, 284, 287, 1, 0, 0, 0, 285,
		283, 1, 0, 0, 0, 285, 286, 1, 0, 0, 0, 286, 289, 1, 0, 0, 0, 287, 285,
		1, 0, 0, 0, 288, 281, 1, 0, 0, 0, 289, 292, 1, 0, 0, 0, 290, 288, 1, 0,
		0, 0, 290, 291, 1, 0, 0, 0, 291, 293, 1, 0, 0, 0, 292, 290, 1, 0, 0, 0,
		293, 297, 3, 24, 12, 0, 294, 296, 5, 41, 0, 0, 295, 294, 1, 0, 0, 0, 296,
		299, 1, 0, 0, 0, 297, 295, 1, 0, 0, 0, 297, 298, 1, 0, 0, 0, 298, 300,
		1, 0, 0, 0, 299, 297, 1, 0, 0, 0, 300, 301, 5, 8, 0, 0, 301, 33, 1, 0,
		0, 0, 302, 303, 3, 90, 45, 0, 303, 35, 1, 0, 0, 0, 304, 305, 5, 15, 0,
		0, 305, 37, 1, 0, 0, 0, 306, 307, 5, 16, 0, 0, 307, 311, 5, 13, 0, 0, 308,
		310, 5, 41, 0, 0, 309, 308, 1, 0, 0, 0, 310, 313, 1, 0, 0, 0, 311, 309,
		1, 0, 0, 0, 311, 312, 1, 0, 0, 0, 312, 324, 1, 0, 0, 0, 313, 311, 1, 0,
		0, 0, 314, 315, 3, 90, 45, 0, 315, 319, 5, 17, 0, 0, 316, 318, 5, 41, 0,
		0, 317, 316, 1, 0, 0, 0, 318, 321, 1, 0, 0, 0, 319, 317, 1, 0, 0, 0, 319,
		320, 1, 0, 0, 0, 320, 323, 1, 0, 0, 0, 321, 319, 1, 0, 0, 0, 322, 314,
		1, 0, 0, 0, 323, 326, 1, 0, 0, 0, 324, 322, 1, 0, 0, 0, 324, 325, 1, 0,
		0, 0, 325, 330, 1, 0, 0, 0, 326, 324, 1, 0, 0, 0, 327, 329, 5, 41, 0, 0,
		328, 327, 1, 0, 0, 0, 329, 332, 1, 0, 0, 0, 330, 328, 1, 0, 0, 0, 330,
		331, 1, 0, 0, 0, 331, 333, 1, 0, 0, 0, 332, 330, 1, 0, 0, 0, 333, 337,
		3, 90, 45, 0, 334, 336, 5, 41, 0, 0, 335, 334, 1, 0, 0, 0, 336, 339, 1,
		0, 0, 0, 337, 335, 1, 0, 0, 0, 337, 338, 1, 0, 0, 0, 338, 340, 1, 0, 0,
		0, 339, 337, 1, 0, 0, 0, 340, 341, 5, 8, 0, 0, 341, 39, 1, 0, 0, 0, 342,
		343, 5, 18, 0, 0, 343, 344, 5, 13, 0, 0, 344, 345, 5, 8, 0, 0, 345, 41,
		1, 0, 0, 0, 346, 347, 5, 19, 0, 0, 347, 348, 5, 13, 0, 0, 348, 349, 5,
		8, 0, 0, 349, 43, 1, 0, 0, 0, 350, 351, 5, 20, 0, 0, 351, 352, 5, 13, 0,
		0, 352, 353, 5, 8, 0, 0, 353, 45, 1, 0, 0, 0, 354, 355, 5, 21, 0, 0, 355,
		356, 3, 92, 46, 0, 356, 357, 5, 8, 0, 0, 357, 47, 1, 0, 0, 0, 358, 359,
		5, 22, 0, 0, 359, 360, 3, 92, 46, 0, 360, 361, 5, 8, 0, 0, 361, 49, 1,
		0, 0, 0, 362, 365, 3, 48, 24, 0, 363, 365, 3, 46, 23, 0, 364, 362, 1, 0,
		0, 0, 364, 363, 1, 0, 0, 0, 365, 51, 1, 0, 0, 0, 366, 367, 5, 23, 0, 0,
		367, 368, 5, 13, 0, 0, 368, 369, 3, 90, 45, 0, 369, 373, 5, 8, 0, 0, 370,
		372, 5, 41, 0, 0, 371, 370, 1, 0, 0, 0, 372, 375, 1, 0, 0, 0, 373, 371,
		1, 0, 0, 0, 373, 374, 1, 0, 0, 0, 374, 385, 1, 0, 0, 0, 375, 373, 1, 0,
		0, 0, 376, 380, 3, 50, 25, 0, 377, 379, 5, 41, 0, 0, 378, 377, 1, 0, 0,
		0, 379, 382, 1, 0, 0, 0, 380, 378, 1, 0, 0, 0, 380, 381, 1, 0, 0, 0, 381,
		384, 1, 0, 0, 0, 382, 380, 1, 0, 0, 0, 383, 376, 1, 0, 0, 0, 384, 387,
		1, 0, 0, 0, 385, 383, 1, 0, 0, 0, 385, 386, 1, 0, 0, 0, 386, 53, 1, 0,
		0, 0, 387, 385, 1, 0, 0, 0, 388, 389, 5, 24, 0, 0, 389, 390, 5, 37, 0,
		0, 390, 394, 5, 8, 0, 0, 391, 393, 5, 41, 0, 0, 392, 391, 1, 0, 0, 0, 393,
		396, 1, 0, 0, 0, 394, 392, 1, 0, 0, 0, 394, 395, 1, 0, 0, 0, 395, 406,
		1, 0, 0, 0, 396, 394, 1, 0, 0, 0, 397, 401, 3, 56, 28, 0, 398, 400, 5,
		41, 0, 0, 399, 398, 1, 0, 0, 0, 400, 403, 1, 0, 0, 0, 401, 399, 1, 0, 0,
		0, 401, 402, 1, 0, 0, 0, 402, 405, 1, 0, 0, 0, 403, 401, 1, 0, 0, 0, 404,
		397, 1, 0, 0, 0, 405, 408, 1, 0, 0, 0, 406, 404, 1, 0, 0, 0, 406, 407,
		1, 0, 0, 0, 407, 55, 1, 0, 0, 0, 408, 406, 1, 0, 0, 0, 409, 413, 3, 16,
		8, 0, 410, 413, 3, 18, 9, 0, 411, 413, 3, 58, 29, 0, 412, 409, 1, 0, 0,
		0, 412, 410, 1, 0, 0, 0, 412, 411, 1, 0, 0, 0, 413, 57, 1, 0, 0, 0, 414,
		415, 5, 25, 0, 0, 415, 416, 5, 38, 0, 0, 416, 417, 5, 8, 0, 0, 417, 59,
		1, 0, 0, 0, 418, 419, 5, 26, 0, 0, 419, 420, 5, 13, 0, 0, 420, 421, 5,
		40, 0, 0, 421, 425, 5, 8, 0, 0, 422, 424, 5, 41, 0, 0, 423, 422, 1, 0,
		0, 0, 424, 427, 1, 0, 0, 0, 425, 423, 1, 0, 0, 0, 425, 426, 1, 0, 0, 0,
		426, 437, 1, 0, 0, 0, 427, 425, 1, 0, 0, 0, 428, 432, 3, 62, 31, 0, 429,
		431, 5, 41, 0, 0, 430, 429, 1, 0, 0, 0, 431, 434, 1, 0, 0, 0, 432, 430,
		1, 0, 0, 0, 432, 433, 1, 0, 0, 0, 433, 436, 1, 0, 0, 0, 434, 432, 1, 0,
		0, 0, 435, 428, 1, 0, 0, 0, 436, 439, 1, 0, 0, 0, 437, 435, 1, 0, 0, 0,
		437, 438, 1, 0, 0, 0, 438, 61, 1, 0, 0, 0, 439, 437, 1, 0, 0, 0, 440, 443,
		3, 64, 32, 0, 441, 443, 3, 66, 33, 0, 442, 440, 1, 0, 0, 0, 442, 441, 1,
		0, 0, 0, 443, 63, 1, 0, 0, 0, 444, 445, 5, 27, 0, 0, 445, 446, 5, 13, 0,
		0, 446, 447, 3, 12, 6, 0, 447, 448, 5, 8, 0, 0, 448, 65, 1, 0, 0, 0, 449,
		450, 5, 28, 0, 0, 450, 451, 5, 13, 0, 0, 451, 452, 3, 12, 6, 0, 452, 453,
		5, 8, 0, 0, 453, 67, 1, 0, 0, 0, 454, 455, 5, 29, 0, 0, 455, 456, 5, 13,
		0, 0, 456, 460, 5, 8, 0, 0, 457, 459, 5, 41, 0, 0, 458, 457, 1, 0, 0, 0,
		459, 462, 1, 0, 0, 0, 460, 458, 1, 0, 0, 0, 460, 461, 1, 0, 0, 0, 461,
		472, 1, 0, 0, 0, 462, 460, 1, 0, 0, 0, 463, 467, 3, 70, 35, 0, 464, 466,
		5, 41, 0, 0, 465, 464, 1, 0, 0, 0, 466, 469, 1, 0, 0, 0, 467, 465, 1, 0,
		0, 0, 467, 468, 1, 0, 0, 0, 468, 471, 1, 0, 0, 0, 469, 467, 1, 0, 0, 0,
		470, 463, 1, 0, 0, 0, 471, 474, 1, 0, 0, 0, 472, 470, 1, 0, 0, 0, 472,
		473, 1, 0, 0, 0, 473, 69, 1, 0, 0, 0, 474, 472, 1, 0, 0, 0, 475, 479, 3,
		16, 8, 0, 476, 479, 3, 18, 9, 0, 477, 479, 3, 20, 10, 0, 478, 475, 1, 0,
		0, 0, 478, 476, 1, 0, 0, 0, 478, 477, 1, 0, 0, 0, 479, 71, 1, 0, 0, 0,
		480, 481, 5, 30, 0, 0, 481, 482, 5, 13, 0, 0, 482, 483, 3, 12, 6, 0, 483,
		487, 5, 8, 0, 0, 484, 486, 5, 41, 0, 0, 485, 484, 1, 0, 0, 0, 486, 489,
		1, 0, 0, 0, 487, 485, 1, 0, 0, 0, 487, 488, 1, 0, 0, 0, 488, 499, 1, 0,
		0, 0, 489, 487, 1, 0, 0, 0, 490, 494, 3, 74, 37, 0, 491, 493, 5, 41, 0,
		0, 492, 491, 1, 0, 0, 0, 493, 496, 1, 0, 0, 0, 494, 492, 1, 0, 0, 0, 494,
		495, 1, 0, 0, 0, 495, 498, 1, 0, 0, 0, 496, 494, 1, 0, 0, 0, 497, 490,
		1, 0, 0, 0, 498, 501, 1, 0, 0, 0, 499, 497, 1, 0, 0, 0, 499, 500, 1, 0,
		0, 0, 500, 73, 1, 0, 0, 0, 501, 499, 1, 0, 0, 0, 502, 505, 3, 22, 11, 0,
		503, 505, 3, 18, 9, 0, 504, 502, 1, 0, 0, 0, 504, 503, 1, 0, 0, 0, 505,
		75, 1, 0, 0, 0, 506, 507, 5, 43, 0, 0, 507, 508, 5, 6, 0, 0, 508, 509,
		3, 78, 39, 0, 509, 77, 1, 0, 0, 0, 510, 516, 3, 82, 41, 0, 511, 516, 3,
		84, 42, 0, 512, 516, 3, 80, 40, 0, 513, 516, 3, 88, 44, 0, 514, 516, 3,
		86, 43, 0, 515, 510, 1, 0, 0, 0, 515, 511, 1, 0, 0, 0, 515, 512, 1, 0,
		0, 0, 515, 513, 1, 0, 0, 0, 515, 514, 1, 0, 0, 0, 516, 79, 1, 0, 0, 0,
		517, 518, 3, 90, 45, 0, 518, 81, 1, 0, 0, 0, 519, 520, 5, 31, 0, 0, 520,
		521, 5, 13, 0, 0, 521, 522, 3, 90, 45, 0, 522, 523, 5, 8, 0, 0, 523, 83,
		1, 0, 0, 0, 524, 528, 5, 32, 0, 0, 525, 527, 5, 41, 0, 0, 526, 525, 1,
		0, 0, 0, 527, 530, 1, 0, 0, 0, 528, 526, 1, 0, 0, 0, 528, 529, 1, 0, 0,
		0, 529, 540, 1, 0, 0, 0, 530, 528, 1, 0, 0, 0, 531, 535, 3, 78, 39, 0,
		532, 534, 5, 41, 0, 0, 533, 532, 1, 0, 0, 0, 534, 537, 1, 0, 0, 0, 535,
		533, 1, 0, 0, 0, 535, 536, 1, 0, 0, 0, 536, 539, 1, 0, 0, 0, 537, 535,
		1, 0, 0, 0, 538, 531, 1, 0, 0, 0, 539, 542, 1, 0, 0, 0, 540, 538, 1, 0,
		0, 0, 540, 541, 1, 0, 0, 0, 541, 546, 1, 0, 0, 0, 542, 540, 1, 0, 0, 0,
		543, 545, 5, 41, 0, 0, 544, 543, 1, 0, 0, 0, 545, 548, 1, 0, 0, 0, 546,
		544, 1, 0, 0, 0, 546, 547, 1, 0, 0, 0, 547, 549, 1, 0, 0, 0, 548, 546,
		1, 0, 0, 0, 549, 550, 5, 33, 0, 0, 550, 85, 1, 0, 0, 0, 551, 552, 5, 34,
		0, 0, 552, 553, 5, 13, 0, 0, 553, 554, 5, 8, 0, 0, 554, 87, 1, 0, 0, 0,
		555, 559, 5, 35, 0, 0, 556, 558, 5, 41, 0, 0, 557, 556, 1, 0, 0, 0, 558,
		561, 1, 0, 0, 0, 559, 557, 1, 0, 0, 0, 559, 560, 1, 0, 0, 0, 560, 562,
		1, 0, 0, 0, 561, 559, 1, 0, 0, 0, 562, 566, 5, 13, 0, 0, 563, 565, 5, 41,
		0, 0, 564, 563, 1, 0, 0, 0, 565, 568, 1, 0, 0, 0, 566, 564, 1, 0, 0, 0,
		566, 567, 1, 0, 0, 0, 567, 569, 1, 0, 0, 0, 568, 566, 1, 0, 0, 0, 569,
		573, 3, 24, 12, 0, 570, 572, 5, 41, 0, 0, 571, 570, 1, 0, 0, 0, 572, 575,
		1, 0, 0, 0, 573, 571, 1, 0, 0, 0, 573, 574, 1, 0, 0, 0, 574, 576, 1, 0,
		0, 0, 575, 573, 1, 0, 0, 0, 576, 580, 5, 8, 0, 0, 577, 579, 5, 41, 0, 0,
		578, 577, 1, 0, 0, 0, 579, 582, 1, 0, 0, 0, 580, 578, 1, 0, 0, 0, 580,
		581, 1, 0, 0, 0, 581, 583, 1, 0, 0, 0, 582, 580, 1, 0, 0, 0, 583, 584,
		3, 78, 39, 0, 584, 89, 1, 0, 0, 0, 585, 586, 5, 43, 0, 0, 586, 588, 5,
		36, 0, 0, 587, 585, 1, 0, 0, 0, 587, 588, 1, 0, 0, 0, 588, 589, 1, 0, 0,
		0, 589, 590, 5, 43, 0, 0, 590, 91, 1, 0, 0, 0, 591, 592, 7, 0, 0, 0, 592,
		93, 1, 0, 0, 0, 61, 98, 103, 113, 120, 127, 132, 138, 149, 156, 163, 168,
		174, 182, 185, 189, 223, 236, 243, 250, 255, 262, 271, 278, 285, 290, 297,
		311, 319, 324, 330, 337, 364, 373, 380, 385, 394, 401, 406, 412, 425, 432,
		437, 442, 460, 467, 472, 478, 487, 494, 499, 504, 515, 528, 535, 540, 546,
		559, 566, 573, 580, 587,
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
	{
		p.SetState(193)
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
		p.SetState(195)
		p.Match(MinecraftMetascriptParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(196)
		p.Match(MinecraftMetascriptParserInt)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(197)
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
		p.SetState(199)
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
		p.SetState(201)
		p.Match(MinecraftMetascriptParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(202)
		p.Number()
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
		p.SetState(205)
		p.Match(MinecraftMetascriptParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(206)
		p.Match(MinecraftMetascriptParserInt)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(207)
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
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinecraftMetascriptParserT__10:
		{
			p.SetState(209)
			p.SurfaceCondition_Not()
		}

	case MinecraftMetascriptParserT__14:
		{
			p.SetState(210)
			p.SurfaceCondition_AboveSurface()
		}

	case MinecraftMetascriptParserT__15:
		{
			p.SetState(211)
			p.SurfaceCondition_Biome()
		}

	case MinecraftMetascriptParserT__17:
		{
			p.SetState(212)
			p.SurfaceCondition_Hole()
		}

	case MinecraftMetascriptParserT__18:
		{
			p.SetState(213)
			p.SurfaceCondition_Steep()
		}

	case MinecraftMetascriptParserT__19:
		{
			p.SetState(214)
			p.SurfaceCondition_Freezing()
		}

	case MinecraftMetascriptParserT__22:
		{
			p.SetState(215)
			p.SurfaceCondition_Noise()
		}

	case MinecraftMetascriptParserT__23:
		{
			p.SetState(216)
			p.SurfaceCondition_StoneDepth()
		}

	case MinecraftMetascriptParserT__28:
		{
			p.SetState(217)
			p.SurfaceCondition_AboveWater()
		}

	case MinecraftMetascriptParserT__29:
		{
			p.SetState(218)
			p.SurfaceCondition_YAbove()
		}

	case MinecraftMetascriptParserIdentifier:
		{
			p.SetState(219)
			p.SurfaceCondition_Reference()
		}

	case MinecraftMetascriptParserT__11:
		{
			p.SetState(220)
			p.SurfaceCondition_And()
		}

	case MinecraftMetascriptParserT__13:
		{
			p.SetState(221)
			p.SurfaceCondition_Or()
		}

	case MinecraftMetascriptParserT__25:
		{
			p.SetState(222)
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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(225)
		p.Match(MinecraftMetascriptParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(226)
		p.Match(MinecraftMetascriptParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(227)
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
		p.SetState(229)
		p.Match(MinecraftMetascriptParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(230)
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
		p.SetState(232)
		p.Match(MinecraftMetascriptParserT__11)
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
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(240)
			p.Match(MinecraftMetascriptParserNL)
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
	}
	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(246)
				p.SurfaceCondition()
			}
			p.SetState(250)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == MinecraftMetascriptParserNL {
				{
					p.SetState(247)
					p.Match(MinecraftMetascriptParserNL)
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
			}

		}
		p.SetState(257)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
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
	{
		p.SetState(265)
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
		p.SetState(267)
		p.Match(MinecraftMetascriptParserT__13)
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

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(268)
			p.Match(MinecraftMetascriptParserNL)
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
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(275)
			p.Match(MinecraftMetascriptParserNL)
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
	}
	p.SetState(290)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(281)
				p.SurfaceCondition()
			}
			p.SetState(285)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == MinecraftMetascriptParserNL {
				{
					p.SetState(282)
					p.Match(MinecraftMetascriptParserNL)
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
			}

		}
		p.SetState(292)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
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
	{
		p.SetState(300)
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
		p.SetState(302)
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
		p.SetState(304)
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
		p.SetState(306)
		p.Match(MinecraftMetascriptParserT__15)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(307)
		p.Match(MinecraftMetascriptParserT__12)
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
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(308)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(313)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(324)
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
				p.SetState(314)
				p.ResourceReference()
			}
			{
				p.SetState(315)
				p.Match(MinecraftMetascriptParserT__16)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(319)
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
						p.SetState(316)
						p.Match(MinecraftMetascriptParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				p.SetState(321)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext())
				if p.HasError() {
					goto errorExit
				}
			}

		}
		p.SetState(326)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(330)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(327)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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
		p.ResourceReference()
	}
	p.SetState(337)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(334)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(339)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(340)
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
		p.SetState(342)
		p.Match(MinecraftMetascriptParserT__17)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(343)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(344)
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
		p.SetState(346)
		p.Match(MinecraftMetascriptParserT__18)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(347)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(348)
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
		p.SetState(350)
		p.Match(MinecraftMetascriptParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(351)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
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
		p.SetState(354)
		p.Match(MinecraftMetascriptParserT__20)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(355)
		p.Number()
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
		p.SetState(358)
		p.Match(MinecraftMetascriptParserT__21)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(359)
		p.Number()
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
	p.SetState(364)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinecraftMetascriptParserT__21:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(362)
			p.SurfaceCondition_NoiseBuilder_Max()
		}

	case MinecraftMetascriptParserT__20:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(363)
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
		p.SetState(366)
		p.Match(MinecraftMetascriptParserT__22)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(367)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(368)
		p.ResourceReference()
	}
	{
		p.SetState(369)
		p.Match(MinecraftMetascriptParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(373)
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
				p.SetState(370)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(375)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(385)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserT__20 || _la == MinecraftMetascriptParserT__21 {
		{
			p.SetState(376)
			p.SurfaceCondition_NoiseBuilder()
		}
		p.SetState(380)
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
					p.SetState(377)
					p.Match(MinecraftMetascriptParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(382)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

		p.SetState(387)
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
		p.SetState(388)
		p.Match(MinecraftMetascriptParserT__23)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(389)
		p.Match(MinecraftMetascriptParserStoneDepthMode)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(390)
		p.Match(MinecraftMetascriptParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
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
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(391)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(396)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(406)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&33555072) != 0 {
		{
			p.SetState(397)
			p.SurfaceCondition_StoneDepthBuilder()
		}
		p.SetState(401)
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
					p.SetState(398)
					p.Match(MinecraftMetascriptParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(403)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext())
			if p.HasError() {
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
	p.SetState(412)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinecraftMetascriptParserT__6:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(409)
			p.SharedBuilder_Offset()
		}

	case MinecraftMetascriptParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(410)
			p.SharedBuilder_Add()
		}

	case MinecraftMetascriptParserT__24:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(411)
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
		p.SetState(414)
		p.Match(MinecraftMetascriptParserT__24)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(415)
		p.Match(MinecraftMetascriptParserInt)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(416)
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
		p.SetState(418)
		p.Match(MinecraftMetascriptParserT__25)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(419)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(420)
		p.Match(MinecraftMetascriptParserString_)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(421)
		p.Match(MinecraftMetascriptParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(425)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(422)
				p.Match(MinecraftMetascriptParserNL)
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
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(437)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserT__26 || _la == MinecraftMetascriptParserT__27 {
		{
			p.SetState(428)
			p.SurfaceCondition_VerticalGradientBuilder()
		}
		p.SetState(432)
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
					p.SetState(429)
					p.Match(MinecraftMetascriptParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(434)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

		p.SetState(439)
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
	p.SetState(442)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinecraftMetascriptParserT__26:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(440)
			p.SurfaceCondition_VerticalGradientBuilder_Top()
		}

	case MinecraftMetascriptParserT__27:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(441)
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
		p.SetState(444)
		p.Match(MinecraftMetascriptParserT__26)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(445)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(446)
		p.VerticalAnchor()
	}
	{
		p.SetState(447)
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
		p.SetState(449)
		p.Match(MinecraftMetascriptParserT__27)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(450)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(451)
		p.VerticalAnchor()
	}
	{
		p.SetState(452)
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
		p.SetState(454)
		p.Match(MinecraftMetascriptParserT__28)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(455)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(456)
		p.Match(MinecraftMetascriptParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(460)
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
				p.SetState(457)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(462)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext())
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

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1664) != 0 {
		{
			p.SetState(463)
			p.SurfaceCondition_AboveWaterBuilder()
		}
		p.SetState(467)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 44, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(464)
					p.Match(MinecraftMetascriptParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(469)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 44, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

		p.SetState(474)
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
	p.SetState(478)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinecraftMetascriptParserT__6:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(475)
			p.SharedBuilder_Offset()
		}

	case MinecraftMetascriptParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(476)
			p.SharedBuilder_Add()
		}

	case MinecraftMetascriptParserT__9:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(477)
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
		p.SetState(480)
		p.Match(MinecraftMetascriptParserT__29)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(481)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(482)
		p.VerticalAnchor()
	}
	{
		p.SetState(483)
		p.Match(MinecraftMetascriptParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(487)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 47, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(484)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(489)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 47, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(499)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserT__8 || _la == MinecraftMetascriptParserT__9 {
		{
			p.SetState(490)
			p.SurfaceCondition_YAboveBuilder()
		}
		p.SetState(494)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 48, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(491)
					p.Match(MinecraftMetascriptParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(496)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 48, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

		p.SetState(501)
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
	p.SetState(504)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinecraftMetascriptParserT__9:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(502)
			p.SharedBuilder_MulInt()
		}

	case MinecraftMetascriptParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(503)
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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(506)
		p.Match(MinecraftMetascriptParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(507)
		p.Match(MinecraftMetascriptParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(508)
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
	p.SetState(515)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinecraftMetascriptParserT__30:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(510)
			p.SurfaceRule_Block()
		}

	case MinecraftMetascriptParserT__31:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(511)
			p.SurfaceRule_Sequence()
		}

	case MinecraftMetascriptParserIdentifier:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(512)
			p.SurfaceRule_Reference()
		}

	case MinecraftMetascriptParserT__34:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(513)
			p.SurfaceRule_If()
		}

	case MinecraftMetascriptParserT__33:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(514)
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
		p.SetState(517)
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
		p.SetState(519)
		p.Match(MinecraftMetascriptParserT__30)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(520)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(521)
		p.ResourceReference()
	}
	{
		p.SetState(522)
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
		p.SetState(524)
		p.Match(MinecraftMetascriptParserT__31)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(528)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 52, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(525)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(530)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 52, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(540)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8854075080704) != 0 {
		{
			p.SetState(531)
			p.SurfaceRule()
		}
		p.SetState(535)
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
					p.SetState(532)
					p.Match(MinecraftMetascriptParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(537)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 53, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

		p.SetState(542)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(546)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(543)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(548)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(549)
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
		p.SetState(551)
		p.Match(MinecraftMetascriptParserT__33)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(552)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(553)
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
		p.SetState(555)
		p.Match(MinecraftMetascriptParserT__34)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(559)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(556)
			p.Match(MinecraftMetascriptParserNL)
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
	}
	{
		p.SetState(562)
		p.Match(MinecraftMetascriptParserT__12)
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

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(563)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(568)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(569)
		p.SurfaceCondition()
	}
	p.SetState(573)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(570)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(575)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(576)
		p.Match(MinecraftMetascriptParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(580)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(577)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(582)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(583)
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
	p.SetState(587)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 60, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(585)
			p.Match(MinecraftMetascriptParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(586)
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
		p.SetState(589)
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
		p.SetState(591)
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
