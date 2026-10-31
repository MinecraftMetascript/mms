package spec

import (
	"os"
	"strings"
	"testing"
)

// TestPrintType prints the types named in $MMS_TYPE (comma separated), for
// exploring the pack. A name without "::" is looked up as a resource kind.
func TestPrintType(t *testing.T) {
	names := os.Getenv("MMS_TYPE")
	if names == "" {
		t.Skip("set MMS_TYPE")
	}
	p := MustDefault()
	for _, n := range strings.Split(names, ",") {
		ty := p.Types[n]
		if r, ok := p.Resources[n]; ok {
			ty = r
		}
		if strings.HasPrefix(n, "dispatch:") {
			parts := strings.SplitN(strings.TrimPrefix(n, "dispatch:"), "/", 2)
			if len(parts) == 2 {
				ty, _ = p.Dispatch(parts[0], parts[1])
			} else {
				t.Logf("%s keys: %v", n, p.DispatchKeys(parts[0]))
				continue
			}
		}
		t.Logf("%s =\n%s", n, Format(p.Resolve(ty), 3))
	}
}
