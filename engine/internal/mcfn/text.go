package mcfn

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MinecraftMetascript/nbt-go"
)

// text renders a text component as the plain text a player would read:
// scores and storage are filled in, a selector shows as written, and
// styles are dropped.
func (vm *VM) text(src string) (string, error) {
	c, err := nbt.ParseSNBT(src)
	if err != nil {
		return "", fmt.Errorf("text component %s: %w", src, err)
	}
	var b strings.Builder
	if err := vm.render(&b, c); err != nil {
		return "", err
	}
	return b.String(), nil
}

func (vm *VM) render(b *strings.Builder, c nbt.Tag) error {
	switch c := c.(type) {
	case nbt.String:
		b.WriteString(string(c))
	case nbt.List:
		for _, x := range c {
			if err := vm.render(b, x); err != nil {
				return err
			}
		}
	case nbt.Compound:
		str := func(k string) string {
			s, _ := c[k].(nbt.String)
			return string(s)
		}
		switch {
		case c["text"] != nil:
			b.WriteString(toText(c["text"]))
		case c["score"] != nil:
			sc, _ := c["score"].(nbt.Compound)
			name, _ := sc["name"].(nbt.String)
			obj, _ := sc["objective"].(nbt.String)
			if v, ok := vm.Score(string(name), string(obj)); ok {
				b.WriteString(strconv.Itoa(int(v)))
			}
		case c["nbt"] != nil:
			path, id := str("nbt"), str("storage")
			if id == "" {
				return fmt.Errorf("an nbt component without storage is not supported")
			}
			v, err := vm.get(id, path)
			if err != nil {
				return err
			}
			// Like the game, a string tag is shown as SNBT (quoted) unless
			// the component interprets it.
			interpret, _ := number(c["interpret"])
			if s, ok := v.(nbt.String); ok && interpret == 0 {
				b.WriteString(s.String())
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
