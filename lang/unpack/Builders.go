package unpack

import (
	"reflect"

	"github.com/minecraftmetascript/mms/lang/spec"
	"github.com/samber/lo"
)

func Builders[T any](v T, builders []spec.FunctionNode) {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	mutableV := reflect.ValueOf(v)
	if mutableV.Kind() == reflect.Ptr {
		mutableV = mutableV.Elem()
	}
	for _, field := range reflect.VisibleFields(t) {
		builderKind := field.Tag.Get("mms_builder")
		if builderKind == "" {
			continue
		}
		mutableField := mutableV.FieldByName(field.Name)
		builder, ok := lo.Find(builders, func(item spec.FunctionNode) bool {
			return item.Name == builderKind
		})
		if !ok {
			continue
		}
		argIdx := getArgIdxForField(field, len(builder.Arguments)-1)
		if argIdx == -1 {
			continue
		}

		if argIdx < len(builder.Arguments) {
			assignValue(
				field, builder.Arguments[argIdx], mutableField,
			)

		}
	}

}
