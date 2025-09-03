package test

import (
	"testing"

	"github.com/minecraftmetascript/mms/lang"
	"github.com/minecraftmetascript/mms/lang/traversal"
)

func TestSimpleDeclaration(t *testing.T) {
	p := lang.NewProject()
	file := p.AddFile("test.mms", "namespace test; blockRule := SurfaceRule.Block minecraft:stone")

	file.Parse()

	if _, ok := file.Project.GlobalScope.Get(*traversal.Ref("test", "blockRule")); !ok {
		t.Error("failed to extract blockRule")
	}
}
