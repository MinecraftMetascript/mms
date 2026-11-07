package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestTestCommand(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "mms.json"), []byte(`{"namespace":"demo"}`), 0o644)
	os.WriteFile(filepath.Join(dir, "main.mms"), []byte(`Global n: int = 0
Function bump {
  n += 1
  /say bumped
}
Test bumps_once {
  run bump()
  expect n == 1
}
Test bumps_twice {
  run bump()
  expect n == 2
}
`), 0o644)
	var out, errb bytes.Buffer
	if code := runTest([]string{dir}, &out, &errb); code != 1 {
		t.Fatalf("exit %d, want 1\n%s%s", code, out.String(), errb.String())
	}
	want := "FAIL  bumps_twice\n    main.mms:12:3: expected n == 2, but n is 1\n    said: bumped\n1 of 2 tests failed\n"
	if out.String() != want {
		t.Errorf("output:\n%s\nwant:\n%s", out.String(), want)
	}
	out.Reset()
	if code := runTest([]string{"-run", "once$", dir}, &out, &errb); code != 0 || !strings.Contains(out.String(), "ok: 1 test passed") {
		t.Errorf("-run: exit %d\n%s", code, out.String())
	}
}
