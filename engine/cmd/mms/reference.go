package main

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/minecraftmetascript/mms/engine/prelude"
	"github.com/minecraftmetascript/mms/engine/sema"
	"github.com/minecraftmetascript/mms/engine/spec"
)

// runReference prints the language reference as JSON: every declaration
// kind, every helper, and the prelude's Templates, Blueprints and Types. The docs site is built from it.
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
	enc := json.NewEncoder(stdout)
	enc.SetIndent("", "  ")
	ref := struct {
		sema.Reference
		// Prelude is the Templates, Blueprints and Types every project
		// can use.
		Prelude []prelude.RefDecl `json:"prelude"`
	}{sema.BuildReference(p, sema.NewKinds(p)), prelude.Reference()}
	if err := enc.Encode(ref); err != nil {
		fmt.Fprintln(stderr, "mms:", err)
		return 1
	}
	return 0
}
