//go:build unix

package main

import (
	"os"
	"path/filepath"
	"syscall"
)

// worldInUse reports whether Minecraft has the world open. The game holds
// a lock on <world>/session.lock while a world is loaded (Java's
// FileChannel lock, which is an fcntl record lock on Linux and macOS), so
// another process asking about the lock sees it.
func worldInUse(world string) bool {
	f, err := os.OpenFile(filepath.Join(world, "session.lock"), os.O_RDWR, 0)
	if err != nil {
		return false
	}
	defer f.Close()
	lk := syscall.Flock_t{Type: syscall.F_WRLCK, Whence: 0, Start: 0, Len: 0}
	if err := syscall.FcntlFlock(f.Fd(), syscall.F_GETLK, &lk); err != nil {
		return false
	}
	return lk.Type != syscall.F_UNLCK
}
