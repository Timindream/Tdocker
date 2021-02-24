package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"tdocker/container"
)

var runCommand = cli.Command {
	Name: "run",
	Usage: `新建一个容器
			tdocker run -ti [command]`,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name: "ti",
			Usage: "enable tty",
		},
	},
	Action: func(context *cli.Context) error{
		if len(context.Args()) < 1 {
			return fmt.Errorf("未输入参数，请重新输入")
		}
		cmd := context.Args().Get(0)
		tty := context.Bool("ti")
		Run(tty,cmd)
		return nil
	},
}

var initCommand = cli.Command{
	Name:  "init",
	Usage: `初始化进程`,

	Action: func(context *cli.Context) error {
		log.Infof("init come on")
		cmd := context.Args().Get(0)
		log.Infof("command %s",cmd)
		err := container.RunContainerInitProcess(cmd,nil)
		return err
	},
}