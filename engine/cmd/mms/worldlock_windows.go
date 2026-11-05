//go:build windows

package main

import (
	"os"
	"path/filepath"
)

// worldInUse reports whether Minecraft has the world open. The game holds
// a lock on <world>/session.lock while a world is loaded, and on Windows
// that lock makes reading the file fail.
func worldInUse(world string) bool {
	p := filepath.Join(world, "session.lock")
	if _, err := os.Stat(p); err != nil {
		return false
	}
	_, err := os.ReadFile(p)
	return err != nil
}
