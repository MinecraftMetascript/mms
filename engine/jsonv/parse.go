package jsonv

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Parse decodes JSON, keeping object key order and number text.
func Parse(data []byte) (Value, error) {
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	v, err := parseValue(dec)
	if err != nil {
		return nil, err
	}
	if dec.More() {
		return nil, fmt.Errorf("jsonv: trailing data")
	}
	return v, nil
}

func parseValue(dec *json.Decoder) (Value, error) {
	tok, err := dec.Token()
	if err != nil {
		return nil, err
	}
	switch t := tok.(type) {
	case json.Delim:
		switch t {
		case '{':
			o := &Object{}
			for dec.More() {
				kt, err := dec.Token()
				if err != nil {
					return nil, err
				}
				k, _ := kt.(string)
				v, err := parseValue(dec)
				if err != nil {
					return nil, err
				}
				o.Members = append(o.Members, Member{Key: k, Value: v})
			}
			_, err := dec.Token()
			return o, err
		case '[':
			a := Array{}
			for dec.More() {
				v, err := parseValue(dec)
				if err != nil {
					return nil, err
				}
				a = append(a, v)
			}
			_, err := dec.Token()
			return a, err
		}
	case string:
		return String(t), nil
	case json.Number:
		return Number(t), nil
	case bool:
		return Bool(t), nil
	case nil:
		return Null{}, nil
	}
	return nil, fmt.Errorf("jsonv: unexpected token %v", tok)
}
