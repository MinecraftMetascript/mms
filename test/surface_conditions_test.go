package test

import (
	"fmt"
	"testing"

	"github.com/minecraftmetascript/mms/lang"
	"github.com/minecraftmetascript/mms/lang/constructs/worldgen/surface_rules"
	"github.com/minecraftmetascript/mms/lang/traversal"
)

func TestAboveSurfaceCondition(t *testing.T) {
	p := lang.NewProject()
	file := p.AddFile("test.mms", "namespace test; t := SurfaceCondition.AboveSurface")

	file.Parse()

	if _, ok := file.Project.GlobalScope.Get(*traversal.Ref("test", "t")); !ok {
		t.Error("failed to extract t")
	}

	if len(file.Diagnostics) != 0 {
		t.Errorf("expected no diagnostics, got: %+v", file.Diagnostics)
		for _, d := range file.Diagnostics {
			t.Errorf("diagnostic: %+v", d)
		}
	}
}

func TestBiomesSurfaceCondition_SingleFq(t *testing.T) {
	p := lang.NewProject()
	file := p.AddFile("test.mms", "namespace test; t := SurfaceCondition.Biome [ minecraft:plains ]")

	file.Parse()

	val, ok := file.Project.GlobalScope.Get(*traversal.Ref("test", "t"))
	if !ok {
		t.Error("failed to extract t")
	} else {
		cond, ok := val.GetValue().(*surface_rules.BiomeCondition)
		if !ok {
			t.Error("t is not a biome condition")
		} else {
			if !cond.Biomes()[0].Equals(*traversal.Ref("minecraft", "plains")) {
				t.Error("biome is incorrect")
			}
		}
	}

	if len(file.Diagnostics) != 0 {
		t.Errorf("expected no diagnostics, got: %+v", file.Diagnostics)
		for _, d := range file.Diagnostics {
			t.Errorf("diagnostic: %+v", d)
		}
	}
}
func TestBiomesSurfaceCondition_SinglePq(t *testing.T) {
	p := lang.NewProject()
	file := p.AddFile("test.mms", "namespace test; t := SurfaceCondition.Biome [ plains ]")

	file.Parse()

	val, ok := file.Project.GlobalScope.Get(*traversal.Ref("test", "t"))
	if !ok {
		t.Error("failed to extract t")
	} else {
		cond, ok := val.GetValue().(*surface_rules.BiomeCondition)
		if !ok {
			t.Error("t is not a biome condition")
		} else {
			if !cond.Biomes()[0].Equals(*traversal.Ref("test", "plains")) {
				t.Error(fmt.Sprintf("biome is incorrect, expected: %s, got: %s", *traversal.Ref("test", "plains"), cond.Biomes()[0]))
			}
		}
	}

	if len(file.Diagnostics) != 0 {
		t.Errorf("expected no diagnostics, got: %+v", file.Diagnostics)
		for _, d := range file.Diagnostics {
			t.Errorf("diagnostic: %+v", d)
		}
	}
}
func TestBiomesSurfaceCondition_Multi(t *testing.T) {
	p := lang.NewProject()
	file := p.AddFile("test.mms", "namespace test; t := SurfaceCondition.Biome [ minecraft:forest plains ]")

	file.Parse()

	expectedRefs := []traversal.Reference{
		*traversal.Ref("minecraft", "forest"),
		*traversal.Ref("test", "plains"),
	}

	val, ok := file.Project.GlobalScope.Get(*traversal.Ref("test", "t"))
	if !ok {
		t.Error("failed to extract t")
	} else {
		cond, ok := val.GetValue().(*surface_rules.BiomeCondition)
		if !ok {
			t.Error("t is not a biome condition")
		} else {
			for i, ref := range cond.Biomes() {
				if !ref.Equals(expectedRefs[i]) {
					t.Error(fmt.Sprintf("biome is incorrect, expected: %s, got: %s", expectedRefs[i], ref))
				}
			}

		}
	}

	if len(file.Diagnostics) != 0 {
		t.Errorf("expected no diagnostics, got: %+v", file.Diagnostics)
		for _, d := range file.Diagnostics {
			t.Errorf("diagnostic: %+v", d)
		}
	}
}

func TestHoleSurfaceCondition(t *testing.T) {
	p := lang.NewProject()
	file := p.AddFile("test.mms", "namespace test; t := SurfaceCondition.Hole")
	file.Parse()
	if _, ok := file.Project.GlobalScope.Get(*traversal.Ref("test", "t")); !ok {
		t.Error("failed to extract t")
	}
	if len(file.Diagnostics) != 0 {
		t.Errorf("expected no diagnostics, got: %+v", file.Diagnostics)
	}
}

func TestSteepSurfaceCondition(t *testing.T) {
	p := lang.NewProject()
	file := p.AddFile("test.mms", "namespace test; t := SurfaceCondition.Steep")
	file.Parse()
	if _, ok := file.Project.GlobalScope.Get(*traversal.Ref("test", "t")); !ok {
		t.Error("failed to extract t")
	}
	if len(file.Diagnostics) != 0 {
		t.Errorf("expected no diagnostics, got: %+v", file.Diagnostics)
	}
}

func TestFreezingSurfaceCondition(t *testing.T) {
	p := lang.NewProject()
	file := p.AddFile("test.mms", "namespace test; t := SurfaceCondition.Freezing")
	file.Parse()
	if _, ok := file.Project.GlobalScope.Get(*traversal.Ref("test", "t")); !ok {
		t.Error("failed to extract t")
	}
	if len(file.Diagnostics) != 0 {
		t.Errorf("expected no diagnostics, got: %+v", file.Diagnostics)
	}
}

func TestStoneDepthSurfaceCondition_Variants(t *testing.T) {
	cases := []string{
		"namespace test; t := SurfaceCondition.StoneDepth Floor 2 add 3",
		"namespace test; t := SurfaceCondition.StoneDepth Ceiling 5 sub 7",
	}
	for i, content := range cases {
		p := lang.NewProject()
		file := p.AddFile("test.mms", content)
		file.Parse()
		if _, ok := file.Project.GlobalScope.Get(*traversal.Ref("test", "t")); !ok {
			t.Errorf("case %d: failed to extract t", i)
		}
		if len(file.Diagnostics) != 0 {
			t.Errorf("case %d: expected no diagnostics, got: %+v", i, file.Diagnostics)
		}
	}
}

func TestAboveWaterSurfaceCondition_AddSub(t *testing.T) {
	cases := []string{
		"namespace test; t := SurfaceCondition.AboveWater 2 0.4 add",
		"namespace test; t := SurfaceCondition.AboveWater 0 1 sub",
	}
	for i, content := range cases {
		p := lang.NewProject()
		file := p.AddFile("test.mms", content)
		file.Parse()
		if _, ok := file.Project.GlobalScope.Get(*traversal.Ref("test", "t")); !ok {
			t.Errorf("case %d: failed to extract t", i)
		}
		if len(file.Diagnostics) != 0 {
			t.Errorf("case %d: expected no diagnostics, got: %+v", i, file.Diagnostics)
		}
	}
}

func TestNoiseSurfaceCondition_VariedRefs(t *testing.T) {
	cases := []string{
		"namespace test; t := SurfaceCondition.Noise minecraft:erosion [ 0.1 , 0.9 ]",
		"namespace test; t := SurfaceCondition.Noise local_noise [ 1 , 2 ]",
	}
	for i, content := range cases {
		p := lang.NewProject()
		file := p.AddFile("test.mms", content)
		file.Parse()
		if _, ok := file.Project.GlobalScope.Get(*traversal.Ref("test", "t")); !ok {
			t.Errorf("case %d: failed to extract t", i)
		}
		if len(file.Diagnostics) != 0 {
			t.Errorf("case %d: expected no diagnostics, got: %+v", i, file.Diagnostics)
		}
	}
}

func TestVerticalGradientSurfaceCondition_Variants(t *testing.T) {
	cases := []string{
		"namespace test; t := SurfaceCondition.VerticalGradient \"seed1\" absolute 5 , below_top 2",
		"namespace test; t := SurfaceCondition.VerticalGradient \"seed2\" above_bottom 10 , absolute 0",
	}
	for i, content := range cases {
		p := lang.NewProject()
		file := p.AddFile("test.mms", content)
		file.Parse()
		if _, ok := file.Project.GlobalScope.Get(*traversal.Ref("test", "t")); !ok {
			t.Errorf("case %d: failed to extract t", i)
		}
		if len(file.Diagnostics) != 0 {
			t.Errorf("case %d: expected no diagnostics, got: %+v", i, file.Diagnostics)
		}
	}
}

func TestYAboveSurfaceCondition_Variants(t *testing.T) {
	cases := []string{
		"namespace test; t := SurfaceCondition.YAbove above_bottom 10 3 add",
		"namespace test; t := SurfaceCondition.YAbove below_top 5 2 sub",
	}
	for i, content := range cases {
		p := lang.NewProject()
		file := p.AddFile("test.mms", content)
		file.Parse()
		if _, ok := file.Project.GlobalScope.Get(*traversal.Ref("test", "t")); !ok {
			t.Errorf("case %d: failed to extract t", i)
		}
		if len(file.Diagnostics) != 0 {
			t.Errorf("case %d: expected no diagnostics, got: %+v", i, file.Diagnostics)
		}
	}
}

func TestReferenceSurfaceCondition_FQandPQ(t *testing.T) {
	cases := []string{
		"namespace test; t := minecraft:some_condition",
		"namespace test; t := local_condition",
	}
	for i, content := range cases {
		p := lang.NewProject()
		file := p.AddFile("test.mms", content)
		file.Parse()
		if _, ok := file.Project.GlobalScope.Get(*traversal.Ref("test", "t")); !ok {
			t.Errorf("case %d: failed to extract t", i)
		}
		if len(file.Diagnostics) != 0 {
			t.Errorf("case %d: expected no diagnostics, got: %+v", i, file.Diagnostics)
		}
	}
}

func TestCompoundSurfaceCondition_FlatAndMultiline(t *testing.T) {
	cases := []string{
		"namespace test; t := & ( SurfaceCondition.AboveSurface !SurfaceCondition.Hole SurfaceCondition.Biome [ minecraft:plains ] )",
		"namespace test; t := & (\n  SurfaceCondition.Freezing\n  ! SurfaceCondition.Steep\n  SurfaceCondition.YAbove absolute 0 1 add\n)",
	}
	for i, content := range cases {
		p := lang.NewProject()
		file := p.AddFile("test.mms", content)
		file.Parse()
		if _, ok := file.Project.GlobalScope.Get(*traversal.Ref("test", "t")); !ok {
			t.Errorf("case %d: failed to extract t", i)
		}
		if len(file.Diagnostics) != 0 {
			t.Errorf("case %d: expected no diagnostics, got: %+v", i, file.Diagnostics)
			for _, d := range file.Diagnostics {
				t.Logf("diagnostic: %+v", d)
			}
		}
	}
}
