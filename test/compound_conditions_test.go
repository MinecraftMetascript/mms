package test

import (
	"encoding/json"
	"fmt"
	"github.com/minecraftmetascript/mms/lang"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"testing"
)

func TestCompoundAnd(t *testing.T) {
	p := lang.NewProject()
	file := p.AddFile("test.mms", `
namespace ns;

rule := SurfaceRule { If ( and( 
		Biome [ forest ] 
		Biome [ plains ] 
	)) 
	Block stone 
}
`)

	file.Parse()

	c, ok := p.GlobalScope.Get(*traversal.Ref("ns", "rule"))
	if !ok {
		t.Fatal("Failed to get cond")
	}

	txt, err := json.MarshalIndent(c.GetValue(), "", "  ")
	if err != nil {
		t.Fatal(err)
	}

	if len(p.Diagnostics) > 0 {
		for _, d := range p.Diagnostics {
			t.Error(d)
		}
	}

	fmt.Println(string(txt))

}
