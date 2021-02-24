package subsystems

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

type MemorySubSystem struct {
}

func (s *MemorySubSystem) set(cgroupPath string, res *ResourceConfig) error {
	if subsysCgrouPath, err := GetCgroupPath(s.Name(), cgroupPath, true); err == nil {
		if res.MemoryLimit != "" {
			if err := ioutil.WriteFile(path.Join(subsysCgrouPath, "memory.limit_in_bytes"), []byte(res.MemoryLimit), 0644); err != nil {
				return fmt.Errorf("创建失败%v", err)
			}
		}
		return nil
	} else {
		return err
	}
}

func (s *MemorySubSystem) Remove() error {
	if subsysCgrouPath, err := GetCgroupPath(s.Name(), cgroupPath, false); err == nil {
		return os.Remove(subsysCgrouPath)
	} else {
		return err
	}
}

func (s *MemorySubSystem) Apply(cgrouPath string, pid int) error {
	if subsysCgrouPath, err := GetCgroupPath(s.Name(), cgroupPath, false); err == nil {
		if err := ioutil.WriteFile(path.Join(subsysCgrouPath, "tasks"), []byte(strconv.Itoa(pid)), 0644); err != nil {
			return fmt.Errorf("创建失败%v", err)
		}
		return nil
	} else {
		return fmt.Errorf("获取cgroup错误%v", err)
	}
}

func (s *MemorySubSystem) Name() string {
	return "memory"
}
