package mcfn

import (
	"fmt"
	"strconv"
	"strings"
)

// text renders a text component as the plain text a player would read:
// scores and storage are filled in, a selector shows as written, and
// styles are dropped.
func (vm *VM) text(src string) (string, error) {
	c, err := ParseSNBT(src)
	if err != nil {
		return "", fmt.Errorf("text component %s: %v", src, err)
	}
	var b strings.Builder
	if err := vm.render(&b, c); err != nil {
		return "", err
	}
	return b.String(), nil
}

func (vm *VM) render(b *strings.Builder, c any) error {
	switch c := c.(type) {
	case string:
		b.WriteString(c)
	case []any:
		for _, x := range c {
			if err := vm.render(b, x); err != nil {
				return err
			}
		}
	case map[string]any:
		switch {
		case c["text"] != nil:
			b.WriteString(toText(c["text"]))
		case c["score"] != nil:
			sc, _ := c["score"].(map[string]any)
			name, _ := sc["name"].(string)
			obj, _ := sc["objective"].(string)
			if v, ok := vm.Score(name, obj); ok {
				b.WriteString(strconv.Itoa(int(v)))
			}
		case c["nbt"] != nil:
			path, _ := c["nbt"].(string)
			id, _ := c["storage"].(string)
			if id == "" {
				return fmt.Errorf("an nbt component without storage is not supported")
			}
			v, err := vm.get(id, path)
			if err != nil {
				return err
			}
			// Like the game, a string tag is shown as SNBT (quoted) unless
			// the component interprets it.
			if s, ok := v.(string); ok && c["interpret"] != int32(1) {
				b.WriteString(SNBT(s))
			} else if v != nil {
				b.WriteString(toText(v))
			}
		case c["selector"] != nil:
			b.WriteString(toText(c["selector"]))
		case c["keybind"] != nil:
			b.WriteString("[" + toText(c["keybind"]) + "]")
		}
		if extra, ok := c["extra"]; ok {
			return vm.render(b, extra)
		}
	default:
		b.WriteString(toText(c))
	}
	return nil
}
