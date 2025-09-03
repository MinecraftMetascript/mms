package test

import (
	"fmt"
	"testing"

	"github.com/minecraftmetascript/mms/lang"
	"github.com/minecraftmetascript/mms/lang/constructs/worldgen/surface_rules"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
)

func TestBlockRule(t *testing.T) {
	p := lang.NewProject()
	file := p.AddFile("test.mms", `
namespace test;

rule := SurfaceRule.Block minecraft:stone
`)

	file.Parse()

	ruleSymbol, ok := p.GlobalScope.Get(*traversal.Ref("test", "rule"))
	if ruleSymbol == nil || !ok {
		t.Error("failed to extract rule")
	}

	if rule, ok := ruleSymbol.GetValue().(*surface_rules.Block); ok {
		if rule.BlockName != "minecraft:stone" {
			t.Error("block name is incorrect")
		}

		rootDir := lib.NewDirLike("project", nil)
		err := rule.ExportSymbol(ruleSymbol, rootDir)
		if err != nil {
			t.Error(err)
		}

		rootDir.PrintDebug()

		outFile := rootDir.Get("test/_debug/surface_rules/rule.json")
		if outFile == nil {
			t.Error("failed to export rule")
			return
		}

		fmt.Println(
			outFile.Content,
		)

	} else {
		t.Error("rule is not a block")
	}

}

func TestBandlandsRule(t *testing.T) {
	p := lang.NewProject()
	file := p.AddFile("test.mms", "namespace test; rule := SurfaceRule.Bandlands")
	file.Parse()
	rule, ok := p.GlobalScope.Get(*traversal.Ref("test", "rule"))
	if rule == nil || !ok {
		t.Error("failed to extract rule")
	}

	if _, ok := rule.GetValue().(*surface_rules.Bandlands); !ok {
		t.Error("rule is not bandlands")
	}
}

func TestConditionalRule(t *testing.T) {
	p := lang.NewProject()
	file := p.AddFile("test.mms", "namespace test; rule := SurfaceRule.If ( SurfaceCondition.AboveSurface ) SurfaceRule.Block minecraft:stone")
	file.Parse()
	rule, ok := p.GlobalScope.Get(*traversal.Ref("test", "rule"))
	if rule == nil || !ok {
		t.Errorf("failed to extract rule. Diagnostics: %+v", file.Diagnostics)

	} else {
		if _, ok := rule.GetValue().(*surface_rules.IfRule); !ok {
			t.Error("rule is not condition")
		}
	}

}
