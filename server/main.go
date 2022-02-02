package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"server/global"
	"server/model"
	"server/service"
)

func main() {
	app := &cli.App{
		Name:    "TinyBlog",
		Version: getVersion(),
		Authors: []*cli.Author{
			{
				Name:  "Paul",
				Email: "panyuxuan@hotmail.com",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Load configuration from `FILE`",
				Value:   "config.yaml",
				EnvVars: []string{"APP_CONFIG"},
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "run",
				Aliases: []string{"r"},
				Usage:   "complete a task on the list",
				Action: func(c *cli.Context) error {
					global.Setup(c.String("config"))
					model.SetupDatabase()
					service.JwtApp.Setup()
					return startWebServer()
				},
			},
			{
				Name:    "install",
				Aliases: []string{"i"},
				Usage:   "Setup database on first run",
				Action: func(c *cli.Context) error {
					global.Setup(c.String("config"))
					return install(c.String("config"))
				},
			},
			{
				Name:    "dev",
				Aliases: []string{"d"},
				Usage:   "development debug",
				Action: func(c *cli.Context) error {
					global.Setup(c.String("config"))
					model.SetupDatabase()
					service.JwtApp.Setup()
					return devDebug()
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

var (
	BuildTimeStamp = "None"
	Version        = "1.0.0"
)

func getVersion() string {
	return Version + "+" + BuildTimeStamp
}
