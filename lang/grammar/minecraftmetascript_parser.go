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
		"']'", "'Bandlands'", "'If'", "':'", "'Noise'", "'Noise('", "'.Octaves('",
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
		"number", "noiseBlock", "noiseDeclaration", "noiseDefinition", "noiseDefinition_Builder",
		"noiseDefinition_Builder_Octaves",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 48, 709, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 1, 0, 1,
		0, 5, 0, 107, 8, 0, 10, 0, 12, 0, 110, 9, 0, 5, 0, 112, 8, 0, 10, 0, 12,
		0, 115, 9, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 5, 2, 122, 8, 2, 10, 2, 12,
		2, 125, 9, 2, 1, 2, 1, 2, 5, 2, 129, 8, 2, 10, 2, 12, 2, 132, 9, 2, 1,
		2, 1, 2, 5, 2, 136, 8, 2, 10, 2, 12, 2, 139, 9, 2, 5, 2, 141, 8, 2, 10,
		2, 12, 2, 144, 9, 2, 1, 2, 5, 2, 147, 8, 2, 10, 2, 12, 2, 150, 9, 2, 1,
		2, 1, 2, 1, 3, 1, 3, 3, 3, 156, 8, 3, 1, 4, 1, 4, 5, 4, 160, 8, 4, 10,
		4, 12, 4, 163, 9, 4, 1, 4, 1, 4, 5, 4, 167, 8, 4, 10, 4, 12, 4, 170, 9,
		4, 1, 4, 1, 4, 5, 4, 174, 8, 4, 10, 4, 12, 4, 177, 9, 4, 5, 4, 179, 8,
		4, 10, 4, 12, 4, 182, 9, 4, 1, 4, 5, 4, 185, 8, 4, 10, 4, 12, 4, 188, 9,
		4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 3, 5, 195, 8, 5, 1, 6, 3, 6, 198, 8, 6,
		1, 6, 1, 6, 3, 6, 202, 8, 6, 1, 7, 1, 7, 1, 7, 5, 7, 207, 8, 7, 10, 7,
		12, 7, 210, 9, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		3, 12, 242, 8, 12, 1, 13, 1, 13, 1, 13, 5, 13, 247, 8, 13, 10, 13, 12,
		13, 250, 9, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 5, 15,
		259, 8, 15, 10, 15, 12, 15, 262, 9, 15, 1, 15, 1, 15, 5, 15, 266, 8, 15,
		10, 15, 12, 15, 269, 9, 15, 1, 15, 1, 15, 5, 15, 273, 8, 15, 10, 15, 12,
		15, 276, 9, 15, 5, 15, 278, 8, 15, 10, 15, 12, 15, 281, 9, 15, 1, 15, 1,
		15, 5, 15, 285, 8, 15, 10, 15, 12, 15, 288, 9, 15, 1, 15, 1, 15, 1, 16,
		1, 16, 5, 16, 294, 8, 16, 10, 16, 12, 16, 297, 9, 16, 1, 16, 1, 16, 5,
		16, 301, 8, 16, 10, 16, 12, 16, 304, 9, 16, 1, 16, 1, 16, 5, 16, 308, 8,
		16, 10, 16, 12, 16, 311, 9, 16, 5, 16, 313, 8, 16, 10, 16, 12, 16, 316,
		9, 16, 1, 16, 1, 16, 5, 16, 320, 8, 16, 10, 16, 12, 16, 323, 9, 16, 1,
		16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 5, 19, 334,
		8, 19, 10, 19, 12, 19, 337, 9, 19, 1, 19, 1, 19, 1, 19, 5, 19, 342, 8,
		19, 10, 19, 12, 19, 345, 9, 19, 5, 19, 347, 8, 19, 10, 19, 12, 19, 350,
		9, 19, 1, 19, 5, 19, 353, 8, 19, 10, 19, 12, 19, 356, 9, 19, 1, 19, 1,
		19, 5, 19, 360, 8, 19, 10, 19, 12, 19, 363, 9, 19, 1, 19, 1, 19, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25,
		3, 25, 389, 8, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 5, 26, 396, 8, 26,
		10, 26, 12, 26, 399, 9, 26, 1, 26, 1, 26, 5, 26, 403, 8, 26, 10, 26, 12,
		26, 406, 9, 26, 5, 26, 408, 8, 26, 10, 26, 12, 26, 411, 9, 26, 1, 27, 1,
		27, 1, 27, 1, 27, 5, 27, 417, 8, 27, 10, 27, 12, 27, 420, 9, 27, 1, 27,
		1, 27, 5, 27, 424, 8, 27, 10, 27, 12, 27, 427, 9, 27, 5, 27, 429, 8, 27,
		10, 27, 12, 27, 432, 9, 27, 1, 28, 1, 28, 1, 28, 3, 28, 437, 8, 28, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 5, 30, 448,
		8, 30, 10, 30, 12, 30, 451, 9, 30, 1, 30, 1, 30, 5, 30, 455, 8, 30, 10,
		30, 12, 30, 458, 9, 30, 5, 30, 460, 8, 30, 10, 30, 12, 30, 463, 9, 30,
		1, 31, 1, 31, 3, 31, 467, 8, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1,
		33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 5, 34, 483,
		8, 34, 10, 34, 12, 34, 486, 9, 34, 1, 34, 1, 34, 5, 34, 490, 8, 34, 10,
		34, 12, 34, 493, 9, 34, 5, 34, 495, 8, 34, 10, 34, 12, 34, 498, 9, 34,
		1, 35, 1, 35, 1, 35, 3, 35, 503, 8, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1,
		36, 5, 36, 510, 8, 36, 10, 36, 12, 36, 513, 9, 36, 1, 36, 1, 36, 5, 36,
		517, 8, 36, 10, 36, 12, 36, 520, 9, 36, 5, 36, 522, 8, 36, 10, 36, 12,
		36, 525, 9, 36, 1, 37, 1, 37, 3, 37, 529, 8, 37, 1, 38, 1, 38, 1, 38, 5,
		38, 534, 8, 38, 10, 38, 12, 38, 537, 9, 38, 1, 38, 1, 38, 1, 39, 1, 39,
		1, 39, 1, 39, 1, 39, 3, 39, 546, 8, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1,
		41, 1, 41, 1, 41, 1, 42, 1, 42, 5, 42, 557, 8, 42, 10, 42, 12, 42, 560,
		9, 42, 1, 42, 1, 42, 5, 42, 564, 8, 42, 10, 42, 12, 42, 567, 9, 42, 5,
		42, 569, 8, 42, 10, 42, 12, 42, 572, 9, 42, 1, 42, 5, 42, 575, 8, 42, 10,
		42, 12, 42, 578, 9, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 43, 1, 44,
		1, 44, 5, 44, 588, 8, 44, 10, 44, 12, 44, 591, 9, 44, 1, 44, 1, 44, 5,
		44, 595, 8, 44, 10, 44, 12, 44, 598, 9, 44, 1, 44, 1, 44, 5, 44, 602, 8,
		44, 10, 44, 12, 44, 605, 9, 44, 1, 44, 1, 44, 5, 44, 609, 8, 44, 10, 44,
		12, 44, 612, 9, 44, 1, 44, 1, 44, 1, 45, 1, 45, 3, 45, 618, 8, 45, 1, 45,
		1, 45, 1, 46, 1, 46, 1, 47, 1, 47, 5, 47, 626, 8, 47, 10, 47, 12, 47, 629,
		9, 47, 1, 47, 1, 47, 5, 47, 633, 8, 47, 10, 47, 12, 47, 636, 9, 47, 1,
		47, 1, 47, 5, 47, 640, 8, 47, 10, 47, 12, 47, 643, 9, 47, 5, 47, 645, 8,
		47, 10, 47, 12, 47, 648, 9, 47, 1, 47, 5, 47, 651, 8, 47, 10, 47, 12, 47,
		654, 9, 47, 1, 47, 1, 47, 1, 48, 1, 48, 1, 48, 1, 48, 1, 49, 1, 49, 5,
		49, 664, 8, 49, 10, 49, 12, 49, 667, 9, 49, 1, 49, 1, 49, 5, 49, 671, 8,
		49, 10, 49, 12, 49, 674, 9, 49, 1, 49, 1, 49, 5, 49, 678, 8, 49, 10, 49,
		12, 49, 681, 9, 49, 1, 49, 1, 49, 5, 49, 685, 8, 49, 10, 49, 12, 49, 688,
		9, 49, 5, 49, 690, 8, 49, 10, 49, 12, 49, 693, 9, 49, 1, 50, 1, 50, 1,
		51, 1, 51, 1, 51, 1, 51, 5, 51, 701, 8, 51, 10, 51, 12, 51, 704, 9, 51,
		1, 51, 1, 51, 1, 51, 1, 51, 0, 0, 52, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18,
		20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54,
		56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90,
		92, 94, 96, 98, 100, 102, 0, 1, 1, 0, 41, 42, 750, 0, 113, 1, 0, 0, 0,
		2, 116, 1, 0, 0, 0, 4, 119, 1, 0, 0, 0, 6, 155, 1, 0, 0, 0, 8, 157, 1,
		0, 0, 0, 10, 194, 1, 0, 0, 0, 12, 201, 1, 0, 0, 0, 14, 203, 1, 0, 0, 0,
		16, 213, 1, 0, 0, 0, 18, 217, 1, 0, 0, 0, 20, 219, 1, 0, 0, 0, 22, 223,
		1, 0, 0, 0, 24, 241, 1, 0, 0, 0, 26, 243, 1, 0, 0, 0, 28, 253, 1, 0, 0,
		0, 30, 256, 1, 0, 0, 0, 32, 291, 1, 0, 0, 0, 34, 326, 1, 0, 0, 0, 36, 328,
		1, 0, 0, 0, 38, 330, 1, 0, 0, 0, 40, 366, 1, 0, 0, 0, 42, 370, 1, 0, 0,
		0, 44, 374, 1, 0, 0, 0, 46, 378, 1, 0, 0, 0, 48, 382, 1, 0, 0, 0, 50, 388,
		1, 0, 0, 0, 52, 390, 1, 0, 0, 0, 54, 412, 1, 0, 0, 0, 56, 436, 1, 0, 0,
		0, 58, 438, 1, 0, 0, 0, 60, 442, 1, 0, 0, 0, 62, 466, 1, 0, 0, 0, 64, 468,
		1, 0, 0, 0, 66, 473, 1, 0, 0, 0, 68, 478, 1, 0, 0, 0, 70, 502, 1, 0, 0,
		0, 72, 504, 1, 0, 0, 0, 74, 528, 1, 0, 0, 0, 76, 530, 1, 0, 0, 0, 78, 545,
		1, 0, 0, 0, 80, 547, 1, 0, 0, 0, 82, 549, 1, 0, 0, 0, 84, 554, 1, 0, 0,
		0, 86, 581, 1, 0, 0, 0, 88, 585, 1, 0, 0, 0, 90, 617, 1, 0, 0, 0, 92, 621,
		1, 0, 0, 0, 94, 623, 1, 0, 0, 0, 96, 657, 1, 0, 0, 0, 98, 661, 1, 0, 0,
		0, 100, 694, 1, 0, 0, 0, 102, 696, 1, 0, 0, 0, 104, 108, 3, 4, 2, 0, 105,
		107, 5, 44, 0, 0, 106, 105, 1, 0, 0, 0, 107, 110, 1, 0, 0, 0, 108, 106,
		1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109, 112, 1, 0, 0, 0, 110, 108, 1, 0,
		0, 0, 111, 104, 1, 0, 0, 0, 112, 115, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0,
		113, 114, 1, 0, 0, 0, 114, 1, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 116, 117,
		5, 1, 0, 0, 117, 118, 5, 46, 0, 0, 118, 3, 1, 0, 0, 0, 119, 123, 3, 2,
		1, 0, 120, 122, 5, 44, 0, 0, 121, 120, 1, 0, 0, 0, 122, 125, 1, 0, 0, 0,
		123, 121, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 126, 1, 0, 0, 0, 125,
		123, 1, 0, 0, 0, 126, 130, 5, 2, 0, 0, 127, 129, 5, 44, 0, 0, 128, 127,
		1, 0, 0, 0, 129, 132, 1, 0, 0, 0, 130, 128, 1, 0, 0, 0, 130, 131, 1, 0,
		0, 0, 131, 142, 1, 0, 0, 0, 132, 130, 1, 0, 0, 0, 133, 137, 3, 6, 3, 0,
		134, 136, 5, 44, 0, 0, 135, 134, 1, 0, 0, 0, 136, 139, 1, 0, 0, 0, 137,
		135, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 141, 1, 0, 0, 0, 139, 137,
		1, 0, 0, 0, 140, 133, 1, 0, 0, 0, 141, 144, 1, 0, 0, 0, 142, 140, 1, 0,
		0, 0, 142, 143, 1, 0, 0, 0, 143, 148, 1, 0, 0, 0, 144, 142, 1, 0, 0, 0,
		145, 147, 5, 44, 0, 0, 146, 145, 1, 0, 0, 0, 147, 150, 1, 0, 0, 0, 148,
		146, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 151, 1, 0, 0, 0, 150, 148,
		1, 0, 0, 0, 151, 152, 5, 3, 0, 0, 152, 5, 1, 0, 0, 0, 153, 156, 3, 8, 4,
		0, 154, 156, 3, 94, 47, 0, 155, 153, 1, 0, 0, 0, 155, 154, 1, 0, 0, 0,
		156, 7, 1, 0, 0, 0, 157, 161, 5, 4, 0, 0, 158, 160, 5, 44, 0, 0, 159, 158,
		1, 0, 0, 0, 160, 163, 1, 0, 0, 0, 161, 159, 1, 0, 0, 0, 161, 162, 1, 0,
		0, 0, 162, 164, 1, 0, 0, 0, 163, 161, 1, 0, 0, 0, 164, 168, 5, 2, 0, 0,
		165, 167, 5, 44, 0, 0, 166, 165, 1, 0, 0, 0, 167, 170, 1, 0, 0, 0, 168,
		166, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169, 180, 1, 0, 0, 0, 170, 168,
		1, 0, 0, 0, 171, 175, 3, 10, 5, 0, 172, 174, 5, 44, 0, 0, 173, 172, 1,
		0, 0, 0, 174, 177, 1, 0, 0, 0, 175, 173, 1, 0, 0, 0, 175, 176, 1, 0, 0,
		0, 176, 179, 1, 0, 0, 0, 177, 175, 1, 0, 0, 0, 178, 171, 1, 0, 0, 0, 179,
		182, 1, 0, 0, 0, 180, 178, 1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181, 186,
		1, 0, 0, 0, 182, 180, 1, 0, 0, 0, 183, 185, 5, 44, 0, 0, 184, 183, 1, 0,
		0, 0, 185, 188, 1, 0, 0, 0, 186, 184, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0,
		187, 189, 1, 0, 0, 0, 188, 186, 1, 0, 0, 0, 189, 190, 5, 3, 0, 0, 190,
		9, 1, 0, 0, 0, 191, 195, 3, 14, 7, 0, 192, 195, 3, 26, 13, 0, 193, 195,
		3, 76, 38, 0, 194, 191, 1, 0, 0, 0, 194, 192, 1, 0, 0, 0, 194, 193, 1,
		0, 0, 0, 195, 11, 1, 0, 0, 0, 196, 198, 5, 5, 0, 0, 197, 196, 1, 0, 0,
		0, 197, 198, 1, 0, 0, 0, 198, 199, 1, 0, 0, 0, 199, 202, 5, 41, 0, 0, 200,
		202, 5, 46, 0, 0, 201, 197, 1, 0, 0, 0, 201, 200, 1, 0, 0, 0, 202, 13,
		1, 0, 0, 0, 203, 204, 5, 46, 0, 0, 204, 208, 5, 6, 0, 0, 205, 207, 5, 44,
		0, 0, 206, 205, 1, 0, 0, 0, 207, 210, 1, 0, 0, 0, 208, 206, 1, 0, 0, 0,
		208, 209, 1, 0, 0, 0, 209, 211, 1, 0, 0, 0, 210, 208, 1, 0, 0, 0, 211,
		212, 3, 12, 6, 0, 212, 15, 1, 0, 0, 0, 213, 214, 5, 7, 0, 0, 214, 215,
		5, 41, 0, 0, 215, 216, 5, 8, 0, 0, 216, 17, 1, 0, 0, 0, 217, 218, 5, 9,
		0, 0, 218, 19, 1, 0, 0, 0, 219, 220, 5, 10, 0, 0, 220, 221, 3, 92, 46,
		0, 221, 222, 5, 8, 0, 0, 222, 21, 1, 0, 0, 0, 223, 224, 5, 10, 0, 0, 224,
		225, 5, 41, 0, 0, 225, 226, 5, 8, 0, 0, 226, 23, 1, 0, 0, 0, 227, 242,
		3, 28, 14, 0, 228, 242, 3, 36, 18, 0, 229, 242, 3, 38, 19, 0, 230, 242,
		3, 40, 20, 0, 231, 242, 3, 42, 21, 0, 232, 242, 3, 44, 22, 0, 233, 242,
		3, 52, 26, 0, 234, 242, 3, 54, 27, 0, 235, 242, 3, 68, 34, 0, 236, 242,
		3, 72, 36, 0, 237, 242, 3, 34, 17, 0, 238, 242, 3, 30, 15, 0, 239, 242,
		3, 32, 16, 0, 240, 242, 3, 60, 30, 0, 241, 227, 1, 0, 0, 0, 241, 228, 1,
		0, 0, 0, 241, 229, 1, 0, 0, 0, 241, 230, 1, 0, 0, 0, 241, 231, 1, 0, 0,
		0, 241, 232, 1, 0, 0, 0, 241, 233, 1, 0, 0, 0, 241, 234, 1, 0, 0, 0, 241,
		235, 1, 0, 0, 0, 241, 236, 1, 0, 0, 0, 241, 237, 1, 0, 0, 0, 241, 238,
		1, 0, 0, 0, 241, 239, 1, 0, 0, 0, 241, 240, 1, 0, 0, 0, 242, 25, 1, 0,
		0, 0, 243, 244, 5, 46, 0, 0, 244, 248, 5, 6, 0, 0, 245, 247, 5, 44, 0,
		0, 246, 245, 1, 0, 0, 0, 247, 250, 1, 0, 0, 0, 248, 246, 1, 0, 0, 0, 248,
		249, 1, 0, 0, 0, 249, 251, 1, 0, 0, 0, 250, 248, 1, 0, 0, 0, 251, 252,
		3, 24, 12, 0, 252, 27, 1, 0, 0, 0, 253, 254, 5, 11, 0, 0, 254, 255, 3,
		24, 12, 0, 255, 29, 1, 0, 0, 0, 256, 260, 5, 12, 0, 0, 257, 259, 5, 44,
		0, 0, 258, 257, 1, 0, 0, 0, 259, 262, 1, 0, 0, 0, 260, 258, 1, 0, 0, 0,
		260, 261, 1, 0, 0, 0, 261, 263, 1, 0, 0, 0, 262, 260, 1, 0, 0, 0, 263,
		267, 5, 13, 0, 0, 264, 266, 5, 44, 0, 0, 265, 264, 1, 0, 0, 0, 266, 269,
		1, 0, 0, 0, 267, 265, 1, 0, 0, 0, 267, 268, 1, 0, 0, 0, 268, 279, 1, 0,
		0, 0, 269, 267, 1, 0, 0, 0, 270, 274, 3, 24, 12, 0, 271, 273, 5, 44, 0,
		0, 272, 271, 1, 0, 0, 0, 273, 276, 1, 0, 0, 0, 274, 272, 1, 0, 0, 0, 274,
		275, 1, 0, 0, 0, 275, 278, 1, 0, 0, 0, 276, 274, 1, 0, 0, 0, 277, 270,
		1, 0, 0, 0, 278, 281, 1, 0, 0, 0, 279, 277, 1, 0, 0, 0, 279, 280, 1, 0,
		0, 0, 280, 282, 1, 0, 0, 0, 281, 279, 1, 0, 0, 0, 282, 286, 3, 24, 12,
		0, 283, 285, 5, 44, 0, 0, 284, 283, 1, 0, 0, 0, 285, 288, 1, 0, 0, 0, 286,
		284, 1, 0, 0, 0, 286, 287, 1, 0, 0, 0, 287, 289, 1, 0, 0, 0, 288, 286,
		1, 0, 0, 0, 289, 290, 5, 8, 0, 0, 290, 31, 1, 0, 0, 0, 291, 295, 5, 14,
		0, 0, 292, 294, 5, 44, 0, 0, 293, 292, 1, 0, 0, 0, 294, 297, 1, 0, 0, 0,
		295, 293, 1, 0, 0, 0, 295, 296, 1, 0, 0, 0, 296, 298, 1, 0, 0, 0, 297,
		295, 1, 0, 0, 0, 298, 302, 5, 13, 0, 0, 299, 301, 5, 44, 0, 0, 300, 299,
		1, 0, 0, 0, 301, 304, 1, 0, 0, 0, 302, 300, 1, 0, 0, 0, 302, 303, 1, 0,
		0, 0, 303, 314, 1, 0, 0, 0, 304, 302, 1, 0, 0, 0, 305, 309, 3, 24, 12,
		0, 306, 308, 5, 44, 0, 0, 307, 306, 1, 0, 0, 0, 308, 311, 1, 0, 0, 0, 309,
		307, 1, 0, 0, 0, 309, 310, 1, 0, 0, 0, 310, 313, 1, 0, 0, 0, 311, 309,
		1, 0, 0, 0, 312, 305, 1, 0, 0, 0, 313, 316, 1, 0, 0, 0, 314, 312, 1, 0,
		0, 0, 314, 315, 1, 0, 0, 0, 315, 317, 1, 0, 0, 0, 316, 314, 1, 0, 0, 0,
		317, 321, 3, 24, 12, 0, 318, 320, 5, 44, 0, 0, 319, 318, 1, 0, 0, 0, 320,
		323, 1, 0, 0, 0, 321, 319, 1, 0, 0, 0, 321, 322, 1, 0, 0, 0, 322, 324,
		1, 0, 0, 0, 323, 321, 1, 0, 0, 0, 324, 325, 5, 8, 0, 0, 325, 33, 1, 0,
		0, 0, 326, 327, 3, 90, 45, 0, 327, 35, 1, 0, 0, 0, 328, 329, 5, 15, 0,
		0, 329, 37, 1, 0, 0, 0, 330, 331, 5, 16, 0, 0, 331, 335, 5, 13, 0, 0, 332,
		334, 5, 44, 0, 0, 333, 332, 1, 0, 0, 0, 334, 337, 1, 0, 0, 0, 335, 333,
		1, 0, 0, 0, 335, 336, 1, 0, 0, 0, 336, 348, 1, 0, 0, 0, 337, 335, 1, 0,
		0, 0, 338, 339, 3, 90, 45, 0, 339, 343, 5, 17, 0, 0, 340, 342, 5, 44, 0,
		0, 341, 340, 1, 0, 0, 0, 342, 345, 1, 0, 0, 0, 343, 341, 1, 0, 0, 0, 343,
		344, 1, 0, 0, 0, 344, 347, 1, 0, 0, 0, 345, 343, 1, 0, 0, 0, 346, 338,
		1, 0, 0, 0, 347, 350, 1, 0, 0, 0, 348, 346, 1, 0, 0, 0, 348, 349, 1, 0,
		0, 0, 349, 354, 1, 0, 0, 0, 350, 348, 1, 0, 0, 0, 351, 353, 5, 44, 0, 0,
		352, 351, 1, 0, 0, 0, 353, 356, 1, 0, 0, 0, 354, 352, 1, 0, 0, 0, 354,
		355, 1, 0, 0, 0, 355, 357, 1, 0, 0, 0, 356, 354, 1, 0, 0, 0, 357, 361,
		3, 90, 45, 0, 358, 360, 5, 44, 0, 0, 359, 358, 1, 0, 0, 0, 360, 363, 1,
		0, 0, 0, 361, 359, 1, 0, 0, 0, 361, 362, 1, 0, 0, 0, 362, 364, 1, 0, 0,
		0, 363, 361, 1, 0, 0, 0, 364, 365, 5, 8, 0, 0, 365, 39, 1, 0, 0, 0, 366,
		367, 5, 18, 0, 0, 367, 368, 5, 13, 0, 0, 368, 369, 5, 8, 0, 0, 369, 41,
		1, 0, 0, 0, 370, 371, 5, 19, 0, 0, 371, 372, 5, 13, 0, 0, 372, 373, 5,
		8, 0, 0, 373, 43, 1, 0, 0, 0, 374, 375, 5, 20, 0, 0, 375, 376, 5, 13, 0,
		0, 376, 377, 5, 8, 0, 0, 377, 45, 1, 0, 0, 0, 378, 379, 5, 21, 0, 0, 379,
		380, 3, 92, 46, 0, 380, 381, 5, 8, 0, 0, 381, 47, 1, 0, 0, 0, 382, 383,
		5, 22, 0, 0, 383, 384, 3, 92, 46, 0, 384, 385, 5, 8, 0, 0, 385, 49, 1,
		0, 0, 0, 386, 389, 3, 48, 24, 0, 387, 389, 3, 46, 23, 0, 388, 386, 1, 0,
		0, 0, 388, 387, 1, 0, 0, 0, 389, 51, 1, 0, 0, 0, 390, 391, 5, 23, 0, 0,
		391, 392, 5, 13, 0, 0, 392, 393, 3, 90, 45, 0, 393, 397, 5, 8, 0, 0, 394,
		396, 5, 44, 0, 0, 395, 394, 1, 0, 0, 0, 396, 399, 1, 0, 0, 0, 397, 395,
		1, 0, 0, 0, 397, 398, 1, 0, 0, 0, 398, 409, 1, 0, 0, 0, 399, 397, 1, 0,
		0, 0, 400, 404, 3, 50, 25, 0, 401, 403, 5, 44, 0, 0, 402, 401, 1, 0, 0,
		0, 403, 406, 1, 0, 0, 0, 404, 402, 1, 0, 0, 0, 404, 405, 1, 0, 0, 0, 405,
		408, 1, 0, 0, 0, 406, 404, 1, 0, 0, 0, 407, 400, 1, 0, 0, 0, 408, 411,
		1, 0, 0, 0, 409, 407, 1, 0, 0, 0, 409, 410, 1, 0, 0, 0, 410, 53, 1, 0,
		0, 0, 411, 409, 1, 0, 0, 0, 412, 413, 5, 24, 0, 0, 413, 414, 5, 40, 0,
		0, 414, 418, 5, 8, 0, 0, 415, 417, 5, 44, 0, 0, 416, 415, 1, 0, 0, 0, 417,
		420, 1, 0, 0, 0, 418, 416, 1, 0, 0, 0, 418, 419, 1, 0, 0, 0, 419, 430,
		1, 0, 0, 0, 420, 418, 1, 0, 0, 0, 421, 425, 3, 56, 28, 0, 422, 424, 5,
		44, 0, 0, 423, 422, 1, 0, 0, 0, 424, 427, 1, 0, 0, 0, 425, 423, 1, 0, 0,
		0, 425, 426, 1, 0, 0, 0, 426, 429, 1, 0, 0, 0, 427, 425, 1, 0, 0, 0, 428,
		421, 1, 0, 0, 0, 429, 432, 1, 0, 0, 0, 430, 428, 1, 0, 0, 0, 430, 431,
		1, 0, 0, 0, 431, 55, 1, 0, 0, 0, 432, 430, 1, 0, 0, 0, 433, 437, 3, 16,
		8, 0, 434, 437, 3, 18, 9, 0, 435, 437, 3, 58, 29, 0, 436, 433, 1, 0, 0,
		0, 436, 434, 1, 0, 0, 0, 436, 435, 1, 0, 0, 0, 437, 57, 1, 0, 0, 0, 438,
		439, 5, 25, 0, 0, 439, 440, 5, 41, 0, 0, 440, 441, 5, 8, 0, 0, 441, 59,
		1, 0, 0, 0, 442, 443, 5, 26, 0, 0, 443, 444, 5, 13, 0, 0, 444, 445, 5,
		43, 0, 0, 445, 449, 5, 8, 0, 0, 446, 448, 5, 44, 0, 0, 447, 446, 1, 0,
		0, 0, 448, 451, 1, 0, 0, 0, 449, 447, 1, 0, 0, 0, 449, 450, 1, 0, 0, 0,
		450, 461, 1, 0, 0, 0, 451, 449, 1, 0, 0, 0, 452, 456, 3, 62, 31, 0, 453,
		455, 5, 44, 0, 0, 454, 453, 1, 0, 0, 0, 455, 458, 1, 0, 0, 0, 456, 454,
		1, 0, 0, 0, 456, 457, 1, 0, 0, 0, 457, 460, 1, 0, 0, 0, 458, 456, 1, 0,
		0, 0, 459, 452, 1, 0, 0, 0, 460, 463, 1, 0, 0, 0, 461, 459, 1, 0, 0, 0,
		461, 462, 1, 0, 0, 0, 462, 61, 1, 0, 0, 0, 463, 461, 1, 0, 0, 0, 464, 467,
		3, 64, 32, 0, 465, 467, 3, 66, 33, 0, 466, 464, 1, 0, 0, 0, 466, 465, 1,
		0, 0, 0, 467, 63, 1, 0, 0, 0, 468, 469, 5, 27, 0, 0, 469, 470, 5, 13, 0,
		0, 470, 471, 3, 12, 6, 0, 471, 472, 5, 8, 0, 0, 472, 65, 1, 0, 0, 0, 473,
		474, 5, 28, 0, 0, 474, 475, 5, 13, 0, 0, 475, 476, 3, 12, 6, 0, 476, 477,
		5, 8, 0, 0, 477, 67, 1, 0, 0, 0, 478, 479, 5, 29, 0, 0, 479, 480, 5, 13,
		0, 0, 480, 484, 5, 8, 0, 0, 481, 483, 5, 44, 0, 0, 482, 481, 1, 0, 0, 0,
		483, 486, 1, 0, 0, 0, 484, 482, 1, 0, 0, 0, 484, 485, 1, 0, 0, 0, 485,
		496, 1, 0, 0, 0, 486, 484, 1, 0, 0, 0, 487, 491, 3, 70, 35, 0, 488, 490,
		5, 44, 0, 0, 489, 488, 1, 0, 0, 0, 490, 493, 1, 0, 0, 0, 491, 489, 1, 0,
		0, 0, 491, 492, 1, 0, 0, 0, 492, 495, 1, 0, 0, 0, 493, 491, 1, 0, 0, 0,
		494, 487, 1, 0, 0, 0, 495, 498, 1, 0, 0, 0, 496, 494, 1, 0, 0, 0, 496,
		497, 1, 0, 0, 0, 497, 69, 1, 0, 0, 0, 498, 496, 1, 0, 0, 0, 499, 503, 3,
		16, 8, 0, 500, 503, 3, 18, 9, 0, 501, 503, 3, 20, 10, 0, 502, 499, 1, 0,
		0, 0, 502, 500, 1, 0, 0, 0, 502, 501, 1, 0, 0, 0, 503, 71, 1, 0, 0, 0,
		504, 505, 5, 30, 0, 0, 505, 506, 5, 13, 0, 0, 506, 507, 3, 12, 6, 0, 507,
		511, 5, 8, 0, 0, 508, 510, 5, 44, 0, 0, 509, 508, 1, 0, 0, 0, 510, 513,
		1, 0, 0, 0, 511, 509, 1, 0, 0, 0, 511, 512, 1, 0, 0, 0, 512, 523, 1, 0,
		0, 0, 513, 511, 1, 0, 0, 0, 514, 518, 3, 74, 37, 0, 515, 517, 5, 44, 0,
		0, 516, 515, 1, 0, 0, 0, 517, 520, 1, 0, 0, 0, 518, 516, 1, 0, 0, 0, 518,
		519, 1, 0, 0, 0, 519, 522, 1, 0, 0, 0, 520, 518, 1, 0, 0, 0, 521, 514,
		1, 0, 0, 0, 522, 525, 1, 0, 0, 0, 523, 521, 1, 0, 0, 0, 523, 524, 1, 0,
		0, 0, 524, 73, 1, 0, 0, 0, 525, 523, 1, 0, 0, 0, 526, 529, 3, 22, 11, 0,
		527, 529, 3, 18, 9, 0, 528, 526, 1, 0, 0, 0, 528, 527, 1, 0, 0, 0, 529,
		75, 1, 0, 0, 0, 530, 531, 5, 46, 0, 0, 531, 535, 5, 6, 0, 0, 532, 534,
		5, 44, 0, 0, 533, 532, 1, 0, 0, 0, 534, 537, 1, 0, 0, 0, 535, 533, 1, 0,
		0, 0, 535, 536, 1, 0, 0, 0, 536, 538, 1, 0, 0, 0, 537, 535, 1, 0, 0, 0,
		538, 539, 3, 78, 39, 0, 539, 77, 1, 0, 0, 0, 540, 546, 3, 82, 41, 0, 541,
		546, 3, 84, 42, 0, 542, 546, 3, 80, 40, 0, 543, 546, 3, 88, 44, 0, 544,
		546, 3, 86, 43, 0, 545, 540, 1, 0, 0, 0, 545, 541, 1, 0, 0, 0, 545, 542,
		1, 0, 0, 0, 545, 543, 1, 0, 0, 0, 545, 544, 1, 0, 0, 0, 546, 79, 1, 0,
		0, 0, 547, 548, 3, 90, 45, 0, 548, 81, 1, 0, 0, 0, 549, 550, 5, 31, 0,
		0, 550, 551, 5, 13, 0, 0, 551, 552, 3, 90, 45, 0, 552, 553, 5, 8, 0, 0,
		553, 83, 1, 0, 0, 0, 554, 558, 5, 32, 0, 0, 555, 557, 5, 44, 0, 0, 556,
		555, 1, 0, 0, 0, 557, 560, 1, 0, 0, 0, 558, 556, 1, 0, 0, 0, 558, 559,
		1, 0, 0, 0, 559, 570, 1, 0, 0, 0, 560, 558, 1, 0, 0, 0, 561, 565, 3, 78,
		39, 0, 562, 564, 5, 44, 0, 0, 563, 562, 1, 0, 0, 0, 564, 567, 1, 0, 0,
		0, 565, 563, 1, 0, 0, 0, 565, 566, 1, 0, 0, 0, 566, 569, 1, 0, 0, 0, 567,
		565, 1, 0, 0, 0, 568, 561, 1, 0, 0, 0, 569, 572, 1, 0, 0, 0, 570, 568,
		1, 0, 0, 0, 570, 571, 1, 0, 0, 0, 571, 576, 1, 0, 0, 0, 572, 570, 1, 0,
		0, 0, 573, 575, 5, 44, 0, 0, 574, 573, 1, 0, 0, 0, 575, 578, 1, 0, 0, 0,
		576, 574, 1, 0, 0, 0, 576, 577, 1, 0, 0, 0, 577, 579, 1, 0, 0, 0, 578,
		576, 1, 0, 0, 0, 579, 580, 5, 33, 0, 0, 580, 85, 1, 0, 0, 0, 581, 582,
		5, 34, 0, 0, 582, 583, 5, 13, 0, 0, 583, 584, 5, 8, 0, 0, 584, 87, 1, 0,
		0, 0, 585, 589, 5, 35, 0, 0, 586, 588, 5, 44, 0, 0, 587, 586, 1, 0, 0,
		0, 588, 591, 1, 0, 0, 0, 589, 587, 1, 0, 0, 0, 589, 590, 1, 0, 0, 0, 590,
		592, 1, 0, 0, 0, 591, 589, 1, 0, 0, 0, 592, 596, 5, 13, 0, 0, 593, 595,
		5, 44, 0, 0, 594, 593, 1, 0, 0, 0, 595, 598, 1, 0, 0, 0, 596, 594, 1, 0,
		0, 0, 596, 597, 1, 0, 0, 0, 597, 599, 1, 0, 0, 0, 598, 596, 1, 0, 0, 0,
		599, 603, 3, 24, 12, 0, 600, 602, 5, 44, 0, 0, 601, 600, 1, 0, 0, 0, 602,
		605, 1, 0, 0, 0, 603, 601, 1, 0, 0, 0, 603, 604, 1, 0, 0, 0, 604, 606,
		1, 0, 0, 0, 605, 603, 1, 0, 0, 0, 606, 610, 5, 8, 0, 0, 607, 609, 5, 44,
		0, 0, 608, 607, 1, 0, 0, 0, 609, 612, 1, 0, 0, 0, 610, 608, 1, 0, 0, 0,
		610, 611, 1, 0, 0, 0, 611, 613, 1, 0, 0, 0, 612, 610, 1, 0, 0, 0, 613,
		614, 3, 78, 39, 0, 614, 89, 1, 0, 0, 0, 615, 616, 5, 46, 0, 0, 616, 618,
		5, 36, 0, 0, 617, 615, 1, 0, 0, 0, 617, 618, 1, 0, 0, 0, 618, 619, 1, 0,
		0, 0, 619, 620, 5, 46, 0, 0, 620, 91, 1, 0, 0, 0, 621, 622, 7, 0, 0, 0,
		622, 93, 1, 0, 0, 0, 623, 627, 5, 37, 0, 0, 624, 626, 5, 44, 0, 0, 625,
		624, 1, 0, 0, 0, 626, 629, 1, 0, 0, 0, 627, 625, 1, 0, 0, 0, 627, 628,
		1, 0, 0, 0, 628, 630, 1, 0, 0, 0, 629, 627, 1, 0, 0, 0, 630, 634, 5, 2,
		0, 0, 631, 633, 5, 44, 0, 0, 632, 631, 1, 0, 0, 0, 633, 636, 1, 0, 0, 0,
		634, 632, 1, 0, 0, 0, 634, 635, 1, 0, 0, 0, 635, 646, 1, 0, 0, 0, 636,
		634, 1, 0, 0, 0, 637, 641, 3, 96, 48, 0, 638, 640, 5, 44, 0, 0, 639, 638,
		1, 0, 0, 0, 640, 643, 1, 0, 0, 0, 641, 639, 1, 0, 0, 0, 641, 642, 1, 0,
		0, 0, 642, 645, 1, 0, 0, 0, 643, 641, 1, 0, 0, 0, 644, 637, 1, 0, 0, 0,
		645, 648, 1, 0, 0, 0, 646, 644, 1, 0, 0, 0, 646, 647, 1, 0, 0, 0, 647,
		652, 1, 0, 0, 0, 648, 646, 1, 0, 0, 0, 649, 651, 5, 44, 0, 0, 650, 649,
		1, 0, 0, 0, 651, 654, 1, 0, 0, 0, 652, 650, 1, 0, 0, 0, 652, 653, 1, 0,
		0, 0, 653, 655, 1, 0, 0, 0, 654, 652, 1, 0, 0, 0, 655, 656, 5, 3, 0, 0,
		656, 95, 1, 0, 0, 0, 657, 658, 5, 46, 0, 0, 658, 659, 5, 6, 0, 0, 659,
		660, 3, 98, 49, 0, 660, 97, 1, 0, 0, 0, 661, 665, 5, 38, 0, 0, 662, 664,
		5, 44, 0, 0, 663, 662, 1, 0, 0, 0, 664, 667, 1, 0, 0, 0, 665, 663, 1, 0,
		0, 0, 665, 666, 1, 0, 0, 0, 666, 668, 1, 0, 0, 0, 667, 665, 1, 0, 0, 0,
		668, 672, 5, 41, 0, 0, 669, 671, 5, 44, 0, 0, 670, 669, 1, 0, 0, 0, 671,
		674, 1, 0, 0, 0, 672, 670, 1, 0, 0, 0, 672, 673, 1, 0, 0, 0, 673, 675,
		1, 0, 0, 0, 674, 672, 1, 0, 0, 0, 675, 679, 5, 8, 0, 0, 676, 678, 5, 44,
		0, 0, 677, 676, 1, 0, 0, 0, 678, 681, 1, 0, 0, 0, 679, 677, 1, 0, 0, 0,
		679, 680, 1, 0, 0, 0, 680, 691, 1, 0, 0, 0, 681, 679, 1, 0, 0, 0, 682,
		686, 3, 100, 50, 0, 683, 685, 5, 44, 0, 0, 684, 683, 1, 0, 0, 0, 685, 688,
		1, 0, 0, 0, 686, 684, 1, 0, 0, 0, 686, 687, 1, 0, 0, 0, 687, 690, 1, 0,
		0, 0, 688, 686, 1, 0, 0, 0, 689, 682, 1, 0, 0, 0, 690, 693, 1, 0, 0, 0,
		691, 689, 1, 0, 0, 0, 691, 692, 1, 0, 0, 0, 692, 99, 1, 0, 0, 0, 693, 691,
		1, 0, 0, 0, 694, 695, 3, 102, 51, 0, 695, 101, 1, 0, 0, 0, 696, 702, 5,
		39, 0, 0, 697, 698, 3, 92, 46, 0, 698, 699, 5, 17, 0, 0, 699, 701, 1, 0,
		0, 0, 700, 697, 1, 0, 0, 0, 701, 704, 1, 0, 0, 0, 702, 700, 1, 0, 0, 0,
		702, 703, 1, 0, 0, 0, 703, 705, 1, 0, 0, 0, 704, 702, 1, 0, 0, 0, 705,
		706, 3, 92, 46, 0, 706, 707, 5, 8, 0, 0, 707, 103, 1, 0, 0, 0, 76, 108,
		113, 123, 130, 137, 142, 148, 155, 161, 168, 175, 180, 186, 194, 197, 201,
		208, 241, 248, 260, 267, 274, 279, 286, 295, 302, 309, 314, 321, 335, 343,
		348, 354, 361, 388, 397, 404, 409, 418, 425, 430, 436, 449, 456, 461, 466,
		484, 491, 496, 502, 511, 518, 523, 528, 535, 545, 558, 565, 570, 576, 589,
		596, 603, 610, 617, 627, 634, 641, 646, 652, 665, 672, 679, 686, 691, 702,
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
	MinecraftMetascriptParserRULE_noiseDefinition                                        = 49
	MinecraftMetascriptParserRULE_noiseDefinition_Builder                                = 50
	MinecraftMetascriptParserRULE_noiseDefinition_Builder_Octaves                        = 51
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
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserT__0 {
		{
			p.SetState(104)
			p.Namespace()
		}
		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MinecraftMetascriptParserNL {
			{
				p.SetState(105)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(110)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(115)
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
		p.SetState(116)
		p.Match(MinecraftMetascriptParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(117)
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
		p.SetState(119)
		p.NamespaceDeclaration()
	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(120)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(126)
		p.Match(MinecraftMetascriptParserT__1)
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
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(127)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserT__3 || _la == MinecraftMetascriptParserT__36 {
		{
			p.SetState(133)
			p.ContentBlocks()
		}
		p.SetState(137)
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
					p.SetState(134)
					p.Match(MinecraftMetascriptParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(139)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

		p.SetState(144)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinecraftMetascriptParserT__3:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(153)
			p.SurfaceBlock()
		}

	case MinecraftMetascriptParserT__36:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(154)
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
		p.SetState(157)
		p.Match(MinecraftMetascriptParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(158)
			p.Match(MinecraftMetascriptParserNL)
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
	}
	{
		p.SetState(164)
		p.Match(MinecraftMetascriptParserT__1)
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
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(165)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(170)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserIdentifier {
		{
			p.SetState(171)
			p.SurfaceStatement()
		}
		p.SetState(175)
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
					p.SetState(172)
					p.Match(MinecraftMetascriptParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(177)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(183)
			p.Match(MinecraftMetascriptParserNL)
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
	}
	{
		p.SetState(189)
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
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(191)
			p.VerticalAnchorDeclaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(192)
			p.SurfaceConditionDeclaration()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(193)
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

	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinecraftMetascriptParserT__4, MinecraftMetascriptParserInt:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(197)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MinecraftMetascriptParserT__4 {
			{
				p.SetState(196)
				p.Match(MinecraftMetascriptParserT__4)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(199)
			p.Match(MinecraftMetascriptParserInt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MinecraftMetascriptParserIdentifier:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(200)
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
		p.SetState(203)
		p.Match(MinecraftMetascriptParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(204)
		p.Match(MinecraftMetascriptParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(205)
			p.Match(MinecraftMetascriptParserNL)
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
	}
	{
		p.SetState(211)
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
		p.SetState(213)
		p.Match(MinecraftMetascriptParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(214)
		p.Match(MinecraftMetascriptParserInt)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(215)
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
		p.SetState(217)
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
		p.SetState(219)
		p.Match(MinecraftMetascriptParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(220)
		p.Number()
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
		p.SetState(223)
		p.Match(MinecraftMetascriptParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(224)
		p.Match(MinecraftMetascriptParserInt)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(225)
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
	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinecraftMetascriptParserT__10:
		{
			p.SetState(227)
			p.SurfaceCondition_Not()
		}

	case MinecraftMetascriptParserT__14:
		{
			p.SetState(228)
			p.SurfaceCondition_AboveSurface()
		}

	case MinecraftMetascriptParserT__15:
		{
			p.SetState(229)
			p.SurfaceCondition_Biome()
		}

	case MinecraftMetascriptParserT__17:
		{
			p.SetState(230)
			p.SurfaceCondition_Hole()
		}

	case MinecraftMetascriptParserT__18:
		{
			p.SetState(231)
			p.SurfaceCondition_Steep()
		}

	case MinecraftMetascriptParserT__19:
		{
			p.SetState(232)
			p.SurfaceCondition_Freezing()
		}

	case MinecraftMetascriptParserT__22:
		{
			p.SetState(233)
			p.SurfaceCondition_NoiseThreshold()
		}

	case MinecraftMetascriptParserT__23:
		{
			p.SetState(234)
			p.SurfaceCondition_StoneDepth()
		}

	case MinecraftMetascriptParserT__28:
		{
			p.SetState(235)
			p.SurfaceCondition_AboveWater()
		}

	case MinecraftMetascriptParserT__29:
		{
			p.SetState(236)
			p.SurfaceCondition_YAbove()
		}

	case MinecraftMetascriptParserIdentifier:
		{
			p.SetState(237)
			p.SurfaceCondition_Reference()
		}

	case MinecraftMetascriptParserT__11:
		{
			p.SetState(238)
			p.SurfaceCondition_And()
		}

	case MinecraftMetascriptParserT__13:
		{
			p.SetState(239)
			p.SurfaceCondition_Or()
		}

	case MinecraftMetascriptParserT__25:
		{
			p.SetState(240)
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
		p.SetState(243)
		p.Match(MinecraftMetascriptParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(244)
		p.Match(MinecraftMetascriptParserT__5)
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
		p.SetState(253)
		p.Match(MinecraftMetascriptParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(254)
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
		p.SetState(256)
		p.Match(MinecraftMetascriptParserT__11)
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

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(257)
			p.Match(MinecraftMetascriptParserNL)
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
	}
	{
		p.SetState(263)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(267)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(264)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(269)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
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

		}
		p.SetState(281)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(282)
		p.SurfaceCondition()
	}
	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(283)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(288)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(289)
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
		p.SetState(291)
		p.Match(MinecraftMetascriptParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(292)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(297)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(298)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(302)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(299)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(304)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(314)
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

		}
		p.SetState(316)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(317)
		p.SurfaceCondition()
	}
	p.SetState(321)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(318)
			p.Match(MinecraftMetascriptParserNL)
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
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(324)
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
		p.SetState(326)
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
		p.SetState(328)
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
		p.SetState(330)
		p.Match(MinecraftMetascriptParserT__15)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(331)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(335)
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
				p.SetState(332)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(337)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(348)
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
				p.SetState(338)
				p.ResourceReference()
			}
			{
				p.SetState(339)
				p.Match(MinecraftMetascriptParserT__16)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
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
			for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				if _alt == 1 {
					{
						p.SetState(340)
						p.Match(MinecraftMetascriptParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				p.SetState(345)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext())
				if p.HasError() {
					goto errorExit
				}
			}

		}
		p.SetState(350)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(354)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(351)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(356)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(357)
		p.ResourceReference()
	}
	p.SetState(361)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(358)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(363)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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
		p.SetState(366)
		p.Match(MinecraftMetascriptParserT__17)
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
		p.SetState(370)
		p.Match(MinecraftMetascriptParserT__18)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(371)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
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
		p.SetState(374)
		p.Match(MinecraftMetascriptParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(375)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(376)
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
		p.SetState(378)
		p.Match(MinecraftMetascriptParserT__20)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(379)
		p.Number()
	}
	{
		p.SetState(380)
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
		p.SetState(382)
		p.Match(MinecraftMetascriptParserT__21)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(383)
		p.Number()
	}
	{
		p.SetState(384)
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
	p.SetState(388)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinecraftMetascriptParserT__21:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(386)
			p.SurfaceCondition_NoiseThresholdBuilder_Max()
		}

	case MinecraftMetascriptParserT__20:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(387)
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
		p.SetState(390)
		p.Match(MinecraftMetascriptParserT__22)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(391)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(392)
		p.ResourceReference()
	}
	{
		p.SetState(393)
		p.Match(MinecraftMetascriptParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(397)
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
				p.SetState(394)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(399)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(409)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserT__20 || _la == MinecraftMetascriptParserT__21 {
		{
			p.SetState(400)
			p.SurfaceCondition_NoiseThresholdBuilder()
		}
		p.SetState(404)
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
					p.SetState(401)
					p.Match(MinecraftMetascriptParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(406)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

		p.SetState(411)
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
		p.SetState(412)
		p.Match(MinecraftMetascriptParserT__23)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(413)
		p.Match(MinecraftMetascriptParserStoneDepthMode)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(414)
		p.Match(MinecraftMetascriptParserT__7)
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
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(415)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(420)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(430)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&33555072) != 0 {
		{
			p.SetState(421)
			p.SurfaceCondition_StoneDepthBuilder()
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

		p.SetState(432)
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
	p.SetState(436)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinecraftMetascriptParserT__6:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(433)
			p.SharedBuilder_Offset()
		}

	case MinecraftMetascriptParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(434)
			p.SharedBuilder_Add()
		}

	case MinecraftMetascriptParserT__24:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(435)
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
		p.SetState(438)
		p.Match(MinecraftMetascriptParserT__24)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(439)
		p.Match(MinecraftMetascriptParserInt)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(440)
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
		p.SetState(442)
		p.Match(MinecraftMetascriptParserT__25)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(443)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(444)
		p.Match(MinecraftMetascriptParserString_)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(445)
		p.Match(MinecraftMetascriptParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(449)
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
				p.SetState(446)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(451)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(461)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserT__26 || _la == MinecraftMetascriptParserT__27 {
		{
			p.SetState(452)
			p.SurfaceCondition_VerticalGradientBuilder()
		}
		p.SetState(456)
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
					p.SetState(453)
					p.Match(MinecraftMetascriptParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(458)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

		p.SetState(463)
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
	p.SetState(466)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinecraftMetascriptParserT__26:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(464)
			p.SurfaceCondition_VerticalGradientBuilder_Top()
		}

	case MinecraftMetascriptParserT__27:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(465)
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
		p.SetState(468)
		p.Match(MinecraftMetascriptParserT__26)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(469)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(470)
		p.VerticalAnchor()
	}
	{
		p.SetState(471)
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
		p.SetState(473)
		p.Match(MinecraftMetascriptParserT__27)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(474)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(475)
		p.VerticalAnchor()
	}
	{
		p.SetState(476)
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
		p.SetState(478)
		p.Match(MinecraftMetascriptParserT__28)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(479)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(480)
		p.Match(MinecraftMetascriptParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(484)
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
				p.SetState(481)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(486)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 46, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(496)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1664) != 0 {
		{
			p.SetState(487)
			p.SurfaceCondition_AboveWaterBuilder()
		}
		p.SetState(491)
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
					p.SetState(488)
					p.Match(MinecraftMetascriptParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(493)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 47, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

		p.SetState(498)
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
	p.SetState(502)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinecraftMetascriptParserT__6:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(499)
			p.SharedBuilder_Offset()
		}

	case MinecraftMetascriptParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(500)
			p.SharedBuilder_Add()
		}

	case MinecraftMetascriptParserT__9:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(501)
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
		p.SetState(504)
		p.Match(MinecraftMetascriptParserT__29)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(505)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(506)
		p.VerticalAnchor()
	}
	{
		p.SetState(507)
		p.Match(MinecraftMetascriptParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(511)
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
				p.SetState(508)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(513)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 50, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(523)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserT__8 || _la == MinecraftMetascriptParserT__9 {
		{
			p.SetState(514)
			p.SurfaceCondition_YAboveBuilder()
		}
		p.SetState(518)
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
					p.SetState(515)
					p.Match(MinecraftMetascriptParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(520)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 51, p.GetParserRuleContext())
			if p.HasError() {
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

errorExit:
	if p.HasError() {
		v := p.GetError()
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
	p.SetState(528)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinecraftMetascriptParserT__9:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(526)
			p.SharedBuilder_MulInt()
		}

	case MinecraftMetascriptParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(527)
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
		p.SetState(530)
		p.Match(MinecraftMetascriptParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(531)
		p.Match(MinecraftMetascriptParserT__5)
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

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(532)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(537)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(538)
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
	p.SetState(545)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinecraftMetascriptParserT__30:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(540)
			p.SurfaceRule_Block()
		}

	case MinecraftMetascriptParserT__31:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(541)
			p.SurfaceRule_Sequence()
		}

	case MinecraftMetascriptParserIdentifier:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(542)
			p.SurfaceRule_Reference()
		}

	case MinecraftMetascriptParserT__34:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(543)
			p.SurfaceRule_If()
		}

	case MinecraftMetascriptParserT__33:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(544)
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
		p.SetState(547)
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
		p.SetState(549)
		p.Match(MinecraftMetascriptParserT__30)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(550)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(551)
		p.ResourceReference()
	}
	{
		p.SetState(552)
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
		p.SetState(554)
		p.Match(MinecraftMetascriptParserT__31)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(558)
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
				p.SetState(555)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(560)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 56, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(570)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&70426726236160) != 0 {
		{
			p.SetState(561)
			p.SurfaceRule()
		}
		p.SetState(565)
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
					p.SetState(562)
					p.Match(MinecraftMetascriptParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(567)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 57, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

		p.SetState(572)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(576)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(573)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
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
	{
		p.SetState(579)
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
		p.SetState(581)
		p.Match(MinecraftMetascriptParserT__33)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(582)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(583)
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
		p.SetState(585)
		p.Match(MinecraftMetascriptParserT__34)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(589)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(586)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(591)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(592)
		p.Match(MinecraftMetascriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(596)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(593)
			p.Match(MinecraftMetascriptParserNL)
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
	}
	{
		p.SetState(599)
		p.SurfaceCondition()
	}
	p.SetState(603)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(600)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(605)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(606)
		p.Match(MinecraftMetascriptParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(610)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(607)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(612)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(613)
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
	p.SetState(617)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 64, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(615)
			p.Match(MinecraftMetascriptParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(616)
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
		p.SetState(619)
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
		p.SetState(621)
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
		p.SetState(623)
		p.Match(MinecraftMetascriptParserT__36)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(627)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(624)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(629)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(630)
		p.Match(MinecraftMetascriptParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(634)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 66, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(631)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(636)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 66, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(646)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserIdentifier {
		{
			p.SetState(637)
			p.NoiseDeclaration()
		}
		p.SetState(641)
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
					p.SetState(638)
					p.Match(MinecraftMetascriptParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(643)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 67, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

		p.SetState(648)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(652)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(649)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
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
	{
		p.SetState(655)
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
	NoiseDefinition() INoiseDefinitionContext

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
		p.SetState(657)
		p.Match(MinecraftMetascriptParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(658)
		p.Match(MinecraftMetascriptParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(659)
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
	Int() antlr.TerminalNode
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode
	AllNoiseDefinition_Builder() []INoiseDefinition_BuilderContext
	NoiseDefinition_Builder(i int) INoiseDefinition_BuilderContext

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
	p.RuleIndex = MinecraftMetascriptParserRULE_noiseDefinition
	return p
}

func InitEmptyNoiseDefinitionContext(p *NoiseDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_noiseDefinition
}

func (*NoiseDefinitionContext) IsNoiseDefinitionContext() {}

func NewNoiseDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NoiseDefinitionContext {
	var p = new(NoiseDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_noiseDefinition

	return p
}

func (s *NoiseDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *NoiseDefinitionContext) Int() antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserInt, 0)
}

func (s *NoiseDefinitionContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(MinecraftMetascriptParserNL)
}

func (s *NoiseDefinitionContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(MinecraftMetascriptParserNL, i)
}

func (s *NoiseDefinitionContext) AllNoiseDefinition_Builder() []INoiseDefinition_BuilderContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INoiseDefinition_BuilderContext); ok {
			len++
		}
	}

	tst := make([]INoiseDefinition_BuilderContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INoiseDefinition_BuilderContext); ok {
			tst[i] = t.(INoiseDefinition_BuilderContext)
			i++
		}
	}

	return tst
}

func (s *NoiseDefinitionContext) NoiseDefinition_Builder(i int) INoiseDefinition_BuilderContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INoiseDefinition_BuilderContext); ok {
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

	return t.(INoiseDefinition_BuilderContext)
}

func (s *NoiseDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NoiseDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NoiseDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterNoiseDefinition(s)
	}
}

func (s *NoiseDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitNoiseDefinition(s)
	}
}

func (p *MinecraftMetascriptParser) NoiseDefinition() (localctx INoiseDefinitionContext) {
	localctx = NewNoiseDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, MinecraftMetascriptParserRULE_noiseDefinition)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(661)
		p.Match(MinecraftMetascriptParserT__37)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(665)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(662)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(667)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(668)
		p.Match(MinecraftMetascriptParserInt)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(672)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserNL {
		{
			p.SetState(669)
			p.Match(MinecraftMetascriptParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(674)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(675)
		p.Match(MinecraftMetascriptParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(679)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 72, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(676)
				p.Match(MinecraftMetascriptParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(681)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 72, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(691)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinecraftMetascriptParserT__38 {
		{
			p.SetState(682)
			p.NoiseDefinition_Builder()
		}
		p.SetState(686)
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
					p.SetState(683)
					p.Match(MinecraftMetascriptParserNL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(688)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 73, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

		p.SetState(693)
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

// INoiseDefinition_BuilderContext is an interface to support dynamic dispatch.
type INoiseDefinition_BuilderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NoiseDefinition_Builder_Octaves() INoiseDefinition_Builder_OctavesContext

	// IsNoiseDefinition_BuilderContext differentiates from other interfaces.
	IsNoiseDefinition_BuilderContext()
}

type NoiseDefinition_BuilderContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNoiseDefinition_BuilderContext() *NoiseDefinition_BuilderContext {
	var p = new(NoiseDefinition_BuilderContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_noiseDefinition_Builder
	return p
}

func InitEmptyNoiseDefinition_BuilderContext(p *NoiseDefinition_BuilderContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_noiseDefinition_Builder
}

func (*NoiseDefinition_BuilderContext) IsNoiseDefinition_BuilderContext() {}

func NewNoiseDefinition_BuilderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NoiseDefinition_BuilderContext {
	var p = new(NoiseDefinition_BuilderContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_noiseDefinition_Builder

	return p
}

func (s *NoiseDefinition_BuilderContext) GetParser() antlr.Parser { return s.parser }

func (s *NoiseDefinition_BuilderContext) NoiseDefinition_Builder_Octaves() INoiseDefinition_Builder_OctavesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INoiseDefinition_Builder_OctavesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INoiseDefinition_Builder_OctavesContext)
}

func (s *NoiseDefinition_BuilderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NoiseDefinition_BuilderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NoiseDefinition_BuilderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterNoiseDefinition_Builder(s)
	}
}

func (s *NoiseDefinition_BuilderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitNoiseDefinition_Builder(s)
	}
}

func (p *MinecraftMetascriptParser) NoiseDefinition_Builder() (localctx INoiseDefinition_BuilderContext) {
	localctx = NewNoiseDefinition_BuilderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, MinecraftMetascriptParserRULE_noiseDefinition_Builder)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(694)
		p.NoiseDefinition_Builder_Octaves()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INoiseDefinition_Builder_OctavesContext is an interface to support dynamic dispatch.
type INoiseDefinition_Builder_OctavesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNumber() []INumberContext
	Number(i int) INumberContext

	// IsNoiseDefinition_Builder_OctavesContext differentiates from other interfaces.
	IsNoiseDefinition_Builder_OctavesContext()
}

type NoiseDefinition_Builder_OctavesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNoiseDefinition_Builder_OctavesContext() *NoiseDefinition_Builder_OctavesContext {
	var p = new(NoiseDefinition_Builder_OctavesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_noiseDefinition_Builder_Octaves
	return p
}

func InitEmptyNoiseDefinition_Builder_OctavesContext(p *NoiseDefinition_Builder_OctavesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinecraftMetascriptParserRULE_noiseDefinition_Builder_Octaves
}

func (*NoiseDefinition_Builder_OctavesContext) IsNoiseDefinition_Builder_OctavesContext() {}

func NewNoiseDefinition_Builder_OctavesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NoiseDefinition_Builder_OctavesContext {
	var p = new(NoiseDefinition_Builder_OctavesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinecraftMetascriptParserRULE_noiseDefinition_Builder_Octaves

	return p
}

func (s *NoiseDefinition_Builder_OctavesContext) GetParser() antlr.Parser { return s.parser }

func (s *NoiseDefinition_Builder_OctavesContext) AllNumber() []INumberContext {
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

func (s *NoiseDefinition_Builder_OctavesContext) Number(i int) INumberContext {
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

func (s *NoiseDefinition_Builder_OctavesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NoiseDefinition_Builder_OctavesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NoiseDefinition_Builder_OctavesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.EnterNoiseDefinition_Builder_Octaves(s)
	}
}

func (s *NoiseDefinition_Builder_OctavesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinecraftMetascriptListener); ok {
		listenerT.ExitNoiseDefinition_Builder_Octaves(s)
	}
}

func (p *MinecraftMetascriptParser) NoiseDefinition_Builder_Octaves() (localctx INoiseDefinition_Builder_OctavesContext) {
	localctx = NewNoiseDefinition_Builder_OctavesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, MinecraftMetascriptParserRULE_noiseDefinition_Builder_Octaves)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(696)
		p.Match(MinecraftMetascriptParserT__38)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(702)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 75, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(697)
				p.Number()
			}
			{
				p.SetState(698)
				p.Match(MinecraftMetascriptParserT__16)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(704)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 75, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(705)
		p.Number()
	}
	{
		p.SetState(706)
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
