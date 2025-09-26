//go:build js && wasm

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"sync"

	"syscall/js"

	"github.com/minecraftmetascript/mms/lang"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
	"github.com/minecraftmetascript/mms/lsp"
)

var logger = log.Default()

type packagedProject struct {
	Source  map[string]string           `json:"source"`
	Files   *lib.FileTreeLike           `json:"files"`
	Symbols map[string]traversal.Symbol `json:"symbols"`
}

func packageProject() (string, error) {
	out := packagedProject{
		Source:  make(map[string]string),
		Files:   project.BuildFsLike("my_mms_project"),
		Symbols: project.GlobalScope.Symbols(),
	}

	for _, file := range project.Files {
		out.Source[file.Path] = file.Content
	}

	serialized, err := json.Marshal(out)
	if err != nil {
		log.Println("[Err]: Failed to package project: ", err)
		return "\"\"", err
	}
	return string(serialized), nil
}

func updateFile(this js.Value, args []js.Value) any {
	if len(args) != 3 {
		log.Println("[Err]: Invalid number of arguments, expected 3, given", len(args))
		return nil
	}
	filename := args[0].String()
	content := args[1].String()
	callback := args[2]
	if callback.Type() != js.TypeFunction {
		log.Println("[Err]: Invalid callback type, expected function, got", callback.Type())
		return nil
	}

	err := project.AddFile(filename, content).Parse()
	if err != nil {
		log.Println("[Err]: Failed to add file:", err)
		return nil
	}
	projectStruct, err := packageProject()
	if err != nil {
		log.Println("[Err]:", err)
	} else {
		callback.Invoke(projectStruct)
	}
	return nil
}

func getFileDiag(this js.Value, args []js.Value) any {
	if len(args) != 2 {
		log.Println("[Err]: Invalid number of arguments, expected 2, given", len(args))
		return nil
	}
	filename := args[0].String()
	callback := args[1]
	if callback.Type() != js.TypeFunction {
		log.Println("[Err]: Invalid callback type, expected function, got", callback.Type())
		return nil
	}
	if file, ok := project.Files[filename]; ok {
		raw, err := json.Marshal(lib.Unique(file.Diagnostics))
		if err != nil {
			log.Println("[Err]:", err)
			return nil
		}
		callback.Invoke(string(raw))
	}
	return nil
}

var project *lang.Project

func main() {
	logger.SetPrefix("[MMS:WASM]: ")
	logger.SetFlags(0)
	logger.Println("MMS WASM loading")

	project = lang.NewProject()

	js.Global().Set("updateFile", js.FuncOf(updateFile))
	logger.Println("updateFile function registered")

	js.Global().Set("getFileDiag", js.FuncOf(getFileDiag))
	logger.Println("getFileDiag function registered")

	logger.Println("MMS WASM loaded")

	// JS function used to deliver bytes to the client (Go -> JS)
	jsLspTo := js.Global().Get("mmsLspRead")
	if jsLspTo.Type() != js.TypeFunction {
		panic("mmsLspRead is not a defined function")
	}

	// Create single duplex stream instance
	stream := NewWasmStream(jsLspTo)

	// Register the function that JS will call to send data to Go (JS -> Go)
	js.Global().Set("mmsLspWrite", js.FuncOf(stream.fromJs))

	// Start LSP using the same stream for both reading and writing

	lsp.StartStreaming(stream)

	select {} // Keep Go WASM running
}

type WasmStream struct {
	// toJs is the JS function to invoke when Go writes bytes to the client.
	toJs js.Value

	// Internal incoming buffer (JS -> Go)
	mu     sync.Mutex
	cond   *sync.Cond
	buf    bytes.Buffer
	closed bool
}

// NewWasmStream constructs a new stream with a ready condition variable.
func NewWasmStream(toJS js.Value) *WasmStream {
	ws := &WasmStream{toJs: toJS}
	ws.cond = sync.NewCond(&ws.mu)
	return ws
}

// fromJs is exposed to JS as mmsLspWrite. It accepts a single string argument
// and appends it to the internal buffer for Read() to consume.
func (w *WasmStream) fromJs(this js.Value, args []js.Value) any {
	if len(args) != 1 {
		log.Println("[Err]: Invalid number of arguments, expected 1, given", len(args))
		return nil
	}
	var data []byte
	switch args[0].Type() {
	case js.TypeString:
		data = []byte(args[0].String())
	case js.TypeObject:
		// Support Uint8Array input
		uint8Array := js.Global().Get("Uint8Array")
		if uint8Array.Truthy() && args[0].InstanceOf(uint8Array) {
			length := args[0].Get("byteLength").Int()
			data = make([]byte, length)
			js.CopyBytesToGo(data, args[0])
		} else {
			log.Println("[Err]: Invalid argument type, expected string or Uint8Array")
			return nil
		}
	default:
		log.Println("[Err]: Invalid argument type, expected string or Uint8Array, got", args[0].Type())
		return nil
	}

	w.mu.Lock()
	defer w.mu.Unlock()

	if w.closed {
		// Ignore incoming data if closed
		return nil
	}

	w.buf.Write(data)
	w.cond.Signal()

	return nil
}

// Read blocks until data is available or the stream is closed.
// Returns io.EOF if closed and no more data is available.
func (w *WasmStream) Read(p []byte) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	for w.buf.Len() == 0 && !w.closed {
		w.cond.Wait()
	}
	if w.buf.Len() == 0 && w.closed {
		return 0, io.EOF
	}
	res, e := w.buf.Read(p)
	return res, e
}

// Write sends bytes to JS via the provided function.
func (w *WasmStream) Write(p []byte) (n int, err error) {
	if w.toJs.Type() != js.TypeFunction {
		return 0, fmt.Errorf("destination JS function is not defined")
	}
	w.toJs.Invoke(string(p))
	return len(p), nil
}

// Close marks the stream as closed and wakes any waiting readers.
func (w *WasmStream) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()
	if !w.closed {
		w.closed = true
		w.cond.Broadcast()
	}
	return nil
}
