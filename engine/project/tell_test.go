package project

import (
	"strings"
	"testing"
)

// tell, title and actionbar write their message as a text component.
func TestTextCommands(t *testing.T) {
	src := `Objective kills: dummy
Global wave = 1
const welcome: Text = "Welcome, ${@s}!"
Template wave_start(n: any): Text = "{red}Wave ${n}{/} begins"
Function f {
  tell(@a, "Hello")
  tell(@s, welcome)
  tell(@a[tag=x, limit=1], wave_start(kills[@s]))
  tell(@a, "You have ${kills[@s]} kills, 5 * 3")
  title(@s, "{gold}The Library", subtitle = "{gray}*mind the shelves*")
  title(@s, "Hi")
  actionbar(@s, "**Full**")
  let n = 3
  let name = "Bob"
  tell(@s, "n is ${n}, twice ${n * 2}, name ${name}")
}
`
	_, r := bookProject(t, src, nil)
	if d := diagText(r); d != "" {
		t.Fatal(d)
	}
	got := emitted(t, r, "data/demo/function/f.mcfunction")
	for _, want := range []string{
		`tellraw @a "Hello"`,
		`tellraw @s ["","Welcome, ",{selector:"@s"},"!"]`,
		`tellraw @a[tag=x,limit=1] `,
		`tellraw @a[tag=x,limit=1] ["",{text:"Wave ",color:"red"},{score:{name:"@s",objective:"kills"},color:"red"}`,
		`{score:{name:"@s",objective:"kills"}}`,
		`5 * 3`,
		`title @s subtitle {`,
		`title @s title {`,
		`title @s title "Hi"`,
		`title @s actionbar {`,
		`bold:true`,
		`{score:{name:"#f.n",objective:"mms.demo"}}`,
		`{score:{name:"#f.-t0",objective:"mms.demo"}}`,
		`{nbt:"l.\"f\".name",storage:"demo:mms"}`,
	} {
		if !strings.Contains(got, want) {
			t.Errorf("output lacks %s:\n%s", want, got)
		}
	}
	if strings.Index(got, "subtitle") > strings.Index(got, "title @s title {") {
		t.Errorf("subtitle after title:\n%s", got)
	}
}

func TestTextCommandErrors(t *testing.T) {
	for _, tc := range []struct{ body, want string }{
		{`tell("Steve", "Hi")`, "expected a selector"},
		{`tell(@a)`, "needs a message"},
		{`tell(@a, 3)`, "expected text"},
		{`tell(@a, "a\n---\nb")`, "page"},
		{`actionbar(@a, "Hi", subtitle = "x")`, "no parameter 'subtitle'"},
	} {
		_, r := bookProject(t, "Function f {\n  "+tc.body+"\n}\n", nil)
		if d := diagText(r); !strings.Contains(d, tc.want) {
			t.Errorf("%s: diagnostics %q lack %q", tc.body, d, tc.want)
		}
	}
}

func TestTellAccidentalMarkup(t *testing.T) {
	_, r := bookProject(t, "Function f {\n  tell(@a, \"5 * 3\")\n}\n", nil)
	for _, d := range r.Diagnostics("main.mms") {
		if d.Code == "lint/accidental-markup" {
			return
		}
	}
	t.Errorf("no accidental-markup hint: %+v", r.Diagnostics("main.mms"))
}
