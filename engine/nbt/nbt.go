// Package nbt reads Minecraft's Named Binary Tag format, the format of
// structure files (.nbt) and level.dat. Values decode to Go types:
// byte, short, int, long, float and double to int8, int16, int32, int64,
// float32 and float64; byte, int and long arrays to []int8, []int32 and
// []int64; strings to string; lists to []any; compounds to
// map[string]any. Files may be gzipped, as the game writes them, or not.
package nbt

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math"
)

// Tag types.
const (
	tagEnd byte = iota
	tagByte
	tagShort
	tagInt
	tagLong
	tagFloat
	tagDouble
	tagByteArray
	tagString
	tagList
	tagCompound
	tagIntArray
	tagLongArray
)

// maxDepth bounds nesting, and maxLen the length of one array, list or
// string, so a damaged file cannot exhaust memory or the stack.
const (
	maxDepth = 512
	maxLen   = 1 << 26
)

// Decode reads a file's root compound. It takes the file gzipped or
// plain.
func Decode(data []byte) (map[string]any, error) {
	var r io.Reader = bytes.NewReader(data)
	if len(data) >= 2 && data[0] == 0x1f && data[1] == 0x8b {
		z, err := gzip.NewReader(r)
		if err != nil {
			return nil, err
		}
		defer z.Close()
		r = z
	}
	d := &decoder{r: bufio.NewReader(r)}
	t, err := d.byte()
	if err != nil {
		return nil, err
	}
	if t != tagCompound {
		return nil, fmt.Errorf("nbt: root is tag %d, not a compound", t)
	}
	if _, err := d.string(); err != nil {
		return nil, err
	}
	v, err := d.payload(tagCompound, 0)
	if err != nil {
		return nil, err
	}
	return v.(map[string]any), nil
}

type decoder struct {
	r   *bufio.Reader
	buf [8]byte
}

func (d *decoder) read(n int) ([]byte, error) {
	if _, err := io.ReadFull(d.r, d.buf[:n]); err != nil {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
		return nil, err
	}
	return d.buf[:n], nil
}

func (d *decoder) byte() (byte, error) {
	b, err := d.read(1)
	if err != nil {
		return 0, err
	}
	return b[0], nil
}

func (d *decoder) length() (int, error) {
	b, err := d.read(4)
	if err != nil {
		return 0, err
	}
	n := int32(binary.BigEndian.Uint32(b))
	if n < 0 {
		n = 0
	}
	if n > maxLen {
		return 0, errors.New("nbt: length too large")
	}
	return int(n), nil
}

// string reads a length-prefixed string. The game writes modified UTF-8,
// which matches UTF-8 for everything but NUL and characters outside the
// BMP; ids and names are ASCII.
func (d *decoder) string() (string, error) {
	b, err := d.read(2)
	if err != nil {
		return "", err
	}
	s := make([]byte, binary.BigEndian.Uint16(b))
	if _, err := io.ReadFull(d.r, s); err != nil {
		return "", io.ErrUnexpectedEOF
	}
	return string(s), nil
}

func (d *decoder) payload(t byte, depth int) (any, error) {
	if depth > maxDepth {
		return nil, errors.New("nbt: nested too deeply")
	}
	switch t {
	case tagByte:
		b, err := d.byte()
		return int8(b), err
	case tagShort:
		b, err := d.read(2)
		if err != nil {
			return nil, err
		}
		return int16(binary.BigEndian.Uint16(b)), nil
	case tagInt:
		b, err := d.read(4)
		if err != nil {
			return nil, err
		}
		return int32(binary.BigEndian.Uint32(b)), nil
	case tagLong:
		b, err := d.read(8)
		if err != nil {
			return nil, err
		}
		return int64(binary.BigEndian.Uint64(b)), nil
	case tagFloat:
		b, err := d.read(4)
		if err != nil {
			return nil, err
		}
		return math.Float32frombits(binary.BigEndian.Uint32(b)), nil
	case tagDouble:
		b, err := d.read(8)
		if err != nil {
			return nil, err
		}
		return math.Float64frombits(binary.BigEndian.Uint64(b)), nil
	case tagByteArray:
		n, err := d.length()
		if err != nil {
			return nil, err
		}
		raw := make([]byte, n)
		if _, err := io.ReadFull(d.r, raw); err != nil {
			return nil, io.ErrUnexpectedEOF
		}
		out := make([]int8, n)
		for i, b := range raw {
			out[i] = int8(b)
		}
		return out, nil
	case tagString:
		return d.string()
	case tagList:
		et, err := d.byte()
		if err != nil {
			return nil, err
		}
		n, err := d.length()
		if err != nil {
			return nil, err
		}
		out := make([]any, 0, min(n, 1024))
		for i := 0; i < n; i++ {
			v, err := d.payload(et, depth+1)
			if err != nil {
				return nil, err
			}
			out = append(out, v)
		}
		return out, nil
	case tagCompound:
		out := map[string]any{}
		for {
			et, err := d.byte()
			if err != nil {
				return nil, err
			}
			if et == tagEnd {
				return out, nil
			}
			name, err := d.string()
			if err != nil {
				return nil, err
			}
			v, err := d.payload(et, depth+1)
			if err != nil {
				return nil, err
			}
			out[name] = v
		}
	case tagIntArray, tagLongArray:
		n, err := d.length()
		if err != nil {
			return nil, err
		}
		if t == tagIntArray {
			out := make([]int32, 0, min(n, 1024))
			for i := 0; i < n; i++ {
				b, err := d.read(4)
				if err != nil {
					return nil, err
				}
				out = append(out, int32(binary.BigEndian.Uint32(b)))
			}
			return out, nil
		}
		out := make([]int64, 0, min(n, 1024))
		for i := 0; i < n; i++ {
			b, err := d.read(8)
			if err != nil {
				return nil, err
			}
			out = append(out, int64(binary.BigEndian.Uint64(b)))
		}
		return out, nil
	}
	return nil, fmt.Errorf("nbt: unknown tag type %d", t)
}
