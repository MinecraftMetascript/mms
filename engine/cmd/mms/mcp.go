package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/minecraftmetascript/mms/engine/mcp"
	"github.com/minecraftmetascript/mms/engine/spec"
)

// runMCP serves the Model Context Protocol server over stdin and stdout,
// for AI agents that write MMS. With -check it tests the server against
// itself instead and says how to set a client up.
func runMCP(args []string, stdin io.Reader, stdout, stderr io.Writer) int {
	fl := flag.NewFlagSet("mcp", flag.ContinueOnError)
	fl.SetOutput(stderr)
	check := fl.Bool("check", false, "test that the server works, print how to connect a client, and exit")
	if err := fl.Parse(args); err != nil {
		return 2
	}
	if fl.NArg() > 0 {
		fmt.Fprintf(stderr, "mms mcp: unknown argument %q\n", fl.Arg(0))
		return 2
	}
	pack, err := spec.Default()
	if err != nil {
		fmt.Fprintln(stderr, "mms mcp:", err)
		return 1
	}
	s := mcp.NewServer(pack)
	s.Version = versionString()
	if *check {
		return checkMCP(s, stdout, stderr)
	}
	s.Log = stderr
	// A client shows this in its log of the server. Someone who ran the
	// command by hand sees why nothing else happens.
	fmt.Fprintf(stderr, "mms mcp: %s is waiting for an MCP client on stdin. To test it by hand, run `mms mcp -check`.\n", versionString())
	if err := s.Run(context.Background(), stdin, stdout); err != nil {
		fmt.Fprintln(stderr, "mms mcp:", err)
		return 1
	}
	return 0
}

// checkMCP runs the server's self-test and prints the entry a client
// needs, with this binary's own path so that it works whatever PATH the
// client starts servers with.
func checkMCP(s *mcp.Server, stdout, stderr io.Writer) int {
	what, err := s.SelfTest(context.Background())
	if err != nil {
		fmt.Fprintln(stderr, "mms mcp: the server is not working (please report this):", err)
		return 1
	}
	self, err := os.Executable()
	if err != nil {
		self = "mms"
	}
	fmt.Fprintf(stdout, `ok: %s; %s

The server works. To connect a client:

  Claude Code, for you in this folder:
    claude mcp add mms -- %s mcp

  Claude Code, shared with the project, in .mcp.json at its root.
  Everyone needs this version of mms on their PATH, and Claude Code asks
  each person to approve the server the next time it starts; until
  then `+"`claude mcp list`"+` shows "Pending approval":
    %s

  Other clients take the same command, %q, with the argument "mcp".

Then `+"`claude mcp list`"+` should show "mms ... Connected", and /mcp inside
Claude Code lists its tools.
`, versionString(), what, quoteArg(self), mcp.ConfigJSON("mms"), self)
	return 0
}

// quoteArg quotes a path for a shell when it has a space.
func quoteArg(s string) string {
	for _, r := range s {
		if r == ' ' || r == '\t' {
			return `"` + s + `"`
		}
	}
	return s
}
