// Package jsonv is an ordered JSON value model with deterministic output.
// Objects keep insertion order, so emitted files follow schema field order
// and rebuilds are byte-identical.
package jsonv

import (
	"bytes"
	"encoding/json"
	"sort"
	"strconv"
	"strings"
)

// Value is one of *Object, Array, String, Number, Bool or Null.
type Value interface{ value() }

type (
	// Object is an ordered set of members.
	Object struct{ Members []Member }
	Member struct {
		Key   string
		Value Value
	}
	Array  []Value
	String string
	// Number keeps its text, so 0.0 stays 0.0 and large ints stay exact.
	Number string
	Bool   bool
	Null   struct{}
)

func (*Object) value() {}
func (Array) value()   {}
func (String) value()  {}
func (Number) value()  {}
func (Bool) value()    {}
func (Null) value()    {}

// Int returns a Number for n.
func Int(n int) Number { return Number(strconv.Itoa(n)) }

// Float returns a Number for f in its shortest form.
func Float(f float64) Number { return Number(strconv.FormatFloat(f, 'g', -1, 64)) }

// Get returns the member value for key, or nil.
func (o *Object) Get(key string) Value {
	for _, m := range o.Members {
		if m.Key == key {
			return m.Value
		}
	}
	return nil
}

// Set replaces key's value or appends it.
func (o *Object) Set(key string, v Value) {
	for i, m := range o.Members {
		if m.Key == key {
			o.Members[i].Value = v
			return
		}
	}
	o.Members = append(o.Members, Member{key, v})
}

// Float reports n's value.
func (n Number) Float() (float64, bool) {
	f, err := strconv.ParseFloat(string(n), 64)
	return f, err == nil
}

// Marshal writes v with two-space indent and a trailing newline.
func Marshal(v Value) []byte {
	var b bytes.Buffer
	write(&b, v, 0)
	b.WriteByte('\n')
	return b.Bytes()
}

// Compact writes v on one line.
func Compact(v Value) string {
	var b bytes.Buffer
	write(&b, v, -1)
	return b.String()
}

func write(b *bytes.Buffer, v Value, indent int) {
	nl := func(n int) {
		if indent < 0 {
			return
		}
		b.WriteByte('\n')
		b.WriteString(strings.Repeat("  ", n))
	}
	next := indent + 1
	if indent < 0 {
		next = -1
	}
	switch v := v.(type) {
	case *Object:
		if len(v.Members) == 0 {
			b.WriteString("{}")
			return
		}
		b.WriteByte('{')
		for i, m := range v.Members {
			if i > 0 {
				b.WriteByte(',')
				if indent < 0 {
					b.WriteByte(' ')
				}
			}
			nl(next)
			writeString(b, m.Key)
			b.WriteString(": ")
			write(b, m.Value, next)
		}
		nl(indent)
		b.WriteByte('}')
	case Array:
		if len(v) == 0 {
			b.WriteString("[]")
			return
		}
		b.WriteByte('[')
		for i, x := range v {
			if i > 0 {
				b.WriteByte(',')
				if indent < 0 {
					b.WriteByte(' ')
				}
			}
			nl(next)
			write(b, x, next)
		}
		nl(indent)
		b.WriteByte(']')
	case String:
		writeString(b, string(v))
	case Number:
		b.WriteString(string(v))
	case Bool:
		b.WriteString(strconv.FormatBool(bool(v)))
	case Null, nil:
		b.WriteString("null")
	}
}

func writeString(b *bytes.Buffer, s string) {
	enc, _ := json.Marshal(s)
	// encoding/json escapes <, > and & for HTML; data packs don't need it.
	enc = bytes.ReplaceAll(enc, []byte(`<`), []byte("<"))
	enc = bytes.ReplaceAll(enc, []byte(`>`), []byte(">"))
	enc = bytes.ReplaceAll(enc, []byte(`&`), []byte("&"))
	b.Write(enc)
}

// FromAny converts a decoded encoding/json value (with UseNumber) into a
// Value, sorting object keys for determinism.
func FromAny(x any) Value {
	switch x := x.(type) {
	case map[string]any:
		keys := make([]string, 0, len(x))
		for k := range x {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		o := &Object{}
		for _, k := range keys {
			o.Members = append(o.Members, Member{k, FromAny(x[k])})
		}
		return o
	case []any:
		a := make(Array, len(x))
		for i, e := range x {
			a[i] = FromAny(e)
		}
		return a
	case string:
		return String(x)
	case json.Number:
		return Number(x)
	case float64:
		return Float(x)
	case bool:
		return Bool(x)
	}
	return Null{}
}
