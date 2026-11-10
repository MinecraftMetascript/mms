package syntax

import (
	"strings"

	"github.com/minecraftmetascript/mms/engine/syntax/parser"
)

// keywords are the lexer's keyword literals, read from the generated lexer
// so this list cannot drift from the grammar.
var keywords = func() map[string]bool {
	out := map[string]bool{}
	for _, lit := range parser.NewMMSLexer(nil).LiteralNames {
		w := strings.Trim(lit, "'")
		if len(lit) > 2 && w != "" && isWord(w) {
			out[w] = true
		}
	}
	return out
}()

// keyWords are the keywords the parser's `word` rule still takes as field
// keys, argument names, member names and declaration names.
var keyWords = map[string]bool{"Type": true, "primary": true, "as": true, "check": true, "emit": true, "setup": true, "expect": true, "run": true, "private": true}

func isWord(w string) bool {
	for _, r := range w {
		if !(r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z') {
			return false
		}
	}
	return true
}

// Reserved reports whether w is a keyword, which can't be written bare
// where a value goes (an id, an enum name, a reference).
func Reserved(w string) bool { return keywords[w] }

// ReservedKey reports whether w can't be written bare as a field key.
func ReservedKey(w string) bool { return keywords[w] && !keyWords[w] }

// codeKinds are the declaration kinds whose body is function code.
var codeKinds = map[string]bool{"Function": true, "Test": true, "OnLoad": true, "OnTick": true, "OnFirstJoin": true, "OnRespawn": true, "OnAdvancement": true}

// IsCodeKind reports whether a declaration of the kind named k has a body
// of function code: statements and commands rather than fields. It is
// the one list the evaluator, the checker and the editor share.
func IsCodeKind(k string) bool { return codeKinds[k] }
