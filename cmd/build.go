package cmd

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path"
	"strings"

	"github.com/minecraftmetascript/mms/lib"
	_project "github.com/minecraftmetascript/mms/project"
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:       "build {input} {output}",
	Short:     "Builds your MMS Project",
	Long:      ``,
	ValidArgs: []cobra.Completion{"Input", "Output"},
	Run: func(cmd *cobra.Command, args []string) {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
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
		fsLike := project.BuildFsLike(outFile)
		if err != nil {
			log.Println("Error exporting project:", err)
			return
		}

		_, err = fs.Stat(os.DirFS("."), outFile)
		if err != nil {
			if strings.HasSuffix(err.Error(), "no such file or directory") {
				err = os.MkdirAll(outFile, 0755)
				if err != nil {
					log.Println("Error creating output directory:", err)
					return
				}
			} else {
				log.Println("Error building project", err)
			}
			return
		}

		flushProject(fsLike, outFile)

		if debugMode {
			r, err := json.MarshalIndent(project.Symbols(), "", "  ")
			log.Println(
				string(r),
				err,
			)
		} else {
			// induce serialization for logging
			_, e := json.MarshalIndent(project.Symbols(), "", "  ")
			if e != nil {
				log.Println("Error serializing project:", e)
			}
		}
	},
}

func flushProject(root *lib.FileTreeLike, rootPath string) {
	log.Println(root.Name)
	for _, file := range root.Children {
		targetPath := path.Join(rootPath, file.Name)
		if file.IsDir {
			err := mkdirIfNotExists(targetPath)
			if err != nil {
				log.Println("Error creating directory:", err)
			}
			flushProject(file, targetPath)
		} else {
			err := os.WriteFile(targetPath, []byte(file.Content), 0644)
			if err != nil {
				log.Println("Error writing file:", err)
			}

		}

	}
}

func mkdirIfNotExists(dirPath string) error {
	_, err := fs.Stat(os.DirFS("."), dirPath)
	if err != nil {
		if strings.HasSuffix(err.Error(), "no such file or directory") {
			err = os.MkdirAll(dirPath, 0755)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
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
