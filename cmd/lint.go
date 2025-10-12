package cmd

import (
	"io/fs"
	"log"
	"os"
	"strings"

	_project "github.com/minecraftmetascript/mms/project"
	"github.com/spf13/cobra"
)

// lintCmd represents the lint command
var lintCmd = &cobra.Command{
	Use:       "lint {input} {output}",
	Short:     "Lints your MMS Project",
	Long:      ``,
	ValidArgs: []cobra.Completion{"Input"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Println("Please provide an input file or directory")
			return
		}
		inFile := args[0]
		if strings.HasSuffix(inFile, "/") {
			inFile = strings.TrimSuffix(inFile, "/")
		}

		project := _project.NewProject()
		stat, err := fs.Stat(os.DirFS("."), inFile)

		if err != nil {
			log.Println("Error building project", err)
			return
		}

		if stat.IsDir() {
			log.Println("Project is a directory")
		} else {
			if content, err := os.ReadFile(inFile); err != nil {
				log.Println("Error reading project:", err)
			} else {
				_, err := project.AddFile(inFile, string(content))

				if err != nil {
					log.Println("Error parsing project:", err)
				}
			}
		}
		for _, diag := range project.AllDiagnostics() {
			log.Println(diag)
		}

		project.BuildFsLike(".")

		//
		//stat, err := fs.Stat(os.DirFS("."), inFile)
		//
		//if err != nil {
		//	log.Println("Error building project.", err)
		//	return
		//}
		//if stat.IsDir() {
		//	log.Println("Project is a directory")
		//} else {
		//	if content, err := os.ReadFile(inFile); err != nil {
		//		log.Println("Error building project:")
		//	} else {
		//		f := project.AddFile(inFile, string(content))
		//		f.Parse()
		//	}
		//}
		//
		//if len(project.Diagnostics()) > 0 {
		//	for _, diag := range project.Diagnostics() {
		//		log.Println(diag)
		//	}
		//} else {
		//	log.Println("Project has no errors or warnings")
		//}

	},
}

func init() {
	rootCmd.AddCommand(lintCmd)
}
