package ide

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/project"
	"github.com/minecraftmetascript/mms/engine/spec"
)

// benchSource is a realistic 2,000-line file: noises, density functions
// with sugar, and surface rules.
func benchSource() string {
	var sb strings.Builder
	for i := 0; sb.Len() == 0 || strings.Count(sb.String(), "\n") < 2000; i++ {
		fmt.Fprintf(&sb, `Noise n%[1]d = {
  base_octave = -7
  amplitude_modifiers = [1, 1, 0.5]
}

DensityFunction d%[1]d = Clamp(Noise(n%[1]d, 1, 0.5) * 2 + 0.5 - Abs(Noise(n%[1]d, 0.25, 0)), -1, 1)

DensityFunction e%[1]d = Abs(Cache(d%[1]d))

SurfaceRule s%[1]d = [
  If(!Water(-1, 0)) { Block(stone) }
  If(StoneDepth(floor, 0, false, 0)) { Block(grass_block) }
]

`, i)
	}
	return sb.String()
}

func benchProject(b *testing.B) (*project.Project, string) {
	src := benchSource()
	p := project.New(spec.MustDefault(), project.Config{Namespace: "bench"})
	p.SetFile("main.mms", src)
	r, err := p.Analyze(context.Background())
	if err != nil {
		b.Fatal(err)
	}
	if r.HasErrors() {
		b.Fatalf("bench source has errors: %v", r.Diagnostics("main.mms")[0])
	}
	return p, src
}

// BenchmarkAnalyzeEdit re-analyses after a one-character edit, the work
// done on every keystroke.
func BenchmarkAnalyzeEdit(b *testing.B) {
	p, src := benchProject(b)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.SetFile("main.mms", src+fmt.Sprintf("// %d\n", i))
		r, err := p.Analyze(context.Background())
		if err != nil {
			b.Fatal(err)
		}
		r.Diagnostics("main.mms") // includes lint, as the server publishes
	}
}

// BenchmarkCompletion completes a method at the end of the file.
func BenchmarkCompletion(b *testing.B) {
	p, src := benchProject(b)
	src += "DensityFunction z = Noise(n0, 1, 1)."
	p.SetFile("main.mms", src)
	r, _ := p.Analyze(context.Background())
	s := &Session{P: p, R: r}
	off := len([]rune(src))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if items, _ := s.Completion("main.mms", off); len(items) == 0 {
			b.Fatal("no items")
		}
	}
}

// BenchmarkAnalyzeEditTop edits the first declaration, so every later
// declaration moves.
func BenchmarkAnalyzeEditTop(b *testing.B) {
	p, src := benchProject(b)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.SetFile("main.mms", fmt.Sprintf("// %d\n", i%2)+strings.Repeat(" ", i%2)+src)
		if _, err := p.Analyze(context.Background()); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkAnalyzeNewName types a new declaration name, which changes the
// set of declared names.
func BenchmarkAnalyzeNewName(b *testing.B) {
	p, src := benchProject(b)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.SetFile("main.mms", src+fmt.Sprintf("DensityFunction new_%d = 1\n", i))
		if _, err := p.Analyze(context.Background()); err != nil {
			b.Fatal(err)
		}
	}
}
