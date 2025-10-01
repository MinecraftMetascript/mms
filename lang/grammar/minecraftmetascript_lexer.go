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
		"", "'{'", "'}'", "'='", "':'", "'('", "','", "')'", "'.'", "'&&'",
		"'||'", "'!'", "'If'", "'['", "']'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "Int", "Float",
		"String", "WS", "NL", "Identifier", "BlockComment", "LineComment",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "Int", "Float", "String",
		"WS", "NL", "Identifier", "BlockComment", "LineComment",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 22, 151, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9,
		1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1,
		14, 3, 14, 78, 8, 14, 1, 14, 4, 14, 81, 8, 14, 11, 14, 12, 14, 82, 1, 15,
		3, 15, 86, 8, 15, 1, 15, 5, 15, 89, 8, 15, 10, 15, 12, 15, 92, 9, 15, 1,
		15, 1, 15, 4, 15, 96, 8, 15, 11, 15, 12, 15, 97, 1, 16, 1, 16, 5, 16, 102,
		8, 16, 10, 16, 12, 16, 105, 9, 16, 1, 16, 1, 16, 1, 17, 4, 17, 110, 8,
		17, 11, 17, 12, 17, 111, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19,
		1, 19, 5, 19, 122, 8, 19, 10, 19, 12, 19, 125, 9, 19, 1, 20, 1, 20, 1,
		20, 1, 20, 5, 20, 131, 8, 20, 10, 20, 12, 20, 134, 9, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 5, 21, 145, 8, 21, 10,
		21, 12, 21, 148, 9, 21, 1, 21, 1, 21, 1, 132, 0, 22, 1, 1, 3, 2, 5, 3,
		7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13,
		27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22,
		1, 0, 6, 1, 0, 48, 57, 2, 0, 10, 10, 13, 13, 2, 0, 9, 9, 32, 32, 1, 0,
		10, 10, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 47, 57, 65, 90, 95, 95, 97,
		122, 160, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1,
		0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15,
		1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0,
		23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0,
		0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0,
		0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 1, 45, 1, 0,
		0, 0, 3, 47, 1, 0, 0, 0, 5, 49, 1, 0, 0, 0, 7, 51, 1, 0, 0, 0, 9, 53, 1,
		0, 0, 0, 11, 55, 1, 0, 0, 0, 13, 57, 1, 0, 0, 0, 15, 59, 1, 0, 0, 0, 17,
		61, 1, 0, 0, 0, 19, 64, 1, 0, 0, 0, 21, 67, 1, 0, 0, 0, 23, 69, 1, 0, 0,
		0, 25, 72, 1, 0, 0, 0, 27, 74, 1, 0, 0, 0, 29, 77, 1, 0, 0, 0, 31, 85,
		1, 0, 0, 0, 33, 99, 1, 0, 0, 0, 35, 109, 1, 0, 0, 0, 37, 115, 1, 0, 0,
		0, 39, 119, 1, 0, 0, 0, 41, 126, 1, 0, 0, 0, 43, 140, 1, 0, 0, 0, 45, 46,
		5, 123, 0, 0, 46, 2, 1, 0, 0, 0, 47, 48, 5, 125, 0, 0, 48, 4, 1, 0, 0,
		0, 49, 50, 5, 61, 0, 0, 50, 6, 1, 0, 0, 0, 51, 52, 5, 58, 0, 0, 52, 8,
		1, 0, 0, 0, 53, 54, 5, 40, 0, 0, 54, 10, 1, 0, 0, 0, 55, 56, 5, 44, 0,
		0, 56, 12, 1, 0, 0, 0, 57, 58, 5, 41, 0, 0, 58, 14, 1, 0, 0, 0, 59, 60,
		5, 46, 0, 0, 60, 16, 1, 0, 0, 0, 61, 62, 5, 38, 0, 0, 62, 63, 5, 38, 0,
		0, 63, 18, 1, 0, 0, 0, 64, 65, 5, 124, 0, 0, 65, 66, 5, 124, 0, 0, 66,
		20, 1, 0, 0, 0, 67, 68, 5, 33, 0, 0, 68, 22, 1, 0, 0, 0, 69, 70, 5, 73,
		0, 0, 70, 71, 5, 102, 0, 0, 71, 24, 1, 0, 0, 0, 72, 73, 5, 91, 0, 0, 73,
		26, 1, 0, 0, 0, 74, 75, 5, 93, 0, 0, 75, 28, 1, 0, 0, 0, 76, 78, 5, 45,
		0, 0, 77, 76, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 80, 1, 0, 0, 0, 79, 81,
		7, 0, 0, 0, 80, 79, 1, 0, 0, 0, 81, 82, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0,
		82, 83, 1, 0, 0, 0, 83, 30, 1, 0, 0, 0, 84, 86, 5, 45, 0, 0, 85, 84, 1,
		0, 0, 0, 85, 86, 1, 0, 0, 0, 86, 90, 1, 0, 0, 0, 87, 89, 7, 0, 0, 0, 88,
		87, 1, 0, 0, 0, 89, 92, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 90, 91, 1, 0, 0,
		0, 91, 93, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 93, 95, 5, 46, 0, 0, 94, 96,
		7, 0, 0, 0, 95, 94, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0,
		97, 98, 1, 0, 0, 0, 98, 32, 1, 0, 0, 0, 99, 103, 5, 34, 0, 0, 100, 102,
		8, 1, 0, 0, 101, 100, 1, 0, 0, 0, 102, 105, 1, 0, 0, 0, 103, 101, 1, 0,
		0, 0, 103, 104, 1, 0, 0, 0, 104, 106, 1, 0, 0, 0, 105, 103, 1, 0, 0, 0,
		106, 107, 5, 34, 0, 0, 107, 34, 1, 0, 0, 0, 108, 110, 7, 2, 0, 0, 109,
		108, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 111, 112,
		1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113, 114, 6, 17, 0, 0, 114, 36, 1, 0,
		0, 0, 115, 116, 7, 3, 0, 0, 116, 117, 1, 0, 0, 0, 117, 118, 6, 18, 1, 0,
		118, 38, 1, 0, 0, 0, 119, 123, 7, 4, 0, 0, 120, 122, 7, 5, 0, 0, 121, 120,
		1, 0, 0, 0, 122, 125, 1, 0, 0, 0, 123, 121, 1, 0, 0, 0, 123, 124, 1, 0,
		0, 0, 124, 40, 1, 0, 0, 0, 125, 123, 1, 0, 0, 0, 126, 127, 5, 47, 0, 0,
		127, 128, 5, 42, 0, 0, 128, 132, 1, 0, 0, 0, 129, 131, 9, 0, 0, 0, 130,
		129, 1, 0, 0, 0, 131, 134, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 132, 130,
		1, 0, 0, 0, 133, 135, 1, 0, 0, 0, 134, 132, 1, 0, 0, 0, 135, 136, 5, 42,
		0, 0, 136, 137, 5, 47, 0, 0, 137, 138, 1, 0, 0, 0, 138, 139, 6, 20, 1,
		0, 139, 42, 1, 0, 0, 0, 140, 141, 5, 47, 0, 0, 141, 142, 5, 47, 0, 0, 142,
		146, 1, 0, 0, 0, 143, 145, 8, 1, 0, 0, 144, 143, 1, 0, 0, 0, 145, 148,
		1, 0, 0, 0, 146, 144, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 149, 1, 0,
		0, 0, 148, 146, 1, 0, 0, 0, 149, 150, 6, 21, 1, 0, 150, 44, 1, 0, 0, 0,
		11, 0, 77, 82, 85, 90, 97, 103, 111, 123, 132, 146, 2, 6, 0, 0, 0, 1, 0,
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
	MinecraftMetascriptLexerBlockComment = 21
	MinecraftMetascriptLexerLineComment  = 22
)
