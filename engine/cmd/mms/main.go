// Command mms is the MMS command-line tool.
package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/minecraftmetascript/mms/engine/spec"
)

// version is set by release builds: -ldflags "-X main.version=v1.0.0".
var version = ""

func versionString() string {
	v := version
	if v == "" {
		v = "dev"
		if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "" && info.Main.Version != "(devel)" {
			v = info.Main.Version
		}
	}
	mc := "?"
	if p, err := spec.Default(); err == nil {
		mc = p.Version
	}
	return fmt.Sprintf("mms %s (Minecraft %s)", v, mc)
}

const usage = `usage: mms <command> [arguments]

commands:
  build   check the project and write the data pack
  check   check the project and report problems
  fix     apply every safe lint fix
  fmt     format .mms files
  import  turn a data pack folder or zip into an MMS project
  init    create mms.json and a starter file
  lsp     run the language server over stdio
  version print the version
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
	case "fix":
		code = runFix(os.Args[2:], os.Stdout, os.Stderr)
	case "fmt":
		code = runFmt(os.Args[2:], os.Stdin, os.Stdout, os.Stderr)
	case "import":
		code = runImport(os.Args[2:], os.Stdout, os.Stderr)
	case "lsp":
		code = runLSP(os.Args[2:], os.Stdin, os.Stdout, os.Stderr)
	case "version", "-v", "--version":
		fmt.Println(versionString())
	case "help", "-h", "--help":
		fmt.Print(usage)
	default:
		fmt.Fprintf(os.Stderr, "mms: unknown command %q\n\n%s", os.Args[1], usage)
		code = 2
	}
	os.Exit(code)
}
