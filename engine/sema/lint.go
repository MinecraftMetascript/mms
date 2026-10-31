package sema

import (
	"regexp"
	"strings"

	"github.com/minecraftmetascript/mms/engine/syntax"
)

// This file is the checker's API for lint rules.

var (
	bareIdent  = regexp.MustCompile(`^[a-z_][a-z0-9_]*$`)
	barePath   = regexp.MustCompile(`^[a-z_][a-z0-9_]*(/[a-z0-9_./-]+)+$`)
	bareResLoc = regexp.MustCompile(`^[a-z0-9_.-]+:[a-z0-9_./-]+$`)
	reserved   = map[string]bool{"let": true, "if": true, "else": true, "while": true, "for": true, "in": true,
		"return": true, "break": true, "continue": true, "true": true, "false": true}
)

// BareString returns how a string literal could be written without
// quotes and mean the same, using what the checker recorded about its
// position: an id (`minecraft:stone` becomes `stone` when no declaration
// in ns would capture the short name), a dispatch key or a string enum
// name. ok is false when the quotes are needed.
func BareString(in *Info, syms *Table, ns string, s *syntax.String) (string, bool) {
	if in == nil || s.Value == "" {
		return "", false
	}
	switch in.Bare {
	case BareKey, BareEnum:
		v := strings.TrimPrefix(s.Value, "minecraft:")
		return v, bareIdent.MatchString(v) && !reserved[v]
	case BareID:
	default:
		return "", false
	}
	text := s.Value
	tag := strings.HasPrefix(text, "#")
	if tag {
		text = text[1:]
	}
	// A quoted id resolves exactly like the bare one, so only the
	// spelling matters; the minecraft: prefix is dropped when the short
	// name cannot mean a declaration instead.
	if path, ok := strings.CutPrefix(text, "minecraft:"); ok && in.Folder != "" && syms.Find(in.Folder, ns+":"+path) == nil && lexesBare(path) {
		text = path
	}
	if !lexesBare(text) {
		return "", false
	}
	if tag {
		return "#" + text, true
	}
	return text, true
}

// ShortName returns name (an id, dispatch key or enum name written bare,
// without '#') with its minecraft: prefix dropped, when the short form
// means the same at this position.
func ShortName(in *Info, syms *Table, ns, name string) (string, bool) {
	path, ok := strings.CutPrefix(name, "minecraft:")
	if in == nil || !ok || !lexesBare(path) {
		return "", false
	}
	switch in.Bare {
	case BareKey, BareEnum:
		return path, bareIdent.MatchString(path) && !reserved[path]
	case BareID:
		return path, in.Folder != "" && syms.Find(in.Folder, ns+":"+path) == nil
	}
	return "", false
}

func lexesBare(s string) bool {
	return (bareIdent.MatchString(s) && !reserved[s]) || barePath.MatchString(s) || bareResLoc.MatchString(s)
}
