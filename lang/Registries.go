package lang

import (
	"github.com/minecraftmetascript/mms/lang/ast"
	"github.com/minecraftmetascript/mms/lang/spec"
	"github.com/minecraftmetascript/mms/lib"
)

var Blocks = spec.NewBlockSpecList()

// JitSymbols is a map of all the symbols present in a project that will be populated immediately prior to serialization
var JitSymbols = map[string]*ast.Namespace{}

func getSymbol(namespace, name string) ast.Symbol {
	if symbolNs, ok := JitSymbols[namespace]; ok {
		return symbolNs.GetDecl(name)
	}
	return nil
}

func getSymbolValue(namespace, name string) any {
	if symbol := getSymbol(namespace, name); !lib.IsNilInterface(symbol) {
		return symbol.ToSerializable()
	}
	return nil
}
