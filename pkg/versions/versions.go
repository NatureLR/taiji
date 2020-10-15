package versions

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
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

// JSON 输出json
func (i *Info) JSON() string {
	ret, err := json.Marshal(i)
	if err != nil {
		log.Print("解析json错误")
		os.Exit(1)
	}
	return string(ret)
}

// Column 以一列的形式
func (i *Info) Column() string {
	return fmt.Sprintf("%s GitCommit:%s GoVersion:%s Built:%s", i.Version, i.GitCommit, i.GoVersion, i.Built)
}

// Strings 输出字符串
func (i *Info) Strings(format string) string {
	var s string
	switch format {
	case "json":
		s = i.JSON()
	case "column":
		s = i.Column()
	default:
		s = i.Column()
	}
	return s
}

func init() {
	Default = New(xVersion, xGitCommit, xBuilt)
}

var (
	// Default 输出的格式
	Default *Info
	// Format 输出的格式
	Format string
)

// Print 将版本输出到控制台
func Print() {
	fmt.Println(Default.Strings(Format))
}

// Strings 输出字符串
func Strings() string {
	return Default.Strings(Format)
}
