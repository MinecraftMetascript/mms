//go:build !js && !wasm

package main

import "mms2/cmd"

func main() {
	cmd.Execute()
	//if len(os.Args) < 2 {
	//	panic("Please provide a filename as the first argument")
	//}
	//filename := os.Args[1]
	//p := lang.NewProject()
	//
	//content, err := os.ReadFile(filename)
	//if err != nil {
	//	panic(err)
	//}
	//f := p.AddFile(filename, string(content))
	//f.Parse()
	//
	//for _, diag := range f.Diagnostics {
	//	fmt.Println(diag)
	//}
	//if len(f.Diagnostics) == 0 {
	//	fmt.Println("No diagnostics generated")
	//	root := lib.NewDirLike("export", nil)
	//	for _, x := range p.GlobalScope.Symbols() {
	//		if x.GetValue() != nil {
	//			x.GetValue().(traversal.Construct).ExportSymbol(x, root)
	//		} else {
	//			fmt.Println("She do be nil")
	//		}
	//	}
	//	root.PrintDebug()
	//}

}
