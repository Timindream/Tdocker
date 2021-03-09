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

func (s *MemorySubSystem) Name() string {
	return "memory"
}

//创建资源
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

//注册将进程ID放入
func (s *MemorySubSystem) Apply(cgroupPath string, pid int) error {
	if subsysCgrouPath, err := GetCgroupPath(s.Name(), cgroupPath, false); err == nil {
		if err := ioutil.WriteFile(path.Join(subsysCgrouPath, "tasks"), []byte(strconv.Itoa(pid)), 0644); err != nil {
			return fmt.Errorf("创建失败%v", err)
		}
		return nil
	} else {
		return fmt.Errorf("获取cgroup错误%v", err)
	}
}

//释放资源
func (s *MemorySubSystem) Remove(cgroupPath string) error {
	if subsysCgrouPath, err := GetCgroupPath(s.Name(), cgroupPath, false); err == nil {
		return os.Remove(subsysCgrouPath)
	} else {
		return err
	}
}