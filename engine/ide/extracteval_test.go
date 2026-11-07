package ide

import (
	"context"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/project"
	"github.com/minecraftmetascript/mms/engine/spec"
)

func extractionTitled(t *testing.T, s *Session, src, sel, title string) (Extraction, bool) {
	t.Helper()
	from := len([]rune(src[:strings.Index(src, sel)]))
	for _, e := range s.Extractions("main.mms", from, from+len([]rune(sel))) {
		if strings.HasSuffix(e.Title, title) {
			return e, true
		}
	}
	return Extraction{}, false
}

// errorsIn reports the errors of src as one project file.
func errorsIn(t *testing.T, src string) string {
	t.Helper()
	p := project.New(spec.MustDefault(), project.Config{Namespace: "test"})
	p.SetFile("main.mms", src)
	r, err := p.Analyze(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	var msgs []string
	for _, d := range r.Diagnostics("main.mms") {
		if d.Severity == 1 {
			msgs = append(msgs, d.Message)
		}
	}
	return strings.Join(msgs, "; ")
}

// A selected value becomes a const, and selected commands a Template;
// the project still builds afterwards.
func TestExtractEval(t *testing.T) {
	src := "Function f {\n  /say one\n  /say two\n  let n = 1\n  /say ${n}\n}\nCustomItem ruby { base = diamond, name = \"Ruby\" }\n"
	s, _ := session(t, src)

	e, ok := extractionTitled(t, s, src, `"Ruby"`, "to a const")
	if !ok {
		t.Fatal("no const extraction")
	}
	got := apply(src, e.Edits)
	if !strings.HasPrefix(got[strings.Index(got, "const"):], "const name = \"Ruby\"\n") || !strings.Contains(got, "name = name") {
		t.Errorf("const extraction:\n%s", got)
	}
	if string([]rune(got)[e.NameAt:e.NameAt+4]) != "name" {
		t.Errorf("NameAt = %d", e.NameAt)
	}
	if errs := errorsIn(t, got); errs != "" {
		t.Errorf("after const extraction: %s\n%s", errs, got)
	}

	e, ok = extractionTitled(t, s, src, "/say one\n  /say two", "to a Template")
	if !ok {
		t.Fatal("no Template extraction")
	}
	got = apply(src, e.Edits)
	if !strings.HasPrefix(got, "Template commands(): Commands {\n  /say one\n  /say two\n}\n\nFunction f {\n  commands()\n  let n") {
		t.Errorf("Template extraction:\n%s", got)
	}
	if errs := errorsIn(t, got); errs != "" {
		t.Errorf("after Template extraction: %s\n%s", errs, got)
	}

	// A command that reads a local stays where it is.
	if _, ok := extractionTitled(t, s, src, "/say ${n}", "to a Template"); ok {
		t.Error("extracted a command reading a local")
	}
}
