package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestInitThenCheck checks that a new project checks and builds clean,
// that init refuses to overwrite one, and that a folder name becomes a
// valid namespace.
func TestInitThenCheck(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "My Pack!")
	var out, errOut bytes.Buffer
	if code := runInit([]string{dir}, &out, &errOut); code != 0 {
		t.Fatalf("init exit %d: %s", code, errOut.String())
	}
	cfg, _ := os.ReadFile(filepath.Join(dir, "mms.json"))
	if !strings.Contains(string(cfg), `"namespace": "my_pack_"`) {
		t.Errorf("mms.json = %s", cfg)
	}
	out.Reset()
	if code := runCheck([]string{dir}, &out, &errOut); code != 0 || !strings.HasPrefix(out.String(), "ok: 1 file") {
		t.Errorf("check exit %d: %s%s", code, out.String(), errOut.String())
	}
	if code := runBuild([]string{dir}, &out, &errOut); code != 0 {
		t.Errorf("build exit %d: %s", code, errOut.String())
	}
	errOut.Reset()
	if code := runInit([]string{dir}, &out, &errOut); code != 1 || !strings.Contains(errOut.String(), "already exists") {
		t.Errorf("second init exit %d: %s", code, errOut.String())
	}
}

func TestSanitizeNS(t *testing.T) {
	for in, want := range map[string]string{
		"pack":      "pack",
		"My Pack":   "my_pack",
		"a.b-c_d":   "a.b-c_d",
		"Ünïcode":   "_n_code",
		"":          "mms",
		"UPPER_123": "upper_123",
	} {
		if got := sanitizeNS(in); got != want {
			t.Errorf("sanitizeNS(%q) = %q, want %q", in, got, want)
		}
	}
}

// TestCheckReportsErrors checks that check exits 1 with the problem and a
// count, and prints hints only with -hints.
func TestCheckReportsErrors(t *testing.T) {
	dir := t.TempDir()
	write(t, filepath.Join(dir, "mms.json"), `{"namespace": "demo"}`)
	write(t, filepath.Join(dir, "main.mms"), "DensityFunction x = Clmp(1, 0, 1)\n")
	var out, errOut bytes.Buffer
	if code := runCheck([]string{dir}, &out, &errOut); code != 1 || !strings.Contains(errOut.String(), "Clamp") || !strings.Contains(errOut.String(), "1 error") {
		t.Errorf("check exit %d: %s", code, errOut.String())
	}

	write(t, filepath.Join(dir, "main.mms"), "DensityFunction x = 2 * 0.25\n")
	errOut.Reset()
	if code := runCheck([]string{dir}, &out, &errOut); code != 0 || strings.Contains(errOut.String(), "constant-expression") {
		t.Errorf("check without -hints exit %d: %s", code, errOut.String())
	}
	errOut.Reset()
	if code := runCheck([]string{"-hints", dir}, &out, &errOut); code != 0 || !strings.Contains(errOut.String(), "constant-expression") {
		t.Errorf("check -hints exit %d: %s", code, errOut.String())
	}

	write(t, filepath.Join(dir, "mms.json"), `{"namepsace": "demo"}`)
	errOut.Reset()
	if code := runCheck([]string{dir}, &out, &errOut); code != 1 || !strings.Contains(errOut.String(), "namepsace") {
		t.Errorf("check with a misspelt key exit %d: %s", code, errOut.String())
	}
}

// TestReference checks that the reference is JSON with kinds, helpers,
// the prelude and every lint rule.
func TestReference(t *testing.T) {
	var out, errOut bytes.Buffer
	if code := runReference(nil, &out, &errOut); code != 0 {
		t.Fatalf("reference exit %d: %s", code, errOut.String())
	}
	var ref map[string]any
	if err := json.Unmarshal(out.Bytes(), &ref); err != nil {
		t.Fatal(err)
	}
	rules, _ := ref["rules"].([]any)
	if len(rules) < 10 {
		t.Errorf("reference has %d rules", len(rules))
	}
	if ref["prelude"] == nil {
		t.Error("reference has no prelude")
	}
	if code := runReference([]string{"extra"}, &out, &errOut); code != 2 {
		t.Errorf("reference with an argument exit %d", code)
	}
}
