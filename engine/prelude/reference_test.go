package prelude

import "testing"

// The reference lists what projects use: Blueprints with their documented
// fields, Templates with their signatures, Types, and not @internal ones.
func TestReference(t *testing.T) {
	by := map[string]RefDecl{}
	for _, r := range Reference() {
		by[r.Name] = r
	}
	ci, ok := by["CustomItem"]
	if !ok || ci.Kind != "Blueprint" || ci.As == "" || len(ci.Fields) < 5 || ci.Doc == "" {
		t.Fatalf("CustomItem = %+v", ci)
	}
	for _, f := range ci.Fields {
		if f.Name == "base" && (f.Type != "Id<item>" || f.Optional || f.Doc == "") {
			t.Errorf("base = %+v", f)
		}
		if f.Name == "components" && (!f.Optional || f.Default != "{}") {
			t.Errorf("components = %+v", f)
		}
	}
	if tell := by["tell"]; tell.Signature != "tell(target: Selector, message: Text): Commands" || !tell.Builtin {
		t.Errorf("tell = %+v", tell)
	}
	if _, ok := by["ItemHook"]; ok {
		t.Error("an @internal Blueprint is listed")
	}
	if _, ok := by["Function"]; ok {
		t.Error("a compiler kind is listed")
	}
	if r := by["Range"]; len(r.Uses) < 3 || r.Uses[0].Type != "IntNumberProviderRef" || r.Uses[0].Doc == "" {
		t.Errorf("Range uses = %+v", r.Uses)
	}
	if lt := by["LangText"]; lt.Signature != "LangText = string | {string: string}" {
		t.Errorf("LangText = %+v", lt)
	}
}
