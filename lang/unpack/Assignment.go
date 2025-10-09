package unpack

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/minecraftmetascript/mms/lang/ast"
	"github.com/minecraftmetascript/mms/lang/spec"
	"github.com/minecraftmetascript/mms/lib"
)

func assignFloat(arg ast.Node, mutableField reflect.Value) {
	if val := spec.GetNumberNodeValue(arg); val != nil {
		mutableField.Set(reflect.ValueOf(*val))
	}
}
func assignString(arg ast.Node, mutableField reflect.Value) {
	if val := spec.GetStringNodeValue(arg); val != nil {
		mutableField.Set(reflect.ValueOf(*val))
	}
}
func assignBool(_ ast.Node, mutableField reflect.Value) {
	// TODO: spec.GetBooleanNodeValue impl
	mutableField.SetBool(true)
}

//func assignSymbol(arg ast.Node, mutableField reflect.Value, kind ast.SymbolKind) {
//	if val := spec.GetReferenceNodeValue(arg, kind); val != nil {
//		mutableField.Set(reflect.ValueOf(*val))
//	}
//}

func assignValue(field reflect.StructField, arg ast.Node, mutableField reflect.Value) {
	switch field.Type.Kind() {
	case reflect.Float64:
		assignFloat(arg, mutableField)
	case reflect.Bool:
		assignBool(arg, mutableField)
	case reflect.String:
		assignString(arg, mutableField)
	default:
		mmsTypeRaw := field.Tag.Get("mms_type")

		if mmsTypeRaw == "" {
			break
		}
		mmsTypeCandidates := strings.Split(mmsTypeRaw, "|")

		for _, candidate := range mmsTypeCandidates {
			mmsTypeParts := strings.Split(candidate, ",")
			mmsType := mmsTypeParts[0]
			switch mmsType {
			case "float":
				assignFloat(arg, mutableField)
			case "string":
				// TODO: Check for symbol type
				assignString(arg, mutableField)
			case "symbol":
				if s, ok := arg.(ast.Symbol); ok {
					serialized := s.ToSerializable()
					if !lib.IsNilInterface(serialized) {
						mutableField.Set(reflect.ValueOf(serialized))
					} else {
						if fn, ok := s.(*spec.FunctionNode); ok {
							if field.Tag.Get("mms_arg") == "" {
								break
							}
							argIdxStr := field.Tag.Get("mms_arg")
							argIdx, err := strconv.Atoi(argIdxStr)
							if err != nil {
								break
							}
							if argIdx > len(fn.Arguments)-1 {
								break
							}

						}
					}
				}
			}
		}

		break
	}
}
