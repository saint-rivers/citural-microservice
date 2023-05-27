package main

import (
	"log"
	"os"

	"github.com/saint-rivers/citural-cli/app/action"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "citural-cli - Basic service developer management CLI"
	app.Usage = "For ease of development and running services"

	app.Commands = []cli.Command{
		{
			Name:        "run",
			HelpName:    "run",
			Action:      action.RunAction,
			ArgsUsage:   ` `,
			Usage:       `runs a specified service in the architecture`,
			Description: ` `,
			Flags: []cli.Flag{
				&cli.UintFlag{
					Name:  "build",
					Usage: "rebuilds containers",
				},
			},
		},
		{
			Name:        "stop",
			HelpName:    "stop",
			Action:      action.StopAction,
			ArgsUsage:   ` `,
			Usage:       `stops a specified service in the architecture`,
			Description: ` `,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
