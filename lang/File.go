package lang

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lang/grammars"
	"github.com/minecraftmetascript/mms/lang/mms_errors"
	"github.com/minecraftmetascript/mms/lang/worldgen/noise"
	"github.com/minecraftmetascript/mms/lang/worldgen/surface"
	"github.com/minecraftmetascript/mms/lang/worldgen/surface/surface_conditions"
	"github.com/minecraftmetascript/mms/lang/worldgen/surface/surface_rules"
	"github.com/minecraftmetascript/mms/lib"
)

type MMSFile struct {
	path       string
	namespace  string
	project    *MMSProject
	tree       grammars.IMmsFileContext
	rawContent string

	errors []mms_errors.MMSError

	Symbols *lib.Namespace
}

func (f *MMSFile) GetNamespace() string {
	return f.namespace
}

func (f *MMSFile) AddError(error mms_errors.MMSError) {
	f.errors = append(f.errors, error)
}

func (f *MMSFile) GetErrors() []mms_errors.MMSError {
	return f.errors
}
func (f *MMSFile) GetRawContent() string {
	return f.rawContent
}

func (f *MMSFile) GetTokenAt(position lib.Location) {

}

func (p *MMSProject) ParseFile(
	path string,
	content string,
) *MMSFile {
	f := &MMSFile{
		rawContent: content,
		path:       path,
		project:    p,
		Symbols:    lib.NewNamespace(),
		namespace:  "",
	}

	tokenStream := antlr.NewInputStream(content)

	listener := mms_errors.NewErrorListener(path, func(err mms_errors.MMSError) {
		f.errors = append(f.errors, err)
	})

	lexer := grammars.NewMMSLexer(tokenStream)
	// Register error listener with the lexer
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(listener)

	parser := grammars.NewMMSParser(antlr.NewCommonTokenStream(lexer, 0))

	surfaceVisitor := surface.NewSurfaceVisitor(f.AddError, path)
	noiseVisitor := noise.NewNoiseVisitor(f.AddError, path)

	fileVisitor := NewMMSFileVisitor(f, []NamespaceAware{surfaceVisitor, noiseVisitor})
	parser.AddParseListener(fileVisitor)
	parser.AddParseListener(surfaceVisitor)
	parser.AddParseListener(noiseVisitor)

	// Register error listener with the parser
	parser.RemoveErrorListeners()
	parser.AddErrorListener(listener)

	f.tree = parser.MmsFile()

	surfaceVisitor.DumpDeclarations(f.Symbols)
	noiseVisitor.DumpDeclarations(f.Symbols)
	// POST PARSE LOGIC

	var namespace string
	if f.GetNamespace() == "" {
		namespace = "unknown"
	} else {
		namespace = f.GetNamespace()
	}
	if _, ok := p.symbols[namespace]; !ok {
		p.symbols[namespace] = f.Symbols
	} else {
		p.symbols[namespace].Merge(f.Symbols)
	}

	if len(surfaceVisitor.RuleDeclarations) > 0 {
		for name, rule := range surfaceVisitor.RuleDeclarations {
			replacement, errs := surface.ReplaceRuleReferences(
				rule.Value,
				ProjectSymbols[surface_conditions.SurfaceCondition](p),
				ProjectSymbols[surface_rules.SurfaceRule](p),
				f.namespace,
			)
			if len(errs) > 0 {
				for _, err := range errs {
					f.AddError(
						mms_errors.SyntaxError(
							f.path, rule.Value.GetLocation(), err.Error(),
						),
					)

				}
			} else {
				rule.Value = replacement
				surfaceVisitor.RuleDeclarations[name] = rule
			}
		}
	}

	return f
}
