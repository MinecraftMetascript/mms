package cmd

import (
	//"github.com/minecraftmetascript/mms/lsp"
	"github.com/spf13/cobra"
)

// lspCmd represents the lsp command
var lspCmd = &cobra.Command{
	Use: "lsp",

	Run: func(cmd *cobra.Command, args []string) {
		//lsp.Start()

	},
}

func init() {
	rootCmd.AddCommand(lspCmd)
}
