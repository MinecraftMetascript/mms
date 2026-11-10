package main

import (
	"context"
	"fmt"
	"io"

	"github.com/minecraftmetascript/mms/engine/mcp"
	"github.com/minecraftmetascript/mms/engine/spec"
)

// runMCP serves the Model Context Protocol server over stdin and stdout,
// for AI agents that write MMS.
func runMCP(args []string, stdin io.Reader, stdout, stderr io.Writer) int {
	if len(args) > 0 {
		fmt.Fprintln(stderr, "usage: mms mcp")
		return 2
	}
	pack, err := spec.Default()
	if err != nil {
		fmt.Fprintln(stderr, "mms mcp:", err)
		return 1
	}
	s := mcp.NewServer(pack)
	s.Version = versionString()
	s.Log = stderr
	if err := s.Run(context.Background(), stdin, stdout); err != nil {
		fmt.Fprintln(stderr, "mms mcp:", err)
		return 1
	}
	return 0
}
