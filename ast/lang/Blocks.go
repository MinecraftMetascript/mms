package lang

import (
	"slices"

	"github.com/minecraftmetascript/mms/ast/lang/spec"
)

var Blocks = []spec.Block{
	NoiseBlock,
}

func GetBlockSpec(kind string) *spec.Block {
	idx := slices.IndexFunc(Blocks, func(s spec.Block) bool {
		return s.Kind == kind
	})
	if idx == -1 {
		return nil
	}
	return &Blocks[idx]
}
