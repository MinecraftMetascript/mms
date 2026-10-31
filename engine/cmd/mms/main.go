// Command mms is the MMS command-line tool.
package main

import (
	"fmt"
	"os"
)

const usage = `usage: mms <command> [arguments]

commands:
  fmt     format .mms files
`

func main() {
	if len(os.Args) < 2 {
		fmt.Fprint(os.Stderr, usage)
		os.Exit(2)
	}
	var code int
	switch os.Args[1] {
	case "fmt":
		code = runFmt(os.Args[2:], os.Stdin, os.Stdout, os.Stderr)
	case "help", "-h", "--help":
		fmt.Print(usage)
	default:
		fmt.Fprintf(os.Stderr, "mms: unknown command %q\n\n%s", os.Args[1], usage)
		code = 2
	}
	os.Exit(code)
}
