package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestFmtFiles(t *testing.T) {
	dir := t.TempDir()
	messy := filepath.Join(dir, "a.mms")
	clean := filepath.Join(dir, "sub", "b.mms")
	broken := filepath.Join(dir, "c.mms")
	os.MkdirAll(filepath.Dir(clean), 0o755)
	os.WriteFile(messy, []byte("Objective kills=deathCount\n"), 0o644)
	os.WriteFile(clean, []byte("Objective kills = deathCount\n"), 0o644)
	os.WriteFile(broken, []byte("LootTable x {\n"), 0o644)

	var out, errOut bytes.Buffer
	if code := runFmt([]string{"-check", dir}, nil, &out, &errOut); code != 1 {
		t.Fatalf("check exit = %d, want 1", code)
	}
	if got := strings.TrimSpace(out.String()); got != messy {
		t.Errorf("check listed %q, want %q", got, messy)
	}
	if !strings.Contains(errOut.String(), "c.mms") {
		t.Errorf("broken file not reported: %q", errOut.String())
	}

	out.Reset()
	errOut.Reset()
	runFmt([]string{messy, clean}, nil, &out, &errOut)
	got, _ := os.ReadFile(messy)
	if string(got) != "Objective kills = deathCount\n" {
		t.Errorf("formatted file = %q", got)
	}
	if code := runFmt([]string{"-check", messy, clean}, nil, &out, &errOut); code != 0 {
		t.Errorf("check after fmt exit = %d, want 0", code)
	}
}

func TestFmtStdin(t *testing.T) {
	var out, errOut bytes.Buffer
	code := runFmt(nil, strings.NewReader("Objective   kills=deathCount"), &out, &errOut)
	if code != 0 || out.String() != "Objective kills = deathCount\n" {
		t.Errorf("stdin: code %d, out %q, err %q", code, out.String(), errOut.String())
	}
}

func TestFix(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "mms.json"), []byte(`{"namespace": "demo"}`), 0o644)
	os.WriteFile(filepath.Join(dir, "a.mms"), []byte(`ItemTag gems = ["minecraft:diamond", minecraft:emerald]`+"\n"), 0o644)
	var out, errOut bytes.Buffer
	if code := runFix([]string{dir}, &out, &errOut); code != 0 {
		t.Fatalf("fix exit %d: %s", code, errOut.String())
	}
	got, _ := os.ReadFile(filepath.Join(dir, "a.mms"))
	if string(got) != "ItemTag gems = [diamond, emerald]\n" {
		t.Errorf("fixed = %q (%s)", got, out.String())
	}
}
