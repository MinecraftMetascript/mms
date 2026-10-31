package main

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/minecraftmetascript/mms/engine/sema"
	"github.com/minecraftmetascript/mms/engine/spec"
)

// runReference prints the language reference as JSON: every declaration
// kind and every helper. The docs site is built from it.
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
	if err := enc.Encode(sema.BuildReference(p, sema.NewKinds(p))); err != nil {
		fmt.Fprintln(stderr, "mms:", err)
		return 1
	}
	return 0
}
