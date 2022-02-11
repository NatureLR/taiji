package versions

import (
	"fmt"
	"runtime"
)

// 通过go build -X 注入的参数
var (
	xVersion   string // Version 程序版本
	xGitCommit string // GitCommit git提交的hash
	xBuilt     string // Built 编译时间
)

// Info 版本信息
type Info struct {
	Version   string `json:"version"`
	GitCommit string `json:"gitCommit"`
	GoVersion string `json:"goVersion"`
	Built     string `json:"built"`
}

// New 返回*info对象
func New(version, gitcommit, built string) *Info {
	return &Info{
		Version:   version,
		GitCommit: gitcommit,
		GoVersion: runtime.Version(),
		Built:     built,
	}
}

func (i *Info) strings() string {
	return fmt.Sprintf("%s GitCommit:%s GoVersion:%s Built:%s", i.Version, i.GitCommit, i.GoVersion, i.Built)
}

func Strings() string {
	return Default.strings()
}

func Print() {
	fmt.Println(Default.strings())
}

func init() {
	Default = New(xVersion, xGitCommit, xBuilt)
}

var Default *Info
