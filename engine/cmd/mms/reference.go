package main

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/minecraftmetascript/mms/engine/prelude"
	"github.com/minecraftmetascript/mms/engine/project"
	"github.com/minecraftmetascript/mms/engine/sema"
	"github.com/minecraftmetascript/mms/engine/spec"
)

// reference is what `mms reference` prints: the language reference
// from sema, plus the prelude and the lint rules.
type reference struct {
	sema.Reference
	// Prelude is the Templates, Blueprints and Types every project
	// can use.
	Prelude []prelude.RefDecl `json:"prelude"`
	Rules   []refRule         `json:"rules"`
}

// refRule is one lint rule, with its default severity.
type refRule struct {
	ID       string `json:"id"`
	Severity string `json:"severity"`
	Doc      string `json:"doc,omitempty"`
}

// runReference prints the language reference as JSON: every declaration
// kind, every helper, the prelude's Templates, Blueprints and Types, and
// every lint rule. The docs site is built from it.
func runReference(args []string, stdout, stderr io.Writer) int {
	if len(args) > 0 {
		fmt.Fprintln(stderr, "usage: mms reference")
		return 2
	}
	p, err := spec.Default()
	if err != nil {
		fmt.Fprintln(stderr, "mms:", err)
		return 1
	}
	ref := reference{Reference: sema.BuildReference(p, sema.NewKinds(p)), Prelude: prelude.Reference()}
	for _, r := range project.Rules {
		ref.Rules = append(ref.Rules, refRule{ID: r.ID, Severity: project.SeverityName(r.Severity), Doc: r.Doc})
	}
	enc := json.NewEncoder(stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(ref); err != nil {
		fmt.Fprintln(stderr, "mms:", err)
		return 1
	}
	return 0
}
