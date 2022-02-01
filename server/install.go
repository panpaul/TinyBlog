package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"server/global"
	"server/model"
	"server/router"
	"server/utils"
	"strconv"
)

func install(configPath string) error {
	global.LOG.Info("Your config file is located at: " + configPath)

	askWeb()
	askRedis()
	askDb()
	global.CONF.Jwt.JwtSecret = utils.RandomString(32)
	global.CONF.Jwt.JwtExpireHour = 12
	global.CONF.Development = false

	fmt.Printf("%+v\n", global.CONF)

	ans := false
	prompt := &survey.Confirm{
		Message: "Are these config correct?",
	}
	_ = survey.AskOne(prompt, &ans)

	if !ans {
		global.LOG.Warn("Bye")
		return nil
	}

	global.CONF.WriteConfig("./config.yaml")

	model.SetupDatabase()
	model.MigrateDatabase()
	router.SetupRouterV1()

	return nil
}

func askWeb() {
	webConfigQs := []*survey.Question{
		{
			Name: "Address",
			Prompt: &survey.Input{
				Message: "WebServer Listen Address:",
				Default: global.CONF.WebServer.Address,
			},
			Validate: survey.Required,
		},
		{
			Name: "Port",
			Prompt: &survey.Input{
				Message: "WebServer Listen Port:",
				Default: strconv.Itoa(global.CONF.WebServer.Port),
			},
			Validate: survey.Required,
		},
	}
	_ = survey.Ask(webConfigQs, &global.CONF.WebServer)
}

func askRedis() {
	redisConfigQs := []*survey.Question{
		{
			Name: "Address",
			Prompt: &survey.Input{
				Message: "Redis Server Address:",
				Default: global.CONF.Redis.Address,
			},
			Validate: survey.Required,
		},
		{
			Name: "Db",
			Prompt: &survey.Input{
				Message: "Redis DB:",
				Default: strconv.Itoa(global.CONF.Redis.Db),
			},
			Validate: survey.Required,
		},
		{
			Name: "Password",
			Prompt: &survey.Input{
				Message: "Redis Password:",
				Default: global.CONF.Redis.Password,
			},
		},
	}
	_ = survey.Ask(redisConfigQs, &global.CONF.Redis)
}

func askDb() {
	dbConfigQs := []*survey.Question{
		{
			Name: "Address",
			Prompt: &survey.Input{
				Message: "Database Server Address:",
				Default: global.CONF.Database.Address,
			},
			Validate: survey.Required,
		},
		{
			Name: "Port",
			Prompt: &survey.Input{
				Message: "Database Server Port:",
				Default: strconv.Itoa(global.CONF.Database.Port),
			},
			Validate: survey.Required,
		},
		{
			Name: "User",
			Prompt: &survey.Input{
				Message: "Database Server Username:",
				Default: global.CONF.Database.User,
			},
			Validate: survey.Required,
		},
		{
			Name: "Password",
			Prompt: &survey.Input{
				Message: "Database Server Password:",
				Default: global.CONF.Database.Password,
			},
			Validate: survey.Required,
		},
		{
			Name: "Database",
			Prompt: &survey.Input{
				Message: "Database Server Database:",
				Default: global.CONF.Database.Database,
			},
			Validate: survey.Required,
		},
		{
			Name: "Prefix",
			Prompt: &survey.Input{
				Message: "Database Server Table Prefix:",
				Default: global.CONF.Database.Prefix,
			},
			Validate: survey.Required,
		},
	}
	_ = survey.Ask(dbConfigQs, &global.CONF.Database)
}
