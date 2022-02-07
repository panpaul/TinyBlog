package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"server/global"
	"server/model"
	"server/service"
	"time"
)

func main() {
	app := &cli.App{
		Name:                 "TinyBlog",
		Compiled:             getBuildTime(),
		Version:              Version,
		EnableBashCompletion: true,
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
					service.ArticleApp.Setup()
					return startWebServer()
				},
			},
			{
				Name:    "setup",
				Aliases: []string{"s"},
				Usage:   "Setup wizard",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "yes",
						Aliases: []string{"y"},
						Usage:   "using default or generated value",
						Value:   false,
					},
				},
				Action: func(c *cli.Context) error {
					global.Setup(c.String("config"))
					return install(c.String("config"), c.Bool("yes"))
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
	BuildTime = "2022-02-01 10:23:00 UTC"
	Version   = "1.0.0"
)

func getBuildTime() time.Time {
	build, err := time.Parse("2006-01-02 03:04:05 MST", BuildTime)
	if err != nil {
		log.Printf("failed to parse build time: %v", err)
		build = time.Now()
	}
	return build
}
