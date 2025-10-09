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
		"", "'{'", "'}'", "'='", "':'", "'.'", "'('", "','", "')'", "'!'", "'&&'",
		"'||'", "'If'", "'['", "']'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "Int", "Float",
		"String", "WS", "NL", "Identifier", "DocString", "BlockComment", "LineComment",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "Int", "Float", "String",
		"WS", "NL", "Identifier", "DocString", "BlockComment", "LineComment",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 23, 180, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1,
		9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12,
		1, 13, 1, 13, 1, 14, 3, 14, 80, 8, 14, 1, 14, 4, 14, 83, 8, 14, 11, 14,
		12, 14, 84, 1, 15, 3, 15, 88, 8, 15, 1, 15, 4, 15, 91, 8, 15, 11, 15, 12,
		15, 92, 1, 15, 1, 15, 5, 15, 97, 8, 15, 10, 15, 12, 15, 100, 9, 15, 1,
		15, 5, 15, 103, 8, 15, 10, 15, 12, 15, 106, 9, 15, 1, 15, 1, 15, 4, 15,
		110, 8, 15, 11, 15, 12, 15, 111, 3, 15, 114, 8, 15, 1, 16, 1, 16, 5, 16,
		118, 8, 16, 10, 16, 12, 16, 121, 9, 16, 1, 16, 1, 16, 1, 17, 4, 17, 126,
		8, 17, 11, 17, 12, 17, 127, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		19, 1, 19, 5, 19, 138, 8, 19, 10, 19, 12, 19, 141, 9, 19, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 5, 20, 148, 8, 20, 10, 20, 12, 20, 151, 9, 20, 1,
		20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 5, 21, 160, 8, 21, 10, 21,
		12, 21, 163, 9, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1,
		22, 1, 22, 5, 22, 174, 8, 22, 10, 22, 12, 22, 177, 9, 22, 1, 22, 1, 22,
		2, 149, 161, 0, 23, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8,
		17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17,
		35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 1, 0, 6, 1, 0, 48, 57,
		2, 0, 10, 10, 13, 13, 2, 0, 9, 9, 32, 32, 1, 0, 10, 10, 3, 0, 65, 90, 95,
		95, 97, 122, 4, 0, 47, 57, 65, 90, 95, 95, 97, 122, 193, 0, 1, 1, 0, 0,
		0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0,
		0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0,
		0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1,
		0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33,
		1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0,
		41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 1, 47, 1, 0, 0, 0,
		3, 49, 1, 0, 0, 0, 5, 51, 1, 0, 0, 0, 7, 53, 1, 0, 0, 0, 9, 55, 1, 0, 0,
		0, 11, 57, 1, 0, 0, 0, 13, 59, 1, 0, 0, 0, 15, 61, 1, 0, 0, 0, 17, 63,
		1, 0, 0, 0, 19, 65, 1, 0, 0, 0, 21, 68, 1, 0, 0, 0, 23, 71, 1, 0, 0, 0,
		25, 74, 1, 0, 0, 0, 27, 76, 1, 0, 0, 0, 29, 79, 1, 0, 0, 0, 31, 113, 1,
		0, 0, 0, 33, 115, 1, 0, 0, 0, 35, 125, 1, 0, 0, 0, 37, 131, 1, 0, 0, 0,
		39, 135, 1, 0, 0, 0, 41, 142, 1, 0, 0, 0, 43, 155, 1, 0, 0, 0, 45, 169,
		1, 0, 0, 0, 47, 48, 5, 123, 0, 0, 48, 2, 1, 0, 0, 0, 49, 50, 5, 125, 0,
		0, 50, 4, 1, 0, 0, 0, 51, 52, 5, 61, 0, 0, 52, 6, 1, 0, 0, 0, 53, 54, 5,
		58, 0, 0, 54, 8, 1, 0, 0, 0, 55, 56, 5, 46, 0, 0, 56, 10, 1, 0, 0, 0, 57,
		58, 5, 40, 0, 0, 58, 12, 1, 0, 0, 0, 59, 60, 5, 44, 0, 0, 60, 14, 1, 0,
		0, 0, 61, 62, 5, 41, 0, 0, 62, 16, 1, 0, 0, 0, 63, 64, 5, 33, 0, 0, 64,
		18, 1, 0, 0, 0, 65, 66, 5, 38, 0, 0, 66, 67, 5, 38, 0, 0, 67, 20, 1, 0,
		0, 0, 68, 69, 5, 124, 0, 0, 69, 70, 5, 124, 0, 0, 70, 22, 1, 0, 0, 0, 71,
		72, 5, 73, 0, 0, 72, 73, 5, 102, 0, 0, 73, 24, 1, 0, 0, 0, 74, 75, 5, 91,
		0, 0, 75, 26, 1, 0, 0, 0, 76, 77, 5, 93, 0, 0, 77, 28, 1, 0, 0, 0, 78,
		80, 5, 45, 0, 0, 79, 78, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 82, 1, 0,
		0, 0, 81, 83, 7, 0, 0, 0, 82, 81, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 82,
		1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 30, 1, 0, 0, 0, 86, 88, 5, 45, 0, 0,
		87, 86, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 90, 1, 0, 0, 0, 89, 91, 7,
		0, 0, 0, 90, 89, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 92,
		93, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 98, 5, 46, 0, 0, 95, 97, 7, 0,
		0, 0, 96, 95, 1, 0, 0, 0, 97, 100, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 98,
		99, 1, 0, 0, 0, 99, 114, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 101, 103, 7,
		0, 0, 0, 102, 101, 1, 0, 0, 0, 103, 106, 1, 0, 0, 0, 104, 102, 1, 0, 0,
		0, 104, 105, 1, 0, 0, 0, 105, 107, 1, 0, 0, 0, 106, 104, 1, 0, 0, 0, 107,
		109, 5, 46, 0, 0, 108, 110, 7, 0, 0, 0, 109, 108, 1, 0, 0, 0, 110, 111,
		1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 114, 1, 0,
		0, 0, 113, 87, 1, 0, 0, 0, 113, 104, 1, 0, 0, 0, 114, 32, 1, 0, 0, 0, 115,
		119, 5, 34, 0, 0, 116, 118, 8, 1, 0, 0, 117, 116, 1, 0, 0, 0, 118, 121,
		1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120, 122, 1, 0,
		0, 0, 121, 119, 1, 0, 0, 0, 122, 123, 5, 34, 0, 0, 123, 34, 1, 0, 0, 0,
		124, 126, 7, 2, 0, 0, 125, 124, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127,
		125, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 130,
		6, 17, 0, 0, 130, 36, 1, 0, 0, 0, 131, 132, 7, 3, 0, 0, 132, 133, 1, 0,
		0, 0, 133, 134, 6, 18, 1, 0, 134, 38, 1, 0, 0, 0, 135, 139, 7, 4, 0, 0,
		136, 138, 7, 5, 0, 0, 137, 136, 1, 0, 0, 0, 138, 141, 1, 0, 0, 0, 139,
		137, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140, 40, 1, 0, 0, 0, 141, 139, 1,
		0, 0, 0, 142, 143, 5, 47, 0, 0, 143, 144, 5, 42, 0, 0, 144, 145, 5, 42,
		0, 0, 145, 149, 1, 0, 0, 0, 146, 148, 9, 0, 0, 0, 147, 146, 1, 0, 0, 0,
		148, 151, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 150,
		152, 1, 0, 0, 0, 151, 149, 1, 0, 0, 0, 152, 153, 5, 42, 0, 0, 153, 154,
		5, 47, 0, 0, 154, 42, 1, 0, 0, 0, 155, 156, 5, 47, 0, 0, 156, 157, 5, 42,
		0, 0, 157, 161, 1, 0, 0, 0, 158, 160, 9, 0, 0, 0, 159, 158, 1, 0, 0, 0,
		160, 163, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 161, 159, 1, 0, 0, 0, 162,
		164, 1, 0, 0, 0, 163, 161, 1, 0, 0, 0, 164, 165, 5, 42, 0, 0, 165, 166,
		5, 47, 0, 0, 166, 167, 1, 0, 0, 0, 167, 168, 6, 21, 1, 0, 168, 44, 1, 0,
		0, 0, 169, 170, 5, 47, 0, 0, 170, 171, 5, 47, 0, 0, 171, 175, 1, 0, 0,
		0, 172, 174, 8, 1, 0, 0, 173, 172, 1, 0, 0, 0, 174, 177, 1, 0, 0, 0, 175,
		173, 1, 0, 0, 0, 175, 176, 1, 0, 0, 0, 176, 178, 1, 0, 0, 0, 177, 175,
		1, 0, 0, 0, 178, 179, 6, 22, 1, 0, 179, 46, 1, 0, 0, 0, 15, 0, 79, 84,
		87, 92, 98, 104, 111, 113, 119, 127, 139, 149, 161, 175, 2, 6, 0, 0, 0,
		1, 0,
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
	MinecraftMetascriptLexerInt          = 15
	MinecraftMetascriptLexerFloat        = 16
	MinecraftMetascriptLexerString_      = 17
	MinecraftMetascriptLexerWS           = 18
	MinecraftMetascriptLexerNL           = 19
	MinecraftMetascriptLexerIdentifier   = 20
	MinecraftMetascriptLexerDocString    = 21
	MinecraftMetascriptLexerBlockComment = 22
	MinecraftMetascriptLexerLineComment  = 23
)
