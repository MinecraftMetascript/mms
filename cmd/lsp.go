package cmd

import (
	"log"

	"github.com/minecraftmetascript/mms/lsp"

	"github.com/spf13/cobra"
)

// lspCmd represents the lsp command
var lspCmd = &cobra.Command{
	Use: "lsp",

	Run: func(cmd *cobra.Command, args []string) {

		if err := lsp.Start(); err != nil {
			log.Fatalf("Failed to start LSP server: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(lspCmd)
}
