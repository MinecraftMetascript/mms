package block_states

import (
	"encoding/json"
	"github.com/minecraftmetascript/mms/lang/traversal"
)

type BlockState struct {
	traversal.BaseConstruct
	Name string
}

func (bs BlockState) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(struct {
		Name string `json:"Name"`
	}{
		Name: bs.Name,
	}, "", "  ")
}
