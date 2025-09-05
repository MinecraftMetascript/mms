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
		"'NoiseThreshold'", "'StoneDepth('", "'.SecondaryDepthRange('", "'VerticalGradient'",
		"'.Top'", "'.Bottom'", "'AboveWater'", "'YAbove'", "'Block'", "'['",
		"']'", "'Bandlands'", "'If'", "':'", "'Noise'", "'Noise('", "'.Amplitudes('",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "StoneDepthMode", "Int", "Float", "String",
		"NL", "WS", "Identifier", "BlockComment", "LineComment",
	}
	staticData.RuleNames = []string{
		"script", "namespaceDeclaration", "namespace", "contentBlocks", "surfaceBlock",
		"surfaceStatement", "verticalAnchor", "verticalAnchorDeclaration", "sharedBuilder_Offset",
		"sharedBuilder_Add", "sharedBuilder_Mul", "sharedBuilder_MulInt", "surfaceCondition",
		"surfaceConditionDeclaration", "surfaceCondition_Not", "surfaceCondition_And",
		"surfaceCondition_Or", "surfaceCondition_Reference", "surfaceCondition_AboveSurface",
		"surfaceCondition_Biome", "surfaceCondition_Hole", "surfaceCondition_Steep",
		"surfaceCondition_Freezing", "surfaceCondition_NoiseThresholdBuilder_Min",
		"surfaceCondition_NoiseThresholdBuilder_Max", "surfaceCondition_NoiseThresholdBuilder",
		"surfaceCondition_NoiseThreshold", "surfaceCondition_StoneDepth", "surfaceCondition_StoneDepthBuilder",
		"surfaceCondition_StoneDepthBuilder_SecondaryDepthRange", "surfaceCondition_VerticalGradient",
		"surfaceCondition_VerticalGradientBuilder", "surfaceCondition_VerticalGradientBuilder_Top",
		"surfaceCondition_VerticalGradientBuilder_Bottom", "surfaceCondition_AboveWater",
		"surfaceCondition_AboveWaterBuilder", "surfaceCondition_YAbove", "surfaceCondition_YAboveBuilder",
		"surfaceRuleDeclaration", "surfaceRule", "surfaceRule_Reference", "surfaceRule_Block",
		"surfaceRule_Sequence", "surfaceRule_Bandlands", "surfaceRule_If", "resourceReference",
		"number", "noiseBlock", "noiseDeclaration", "noise", "noise_Builder",
		"noise_Builder_Amplitudes",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 48, 715, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 1, 0, 5,
		0, 106, 8, 0, 10, 0, 12, 0, 109, 9, 0, 1, 0, 1, 0, 5, 0, 113, 8, 0, 10,
		0, 12, 0, 116, 9, 0, 5, 0, 118, 8, 0, 10, 0, 12, 0, 121, 9, 0, 1, 1, 1,
		1, 1, 1, 1, 2, 1, 2, 5, 2, 128, 8, 2, 10, 2, 12, 2, 131, 9, 2, 1, 2, 1,
		2, 5, 2, 135, 8, 2, 10, 2, 12, 2, 138, 9, 2, 1, 2, 1, 2, 5, 2, 142, 8,
		2, 10, 2, 12, 2, 145, 9, 2, 5, 2, 147, 8, 2, 10, 2, 12, 2, 150, 9, 2, 1,
		2, 5, 2, 153, 8, 2, 10, 2, 12, 2, 156, 9, 2, 1, 2, 1, 2, 1, 3, 1, 3, 3,
		3, 162, 8, 3, 1, 4, 1, 4, 5, 4, 166, 8, 4, 10, 4, 12, 4, 169, 9, 4, 1,
		4, 1, 4, 5, 4, 173, 8, 4, 10, 4, 12, 4, 176, 9, 4, 1, 4, 1, 4, 5, 4, 180,
		8, 4, 10, 4, 12, 4, 183, 9, 4, 5, 4, 185, 8, 4, 10, 4, 12, 4, 188, 9, 4,
		1, 4, 5, 4, 191, 8, 4, 10, 4, 12, 4, 194, 9, 4, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 5, 3, 5, 201, 8, 5, 1, 6, 3, 6, 204, 8, 6, 1, 6, 1, 6, 3, 6, 208, 8,
		6, 1, 7, 1, 7, 1, 7, 5, 7, 213, 8, 7, 10, 7, 12, 7, 216, 9, 7, 1, 7, 1,
		7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 248, 8, 12, 1, 13,
		1, 13, 1, 13, 5, 13, 253, 8, 13, 10, 13, 12, 13, 256, 9, 13, 1, 13, 1,
		13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 5, 15, 265, 8, 15, 10, 15, 12, 15,
		268, 9, 15, 1, 15, 1, 15, 5, 15, 272, 8, 15, 10, 15, 12, 15, 275, 9, 15,
		1, 15, 1, 15, 5, 15, 279, 8, 15, 10, 15, 12, 15, 282, 9, 15, 5, 15, 284,
		8, 15, 10, 15, 12, 15, 287, 9, 15, 1, 15, 1, 15, 5, 15, 291, 8, 15, 10,
		15, 12, 15, 294, 9, 15, 1, 15, 1, 15, 1, 16, 1, 16, 5, 16, 300, 8, 16,
		10, 16, 12, 16, 303, 9, 16, 1, 16, 1, 16, 5, 16, 307, 8, 16, 10, 16, 12,
		16, 310, 9, 16, 1, 16, 1, 16, 5, 16, 314, 8, 16, 10, 16, 12, 16, 317, 9,
		16, 5, 16, 319, 8, 16, 10, 16, 12, 16, 322, 9, 16, 1, 16, 1, 16, 5, 16,
		326, 8, 16, 10, 16, 12, 16, 329, 9, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1,
		18, 1, 18, 1, 19, 1, 19, 1, 19, 5, 19, 340, 8, 19, 10, 19, 12, 19, 343,
		9, 19, 1, 19, 1, 19, 1, 19, 5, 19, 348, 8, 19, 10, 19, 12, 19, 351, 9,
		19, 5, 19, 353, 8, 19, 10, 19, 12, 19, 356, 9, 19, 1, 19, 5, 19, 359, 8,
		19, 10, 19, 12, 19, 362, 9, 19, 1, 19, 1, 19, 5, 19, 366, 8, 19, 10, 19,
		12, 19, 369, 9, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 3, 25, 395, 8, 25, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 5, 26, 402, 8, 26, 10, 26, 12, 26, 405, 9, 26,
		1, 26, 1, 26, 5, 26, 409, 8, 26, 10, 26, 12, 26, 412, 9, 26, 5, 26, 414,
		8, 26, 10, 26, 12, 26, 417, 9, 26, 1, 27, 1, 27, 1, 27, 1, 27, 5, 27, 423,
		8, 27, 10, 27, 12, 27, 426, 9, 27, 1, 27, 1, 27, 5, 27, 430, 8, 27, 10,
		27, 12, 27, 433, 9, 27, 5, 27, 435, 8, 27, 10, 27, 12, 27, 438, 9, 27,
		1, 28, 1, 28, 1, 28, 3, 28, 443, 8, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1,
		30, 1, 30, 1, 30, 1, 30, 1, 30, 5, 30, 454, 8, 30, 10, 30, 12, 30, 457,
		9, 30, 1, 30, 1, 30, 5, 30, 461, 8, 30, 10, 30, 12, 30, 464, 9, 30, 5,
		30, 466, 8, 30, 10, 30, 12, 30, 469, 9, 30, 1, 31, 1, 31, 3, 31, 473, 8,
		31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33,
		1, 34, 1, 34, 1, 34, 1, 34, 5, 34, 489, 8, 34, 10, 34, 12, 34, 492, 9,
		34, 1, 34, 1, 34, 5, 34, 496, 8, 34, 10, 34, 12, 34, 499, 9, 34, 5, 34,
		501, 8, 34, 10, 34, 12, 34, 504, 9, 34, 1, 35, 1, 35, 1, 35, 3, 35, 509,
		8, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 5, 36, 516, 8, 36, 10, 36, 12,
		36, 519, 9, 36, 1, 36, 1, 36, 5, 36, 523, 8, 36, 10, 36, 12, 36, 526, 9,
		36, 5, 36, 528, 8, 36, 10, 36, 12, 36, 531, 9, 36, 1, 37, 1, 37, 3, 37,
		535, 8, 37, 1, 38, 1, 38, 1, 38, 5, 38, 540, 8, 38, 10, 38, 12, 38, 543,
		9, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 3, 39, 552, 8,
		39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 42, 1, 42, 5, 42,
		563, 8, 42, 10, 42, 12, 42, 566, 9, 42, 1, 42, 1, 42, 5, 42, 570, 8, 42,
		10, 42, 12, 42, 573, 9, 42, 5, 42, 575, 8, 42, 10, 42, 12, 42, 578, 9,
		42, 1, 42, 5, 42, 581, 8, 42, 10, 42, 12, 42, 584, 9, 42, 1, 42, 1, 42,
		1, 43, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44, 5, 44, 594, 8, 44, 10, 44, 12,
		44, 597, 9, 44, 1, 44, 1, 44, 5, 44, 601, 8, 44, 10, 44, 12, 44, 604, 9,
		44, 1, 44, 1, 44, 5, 44, 608, 8, 44, 10, 44, 12, 44, 611, 9, 44, 1, 44,
		1, 44, 5, 44, 615, 8, 44, 10, 44, 12, 44, 618, 9, 44, 1, 44, 1, 44, 1,
		45, 1, 45, 3, 45, 624, 8, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 47, 1, 47,
		5, 47, 632, 8, 47, 10, 47, 12, 47, 635, 9, 47, 1, 47, 1, 47, 5, 47, 639,
		8, 47, 10, 47, 12, 47, 642, 9, 47, 1, 47, 1, 47, 5, 47, 646, 8, 47, 10,
		47, 12, 47, 649, 9, 47, 5, 47, 651, 8, 47, 10, 47, 12, 47, 654, 9, 47,
		1, 47, 5, 47, 657, 8, 47, 10, 47, 12, 47, 660, 9, 47, 1, 47, 1, 47, 1,
		48, 1, 48, 1, 48, 1, 48, 1, 49, 1, 49, 5, 49, 670, 8, 49, 10, 49, 12, 49,
		673, 9, 49, 1, 49, 1, 49, 5, 49, 677, 8, 49, 10, 49, 12, 49, 680, 9, 49,
		1, 49, 1, 49, 5, 49, 684, 8, 49, 10, 49, 12, 49, 687, 9, 49, 1, 49, 1,
		49, 5, 49, 691, 8, 49, 10, 49, 12, 49, 694, 9, 49, 5, 49, 696, 8, 49, 10,
		49, 12, 49, 699, 9, 49, 1, 50, 1, 50, 1, 51, 1, 51, 1, 51, 1, 51, 5, 51,
		707, 8, 51, 10, 51, 12, 51, 710, 9, 51, 1, 51, 1, 51, 1, 51, 1, 51, 0,
		0, 52, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
		36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70,
		72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 0, 1,
		1, 0, 41, 42, 757, 0, 107, 1, 0, 0, 0, 2, 122, 1, 0, 0, 0, 4, 125, 1, 0,
		0, 0, 6, 161, 1, 0, 0, 0, 8, 163, 1, 0, 0, 0, 10, 200, 1, 0, 0, 0, 12,
		207, 1, 0, 0, 0, 14, 209, 1, 0, 0, 0, 16, 219, 1, 0, 0, 0, 18, 223, 1,
		0, 0, 0, 20, 225, 1, 0, 0, 0, 22, 229, 1, 0, 0, 0, 24, 247, 1, 0, 0, 0,
		26, 249, 1, 0, 0, 0, 28, 259, 1, 0, 0, 0, 30, 262, 1, 0, 0, 0, 32, 297,
		1, 0, 0, 0, 34, 332, 1, 0, 0, 0, 36, 334, 1, 0, 0, 0, 38, 336, 1, 0, 0,
		0, 40, 372, 1, 0, 0, 0, 42, 376, 1, 0, 0, 0, 44, 380, 1, 0, 0, 0, 46, 384,
		1, 0, 0, 0, 48, 388, 1, 0, 0, 0, 50, 394, 1, 0, 0, 0, 52, 396, 1, 0, 0,
		0, 54, 418, 1, 0, 0, 0, 56, 442, 1, 0, 0, 0, 58, 444, 1, 0, 0, 0, 60, 448,
		1, 0, 0, 0, 62, 472, 1, 0, 0, 0, 64, 474, 1, 0, 0, 0, 66, 479, 1, 0, 0,
		0, 68, 484, 1, 0, 0, 0, 70, 508, 1, 0, 0, 0, 72, 510, 1, 0, 0, 0, 74, 534,
		1, 0, 0, 0, 76, 536, 1, 0, 0, 0, 78, 551, 1, 0, 0, 0, 80, 553, 1, 0, 0,
		0, 82, 555, 1, 0, 0, 0, 84, 560, 1, 0, 0, 0, 86, 587, 1, 0, 0, 0, 88, 591,
		1, 0, 0, 0, 90, 623, 1, 0, 0, 0, 92, 627, 1, 0, 0, 0, 94, 629, 1, 0, 0,
		0, 96, 663, 1, 0, 0, 0, 98, 667, 1, 0, 0, 0, 100, 700, 1, 0, 0, 0, 102,
		702, 1, 0, 0, 0, 104, 106, 5, 44, 0, 0, 105, 104, 1, 0, 0, 0, 106, 109,
		1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 119, 1, 0,
		0, 0, 109, 107, 1, 0, 0, 0, 110, 114, 3, 4, 2, 0, 111, 113, 5, 44, 0, 0,
		112, 111, 1, 0, 0, 0, 113, 116, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 114,
		115, 1, 0, 0, 0, 115, 118, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 117, 110,
		1, 0, 0, 0, 118, 121, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 119, 120, 1, 0,
		0, 0, 120, 1, 1, 0, 0, 0, 121, 119, 1, 0, 0, 0, 122, 123, 5, 1, 0, 0, 123,
		124, 5, 46, 0, 0, 124, 3, 1, 0, 0, 0, 125, 129, 3, 2, 1, 0, 126, 128, 5,
		44, 0, 0, 127, 126, 1, 0, 0, 0, 128, 131, 1, 0, 0, 0, 129, 127, 1, 0, 0,
		0, 129, 130, 1, 0, 0, 0, 130, 132, 1, 0, 0, 0, 131, 129, 1, 0, 0, 0, 132,
		136, 5, 2, 0, 0, 133, 135, 5, 44, 0, 0, 134, 133, 1, 0, 0, 0, 135, 138,
		1, 0, 0, 0, 136, 134, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137, 148, 1, 0,
		0, 0, 138, 136, 1, 0, 0, 0, 139, 143, 3, 6, 3, 0, 140, 142, 5, 44, 0, 0,
		141, 140, 1, 0, 0, 0, 142, 145, 1, 0, 0, 0, 143, 141, 1, 0, 0, 0, 143,
		144, 1, 0, 0, 0, 144, 147, 1, 0, 0, 0, 145, 143, 1, 0, 0, 0, 146, 139,
		1, 0, 0, 0, 147, 150, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 148, 149, 1, 0,
		0, 0, 149, 154, 1, 0, 0, 0, 150, 148, 1, 0, 0, 0, 151, 153, 5, 44, 0, 0,
		152, 151, 1, 0, 0, 0, 153, 156, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 154,
		155, 1, 0, 0, 0, 155, 157, 1, 0, 0, 0, 156, 154, 1, 0, 0, 0, 157, 158,
		5, 3, 0, 0, 158, 5, 1, 0, 0, 0, 159, 162, 3, 8, 4, 0, 160, 162, 3, 94,
		47, 0, 161, 159, 1, 0, 0, 0, 161, 160, 1, 0, 0, 0, 162, 7, 1, 0, 0, 0,
		163, 167, 5, 4, 0, 0, 164, 166, 5, 44, 0, 0, 165, 164, 1, 0, 0, 0, 166,
		169, 1, 0, 0, 0, 167, 165, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 170,
		1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 170, 174, 5, 2, 0, 0, 171, 173, 5, 44,
		0, 0, 172, 171, 1, 0, 0, 0, 173, 176, 1, 0, 0, 0, 174, 172, 1, 0, 0, 0,
		174, 175, 1, 0, 0, 0, 175, 186, 1, 0, 0, 0, 176, 174, 1, 0, 0, 0, 177,
		181, 3, 10, 5, 0, 178, 180, 5, 44, 0, 0, 179, 178, 1, 0, 0, 0, 180, 183,
		1, 0, 0, 0, 181, 179, 1, 0, 0, 0, 181, 182, 1, 0, 0, 0, 182, 185, 1, 0,
		0, 0, 183, 181, 1, 0, 0, 0, 184, 177, 1, 0, 0, 0, 185, 188, 1, 0, 0, 0,
		186, 184, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 192, 1, 0, 0, 0, 188,
		186, 1, 0, 0, 0, 189, 191, 5, 44, 0, 0, 190, 189, 1, 0, 0, 0, 191, 194,
		1, 0, 0, 0, 192, 190, 1, 0, 0, 0, 192, 193, 1, 0, 0, 0, 193, 195, 1, 0,
		0, 0, 194, 192, 1, 0, 0, 0, 195, 196, 5, 3, 0, 0, 196, 9, 1, 0, 0, 0, 197,
		201, 3, 14, 7, 0, 198, 201, 3, 26, 13, 0, 199, 201, 3, 76, 38, 0, 200,
		197, 1, 0, 0, 0, 200, 198, 1, 0, 0, 0, 200, 199, 1, 0, 0, 0, 201, 11, 1,
		0, 0, 0, 202, 204, 5, 5, 0, 0, 203, 202, 1, 0, 0, 0, 203, 204, 1, 0, 0,
		0, 204, 205, 1, 0, 0, 0, 205, 208, 5, 41, 0, 0, 206, 208, 5, 46, 0, 0,
		207, 203, 1, 0, 0, 0, 207, 206, 1, 0, 0, 0, 208, 13, 1, 0, 0, 0, 209, 210,
		5, 46, 0, 0, 210, 214, 5, 6, 0, 0, 211, 213, 5, 44, 0, 0, 212, 211, 1,
		0, 0, 0, 213, 216, 1, 0, 0, 0, 214, 212, 1, 0, 0, 0, 214, 215, 1, 0, 0,
		0, 215, 217, 1, 0, 0, 0, 216, 214, 1, 0, 0, 0, 217, 218, 3, 12, 6, 0, 218,
		15, 1, 0, 0, 0, 219, 220, 5, 7, 0, 0, 220, 221, 5, 41, 0, 0, 221, 222,
		5, 8, 0, 0, 222, 17, 1, 0, 0, 0, 223, 224, 5, 9, 0, 0, 224, 19, 1, 0, 0,
		0, 225, 226, 5, 10, 0, 0, 226, 227, 3, 92, 46, 0, 227, 228, 5, 8, 0, 0,
		228, 21, 1, 0, 0, 0, 229, 230, 5, 10, 0, 0, 230, 231, 5, 41, 0, 0, 231,
		232, 5, 8, 0, 0, 232, 23, 1, 0, 0, 0, 233, 248, 3, 28, 14, 0, 234, 248,
		3, 36, 18, 0, 235, 248, 3, 38, 19, 0, 236, 248, 3, 40, 20, 0, 237, 248,
		3, 42, 21, 0, 238, 248, 3, 44, 22, 0, 239, 248, 3, 52, 26, 0, 240, 248,
		3, 54, 27, 0, 241, 248, 3, 68, 34, 0, 242, 248, 3, 72, 36, 0, 243, 248,
		3, 34, 17, 0, 244, 248, 3, 30, 15, 0, 245, 248, 3, 32, 16, 0, 246, 248,
		3, 60, 30, 0, 247, 233, 1, 0, 0, 0, 247, 234, 1, 0, 0, 0, 247, 235, 1,
		0, 0, 0, 247, 236, 1, 0, 0, 0, 247, 237, 1, 0, 0, 0, 247, 238, 1, 0, 0,
		0, 247, 239, 1, 0, 0, 0, 247, 240, 1, 0, 0, 0, 247, 241, 1, 0, 0, 0, 247,
		242, 1, 0, 0, 0, 247, 243, 1, 0, 0, 0, 247, 244, 1, 0, 0, 0, 247, 245,
		1, 0, 0, 0, 247, 246, 1, 0, 0, 0, 248, 25, 1, 0, 0, 0, 249, 250, 5, 46,
		0, 0, 250, 254, 5, 6, 0, 0, 251, 253, 5, 44, 0, 0, 252, 251, 1, 0, 0, 0,
		253, 256, 1, 0, 0, 0, 254, 252, 1, 0, 0, 0, 254, 255, 1, 0, 0, 0, 255,
		257, 1, 0, 0, 0, 256, 254, 1, 0, 0, 0, 257, 258, 3, 24, 12, 0, 258, 27,
		1, 0, 0, 0, 259, 260, 5, 11, 0, 0, 260, 261, 3, 24, 12, 0, 261, 29, 1,
		0, 0, 0, 262, 266, 5, 12, 0, 0, 263, 265, 5, 44, 0, 0, 264, 263, 1, 0,
		0, 0, 265, 268, 1, 0, 0, 0, 266, 264, 1, 0, 0, 0, 266, 267, 1, 0, 0, 0,
		267, 269, 1, 0, 0, 0, 268, 266, 1, 0, 0, 0, 269, 273, 5, 13, 0, 0, 270,
		272, 5, 44, 0, 0, 271, 270, 1, 0, 0, 0, 272, 275, 1, 0, 0, 0, 273, 271,
		1, 0, 0, 0, 273, 274, 1, 0, 0, 0, 274, 285, 1, 0, 0, 0, 275, 273, 1, 0,
		0, 0, 276, 280, 3, 24, 12, 0, 277, 279, 5, 44, 0, 0, 278, 277, 1, 0, 0,
		0, 279, 282, 1, 0, 0, 0, 280, 278, 1, 0, 0, 0, 280, 281, 1, 0, 0, 0, 281,
		284, 1, 0, 0, 0, 282, 280, 1, 0, 0, 0, 283, 276, 1, 0, 0, 0, 284, 287,
		1, 0, 0, 0, 285, 283, 1, 0, 0, 0, 285, 286, 1, 0, 0, 0, 286, 288, 1, 0,
		0, 0, 287, 285, 1, 0, 0, 0, 288, 292, 3, 24, 12, 0, 289, 291, 5, 44, 0,
		0, 290, 289, 1, 0, 0, 0, 291, 294, 1, 0, 0, 0, 292, 290, 1, 0, 0, 0, 292,
		293, 1, 0, 0, 0, 293, 295, 1, 0, 0, 0, 294, 292, 1, 0, 0, 0, 295, 296,
		5, 8, 0, 0, 296, 31, 1, 0, 0, 0, 297, 301, 5, 14, 0, 0, 298, 300, 5, 44,
		0, 0, 299, 298, 1, 0, 0, 0, 300, 303, 1, 0, 0, 0, 301, 299, 1, 0, 0, 0,
		301, 302, 1, 0, 0, 0, 302, 304, 1, 0, 0, 0, 303, 301, 1, 0, 0, 0, 304,
		308, 5, 13, 0, 0, 305, 307, 5, 44, 0, 0, 306, 305, 1, 0, 0, 0, 307, 310,
		1, 0, 0, 0, 308, 306, 1, 0, 0, 0, 308, 309, 1, 0, 0, 0, 309, 320, 1, 0,
		0, 0, 310, 308, 1, 0, 0, 0, 311, 315, 3, 24, 12, 0, 312, 314, 5, 44, 0,
		0, 313, 312, 1, 0, 0, 0, 314, 317, 1, 0, 0, 0, 315, 313, 1, 0, 0, 0, 315,
		316, 1, 0, 0, 0, 316, 319, 1, 0, 0, 0, 317, 315, 1, 0, 0, 0, 318, 311,
		1, 0, 0, 0, 319, 322, 1, 0, 0, 0, 320, 318, 1, 0, 0, 0, 320, 321, 1, 0,
		0, 0, 321, 323, 1, 0, 0, 0, 322, 320, 1, 0, 0, 0, 323, 327, 3, 24, 12,
		0, 324, 326, 5, 44, 0, 0, 325, 324, 1, 0, 0, 0, 326, 329, 1, 0, 0, 0, 327,
		325, 1, 0, 0, 0, 327, 328, 1, 0, 0, 0, 328, 330, 1, 0, 0, 0, 329, 327,
		1, 0, 0, 0, 330, 331, 5, 8, 0, 0, 331, 33, 1, 0, 0, 0, 332, 333, 3, 90,
		45, 0, 333, 35, 1, 0, 0, 0, 334, 335, 5, 15, 0, 0, 335, 37, 1, 0, 0, 0,
		336, 337, 5, 16, 0, 0, 337, 341, 5, 13, 0, 0, 338, 340, 5, 44, 0, 0, 339,
		338, 1, 0, 0, 0, 340, 343, 1, 0, 0, 0, 341, 339, 1, 0, 0, 0, 341, 342,
		1, 0, 0, 0, 342, 354, 1, 0, 0, 0, 343, 341, 1, 0, 0, 0, 344, 345, 3, 90,
		45, 0, 345, 349, 5, 17, 0, 0, 346, 348, 5, 44, 0, 0, 347, 346, 1, 0, 0,
		0, 348, 351, 1, 0, 0, 0, 349, 347, 1, 0, 0, 0, 349, 350, 1, 0, 0, 0, 350,
		353, 1, 0, 0, 0, 351, 349, 1, 0, 0, 0, 352, 344, 1, 0, 0, 0, 353, 356,
		1, 0, 0, 0, 354, 352, 1, 0, 0, 0, 354, 355, 1, 0, 0, 0, 355, 360, 1, 0,
		0, 0, 356, 354, 1, 0, 0, 0, 357, 359, 5, 44, 0, 0, 358, 357, 1, 0, 0, 0,
		359, 362, 1, 0, 0, 0, 360, 358, 1, 0, 0, 0, 360, 361, 1, 0, 0, 0, 361,
		363, 1, 0, 0, 0, 362, 360, 1, 0, 0, 0, 363, 367, 3, 90, 45, 0, 364, 366,
		5, 44, 0, 0, 365, 364, 1, 0, 0, 0, 366, 369, 1, 0, 0, 0, 367, 365, 1, 0,
		0, 0, 367, 368, 1, 0, 0, 0, 368, 370, 1, 0, 0, 0, 369, 367, 1, 0, 0, 0,
		370, 371, 5, 8, 0, 0, 371, 39, 1, 0, 0, 0, 372, 373, 5, 18, 0, 0, 373,
		374, 5, 13, 0, 0, 374, 375, 5, 8, 0, 0, 375, 41, 1, 0, 0, 0, 376, 377,
		5, 19, 0, 0, 377, 378, 5, 13, 0, 0, 378, 379, 5, 8, 0, 0, 379, 43, 1, 0,
		0, 0, 380, 381, 5, 20, 0, 0, 381, 382, 5, 13, 0, 0, 382, 383, 5, 8, 0,
		0, 383, 45, 1, 0, 0, 0, 384, 385, 5, 21, 0, 0, 385, 386, 3, 92, 46, 0,
		386, 387, 5, 8, 0, 0, 387, 47, 1, 0, 0, 0, 388, 389, 5, 22, 0, 0, 389,
		390, 3, 92, 46, 0, 390, 391, 5, 8, 0, 0, 391, 49, 1, 0, 0, 0, 392, 395,
		3, 48, 24, 0, 393, 395, 3, 46, 23, 0, 394, 392, 1, 0, 0, 0, 394, 393, 1,
		0, 0, 0, 395, 51, 1, 0, 0, 0, 396, 397, 5, 23, 0, 0, 397, 398, 5, 13, 0,
		0, 398, 399, 3, 90, 45, 0, 399, 403, 5, 8, 0, 0, 400, 402, 5, 44, 0, 0,
		401, 400, 1, 0, 0, 0, 402, 405, 1, 0, 0, 0, 403, 401, 1, 0, 0, 0, 403,
		404, 1, 0, 0, 0, 404, 415, 1, 0, 0, 0, 405, 403, 1, 0, 0, 0, 406, 410,
		3, 50, 25, 0, 407, 409, 5, 44, 0, 0, 408, 407, 1, 0, 0, 0, 409, 412, 1,
		0, 0, 0, 410, 408, 1, 0, 0, 0, 410, 411, 1, 0, 0, 0, 411, 414, 1, 0, 0,
		0, 412, 410, 1, 0, 0, 0, 413, 406, 1, 0, 0, 0, 414, 417, 1, 0, 0, 0, 415,
		413, 1, 0, 0, 0, 415, 416, 1, 0, 0, 0, 416, 53, 1, 0, 0, 0, 417, 415, 1,
		0, 0, 0, 418, 419, 5, 24, 0, 0, 419, 420, 5, 40, 0, 0, 420, 424, 5, 8,
		0, 0, 421, 423, 5, 44, 0, 0, 422, 421, 1, 0, 0, 0, 423, 426, 1, 0, 0, 0,
		424, 422, 1, 0, 0, 0, 424, 425, 1, 0, 0, 0, 425, 436, 1, 0, 0, 0, 426,
		424, 1, 0, 0, 0, 427, 431, 3, 56, 28, 0, 428, 430, 5, 44, 0, 0, 429, 428,
		1, 0, 0, 0, 430, 433, 1, 0, 0, 0, 431, 429, 1, 0, 0, 0, 431, 432, 1, 0,
		0, 0, 432, 435, 1, 0, 0, 0, 433, 431, 1, 0, 0, 0, 434, 427, 1, 0, 0, 0,
		435, 438, 1, 0, 0, 0, 436, 434, 1, 0, 0, 0, 436, 437, 1, 0, 0, 0, 437,
		55, 1, 0, 0, 0, 438, 436, 1, 0, 0, 0, 439, 443, 3, 16, 8, 0, 440, 443,
		3, 18, 9, 0, 441, 443, 3, 58, 29, 0, 442, 439, 1, 0, 0, 0, 442, 440, 1,
		0, 0, 0, 442, 441, 1, 0, 0, 0, 443, 57, 1, 0, 0, 0, 444, 445, 5, 25, 0,
		0, 445, 446, 5, 41, 0, 0, 446, 447, 5, 8, 0, 0, 447, 59, 1, 0, 0, 0, 448,
		449, 5, 26, 0, 0, 449, 450, 5, 13, 0, 0, 450, 451, 5, 43, 0, 0, 451, 455,
		5, 8, 0, 0, 452, 454, 5, 44, 0, 0, 453, 452, 1, 0, 0, 0, 454, 457, 1, 0,
		0, 0, 455, 453, 1, 0, 0, 0, 455, 456, 1, 0, 0, 0, 456, 467, 1, 0, 0, 0,
		457, 455, 1, 0, 0, 0, 458, 462, 3, 62, 31, 0, 459, 461, 5, 44, 0, 0, 460,
		459, 1, 0, 0, 0, 461, 464, 1, 0, 0, 0, 462, 460, 1, 0, 0, 0, 462, 463,
		1, 0, 0, 0, 463, 466, 1, 0, 0, 0, 464, 462, 1, 0, 0, 0, 465, 458, 1, 0,
		0, 0, 466, 469, 1, 0, 0, 0, 467, 465, 1, 0, 0, 0, 467, 468, 1, 0, 0, 0,
		468, 61, 1, 0, 0, 0, 469, 467, 1, 0, 0, 0, 470, 473, 3, 64, 32, 0, 471,
		473, 3, 66, 33, 0, 472, 470, 1, 0, 0, 0, 472, 471, 1, 0, 0, 0, 473, 63,
		1, 0, 0, 0, 474, 475, 5, 27, 0, 0, 475, 476, 5, 13, 0, 0, 476, 477, 3,
		12, 6, 0, 477, 478, 5, 8, 0, 0, 478, 65, 1, 0, 0, 0, 479, 480, 5, 28, 0,
		0, 480, 481, 5, 13, 0, 0, 481, 482, 3, 12, 6, 0, 482, 483, 5, 8, 0, 0,
		483, 67, 1, 0, 0, 0, 484, 485, 5, 29, 0, 0, 485, 486, 5, 13, 0, 0, 486,
		490, 5, 8, 0, 0, 487, 489, 5, 44, 0, 0, 488, 487, 1, 0, 0, 0, 489, 492,
		1, 0, 0, 0, 490, 488, 1, 0, 0, 0, 490, 491, 1, 0, 0, 0, 491, 502, 1, 0,
		0, 0, 492, 490, 1, 0, 0, 0, 493, 497, 3, 70, 35, 0, 494, 496, 5, 44, 0,
		0, 495, 494, 1, 0, 0, 0, 496, 499, 1, 0, 0, 0, 497, 495, 1, 0, 0, 0, 497,
		498, 1, 0, 0, 0, 498, 501, 1, 0, 0, 0, 499, 497, 1, 0, 0, 0, 500, 493,
		1, 0, 0, 0, 501, 504, 1, 0, 0, 0, 502, 500, 1, 0, 0, 0, 502, 503, 1, 0,
		0, 0, 503, 69, 1, 0, 0, 0, 504, 502, 1, 0, 0, 0, 505, 509, 3, 16, 8, 0,
		506, 509, 3, 18, 9, 0, 507, 509, 3, 20, 10, 0, 508, 505, 1, 0, 0, 0, 508,
		506, 1, 0, 0, 0, 508, 507, 1, 0, 0, 0, 509, 71, 1, 0, 0, 0, 510, 511, 5,
		30, 0, 0, 511, 512, 5, 13, 0, 0, 512, 513, 3, 12, 6, 0, 513, 517, 5, 8,
		0, 0, 514, 516, 5, 44, 0, 0, 515, 514, 1, 0, 0, 0, 516, 519, 1, 0, 0, 0,
		517, 515, 1, 0, 0, 0, 517, 518, 1, 0, 0, 0, 518, 529, 1, 0, 0, 0, 519,
		517, 1, 0, 0, 0, 520, 524, 3, 74, 37, 0, 521, 523, 5, 44, 0, 0, 522, 521,
		1, 0, 0, 0, 523, 526, 1, 0, 0, 0, 524, 522, 1, 0, 0, 0, 524, 525, 1, 0,
		0, 0, 525, 528, 1, 0, 0, 0, 526, 524, 1, 0, 0, 0, 527, 520, 1, 0, 0, 0,
		528, 531, 1, 0, 0, 0, 529, 527, 1, 0, 0, 0, 529, 530, 1, 0, 0, 0, 530,
		73, 1, 0, 0, 0, 531, 529, 1, 0, 0, 0, 532, 535, 3, 22, 11, 0, 533, 535,
		3, 18, 9, 0, 534, 532, 1, 0, 0, 0, 534, 533, 1, 0, 0, 0, 535, 75, 1, 0,
		0, 0, 536, 537, 5, 46, 0, 0, 537, 541, 5, 6, 0, 0, 538, 540, 5, 44, 0,
		0, 539, 538, 1, 0, 0, 0, 540, 543, 1, 0, 0, 0, 541, 539, 1, 0, 0, 0, 541,
		542, 1, 0, 0, 0, 542, 544, 1, 0, 0, 0, 543, 541, 1, 0, 0, 0, 544, 545,
		3, 78, 39, 0, 545, 77, 1, 0, 0, 0, 546, 552, 3, 82, 41, 0, 547, 552, 3,
		84, 42, 0, 548, 552, 3, 80, 40, 0, 549, 552, 3, 88, 44, 0, 550, 552, 3,
		86, 43, 0, 551, 546, 1, 0, 0, 0, 551, 547, 1, 0, 0, 0, 551, 548, 1, 0,
		0, 0, 551, 549, 1, 0, 0, 0, 551, 550, 1, 0, 0, 0, 552, 79, 1, 0, 0, 0,
		553, 554, 3, 90, 45, 0, 554, 81, 1, 0, 0, 0, 555, 556, 5, 31, 0, 0, 556,
		557, 5, 13, 0, 0, 557, 558, 3, 90, 45, 0, 558, 559, 5, 8, 0, 0, 559, 83,
		1, 0, 0, 0, 560, 564, 5, 32, 0, 0, 561, 563, 5, 44, 0, 0, 562, 561, 1,
		0, 0, 0, 563, 566, 1, 0, 0, 0, 564, 562, 1, 0, 0, 0, 564, 565, 1, 0, 0,
		0, 565, 576, 1, 0, 0, 0, 566, 564, 1, 0, 0, 0, 567, 571, 3, 78, 39, 0,
		568, 570, 5, 44, 0, 0, 569, 568, 1, 0, 0, 0, 570, 573, 1, 0, 0, 0, 571,
		569, 1, 0, 0, 0, 571, 572, 1, 0, 0, 0, 572, 575, 1, 0, 0, 0, 573, 571,
		1, 0, 0, 0, 574, 567, 1, 0, 0, 0, 575, 578, 1, 0, 0, 0, 576, 574, 1, 0,
		0, 0, 576, 577, 1, 0, 0, 0, 577, 582, 1, 0, 0, 0, 578, 576, 1, 0, 0, 0,
		579, 581, 5, 44, 0, 0, 580, 579, 1, 0, 0, 0, 581, 584, 1, 0, 0, 0, 582,
		580, 1, 0, 0, 0, 582, 583, 1, 0, 0, 0, 583, 585, 1, 0, 0, 0, 584, 582,
		1, 0, 0, 0, 585, 586, 5, 33, 0, 0, 586, 85, 1, 0, 0, 0, 587, 588, 5, 34,
		0, 0, 588, 589, 5, 13, 0, 0, 589, 590, 5, 8, 0, 0, 590, 87, 1, 0, 0, 0,
		591, 595, 5, 35, 0, 0, 592, 594, 5, 44, 0, 0, 593, 592, 1, 0, 0, 0, 594,
		597, 1, 0, 0, 0, 595, 593, 1, 0, 0, 0, 595, 596, 1, 0, 0, 0, 596, 598,
		1, 0, 0, 0, 597, 595, 1, 0, 0, 0, 598, 602, 5, 13, 0, 0, 599, 601, 5, 44,
		0, 0, 600, 599, 1, 0, 0, 0, 601, 604, 1, 0, 0, 0, 602, 600, 1, 0, 0, 0,
		602, 603, 1, 0, 0, 0, 603, 605, 1, 0, 0, 0, 604, 602, 1, 0, 0, 0, 605,
		609, 3, 24, 12, 0, 606, 608, 5, 44, 0, 0, 607, 606, 1, 0, 0, 0, 608, 611,
		1, 0, 0, 0, 609, 607, 1, 0, 0, 0, 609, 610, 1, 0, 0, 0, 610, 612, 1, 0,
		0, 0, 611, 609, 1, 0, 0, 0, 612, 616, 5, 8, 0, 0, 613, 615, 5, 44, 0, 0,
		614, 613, 1, 0, 0, 0, 615, 618, 1, 0, 0, 0, 616, 614, 1, 0, 0, 0, 616,
		617, 1, 0, 0, 0, 617, 619, 1, 0, 0, 0, 618, 616, 1, 0, 0, 0, 619, 620,
		3, 78, 39, 0, 620, 89, 1, 0, 0, 0, 621, 622, 5, 46, 0, 0, 622, 624, 5,
		36, 0, 0, 623, 621, 1, 0, 0, 0, 623, 624, 1, 0, 0, 0, 624, 625, 1, 0, 0,
		0, 625, 626, 5, 46, 0, 0, 626, 91, 1, 0, 0, 0, 627, 628, 7, 0, 0, 0, 628,
		93, 1, 0, 0, 0, 629, 633, 5, 37, 0, 0, 630, 632, 5, 44, 0, 0, 631, 630,
		1, 0, 0, 0, 632, 635, 1, 0, 0, 0, 633, 631, 1, 0, 0, 0, 633, 634, 1, 0,
		0, 0, 634, 636, 1, 0, 0, 0, 635, 633, 1, 0, 0, 0, 636, 640, 5, 2, 0, 0,
		637, 639, 5, 44, 0, 0, 638, 637, 1, 0, 0, 0, 639, 642, 1, 0, 0, 0, 640,
		638, 1, 0, 0, 0, 640, 641, 1, 0, 0, 0, 641, 652, 1, 0, 0, 0, 642, 640,
		1, 0, 0, 0, 643, 647, 3, 96, 48, 0, 644, 646, 5, 44, 0, 0, 645, 644, 1,
		0, 0, 0, 646, 649, 1, 0, 0, 0, 647, 645, 1, 0, 0, 0, 647, 648, 1, 0, 0,
		0, 648, 651, 1, 0, 0, 0, 649, 647, 1, 0, 0, 0, 650, 643, 1, 0, 0, 0, 651,
		654, 1, 0, 0, 0, 652, 650, 1, 0, 0, 0, 652, 653, 1, 0, 0, 0, 653, 658,
		1, 0, 0, 0, 654, 652, 1, 0, 0, 0, 655, 657, 5, 44, 0, 0, 656, 655, 1, 0,
		0, 0, 657, 660, 1, 0, 0, 0, 658, 656, 1, 0, 0, 0, 658, 659, 1, 0, 0, 0,
		659, 661, 1, 0, 0, 0, 660, 658, 1, 0, 0, 0, 661, 662, 5, 3, 0, 0, 662,
		95, 1, 0, 0, 0, 663, 664, 5, 46, 0, 0, 664, 665, 5, 6, 0, 0, 665, 666,
		3, 98, 49, 0, 666, 97, 1, 0, 0, 0, 667, 671, 5, 38, 0, 0, 668, 670, 5,
		44, 0, 0, 669, 668, 1, 0, 0, 0, 670, 673, 1, 0, 0, 0, 671, 669, 1, 0, 0,
		0, 671, 672, 1, 0, 0, 0, 672, 674, 1, 0, 0, 0, 673, 671, 1, 0, 0, 0, 674,
		678, 5, 41, 0, 0, 675, 677, 5, 44, 0, 0, 676, 675, 1, 0, 0, 0, 677, 680,
		1, 0, 0, 0, 678, 676, 1, 0, 0, 0, 678, 679, 1, 0, 0, 0, 679, 681, 1, 0,
		0, 0, 680, 678, 1, 0, 0, 0, 681, 685, 5, 8, 0, 0, 682, 684, 5, 44, 0, 0,
		683, 682, 1, 0, 0, 0, 684, 687, 1, 0, 0, 0, 685, 683, 1, 0, 0, 0, 685,
		686, 1, 0, 0, 0, 686, 697, 1, 0, 0, 0, 687, 685, 1, 0, 0, 0, 688, 692,
		3, 100, 50, 0, 689, 691, 5, 44, 0, 0, 690, 689, 1, 0, 0, 0, 691, 694, 1,
		0, 0, 0, 692, 690, 1, 0, 0, 0, 692, 693, 1, 0, 0, 0, 693, 696, 1, 0, 0,
		0, 694, 692, 1, 0, 0, 0, 695, 688, 1, 0, 0, 0, 696, 699, 1, 0, 0, 0, 697,
		695, 1, 0, 0, 0, 697, 698, 1, 0, 0, 0, 698, 99, 1, 0, 0, 0, 699, 697, 1,
		0, 0, 0, 700, 701, 3, 102, 51, 0, 701, 101, 1, 0, 0, 0, 702, 708, 5, 39,
		0, 0, 703, 704, 3, 92, 46, 0, 704, 705, 5, 17, 0, 0, 705, 707, 1, 0, 0,
		0, 706, 703, 1, 0, 0, 0, 707, 710, 1, 0, 0, 0, 708, 706, 1, 0, 0, 0, 708,
		709, 1, 0, 0, 0, 709, 711, 1, 0, 0, 0, 710, 708, 1, 0, 0, 0, 711, 712,
		3, 92, 46, 0, 712, 713, 5, 8, 0, 0, 713, 103, 1, 0, 0, 0, 77, 107, 114,
		119, 129, 136, 143, 148, 154, 161, 167, 174, 181, 186, 192, 200, 203, 207,
		214, 247, 254, 266, 273, 280, 285, 292, 301, 308, 315, 320, 327, 341, 349,
		354, 360, 367, 394, 403, 410, 415, 424, 431, 436, 442, 455, 462, 467, 472,
		490, 497, 502, 508, 517, 524, 529, 534, 541, 551, 564, 571, 576, 582, 595,
		602, 609, 616, 623, 633, 640, 647, 652, 658, 671, 678, 685, 692, 697, 708,
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
	MinecraftMetascriptParserT__36          = 37
	MinecraftMetascriptParserT__37          = 38
	MinecraftMetascriptParserT__38          = 39
	MinecraftMetascriptParserStoneDepthMode = 40
	MinecraftMetascriptParserInt            = 41
	MinecraftMetascriptParserFloat          = 42
	MinecraftMetascriptParserString_        = 43
	MinecraftMetascriptParserNL             = 44
	MinecraftMetascriptParserWS             = 45
	MinecraftMetascriptParserIdentifier     = 46
	MinecraftMetascriptParserBlockComment   = 47
	MinecraftMetascriptParserLineComment    = 48
)

// MinecraftMetascriptParser rules.
const (
	MinecraftMetascriptParserRULE_script                                                 = 0
	MinecraftMetascriptParserRULE_namespaceDeclaration                                   = 1
	MinecraftMetascriptParserRULE_namespace                                              = 2
	MinecraftMetascriptParserRULE_contentBlocks                                          = 3
	MinecraftMetascriptParserRULE_surfaceBlock                                           = 4
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
	MinecraftMetascriptParserRULE_surfaceCondition_NoiseThresholdBuilder_Min             = 23
	MinecraftMetascriptParserRULE_surfaceCondition_NoiseThresholdBuilder_Max             = 24
	MinecraftMetascriptParserRULE_surfaceCondition_NoiseThresholdBuilder                 = 25
	MinecraftMetascriptParserRULE_surfaceCondition_NoiseThreshold                        = 26
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
	MinecraftMetascriptParserRULE_noiseBlock                                             = 47
	MinecraftMetascriptParserRULE_noiseDeclaration                                       = 48
	MinecraftMetascriptParserRULE_noise                                                  = 49
	MinecraftMetascriptParserRULE_noise_Builder                                          = 50
	MinecraftMetascriptParserRULE_noise_Builder_Amplitudes                               = 51
)

// IScriptContext is an interface to support dynamic dispatch.
type IScriptContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode
	AllNamespace() []INamespaceContext
	Namespace(i int) INamespaceContext

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

func (s *ScriptContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(MinecraftMetascriptParserNL)
}

func (s *ScriptContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserNL, i)
}

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
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(104)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserT__0 {
		{
			p.SetState(110)
			p.Namespace()
		}
		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MinecraftMetascriptParserNL {
			{
				p.SetState(111)
				p.Match(MinecraftMetascriptParserNL)
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

		p.SetState(121)
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
		p.SetState(122)
		p.Match(MinecraftMetascriptParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(123)
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
		p.SetState(125)
		p.NamespaceDeclaration()
	}
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(126)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(132)
		p.Match(MinecraftMetascriptParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(136)
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
				p.SetState(133)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(138)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserT__3 || _la == MinecraftMetascriptParserT__36 {
		{
			p.SetState(139)
			p.ContentBlocks()
		}
		p.SetState(143)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(140)
					p.Match(MinecraftMetascriptParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(145)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
			if p.HasError() {
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
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(151)
			p.Match(MinecraftMetascriptParserNL)
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
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(157)
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
	SurfaceBlock() ISurfaceBlockContext
	NoiseBlock() INoiseBlockContext

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

func (s *ContentBlocksContext) SurfaceBlock() ISurfaceBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISurfaceBlockContext)
}

func (s *ContentBlocksContext) NoiseBlock() INoiseBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INoiseBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INoiseBlockContext)
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
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinecraftMetascriptParserT__3:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(159)
			p.SurfaceBlock()
		}

	case MinecraftMetascriptParserT__36:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(160)
			p.NoiseBlock()
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

// ISurfaceBlockContext is an interface to support dynamic dispatch.
type ISurfaceBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode
	AllSurfaceStatement() []ISurfaceStatementContext
	SurfaceStatement(i int) ISurfaceStatementContext

	// IsSurfaceBlockContext differentiates from other interfaces.
	IsSurfaceBlockContext()
}

type SurfaceBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceBlockContext() *SurfaceBlockContext {
	var p = new(SurfaceBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceBlock
	return p
}

func InitEmptySurfaceBlockContext(p *SurfaceBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceBlock
}

func (*SurfaceBlockContext) IsSurfaceBlockContext() {}

func NewSurfaceBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceBlockContext {
	var p = new(SurfaceBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceBlock

	return p
}

func (s *SurfaceBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceBlockContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(MinecraftMetascriptParserNL)
}

func (s *SurfaceBlockContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserNL, i)
}

func (s *SurfaceBlockContext) AllSurfaceStatement() []ISurfaceStatementContext {
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

func (s *SurfaceBlockContext) SurfaceStatement(i int) ISurfaceStatementContext {
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

func (s *SurfaceBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterSurfaceBlock(s)
	}
}

func (s *SurfaceBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitSurfaceBlock(s)
	}
}

func (p *MinecraftMetascriptParser) SurfaceBlock() (localctx ISurfaceBlockContext) {
	localctx = NewSurfaceBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MinecraftMetascriptParserRULE_surfaceBlock)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)
		p.Match(MinecraftMetascriptParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(164)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(170)
		p.Match(MinecraftMetascriptParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(171)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserIdentifier {
		{
			p.SetState(177)
			p.SurfaceStatement()
		}
		p.SetState(181)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(178)
					p.Match(MinecraftMetascriptParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(183)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

		p.SetState(188)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(189)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(194)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(195)
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
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(197)
			p.VerticalAnchorDeclaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(198)
			p.SurfaceConditionDeclaration()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(199)
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

	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinecraftMetascriptParserT__4, MinecraftMetascriptParserInt:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(203)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MinecraftMetascriptParserT__4 {
			{
				p.SetState(202)
				p.Match(MinecraftMetascriptParserT__4)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(205)
			p.Match(MinecraftMetascriptParserInt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MinecraftMetascriptParserIdentifier:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(206)
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
		p.SetState(209)
		p.Match(MinecraftMetascriptParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(210)
		p.Match(MinecraftMetascriptParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(211)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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
		p.SetState(219)
		p.Match(MinecraftMetascriptParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(220)
		p.Match(MinecraftMetascriptParserInt)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(221)
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
		p.SetState(223)
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
		p.SetState(225)
		p.Match(MinecraftMetascriptParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(226)
		p.Number()
	}
	{
		p.SetState(227)
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
		p.SetState(229)
		p.Match(MinecraftMetascriptParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(230)
		p.Match(MinecraftMetascriptParserInt)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(231)
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
	SurfaceCondition_NoiseThreshold() ISurfaceCondition_NoiseThresholdContext
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

func (s *SurfaceConditionContext) SurfaceCondition_NoiseThreshold() ISurfaceCondition_NoiseThresholdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceCondition_NoiseThresholdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISurfaceCondition_NoiseThresholdContext)
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
	p.SetState(247)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinecraftMetascriptParserT__10:
		{
			p.SetState(233)
			p.SurfaceCondition_Not()
		}

	case MinecraftMetascriptParserT__14:
		{
			p.SetState(234)
			p.SurfaceCondition_AboveSurface()
		}

	case MinecraftMetascriptParserT__15:
		{
			p.SetState(235)
			p.SurfaceCondition_Biome()
		}

	case MinecraftMetascriptParserT__17:
		{
			p.SetState(236)
			p.SurfaceCondition_Hole()
		}

	case MinecraftMetascriptParserT__18:
		{
			p.SetState(237)
			p.SurfaceCondition_Steep()
		}

	case MinecraftMetascriptParserT__19:
		{
			p.SetState(238)
			p.SurfaceCondition_Freezing()
		}

	case MinecraftMetascriptParserT__22:
		{
			p.SetState(239)
			p.SurfaceCondition_NoiseThreshold()
		}

	case MinecraftMetascriptParserT__23:
		{
			p.SetState(240)
			p.SurfaceCondition_StoneDepth()
		}

	case MinecraftMetascriptParserT__28:
		{
			p.SetState(241)
			p.SurfaceCondition_AboveWater()
		}

	case MinecraftMetascriptParserT__29:
		{
			p.SetState(242)
			p.SurfaceCondition_YAbove()
		}

	case MinecraftMetascriptParserIdentifier:
		{
			p.SetState(243)
			p.SurfaceCondition_Reference()
		}

	case MinecraftMetascriptParserT__11:
		{
			p.SetState(244)
			p.SurfaceCondition_And()
		}

	case MinecraftMetascriptParserT__13:
		{
			p.SetState(245)
			p.SurfaceCondition_Or()
		}

	case MinecraftMetascriptParserT__25:
		{
			p.SetState(246)
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
		p.SetState(249)
		p.Match(MinecraftMetascriptParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(250)
		p.Match(MinecraftMetascriptParserT__5)
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

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(251)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(256)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(257)
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
		p.SetState(259)
		p.Match(MinecraftMetascriptParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(260)
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
		p.SetState(262)
		p.Match(MinecraftMetascriptParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(263)
			p.Match(MinecraftMetascriptParserNL)
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
	}
	{
		p.SetState(269)
		p.Match(MinecraftMetascriptParserT__12)
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

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(270)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(275)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(285)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(276)
				p.SurfaceCondition()
			}
			p.SetState(280)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == MinecraftMetascriptParserNL {
				{
					p.SetState(277)
					p.Match(MinecraftMetascriptParserNL)
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

		}
		p.SetState(287)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(288)
		p.SurfaceCondition()
	}
	p.SetState(292)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(289)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(294)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(295)
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
		p.SetState(297)
		p.Match(MinecraftMetascriptParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(301)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(298)
			p.Match(MinecraftMetascriptParserNL)
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
	}
	{
		p.SetState(304)
		p.Match(MinecraftMetascriptParserT__12)
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

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(305)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(310)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(320)
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
				p.SetState(311)
				p.SurfaceCondition()
			}
			p.SetState(315)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == MinecraftMetascriptParserNL {
				{
					p.SetState(312)
					p.Match(MinecraftMetascriptParserNL)
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
			}

		}
		p.SetState(322)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(323)
		p.SurfaceCondition()
	}
	p.SetState(327)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(324)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(329)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(330)
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
		p.SetState(332)
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
		p.SetState(334)
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
		p.SetState(336)
		p.Match(MinecraftMetascriptParserT__15)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(337)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(341)
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
				p.SetState(338)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(343)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(354)
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
				p.SetState(344)
				p.ResourceReference()
			}
			{
				p.SetState(345)
				p.Match(MinecraftMetascriptParserT__16)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(349)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
			for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				if _alt == 1 {
					{
						p.SetState(346)
						p.Match(MinecraftMetascriptParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				p.SetState(351)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext())
				if p.HasError() {
					goto errorExit
				}
			}

		}
		p.SetState(356)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(360)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(357)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(362)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(363)
		p.ResourceReference()
	}
	p.SetState(367)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(364)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(369)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(370)
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
		p.SetState(372)
		p.Match(MinecraftMetascriptParserT__17)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(373)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(374)
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
		p.SetState(376)
		p.Match(MinecraftMetascriptParserT__18)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(377)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(378)
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
		p.SetState(380)
		p.Match(MinecraftMetascriptParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(381)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(382)
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

// ISurfaceCondition_NoiseThresholdBuilder_MinContext is an interface to support dynamic dispatch.
type ISurfaceCondition_NoiseThresholdBuilder_MinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Number() INumberContext

	// IsSurfaceCondition_NoiseThresholdBuilder_MinContext differentiates from other interfaces.
	IsSurfaceCondition_NoiseThresholdBuilder_MinContext()
}

type SurfaceCondition_NoiseThresholdBuilder_MinContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceCondition_NoiseThresholdBuilder_MinContext() *SurfaceCondition_NoiseThresholdBuilder_MinContext {
	var p = new(SurfaceCondition_NoiseThresholdBuilder_MinContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_NoiseThresholdBuilder_Min
	return p
}

func InitEmptySurfaceCondition_NoiseThresholdBuilder_MinContext(p *SurfaceCondition_NoiseThresholdBuilder_MinContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_NoiseThresholdBuilder_Min
}

func (*SurfaceCondition_NoiseThresholdBuilder_MinContext) IsSurfaceCondition_NoiseThresholdBuilder_MinContext() {
}

func NewSurfaceCondition_NoiseThresholdBuilder_MinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_NoiseThresholdBuilder_MinContext {
	var p = new(SurfaceCondition_NoiseThresholdBuilder_MinContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_NoiseThresholdBuilder_Min

	return p
}

func (s *SurfaceCondition_NoiseThresholdBuilder_MinContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceCondition_NoiseThresholdBuilder_MinContext) Number() INumberContext {
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

func (s *SurfaceCondition_NoiseThresholdBuilder_MinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_NoiseThresholdBuilder_MinContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_NoiseThresholdBuilder_MinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterSurfaceCondition_NoiseThresholdBuilder_Min(s)
	}
}

func (s *SurfaceCondition_NoiseThresholdBuilder_MinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitSurfaceCondition_NoiseThresholdBuilder_Min(s)
	}
}

func (p *MinecraftMetascriptParser) SurfaceCondition_NoiseThresholdBuilder_Min() (localctx ISurfaceCondition_NoiseThresholdBuilder_MinContext) {
	localctx = NewSurfaceCondition_NoiseThresholdBuilder_MinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, MinecraftMetascriptParserRULE_surfaceCondition_NoiseThresholdBuilder_Min)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(384)
		p.Match(MinecraftMetascriptParserT__20)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(385)
		p.Number()
	}
	{
		p.SetState(386)
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

// ISurfaceCondition_NoiseThresholdBuilder_MaxContext is an interface to support dynamic dispatch.
type ISurfaceCondition_NoiseThresholdBuilder_MaxContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Number() INumberContext

	// IsSurfaceCondition_NoiseThresholdBuilder_MaxContext differentiates from other interfaces.
	IsSurfaceCondition_NoiseThresholdBuilder_MaxContext()
}

type SurfaceCondition_NoiseThresholdBuilder_MaxContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceCondition_NoiseThresholdBuilder_MaxContext() *SurfaceCondition_NoiseThresholdBuilder_MaxContext {
	var p = new(SurfaceCondition_NoiseThresholdBuilder_MaxContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_NoiseThresholdBuilder_Max
	return p
}

func InitEmptySurfaceCondition_NoiseThresholdBuilder_MaxContext(p *SurfaceCondition_NoiseThresholdBuilder_MaxContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_NoiseThresholdBuilder_Max
}

func (*SurfaceCondition_NoiseThresholdBuilder_MaxContext) IsSurfaceCondition_NoiseThresholdBuilder_MaxContext() {
}

func NewSurfaceCondition_NoiseThresholdBuilder_MaxContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_NoiseThresholdBuilder_MaxContext {
	var p = new(SurfaceCondition_NoiseThresholdBuilder_MaxContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_NoiseThresholdBuilder_Max

	return p
}

func (s *SurfaceCondition_NoiseThresholdBuilder_MaxContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceCondition_NoiseThresholdBuilder_MaxContext) Number() INumberContext {
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

func (s *SurfaceCondition_NoiseThresholdBuilder_MaxContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_NoiseThresholdBuilder_MaxContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_NoiseThresholdBuilder_MaxContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterSurfaceCondition_NoiseThresholdBuilder_Max(s)
	}
}

func (s *SurfaceCondition_NoiseThresholdBuilder_MaxContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitSurfaceCondition_NoiseThresholdBuilder_Max(s)
	}
}

func (p *MinecraftMetascriptParser) SurfaceCondition_NoiseThresholdBuilder_Max() (localctx ISurfaceCondition_NoiseThresholdBuilder_MaxContext) {
	localctx = NewSurfaceCondition_NoiseThresholdBuilder_MaxContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, MinecraftMetascriptParserRULE_surfaceCondition_NoiseThresholdBuilder_Max)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(388)
		p.Match(MinecraftMetascriptParserT__21)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(389)
		p.Number()
	}
	{
		p.SetState(390)
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

// ISurfaceCondition_NoiseThresholdBuilderContext is an interface to support dynamic dispatch.
type ISurfaceCondition_NoiseThresholdBuilderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SurfaceCondition_NoiseThresholdBuilder_Max() ISurfaceCondition_NoiseThresholdBuilder_MaxContext
	SurfaceCondition_NoiseThresholdBuilder_Min() ISurfaceCondition_NoiseThresholdBuilder_MinContext

	// IsSurfaceCondition_NoiseThresholdBuilderContext differentiates from other interfaces.
	IsSurfaceCondition_NoiseThresholdBuilderContext()
}

type SurfaceCondition_NoiseThresholdBuilderContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceCondition_NoiseThresholdBuilderContext() *SurfaceCondition_NoiseThresholdBuilderContext {
	var p = new(SurfaceCondition_NoiseThresholdBuilderContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_NoiseThresholdBuilder
	return p
}

func InitEmptySurfaceCondition_NoiseThresholdBuilderContext(p *SurfaceCondition_NoiseThresholdBuilderContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_NoiseThresholdBuilder
}

func (*SurfaceCondition_NoiseThresholdBuilderContext) IsSurfaceCondition_NoiseThresholdBuilderContext() {
}

func NewSurfaceCondition_NoiseThresholdBuilderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_NoiseThresholdBuilderContext {
	var p = new(SurfaceCondition_NoiseThresholdBuilderContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_NoiseThresholdBuilder

	return p
}

func (s *SurfaceCondition_NoiseThresholdBuilderContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceCondition_NoiseThresholdBuilderContext) SurfaceCondition_NoiseThresholdBuilder_Max() ISurfaceCondition_NoiseThresholdBuilder_MaxContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceCondition_NoiseThresholdBuilder_MaxContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISurfaceCondition_NoiseThresholdBuilder_MaxContext)
}

func (s *SurfaceCondition_NoiseThresholdBuilderContext) SurfaceCondition_NoiseThresholdBuilder_Min() ISurfaceCondition_NoiseThresholdBuilder_MinContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceCondition_NoiseThresholdBuilder_MinContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISurfaceCondition_NoiseThresholdBuilder_MinContext)
}

func (s *SurfaceCondition_NoiseThresholdBuilderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_NoiseThresholdBuilderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_NoiseThresholdBuilderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterSurfaceCondition_NoiseThresholdBuilder(s)
	}
}

func (s *SurfaceCondition_NoiseThresholdBuilderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitSurfaceCondition_NoiseThresholdBuilder(s)
	}
}

func (p *MinecraftMetascriptParser) SurfaceCondition_NoiseThresholdBuilder() (localctx ISurfaceCondition_NoiseThresholdBuilderContext) {
	localctx = NewSurfaceCondition_NoiseThresholdBuilderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, MinecraftMetascriptParserRULE_surfaceCondition_NoiseThresholdBuilder)
	p.SetState(394)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinecraftMetascriptParserT__21:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(392)
			p.SurfaceCondition_NoiseThresholdBuilder_Max()
		}

	case MinecraftMetascriptParserT__20:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(393)
			p.SurfaceCondition_NoiseThresholdBuilder_Min()
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

// ISurfaceCondition_NoiseThresholdContext is an interface to support dynamic dispatch.
type ISurfaceCondition_NoiseThresholdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ResourceReference() IResourceReferenceContext
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode
	AllSurfaceCondition_NoiseThresholdBuilder() []ISurfaceCondition_NoiseThresholdBuilderContext
	SurfaceCondition_NoiseThresholdBuilder(i int) ISurfaceCondition_NoiseThresholdBuilderContext

	// IsSurfaceCondition_NoiseThresholdContext differentiates from other interfaces.
	IsSurfaceCondition_NoiseThresholdContext()
}

type SurfaceCondition_NoiseThresholdContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySurfaceCondition_NoiseThresholdContext() *SurfaceCondition_NoiseThresholdContext {
	var p = new(SurfaceCondition_NoiseThresholdContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_NoiseThreshold
	return p
}

func InitEmptySurfaceCondition_NoiseThresholdContext(p *SurfaceCondition_NoiseThresholdContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_NoiseThreshold
}

func (*SurfaceCondition_NoiseThresholdContext) IsSurfaceCondition_NoiseThresholdContext() {}

func NewSurfaceCondition_NoiseThresholdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceCondition_NoiseThresholdContext {
	var p = new(SurfaceCondition_NoiseThresholdContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_surfaceCondition_NoiseThreshold

	return p
}

func (s *SurfaceCondition_NoiseThresholdContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceCondition_NoiseThresholdContext) ResourceReference() IResourceReferenceContext {
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

func (s *SurfaceCondition_NoiseThresholdContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(MinecraftMetascriptParserNL)
}

func (s *SurfaceCondition_NoiseThresholdContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserNL, i)
}

func (s *SurfaceCondition_NoiseThresholdContext) AllSurfaceCondition_NoiseThresholdBuilder() []ISurfaceCondition_NoiseThresholdBuilderContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISurfaceCondition_NoiseThresholdBuilderContext); ok {
			len++
		}
	}

	tst := make([]ISurfaceCondition_NoiseThresholdBuilderContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISurfaceCondition_NoiseThresholdBuilderContext); ok {
			tst[i] = t.(ISurfaceCondition_NoiseThresholdBuilderContext)
			i++
		}
	}

	return tst
}

func (s *SurfaceCondition_NoiseThresholdContext) SurfaceCondition_NoiseThresholdBuilder(i int) ISurfaceCondition_NoiseThresholdBuilderContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISurfaceCondition_NoiseThresholdBuilderContext); ok {
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

	return t.(ISurfaceCondition_NoiseThresholdBuilderContext)
}

func (s *SurfaceCondition_NoiseThresholdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceCondition_NoiseThresholdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceCondition_NoiseThresholdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterSurfaceCondition_NoiseThreshold(s)
	}
}

func (s *SurfaceCondition_NoiseThresholdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitSurfaceCondition_NoiseThreshold(s)
	}
}

func (p *MinecraftMetascriptParser) SurfaceCondition_NoiseThreshold() (localctx ISurfaceCondition_NoiseThresholdContext) {
	localctx = NewSurfaceCondition_NoiseThresholdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, MinecraftMetascriptParserRULE_surfaceCondition_NoiseThreshold)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(396)
		p.Match(MinecraftMetascriptParserT__22)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(397)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(398)
		p.ResourceReference()
	}
	{
		p.SetState(399)
		p.Match(MinecraftMetascriptParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
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
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(400)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(405)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(415)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserT__20 || _la == MinecraftMetascriptParserT__21 {
		{
			p.SetState(406)
			p.SurfaceCondition_NoiseThresholdBuilder()
		}
		p.SetState(410)
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
					p.SetState(407)
					p.Match(MinecraftMetascriptParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(412)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

		p.SetState(417)
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
		p.SetState(418)
		p.Match(MinecraftMetascriptParserT__23)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(419)
		p.Match(MinecraftMetascriptParserStoneDepthMode)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(420)
		p.Match(MinecraftMetascriptParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(424)
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
				p.SetState(421)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(426)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(436)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&33555072) != 0 {
		{
			p.SetState(427)
			p.SurfaceCondition_StoneDepthBuilder()
		}
		p.SetState(431)
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
					p.SetState(428)
					p.Match(MinecraftMetascriptParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(433)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

		p.SetState(438)
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
	p.SetState(442)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinecraftMetascriptParserT__6:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(439)
			p.SharedBuilder_Offset()
		}

	case MinecraftMetascriptParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(440)
			p.SharedBuilder_Add()
		}

	case MinecraftMetascriptParserT__24:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(441)
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
		p.SetState(444)
		p.Match(MinecraftMetascriptParserT__24)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(445)
		p.Match(MinecraftMetascriptParserInt)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(446)
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
		p.SetState(448)
		p.Match(MinecraftMetascriptParserT__25)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(449)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(450)
		p.Match(MinecraftMetascriptParserString_)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(451)
		p.Match(MinecraftMetascriptParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(455)
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
				p.SetState(452)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(457)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(467)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserT__26 || _la == MinecraftMetascriptParserT__27 {
		{
			p.SetState(458)
			p.SurfaceCondition_VerticalGradientBuilder()
		}
		p.SetState(462)
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
					p.SetState(459)
					p.Match(MinecraftMetascriptParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(464)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 44, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

		p.SetState(469)
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
	p.SetState(472)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinecraftMetascriptParserT__26:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(470)
			p.SurfaceCondition_VerticalGradientBuilder_Top()
		}

	case MinecraftMetascriptParserT__27:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(471)
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
		p.SetState(474)
		p.Match(MinecraftMetascriptParserT__26)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(475)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(476)
		p.VerticalAnchor()
	}
	{
		p.SetState(477)
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
		p.SetState(479)
		p.Match(MinecraftMetascriptParserT__27)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(480)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(481)
		p.VerticalAnchor()
	}
	{
		p.SetState(482)
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
		p.SetState(484)
		p.Match(MinecraftMetascriptParserT__28)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(485)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(486)
		p.Match(MinecraftMetascriptParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(490)
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
				p.SetState(487)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(492)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 47, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(502)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1664) != 0 {
		{
			p.SetState(493)
			p.SurfaceCondition_AboveWaterBuilder()
		}
		p.SetState(497)
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
					p.SetState(494)
					p.Match(MinecraftMetascriptParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(499)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 48, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

		p.SetState(504)
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
	p.SetState(508)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinecraftMetascriptParserT__6:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(505)
			p.SharedBuilder_Offset()
		}

	case MinecraftMetascriptParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(506)
			p.SharedBuilder_Add()
		}

	case MinecraftMetascriptParserT__9:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(507)
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
		p.SetState(510)
		p.Match(MinecraftMetascriptParserT__29)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(511)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(512)
		p.VerticalAnchor()
	}
	{
		p.SetState(513)
		p.Match(MinecraftMetascriptParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(517)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 51, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(514)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(519)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 51, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(529)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserT__8 || _la == MinecraftMetascriptParserT__9 {
		{
			p.SetState(520)
			p.SurfaceCondition_YAboveBuilder()
		}
		p.SetState(524)
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
					p.SetState(521)
					p.Match(MinecraftMetascriptParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(526)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 52, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

		p.SetState(531)
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
	p.SetState(534)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinecraftMetascriptParserT__9:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(532)
			p.SharedBuilder_MulInt()
		}

	case MinecraftMetascriptParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(533)
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
		p.SetState(536)
		p.Match(MinecraftMetascriptParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(537)
		p.Match(MinecraftMetascriptParserT__5)
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

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(538)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(543)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(544)
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
	p.SetState(551)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinecraftMetascriptParserT__30:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(546)
			p.SurfaceRule_Block()
		}

	case MinecraftMetascriptParserT__31:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(547)
			p.SurfaceRule_Sequence()
		}

	case MinecraftMetascriptParserIdentifier:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(548)
			p.SurfaceRule_Reference()
		}

	case MinecraftMetascriptParserT__34:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(549)
			p.SurfaceRule_If()
		}

	case MinecraftMetascriptParserT__33:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(550)
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
		p.SetState(553)
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
		p.SetState(555)
		p.Match(MinecraftMetascriptParserT__30)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(556)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(557)
		p.ResourceReference()
	}
	{
		p.SetState(558)
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
		p.SetState(560)
		p.Match(MinecraftMetascriptParserT__31)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(564)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 57, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(561)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(566)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 57, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(576)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&70426726236160) != 0 {
		{
			p.SetState(567)
			p.SurfaceRule()
		}
		p.SetState(571)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 58, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(568)
					p.Match(MinecraftMetascriptParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(573)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 58, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

		p.SetState(578)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(582)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(579)
			p.Match(MinecraftMetascriptParserNL)
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
	}
	{
		p.SetState(585)
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
		p.SetState(587)
		p.Match(MinecraftMetascriptParserT__33)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(588)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(589)
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
		p.SetState(591)
		p.Match(MinecraftMetascriptParserT__34)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(595)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(592)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(597)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(598)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(602)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(599)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(604)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(605)
		p.SurfaceCondition()
	}
	p.SetState(609)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(606)
			p.Match(MinecraftMetascriptParserNL)
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
	}
	{
		p.SetState(612)
		p.Match(MinecraftMetascriptParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(616)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(613)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(618)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(619)
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
	p.SetState(623)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 65, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(621)
			p.Match(MinecraftMetascriptParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(622)
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
		p.SetState(625)
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
		p.SetState(627)
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

// INoiseBlockContext is an interface to support dynamic dispatch.
type INoiseBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode
	AllNoiseDeclaration() []INoiseDeclarationContext
	NoiseDeclaration(i int) INoiseDeclarationContext

	// IsNoiseBlockContext differentiates from other interfaces.
	IsNoiseBlockContext()
}

type NoiseBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNoiseBlockContext() *NoiseBlockContext {
	var p = new(NoiseBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_noiseBlock
	return p
}

func InitEmptyNoiseBlockContext(p *NoiseBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_noiseBlock
}

func (*NoiseBlockContext) IsNoiseBlockContext() {}

func NewNoiseBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NoiseBlockContext {
	var p = new(NoiseBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_noiseBlock

	return p
}

func (s *NoiseBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *NoiseBlockContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(MinecraftMetascriptParserNL)
}

func (s *NoiseBlockContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserNL, i)
}

func (s *NoiseBlockContext) AllNoiseDeclaration() []INoiseDeclarationContext {
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

func (s *NoiseBlockContext) NoiseDeclaration(i int) INoiseDeclarationContext {
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

func (s *NoiseBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NoiseBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NoiseBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterNoiseBlock(s)
	}
}

func (s *NoiseBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitNoiseBlock(s)
	}
}

func (p *MinecraftMetascriptParser) NoiseBlock() (localctx INoiseBlockContext) {
	localctx = NewNoiseBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, MinecraftMetascriptParserRULE_noiseBlock)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(629)
		p.Match(MinecraftMetascriptParserT__36)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(633)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(630)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(635)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(636)
		p.Match(MinecraftMetascriptParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(640)
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
				p.SetState(637)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(642)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 67, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(652)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserIdentifier {
		{
			p.SetState(643)
			p.NoiseDeclaration()
		}
		p.SetState(647)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 68, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(644)
					p.Match(MinecraftMetascriptParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(649)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 68, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

		p.SetState(654)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(658)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(655)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(660)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(661)
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

// INoiseDeclarationContext is an interface to support dynamic dispatch.
type INoiseDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	Noise() INoiseContext

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
	p.RuleIndex = MinecraftMetascriptParserRULE_noiseDeclaration
	return p
}

func InitEmptyNoiseDeclarationContext(p *NoiseDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_noiseDeclaration
}

func (*NoiseDeclarationContext) IsNoiseDeclarationContext() {}

func NewNoiseDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NoiseDeclarationContext {
	var p = new(NoiseDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_noiseDeclaration

	return p
}

func (s *NoiseDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *NoiseDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserIdentifier, 0)
}

func (s *NoiseDeclarationContext) Noise() INoiseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INoiseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INoiseContext)
}

func (s *NoiseDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NoiseDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NoiseDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterNoiseDeclaration(s)
	}
}

func (s *NoiseDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitNoiseDeclaration(s)
	}
}

func (p *MinecraftMetascriptParser) NoiseDeclaration() (localctx INoiseDeclarationContext) {
	localctx = NewNoiseDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, MinecraftMetascriptParserRULE_noiseDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(663)
		p.Match(MinecraftMetascriptParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(664)
		p.Match(MinecraftMetascriptParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(665)
		p.Noise()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INoiseContext is an interface to support dynamic dispatch.
type INoiseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Int() antlr.TerminalNode
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode
	AllNoise_Builder() []INoise_BuilderContext
	Noise_Builder(i int) INoise_BuilderContext

	// IsNoiseContext differentiates from other interfaces.
	IsNoiseContext()
}

type NoiseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNoiseContext() *NoiseContext {
	var p = new(NoiseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_noise
	return p
}

func InitEmptyNoiseContext(p *NoiseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_noise
}

func (*NoiseContext) IsNoiseContext() {}

func NewNoiseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NoiseContext {
	var p = new(NoiseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_noise

	return p
}

func (s *NoiseContext) GetParser() antlr.Parser { return s.parser }

func (s *NoiseContext) Int() antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserInt, 0)
}

func (s *NoiseContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(MinecraftMetascriptParserNL)
}

func (s *NoiseContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserNL, i)
}

func (s *NoiseContext) AllNoise_Builder() []INoise_BuilderContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INoise_BuilderContext); ok {
			len++
		}
	}

	tst := make([]INoise_BuilderContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INoise_BuilderContext); ok {
			tst[i] = t.(INoise_BuilderContext)
			i++
		}
	}

	return tst
}

func (s *NoiseContext) Noise_Builder(i int) INoise_BuilderContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INoise_BuilderContext); ok {
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

	return t.(INoise_BuilderContext)
}

func (s *NoiseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NoiseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NoiseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterNoise(s)
	}
}

func (s *NoiseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitNoise(s)
	}
}

func (p *MinecraftMetascriptParser) Noise() (localctx INoiseContext) {
	localctx = NewNoiseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, MinecraftMetascriptParserRULE_noise)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(667)
		p.Match(MinecraftMetascriptParserT__37)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(671)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(668)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(673)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(674)
		p.Match(MinecraftMetascriptParserInt)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(678)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(675)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(680)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(681)
		p.Match(MinecraftMetascriptParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(685)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 73, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(682)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(687)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 73, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(697)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserT__38 {
		{
			p.SetState(688)
			p.Noise_Builder()
		}
		p.SetState(692)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 74, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(689)
					p.Match(MinecraftMetascriptParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(694)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 74, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

		p.SetState(699)
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

// INoise_BuilderContext is an interface to support dynamic dispatch.
type INoise_BuilderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Noise_Builder_Amplitudes() INoise_Builder_AmplitudesContext

	// IsNoise_BuilderContext differentiates from other interfaces.
	IsNoise_BuilderContext()
}

type Noise_BuilderContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNoise_BuilderContext() *Noise_BuilderContext {
	var p = new(Noise_BuilderContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_noise_Builder
	return p
}

func InitEmptyNoise_BuilderContext(p *Noise_BuilderContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_noise_Builder
}

func (*Noise_BuilderContext) IsNoise_BuilderContext() {}

func NewNoise_BuilderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Noise_BuilderContext {
	var p = new(Noise_BuilderContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_noise_Builder

	return p
}

func (s *Noise_BuilderContext) GetParser() antlr.Parser { return s.parser }

func (s *Noise_BuilderContext) Noise_Builder_Amplitudes() INoise_Builder_AmplitudesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INoise_Builder_AmplitudesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INoise_Builder_AmplitudesContext)
}

func (s *Noise_BuilderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Noise_BuilderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Noise_BuilderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterNoise_Builder(s)
	}
}

func (s *Noise_BuilderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitNoise_Builder(s)
	}
}

func (p *MinecraftMetascriptParser) Noise_Builder() (localctx INoise_BuilderContext) {
	localctx = NewNoise_BuilderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, MinecraftMetascriptParserRULE_noise_Builder)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(700)
		p.Noise_Builder_Amplitudes()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INoise_Builder_AmplitudesContext is an interface to support dynamic dispatch.
type INoise_Builder_AmplitudesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNumber() []INumberContext
	Number(i int) INumberContext

	// IsNoise_Builder_AmplitudesContext differentiates from other interfaces.
	IsNoise_Builder_AmplitudesContext()
}

type Noise_Builder_AmplitudesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNoise_Builder_AmplitudesContext() *Noise_Builder_AmplitudesContext {
	var p = new(Noise_Builder_AmplitudesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_noise_Builder_Amplitudes
	return p
}

func InitEmptyNoise_Builder_AmplitudesContext(p *Noise_Builder_AmplitudesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_noise_Builder_Amplitudes
}

func (*Noise_Builder_AmplitudesContext) IsNoise_Builder_AmplitudesContext() {}

func NewNoise_Builder_AmplitudesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Noise_Builder_AmplitudesContext {
	var p = new(Noise_Builder_AmplitudesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_noise_Builder_Amplitudes

	return p
}

func (s *Noise_Builder_AmplitudesContext) GetParser() antlr.Parser { return s.parser }

func (s *Noise_Builder_AmplitudesContext) AllNumber() []INumberContext {
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

func (s *Noise_Builder_AmplitudesContext) Number(i int) INumberContext {
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

func (s *Noise_Builder_AmplitudesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Noise_Builder_AmplitudesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Noise_Builder_AmplitudesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterNoise_Builder_Amplitudes(s)
	}
}

func (s *Noise_Builder_AmplitudesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitNoise_Builder_Amplitudes(s)
	}
}

func (p *MinecraftMetascriptParser) Noise_Builder_Amplitudes() (localctx INoise_Builder_AmplitudesContext) {
	localctx = NewNoise_Builder_AmplitudesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, MinecraftMetascriptParserRULE_noise_Builder_Amplitudes)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(702)
		p.Match(MinecraftMetascriptParserT__38)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(708)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 76, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(703)
				p.Number()
			}
			{
				p.SetState(704)
				p.Match(MinecraftMetascriptParserT__16)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(710)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 76, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(711)
		p.Number()
	}
	{
		p.SetState(712)
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
