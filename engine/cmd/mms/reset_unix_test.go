//go:build unix

package main

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"testing"
)

// TestHelperHoldLock is not a test: TestBuildResetOpenWorld runs the test
// binary with it to hold a world's session.lock from another process, as
// Minecraft does while the world is open.
func TestHelperHoldLock(t *testing.T) {
	p := os.Getenv("MMS_HOLD_LOCK")
	if p == "" {
		t.Skip("helper")
	}
	f, err := os.OpenFile(p, os.O_RDWR, 0)
	if err != nil {
		t.Fatal(err)
	}
	lk := syscall.Flock_t{Type: syscall.F_WRLCK}
	if err := syscall.FcntlFlock(f.Fd(), syscall.F_SETLKW, &lk); err != nil {
		t.Fatal(err)
	}
	os.Stdout.WriteString("locked\n")
	io.ReadAll(os.Stdin)
}

// TestBuildResetOpenWorld checks that a reset is refused, and nothing
// written, while the game has the world open.
func TestBuildResetOpenWorld(t *testing.T) {
	dir, world := copyExample(t, "sky-islands"), t.TempDir()
	os.WriteFile(filepath.Join(world, "level.dat"), nil, 0o644)
	lock := filepath.Join(world, "session.lock")
	os.WriteFile(lock, []byte("\xe2\x98\x83"), 0o644)
	if worldInUse(world) {
		t.Fatal("an unlocked world reads as open")
	}

	cmd := exec.Command(os.Args[0], "-test.run=^TestHelperHoldLock$")
	cmd.Env = append(os.Environ(), "MMS_HOLD_LOCK="+lock)
	stdin, _ := cmd.StdinPipe()
	stdout, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}
	defer func() { stdin.Close(); cmd.Wait() }()
	buf := make([]byte, 7)
	if _, err := io.ReadFull(stdout, buf); err != nil || string(buf) != "locked\n" {
		t.Fatalf("helper did not lock: %q %v", buf, err)
	}
	if !worldInUse(world) {
		t.Fatal("a locked world reads as closed")
	}

	var out, errOut bytes.Buffer
	if code := runBuild([]string{"-world", world, "-reset-dimensions", dir}, &out, &errOut); code != 1 || !strings.Contains(errOut.String(), "is open in Minecraft; leave the world to reset sky:islands") {
		t.Errorf("exit %d: %s", code, errOut.String())
	}
	if _, err := os.Stat(filepath.Join(world, "datapacks")); !os.IsNotExist(err) {
		t.Error("the refused build wrote the data pack")
	}
	// Without a reset to do, building into an open world is fine.
	if code := runBuild([]string{"-world", world, dir}, &out, &errOut); code != 0 {
		t.Errorf("plain build into an open world: exit %d: %s", code, errOut.String())
	}
}
