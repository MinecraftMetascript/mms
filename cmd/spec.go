package cmd

import (
	"fmt"

	"github.com/minecraftmetascript/mms/lang"
	"github.com/minecraftmetascript/mms/lang/spec"
	"github.com/spf13/cobra"
)

// specCmd represents the spec command
var specCmd = &cobra.Command{
	Use:   "spec",
	Short: "Generate a markdown-formatted language specification",
	Long: `Generate a markdown-formatted language specification.

This will leverage the built-in spec package to generate usage strings for all registered
block types, and the values they may contain, along with any help that has been specified.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(spec.GenerateSpecString(lang.Blocks))
	},
}

func init() {
	rootCmd.AddCommand(specCmd)
}
