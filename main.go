//go:build !js && !wasm

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/minecraftmetascript/mms/lang"
	"github.com/minecraftmetascript/mms/lsp"

	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		Name: "mms",

		Commands: []*cli.Command{
			{
				Name: "version",
				Action: func(c context.Context, cmd *cli.Command) error {
					log.Println("mms version 0.1.0")
					return nil
				},
			},
			{
				Name: "lsp",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name: "stdio",
					},
				},
				Description: "Start the MMS language server",
				Action: func(c context.Context, cmd *cli.Command) error {
					log.Println("Starting MMS language server...")
					return lsp.Start()
				},
			},
			{
				Name: "build",
				Arguments: []cli.Argument{
					&cli.StringArg{
						Name:      "file",
						UsageText: "The root directory or file to build",
					},
				},

				Action: func(c context.Context, cmd *cli.Command) error {
					file := cmd.StringArg("file")
					project := lang.NewMMSProject()
					err := project.ParseFiles(file)
					if err != nil {
						return err
					}

					data, err := json.Marshal(project.Export())
					if err != nil {
						return err
					}
					fmt.Println(string(data))

					return nil
				},
			},
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

}
