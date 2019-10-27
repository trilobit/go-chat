package main

import (
	"github.com/trilobit/go-chat/src"
	"github.com/urfave/cli"
	"os"
)

var (
	Version = "development"
)

func main() {
	app := cli.NewApp()
	app.Name = "GoChat"
	app.Version = Version
	app.Commands = []cli.Command{
		{
			Name:      "serve",
			ShortName: "s",
			Usage:     "start web server",
			Action: func(ctx *cli.Context) {
				src.Run()
			},
		},
		{
			Name:  "migrate:up",
			Usage: "up all new migrations",
			//Flags: []cli.Flag{
			//	cli.StringFlag{
			//		Name:  "dir",
			//		Value: "./migrations",
			//	},
			//},
			Action: func(ctx *cli.Context) {
				//dir := ctx.String("dir")
				src.Migrate("up")
			},
		},
		{
			Name:  "migrate:reset",
			Usage: "rollback all applied migrations",
			Action: func(ctx *cli.Context) {
				src.Migrate("reset")
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
