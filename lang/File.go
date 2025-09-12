package lang

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
)

type File struct {
	Content     string
	Path        string
	Project     *Project
	Diagnostics []traversal.Diagnostic
	Script      *grammar.ScriptContext
}

func (f *File) Parse() error {
	err := f.Project.GlobalScope.PurgeFile(f.Path)
	f.Diagnostics = make([]traversal.Diagnostic, 0)
	if err != nil {
		return err
	}
	parser := traversal.NewParser(f.Content, f.Path, f.Project.GlobalScope, &f.Diagnostics)
	script, err := parser.Parse()
	if err != nil {
		return err
	}
	f.Script = script

	return nil
}

func rulesAt(parent antlr.ParserRuleContext, position int) ([]antlr.ParserRuleContext, error) {
	out := make([]antlr.ParserRuleContext, 0)
	out = append(out, parent)
	if parent.GetStart().GetStart() > position || position > parent.GetStop().GetStop() {
		return nil, nil
	}

	for _, child := range parent.GetChildren() {
		if rule, ok := child.(antlr.ParserRuleContext); ok {
			if rule.GetStart().GetStart() < position && position < rule.GetStop().GetStop() {
				// We are within this rule
				result, err := rulesAt(rule, position)
				if err != nil {
					return nil, err
				}
				if result != nil {
					out = append(out, result...)
				}
			}
		}
	}
	return out, nil

}

func (f *File) GetRulesAtPosition(position int) []antlr.ParserRuleContext {
	rules, _ := rulesAt(f.Script, position)
	return rules

}
