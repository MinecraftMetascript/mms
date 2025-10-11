// Code generated from ./grammar/MinecraftMetascript.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type MinecraftMetascriptLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var MinecraftMetascriptLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func minecraftmetascriptlexerLexerInit() {
	staticData := &MinecraftMetascriptLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'{'", "'}'", "'='", "':'", "'#'", "'.'", "'('", "','", "')'", "'!'",
		"'&&'", "'||'", "'If'", "'['", "']'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "Int",
		"Float", "String", "WS", "NL", "Identifier", "DocString", "BlockComment",
		"LineComment",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "Int", "Float",
		"String", "WS", "NL", "Identifier", "DocString", "BlockComment", "LineComment",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 24, 184, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1,
		8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1,
		12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 3, 15, 84, 8, 15, 1, 15,
		4, 15, 87, 8, 15, 11, 15, 12, 15, 88, 1, 16, 3, 16, 92, 8, 16, 1, 16, 4,
		16, 95, 8, 16, 11, 16, 12, 16, 96, 1, 16, 1, 16, 5, 16, 101, 8, 16, 10,
		16, 12, 16, 104, 9, 16, 1, 16, 5, 16, 107, 8, 16, 10, 16, 12, 16, 110,
		9, 16, 1, 16, 1, 16, 4, 16, 114, 8, 16, 11, 16, 12, 16, 115, 3, 16, 118,
		8, 16, 1, 17, 1, 17, 5, 17, 122, 8, 17, 10, 17, 12, 17, 125, 9, 17, 1,
		17, 1, 17, 1, 18, 4, 18, 130, 8, 18, 11, 18, 12, 18, 131, 1, 18, 1, 18,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 5, 20, 142, 8, 20, 10, 20, 12,
		20, 145, 9, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 5, 21, 152, 8, 21, 10,
		21, 12, 21, 155, 9, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22,
		5, 22, 164, 8, 22, 10, 22, 12, 22, 167, 9, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 5, 23, 178, 8, 23, 10, 23, 12, 23,
		181, 9, 23, 1, 23, 1, 23, 2, 153, 165, 0, 24, 1, 1, 3, 2, 5, 3, 7, 4, 9,
		5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14,
		29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23,
		47, 24, 1, 0, 6, 1, 0, 48, 57, 2, 0, 10, 10, 13, 13, 2, 0, 9, 9, 32, 32,
		1, 0, 10, 10, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 47, 57, 65, 90, 95,
		95, 97, 122, 197, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0,
		0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0,
		0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0,
		0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1,
		0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37,
		1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0,
		45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 1, 49, 1, 0, 0, 0, 3, 51, 1, 0, 0, 0,
		5, 53, 1, 0, 0, 0, 7, 55, 1, 0, 0, 0, 9, 57, 1, 0, 0, 0, 11, 59, 1, 0,
		0, 0, 13, 61, 1, 0, 0, 0, 15, 63, 1, 0, 0, 0, 17, 65, 1, 0, 0, 0, 19, 67,
		1, 0, 0, 0, 21, 69, 1, 0, 0, 0, 23, 72, 1, 0, 0, 0, 25, 75, 1, 0, 0, 0,
		27, 78, 1, 0, 0, 0, 29, 80, 1, 0, 0, 0, 31, 83, 1, 0, 0, 0, 33, 117, 1,
		0, 0, 0, 35, 119, 1, 0, 0, 0, 37, 129, 1, 0, 0, 0, 39, 135, 1, 0, 0, 0,
		41, 139, 1, 0, 0, 0, 43, 146, 1, 0, 0, 0, 45, 159, 1, 0, 0, 0, 47, 173,
		1, 0, 0, 0, 49, 50, 5, 123, 0, 0, 50, 2, 1, 0, 0, 0, 51, 52, 5, 125, 0,
		0, 52, 4, 1, 0, 0, 0, 53, 54, 5, 61, 0, 0, 54, 6, 1, 0, 0, 0, 55, 56, 5,
		58, 0, 0, 56, 8, 1, 0, 0, 0, 57, 58, 5, 35, 0, 0, 58, 10, 1, 0, 0, 0, 59,
		60, 5, 46, 0, 0, 60, 12, 1, 0, 0, 0, 61, 62, 5, 40, 0, 0, 62, 14, 1, 0,
		0, 0, 63, 64, 5, 44, 0, 0, 64, 16, 1, 0, 0, 0, 65, 66, 5, 41, 0, 0, 66,
		18, 1, 0, 0, 0, 67, 68, 5, 33, 0, 0, 68, 20, 1, 0, 0, 0, 69, 70, 5, 38,
		0, 0, 70, 71, 5, 38, 0, 0, 71, 22, 1, 0, 0, 0, 72, 73, 5, 124, 0, 0, 73,
		74, 5, 124, 0, 0, 74, 24, 1, 0, 0, 0, 75, 76, 5, 73, 0, 0, 76, 77, 5, 102,
		0, 0, 77, 26, 1, 0, 0, 0, 78, 79, 5, 91, 0, 0, 79, 28, 1, 0, 0, 0, 80,
		81, 5, 93, 0, 0, 81, 30, 1, 0, 0, 0, 82, 84, 5, 45, 0, 0, 83, 82, 1, 0,
		0, 0, 83, 84, 1, 0, 0, 0, 84, 86, 1, 0, 0, 0, 85, 87, 7, 0, 0, 0, 86, 85,
		1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0,
		89, 32, 1, 0, 0, 0, 90, 92, 5, 45, 0, 0, 91, 90, 1, 0, 0, 0, 91, 92, 1,
		0, 0, 0, 92, 94, 1, 0, 0, 0, 93, 95, 7, 0, 0, 0, 94, 93, 1, 0, 0, 0, 95,
		96, 1, 0, 0, 0, 96, 94, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 98, 1, 0, 0,
		0, 98, 102, 5, 46, 0, 0, 99, 101, 7, 0, 0, 0, 100, 99, 1, 0, 0, 0, 101,
		104, 1, 0, 0, 0, 102, 100, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 118,
		1, 0, 0, 0, 104, 102, 1, 0, 0, 0, 105, 107, 7, 0, 0, 0, 106, 105, 1, 0,
		0, 0, 107, 110, 1, 0, 0, 0, 108, 106, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0,
		109, 111, 1, 0, 0, 0, 110, 108, 1, 0, 0, 0, 111, 113, 5, 46, 0, 0, 112,
		114, 7, 0, 0, 0, 113, 112, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 113,
		1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116, 118, 1, 0, 0, 0, 117, 91, 1, 0,
		0, 0, 117, 108, 1, 0, 0, 0, 118, 34, 1, 0, 0, 0, 119, 123, 5, 34, 0, 0,
		120, 122, 8, 1, 0, 0, 121, 120, 1, 0, 0, 0, 122, 125, 1, 0, 0, 0, 123,
		121, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 126, 1, 0, 0, 0, 125, 123,
		1, 0, 0, 0, 126, 127, 5, 34, 0, 0, 127, 36, 1, 0, 0, 0, 128, 130, 7, 2,
		0, 0, 129, 128, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131, 129, 1, 0, 0, 0,
		131, 132, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 134, 6, 18, 0, 0, 134,
		38, 1, 0, 0, 0, 135, 136, 7, 3, 0, 0, 136, 137, 1, 0, 0, 0, 137, 138, 6,
		19, 1, 0, 138, 40, 1, 0, 0, 0, 139, 143, 7, 4, 0, 0, 140, 142, 7, 5, 0,
		0, 141, 140, 1, 0, 0, 0, 142, 145, 1, 0, 0, 0, 143, 141, 1, 0, 0, 0, 143,
		144, 1, 0, 0, 0, 144, 42, 1, 0, 0, 0, 145, 143, 1, 0, 0, 0, 146, 147, 5,
		47, 0, 0, 147, 148, 5, 42, 0, 0, 148, 149, 5, 42, 0, 0, 149, 153, 1, 0,
		0, 0, 150, 152, 9, 0, 0, 0, 151, 150, 1, 0, 0, 0, 152, 155, 1, 0, 0, 0,
		153, 154, 1, 0, 0, 0, 153, 151, 1, 0, 0, 0, 154, 156, 1, 0, 0, 0, 155,
		153, 1, 0, 0, 0, 156, 157, 5, 42, 0, 0, 157, 158, 5, 47, 0, 0, 158, 44,
		1, 0, 0, 0, 159, 160, 5, 47, 0, 0, 160, 161, 5, 42, 0, 0, 161, 165, 1,
		0, 0, 0, 162, 164, 9, 0, 0, 0, 163, 162, 1, 0, 0, 0, 164, 167, 1, 0, 0,
		0, 165, 166, 1, 0, 0, 0, 165, 163, 1, 0, 0, 0, 166, 168, 1, 0, 0, 0, 167,
		165, 1, 0, 0, 0, 168, 169, 5, 42, 0, 0, 169, 170, 5, 47, 0, 0, 170, 171,
		1, 0, 0, 0, 171, 172, 6, 22, 1, 0, 172, 46, 1, 0, 0, 0, 173, 174, 5, 47,
		0, 0, 174, 175, 5, 47, 0, 0, 175, 179, 1, 0, 0, 0, 176, 178, 8, 1, 0, 0,
		177, 176, 1, 0, 0, 0, 178, 181, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0, 179,
		180, 1, 0, 0, 0, 180, 182, 1, 0, 0, 0, 181, 179, 1, 0, 0, 0, 182, 183,
		6, 23, 1, 0, 183, 48, 1, 0, 0, 0, 15, 0, 83, 88, 91, 96, 102, 108, 115,
		117, 123, 131, 143, 153, 165, 179, 2, 6, 0, 0, 0, 1, 0,
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

// MinecraftMetascriptLexerInit initializes any static state used to implement MinecraftMetascriptLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewMinecraftMetascriptLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func MinecraftMetascriptLexerInit() {
	staticData := &MinecraftMetascriptLexerLexerStaticData
	staticData.once.Do(minecraftmetascriptlexerLexerInit)
}

// NewMinecraftMetascriptLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewMinecraftMetascriptLexer(input antlr.CharStream) *MinecraftMetascriptLexer {
	MinecraftMetascriptLexerInit()
	l := new(MinecraftMetascriptLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &MinecraftMetascriptLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "MinecraftMetascript.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// MinecraftMetascriptLexer tokens.
const (
	MinecraftMetascriptLexerT__0         = 1
	MinecraftMetascriptLexerT__1         = 2
	MinecraftMetascriptLexerT__2         = 3
	MinecraftMetascriptLexerT__3         = 4
	MinecraftMetascriptLexerT__4         = 5
	MinecraftMetascriptLexerT__5         = 6
	MinecraftMetascriptLexerT__6         = 7
	MinecraftMetascriptLexerT__7         = 8
	MinecraftMetascriptLexerT__8         = 9
	MinecraftMetascriptLexerT__9         = 10
	MinecraftMetascriptLexerT__10        = 11
	MinecraftMetascriptLexerT__11        = 12
	MinecraftMetascriptLexerT__12        = 13
	MinecraftMetascriptLexerT__13        = 14
	MinecraftMetascriptLexerT__14        = 15
	MinecraftMetascriptLexerInt          = 16
	MinecraftMetascriptLexerFloat        = 17
	MinecraftMetascriptLexerString_      = 18
	MinecraftMetascriptLexerWS           = 19
	MinecraftMetascriptLexerNL           = 20
	MinecraftMetascriptLexerIdentifier   = 21
	MinecraftMetascriptLexerDocString    = 22
	MinecraftMetascriptLexerBlockComment = 23
	MinecraftMetascriptLexerLineComment  = 24
)
