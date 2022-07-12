package subsystems

import (
	"fmt"
	"strings"
	"os"
	"path"
	"bufio"
)


// 获取 Cgroup 挂载点
func FindCgroupMountpoint(subsystem string) string {
	// 打开挂载文件
	f, err := os.Open("/proc/self/mountinfo")
	if err != nil {
		return ""
	}
	defer f.Close()

	// 扫描文件内容
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		//
		txt := scanner.Text()
		// 单词
		fields := strings.Split(txt, " ")
		// todo len-1 ?
		for _, opt := range strings.Split(fields[len(fields)-1], ",") {
			if opt == subsystem {
				return fields[4]
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return ""
	}

	return ""
}

// 获取Cgroup Path
func GetCgroupPath(subsystem string, cgroupPath string, autoCreate bool) (string, error) {
	cgroupRoot := FindCgroupMountpoint(subsystem)
	if _, err := os.Stat(path.Join(cgroupRoot, cgroupPath)); err == nil || (autoCreate && os.IsNotExist(err)) {
		if os.IsNotExist(err) {
			if err := os.Mkdir(path.Join(cgroupRoot, cgroupPath), 0755); err == nil {
			} else {
				return "", fmt.Errorf("error create cgroup %v", err)
			}
		}
		return path.Join(cgroupRoot, cgroupPath), nil
	} else {
		return "", fmt.Errorf("cgroup path error %v", err)
	}
}