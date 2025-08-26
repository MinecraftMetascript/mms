package lang

import (
	"github.com/minecraftmetascript/mms/lang/grammars"
	"github.com/minecraftmetascript/mms/lang/mms_errors"
	"github.com/minecraftmetascript/mms/lib"
)

type NamespaceAware interface {
	Namespace() string
	SetNamespace(string)
}

type MMSFileVisitor struct {
	grammars.BaseMMSParserListener

	file     *MMSFile
	siblings []NamespaceAware
}

func NewMMSFileVisitor(file *MMSFile, siblings []NamespaceAware) *MMSFileVisitor {
	return &MMSFileVisitor{
		file:     file,
		siblings: siblings,
	}
}
func (v *MMSFileVisitor) ExitNamespaceDeclaration(ctx *grammars.NamespaceDeclarationContext) {
	if id := ctx.ResourceReference(); id != nil {
		v.file.namespace = id.GetText()
		for _, s := range v.siblings {
			s.SetNamespace(v.file.namespace)
		}
	} else {
		v.file.errors = append(v.file.errors, mms_errors.SyntaxError(
			v.file.path,
			lib.GetRuleLocation(ctx),
			"Namespace declaration must have an identifier",
		))
	}
}

func (v *MMSFileVisitor) ExitMmsFile(ctx *grammars.MmsFileContext) {
	v.file.tree = ctx
}
