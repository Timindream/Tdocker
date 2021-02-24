package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
	"tdocker/cgroups"
	"tdocker/container"
)

func Run(tty bool, command string) {
	parent := container.NewProcess(tty, command)
	if err := parent.Start(); err != nil {
		log.Error()
	}
	cgroupManager := cgroups.CgroupManager("tdocker-cgroup")
	defer cgroupManager.Destroy()
	cgroupManager.Apply(parent.Process.Pid)
	sendInitCommand(comArray, writePipe)
	parent.Wait()
}

func sendInitCommand(comArray []string, writePipe *os.File) {
	command := strings.Join(comArray, " ")
	log.Infof("command all is %s", command)
	writePipe.WriteString(command)
	writePipe.Close()
}
