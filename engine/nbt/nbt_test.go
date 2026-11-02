package nbt

import (
	"bytes"
	"compress/gzip"
	"reflect"
	"testing"
)

func TestRoundTrip(t *testing.T) {
	root := map[string]any{
		"DataVersion": int32(4440),
		"size":        []any{int32(3), int32(2), int32(1)},
		"palette":     []any{map[string]any{"Name": "minecraft:jigsaw", "Properties": map[string]any{"orientation": "north_up"}}},
		"blocks": []any{map[string]any{
			"pos":   []any{int32(1), int32(0), int32(0)},
			"state": int32(0),
			"nbt":   map[string]any{"id": "minecraft:jigsaw", "pool": "demo:tops", "name": "demo:bottom", "target": "demo:top"},
		}},
		"b": int8(-1), "s": int16(300), "l": int64(1) << 40, "f": float32(1.5), "d": 2.25,
		"ba": []int8{1, -2}, "ia": []int32{7, 8}, "la": []int64{9}, "empty": []any{},
	}
	data, err := Encode(root)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Decode(data)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, root) {
		t.Fatalf("round trip:\n got %#v\nwant %#v", got, root)
	}
	// Uncompressed files read the same.
	z, _ := gzip.NewReader(bytes.NewReader(data))
	var plain bytes.Buffer
	plain.ReadFrom(z)
	if got, err := Decode(plain.Bytes()); err != nil || !reflect.DeepEqual(got, root) {
		t.Fatalf("plain: %v", err)
	}
}

func TestDamaged(t *testing.T) {
	data, _ := Encode(map[string]any{"a": "hello", "b": []any{int32(1), int32(2)}})
	z, _ := gzip.NewReader(bytes.NewReader(data))
	var plain bytes.Buffer
	plain.ReadFrom(z)
	raw := plain.Bytes()
	for n := 0; n < len(raw); n++ {
		if _, err := Decode(raw[:n]); err == nil {
			t.Fatalf("truncated to %d bytes: no error", n)
		}
	}
	for _, bad := range [][]byte{nil, {0x1f, 0x8b, 8, 0}, {9, 0, 0}, {10, 0, 0, 99}} {
		if _, err := Decode(bad); err == nil {
			t.Fatalf("%v: no error", bad)
		}
	}
}
