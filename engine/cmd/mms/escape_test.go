package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestBuildStaysInsideProject checks that nothing in a project (a
// declaration name, mms.json's output folders or namespace, or a symlink)
// can make `mms build` write or delete outside its output folder.
func TestBuildStaysInsideProject(t *testing.T) {
	cases := []struct {
		name, config, src, symlink, want string
	}{
		{name: "dotdot name", config: `{"namespace":"demo"}`, src: "Function a/../../../evil {\n  /say hi\n}\n", want: "can't be empty, '.' or '..'"},
		{name: "output dot", config: `{"namespace":"demo","output":"."}`, want: "must be a folder inside the project"},
		{name: "output parent", config: `{"namespace":"demo","output":"../victim"}`, want: "must be a folder inside the project"},
		{name: "resource output parent", config: `{"namespace":"demo","resourcePack":{"output":"../victim"}}`, want: "must be a folder inside the project"},
		{name: "namespace dotdot", config: `{"namespace":".."}`, want: "namespace"},
		{name: "symlinked output", config: `{"namespace":"demo"}`, symlink: "build", want: "outside the project"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			root := t.TempDir()
			dir := filepath.Join(root, "proj")
			victim := filepath.Join(root, "victim")
			keep := filepath.Join(victim, "data", "keep.txt")
			for _, d := range []string{dir, filepath.Dir(keep)} {
				if err := os.MkdirAll(d, 0o755); err != nil {
					t.Fatal(err)
				}
			}
			if err := os.WriteFile(keep, []byte("keep"), 0o644); err != nil {
				t.Fatal(err)
			}
			// The project's own data/ must survive too ("output": ".").
			own := filepath.Join(dir, "data", "mine.txt")
			if err := os.MkdirAll(filepath.Dir(own), 0o755); err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(own, []byte("mine"), 0o644); err != nil {
				t.Fatal(err)
			}
			src := tc.src
			if src == "" {
				src = "Function hello {\n  /say hi\n}\n"
			}
			write(t, filepath.Join(dir, "mms.json"), tc.config)
			write(t, filepath.Join(dir, "main.mms"), src)
			if tc.symlink != "" {
				if err := os.Symlink(victim, filepath.Join(dir, tc.symlink)); err != nil {
					t.Skip("symlinks unavailable:", err)
				}
			}
			var out, errOut bytes.Buffer
			if code := runBuild([]string{dir}, &out, &errOut); code == 0 {
				t.Fatalf("build succeeded; want an error containing %q\n%s", tc.want, out.String())
			}
			if !strings.Contains(errOut.String()+out.String(), tc.want) {
				t.Errorf("error %q lacks %q", errOut.String()+out.String(), tc.want)
			}
			for _, f := range []string{keep, own} {
				if _, err := os.Stat(f); err != nil {
					t.Errorf("%s is gone: %v", f, err)
				}
			}
			if _, err := os.Stat(filepath.Join(root, "evil.mcfunction")); err == nil {
				t.Error("wrote evil.mcfunction outside the project")
			}
		})
	}
}

func write(t *testing.T, p, text string) {
	t.Helper()
	if err := os.WriteFile(p, []byte(text), 0o644); err != nil {
		t.Fatal(err)
	}
}
