package main

import (
	"log"
	"os"
	"install"
	"github.com/urfave/cli/v2"
)

func main() {
	var ff bool
	var force bool
	var colors bool
	var mirror string

	app := &cli.App{
		Name:  "TumCat",
		Usage: "The universal managing cat - Install Packages, Apps, Manage Projects, Create Projects off Templates, etc.",
		// Setting all of the flags
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "ff",
				Value:       false,
				Usage:       "Fast-Forward all of the annoying questions (Uses default values)",
				Destination: &ff,
				Aliases:     []string{"fastforward", "fafo"},
			},
			&cli.BoolFlag{
				Name:        "force",
				Value:       false,
				Usage:       "Force an operation",
				Destination: &force,
				Aliases:     []string{"f"},
			},
			&cli.BoolFlag{
				Name:        "colors",
				Value:       true,
				Usage:       "Do you want to see colors in the output?",
				Destination: &colors,
				Aliases:     []string{"c", "col"},
			},
			&cli.StringFlag{
				Name:        "mirror",
				Value:       "random",
				Usage:       "Select a mirror. Not recommended as TumCat selects a random one and tries all of them.",
				Destination: &mirror,
				Aliases:     []string{"m", "mir"},
			},
		},
		// Setting all of the Commands
		Commands: []*cli.Command{
			{
				Name: "install",
				Aliases: []string{"i", "inst"},
				Usage: "Use: tumcat i <package-name> | sInstalls an application from a random mirror"
				Action: func(c *cli.Context) error {
					return nil
				},
			}
		},
		Action: func(c *cli.Context) error {
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func install(pkgname string, mirror string)
{

}