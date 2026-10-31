package command

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

// TestCorpus checks every command in MMS_COMMAND_CORPUS (a folder of
// .mcfunction files) and prints the problems found, most common first.
func TestCorpus(t *testing.T) {
	dir := os.Getenv("MMS_COMMAND_CORPUS")
	if dir == "" {
		t.Skip("MMS_COMMAND_CORPUS is not set")
	}
	c := checker(t)
	counts := map[string]int{}
	examples := map[string]string{}
	total, bad := 0, 0
	filepath.WalkDir(dir, func(p string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(p, ".mcfunction") {
			return nil
		}
		f, _ := os.Open(p)
		defer f.Close()
		sc := bufio.NewScanner(f)
		sc.Buffer(nil, 1<<20)
		pending := ""
		for sc.Scan() {
			line := strings.TrimSpace(sc.Text())
			if strings.HasSuffix(line, "\\") {
				pending += strings.TrimSuffix(line, "\\")
				continue
			}
			line, pending = pending+line, ""
			if line == "" || line[0] == '#' || strings.Contains(line, "{{") || strings.Contains(line, "{%") {
				continue
			}
			total++
			if r := c.Check(line); r.Problem != nil {
				bad++
				msg := r.Problem.Message
				if strings.HasPrefix(msg, "unknown command") {
					msg = strings.SplitN(msg, ";", 2)[0]
				} else if i := strings.IndexAny(msg, "'<"); i > 0 {
					msg = msg[:i]
				}
				counts[msg]++
				if _, ok := examples[msg]; !ok || len(line) < len(examples[msg]) {
					examples[msg] = fmt.Sprintf("%s\n      at %d: %s", line, r.Problem.Start, r.Problem.Message)
				}
			}
		}
		return nil
	})
	var keys []string
	for k := range counts {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return counts[keys[i]] > counts[keys[j]] })
	for _, k := range keys {
		t.Logf("%5d %s\n    e.g. %s", counts[k], k, examples[k])
	}
	t.Logf("%d of %d commands have problems", bad, total)
}
