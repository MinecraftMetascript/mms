package nbt

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"fmt"
	"math"
	"sort"
)

// Encode writes a root compound, gzipped as the game writes structure
// files. It takes the types Decode returns; a list's element type comes
// from its first element. Compound keys are written sorted, so the output
// is reproducible.
func Encode(root map[string]any) ([]byte, error) {
	var b bytes.Buffer
	b.WriteByte(tagCompound)
	writeString(&b, "")
	if err := writePayload(&b, root); err != nil {
		return nil, err
	}
	var z bytes.Buffer
	w := gzip.NewWriter(&z)
	w.Write(b.Bytes())
	w.Close()
	return z.Bytes(), nil
}

func tagOf(v any) (byte, error) {
	switch v.(type) {
	case int8:
		return tagByte, nil
	case int16:
		return tagShort, nil
	case int32:
		return tagInt, nil
	case int64:
		return tagLong, nil
	case float32:
		return tagFloat, nil
	case float64:
		return tagDouble, nil
	case []int8:
		return tagByteArray, nil
	case string:
		return tagString, nil
	case []any:
		return tagList, nil
	case map[string]any:
		return tagCompound, nil
	case []int32:
		return tagIntArray, nil
	case []int64:
		return tagLongArray, nil
	}
	return 0, fmt.Errorf("nbt: cannot encode %T", v)
}

func writeString(b *bytes.Buffer, s string) {
	binary.Write(b, binary.BigEndian, uint16(len(s)))
	b.WriteString(s)
}

func writePayload(b *bytes.Buffer, v any) error {
	switch v := v.(type) {
	case int8:
		b.WriteByte(byte(v))
	case int16, int32, int64:
		binary.Write(b, binary.BigEndian, v)
	case float32:
		binary.Write(b, binary.BigEndian, math.Float32bits(v))
	case float64:
		binary.Write(b, binary.BigEndian, math.Float64bits(v))
	case []int8:
		binary.Write(b, binary.BigEndian, int32(len(v)))
		for _, x := range v {
			b.WriteByte(byte(x))
		}
	case string:
		writeString(b, v)
	case []any:
		et := tagEnd
		if len(v) > 0 {
			t, err := tagOf(v[0])
			if err != nil {
				return err
			}
			et = t
		}
		b.WriteByte(et)
		binary.Write(b, binary.BigEndian, int32(len(v)))
		for _, x := range v {
			if err := writePayload(b, x); err != nil {
				return err
			}
		}
	case map[string]any:
		keys := make([]string, 0, len(v))
		for k := range v {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			t, err := tagOf(v[k])
			if err != nil {
				return err
			}
			b.WriteByte(t)
			writeString(b, k)
			if err := writePayload(b, v[k]); err != nil {
				return err
			}
		}
		b.WriteByte(tagEnd)
	case []int32:
		binary.Write(b, binary.BigEndian, int32(len(v)))
		binary.Write(b, binary.BigEndian, v)
	case []int64:
		binary.Write(b, binary.BigEndian, int32(len(v)))
		binary.Write(b, binary.BigEndian, v)
	default:
		return fmt.Errorf("nbt: cannot encode %T", v)
	}
	return nil
}
