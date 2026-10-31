//go:build js && wasm

// Command mms-wasm is the browser build of the language server. It runs in
// a Web Worker: the page posts JSON-RPC messages to globalThis.mmsSend(text)
// and receives the server's messages through globalThis.mmsReceive(text),
// which the worker script defines before starting the program.
package main

import (
	"context"
	"fmt"
	"syscall/js"

	"github.com/minecraftmetascript/mms/engine/lsp"
	"github.com/minecraftmetascript/mms/engine/spec"
)

func main() {
	pack := spec.MustDefault()
	pipe := lsp.NewPipe()
	send := js.FuncOf(func(_ js.Value, args []js.Value) any {
		if len(args) > 0 {
			pipe.Send([]byte(args[0].String()))
		}
		return nil
	})
	defer send.Release()
	global := js.Global()
	global.Set("mmsSend", send)
	receive := global.Get("mmsReceive")
	go func() {
		for {
			m, ok := pipe.Recv()
			if !ok {
				return
			}
			receive.Invoke(string(m))
		}
	}()
	if ready := global.Get("mmsReady"); ready.Type() == js.TypeFunction {
		ready.Invoke(pack.Version)
	}
	s := lsp.NewServer(pipe, pack)
	if err := s.Run(context.Background()); err != nil {
		fmt.Println("mms-wasm:", err)
	}
	pipe.Close()
}
