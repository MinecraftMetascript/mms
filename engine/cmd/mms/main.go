// Command mms is the MMS command-line tool.
package main

import (
	"fmt"
	"os"
)

const usage = `usage: mms <command> [arguments]

commands:
  build   check the project and write the data pack
  check   check the project and report problems
  fmt     format .mms files
  import  turn a data pack folder or zip into an MMS project
  init    create mms.json and a starter file
  lsp     run the language server over stdio
`

func main() {
	if len(os.Args) < 2 {
		fmt.Fprint(os.Stderr, usage)
		os.Exit(2)
	}
	var code int
	switch os.Args[1] {
	case "build":
		code = runBuild(os.Args[2:], os.Stdout, os.Stderr)
	case "check":
		code = runCheck(os.Args[2:], os.Stdout, os.Stderr)
	case "init":
		code = runInit(os.Args[2:], os.Stdout, os.Stderr)
	case "fmt":
		code = runFmt(os.Args[2:], os.Stdin, os.Stdout, os.Stderr)
	case "import":
		code = runImport(os.Args[2:], os.Stdout, os.Stderr)
	case "lsp":
		code = runLSP(os.Args[2:], os.Stdin, os.Stdout, os.Stderr)
	case "help", "-h", "--help":
		fmt.Print(usage)
	default:
		fmt.Fprintf(os.Stderr, "mms: unknown command %q\n\n%s", os.Args[1], usage)
		code = 2
	}
	os.Exit(code)
}
