package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app:=cli.NewApp()
	app.Name = "tdocker"
	app.Usage = usage
	app.Commands = []cli.Commands{
		initCommand,
		runCommand,
	}
	app.Before = func(context *cli.Context) error {
		log.SetFormatter(&log.JSONFormatter{})

		log.SetFormatter(os.Stdout)
		return nil
	}

	if err:=app.Run(os.Args);err!=nil {
		log.Fatal(err)
	}
}

