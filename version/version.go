package version

import (
	"fmt"
)

var (
	// Version 程序版本
	Version string

	// GitCommit git提交的hash
	GitCommit string

	// GoVersion 编译的go版本
	GoVersion string

	// Built 编译时间
	Built string
)

// Print 打印版本
func Print() string {
	return fmt.Sprintf("%s GitCommit:%s GoVersion:%s Built:%s", Version, GitCommit, GoVersion, Built)
}
