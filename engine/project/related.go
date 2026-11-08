package project

import (
	"github.com/minecraftmetascript/mms/engine/eval"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// relate points each diagnostic on text written elsewhere, such as a
// value a Template body wrote, at that text: "in Template Piece". The
// diagnostic stays on the user's span, where the value went in, and the
// related location is the line to read or fix. It is worked out on each
// call from this analysis's expansion, so it never lags an edit.
func (r *Result) relate(path string, ds []syntax.Diagnostic) {
	if r.Eval == nil {
		return
	}
	r.relOnce.Do(func() {
		r.written = map[eval.Origin]eval.Prov{}
		vague := map[eval.Origin]bool{}
		for _, p := range r.Eval.Nodes {
			if p.Def == nil || *p.Def == p.At || vague[p.At] {
				continue
			}
			old, ok := r.written[p.At]
			if !ok || *old.Def == *p.Def {
				r.written[p.At] = p
				continue
			}
			// Several lines wrote values that went in at this span, such
			// as a Template call's whole output, and a diagnostic only
			// has its span. Point at what they share: the Template,
			// Blueprint or const itself.
			if f, ok := sharedFrame(old.Frames, p.Frames); ok {
				def := f.Def
				r.written[p.At] = eval.Prov{At: p.At, Def: &def, Frames: []eval.Frame{f}}
			} else {
				delete(r.written, p.At)
			}
			vague[p.At] = true
		}
	})
	for i, d := range ds {
		if len(d.Related) > 0 {
			continue
		}
		// The statement the checker found the problem in names the
		// line exactly; failing that, the span names what wrote there.
		p, ok := r.Eval.Lines[d.At]
		if !ok || p.At.Path != path || p.At.Sp != d.Sp {
			if p, ok = r.written[eval.Origin{Path: path, Sp: d.Sp}]; !ok {
				continue
			}
		}
		ds[i].Related = []syntax.Related{{Path: p.Def.Path, Sp: p.Def.Sp, Message: eval.WrittenIn(p.Frames)}}
	}
}

// sharedFrame is the innermost named frame two expansions both came
// through.
func sharedFrame(a, b []eval.Frame) (eval.Frame, bool) {
	var out eval.Frame
	ok := false
	for i := 0; i < len(a) && i < len(b) && a[i] == b[i]; i++ {
		if a[i].Kind != eval.FrameFor && a[i].Name != "" {
			out, ok = a[i], true
		}
	}
	return out, ok
}
