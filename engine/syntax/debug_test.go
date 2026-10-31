package syntax

import (
	"os"
	"testing"
)

// TestDumpFile prints the outline of $MMS_DUMP, for debugging.
func TestDumpFile(t *testing.T) {
	path := os.Getenv("MMS_DUMP")
	if path == "" {
		t.Skip("set MMS_DUMP to a file")
	}
	src, _ := os.ReadFile(path)
	t.Log("\n" + Dump(Parse(string(src))))
}
