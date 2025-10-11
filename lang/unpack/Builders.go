package unpack

import (
	"reflect"
	"strconv"
	"strings"

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
		builderTagValue := field.Tag.Get("mms_builder")
		if builderTagValue == "" {
			continue
		}

		builderTagParts := strings.Split(builderTagValue, ",")
		builderKind := builderTagParts[0]

		mutableField := mutableV.FieldByName(field.Name)
		builder, ok := lo.Find(builders, func(item spec.FunctionNode) bool {
			return item.Name == builderKind
		})
		if !ok {
			continue
		}
		argIdx := 0
		if len(builderTagParts) > 1 {
			argIdxStr := builderTagParts[1]

			if i, err := strconv.Atoi(argIdxStr); err == nil {
				if i >= len(builder.Arguments) {
					argIdx = -1
				} else {
					argIdx = i
				}
			}
		}

		if argIdx == -1 {
			continue
		}

		if argIdx == 0 && len(builder.Arguments) == 0 {
			// Presence of a builder without arguments (e.g., flag builders)
			assignValue(
				field, nil, mutableField,
			)
		}

		if argIdx < len(builder.Arguments) {
			assignValue(
				field, builder.Arguments[argIdx], mutableField,
			)

		}
	}

}
