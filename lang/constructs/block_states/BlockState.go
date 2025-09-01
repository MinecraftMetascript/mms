package block_states

import (
	"encoding/json"
	"mms2/lang/traversal"
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
