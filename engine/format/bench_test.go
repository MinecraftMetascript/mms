package format

import (
	"strings"
	"testing"
)

var benchSrc = strings.Repeat(`LootTable x {
  pools = [{ rolls = 1, entries = [{ type = "item", name = ruby }] }] // c
  functions = [SetCount(1..3), Foo(a = 1)]
}

Function f {
  let x = 1 + 2 * (3 - y)
  /say hi
  if (x > 1) { x += 1 }
}
`, 100)

// BenchmarkSource formats a 1,000-line file, including the safety check.
func BenchmarkSource(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := Source(benchSrc); err != nil {
			b.Fatal(err)
		}
	}
}
