package lib

import (
	"os"
	"path/filepath"
)

func RecursiveFilesInDir(root string) ([]string, error) {
	files := make([]string, 0)
	entries, err := os.ReadDir(root)
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			subFiles, err := RecursiveFilesInDir(filepath.Join(root, entry.Name()))
			if err != nil {
				return nil, err
			}
			files = append(files, subFiles...)
		} else {
			files = append(files, root+"/"+entry.Name())
		}
	}
	return files, nil
}
