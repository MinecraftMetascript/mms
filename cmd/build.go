package cmd

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"mms2/lang"
	"os"

	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:       "build {input} {output}",
	Short:     "Builds your MMS Project",
	Long:      ``,
	ValidArgs: []cobra.Completion{"Input", "Output"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Println("Please provide an input file or directory")
			return
		}
		if len(args) < 2 {
			log.Println("Please provide an output directory")
			return
		}
		inFile := args[0]
		outFile := args[1]

		debugMode, err := cmd.Flags().GetBool("debug")
		if err != nil {
			log.Fatal(err)
		}

		if debugMode {
			log.Println(
				fmt.Sprintf("Building your project from %s to %s", inFile, outFile),
			)

		}

		project := lang.NewProject()

		stat, err := fs.Stat(os.DirFS("."), inFile)

		if err != nil {
			log.Println("Error building project.", err)
			return
		}
		if stat.IsDir() {
			log.Println("Project is a directory")
		} else {
			if content, err := os.ReadFile(inFile); err != nil {
				log.Println("Error building project:")
			} else {
				f := project.AddFile(inFile, string(content))
				f.Parse()
			}
		}
		r, err := json.MarshalIndent(project.BuildFsLike(outFile), "", "  ")

		if debugMode {
			log.Println(
				string(r),
				err,
			)
		}

		if len(project.Diagnostics) > 0 {
			for _, diag := range project.Diagnostics {
				log.Println(diag)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//buildCmd.Flags().StringP("outdir", "o", "", "Output Directory")
	buildCmd.ArgAliases = append(buildCmd.ArgAliases, "input file or directory")
	buildCmd.ArgAliases = append(buildCmd.ArgAliases, "output file")

	buildCmd.Flags().BoolP("debug", "d", false, "Debug mode")
}
