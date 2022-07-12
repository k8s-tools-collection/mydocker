package subsystems

// 资源配置
type ResourceConfig struct {
	MemoryLimit string // 内存限制
	CpuShare    string // cpu
	CpuSet      string // cpu
}

type Subsystem interface {
	Name() string
	Set(path string, res *ResourceConfig) error
	Apply(path string, pid int) error
	Remove(path string) error
}

var (
	// 子系统
	SubsystemsIns = []Subsystem{
		&CpusetSubSystem{},
		&MemorySubSystem{},
		&CpuSubSystem{},
	}
)
