package noise

import (
	"encoding/json"
	"fmt"

	"github.com/minecraftmetascript/mms/lib"
)

func Export(ns *lib.Namespace, rootDir *lib.FileTreeLike) error {
	dir := rootDir.MkDir("worldgen", nil).MkDir("noise", nil)

	for name, noise := range lib.AllOf[Noise](ns) {
		contents, err := json.MarshalIndent(noise.Value, "", "  ")
		if err != nil {
			continue
		}
		dir.MkFile(
			fmt.Sprintf("%s.json", name),
			string(contents),
			noise,
		)
	}

	return nil
}
