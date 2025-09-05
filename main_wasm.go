//go:build js && wasm

package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/minecraftmetascript/mms/lang"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
)
import "syscall/js"

var logger = log.Default()

type packagedProject struct {
	Source  map[string]string           `json:"source"`
	Files   *lib.FileTreeLike           `json:"files"`
	Symbols map[string]traversal.Symbol `json:"symbols"`
}

func packageProject() (string, error) {
	out := packagedProject{
		Source:  make(map[string]string),
		Files:   project.BuildFsLike("wasm"),
		Symbols: project.GlobalScope.Symbols(),
	}

	for _, file := range project.Files {
		out.Source[file.Path] = file.Content
	}

	serialized, err := json.Marshal(out)
	if err != nil {
		fmt.Println("error packaging project: ", err)
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
		fmt.Println("[Err]: Invalid callback type, expected function, got", callback.Type())
		return nil
	}

	err := project.AddFile(filename, content).Parse()
	if err != nil {
		log.Println("[Err]: Failed to add file:", err)
		return nil
	}
	projectStruct, err := packageProject()
	if err != nil {
		fmt.Println("[Err]:", err)
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
		fmt.Println("[Err]: Invalid callback type, expected function, got", callback.Type())
		return nil
	}
	if file, ok := project.Files[filename]; ok {
		raw, err := json.Marshal(lib.Unique(file.Diagnostics))
		if err != nil {
			fmt.Println("[Err]:", err)
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
	select {} // Keep Go WASM running
}
