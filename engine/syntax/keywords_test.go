package syntax

import "testing"

// TestReservedMatchesGrammar checks Reserved and ReservedKey against the
// parser: a word is reserved exactly when writing it bare fails to parse,
// or (true, false) parses as something other than a name.
func TestReservedMatchesGrammar(t *testing.T) {
	words := []string{"stone", "setup", "run", "check", "const", "extern", "let", "true", "Type", "primary", "as", "emit", "expect", "in", "return"}
	for w := range keywords {
		words = append(words, w)
	}
	for _, w := range words {
		value := len(Parse("FunctionTag t = ["+w+"]\n").Diags) > 0 || w == "true" || w == "false"
		if value != Reserved(w) {
			t.Errorf("%q: bare value fails to parse = %v, Reserved = %v", w, value, Reserved(w))
		}
		key := len(Parse("LootTable t {\n  "+w+" = 1\n}\n").Diags) > 0 || w == "true" || w == "false"
		if key != ReservedKey(w) {
			t.Errorf("%q: bare key fails to parse = %v, ReservedKey = %v", w, key, ReservedKey(w))
		}
	}
}

func TestIsCodeKind(t *testing.T) {
	for _, k := range []string{"Function", "Test", "OnLoad", "OnTick", "OnFirstJoin", "OnRespawn", "OnAdvancement"} {
		if !IsCodeKind(k) {
			t.Errorf("%s is code", k)
		}
	}
	for _, k := range []string{"LootTable", "CustomItem", "Template", ""} {
		if IsCodeKind(k) {
			t.Errorf("%s is not code", k)
		}
	}
}
