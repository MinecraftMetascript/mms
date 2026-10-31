package main

import (
	"context"
	"fmt"
	"io"

	"github.com/minecraftmetascript/mms/engine/lsp"
	"github.com/minecraftmetascript/mms/engine/spec"
)

// runLSP serves the language server over stdin and stdout.
func runLSP(args []string, stdin io.Reader, stdout, stderr io.Writer) int {
	for _, a := range args {
		if a != "--stdio" {
			fmt.Fprintf(stderr, "mms lsp: unknown argument %q\n", a)
			return 2
		}
	}
	pack, err := spec.Default()
	if err != nil {
		fmt.Fprintln(stderr, "mms lsp:", err)
		return 1
	}
	s := lsp.NewServer(lsp.NewStream(stdin, stdout), pack)
	s.Log = stderr
	if err := s.Run(context.Background()); err != nil {
		fmt.Fprintln(stderr, "mms lsp:", err)
		return 1
	}
	return 0
}
