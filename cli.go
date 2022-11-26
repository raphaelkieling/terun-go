package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func startCLI() {
	app := &cli.App{
		Name:  "terun",
		Usage: "Tool to transport template files",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "base",
				Usage: "It's to set the terun.yml folder base. Default: current folder",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "make",
				Aliases: []string{"m"},
				Usage:   "Transform files based on command",
				Action: func(cCtx *cli.Context) error {
					var basePath string
					basePath = cCtx.String("base")

					// Current folder
					if basePath == "" {
						currentFolderPath, err := os.Getwd()
						basePath = currentFolderPath
						if err != nil {
							return err
						}
					}

					// Make the transportation files
					terun := createTerun(basePath)
					err := terun.Make(cCtx.Args().First())
					if err != nil {
						return err
					}

					return nil
				},
			},
			{
				Name:    "init",
				Aliases: []string{"i"},
				Usage:   "Create the base file for commands",
				Action: func(cCtx *cli.Context) error {
					var basePath string
					basePath = cCtx.String("base")

					// Current folder
					if basePath == "" {
						currentFolderPath, err := os.Getwd()
						basePath = currentFolderPath
						if err != nil {
							return err
						}
					}

					// Make the transportation files
					terun := createTerun(basePath)
					err := terun.Init()
					if err != nil {
						return err
					}

					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
