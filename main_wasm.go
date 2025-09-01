//go:build js && wasm

package main

import "log"
import "syscall/js"

func parseLiteral(this js.Value, args []js.Value) any {
	return "Hello World!"
}

func main() {
	log.Default().Println("MMS WASM loading")
	js.Global().Set("parseLiteral", js.FuncOf(parseLiteral))
	log.Default().Println("parseLiteral function registered")

	log.Default().Println("MMS WASM loaded")
	select {} // Keep Go WASM running
}
