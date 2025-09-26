package lib

import "reflect"

func DerefType(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Ptr {
		return t.Elem()
	}
	return t
}

func Ptr[T any](t T) *T {
	return &t
}
