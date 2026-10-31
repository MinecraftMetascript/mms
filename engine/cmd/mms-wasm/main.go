//go:build js && wasm

// Command mms-wasm is the browser build of the engine. For now it only
// loads the spec pack and reports how long that took; the LSP transport
// arrives in M3.
package main

import (
	"fmt"
	"time"

	"github.com/minecraftmetascript/mms/engine/spec"
)

func main() {
	start := time.Now()
	p := spec.MustDefault()
	fmt.Printf("mms-wasm: spec pack %s loaded (%d types) in %s\n", p.Version, len(p.Types), time.Since(start).Round(time.Millisecond))
}
