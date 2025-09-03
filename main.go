//go:build !js && !wasm

package main

import "github.com/minecraftmetascript/mms/cmd"

func main() {
	cmd.Execute()
}
