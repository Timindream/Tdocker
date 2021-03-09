package cgroups

import (
	log "github.com/sirupsen/logrus"
	"tdocker/cgroups/subsystems"
)

type CgroupManager struct {

	//cgroup在hierarchy中的路径
	Path     string

	//资源配置
	Resource *subsystems.ResourceConfig
}

//创建Cgroup
func NewCgroupManager(path string) *CgroupManager {
	return &CgroupManager{
		Path: path,
	}
}

//将进程的PID写入每个cgroup中
func (c *CgroupManager) Apply(pid int) error {
	for _, subSysIns := range subsystems.SubsystemsIns {
		subSysIns.Apply(c.Path, pid)
	}
	return nil
}

//设置各个subsystem挂在中的cgroup资源限制
func (c *CgroupManager) Set(res *subsystems.ResourceConfig) error {
	for _, subSysIns := range subsystems.SubsystemsIns {
		subSysIns.Set(c.Path, res)
	}
	return nil
}

//释放各个subsystem挂载中的cgroup
func (c *CgroupManager) Destroy() error {
	for _, subSysIns := range subsystems.SubsystemsIns {
		if err := subSysIns.Remove(c.Path); err != nil {
			log.Warnf("删除cgroup失败%v", err)
		}
	}
	return nil
}
