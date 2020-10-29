package main

import (
	"github.com/micro-community/micro-starter/config"
	"github.com/micro-community/micro-starter/pubsub"
	"github.com/micro/micro/v3/cmd"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"github.com/urfave/cli/v2"

	//load profile
	"github.com/micro-community/micro-starter/profile"
)

func main() {

	srv := service.New(
		service.Name("micro-v3-starter"),
		service.Version("latest"),
	)

	// add customer Flags
	cmdFlags := []cli.Flag{
		&cli.BoolFlag{
			Name:    "debug",
			Usage:   "run in debug mode",
			EnvVars: []string{"MICRO_STARTER_DEBUG_MODE"},
			Value:   true,
		},
	}

	cmdFlags = append(cmd.DefaultCmd.App().Flags, cmdFlags...)
	cmdOption := cmd.Flags(cmdFlags...)

	cmd.DefaultCmd.Init(cmdOption)

	srv.Init()

	config.LoadConfigWithDefault(nil)

	profile.BuildingStartupService(srv, config.Default)
	//handle pub/sub message
	pubsub.RegisterSubscription(srv, config.Default.Pubsub)

	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
