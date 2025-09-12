package cmd

import (
	"fmt"

	"github.com/minecraftmetascript/mms/lsp"
	"github.com/spf13/cobra"
)

// lspCmd represents the lsp command
var lspCmd = &cobra.Command{
	Use: "lsp",

	Run: func(cmd *cobra.Command, args []string) {
		if lsp.Start() != nil {
			panic("Failed to start LSP server")
		} else {
			fmt.Println("LSP Started")
		}
	},
}

func init() {
	rootCmd.AddCommand(lspCmd)
}
