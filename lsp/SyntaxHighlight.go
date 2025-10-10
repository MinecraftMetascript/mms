package lsp

import (
	"log"
	"strings"
	"unicode/utf8"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/tliron/glsp"

	// antlr v4 Go runtime for token APIs
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// --- Legend mapping (indices must match your advertised legend) ---
var tokenTypeMapping = map[string]uint32{
	// Core identifiers and literals from your ANTLR lexer:
	"Identifier":   1, // variable
	"Int":          3, // number
	"Float":        3, // number
	"String":       4, // string
	"LineComment":  5, // comment
	"BlockComment": 5, // comment
	"DocString":    5, // comment

	// If your grammar names specific keywords by symbolic name, map them to 0:
	// "IF": 0, "ELSE": 0, "RETURN": 0, ... // keyword

	// If operators are named as tokens, map to 6:
	// "PLUS": 6, "MINUS": 6, "STAR": 6, "SLASH": 6, ...
}

// --- UTF-16 helpers ---

// utf16Len returns the number of UTF-16 code units in s.
func utf16Len(s string) int {
	// Count runes, then count how many code units they'd produce
	// (BMP = 1, surrogate pair = 2).
	n := 0
	for _, r := range s {
		if r <= 0xFFFF {
			n += 1
		} else {
			n += 2
		}
	}
	return n
}

// utf16Column computes the UTF-16 column (0-based) given the full source,
// the token start byte offset, and the start-of-line byte offset.
func utf16Column(src string, tokStart, lineStart int) int {
	// Column is the count of UTF-16 units from lineStart to tokStart.
	// Bounds checks to be safe.
	if tokStart < lineStart {
		tokStart = lineStart
	}
	if tokStart > len(src) {
		tokStart = len(src)
	}
	segment := src[lineStart:tokStart]
	return utf16Len(segment)
}

// buildLineStarts returns byte offsets for the start of each line.
// Line 0 starts at 0. For N lines, returns length N (no sentinel).
func buildLineStarts(src string) []int {
	starts := []int{0}
	// We treat \n as line separator. If your files may have \r\n, this still works.
	for i := 0; i < len(src); {
		r, size := utf8.DecodeRuneInString(src[i:])
		if r == '\n' {
			starts = append(starts, i+size)
		}
		i += size
	}
	return starts
}

// byteOffsetFromTokenStart attempts to get the byte offset where a token starts.
// If you already track positions elsewhere, use that instead.
func byteOffsetFromTokenStart(t antlr.Token) int {
	// ANTLR GetStart() returns a character index into the input stream.
	// In Go runtime this is a byte index for UTF-8 input.
	// If your input stream is not plain Go string, adapt accordingly.
	return t.GetStart()
}

// --- Semantic Tokens (Full) ---

func (s *LanguageServer) TextDocumentSemanticTokensFull(_ *glsp.Context, params *protocol.SemanticTokensParams) (*protocol.SemanticTokens, error) {
	uri := params.TextDocument.URI
	file := s.project.File(uri)
	if file == nil {
		return &protocol.SemanticTokens{Data: []uint32{}}, nil
	}

	// Get full document text (adjust to your file type)
	src := file.Content() // <-- replace with your source accessor if different

	// Token stream should already exist from your parse/lex phase
	tokenStream := file.Tokens()
	if tokenStream == nil || tokenStream.Size() == 0 {
		return &protocol.SemanticTokens{Data: []uint32{}}, nil
	}

	// Precompute line start byte offsets to convert columns to UTF-16
	lineStarts := buildLineStarts(src)

	data := make([]uint32, 0, tokenStream.Size()*5)

	prevLine := 0
	prevChar := 0 // UTF-16 units

	// Your legend of modifiers is empty in the snippets, so always 0.
	const tokenModifiers uint32 = 0

	// Walk tokens in lexical order; Helix expects sorted + delta-encoded
	for i := 0; i < tokenStream.Size(); i++ {
		t := tokenStream.Get(i)
		if t == nil {
			continue
		}

		// Skip hidden channel tokens unless you want comments colored (we do)
		// If your lexer marks comments as a different channel, keep them.
		// Otherwise, you can conditionally skip whitespace here:
		if t.GetTokenType() == antlr.TokenEOF {
			continue
		}

		// Map token type
		symName := grammar.MinecraftMetascriptLexerLexerStaticData.SymbolicNames[t.GetTokenType()] // implement to return ANTLR symbolic name
		tokenType, ok := tokenTypeMapping[symName]
		if !ok {
			// Fall back: keywords often aren’t identifiers. If your grammar
			// lowercases keyword literals, detect via a keyword set here.
			// Otherwise, skip unmapped tokens to reduce noise:
			continue
		}

		// Compute line (0-based) from ANTLR token.GetLine() (1-based)
		line := t.GetLine() - 1
		if line < 0 {
			line = 0
		}
		if line >= len(lineStarts) {
			// In case of mismatch, clamp to last line.
			line = len(lineStarts) - 1
			if line < 0 {
				line = 0
			}
		}

		// Byte offsets
		startByte := byteOffsetFromTokenStart(t)
		lineStartByte := lineStarts[line]

		// UTF-16 column
		charUTF16 := utf16Column(src, startByte, lineStartByte)

		// UTF-16 length of token text
		tokText := t.GetText()
		if tokText == "" {
			// Reconstruct token text if ANTLR text is unavailable
			// (rare; usually GetText() works)
			// Attempt to slice from input using start/stop indices.
			// Guard against invalid indices.
			start := t.GetStart()
			stop := t.GetStop()
			if start >= 0 && stop >= start && stop < len(src) {
				tokText = src[start : stop+1]
			}
		}
		lengthUTF16 := utf16Len(tokText)
		if lengthUTF16 == 0 {
			continue
		}

		// Delta encode (UTF-16 aware)
		var deltaLine, deltaChar int
		if i == 0 {
			deltaLine = line
			deltaChar = charUTF16
		} else {
			if line == prevLine {
				deltaLine = 0
				deltaChar = charUTF16 - prevChar
				if deltaChar < 0 {
					// Out-of-order or miscomputed column; skip to be safe
					continue
				}
			} else {
				deltaLine = line - prevLine
				deltaChar = charUTF16
				if deltaLine < 0 {
					// Out-of-order token; skip
					continue
				}
			}
		}

		log.Println(uint32(deltaLine),
			uint32(deltaChar),
			uint32(lengthUTF16),
			uint32(tokenType),
			tokenModifiers)
		data = append(data,
			uint32(deltaLine),
			uint32(deltaChar),
			uint32(lengthUTF16),
			uint32(tokenType),
			tokenModifiers,
		)

		prevLine = line
		prevChar = charUTF16
	}

	return &protocol.SemanticTokens{Data: data}, nil
}

// --- Optional: simple keyword detector if your grammar doesn't name them ---
// If your grammar uses literal tokens for keywords and they come through as
// 'Identifier' text, you can map them like this. Call inside the loop to
// override tokenType when needed.

var keywordSet = map[string]struct{}{
	"if": {}, "else": {}, "return": {}, "for": {}, "while": {},
	// add your language keywords...
}

func isKeyword(text string) bool {
	_, ok := keywordSet[strings.ToLower(text)]
	return ok
}
