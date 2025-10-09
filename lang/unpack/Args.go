package unpack

import (
	"reflect"
	"strconv"

	"github.com/minecraftmetascript/mms/lang/ast"
	"github.com/minecraftmetascript/mms/lib"
)

func getArgIdxForField(field reflect.StructField, max int) int {
	argIdxStr := field.Tag.Get("mms_arg")
	if argIdxStr == "" {
		return 0
	}
	argIdx, err := strconv.Atoi(argIdxStr)
	if err != nil {
		return 0
	}
	if argIdx > max {
		return -1
	}
	return argIdx
}

func Args[T any](v T, args []ast.Node) {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	mutableV := reflect.ValueOf(v)
	if mutableV.Kind() == reflect.Ptr {
		mutableV = mutableV.Elem()
	}
	for _, field := range reflect.VisibleFields(t) {
		// We don't have an arg index, skip
		if field.Tag.Get("mms_arg") == "" {
			continue
		}

		argIdx := getArgIdxForField(field, len(args)-1)
		if argIdx == -1 {
			continue
		}
		if argIdx < len(args) {

			arg := args[argIdx]
			mutableField := mutableV.FieldByName(field.Name)
			if !lib.IsNilInterface(mutableField) {
				assignValue(field, arg, mutableField)
			}
		}

	}
}
