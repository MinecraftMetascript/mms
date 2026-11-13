package project

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/minecraftmetascript/mms/engine/eval"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// mustUseTag marks, in the doc comment of a const or declaration in a
// Blueprint's emit block, an output that does nothing unless the project
// reads it: `@mustUse no biome source places the biome`.
const mustUseTag = "mustUse"

// mustUseText is the text of a doc comment's @mustUse tag, and whether
// it has one.
func mustUseText(doc string) (string, bool) {
	if doc == "" || !strings.Contains(strings.ToLower(doc), "@mustuse") {
		return "", false
	}
	for _, t := range syntax.ParseDoc(doc).Other {
		if strings.EqualFold(t.Name, mustUseTag) {
			return strings.TrimSpace(t.Text), true
		}
	}
	return "", false
}

// mustUseDecls are the declarations Blueprint instances in the project
// emitted from an emit block declaration marked @mustUse, with the
// declaration that emitted them.
func mustUseDecls(r *Result) map[*syntax.Decl]*eval.Decl {
	out := map[*syntax.Decl]*eval.Decl{}
	if r.Eval == nil {
		return out
	}
	for _, f := range r.Eval.Files {
		for _, od := range f.Decls {
			if od.Emit == nil || od.Instance == nil || !od.Instance.User || od.Instance.Name == nil {
				continue
			}
			if _, ok := mustUseText(od.Emit.Doc); ok {
				out[od.D] = od
			}
		}
	}
	return out
}

// memberRe finds `name.const` inside a ${...} piece of a string or command.
var memberRe = regexp.MustCompile(`([A-Za-z_][\w]*)\.([A-Za-z_][\w]*)`)

// lintMustUse reports the @mustUse consts and declarations of Blueprint
// instances that nothing reads: the instance is written, but has no
// effect. Reads are taken from the source as written, `ash.noise_entries`
// or ${ash.noise_entries}, so they hold however the files were expanded.
func lintMustUse(l *lintRun) {
	r := l.r
	if r.Eval == nil {
		return
	}
	reads := map[string]bool{}
	for _, f := range r.Source {
		syntax.Walk(f, func(n syntax.Node) bool {
			switch n := n.(type) {
			case *syntax.Member:
				if id, ok := n.X.(*syntax.Ident); ok && n.Name != nil {
					reads[id.Name+"."+n.Name.Text] = true
				}
			case *syntax.String, *syntax.CommandStmt:
				sp := n.Span()
				if sp.End <= len(f.Src) {
					for _, piece := range evalPieceRe.FindAllString(string(f.Src[sp.Start:sp.End]), -1) {
						for _, m := range memberRe.FindAllStringSubmatch(piece, -1) {
							reads[m[1]+"."+m[2]] = true
						}
					}
				}
			}
			return true
		})
	}
	report := func(in *eval.Instance, msg string) {
		l.out[in.Path] = append(l.out[in.Path], syntax.Diagnostic{Code: "lint/" + l.rule.ID, Severity: l.sev, Sp: in.Name.Sp, Message: msg})
	}
	for _, in := range r.Eval.Instances {
		if !in.User || in.Name == nil || in.ConstDefs == nil || r.Unit(in.Path) == nil {
			continue
		}
		bp := in.Blueprint.D.NameText()
		for _, k := range in.Consts.Keys {
			def := in.ConstDefs[k]
			text, ok := mustUseText(def.Doc)
			if !ok || reads[in.Name.Text+"."+k] {
				continue
			}
			report(in, fmt.Sprintf("%s '%s': nothing reads %s.%s%s", bp, in.Name.Text, in.Name.Text, k, because(text)))
		}
	}
	used := l.usedIDs()
	for d, od := range mustUseDecls(r) {
		pd := r.Decls[d]
		if pd == nil || r.Unit(od.Instance.Path) == nil {
			continue
		}
		id := pd.ID()
		if id == "" || used[pd.Kind.Folder()+"|"+id] || used[id] {
			continue
		}
		text, _ := mustUseText(od.Emit.Doc)
		report(od.Instance, fmt.Sprintf("%s '%s': nothing references its %s '%s'%s", od.Instance.Blueprint.D.NameText(), od.Instance.Name.Text, pd.Kind.Name, id, because(text)))
	}
}

func because(text string) string {
	if text == "" {
		return ""
	}
	return ", so " + text
}
