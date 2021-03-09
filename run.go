package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
	"tdocker/container"
)

func Run(tty bool, comArray []string) {

	parent, writePipe := container.NewParentProcess(tty)
	if parent == nil {
		log.Errorf("进程错误")
		return
	}
	if err := parent.Start(); err != nil {
		log.Error()
	}

	sendInitCommand(comArray, writePipe)
	parent.Wait()
	os.Exit(0)
}

func sendInitCommand(comArray []string, writePipe *os.File) {
	command := strings.Join(comArray, " ")
	log.Infof("command all is %s", command)
	writePipe.WriteString(command)
	writePipe.Close()
}
