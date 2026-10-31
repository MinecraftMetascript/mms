package format

import (
	"os"
	"testing"

	"github.com/minecraftmetascript/mms/engine/syntax"
)

// TestDebugFile formats $MMS_FMT without the safety check and shows both
// outlines, for debugging.
func TestDebugFile(t *testing.T) {
	path := os.Getenv("MMS_FMT")
	if path == "" {
		t.Skip("set MMS_FMT to a file")
	}
	src, _ := os.ReadFile(path)
	f := syntax.Parse(string(src))
	out := File(f)
	t.Log("\n" + out)
	re := syntax.Parse(out)
	if a, b := syntax.Dump(f), syntax.Dump(re); a != b {
		t.Log("outline differs\n--- in\n" + a + "--- out\n" + b)
	}
	t.Logf("comments in %q\ncomments out %q", commentTexts(f), commentTexts(re))
}
